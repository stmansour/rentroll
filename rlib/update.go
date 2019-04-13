package rlib

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"extres"
	"runtime/debug"
)

// updateSessionProblem is a convenience function that replaces 8 lines
// of code with about 4. Since these lines are needed for every update call
// it saves a lot of lines.  Added this routine at the time Task,TaskList,
// TaskDescriptor and  TaskListDefinition were added.
//-----------------------------------------------------------------------------
func updateSessionProblem(ctx context.Context, id1, id2 *int64) error {
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
		(*id1) = sess.UID
		(*id2) = sess.UID
		return nil
	}
	return nil
}

func updateError(err error, n string, a interface{}) error {
	if nil != err {
		Ulog("Update%s: error updating %s:  %v\n", n, n, err)
		Ulog("%s = %#v\n", n, a)
	}
	return err
}

// UpdateAR updates an AR record
func UpdateAR(ctx context.Context, a *AR) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Name, a.ARType, a.DebitLID, a.CreditLID, a.Description, a.RARequired, a.DtStart, a.DtStop, a.FLAGS, a.DefaultAmount, a.DefaultRentCycle, a.DefaultProrationCycle, a.LastModBy, a.ARID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateAR)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateAR.Exec(fields...)
	}
	return updateError(err, "AR", *a)
}

// UpdateAssessment updates an Assessment record
func UpdateAssessment(ctx context.Context, a *Assessment) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	// DEBUG
	// just looking for where a problem is coming from
	if a.Stop.Before(a.Start) {
		Console("\n\n **** FATAL ****   **** FATAL ****  stop date prior to start\n")
		debug.PrintStack()
		Console("\n **** FATAL ****   **** FATAL ****  stop date prior to start\n\n")
	}

	a.Amount = Round(a.Amount, .5, 2)
	fields := []interface{}{
		a.PASMID,
		a.RPASMID,
		a.AGRCPTID,
		a.BID,
		a.RID,
		a.AssocElemType,
		a.AssocElemID,
		a.RAID,
		a.Amount,
		a.Start,
		a.Stop,
		a.RentCycle,
		a.ProrationCycle,
		a.InvoiceNo,
		a.AcctRule,
		a.ARID,
		a.FLAGS,
		a.Comment,
		a.LastModBy,
		a.ASMID,
	}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateAssessment)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateAssessment.Exec(fields...)
	}
	return updateError(err, "Assessment", *a)
}

// UpdateBusiness updates an Business record
func UpdateBusiness(ctx context.Context, a *Business) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	// TODO(Sudip): keep mind this FLAGS insertion in fields, this might be removed in the future
	fields := []interface{}{a.Designation, a.Name, a.DefaultRentCycle, a.DefaultProrationCycle, a.DefaultGSRPC, a.ClosePeriodTLID, a.FLAGS, a.LastModBy, a.BID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateBusiness)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateBusiness.Exec(fields...)
	}

	// build business list and cache again
	RRdb.BUDlist, RRdb.BizCache = BuildBusinessDesignationMap(ctx)

	return updateError(err, "Business", *a)
}

// UpdateBusinessPropertiesData updates the flow Data json column
func UpdateBusinessPropertiesData(ctx context.Context, jsonDataKey string, jsonData []byte, a *BusinessProperties) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	// make sure that json is valid before inserting it in database
	if !(IsByteDataValidJSON(jsonData)) {
		return ErrFlowInvalidJSONData
	}

	// as a.Data is type of json.RawMessage - convert it to byte stream so that it can be inserted
	// in mysql `json` type column
	fields := []interface{}{jsonDataKey, jsonData, a.BPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateBusinessPropertiesData)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateBusinessPropertiesData.Exec(fields...)
	}
	return updateError(err, "BusinessProperties", *a)
}

// UpdateClosePeriod updates an ClosePeriod record
func UpdateClosePeriod(ctx context.Context, a *ClosePeriod) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.TLID, a.Dt, a.CreateBy, a.LastModBy, a.CPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateClosePeriod)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateClosePeriod.Exec(fields...)
	}

	return updateError(err, "ClosePeriod", *a)
}

// UpdateCustomAttribute updates an CustomAttribute record
func UpdateCustomAttribute(ctx context.Context, a *CustomAttribute) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Type, a.Name, a.Value, a.Units, a.LastModBy, a.CID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateCustomAttribute)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateCustomAttribute.Exec(fields...)
	}
	return updateError(err, "CustomAttribute", *a)
}

// UpdateDemandSource updates a DemandSource record in the database
func UpdateDemandSource(ctx context.Context, a *DemandSource) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.Name, a.Industry, a.LastModBy, a.SourceSLSID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateDemandSource)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateDemandSource.Exec(fields...)
	}
	return updateError(err, "DemandSource", *a)
}

// UpdateDeposit updates a Deposit record
func UpdateDeposit(ctx context.Context, a *Deposit) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.DEPID, a.DPMID, a.Dt, a.Amount, a.ClearedAmount, a.FLAGS, a.LastModBy, a.DID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateDeposit)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateDeposit.Exec(fields...)
	}
	return updateError(err, "Deposit", *a)
}

// UpdateDepository updates a Depository record
func UpdateDepository(ctx context.Context, a *Depository) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.LID, a.Name, a.AccountNo, a.LastModBy, a.DEPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateDepository)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateDepository.Exec(fields...)
	}
	return updateError(err, "Depository", *a)
}

// UpdateDepositMethod updates a DepositMethod record
func UpdateDepositMethod(ctx context.Context, a *DepositMethod) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Method, a.LastModBy, a.DPMID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateDepositMethod)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateDepositMethod.Exec(fields...)
	}
	return updateError(err, "DepositMethod", *a)
}

// UpdateDepositPart updates a DepositPart record
func UpdateDepositPart(ctx context.Context, a *DepositPart) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.DID, a.BID, a.RCPTID, a.LastModBy, a.DPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateDepositPart)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateDepositPart.Exec(fields...)
	}
	return updateError(err, "DepositPart", *a)
}

// UpdateExpense updates a Expense record
func UpdateExpense(ctx context.Context, a *Expense) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	a.Amount = Round(a.Amount, .5, 2)
	fields := []interface{}{a.RPEXPID, a.BID, a.RID, a.RAID, a.Amount, a.Dt, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.LastModBy, a.EXPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateExpense)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateExpense.Exec(fields...)
	}
	return updateError(err, "Expense", *a)
}

// UpdateRAFlowWithInitState updates the flow record with resetting it's state
// to application being complete
func UpdateRAFlowWithInitState(ctx context.Context, a *Flow) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	// make sure that json is valid before inserting it in database
	if !(IsByteDataValidJSON(a.Data)) {
		return ErrFlowInvalidJSONData
	}

	// -------------------------------------------------------------------------
	// RESET META - CHANGE STATE TO APPLICATION BEING COMPLETED
	// -------------------------------------------------------------------------
	flowData := RAFlowJSONData{}
	err = json.Unmarshal([]byte(a.Data), &flowData)
	if err != nil {
		return err
	}

	//  IF DATA HAS BEEN CHANGED, RESET META AND SET STATE TO APP BEING COMPLETED
	resetMeta := RAFlowMetaInfo{
		RAID:         flowData.Meta.RAID,
		BID:          flowData.Meta.BID,
		LastTMPPETID: flowData.Meta.LastTMPPETID,
		LastTMPVID:   flowData.Meta.LastTMPVID,
		LastTMPTCID:  flowData.Meta.LastTMPTCID,
		LastTMPASMID: flowData.Meta.LastTMPASMID,
		HavePets:     flowData.Meta.HavePets,
		HaveVehicles: flowData.Meta.HaveVehicles,
	}
	flowData.Meta = resetMeta

	var modFlowJSONData []byte
	modFlowJSONData, err = json.Marshal(&flowData)
	if err != nil {
		return err
	}
	// -------------------------------------------------------------------------

	fields := []interface{}{a.BID, a.UserRefNo, a.FlowType, a.ID, modFlowJSONData, a.LastModBy, a.FlowID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateFlow)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateFlow.Exec(fields...)
	}
	return updateError(err, "Flow", *a)
}

// UpdateFlow updates the flow record
func UpdateFlow(ctx context.Context, a *Flow) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	// make sure that json is valid before inserting it in database
	if !(IsByteDataValidJSON(a.Data)) {
		return ErrFlowInvalidJSONData
	}

	fields := []interface{}{a.BID, a.UserRefNo, a.FlowType, a.ID, []byte(a.Data), a.LastModBy, a.FlowID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateFlow)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateFlow.Exec(fields...)
	}
	return updateError(err, "Flow", *a)
}

// UpdateFlowPartData updates the flow Data json column
func UpdateFlowPartData(ctx context.Context, jsonDataKey string, jsonData []byte, a *Flow) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	// make sure that json is valid before inserting it in database
	if !(IsByteDataValidJSON(jsonData)) {
		return ErrFlowInvalidJSONData
	}

	// as a.Data is type of json.RawMessage - convert it to byte stream so that it can be inserted
	// in mysql `json` type column
	fields := []interface{}{jsonDataKey, jsonData, a.FlowID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateFlowPartData)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateFlowPartData.Exec(fields...)
	}
	return updateError(err, "FlowPartData", *a)
}

// UpdateInvoice updates a Invoice record
func UpdateInvoice(ctx context.Context, a *Invoice) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Dt, a.DtDue, a.Amount, a.DeliveredBy, a.LastModBy, a.InvoiceNo}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateInvoice)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateInvoice.Exec(fields...)
	}
	return updateError(err, "Invoice", *a)
}

// UpdateLedgerMarker updates a LedgerMarker record
func UpdateLedgerMarker(ctx context.Context, a *LedgerMarker) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.LID, a.BID, a.RAID, a.RID, a.TCID, a.Dt, a.Balance, a.State, a.LastModBy, a.LMID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateLedgerMarker)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateLedgerMarker.Exec(fields...)
	}
	return updateError(err, "LedgerMarker", *a)
}

// UpdateLedger updates a Ledger record
func UpdateLedger(ctx context.Context, a *GLAccount) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.PLID, a.BID, a.RAID, a.TCID, a.GLNumber, a.Name, a.AcctType, a.AllowPost, a.FLAGS, a.Description, a.LastModBy, a.LID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateLedger)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateLedger.Exec(fields...)
	}
	return updateError(err, "GLAccount", *a)
}

// UpdateJournalAllocation updates a JournalAllocation record
func UpdateJournalAllocation(ctx context.Context, a *JournalAllocation) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.JID, a.RID, a.RAID, a.TCID, a.RCPTID, a.Amount, a.ASMID, a.EXPID, a.AcctRule, a.LastModBy, a.JAID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateJournalAllocation)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateJournalAllocation.Exec(fields...)
	}
	return updateError(err, "JournalAllocation", *a)
}

// UpdatePaymentType updates a PaymentType record in the database
func UpdatePaymentType(ctx context.Context, a *PaymentType) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Name, a.Description, a.LastModBy, a.PMTID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdatePaymentType)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdatePaymentType.Exec(fields...)
	}
	return updateError(err, "PaymentType", *a)
}

// UpdatePayor updates a Payor record in the database
func UpdatePayor(ctx context.Context, a *Payor) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}
	t1, err := Encrypt(a.TaxpayorID)
	if err != nil {
		return err
	}
	t := hex.EncodeToString(t1)

	d1, err := Encrypt(a.DriversLicense)
	if err != nil {
		return err
	}
	d := hex.EncodeToString(d1)

	fields := []interface{}{a.TCID, a.BID, a.CreditLimit, t, a.EligibleFuturePayor,
		a.FLAGS, d, a.GrossIncome, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdatePayor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdatePayor.Exec(fields...)
	}
	return updateError(err, "Payor", *a)
}

// UpdateProspect updates a Prospect record in the database
func UpdateProspect(ctx context.Context, a *Prospect) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{
		a.BID,
		a.CompanyAddress,
		a.CompanyCity,
		a.CompanyState,
		a.CompanyPostalCode,
		a.CompanyEmail,
		a.CompanyPhone,
		a.Occupation,
		a.EvictedDes,
		a.ConvictedDes,
		a.BankruptcyDes,
		a.FollowUpDate,
		a.FLAGS,
		a.OtherPreferences,
		a.SpecialNeeds,
		a.CurrentAddress,
		a.CurrentLandLordName,
		a.CurrentLandLordPhoneNo,
		a.CurrentReasonForMoving,
		a.CurrentLengthOfResidency,
		a.PriorAddress,
		a.PriorLandLordName,
		a.PriorLandLordPhoneNo,
		a.PriorReasonForMoving,
		a.PriorLengthOfResidency,
		a.CommissionableThirdParty,
		a.ThirdPartySource,
		a.LastModBy,
		a.TCID,
	}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateProspect)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateProspect.Exec(fields...)
	}
	return updateError(err, "Prospect", *a)
}

// UpdateRentable updates a Rentable record in the database
func UpdateRentable(ctx context.Context, a *Rentable) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.PRID, a.RentableName, a.AssignmentTime, a.MRStatus, a.DtMRStart, a.Comment, a.LastModBy, a.RID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentable)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentable.Exec(fields...)
	}
	return updateError(err, "Rentable", *a)
}

// UpdateRentableLeaseStatus updates a RentableLeaseStatus record in the database
func UpdateRentableLeaseStatus(ctx context.Context, a *RentableLeaseStatus) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	if a.DtStop.Before(a.DtStart) {
		debug.PrintStack()
	}

	fields := []interface{}{a.RID, a.BID, a.DtStart, a.DtStop, a.LeaseStatus, a.Comment, a.FirstName, a.LastName, a.Email, a.Phone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.CCName, a.CCType, a.CCNumber, a.CCExpMonth, a.LastModBy, a.RLID}

	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableLeaseStatus)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableLeaseStatus.Exec(fields...)
	}
	return updateError(err, "RentableLeaseStatus", *a)
}

// UpdateRentableUseStatus updates a RentableUseStatus record in the database
func UpdateRentableUseStatus(ctx context.Context, a *RentableUseStatus) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.RID, a.BID, a.DtStart, a.DtStop, a.Comment, a.UseStatus /*a.LeaseStatus,*/, a.LastModBy, a.RSID}

	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableUseStatus)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableUseStatus.Exec(fields...)
	}
	return updateError(err, "RentableUseStatus", *a)
}

// UpdateRentableUseType updates a RentableUseType record in the database
func UpdateRentableUseType(ctx context.Context, a *RentableUseType) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.RID, a.BID, a.DtStart, a.DtStop, a.Comment, a.UseType, a.LastModBy, a.UTID}

	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableUseType)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableUseType.Exec(fields...)
	}
	return updateError(err, "RentableUseType", *a)
}

// UpdateRatePlan updates a RatePlan record in the database
func UpdateRatePlan(ctx context.Context, a *RatePlan) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Name, a.LastModBy, a.RPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRatePlan)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRatePlan.Exec(fields...)
	}
	return updateError(err, "RatePlan", *a)
}

// UpdateRatePlanRef updates a RatePlanRef record in the database
func UpdateRatePlanRef(ctx context.Context, a *RatePlanRef) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.RPID, a.DtStart, a.DtStop, a.FeeAppliesAge, a.MaxNoFeeUsers, a.AdditionalUserFee, a.PromoCode, a.CancellationFee, a.FLAGS, a.LastModBy, a.RPRID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRatePlanRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRatePlanRef.Exec(fields...)
	}
	return updateError(err, "RatePlanRef", *a)
}

// UpdateRatePlanRefRTRate updates a RatePlanRefRTRate record in the database
func UpdateRatePlanRefRTRate(ctx context.Context, a *RatePlanRefRTRate) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.FLAGS, a.Val, a.LastModBy, a.RPRID, a.RTID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRatePlanRefRTRate)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRatePlanRefRTRate.Exec(fields...)
	}
	return updateError(err, "RatePlanRefRTRate", *a)
}

// UpdateRatePlanRefSPRate updates a RatePlanRefSPRate record in the database
func UpdateRatePlanRefSPRate(ctx context.Context, a *RatePlanRefSPRate) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.FLAGS, a.Val, a.LastModBy, a.RPRID, a.RTID, a.RSPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRatePlanRefSPRate)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRatePlanRefSPRate.Exec(fields...)
	}
	return updateError(err, "RatePlanRefSPRate", *a)
}

// UpdateReceipt updates a Receipt record in the database
func UpdateReceipt(ctx context.Context, a *Receipt) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	a.Amount = Round(a.Amount, .5, 2)
	fields := []interface{}{a.PRCPTID, a.BID, a.TCID, a.PMTID, a.DEPID, a.DID, a.RAID, a.Dt, a.DocNo, a.Amount, a.AcctRuleReceive, a.ARID, a.AcctRuleApply, a.FLAGS, a.Comment, a.OtherPayorName, a.LastModBy, a.RCPTID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateReceipt)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateReceipt.Exec(fields...)
	}
	return updateError(err, "Receipt", *a)
}

// UpdateReceiptAllocation updates a ReceiptAllocation record in the database
func UpdateReceiptAllocation(ctx context.Context, a *ReceiptAllocation) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	a.Amount = Round(a.Amount, .5, 2)
	fields := []interface{}{a.RCPTID, a.BID, a.RAID, a.Dt, a.Amount, a.ASMID, a.FLAGS, a.AcctRule, a.LastModBy, a.RCPAID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateReceiptAllocation)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateReceiptAllocation.Exec(fields...)
	}
	return updateError(err, "ReceiptAllocation", *a)
}

// UpdateRentalAgreement updates a RentalAgreement record in the database
func UpdateRentalAgreement(ctx context.Context, a *RentalAgreement) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{
		a.PRAID,
		a.ORIGIN,
		a.RATID,
		a.BID,
		a.NLID,
		a.DocumentDate,
		a.AgreementStart,
		a.AgreementStop,
		a.PossessionStart,
		a.PossessionStop,
		a.RentStart,
		a.RentStop,
		a.RentCycleEpoch,
		a.UnspecifiedAdults,
		a.UnspecifiedChildren,
		a.Renewal,
		a.SpecialProvisions,
		a.LeaseType,
		a.ExpenseAdjustmentType,
		a.ExpensesStop,
		a.ExpenseStopCalculation,
		a.BaseYearEnd,
		a.ExpenseAdjustment,
		a.EstimatedCharges,
		a.RateChange,
		a.CSAgent,
		a.NextRateChange,
		a.PermittedUses,
		a.ExclusiveUses,
		a.ExtensionOption,
		a.ExtensionOptionNotice,
		a.ExpansionOption,
		a.ExpansionOptionNotice,
		a.RightOfFirstRefusal,
		a.DesiredUsageStartDate,
		a.RentableTypePreference,
		a.FLAGS,
		a.ApplicationReadyUID,
		a.ApplicationReadyDate,
		a.Approver1,
		a.DecisionDate1,
		a.DeclineReason1,
		a.Approver2,
		a.DecisionDate2,
		a.DeclineReason2,
		a.MoveInUID,
		a.MoveInDate,
		a.ActiveUID,
		a.ActiveDate,
		a.Outcome,
		a.NoticeToMoveUID,
		a.NoticeToMoveDate,
		a.NoticeToMoveReported,
		a.TerminatorUID,
		a.TerminationDate,
		a.LeaseTerminationReason,
		a.LastModBy,
		a.RAID,
	}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentalAgreement)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentalAgreement.Exec(fields...)
	}

	return updateError(err, "RentalAgreement", *a)
}

// UpdateRentalAgreementPayor updates a RentalAgreementPayor record in the database
func UpdateRentalAgreementPayor(ctx context.Context, a *RentalAgreementPayor) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.RAID, a.BID, a.TCID, a.DtStart, a.DtStop, a.FLAGS, a.LastModBy, a.RAPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentalAgreementPayor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentalAgreementPayor.Exec(fields...)
	}
	return updateError(err, "UpdateRentalAgreementPayor", *a)
}

// UpdateRentalAgreementPayorByRBT updates a RentalAgreementPayor record in the database
func UpdateRentalAgreementPayorByRBT(ctx context.Context, a *RentalAgreementPayor) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.DtStart, a.DtStop, a.FLAGS, a.LastModBy, a.RAID, a.BID, a.TCID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentalAgreementPayorByRBT)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentalAgreementPayorByRBT.Exec(fields...)
	}
	return updateError(err, "UpdateRentalAgreementPayorByRBT", *a)
}

// UpdatePet updates a Pet record in the database
func UpdatePet(ctx context.Context, a *Pet) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.RAID, a.TCID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.LastModBy, a.PETID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdatePet)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdatePet.Exec(fields...)
	}
	return updateError(err, "UpdatePet", *a)
}

// UpdateRentalAgreementRentable updates a RentalAgreementRentable record in the database
func UpdateRentalAgreementRentable(ctx context.Context, a *RentalAgreementRentable) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.RAID, a.BID, a.RID, a.CLID, a.ContractRent, a.RARDtStart, a.RARDtStop, a.LastModBy, a.RARID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentalAgreementRentable)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentalAgreementRentable.Exec(fields...)
	}
	return updateError(err, "RentalAgreementRentable", *a)
}

// UpdateRentableSpecialtyRef updates a RentableSpecialtyRef record in the database
func UpdateRentableSpecialtyRef(ctx context.Context, a *RentableSpecialtyRef) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.RID, a.RSPID, a.DtStart, a.DtStop, a.LastModBy, a.RID, a.DtStart, a.DtStop}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableSpecialtyRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableSpecialtyRef.Exec(fields...)
	}
	return updateError(err, "RentableSpecialtyRef", *a)
}

// UpdateRentableMarketRate updates the given instance of RentableMarketRate
func UpdateRentableMarketRate(ctx context.Context, a *RentableMarketRate) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.RTID, a.BID, a.MarketRate, a.DtStart, a.DtStop, a.LastModBy, a.RMRID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableMarketRate)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableMarketRate.Exec(fields...)
	}
	return updateError(err, "RentableMarketRate", *a)
}

// UpdateRentableType updates a RentableType record in the database
func UpdateRentableType(ctx context.Context, a *RentableType) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Style, a.Name, a.RentCycle, a.Proration, a.GSRPC, a.ARID, a.FLAGS, a.LastModBy, a.RTID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableType)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableType.Exec(fields...)
	}
	return updateError(err, "RentableType", *a)
}

// UpdateRentableTypeToActive makes a rentabletype as active
func UpdateRentableTypeToActive(ctx context.Context, a *RentableType) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}
	fields := []interface{}{a.LastModBy, a.RTID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableTypeToActive)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableTypeToActive.Exec(fields...)
	}
	return updateError(err, "RentableType", *a)
}

// UpdateRentableTypeToInactive makes a rentabletype inactive
func UpdateRentableTypeToInactive(ctx context.Context, a *RentableType) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}
	fields := []interface{}{a.LastModBy, a.RTID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableTypeToInactive)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableTypeToInactive.Exec(fields...)
	}
	return updateError(err, "RentableType", *a)
}

// UpdateRentableTypeRef updates a RentableTypeRef record in the database
func UpdateRentableTypeRef(ctx context.Context, a *RentableTypeRef) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	//  SET BID=?,RTID=?,OverrideRentCycle=?,OverrideProrationCycle=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?"
	fields := []interface{}{a.RID, a.BID, a.RTID, a.OverrideRentCycle, a.OverrideProrationCycle, a.DtStart, a.DtStop, a.LastModBy, a.RTRID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableTypeRef)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableTypeRef.Exec(fields...)
	}
	return updateError(err, "RentableTypeRef", *a)
}

// UpdateRentableUser updates a RentableUser record in the database
func UpdateRentableUser(ctx context.Context, a *RentableUser) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.RID, a.BID, a.TCID, a.DtStart, a.DtStop, a.LastModBy, a.RUID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableUser)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableUser.Exec(fields...)
	}
	return updateError(err, "RentableUser", *a)
}

// UpdateRentableUserByRBT updates a RentableUser record in the database
func UpdateRentableUserByRBT(ctx context.Context, a *RentableUser) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.DtStart, a.DtStop, a.LastModBy, a.RID, a.BID, a.TCID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateRentableUserByRBT)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateRentableUserByRBT.Exec(fields...)
	}
	return updateError(err, "RentableUser", *a)
}

// UpdateStringList updates a StringList record in the database. It also updates the string list. It does this by
// deleting all the strings first, then inserting the ones it has.
func UpdateStringList(ctx context.Context, a *StringList) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Name, a.LastModBy, a.SLID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateStringList)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateStringList.Exec(fields...)
	}
	updateError(err, "StringList", *a)
	DeleteSLStrings(ctx, a.SLID)
	InsertSLStrings(ctx, a)
	return err
}

// UpdateSLString updates a SLString record in the database
func UpdateSLString(ctx context.Context, a *SLString) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.SLID, a.Value, a.LastModBy, a.SLSID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateSLString)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateSLString.Exec(fields...)
	}
	return updateError(err, "SLString", *a)
}

// UpdateSubAR updates a SubAR record in the database
func UpdateSubAR(ctx context.Context, a *SubAR) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.ARID, a.SubARID, a.BID, a.LastModBy, a.SARID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateSubAR)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateSubAR.Exec(fields...)
	}
	return updateError(err, "SubAR", *a)
}

//**************************************************************
// TASK, TASKLIST, TASK DESCRIPTOR, TASK LIST DEFINITION
//**************************************************************

func authProblem(ctx context.Context, uid *int64) bool {
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return true
		}
		*uid = sess.UID
	}
	return false
}

// UpdateTask updates a Task record in the database
func UpdateTask(ctx context.Context, a *Task) error {
	var err error
	if authProblem(ctx, &a.LastModBy) {
		return ErrSessionRequired
	}

	fields := []interface{}{a.BID, a.TLID, a.Name, a.Worker, a.DtDue, a.DtPreDue, a.DtDone, a.DtPreDone, a.FLAGS, a.DoneUID, a.PreDoneUID, a.Comment, a.LastModBy, a.TID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateTask)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateTask.Exec(fields...)
	}
	return updateError(err, "Task", *a)
}

// UpdateTaskList updates a TaskList record in the database
func UpdateTaskList(ctx context.Context, a *TaskList) error {
	var err error
	if authProblem(ctx, &a.LastModBy) {
		return ErrSessionRequired
	}
	fields := []interface{}{a.BID, a.PTLID, a.TLDID, a.Name, a.Cycle, a.DtDue, a.DtPreDue, a.DtDone, a.DtPreDone, a.FLAGS, a.DoneUID, a.PreDoneUID, a.EmailList, a.DtLastNotify, a.DurWait, a.Comment, a.LastModBy, a.TLID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateTaskList)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateTaskList.Exec(fields...)
	}
	return updateError(err, "TaskList", *a)
}

// UpdateTaskDescriptor updates a TaskDescriptor record in the database
func UpdateTaskDescriptor(ctx context.Context, a *TaskDescriptor) error {
	var err error
	if authProblem(ctx, &a.LastModBy) {
		return ErrSessionRequired
	}

	fields := []interface{}{a.BID, a.TLDID, a.Name, a.Worker, a.EpochDue, a.EpochPreDue, a.FLAGS, a.Comment, a.LastModBy, a.TDID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateTaskDescriptor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateTaskDescriptor.Exec(fields...)
	}
	return updateError(err, "TaskDescriptor", *a)
}

// UpdateTaskListDefinition updates a TaskListDefinition record in the database
func UpdateTaskListDefinition(ctx context.Context, a *TaskListDefinition) error {
	var err error
	if authProblem(ctx, &a.LastModBy) {
		return ErrSessionRequired
	}
	fields := []interface{}{a.BID, a.Name, a.Cycle, a.Epoch, a.EpochDue, a.EpochPreDue, a.FLAGS, a.EmailList, a.DurWait, a.Comment, a.LastModBy, a.TLDID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateTaskListDefinition)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateTaskListDefinition.Exec(fields...)
	}
	return updateError(err, "TaskListDefinition", *a)
}

//*****************************************************************************
//    TBIND
//*****************************************************************************

// UpdateTBind updates a TBind record in the database
func UpdateTBind(ctx context.Context, a *TBind) error {
	var err error
	if authProblem(ctx, &a.LastModBy) {
		return ErrSessionRequired
	}
	fields := []interface{}{
		a.BID, a.SourceElemType, a.SourceElemID, a.AssocElemType, a.AssocElemID,
		a.DtStart, a.DtStop, a.FLAGS, a.LastModBy, a.TBID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateTBind)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateTBind.Exec(fields...)
	}
	return updateError(err, "TBind", *a)
}

//*****************************************************************************
//    TRANSACTANT
//*****************************************************************************

// UpdateTransactant updates a Transactant record in the database
func UpdateTransactant(ctx context.Context, a *Transactant) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.NLID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName,
		a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone,
		a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.Comment, a.FLAGS,
		a.LastModBy, a.TCID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateTransactant)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateTransactant.Exec(fields...)
	}
	return updateError(err, "Transactant", *a)
}

// UpdateUser updates a User record in the database
func UpdateUser(ctx context.Context, a *User) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{a.BID, a.Points, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress,
		a.EmergencyContactTelephone, a.EmergencyContactEmail, a.AlternateEmailAddress, a.EligibleFutureUser, a.FLAGS,
		a.Industry, a.SourceSLSID, a.LastModBy, a.TCID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateUser)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateUser.Exec(fields...)
	}
	return updateError(err, "User", *a)
}

// UpdateVehicle updates a Vehicle record in the database
func UpdateVehicle(ctx context.Context, a *Vehicle) error {
	var err error

	if err = updateSessionProblem(ctx, &a.CreateBy, &a.LastModBy); err != nil {
		return err
	}

	fields := []interface{}{
		a.TCID,
		a.BID,
		a.VehicleType,
		a.VehicleMake,
		a.VehicleModel,
		a.VehicleColor,
		a.VIN,
		a.VehicleYear,
		a.LicensePlateState,
		a.LicensePlateNumber,
		a.ParkingPermitNumber,
		a.DtStart,
		a.DtStop,
		a.LastModBy,
		a.VID,
	}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.UpdateVehicle)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.UpdateVehicle.Exec(fields...)
	}
	return updateError(err, "Vehicle", *a)
}
