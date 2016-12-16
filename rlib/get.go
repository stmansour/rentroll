package rlib

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

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

	err := RRdb.Prepstmt.FindAgreementByRentable.QueryRow(rid, d1, d2).Scan(&a.RAID, &a.RID, &a.CLID, &a.ContractRent, &a.DtStart, &a.DtStop)
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
		ReadAssessments(rows, &a)
		t = append(t, a)
	}
	return t
}

// GetAssessment returns the Assessment struct for the account with the supplied asmid
func GetAssessment(asmid int64) (Assessment, error) {
	var a Assessment
	row := RRdb.Prepstmt.GetAssessment.QueryRow(asmid)
	ReadAssessment(row, &a)
	return a, nil
}

// GetAssessmentInstance returns the Assessment struct for the account with the supplied asmid
func GetAssessmentInstance(start *time.Time, pasmid int64) (Assessment, error) {
	var a Assessment
	row := RRdb.Prepstmt.GetAssessmentInstance.QueryRow(start, pasmid)
	ReadAssessment(row, &a)
	return a, nil
}

// GetAssessmentDuplicate returns the Assessment struct for the account with the supplied asmid
func GetAssessmentDuplicate(start *time.Time, amt float64, pasmid, rid, raid, atypelid int64) Assessment {
	var a Assessment
	row := RRdb.Prepstmt.GetAssessmentDuplicate.QueryRow(start, amt, pasmid, rid, raid, atypelid)
	ReadAssessment(row, &a)
	return a
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

// GetAllBusinesses generates a report of all Businesses defined in the database.
func GetAllBusinesses() ([]Business, error) {
	var m []Business
	rows, err := RRdb.Prepstmt.GetAllBusinesses.Query()
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p Business
		ReadBusinesses(rows, &p)
		m = append(m, p)
	}
	Errcheck(rows.Err())
	return m, err
}

// GetBusiness loads the Business struct for the supplied Business id
func GetBusiness(bid int64, a *Business) {
	row := RRdb.Prepstmt.GetBusiness.QueryRow(bid)
	ReadBusiness(row, a)
}

// GetBusinessByDesignation loads the Business struct for the supplied designation
func GetBusinessByDesignation(des string) Business {
	var a Business
	row := RRdb.Prepstmt.GetBusinessByDesignation.QueryRow(des)
	ReadBusiness(row, &a)
	return a
}

// GetXBusiness loads the XBusiness struct for the supplied Business id.
func GetXBusiness(bid int64, xbiz *XBusiness) {
	if xbiz.P.BID == 0 && bid > 0 {
		GetBusiness(bid, &xbiz.P)
	}
	xbiz.RT = GetBusinessRentableTypes(bid)
	xbiz.US = make(map[int64]RentableSpecialty, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessSpecialtyTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableSpecialty
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
func GetCustomAttribute(id int64) CustomAttribute {
	var a CustomAttribute
	row := RRdb.Prepstmt.GetCustomAttribute.QueryRow(id)
	ReadCustomAttribute(row, &a)
	return a
}

// GetCustomAttributeByVals reads a CustomAttribute structure based on the supplied attributes
// t = data type (CUSTSTRING, CUSTINT, CUSTUINT, CUSTFLOAT, CUSTDATE
// n = name of this custom attribute
// v = the value of this attribute
// u = units (if not applicable then "")
func GetCustomAttributeByVals(t int64, n, v, u string) CustomAttribute {
	var a CustomAttribute
	row := RRdb.Prepstmt.GetCustomAttributeByVals.QueryRow(t, n, v, u)
	ReadCustomAttribute(row, &a)
	return a
}

// GetAllCustomAttributes returns a list of CustomAttributes for the supplied elementid and instanceid
func GetAllCustomAttributes(elemid, id int64) (map[string]CustomAttribute, error) {
	var t []int64
	var m map[string]CustomAttribute
	m = make(map[string]CustomAttribute, 0)
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
		c := GetCustomAttribute(t[i])
		m[c.Name] = c
	}

	return m, err
}

// GetCustomAttributeRef reads a CustomAttribute structure for the supplied ElementType, ID, and CID
func GetCustomAttributeRef(e, i, c int64) CustomAttributeRef {
	var a CustomAttributeRef
	row := RRdb.Prepstmt.GetCustomAttributeRef.QueryRow(e, i, c)
	ReadCustomAttributeRef(row, &a)
	return a
}

// LoadRentableTypeCustomaAttributes adds all the custom attributes to each RentableType
func LoadRentableTypeCustomaAttributes(xbiz *XBusiness) {
	var err error
	for k, v := range xbiz.RT {
		var tmp = xbiz.RT[k]
		tmp.CA, err = GetAllCustomAttributes(ELEMRENTABLETYPE, v.RTID)
		if err != nil {
			Ulog("LoadRentableTypeCustomaAttributes: error reading custom attributes elementtype=%d, id=%d, err = %s\n", ELEMRENTABLETYPE, v.RTID, err.Error())
		}
		xbiz.RT[k] = tmp // this workaround (assigning to tmp) instead of just directly assigning the .CA member is a known issue in go
	}
}

//=======================================================
//  DEMAND SOURCE
//=======================================================

// GetDemandSource reads a DemandSource structure based on the supplied DemandSource id
func GetDemandSource(id int64, t *DemandSource) {
	ReadDemandSource(RRdb.Prepstmt.GetDemandSource.QueryRow(id), t)
}

// GetDemandSourceByName reads a DemandSource structure based on the supplied DemandSource id
func GetDemandSourceByName(bid int64, name string, t *DemandSource) {
	ReadDemandSource(RRdb.Prepstmt.GetDemandSourceByName.QueryRow(bid, name), t)
}

// GetAllDemandSources returns an array of DemandSource structures containing all sources for the supplied BID
func GetAllDemandSources(id int64) ([]DemandSource, error) {
	var m []DemandSource
	rows, err := RRdb.Prepstmt.GetAllDemandSources.Query(id)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var s DemandSource
		ReadDemandSources(rows, &s)
		m = append(m, s)
	}
	Errcheck(rows.Err())
	return m, err
}

//=======================================================
//  DEPOSIT
//  Deposit, Depository, Deposit Method, DepositPart
//=======================================================

// GetDeposit reads a Deposit structure based on the supplied Deposit id
func GetDeposit(id int64) (Deposit, error) {
	var a Deposit
	err := RRdb.Prepstmt.GetDeposit.QueryRow(id).Scan(&a.DID, &a.BID, &a.DEPID, &a.DPMID, &a.Dt, &a.Amount, &a.LastModTime, &a.LastModBy)
	return a, err
}

// GetAllDepositsInRange returns an array of all Deposits for bid between the supplied dates
func GetAllDepositsInRange(bid int64, d1, d2 *time.Time) []Deposit {
	var t []Deposit
	rows, err := RRdb.Prepstmt.GetAllDepositsInRange.Query(bid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a Deposit
		Errcheck(rows.Scan(&a.DID, &a.BID, &a.DEPID, &a.DPMID, &a.Dt, &a.Amount, &a.LastModTime, &a.LastModBy))
		a.DP, err = GetDepositParts(a.DID)
		Errcheck(err)
		t = append(t, a)
	}
	Errcheck(rows.Err())
	return t
}

// GetDepository reads a Depository structure based on the supplied Depository id
func GetDepository(id int64) (Depository, error) {
	var a Depository
	err := RRdb.Prepstmt.GetDepository.QueryRow(id).Scan(&a.DEPID, &a.BID, &a.Name, &a.AccountNo, &a.LastModTime, &a.LastModBy)
	return a, err
}

// GetDepositoryByAccount reads a Depository structure based on the supplied Account id
func GetDepositoryByAccount(bid int64, acct string) Depository {
	var a Depository
	row := RRdb.Prepstmt.GetDepositoryByAccount.QueryRow(bid, acct)
	ReadDepository(row, &a)
	return a
}

// GetAllDepositories returns an array of all Depositories for the supplied business
func GetAllDepositories(bid int64) []Depository {
	var t []Depository
	rows, err := RRdb.Prepstmt.GetAllDepositories.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r Depository
		Errcheck(rows.Scan(&r.DEPID, &r.BID, &r.Name, &r.AccountNo, &r.LastModTime, &r.LastModBy))
		t = append(t, r)
	}
	Errcheck(rows.Err())
	return t
}

// GetDepositParts reads a DepositPart structure based on the supplied DepositPart DID
func GetDepositParts(id int64) ([]DepositPart, error) {
	var m []DepositPart
	rows, err := RRdb.Prepstmt.GetDepositParts.Query(id)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a DepositPart
		Errcheck(rows.Scan(&a.DID, &a.RCPTID))
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m, err
}

// GetDepositMethod reads a DepositMethod structure based on the supplied Deposit id
func GetDepositMethod(id int64) (DepositMethod, error) {
	var a DepositMethod
	err := RRdb.Prepstmt.GetDepositMethod.QueryRow(id).Scan(&a.DPMID, &a.BID, &a.Name)
	return a, err
}

// GetDepositMethodByName reads a DepositMethod structure based on the supplied BID and Name
func GetDepositMethodByName(bid int64, name string) (DepositMethod, error) {
	var a DepositMethod
	err := RRdb.Prepstmt.GetDepositMethodByName.QueryRow(bid, name).Scan(&a.DPMID, &a.BID, &a.Name)
	return a, err
}

// GetAllDepositMethods returns an array of all DepositMethods for the supplied business
func GetAllDepositMethods(bid int64) []DepositMethod {
	var t []DepositMethod
	rows, err := RRdb.Prepstmt.GetAllDepositMethods.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r DepositMethod
		Errcheck(rows.Scan(&r.DPMID, &r.BID, &r.Name))
		t = append(t, r)
	}
	Errcheck(rows.Err())
	return t
}

//=======================================================
//  I N V O I C E
//=======================================================

// GetInvoice reads a Invoice structure based on the supplied Invoice id
func GetInvoice(id int64) (Invoice, error) {
	var a Invoice
	var err error
	row := RRdb.Prepstmt.GetInvoice.QueryRow(id)
	ReadInvoice(row, &a)
	if err == nil {
		a.A, err = GetInvoiceAssessments(id)
		if err == nil {
			a.P, err = GetInvoicePayors(id)
		}
	}
	return a, err
}

// GetAllInvoicesInRange returns an array of all Invoices for bid between the supplied dates
func GetAllInvoicesInRange(bid int64, d1, d2 *time.Time) []Invoice {
	var t []Invoice
	rows, err := RRdb.Prepstmt.GetAllInvoicesInRange.Query(bid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a Invoice
		ReadInvoices(rows, &a)
		a.A, err = GetInvoiceAssessments(a.InvoiceNo)
		Errcheck(err)
		a.P, err = GetInvoicePayors(a.InvoiceNo)
		t = append(t, a)
		Errcheck(err)
	}
	Errcheck(rows.Err())
	return t
}

// GetInvoiceAssessments reads a InvoiceAssessment structure based on the supplied InvoiceAssessment DID
func GetInvoiceAssessments(id int64) ([]InvoiceAssessment, error) {
	var m []InvoiceAssessment
	rows, err := RRdb.Prepstmt.GetInvoiceAssessments.Query(id)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a InvoiceAssessment
		Errcheck(rows.Scan(&a.InvoiceNo, &a.ASMID))
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m, err
}

// GetInvoicePayors reads an InvoicePayor structure based on the supplied InvoiceNo (id)
func GetInvoicePayors(id int64) ([]InvoicePayor, error) {
	var m []InvoicePayor
	rows, err := RRdb.Prepstmt.GetInvoicePayors.Query(id)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a InvoicePayor
		Errcheck(rows.Scan(&a.InvoiceNo, &a.PID))
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m, err
}

//=======================================================
//  JOURNAL
//=======================================================

// GetJournal returns the Journal struct for the journal entry with the supplied id
func GetJournal(jid int64) Journal {
	var r Journal
	row := RRdb.Prepstmt.GetJournal.QueryRow(jid)
	ReadJournal(row, &r)
	return r
}

// // GetJournalInstance returns the Journal struct for entries that were created with the assumption that
// // they are idempotent -- essentially: instances of recurring assessments and vacancy instances.  This call
// // is made prior to generating new ones to ensure that we don't have double entries for the same thing.
// func GetJournalInstance(id int64, dt1, dt2 *time.Time) Journal {
// 	var r Journal
// 	row := RRdb.Prepstmt.GetJournalInstance.QueryRow(id, dt1, dt2)
// 	ReadJournal(row, &r)
// 	return r
// }

// GetJournalVacancy returns the Journal struct for entries that were created with the assumption that
// they are idempotent -- essentially: instances of recurring assessments and vacancy instances.  This call
// is made prior to generating new ones to ensure that we don't have double entries for the same thing.
func GetJournalVacancy(id int64, dt1, dt2 *time.Time) Journal {
	var r Journal
	row := RRdb.Prepstmt.GetJournalVacancy.QueryRow(id, dt1, dt2)
	ReadJournal(row, &r)
	return r
}

// GetJournalByReceiptID returns the Journal struct for a Journal Entry that references the supplied
// receiptID
func GetJournalByReceiptID(id int64) Journal {
	var r Journal
	row := RRdb.Prepstmt.GetJournalByReceiptID.QueryRow(id)
	ReadJournal(row, &r)
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
	// for i := 0; i < len(j.JA)-1; i++ {
	// 	for k := i + 1; k < len(j.JA); k++ {
	// 		if j.JA[i].Amount < j.JA[k].Amount {
	// 			j.JA[i].Amount, j.JA[k].Amount = j.JA[k].Amount, j.JA[i].Amount
	// 		}
	// 	}
	// }
}

//=======================================================
//  L E D G E R   M A R K E R
//=======================================================

// GetLatestLedgerMarkerByLID returns the LedgerMarker struct for the GLAccount with the supplied LID
func GetLatestLedgerMarkerByLID(bid, lid int64) LedgerMarker {
	var r LedgerMarker
	row := RRdb.Prepstmt.GetLatestLedgerMarkerByLID.QueryRow(bid, lid)
	ReadLedgerMarker(row, &r)
	return r
}

// GetLedgerMarkerOnOrBefore returns the LedgerMarker struct for the GLAccount with the supplied LID
func GetLedgerMarkerOnOrBefore(bid, lid int64, dt *time.Time) LedgerMarker {
	var r LedgerMarker
	row := RRdb.Prepstmt.GetLedgerMarkerOnOrBefore.QueryRow(bid, lid, dt)
	ReadLedgerMarker(row, &r)
	return r
}

// GetRALedgerMarkerOnOrBefore returns the LedgerMarker struct for the GLAccount with
// the supplied LID filtered for the supplied RentalAgreement raid
func GetRALedgerMarkerOnOrBefore(bid, lid, raid int64, dt *time.Time) LedgerMarker {
	var r LedgerMarker
	row := RRdb.Prepstmt.GetRALedgerMarkerOnOrBefore.QueryRow(bid, lid, raid, dt)
	ReadLedgerMarker(row, &r)
	return r
}

// GetRentableLedgerMarkerOnOrBefore returns the LedgerMarker struct for the GLAccount with
// the supplied LID filtered for the supplied Rentable rid
func GetRentableLedgerMarkerOnOrBefore(bid, lid, rid int64, dt *time.Time) LedgerMarker {
	var r LedgerMarker
	row := RRdb.Prepstmt.GetRentableLedgerMarkerOnOrBefore.QueryRow(bid, lid, rid, dt)
	ReadLedgerMarker(row, &r)
	return r
}

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
func LoadRALedgerMarker(bid, lid, raid int64, dt *time.Time) LedgerMarker {
	lm := GetRALedgerMarkerOnOrBefore(bid, lid, raid, dt)
	if lm.LMID == 0 {
		lm.LID = lid
		lm.BID = bid
		lm.RAID = raid
		lm.Dt = *dt
		lm.State = MARKERSTATEORIGIN
		err := InsertLedgerMarker(&lm)
		if nil != err {
			fmt.Printf("LoadRALedgerMarker: Error creating LedgerMarker: %s\n", err.Error())
		}
	}
	return lm
}

// GetLatestLedgerMarkerByGLNo returns the LedgerMarker struct for the GLNo with the supplied name
func GetLatestLedgerMarkerByGLNo(bid int64, s string) LedgerMarker {
	l := GetLedgerByGLNo(bid, s)
	if l.LID == 0 {
		var r LedgerMarker
		return r
	}
	return GetLatestLedgerMarkerByLID(bid, l.LID)
}

// GetLatestLedgerMarkerByType returns the LedgerMarker struct for the supplied type
func GetLatestLedgerMarkerByType(bid int64, t int64) LedgerMarker {
	var r LedgerMarker
	l := GetLedgerByType(bid, t)
	if 0 == l.LID {
		return r
	}
	return GetLatestLedgerMarkerByLID(bid, l.LID)
}

// GetAllLedgerMarkersOnOrBefore returns a map of all ledgermarkers for the supplied business and dat
func GetAllLedgerMarkersOnOrBefore(bid int64, dt *time.Time) map[int64]LedgerMarker {
	var t map[int64]LedgerMarker
	t = make(map[int64]LedgerMarker, 0) // this line is absolutely necessary
	rows, err := RRdb.Prepstmt.GetAllLedgerMarkersOnOrBefore.Query(bid, dt)
	Errcheck(err)
	defer rows.Close()
	// fmt.Printf("%4s  %4s  %4s  %5s  %10s  %8s\n", "LMID", "LID", "BID", "State", "Dt", "Balance")
	for rows.Next() {
		var r LedgerMarker
		ReadLedgerMarkers(rows, &r)
		t[r.LID] = r
		// fmt.Printf("%4d  %4d  %4d  %5d  %10s  %8.2f\n", r.LMID, r.LID, r.BID, r.State, r.Dt, r.Balance)
	}
	Errcheck(rows.Err())
	return t
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
		ReadGLAccounts(rows, &r)
		t = append(t, r)
	}
	return t
}

// GetGLAccountMap returns a map of all GLAccounts for the supplied business
func GetGLAccountMap(bid int64) map[int64]GLAccount {
	rows, err := RRdb.Prepstmt.GetLedgerList.Query(bid)
	Errcheck(err)
	defer rows.Close()
	var t map[int64]GLAccount
	t = make(map[int64]GLAccount)
	for rows.Next() {
		var r GLAccount
		ReadGLAccounts(rows, &r)
		t[r.LID] = r
	}
	return t
}

// GetLedger returns the GLAccount struct for the supplied LID
func GetLedger(lid int64) GLAccount {
	var a GLAccount
	row := RRdb.Prepstmt.GetLedger.QueryRow(lid)
	ReadGLAccount(row, &a)
	return a
}

// GetLedgerEntryByJAID returns the GLAccount struct for the supplied LID
func GetLedgerEntryByJAID(bid, lid, jaid int64) LedgerEntry {
	var a LedgerEntry
	row := RRdb.Prepstmt.GetLedgerEntryByJAID.QueryRow(bid, lid, jaid)
	ReadLedgerEntry(row, &a)
	return a
}

// GetLedgerByGLNo returns the GLAccount struct for the supplied GLNo
func GetLedgerByGLNo(bid int64, s string) GLAccount {
	var a GLAccount
	row := RRdb.Prepstmt.GetLedgerByGLNo.QueryRow(bid, s)
	ReadGLAccount(row, &a)
	return a
}

// GetLedgerByType returns the GLAccount struct for the supplied Type
func GetLedgerByType(bid, t int64) GLAccount {
	var a GLAccount
	row := RRdb.Prepstmt.GetLedgerByType.QueryRow(bid, t)
	ReadGLAccount(row, &a)
	return a
}

// // GetRABalanceLedger returns the GLAccount struct for the supplied Type
// func GetRABalanceLedger(bid, RAID int64) GLAccount {
// 	var a GLAccount
// 	var err error
// 	row := RRdb.Prepstmt.GetRABalanceLedger.QueryRow(bid)
// 	ReadGLAccount(row, &a)
// 	return a
// }

// // GetSecDepBalanceLedger returns the GLAccount struct for the supplied Type
// func GetSecDepBalanceLedger(bid, RAID int64) GLAccount {
// 	var a GLAccount
// 	var err error
// 	row := RRdb.Prepstmt.GetSecDepBalanceLedger.QueryRow(bid, RAID)
// 	ReadGLAccount(row, &a)
// 	return a
// }

// GetDefaultLedgers loads the default GLAccount for the supplied Business bid
func GetDefaultLedgers(bid int64) {
	rows, err := RRdb.Prepstmt.GetDefaultLedgers.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var r GLAccount
		ReadGLAccounts(rows, &r)
		RRdb.BizTypes[bid].DefaultAccts[r.Type] = &r
	}
}

//=======================================================
//  LEDGER ENTRY
//=======================================================

// GetLedgerEntryArray returns a list of Ledger Entries for the supplied rows value
func GetLedgerEntryArray(rows *sql.Rows) ([]LedgerEntry, error) {
	var m []LedgerEntry
	for rows.Next() {
		var le LedgerEntry
		ReadLedgerEntries(rows, &le)
		m = append(m, le)
	}
	return m, rows.Err()
}

// GetLedgerEntriesInRange returns a list of Ledger Entries for the supplied Ledger during the supplied range
func GetLedgerEntriesInRange(d1, d2 *time.Time, bid, lid int64) ([]LedgerEntry, error) {
	rows, err := RRdb.Prepstmt.GetLedgerEntriesInRangeByLID.Query(bid, lid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	return GetLedgerEntryArray(rows)
}

// GetLedgerEntriesForRAID returns a list of Ledger Entries for the supplied RentalAgreement and Ledger
func GetLedgerEntriesForRAID(d1, d2 *time.Time, raid, lid int64) ([]LedgerEntry, error) {
	rows, err := RRdb.Prepstmt.GetLedgerEntriesForRAID.Query(d1, d2, raid, lid)
	Errcheck(err)
	defer rows.Close()
	return GetLedgerEntryArray(rows)
}

// GetLedgerEntriesForRentable returns a list of Ledger Entries for the supplied Rentable (rid) and Ledger (lid)
func GetLedgerEntriesForRentable(d1, d2 *time.Time, rid, lid int64) ([]LedgerEntry, error) {
	rows, err := RRdb.Prepstmt.GetLedgerEntriesForRentable.Query(d1, d2, rid, lid)
	Errcheck(err)
	defer rows.Close()
	return GetLedgerEntryArray(rows)
}

// GetAllLedgerEntriesForRAID returns a list of Ledger Entries for the supplied RentalAgreement
func GetAllLedgerEntriesForRAID(d1, d2 *time.Time, raid int64) ([]LedgerEntry, error) {
	rows, err := RRdb.Prepstmt.GetAllLedgerEntriesForRAID.Query(d1, d2, raid)
	Errcheck(err)
	defer rows.Close()
	return GetLedgerEntryArray(rows)
}

// GetAllLedgerEntriesForRID returns a list of Ledger Entries for the supplied Rentable ID
func GetAllLedgerEntriesForRID(d1, d2 *time.Time, rid int64) ([]LedgerEntry, error) {
	rows, err := RRdb.Prepstmt.GetAllLedgerEntriesForRID.Query(d1, d2, rid)
	Errcheck(err)
	defer rows.Close()
	return GetLedgerEntryArray(rows)
}

// GetAllLedgerEntriesInRange returns a list of Ledger Entries for the supplied business and time period
func GetAllLedgerEntriesInRange(bid int64, d1, d2 *time.Time) ([]LedgerEntry, error) {
	rows, err := RRdb.Prepstmt.GetAllLedgerEntriesInRange.Query(bid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	return GetLedgerEntryArray(rows)
}

// // GetLedgerEntriesInRange returns a list of Ledger Entries for the supplied business and time period
// func GetLedgerEntriesInRange(bid, lid, raid int64, d1, d2 *time.Time) ([]LedgerEntry, error) {
// 	rows, err := RRdb.Prepstmt.GetLedgerEntriesInRange.Query(bid, lid, raid, d1, d2)
// 	Errcheck(err)
// 	defer rows.Close()
// 	return GetLedgerEntryArray(rows)
// }

//=======================================================
//  NOTES
//=======================================================

// GetNote reads a Note structure based on the supplied Note id
func GetNote(tid int64, t *Note) {
	ReadNote(RRdb.Prepstmt.GetNote.QueryRow(tid), t)
}

// GetNoteAndChildNotes reads a Note structure based on the supplied Note id, then it reads all its child notes, organizes them by Date
// and returns them in an array
func GetNoteAndChildNotes(nid int64) Note {
	var n Note
	GetNote(nid, &n)
	rows, err := RRdb.Prepstmt.GetNoteAndChildNotes.Query(nid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p Note
		ReadNotes(rows, &p)
		n.CN = append(n.CN, p)
	}
	Errcheck(rows.Err())
	return n
}

//=======================================================
//  NOTELIST
//=======================================================

// GetNoteList reads a NoteList structure based on the supplied NoteList id
func GetNoteList(nlid int64) NoteList {
	var m NoteList
	Errcheck(RRdb.Prepstmt.GetNoteList.QueryRow(nlid).Scan(&m.NLID, &m.LastModTime, &m.LastModBy))
	rows, err := RRdb.Prepstmt.GetNoteListMembers.Query(nlid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var nid int64
		Errcheck(rows.Scan(&nid))
		p := GetNoteAndChildNotes(nid)
		m.N = append(m.N, p)
	}
	Errcheck(rows.Err())
	return m
}

//=======================================================
//  NOTE TYPE
//=======================================================

// GetNoteType reads a NoteType structure based on the supplied NoteType id
func GetNoteType(ntid int64, t *NoteType) {
	Errcheck(RRdb.Prepstmt.GetNoteType.QueryRow(ntid).Scan(&t.NTID, &t.BID, &t.Name, &t.LastModTime, &t.LastModBy))
}

// GetAllNoteTypes reads a NoteType structure based for all NoteTypes in the supplied bid
func GetAllNoteTypes(bid int64) []NoteType {
	var m []NoteType
	rows, err := RRdb.Prepstmt.GetAllNoteTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p NoteType
		Errcheck(rows.Scan(&p.NTID, &p.BID, &p.Name, &p.LastModTime, &p.LastModBy))
		m = append(m, p)
	}
	Errcheck(rows.Err())
	return m
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
		ReadPaymentTypes(rows, &a)
		// Errcheck(rows.Scan(&a.PMTID, &a.BID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.PMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

// GetPaymentTypeByName reads a PaymentType structure based on the supplied bid and na
func GetPaymentTypeByName(bid int64, name string, a *PaymentType) {
	ReadPaymentType(RRdb.Prepstmt.GetPaymentTypeByName.QueryRow(bid, name), a)
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
		ReadPaymentTypes(rows, &a)
		// Errcheck(rows.Scan(&a.PMTID, &a.BID, &a.Name, &a.Description, &a.LastModTime, &a.LastModBy))
		t[a.PMTID] = a
	}
	Errcheck(rows.Err())
	return t
}

//=======================================================
//  RATE PLAN
//=======================================================

// GetRatePlan reads a RatePlan structure based on the supplied RatePlan id
func GetRatePlan(id int64, a *RatePlan) {
	ReadRatePlan(RRdb.Prepstmt.GetRatePlan.QueryRow(id), a)
}

// GetRatePlanByName reads a RatePlan structure based on the supplied RatePlan id
func GetRatePlanByName(id int64, s string, a *RatePlan) {
	ReadRatePlan(RRdb.Prepstmt.GetRatePlanByName.QueryRow(id, s), a)
}

// GetAllRatePlans reads all RatePlan structures based on the supplied bid
func GetAllRatePlans(id int64) []RatePlan {
	var m []RatePlan
	rows, err := RRdb.Prepstmt.GetAllRatePlans.Query(id)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p RatePlan
		ReadRatePlans(rows, &p)
		m = append(m, p)
	}
	Errcheck(rows.Err())
	return m
}

// GetRatePlanRef reads a RatePlanRef structure based on the supplied RatePlanRef id
func GetRatePlanRef(id int64, a *RatePlanRef) {
	ReadRatePlanRef(RRdb.Prepstmt.GetRatePlanRef.QueryRow(id), a)
}

// GetRatePlanRefFull reads a RatePlanRef structure based on the supplied RatePlanRef id and
// pulls in all the RTRate and SPRate structure arrays
func GetRatePlanRefFull(id int64, a *RatePlanRef) {
	if a.RPRID == 0 {
		ReadRatePlanRef(RRdb.Prepstmt.GetRatePlanRef.QueryRow(id), a)
	}
	// now load all its rates
	rows, err := RRdb.Prepstmt.GetAllRatePlanRefRTRates.Query(id)
	if err != nil {
		Ulog("GetRatePlanRefFull:   GetAllRatePlanRefRTRates - error = %s\n", err.Error())
		return
	}
	for rows.Next() {
		var p RatePlanRefRTRate
		ReadRatePlanRefRTRates(rows, &p)
		a.RT = append(a.RT, p)
	}
	// now load all Specialty rates
	rows, err = RRdb.Prepstmt.GetAllRatePlanRefSPRates.Query(a.RPRID, a.RPID)
	if err != nil {
		Ulog("GetRatePlanRefFull: GetAllRatePlanRefSPRates - error = %s\n", err.Error())
		return
	}
	for rows.Next() {
		var p RatePlanRefSPRate
		ReadRatePlanRefSPRates(rows, &p)
		a.SP = append(a.SP, p)
	}
}

// GetRatePlanRefsInRange reads a RatePlanRef structure based on the supplied RatePlan id and the date.
func GetRatePlanRefsInRange(id int64, d1, d2 *time.Time) []RatePlanRef {
	var m []RatePlanRef
	rows, err := RRdb.Prepstmt.GetRatePlanRefsInRange.Query(id, d1, d2)
	if err != nil {
		Ulog("GetRatePlanRefsInRange: error = %s\n", err.Error())
		return m
	}
	for rows.Next() {
		var a RatePlanRef
		ReadRatePlanRefs(rows, &a)
		m = append(m, a)
	}
	return m
}

// GetAllRatePlanRefsInRange reads all RatePlanRef structure based on the supplied date range
func GetAllRatePlanRefsInRange(d1, d2 *time.Time) []RatePlanRef {
	var m []RatePlanRef
	rows, err := RRdb.Prepstmt.GetAllRatePlanRefsInRange.Query(d1, d2)
	if err != nil {
		Ulog("GetAllRatePlanRefsInRange: error = %s\n", err.Error())
		return m
	}
	for rows.Next() {
		var a RatePlanRef
		ReadRatePlanRefs(rows, &a)
		m = append(m, a)
	}
	return m
}

// GetRatePlanRefRTRate reads the RatePlanRefRTRate struct for the supplied rprid and rtid
func GetRatePlanRefRTRate(rprid, rtid int64, a *RatePlanRefRTRate) {
	row := RRdb.Prepstmt.GetRatePlanRefRTRate.QueryRow(rprid, rtid)
	ReadRatePlanRefRTRate(row, a)
}

// GetRatePlanRefSPRate reads the RatePlanRefSPRate struct for the supplied rprid and rtid
func GetRatePlanRefSPRate(rprid, rtid int64, a *RatePlanRefSPRate) {
	row := RRdb.Prepstmt.GetRatePlanRefSPRate.QueryRow(rprid, rtid)
	ReadRatePlanRefSPRate(row, a)
}

// GetAllRatePlanRefSPRates reads all RatePlanRefSPRate structures based on the supplied RatePlan id and the date.
func GetAllRatePlanRefSPRates(rprid, rtid int64) []RatePlanRefSPRate {
	var m []RatePlanRefSPRate
	rows, err := RRdb.Prepstmt.GetAllRatePlanRefSPRates.Query(rprid, rtid)
	if err != nil {
		Ulog("GetAllRatePlanRefSPRates: error = %s\n", err.Error())
		return m
	}
	for rows.Next() {
		var a RatePlanRefSPRate
		ReadRatePlanRefSPRates(rows, &a)
		m = append(m, a)
	}
	return m
}

//=======================================================
//  RECEIPT ALLOCATION
//=======================================================

// GetReceipt returns a Receipt structure for the supplied RCPTID
func GetReceipt(rcptid int64) Receipt {
	var r Receipt
	row := RRdb.Prepstmt.GetReceipt.QueryRow(rcptid)
	ReadReceipt(row, &r)
	GetReceiptAllocations(rcptid, &r)
	return r
}

// GetReceiptDuplicate returns a Receipt structure for the supplied RCPTID
func GetReceiptDuplicate(dt *time.Time, amt float64, docno string) Receipt {
	var r Receipt
	row := RRdb.Prepstmt.GetReceiptDuplicate.QueryRow(dt, amt, docno)
	ReadReceipt(row, &r)
	// GetReceiptAllocations(rcptid, &r)  // I don't think we need to spend time on this operation for duplicate checking, but uncomment if needed
	return r
}

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
		ReadReceipts(rows, &r)
		r.RA = make([]ReceiptAllocation, 0)
		GetReceiptAllocations(r.RCPTID, &r)
		t = append(t, r)
	}
	return t
}

// GetReceiptsInRAIDDateRange for the supplied RentalAgreement in date range [d1 - d2)
func GetReceiptsInRAIDDateRange(bid, raid int64, d1, d2 *time.Time) []Receipt {
	rows, err := RRdb.Prepstmt.GetReceiptsInRAIDDateRange.Query(bid, raid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []Receipt
	t = make([]Receipt, 0)
	for rows.Next() {
		var r Receipt
		ReadReceipts(rows, &r)
		r.RA = make([]ReceiptAllocation, 0)
		GetReceiptAllocations(r.RCPTID, &r)
		t = append(t, r)
	}
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
func GetRentableSpecialtyTypeByName(bid int64, name string) RentableSpecialty {
	var rsp RentableSpecialty
	err := RRdb.Prepstmt.GetRentableSpecialtyTypeByName.QueryRow(bid, name).Scan(&rsp.RSPID, &rsp.BID, &rsp.Name, &rsp.Fee, &rsp.Description)
	if err != nil {
		s := err.Error()
		if !strings.Contains(s, "no rows") {
			fmt.Printf("GetRentableSpecialtyTypeByName: err = %v\n", err)
		}
	}
	return rsp
}

// GetRentableSpecialtyType returns the RentableSpecialty record for the supplied RSPID
func GetRentableSpecialtyType(rspid int64) (RentableSpecialty, error) {
	var rs RentableSpecialty
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
// func SelectRentableTypeRefForDate(rsa *[]RentableSpecialty, dt *time.Time) RentableSpecialty {
// 	for i := 0; i < len(*rsa); i++ {
// 		if DateInRange(dt, &(*rsa)[i]. , &(*rsa)[i].DtStop) {
// 			return (*rsa)[i]
// 		}
// 	}
// 	var r RentableSpecialty
// 	return r // nothing matched
// }

// GetRentableSpecialtyTypesForRentableByRange returns an array of RentableSpecialty structures that
// apply for the supplied time range d1,d2
func GetRentableSpecialtyTypesForRentableByRange(r *Rentable, d1, d2 *time.Time) ([]RentableSpecialty, error) {
	var err error
	var rsta []RentableSpecialty
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
		Errcheck(rows.Scan(&a.RID, &a.RTID, &a.OverrideRentCycle, &a.OverrideProrationCycle, &a.DtStart, &a.DtStop, &a.LastModTime, &a.LastModBy))
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
		Errcheck(rows.Scan(&a.RID, &a.DtStart, &a.DtStop, &a.DtNoticeToVacate, &a.Status, &a.LastModTime, &a.LastModBy))
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
		&rt.Proration, &rt.GSRPC, &rt.ManageToBudget, &rt.LastModTime, &rt.LastModBy)
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
	err := RRdb.Prepstmt.GetRentableTypeByStyle.QueryRow(name, bid).Scan(&rt.RTID, &rt.BID, &rt.Style, &rt.Name, &rt.RentCycle, &rt.Proration, &rt.GSRPC, &rt.ManageToBudget, &rt.LastModTime, &rt.LastModBy)
	return rt, err
}

// GetBusinessRentableTypes returns a slice of RentableType indexed by the RTID
func GetBusinessRentableTypes(bid int64) map[int64]RentableType {
	var t map[int64]RentableType
	t = make(map[int64]RentableType, 0)
	rows, err := RRdb.Prepstmt.GetAllBusinessRentableTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableType
		Errcheck(rows.Scan(&a.RTID, &a.BID, &a.Style, &a.Name, &a.RentCycle, &a.Proration, &a.GSRPC, &a.ManageToBudget, &a.LastModTime, &a.LastModBy))
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
	for i := 0; i < len(mr); i++ {
		if DateRangeOverlap(d1, d2, &mr[i].DtStart, &mr[i].DtStop) {
			return mr[i].MarketRate
		}
	}
	return float64(0)
}

// GetRentableUsers returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentableUsers(rid int64, d1, d2 *time.Time) []RentableUser {
	rows, err := RRdb.Prepstmt.GetRentableUsers.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentableUser
	// t = make([]RentableUser, 0)
	for rows.Next() {
		var r RentableUser
		Errcheck(rows.Scan(&r.RID, &r.TCID, &r.DtStart, &r.DtStop))
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
	row := RRdb.Prepstmt.GetRentalAgreement.QueryRow(raid)
	err := ReadRentalAgreement(row, &r)
	return r, err
}

// LoadXRentalAgreement is like GetXRentalAgreement except that it assumes that some of the structure may
// already be loaded. It only loads those portions that appear not to already be loaded.
func LoadXRentalAgreement(raid int64, r *RentalAgreement, d1, d2 *time.Time) error {
	var err error
	if r.RAID != raid {
		(*r), err = GetRentalAgreement(raid)
	}

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
		var xp XPerson
		GetXPerson(m[i].TCID, &xp)
		r.P = append(r.P, xp)
	}

	n := GetRentableUsers(raid, d1, d2)
	r.T = make([]XPerson, 0)
	for i := 0; i < len(n); i++ {
		var xp XPerson
		GetXPerson(n[i].TCID, &xp)
		r.T = append(r.T, xp)
	}
	return err
}

// GetXRentalAgreement gets the RentalAgreement plus the associated rentables and payors for the
// time period specified
func GetXRentalAgreement(raid int64, d1, d2 *time.Time) (RentalAgreement, error) {
	var ra RentalAgreement
	err := LoadXRentalAgreement(raid, &ra, d1, d2)
	return ra, err
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
	rows, err := RRdb.Prepstmt.GetRentalAgreementsForRentable.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentalAgreementRentable
	for rows.Next() {
		var r RentalAgreementRentable
		Errcheck(rows.Scan(&r.RAID, &r.RID, &r.CLID, &r.ContractRent, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRentalAgreementRentables returns an array of RentalAgreementRentables associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentalAgreementRentables(raid int64, d1, d2 *time.Time) []RentalAgreementRentable {
	rows, err := RRdb.Prepstmt.GetRentalAgreementRentables.Query(raid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentalAgreementRentable
	for rows.Next() {
		var r RentalAgreementRentable
		Errcheck(rows.Scan(&r.RAID, &r.RID, &r.CLID, &r.ContractRent, &r.DtStart, &r.DtStop))
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
		ReadRentalAgreementPayors(rows, &r)
		t = append(t, r)
	}
	return t
}

//=======================================================
//  RENTAL AGREEMENT TEMPLATE
//=======================================================

// GetRentalAgreementTemplate returns the RentalAgreementTemplate struct for the supplied rental agreement id
func GetRentalAgreementTemplate(ratid int64) RentalAgreementTemplate {
	var r RentalAgreementTemplate
	row := RRdb.Prepstmt.GetRentalAgreementTemplate.QueryRow(ratid)
	ReadRentalAgreementTemplate(row, &r)
	return r
}

// GetRentalAgreementByRATemplateName returns the RentalAgreementTemplate struct for the supplied rental agreement id
func GetRentalAgreementByRATemplateName(ref string) RentalAgreementTemplate {
	var r RentalAgreementTemplate
	row := RRdb.Prepstmt.GetRentalAgreementByRATemplateName.QueryRow(ref)
	ReadRentalAgreementTemplate(row, &r)
	return r
}

//=======================================================
//  STRING LIST
//=======================================================

// GetStringList reads a StringList structure based on the supplied StringList id
func GetStringList(id int64, a *StringList) {
	ReadStringList(RRdb.Prepstmt.GetStringList.QueryRow(id), a)
	GetSLStrings(a.SLID, a)
}

// GetAllStringLists reads all StringList structures belonging to the business with the the supplied id
func GetAllStringLists(id int64) []StringList {
	var m []StringList
	rows, err := RRdb.Prepstmt.GetAllStringLists.Query(id)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a StringList
		ReadStringLists(rows, &a)
		GetSLStrings(a.SLID, &a)
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m
}

// GetStringListByName reads a StringList structure based on the supplied StringList id
func GetStringListByName(bid int64, s string, a *StringList) {
	ReadStringList(RRdb.Prepstmt.GetStringListByName.QueryRow(bid, s), a)
	if a.SLID != 0 {
		GetSLStrings(a.SLID, a)
	}
}

// GetSLStrings reads all strings with the supplid SLID into a
func GetSLStrings(id int64, a *StringList) {
	if id == 0 {
		return
	}
	rows, err := RRdb.Prepstmt.GetSLStrings.Query(id)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p SLString
		ReadSLStrings(rows, &p)
		a.S = append(a.S, p)
	}
	Errcheck(rows.Err())
}

//=======================================================
//  TRANSACTANT
//  Transactant, Prospect, User, Payor, XPerson
//=======================================================

// GetTransactantByPhoneOrEmail searches for a transactoant match on the phone number or email
func GetTransactantByPhoneOrEmail(BID int64, s string) Transactant {
	var t Transactant
	p := fmt.Sprintf("SELECT "+TRNSfields+" FROM Transactant where BID=%d AND (WorkPhone=\"%s\" or CellPhone=\"%s\" or PrimaryEmail=\"%s\" or SecondaryEmail=\"%s\")", BID, s, s, s, s)

	// there could be multiple people with the same identifying number...
	// to safeguard, we'll read as a query and accept first match
	rows, err := RRdb.Dbrr.Query(p)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		ReadTransactants(rows, &t)
		return t
	}
	//ReadTransactant(RRdb.Dbrr.QueryRow(p), &t)
	return t
}

// GetTransactant reads a Transactant structure based on the supplied Transactant id
func GetTransactant(tid int64, t *Transactant) {
	ReadTransactant(RRdb.Prepstmt.GetTransactant.QueryRow(tid), t)
}

// GetProspect reads a Prospect structure based on the supplied Transactant id
func GetProspect(id int64, p *Prospect) {
	ReadProspect(RRdb.Prepstmt.GetProspect.QueryRow(id), p)
}

// GetUser reads a User structure based on the supplied User id.
// This call does not load the vehicle list.  You can use GetVehiclesByTransactant()
// if you need them.  Or you can call GetXPerson, which loads all details about a Transactant.
func GetUser(tcid int64, t *User) {
	ReadUser(RRdb.Prepstmt.GetUser.QueryRow(tcid), t)
}

// GetPayor reads a Payor structure based on the supplied Transactant id
func GetPayor(pid int64, p *Payor) {
	ReadPayor(RRdb.Prepstmt.GetPayor.QueryRow(pid), p)
}

// GetVehicle reads a Vehicle structure based on the supplied Vehicle id
func GetVehicle(vid int64, t *Vehicle) {
	ReadVehicle(RRdb.Prepstmt.GetVehicle.QueryRow(vid), t)
}

func getVehicleList(rows *sql.Rows) []Vehicle {
	var m []Vehicle
	for rows.Next() {
		var a Vehicle
		ReadVehicles(rows, &a)
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m
}

// GetVehiclesByLicensePlate reads a Vehicle structure based on the supplied Vehicle id
func GetVehiclesByLicensePlate(s string) []Vehicle {
	rows, err := RRdb.Prepstmt.GetVehiclesByLicensePlate.Query(s)
	Errcheck(err)
	defer rows.Close()
	return getVehicleList(rows)
}

// GetVehiclesByTransactant reads a Vehicle structure based on the supplied Vehicle id
func GetVehiclesByTransactant(tcid int64) []Vehicle {
	rows, err := RRdb.Prepstmt.GetVehiclesByTransactant.Query(tcid)
	Errcheck(err)
	defer rows.Close()
	return getVehicleList(rows)
}

// GetVehiclesByBID reads a Vehicle structure based on the supplied Vehicle id
func GetVehiclesByBID(bid int64) []Vehicle {
	rows, err := RRdb.Prepstmt.GetVehiclesByBID.Query(bid)
	Errcheck(err)
	defer rows.Close()
	return getVehicleList(rows)
}

// GetXPerson will load a full XPerson given the trid
func GetXPerson(tcid int64, x *XPerson) {
	if 0 == x.Trn.TCID {
		GetTransactant(tcid, &x.Trn)
	}
	if 0 == x.Psp.TCID {
		GetProspect(tcid, &x.Psp)
	}
	if 0 == x.Usr.TCID {
		GetUser(tcid, &x.Usr)
		x.Usr.Vehicles = GetVehiclesByTransactant(tcid)
	}
	if 0 == x.Pay.TCID {
		GetPayor(tcid, &x.Pay)
	}
}

// GetDateOfLedgerMarkerOnOrBefore returns the Dt value of the last LedgerMarker set generated on or before d1
func GetDateOfLedgerMarkerOnOrBefore(bid int64, d1 *time.Time) time.Time {
	GenRcvLID := RRdb.BizTypes[bid].DefaultAccts[GLGENRCV].LID
	lm := GetLedgerMarkerOnOrBefore(bid, GenRcvLID, d1)
	if lm.LMID == 0 {
		Ulog("%s - SEVERE ERROR - unable to locate a LedgerMarker on or before %s\n", d1.Format(RRDATEFMT4))
	}
	return lm.Dt
}
