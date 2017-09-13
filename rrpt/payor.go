package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// PayorStatement builds a statement for a Payor for a time period
// params
//  bid      - biz id
//  tcid     - tcid of payor
//  d1 - d2  - time range for report
//  internal - true = internal view (show Unapplied Funds section -- receipts
//             from other payors containing unapplied funds)
//             false = external view (do not show Unapplied Funds section)
//============================================================================
func PayorStatement(bid, tcid int64, d1, d2 *time.Time, internal bool) gotable.Table {
	var t gotable.Table
	var xbiz rlib.XBusiness

	const (
		Date           = 0
		Payor          = iota
		Description    = iota
		RAID           = iota
		ASMID          = iota
		RCPTID         = iota
		Rentable       = iota
		UnappliedFunds = iota
		AppliedFunds   = iota
		Assessment     = iota
		Balance        = iota
	)

	//
	// UGH!
	//=======================================================================
	rlib.InitBizInternals(bid, &xbiz)
	// rlib.Console("bid = %d\n", bid)
	_, ok := rlib.RRdb.BizTypes[bid]
	if !ok {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d]", bid)
		t.SetSection3(e.Error())
		return t
	}
	if len(rlib.RRdb.BizTypes[bid].GLAccounts) == 0 {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d].GLAccounts", bid)
		t.SetSection3(e.Error())
		return t
	}
	//=======================================================================
	// UGH!

	t.Init()
	t.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Payor(s)", 25, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Description", 35, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RAID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("ASMID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("RCPTID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Rentable", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Unapplied Funds", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Applied Funds", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Assessment", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Balance", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)

	payors := []int64{tcid}
	payorcache := map[int64]rlib.Transactant{}

	t.SetTitle("Payor Statement\n")
	payorName := rlib.GetNameFromTransactantCache(tcid, payorcache)
	t.SetSection1(fmt.Sprintf("Statement for: %s", payorName))

	var section2 string // includes address and date range
	tr := rlib.Transactant{}
	err := rlib.GetTransactant(tcid, &tr)
	if err != nil {
		t.SetSection3("Unable to get Payor info: " + err.Error())
		return t
	}
	addr := tr.SingleLineAddress()
	section2 += fmt.Sprintf("%s\n%s - %s", addr, d1.Format(rlib.RRDATEREPORTFMT), d2.Format(rlib.RRDATEREPORTFMT))
	t.SetSection2(section2)

	m, err := rlib.PayorsStatement(bid, payors, d1, d2)
	if err != nil {
		t.SetSection3("Error from PayorsStatement: " + err.Error())
		return t
	}

	//------------------------------------------------------
	// Generate the Receipt Summary
	//------------------------------------------------------
	t.AddRow()
	t.Puts(-1, Description, "*** RECEIPT SUMMARY ***")
	lenmRL := len(m.RL)
	if len(m.RL) == 0 {
		t.Puts(-1, Description, "No receipts this period")
	} else {
		for i := 0; i < len(m.RL); i++ {
			if m.RL[i].R.TCID != tcid {
				continue
			}
			t.AddRow()
			t.Putd(-1, Date, m.RL[i].R.Dt)
			t.Puts(-1, Payor, rlib.GetNameFromTransactantCache(m.RL[i].R.TCID, payorcache))
			t.Puts(-1, RCPTID, rlib.IDtoShortString("RCPT", m.RL[i].R.RCPTID))
			t.Puts(-1, Description, "Receipt "+m.RL[i].R.DocNo)
			t.Putf(-1, UnappliedFunds, m.RL[i].Unallocated)
			t.Putf(-1, AppliedFunds, m.RL[i].Allocated)
			t.Putf(-1, Balance, m.RL[i].R.Amount)
		}
	}
	t.AddRow()

	//------------------------------------------------------
	// Unapplied Funds...
	//------------------------------------------------------
	if internal {
		t.AddRow()
		t.Puts(-1, Description, "*** UNAPPLIED FUNDS ***")
		if len(m.RL) == 0 {
			t.Puts(-1, Description, "No allocations this period")
		} else {
			for i := 0; i < lenmRL; i++ {
				if m.RL[i].R.TCID == tcid {
					continue
				}
				t.AddRow()
				t.Putd(-1, Date, m.RL[i].R.Dt)
				t.Puts(-1, Payor, rlib.GetNameFromTransactantCache(m.RL[i].R.TCID, payorcache))

				//----------------------------------------------------
				// If the payor only has one RAID and it is THIS one
				// then we can list the details of the receipt
				//----------------------------------------------------
				l1 := rlib.GetRentalAgreementsByPayorRange(bid, m.RL[i].R.TCID, d1, d2)
				if len(l1) == 1 {
					t.Puts(-1, RAID, rlib.IDtoShortString("RA", l1[0].RAID))
					t.Puts(-1, RCPTID, rlib.IDtoShortString("RCPT", m.RL[i].R.RCPTID))
					t.Puts(-1, Description, "Receipt "+m.RL[i].R.DocNo)
					t.Putf(-1, UnappliedFunds, m.RL[i].Unallocated)
					t.Putf(-1, AppliedFunds, m.RL[i].Allocated)
					t.Putf(-1, Balance, m.RL[i].R.Amount)
				} else {
					t.Puts(-1, Description, "TBD")
				}
			}
		}
	}
	t.AddRow()

	//------------------------------------------------------
	// Generate the per-Rental-Agreement information...
	//------------------------------------------------------
	for i := 0; i < len(m.RAB); i++ { // for each RA
		raidstr := rlib.IDtoShortString("RA", m.RAB[i].RAID)
		ra, err := rlib.GetRentalAgreement(m.RAB[i].RAID)
		if err != nil {
			rlib.LogAndPrintError("PayorStatement", err)
			continue
		}
		rentableName := ra.GetTheRentableName(d1, d2)
		t.AddRow()
		t.Puts(-1, Description, fmt.Sprintf("*** RENTAL AGREEMENT %d ***", m.RAB[i].RAID))
		t.AddRow()
		t.Puts(-1, Description, "Opening balance")
		t.Putd(-1, Date, m.RAB[i].DtStart)
		t.Putf(-1, Balance, m.RAB[i].OpeningBal)

		//------------------------
		// init running totals
		//------------------------
		bal := m.RAB[i].OpeningBal
		asmts := float64(0)
		applied := asmts
		// unapplied := asmts

		for j := 0; j < len(m.RAB[i].Stmt); j++ { // for each line in the statement
			t.AddRow()
			t.Putd(-1, Date, m.RAB[i].Stmt[j].Dt)
			t.Puts(-1, RAID, raidstr)
			t.Puts(-1, Rentable, rentableName)
			if m.RAB[i].Stmt[j].TCID > 0 {
				t.Puts(-1, Payor, rlib.IDtoShortString("TC", m.RAB[i].Stmt[j].TCID))
			}

			descr := ""
			if m.RAB[i].Stmt[j].Reverse {
				descr += "REVERSAL: "
			}
			amt := m.RAB[i].Stmt[j].Amt

			switch m.RAB[i].Stmt[j].T {
			case 1: // assessments
				t.Putf(-1, Assessment, amt)
				if m.RAB[i].Stmt[j].A.ARID > 0 { // The description will be the name of the Account Rule...
					descr += rlib.RRdb.BizTypes[bid].AR[m.RAB[i].Stmt[j].A.ARID].Name
				} else {
					descr += rlib.RRdb.BizTypes[bid].GLAccounts[m.RAB[i].Stmt[j].A.ATypeLID].Name
				}
				if m.RAB[i].Stmt[j].A.ASMID > 0 {
					t.Puts(-1, ASMID, rlib.IDtoShortString("ASM", m.RAB[i].Stmt[j].A.ASMID))
				}
				if m.RAB[i].Stmt[j].A.RAID > 0 { // Payor(s) = all payors associated with RentalAgreement
					pyrs := rlib.GetRentalAgreementPayorsInRange(m.RAB[i].Stmt[j].A.RAID, d1, d2)
					sa := []string{}
					for k := 0; k < len(pyrs); k++ {
						sa = append(sa, rlib.GetNameFromTransactantCache(pyrs[k].TCID, payorcache))
					}
					t.Puts(-1, Payor, strings.Join(sa, ","))
				}
				if !m.RAB[i].Stmt[j].Reverse { // update running totals if not a reversal
					asmts += amt
					bal += amt
				} else {
					descr += " (" + m.RAB[i].Stmt[j].A.Comment + ")"
				}
			case 2: // receipts
				t.Putf(-1, AppliedFunds, amt)
				rcptid := m.RAB[i].Stmt[j].R.RCPTID
				descr += "Receipt allocation"
				if rcptid > 0 {
					t.Puts(-1, RCPTID, rlib.IDtoShortString("RCPT", rcptid))
					rcpt := rlib.GetReceipt(rcptid)
					if rcpt.RCPTID > 0 {
						name := rlib.GetNameFromTransactantCache(rcpt.TCID, payorcache)
						t.Puts(-1, Payor, name)
					}
				}
				if m.RAB[i].Stmt[j].A.ASMID > 0 {
					t.Puts(-1, ASMID, rlib.IDtoShortString("ASM", m.RAB[i].Stmt[j].A.ASMID))
				}
				if !m.RAB[i].Stmt[j].Reverse {
					applied += amt
					bal -= amt
				} else {
					rcpt := rlib.GetReceipt(m.RAB[i].Stmt[j].R.RCPTID)
					if rcpt.RCPTID > 0 && len(rcpt.Comment) > 0 {
						descr += " (" + rcpt.Comment + ")"
					}
				}
			}
			t.Putf(-1, Balance, bal)
			t.Puts(-1, Description, descr)
		}
		t.AddLineAfter(len(t.Row) - 1)
		t.AddRow()
		t.Putd(-1, Date, m.RAB[i].DtStop)
		t.Puts(-1, Description, "Closing balance")
		t.Putf(-1, AppliedFunds, applied)
		t.Putf(-1, Assessment, asmts)
		t.Putf(-1, Balance, m.RAB[i].ClosingBal)
		t.AddRow()
	}
	return t
}

// RRPayorStatement is a report regarding payor statement, used to be download in pdf, csv format
func RRPayorStatement(ri *ReporterInfo) gotable.Table {
	// find tcid from query params
	var (
		tcid     int64
		internal bool
	)
	tcidStr := ri.QueryParams.Get("tcid")
	tcid, _ = strconv.ParseInt(tcidStr, 10, 64)

	internalStr := ri.QueryParams.Get("internal")
	if internalStr == "1" { //only internal if internal=1
		internal = true
	}

	return PayorStatement(ri.Bid, tcid, &ri.D1, &ri.D2, internal)
}
