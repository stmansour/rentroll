package rcsv

import (
	"fmt"
	"math"
	"rentroll/rlib"
	"strings"
)

// CVS record format:
// 0    1         2       3        4       5            6
// BUD, Date,    Payor,   DateDue, Amount, DeliveredBy, Assessments
// REX, 6/1/16,  DEP001,  7/1/16  2005.37, 1,           "ASM00001,2"

// CreateInvoicesFromCSV reads an invoice type string array and creates a database record
func CreateInvoicesFromCSV(sa []string, lineno int) {
	funcname := "CreateInvoicesFromCSV"
	var err error
	var inv rlib.Invoice
	var ok bool

	bud := strings.ToLower(strings.TrimSpace(sa[0]))
	if bud == "bud" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 7
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(bud) > 0 {
		b1, _ := rlib.GetBusinessByDesignation(bud)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: line %d - Business with designation %s does net exist\n", funcname, lineno, sa[0])
			return
		}
		inv.BID = b1.BID
	}

	//-------------------------------------------------------------------
	// Date
	//-------------------------------------------------------------------
	inv.Dt, err = rlib.StringToDate(sa[1])
	if err != nil {
		fmt.Printf("%s: line %d - invalid start date:  %s\n", funcname, lineno, sa[1])
		return
	}

	//-------------------------------------------------------------------
	// Payors
	//-------------------------------------------------------------------
	t, err := CSVLoaderTransactantList(sa[2])
	if err != nil {
		fmt.Printf("%s: line %d - invalid payor list:  %s\n", funcname, lineno, sa[2])
		return
	}

	//-------------------------------------------------------------------
	// Date Due
	//-------------------------------------------------------------------
	inv.DtDue, err = rlib.StringToDate(sa[3])
	if err != nil {
		fmt.Printf("%s: line %d - invalid due date:  %s\n", funcname, lineno, sa[3])
		return
	}

	//-------------------------------------------------------------------
	// Amount
	//-------------------------------------------------------------------
	inv.Amount, ok = rlib.FloatFromString(sa[4], "deposit Amount is invalid")
	if !ok {
		return
	}

	//-------------------------------------------------------------------
	// DeliveredBy
	//-------------------------------------------------------------------
	inv.DeliveredBy = strings.TrimSpace(sa[5])

	//-------------------------------------------------------------------
	// Assessments - comma separated list of ASMIDs. Could be of the form
	// ASM00001 or simply 1.
	//-------------------------------------------------------------------
	var asmts []int64
	var mm []rlib.Assessment
	var tot = float64(0)

	s := strings.TrimSpace(sa[6])
	ssa := strings.Split(s, ",")
	if len(ssa) == 0 {
		rlib.Ulog("%s: line %d - no assessments found. You must supply at least one assessment\n", funcname, lineno)
		return
	}
	for i := 0; i < len(ssa); i++ {
		id := CSVLoaderGetASMID(ssa[i])
		if 0 == id {
			rlib.Ulog("%s: line %d - invalid assessment number: %s\n", funcname, lineno, ssa[i])
			return
		}
		asmts = append(asmts, id)
		// load each assessment so that we can total the amount and see if it matches Amount
		a, err := rlib.GetAssessment(id)
		if err != nil {
			fmt.Printf("%s: line %d -  error getting Assessment %d: %v\n", funcname, lineno, id, err)
			return
		}
		tot += a.Amount
		mm = append(mm, a) // may need this later
	}
	if math.Abs(tot-inv.Amount) > 0.005 {
		rlib.Ulog("%s: line %d - Total of all assessments found to be %8.2f, but Amount was specified as %8.2f. Please correct.\n", funcname, lineno, tot, inv.Amount)
		return
	}

	//-------------------------------------------------------------------
	// We have all we need. Write the records.  First, the Invoice itself
	//-------------------------------------------------------------------
	id, err := rlib.InsertInvoice(&inv)
	if err != nil {
		fmt.Printf("%s: line %d -  error inserting invoice: %v\n", funcname, lineno, err)
		return
	}
	// Next, its associated Assessments
	for i := 0; i < len(asmts); i++ {
		var a rlib.InvoiceAssessment
		a.InvoiceNo = id
		a.ASMID = asmts[i]
		err = rlib.InsertInvoiceAssessment(&a)
		if nil != err {
			fmt.Printf("%s: line %d -  error inserting invoice part: %v\n", funcname, lineno, err)
			rlib.DeleteInvoice(id)
			return
		}
	}
	// Finally, the payors
	for i := 0; i < len(t); i++ {
		var a rlib.InvoicePayor
		a.InvoiceNo = id
		a.PID = t[i].TCID
		err = rlib.InsertInvoicePayor(&a)
		if nil != err {
			fmt.Printf("%s: line %d -  error inserting invoice payor: %v\n", funcname, lineno, err)
			rlib.DeleteInvoice(id)
			return
		}
	}
}

// LoadInvoicesCSV loads a csv file with deposits and creates Invoice records
func LoadInvoicesCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateInvoicesFromCSV(t[i], i+1)
	}
}
