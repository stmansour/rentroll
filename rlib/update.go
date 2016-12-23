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

// UpdateBusiness updates an Business record
func UpdateBusiness(a *Business) error {
	_, err := RRdb.Prepstmt.UpdateBusiness.Exec(a.Designation, a.Name, a.DefaultRentCycle, a.DefaultProrationCycle, a.DefaultGSRPC, a.LastModBy, a.BID)
	if nil != err {
		Ulog("UpdateBusiness: error updating Business:  %v\n", err)
		Ulog("Business = %#v\n", *a)
	}
	return err
}

// UpdateDemandSource updates a DemandSource record in the database
func UpdateDemandSource(a *DemandSource) error {
	_, err := RRdb.Prepstmt.UpdateDemandSource.Exec(a.Name, a.Industry, a.LastModBy, a.SourceSLSID)
	if nil != err {
		Ulog("UpdateDemandSource: error updating DemandSource:  %v\n", err)
		Ulog("DemandSource = %#v\n", *a)
	}
	return err
}

// UpdateDeposit updates a Deposit record
func UpdateDeposit(a *Deposit) error {
	_, err := RRdb.Prepstmt.UpdateDeposit.Exec(a.BID, a.DEPID, a.DPMID, a.Dt, a.Amount, a.LastModBy, a.DID)
	if nil != err {
		Ulog("UpdateDeposit: error updating Deposit:  %v\n", err)
		Ulog("Deposit = %#v\n", *a)
	}
	return err
}

// UpdateDepositMethod updates a DepositMethod record
func UpdateDepositMethod(a *DepositMethod) error {
	_, err := RRdb.Prepstmt.UpdateDepositMethod.Exec(a.BID, a.Name, a.DPMID)
	if nil != err {
		Ulog("UpdateDepositMethod: error updating DepositMethod:  %v\n", err)
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

// UpdateInvoice updates a Invoice record
func UpdateInvoice(a *Invoice) error {
	_, err := RRdb.Prepstmt.UpdateInvoice.Exec(a.BID, a.Dt, a.DtDue, a.Amount, a.DeliveredBy, a.LastModBy, a.InvoiceNo)
	if nil != err {
		Ulog("UpdateInvoice: error updating Invoice:  %v\n", err)
		Ulog("Deposit = %#v\n", *a)
	}
	return err
}

// UpdateLedgerMarker updates a LedgerMarker record
func UpdateLedgerMarker(lm *LedgerMarker) error {
	// if lm.RAID > 0 {
	// 	fmt.Printf("UpdateLedgerMarker: lm.RAID = %d\n", lm.RAID)
	// 	debug.PrintStack()
	// }
	_, err := RRdb.Prepstmt.UpdateLedgerMarker.Exec(lm.LID, lm.BID, lm.RAID, lm.RID, lm.Dt, lm.Balance, lm.State, lm.LastModBy, lm.LMID)
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

// UpdatePayor updates a Payor record in the database
func UpdatePayor(a *Payor) error {
	_, err := RRdb.Prepstmt.UpdatePayor.Exec(a.CreditLimit, a.TaxpayorID, a.AccountRep, a.EligibleFuturePayor, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdatePayor: error updating pet:  %v\n", err)
		Ulog("Payor = %#v\n", *a)
	}
	return err
}

// UpdateProspect updates a Prospect record in the database
func UpdateProspect(a *Prospect) error {
	_, err := RRdb.Prepstmt.UpdateProspect.Exec(a.EmployerName, a.EmployerStreetAddress, a.EmployerCity,
		a.EmployerState, a.EmployerPostalCode, a.EmployerEmail, a.EmployerPhone, a.Occupation, a.ApplicationFee,
		a.DesiredUsageStartDate, a.RentableTypePreference, a.FLAGS, a.Approver, a.DeclineReasonSLSID, a.OtherPreferences,
		a.FollowUpDate, a.CSAgent, a.OutcomeSLSID, a.FloatingDeposit, a.RAID, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateProspect: error updating pet:  %v\n", err)
		Ulog("Prospect = %#v\n", *a)
	}
	return err
}

// UpdateRatePlan updates a RatePlan record in the database
func UpdateRatePlan(a *RatePlan) error {
	_, err := RRdb.Prepstmt.UpdateRatePlan.Exec(a.BID, a.Name, a.LastModBy, a.RPID)
	if nil != err {
		Ulog("UpdateRatePlan: error:  %v\n", err)
		Ulog("RatePlan = %#v\n", *a)
	}
	return err
}

// UpdateRatePlanRef updates a RatePlanRef record in the database
func UpdateRatePlanRef(a *RatePlanRef) error {
	_, err := RRdb.Prepstmt.UpdateRatePlanRef.Exec(a.RPID, a.DtStart, a.DtStop, a.FeeAppliesAge, a.MaxNoFeeUsers, a.AdditionalUserFee, a.PromoCode, a.CancellationFee, a.FLAGS, a.LastModBy, a.RPRID)
	if nil != err {
		Ulog("UpdateRatePlanRef: error:  %v\n", err)
		Ulog("RatePlanRef = %#v\n", *a)
	}
	return err
}

// UpdateRatePlanRefRTRate updates a RatePlanRefRTRate record in the database
func UpdateRatePlanRefRTRate(a *RatePlanRefRTRate) error {
	_, err := RRdb.Prepstmt.UpdateRatePlanRefRTRate.Exec(a.FLAGS, a.Val, a.RPRID, a.RTID)
	if nil != err {
		Ulog("UpdateRatePlanRefRTRate: error:  %v\n", err)
		Ulog("RatePlanRefRTRate = %#v\n", *a)
	}
	return err
}

// UpdateRatePlanRefSPRate updates a RatePlanRefSPRate record in the database
func UpdateRatePlanRefSPRate(a *RatePlanRefSPRate) error {
	_, err := RRdb.Prepstmt.UpdateRatePlanRefSPRate.Exec(a.FLAGS, a.Val, a.RPRID, a.RTID, a.RSPID)
	if nil != err {
		Ulog("UpdateRatePlanRefSPRate: error:  %v\n", err)
		Ulog("RatePlanRefSPRate = %#v\n", *a)
	}
	return err
}

// UpdateRentalAgreement updates a RentalAgreement record in the database
func UpdateRentalAgreement(a *RentalAgreement) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreement.Exec(a.RATID, a.BID, a.NLID, a.AgreementStart, a.AgreementStop, a.PossessionStart, a.PossessionStop, a.RentStart, a.RentStop, a.RentCycleEpoch, a.UnspecifiedAdults, a.UnspecifiedChildren, a.Renewal, a.SpecialProvisions, a.LeaseType, a.ExpenseAdjustmentType, a.ExpensesStop, a.ExpenseStopCalculation, a.BaseYearEnd, a.ExpenseAdjustment, a.EstimatedCharges, a.RateChange, a.NextRateChange, a.PermittedUses, a.ExclusiveUses, a.ExtensionOption, a.ExtensionOptionNotice, a.ExpansionOption, a.ExpansionOptionNotice, a.RightOfFirstRefusal, a.LastModBy, a.RAID)
	if nil != err {
		Ulog("UpdateRentalAgreement: error updating :  %v\n", err)
		Ulog("RentalAgreement = %#v\n", *a)
	}
	return err
}

// UpdateRentalAgreementPet updates a RentalAgreementPet record in the database
func UpdateRentalAgreementPet(a *RentalAgreementPet) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreementPet.Exec(a.RAID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.LastModBy, a.PETID)
	if nil != err {
		Ulog("UpdateRentalAgreementPet: error updating pet:  %v\n", err)
		Ulog("RentalAgreementPet = %#v\n", *a)
	}
	return err
}

// UpdateRentableSpecialtyRef updates a RentableSpecialtyRef record in the database
func UpdateRentableSpecialtyRef(a *RentableSpecialtyRef) error {
	_, err := RRdb.Prepstmt.UpdateRentableSpecialtyRef.Exec(a.RSPID, a.LastModBy, a.RID, a.DtStart, a.DtStop)
	if nil != err {
		Ulog("UpdateRentableSpecialtyRef: error updating RentableSpecialtyRef:  %v\n", err)
		Ulog("RentableSpecialtyRef = %#v\n", *a)
	}
	return err
}

// UpdateRentableTypeRef updates a RentableTypeRef record in the database
func UpdateRentableTypeRef(a *RentableTypeRef) error {
	_, err := RRdb.Prepstmt.UpdateRentableTypeRef.Exec(a.RTID, a.OverrideRentCycle, a.OverrideProrationCycle, a.LastModBy, a.RID, a.DtStart, a.DtStop)
	if nil != err {
		Ulog("UpdateRentableTypeRef: error updating RentableTypeRef:  %v\n", err)
		Ulog("RentableTypeRef = %#v\n", *a)
	}
	return err
}

// UpdateStringList updates a StringList record in the database. It also updates the string list. It does this by
// deleting all the strings first, then inserting the ones it has.
func UpdateStringList(a *StringList) error {
	_, err := RRdb.Prepstmt.UpdateStringList.Exec(a.BID, a.Name, a.LastModBy, a.SLID)
	if nil != err {
		Ulog("UpdateStringList: error:  %v\n", err)
		Ulog("StringList = %#v\n", *a)
	}
	DeleteSLStrings(a.SLID)
	InsertSLStrings(a)
	return err
}

// UpdateSLString updates a SLString record in the database
func UpdateSLString(a *SLString) error {
	_, err := RRdb.Prepstmt.UpdateSLString.Exec(a.SLID, a.Value, a.LastModBy, a.SLSID)
	if nil != err {
		Ulog("UpdateSLString: error:  %v\n", err)
		Ulog("SLString = %#v\n", *a)
	}
	return err
}

// UpdateTransactant updates a Transactant record in the database
func UpdateTransactant(a *Transactant) error {
	_, err := RRdb.Prepstmt.UpdateTransactant.Exec(a.BID, a.NLID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateTransactant: error updating Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return err
}

// UpdateUser updates a User record in the database
func UpdateUser(a *User) error {
	_, err := RRdb.Prepstmt.UpdateUser.Exec(a.Points, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress, a.EmergencyContactTelephone, a.EmergencyEmail, a.AlternateAddress, a.EligibleFutureUser, a.Industry, a.SourceSLSID, a.LastModBy, a.TCID)
	if nil != err {
		Ulog("UpdateUser: error updating User:  %v\n", err)
		Ulog("User = %#v\n", *a)
	}
	return err
}

// UpdateVehicle updates a Vehicle record in the database
func UpdateVehicle(a *Vehicle) error {
	_, err := RRdb.Prepstmt.UpdateVehicle.Exec(a.TCID, a.BID, a.VehicleType, a.VehicleMake, a.VehicleModel, a.VehicleColor, a.VehicleYear, a.LicensePlateState, a.LicensePlateNumber, a.ParkingPermitNumber, a.DtStart, a.DtStop, a.LastModBy, a.VID)
	if nil != err {
		Ulog("UpdateVehicle: error updating Vehicle:  %v\n", err)
		Ulog("Vehicle = %#v\n", *a)
	}
	return err
}
