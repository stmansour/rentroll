package rlib

import (
	"fmt"
	"time"
)

// GetTransactant reads a Transactant structure based on the supplied transactant id
func GetTransactant(tid int64, t *Transactant) {
	Errcheck(RRdb.Prepstmt.GetTransactant.QueryRow(tid).Scan(
		&t.TCID, &t.TID, &t.PID, &t.PRSPID, &t.FirstName, &t.MiddleName, &t.LastName, &t.PrimaryEmail,
		&t.SecondaryEmail, &t.WorkPhone, &t.CellPhone, &t.Address, &t.Address2, &t.City, &t.State,
		&t.PostalCode, &t.Country, &t.LastModTime, &t.LastModBy))
}

// GetProspect reads a Prospect structure based on the supplied transactant id
func GetProspect(prspid int64, p *Prospect) {
	Errcheck(RRdb.Prepstmt.GetProspect.QueryRow(prspid).Scan(&p.PRSPID, &p.TCID, &p.ApplicationFee))
}

// GetTenant reads a Tenant structure based on the supplied tenant id
func GetTenant(tcid int64, t *Tenant) {
	Errcheck(RRdb.Prepstmt.GetTransactant.QueryRow(tcid).Scan(&t.TID, &t.TCID, &t.Points,
		&t.CarMake, &t.CarModel, &t.CarColor, &t.CarYear, &t.LicensePlateState, &t.LicensePlateNumber,
		&t.ParkingPermitNumber, &t.AccountRep, &t.DateofBirth, &t.EmergencyContactName, &t.EmergencyContactAddress,
		&t.EmergencyContactTelephone, &t.EmergencyAddressEmail, &t.AlternateAddress, &t.ElibigleForFutureOccupancy,
		&t.Industry, &t.Source, &t.InvoicingCustomerNumber))
}

// GetPayor reads a Payor structure based on the supplied transactant id
func GetPayor(pid int64, p *Payor) {
	Errcheck(RRdb.Prepstmt.GetPayor.QueryRow(pid).Scan(
		&p.PID, &p.TCID, &p.CreditLimit, &p.EmployerName, &p.EmployerStreetAddress, &p.EmployerCity,
		&p.EmployerState, &p.EmployerZipcode, &p.Occupation, &p.LastModTime, &p.LastModBy))
}

// GetXPerson will load a full XPerson given the trid
func GetXPerson(tcid int64, x *XPerson) {
	if 0 == x.Trn.TCID {
		GetTransactant(tcid, &x.Trn)
	}
	if 0 == x.Psp.PRSPID && x.Trn.PRSPID > 0 {
		GetProspect(x.Trn.PRSPID, &x.Psp)
	}
	if 0 == x.Tnt.TID && x.Trn.TID > 0 {
		GetTenant(x.Trn.TID, &x.Tnt)
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

// GetRentableByID reads a Rentable structure based on the supplied rentable id
func GetRentableByID(rid int64, r *Rentable) {
	Errcheck(RRdb.Prepstmt.GetRentable.QueryRow(rid).Scan(&r.RID, &r.LID, &r.RTID, &r.BID, &r.UNITID, &r.Name, &r.Assignment, &r.Report, &r.DefaultOccType, &r.OccType, &r.LastModTime, &r.LastModBy))
}

// GetRentable reads and returns a Rentable structure based on the supplied rentable id
func GetRentable(rid int64) Rentable {
	var r Rentable
	Errcheck(RRdb.Prepstmt.GetRentable.QueryRow(rid).Scan(&r.RID, &r.LID, &r.RTID, &r.BID, &r.UNITID, &r.Name, &r.Assignment, &r.Report, &r.DefaultOccType, &r.OccType, &r.LastModTime, &r.LastModBy))
	return r
}

// GetUnit reads a Unit structure based on the supplied unit id
func GetUnit(uid int64, u *Unit) {
	Errcheck(RRdb.Prepstmt.GetUnit.QueryRow(uid).Scan(
		&u.UNITID, &u.BLDGID, &u.UTID, &u.RID, &u.AVAILID,
		&u.LastModTime, &u.LastModBy))
}

// GetXUnit reads an XUnit structure based on the RID.
func GetXUnit(rid int64, x *XUnit) {
	if x.R.RID == 0 && rid > 0 {
		GetRentableByID(rid, &x.R)
	}
	if x.U.UNITID == 0 && x.R.UNITID > 0 {
		GetUnit(x.R.UNITID, &x.U)
	}
	// fmt.Printf("GetXUnit:  bid = %d,  unitid = %d\n", x.R.BID, x.U.UNITID)
	x.S = GetUnitSpecialties(x.R.BID, x.U.UNITID)
}

// GetUnitSpecialties returns a list of specialties associated with the supplied unit
func GetUnitSpecialties(bid, unitid int64) []int64 {
	// first, get the specialties for this unit
	var m []int64
	rows, err := RRdb.Prepstmt.GetUnitSpecialties.Query(bid, unitid)
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

// GetUnitSpecialtyType returns a list of specialties associated with the supplied unit
func GetUnitSpecialtyType(uspid int64, ust *UnitSpecialtyType) {
	Errcheck(RRdb.Prepstmt.GetUnitSpecialtyType.QueryRow(uspid).Scan(&ust.USPID, &ust.BID, &ust.Name, &ust.Fee, &ust.Description))
}

// getUnitSpecialtiesTypes returns a list of UnitSpecialtyType structs associated with the supplied business
func getUnitSpecialtiesTypes(m *[]int64) map[int64]UnitSpecialtyType {
	// first, get the specialties for this unit
	var t map[int64]UnitSpecialtyType
	t = make(map[int64]UnitSpecialtyType, 0)

	for i := 0; i < len(*m); i++ {
		if _, ok := t[(*m)[i]]; !ok {
			var u UnitSpecialtyType
			GetUnitSpecialtyType((*m)[i], &u)
			t[(*m)[i]] = u
		}
	}
	return t
}

// GetRentableType returns characteristics of the unit
func GetRentableType(rtid int64, rt *RentableType) {
	Errcheck(RRdb.Prepstmt.GetRentableType.QueryRow(rtid).Scan(&rt.RTID, &rt.BID, &rt.Name, &rt.Frequency,
		&rt.Proration, &rt.Report, &rt.ManageToBudget, &rt.LastModTime, &rt.LastModBy))
}

// GetRentableTypeByName returns characteristics of the unit
func GetRentableTypeByName(name string, bid int64) (RentableType, error) {
	var rt RentableType
	err := RRdb.Prepstmt.GetRentableTypeByName.QueryRow(name, bid).Scan(&rt.RTID, &rt.BID, &rt.Name, &rt.Frequency,
		&rt.Proration, &rt.Report, &rt.ManageToBudget, &rt.LastModTime, &rt.LastModBy)
	return rt, err
}

// GetUnitType returns characteristics of the unit
func GetUnitType(utid int64, ut *UnitType) {
	Errcheck(RRdb.Prepstmt.GetUnitType.QueryRow(utid).Scan(&ut.UTID, &ut.BID, &ut.Style, &ut.Name, &ut.SqFt, &ut.Frequency, &ut.Proration, &ut.LastModTime, &ut.LastModBy))
}

// GetAssessmentTypeByName returns the record for the assessment type with the supplied name. If no such record exists or a database error occurred,
// the return structure will be empty
func GetAssessmentTypeByName(name string) (AssessmentType, error) {
	var t AssessmentType
	err := RRdb.Prepstmt.GetUnitType.QueryRow(name).Scan(&t.ASMTID, &t.Name, &t.Description, &t.LastModTime, &t.LastModBy)
	return t, err
}

// GetAssessmentTypes returns a slice of assessment types indexed by the ASMTID
func GetAssessmentTypes() map[int64]AssessmentType {
	var t map[int64]AssessmentType
	t = make(map[int64]AssessmentType, 0)
	rows, err := RRdb.dbrr.Query("SELECT ASMTID,Name,Description,LastModTime,LastModBy FROM assessmenttypes")
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a AssessmentType
		Errcheck(rows.Scan(&a.ASMTID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.ASMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

// GetSecurityDepositAssessments returns all the security deposit assessments for the supplied unit
func GetSecurityDepositAssessments(unitid int64) []Assessment {
	var m []Assessment
	rows, err := RRdb.Prepstmt.GetSecurityDepositAssessment.Query(unitid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a Assessment
		Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID,
			&a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency,
			&a.ProrationMethod, &a.AcctRule, &a.Comment))
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m
}

// GetPaymentTypes returns a slice of payment types indexed by the PMTID
func GetPaymentTypes() map[int64]PaymentType {
	var t map[int64]PaymentType
	t = make(map[int64]PaymentType, 0)
	rows, err := RRdb.dbrr.Query("SELECT PMTID,Name,Description,LastModTime,LastModBy FROM paymenttypes")
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a PaymentType
		Errcheck(rows.Scan(&a.PMTID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.PMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

// GetRentableMarketRates loads all the MarketRate rent information for this rentable into an array
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

// GetBusinessRentableTypes returns a slice of payment types indexed by the PMTID
func GetBusinessRentableTypes(bid int64) map[int64]RentableType {
	var t map[int64]RentableType
	t = make(map[int64]RentableType, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessRentableTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableType
		Errcheck(rows.Scan(&a.RTID, &a.BID, &a.Name, &a.Frequency, &a.Proration, &a.Report, &a.ManageToBudget, &a.LastModTime, &a.LastModBy))
		a.MR = make([]RentableMarketRate, 0)
		GetRentableMarketRates(&a)
		t[a.RTID] = a
	}
	Errcheck(rows.Err())

	return t
}

// GetRentableMarketRate returns the market-rate rent amount for r during the given time range. If the time range
// is large and spans multiple price changes, the chronologically earliest price that fits in the time range will be
// returned. It is best to provide as small a timerange d1-d2 as possible to minimize risk of overlap
func GetRentableMarketRate(xbiz *XBusiness, r *Rentable, d1, d2 *time.Time) float64 {
	if r.UNITID > 0 {
		var u Unit
		GetUnit(r.UNITID, &u)
		return GetUnitMarketRate(xbiz, &u, d1, d2)
	}
	// fmt.Printf("Get Market Rate for RTID = %d\n", r.RTID)
	mr := xbiz.RT[r.RTID].MR
	for i := 0; i < len(mr); i++ {
		if DateRangeOverlap(d1, d2, &mr[i].DtStart, &mr[i].DtStop) {
			return mr[i].MarketRate
		}
	}
	return float64(0)
}

// GetUnitMarketRates loads all the MarketRate rent information for this unit into an array
func GetUnitMarketRates(rt *UnitType) {
	// now get all the MarketRate rent info...
	rows, err := RRdb.Prepstmt.GetUnitMarketRates.Query(rt.UTID)
	Errcheck(err)
	defer rows.Close()
	LatestMRDTStart := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var a UnitMarketRate
		Errcheck(rows.Scan(&a.UTID, &a.MarketRate, &a.DtStart, &a.DtStop))
		if a.DtStart.After(LatestMRDTStart) {
			LatestMRDTStart = a.DtStart
			rt.MRCurrent = a.MarketRate
		}
		rt.MR = append(rt.MR, a)
	}
	Errcheck(rows.Err())
}

// GetBusinessUnitTypes returns a slice of payment types indexed by the PMTID
func GetBusinessUnitTypes(bid int64) map[int64]UnitType {
	var t map[int64]UnitType
	t = make(map[int64]UnitType, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessUnitTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a UnitType
		Errcheck(rows.Scan(&a.UTID, &a.BID, &a.Style, &a.Name, &a.SqFt, &a.Frequency, &a.Proration, &a.LastModTime, &a.LastModBy))
		GetUnitMarketRates(&a)
		t[a.UTID] = a
	}
	Errcheck(rows.Err())
	return t
}

// GetUnitMarketRate returns the market-rate rent amount for u during the given time range. If the time range
// is large and spans multiple price changes, the chronologically earliest price that fits in the time range will be
// returned. It is best to provide as small a timerange d1-d2 as possible to minimize risk of overlap
func GetUnitMarketRate(xbiz *XBusiness, u *Unit, d1, d2 *time.Time) float64 {
	mr := xbiz.UT[u.UTID].MR
	for i := 0; i < len(mr); i++ {
		if DateRangeOverlap(d1, d2, &mr[i].DtStart, &mr[i].DtStop) {
			// fmt.Printf("GetUnitMarketRate: returnning %f\n", mr[i].MarketRate)
			return mr[i].MarketRate
		}
	}
	return float64(0)
}

// GetBusiness loads the Business struct for the supplied business id
func GetBusiness(bid int64, p *Business) {
	Errcheck(RRdb.Prepstmt.GetBusiness.QueryRow(bid).Scan(&p.BID, &p.Designation,
		&p.Name, &p.DefaultOccupancyType, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy))
}

// GetBusinessByDesignation loads the Business struct for the supplied designation
func GetBusinessByDesignation(des string) (Business, error) {
	var p Business
	err := RRdb.Prepstmt.GetBusinessByDesignation.QueryRow(des).Scan(&p.BID, &p.Designation,
		&p.Name, &p.DefaultOccupancyType, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy)
	return p, err
}

// GetXBusiness loads the XBusiness struct for the supplied business id.
func GetXBusiness(bid int64, xbiz *XBusiness) {
	if xbiz.P.BID == 0 && bid > 0 {
		GetBusiness(bid, &xbiz.P)
	}
	xbiz.RT = GetBusinessRentableTypes(bid)
	xbiz.UT = GetBusinessUnitTypes(bid)
	xbiz.US = make(map[int64]UnitSpecialtyType, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessSpecialtyTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a UnitSpecialtyType
		Errcheck(rows.Scan(&a.USPID, &a.BID, &a.Name, &a.Fee, &a.Description))
		xbiz.US[a.USPID] = a
	}
	Errcheck(rows.Err())
}

// GetAllRentableAssessments for the supplied RID and date range
func GetAllRentableAssessments(RID int64, d1, d2 *time.Time) []Assessment {
	rows, err := RRdb.Prepstmt.GetAllRentableAssessments.Query(RID, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []Assessment
	t = make([]Assessment, 0)
	for i := 0; rows.Next(); i++ {
		var a Assessment
		Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID, &a.ASMTID,
			&a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod,
			&a.AcctRule, &a.Comment, &a.LastModTime, &a.LastModBy))
		t = append(t, a)
	}
	return t
}

// GetDefaultLedgerMarkers loads the default LedgerMarkers for the supplied Business bid
func GetDefaultLedgerMarkers(bid int64) LedgerMarker {
	var r LedgerMarker
	rows, err := RRdb.Prepstmt.GetDefaultLedgerMarkers.Query(bid)
	Errcheck(err)
	defer rows.Close()

	// fmt.Printf("GetDefaultLedgerMarkers: RRdb.BizTypes[%d].DefaultAccts = %#v\n", bid, RRdb.BizTypes[bid].DefaultAccts)
	// if nil == RRdb.BizTypes[bid].DefaultAccts {
	// 	RRdb.BizTypes[bid].DefaultAccts = make(map[int64]*LedgerMarker)
	// 	// fmt.Printf("GetDefaultLedgerMarkers: Setting RRdb.BizTypes[%d].DefaultAccts, val = %#v\n", bid, RRdb.BizTypes[bid].DefaultAccts)
	// }

	for rows.Next() {
		var r LedgerMarker
		Errcheck(rows.Scan(&r.LMID, &r.BID, &r.PID, &r.GLNumber, &r.State, &r.DtStart, &r.DtStop, &r.Balance, &r.Type, &r.Name))
		RRdb.BizTypes[bid].DefaultAccts[r.Type] = &r

		// fmt.Printf("GetDefaultLedgerMarkers just added: RRdb.BizTypes[%d].DefaultAccts[%d]\n", bid, r.Type)
		// pr := RRdb.BizTypes[bid].DefaultAccts[r.Type]
		// fmt.Printf("value = %#v\n", *pr)
	}
	return r
}

// GetAgreementsForRentable returns an array of AgreementRentables associated with the supplied RentableID
// during the time range d1-d2
func GetAgreementsForRentable(rid int64, d1, d2 *time.Time) []AgreementRentable {
	rows, err := RRdb.Prepstmt.GetAgreementRentables.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []AgreementRentable
	for rows.Next() {
		var r AgreementRentable
		Errcheck(rows.Scan(&r.RAID, &r.RID, &r.UNITID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetAgreementRentables returns an array of AgreementRentables associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetAgreementRentables(rid int64, d1, d2 *time.Time) []AgreementRentable {
	rows, err := RRdb.Prepstmt.GetAgreementRentables.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []AgreementRentable
	for rows.Next() {
		var r AgreementRentable
		Errcheck(rows.Scan(&r.RAID, &r.RID, &r.UNITID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetAgreementPayors returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetAgreementPayors(raid int64, d1, d2 *time.Time) []AgreementPayor {
	rows, err := RRdb.Prepstmt.GetAgreementPayors.Query(raid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []AgreementPayor
	t = make([]AgreementPayor, 0)
	for rows.Next() {
		var r AgreementPayor
		Errcheck(rows.Scan(&r.RAID, &r.PID, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRentalAgreement returns the RentalAgreement struct for the supplied rental agreement id
func GetRentalAgreement(raid int64) (RentalAgreement, error) {
	var r RentalAgreement
	// fmt.Printf("Ledger = %s\n", s)
	err := RRdb.Prepstmt.GetRentalAgreement.QueryRow(raid).Scan(&r.RAID, &r.RATID, &r.BID,
		/*&r.RID, &r.UNITID, &r.PID, &r.LID, */
		&r.PrimaryTenant, &r.RentalStart,
		&r.RentalStop, &r.Renewal, &r.SpecialProvisions, &r.LastModTime, &r.LastModBy)
	if nil != err {
		fmt.Printf("GetRentalAgreement: could not get rental agreement with raid = %d,  err = %v\n", raid, err)
	}
	return r, err
}

// GetXRentalAgreement gets the RentalAgreement plus the associated rentables and payors for the
// time period specified
func GetXRentalAgreement(raid int64, d1, d2 *time.Time) (RentalAgreement, error) {
	r, err := GetRentalAgreement(raid)
	t := GetAgreementRentables(raid, d1, d2)
	r.R = make([]XUnit, 0)
	for i := 0; i < len(t); i++ {
		var xu XUnit
		GetXUnit(t[i].RID, &xu)
		r.R = append(r.R, xu)
	}

	m := GetAgreementPayors(raid, d1, d2)
	r.P = make([]XPerson, 0)
	for i := 0; i < len(m); i++ {
		xp := GetXPersonByPID(m[i].PID)
		r.P = append(r.P, xp)
	}
	return r, err
}

// GetReceiptAllocations loads all receipt allocations associated with the supplied receipt id into
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

// GetReceipts for the supplied business (bid) in date range [d1 - d2)
func GetReceipts(bid int64, d1, d2 *time.Time) []Receipt {
	rows, err := RRdb.Prepstmt.GetReceiptsInDateRange.Query(bid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []Receipt
	t = make([]Receipt, 0)
	for rows.Next() {
		var r Receipt
		Errcheck(rows.Scan(
			&r.RCPTID, &r.BID, &r.RAID, &r.PMTID, &r.Dt, &r.Amount, &r.AcctRule, &r.Comment))
		r.RA = make([]ReceiptAllocation, 0)
		GetReceiptAllocations(r.RCPTID, &r)
		t = append(t, r)
	}
	return t
}

// GetReceipt returns a receipt structure for the supplied RCPTID
func GetReceipt(rcptid int64) Receipt {
	var r Receipt
	Errcheck(RRdb.Prepstmt.GetReceipt.QueryRow(rcptid).Scan(
		&r.RCPTID, &r.BID, &r.RAID, &r.PMTID, &r.Dt, &r.Amount, &r.AcctRule, &r.Comment))
	GetReceiptAllocations(rcptid, &r)
	return r
}

// GetAssessment returns the Assessment struct for the account with the supplied asmid
func GetAssessment(asmid int64) (Assessment, error) {
	var a Assessment
	err := RRdb.Prepstmt.GetAssessment.QueryRow(asmid).Scan(&a.ASMID, &a.BID, &a.RID,
		&a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency,
		&a.ProrationMethod, &a.AcctRule, &a.Comment, &a.LastModTime, &a.LastModBy)
	if nil != err {
		fmt.Printf("GetAssessment: could not get assessment with asmid = %d,  err = %v\n", asmid, err)
	}
	return a, err
}

// GetXType returns the RentalType structure and if it exists the UnitType structure is also returned
func GetXType(rtid, utid int64) XType {
	var xt XType
	GetRentableType(rtid, &xt.RT)
	GetUnitType(utid, &xt.UT)
	return xt
}

// GetJournalMarkers loads the last n journal markers
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

// GetLastJournalMarker returns the last journal marker or nil if no journal markers exist
func GetLastJournalMarker() JournalMarker {
	t := GetJournalMarkers(1)
	return t[0]
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
		fmt.Printf("GetJournal: could not get journal entry with jid = %d,  err = %v\n", jid, err)
	}
	return r, err
}

// GetLedgerMarkerInitList loads the Inititial LedgerMarker for all ledgers
// this is essentially a way to get the exhaustive list of ledger numbers for a business
func GetLedgerMarkerInitList(bid int64) []LedgerMarker {
	rows, err := RRdb.Prepstmt.GetLedgerMarkerInitList.Query(bid)
	Errcheck(err)
	defer rows.Close()
	var t []LedgerMarker
	t = make([]LedgerMarker, 0)
	for rows.Next() {
		var r LedgerMarker
		Errcheck(rows.Scan(&r.LMID, &r.BID, &r.PID, &r.GLNumber, &r.Status, &r.State, &r.DtStart, &r.DtStop, &r.Balance, &r.Type, &r.Name))
		t = append(t, r)
	}
	return t
}

// GetLedgerMarkers loads the last n Ledger markers for business BID
// func GetLedgerMarkers(bid, n int64) []LedgerMarker {
// 	rows, err := RRdb.Prepstmt.GetLedgerMarkerAcctList.Query(bid, n)
// 	Errcheck(err)
// 	defer rows.Close()
// 	var t []LedgerMarker
// 	t = make([]LedgerMarker, 0)
// 	for rows.Next() {
// 		var r LedgerMarker
// 		Errcheck(rows.Scan(&r.LMID, &r.BID, &r.PID, &r.GLNumber, &r.Status, &r.State, &r.DtStart, &r.DtStop, &r.Balance, &r.Type, &r.Name))
// 		t = append(t, r)
// 	}
// 	return t
// }

// // GetLastLedgerMarker returns the last Ledger Marker or nil if no journal markers exist
// func GetLastLedgerMarker(bid int64) LedgerMarker {
// 	t := GetLedgerMarkers(bid, 1)
// 	return t[0]
// }

// GetLedgerMarkerByGLNo returns the LedgerMarker struct for the GLNo with the supplied name
func GetLedgerMarkerByGLNo(bid int64, s string) LedgerMarker {
	var r LedgerMarker
	// fmt.Printf("Ledger = %s\n", s)
	err := RRdb.Prepstmt.GetLedgerMarkerByGLNo.QueryRow(bid, s).Scan(&r.LMID, &r.BID, &r.PID, &r.GLNumber, &r.Status, &r.State, &r.DtStart, &r.DtStop, &r.Balance, &r.Type, &r.Name)
	if nil != err {
		fmt.Printf("GetLedgerMarkerByGLNo: Could not find ledgermarker for GLNumber \"%s\".\n", s)
		fmt.Printf("err = %v\n", err)
	}
	return r
}

// GetLedgerMarkerByGLNoDateRange returns the LedgerMarker struct for the supplied time range
func GetLedgerMarkerByGLNoDateRange(bid int64, s string, d1, d2 *time.Time) LedgerMarker {
	var r LedgerMarker
	// fmt.Printf("Ledger = %s\n", s)
	err := RRdb.Prepstmt.GetLedgerMarkerByGLNoDateRange.QueryRow(bid, s, d1, d2).Scan(&r.LMID, &r.BID, &r.PID, &r.GLNumber, &r.Status, &r.State, &r.DtStart, &r.DtStop, &r.Balance, &r.Type, &r.Name)
	if nil != err {
		fmt.Printf("GetLedgerMarkerByGLNoDateRange: Could not find ledgermarker for GLNumber \"%s\".\n", s)
		fmt.Printf("err = %v\n", err)
	}
	return r
}

// GetLatestLedgerMarkerByGLNo returns the LedgerMarker struct for the GLNo with the supplied name
func GetLatestLedgerMarkerByGLNo(bid int64, s string) LedgerMarker {
	var r LedgerMarker
	// fmt.Printf("Ledger = %s\n", s)
	err := RRdb.Prepstmt.GetLatestLedgerMarkerByGLNo.QueryRow(bid, s).Scan(&r.LMID, &r.BID, &r.PID, &r.GLNumber, &r.Status, &r.State, &r.DtStart, &r.DtStop, &r.Balance, &r.Type, &r.Name)
	if nil != err {
		fmt.Printf("GetLatestLedgerMarkerByGLNo: Could not find ledgermarker for GLNumber \"%s\".\n", s)
		fmt.Printf("err = %v\n", err)
	}
	return r
}
