package rlib

import (
	"context"
	"database/sql"
	"extres"
	"fmt"
	"strconv"
	"time"
)

// sessionCheck encapsulates 6 lines of code that was repeated in every call
//
// INPUTS
//  ctx  the context, which should have session
//
// RETURNS
//  true - session was required but not found
//  false - session was found or session not required
//-----------------------------------------------------------------------------
func sessionCheck(ctx context.Context) bool {
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		return !ok
	}
	return false
}

// GetCountByTableName returns the count of records in table t
// that belong to business bid
//------------------------------------------------------------------
func GetCountByTableName(ctx context.Context, t string, bid int64) (int, error) {

	var (
		err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	q := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE BID=%d", t, bid)
	row := RRdb.Dbrr.QueryRow(q)
	err = row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, err
}

//=======================================================
//  AR
//=======================================================

// GetAR reads a AR the structure for the supplied id
func GetAR(ctx context.Context, id int64) (AR, error) {

	var (
		// err error
		a AR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAR)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetAR.QueryRow(fields...)
	}
	return a, ReadAR(row, &a)
}

// GetARByName reads a AR the structure for the supplied bid and name
func GetARByName(ctx context.Context, bid int64, name string) (AR, error) {

	var (
		// err error
		a AR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetARByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetARByName.QueryRow(fields...)
	}
	return a, ReadAR(row, &a)
}

// getARsForRows uses the supplied rows param, gets all the AR records
// and returns them in a slice of AR structs
func getARsForRows(ctx context.Context, rows *sql.Rows) ([]AR, error) {

	var (
		err error
		t   []AR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var a AR
		err = ReadARs(rows, &a)
		if err != nil {
			return t, err
		}
		t = append(t, a)
	}

	return t, rows.Err()
}

// getARMap returns a map of all account rules for the supplied bid
func getARMap(bid int64) (map[int64]AR, error) {

	var (
		err error
		t   = make(map[int64]AR)
	)

	var rows *sql.Rows
	fields := []interface{}{bid}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllARs)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllARs.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetAllARs.Query(fields...)

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var a AR
		err = ReadARs(rows, &a)
		if err != nil {
			return t, err
		}
		t[a.ARID] = a
	}

	return t, rows.Err()
}

// GetARMap returns a map of all account rules for the supplied bid
func GetARMap(ctx context.Context, bid int64) (map[int64]AR, error) {

	var (
		// err error
		t = make(map[int64]AR)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	return getARMap(bid)
}

// GetAllARs reads all AccountRules for the supplied BID
func GetAllARs(ctx context.Context, BID int64) ([]AR, error) {

	var (
		err error
		t   []AR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{BID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllARs)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllARs.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getARsForRows(ctx, rows)
}

// GetARsByType reads all AccountRules for the supplied BID of type artype
func GetARsByType(ctx context.Context, bid int64, artype int) ([]AR, error) {

	var (
		err error
		t   []AR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, artype}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetARsByType)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetARsByType.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getARsForRows(ctx, rows)
}

// GetARsByFLAGS reads all AccountRules for the supplied BID with FLAGS
func GetARsByFLAGS(ctx context.Context, bid int64, FLAGS uint64) ([]AR, error) {

	var (
		err error
		t   []AR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, FLAGS, FLAGS}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetARsByFLAGS)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetARsByFLAGS.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getARsForRows(ctx, rows)
}

//=======================================================
//  A G R E E M E N T   P E T S
//=======================================================

// GetPetsByTransactant reads all Pet records for the supplied TCID
func GetPetsByTransactant(ctx context.Context, TCID int64) ([]RentalAgreementPet, error) {

	var (
		err error
		t   []RentalAgreementPet
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{TCID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPetsByTransactant)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetPetsByTransactant.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var a RentalAgreementPet
		err = ReadRentalAgreementPets(rows, &a)
		if err != nil {
			return t, err
		}
		t = append(t, a)
	}

	return t, rows.Err()
}

// GetRentalAgreementPet reads a Pet the structure for the supplied PETID
func GetRentalAgreementPet(ctx context.Context, petid int64) (RentalAgreementPet, error) {

	var (
		// err error
		a RentalAgreementPet
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{petid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementPet)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreementPet.QueryRow(fields...)
	}
	return a, ReadRentalAgreementPet(row, &a)
}

// GetAllRentalAgreementPets reads all Pet records for the supplied rental agreement id
func GetAllRentalAgreementPets(ctx context.Context, raid int64) ([]RentalAgreementPet, error) {

	var (
		err error
		t   []RentalAgreementPet
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRentalAgreementPets)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRentalAgreementPets.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var a RentalAgreementPet
		err = ReadRentalAgreementPets(rows, &a)
		if err != nil {
			return t, err
		}
		t = append(t, a)
	}

	return t, rows.Err()
}

//=======================================================
//  A G R E E M E N T   R E N T A B L E
//=======================================================

// FindAgreementByRentable reads a Prospect structure based on the supplied Transactant id
func FindAgreementByRentable(ctx context.Context, rid int64, d1, d2 *time.Time) (RentalAgreementRentable, error) {

	var (
		// err error
		a RentalAgreementRentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	// SELECT RAID,BID,RID,DtStart,DtStop from RentalAgreementRentables where RID=? and DtStop>=? and DtStart<=?

	var row *sql.Row
	fields := []interface{}{rid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.FindAgreementByRentable)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.FindAgreementByRentable.QueryRow(fields...)
	}
	return a, ReadRentalAgreementRentable(row, &a)
}

//=======================================================
//  A S S E S S M E N T S
//=======================================================

// GetAllRentableAssessments for the supplied RID and date range
func GetAllRentableAssessments(ctx context.Context, RID int64, d1, d2 *time.Time) ([]Assessment, error) {

	var (
		err error
		t   []Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{RID, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRentableAssessments)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRentableAssessments.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getAssessmentsByRows(ctx, rows)
}

// GetUnpaidAssessmentsByRAID for the supplied RAID
func GetUnpaidAssessmentsByRAID(ctx context.Context, RAID int64) ([]Assessment, error) {

	var (
		err error
		t   []Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{RAID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetUnpaidAssessmentsByRAID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetUnpaidAssessmentsByRAID.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getAssessmentsByRows(ctx, rows)
}

// GetEpochAssessmentsByRentalAgreement for the supplied RAID
// INPUTS
// ctx  - context
// RAID - Rental Agreement id of interest
//
// RETURNS
//    array of recurring assessment definitions and non-recurring single instances
//    which are not part of a recurring series
//-----------------------------------------------------------------------------
func GetEpochAssessmentsByRentalAgreement(ctx context.Context, RAID int64) ([]Assessment, error) {

	var (
		err error
		t   []Assessment
	)
	if sessionCheck(ctx) {
		return t, ErrSessionRequired
	}

	var rows *sql.Rows
	fields := []interface{}{RAID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetEpochAssessmentsByRentalAgreement)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetEpochAssessmentsByRentalAgreement.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getAssessmentsByRows(ctx, rows)
}

// GetAssessmentsByRAIDRID gets all the Assessments associated with the
// supplied BID/RAID/RID combination.  It returns the epoch instance of
// recurring assessments. It will not return individual instances unless
// they are epoch instances (non recurring assessments)
//
// INPUTS
//    ctx  - context
//    bid  - the biz
//    raid - Rental Agreement id of interest
//    rid  - the rentable of interest
//
// RETURNS
//    array of assessments suitable for an assessment list in a RentalAgreement
//-----------------------------------------------------------------------------
func GetAssessmentsByRAIDRID(ctx context.Context, bid, raid, rid int64) ([]Assessment, error) {
	var err error
	var t []Assessment
	if sessionCheck(ctx) {
		return t, ErrSessionRequired
	}

	var rows *sql.Rows
	fields := []interface{}{bid, raid, rid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAssessmentsByRAIDRID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAssessmentsByRAIDRID.Query(fields...)
	}
	if err != nil {
		return t, err
	}
	return getAssessmentsByRows(ctx, rows)
}

// GetAssessmentInstancesByParent for the supplied RAID
// INPUTS
//    id - id of Parent Assessment
// d1-d2 - date range for search
//
// RETURNS
//    array of matching assessments
//-----------------------------------------------------------------------------
func GetAssessmentInstancesByParent(ctx context.Context, id int64, d1, d2 *time.Time) ([]Assessment, error) {

	var (
		err error
		t   []Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAssessmentInstancesByParent)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAssessmentInstancesByParent.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getAssessmentsByRows(ctx, rows)
}

// getAssessmentsByRows for the supplied sql.Rows
func getAssessmentsByRows(ctx context.Context, rows *sql.Rows) ([]Assessment, error) {

	var (
		err error
		t   []Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var a Assessment
		err = ReadAssessments(rows, &a)
		if err != nil {
			return t, err
		}
		t = append(t, a)
	}

	return t, rows.Err()
}

// GetAssessment returns the Assessment struct for the account with the supplied asmid
func GetAssessment(ctx context.Context, asmid int64) (Assessment, error) {

	var (
		// err error
		a Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{asmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAssessment)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetAssessment.QueryRow(fields...)
	}
	return a, ReadAssessment(row, &a)
}

// GetAssessmentInstance returns the Assessment struct for the account with the supplied asmid
func GetAssessmentInstance(ctx context.Context, start *time.Time, pasmid int64) (Assessment, error) {

	var (
		// err error
		a Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{start, pasmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAssessmentInstance)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetAssessmentInstance.QueryRow(fields...)
	}
	return a, ReadAssessment(row, &a)
}

// GetAssessmentFirstInstance returns the Assessment struct for the first instance of the
// recurring series with PASMID = pasmid
func GetAssessmentFirstInstance(ctx context.Context, pasmid int64) (Assessment, error) {

	var (
		// err error
		a Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{pasmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAssessmentFirstInstance)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetAssessmentFirstInstance.QueryRow(fields...)
	}
	return a, ReadAssessment(row, &a)
}

// GetAssessmentDuplicate returns the Assessment struct for the account with the supplied asmid
func GetAssessmentDuplicate(ctx context.Context, start *time.Time, amt float64, pasmid, rid, raid, atypelid int64) (Assessment, error) {

	var (
		// err error
		a Assessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{start, amt, pasmid, rid, raid, atypelid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAssessmentDuplicate)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetAssessmentDuplicate.QueryRow(fields...)
	}
	return a, ReadAssessment(row, &a)
}

//=======================================================
//  B U I L D I N G
//=======================================================

// GetBuilding returns the record for supplied bldg id. If no such record exists or a database error occurred,
// the return structure will be empty
func GetBuilding(ctx context.Context, id int64) (Building, error) {

	var (
		// err error
		t Building
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetBuilding)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetBuilding.QueryRow(fields...)
	}
	return t, ReadBuildingData(row, &t)
}

//=======================================================
//  B U S I N E S S
//=======================================================

// GetAllBiz generates a slice of all Businesses defined in the database
// without authentication
func GetAllBiz() ([]Business, error) {

	var (
		err error
		m   []Business
	)

	var rows *sql.Rows
	fields := []interface{}{}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllBusinesses)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllBusinesses.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetAllBusinesses.Query(fields...)

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Business
		err = ReadBusinesses(rows, &p)
		if err != nil {
			return m, err
		}
		m = append(m, p)
	}

	return m, rows.Err()
}

// GetAllBusinesses generates a report of all Businesses defined in the database.
func GetAllBusinesses(ctx context.Context) ([]Business, error) {

	var (
		// err error
		m []Business
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	return GetAllBiz()
}

// getBiz loads the Business struct for the supplied Business id
func getBiz(bid int64, a *Business) error {
	var row *sql.Row
	fields := []interface{}{bid}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetBusiness)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetBusiness.QueryRow(fields...)
	}*/
	row = RRdb.Prepstmt.GetBusiness.QueryRow(fields...)
	return ReadBusiness(row, a)
}

// GetBusiness loads the Business struct for the supplied Business id
func GetBusiness(ctx context.Context, bid int64, a *Business) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	return getBiz(bid, a)
}

// GetBizByDesignation loads the Business struct for the supplied designation
func GetBizByDesignation(des string) (Business, error) {

	var (
		// err error
		a Business
	)

	var row *sql.Row
	fields := []interface{}{des}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetBusinessByDesignation)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetBusinessByDesignation.QueryRow(fields...)
	}*/
	row = RRdb.Prepstmt.GetBusinessByDesignation.QueryRow(fields...)
	return a, ReadBusiness(row, &a)
}

// GetBusinessByDesignation loads the Business struct for the supplied designation
func GetBusinessByDesignation(ctx context.Context, des string) (Business, error) {

	var (
		// err error
		a Business
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	return GetBizByDesignation(des)
}

// GetXBiz loads the XBusiness struct for the supplied Business id.
func GetXBiz(bid int64, xbiz *XBusiness) error {

	var (
		err error
	)

	if xbiz.P.BID == 0 && bid > 0 {
		err = getBiz(bid, &xbiz.P)
		if err != nil {
			return err
		}
	}

	xbiz.RT, err = getBizRentableTypes(bid)
	if err != nil {
		return err
	}

	xbiz.US = make(map[int64]RentableSpecialty)

	var rows *sql.Rows
	fields := []interface{}{bid}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllBusinessRentableSpecialtyTypes)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllBusinessRentableSpecialtyTypes.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetAllBusinessRentableSpecialtyTypes.Query(fields...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var a RentableSpecialty
		err = ReadRentableSpecialties(rows, &a)
		if err != nil {
			return err
		}
		xbiz.US[a.RSPID] = a
	}

	return rows.Err()
}

// GetXBusiness loads the XBusiness struct for the supplied Business id.
func GetXBusiness(ctx context.Context, bid int64, xbiz *XBusiness) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	return GetXBiz(bid, xbiz)
}

//=======================================================
//  CLOSE PERIOD
//=======================================================

// GetClosePeriod reads specific ClosePeriod record
//-----------------------------------------------------------------------------
func GetClosePeriod(ctx context.Context, id int64) (ClosePeriod, error) {
	var a ClosePeriod
	var row *sql.Row

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetClosePeriod)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetClosePeriod.QueryRow(fields...)
	}
	return a, ReadClosePeriod(row, &a)
}

// GetLastClosePeriod reads the last period closed
//
// INPUTS
//  id  = BID
//-----------------------------------------------------------------------------
func GetLastClosePeriod(ctx context.Context, id int64) (ClosePeriod, error) {
	var a ClosePeriod
	var row *sql.Row

	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLastClosePeriod)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLastClosePeriod.QueryRow(fields...)
	}
	return a, ReadClosePeriod(row, &a)
}

//=======================================================
//  C U S T O M   A T T R I B U T E
//  CustomAttribute, CustomAttributeRef
//=======================================================

// getCustomAttribute reads a CustomAttribute structure based on the supplied CustomAttribute id
func getCustomAttribute(id int64) (CustomAttribute, error) {

	var (
		a CustomAttribute
	)

	var row *sql.Row
	fields := []interface{}{id}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetCustomAttribute)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetCustomAttribute.QueryRow(fields...)
	}*/
	row = RRdb.Prepstmt.GetCustomAttribute.QueryRow(fields...)
	return a, ReadCustomAttribute(row, &a)
}

// GetCustomAttribute reads a CustomAttribute structure based on the supplied CustomAttribute id
func GetCustomAttribute(ctx context.Context, id int64) (CustomAttribute, error) {

	var (
		// err error
		a CustomAttribute
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	return getCustomAttribute(id)
}

// GetCustomAttributeByVals reads a CustomAttribute structure based on the supplied attributes
// t = data type (CUSTSTRING, CUSTINT, CUSTUINT, CUSTFLOAT, CUSTDATE
// n = name of this custom attribute
// v = the value of this attribute
// u = units (if not applicable then "")
func GetCustomAttributeByVals(ctx context.Context, t int64, n, v, u string) (CustomAttribute, error) {

	var (
		// err error
		a CustomAttribute
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{t, n, v, u}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetCustomAttributeByVals)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetCustomAttributeByVals.QueryRow(fields...)
	}
	return a, ReadCustomAttribute(row, &a)
}

// getAllCustomAttributes returns a list of CustomAttributes for the supplied elementid and instanceid
func getAllCustomAttributes(elemid, id int64) (map[string]CustomAttribute, error) {

	var (
		err error
		t   []int64
		m   = make(map[string]CustomAttribute)
	)

	var rows *sql.Rows
	fields := []interface{}{elemid, id}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetCustomAttributeRefs)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetCustomAttributeRefs.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetCustomAttributeRefs.Query(fields...)

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var cid int64
		err = rows.Scan(&cid)
		if err != nil {
			return m, err
		}
		t = append(t, cid)
	}

	err = rows.Err()
	if err != nil {
		return m, err
	}

	for i := 0; i < len(t); i++ {
		c, err := getCustomAttribute(t[i])
		if err != nil {
			return m, err
		}
		m[c.Name] = c
	}

	return m, err
}

// GetAllCustomAttributes returns a list of CustomAttributes for the supplied elementid and instanceid
func GetAllCustomAttributes(ctx context.Context, elemid, id int64) (map[string]CustomAttribute, error) {

	var (
		// err error
		m = make(map[string]CustomAttribute)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	return getAllCustomAttributes(elemid, id)
}

// GetCustomAttributeRef reads a CustomAttribute structure for the supplied ElementType, ID, and CID
func GetCustomAttributeRef(ctx context.Context, e, i, c int64) (CustomAttributeRef, error) {

	var (
		// err error
		a CustomAttributeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{e, i, c}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetCustomAttributeRef)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetCustomAttributeRef.QueryRow(fields...)
	}
	return a, ReadCustomAttributeRef(row, &a)
}

// loadRentableTypeCustomaAttributes adds all the custom attributes to each RentableType
func loadRentableTypeCustomaAttributes(xbiz *XBusiness) error {

	var (
		err error
	)

	for k, v := range xbiz.RT {
		var tmp = xbiz.RT[k]
		tmp.CA, err = getAllCustomAttributes(ELEMRENTABLETYPE, v.RTID)
		if err != nil {
			Ulog("LoadRentableTypeCustomaAttributes: error reading custom attributes elementtype=%d, id=%d, err = %s\n", ELEMRENTABLETYPE, v.RTID, err.Error())
		}
		xbiz.RT[k] = tmp // this workaround (assigning to tmp) instead of just directly assigning the .CA member is a known issue in go
	}

	return err
}

//=======================================================
//  DEMAND SOURCE
//=======================================================

// GetDemandSource reads a DemandSource structure based on the supplied DemandSource id
func GetDemandSource(ctx context.Context, id int64, t *DemandSource) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDemandSource)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDemandSource.QueryRow(fields...)
	}

	return ReadDemandSource(row, t)
}

// GetDemandSourceByName reads a DemandSource structure based on the supplied DemandSource id
func GetDemandSourceByName(ctx context.Context, bid int64, name string, t *DemandSource) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDemandSourceByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDemandSourceByName.QueryRow(fields...)
	}

	return ReadDemandSource(row, t)
}

// GetAllDemandSources returns an array of DemandSource structures containing all sources for the supplied BID
func GetAllDemandSources(ctx context.Context, id int64) ([]DemandSource, error) {

	var (
		err error
		m   []DemandSource
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllDemandSources)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllDemandSources.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var s DemandSource
		err = ReadDemandSources(rows, &s)
		if err != nil {
			return m, err
		}
		m = append(m, s)
	}

	return m, rows.Err()
}

//=======================================================
//  DEPOSIT
//  Deposit, Depository, Deposit Method, DepositPart
//=======================================================

// GetDeposit reads a Deposit structure based on the supplied Deposit id
func GetDeposit(ctx context.Context, id int64) (Deposit, error) {

	var (
		// err error
		a Deposit
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDeposit)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDeposit.QueryRow(fields...)
	}
	return a, ReadDeposit(row, &a)
}

// GetAllDepositsInRange returns an array of all Deposits for bid between the supplied dates
func GetAllDepositsInRange(ctx context.Context, bid int64, d1, d2 *time.Time) ([]Deposit, error) {

	var (
		err error
		t   []Deposit
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllDepositsInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllDepositsInRange.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Deposit
		err = ReadDeposits(rows, &a)
		if err != nil {
			return t, err
		}
		a.DP, err = GetDepositParts(ctx, a.DID)
		if err != nil {
			return t, err
		}
		//Console("GetAllDepositsInRange: a.DID = %d, len(a.DP) =  %d\n", a.DID, len(a.DP))
		t = append(t, a)
	}

	return t, rows.Err()
}

// GetDepository reads a Depository structure based on the supplied Depository id
func GetDepository(ctx context.Context, id int64) (Depository, error) {

	var (
		// err error
		a Depository
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepository)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDepository.QueryRow(fields...)
	}
	return a, ReadDepository(row, &a)
}

// GetDepositoryByAccount reads a Depository structure based on the supplied Account id
func GetDepositoryByAccount(ctx context.Context, bid int64, acct string) (Depository, error) {

	var (
		// err error
		a Depository
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, acct}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepositoryByAccount)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDepositoryByAccount.QueryRow(fields...)
	}
	return a, ReadDepository(row, &a)
}

// GetDepositoryByName reads a Depository structure based on the supplied Name id
func GetDepositoryByName(ctx context.Context, bid int64, name string) (Depository, error) {

	var (
		// err error
		a Depository
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepositoryByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDepositoryByName.QueryRow(fields...)
	}
	return a, ReadDepository(row, &a)
}

// GetDepositoryByLID reads a Depository structure based on the supplied LID id
func GetDepositoryByLID(ctx context.Context, bid int64, id int64) (Depository, error) {

	var (
		// err error
		a Depository
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepositoryByLID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDepositoryByLID.QueryRow(fields...)
	}
	return a, ReadDepository(row, &a)
}

// GetAllDepositories returns an array of all Depositories for the supplied business
func GetAllDepositories(ctx context.Context, bid int64) ([]Depository, error) {

	var (
		err error
		t   []Depository
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllDepositories)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllDepositories.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Depository
		err = ReadDepositories(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetDepositParts reads a DepositPart structure based on the supplied DepositPart DID
func GetDepositParts(ctx context.Context, id int64) ([]DepositPart, error) {

	var (
		err error
		m   []DepositPart
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepositParts)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetDepositParts.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a DepositPart
		err = ReadDepositParts(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetDepositMethod reads a DepositMethod structure based on the supplied Deposit id
func GetDepositMethod(ctx context.Context, id int64) (DepositMethod, error) {

	var (
		// err error
		a DepositMethod
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepositMethod)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDepositMethod.QueryRow(fields...)
	}
	return a, ReadDepositMethod(row, &a)
}

// GetDepositMethodByName reads a DepositMethod structure based on the supplied BID and Name
func GetDepositMethodByName(ctx context.Context, bid int64, name string) (DepositMethod, error) {

	var (
		// err error
		a DepositMethod
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetDepositMethodByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetDepositMethodByName.QueryRow(fields...)
	}
	return a, ReadDepositMethod(row, &a)
}

// GetAllDepositMethods returns an array of all DepositMethods for the supplied business
func GetAllDepositMethods(ctx context.Context, bid int64) ([]DepositMethod, error) {

	var (
		err error
		t   []DepositMethod
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllDepositMethods)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllDepositMethods.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var a DepositMethod
		err = ReadDepositMethods(rows, &a)
		if err != nil {
			return t, err
		}
		t = append(t, a)
	}

	return t, rows.Err()
}

//=======================================================
//  EXPENSE
//=======================================================

// GetExpense reads a Expense structure based on the supplied Expense id
func GetExpense(ctx context.Context, id int64) (Expense, error) {

	var (
		// err error
		a Expense
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetExpense)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetExpense.QueryRow(fields...)
	}
	return a, ReadExpense(row, &a)
}

//=======================================================
//  I N V O I C E
//=======================================================

// GetInvoice reads a Invoice structure based on the supplied Invoice id
func GetInvoice(ctx context.Context, id int64) (Invoice, error) {

	var (
		err error
		a   Invoice
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetInvoice)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetInvoice.QueryRow(fields...)
	}
	err = ReadInvoice(row, &a)
	if err != nil {
		return a, err
	}

	a.A, err = GetInvoiceAssessments(ctx, id)
	if err != nil {
		return a, err
	}

	a.P, err = GetInvoicePayors(ctx, id)
	if err != nil {
		return a, err
	}

	return a, err
}

// GetAllInvoicesInRange returns an array of all Invoices for bid between the supplied dates
func GetAllInvoicesInRange(ctx context.Context, bid int64, d1, d2 *time.Time) ([]Invoice, error) {

	var (
		err error
		t   []Invoice
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllInvoicesInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllInvoicesInRange.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Invoice
		err = ReadInvoices(rows, &a)
		if err != nil {
			return t, err
		}

		a.A, err = GetInvoiceAssessments(ctx, a.InvoiceNo)
		if err != nil {
			return t, err
		}

		a.P, err = GetInvoicePayors(ctx, a.InvoiceNo)
		t = append(t, a)
		if err != nil {
			return t, err
		}
	}

	return t, rows.Err()
}

// GetInvoiceAssessments reads a InvoiceAssessment structure based on the supplied InvoiceAssessment DID
func GetInvoiceAssessments(ctx context.Context, id int64) ([]InvoiceAssessment, error) {

	var (
		err error
		m   []InvoiceAssessment
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetInvoiceAssessments)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetInvoiceAssessments.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a InvoiceAssessment
		err = ReadInvoiceAssessments(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetInvoicePayors reads an InvoicePayor structure based on the supplied InvoiceNo (id)
func GetInvoicePayors(ctx context.Context, id int64) ([]InvoicePayor, error) {

	var (
		err error
		m   []InvoicePayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetInvoicePayors)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetInvoicePayors.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a InvoicePayor
		err = ReadInvoicePayors(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

//=======================================================
//  JOURNAL
//=======================================================

// GetJournal returns the Journal struct for the journal entry with the supplied id
func GetJournal(ctx context.Context, jid int64) (Journal, error) {

	var (
		// err error
		r Journal
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{jid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournal)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetJournal.QueryRow(fields...)
	}
	return r, ReadJournal(row, &r)
}

/*// GetJournalInstance returns the Journal struct for entries that were created with the assumption that
// they are idempotent -- essentially: instances of recurring assessments and vacancy instances.  This call
// is made prior to generating new ones to ensure that we don't have double entries for the same thing.
func GetJournalInstance(ctx context.Context, id int64, dt1, dt2 *time.Time) (Journal, error) {
	var (
		// err error
		r Journal
	)
	var row *sql.Row
	fields := []interface{}{id, dt1, dt2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalInstance)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetJournalInstance.QueryRow(fields...)
	}
	return r, ReadJournal(row, &r)
}*/

// GetJournalVacancy returns the Journal struct for entries that were created with the assumption that
// they are idempotent -- essentially: instances of recurring assessments and vacancy instances.  This call
// is made prior to generating new ones to ensure that we don't have double entries for the same thing.
func GetJournalVacancy(ctx context.Context, id int64, dt1, dt2 *time.Time) (Journal, error) {

	var (
		// err error
		r Journal
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id, dt1, dt2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalVacancy)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetJournalVacancy.QueryRow(fields...)
	}
	return r, ReadJournal(row, &r)
}

// GetJournalByTypeAndID returns the Journal struct for entries match the supplied
// Type and ID fields
func GetJournalByTypeAndID(ctx context.Context, t, id int64) (Journal, error) {

	var (
		// err error
		r Journal
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{t, id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalByTypeAndID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetJournalByTypeAndID.QueryRow(fields...)
	}
	return r, ReadJournal(row, &r)
}

// GetJournalByReceiptID returns the Journal struct for a Journal Entry that references the supplied
// receiptID
func GetJournalByReceiptID(ctx context.Context, id int64) (Journal, error) {

	var (
		// err error
		r Journal
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalByReceiptID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetJournalByReceiptID.QueryRow(fields...)
	}
	return r, ReadJournal(row, &r)
}

// GetJournalsByReceiptID returns a slice of Journal structs where it references the supplied
// receiptID
func GetJournalsByReceiptID(ctx context.Context, id int64) ([]Journal, error) {

	var (
		err error
		t   []Journal
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalByReceiptID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetJournalByReceiptID.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Journal
		err = ReadJournals(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetJournalMarkers loads the last n Journal markers
func GetJournalMarkers(ctx context.Context, n int64) ([]JournalMarker, error) {

	var (
		err error
		t   []JournalMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{n}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalMarkers)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetJournalMarkers.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r JournalMarker
		err = ReadJournalMarkers(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetLastJournalMarker returns the last Journal marker or nil if no Journal markers exist
func GetLastJournalMarker(ctx context.Context) (JournalMarker, error) {

	var (
		err error
		j   JournalMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return j, ErrSessionRequired
		}
	}

	t, err := GetJournalMarkers(ctx, 1)
	if err != nil {
		return j, err
	}

	if len(t) > 0 {
		j = t[0]
	}

	return j, err
}

//=======================================================
//  JOURNAL ALLOCATION
//=======================================================

// GetJournalAllocation returns the Journal allocation for the supplied JAID
func GetJournalAllocation(ctx context.Context, jaid int64) (JournalAllocation, error) {

	var (
		// err error
		a JournalAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{jaid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalAllocation)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetJournalAllocation.QueryRow(fields...)
	}
	return a, ReadJournalAllocation(row, &a)
}

// GetJournalAllocations loads all Journal allocations associated with the supplied Journal id into
// the RA array within a Journal structure
func GetJournalAllocations(ctx context.Context, j *Journal) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{j.JID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalAllocations)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetJournalAllocations.Query(fields...)
	}

	if err != nil {
		return err
	}

	// get all journal allocation rows in j.JA field
	jaRows, err := getJournalAllocationRows(ctx, rows)
	if err != nil {
		return err
	}

	j.JA = jaRows
	return err
}

// GetJournalAllocationByASMID returns an array of JournalAllocation records that reference
// the supplied ASMID.
func GetJournalAllocationByASMID(ctx context.Context, id int64) ([]JournalAllocation, error) {

	var (
		err    error
		jaRows []JournalAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return jaRows, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalAllocationsByASMID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetJournalAllocationsByASMID.Query(fields...)
	}

	if err != nil {
		return jaRows, err
	}
	return getJournalAllocationRows(ctx, rows)
}

// GetJournalAllocationByASMandRCPTID returns an array of JournalAllocation
// records that reference the supplied RCPTID and that have a non-zero ASMID.
// These are the JournalAllocation entries created for Receipts that had
// SubARs automatically generate an associated Assessment.
//----------------------------------------------------------------------------
func GetJournalAllocationByASMandRCPTID(ctx context.Context, id int64) ([]JournalAllocation, error) {

	var (
		err    error
		jaRows []JournalAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return jaRows, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetJournalAllocationsByASMandRCPTID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetJournalAllocationsByASMandRCPTID.Query(fields...)
	}

	if err != nil {
		return jaRows, err
	}
	return getJournalAllocationRows(ctx, rows)
}

func getJournalAllocationRows(ctx context.Context, rows *sql.Rows) ([]JournalAllocation, error) {

	var (
		err error
		ja  []JournalAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ja, ErrSessionRequired
		}
	}

	defer rows.Close()

	for rows.Next() {
		var a JournalAllocation
		err = ReadJournalAllocations(rows, &a)
		if err != nil {
			return ja, err
		}
		ja = append(ja, a)
	}

	return ja, rows.Err()
}

//=======================================================
//  L E D G E R   M A R K E R
//=======================================================

// GetLatestLedgerMarkerByLID returns the latest LedgerMarker struct for the GLAccount with the supplied LID
func GetLatestLedgerMarkerByLID(ctx context.Context, bid, lid int64) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, lid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLatestLedgerMarkerByLID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLatestLedgerMarkerByLID.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetInitialLedgerMarkerByRAID returns the LedgerMarker struct for the GLAccount with the supplied LID
func GetInitialLedgerMarkerByRAID(ctx context.Context, raid int64) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetInitialLedgerMarkerByRAID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetInitialLedgerMarkerByRAID.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetInitialLedgerMarkerByRID returns the LedgerMarker struct for the GLAccount with the supplied LID
func GetInitialLedgerMarkerByRID(ctx context.Context, id int64) (LedgerMarker, error) {

	var (
		// err error
		a LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetInitialLedgerMarkerByRID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetInitialLedgerMarkerByRID.QueryRow(fields...)
	}
	return a, ReadLedgerMarker(row, &a)
}

// GetLedgerMarkerOnOrBefore returns the LedgerMarker struct for the GLAccount with the supplied LID
func GetLedgerMarkerOnOrBefore(ctx context.Context, bid, lid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, lid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerMarkerOnOrBefore)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLedgerMarkerOnOrBefore.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

/*// GetPayorLedgerMarkerOnOrBefore returns the LedgerMarker struct for the TCID
func GetPayorLedgerMarkerOnOrBefore(ctx context.Context, bid, tcid int64, dt *time.Time) LedgerMarker {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, tcid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPayorLedgerMarkerOnOrBefore)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetPayorLedgerMarkerOnOrBefore.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}*/

// GetRALedgerMarkerOnOrBeforeDeprecated returns the LedgerMarker struct for the GLAccount with
// the supplied LID filtered for the supplied RentalAgreement raid
// THIS HAS BEEN DEPRECATED  7/27/2017
func GetRALedgerMarkerOnOrBeforeDeprecated(ctx context.Context, bid, lid, raid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, lid, raid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRALedgerMarkerOnOrBeforeDeprecated)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRALedgerMarkerOnOrBeforeDeprecated.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetRALedgerMarkerOnOrBefore returns the LedgerMarker struct for the RAID with
// the supplied LID filtered for the supplied RentalAgreement raid
//=============================================================================
func GetRALedgerMarkerOnOrBefore(ctx context.Context, raid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRALedgerMarkerOnOrBefore)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRALedgerMarkerOnOrBefore.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetRALedgerMarkerOnOrAfter returns the LedgerMarker struct for the RAID with
// the supplied LID filtered for the supplied RentalAgreement raid
//=============================================================================
func GetRALedgerMarkerOnOrAfter(ctx context.Context, raid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRALedgerMarkerOnOrAfter)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRALedgerMarkerOnOrAfter.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetTCLedgerMarkerOnOrBefore returns the LedgerMarker struct for the TCID
// filtered for the supplied date
//=============================================================================
func GetTCLedgerMarkerOnOrBefore(ctx context.Context, tcid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{tcid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTCLedgerMarkerOnOrBefore)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTCLedgerMarkerOnOrBefore.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetTCLedgerMarkerOnOrAfter returns the LedgerMarker struct for the TCID
// filtered for the supplied date
//=============================================================================
func GetTCLedgerMarkerOnOrAfter(ctx context.Context, tcid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{tcid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTCLedgerMarkerOnOrAfter)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTCLedgerMarkerOnOrAfter.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetRentableLedgerMarkerOnOrBefore returns the LedgerMarker struct for the GLAccount with
// the supplied LID filtered for the supplied Rentable rid
func GetRentableLedgerMarkerOnOrBefore(ctx context.Context, bid, lid, rid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, lid, rid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableLedgerMarkerOnOrBefore)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableLedgerMarkerOnOrBefore.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

// GetRARentableLedgerMarkerOnOrBefore returns the LedgerMarker struct for the GLAccount with
// the supplied LID filtered for the supplied Rentable rid
func GetRARentableLedgerMarkerOnOrBefore(ctx context.Context, raid, rid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		// err error
		r LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid, rid, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRARentableLedgerMarkerOnOrBefore)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRARentableLedgerMarkerOnOrBefore.QueryRow(fields...)
	}
	return r, ReadLedgerMarker(row, &r)
}

/*// LoadPayorLedgerMarker returns the LedgerMarker for the supplied bid,tcid
// values. It loads the marker on-or-before dt.  If no such LedgerMarker exists,
// then one will be created.
//
// TODO:  If a subsequent LedgerMarker exists and it is marked as the epoch (3) then
// then it will be updated to normal status as the LedgerMarker just being will
// created will be the new epoch.
//
// INPUTS
//		bid  - business id
//		tcid - which payor
//		dt   - the ledger marker on this date, or the first prior LedgerMarker
//			   will be loaded and returned.
//-----------------------------------------------------------------------------
func LoadPayorLedgerMarker(ctx context.Context, bid, tcid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		err error
		lm LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return lm, ErrSessionRequired
		}
	}

	lm, err = GetPayorLedgerMarkerOnOrBefore(ctx, bid, tcid, dt)
	if err != nil {
		lm.BID = bid
		lm.TCID = tcid
		lm.Dt = *dt
		lm.State = LMINITIAL
		err = InsertLedgerMarker(ctx, &lm)
		if nil != err {
			fmt.Printf("LoadRALedgerMarker: Error creating LedgerMarker: %s\n", err.Error())
		}
	}
	return lm, err
}*/

// LoadRALedgerMarker returns the LedgerMarker for the supplied bid,lid,raid
// values. It loads the marker on-or-before dt.  If no such LedgerMarker exists,
// then one will be created.
//
// TODO:  If a subsequent LedgerMarker exists and it is marked as the epoch (3) then
// then it will be updated to normal status as the LedgerMarker just being will
// created will be the new epoch.
//
// INPUTS
//		bid  - business id
//		lid  - parent ledger id
//		raid - which RentalAgreement
//		dt   - the ledger marker on this date, or the first prior LedgerMarker
//			   will be loaded and returned.
//-----------------------------------------------------------------------------
func LoadRALedgerMarker(ctx context.Context, bid, lid, raid int64, dt *time.Time) (LedgerMarker, error) {

	var (
		err error
		lm  LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return lm, ErrSessionRequired
		}
	}

	lm, err = GetRALedgerMarkerOnOrBeforeDeprecated(ctx, bid, lid, raid, dt)
	if err != nil {
		lm.LID = lid
		lm.BID = bid
		lm.RAID = raid
		lm.Dt = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		lm.State = LMINITIAL
		_, err = InsertLedgerMarker(ctx, &lm)
		if nil != err {
			fmt.Printf("LoadRALedgerMarker: Error creating LedgerMarker: %s\n", err.Error())
		}
	}
	return lm, err
}

// GetLatestLedgerMarkerByGLNo returns the LedgerMarker struct for the GLNo with the supplied name
func GetLatestLedgerMarkerByGLNo(ctx context.Context, bid int64, s string) (LedgerMarker, error) {

	var (
		err error
		lm  LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return lm, ErrSessionRequired
		}
	}

	l, err := GetLedgerByGLNo(ctx, bid, s)
	if err != nil {
		return lm, err
	}
	return GetLatestLedgerMarkerByLID(ctx, bid, l.LID)
}

/*// GetLatestLedgerMarkerByType returns the LedgerMarker struct for the supplied type
func GetLatestLedgerMarkerByType(ctx context.Context, bid int64, t int64) (LedgerMarker, error) {

	var (
		err error
		lm LedgerMarker
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return lm, ErrSessionRequired
		}
	}

	l, err := GetLedgerByType(ctx, bid, t)
	if err != nil {
		return lm, err
	}
	if 0 == l.LID {
		return lm, err
	}
	return GetLatestLedgerMarkerByLID(ctx, bid, l.LID)
}*/

/*// GetAllLedgerMarkersOnOrBefore returns a map of all ledgermarkers for the supplied business and dat
func GetAllLedgerMarkersOnOrBefore(ctx context.Context, bid int64, dt *time.Time) (map[int64]LedgerMarker, error) {

	var (
		err error
		t   = make(map[int64]LedgerMarker)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}

	}


	var rows *sql.Rows
	(bid, dt)
	fields := []interface{}{}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(Prepstmt.GetAllLedgerMarkersOnOrBefore.Query)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = Prepstmt.GetAllLedgerMarkersOnOrBefore.Query.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	// fmt.Printf("%4s  %4s  %4s  %5s  %10s  %8s\n", "LMID", "LID", "BID", "State", "Dt", "Balance")
	for rows.Next() {
		var r LedgerMarker
		err = ReadLedgerMarkers(rows, &r)
		if err != nil {
			return t, err
		}
		t[r.LID] = r
		// fmt.Printf("%4d  %4d  %4d  %5d  %10s  %8.2f\n", r.LMID, r.LID, r.BID, r.State, r.Dt, r.Balance)
	}

	return t, rows.Err()
}*/

//=======================================================
//  L E D G E R
//=======================================================

// GetLedgerList returns an array of all GLAccount
// this is essentially a way to get the exhaustive list of GLAccount numbers for a Business
func GetLedgerList(ctx context.Context, bid int64) ([]GLAccount, error) {

	var (
		err error
		t   []GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerList)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetLedgerList.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r GLAccount
		err = ReadGLAccounts(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// getLedgerMap returns a map of all GLAccounts for the supplied business
func getLedgerMap(bid int64) (map[int64]GLAccount, error) {

	var (
		err error
		t   = make(map[int64]GLAccount)
	)

	var rows *sql.Rows
	fields := []interface{}{bid}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerList)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetLedgerList.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetLedgerList.Query(fields...)

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r GLAccount
		err = ReadGLAccounts(rows, &r)
		if err != nil {
			return t, err
		}
		t[r.LID] = r
	}

	return t, rows.Err()
}

// GetGLAccountMap returns a map of all GLAccounts for the supplied business
func GetGLAccountMap(ctx context.Context, bid int64) (map[int64]GLAccount, error) {

	var (
		// err error
		t = make(map[int64]GLAccount)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	return getLedgerMap(bid)
}

// GetLedger returns the GLAccount struct for the supplied LID
func GetLedger(ctx context.Context, lid int64) (GLAccount, error) {

	var (
		// err error
		a GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{lid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedger)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLedger.QueryRow(fields...)
	}
	return a, ReadGLAccount(row, &a)
}

// GetLedgerEntryByJAID returns the GLAccount struct for the supplied LID
func GetLedgerEntryByJAID(ctx context.Context, bid, lid, jaid int64) (LedgerEntry, error) {

	var (
		// err error
		a LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, lid, jaid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerEntryByJAID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLedgerEntryByJAID.QueryRow(fields...)
	}
	return a, ReadLedgerEntry(row, &a)
}

// GetLedgerEntriesByJAID returns the GLAccount struct for the supplied LID
func GetLedgerEntriesByJAID(ctx context.Context, bid, jaid int64) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, jaid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerEntriesByJAID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetLedgerEntriesByJAID.Query(fields...)
	}

	if err != nil {
		return m, err
	}

	for rows.Next() {
		var le LedgerEntry
		err = ReadLedgerEntries(rows, &le)
		if err != nil {
			return m, err
		}
		m = append(m, le)
	}

	return m, rows.Err()
}

// GetCountLedgerEntries get total count for RentableTypes
// with particular associated business
func GetCountLedgerEntries(ctx context.Context, lid, bid int64) (int64, error) {

	var (
		// err   error
		count int64
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{lid, bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountLedgerEntries)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountLedgerEntries.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetLedgerByGLNo returns the GLAccount struct for the supplied GLNo
func GetLedgerByGLNo(ctx context.Context, bid int64, s string) (GLAccount, error) {

	var (
		// err error
		a GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, s}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerByGLNo)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLedgerByGLNo.QueryRow(fields...)
	}
	return a, ReadGLAccount(row, &a)
}

// GetLedgerByName returns the GLAccount struct for the supplied Name
func GetLedgerByName(ctx context.Context, bid int64, s string) (GLAccount, error) {

	var (
		// err error
		a GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, s}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLedgerByName.QueryRow(fields...)
	}
	return a, ReadGLAccount(row, &a)
}

/*// GetLedgerByType returns the GLAccount struct for the supplied Type
func GetLedgerByType(ctx context.Context, bid, t int64) (GLAccount, error) {

	var (
		// err error
		a GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, t}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerByType)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLedgerByType.QueryRow(fields...)
	}
	return a, ReadGLAccount(row, &a)
}*/

/*// GetRABalanceLedger returns the GLAccount struct for the supplied Type
func GetRABalanceLedger(ctx context.Context, bid, RAID int64) (GLAccount, error) {

	var (
		// err error
		a   GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRABalanceLedger)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRABalanceLedger.QueryRow(fields...)
	}
	return a, ReadGLAccount(row, &a)
}*/

/*// GetSecDepBalanceLedger returns the GLAccount struct for the supplied Type
func GetSecDepBalanceLedger(ctx context.Context, bid, RAID int64) (GLAccount, error) {

	var (
		// err error
		a   GLAccount
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, RAID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetSecDepBalanceLedger)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetSecDepBalanceLedger.QueryRow(fields...)
	}
	return a, ReadGLAccount(row, &a)
}*/

/*// GetDefaultLedgers loads the default GLAccount for the supplied Business bid
func GetDefaultLedgers(ctx context.Context, bid int64) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}


	var rows *sql.Rows
	(bid)
	fields := []interface{}{}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(Prepstmt.GetDefaultLedgers.Query)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = Prepstmt.GetDefaultLedgers.Query.Query(fields...)
	}

	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var r GLAccount
		err = ReadGLAccounts(rows, &r)
		if err != nil {
			return err
		}
		RRdb.BizTypes[bid].DefaultAccts[r.Type] = &r
	}

	return err
}*/

//=======================================================
//  LEDGER ENTRY
//=======================================================

// getLedgerEntryArray returns a list of Ledger Entries for the supplied rows value
func getLedgerEntryArray(ctx context.Context, rows *sql.Rows) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	for rows.Next() {
		var le LedgerEntry
		err = ReadLedgerEntries(rows, &le)
		if err != nil {
			return m, err
		}
		m = append(m, le)
	}

	return m, rows.Err()
}

// GetLedgerEntriesInRange returns a list of Ledger Entries for the supplied Ledger during the supplied range
func GetLedgerEntriesInRange(ctx context.Context, d1, d2 *time.Time, bid, lid int64) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, lid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerEntriesInRangeByLID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetLedgerEntriesInRangeByLID.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}

// GetLedgerEntriesForRAID returns a list of Ledger Entries for the supplied RentalAgreement and Ledger
func GetLedgerEntriesForRAID(ctx context.Context, d1, d2 *time.Time, raid, lid int64) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{d1, d2, raid, lid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerEntriesForRAID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetLedgerEntriesForRAID.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}

// GetLedgerEntriesForRentable returns a list of Ledger Entries for the supplied Rentable (rid) and Ledger (lid)
func GetLedgerEntriesForRentable(ctx context.Context, d1, d2 *time.Time, rid, lid int64) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{d1, d2, rid, lid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLedgerEntriesForRentable)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetLedgerEntriesForRentable.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}

// GetAllLedgerEntriesForRAID returns a list of Ledger Entries for the supplied RentalAgreement
func GetAllLedgerEntriesForRAID(ctx context.Context, d1, d2 *time.Time, raid int64) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{d1, d2, raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllLedgerEntriesForRAID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllLedgerEntriesForRAID.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}

// GetAllLedgerEntriesForRID returns a list of Ledger Entries for the supplied Rentable ID
func GetAllLedgerEntriesForRID(ctx context.Context, d1, d2 *time.Time, rid int64) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{d1, d2, rid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllLedgerEntriesForRID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllLedgerEntriesForRID.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}

// GetAllLedgerEntriesInRange returns a list of Ledger Entries for the supplied business and time period
func GetAllLedgerEntriesInRange(ctx context.Context, bid int64, d1, d2 *time.Time) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllLedgerEntriesInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllLedgerEntriesInRange.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}

/*// GetLedgerEntriesInRange returns a list of Ledger Entries for the supplied business and time period
func GetLedgerEntriesInRange(ctx context.Context, bid, lid, raid int64, d1, d2 *time.Time) ([]LedgerEntry, error) {

	var (
		err error
		m   []LedgerEntry
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}


	var rows *sql.Rows
	(bid, lid, raid, d1, d2)
	fields := []interface{}{}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(Prepstmt.GetLedgerEntriesInRange.Query)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = Prepstmt.GetLedgerEntriesInRange.Query.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	return getLedgerEntryArray(ctx, rows)
}*/

//=======================================================
//  NOTES
//=======================================================

// GetNote reads a Note structure based on the supplied Note id
func GetNote(ctx context.Context, tid int64, t *Note) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{tid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetNote)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetNote.QueryRow(fields...)
	}
	return ReadNote(row, t)
}

// GetNoteAndChildNotes reads a Note structure based on the supplied Note id, then it reads all its child notes, organizes them by Date
// and returns them in an array
func GetNoteAndChildNotes(ctx context.Context, nid int64) (Note, error) {

	var (
		err error
		n   Note
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return n, ErrSessionRequired
		}
	}

	err = GetNote(ctx, nid, &n)
	if err != nil {
		return n, err
	}

	var rows *sql.Rows
	fields := []interface{}{nid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetNoteAndChildNotes)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetNoteAndChildNotes.Query(fields...)
	}

	if err != nil {
		return n, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Note
		err = ReadNotes(rows, &p)
		if err != nil {
			return n, err
		}
		n.CN = append(n.CN, p)
	}

	return n, rows.Err()
}

//=======================================================
//  NOTELIST
//=======================================================

// GetNoteList reads a NoteList structure based on the supplied NoteList id
func GetNoteList(ctx context.Context, nlid int64) (NoteList, error) {

	var (
		err error
		m   NoteList
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{nlid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetNoteList)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetNoteList.QueryRow(fields...)
	}
	err = ReadNoteList(row, &m)
	if err != nil {
		return m, err
	}

	var rows *sql.Rows
	fields = []interface{}{nlid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetNoteListMembers)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetNoteListMembers.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var nid int64

		err = rows.Scan(&nid)
		if err != nil {
			return m, err
		}
		p, err := GetNoteAndChildNotes(ctx, nid)
		if err != nil {
			return m, err
		}
		m.N = append(m.N, p)
	}

	return m, rows.Err()
}

//=======================================================
//  NOTE TYPE
//=======================================================

// GetNoteType reads a NoteType structure based on the supplied NoteType id
func GetNoteType(ctx context.Context, ntid int64, t *NoteType) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{ntid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetNoteType)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetNoteType.QueryRow(fields...)
	}
	return ReadNoteType(row, t)
}

// getBusinessAllNoteTypes reads a NoteType structure based for all NoteTypes in the supplied bid
func getBusinessAllNoteTypes(bid int64) ([]NoteType, error) {

	var (
		err error
		m   []NoteType
	)

	var rows *sql.Rows
	fields := []interface{}{bid}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllNoteTypes)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllNoteTypes.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetAllNoteTypes.Query(fields...)

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var p NoteType
		err = ReadNoteTypes(rows, &p)
		if err != nil {
			return m, err
		}
		m = append(m, p)
	}

	return m, rows.Err()
}

// GetAllNoteTypes reads a NoteType structure based for all NoteTypes in the supplied bid
func GetAllNoteTypes(ctx context.Context, bid int64) ([]NoteType, error) {

	var (
		// err error
		m []NoteType
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	return getBusinessAllNoteTypes(bid)
}

//=======================================================
//  P A Y M E N T   T Y P E S
//=======================================================

// GetPaymentType reads a PaymentType structure based on the supplied bid and na
func GetPaymentType(ctx context.Context, id int64, a *PaymentType) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPaymentType)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetPaymentType.QueryRow(fields...)
	}
	return ReadPaymentType(row, a)
}

// GetPaymentTypeByName reads a PaymentType structure based on the supplied bid and na
func GetPaymentTypeByName(ctx context.Context, bid int64, name string, a *PaymentType) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPaymentTypeByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetPaymentTypeByName.QueryRow(fields...)
	}
	return ReadPaymentType(row, a)
}

// GetPaymentTypesByBusiness returns a slice of payment types indexed by the PMTID for the supplied Business
func GetPaymentTypesByBusiness(ctx context.Context, bid int64) (map[int64]PaymentType, error) {

	var (
		err error
		t   = make(map[int64]PaymentType)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPaymentTypesByBusiness)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetPaymentTypesByBusiness.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var a PaymentType
		err = ReadPaymentTypes(rows, &a)
		if err != nil {
			return t, err
		}
		t[a.PMTID] = a
	}

	return t, rows.Err()
}

//=======================================================
//  RATE PLAN
//=======================================================

// GetRatePlan reads a RatePlan structure based on the supplied RatePlan id
func GetRatePlan(ctx context.Context, id int64, a *RatePlan) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlan)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRatePlan.QueryRow(fields...)
	}
	return ReadRatePlan(row, a)
}

// GetRatePlanByName reads a RatePlan structure based on the supplied RatePlan id
func GetRatePlanByName(ctx context.Context, id int64, s string, a *RatePlan) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id, s}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlanByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRatePlanByName.QueryRow(fields...)
	}
	return ReadRatePlan(row, a)
}

// GetAllRatePlans reads all RatePlan structures based on the supplied bid
func GetAllRatePlans(ctx context.Context, id int64) ([]RatePlan, error) {

	var (
		err error
		m   []RatePlan
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRatePlans)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRatePlans.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var p RatePlan
		err = ReadRatePlans(rows, &p)
		if err != nil {
			return m, err
		}
		m = append(m, p)
	}

	return m, rows.Err()
}

// GetRatePlanRef reads a RatePlanRef structure based on the supplied RatePlanRef id
func GetRatePlanRef(ctx context.Context, id int64, a *RatePlanRef) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlanRef)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRatePlanRef.QueryRow(fields...)
	}
	return ReadRatePlanRef(row, a)
}

// GetRatePlanRefFull reads a RatePlanRef structure based on the supplied RatePlanRef id and
// pulls in all the RTRate and SPRate structure arrays
func GetRatePlanRefFull(ctx context.Context, id int64, a *RatePlanRef) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	if a.RPRID == 0 {
		var row *sql.Row
		fields := []interface{}{id}
		if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
			stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlanRef)
			defer stmt.Close()
			row = stmt.QueryRow(fields...)
		} else {
			row = RRdb.Prepstmt.GetRatePlanRef.QueryRow(fields...)
		}
		err = ReadRatePlanRef(row, a)
		if err != nil {
			return err
		}
	}

	// =====================================
	// LOAD ALL Rate Plan RT RATES
	// =====================================

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRatePlanRefRTRates)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRatePlanRefRTRates.Query(fields...)
	}

	if err != nil {
		Ulog("GetRatePlanRefFull:   GetAllRatePlanRefRTRates - error = %s\n", err.Error())
		return err
	}

	for rows.Next() {
		var p RatePlanRefRTRate
		err = ReadRatePlanRefRTRates(rows, &p)
		if err != nil {
			return err
		}
		a.RT = append(a.RT, p)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	rows.Close() // REMEMBER TO CLOSE IT HERE, Before re-assinging something new in rows itself

	// =====================================
	// LOAD ALL Rate Plan SPECIALTY RATES
	// =====================================

	// var rows *sql.Rows
	fields = []interface{}{a.RPRID, a.RPID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRatePlanRefSPRates)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRatePlanRefSPRates.Query(fields...)
	}

	if err != nil {
		Ulog("GetRatePlanRefFull: GetAllRatePlanRefSPRates - error = %s\n", err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var p RatePlanRefSPRate
		err = ReadRatePlanRefSPRates(rows, &p)
		if err != nil {
			return err
		}
		a.SP = append(a.SP, p)
	}
	return rows.Err()
}

// GetRatePlanRefsInRange reads a RatePlanRef structure based on the supplied RatePlan id and the date.
func GetRatePlanRefsInRange(ctx context.Context, id int64, d1, d2 *time.Time) ([]RatePlanRef, error) {

	var (
		err error
		m   []RatePlanRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlanRefsInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRatePlanRefsInRange.Query(fields...)
	}

	if err != nil {
		Ulog("GetRatePlanRefsInRange: error = %s\n", err.Error())
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a RatePlanRef
		err = ReadRatePlanRefs(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetAllRatePlanRefsInRange reads all RatePlanRef structure based on the supplied date range
func GetAllRatePlanRefsInRange(ctx context.Context, d1, d2 *time.Time) ([]RatePlanRef, error) {

	var (
		err error
		m   []RatePlanRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRatePlanRefsInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRatePlanRefsInRange.Query(fields...)
	}

	if err != nil {
		Ulog("GetAllRatePlanRefsInRange: error = %s\n", err.Error())
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a RatePlanRef
		err = ReadRatePlanRefs(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetRatePlanRefRTRate reads the RatePlanRefRTRate struct for the supplied rprid and rtid
func GetRatePlanRefRTRate(ctx context.Context, rprid, rtid int64, a *RatePlanRefRTRate) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rprid, rtid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlanRefRTRate)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRatePlanRefRTRate.QueryRow(fields...)
	}
	return ReadRatePlanRefRTRate(row, a)
}

// GetRatePlanRefSPRate reads the RatePlanRefSPRate struct for the supplied rprid and rtid
func GetRatePlanRefSPRate(ctx context.Context, rprid, rtid int64, a *RatePlanRefSPRate) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rprid, rtid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRatePlanRefSPRate)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRatePlanRefSPRate.QueryRow(fields...)
	}
	return ReadRatePlanRefSPRate(row, a)
}

// GetAllRatePlanRefSPRates reads all RatePlanRefSPRate structures based on the supplied RatePlan id and the date.
func GetAllRatePlanRefSPRates(ctx context.Context, rprid, rtid int64) ([]RatePlanRefSPRate, error) {

	var (
		err error
		m   []RatePlanRefSPRate
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{rprid, rtid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRatePlanRefSPRates)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRatePlanRefSPRates.Query(fields...)
	}

	if err != nil {
		Ulog("GetAllRatePlanRefSPRates: error = %s\n", err.Error())
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a RatePlanRefSPRate
		err = ReadRatePlanRefSPRates(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

//=======================================================
//  RECEIPT ALLOCATION
//=======================================================

// GetReceipt returns a Receipt structure for the supplied RCPTID
func GetReceipt(ctx context.Context, rcptid int64) (Receipt, error) {

	var (
		err error
		r   Receipt
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	r, err = GetReceiptNoAllocations(ctx, rcptid)
	if err != nil {
		return r, err
	}

	return r, GetReceiptAllocations(ctx, rcptid, &r)
}

// GetReceiptAllocation returns a ReceiptAllocation structure for the supplied RCPTID
func GetReceiptAllocation(ctx context.Context, id int64) (ReceiptAllocation, error) {

	var (
		// err error
		r ReceiptAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceiptAllocation)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetReceiptAllocation.QueryRow(fields...)
	}
	return r, ReadReceiptAllocation(row, &r)
}

// GetReceiptNoAllocations returns a Receipt structure for the supplied RCPTID.
// It does not get the receipt allocations
func GetReceiptNoAllocations(ctx context.Context, rcptid int64) (Receipt, error) {

	var (
		// err error
		r Receipt
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rcptid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceipt)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetReceipt.QueryRow(fields...)
	}
	return r, ReadReceipt(row, &r)
}

// GetReceiptDuplicate returns a Receipt structure for the supplied RCPTID
func GetReceiptDuplicate(ctx context.Context, dt *time.Time, amt float64, docno string) (Receipt, error) {

	var (
		// err error
		r Receipt
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{dt, amt, docno}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceiptDuplicate)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetReceiptDuplicate.QueryRow(fields...)
	}
	return r, ReadReceipt(row, &r)
}

// GetReceiptAllocations loads all Receipt allocations associated with the supplied Receipt id into
// the RA array within a Receipt structure
func GetReceiptAllocations(ctx context.Context, rcptid int64, r *Receipt) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{rcptid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceiptAllocations)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetReceiptAllocations.Query(fields...)
	}

	if err != nil {
		return err
	}
	defer rows.Close()

	r.RA = make([]ReceiptAllocation, 0)
	for rows.Next() {
		var a ReceiptAllocation
		err = ReadReceiptAllocations(rows, &a)
		if err != nil {
			return err
		}
		r.RA = append(r.RA, a)
	}

	return rows.Err()
}

// GetReceipts for the supplied Business (bid) in date range [d1 - d2)
func GetReceipts(ctx context.Context, bid int64, d1, d2 *time.Time) ([]Receipt, error) {

	var (
		err error
		t   []Receipt
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceiptsInDateRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetReceiptsInDateRange.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Receipt
		err = ReadReceipts(rows, &r)
		if err != nil {
			return t, err
		}

		r.RA = make([]ReceiptAllocation, 0)
		err = GetReceiptAllocations(ctx, r.RCPTID, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// getReceiptAllocationList for the supplied rows variable
func getReceiptAllocationList(ctx context.Context, rows *sql.Rows) ([]ReceiptAllocation, error) {

	var (
		err error
		t   []ReceiptAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	defer rows.Close()

	for rows.Next() {
		var r ReceiptAllocation
		err = ReadReceiptAllocations(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetASMReceiptAllocationsInRAIDDateRange for the supplied RentalAgreement in date range [d1 - d2).
// To do this we select all the ReceiptAllocations that occurred during d1-d2 that involved
// raid.
func GetASMReceiptAllocationsInRAIDDateRange(ctx context.Context, raid int64, d1, d2 *time.Time) ([]ReceiptAllocation, error) {

	var (
		err error
		t   []ReceiptAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{raid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetASMReceiptAllocationsInRAIDDateRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetASMReceiptAllocationsInRAIDDateRange.Query(fields...)
	}

	if err != nil {
		return t, err
	}

	return getReceiptAllocationList(ctx, rows)
}

// GetReceiptAllocationsByASMID returns any payment allocation on targeted at the supplied ASMID.
// This call is used primarily to determine how much payment is left to make on a partially paid
// assessment.
func GetReceiptAllocationsByASMID(ctx context.Context, bid, asmid int64) ([]ReceiptAllocation, error) {

	var (
		err error
		t   []ReceiptAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, asmid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceiptAllocationsByASMID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetReceiptAllocationsByASMID.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getReceiptAllocationList(ctx, rows)
}

// GetReceiptAllocationsThroughDate selects the ReceiptAllocations associated with receipt id
// and that happened on or before the supplied date
// @params
//	 id = RCPTID of desired allocations
//   dt = date for all allocations to be on or prior to
// @returns  []ReceiptAllocation
func GetReceiptAllocationsThroughDate(ctx context.Context, id int64, dt *time.Time) ([]ReceiptAllocation, error) {

	var (
		err error
		t   []ReceiptAllocation
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetReceiptAllocationsThroughDate)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetReceiptAllocationsThroughDate.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getReceiptAllocationList(ctx, rows)
}

// GetReceiptAllocationAmountsOnDate returns the amount of unallocated funds in id on the
// supplied date
// @params
//	 id = RCPTID of the receipt in question
//   dt = date on which the unallocated amount is desired
// @returns  float64 of:
//   receipt amount
//   amount allocated as of dt
//   amount unallocated as of dt
func GetReceiptAllocationAmountsOnDate(ctx context.Context, id int64, dt *time.Time) (float64, float64, float64, error) {

	var (
		err     error
		amt     float64
		alloc   float64
		unalloc float64
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return amt, alloc, unalloc, ErrSessionRequired
		}
	}

	// get receipt
	rcpt, err := GetReceipt(ctx, id)
	if err != nil {
		return amt, alloc, unalloc, err
	}
	amt = rcpt.Amount
	unalloc = rcpt.Amount

	m, err := GetReceiptAllocationsThroughDate(ctx, id, dt)
	if err != nil {
		return amt, alloc, unalloc, err
	}

	for i := 0; i < len(m); i++ {
		unalloc -= m[i].Amount
		alloc += m[i].Amount
	}
	return amt, alloc, unalloc, err
}

// GetUnallocatedReceiptsByPayor returns the receipts paid by the supplied payor tcid that
// have not yet been fully allocated.
func GetUnallocatedReceiptsByPayor(ctx context.Context, bid, tcid int64) ([]Receipt, error) {

	var (
		err error
		t   []Receipt
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetUnallocatedReceiptsByPayor)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetUnallocatedReceiptsByPayor.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Receipt
		err = ReadReceipts(rows, &r)
		if err != nil {
			return t, err
		}

		r.RA = make([]ReceiptAllocation, 0) // the receipt may be partially allocated
		err = GetReceiptAllocations(ctx, r.RCPTID, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetPayorUnallocatedReceiptsCount returns a count of unallocated receipts for the supplied bid & tcid
func GetPayorUnallocatedReceiptsCount(ctx context.Context, bid, tcid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPayorUnallocatedReceiptsCount)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetPayorUnallocatedReceiptsCount.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

//=======================================================
//  R E N T A B L E
//=======================================================

// GetRentableByID reads a Rentable structure based on the supplied Rentable id
func GetRentableByID(ctx context.Context, rid int64, r *Rentable) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentable)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentable.QueryRow(fields...)
	}
	return ReadRentable(row, r)
}

// GetRentable reads and returns a Rentable structure based on the supplied Rentable id
func GetRentable(ctx context.Context, rid int64) (Rentable, error) {

	var (
		// err error
		r Rentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	return r, GetRentableByID(ctx, rid, &r)
}

// GetRentableByName reads and returns a Rentable structure based on the supplied Rentable id
func GetRentableByName(ctx context.Context, name string, bid int64) (Rentable, error) {

	var (
		// err error
		r Rentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{name, bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableByName.QueryRow(fields...)
	}
	return r, ReadRentable(row, &r)
}

// GetRentableTypeDown returns the values needed for typedown controls:
// input:   bid - business
//            s - string or substring to search for
//        limit - return no more than this many matches
// return a slice of RentableTypeDowns and an error.
func GetRentableTypeDown(ctx context.Context, bid int64, s string, limit int) ([]RentableTypeDown, error) {

	var (
		err error
		m   []RentableTypeDown
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	s = "%" + s + "%"

	var rows *sql.Rows
	fields := []interface{}{bid, s, limit}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableTypeDown)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableTypeDown.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var t RentableTypeDown
		err = ReadRentableTypeDown(rows, &t)
		t.Recid = t.RID
		if err != nil {
			return m, err
		}
		m = append(m, t)
	}

	return m, rows.Err()
}

// GetXRentable reads an XRentable structure based on the RID.
func GetXRentable(ctx context.Context, rid int64, x *XRentable) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	if x.R.RID == 0 && rid > 0 {
		err = GetRentableByID(ctx, rid, &x.R)
		if err != nil {
			return err
		}
	}

	x.S, err = GetAllRentableSpecialtyRefs(ctx, x.R.BID, x.R.RID)
	return err
}

// GetRentableUser returns a Rentable User record with the supplied RUID
func GetRentableUser(ctx context.Context, ruid int64) (RentableUser, error) {

	var (
		// err error
		r RentableUser
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{ruid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableUser)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableUser.QueryRow(fields...)
	}
	return r, ReadRentableUser(row, &r)
}

// GetRentableUserByRBT returns a Rentable User record matching the supplied
// RID, BID, TCID
func GetRentableUserByRBT(ctx context.Context, rid, bid, tcid int64) (RentableUser, error) {

	var (
		// err error
		r RentableUser
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rid, bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableUserByRBT)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableUserByRBT.QueryRow(fields...)
	}
	return r, ReadRentableUser(row, &r)
}

// GetRentableSpecialtyTypeByName returns a list of specialties associated with the supplied Rentable
func GetRentableSpecialtyTypeByName(ctx context.Context, bid int64, name string) (RentableSpecialty, error) {

	var (
		// err error
		rsp RentableSpecialty
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rsp, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableSpecialtyTypeByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableSpecialtyTypeByName.QueryRow(fields...)
	}
	return rsp, ReadRentableSpecialty(row, &rsp)
}

// GetRentableSpecialtyType returns the RentableSpecialty record for the supplied RSPID
func GetRentableSpecialtyType(ctx context.Context, rspid int64) (RentableSpecialty, error) {

	var (
		// err error
		rs RentableSpecialty
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rspid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableSpecialtyType)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableSpecialtyType.QueryRow(fields...)
	}
	return rs, ReadRentableSpecialty(row, &rs)
}

// GetAllRentableSpecialtyRefs returns a list of specialties associated with the supplied Rentable
func GetAllRentableSpecialtyRefs(ctx context.Context, bid, rid int64) ([]int64, error) {

	var (
		err error
		m   []int64
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	// first, get the specialties for this Rentable

	var rows *sql.Rows
	fields := []interface{}{bid, rid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableSpecialtyRefs)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableSpecialtyRefs.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var uspid int64
		err = rows.Scan(&uspid)
		if err != nil {
			return m, err
		}
		m = append(m, uspid)
	}

	return m, rows.Err()
}

/*// SelectRentableTypeRefForDate returns the first RTID of the list where the supplied date falls in range
func SelectRentableTypeRefForDate(ctx context.Context, rsa *[]RentableSpecialty, dt *time.Time) (RentableSpecialty, error) {

	var (
		err error
		r RentableSpecialty
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	for i := 0; i < len(*rsa); i++ {
		if DateInRange(dt, &(*rsa)[i].DtStart, &(*rsa)[i].DtStop) {
			return (*rsa)[i], err
		}
	}
	return r, err // nothing matched
}*/

// GetRentableSpecialtyTypesForRentableByRange returns an array of RentableSpecialty structures that
// apply for the supplied time range d1,d2
func GetRentableSpecialtyTypesForRentableByRange(ctx context.Context, bid, rid int64, d1, d2 *time.Time) ([]RentableSpecialty, error) {

	var (
		err  error
		rsta []RentableSpecialty
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rsta, ErrSessionRequired
		}
	}

	rsrefs, err := GetRentableSpecialtyRefsByRange(ctx, bid, rid, d1, d2)
	if err != nil {
		return rsta, err
	}

	for i := 0; i < len(rsrefs); i++ {
		rs, err := GetRentableSpecialtyType(ctx, rsrefs[i].RSPID)
		if err != nil {
			return rsta, err
		}
		rsta = append(rsta, rs)
	}

	return rsta, err
}

// GetRentableSpecialtyRefsByRange loads all the RentableSpecialtyRef records that overlap the supplied time range
// and returns them in an array
func GetRentableSpecialtyRefsByRange(ctx context.Context, bid, rid int64, d1, d2 *time.Time) ([]RentableSpecialtyRef, error) {

	var (
		err error
		rs  []RentableSpecialtyRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, rid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableSpecialtyRefsByRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableSpecialtyRefsByRange.Query(fields...)
	}

	if err != nil {
		return rs, err
	}
	defer rows.Close()

	for rows.Next() {
		var a RentableSpecialtyRef
		err = ReadRentableSpecialtyRefs(rows, &a)
		if err != nil {
			return rs, err
		}
		rs = append(rs, a)
	}

	return rs, rows.Err()
}

// GetRentableTypeRef gets RentableTypeRef record for given RTRID -- RentableTypeRef ID (unique ID)
func GetRentableTypeRef(ctx context.Context, rtrid int64) (RentableTypeRef, error) {

	var (
		// err error
		rtr RentableTypeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rtr, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rtrid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableTypeRef)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableTypeRef.QueryRow(fields...)
	}
	return rtr, ReadRentableTypeRef(row, &rtr)
}

// SelectRentableTypeRefForDate returns the first RTID of the list where the supplied date falls in range
func SelectRentableTypeRefForDate(ctx context.Context, rta *[]RentableTypeRef, dt *time.Time) (RentableTypeRef, error) {

	var (
		err error
		r   RentableTypeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	for i := 0; i < len(*rta); i++ {
		if DateInRange(dt, &(*rta)[i].DtStart, &(*rta)[i].DtStop) {
			return (*rta)[i], err
		}
	}

	return r, err // nothing matched
}

// getRTRefs performs the query over the supplied rows and returns a slice of result records
func getRTRefs(ctx context.Context, rows *sql.Rows) ([]RentableTypeRef, error) {

	var (
		err error
		rs  []RentableTypeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}
	defer rows.Close()

	for rows.Next() {
		var a RentableTypeRef
		err = ReadRentableTypeRefs(rows, &a)
		if err != nil {
			return rs, err
		}
		rs = append(rs, a)
	}

	return rs, rows.Err()
}

// GetRentableTypeRefsByRange loads all the RentableTypeRef records that overlap the supplied time range
// and returns them in an array
func GetRentableTypeRefsByRange(ctx context.Context, RID int64, d1, d2 *time.Time) ([]RentableTypeRef, error) {

	var (
		err error
		rtr []RentableTypeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rtr, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{RID, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableTypeRefsByRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableTypeRefsByRange.Query(fields...)
	}

	if err != nil {
		return rtr, err
	}
	return getRTRefs(ctx, rows)
}

// GetRentableTypeRefs loads all the RentableTypeRef records for a particular
func GetRentableTypeRefs(ctx context.Context, RID int64) ([]RentableTypeRef, error) {

	var (
		err error
		rtr []RentableTypeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rtr, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{RID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableTypeRefs)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableTypeRefs.Query(fields...)
	}

	if err != nil {
		return rtr, err
	}

	return getRTRefs(ctx, rows)
}

// GetRTIDForDate returns the RTID in effect on the supplied date
func GetRTIDForDate(ctx context.Context, RID int64, d1 *time.Time) (int64, error) {

	var (
		err  error
		rtid int64
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rtid, ErrSessionRequired
		}
	}

	DtStop, _ := StringToDate("1/1/9999")
	m, err := GetRentableTypeRefsByRange(ctx, RID, d1, &DtStop)
	if err != nil {
		return rtid, err
	}

	if len(m) > 0 {
		rtid = m[0].RTID
	}
	return rtid, err
}

// GetRentableTypeRefForDate returns the RTID in effect on the supplied date
func GetRentableTypeRefForDate(ctx context.Context, RID int64, d1 *time.Time) (RentableTypeRef, error) {

	var (
		err error
		r   RentableTypeRef
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	DtStop, _ := StringToDate("1/1/9999")
	m, err := GetRentableTypeRefsByRange(ctx, RID, d1, &DtStop)
	if err != nil {
		return r, err
	}

	if len(m) > 0 {
		r = m[0]
	}
	return r, err
}

// GetRentableStatus gets RentableStatus record for given RSID -- RentableStatus ID (unique ID)
func GetRentableStatus(ctx context.Context, rsid int64) (RentableStatus, error) {

	var (
		// err error
		rs RentableStatus
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rsid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableStatus)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableStatus.QueryRow(fields...)
	}
	return rs, ReadRentableStatus(row, &rs)
}

// getRentableStatusRows loads all the RentableStatus records for rows
func getRentableStatusRows(ctx context.Context, rows *sql.Rows) ([]RentableStatus, error) {

	var (
		err error
		rs  []RentableStatus
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}
	defer rows.Close()

	for rows.Next() {
		var a RentableStatus
		err = ReadRentableStatuss(rows, &a)
		if err != nil {
			return rs, err
		}
		rs = append(rs, a)
	}

	return rs, rows.Err()
}

// GetRentableStatusOnOrAfter loads all the RentableStatus records that overlap the supplied time range
func GetRentableStatusOnOrAfter(ctx context.Context, RID int64, dt *time.Time) (RentableStatus, error) {

	var (
		// err error
		a RentableStatus
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{RID, dt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableStatusOnOrAfter)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableStatusOnOrAfter.QueryRow(fields...)
	}
	return a, ReadRentableStatus(row, &a)
}

// GetRentableStatusByRange loads all the RentableStatus records that overlap the supplied time range
func GetRentableStatusByRange(ctx context.Context, RID int64, d1, d2 *time.Time) ([]RentableStatus, error) {

	var (
		err error
		rs  []RentableStatus
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{RID, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableStatusByRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableStatusByRange.Query(fields...)
	}

	if err != nil {
		return rs, err
	}
	return getRentableStatusRows(ctx, rows)
}

// GetAllRentableStatus loads all the RentableStatus records that overlap the supplied time range
func GetAllRentableStatus(ctx context.Context, RID int64) ([]RentableStatus, error) {

	var (
		err error
		rs  []RentableStatus
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rs, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{RID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRentableStatus)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRentableStatus.Query(fields...)
	}

	if err != nil {
		return rs, err
	}
	return getRentableStatusRows(ctx, rows)
}

//=======================================================
//  R E N T A B L E   T Y P E
//=======================================================

// GetRentableType returns characteristics of the Rentable
func GetRentableType(ctx context.Context, rtid int64, rt *RentableType) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rtid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableType)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableType.QueryRow(fields...)
	}
	err = ReadRentableType(row, rt)
	if err != nil {
		return err
	}

	var cerr error
	rt.CA, cerr = GetAllCustomAttributes(ctx, ELEMRENTABLETYPE, rtid)
	if cerr != nil { // it's not really an error if we don't find any custom attributes
		err = cerr
	}

	return err
}

// GetRentableTypeByStyle returns characteristics of the Rentable
func GetRentableTypeByStyle(ctx context.Context, name string, bid int64) (RentableType, error) {

	var (
		// err error
		rt RentableType
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rt, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{name, bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableTypeByStyle)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableTypeByStyle.QueryRow(fields...)
	}
	return rt, ReadRentableType(row, &rt)
}

// GetRentableTypeByName returns characteristics of the Rentable
func GetRentableTypeByName(ctx context.Context, name string, bid int64) (RentableType, error) {

	var (
		// err error
		rt RentableType
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rt, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{name, bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableTypeByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableTypeByName.QueryRow(fields...)
	}
	return rt, ReadRentableType(row, &rt)
}

// getBizRentableTypes returns a slice of RentableType indexed by the RTID
func getBizRentableTypes(bid int64) (map[int64]RentableType, error) {

	var (
		err error
		t   = make(map[int64]RentableType)
	)

	var rows *sql.Rows
	fields := []interface{}{bid}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllBusinessRentableTypes)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllBusinessRentableTypes.Query(fields...)
	}*/
	rows, err = RRdb.Prepstmt.GetAllBusinessRentableTypes.Query(fields...)
	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var a RentableType
		err = ReadRentableTypes(rows, &a)
		if err != nil {
			return t, err
		}
		a.MR = []RentableMarketRate{}
		err = getRentableMarketRates(&a)
		if err != nil {
			return t, err
		}
		t[a.RTID] = a
	}

	return t, rows.Err()
}

// GetBusinessRentableTypes returns a slice of RentableType indexed by the RTID
func GetBusinessRentableTypes(ctx context.Context, bid int64) (map[int64]RentableType, error) {

	var (
		// err error
		t = make(map[int64]RentableType)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	return getBizRentableTypes(bid)
}

// getRentableMarketRates loads all the MarketRate rent information for this Rentable into an array
func getRentableMarketRates(rt *RentableType) error {

	var (
		err error
	)

	// now get all the MarketRate rent info...

	var rows *sql.Rows
	fields := []interface{}{rt.RTID}
	/*if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableMarketRates)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableMarketRates.Query(fields...)
	}*/

	rows, err = RRdb.Prepstmt.GetRentableMarketRates.Query(fields...)
	if err != nil {
		return err
	}
	defer rows.Close()

	LatestMRDTStart := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var a RentableMarketRate
		err = ReadRentableMarketRates(rows, &a)
		if err != nil {
			return err
		}
		if a.DtStart.After(LatestMRDTStart) {
			LatestMRDTStart = a.DtStart
			rt.MRCurrent = a.MarketRate
		}
		rt.MR = append(rt.MR, a)
	}

	return rows.Err()
}

// GetRentableMarketRates loads all the MarketRate rent information for this Rentable into an array
func GetRentableMarketRates(ctx context.Context, rt *RentableType) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	return getRentableMarketRates(rt)
}

// GetRentableMarketRateInstance returns instance of rentableMarketRate for given RMRID
func GetRentableMarketRateInstance(ctx context.Context, rmrid int64) (RentableMarketRate, error) {

	var (
		// err error
		rmr RentableMarketRate
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return rmr, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rmrid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableMarketRateInstance)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentableMarketRateInstance.QueryRow(fields...)
	}
	return rmr, ReadRentableMarketRate(row, &rmr)
}

// GetRentableMarketRate returns the market-rate rent amount for r during the
// given time range. If the time range is large and spans multiple price
// changes, the chronologically earliest price that fits in the time range
// will be returned. It is best to provide as small a timerange d1-d2 as
// possible to minimize risk of overlap
//-----------------------------------------------------------------------------
func GetRentableMarketRate(ctx context.Context, xbiz *XBusiness, RID int64, d1, d2 *time.Time) (float64, error) {

	var (
		err             error
		marketRateValue float64
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return marketRateValue, ErrSessionRequired
		}
	}

	rtid, err := GetRTIDForDate(ctx, RID, d1) // first thing... find the RTID for this time range
	if err != nil {
		return marketRateValue, err
	}

	mr := xbiz.RT[rtid].MR
	// Console("GetRentableMarketRate: len(mr) is %d\n", len(mr))
	for i := 0; i < len(mr); i++ {
		if DateRangeOverlap(d1, d2, &mr[i].DtStart, &mr[i].DtStop) {
			return mr[i].MarketRate, err
		}
	}
	return marketRateValue, err
}

// GetRentableUsersInRange returns an array of user (in the form of user)
// associated with the supplied RentalAgreement ID during the time range d1-d2
//-----------------------------------------------------------------------------
func GetRentableUsersInRange(ctx context.Context, rid int64, d1, d2 *time.Time) ([]RentableUser, error) {
	var err error
	var t []RentableUser

	if sessionCheck(ctx) {
		return t, ErrSessionRequired
	}

	var rows *sql.Rows
	fields := []interface{}{rid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentableUsersInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentableUsersInRange.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r RentableUser
		err = ReadRentableUsers(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

//=======================================================
//  R E N T A L   A G R E E M E N T
//=======================================================

// GetRentalAgreement returns the RentalAgreement struct for the supplied rental agreement id
func GetRentalAgreement(ctx context.Context, raid int64) (RentalAgreement, error) {

	var (
		// err error
		r RentalAgreement
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreement)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreement.QueryRow(fields...)
	}
	return r, ReadRentalAgreement(row, &r)
}

// LoadXRentalAgreement is like GetXRentalAgreement except that it assumes that some of the structure may
// already be loaded. It only loads those portions that appear not to already be loaded.
func LoadXRentalAgreement(ctx context.Context, raid int64, r *RentalAgreement, d1, d2 *time.Time) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	if r.RAID != raid {
		(*r), err = GetRentalAgreement(ctx, raid)
		if err != nil {
			return err
		}
	}

	t, err := GetRentalAgreementRentables(ctx, raid, d1, d2)
	if err != nil {
		return err
	}

	r.R = make([]XRentable, 0)
	for i := 0; i < len(t); i++ {
		var xu XRentable
		err = GetXRentable(ctx, t[i].RID, &xu)
		if err != nil {
			return err
		}
		r.R = append(r.R, xu)
	}

	m, err := GetRentalAgreementPayorsInRange(ctx, raid, d1, d2)
	if err != nil {
		return err
	}

	r.P = make([]XPerson, 0)
	for i := 0; i < len(m); i++ {
		var xp XPerson
		err = GetXPerson(ctx, m[i].TCID, &xp)
		if err != nil {
			return err
		}

		r.P = append(r.P, xp)
	}

	for j := 0; j < len(r.R); j++ {
		n, err := GetRentableUsersInRange(ctx, r.R[j].R.RID, d1, d2)
		if err != nil {
			return err
		}

		r.T = make([]XPerson, 0)
		for i := 0; i < len(n); i++ {
			var xp XPerson
			err = GetXPerson(ctx, n[i].TCID, &xp)
			if err != nil {
				return err
			}

			r.T = append(r.T, xp)
		}
	}

	return err
}

// GetRentalAgreementEarliestDate returns the earliest of
// AgreementStart, PossessionStart, and RentStart
func GetRentalAgreementEarliestDate(ctx context.Context, a *RentalAgreement) (time.Time, error) {

	var (
		err error
		dt  time.Time
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return dt, ErrSessionRequired
		}
	}

	dt = a.AgreementStart
	if a.PossessionStart.Before(dt) {
		dt = a.PossessionStart
	}
	if a.RentStart.Before(dt) {
		dt = a.RentStart
	}
	return dt, err
}

// GetXRentalAgreement gets the RentalAgreement plus the associated rentables and payors for the
// time period specified
func GetXRentalAgreement(ctx context.Context, raid int64, d1, d2 *time.Time) (RentalAgreement, error) {

	var (
		// err error
		ra RentalAgreement
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ra, ErrSessionRequired
		}
	}

	return ra, LoadXRentalAgreement(ctx, raid, &ra, d1, d2)
}

// GetRentalAgreementsFromList takes an array of RentalAgreementRentables and returns map of
// all the rental agreements referenced. The map is indexed by the RAID
func GetRentalAgreementsFromList(ctx context.Context, raa *[]RentalAgreementRentable) (map[int64]RentalAgreement, error) {

	var (
		err error
		t   = make(map[int64]RentalAgreement)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	for i := 0; i < len(*raa); i++ {
		ra, err := GetRentalAgreement(ctx, (*raa)[i].RAID)
		if err != nil {
			return t, err
		}

		if ra.RAID > 0 {
			t[ra.RAID] = ra
		}
	}

	return t, err
}

// GetAgreementsForRentable returns an array of RentalAgreementRentables associated with the supplied RentableID
// during the time range d1-d2
func GetAgreementsForRentable(ctx context.Context, rid int64, d1, d2 *time.Time) ([]RentalAgreementRentable, error) {

	var (
		err error
		t   []RentalAgreementRentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{rid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementsForRentable)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentalAgreementsForRentable.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r RentalAgreementRentable
		err = ReadRentalAgreementRentables(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetRARentableForDate gets the RentalAgreementRentable plus the associated rentables and payors for the
// time period specified
func GetRARentableForDate(ctx context.Context, raid int64, d1 *time.Time, rar *RentalAgreementRentable) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid, d1, d1}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRARentableForDate)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRARentableForDate.QueryRow(fields...)
	}
	return ReadRentalAgreementRentable(row, rar)
}

// GetRentalAgreementRentable returns Rentable record matching the supplied RARID
func GetRentalAgreementRentable(ctx context.Context, rarid int64) (RentalAgreementRentable, error) {

	var (
		// err error
		r RentalAgreementRentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{rarid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementRentable)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreementRentable.QueryRow(fields...)
	}
	return r, ReadRentalAgreementRentable(row, &r)
}

// GetRentalAgreementRentables returns an array of RentalAgreementRentables associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentalAgreementRentables(ctx context.Context, raid int64, d1, d2 *time.Time) ([]RentalAgreementRentable, error) {

	var (
		err error
		t   []RentalAgreementRentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{raid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementRentables)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentalAgreementRentables.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r RentalAgreementRentable
		err = ReadRentalAgreementRentables(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetAllRentalAgreementRentables returns an array of RentalAgreementRentables
// associated with the supplied RentalAgreement ID
//
// INPUTS
//  ctx - context for txn
//  raid - which rental agreement
//
// RETURNS
//  an array of rental agr. rentables
//  any error encountered
//-----------------------------------------------------------------------------
func GetAllRentalAgreementRentables(ctx context.Context, raid int64) ([]RentalAgreementRentable, error) {

	var (
		err error
		t   []RentalAgreementRentable
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{raid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllRentalAgreementRentables)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllRentalAgreementRentables.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for rows.Next() {
		var r RentalAgreementRentable
		err = ReadRentalAgreementRentables(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

// GetRentalAgreementPayorByRBT returns Rental Agreement Payor record matching the supplied
// RAID, BID, TCID
func GetRentalAgreementPayorByRBT(ctx context.Context, raid, bid, tcid int64) (RentalAgreementPayor, error) {

	var (
		// err error
		r RentalAgreementPayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{raid, bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementPayorByRBT)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreementPayorByRBT.QueryRow(fields...)
	}
	return r, ReadRentalAgreementPayor(row, &r)
}

// GetRentalAgreementPayor returns Rental Agreement Payor record matching the supplied id
func GetRentalAgreementPayor(ctx context.Context, id int64) (RentalAgreementPayor, error) {

	var (
		// err error
		r RentalAgreementPayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementPayor)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreementPayor.QueryRow(fields...)
	}
	return r, ReadRentalAgreementPayor(row, &r)
}

// GetRentalAgreementPayorsInRange returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentalAgreementPayorsInRange(ctx context.Context, raid int64, d1, d2 *time.Time) ([]RentalAgreementPayor, error) {

	var (
		err error
		t   []RentalAgreementPayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{raid, d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementPayorsInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentalAgreementPayorsInRange.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getRentalAgreementPayorsByRows(ctx, rows)
}

// GetRentalAgreementsByPayor returns an array of RentalAgreementPayor where the supplied
// TCID is a payor on the specified date
func GetRentalAgreementsByPayor(ctx context.Context, bid, tcid int64, dt *time.Time) ([]RentalAgreementPayor, error) {

	var (
		err error
		t   []RentalAgreementPayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementsByPayor)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentalAgreementsByPayor.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getRentalAgreementPayorsByRows(ctx, rows)
}

// GetRentalAgreementsByPayorRange returns an array of RentalAgreementPayor where the supplied
// TCID is a payor within the supplied range
func GetRentalAgreementsByPayorRange(ctx context.Context, bid, tcid int64, d1, d2 *time.Time) ([]RentalAgreementPayor, error) {

	var (
		err error
		t   []RentalAgreementPayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid, tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementsByPayor)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetRentalAgreementsByPayor.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	return getRentalAgreementPayorsByRows(ctx, rows)
}

// getRentalAgreementPayorsByRows returns an array of RentalAgreementPayor records
// that were matched by the supplied sql.Rows
func getRentalAgreementPayorsByRows(ctx context.Context, rows *sql.Rows) ([]RentalAgreementPayor, error) {

	var (
		err error
		t   []RentalAgreementPayor
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}
	defer rows.Close()

	for rows.Next() {
		var r RentalAgreementPayor
		err = ReadRentalAgreementPayors(rows, &r)
		if err != nil {
			return t, err
		}
		t = append(t, r)
	}

	return t, rows.Err()
}

//=======================================================
//  RENTAL AGREEMENT TEMPLATE
//=======================================================

// GetRentalAgreementTemplate returns the RentalAgreementTemplate struct for the supplied rental agreement id
func GetRentalAgreementTemplate(ctx context.Context, ratid int64) (RentalAgreementTemplate, error) {

	var (
		// err error
		r RentalAgreementTemplate
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{ratid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementTemplate)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreementTemplate.QueryRow(fields...)
	}
	return r, ReadRentalAgreementTemplate(row, &r)
}

// GetRentalAgreementByRATemplateName returns the RentalAgreementTemplate struct for the supplied rental agreement id
func GetRentalAgreementByRATemplateName(ctx context.Context, ref string) (RentalAgreementTemplate, error) {

	var (
		// err error
		r RentalAgreementTemplate
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return r, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{ref}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetRentalAgreementByRATemplateName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetRentalAgreementByRATemplateName.QueryRow(fields...)
	}
	return r, ReadRentalAgreementTemplate(row, &r)
}

//=======================================================
//  STRING LIST
//=======================================================

// GetStringList reads a StringList structure based on the supplied StringList id
func GetStringList(ctx context.Context, id int64, a *StringList) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetStringList)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetStringList.QueryRow(fields...)
	}
	err = ReadStringList(row, a)
	if err != nil {
		return err
	}
	return getSLStrings(ctx, a.SLID, a)
}

// GetAllStringLists reads all StringList structures belonging to the business with the the supplied id
func GetAllStringLists(ctx context.Context, id int64) ([]StringList, error) {

	var (
		err error
		m   []StringList
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllStringLists)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllStringLists.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a StringList
		err = ReadStringLists(rows, &a)
		if err != nil {
			return m, err
		}

		err = getSLStrings(ctx, a.SLID, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetStringListByName reads a StringList structure based on the supplied StringList id
func GetStringListByName(ctx context.Context, bid int64, s string, a *StringList) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid, s}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetStringListByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetStringListByName.QueryRow(fields...)
	}
	err = ReadStringList(row, a)
	if err != nil {
		return err
	}

	return getSLStrings(ctx, a.SLID, a)
}

// GetSLStrings reads all strings with the supplid SLID into a
func getSLStrings(ctx context.Context, id int64, a *StringList) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetSLStrings)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetSLStrings.Query(fields...)
	}

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var p SLString
		err = ReadSLStrings(rows, &p)
		if err != nil {
			return err
		}
		a.S = append(a.S, p)
	}

	return rows.Err()
}

//=======================================================
//  SubAR
//=======================================================

// GetSubAR reads a SubAR structure based on the supplied SubAR id
func GetSubAR(ctx context.Context, id int64, a *SubAR) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetSubAR)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetSubAR.QueryRow(fields...)
	}
	return ReadSubAR(row, a)
}

// GetSubARs reads all SubAR structures belonging to the business with the the supplied id
func GetSubARs(ctx context.Context, id int64) ([]SubAR, error) {

	var (
		err error
		m   []SubAR
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetSubARs)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetSubARs.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a SubAR
		err = ReadSubARs(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

//============================================================
//  TASKS
//  TaskListDefintion, TaskListDescriptor, TaskList, Task
//============================================================

// GetTask returns the task with the supplied id
func GetTask(ctx context.Context, id int64) (Task, error) {
	var a Task
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTask)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTask.QueryRow(fields...)
	}
	return a, ReadTask(row, &a)
}

// GetLatestCompletedTaskList returns the latest completed task list
// with the parent or epoch equal to id
func GetLatestCompletedTaskList(ctx context.Context, id int64) (TaskList, error) {
	var a TaskList
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{id, id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetLatestCompletedTaskList)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetLatestCompletedTaskList.QueryRow(fields...)
	}
	return a, ReadTaskList(row, &a)
}

// GetTasks returns a slice of tasks with the supplied id
func GetTasks(ctx context.Context, id int64) ([]Task, error) {
	var m []Task
	var err error
	if sessionCheck(ctx) {
		return m, ErrSessionRequired
	}
	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTasks)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetTasks.Query(fields...)
	}
	if err != nil {
		return m, err
	}
	defer rows.Close()
	for rows.Next() {
		var a Task
		if err = ReadTasks(rows, &a); err != nil {
			return m, err
		}
		m = append(m, a)
	}
	return m, rows.Err()
}

// GetTaskList returns the tasklist with the supplied id
func GetTaskList(ctx context.Context, id int64) (TaskList, error) {
	var a TaskList
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTaskList)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTaskList.QueryRow(fields...)
	}
	return a, ReadTaskList(row, &a)
}

// GetTaskListInstanceInRange returns a tasklist instance where
// the PTLID matches the supplied ptlid and the due date falls in the
// supplied date range.
//
// INPUTS
//     ctx       - for transactions
//     id        = Parent TLID - the head of the task list instances
//     dt1 - dt2 = time period of instance for the due date
//
// RETURNS
//     the tasklist if found, or an empty task list if not found
//-----------------------------------------------------------------------------
func GetTaskListInstanceInRange(ctx context.Context, id int64, dt1, dt2 *time.Time) (TaskList, error) {
	var a TaskList
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{id, id, dt1, dt2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTaskListInstanceInRange)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTaskListInstanceInRange.QueryRow(fields...)
	}
	return a, ReadTaskList(row, &a)
}

// GetTaskDescriptor returns the tasklist with the supplied id
func GetTaskDescriptor(ctx context.Context, id int64) (TaskDescriptor, error) {
	var a TaskDescriptor
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTaskDescriptor)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTaskDescriptor.QueryRow(fields...)
	}
	return a, ReadTaskDescriptor(row, &a)
}

// GetTaskListDescriptors reads all TaskListDescriptor structures belonging to
// the TaskListDefinition with the the supplied id
//-----------------------------------------------------------------------------
func GetTaskListDescriptors(ctx context.Context, id int64) ([]TaskDescriptor, error) {
	var err error
	var m []TaskDescriptor
	if sessionCheck(ctx) {
		return m, ErrSessionRequired
	}

	var rows *sql.Rows
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTaskListDescriptors)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetTaskListDescriptors.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a TaskDescriptor
		if err = ReadTaskDescriptors(rows, &a); err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetAllTaskListDefinitions returns the active tasklist definitions for the
// supplied bid
//
// INPUTS:
//	id = BID of the business of interest
//
// RETURNS
//  slice of active TaskListDefintions defined for the business
//  any error encountered
//-----------------------------------------------------------------------------
func GetAllTaskListDefinitions(ctx context.Context, id int64) ([]TaskListDefinition, error) {
	var m []TaskListDefinition
	var err error

	if sessionCheck(ctx) {
		return m, ErrSessionRequired
	}

	var rows *sql.Rows
	fields := []interface{}{id}

	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetAllTaskListDefinitions)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetAllTaskListDefinitions.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var a TaskListDefinition
		if err = ReadTaskListDefinitions(rows, &a); err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetTaskListDefinition returns the tasklist with the supplied id
func GetTaskListDefinition(ctx context.Context, id int64) (TaskListDefinition, error) {
	var a TaskListDefinition
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTaskListDefinition)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTaskListDefinition.QueryRow(fields...)
	}
	return a, ReadTaskListDefinition(row, &a)
}

// GetTaskListDefinitionByName returns the tasklist with the supplied namd in the BID
func GetTaskListDefinitionByName(ctx context.Context, bid int64, name string) (TaskListDefinition, error) {
	var a TaskListDefinition
	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}
	var row *sql.Row
	fields := []interface{}{bid, name}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTaskListDefinitionByName)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTaskListDefinitionByName.QueryRow(fields...)
	}
	return a, ReadTaskListDefinition(row, &a)
}

// CheckForTLDInstances returns true if there are any instances of the
// supplied TLDID or false if there are none.
//
// INPUTS:
//    id = the TLDID to search for
//
// RETURNS
//    count: false if no instances found, true otherwise
//    error: any error encountered.
//-----------------------------------------------------------------------------
func CheckForTLDInstances(ctx context.Context, id int64) (bool, error) {
	var count int

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return false, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CheckForTLDInstances)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CheckForTLDInstances.QueryRow(fields...)
	}
	return (count > 0), row.Scan(&count)
}

//=======================================================
//  TRANSACTANT
//  Transactant, Prospect, User, Payor, XPerson
//=======================================================

// GetTransactantTypeDown returns the values needed for typedown controls:
// input:   bid - business
//            s - string or substring to search for
//        limit - return no more than this many matches
// return a slice of TransactantTypeDowns and an error.
func GetTransactantTypeDown(ctx context.Context, bid int64, s string, limit int) ([]TransactantTypeDown, error) {

	var (
		err error
		m   []TransactantTypeDown
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	s = "%" + s + "%"

	var rows *sql.Rows
	fields := []interface{}{bid, s, s, s, s, limit}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTransactantTypeDown)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetTransactantTypeDown.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var t TransactantTypeDown
		err = ReadTransactantTypeDowns(rows, &t)
		if err != nil {
			return m, err
		}
		m = append(m, t)
	}

	return m, rows.Err()
}

// GetTCIDByNote used to get TCID from Note Comment
// originally to get it from people csv Notes field
func GetTCIDByNote(ctx context.Context, cmt string) (int64, error) {

	var (
		// err  error
		tcid int64
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return tcid, ErrSessionRequired
		}
	}

	// just return first, in case of duplicate
	// TODO: need to verify
	var row *sql.Row
	fields := []interface{}{cmt}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.FindTCIDByNote)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.FindTCIDByNote.QueryRow(fields...)
	}
	return tcid, row.Scan(&tcid)
}

// GetTransactantByPhoneOrEmail searches for a transactoant match on the phone number or email
func GetTransactantByPhoneOrEmail(ctx context.Context, BID int64, s string) (Transactant, error) {

	var (
		// err error
		t Transactant
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	tq := `
	SELECT
		{{.SelectClause}}
	FROM
		Transactant
	WHERE
		BID={{.BID}} AND (
			WorkPhone="{{.WorkPhone}}" OR
			CellPhone="{{.CellPhone}}" OR
			PrimaryEmail="{{.PrimaryEmail}}" OR
			SecondaryEmail="{{.SecondaryEmail}}"
		);`

	qc := QueryClause{
		"SelectClause":   TRNSfields,
		"BID":            strconv.FormatInt(BID, 10),
		"WorkPhone":      s,
		"CellPhone":      s,
		"PrimaryEmail":   s,
		"SecondaryEmail": s,
	}

	// get formatted query
	p := RenderSQLQuery(tq, qc)

	// there could be multiple people with the same identifying number...
	// to safeguard, we'll read as a query and accept first match
	row := RRdb.Dbrr.QueryRow(p)
	return t, ReadTransactant(row, &t)
}

// GetTransactant reads a Transactant structure based on the supplied Transactant id
func GetTransactant(ctx context.Context, tid int64, t *Transactant) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{tid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetTransactant)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetTransactant.QueryRow(fields...)
	}
	return ReadTransactant(row, t)
}

// GetProspect reads a Prospect structure based on the supplied Transactant id
func GetProspect(ctx context.Context, id int64, p *Prospect) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{id}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetProspect)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetProspect.QueryRow(fields...)
	}
	return ReadProspect(row, p)
}

// GetUser reads a User structure based on the supplied User id.
// This call does not load the vehicle list.  You can use GetVehiclesByTransactant()
// if you need them.  Or you can call GetXPerson, which loads all details about a Transactant.
func GetUser(ctx context.Context, tcid int64, t *User) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetUser)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetUser.QueryRow(fields...)
	}
	return ReadUser(row, t)
}

// GetPayor reads a Payor structure based on the supplied Transactant id
func GetPayor(ctx context.Context, pid int64, p *Payor) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{pid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetPayor)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetPayor.QueryRow(fields...)
	}
	return ReadPayor(row, p)
}

/*// GetRentalAgreementGridInfo returns the array of rental agreement for grid
func GetRentalAgreementGridInfo(ctx context.Context, raid int64, d1, d2 *time.Time) ([]RentalAgreementGrid, error) {

	var (
		err error
		m   []RentalAgreementGrid
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	rows, err :=RRdb.Prepstmt.UIRAGrid(raid, d1, d2)
	if err != nil {
		return m, err
	}
	defer rows.Close()

	for rows.Next() {
		var t RentalAgreementGrid
		err = ReadRentalAgreementGrids(rows, &t)
		if err != nil {
			return m, err
		}
		m = append(m, &t)
	}

	return m, rows.Err()
}*/

// GetVehicle reads a Vehicle structure based on the supplied Vehicle id
func GetVehicle(ctx context.Context, vid int64, t *Vehicle) error {

	var (
	// err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{vid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetVehicle)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetVehicle.QueryRow(fields...)
	}
	return ReadVehicle(row, t)
}

func getVehicleList(ctx context.Context, rows *sql.Rows) ([]Vehicle, error) {

	var (
		err error
		m   []Vehicle
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return m, ErrSessionRequired
		}
	}

	for rows.Next() {
		var a Vehicle
		err = ReadVehicles(rows, &a)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}

	return m, rows.Err()
}

// GetVehiclesByLicensePlate reads a Vehicle structure based on the supplied Vehicle id
func GetVehiclesByLicensePlate(ctx context.Context, s string) ([]Vehicle, error) {

	var (
		err error
		t   []Vehicle
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{s}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetVehiclesByLicensePlate)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetVehiclesByLicensePlate.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	return getVehicleList(ctx, rows)
}

// GetVehiclesByTransactant reads a Vehicle structure based on the supplied Vehicle id
func GetVehiclesByTransactant(ctx context.Context, tcid int64) ([]Vehicle, error) {

	var (
		err error
		t   []Vehicle
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{tcid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetVehiclesByTransactant)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetVehiclesByTransactant.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	return getVehicleList(ctx, rows)
}

// GetVehiclesByBID reads a Vehicle structure based on the supplied Vehicle id
func GetVehiclesByBID(ctx context.Context, bid int64) ([]Vehicle, error) {

	var (
		err error
		t   []Vehicle
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetVehiclesByBID)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetVehiclesByBID.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	return getVehicleList(ctx, rows)
}

// GetXPerson will load a full XPerson given the trid
func GetXPerson(ctx context.Context, tcid int64, x *XPerson) error {

	var (
		err error
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return ErrSessionRequired
		}
	}

	if 0 == x.Trn.TCID {
		err = GetTransactant(ctx, tcid, &x.Trn)
		if err != nil {
			return err
		}
	}
	if 0 == x.Psp.TCID {
		err = GetProspect(ctx, tcid, &x.Psp)
		if err != nil {
			return err
		}
	}
	if 0 == x.Usr.TCID {
		err = GetUser(ctx, tcid, &x.Usr)
		if err != nil {
			return err
		}

		x.Usr.Vehicles, err = GetVehiclesByTransactant(ctx, tcid)
		if err != nil {
			return err
		}
	}
	if 0 == x.Pay.TCID {
		err = GetPayor(ctx, tcid, &x.Pay)
		if err != nil {
			return err
		}
	}

	return err
}

/*// GetDateOfLedgerMarkerOnOrBefore returns the Dt value of the last LedgerMarker set generated on or before d1
func GetDateOfLedgerMarkerOnOrBefore(ctx context.Context, bid int64, d1 *time.Time) (time.Time, err) {

	var (
		err error
		dt  time.Time
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return dt, ErrSessionRequired
		}
	}

	var lm LedgerMarker
	GenRcvLID := RRdb.BizTypes[bid].DefaultAccts[GLGENRCV].LID
	lm, err = GetLedgerMarkerOnOrBefore(ctx, bid, GenRcvLID, d1)

	// log the error
	if err != nil {
		Ulog("%s - SEVERE ERROR - unable to locate a LedgerMarker on or before %s\n", d1.Format(RRDATEFMT4))
	}

	return lm.Dt, err
}*/

// GetCountBusinessCustomAttrRefs get total count for CustomAttrRefs
// with particular associated business
func GetCountBusinessCustomAttrRefs(ctx context.Context, bid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountBusinessCustomAttrRefs)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountBusinessCustomAttrRefs.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetCountBusinessCustomAttributes get total count for CustomAttributes
// with particular associated business
func GetCountBusinessCustomAttributes(ctx context.Context, bid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountBusinessCustomAttributes)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountBusinessCustomAttributes.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetCountBusinessRentableTypes get total count for RentableTypes
// with particular associated business
func GetCountBusinessRentableTypes(ctx context.Context, bid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountBusinessRentableTypes)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountBusinessRentableTypes.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetCountBusinessTransactants get total count for Transactants
// with particular associated business
func GetCountBusinessTransactants(ctx context.Context, bid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountBusinessTransactants)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountBusinessTransactants.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetCountBusinessRentables get total count for Rentables
// with particular associated business
func GetCountBusinessRentables(ctx context.Context, bid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountBusinessRentables)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountBusinessRentables.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetCountBusinessRentalAgreements get total count for RentalAgreements
// with particular associated business
func GetCountBusinessRentalAgreements(ctx context.Context, bid int64) (int, error) {

	var (
		// err   error
		count int
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return count, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{bid}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.CountBusinessRentalAgreements)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.CountBusinessRentalAgreements.QueryRow(fields...)
	}
	return count, row.Scan(&count)
}

// GetFlow reads a Flow structure based on the supplied flowId
func GetFlow(ctx context.Context, flowID int64) (Flow, error) {

	var (
		// err error
		a Flow
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return a, ErrSessionRequired
		}
	}

	var row *sql.Row
	fields := []interface{}{flowID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetFlow)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetFlow.QueryRow(fields...)
	}
	return a, ReadFlow(row, &a)
}

// GetFlowForRAID reads a Flow structure based on the supplied
// FlowType and ID
//
// INPUTS:
//     flowtype = type of flow. "RA" or whatever
//           ID - the id that refers to a permanent table association.
//                for FlowType "RA", ID is the RAID
//
// RETURNS
//     The Flow struct
//     Any error encountered
//-----------------------------------------------------------------------------
func GetFlowForRAID(ctx context.Context, flowtype string, ID int64) (Flow, error) {
	var a Flow

	if sessionCheck(ctx) {
		return a, ErrSessionRequired
	}

	var row *sql.Row
	fields := []interface{}{flowtype, ID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetFlowForRAID)
		defer stmt.Close()
		row = stmt.QueryRow(fields...)
	} else {
		row = RRdb.Prepstmt.GetFlowForRAID.QueryRow(fields...)
	}
	return a, ReadFlow(row, &a)
}

// GetFlowsByFlowType reads all flowID for the current user
func GetFlowsByFlowType(ctx context.Context, flowType string) ([]Flow, error) {

	var (
		err error
		t   []Flow
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
	}

	var rows *sql.Rows
	fields := []interface{}{flowType}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetFlowsByFlowType)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetFlowsByFlowType.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var f Flow
		err = ReadFlows(rows, &f)
		if err != nil {
			return t, err
		}
		t = append(t, f)
	}
	return t, rows.Err()
}

// GetFlowIDsByUser reads all flowID for the current user
func GetFlowIDsByUser(ctx context.Context) ([]int64, error) {

	var (
		err error
		t   []int64
		UID = int64(0)
	)

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		sess, ok := SessionFromContext(ctx)
		if !ok {
			return t, ErrSessionRequired
		}
		UID = sess.UID
	}

	var rows *sql.Rows
	fields := []interface{}{UID}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetFlowIDsByUser)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetFlowIDsByUser.Query(fields...)
	}

	if err != nil {
		return t, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return t, err
		}
		t = append(t, id)
	}
	return t, rows.Err()
}

// GetFlowMetaDataInRange reads all flows struct between the supplied dates.
// The returned data does NOT include the JSON
//
// INPUTS
//     d1,d2 date range
//
// RETURNS
//     a slice of flows
//     any error encountered
//-----------------------------------------------------------------------------
func GetFlowMetaDataInRange(ctx context.Context, d1, d2 *time.Time) ([]Flow, error) {
	var err error
	var m []Flow

	if sessionCheck(ctx) {
		return m, ErrSessionRequired
	}

	var rows *sql.Rows
	fields := []interface{}{d1, d2}
	if tx, ok := DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(RRdb.Prepstmt.GetFlowMetaDataInRange)
		defer stmt.Close()
		rows, err = stmt.Query(fields...)
	} else {
		rows, err = RRdb.Prepstmt.GetFlowMetaDataInRange.Query(fields...)
	}

	if err != nil {
		return m, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var a Flow
		err = rows.Scan(&a.FlowID, &a.BID, &a.UserRefNo, &a.FlowType, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy)
		if err != nil {
			return m, err
		}
		m = append(m, a)
	}
	return m, rows.Err()
}
