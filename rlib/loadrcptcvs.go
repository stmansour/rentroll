package rlib

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// // AcctRule is a structure of the 3-tuple that makes up a whole part of an AcctRule
// type CSVAcctRule struct {
// 	Action  string // "d" = debit, "c" = credit
// 	Account string // GL No for the account
// 	Amount  string // use the entire amount of the assessment or deposit, otherwise the amount to use
// 	ASMID   string // Used only for ReceiptAllocation; the assessment that caused this payment
// }

// CVS record format:
// 0    1           2      3             4        5         6
// BID, RAID,       PMTID, Dt,           Amount,  AcctRule, Comment
// REH, RA00000001, 2,     "2004-01-01", 1000.00, "ASM(7) d ${DFLTCASH} _, ASM(7) c 11002 _"
// REH, RA00000001, 1,     "2015-11-21",  294.66, "ASM(1) c ${DFLTGENRCV} 266.67, ASM(1) d ${DFLTCASH} 266.67, ASM(3) c ${DFLTGENRCV} 13.33, ASM(3) d ${DFLTCASH} 13.33, ASM(4) c ${DFLTGENRCV} 5.33, ASM(4) d ${DFLTCASH} 5.33, ASM(9) c ${DFLTGENRCV} 9.33,ASM(9) d ${DFLTCASH} 9.33", "I am a comment"

// GenerateReceiptAllocations processes the AcctRule for the supplied Receipt and generates ReceiptAllocation records
func GenerateReceiptAllocations(rcpt *Receipt, xbiz *XBusiness) error {
	var d1 = time.Date(rcpt.Dt.Year(), rcpt.Dt.Month(), 1, 0, 0, 0, 0, time.UTC)
	var d2 = d1.AddDate(0, 0, 31)
	t := ParseAcctRule(xbiz, 0, &d1, &d2, rcpt.AcctRule, rcpt.Amount, 1.0)
	u := make(map[int64][]int64)

	// First, group together all entries that refer to a single ASMID into a list of lists
	for i := int64(0); i < int64(len(t)); i++ {
		u[t[i].ASMID] = append(u[t[i].ASMID], i)
	}
	// Process each list in the list of lists.
	for k, v := range u {
		var a ReceiptAllocation
		a.AcctRule = ""
		a.ASMID = k
		a.Amount = t[int(v[0])].Amount
		a.RCPTID = rcpt.RCPTID

		// make sure the referenced assessment actually exists
		_, err := GetAssessment(a.ASMID)
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
		InsertReceiptAllocation(&a)
	}
	return nil
}

// CreateReceiptsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateReceiptsFromCSV(sa []string, PmtTypes *map[int64]PaymentType) {
	var xbiz XBusiness
	var r Receipt
	var err error
	des := strings.ToLower(strings.TrimSpace(sa[0]))
	if des == "designation" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Make sure the business is in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		b1, _ := GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			Ulog("CreateLedgerMarkers: business with designation %s does net exist\n", sa[0])
			return
		}
		r.BID = b1.BID
		GetXBusiness(r.BID, &xbiz)
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
	r.RAID, _ = IntFromString(s, "Rental Agreement number is invalid")
	_, err = GetRentalAgreement(r.RAID)
	if nil != err {
		fmt.Printf("CreateReceiptsFromCSV: error loading Rental Agreement %s, err = %v\n", sa[1], err)
		return
	}

	//-------------------------------------------------------------------
	// Get the PaymentType
	//-------------------------------------------------------------------
	r.PMTID, _ = IntFromString(sa[2], "Payment type is invalid")
	_, ok := (*PmtTypes)[r.PMTID]
	if !ok {
		fmt.Printf("CreateReceiptsFromCSV: Payment type is invalid: %s\n", sa[2])
		return
	}

	//-------------------------------------------------------------------
	// Get the date
	//-------------------------------------------------------------------

	Dt, err := time.Parse(RRDATEINPFMT, strings.TrimSpace(sa[3]))
	if err != nil {
		fmt.Printf("CreateReceiptsFromCSV: invalid receipt date:  %s\n", sa[3])
		return
	}
	r.Dt = Dt

	//-------------------------------------------------------------------
	// Determine the amount
	//-------------------------------------------------------------------
	r.Amount, _ = FloatFromString(sa[4], "Receipt Amount is invalid")

	//-------------------------------------------------------------------
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	r.AcctRule = strings.TrimSpace(sa[5])

	r.Comment = strings.TrimSpace(sa[6])

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(r.AcctRule) == 0 || r.PMTID == 0 ||
		r.Amount == 0 || r.RAID == 0 || r.BID == 0 {
		fmt.Printf("Skipping this record\n")
		return
	}

	rcptid, err := InsertReceipt(&r)
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
		DeleteReceipt(r.RCPTID)
		DeleteReceiptAllocations(r.RCPTID)
	}

}

// LoadReceiptsCSV loads a csv file with a chart of accounts and creates ledger markers for each
func LoadReceiptsCSV(fname string, PmtTypes *map[int64]PaymentType) {
	t := LoadCSV(fname)
	if len(t) > 1 {
		//-------------------------------------------------------------------
		// Check to see if this rental specialty type is already in the database
		//-------------------------------------------------------------------
		des := strings.TrimSpace(t[1][0])
		if len(des) > 0 {
			b, _ := GetBusinessByDesignation(des)
			if b.BID < 1 {
				Ulog("LoadReceiptsCSV: Business named %s not found\n", des)
				return
			}
			InitBusinessFields(b.BID)
			GetDefaultLedgerMarkers(b.BID) // the actually loads the RRdb.BizTypes array which is needed by rpn
		}
	}
	for i := 0; i < len(t); i++ {
		CreateReceiptsFromCSV(t[i], PmtTypes)
	}
}
