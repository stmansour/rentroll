package rrpt

import (
	"fmt"
	"gotable"
	"rentroll/rlib"
	"strconv"
)

// RRTextReport prints a text-based RentRoll report
// for the business in xbiz and timeframe d1 to d2 to stdout
func RRTextReport(ri *ReporterInfo) {
	fmt.Print(RRReport(ri))
}

// RRReport returns a string containin a text-based RentRoll report
// for the business in xbiz and timeframe d1 to d2.
func RRReport(ri *ReporterInfo) string {
	tbl := RRReportTable(ri)
	return ReportToString(&tbl, ri)
}

// RRReportTable returns the gotable representation for rentroll report
func RRReportTable(ri *ReporterInfo) gotable.Table {
	const funcname = "RRReportTable"
	var (
		err error
		tbl = getRRTable() // gotable init for this report
	)
	rlib.Console("Entered in %s", funcname)

	// use section3 for errors and apply red color
	cssListSection3 := []*gotable.CSSProperty{
		{Name: "color", Value: "red"},
		{Name: "font-family", Value: "monospace"},
	}
	tbl.SetSection3CSS(cssListSection3)

	// set table title, sections
	err = TableReportHeaderBlock(&tbl, "Rentroll", funcname, ri)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return tbl
	}

	// Add columns to the table
	tbl.AddColumn("Rentable", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                    // column for the Rentable name
	tbl.AddColumn("Rentable Type", 15, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)               // RentableType name
	tbl.AddColumn("SqFt", 5, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                        // the Custom Attribute "Square Feet"
	tbl.AddColumn("Description", 20, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                 // the Custom Attribute "Square Feet"
	tbl.AddColumn("Users", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                       // Users of this rentable
	tbl.AddColumn("Payors", 30, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                      // Users of this rentable
	tbl.AddColumn("Rental Agreement", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)            // the Rental Agreement id
	tbl.AddColumn("Use Period", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                  // the use period
	tbl.AddColumn("Rent Period", 10, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                 // the rent period
	tbl.AddColumn("Rent Cycle", 12, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)                  // the rent cycle
	tbl.AddColumn("GSR", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                        // gross scheduled rent
	tbl.AddColumn("Period GSR", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                 // gross scheduled rent
	tbl.AddColumn("Income Offsets", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)             // GL Account
	tbl.AddColumn("Amount Due", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)                 // Amount due
	tbl.AddColumn("Payments Applied", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)           // contract rent amounts
	tbl.AddColumn("Beginning Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Change In Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)       // account for the associated RentalAgreement
	tbl.AddColumn("Ending Receivable", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)          // account for the associated RentalAgreement
	tbl.AddColumn("Beginning Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Change In Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT) // account for the associated RentalAgreement
	tbl.AddColumn("Ending Security Deposit", 10, gotable.CELLSTRING, gotable.COLJUSTIFYRIGHT)    // account for the associated RentalAgreement

	// NOW GET THE ROWS FOR RENTROLL ROUTINE
	rows, _, _, err := rlib.GetRentRollRows(
		ri.Bid, ri.D1, ri.D2, // BID, startDate, stopDate
		-1, -1, // offset, limit
	)

	// if any error encountered then just set it to section3
	if err != nil {
		tbl.SetSection3(err.Error())
		return tbl
	}

	for index, row := range rows {
		if (row.FLAGS & rlib.RentRollSubTotalRow) > 0 { // add line before subtotal Row
			// tbl.AddLineBefore(index) // AddLineBefore is not working
			tbl.AddLineAfter(index - 1)
		}
		rrTableAddRow(&tbl, row)
	}
	tbl.AddLineAfter(len(tbl.Row) - 2) // Grand Total line, Rows index start from zero

	return tbl
}

// rrTableAddRow adds row in gotable struct with information
// given by RentRollViewRow struct
func rrTableAddRow(tbl *gotable.Table, q rlib.RentRollStaticInfo) {

	// column numbers for gotable report
	const (
		RName       = 0
		RType       = iota
		SqFt        = iota
		Descr       = iota
		Users       = iota
		Payors      = iota
		RAgr        = iota
		UsePeriod   = iota
		RentPeriod  = iota
		RentCycle   = iota
		GSR         = iota
		PeriodGSR   = iota
		IncOff      = iota
		AmtDue      = iota
		PmtRcvd     = iota
		BeginRcv    = iota
		ChgRcv      = iota
		EndRcv      = iota
		BeginSecDep = iota
		ChgSecDep   = iota
		EndSecDep   = iota
	)

	// Add new row in gotable to show the data
	tbl.AddRow()

	// --------------------------------------
	// now, start to put values in the column
	// --------------------------------------

	// 1. RentableName
	if q.RentableName.Valid {
		tbl.Puts(-1, RName, q.RentableName.String)
	} else {
		tbl.Puts(-1, RName, "")
	}

	// 2. RentableType
	if q.RentableType.Valid {
		tbl.Puts(-1, RType, q.RentableType.String)
	} else {
		tbl.Puts(-1, RType, "")
	}

	// 3. SqFt
	if q.SqFt.Valid {
		tbl.Puts(-1, SqFt, int64ToStr(q.SqFt.Int64, true))
	} else {
		tbl.Puts(-1, SqFt, "")
	}

	// 4. Description
	if q.Description.Valid {
		tbl.Puts(-1, Descr, q.Description.String)
	} else {
		tbl.Puts(-1, Descr, "")
	}

	// 5. Users
	if q.Users.Valid {
		tbl.Puts(-1, Users, q.Users.String)
	} else {
		tbl.Puts(-1, Users, "")
	}

	// 6. Payors
	if q.Payors.Valid {
		tbl.Puts(-1, Payors, q.Payors.String)
	} else {
		tbl.Puts(-1, Payors, "")
	}

	// 7. Rental Agreement
	tbl.Puts(-1, RAgr, q.RAIDREP)

	// 8. Use Period
	useTimeStr := ""
	if q.PossessionStart.Valid && q.PossessionStop.Valid {
		useTimeStr = q.PossessionStart.Time.Format(rlib.RRDATEFMT3) +
			"-" + q.PossessionStop.Time.Format(rlib.RRDATEFMT3)
	}
	tbl.Puts(-1, UsePeriod, useTimeStr)

	// 9. Rent Period
	rentTimeStr := ""
	if q.RentStart.Valid && q.RentStop.Valid {
		rentTimeStr = q.RentStart.Time.Format(rlib.RRDATEFMT3) +
			"-" + q.RentStop.Time.Format(rlib.RRDATEFMT3)
	}
	tbl.Puts(-1, RentPeriod, rentTimeStr)

	// 10. Rent Cycle
	tbl.Puts(-1, RentCycle, q.RentCycleREP) // Rent Cycle

	// --------------------------------------
	// Rest of columns are amount columns,
	// those need be shown based on row type
	// --------------------------------------
	var GSRREP, PeriodGSRREP,
		IncomeOffsetsREP, AmountDueREP, PaymentsAppliedREP,
		BeginReceivableREP, DeltaReceivableREP, EndReceivableREP,
		BeginSecDepREP, DeltaSecDepREP, EndSecDepREP string

	/*
	   For Blank Row, we don't have to show data in all columns
	   so we're only doing data representation for non-blank row
	*/

	// if it is NOT blank row then
	if !((q.FLAGS & rlib.RentRollBlankRow) > 0) {

		// GSR, by default it is blank string, no need else clause
		if q.RentCycleGSR.Valid {
			GSRREP = float64ToStr(q.RentCycleGSR.Float64, false)
		}

		// Period GSR
		if q.PeriodGSR.Valid {
			PeriodGSRREP = float64ToStr(q.PeriodGSR.Float64, false)
		}

		// Income Offsets
		if q.IncomeOffsets.Valid {
			IncomeOffsetsREP = float64ToStr(q.IncomeOffsets.Float64, false)
		}

		// Amount Due
		if q.AmountDue.Valid {
			AmountDueREP = float64ToStr(q.AmountDue.Float64, false)
		}

		// Payments Applied
		if q.PaymentsApplied.Valid {
			PaymentsAppliedREP = float64ToStr(q.PaymentsApplied.Float64, false)
		}

		// ------------- LAST SIX COLUMS ----------------

		// get the values in last six columns for
		// subtotal as well as grand total row
		if (q.FLAGS&rlib.RentRollSubTotalRow) > 0 || (q.FLAGS&rlib.RentRollGrandTotalRow) > 0 {
			BeginReceivableREP = float64ToStr(q.BeginReceivable, false)
			DeltaReceivableREP = float64ToStr(q.DeltaReceivable, false)
			EndReceivableREP = float64ToStr(q.EndReceivable, false)
			BeginSecDepREP = float64ToStr(q.BeginSecDep, false)
			DeltaSecDepREP = float64ToStr(q.DeltaSecDep, false)
			EndSecDepREP = float64ToStr(q.EndSecDep, false)
		} else {

			// for normal row, last six columns should have be blank
			// also, those need to be greyish

			// set the background color for these cells
			rowIndex := tbl.RowCount() - 1

			greyCellCSS := []*gotable.CSSProperty{
				{Name: "background-color", Value: "#CCC"},
			}

			tbl.SetCellCSS(rowIndex, BeginRcv, greyCellCSS)
			tbl.SetCellCSS(rowIndex, ChgRcv, greyCellCSS)
			tbl.SetCellCSS(rowIndex, EndRcv, greyCellCSS)
			tbl.SetCellCSS(rowIndex, BeginSecDep, greyCellCSS)
			tbl.SetCellCSS(rowIndex, ChgSecDep, greyCellCSS)
			tbl.SetCellCSS(rowIndex, EndSecDep, greyCellCSS)
		}
	}

	// FEED all last formatted amount columns finallay in the row
	tbl.Puts(-1, GSR, GSRREP)
	tbl.Puts(-1, PeriodGSR, PeriodGSRREP)
	tbl.Puts(-1, IncOff, IncomeOffsetsREP)
	tbl.Puts(-1, AmtDue, AmountDueREP)
	tbl.Puts(-1, PmtRcvd, PaymentsAppliedREP)
	tbl.Puts(-1, BeginRcv, BeginReceivableREP)
	tbl.Puts(-1, ChgRcv, DeltaReceivableREP)
	tbl.Puts(-1, EndRcv, EndReceivableREP)
	tbl.Puts(-1, BeginSecDep, BeginSecDepREP)
	tbl.Puts(-1, ChgSecDep, DeltaSecDepREP)
	tbl.Puts(-1, EndSecDep, EndSecDepREP)
}

// int64ToStr returns the string represenation of int64 type number
// if blank is set to true, then it will returns blank string otherwise returns 0
func int64ToStr(number int64, blank bool) string {
	nStr := strconv.FormatInt(number, 10)
	if nStr == "0" {
		if blank {
			return ""
		}
	}
	return nStr
}

// float64ToStr returns the string represenation of float64 type number
// if blank is set to true, then it will returns blank string otherwise returns 0.00
func float64ToStr(number float64, blank bool) string {
	nStr := strconv.FormatFloat(number, 'f', 2, 64)
	if nStr == "0.00" {
		if blank {
			return ""
		}
	}
	return nStr
}
