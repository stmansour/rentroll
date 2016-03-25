package main

import (
	"fmt"
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

// GetXPersonByPID will load a full XPerson given the PID
func GetXPersonByPID(pid int) XPerson {
	var xp XPerson
	GetPayor(pid, &xp.pay)
	GetXPerson(xp.pay.TCID, &xp)
	return xp
}

// GetRentableByID reads a Rentable structure based on the supplied rentable id
func GetRentableByID(rid int, r *Rentable) {
	rlib.Errcheck(App.prepstmt.getRentable.QueryRow(rid).Scan(&r.RID, &r.LID, &r.RTID, &r.BID, &r.UNITID, &r.Name, &r.Assignment, &r.Report, &r.DefaultOccType, &r.OccType, &r.LastModTime, &r.LastModBy))
}

// GetRentable reads and returns a Rentable structure based on the supplied rentable id
func GetRentable(rid int) Rentable {
	var r Rentable
	rlib.Errcheck(App.prepstmt.getRentable.QueryRow(rid).Scan(&r.RID, &r.LID, &r.RTID, &r.BID, &r.UNITID, &r.Name, &r.Assignment, &r.Report, &r.DefaultOccType, &r.OccType, &r.LastModTime, &r.LastModBy))
	return r
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
		GetRentableByID(rid, &x.R)
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
func GetRentableType(rtid int, rt *RentableType) {
	rlib.Errcheck(App.prepstmt.getRentableType.QueryRow(rtid).Scan(&rt.RTID, &rt.BID, &rt.Name, &rt.Frequency, &rt.Proration, &rt.LastModTime, &rt.LastModBy))
}

// GetUnitType returns characteristics of the unit
func GetUnitType(utid int, ut *UnitType) {
	rlib.Errcheck(App.prepstmt.getUnitType.QueryRow(utid).Scan(&ut.UTID, &ut.BID, &ut.Style, &ut.Name, &ut.SqFt, &ut.Frequency, &ut.Proration, &ut.LastModTime, &ut.LastModBy))
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
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.AcctRule))
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

// GetRentableMarketRates loads all the MarketRate rent information for this rentable into an array
func GetRentableMarketRates(rt *RentableType) {
	// now get all the MarketRate rent info...
	rows, err := App.prepstmt.getRentableMarketRates.Query(rt.RTID)
	rlib.Errcheck(err)
	defer rows.Close()
	LatestMRDTStart := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var a RentableMarketRate
		rlib.Errcheck(rows.Scan(&a.RTID, &a.MarketRate, &a.DtStart, &a.DtStop))
		if a.DtStart.After(LatestMRDTStart) {
			LatestMRDTStart = a.DtStart
			rt.MRCurrent = a.MarketRate
		}
		rt.MR = append(rt.MR, a)
	}
	rlib.Errcheck(rows.Err())
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
		rlib.Errcheck(rows.Scan(&a.RTID, &a.BID, &a.Name, &a.Frequency, &a.Proration, &a.LastModTime, &a.LastModBy))
		a.MR = make([]RentableMarketRate, 0)
		GetRentableMarketRates(&a)
		t[a.RTID] = a
	}
	rlib.Errcheck(rows.Err())

	return t
}

// GetUnitMarketRates loads all the MarketRate rent information for this unit into an array
func GetUnitMarketRates(rt *UnitType) {
	// now get all the MarketRate rent info...
	rows, err := App.prepstmt.getUnitMarketRates.Query(rt.UTID)
	rlib.Errcheck(err)
	defer rows.Close()
	LatestMRDTStart := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	for rows.Next() {
		var a UnitMarketRate
		rlib.Errcheck(rows.Scan(&a.UTID, &a.MarketRate, &a.DtStart, &a.DtStop))
		if a.DtStart.After(LatestMRDTStart) {
			LatestMRDTStart = a.DtStart
			rt.MRCurrent = a.MarketRate
		}
		rt.MR = append(rt.MR, a)
	}
	rlib.Errcheck(rows.Err())
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
		rlib.Errcheck(rows.Scan(&a.UTID, &a.BID, &a.Style, &a.Name, &a.SqFt, &a.Frequency, &a.Proration, &a.LastModTime, &a.LastModBy))
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
		rlib.Errcheck(rows.Scan(&a.ASMID, &a.BID, &a.RID, &a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency, &a.ProrationMethod, &a.AcctRule, &a.LastModTime, &a.LastModBy))
		t = append(t, a)
	}
	return t
}

// GetLedgerMarkerByGLNo returns the LedgerMarker struct for the GLNo with the supplied name
func GetLedgerMarkerByGLNo(s string) LedgerMarker {
	var r LedgerMarker
	// fmt.Printf("Ledger = %s\n", s)
	err := App.prepstmt.getLedgerMarkerByGLNo.QueryRow(s).Scan(&r.LMID, &r.BID, &r.GLNumber, &r.State, &r.Dt, &r.Balance, &r.DefaultAcct, &r.Name)
	if nil != err {
		fmt.Printf("GetLedgerMarkerByGLNo: Could not find ledgermarker for GLNumber \"%s\".\n", s)
		fmt.Printf("err = %v\n", err)
	}
	return r
}

// GetDefaultCashLedgerMarker returns the LedgerMarker the default cash account associated with Business bid
func GetDefaultCashLedgerMarker(bid int) LedgerMarker {
	var r LedgerMarker
	err := App.prepstmt.getDefaultCashLedgerMarker.QueryRow(bid).Scan(&r.LMID, &r.BID, &r.GLNumber, &r.State, &r.Dt, &r.Balance, &r.DefaultAcct, &r.Name)
	if nil != err {
		fmt.Printf("GetDefaultCashLedgerMarker: Could not find default cash acount for Business %d.\n", bid)
		fmt.Printf("err = %v\n", err)
	}
	return r
}

// GetRentalAgreement returns the Ledger struct for the account with the supplied name
func GetRentalAgreement(raid int) (RentalAgreement, error) {
	var r RentalAgreement
	// fmt.Printf("Ledger = %s\n", s)
	err := App.prepstmt.getRentalAgreement.QueryRow(raid).Scan(&r.RAID, &r.RATID,
		&r.BID, &r.RID, &r.UNITID, &r.PID, &r.PrimaryTenant, &r.RentalStart,
		&r.RentalStop, &r.Renewal, &r.SpecialProvisions, &r.LastModTime, &r.LastModBy)
	if nil != err {
		fmt.Printf("GetRentalAgreement: could not get rental agreement with raid = %d,  err = %v\n", raid, err)
	}
	return r, err
}

// GetReceiptAllocations loads all receipt allocations associated with the supplied receipt id into
// the RA array within a Receipt structure
func GetReceiptAllocations(rcptid int, r *Receipt) {
	rows, err := App.prepstmt.getReceiptAllocations.Query(rcptid)
	rlib.Errcheck(err)
	defer rows.Close()
	r.RA = make([]ReceiptAllocation, 0)
	for rows.Next() {
		var a ReceiptAllocation
		rlib.Errcheck(rows.Scan(&a.RCPTID, &a.Amount, &a.ASMID, &a.AcctRule))
		r.RA = append(r.RA, a)
	}
}

// GetReceipts for the supplied business (bid) in date range [d1 - d2)
func GetReceipts(bid int, d1, d2 *time.Time) []Receipt {
	rows, err := App.prepstmt.getReceiptsInDateRange.Query(bid, d1, d2)
	rlib.Errcheck(err)
	defer rows.Close()
	var t []Receipt
	t = make([]Receipt, 0)
	for rows.Next() {
		var r Receipt
		rlib.Errcheck(rows.Scan(&r.RCPTID, &r.BID, &r.RAID, &r.PMTID, &r.Dt, &r.Amount, &r.AcctRule))
		r.RA = make([]ReceiptAllocation, 0)
		GetReceiptAllocations(r.RCPTID, &r)
		t = append(t, r)
	}
	return t
}

// GetReceipt returns a receipt structure for the supplied RCPTID
func GetReceipt(rcptid int) Receipt {
	var r Receipt
	rlib.Errcheck(App.prepstmt.getReceipt.QueryRow(rcptid).Scan(&r.RCPTID, &r.BID, &r.RAID, &r.PMTID, &r.Dt, &r.Amount, &r.AcctRule))
	GetReceiptAllocations(rcptid, &r)
	return r
}

// GetAssessment returns the Assessment struct for the account with the supplied asmid
func GetAssessment(asmid int) (Assessment, error) {
	var a Assessment
	err := App.prepstmt.getAssessment.QueryRow(asmid).Scan(&a.ASMID, &a.BID, &a.RID,
		&a.UNITID, &a.ASMTID, &a.RAID, &a.Amount, &a.Start, &a.Stop, &a.Frequency,
		&a.ProrationMethod, &a.AcctRule, &a.LastModTime, &a.LastModBy)
	if nil != err {
		fmt.Printf("GetAssessment: could not get assessment with asmid = %d,  err = %v\n", asmid, err)
	}
	return a, err
}

// GetXType returns the RentalType structure and if it exists the UnitType structure is also returned
func GetXType(rtid, utid int) XType {
	var xt XType
	GetRentableType(rtid, &xt.RT)
	GetUnitType(utid, &xt.UT)
	return xt
}

// GetJournalMarkers loads the last n journal markers
func GetJournalMarkers(n int) []JournalMarker {
	rows, err := App.prepstmt.getJournalMarkers.Query(n)
	rlib.Errcheck(err)
	defer rows.Close()
	var t []JournalMarker
	t = make([]JournalMarker, 0)
	for rows.Next() {
		var r JournalMarker
		rlib.Errcheck(rows.Scan(&r.JMID, &r.BID, &r.State, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetLastJournalMarker returns the last journal marker or nil if no journal markers exist
func GetLastJournalMarker() JournalMarker {
	t := GetJournalMarkers(1)
	return t[0]
}

// GetJournalAllocations loads all Journal allocations associated with the supplied Journal id into
// the RA array within a Journal structure
func GetJournalAllocations(jid int, j *Journal) {
	rows, err := App.prepstmt.getJournalAllocations.Query(jid)
	rlib.Errcheck(err)
	defer rows.Close()
	j.JA = make([]JournalAllocation, 0)
	for rows.Next() {
		var a JournalAllocation
		rlib.Errcheck(rows.Scan(&a.JID, &a.Amount, &a.ASMID, &a.AcctRule))
		j.JA = append(j.JA, a)
	}
}
