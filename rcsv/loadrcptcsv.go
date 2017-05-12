package rcsv

import (
	"fmt"
	"os"
	"rentroll/rlib"
	"strings"
	"time"
)

// // AcctRule is a structure of the 3-tuple that makes up a whole part of an AcctRule
// type CSVAcctRule struct {
// 	Action  string // "d" = debit, "c" = credit
// 	Account string // GL No for the account
// 	Amount  string // use the entire amount of the assessment or deposit, otherwise the amount to use
// 	ASMID   string // Used only for rlib.ReceiptAllocation; the assessment that caused this payment
// }

// CVS record format:
// 0    1           2      3      4            5              6        7                     8                                           9
// BID, RAID,       PMTID, DEPID, Dt,           DocNo,        Amount,  AR,                   AcctRule,                                   Comment
// REH, RA00000001, 2,     1,     "2004-01-01", 1254,         1000.00, "Rent Payment Check", "ASM(7) d ${rlib.DFLT} _, ASM(7) c 11002 _",
// REH, RA00000001, 1,     1,     "2015-11-21", 883789238746, 294.66,  "Rent Payment Check", "ASM(1) c ${GLGENRCV} 266.67, ASM(1) d ${rlib.DFLT} 266.67, ASM(3) c ${GLGENRCV} 13.33, ASM(3) d ${rlib.DFLT} 13.33, ASM(4) c ${GLGENRCV} 5.33, ASM(4) d ${rlib.DFLT} 5.33, ASM(9) c ${GLGENRCV} 9.33,ASM(9) d ${rlib.DFLT} 9.33", "I am a comment"

// GenerateReceiptAllocations processes the AcctRule for the supplied rlib.Receipt and generates rlib.ReceiptAllocation records
func GenerateReceiptAllocations(rcpt *rlib.Receipt, raid int64, xbiz *rlib.XBusiness) error {
	var d1 = time.Date(rcpt.Dt.Year(), rcpt.Dt.Month(), 1, 0, 0, 0, 0, time.UTC)
	var d2 = d1.AddDate(0, 0, 31)
	t := rlib.ParseAcctRule(xbiz, 0, &d1, &d2, rcpt.AcctRule, rcpt.Amount, 1.0)
	u := make(map[int64][]int64)

	// First, group together all entries that refer to a single ASMID into a list of lists
	for i := int64(0); i < int64(len(t)); i++ {
		u[t[i].ASMID] = append(u[t[i].ASMID], i)
	}
	// Process each list in the list of lists.
	for k, v := range u {
		var a rlib.ReceiptAllocation
		a.AcctRule = ""
		a.ASMID = k
		a.Amount = t[int(v[0])].Amount
		a.RCPTID = rcpt.RCPTID
		a.RAID = raid
		a.Dt = rcpt.Dt

		// make sure the referenced assessment actually exists
		a1, _ := rlib.GetAssessment(a.ASMID)
		if a1.ASMID == 0 {
			return fmt.Errorf("GenerateReceiptAllocations: Referenced assessment ID %d does not exist", a.ASMID)
		}

		// for each index in the list, build its part of the AcctRule
		lim := int64(len(v))
		for i := int64(0); i < lim; i++ {
			j := int(v[i])
			a.AcctRule += fmt.Sprintf("ASM(%d) %s %s %s", t[j].ASMID, t[j].Action, t[j].AcctExpr, t[j].Expr)
			if i+1 < lim {
				a.AcctRule += ","
			}
		}
		a.BID = rcpt.BID
		_, err := rlib.InsertReceiptAllocation(&a)
		if err != nil {
			fmt.Printf("GenerateReceiptAllocations: Error inserting ReceiptAllocation: %s\n", err.Error())
			os.Exit(1)
		}
		rcpt.RA = append(rcpt.RA, a)
	}

	return nil
}

//var pmtTypes = map[int64]rlib.PaymentType{}

// CreateReceiptsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateReceiptsFromCSV(sa []string, lineno int) (int, error) {
	funcname := "CreateReceiptsFromCSV"
	var xbiz rlib.XBusiness
	var r rlib.Receipt
	var err error
	bud := strings.ToLower(strings.TrimSpace(sa[0]))

	const (
		BUD      = 0
		RAID     = iota
		PMTID    = iota
		DEPID    = iota
		Dt       = iota
		DocNo    = iota
		Amount   = iota
		AR       = iota
		AcctRule = iota
		Comment  = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RAID", RAID},
		{"PMTID", PMTID},
		{"DEPID", DEPID},
		{"Dt", Dt},
		{"DocNo", DocNo},
		{"Amount", Amount},
		{"AR", AR},
		{"AcctRule", AcctRule},
		{"Comment", Comment},
	}

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(bud) > 0 {
		b1 := rlib.GetBusinessByDesignation(bud)
		if len(b1.Designation) == 0 {
			return CsvErrorSensitivity, fmt.Errorf("CreateLedgerMarkers: rlib.Business with designation %s does not exist", sa[0])
		}
		r.BID = b1.BID
		rlib.GetXBusiness(r.BID, &xbiz)
	}

	pmtTypes := rlib.GetPaymentTypesByBusiness(r.BID)

	//-------------------------------------------------------------------
	// Find Rental Agreement
	//-------------------------------------------------------------------
	raid := CSVLoaderGetRAID(sa[RAID]) // this should probably go away, we should select it from an Assessment in the AcctRule

	_, err = rlib.GetRentalAgreement(raid)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error loading Rental Agreement %s, err = %v", funcname, lineno, sa[RAID], err)
	}

	//-------------------------------------------------------------------
	// Get the rlib.PaymentType
	//-------------------------------------------------------------------
	r.PMTID, _ = rlib.IntFromString(sa[PMTID], "Payment type is invalid")
	_, ok := pmtTypes[r.PMTID]
	if !ok {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  Payment type is invalid: %s", funcname, lineno, sa[PMTID])
	}

	//-------------------------------------------------------------------
	// Get the Depository
	//-------------------------------------------------------------------
	r.DEPID, err = rlib.IntFromString(sa[DEPID], "Depository ID is invalid")
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  Depository ID is invalid: %s", funcname, lineno, sa[DEPID])
	}

	//-------------------------------------------------------------------
	// Get the date
	//-------------------------------------------------------------------
	dt, err := rlib.StringToDate(sa[Dt])
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  invalid rlib.Receipt date:  %s", funcname, lineno, sa[Dt])
	}
	r.Dt = dt

	//-------------------------------------------------------------------
	// Determine the DocNo
	//-------------------------------------------------------------------
	r.DocNo = strings.TrimSpace(sa[DocNo])

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	r.Amount, _ = rlib.FloatFromString(sa[Amount], "rlib.Receipt Amount is invalid")

	//-------------------------------------------------------------------
	// Set the ARID
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[AR])
	if len(s) > 0 {
		rule, err := rlib.GetARByName(r.BID, s)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Could not load AR named %s: %s", funcname, lineno, s, err.Error())
		}
		r.ARID = rule.ARID
	}

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	r.AcctRule = strings.TrimSpace(sa[AcctRule])

	r.Comment = strings.TrimSpace(sa[Comment])

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(r.AcctRule) == 0 || r.PMTID == 0 || r.Amount == 0 || r.BID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("Skipping this record")
	}

	//-------------------------------------------------------------------
	// Make sure there's no duplicate...
	//-------------------------------------------------------------------
	rdup := rlib.GetReceiptDuplicate(&r.Dt, r.Amount, r.DocNo)
	if rdup.RCPTID != 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - this is a duplicate of an existing receipt: %s", funcname, lineno, rdup.IDtoString())
	}

	rcptid, err := rlib.InsertReceipt(&r)
	if err != nil {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error inserting receipt: %v", funcname, lineno, err)
	}
	r.RCPTID = rcptid

	//-------------------------------------------------------------------
	// Create the allocations...
	//-------------------------------------------------------------------
	err = GenerateReceiptAllocations(&r, raid, &xbiz)
	if err != nil {
		rlib.DeleteReceipt(r.RCPTID)
		rlib.DeleteReceiptAllocations(r.RCPTID)
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error processing receipt: %s", funcname, lineno, err.Error())
	}

	// fmt.Printf("Completed generating receipt %d.\n", r.RCPTID)
	//-------------------------------------------------------------------
	// first, make a complete pass through the Assessments to see if any
	// of them have already been marked as paid
	//-------------------------------------------------------------------
	for i := 0; i < len(r.RA); i++ {
		// fmt.Printf("Checking receipt allocation: %#v\n", r.RA[i])
		a, err := rlib.GetAssessment(r.RA[i].ASMID)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error marking assessments as paid: %s", funcname, lineno, err.Error())
		}
		// fmt.Printf("a.FLAGS = 0x%x\n", a.FLAGS)
		if a.FLAGS&1<<0 != 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  assessment %d is already marked as paid", funcname, lineno, a.ASMID)
		}
	}
	//-------------------------------------------------------------------
	// Now mark the allocated assessments as paid
	//-------------------------------------------------------------------
	for i := 0; i < len(r.RA); i++ {
		a, err := rlib.GetAssessment(r.RA[i].ASMID)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error marking assessments as paid: %s", funcname, lineno, err.Error())
		}
		a.FLAGS |= 1 << 0 // bit 0 is the "paid" flag
		err = rlib.UpdateAssessment(&a)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d -  error marking assessments as paid: %s", funcname, lineno, err.Error())
		}
	}

	//-------------------------------------------------------------------
	// Process the receipt...
	//-------------------------------------------------------------------
	rlib.ProcessNewReceipt(Rcsv.Xbiz, &Rcsv.DtStart, &Rcsv.DtStop, &r)

	return 0, nil
}

// LoadReceiptsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadReceiptsCSV(fname string) []error {
	var m []error
	// pmtTypes = rlib.GetPaymentTypes()
	t := rlib.LoadCSV(fname)
	if len(t) > 1 {
		//-------------------------------------------------------------------
		// Check to see if this rental specialty type is already in the database
		//-------------------------------------------------------------------
		des := strings.TrimSpace(t[1][0])
		if len(des) > 0 {
			b := rlib.GetBusinessByDesignation(des)
			if b.BID < 1 {
				err := fmt.Errorf("LoadReceiptsCSV: rlib.Business named %s not found", des)
				m = append(m, err)
				return m
			}
			rlib.InitBusinessFields(b.BID)
			rlib.GetDefaultLedgers(b.BID) // the actually loads the RRdb.BizTypes array which is needed by rpn
		}
	}

	return LoadRentRollCSV(fname, CreateReceiptsFromCSV)
}
