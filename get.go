package main

import (
	"rentroll/rlib"
	"time"
)

// GetTransactant reads a Transactant structure based on the supplied transactant id
func GetTransactant(tid int, t *Transactant) {
	rlib.Errcheck(App.prepstmt.getTransactant.QueryRow(tid).Scan(
		&t.TCID, &t.TID, &t.PID, &t.PRSPID, &t.FirstName, &t.MiddleName, &t.LastName, &t.PrimaryEmail,
		&t.SecondaryEmail, &t.WorkPhone, &t.CellPhone, &t.Address, &t.Address2, &t.City, &t.State,
		&t.PostalCode, &t.Country, &t.LastModTime, &t.LastModBy))
}

// GetProspect reads a Prospect structure based on the supplied transactant id
func GetProspect(prspid int, p *Prospect) {
	rlib.Errcheck(App.prepstmt.getProspect.QueryRow(prspid).Scan(&p.PRSPID, &p.TCID, &p.ApplicationFee))
}

// GetTenant reads a Tenant structure based on the supplied tenant id
func GetTenant(tcid int, t *Tenant) {
	rlib.Errcheck(App.prepstmt.getTransactant.QueryRow(tcid).Scan(&t.TID, &t.TCID, &t.Points,
		&t.CarMake, &t.CarModel, &t.CarColor, &t.CarYear, &t.LicensePlateState, &t.LicensePlateNumber,
		&t.ParkingPermitNumber, &t.AccountRep, &t.DateofBirth, &t.EmergencyContactName, &t.EmergencyContactAddress,
		&t.EmergencyContactTelephone, &t.EmergencyAddressEmail, &t.AlternateAddress, &t.ElibigleForFutureOccupancy,
		&t.Industry, &t.Source, &t.InvoicingCustomerNumber))
}

// GetPayor reads a Payor structure based on the supplied transactant id
func GetPayor(pid int, p *Payor) {
	rlib.Errcheck(App.prepstmt.getPayor.QueryRow(pid).Scan(
		&p.PID, &p.TCID, &p.CreditLimit, &p.EmployerName, &p.EmployerStreetAddress, &p.EmployerCity,
		&p.EmployerState, &p.EmployerZipcode, &p.Occupation, &p.LastModTime, &p.LastModBy))
}

// GetXPerson will load a full XPerson given the trid
func GetXPerson(tcid int, x *XPerson) {
	if 0 == x.trn.TCID {
		GetTransactant(tcid, &x.trn)
	}
	if 0 == x.psp.PRSPID && x.trn.PRSPID > 0 {
		GetProspect(x.trn.PRSPID, &x.psp)
	}
	if 0 == x.tnt.TID && x.trn.TID > 0 {
		GetTenant(x.trn.TID, &x.tnt)
	}
	if 0 == x.pay.PID && x.trn.PID > 0 {
		GetPayor(x.trn.PID, &x.pay)
	}
}

// GetRentable reads a Rentable structure based on the supplied rentable id
func GetRentable(rid int, r *Rentable) {
	rlib.Errcheck(App.prepstmt.getRentable.QueryRow(rid).Scan(&r.RID, &r.LID, &r.RTID, &r.BID, &r.PID, &r.RAID, &r.UNITID, &r.Name, &r.Assignment, &r.Report, &r.DefaultOccType, &r.OccType, &r.LastModTime, &r.LastModBy))
}

// GetUnit reads a Unit structure based on the supplied unit id
func GetUnit(uid int, u *Unit) {
	rlib.Errcheck(App.prepstmt.getUnit.QueryRow(uid).Scan(
		&u.UNITID, &u.BLDGID, &u.UTID, &u.RID, &u.AVAILID,
		&u.LastModTime, &u.LastModBy))
}

// GetXUnit reads an XUnit structure based on the RID.
func GetXUnit(rid int, x *XUnit) {
	if x.R.RID == 0 && rid > 0 {
		GetRentable(rid, &x.R)
	}
	if x.U.UNITID == 0 && x.R.UNITID > 0 {
		GetUnit(x.R.UNITID, &x.U)
	}
	// fmt.Printf("GetXUnit:  bid = %d,  unitid = %d\n", x.R.BID, x.U.UNITID)
	x.S = GetUnitSpecialties(x.R.BID, x.U.UNITID)
}

// GetUnitSpecialties returns a list of specialties associated with the supplied unit
func GetUnitSpecialties(bid, unitid int) []int {
	// first, get the specialties for this unit
	var m []int
	rows, err := App.prepstmt.getUnitSpecialties.Query(bid, unitid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var uspid int
		rlib.Errcheck(rows.Scan(&uspid))
		m = append(m, uspid)
	}
	rlib.Errcheck(rows.Err())
	return m
}

// GetUnitSpecialtyType returns a list of specialties associated with the supplied unit
func GetUnitSpecialtyType(uspid int, ust *UnitSpecialtyType) {
	rlib.Errcheck(App.prepstmt.getUnitSpecialtyType.QueryRow(uspid).Scan(&ust.USPID, &ust.BID, &ust.Name, &ust.Fee, &ust.Description))
}

// getUnitSpecialtiesTypes returns a list of UnitSpecialtyType structs associated with the supplied business
func getUnitSpecialtiesTypes(m *[]int) map[int]UnitSpecialtyType {
	// first, get the specialties for this unit
	var t map[int]UnitSpecialtyType
	t = make(map[int]UnitSpecialtyType, 0)

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
func GetRentableType(utid int, ut *RentableType) {
	rlib.Errcheck(App.prepstmt.getRentableType.QueryRow(utid).Scan(&ut.RTID, &ut.BID, &ut.Name, &ut.MarketRate, &ut.Frequency, &ut.Proration, &ut.LastModTime, &ut.LastModBy))
}

// GetUnitType returns characteristics of the unit
func GetUnitType(utid int, ut *UnitType) {
	rlib.Errcheck(App.prepstmt.getUnitType.QueryRow(utid).Scan(&ut.UTID, &ut.BID, &ut.Style, &ut.Name, &ut.SqFt, &ut.MarketRate, &ut.Frequency, &ut.Proration, &ut.LastModTime, &ut.LastModBy))
}

// GetAssessmentTypes returns a slice of assessment types indexed by the ASMTID
func GetAssessmentTypes() map[int]AssessmentType {
	var t map[int]AssessmentType
	t = make(map[int]AssessmentType, 0)
	rows, err := App.dbrr.Query("SELECT ASMTID,Name,Type,LastModTime,LastModBy FROM assessmenttypes")
	rlib.Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a AssessmentType
		rlib.Errcheck(rows.Scan(&a.ASMTID, &a.Name, &a.Type, &a.LastModTime, &a.LastModBy))
		t[a.ASMTID] = a
	}
	rlib.Errcheck(rows.Err())
	return t
}

// GetSecurityDepositAssessments returns all the security deposit assessments for the supplied unit
func GetSecurityDepositAssessments(unitid int) []Assessment {
	var m []Assessment
	rows, err := App.prepstmt.getSecurityDepositAssessment.Query(unitid)
	rlib.Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a Assessment
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod))
		m = append(m, a)
	}
	rlib.Errcheck(rows.Err())
	return m
}

// GetPaymentTypes returns a slice of payment types indexed by the PMTID
func GetPaymentTypes() map[int]PaymentType {
	var t map[int]PaymentType
	t = make(map[int]PaymentType, 0)
	rows, err := App.dbrr.Query("SELECT PMTID,Name,Description,LastModTime,LastModBy FROM paymenttypes")
	rlib.Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a PaymentType
		rlib.Errcheck(rows.Scan(&a.PMTID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.PMTID] = a
	}
	rlib.Errcheck(rows.Err())
	return t
}

// GetBusinessRentableTypes returns a slice of payment types indexed by the PMTID
func GetBusinessRentableTypes(bid int) map[int]RentableType {
	var t map[int]RentableType
	t = make(map[int]RentableType, 0)
	rows, err := App.prepstmt.getAllBusinessRentableTypes.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableType
		rlib.Errcheck(rows.Scan(&a.RTID, &a.BID, &a.Name, &a.MarketRate, &a.Frequency, &a.Proration, &a.LastModTime, &a.LastModBy))
		t[a.RTID] = a
	}
	rlib.Errcheck(rows.Err())
	return t
}

// GetBusinessUnitTypes returns a slice of payment types indexed by the PMTID
func GetBusinessUnitTypes(bid int) map[int]UnitType {
	var t map[int]UnitType
	t = make(map[int]UnitType, 0)
	rows, err := App.prepstmt.getAllBusinessUnitTypes.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a UnitType
		rlib.Errcheck(rows.Scan(&a.UTID, &a.BID, &a.Style, &a.Name, &a.SqFt, &a.MarketRate, &a.Frequency, &a.Proration, &a.LastModTime, &a.LastModBy))
		t[a.UTID] = a
	}
	rlib.Errcheck(rows.Err())
	return t
}

// GetBusiness loads the Business struct for the supplied business id
func GetBusiness(bid int, p *Business) {
	rlib.Errcheck(App.prepstmt.getPayor.QueryRow(bid).Scan(&p.BID, &p.Address, &p.Address2, &p.City,
		&p.State, &p.PostalCode, &p.Country, &p.Phone, &p.Name, &p.DefaultOccupancyType, &p.ParkingPermitInUse, &p.LastModTime, &p.LastModBy))
}

// GetXBusiness loads the XBusiness struct for the supplied business id.
func GetXBusiness(bid int, xprop *XBusiness) {
	if xprop.P.BID == 0 && bid > 0 {
		GetBusiness(bid, &xprop.P)
	}
	xprop.RT = GetBusinessRentableTypes(bid)
	xprop.UT = GetBusinessUnitTypes(bid)
	xprop.US = make(map[int]UnitSpecialtyType, 0)
	rows, err := App.prepstmt.getAllBusinessSpecialtyTypes.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a UnitSpecialtyType
		rlib.Errcheck(rows.Scan(&a.USPID, &a.BID, &a.Name, &a.Fee, &a.Description))
		xprop.US[a.USPID] = a
	}
	rlib.Errcheck(rows.Err())
}

// GetAllRentableAssessments for the supplied RID and date range
func GetAllRentableAssessments(RID int, d1, d2 *time.Time) []Assessment {
	rows, err := App.prepstmt.getAllRentableAssessments.Query(RID, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	var t []Assessment
	t = make([]Assessment, 0)
	for i := 0; rows.Next(); i++ {
		var a Assessment
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.RID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.LastModTime, &a.LastModBy))
		t = append(t, a)
	}
	return t
}
