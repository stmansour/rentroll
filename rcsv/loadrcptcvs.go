package rcsv

import (
	"fmt"
	"regexp"
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
// 0    1           2      3             4             5        6                                           7
// BID, RAID,       PMTID, Dt,           DocNo,        Amount,  AcctRule,                                   Comment
// REH, RA00000001, 2,     "2004-01-01", 1254,         1000.00, "ASM(7) d ${rlib.DFLT} _, ASM(7) c 11002 _",
// REH, RA00000001, 1,     "2015-11-21", 883789238746, 294.66,  "ASM(1) c ${GLGENRCV} 266.67, ASM(1) d ${rlib.DFLT} 266.67, ASM(3) c ${GLGENRCV} 13.33, ASM(3) d ${rlib.DFLT} 13.33, ASM(4) c ${GLGENRCV} 5.33, ASM(4) d ${rlib.DFLT} 5.33, ASM(9) c ${GLGENRCV} 9.33,ASM(9) d ${rlib.DFLT} 9.33", "I am a comment"

// GenerateReceiptAllocations processes the AcctRule for the supplied rlib.Receipt and generates rlib.ReceiptAllocation records
func GenerateReceiptAllocations(rcpt *rlib.Receipt, xbiz *rlib.XBusiness) error {
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

		// make sure the referenced assessment actually exists
		_, err := rlib.GetAssessment(a.ASMID)
		if err != nil {
			return fmt.Errorf("GenerateReceiptAllocations: Referenced assessment ID %d does not exist\n", a.ASMID)
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
		rlib.InsertReceiptAllocation(&a)
	}
	return nil
}

// CreateReceiptsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateReceiptsFromCSV(sa []string, PmtTypes *map[int64]rlib.PaymentType, lineno int) {
	funcname := "CreateReceiptsFromCSV"
	var xbiz rlib.XBusiness
	var r rlib.Receipt
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 8
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("CreateLedgerMarkers: rlib.Business with designation %s does net exist\n", sa[0])
			return
		}
		r.BID = b1.BID
		rlib.GetXBusiness(r.BID, &xbiz)
	}

	//-------------------------------------------------------------------
	// Find Rental Agreement
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[1])
	re, _ := regexp.Compile("^RA0*(.*)")
	m := re.FindStringSubmatch(s) // returns this pattern:  ["RA0000001" "1"]
	if len(m) > 0 {               // if the prefix was "RA", m will have 2 elements, our number should be the second element
		s = m[1]
	}
	r.RAID, _ = rlib.IntFromString(s, "Rental Agreement number is invalid")
	_, err = rlib.GetRentalAgreement(r.RAID)
	if nil != err {
		fmt.Printf("CreateReceiptsFromCSV: error loading Rental Agreement %s, err = %v\n", sa[1], err)
		return
	}

	//-------------------------------------------------------------------
	// Get the rlib.PaymentType
	//-------------------------------------------------------------------
	r.PMTID, _ = rlib.IntFromString(sa[2], "Payment type is invalid")
	_, ok := (*PmtTypes)[r.PMTID]
	if !ok {
		fmt.Printf("CreateReceiptsFromCSV: Payment type is invalid: %s\n", sa[2])
		return
	}

	//-------------------------------------------------------------------
	// Get the date
	//-------------------------------------------------------------------
	Dt, err := StringToDate(sa[3])
	if err != nil {
		fmt.Printf("CreateReceiptsFromCSV: invalid rlib.Receipt date:  %s\n", sa[3])
		return
	}
	r.Dt = Dt

	//-------------------------------------------------------------------
	// Determine the DocNo
	//-------------------------------------------------------------------
	r.DocNo = strings.TrimSpace(sa[4])
	//fmt.Printf("r.DocNo = %s\n", r.DocNo)

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	r.Amount, _ = rlib.FloatFromString(sa[5], "rlib.Receipt Amount is invalid")

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	r.AcctRule = strings.TrimSpace(sa[6])

	r.Comment = strings.TrimSpace(sa[7])

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(r.AcctRule) == 0 || r.PMTID == 0 ||
		r.Amount == 0 || r.RAID == 0 || r.BID == 0 {
		fmt.Printf("Skipping this record\n")
		return
	}

	rcptid, err := rlib.InsertReceipt(&r)
	if err != nil {
		fmt.Printf("CreateReceiptsFromCSV: error inserting assessment: %v\n", err)
	}
	r.RCPTID = rcptid

	//-------------------------------------------------------------------
	// Create the allocations...
	//-------------------------------------------------------------------
	err = GenerateReceiptAllocations(&r, &xbiz)
	if err != nil {
		fmt.Printf("CreateReceiptsFromCSV: error processing payments: %s\n", err.Error())
		rlib.DeleteReceipt(r.RCPTID)
		rlib.DeleteReceiptAllocations(r.RCPTID)
	}

}

// LoadReceiptsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadReceiptsCSV(fname string, PmtTypes *map[int64]rlib.PaymentType) {
	t := rlib.LoadCSV(fname)
	if len(t) > 1 {
		//-------------------------------------------------------------------
		// Check to see if this rental specialty type is already in the database
		//-------------------------------------------------------------------
		des := strings.TrimSpace(t[1][0])
		if len(des) > 0 {
			b, _ := rlib.GetBusinessByDesignation(des)
			if b.BID < 1 {
				rlib.Ulog("LoadReceiptsCSV: rlib.Business named %s not found\n", des)
				return
			}
			rlib.InitBusinessFields(b.BID)
			rlib.GetDefaultLedgers(b.BID) // the actually loads the RRdb.BizTypes array which is needed by rpn
		}
	}
	for i := 0; i < len(t); i++ {
		CreateReceiptsFromCSV(t[i], PmtTypes, i+1)
	}
}
