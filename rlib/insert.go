package rlib

import (
	"context"
	"database/sql"
	"extres"
)

// insertSessionProblem is a convenience function that replaces 8 lines
// of code with about 4. Since these lines are needed for every insert call
// it saves a lot of lines.  Added this routine at the time Task,TaskList,
// TaskDescriptor and  TaskListDefinition were added.
//-----------------------------------------------------------------------------
func insertSessionProblem(ctx context.Context, id *int64) error {
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
		(*id) = sess.UID
		return nil
	}
	return nil
}

func insertError(err error, n string, a interface{}) error {
	if nil != err {
		Ulog("Insert%s: error inserting %s:  %v\n", n, n, err)
		Ulog("%s = %#v\n", n, a)
	}
	return err
}

// InsertAR writes a new AR record to the database. If the record is successfully written,
// the ARID field is set to its new value.
func InsertAR(ctx context.Context, a *AR) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.ARType, a.DebitLID, a.CreditLID, a.Description, a.RARequired, a.DtStart, a.DtStop, a.FLAGS, a.DefaultAmount, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertAR)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertAR.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.ARID = rid
		}
	} else {
		err = insertError(err, "AR", *a)
	}
	return rid, err
}

// InsertAssessment writes a new assessmenttype record to the database. If the record is successfully written,
// the ASMID field is set to its new value.
func InsertAssessment(ctx context.Context, a *Assessment) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// ROUND OFF Amount upto 2 decimals
	a.Amount = Round(a.Amount, .5, 2)

	// transaction... context
	fields := []interface{}{a.PASMID, a.RPASMID, a.AGRCPTID, a.BID, a.RID, a.ATypeLID, a.RAID, a.Amount, a.Start, a.Stop, a.RentCycle, a.ProrationCycle, a.InvoiceNo, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertAssessment)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertAssessment.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.ASMID = rid
		}
	} else {
		err = insertError(err, "Insert", *a)
	}
	return rid, err
}

// InsertBuilding writes a new Building record to the database
func InsertBuilding(ctx context.Context, a *Building) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertBuilding)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertBuilding.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.BLDGID = rid
		}
	} else {
		err = insertError(err, "Building", *a)
	}
	return rid, err
}

// InsertBuildingWithID writes a new Building record to the database with the supplied bldgid
// the Building ID must be set in the supplied Building struct ptr (a.BLDGID).
func InsertBuildingWithID(ctx context.Context, a *Building) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BLDGID, a.BID, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertBuildingWithID)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertBuildingWithID.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.BLDGID = rid
		}
	} else {
		err = insertError(err, "Building", *a)
	}
	return rid, err
}

// InsertBusiness writes a new Business record.
// returns the new Business ID and any associated error
func InsertBusiness(ctx context.Context, a *Business) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context

	// TODO(Sudip): keep mind this FLAGS insertion in fields, this might be removed in the future
	fields := []interface{}{a.Designation, a.Name, a.DefaultRentCycle, a.DefaultProrationCycle, a.DefaultGSRPC, a.FLAGS, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertBusiness)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertBusiness.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.BID = rid
		}

		/*// Need to update this BUD list memory cache
		RRdb.BUDlist[a.Designation] = a.BID*/

		// build business list and cache again
		RRdb.BUDlist, RRdb.BizCache = BuildBusinessDesignationMap()
	}
	return rid, err
}

// InsertCustomAttribute writes a new User record to the database
func InsertCustomAttribute(ctx context.Context, a *CustomAttribute) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Type, a.Name, a.Value, a.Units, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertCustomAttribute)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertCustomAttribute.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.CID = rid
		}
	} else {
		err = insertError(err, "CustomAttribute", *a)
	}
	return rid, err
}

// InsertCustomAttributeRef writes a new assessmenttype record to the database
func InsertCustomAttributeRef(ctx context.Context, a *CustomAttributeRef) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.ElementType, a.BID, a.ID, a.CID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertCustomAttributeRef)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertCustomAttributeRef.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.CARID = rid
		}
	} else {
		err = insertError(err, "CustomAttributeRef", *a)
	}
	return rid, err
}

// InsertDemandSource writes a new DemandSource record to the database
func InsertDemandSource(ctx context.Context, a *DemandSource) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.Industry, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertDemandSource)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertDemandSource.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.SourceSLSID = rid
		}
	} else {
		err = insertError(err, "DemandSource", *a)
	}
	return rid, err
}

// InsertDeposit writes a new Deposit record to the database
func InsertDeposit(ctx context.Context, a *Deposit) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.DEPID, a.DPMID, a.Dt, a.Amount, a.ClearedAmount, a.FLAGS, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertDeposit)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertDeposit.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.DID = rid
		}
	} else {
		err = insertError(err, "Deposit", *a)
	}
	return rid, err
}

// InsertDepositMethod writes a new DepositMethod record to the database
func InsertDepositMethod(ctx context.Context, a *DepositMethod) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Method, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertDepositMethod)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertDepositMethod.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.DPMID = rid
		}
	} else {
		err = insertError(err, "DepositMethod", *a)
	}
	return rid, err
}

// InsertDepositPart writes a new DepositPart record to the database
func InsertDepositPart(ctx context.Context, a *DepositPart) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.DID, a.BID, a.RCPTID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertDepositPart)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertDepositPart.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.DPID = rid
		}
	} else {
		err = insertError(err, "DepositPart", *a)
	}
	return rid, err
}

// InsertDepository writes a new Depository record to the database
func InsertDepository(ctx context.Context, a *Depository) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.LID, a.Name, a.AccountNo, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertDepository)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertDepository.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.DEPID = rid
		}
	} else {
		err = insertError(err, "Depository", *a)
	}
	return rid, err
}

//======================================
//  EXPENSE
//======================================

// InsertExpense writes a new Expense record to the database
func InsertExpense(ctx context.Context, a *Expense) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	a.Amount = Round(a.Amount, .5, 2)
	// transaction... context
	fields := []interface{}{a.RPEXPID, a.BID, a.RID, a.RAID, a.Amount, a.Dt, a.AcctRule, a.ARID, a.FLAGS, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertExpense)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertExpense.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.EXPID = rid
		}
	} else {
		err = insertError(err, "Expense", *a)
	}
	return rid, err
}

//======================================
//  INVOICE
//======================================

// InsertInvoice writes a new Invoice record to the database
func InsertInvoice(ctx context.Context, a *Invoice) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Dt, a.DtDue, a.Amount, a.DeliveredBy, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertInvoice)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertInvoice.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.InvoiceNo = rid
		}
	} else {
		err = insertError(err, "Invoice", *a)
	}
	return rid, err
}

// InsertInvoiceAssessment writes a new InvoiceAssessment record to the database
func InsertInvoiceAssessment(ctx context.Context, a *InvoiceAssessment) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.InvoiceNo, a.BID, a.ASMID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertInvoiceAssessment)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertInvoiceAssessment.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.InvoiceASMID = rid
		}
	} else {
		err = insertError(err, "InvoiceAssessment", *a)
	}
	return rid, err
}

// InsertInvoicePayor writes a new InvoicePayor record to the database
func InsertInvoicePayor(ctx context.Context, a *InvoicePayor) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.InvoiceNo, a.BID, a.PID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertInvoicePayor)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertInvoicePayor.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.InvoicePayorID = rid
		}
	} else {
		err = insertError(err, "InvoicePayor", *a)
	}
	return rid, err
}

// InsertJournal writes a new Journal entry to the database
func InsertJournal(ctx context.Context, a *Journal) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Dt, a.Amount, a.Type, a.ID, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertJournal)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertJournal.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.JID = rid
		}
	} else {
		err = insertError(err, "Journal", *a)
	}
	return rid, err
}

// InsertJournalAllocationEntry writes a new JournalAllocation record to the database. Also sets JAID with its
// newly assigned id.
func InsertJournalAllocationEntry(ctx context.Context, a *JournalAllocation) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// debug.PrintStack()
	// transaction... context
	fields := []interface{}{a.BID, a.JID, a.RID, a.RAID, a.TCID, a.RCPTID, a.Amount, a.ASMID, a.EXPID, a.AcctRule, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertJournalAllocation)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertJournalAllocation.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.JAID = rid
		}
	} else {
		err = insertError(err, "JournalAllocation", *a)
	}
	return rid, err
}

// InsertJournalMarker writes a new JournalMarker record to the database
func InsertJournalMarker(ctx context.Context, a *JournalMarker) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.State, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertJournalMarker)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertJournalMarker.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.JMID = rid
		}
	} else {
		err = insertError(err, "JournalMarker", *a)
	}

	// After getting result...
	return rid, err
}

//======================================
//  LEDGER MARKER
//======================================

// InsertLedgerMarker writes a new LedgerMarker record to the database
func InsertLedgerMarker(ctx context.Context, a *LedgerMarker) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.LID, a.BID, a.RAID, a.RID, a.TCID, a.Dt, a.Balance, a.State, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertLedgerMarker)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertLedgerMarker.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			a.LMID = int64(x)
		}
	} else {
		err = insertError(err, "LedgerMarker", *a)
	}
	return rid, err
}

// InsertLedgerEntry writes a new LedgerEntry to the database
func InsertLedgerEntry(ctx context.Context, a *LedgerEntry) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.JID, a.JAID, a.LID, a.RAID, a.RID, a.TCID, a.Dt, a.Amount, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertLedgerEntry)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertLedgerEntry.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.LEID = rid
		}
	} else {
		err = insertError(err, "LedgerEntry", *a)
	}
	return rid, err
}

// InsertLedger writes a new GLAccount to the database
func InsertLedger(ctx context.Context, a *GLAccount) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	//                                            PLID, BID,     RAID,  TCID,   GLNumber,   Status,   Name,   AcctType,   AllowPost,  FLAGS,   Description, CreateBy, LastModBy
	// transaction... context
	fields := []interface{}{a.PLID, a.BID, a.RAID, a.TCID, a.GLNumber, a.Status, a.Name, a.AcctType, a.AllowPost, a.FLAGS, a.Description, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertLedger)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertLedger.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.LID = rid
		}
	} else {
		err = insertError(err, "Ledger", *a)
	}
	return rid, err
}

//======================================
// NOTE
//======================================

// InsertNote writes a new Note to the database
func InsertNote(ctx context.Context, a *Note) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.NLID, a.PNID, a.NTID, a.RID, a.RAID, a.TCID, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertNote)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertNote.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.NID = rid
		}
	} else {
		err = insertError(err, "Note", *a)
	}
	return rid, err
}

//======================================
// NOTE LIST
//======================================

// InsertNoteList inserts a new wrapper for a notelist into the database
func InsertNoteList(ctx context.Context, a *NoteList) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertNoteList)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertNoteList.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.NLID = rid
		}
	} else {
		err = insertError(err, "NoteList", *a)
	}
	return rid, err
}

//======================================
// NOTE TYPE
//======================================

// InsertNoteType writes a new NoteType to the database
func InsertNoteType(ctx context.Context, a *NoteType) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertNoteType)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertNoteType.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.NTID = rid
		}
	} else {
		err = insertError(err, "NoteType", *a)
	}
	return rid, err
}

//=======================================================
//  RATE PLAN
//=======================================================

// InsertRatePlan writes a new RatePlan record to the database
func InsertRatePlan(ctx context.Context, a *RatePlan) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRatePlan)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRatePlan.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RPID = rid
		}
	} else {
		err = insertError(err, "RatePlan", *a)
	}
	return rid, err
}

// InsertRatePlanRef writes a new RatePlanRef record to the database
func InsertRatePlanRef(ctx context.Context, a *RatePlanRef) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.RPID, a.DtStart, a.DtStop, a.FeeAppliesAge, a.MaxNoFeeUsers, a.AdditionalUserFee, a.PromoCode, a.CancellationFee, a.FLAGS, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRatePlanRef)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRatePlanRef.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RPRID = rid
		}
	} else {
		err = insertError(err, "RatePlanRef", *a)
	}
	return rid, err
}

// InsertRatePlanRefRTRate writes a new RatePlanRefRTRate record to the database
func InsertRatePlanRefRTRate(ctx context.Context, a *RatePlanRefRTRate) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RPRID, a.BID, a.RTID, a.FLAGS, a.Val, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRatePlanRefRTRate)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRatePlanRefRTRate.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RPRRTRateID = rid
		}
	} else {
		err = insertError(err, "RatePlanRefRTRate", *a)
	}
	return rid, err
}

// InsertRatePlanRefSPRate writes a new RatePlanRefSPRate record to the database
func InsertRatePlanRefSPRate(ctx context.Context, a *RatePlanRefSPRate) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RPRID, a.BID, a.RTID, a.RSPID, a.FLAGS, a.Val, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRatePlanRefSPRate)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRatePlanRefSPRate.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RPRSPRateID = rid
		}
	} else {
		err = insertError(err, "RatePlanRefSPRate", *a)
	}
	return rid, err
}

//=======================================================
//  PAYMENT
//=======================================================

// InsertPaymentType writes a new assessmenttype record to the database
func InsertPaymentType(ctx context.Context, a *PaymentType) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.Description, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertPaymentType)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertPaymentType.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.PMTID = rid
		}
	} else {
		err = insertError(err, "PaymentType", *a)
	}
	return rid, err
}

// InsertRentable writes a new Rentable record to the database
func InsertRentable(ctx context.Context, a *Rentable) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.RentableName, a.AssignmentTime, a.MRStatus, a.DtMRStart, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentable)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentable.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RID = rid
		}
	} else {
		err = insertError(err, "Rentable", *a)
	}
	return rid, err
}

//=======================================================
//  R E C E I P T
//=======================================================

// InsertReceipt writes a new Receipt record to the database. If the record is successfully written,
// the RCPTID field is set to its new value.
func InsertReceipt(ctx context.Context, a *Receipt) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	a.Amount = Round(a.Amount, .5, 2)
	// transaction... context
	fields := []interface{}{a.PRCPTID, a.BID, a.TCID, a.PMTID, a.DEPID, a.DID, a.RAID, a.Dt, a.DocNo, a.Amount, a.AcctRuleReceive, a.ARID, a.AcctRuleApply, a.FLAGS, a.Comment, a.OtherPayorName, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertReceipt)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertReceipt.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RCPTID = rid
		}
	} else {
		err = insertError(err, "Receipt", *a)
	}
	return rid, err
}

// InsertReceiptAllocation writes a new ReceiptAllocation record to the database
func InsertReceiptAllocation(ctx context.Context, a *ReceiptAllocation) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	a.Amount = Round(a.Amount, .5, 2)
	// transaction... context
	fields := []interface{}{a.RCPTID, a.BID, a.RAID, a.Dt, a.Amount, a.ASMID, a.FLAGS, a.AcctRule, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertReceiptAllocation)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertReceiptAllocation.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RCPAID = rid
		}
	} else {
		err = insertError(err, "ReceiptAllocation", *a)
	}
	return rid, err
}

// InsertRentalAgreement writes a new RentalAgreement record to the database
func InsertRentalAgreement(ctx context.Context, a *RentalAgreement) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RATID, a.BID, a.NLID, a.AgreementStart, a.AgreementStop, a.PossessionStart, a.PossessionStop, a.RentStart, a.RentStop, a.RentCycleEpoch, a.UnspecifiedAdults, a.UnspecifiedChildren, a.Renewal, a.SpecialProvisions, a.LeaseType, a.ExpenseAdjustmentType, a.ExpensesStop, a.ExpenseStopCalculation, a.BaseYearEnd, a.ExpenseAdjustment, a.EstimatedCharges, a.RateChange, a.NextRateChange, a.PermittedUses, a.ExclusiveUses, a.ExtensionOption, a.ExtensionOptionNotice, a.ExpansionOption, a.ExpansionOptionNotice, a.RightOfFirstRefusal, a.FLAGS, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentalAgreement)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentalAgreement.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RAID = rid
		}
	} else {
		err = insertError(err, "RentalAgreement", *a)
	}
	return rid, err
}

// InsertRentalAgreementPayor writes a new User record to the database
func InsertRentalAgreementPayor(ctx context.Context, a *RentalAgreementPayor) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RAID, a.BID, a.TCID, a.DtStart, a.DtStop, a.FLAGS, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentalAgreementPayor)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentalAgreementPayor.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RAPID = rid
		}
	} else {
		err = insertError(err, "RentalAgreementPayor", *a)
	}
	return rid, err
}

// InsertRentalAgreementPet writes a new User record to the database
func InsertRentalAgreementPet(ctx context.Context, a *RentalAgreementPet) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.RAID, a.Type, a.Breed, a.Color, a.Weight, a.Name, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentalAgreementPet)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentalAgreementPet.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.PETID = rid
		}
	} else {
		err = insertError(err, "RentalAgreementPet", *a)
	}
	return rid, err
}

// InsertRentalAgreementRentable writes a new User record to the database
func InsertRentalAgreementRentable(ctx context.Context, a *RentalAgreementRentable) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RAID, a.BID, a.RID, a.CLID, a.ContractRent, a.RARDtStart, a.RARDtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentalAgreementRentable)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentalAgreementRentable.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RARID = rid
		}
	} else {
		err = insertError(err, "RentalAgreementRentable", *a)
	}
	return rid, err
}

//=======================================================
//  RENTAL AGREEMENT TEMPLATE
//=======================================================

// InsertRentalAgreementTemplate writes a new User record to the database
func InsertRentalAgreementTemplate(ctx context.Context, a *RentalAgreementTemplate) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.RATemplateName, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentalAgreementTemplate)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentalAgreementTemplate.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RATID = rid
		}
	} else {
		err = insertError(err, "RentalAgreementTemplate", *a)
	}
	return rid, err
}

// InsertRentableSpecialty writes a new RentableSpecialty record to the database
func InsertRentableSpecialty(ctx context.Context, a *RentableSpecialty) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.Fee, a.Description, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableSpecialtyType)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableSpecialtyType.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RSPID = rid
		}
	} else {
		err = insertError(err, "RentableSpecialty", *a)
	}
	return rid, err
}

// InsertRentableMarketRates writes a new marketrate record to the database
func InsertRentableMarketRates(ctx context.Context, a *RentableMarketRate) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RTID, a.BID, a.MarketRate, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableMarketRates)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableMarketRates.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RMRID = rid
		}
	} else {
		err = insertError(err, "RentableMarketRate", *a)
	}
	return rid, err
}

// InsertRentableType writes a new RentableType record to the database
func InsertRentableType(ctx context.Context, a *RentableType) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Style, a.Name, a.RentCycle, a.Proration, a.GSRPC, a.ManageToBudget, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableType)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableType.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RTID = rid
		}
	} else {
		err = insertError(err, "RentableType", *a)
	}
	return rid, err
}

// InsertRentableSpecialtyRef writes a new RentableSpecialty record to the database
func InsertRentableSpecialtyRef(ctx context.Context, a *RentableSpecialtyRef) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.RID, a.RSPID, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableSpecialtyRef)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableSpecialtyRef.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RSPRefID = rid
		}
	} else {
		err = insertError(err, "RentableSpecialtyRef", *a)
	}
	return rid, err
}

// InsertRentableStatus writes a new RentableStatus record to the database
func InsertRentableStatus(ctx context.Context, a *RentableStatus) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RID, a.BID, a.DtStart, a.DtStop, a.DtNoticeToVacate, a.UseStatus, a.LeaseStatus, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableStatus)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableStatus.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RSID = rid
		}
	} else {
		err = insertError(err, "RentableStatus", *a)
	}
	return rid, err

}

// InsertRentableTypeRef writes a new RentableTypeRef record to the database
func InsertRentableTypeRef(ctx context.Context, a *RentableTypeRef) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RID, a.BID, a.RTID, a.OverrideRentCycle, a.OverrideProrationCycle, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableTypeRef)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableTypeRef.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RTRID = rid
		}
	} else {
		err = insertError(err, "RentableTypeRef", *a)
	}
	return rid, err
}

// InsertRentableUser writes a new User record to the database
func InsertRentableUser(ctx context.Context, a *RentableUser) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.RID, a.BID, a.TCID, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertRentableUser)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertRentableUser.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.RUID = rid
		}
	} else {
		err = insertError(err, "RentableUser", *a)
	}
	return rid, err
}

// InsertStringList writes a new StringList record to the database
func InsertStringList(ctx context.Context, a *StringList) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.Name, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertStringList)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertStringList.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.SLID = rid
		}
	} else {
		err = insertError(err, "StringList", *a)
	}

	// Before return, insert string list with context
	InsertSLStrings(ctx, a)

	return rid, err
}

// InsertSLStrings writes a the list of strings in a StringList to the database
// This one conducts the write operation in bulk mode
// So, if transaction being found from context then it will consider that
// Otherwise it creates new transaction and executes bulk write and commit it
// TAGS: BULK-WRITE,
func InsertSLStrings(ctx context.Context, a *StringList) (int64, error) {

	var (
		rid = int64(0)
		err error
		// res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// SPECIAL CASE
	var (
		insertStmt *sql.Stmt
		newTx      bool
		tx         *sql.Tx
		ok         bool
	)

	if tx, ok = DBTxFromContext(ctx); ok { // if transaction is supplied
		insertStmt = tx.Stmt(RRdb.Prepstmt.InsertSLString)
	} else {
		newTx = true
		tx, err = RRdb.Dbrr.Begin()
		if err != nil {
			return rid, err
		}
		insertStmt = tx.Stmt(RRdb.Prepstmt.InsertSLString)
	}
	defer insertStmt.Close()

	for i := 0; i < len(a.S); i++ {
		a.S[i].SLID = a.SLID

		// transaction... context
		fields := []interface{}{a.BID, a.SLID, a.S[i].Value, a.CreateBy, a.S[i].LastModBy}
		insertStmt.Exec(fields...)

		// After getting result...
		if nil != err {
			Ulog("Error while inserting SLString BULK-WRITE: %s\n", err.Error())
		}
	}

	if newTx { // if new transaction then commit it
		// if error then rollback
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			Ulog("Error while Committing transaction | inserting SLString BULK-WRITE: %s\n", err.Error())
			err = insertError(err, "SLStrings", *a)
			return rid, err
		}
	}

	return rid, err
}

// InsertSubAR writes a SubAR to the database
func InsertSubAR(ctx context.Context, a *SubAR) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.ARID, a.SubARID, a.BID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertSubAR)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertSubAR.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.SARID = rid
		}
	} else {
		err = insertError(err, "InsertSubAR", *a)
	}
	return rid, err
}

//*****************************************************************
// TASK, TASKLIST, TASK LIST DEFINITION, TASK LIST DESCRIPTOR
//*****************************************************************

// InsertTask writes a new Task record to the database
func InsertTask(ctx context.Context, a *Task) error {
	var id = int64(0)
	var err error
	var res sql.Result

	if err = insertSessionProblem(ctx, &a.CreateBy); err != nil {
		return err
	}
	a.LastModBy = a.CreateBy

	fields := []interface{}{a.BID, a.TLID, a.Name, a.Worker, a.DtDue, a.DtPreDue, a.DtDone, a.DtPreDone, a.FLAGS, a.DoneUID, a.PreDoneUID, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertTask)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertTask.Exec(fields...)
	}

	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			id = int64(x)
			a.TID = id
		}
	} else {
		err = insertError(err, "Task", *a)
	}
	return err
}

// InsertTaskList writes a new TaskList record to the database
func InsertTaskList(ctx context.Context, a *TaskList) error {
	var id = int64(0)
	var err error
	var res sql.Result

	if err = insertSessionProblem(ctx, &a.CreateBy); err != nil {
		return err
	}
	a.LastModBy = a.CreateBy
	fields := []interface{}{a.BID, a.Name, a.Cycle, a.DtDue, a.DtPreDue, a.DtDone, a.DtPreDone, a.FLAGS, a.DoneUID, a.PreDoneUID, a.Comment, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertTaskList)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertTaskList.Exec(fields...)
	}
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			id = int64(x)
			a.TLID = id
		}
	} else {
		err = insertError(err, "TaskList", *a)
	}
	return err
}

// InsertTaskDescriptor writes a new TaskDescriptor record to the database
func InsertTaskDescriptor(ctx context.Context, a *TaskDescriptor) error {
	var id = int64(0)
	var err error
	var res sql.Result

	if err = insertSessionProblem(ctx, &a.CreateBy); err != nil {
		return err
	}
	a.LastModBy = a.CreateBy
	fields := []interface{}{a.BID, a.TLDID, a.Name, a.Worker, a.EpochDue, a.EpochPreDue, a.FLAGS, a.LastModBy, a.TDID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertTaskDescriptor)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertTaskDescriptor.Exec(fields...)
	}
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			id = int64(x)
			a.TDID = id
		}
	} else {
		err = insertError(err, "TaskDescriptor", *a)
	}
	return err
}

// InsertTaskListDefinition writes a new TaskListDefinition record to the database
func InsertTaskListDefinition(ctx context.Context, a *TaskListDefinition) error {
	var id = int64(0)
	var err error
	var res sql.Result

	if err = insertSessionProblem(ctx, &a.CreateBy); err != nil {
		return err
	}
	a.LastModBy = a.CreateBy

	fields := []interface{}{a.BID, a.Name, a.Cycle, a.Epoch, a.EpochDue, a.EpochPreDue, a.FLAGS, a.LastModBy, a.TLDID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertTaskListDefinition)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertTaskListDefinition.Exec(fields...)
	}

	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			id = int64(x)
			a.TLDID = id
		}
	} else {
		err = insertError(err, "TaskListDefinition", *a)
	}
	return err
}

//*****************************************************************************
//  TRANSACTANT, PAYOR, USER, PROSPECT
//*****************************************************************************

// InsertTransactant writes a new Transactant record to the database
func InsertTransactant(ctx context.Context, a *Transactant) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.BID, a.NLID, a.FirstName, a.MiddleName, a.LastName, a.PreferredName, a.CompanyName, a.IsCompany, a.PrimaryEmail, a.SecondaryEmail, a.WorkPhone, a.CellPhone, a.Address, a.Address2, a.City, a.State, a.PostalCode, a.Country, a.Website, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertTransactant)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertTransactant.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.TCID = rid
		}
	} else {
		err = insertError(err, "Transactant", *a)
	}
	return rid, err
}

// InsertPayor writes a new User record to the database
func InsertPayor(ctx context.Context, a *Payor) (int64, error) {

	var (
		rid = int64(0)
		err error
		// res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.TCID, a.BID, a.CreditLimit, a.TaxpayorID, a.AccountRep, a.EligibleFuturePayor, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertPayor)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.InsertPayor.Exec(fields...)
	}

	// After getting result...
	/*if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.PayorID = rid
		}
	} else {
		err = insertError(err, "Payor", *a)
	}*/
	if err != nil {
		err = insertError(err, "Payor", *a)
	}
	return rid, err
}

// InsertProspect writes a new User record to the database
func InsertProspect(ctx context.Context, a *Prospect) (int64, error) {

	var (
		rid = int64(0)
		err error
		// res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.TCID, a.BID, a.EmployerName, a.EmployerStreetAddress, a.EmployerCity,
		a.EmployerState, a.EmployerPostalCode, a.EmployerEmail, a.EmployerPhone, a.Occupation, a.ApplicationFee,
		a.DesiredUsageStartDate, a.RentableTypePreference, a.FLAGS, a.Approver, a.DeclineReasonSLSID, a.OtherPreferences,
		a.FollowUpDate, a.CSAgent, a.OutcomeSLSID, a.FloatingDeposit, a.RAID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertProspect)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.InsertProspect.Exec(fields...)
	}

	// After getting result...
	/*if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.ProspectID = rid
		}
	} else {
		err = insertError(err, "Prospect", *a)
	}*/
	if err != nil {
		err = insertError(err, "Prospect", *a)
	}
	return rid, err
}

// InsertUser writes a new User record to the database
func InsertUser(ctx context.Context, a *User) (int64, error) {

	var (
		rid = int64(0)
		err error
		// res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.TCID, a.BID, a.Points, a.DateofBirth, a.EmergencyContactName, a.EmergencyContactAddress, a.EmergencyContactTelephone, a.EmergencyEmail, a.AlternateAddress, a.EligibleFutureUser, a.Industry, a.SourceSLSID, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertUser)
		defer stmt.Close()
		_, err = stmt.Exec(fields...)
	} else {
		_, err = RRdb.Prepstmt.InsertUser.Exec(fields...)
	}

	// After getting result...
	/*if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.UserID = rid
		}
	} else {
		err = insertError(err, "User", *a)
	}*/
	if err != nil {
		err = insertError(err, "User", *a)
	}
	return rid, err
}

// InsertVehicle writes a new Vehicle record to the database
func InsertVehicle(ctx context.Context, a *Vehicle) (int64, error) {

	var (
		rid = int64(0)
		err error
		res sql.Result
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return rid, ErrSessionRequired
		}

		// user from session, CreateBy, LastModBy
		a.CreateBy = sess.UID
		a.LastModBy = a.CreateBy
	}

	// transaction... context
	fields := []interface{}{a.TCID, a.BID, a.VehicleType, a.VehicleMake, a.VehicleModel, a.VehicleColor, a.VehicleYear, a.LicensePlateState, a.LicensePlateNumber, a.ParkingPermitNumber, a.DtStart, a.DtStop, a.CreateBy, a.LastModBy}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.InsertVehicle)
		defer stmt.Close()
		res, err = stmt.Exec(fields...)
	} else {
		res, err = RRdb.Prepstmt.InsertVehicle.Exec(fields...)
	}

	// After getting result...
	if nil == err {
		x, err := res.LastInsertId()
		if err == nil {
			rid = int64(x)
			a.VID = rid
		}
	} else {
		err = insertError(err, "Vehicle", *a)
	}
	return rid, err
}
