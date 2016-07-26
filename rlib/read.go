package rlib

import "database/sql"

// As the database structures change, having the calls that Read from the database into these structures located
// in one place simplifies maintenance

// ReadAssessment reads a full Assessment structure of data from the database based on the supplied Rows pointer.
func ReadAssessment(rows *sql.Rows, a *Assessment) {
	Errcheck(rows.Scan(&a.ASMID, &a.PASMID, &a.BID, &a.RID, &a.ATypeLID, &a.RAID, &a.Amount,
		&a.Start, &a.Stop, &a.RentCycle, &a.ProrationCycle, &a.InvoiceNo, &a.AcctRule, &a.Comment,
		&a.LastModTime, &a.LastModBy))
}

// ReadGLAccount reads a full GLAccount structure of data from the database based on the supplied Rows pointer.
func ReadGLAccount(rows *sql.Rows, a *GLAccount) {
	Errcheck(rows.Scan(&a.LID, &a.PLID, &a.BID, &a.RAID, &a.GLNumber, &a.Status, &a.Type, &a.Name, &a.AcctType,
		&a.RAAssociated, &a.AllowPost, &a.RARequired, &a.ManageToBudget, &a.Description, &a.LastModTime, &a.LastModBy))
}

// ReadRentalAgreement reads a full RentalAgreement structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreement(rows *sql.Rows, a *RentalAgreement) {
	Errcheck(rows.Scan(&a.RAID, &a.RATID, &a.BID, &a.NLID, &a.RentalStart, &a.RentalStop, &a.PossessionStart, &a.PossessionStop,
		&a.Renewal, &a.SpecialProvisions, &a.LastModTime, &a.LastModBy))
}

// ReadReceipt reads a full Receipt structure of data from the database based on the supplied Rows pointer.
func ReadReceipt(rows *sql.Rows, a *Receipt) {
	Errcheck(rows.Scan(
		&a.RCPTID, &a.PRCPTID, &a.BID, &a.RAID, &a.PMTID, &a.Dt, &a.DocNo, &a.Amount, &a.AcctRule, &a.Comment, &a.OtherPayorName, &a.LastModTime, &a.LastModBy))
}

// ReadSource reads a full Source structure from the database ased on the supplied row
func ReadSource(row *sql.Row, a *Source) {
	Errcheck(row.Scan(&a.SID, &a.Name, &a.Industry, &a.LastModTime, &a.LastModBy))
}
