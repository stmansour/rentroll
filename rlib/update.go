package rlib

// UpdateAssessment updates an Assessment record
func UpdateAssessment(a *Assessment) error {
	_, err := RRdb.Prepstmt.UpdateAssessment.Exec(a.PASMID, a.BID, a.RID, a.ATypeLID, a.RAID, a.Amount, a.Start, a.Stop, a.RentCycle, a.ProrationCycle, a.InvoiceNo, a.AcctRule, a.Comment, a.LastModBy, a.ASMID)
	if nil != err {
		Ulog("UpdateAssessment: error updating Assessment:  %v\n", err)
		Ulog("Assessment = %#v\n", *a)
	}
	return err
}

// UpdateDeposit updates a Deposit record
func UpdateDeposit(a *Deposit) error {
	_, err := RRdb.Prepstmt.UpdateDeposit.Exec(a.BID, a.DEPID, a.Dt, a.Amount, a.LastModBy, a.DID)
	if nil != err {
		Ulog("UpdateDeposit: error updating Deposit:  %v\n", err)
		Ulog("Deposit = %#v\n", *a)
	}
	return err
}

// UpdateDepository updates a Depository record
func UpdateDepository(a *Depository) error {
	_, err := RRdb.Prepstmt.UpdateDepository.Exec(a.BID, a.Name, a.AccountNo, a.LastModBy, a.DEPID)
	if nil != err {
		Ulog("UpdateDepository: error updating Depository:  %v\n", err)
		Ulog("Depository = %#v\n", *a)
	}
	return err
}

// UpdateLedgerMarker updates a LedgerMarker record
func UpdateLedgerMarker(lm *LedgerMarker) error {
	_, err := RRdb.Prepstmt.UpdateLedgerMarker.Exec(lm.LMID, lm.LID, lm.BID, lm.DtStart, lm.DtStop, lm.Balance, lm.State, lm.LastModBy, lm.LMID)
	if nil != err {
		Ulog("UpdateLedgerMarker: error updating LedgerMarker:  %v\n", err)
		Ulog("LedgerMarker = %#v\n", *lm)
	}
	return err
}

// UpdateLedger updates a Ledger record
func UpdateLedger(l *GLAccount) error {
	_, err := RRdb.Prepstmt.UpdateLedger.Exec(l.PLID, l.BID, l.RAID, l.GLNumber, l.Status, l.Type, l.Name, l.AcctType, l.RAAssociated, l.AllowPost, l.RARequired, l.ManageToBudget, l.Description, l.LastModBy, l.LID)
	if nil != err {
		Ulog("UpdateLedger: error updating GLAccount:  %v\n", err)
		Ulog("GLAccount = %#v\n", *l)
	}
	return err
}

// UpdateTransactant updates a Transactant record in the database
func UpdateTransactant(a *Transactant) error {
	_, err := RRdb.Prepstmt.UpdateTransactant.Exec(a.USERID, a.PID, a.PRSPID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.Notes, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateTransactant: error updating Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return err
}

// UpdateRentalAgreementPet updates a Transactant record in the database
func UpdateRentalAgreementPet(a *RentalAgreementPet) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreementPet.Exec(a.RAID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.LastModBy, a.PETID)
	if nil != err {
		Ulog("UpdateRentalAgreementPet: error updating pet:  %v\n", err)
		Ulog("RentalAgreementPet = %#v\n", *a)
	}
	return err
}

// UpdateRentableTypeRef updates a Transactant record in the database
func UpdateRentableTypeRef(a *RentableTypeRef) error {
	_, err := RRdb.Prepstmt.UpdateRentableTypeRef.Exec(a.RTID, a.RentCycle, a.ProrationCycle, a.LastModBy, a.RID, a.DtStart, a.DtStop)
	if nil != err {
		Ulog("UpdateRentableTypeRef: error updating pet:  %v\n", err)
		Ulog("RentableTypeRef = %#v\n", *a)
	}
	return err
}

// UpdateRentableSpecialtyRef updates a Transactant record in the database
func UpdateRentableSpecialtyRef(a *RentableSpecialtyRef) error {
	_, err := RRdb.Prepstmt.UpdateRentableSpecialtyRef.Exec(a.RSPID, a.LastModBy, a.RID, a.DtStart, a.DtStop)
	if nil != err {
		Ulog("UpdateRentableSpecialtyRef: error updating pet:  %v\n", err)
		Ulog("RentableSpecialtyRef = %#v\n", *a)
	}
	return err
}
