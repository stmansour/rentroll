package rlib

// UpdateLedgerMarker updates a ledger marker record
func UpdateLedgerMarker(lm *LedgerMarker) error {
	_, err := RRdb.Prepstmt.UpdateLedgerMarker.Exec(lm.LMID, lm.LID, lm.BID, lm.DtStart, lm.DtStop, lm.Balance, lm.State, lm.LastModBy, lm.LMID)
	if nil != err {
		Ulog("UpdateLedgerMarker: error inserting LedgerMarker:  %v\n", err)
		Ulog("LedgerMarker = %#v\n", *lm)
	}
	return err
}

// UpdateLedger updates a ledger marker record
func UpdateLedger(l *Ledger) error {
	_, err := RRdb.Prepstmt.UpdateLedger.Exec(l.BID, l.RAID, l.GLNumber, l.Status, l.Type, l.Name, l.AcctType, l.RAAssociated, l.LastModBy, l.LID)
	if nil != err {
		Ulog("UpdateLedger: error inserting Ledger:  %v\n", err)
		Ulog("Ledger = %#v\n", *l)
	}
	return err
}

// UpdateTransactant updates a transactant record in the database
func UpdateTransactant(a *Transactant) error {
	_, err := RRdb.Prepstmt.UpdateTransactant.Exec(a.TID, a.PID, a.PRSPID, a.FirstName, a.MiddleName, a.LastName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateTransactant: error inserting Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return err
}
