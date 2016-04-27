package rlib

// UpdateTransactant updates a new transactant record to the database
func UpdateTransactant(a *Transactant) error {
	_, err := RRdb.Prepstmt.UpdateTransactant.Exec(a.TID, a.PID, a.PRSPID, a.FirstName, a.MiddleName, a.LastName, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateTransactant: error inserting Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return err
}
