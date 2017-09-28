package rlib

func updateError(err error, n string, a interface{}) error {
	if nil != err {
		Ulog("Update%s: error updating %s:  %v\n", n, n, err)
		Ulog("%s = %#v\n", n, a)
	}
	return err
}

// UpdateAR updates an AR record
func UpdateAR(a *AR) error {
	_, err := RRdb.Prepstmt.UpdateAR.Exec(a.BID, a.Name, a.ARType, a.DebitLID, a.CreditLID, a.Description, a.RARequired, a.DtStart, a.DtStop, a.FLAGS, a.DefaultAmount, a.LastModBy, a.ARID)
	return updateError(err, "AR", *a)
}

// UpdateAssessment updates an Assessment record
func UpdateAssessment(a *Assessment) error {
	// debug.PrintStack()
	_, err := RRdb.Prepstmt.UpdateAssessment.Exec(a.PASMID, a.RPASMID, a.BID, a.RID, a.ATypeLID, a.RAID, a.Amount, a.Start, a.Stop, a.RentCycle, a.ProrationCycle, a.InvoiceNo, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.LastModBy, a.ASMID)
	return updateError(err, "Assessment", *a)
}

// UpdateBusiness updates an Business record
func UpdateBusiness(a *Business) error {
	_, err := RRdb.Prepstmt.UpdateBusiness.Exec(a.Designation, a.Name, a.DefaultRentCycle, a.DefaultProrationCycle, a.DefaultGSRPC, a.LastModBy, a.BID)
	return updateError(err, "Business", *a)
}

// UpdateCustomAttribute updates an CustomAttribute record
func UpdateCustomAttribute(a *CustomAttribute) error {
	_, err := RRdb.Prepstmt.UpdateCustomAttribute.Exec(a.BID, a.Type, a.Name, a.Value, a.Units, a.LastModBy, a.CID)
	return updateError(err, "CustomAttribute", *a)
}

// UpdateDemandSource updates a DemandSource record in the database
func UpdateDemandSource(a *DemandSource) error {
	_, err := RRdb.Prepstmt.UpdateDemandSource.Exec(a.Name, a.Industry, a.LastModBy, a.SourceSLSID)
	return updateError(err, "DemandSource", *a)
}

// UpdateDeposit updates a Deposit record
func UpdateDeposit(a *Deposit) error {
	_, err := RRdb.Prepstmt.UpdateDeposit.Exec(a.BID, a.DEPID, a.DPMID, a.Dt, a.Amount, a.ClearedAmount, a.FLAGS, a.LastModBy, a.DID)
	return updateError(err, "Deposit", *a)
}

// UpdateDepository updates a Depository record
func UpdateDepository(a *Depository) error {
	_, err := RRdb.Prepstmt.UpdateDepository.Exec(a.BID, a.LID, a.Name, a.AccountNo, a.LastModBy, a.DEPID)
	return updateError(err, "Depository", *a)
}

// UpdateDepositMethod updates a DepositMethod record
func UpdateDepositMethod(a *DepositMethod) error {
	_, err := RRdb.Prepstmt.UpdateDepositMethod.Exec(a.BID, a.Method, a.LastModBy, a.DPMID)
	return updateError(err, "DepositMethod", *a)
}

// UpdateDepositPart updates a DepositPart record
func UpdateDepositPart(a *DepositPart) error {
	_, err := RRdb.Prepstmt.UpdateDepositPart.Exec(a.DID, a.BID, a.RCPTID, a.LastModBy, a.DPID)
	return updateError(err, "DepositPart", *a)
}

// UpdateExpense updates a Expense record
func UpdateExpense(a *Expense) error {
	_, err := RRdb.Prepstmt.UpdateExpense.Exec(a.RPEXPID, a.BID, a.RID, a.RAID, a.Amount, a.Dt, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.LastModBy, a.EXPID)
	return updateError(err, "Expense", *a)
}

// UpdateInvoice updates a Invoice record
func UpdateInvoice(a *Invoice) error {
	_, err := RRdb.Prepstmt.UpdateInvoice.Exec(a.BID, a.Dt, a.DtDue, a.Amount, a.DeliveredBy, a.LastModBy, a.InvoiceNo)
	return updateError(err, "Invoice", *a)
}

// UpdateLedgerMarker updates a LedgerMarker record
func UpdateLedgerMarker(a *LedgerMarker) error {
	_, err := RRdb.Prepstmt.UpdateLedgerMarker.Exec(a.LID, a.BID, a.RAID, a.RID, a.TCID, a.Dt, a.Balance, a.State, a.LastModBy, a.LMID)
	return updateError(err, "LedgerMarker", *a)
}

// UpdateLedger updates a Ledger record
func UpdateLedger(a *GLAccount) error {
	_, err := RRdb.Prepstmt.UpdateLedger.Exec(a.PLID, a.BID, a.RAID, a.TCID, a.GLNumber, a.Status, a.Name, a.AcctType, a.AllowPost, a.FLAGS, a.Description, a.LastModBy, a.LID)
	return updateError(err, "GLAccount", *a)
}

// UpdateJournalAllocation updates a JournalAllocation record
func UpdateJournalAllocation(a *JournalAllocation) error {
	_, err := RRdb.Prepstmt.UpdateJournalAllocation.Exec(a.BID, a.JID, a.RID, a.RAID, a.TCID, a.RCPTID, a.Amount, a.ASMID, a.EXPID, a.AcctRule, a.JAID)
	return updateError(err, "JournalAllocation", *a)
}

// UpdatePaymentType updates a PaymentType record in the database
func UpdatePaymentType(a *PaymentType) error {
	_, err := RRdb.Prepstmt.UpdatePaymentType.Exec(a.BID, a.Name, a.Description, a.LastModBy, a.PMTID)
	return updateError(err, "PaymentType", *a)
}

// UpdatePayor updates a Payor record in the database
func UpdatePayor(a *Payor) error {
	_, err := RRdb.Prepstmt.UpdatePayor.Exec(a.BID, a.CreditLimit, a.TaxpayorID, a.AccountRep, a.EligibleFuturePayor, a.LastModBy, a.TCID)
	return updateError(err, "Payor", *a)
}

// UpdateProspect updates a Prospect record in the database
func UpdateProspect(a *Prospect) error {
	_, err := RRdb.Prepstmt.UpdateProspect.Exec(a.BID, a.EmployerName, a.EmployerStreetAddress, a.EmployerCity,
		a.EmployerState, a.EmployerPostalCode, a.EmployerEmail, a.EmployerPhone, a.Occupation, a.ApplicationFee,
		a.DesiredUsageStartDate, a.RentableTypePreference, a.FLAGS, a.Approver, a.DeclineReasonSLSID, a.OtherPreferences,
		a.FollowUpDate, a.CSAgent, a.OutcomeSLSID, a.FloatingDeposit, a.RAID, a.LastModBy, a.TCID)
	return updateError(err, "Prospect", *a)
}

// UpdateRentable updates a Rentable record in the database
func UpdateRentable(a *Rentable) error {
	_, err := RRdb.Prepstmt.UpdateRentable.Exec(a.BID, a.RentableName, a.AssignmentTime, a.LastModBy, a.RID)
	return updateError(err, "Rentable", *a)
}

// UpdateRentableStatus updates a RentableStatus record in the database
func UpdateRentableStatus(a *RentableStatus) error {
	_, err := RRdb.Prepstmt.UpdateRentableStatus.Exec(a.RID, a.BID, a.DtStart, a.DtStop, a.DtNoticeToVacate, a.Status, a.LastModBy, a.RSID)
	return updateError(err, "RentableStatus", *a)
}

// UpdateRatePlan updates a RatePlan record in the database
func UpdateRatePlan(a *RatePlan) error {
	_, err := RRdb.Prepstmt.UpdateRatePlan.Exec(a.BID, a.Name, a.LastModBy, a.RPID)
	return updateError(err, "RatePlan", *a)
}

// UpdateRatePlanRef updates a RatePlanRef record in the database
func UpdateRatePlanRef(a *RatePlanRef) error {
	_, err := RRdb.Prepstmt.UpdateRatePlanRef.Exec(a.BID, a.RPID, a.DtStart, a.DtStop, a.FeeAppliesAge, a.MaxNoFeeUsers, a.AdditionalUserFee, a.PromoCode, a.CancellationFee, a.FLAGS, a.LastModBy, a.RPRID)
	return updateError(err, "RatePlanRef", *a)
}

// UpdateRatePlanRefRTRate updates a RatePlanRefRTRate record in the database
func UpdateRatePlanRefRTRate(a *RatePlanRefRTRate) error {
	_, err := RRdb.Prepstmt.UpdateRatePlanRefRTRate.Exec(a.BID, a.FLAGS, a.Val, a.RPRID, a.RTID)
	return updateError(err, "RatePlanRefRTRate", *a)
}

// UpdateRatePlanRefSPRate updates a RatePlanRefSPRate record in the database
func UpdateRatePlanRefSPRate(a *RatePlanRefSPRate) error {
	_, err := RRdb.Prepstmt.UpdateRatePlanRefSPRate.Exec(a.BID, a.FLAGS, a.Val, a.RPRID, a.RTID, a.RSPID)
	return updateError(err, "RatePlanRefSPRate", *a)
}

// UpdateReceipt updates a Receipt record in the database
func UpdateReceipt(a *Receipt) error {
	_, err := RRdb.Prepstmt.UpdateReceipt.Exec(a.PRCPTID, a.BID, a.TCID, a.PMTID, a.DEPID, a.DID, a.RAID, a.Dt, a.DocNo, a.Amount, a.AcctRuleReceive, a.ARID, a.AcctRuleApply, a.FLAGS, a.Comment, a.OtherPayorName, a.LastModBy, a.RCPTID)
	return updateError(err, "Receipt", *a)
}

// UpdateReceiptAllocation updates a ReceiptAllocation record in the database
func UpdateReceiptAllocation(a *ReceiptAllocation) error {
	_, err := RRdb.Prepstmt.UpdateReceiptAllocation.Exec(a.RCPTID, a.BID, a.RAID, a.Dt, a.Amount, a.ASMID, a.FLAGS, a.AcctRule, a.LastModBy, a.RCPAID)
	return updateError(err, "ReceiptAllocation", *a)
}

// UpdateRentalAgreement updates a RentalAgreement record in the database
func UpdateRentalAgreement(a *RentalAgreement) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreement.Exec(a.RATID, a.BID, a.NLID, a.AgreementStart, a.AgreementStop, a.PossessionStart, a.PossessionStop, a.RentStart, a.RentStop, a.RentCycleEpoch, a.UnspecifiedAdults, a.UnspecifiedChildren, a.Renewal, a.SpecialProvisions, a.LeaseType, a.ExpenseAdjustmentType, a.ExpensesStop, a.ExpenseStopCalculation, a.BaseYearEnd, a.ExpenseAdjustment, a.EstimatedCharges, a.RateChange, a.NextRateChange, a.PermittedUses, a.ExclusiveUses, a.ExtensionOption, a.ExtensionOptionNotice, a.ExpansionOption, a.ExpansionOptionNotice, a.RightOfFirstRefusal, a.FLAGS, a.LastModBy, a.RAID)

	return updateError(err, "RentalAgreement", *a)
}

// UpdateRentalAgreementPayor updates a RentalAgreementPayor record in the database
func UpdateRentalAgreementPayor(a *RentalAgreementPayor) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreementPayor.Exec(a.RAID, a.BID, a.TCID, a.DtStart, a.DtStop, a.FLAGS, a.RAPID)
	return updateError(err, "UpdateRentalAgreementPayor", *a)
}

// UpdateRentalAgreementPayorByRBT updates a RentalAgreementPayor record in the database
func UpdateRentalAgreementPayorByRBT(a *RentalAgreementPayor) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreementPayorByRBT.Exec(a.DtStart, a.DtStop, a.FLAGS, a.RAID, a.BID, a.TCID)
	return updateError(err, "UpdateRentalAgreementPayorByRBT", *a)
}

// UpdateRentalAgreementPet updates a RentalAgreementPet record in the database
func UpdateRentalAgreementPet(a *RentalAgreementPet) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreementPet.Exec(a.BID, a.RAID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.LastModBy, a.PETID)
	return updateError(err, "UpdateRentalAgreementPet", *a)
}

// UpdateRentalAgreementRentable updates a RentalAgreementRentable record in the database
func UpdateRentalAgreementRentable(a *RentalAgreementRentable) error {
	_, err := RRdb.Prepstmt.UpdateRentalAgreementRentable.Exec(a.RAID, a.BID, a.RID, a.CLID, a.ContractRent, a.RARDtStart, a.RARDtStop, a.RARID)
	return updateError(err, "RentalAgreementRentable", *a)
}

// UpdateRentableSpecialtyRef updates a RentableSpecialtyRef record in the database
func UpdateRentableSpecialtyRef(a *RentableSpecialtyRef) error {
	_, err := RRdb.Prepstmt.UpdateRentableSpecialtyRef.Exec(a.RSPID, a.LastModBy, a.RID, a.DtStart, a.DtStop)
	return updateError(err, "RentableSpecialtyRef", *a)
}

// UpdateRentableMarketRateInstance updates the given instance of RentableMarketRate
func UpdateRentableMarketRateInstance(a *RentableMarketRate) error {
	_, err := RRdb.Prepstmt.UpdateRentableMarketRateInstance.Exec(a.RTID, a.BID, a.MarketRate, a.DtStart, a.DtStop, a.RMRID)
	return updateError(err, "RentableMarketRate", *a)
}

// UpdateRentableType updates a RentableType record in the database
func UpdateRentableType(a *RentableType) error {
	_, err := RRdb.Prepstmt.UpdateRentableType.Exec(a.BID, a.Style, a.Name, a.RentCycle, a.Proration, a.GSRPC, a.ManageToBudget, a.LastModBy, a.RTID)
	return updateError(err, "RentableType", *a)
}

// ReactivateRentableType reactivates a RentableType record in the database
func ReactivateRentableType(a *RentableType) error {
	_, err := RRdb.Prepstmt.ReactivateRentableType.Exec(a.RTID)
	return updateError(err, "RentableType", *a)
}

// UpdateRentableTypeRef updates a RentableTypeRef record in the database
func UpdateRentableTypeRef(a *RentableTypeRef) error {
	//  SET BID=?,RTID=?,OverrideRentCycle=?,OverrideProrationCycle=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?"
	_, err := RRdb.Prepstmt.UpdateRentableTypeRef.Exec(a.RID, a.BID, a.RTID, a.OverrideRentCycle, a.OverrideProrationCycle, a.DtStart, a.DtStop, a.LastModBy, a.RTRID)
	return updateError(err, "RentableTypeRef", *a)
}

// UpdateRentableUser updates a RentableUser record in the database
func UpdateRentableUser(a *RentableUser) error {
	_, err := RRdb.Prepstmt.UpdateRentableUser.Exec(a.RID, a.BID, a.TCID, a.DtStart, a.DtStop, a.RUID)
	return updateError(err, "RentableUser", *a)
}

// UpdateRentableUserByRBT updates a RentableUser record in the database
func UpdateRentableUserByRBT(a *RentableUser) error {
	_, err := RRdb.Prepstmt.UpdateRentableUserByRBT.Exec(a.DtStart, a.DtStop, a.RID, a.BID, a.TCID)
	return updateError(err, "RentableUser", *a)
}

// UpdateStringList updates a StringList record in the database. It also updates the string list. It does this by
// deleting all the strings first, then inserting the ones it has.
func UpdateStringList(a *StringList) error {
	_, err := RRdb.Prepstmt.UpdateStringList.Exec(a.BID, a.Name, a.LastModBy, a.SLID)
	updateError(err, "StringList", *a)
	DeleteSLStrings(a.SLID)
	InsertSLStrings(a)
	return err
}

// UpdateSLString updates a SLString record in the database
func UpdateSLString(a *SLString) error {
	_, err := RRdb.Prepstmt.UpdateSLString.Exec(a.SLID, a.Value, a.LastModBy, a.SLSID)
	return updateError(err, "SLString", *a)
}

// UpdateTransactant updates a Transactant record in the database
func UpdateTransactant(a *Transactant) error {
	_, err := RRdb.Prepstmt.UpdateTransactant.Exec(a.BID, a.NLID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.LastModBy, a.TCID)
	return updateError(err, "Transactant", *a)
}

// UpdateUser updates a User record in the database
func UpdateUser(a *User) error {
	_, err := RRdb.Prepstmt.UpdateUser.Exec(a.BID, a.Points, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress, a.EmergencyContactTelephone, a.EmergencyEmail, a.AlternateAddress, a.EligibleFutureUser, a.Industry, a.SourceSLSID, a.LastModBy, a.TCID)
	return updateError(err, "User", *a)
}

// UpdateVehicle updates a Vehicle record in the database
func UpdateVehicle(a *Vehicle) error {
	_, err := RRdb.Prepstmt.UpdateVehicle.Exec(a.TCID, a.BID, a.VehicleType, a.VehicleMake, a.VehicleModel, a.VehicleColor, a.VehicleYear, a.LicensePlateState, a.LicensePlateNumber, a.ParkingPermitNumber, a.DtStart, a.DtStop, a.LastModBy, a.VID)
	return updateError(err, "Vehicle", *a)
}
