package rlib

import (
	"fmt"
	"strings"
	"time"
)

// GetName for RentalAgreements returns a unique identifier string.
func (ra RentalAgreement) GetName() string {
	return fmt.Sprintf("RA%08d", ra.RAID)
}

//=======================================================
//  A G R E E M E N T   P E T S
//=======================================================

// GetRentalAgreementPet reads a Pet the structure for the supplied PETID
func GetRentalAgreementPet(petid int64) (RentalAgreementPet, error) {
	var a RentalAgreementPet

	err := RRdb.Prepstmt.GetRentalAgreementPet.QueryRow(petid).Scan(&a.PETID, &a.RAID, &a.Type, &a.Breed, &a.Color, &a.Weight, &a.Name, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy)
	return a, err
}

// GetAllRentalAgreementPets reads all Pet records for the supplied rental agreement id
func GetAllRentalAgreementPets(raid int64) []RentalAgreementPet {
	rows, err := RRdb.Prepstmt.GetAllRentalAgreementPets.Query(raid)
	Errcheck(err)
	defer rows.Close()
	var t []RentalAgreementPet
	for i := 0; rows.Next(); i++ {
		var a RentalAgreementPet
		Errcheck(rows.Scan(&a.PETID, &a.RAID, &a.Type, &a.Breed, &a.Color, &a.Weight, &a.Name, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy))
		t = append(t, a)
	}
	return t
}

//=======================================================
//  A G R E E M E N T   R E N T A B L E
//=======================================================

// FindAgreementByRentable reads a Prospect structure based on the supplied Transactant id
func FindAgreementByRentable(rid int64, d1, d2 *time.Time) (RentalAgreementRentable, error) {
	var a RentalAgreementRentable

	// SELECT RAID,RID,DtStart,DtStop from RentalAgreementRentables where RID=? and DtStop>=? and DtStart<=?

	err := RRdb.Prepstmt.FindAgreementByRentable.QueryRow(rid, d1, d2).Scan(&a.RAID, &a.RID, &a.DtStart, &a.DtStop)
	return a, err
}

//=======================================================
//  A S S E S S M E N T S
//=======================================================

// GetAllRentableAssessments for the supplied RID and date range
func GetAllRentableAssessments(RID int64, d1, d2 *time.Time) []Assessment {
	rows, err := RRdb.Prepstmt.GetAllRentableAssessments.Query(RID, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []Assessment
	for i := 0; rows.Next(); i++ {
		var a Assessment
		Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.ASMTID,
			&a.RAID, &a.Amount, &a.Start, &a.Stop, &a.RecurCycle, &a.ProrationCycle,
			&a.AcctRule, &a.Comment, &a.LastModTime, &a.LastModBy))
		t = append(t, a)
	}
	return t
}

// GetAssessment returns the Assessment struct for the account with the supplied asmid
func GetAssessment(asmid int64) (Assessment, error) {
	var a Assessment
	err := RRdb.Prepstmt.GetAssessment.QueryRow(asmid).Scan(&a.ASMID, &a.BID, &a.RID,
		&a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.RecurCycle,
		&a.ProrationCycle, &a.AcctRule, &a.Comment, &a.LastModTime, &a.LastModBy)
	if nil != err {
		Ulog("GetAssessment: could not get assessment with asmid = %d,  err = %v\n", asmid, err)
	}
	return a, err
}

//=======================================================
//  A S S E S S M E N T   T Y P E S
//=======================================================

// GetAssessmentTypeByName returns the record for the assessment type with the supplied name. If no such record exists or a database error occurred,
// the return structure will be empty
func GetAssessmentTypeByName(name string) (AssessmentType, error) {
	var t AssessmentType
	err := RRdb.Prepstmt.GetAssessmentTypeByName.QueryRow(name).Scan(&t.ASMTID, &t.RARequired, &t.Name, &t.Description, &t.LastModTime, &t.LastModBy)
	return t, err
}

// GetAssessmentTypes returns a slice of assessment types indexed by the ASMTID
func GetAssessmentTypes() map[int64]AssessmentType {
	var t map[int64]AssessmentType
	t = make(map[int64]AssessmentType, 0)
	rows, err := RRdb.Dbrr.Query("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM AssessmentTypes")
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a AssessmentType
		Errcheck(rows.Scan(&a.ASMTID, &a.RARequired, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.ASMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

//=======================================================
//  B U I L D I N G
//=======================================================

// GetBuilding returns the record for supplied bldg id. If no such record exists or a database error occurred,
// the return structure will be empty
func GetBuilding(id int64) Building {
	var t Building
	err := RRdb.Prepstmt.GetBuilding.QueryRow(id).Scan(&t.BLDGID, &t.BID, &t.Address, &t.Address2, &t.City, &t.State, &t.PostalCode, &t.Country, &t.LastModTime, &t.LastModBy)
	if err != nil {
		Ulog("GetBuilding: err = %v\n", err)
	}
	return t
}

//=======================================================
//  B U S I N E S S
//=======================================================

// GetBusiness loads the Business struct for the supplied Business id
func GetBusiness(bid int64, p *Business) {
	Errcheck(RRdb.Prepstmt.GetBusiness.QueryRow(bid).Scan(&p.BID, &p.Designation,
		&p.Name, &p.DefaultRentalPeriod, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy))
}

// GetBusinessByDesignation loads the Business struct for the supplied designation
func GetBusinessByDesignation(des string) (Business, error) {
	var p Business
	err := RRdb.Prepstmt.GetBusinessByDesignation.QueryRow(des).Scan(&p.BID, &p.Designation,
		&p.Name, &p.DefaultRentalPeriod, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy)
	return p, err
}

// GetXBusiness loads the XBusiness struct for the supplied Business id.
func GetXBusiness(bid int64, xbiz *XBusiness) {
	if xbiz.P.BID == 0 && bid > 0 {
		GetBusiness(bid, &xbiz.P)
	}
	xbiz.RT = GetBusinessRentableTypes(bid)
	xbiz.US = make(map[int64]RentableSpecialtyType, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessSpecialtyTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableSpecialtyType
		Errcheck(rows.Scan(&a.RSPID, &a.BID, &a.Name, &a.Fee, &a.Description))
		xbiz.US[a.RSPID] = a
	}
	Errcheck(rows.Err())
}

//=======================================================
//  C U S T O M   A T T R I B U T E
//  CustomAttribute, CustomAttributeRef
//=======================================================

// GetCustomAttribute reads a CustomAttribute structure based on the supplied CustomAttribute id
func GetCustomAttribute(cid int64) (CustomAttribute, error) {
	var a CustomAttribute
	err := RRdb.Prepstmt.GetCustomAttribute.QueryRow(cid).Scan(&a.CID, &a.Type, &a.Name, &a.Value, &a.LastModTime, &a.LastModBy)
	return a, err
}

// GetAllCustomAttributes returns a list of CustomAttributes for the supplied elementid and instanceid
func GetAllCustomAttributes(elemid, id int64) ([]CustomAttribute, error) {
	var t []int64
	var m []CustomAttribute
	rows, err := RRdb.Prepstmt.GetCustomAttributeRefs.Query(elemid, id)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var cid int64
		Errcheck(rows.Scan(&cid))
		t = append(t, cid)
	}
	Errcheck(rows.Err())

	for i := 0; i < len(t); i++ {
		var c CustomAttribute
		c, err := GetCustomAttribute(t[i])
		Errcheck(err)
		m = append(m, c)
	}

	return m, err
}

//=======================================================
//  L E D G E R   M A R K E R
//=======================================================

// GetLedgerMarkerByGLNoDateRange returns the LedgerMarker struct for the supplied time range
func GetLedgerMarkerByGLNoDateRange(bid int64, s string, d1, d2 *time.Time) (LedgerMarker, error) {
	var r LedgerMarker
	l, err := GetLedgerByGLNo(bid, s)
	if err == nil {
		err = RRdb.Prepstmt.GetLedgerMarkerByDateRange.QueryRow(bid, l.LID, d1, d2).Scan(&r.LMID, &r.LID, &r.BID, &r.DtStart, &r.DtStop, &r.Balance, &r.State, &r.LastModTime, &r.LastModBy)
		// if nil != err {
		// 	fmt.Printf("GetLedgerMarkerByGLNoDateRange: Could not find LedgerMarker for GLNumber \"%s\".\n", s)
		// 	fmt.Printf("err = %v\n", err)
		// }
	}
	return r, err
}

// GetLatestLedgerMarkerByLID returns the LedgerMarker struct for the GLAccount with the supplied LID
func GetLatestLedgerMarkerByLID(bid, lid int64) (LedgerMarker, error) {
	var r LedgerMarker
	err := RRdb.Prepstmt.GetLatestLedgerMarkerByLID.QueryRow(bid, lid).Scan(&r.LMID, &r.LID, &r.BID, &r.DtStart, &r.DtStop, &r.Balance, &r.State, &r.LastModTime, &r.LastModBy)
	if nil != err {
		if !IsSQLNoResultsError(err) {
			Ulog("GetLatestLedgerMarkerByGLNo: err = %v\n", err)
		}
	}
	return r, err
}

// GetLatestLedgerMarkerByGLNo returns the LedgerMarker struct for the GLNo with the supplied name
func GetLatestLedgerMarkerByGLNo(bid int64, s string) (LedgerMarker, error) {
	l, err := GetLedgerByGLNo(bid, s)
	if err != nil {
		var r LedgerMarker
		return r, err
	}
	return GetLatestLedgerMarkerByLID(bid, l.LID)
}

// GetLatestLedgerMarkerByType returns the LedgerMarker struct for the supplied type
func GetLatestLedgerMarkerByType(bid int64, t int64) (LedgerMarker, error) {
	var r LedgerMarker
	l, err := GetLedgerByType(bid, t)
	if err != nil {
		return r, err
	}
	return GetLatestLedgerMarkerByLID(bid, l.LID)
}

//=======================================================
//  P A Y M E N T   T Y P E S
//=======================================================

// GetPaymentTypes returns a slice of payment types indexed by the PMTID
func GetPaymentTypes() map[int64]PaymentType {
	var t map[int64]PaymentType
	t = make(map[int64]PaymentType, 0)
	rows, err := RRdb.Dbrr.Query("SELECT PMTID,BID,Name,Description,LastModTime,LastModBy FROM PaymentTypes")
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a PaymentType
		Errcheck(rows.Scan(&a.PMTID, &a.BID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.PMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

// GetPaymentTypesByBusiness returns a slice of payment types indexed by the PMTID for the supplied Business
func GetPaymentTypesByBusiness(bid int64) map[int64]PaymentType {
	var t map[int64]PaymentType
	t = make(map[int64]PaymentType, 0)
	rows, err := RRdb.Prepstmt.GetPaymentTypesByBusiness.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a PaymentType
		Errcheck(rows.Scan(&a.PMTID, &a.BID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.PMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

//=======================================================
//  R E N T A B L E
//=======================================================

// GetRentableByID reads a Rentable structure based on the supplied Rentable id
func GetRentableByID(rid int64, r *Rentable) {
	Errcheck(RRdb.Prepstmt.GetRentable.QueryRow(rid).Scan(&r.RID, &r.BID, &r.Name, &r.AssignmentTime, &r.LastModTime, &r.LastModBy))
}

// GetRentable reads and returns a Rentable structure based on the supplied Rentable id
func GetRentable(rid int64) Rentable {
	var r Rentable
	Errcheck(RRdb.Prepstmt.GetRentable.QueryRow(rid).Scan(&r.RID, &r.BID, &r.Name, &r.AssignmentTime, &r.LastModTime, &r.LastModBy))
	return r
}

// GetRentableByName reads and returns a Rentable structure based on the supplied Rentable id
func GetRentableByName(name string, bid int64) (Rentable, error) {
	var r Rentable
	err := RRdb.Prepstmt.GetRentableByName.QueryRow(name, bid).Scan(&r.RID, &r.BID, &r.Name, &r.AssignmentTime, &r.LastModTime, &r.LastModBy)
	return r, err
}

// GetXRentable reads an XRentable structure based on the RID.
func GetXRentable(rid int64, x *XRentable) {
	if x.R.RID == 0 && rid > 0 {
		GetRentableByID(rid, &x.R)
	}
	x.S = GetAllRentableSpecialtyRefs(x.R.BID, x.R.RID)
}

// GetRentableSpecialtyTypeByName returns a list of specialties associated with the supplied Rentable
func GetRentableSpecialtyTypeByName(bid int64, name string) RentableSpecialtyType {
	var rsp RentableSpecialtyType
	err := RRdb.Prepstmt.GetRentableSpecialtyTypeByName.QueryRow(bid, name).Scan(&rsp.RSPID, &rsp.BID, &rsp.Name, &rsp.Fee, &rsp.Description)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			fmt.Printf("GetRentableSpecialtyTypeByName: err = %v\n", err)
		}
	}
	return rsp
}

// GetRentableSpecialtyType returns the RentableSpecialtyType record for the supplied RSPID
func GetRentableSpecialtyType(rspid int64) (RentableSpecialtyType, error) {
	var rs RentableSpecialtyType
	err := RRdb.Prepstmt.GetRentableSpecialtyType.QueryRow(rspid).Scan(&rs.RSPID, &rs.BID, &rs.Name, &rs.Fee, &rs.Description)
	return rs, err
}

// GetAllRentableSpecialtyRefs returns a list of specialties associated with the supplied Rentable
func GetAllRentableSpecialtyRefs(bid, rid int64) []int64 {
	// first, get the specialties for this Rentable
	var m []int64
	rows, err := RRdb.Prepstmt.GetRentableSpecialtyRefs.Query(bid, rid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var uspid int64
		Errcheck(rows.Scan(&uspid))
		m = append(m, uspid)
	}
	Errcheck(rows.Err())
	return m
}

// // SelectRentableTypeRefForDate returns the first RTID of the list where the supplied date falls in range
// func SelectRentableTypeRefForDate(rsa *[]RentableSpecialtyType, dt *time.Time) RentableSpecialtyType {
// 	for i := 0; i < len(*rsa); i++ {
// 		if DateInRange(dt, &(*rsa)[i]. , &(*rsa)[i].DtStop) {
// 			return (*rsa)[i]
// 		}
// 	}
// 	var r RentableSpecialtyType
// 	return r // nothing matched
// }

// GetRentableSpecialtyTypesForRentableByRange returns an array of RentableSpecialtyType structures that
// apply for the supplied time range d1,d2
func GetRentableSpecialtyTypesForRentableByRange(r *Rentable, d1, d2 *time.Time) ([]RentableSpecialtyType, error) {
	var err error
	var rsta []RentableSpecialtyType
	rsrefs := GetRentableSpecialtyRefsByRange(r, d1, d2)
	for i := 0; i < len(rsrefs); i++ {
		rs, err := GetRentableSpecialtyType(rsrefs[i].RSPID)
		if err != nil {
			Ulog(err.Error())
			return rsta, err
		}
		rsta = append(rsta, rs)
	}
	return rsta, err
}

// GetRentableSpecialtyRefsByRange loads all the RentableSpecialtyRef records that overlap the supplied time range
// and returns them in an array
func GetRentableSpecialtyRefsByRange(r *Rentable, d1, d2 *time.Time) []RentableSpecialtyRef {
	var rs []RentableSpecialtyRef
	rows, err := RRdb.Prepstmt.GetRentableSpecialtyRefsByRange.Query(r.BID, r.RID, d1, d2)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableSpecialtyRef
		Errcheck(rows.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy))
		rs = append(rs, a)
	}
	Errcheck(rows.Err())
	return rs
}

// SelectRentableTypeRefForDate returns the first RTID of the list where the supplied date falls in range
func SelectRentableTypeRefForDate(rta *[]RentableTypeRef, dt *time.Time) RentableTypeRef {
	for i := 0; i < len(*rta); i++ {
		if DateInRange(dt, &(*rta)[i].DtStart, &(*rta)[i].DtStop) {
			return (*rta)[i]
		}
	}
	var r RentableTypeRef
	return r // nothing matched
}

// GetRentableTypeRefsByRange loads all the RentableTypeRef records that overlap the supplied time range
// and returns them in an array
func GetRentableTypeRefsByRange(RID int64, d1, d2 *time.Time) []RentableTypeRef {
	var rs []RentableTypeRef
	rows, err := RRdb.Prepstmt.GetRentableTypeRefsByRange.Query(RID, d1, d2)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableTypeRef
		Errcheck(rows.Scan(&a.RID, &a.RTID, &a.RentCycle, &a.ProrationCycle, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy))
		rs = append(rs, a)
	}
	Errcheck(rows.Err())
	return rs
}

// GetRTIDForDate returns the RTID in effect on the supplied date
func GetRTIDForDate(RID int64, d1 *time.Time) int64 {
	rtid := int64(0)
	DtStop, _ := StringToDate("1/1/9999")
	m := GetRentableTypeRefsByRange(RID, d1, &DtStop)
	if len(m) > 0 {
		rtid = m[0].RTID
	}
	return rtid
}

// GetRentableTypeRefForDate returns the RTID in effect on the supplied date
func GetRentableTypeRefForDate(RID int64, d1 *time.Time) RentableTypeRef {
	DtStop, _ := StringToDate("1/1/9999")
	m := GetRentableTypeRefsByRange(RID, d1, &DtStop)
	if len(m) > 0 {
		return m[0]
	}
	var r RentableTypeRef
	return r
}

// GetRentableStatusByRange loads all the RentableStatus records that overlap the supplied time range
func GetRentableStatusByRange(RID int64, d1, d2 *time.Time) []RentableStatus {
	var rs []RentableStatus
	rows, err := RRdb.Prepstmt.GetRentableStatusByRange.Query(RID, d1, d2)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableStatus
		Errcheck(rows.Scan(&a.RID, &a.DtStart, &a.DtStop, &a.Status, &a.LastModTime, &a.LastModBy))
		rs = append(rs, a)
	}
	Errcheck(rows.Err())
	return rs
}

//=======================================================
//  R E N T A B L E   T Y P E
//=======================================================

// GetRentableType returns characteristics of the Rentable
func GetRentableType(rtid int64, rt *RentableType) error {
	err := RRdb.Prepstmt.GetRentableType.QueryRow(rtid).Scan(&rt.RTID, &rt.BID, &rt.Style, &rt.Name, &rt.RentCycle,
		&rt.Proration, &rt.GSPRC, &rt.ManageToBudget, &rt.LastModTime, &rt.LastModBy)
	if nil == err {
		var cerr error
		rt.CA, cerr = GetAllCustomAttributes(ELEMRENTABLETYPE, rtid)
		if !IsSQLNoResultsError(cerr) { // it's not really an error if we don't find any custom attributes
			err = cerr
		}
	}
	return err
}

// GetRentableTypeByStyle returns characteristics of the Rentable
func GetRentableTypeByStyle(name string, bid int64) (RentableType, error) {
	var rt RentableType
	err := RRdb.Prepstmt.GetRentableTypeByStyle.QueryRow(name, bid).Scan(&rt.RTID, &rt.BID, &rt.Style, &rt.Name, &rt.RentCycle, &rt.Proration, &rt.GSPRC, &rt.ManageToBudget, &rt.LastModTime, &rt.LastModBy)
	return rt, err
}

// GetBusinessRentableTypes returns a slice of payment types indexed by the PMTID
func GetBusinessRentableTypes(bid int64) map[int64]RentableType {
	var t map[int64]RentableType
	t = make(map[int64]RentableType, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessRentableTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableType
		Errcheck(rows.Scan(&a.RTID, &a.BID, &a.Style, &a.Name, &a.RentCycle, &a.Proration, &a.GSPRC, &a.ManageToBudget, &a.LastModTime, &a.LastModBy))
		a.MR = make([]RentableMarketRate, 0)
		GetRentableMarketRates(&a)
		t[a.RTID] = a
	}
	Errcheck(rows.Err())

	return t
}

// GetRentableMarketRates loads all the MarketRate rent information for this Rentable into an array
func GetRentableMarketRates(rt *RentableType) {
	// now get all the MarketRate rent info...
	rows, err := RRdb.Prepstmt.GetRentableMarketRates.Query(rt.RTID)
	Errcheck(err)
	defer rows.Close()
	LatestMRDTStart := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var a RentableMarketRate
		Errcheck(rows.Scan(&a.RTID, &a.MarketRate, &a.DtStart, &a.DtStop))
		if a.DtStart.After(LatestMRDTStart) {
			LatestMRDTStart = a.DtStart
			rt.MRCurrent = a.MarketRate
		}
		rt.MR = append(rt.MR, a)
	}
	Errcheck(rows.Err())
}

// GetRentableMarketRate returns the market-rate rent amount for r during the given time range. If the time range
// is large and spans multiple price changes, the chronologically earliest price that fits in the time range will be
// returned. It is best to provide as small a timerange d1-d2 as possible to minimize risk of overlap
func GetRentableMarketRate(xbiz *XBusiness, r *Rentable, d1, d2 *time.Time) float64 {
	rtid := GetRTIDForDate(r.RID, d1) // first thing... find the RTID for this time range
	mr := xbiz.RT[rtid].MR
	// fmt.Printf("GetRentableMarketRate: Get Market Rate for RTID = %d, %s - %s\n", rtid, d1.Format(RRDATEINPFMT), d2.Format(RRDATEINPFMT))
	for i := 0; i < len(mr); i++ {
		if DateRangeOverlap(d1, d2, &mr[i].DtStart, &mr[i].DtStop) {
			return mr[i].MarketRate
		}
	}
	return float64(0)
}

// GetRentalAgreementsFromList takes an array of RentalAgreementRentables and returns map of
// all the rental agreements referenced. The map is indexed by the RAID
func GetRentalAgreementsFromList(raa *[]RentalAgreementRentable) map[int64]RentalAgreement {
	var t map[int64]RentalAgreement
	for i := 0; i < len(*raa); i++ {
		ra, err := GetRentalAgreement((*raa)[i].RAID)
		Errlog(err)
		if ra.RAID > 0 {
			t[ra.RAID] = ra
		}
	}
	return t
}

// GetAgreementsForRentable returns an array of RentalAgreementRentables associated with the supplied RentableID
// during the time range d1-d2
func GetAgreementsForRentable(rid int64, d1, d2 *time.Time) []RentalAgreementRentable {
	rows, err := RRdb.Prepstmt.GetRentalAgreementRentables.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentalAgreementRentable
	for rows.Next() {
		var r RentalAgreementRentable
		Errcheck(rows.Scan(&r.RAID, &r.RID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRentalAgreementRentables returns an array of RentalAgreementRentables associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentalAgreementRentables(rid int64, d1, d2 *time.Time) []RentalAgreementRentable {
	rows, err := RRdb.Prepstmt.GetRentalAgreementRentables.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentalAgreementRentable
	for rows.Next() {
		var r RentalAgreementRentable
		Errcheck(rows.Scan(&r.RAID, &r.RID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRentalAgreementPayors returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentalAgreementPayors(raid int64, d1, d2 *time.Time) []RentalAgreementPayor {
	rows, err := RRdb.Prepstmt.GetRentalAgreementPayors.Query(raid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentalAgreementPayor
	t = make([]RentalAgreementPayor, 0)
	for rows.Next() {
		var r RentalAgreementPayor
		Errcheck(rows.Scan(&r.RAID, &r.PID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRentableUsers returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentableUsers(raid int64, d1, d2 *time.Time) []RentableUser {
	rows, err := RRdb.Prepstmt.GetRentableUsers.Query(raid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentableUser
	// t = make([]RentableUser, 0)
	for rows.Next() {
		var r RentableUser
		Errcheck(rows.Scan(&r.RID, &r.USERID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

//=======================================================
//  R E N T A L   A G R E E M E N T
//=======================================================

// GetRentalAgreement returns the RentalAgreement struct for the supplied rental agreement id
func GetRentalAgreement(raid int64) (RentalAgreement, error) {
	var r RentalAgreement
	err := RRdb.Prepstmt.GetRentalAgreement.QueryRow(raid).Scan(&r.RAID, &r.RATID, &r.BID,
		&r.RentalStart, &r.RentalStop, &r.PossessionStart, &r.PossessionStop,
		&r.Renewal, &r.SpecialProvisions, &r.LastModTime, &r.LastModBy)
	if nil != err && !IsSQLNoResultsError(err) {
		fmt.Printf("GetRentalAgreement: could not get rental agreement with raid = %d,  err = %v\n", raid, err)
	}
	return r, err
}

// GetXRentalAgreement gets the RentalAgreement plus the associated rentables and payors for the
// time period specified
func GetXRentalAgreement(raid int64, d1, d2 *time.Time) (RentalAgreement, error) {
	r, err := GetRentalAgreement(raid)

	t := GetRentalAgreementRentables(raid, d1, d2)
	r.R = make([]XRentable, 0)
	for i := 0; i < len(t); i++ {
		var xu XRentable
		GetXRentable(t[i].RID, &xu)
		r.R = append(r.R, xu)
	}

	m := GetRentalAgreementPayors(raid, d1, d2)
	r.P = make([]XPerson, 0)
	for i := 0; i < len(m); i++ {
		xp := GetXPersonByPID(m[i].PID)
		r.P = append(r.P, xp)
	}

	n := GetRentableUsers(raid, d1, d2)
	r.T = make([]XPerson, 0)
	for i := 0; i < len(n); i++ {
		xp := GetXPersonByTID(n[i].USERID)
		r.T = append(r.T, xp)
	}
	return r, err
}

//=======================================================
//  RENTAL AGREEMENT TEMPLATE
//=======================================================

// GetRentalAgreementTemplate returns the RentalAgreementTemplate struct for the supplied rental agreement id
func GetRentalAgreementTemplate(ratid int64) (RentalAgreementTemplate, error) {
	var r RentalAgreementTemplate
	err := RRdb.Prepstmt.GetRentalAgreementTemplate.QueryRow(ratid).Scan(&r.RATID, &r.BID, &r.RentalTemplateNumber, &r.LastModTime, &r.LastModBy)
	if nil != err {
		Ulog("GetRentalAgreementTemplate: could not get rental agreement template with RATID = %d,  err = %v\n", ratid, err)
	}
	return r, err
}

// GetRentalAgreementByRentalTemplateNumber returns the RentalAgreementTemplate struct for the supplied rental agreement id
func GetRentalAgreementByRentalTemplateNumber(ref string) (RentalAgreementTemplate, error) {
	var r RentalAgreementTemplate
	err := RRdb.Prepstmt.GetRentalAgreementByRentalTemplateNumber.QueryRow(ref).Scan(&r.RATID, &r.BID, &r.RentalTemplateNumber, &r.LastModTime, &r.LastModBy)
	return r, err
}

//=======================================================
//  RECEIPT ALLOCATION
//=======================================================

// GetReceiptAllocations loads all Receipt allocations associated with the supplied Receipt id into
// the RA array within a Receipt structure
func GetReceiptAllocations(rcptid int64, r *Receipt) {
	rows, err := RRdb.Prepstmt.GetReceiptAllocations.Query(rcptid)
	Errcheck(err)
	defer rows.Close()
	r.RA = make([]ReceiptAllocation, 0)
	for rows.Next() {
		var a ReceiptAllocation
		Errcheck(rows.Scan(&a.RCPTID, &a.Amount, &a.ASMID, &a.AcctRule))
		r.RA = append(r.RA, a)
	}
}

// GetReceipts for the supplied Business (bid) in date range [d1 - d2)
func GetReceipts(bid int64, d1, d2 *time.Time) []Receipt {
	rows, err := RRdb.Prepstmt.GetReceiptsInDateRange.Query(bid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []Receipt
	t = make([]Receipt, 0)
	for rows.Next() {
		var r Receipt
		Errcheck(rows.Scan(
			&r.RCPTID, &r.BID, &r.RAID, &r.PMTID, &r.Dt, &r.Amount, &r.AcctRule, &r.Comment, &r.LastModTime, &r.LastModBy))
		r.RA = make([]ReceiptAllocation, 0)
		GetReceiptAllocations(r.RCPTID, &r)
		t = append(t, r)
	}
	return t
}

// GetReceipt returns a Receipt structure for the supplied RCPTID
func GetReceipt(rcptid int64) Receipt {
	var r Receipt
	Errcheck(RRdb.Prepstmt.GetReceipt.QueryRow(rcptid).Scan(
		&r.RCPTID, &r.BID, &r.RAID, &r.PMTID, &r.Dt, &r.Amount, &r.AcctRule, &r.Comment, &r.LastModTime, &r.LastModBy))
	GetReceiptAllocations(rcptid, &r)
	return r
}

// GetJournalMarkers loads the last n Journal markers
func GetJournalMarkers(n int64) []JournalMarker {
	rows, err := RRdb.Prepstmt.GetJournalMarkers.Query(n)
	Errcheck(err)
	defer rows.Close()
	var t []JournalMarker
	t = make([]JournalMarker, 0)
	for rows.Next() {
		var r JournalMarker
		Errcheck(rows.Scan(&r.JMID, &r.BID, &r.State, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetLastJournalMarker returns the last Journal marker or nil if no Journal markers exist
func GetLastJournalMarker() JournalMarker {
	t := GetJournalMarkers(1)
	if len(t) > 0 {
		return t[0]
	}
	var j JournalMarker
	return j
}

// GetJournalAllocation returns the Journal allocation for the supplied JAID
func GetJournalAllocation(jaid int64) (JournalAllocation, error) {
	var a JournalAllocation
	err := RRdb.Prepstmt.GetJournalAllocation.QueryRow(jaid).Scan(&a.JAID, &a.JID, &a.RID, &a.Amount, &a.ASMID, &a.AcctRule)
	if err != nil {
		Ulog("Error getting JournalAllocation jaid = %d:  error = %v\n", jaid, err)
	}
	return a, err
}

// GetJournalAllocations loads all Journal allocations associated with the supplied Journal id into
// the RA array within a Journal structure
func GetJournalAllocations(jid int64, j *Journal) {
	rows, err := RRdb.Prepstmt.GetJournalAllocations.Query(jid)
	Errcheck(err)
	defer rows.Close()
	j.JA = make([]JournalAllocation, 0)
	for rows.Next() {
		var a JournalAllocation
		Errcheck(rows.Scan(&a.JAID, &a.JID, &a.RID, &a.Amount, &a.ASMID, &a.AcctRule))
		j.JA = append(j.JA, a)
	}
}

// GetJournal returns the Journal struct for the account with the supplied name
func GetJournal(jid int64) (Journal, error) {
	var r Journal
	err := RRdb.Prepstmt.GetJournal.QueryRow(jid).Scan(&r.JID, &r.BID, &r.RAID,
		&r.Dt, &r.Amount, &r.Type, &r.ID, &r.Comment, &r.LastModTime, &r.LastModBy)
	if nil != err {
		fmt.Printf("GetJournal: could not get Journal entry with jid = %d,  err = %v\n", jid, err)
	}
	return r, err
}

//=======================================================
//  L E D G E R
//=======================================================

// GetLedgerList returns an array of all GLAccount
// this is essentially a way to get the exhaustive list of GLAccount numbers for a Business
func GetLedgerList(bid int64) []GLAccount {
	rows, err := RRdb.Prepstmt.GetLedgerList.Query(bid)
	Errcheck(err)
	defer rows.Close()
	var t []GLAccount
	for rows.Next() {
		var r GLAccount
		Errcheck(rows.Scan(&r.LID, &r.PLID, &r.BID, &r.RAID, &r.GLNumber, &r.Status, &r.Type, &r.Name, &r.AcctType, &r.RAAssociated, &r.AllowPost, &r.LastModTime, &r.LastModBy))
		t = append(t, r)
	}
	return t
}

// GetLedger returns the GLAccount struct for the supplied LID
func GetLedger(lid int64) (GLAccount, error) {
	var r GLAccount
	err := RRdb.Prepstmt.GetLedger.QueryRow(lid).Scan(&r.LID, &r.PLID, &r.BID, &r.RAID, &r.GLNumber, &r.Status, &r.Type, &r.Name, &r.AcctType, &r.RAAssociated, &r.AllowPost, &r.LastModTime, &r.LastModBy)
	return r, err
}

// GetLedgerByGLNo returns the GLAccount struct for the supplied GLNo
func GetLedgerByGLNo(bid int64, s string) (GLAccount, error) {
	var r GLAccount
	err := RRdb.Prepstmt.GetLedgerByGLNo.QueryRow(bid, s).Scan(&r.LID, &r.PLID, &r.BID, &r.RAID, &r.GLNumber, &r.Status, &r.Type, &r.Name, &r.AcctType, &r.RAAssociated, &r.AllowPost, &r.LastModTime, &r.LastModBy)
	return r, err
}

// GetLedgerByType returns the GLAccount struct for the supplied Type
func GetLedgerByType(bid, t int64) (GLAccount, error) {
	var r GLAccount
	err := RRdb.Prepstmt.GetLedgerByType.QueryRow(bid, t).Scan(&r.LID, &r.PLID, &r.BID, &r.RAID, &r.GLNumber, &r.Status, &r.Type, &r.Name, &r.AcctType, &r.RAAssociated, &r.AllowPost, &r.LastModTime, &r.LastModBy)
	return r, err
}

// GetRABalanceLedger returns the GLAccount struct for the supplied Type
func GetRABalanceLedger(bid, RAID int64) (GLAccount, error) {
	var r GLAccount
	err := RRdb.Prepstmt.GetRABalanceLedger.QueryRow(bid, RAID).Scan(&r.LID, &r.PLID, &r.BID, &r.RAID, &r.GLNumber, &r.Status, &r.Type, &r.Name, &r.AcctType, &r.RAAssociated, &r.AllowPost, &r.LastModTime, &r.LastModBy)
	return r, err
}

// GetDefaultLedgers loads the default GLAccount for the supplied Business bid
func GetDefaultLedgers(bid int64) {
	rows, err := RRdb.Prepstmt.GetDefaultLedgers.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r GLAccount
		Errcheck(rows.Scan(&r.LID, &r.PLID, &r.BID, &r.RAID, &r.GLNumber, &r.Status, &r.Type, &r.Name, &r.AcctType, &r.RAAssociated, &r.AllowPost, &r.LastModTime, &r.LastModBy))
		RRdb.BizTypes[bid].DefaultAccts[r.Type] = &r
	}
}

//=======================================================
//  LEDGER ENTRUY
//=======================================================

// GetLedgerEntriesForRAID returns a list of CustomAttributes for the supplied elementid and instanceid
func GetLedgerEntriesForRAID(d1, d2 *time.Time, raid, lid int64) ([]LedgerEntry, error) {
	var m []LedgerEntry
	rows, err := RRdb.Prepstmt.GetLedgerEntriesForRAID.Query(d1, d2, raid, lid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var le LedgerEntry
		Errcheck(rows.Scan(&le.LEID, &le.BID, &le.JID, &le.JAID, &le.LID, &le.RAID, &le.Dt, &le.Amount, &le.Comment, &le.LastModTime, &le.LastModBy))
		m = append(m, le)
	}
	Errcheck(rows.Err())
	return m, err
}

//=======================================================
//  TRANSACTANT
//  Transactant, Prospect, User, Payor, XPerson
//=======================================================

// GetTransactantByPhoneOrEmail searches for a transactoant match on the phone number or email
func GetTransactantByPhoneOrEmail(s string) (Transactant, error) {
	var t Transactant
	p := fmt.Sprintf("SELECT "+TRNSfields+" FROM Transactant where WorkPhone=\"%s\" or CellPhone=\"%s\" or PrimaryEmail=\"%s\" or SecondaryEmail=\"%s\"", s, s, s, s)
	err := RRdb.Dbrr.QueryRow(p).Scan(&t.TCID, &t.USERID, &t.PID, &t.PRSPID, &t.FirstName, &t.MiddleName, &t.LastName,
		&t.PreferredName, &t.CompanyName, &t.IsCompany, &t.PrimaryEmail, &t.SecondaryEmail, &t.WorkPhone, &t.CellPhone,
		&t.Address, &t.Address2, &t.City, &t.State, &t.PostalCode, &t.Country, &t.Website, &t.Notes, &t.LastModTime, &t.LastModBy)
	return t, err
}

// GetTransactant reads a Transactant structure based on the supplied Transactant id
func GetTransactant(tid int64, t *Transactant) {
	Errcheck(RRdb.Prepstmt.GetTransactant.QueryRow(tid).Scan(&t.TCID, &t.USERID, &t.PID, &t.PRSPID, &t.FirstName, &t.MiddleName, &t.LastName, &t.PreferredName, &t.CompanyName, &t.IsCompany, &t.PrimaryEmail, &t.SecondaryEmail, &t.WorkPhone, &t.CellPhone, &t.Address, &t.Address2, &t.City, &t.State, &t.PostalCode, &t.Country, &t.Website, &t.Notes, &t.LastModTime, &t.LastModBy))
}

// GetProspect reads a Prospect structure based on the supplied Transactant id
func GetProspect(prspid int64, p *Prospect) {
	Errcheck(RRdb.Prepstmt.GetProspect.QueryRow(prspid).Scan(&p.PRSPID, &p.TCID, &p.EmployerName, &p.EmployerStreetAddress,
		&p.EmployerCity, &p.EmployerState, &p.EmployerPostalCode, &p.EmployerEmail, &p.EmployerPhone, &p.Occupation,
		&p.ApplicationFee, &p.LastModTime, &p.LastModBy))
}

// GetUser reads a User structure based on the supplied User id
func GetUser(tcid int64, t *User) {
	Errcheck(RRdb.Prepstmt.GetUser.QueryRow(tcid).Scan(&t.USERID, &t.TCID, &t.Points, &t.CarMake, &t.CarModel, &t.CarColor, &t.CarYear, &t.LicensePlateState, &t.LicensePlateNumber, &t.ParkingPermitNumber, &t.DateofBirth, &t.EmergencyContactName, &t.EmergencyContactAddress, &t.EmergencyContactTelephone, &t.EmergencyEmail, &t.AlternateAddress, &t.EligibleFutureUser, &t.Industry, &t.Source, &t.LastModTime, &t.LastModBy))
}

// GetPayor reads a Payor structure based on the supplied Transactant id
func GetPayor(pid int64, p *Payor) {
	Errcheck(RRdb.Prepstmt.GetPayor.QueryRow(pid).Scan(&p.PID, &p.TCID, &p.CreditLimit, &p.TaxpayorID, &p.AccountRep, &p.EligibleFuturePayor, &p.LastModTime, &p.LastModBy))
}

// GetXPerson will load a full XPerson given the trid
func GetXPerson(tcid int64, x *XPerson) {
	if 0 == x.Trn.TCID {
		GetTransactant(tcid, &x.Trn)
	}
	if 0 == x.Psp.PRSPID && x.Trn.PRSPID > 0 {
		GetProspect(x.Trn.PRSPID, &x.Psp)
	}
	if 0 == x.Tnt.USERID && x.Trn.USERID > 0 {
		GetUser(x.Trn.USERID, &x.Tnt)
	}
	if 0 == x.Pay.PID && x.Trn.PID > 0 {
		GetPayor(x.Trn.PID, &x.Pay)
	}
}

// GetXPersonByPID will load a full XPerson given the PID
func GetXPersonByPID(pid int64) XPerson {
	var xp XPerson
	GetPayor(pid, &xp.Pay)
	GetXPerson(xp.Pay.TCID, &xp)
	return xp
}

// GetXPersonByTID will load a full XPerson given the USERID
func GetXPersonByTID(tid int64) XPerson {
	var xp XPerson
	GetUser(tid, &xp.Tnt)
	GetXPerson(xp.Tnt.TCID, &xp)
	return xp
}
