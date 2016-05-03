package rlib

// UpdateLedgerMarker updates a ledger marker record
func UpdateLedgerMarker(lm *LedgerMarker) error {
	_, err := RRdb.Prepstmt.UpdateLedgerMarker.Exec(lm.LMID, lm.BID, lm.PID, lm.GLNumber, lm.Status, lm.State, lm.DtStart, lm.DtStop, lm.Balance, lm.Type, lm.Name, lm.AcctType, lm.RAAssociated, lm.LastModBy, lm.LMID)
	if nil != err {
		Ulog("UpdateLedgerMarker: error inserting LedgerMarker:  %v\n", err)
		Ulog("LedgerMarker = %#v\n", *lm)
	}
	return err
}

// UpdateTransactant updates a transactant record in the database
func UpdateTransactant(a *Transactant) error {
	_, err := RRdb.Prepstmt.UpdateTransactant.Exec(a.TID, a.PID, a.PRSPID, a.FirstName, a.MiddleName, a.LastName, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateTransactant: error inserting Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return err
}
