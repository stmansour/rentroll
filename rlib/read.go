package rlib

import "database/sql"

// As the database structures change, having the calls that Read from the database into these structures located
// in one place simplifies maintenance

// ReadAssessments reads a full Assessment structure of data from the database based on the supplied Rows pointer.
func ReadAssessments(rows *sql.Rows, a *Assessment) {
	Errcheck(rows.Scan(&a.ASMID, &a.PASMID, &a.BID, &a.RID, &a.ATypeLID, &a.RAID, &a.Amount,
		&a.Start, &a.Stop, &a.RentCycle, &a.ProrationCycle, &a.InvoiceNo, &a.AcctRule, &a.Comment,
		&a.LastModTime, &a.LastModBy))
}

// ReadGLAccounts reads a full GLAccount structure of data from the database based on the supplied Rows pointer.
func ReadGLAccounts(rows *sql.Rows, a *GLAccount) {
	Errcheck(rows.Scan(&a.LID, &a.PLID, &a.BID, &a.RAID, &a.GLNumber, &a.Status, &a.Type, &a.Name, &a.AcctType,
		&a.RAAssociated, &a.AllowPost, &a.RARequired, &a.ManageToBudget, &a.Description, &a.LastModTime, &a.LastModBy))
}

// ReadNote reads a full Note structure from the database based on the supplied row object
func ReadNote(row *sql.Row, a *Note) {
	Errcheck(row.Scan(&a.NID, &a.NLID, &a.PNID, &a.NTID, &a.RID, &a.RAID, &a.TCID, &a.Comment, &a.LastModTime, &a.LastModBy))
}

// ReadNotes reads a full Note structure from the database based on the supplied row object
func ReadNotes(rows *sql.Rows, a *Note) {
	Errcheck(rows.Scan(&a.NID, &a.NLID, &a.PNID, &a.NTID, &a.RID, &a.RAID, &a.TCID, &a.Comment, &a.LastModTime, &a.LastModBy))
}

// ReadRentalAgreement reads a full RentalAgreement structure of data from the database based on the supplied Row pointer.
func ReadRentalAgreement(row *sql.Row, a *RentalAgreement) error {
	return row.Scan(&a.RAID, &a.RATID, &a.BID, &a.NLID, &a.AgreementStart, &a.AgreementStop, &a.PossessionStart, &a.PossessionStop,
		&a.RentStart, &a.RentStop, &a.Renewal, &a.SpecialProvisions, &a.LastModTime, &a.LastModBy)
}

// ReadRentalAgreements reads a full RentalAgreement structure of data from the database based on the supplied Rows pointer.
func ReadRentalAgreements(rows *sql.Rows, a *RentalAgreement) error {
	return rows.Scan(&a.RAID, &a.RATID, &a.BID, &a.NLID, &a.AgreementStart, &a.AgreementStop, &a.PossessionStart, &a.PossessionStop,
		&a.RentStart, &a.RentStop, &a.Renewal, &a.SpecialProvisions, &a.LastModTime, &a.LastModBy)
}

// ReadReceipts reads a full Receipt structure of data from the database based on the supplied Rows pointer.
func ReadReceipts(rows *sql.Rows, a *Receipt) {
	Errcheck(rows.Scan(&a.RCPTID, &a.PRCPTID, &a.BID, &a.RAID, &a.PMTID, &a.Dt, &a.DocNo, &a.Amount, &a.AcctRule, &a.Comment, &a.OtherPayorName, &a.LastModTime, &a.LastModBy))
}

// ReadDemandSource reads a full DemandSource structure from the database based on the supplied row object
func ReadDemandSource(row *sql.Row, a *DemandSource) {
	Errcheck(row.Scan(&a.DSID, &a.BID, &a.Name, &a.Industry, &a.LastModTime, &a.LastModBy))
}

// ReadDemandSources reads a full DemandSource structure from the database based on the supplied rows object
func ReadDemandSources(rows *sql.Rows, a *DemandSource) {
	Errcheck(rows.Scan(&a.DSID, &a.BID, &a.Name, &a.Industry, &a.LastModTime, &a.LastModBy))
}

// ReadStringList reads a full StringList structure from the database based on the supplied row object
func ReadStringList(row *sql.Row, a *StringList) {
	Errcheck(row.Scan(&a.SLID, &a.BID, &a.Name, &a.LastModTime, &a.LastModBy))
}

// ReadStringLists reads a full StringList structure from the database based on the supplied row object
func ReadStringLists(rows *sql.Rows, a *StringList) {
	Errcheck(rows.Scan(&a.SLID, &a.BID, &a.Name, &a.LastModTime, &a.LastModBy))
}
