package rcsv

import (
	"fmt"
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
		a1, _ := rlib.GetAssessment(a.ASMID)
		if a1.ASMID == 0 {
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
		rcpt.RA = append(rcpt.RA, a)
	}
	return nil
}

// CreateReceiptsFromCSV reads an assessment type string array and creates a database record for the assessment type
func CreateReceiptsFromCSV(sa []string, PmtTypes *map[int64]rlib.PaymentType, lineno int) (string, int) {
	funcname := "CreateReceiptsFromCSV"
	var xbiz rlib.XBusiness
	var r rlib.Receipt
	var err error
	bud := strings.ToLower(strings.TrimSpace(sa[0]))
	rs := ""

	const (
		BUD      = 0
		RAID     = iota
		PMTID    = iota
		Dt       = iota
		DocNo    = iota
		Amount   = iota
		AcctRule = iota
		Comment  = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"RAID", RAID},
		{"PMTID", PMTID},
		{"Dt", Dt},
		{"DocNo", DocNo},
		{"Amount", Amount},
		{"AcctRule", AcctRule},
		{"Comment", Comment},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	//-------------------------------------------------------------------
	// Make sure the rlib.Business is in the database
	//-------------------------------------------------------------------
	if len(bud) > 0 {
		b1 := rlib.GetBusinessByDesignation(bud)
		if len(b1.Designation) == 0 {
			rs := fmt.Sprintf("CreateLedgerMarkers: rlib.Business with designation %s does not exist\n", sa[0])
			return rs, CsvErrorSensitivity
		}
		r.BID = b1.BID
		rlib.GetXBusiness(r.BID, &xbiz)
	}

	//-------------------------------------------------------------------
	// Find Rental Agreement
	//-------------------------------------------------------------------
	r.RAID = CSVLoaderGetRAID(sa[RAID]) // this should probably go away, we should select it from an Assessment in the AcctRule

	_, err = rlib.GetRentalAgreement(r.RAID)
	if nil != err {
		rs := fmt.Sprintf("%s: line %d -  error loading Rental Agreement %s, err = %v\n", funcname, lineno, sa[RAID], err)
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Get the rlib.PaymentType
	//-------------------------------------------------------------------
	r.PMTID, _ = rlib.IntFromString(sa[PMTID], "Payment type is invalid")
	_, ok := (*PmtTypes)[r.PMTID]
	if !ok {
		rs := fmt.Sprintf("%s: line %d -  Payment type is invalid: %s\n", funcname, lineno, sa[PMTID])
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Get the date
	//-------------------------------------------------------------------
	dt, err := rlib.StringToDate(sa[Dt])
	if err != nil {
		rs := fmt.Sprintf("%s: line %d -  invalid rlib.Receipt date:  %s\n", funcname, lineno, sa[Dt])
		return rs, CsvErrorSensitivity
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
	// Set the AcctRule.  No checking for now...
	//-------------------------------------------------------------------
	r.AcctRule = strings.TrimSpace(sa[AcctRule])

	r.Comment = strings.TrimSpace(sa[Comment])

	//-------------------------------------------------------------------
	// Make sure everything that needs to be set actually got set...
	//-------------------------------------------------------------------
	if len(r.AcctRule) == 0 || r.PMTID == 0 ||
		r.Amount == 0 || r.RAID == 0 || r.BID == 0 {
		rs := fmt.Sprintf("Skipping this record\n")
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Make sure there's no duplicate...
	//-------------------------------------------------------------------
	rdup := rlib.GetReceiptDuplicate(&r.Dt, r.Amount, r.DocNo)
	if rdup.RCPTID != 0 {
		rs = fmt.Sprintf("%s: line %d - this is a duplicate of an existing receipt: %s\n", funcname, lineno, rdup.IDtoString())
		return rs, CsvErrorSensitivity
	}

	rcptid, err := rlib.InsertReceipt(&r)
	if err != nil {
		rs := fmt.Sprintf("%s: line %d -  error inserting receipt: %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}
	r.RCPTID = rcptid

	//-------------------------------------------------------------------
	// Create the allocations...
	//-------------------------------------------------------------------
	err = GenerateReceiptAllocations(&r, &xbiz)
	if err != nil {
		rs := fmt.Sprintf("%s: line %d -  error processing receipt: %s\n", funcname, lineno, err.Error())
		rlib.DeleteReceipt(r.RCPTID)
		rlib.DeleteReceiptAllocations(r.RCPTID)
		return rs, CsvErrorSensitivity
	}

	//-------------------------------------------------------------------
	// Process the receipt...
	//-------------------------------------------------------------------
	rlib.ProcessNewReceipt(Rcsv.Xbiz, &Rcsv.DtStart, &Rcsv.DtStop, &r)

	return rs, 0
}

// LoadReceiptsCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadReceiptsCSV(fname string) string {
	rs := ""
	PmtTypes := rlib.GetPaymentTypes()
	t := rlib.LoadCSV(fname)
	if len(t) > 1 {
		//-------------------------------------------------------------------
		// Check to see if this rental specialty type is already in the database
		//-------------------------------------------------------------------
		des := strings.TrimSpace(t[1][0])
		if len(des) > 0 {
			b := rlib.GetBusinessByDesignation(des)
			if b.BID < 1 {
				rs += fmt.Sprintf("LoadReceiptsCSV: rlib.Business named %s not found\n", des)
				return rs
			}
			rlib.InitBusinessFields(b.BID)
			rlib.GetDefaultLedgers(b.BID) // the actually loads the RRdb.BizTypes array which is needed by rpn
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i][0] == "#" {
			continue
		}
		s, err := CreateReceiptsFromCSV(t[i], &PmtTypes, i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
