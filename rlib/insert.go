package rlib

// InsertAssessment writes a new assessmenttype record to the database. If the record is successfully written,
// the ASMID field is set to its new value.
func InsertAssessment(a *Assessment) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertAssessment.Exec(a.PASMID, a.BID, a.RID, a.ATypeLID, a.RAID, a.Amount, a.Start, a.Stop, a.RentCycle, a.ProrationCycle, a.InvoiceNo, a.AcctRule, a.Comment, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			a.ASMID = rid
		}
	} else {
		Ulog("InsertAssessment: error inserting Assessment:  %v\n", err)
		Ulog("Assessment = %#v\n", *a)
	}
	return rid, err
}

// InsertBuilding writes a new Building record to the database
func InsertBuilding(a *Building) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertBuilding.Exec(a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Building:  %v\n", err)
		Ulog("Bldg = %#v\n", *a)
	}
	return rid, err
}

// InsertBuildingWithID writes a new Building record to the database with the supplied bldgid
// the Building ID must be set in the supplied Building struct ptr (a.BLDGID).
func InsertBuildingWithID(a *Building) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertBuildingWithID.Exec(a.BLDGID, a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("InsertBuildingWithID: error inserting Building:  %v\n", err)
		Ulog("Bldg = %#v\n", *a)
	}
	return rid, err
}

// InsertBusiness writes a new Business record.
// returns the new Business ID and any associated error
func InsertBusiness(b *Business) (int64, error) {
	var bid = int64(0)
	res, err := RRdb.Prepstmt.InsertBusiness.Exec(b.Designation, b.Name, b.DefaultRentCycle, b.DefaultProrationCycle, b.DefaultGSRPC, b.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			bid = int64(id)
		}
	}
	return bid, err
}

// InsertCustomAttribute writes a new User record to the database
func InsertCustomAttribute(a *CustomAttribute) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertCustomAttribute.Exec(a.Type, a.Name, a.Value, a.Units, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertCustomAttribute: error inserting CustomAttribute:  %v\n", err)
		Ulog("CustomAttribute = %#v\n", *a)
	}
	return tid, err
}

// InsertCustomAttributeRef writes a new assessmenttype record to the database
func InsertCustomAttributeRef(a *CustomAttributeRef) error {
	_, err := RRdb.Prepstmt.InsertCustomAttributeRef.Exec(a.ElementType, a.ID, a.CID)
	return err
}

// InsertDemandSource writes a new DemandSource record to the database
func InsertDemandSource(a *DemandSource) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertDemandSource.Exec(a.BID, a.Name, a.Industry, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertDemandSource: error inserting DemandSource:  %v\n", err)
		Ulog("DemandSource = %#v\n", *a)
	}
	return tid, err
}

// InsertDeposit writes a new Deposit record to the database
func InsertDeposit(a *Deposit) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertDeposit.Exec(a.BID, a.DEPID, a.DPMID, a.Dt, a.Amount, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Deposit:  %v\n", err)
		Ulog("Deposit = %#v\n", *a)
	}
	return rid, err
}

// InsertDepositMethod writes a new DepositMethod record to the database
func InsertDepositMethod(a *DepositMethod) error {
	_, err := RRdb.Prepstmt.InsertDepositMethod.Exec(a.BID, a.Name)
	if nil != err {
		Ulog("Error inserting DepositMethod:  %v\n", err)
		Ulog("DepositMethod = %#v\n", *a)
	}
	return err
}

// InsertDepositPart writes a new DepositPart record to the database
func InsertDepositPart(a *DepositPart) error {
	_, err := RRdb.Prepstmt.InsertDepositPart.Exec(a.DID, a.RCPTID)
	if nil != err {
		Ulog("Error inserting DepositPart:  %v\n", err)
		Ulog("DepositPart = %#v\n", *a)
	}
	return err
}

// InsertDepository writes a new Depository record to the database
func InsertDepository(a *Depository) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertDepository.Exec(a.BID, a.Name, a.AccountNo, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Depository:  %v\n", err)
		Ulog("Depository = %#v\n", *a)
	}
	return rid, err
}

//======================================
//  INVOICE
//======================================

// InsertInvoice writes a new Invoice record to the database
func InsertInvoice(a *Invoice) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertInvoice.Exec(a.BID, a.Dt, a.DtDue, a.Amount, a.DeliveredBy, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Invoice:  %v\n", err)
		Ulog("Invoice = %#v\n", *a)
	}
	return rid, err
}

// InsertInvoiceAssessment writes a new InvoiceAssessment record to the database
func InsertInvoiceAssessment(a *InvoiceAssessment) error {
	_, err := RRdb.Prepstmt.InsertInvoiceAssessment.Exec(a.InvoiceNo, a.ASMID)
	if nil != err {
		Ulog("Error inserting InvoiceAssessment:  %v\n", err)
		Ulog("DepositPart = %#v\n", *a)
	}
	return err
}

// InsertInvoicePayor writes a new InvoicePayor record to the database
func InsertInvoicePayor(a *InvoicePayor) error {
	_, err := RRdb.Prepstmt.InsertInvoicePayor.Exec(a.InvoiceNo, a.PID)
	if nil != err {
		Ulog("Error inserting InvoicePayor:  %v\n", err)
		Ulog("DepositPayor = %#v\n", *a)
	}
	return err
}

// InsertJournalEntry writes a new Journal entry to the database
func InsertJournalEntry(j *Journal) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertJournal.Exec(j.BID, j.RAID, j.Dt, j.Amount, j.Type, j.ID, j.Comment, j.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			j.JID = rid
		}
	}
	return rid, err
}

// InsertJournalAllocationEntry writes a new JournalAllocation record to the database
func InsertJournalAllocationEntry(ja *JournalAllocation) error {
	_, err := RRdb.Prepstmt.InsertJournalAllocation.Exec(ja.JID, ja.RID, ja.Amount, ja.ASMID, ja.AcctRule)
	return err
}

// InsertJournalMarker writes a new JournalMarker record to the database
func InsertJournalMarker(jm *JournalMarker) error {
	_, err := RRdb.Prepstmt.InsertJournalMarker.Exec(jm.BID, jm.State, jm.DtStart, jm.DtStop)
	return err
}

//======================================
//  LEDGER MARKER
//======================================

// InsertLedgerMarker writes a new LedgerMarker record to the database
func InsertLedgerMarker(l *LedgerMarker) error {
	res, err := RRdb.Prepstmt.InsertLedgerMarker.Exec(l.LID, l.BID, l.RAID, l.RID, l.Dt, l.Balance, l.State, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			l.LMID = int64(id)
		}
	} else {
		Ulog("InsertLedgerMarker: err = %#v\n", err)
	}
	return err
}

// InsertLedgerEntry writes a new LedgerEntry to the database
func InsertLedgerEntry(l *LedgerEntry) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertLedgerEntry.Exec(l.BID, l.JID, l.JAID, l.LID, l.RAID, l.RID, l.Dt, l.Amount, l.Comment, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting LedgerEntry:  %v\n", err)
	}
	return rid, err
}

// InsertLedger writes a new GLAccount to the database
func InsertLedger(l *GLAccount) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertLedger.Exec(l.PLID, l.BID, l.RAID, l.GLNumber, l.Status, l.Type, l.Name, l.AcctType, l.RAAssociated, l.AllowPost, l.RARequired, l.ManageToBudget, l.Description, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
			l.LID = rid
		}
	} else {
		Ulog("Error inserting GLAccount:  %v\n", err)
	}
	return rid, err
}

//======================================
// NOTE
//======================================

// InsertNote writes a new Note to the database
func InsertNote(a *Note) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertNote.Exec(a.NLID, a.PNID, a.NTID, a.RID, a.RAID, a.TCID, a.Comment, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting Note:  %v\n", err)
	}
	return rid, err
}

//======================================
// NOTE LIST
//======================================

// InsertNoteList inserts a new wrapper for a notelist into the database
func InsertNoteList(a *NoteList) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertNoteList.Exec(a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting NoteList:  %v\n", err)
	}
	return rid, err
}

//======================================
// NOTE TYPE
//======================================

// InsertNoteType writes a new NoteType to the database
func InsertNoteType(a *NoteType) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertNoteType.Exec(a.BID, a.Name, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting NoteType:  %v\n", err)
	}
	return rid, err
}

//=======================================================
//  RATE PLAN
//=======================================================

// InsertRatePlan writes a new RatePlan record to the database
func InsertRatePlan(a *RatePlan) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRatePlan.Exec(a.BID, a.Name, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRatePlan: error inserting RatePlan:  %v\n", err)
		Ulog("RatePlan = %#v\n", *a)
	}
	a.RPID = tid
	return tid, err
}

// InsertRatePlanRef writes a new RatePlanRef record to the database
func InsertRatePlanRef(a *RatePlanRef) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRatePlanRef.Exec(a.RPID, a.DtStart, a.DtStop, a.FeeAppliesAge, a.MaxNoFeeUsers, a.AdditionalUserFee, a.PromoCode, a.CancellationFee, a.FLAGS, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRatePlanRef: error inserting RatePlanRef:  %v\n", err)
		Ulog("RatePlanRef = %#v\n", *a)
	}
	a.RPRID = tid
	return tid, err
}

// InsertRatePlanRefRTRate writes a new RatePlanRefRTRate record to the database
func InsertRatePlanRefRTRate(a *RatePlanRefRTRate) error {
	_, err := RRdb.Prepstmt.InsertRatePlanRefRTRate.Exec(a.RPRID, a.RTID, a.FLAGS, a.Val)
	if nil != err {
		Ulog("InsertRatePlanRefRTRate: error inserting RatePlanRefRTRate:  %v\n", err)
		Ulog("RatePlanRefRTRate = %#v\n", *a)
	}
	return err
}

// InsertRatePlanRefSPRate writes a new RatePlanRefSPRate record to the database
func InsertRatePlanRefSPRate(a *RatePlanRefSPRate) error {
	_, err := RRdb.Prepstmt.InsertRatePlanRefSPRate.Exec(a.RPRID, a.RTID, a.RSPID, a.FLAGS, a.Val)
	if nil != err {
		Ulog("InsertRatePlanRefSPRate: error inserting RatePlanRefSPRate:  %v\n", err)
		Ulog("RatePlanRefSPRate = %#v\n", *a)
	}
	return err
}

//=======================================================
//  PAYMENT
//=======================================================

// InsertPaymentType writes a new assessmenttype record to the database
func InsertPaymentType(a *PaymentType) error {
	_, err := RRdb.Prepstmt.InsertPaymentType.Exec(a.BID, a.Name, a.Description, a.LastModBy)
	return err
}

// InsertPayor writes a new User record to the database
func InsertPayor(a *Payor) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertPayor.Exec(a.TCID, a.CreditLimit, a.TaxpayorID, a.AccountRep, a.EligibleFuturePayor, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertPayor: error inserting Payor:  %v\n", err)
		Ulog("Payor = %#v\n", *a)
	}
	return tid, err
}

// InsertProspect writes a new User record to the database
func InsertProspect(a *Prospect) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertProspect.Exec(a.TCID, a.EmployerName, a.EmployerStreetAddress, a.EmployerCity,
		a.EmployerState, a.EmployerPostalCode, a.EmployerEmail, a.EmployerPhone, a.Occupation, a.ApplicationFee,
		a.DesiredUsageStartDate, a.RentableTypePreference, a.FLAGS, a.Approver, a.DeclineReasonSLSID, a.OtherPreferences,
		a.FollowUpDate, a.CSAgent, a.OutcomeSLSID, a.FloatingDeposit, a.RAID, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertProspect: error inserting Prospect:  %v\n", err)
		Ulog("Prospect = %#v\n", *a)
	}
	return tid, err
}

// InsertRentable writes a new Rentable record to the database
func InsertRentable(a *Rentable) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentable.Exec(a.BID, a.Name, a.AssignmentTime, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("InsertRentable: error inserting Building:  %v\n", err)
		Ulog("Rentable = %#v\n", *a)
	}
	return rid, err
}

//=======================================================
//  R E C E I P T
//=======================================================

// InsertReceipt writes a new Receipt record to the database. If the record is successfully written,
// the RCPTID field is set to its new value.
func InsertReceipt(r *Receipt) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertReceipt.Exec(r.PRCPTID, r.BID, r.RAID, r.PMTID, r.Dt, r.DocNo, r.Amount, r.AcctRule, r.Comment, r.OtherPayorName, r.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
			r.RCPTID = tid
		}
	} else {
		Ulog("InsertReceipt: error inserting Receipt:  %v\n", err)
		Ulog("Receipt = %#v\n", *r)
	}
	return tid, err
}

// InsertReceiptAllocation writes a new ReceiptAllocation record to the database
func InsertReceiptAllocation(r *ReceiptAllocation) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertReceiptAllocation.Exec(r.RCPTID, r.Amount, r.ASMID, r.AcctRule)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertReceiptAllocation: error inserting ReceiptAllocation:  %v\n", err)
		Ulog("ReceiptAllocation = %#v\n", *r)
	}
	return tid, err
}

// InsertRentalAgreement writes a new RentalAgreement record to the database
func InsertRentalAgreement(a *RentalAgreement) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreement.Exec(a.RATID, a.BID, a.NLID, a.AgreementStart, a.AgreementStop, a.PossessionStart, a.PossessionStop, a.RentStart, a.RentStop, a.RentCycleEpoch, a.UnspecifiedAdults, a.UnspecifiedChildren, a.Renewal, a.SpecialProvisions, a.LeaseType, a.ExpenseAdjustmentType, a.ExpensesStop, a.ExpenseStopCalculation, a.BaseYearEnd, a.ExpenseAdjustment, a.EstimatedCharges, a.RateChange, a.NextRateChange, a.PermittedUses, a.ExclusiveUses, a.ExtensionOption, a.ExtensionOptionNotice, a.ExpansionOption, a.ExpansionOptionNotice, a.RightOfFirstRefusal, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRentalAgreement: error inserting RentalAgreement:  %v\n", err)
		Ulog("RentalAgreement = %#v\n", *a)
	}
	return tid, err
}

// InsertRentalAgreementPayor writes a new User record to the database
func InsertRentalAgreementPayor(a *RentalAgreementPayor) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementPayor.Exec(a.RAID, a.TCID, a.DtStart, a.DtStop, a.FLAGS)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRentalAgreementPayor: error inserting RentalAgreementRentable:  %v\n", err)
		Ulog("RentalAgreementPayor = %#v\n", *a)
	}
	return tid, err
}

// InsertRentalAgreementPet writes a new User record to the database
func InsertRentalAgreementPet(a *RentalAgreementPet) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementPet.Exec(a.RAID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRentalAgreementPet: error inserting RentalAgreementRentable:  %v\n", err)
		Ulog("RentalAgreementPet = %#v\n", *a)
	}
	return tid, err
}

// InsertRentalAgreementRentable writes a new User record to the database
func InsertRentalAgreementRentable(a *RentalAgreementRentable) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementRentable.Exec(a.RAID, a.RID, a.CLID, a.ContractRent, a.DtStart, a.DtStop)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRentalAgreementRentable: error inserting RentalAgreementRentable:  %v\n", err)
		Ulog("RentalAgreementRentable = %#v\n", *a)
	}
	return tid, err
}

//=======================================================
//  RENTAL AGREEMENT TEMPLATE
//=======================================================

// InsertRentalAgreementTemplate writes a new User record to the database
func InsertRentalAgreementTemplate(a *RentalAgreementTemplate) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentalAgreementTemplate.Exec(a.BID, a.RATemplateName, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRentalAgreementTemplate: error inserting RentalAgreementTemplate:  %v\n", err)
		Ulog("RentalAgreementTemplate = %#v\n", *a)
	}
	return tid, err
}

// InsertRentableSpecialty writes a new RentableSpecialty record to the database
func InsertRentableSpecialty(a *RentableSpecialty) error {
	_, err := RRdb.Prepstmt.InsertRentableSpecialtyType.Exec(a.RSPID, a.BID, a.Name, a.Fee, a.Description)
	return err
}

// InsertRentableMarketRates writes a new marketrate record to the database
func InsertRentableMarketRates(r *RentableMarketRate) error {
	_, err := RRdb.Prepstmt.InsertRentableMarketRates.Exec(r.RTID, r.MarketRate, r.DtStart, r.DtStop)
	return err
}

// InsertRentableType writes a new RentableType record to the database
func InsertRentableType(a *RentableType) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentableType.Exec(a.BID, a.Style, a.Name, a.RentCycle, a.Proration, a.GSRPC, a.ManageToBudget, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting RentableType:  %v\n", err)
	}
	return rid, err
}

// InsertRentableSpecialtyRef writes a new RentableSpecialty record to the database
func InsertRentableSpecialtyRef(a *RentableSpecialtyRef) error {
	_, err := RRdb.Prepstmt.InsertRentableSpecialtyRef.Exec(a.BID, a.RID, a.RSPID, a.DtStart, a.DtStop, a.LastModBy)
	return err
}

// InsertRentableStatus writes a new RentableStatus record to the database
func InsertRentableStatus(a *RentableStatus) error {
	_, err := RRdb.Prepstmt.InsertRentableStatus.Exec(a.RID, a.DtStart, a.DtStop, a.DtNoticeToVacate, a.Status, a.LastModBy)
	return err
}

// InsertRentableTypeRef writes a new RentableTypeRef record to the database
func InsertRentableTypeRef(a *RentableTypeRef) error {
	_, err := RRdb.Prepstmt.InsertRentableTypeRef.Exec(a.RID, a.RTID, a.OverrideRentCycle, a.OverrideProrationCycle, a.DtStart, a.DtStop, a.LastModBy)
	return err
}

// InsertRentableUser writes a new User record to the database
func InsertRentableUser(a *RentableUser) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertRentableUser.Exec(a.RID, a.TCID, a.DtStart, a.DtStop)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertRentableUser: error inserting RentableUser:  %v\n", err)
		Ulog("RentableUser = %#v\n", *a)
	}
	return tid, err
}

// InsertStringList writes a new StringList record to the database
func InsertStringList(a *StringList) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertStringList.Exec(a.BID, a.Name, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertStringList: error inserting StringList:  %v\n", err)
		Ulog("StringList = %#v\n", *a)
	}
	a.SLID = tid
	InsertSLStrings(a)
	return tid, err
}

// InsertSLStrings writes a the list of strings in a StringList to the database
// THIS SHOULD BE PUT IN A TRANSACTION
func InsertSLStrings(a *StringList) {
	// DeleteSLStrings(a.SLID)
	for i := 0; i < len(a.S); i++ {
		a.S[i].SLID = a.SLID
		_, err := RRdb.Prepstmt.InsertSLString.Exec(a.SLID, a.S[i].Value, a.S[i].LastModBy)
		if nil != err {
			Ulog("InsertSLString: error:  %v\n", err)
		}
	}
}

// InsertTransactant writes a new Transactant record to the database
func InsertTransactant(a *Transactant) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertTransactant.Exec(a.BID, a.NLID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertTransactant: error inserting Transactant:  %v\n", err)
		Ulog("Transactant = %#v\n", *a)
	}
	return tid, err
}

// InsertUser writes a new User record to the database
func InsertUser(a *User) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertUser.Exec(a.TCID, a.Points, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress, a.EmergencyContactTelephone, a.EmergencyEmail, a.AlternateAddress, a.EligibleFutureUser, a.Industry, a.SourceSLSID, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertUser: error inserting User:  %v\n", err)
		Ulog("User = %#v\n", *a)
	}
	return tid, err
}

// InsertVehicle writes a new Vehicle record to the database
func InsertVehicle(a *Vehicle) (int64, error) {
	var tid = int64(0)
	res, err := RRdb.Prepstmt.InsertVehicle.Exec(a.TCID, a.BID, a.VehicleType, a.VehicleMake, a.VehicleModel, a.VehicleColor, a.VehicleYear, a.LicensePlateState, a.LicensePlateNumber, a.ParkingPermitNumber, a.DtStart, a.DtStop, a.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			tid = int64(id)
		}
	} else {
		Ulog("InsertVehicle: error inserting Vehicle:  %v\n", err)
		Ulog("Vehicle = %#v\n", *a)
	}
	return tid, err
}
