package rlib

import "database/sql"

// As the database structures change, having the calls that Read from the database into these structures located
// in one place simplifies maintenance

// ReadAssessment reads a full Assessment structure of data from the database based on the supplied Rows pointer.
func ReadAssessment(rows *sql.Rows, a *Assessment) {
	Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.ASMTID, &a.RAID, &a.Amount,
		&a.Start, &a.Stop, &a.RecurCycle, &a.ProrationCycle, &a.AcctRule, &a.Comment,
		&a.LastModTime, &a.LastModBy))
}

// ReadRentalAgreement reads a full RentalAgreement structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreement(rows *sql.Rows, a *RentalAgreement) {
	Errcheck(rows.Scan(&a.RAID, &a.RATID, &a.BID, &a.RentalStart, &a.RentalStop, &a.PossessionStart, &a.PossessionStop,
		&a.Renewal, &a.SpecialProvisions, &a.LastModTime, &a.LastModBy))
}

// ReadReceipt reads a full Receipt structure of data from the database based on the supplied Rows pointer.
func ReadReceipt(rows *sql.Rows, a *Receipt) {
	Errcheck(rows.Scan(
		&a.RCPTID, &a.BID, &a.RAID, &a.PMTID, &a.Dt, &a.Amount, &a.AcctRule, &a.Comment, &a.LastModTime, &a.LastModBy))
}
