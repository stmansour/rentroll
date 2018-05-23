package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
	"time"
)

type jprintctx struct {
	ReportStart time.Time
	ReportStop  time.Time
}

// func setTitle(tbl *gotable.Table, xbiz *rlib.XBusiness, d1, d2 *time.Time) {
// 	s := "JOURNAL\n"
// 	s += fmt.Sprintf("Business: %-13s\n", xbiz.P.Name)
// 	s += fmt.Sprintf("Period:   %s - %s\n\n", d1.Format(rlib.RRDATEFMT), d2.AddDate(0, 0, -1).Format(rlib.RRDATEFMT))
// 	tbl.SetTitle(s)
// }

func processAcctRuleAmount(ctx context.Context, tbl *gotable.Table, xbiz *rlib.XBusiness, rid int64, d time.Time, rule string, raid int64, r *rlib.Rentable, amt float64) error {
	const funcname = "processAcctRuleAmount"
	var (
		err error
	)

	m, err := rlib.ParseAcctRule(ctx, xbiz, rid, &d, &d, rule, amt, float64(1))
	if err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {
		amt := m[i].Amount
		if m[i].Action == "c" {
			amt = -amt
		}

		// ---------------------------------------------------------
		// This code essentially skips amounts that calculate
		// to something less than .0001 cents (a rounding error)
		x := amt
		if x < 0 {
			x = -x
		}
		if x < 0.0001 {
			continue
		}
		// ---------------------------------------------------------

		l, err := rlib.GetLedgerByGLNo(ctx, xbiz.P.BID, m[i].Account)
		if err != nil {
			// TODO(Steve): in case of error do we want to continue?
			rlib.LogAndPrintError(funcname, err)
			// debug.PrintStack()
			rlib.LogAndPrint("%s: Could not get GLAccount named %s in Business %d\n", funcname, m[i].Account, r.BID)
			rlib.LogAndPrint("%s: rule = \"%s\"\n", funcname, rule)
			continue
		}

		if 0 == l.LID {
			// debug.PrintStack()
			rlib.LogAndPrint("%s: Could not get GLAccount named %s in Business %d\n", funcname, m[i].Account, r.BID)
			rlib.LogAndPrint("%s: rule = \"%s\"\n", funcname, rule)
			continue
		}

		// printDatedJournalEntryRJ(l.Name, d, fmt.Sprintf("%d", raid), r.RentableName, m[i].Account, amt)
		tbl.AddRow()
		tbl.Puts(-1, 1, l.Name)
		tbl.Putd(-1, 2, d)
		tbl.Puts(-1, 3, rlib.IDtoShortString("RA", raid))
		tbl.Puts(-1, 4, r.RentableName)
		tbl.Puts(-1, 5, m[i].Account)
		tbl.Putf(-1, 6, amt)
	}

	return err
}

func textPrintJournalAssessment(ctx context.Context, tbl *gotable.Table, jctx *jprintctx, xbiz *rlib.XBusiness, j *rlib.Journal, a *rlib.Assessment, r *rlib.Rentable, rentDuration, assessmentDuration int64) {
	const funcname = "textPrintJournalAssessment"

	var (
		err  error
		s    string
		rtid int64
	)

	if a.ARID > 0 {
		s = rlib.RRdb.BizTypes[xbiz.P.BID].AR[a.ARID].Name
	} else if a.ATypeLID > 0 {
		s = rlib.RRdb.BizTypes[xbiz.P.BID].GLAccounts[a.ATypeLID].Name
	}

	//-------------------------------------------------------------------------------------
	// For reporting, we want to show any proration that needs to take place. To determine
	// whether or not there is any proration:
	// 1. Check to see if the Accrual period for the Rentable in question is greater than
	//    the ProrationCycle.
	//        *  NO: there is no proration, we don't need to report anything, pf = 1
	// 2. What percent of the accrual period was the Rentable "occupied" during the range of interest
	//        *  create a time range equal to the report period [reportDtStart - reportDtStop]
	//        *  if this range is > Accrual Period, trim the range accordingly
	//        *  if this range is > "occupiedrange", trim the range accordingly
	//        *  if the resulting range == Accrual Period then we don't need to report anything, pf = 1
	// 3. Report the prorate factor numerator and denominator:
	//           pf = (resulting range duration)/AccrualPeriod (both in units of the ProrationCycle)
	//-------------------------------------------------------------------------------------
	if r.RID > 0 {
		_, pro, rti, err := rlib.GetRentCycleAndProration(ctx, r, &a.Start, xbiz)
		rtid = rti
		if err != nil {
			rlib.Ulog("textPrintJournalAssessment: error getting RentCycle and Proration: err = %s\n", err.Error())
			return
		}
		if a.RentCycle > pro && pro != 0 && a.ProrationCycle != 0 { // if accrual > proration then we *may* need to show prorate info
			d1 := jctx.ReportStart // start with the report range
			d2 := jctx.ReportStop  // start with the report range
			if j.Dt.After(d1) {    // if this assessment is later move the start time
				d1 = j.Dt
			}
			tmp := d1.Add(rlib.CycleDuration(a.RentCycle, d1)) // start + accrual duration
			if tmp.Before(d2) {                                // if this occurs prior to the range end...
				d2 = tmp // snap the range end
			}
			ra, err := rlib.GetRentalAgreement(ctx, a.RAID) // need rental agreement to find Possession time
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				return
			}

			if ra.RAID > 0 { // if we found the rental agreement
				if ra.RentStart.After(d1) { // if possession started after d1
					d1 = ra.RentStart // snap the begin time
				}
				if ra.RentStop.Before(d2) { // if possession ended prior to d2
					d2 = ra.RentStop // snap the end time
				}
			}

			units := rlib.CycleDuration(pro, d1) // duration of the unit for proration
			numerator := d2.Sub(d1)
			denominator := rlib.GetProrationRange(d1, d2, a.RentCycle, pro)

			if numerator != denominator {
				s += fmt.Sprintf(" (%d/%d %s)", numerator/units, denominator/units, rlib.ProrationUnits(pro))
			}
		}
	}

	if rtid > 0 {
		s += fmt.Sprintf("  %s", r.RentableName) + " [" + xbiz.RT[rtid].Style
		if a.RentCycle > rlib.RECURNONE {
			s += ", " + rlib.RentalPeriodToString(a.RentCycle)
		}
		s += "] " + j.Comment
	}

	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoShortString())
	tbl.Puts(-1, 1, s)

	for i := 0; i < len(j.JA); i++ {
		err = processAcctRuleAmount(ctx, tbl, xbiz, r.RID, j.Dt, j.JA[i].AcctRule, j.JA[i].RAID, r, j.JA[i].Amount)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			continue
		}
	}

	tbl.AddRow() // nothing in this line, it's blank
}

func printJournalExpense(ctx context.Context, tbl *gotable.Table, xbiz *rlib.XBusiness, j *rlib.Journal, a *rlib.Expense, r *rlib.Rentable) {
	const funcname = "printJournalExpense"
	var (
	// err error
	)

	s := rlib.RRdb.BizTypes[xbiz.P.BID].AR[a.ARID].Name
	if a.RID > 0 {
		rtr, err := rlib.GetRentableTypeRefForDate(ctx, r.RID, &j.Dt)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			return
		}

		s += fmt.Sprintf("  %s [%s]", r.RentableName, xbiz.RT[rtr.RTID].Style)
	}
	s += " " + j.Comment

	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoShortString())
	tbl.Puts(-1, 1, s)

	for i := 0; i < len(j.JA); i++ {
		clid := rlib.RRdb.BizTypes[j.BID].AR[a.ARID].CreditLID
		dlid := rlib.RRdb.BizTypes[j.BID].AR[a.ARID].DebitLID
		rn := ""
		if j.JA[i].RID > 0 {
			rn = r.RentableName
		}
		raid := ""
		if j.JA[i].RAID > 0 {
			raid = rlib.IDtoShortString("RA", j.JA[i].RAID)
		}
		tbl.AddRow()
		tbl.Puts(-1, 1, rlib.RRdb.BizTypes[j.BID].GLAccounts[dlid].Name)
		tbl.Putd(-1, 2, j.Dt)
		tbl.Puts(-1, 3, raid)
		tbl.Puts(-1, 4, rn)
		tbl.Puts(-1, 5, rlib.RRdb.BizTypes[j.BID].GLAccounts[dlid].GLNumber)
		tbl.Putf(-1, 6, j.Amount)

		tbl.AddRow()
		tbl.Puts(-1, 1, rlib.RRdb.BizTypes[j.BID].GLAccounts[clid].Name)
		tbl.Putd(-1, 2, j.Dt)
		tbl.Puts(-1, 3, raid)
		tbl.Puts(-1, 4, rn)
		tbl.Puts(-1, 5, rlib.RRdb.BizTypes[j.BID].GLAccounts[clid].GLNumber)
		tbl.Putf(-1, 6, -j.Amount)
	}

	tbl.AddRow() // nothing in this line, it's blank
}

func textPrintJournalReceipt(ctx context.Context, tbl *gotable.Table, ri *ReporterInfo, jctx *jprintctx, j *rlib.Journal, rcpt *rlib.Receipt) {
	const funcname = "textPrintJournalReceipt"
	// fmt.Printf("Entered: %s,   JID = %d, RCPTID = %d\n", funcname, j.JID, rcpt.RCPTID)
	// The receipt has the payor TCID.  We get the payor name from the receipt
	var (
		t  rlib.Transactant
		ps string
	)

	if err := rlib.GetTransactant(ctx, rcpt.TCID, &t); err != nil {
		// fmt.Printf("<< rcpt.TCID = %d   db err = %s>>\n", rcpt.TCID, err.Error())
		// No transactant ID.  See if there is an OtherPayor. If so use it, if not, get the payors associated with this journal entry...
		if len(rcpt.OtherPayorName) > 0 {
			ps = rcpt.OtherPayorName
		} else {
			// fmt.Printf("Will look up all RAIDs in ReceiptAllocations. len(rcpt.RA) = %d\n", len(rcpt.RA))
			var mm = map[int64]int64{} // mm[raid]raid -- just to keep track of what we've found
			for i := 0; i < len(rcpt.RA); i++ {
				raid := rcpt.RA[i].RAID
				// fmt.Printf("rcpt.RA[i].RAID = %d\n", raid)
				_, ok := mm[raid]
				if !ok {
					// fmt.Printf("mm[raid] was not found. Will search RAID %d for payors on %s\n", raid, j.Dt.Format(rlib.RRDATEFMT4))
					mm[raid] = raid
					n, err := rlib.GetRentalAgreementPayorsInRange(ctx, raid, &j.Dt, &j.Dt)
					if err != nil {
						rlib.LogAndPrintError(funcname, err)
						continue
					}

					// fmt.Printf("found %d payors\n", len(n))
					for j := 0; j < len(n); j++ {
						var t rlib.Transactant
						if err := rlib.GetTransactant(ctx, n[j].TCID, &t); err != nil {
							rlib.LogAndPrintError(funcname, err)
							continue
						}
						if len(ps) > 0 {
							ps += ","
						}
						ps += t.GetTransactantLastName() // could be multiple names, use Lastname only
					}
					// fmt.Printf("ps = %s\n", ps)
					if len(ps) == 0 {
						ps = fmt.Sprintf("No payors for RAID %d on %s\n", raid, j.Dt.Format(rlib.RRDATEFMT4))
					}
				}
			}
		}
	} else {
		ps = t.GetFullTransactantName() // we know this will be only one name, so we should have the space for the full name
	}

	s := fmt.Sprintf("Payment - %s   #%s  %.2f", ps, rcpt.DocNo, rcpt.Amount)
	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoShortString())
	tbl.Puts(-1, 1, s)

	// PROCESS EVERY RECEIPT ALLOCATION IN OUR DATE RANGE...
	for i := 0; i < len(rcpt.RA); i++ {
		processThis := false // assume this Journal entry applies to this receipt allocation
		for k := 0; k < len(j.JA); k++ {
			if rcpt.RA[i].AcctRule == j.JA[k].AcctRule { // find the account rule that matches...
				processThis = true
				break
			}
		}
		if !processThis { // this account rule is described by a different
			continue
		}
		// first do a quick reject test -- only show those that happen in the time range of the report
		rdt := rcpt.RA[i].Dt
		if !((ri.D1.Equal(rdt) || ri.D1.Before(rdt)) && ri.D2.After(rdt)) {
			continue
		}
		// TODO(Steve): should we ignore error?
		a, _ := rlib.GetAssessment(ctx, rcpt.RA[i].ASMID)
		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			continue
		}

		// if r.RID == 0 {
		// 	rlib.LogAndPrint("%s: rcpt.RA[%d].RCPAID = %d, r.RID = 0, rcpt.RA[i].ASMID = %d, a.RID = %d\n", funcname, i, rcpt.RA[i].RCPAID, rcpt.RA[i].ASMID, a.RID)
		// 	continue
		// }
		m, err := rlib.ParseAcctRule(ctx, ri.Xbiz, r.RID, &jctx.ReportStart, &jctx.ReportStop, rcpt.RA[i].AcctRule, rcpt.RA[i].Amount, 1.0)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			continue
		}
		// fmt.Printf("%s: acctrule = %s     Amount = %.2f\n", funcname, rcpt.RA[i].AcctRule, rcpt.RA[i].Amount)
		// for k := 0; k < len(m); k++ {
		// 	fmt.Printf("%d. .Account = %s, .Amount = %.2f   .ASMID = %d\n", k, m[k].Account, m[k].Amount, m[k].ASMID)
		// }
		// printJournalSubtitle("\t" + rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[a.ATypeLID].Name)
		// fmt.Printf("rcpt.RA[i].ASMID = %d, a.ASMID = %d, a.RID = %d\n", rcpt.RA[i].ASMID, a.ASMID, a.RID)
		// if r.BID == 0 {
		// 	fmt.Printf("r.BID == 0:  r.RID = %d\n", r.RID)
		// }

		// tbl.AddRow()
		// tbl.Puts(-1, 1, rlib.RRdb.BizTypes[ri.Xbiz.P.BID].GLAccounts[a.ATypeLID].Name)

		for k := 0; k < len(m); k++ {
			l, err := rlib.GetLedgerByGLNo(ctx, j.BID, m[k].Account)
			if err != nil {
				rlib.LogAndPrintError(funcname, err)
				rlib.LogAndPrint("%s: Could not get GLAccount named %s in Business %d\n", funcname, m[i].Account, r.BID)
				rlib.LogAndPrint("%s: rule = \"%s\"\n", funcname, rcpt.RA[i].AcctRule)
				continue
			}
			if 0 == l.LID {
				rlib.LogAndPrint("%s: Could not get GLAccount named %s in Business %d\n", funcname, m[i].Account, r.BID)
				rlib.LogAndPrint("%s: rule = \"%s\"\n", funcname, rcpt.RA[i].AcctRule)
				continue
			}
			amt := m[k].Amount
			if m[k].Action == "c" {
				amt = -amt
			}
			// s := fmt.Sprintf("%d", a.RAID)
			// printDatedJournalEntryRJ(l.Name, rcpt.Dt, s, r.RentableName, m[k].Account, amt)
			rs := ""
			if a.RAID > 0 {
				rs = rlib.IDtoShortString("RA", a.RAID)
			}
			tbl.AddRow()
			tbl.Puts(-1, 1, l.Name)
			tbl.Putd(-1, 2, rcpt.Dt)
			tbl.Puts(-1, 3, rs)
			tbl.Puts(-1, 4, r.RentableName)
			tbl.Puts(-1, 5, m[k].Account)
			tbl.Putf(-1, 6, amt)
		}
	}
	tbl.AddRow() // nothing in this line, it's blank
}

func textPrintJournalUnassociated(ctx context.Context, tbl *gotable.Table, xbiz *rlib.XBusiness, jctx *jprintctx, j *rlib.Journal) {
	const funcname = "textPrintJournalUnassociated"

	var (
		r   rlib.Rentable
		err error
	)
	// rlib.Console("textPrintJournalUnassociated\n")
	err = rlib.GetRentableByID(ctx, j.ID, &r) // j.ID is RID when it is unassociated (RAID == 0)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return
	}
	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoShortString())
	tbl.Puts(-1, 1, fmt.Sprintf("Unassociated: %s %s", r.RentableName, j.Comment))
	for i := 0; i < len(j.JA); i++ {
		err = processAcctRuleAmount(ctx, tbl, xbiz, j.JA[i].RID, j.Dt, j.JA[i].AcctRule, 0, &r, j.JA[i].Amount)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			continue
		}
	}
	tbl.AddRow() // separater line
}

func textPrintJournalXfer(tbl *gotable.Table, ri *ReporterInfo, jctx *jprintctx, j *rlib.Journal) {
	tbl.AddRow()
	tbl.Puts(-1, 0, j.IDtoShortString())
	s := "Transfer"
	if len(j.Comment) > 0 {
		s += fmt.Sprintf(" (%s)", j.Comment)
	}
	tbl.Puts(-1, 1, s)

	if len(j.JA[0].AcctRule) > 0 {
		var clid, dlid int64
		m := rlib.ParseSimpleAcctRule(j.JA[0].AcctRule)
		for i := 0; i < len(m); i++ {
			var lid int64
			found := false
			for _, v := range rlib.RRdb.BizTypes[j.BID].GLAccounts {
				if m[i].Account == v.GLNumber {
					lid = v.LID
					found = true
				}
			}
			if !found {
				continue
			}
			if m[i].Action == "d" {
				dlid = lid
			} else {
				clid = lid
			}
		}

		if dlid > 0 && clid > 0 {
			tbl.AddRow()
			tbl.Puts(-1, 1, "from "+rlib.RRdb.BizTypes[j.BID].GLAccounts[clid].Name)
			tbl.Putd(-1, 2, j.Dt)
			tbl.Puts(-1, 3, rlib.IDtoShortString("RA", j.JA[0].RAID))
			tbl.Puts(-1, 5, rlib.RRdb.BizTypes[j.BID].GLAccounts[clid].GLNumber)
			tbl.Putf(-1, 6, -j.Amount)

			tbl.AddRow()
			tbl.Puts(-1, 1, "to "+rlib.RRdb.BizTypes[j.BID].GLAccounts[dlid].Name)
			tbl.Putd(-1, 2, j.Dt)
			tbl.Puts(-1, 3, rlib.IDtoShortString("RA", j.JA[0].RAID))
			tbl.Puts(-1, 5, rlib.RRdb.BizTypes[j.BID].GLAccounts[dlid].GLNumber)
			tbl.Putf(-1, 6, j.Amount)
		}
	}
	tbl.AddRow() // nothing in this line, it's blank
}

func textPrintJournalEntry(ctx context.Context, tbl *gotable.Table, ri *ReporterInfo, jctx *jprintctx, j *rlib.Journal, rentDuration, assessmentDuration int64) {
	switch j.Type {
	case rlib.JNLTYPEUNAS:
		textPrintJournalUnassociated(ctx, tbl, ri.Xbiz, jctx, j)
	case rlib.JNLTYPERCPT:
		rcpt, err := rlib.GetReceipt(ctx, j.ID)
		if err != nil {
			rlib.LogAndPrint("Failed to get receipt for j.ID = %d,  j.JID = %d\n", j.ID, j.JID)
			return
		}
		if rcpt.RCPTID == 0 {
			rlib.LogAndPrint("Failed to get receipt for j.ID = %d,  j.JID = %d\n", j.ID, j.JID)
			return
		}
		textPrintJournalReceipt(ctx, tbl, ri, jctx, j, &rcpt)
	case rlib.JNLTYPEXFER:
		textPrintJournalXfer(tbl, ri, jctx, j)
	case rlib.JNLTYPEASMT:
		a, _ := rlib.GetAssessment(ctx, j.ID) // TODO(Steve): ignore error?
		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			rlib.LogAndPrint("Failed to get rentable for j.ID = %d,  a.RID = %d\n", j.ID, a.RID)
			return
		}
		textPrintJournalAssessment(ctx, tbl, jctx, ri.Xbiz, j, &a, &r, rentDuration, assessmentDuration)
	case rlib.JNLTYPEEXP:
		a, _ := rlib.GetExpense(ctx, j.ID) // TODO(Steve): ignore error?
		r, err := rlib.GetRentable(ctx, a.RID)
		if err != nil {
			rlib.LogAndPrint("Failed to get rentable for j.ID = %d,  a.RID = %d\n", j.ID, a.RID)
			return
		}

		printJournalExpense(ctx, tbl, ri.Xbiz, j, &a, &r)
	default:
		rlib.LogAndPrint("printJournalEntry: unrecognized type: %d\n", j.Type)
	}
}

func textReportJournalEntry(ctx context.Context, tbl *gotable.Table, ri *ReporterInfo, j *rlib.Journal, jctx *jprintctx) {
	//-------------------------------------------------------------------------------------
	// over what range of time does this rental apply between jctx.ReportStart & jctx.ReportStop?
	// the rental possession dates may be different than the report range...
	//--------------------------------------------------------------------------------------
	start := jctx.ReportStart // start with the report range
	stop := jctx.ReportStop   // start with the report range

	// TODO:  THIS NEEDS TO BE BETTER GENERALIZED...

	if len(j.JA) > 0 { // is there an associated rental agreement?
		// TODO(Steve): ignore error?
		ra, _ := rlib.GetRentalAgreement(ctx, j.JA[0].RAID) // if so, get it
		if ra.AgreementStart.After(start) {                 // if possession of rental starts later...
			start = ra.AgreementStart // ...then make an adjustment
		}
		stop = ra.AgreementStop // .Add(24 * 60 * time.Minute) -- removing this as all ranges should be NON-INCLUSIVE
		if stop.After(jctx.ReportStop) {
			stop = jctx.ReportStop
		}
	}

	//-------------------------------------------------------------------------------------------
	// this code needs to be generalized based on the recurrence period and the proration period
	//-------------------------------------------------------------------------------------------
	fullAccrualPeriod := int64(jctx.ReportStop.Sub(jctx.ReportStart).Hours() / 24)
	thisAccrualPeriod := int64(stop.Sub(start).Hours() / 24)

	// fmt.Printf("start = %s, stop = %s, fullAccrualPeriod, thisAccrualPeriod =  %d, %d\n", start.Format(rlib.RRDATEINPFMT), stop.Format(rlib.RRDATEINPFMT), fullAccrualPeriod, thisAccrualPeriod)
	textPrintJournalEntry(ctx, tbl, ri, jctx, j, thisAccrualPeriod, fullAccrualPeriod)

}

// JournalReportTable returns a Journal report in a gotable.Table for the supplied Business and time range
func JournalReportTable(ctx context.Context, ri *ReporterInfo) gotable.Table {
	const funcname = "JournalReportTable"
	var (
		err error
	)

	// init and prepare some values before table init
	ri.RptHeaderD1 = true
	ri.RptHeaderD2 = true

	// table init
	tbl := getRRTable()

	tbl.AddColumn("Journal ID", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)  // 0
	tbl.AddColumn("Description", 70, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT) // 1
	tbl.AddColumn("Date", 10, gotable.CELLDATE, gotable.COLJUSTIFYLEFT)          // 2
	tbl.AddColumn("RntAgr", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)      // 3
	tbl.AddColumn("Rentable", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)    // 4
	tbl.AddColumn("GLAccount", 8, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)    // 5
	tbl.AddColumn("Amount", 12, gotable.CELLFLOAT, gotable.COLJUSTIFYRIGHT)      // 6

	// prepare table's title, sections
	err = TableReportHeaderBlock(ctx, &tbl, "Journal", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(err.Error())
		return tbl
	}

	// get records from db
	rows, err := rlib.RRdb.Prepstmt.GetAllJournalsInRange.Query(ri.Xbiz.P.BID, &ri.D1, &ri.D2)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}
	defer rows.Close()

	jctx := jprintctx{ri.D1, ri.D2}
	// setTitle(&tbl, ri.Xbiz, &ri.D1, &ri.D2)

	for rows.Next() {
		var j rlib.Journal
		err = rlib.ReadJournals(rows, &j)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(NoRecordsFoundMsg)
			return tbl
		}

		err = rlib.GetJournalAllocations(ctx, &j)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			tbl.SetSection3(NoRecordsFoundMsg)
			return tbl
		}

		textReportJournalEntry(ctx, &tbl, ri, &j, &jctx)
	}
	err = rows.Err()
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		tbl.SetSection3(NoRecordsFoundMsg)
		return tbl
	}

	return tbl
}

// JournalReport generates a text-based report based on JournalReportTable table object
func JournalReport(ctx context.Context, ri *ReporterInfo) string {
	tbl := JournalReportTable(ctx, ri)
	return ReportToString(&tbl, ri)
}
