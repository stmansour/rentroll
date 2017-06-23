package rlib

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

//=======================================================
//  AR
//=======================================================

// GetAR reads a AR the structure for the supplied id
func GetAR(id int64) (AR, error) {
	var a AR
	row := RRdb.Prepstmt.GetAR.QueryRow(id)
	err := ReadAR(row, &a)
	return a, err
}

// GetARByName reads a AR the structure for the supplied bid and name
func GetARByName(id int64, name string) (AR, error) {
	var a AR
	row := RRdb.Prepstmt.GetARByName.QueryRow(id, name)
	err := ReadAR(row, &a)
	return a, err
}

// GetARsForRows uses the supplied rows param, gets all the AR records
// and returns them in a slice of AR structs
func GetARsForRows(rows *sql.Rows) []AR {
	defer rows.Close()
	var t []AR
	for i := 0; rows.Next(); i++ {
		var a AR
		ReadARs(rows, &a)
		t = append(t, a)
	}
	return t
}

// GetARMap returns a map of all account rules for the supplied bid
func GetARMap(bid int64) map[int64]AR {
	rows, err := RRdb.Prepstmt.GetAllARs.Query(bid)
	Errcheck(err)
	defer rows.Close()
	var t map[int64]AR
	t = make(map[int64]AR)
	for rows.Next() {
		var a AR
		ReadARs(rows, &a)
		t[a.ARID] = a
	}
	return t
}

// GetAllARs reads all AccountRules for the supplied BID
func GetAllARs(id int64) []AR {
	rows, err := RRdb.Prepstmt.GetAllARs.Query(id)
	Errcheck(err)
	return GetARsForRows(rows)
}

// GetARsByType reads all AccountRules for the supplied BID of type artype
func GetARsByType(id int64, artype int) []AR {
	rows, err := RRdb.Prepstmt.GetARsByType.Query(id, artype)
	Errcheck(err)
	return GetARsForRows(rows)
}

//=======================================================
//  A G R E E M E N T   P E T S
//=======================================================

// GetRentalAgreementPet reads a Pet the structure for the supplied PETID
func GetRentalAgreementPet(petid int64) (RentalAgreementPet, error) {
	var a RentalAgreementPet
	row := RRdb.Prepstmt.GetRentalAgreementPet.QueryRow(petid)
	err := ReadRentalAgreementPet(row, &a)
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
		ReadRentalAgreementPets(rows, &a)
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

	// SELECT RAID,BID,RID,DtStart,DtStop from RentalAgreementRentables where RID=? and DtStop>=? and DtStart<=?

	row := RRdb.Prepstmt.FindAgreementByRentable.QueryRow(rid, d1, d2)
	err := ReadRentalAgreementRentable(row, &a)
	return a, err
}

//=======================================================
//  A S S E S S M E N T S
//=======================================================

// GetAllRentableAssessments for the supplied RID and date range
func GetAllRentableAssessments(RID int64, d1, d2 *time.Time) []Assessment {
	rows, err := RRdb.Prepstmt.GetAllRentableAssessments.Query(RID, d1, d2)
	Errcheck(err)
	return GetAssessmentsByRows(rows)
}

// GetUnpaidAssessmentsByRAID for the supplied RAID
func GetUnpaidAssessmentsByRAID(RAID int64) []Assessment {
	rows, err := RRdb.Prepstmt.GetUnpaidAssessmentsByRAID.Query(RAID)
	Errcheck(err)
	return GetAssessmentsByRows(rows)
}

// GetAssessmentInstancesByParent for the supplied RAID
func GetAssessmentInstancesByParent(id int64, d1, d2 *time.Time) []Assessment {
	rows, err := RRdb.Prepstmt.GetAssessmentInstancesByParent.Query(id, d1, d2)
	Errcheck(err)
	return GetAssessmentsByRows(rows)
}

// GetAssessmentsByRows for the supplied sql.Rows
func GetAssessmentsByRows(rows *sql.Rows) []Assessment {
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
	err := RRdb.Prepstmt.GetBuilding.QueryRow(id).Scan(&t.BLDGID, &t.BID, &t.Address, &t.Address2, &t.City, &t.State, &t.PostalCode, &t.Country, &t.CreateTS, &t.CreateBy, &t.LastModTime, &t.LastModBy)
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
	xbiz.US = make(map[int64]RentableSpecialty)
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
	m = make(map[string]CustomAttribute)
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
	row := RRdb.Prepstmt.GetDeposit.QueryRow(id)
	err := ReadDeposit(row, &a)
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
		ReadDeposits(rows, &a)
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
	row := RRdb.Prepstmt.GetDepository.QueryRow(id)
	err := ReadDepository(row, &a)
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
		Errcheck(ReadDepositories(rows, &r))
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
		ReadDepositParts(rows, &a)
		m = append(m, a)
	}
	Errcheck(rows.Err())
	return m, err
}

// GetDepositMethod reads a DepositMethod structure based on the supplied Deposit id
func GetDepositMethod(id int64) (DepositMethod, error) {
	var a DepositMethod
	err := RRdb.Prepstmt.GetDepositMethod.QueryRow(id).Scan(&a.DPMID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy)
	return a, err
}

// GetDepositMethodByName reads a DepositMethod structure based on the supplied BID and Name
func GetDepositMethodByName(bid int64, name string) (DepositMethod, error) {
	var a DepositMethod
	err := RRdb.Prepstmt.GetDepositMethodByName.QueryRow(bid, name).Scan(&a.DPMID, &a.BID, &a.Name, &a.CreateTS, &a.CreateBy)
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
		Errcheck(rows.Scan(&r.DPMID, &r.BID, &r.Name, &r.CreateTS, &r.CreateBy))
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
		ReadInvoiceAssessments(rows, &a)
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
		ReadInvoicePayors(rows, &a)
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
		Errcheck(rows.Scan(&r.JMID, &r.BID, &r.State, &r.DtStart, &r.DtStop, &r.CreateTS, &r.CreateBy))
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
func GetJournalAllocation(jaid int64) JournalAllocation {
	var a JournalAllocation
	row := RRdb.Prepstmt.GetJournalAllocation.QueryRow(jaid)
	ReadJournalAllocation(row, &a)
	return a
}

// GetJournalAllocations loads all Journal allocations associated with the supplied Journal id into
// the RA array within a Journal structure
func GetJournalAllocations(j *Journal) {
	rows, err := RRdb.Prepstmt.GetJournalAllocations.Query(j.JID)
	Errcheck(err)
	defer rows.Close()
	j.JA = make([]JournalAllocation, 0)
	for rows.Next() {
		var a JournalAllocation
		ReadJournalAllocations(rows, &a)
		j.JA = append(j.JA, a)
	}
}

// GetJournalAllocationByASMID gets the journal allocation record that references
// the supplied ASMID.
func GetJournalAllocationByASMID(id int64) JournalAllocation {
	var a JournalAllocation
	row := RRdb.Prepstmt.GetJournalAllocationByASMID.QueryRow(id)
	ReadJournalAllocation(row, &a)
	return a
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

// // GetPayorLedgerMarkerOnOrBefore returns the LedgerMarker struct for the TCID
// func GetPayorLedgerMarkerOnOrBefore(bid, tcid int64, dt *time.Time) LedgerMarker {
// 	var r LedgerMarker
// 	row := RRdb.Prepstmt.GetPayorLedgerMarkerOnOrBefore.QueryRow(bid, tcid, dt)
// 	ReadLedgerMarker(row, &r)
// 	return r
// }

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

// // LoadPayorLedgerMarker returns the LedgerMarker for the supplied bid,tcid
// // values. It loads the marker on-or-before dt.  If no such LedgerMarker exists,
// // then one will be created.
// //
// // TODO:  If a subsequent LedgerMarker exists and it is marked as the epoch (3) then
// // then it will be updated to normal status as the LedgerMarker just being will
// // created will be the new epoch.
// //
// // INPUTS
// //		bid  - business id
// //		tcid - which payor
// //		dt   - the ledger marker on this date, or the first prior LedgerMarker
// //			   will be loaded and returned.
// //-----------------------------------------------------------------------------
// func LoadPayorLedgerMarker(bid, tcid int64, dt *time.Time) LedgerMarker {
// 	lm := GetPayorLedgerMarkerOnOrBefore(bid, tcid, dt)
// 	if lm.LMID == 0 {
// 		lm.BID = bid
// 		lm.TCID = tcid
// 		lm.Dt = *dt
// 		lm.State = MARKERSTATEORIGIN
// 		err := InsertLedgerMarker(&lm)
// 		if nil != err {
// 			fmt.Printf("LoadRALedgerMarker: Error creating LedgerMarker: %s\n", err.Error())
// 		}
// 	}
// 	return lm
// }

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

// // GetAllLedgerMarkersOnOrBefore returns a map of all ledgermarkers for the supplied business and dat
// func GetAllLedgerMarkersOnOrBefore(bid int64, dt *time.Time) map[int64]LedgerMarker {
// 	var t map[int64]LedgerMarker
// 	t = make(map[int64]LedgerMarker) // this line is absolutely necessary
// 	rows, err := RRdb.Prepstmt.GetAllLedgerMarkersOnOrBefore.Query(bid, dt)
// 	Errcheck(err)
// 	defer rows.Close()
// 	// fmt.Printf("%4s  %4s  %4s  %5s  %10s  %8s\n", "LMID", "LID", "BID", "State", "Dt", "Balance")
// 	for rows.Next() {
// 		var r LedgerMarker
// 		ReadLedgerMarkers(rows, &r)
// 		t[r.LID] = r
// 		// fmt.Printf("%4d  %4d  %4d  %5d  %10s  %8.2f\n", r.LMID, r.LID, r.BID, r.State, r.Dt, r.Balance)
// 	}
// 	Errcheck(rows.Err())
// 	return t
// }

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

// GetLedgerEntriesByJAID returns the GLAccount struct for the supplied LID
func GetLedgerEntriesByJAID(bid, jaid int64) []LedgerEntry {
	rows, err := RRdb.Prepstmt.GetLedgerEntriesByJAID.Query(bid, jaid)
	Errcheck(err)
	var m []LedgerEntry
	for rows.Next() {
		var le LedgerEntry
		ReadLedgerEntries(rows, &le)
		m = append(m, le)
	}
	return m
}

// GetLedgerByGLNo returns the GLAccount struct for the supplied GLNo
func GetLedgerByGLNo(bid int64, s string) GLAccount {
	var a GLAccount
	row := RRdb.Prepstmt.GetLedgerByGLNo.QueryRow(bid, s)
	ReadGLAccount(row, &a)
	return a
}

// GetLedgerByName returns the GLAccount struct for the supplied Name
func GetLedgerByName(bid int64, s string) GLAccount {
	var a GLAccount
	row := RRdb.Prepstmt.GetLedgerByName.QueryRow(bid, s)
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
	Errcheck(RRdb.Prepstmt.GetNoteList.QueryRow(nlid).Scan(&m.NLID, &m.BID, &m.CreateTS, &m.CreateBy, &m.LastModTime, &m.LastModBy))
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
	Errcheck(RRdb.Prepstmt.GetNoteType.QueryRow(ntid).Scan(&t.NTID, &t.BID, &t.Name, &t.CreateTS, &t.CreateBy, &t.LastModTime, &t.LastModBy))
}

// GetAllNoteTypes reads a NoteType structure based for all NoteTypes in the supplied bid
func GetAllNoteTypes(bid int64) []NoteType {
	var m []NoteType
	rows, err := RRdb.Prepstmt.GetAllNoteTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var p NoteType
		Errcheck(rows.Scan(&p.NTID, &p.BID, &p.Name, &p.CreateTS, &p.CreateBy, &p.LastModTime, &p.LastModBy))
		m = append(m, p)
	}
	Errcheck(rows.Err())
	return m
}

//=======================================================
//  P A Y M E N T   T Y P E S
//=======================================================

// // GetPaymentTypes returns a slice of payment types indexed by the PMTID
// func GetPaymentTypes() map[int64]PaymentType {
// 	var t map[int64]PaymentType
// 	t = make(map[int64]PaymentType)
// 	rows, err := RRdb.Dbrr.Query("SELECT PMTID,BID,Name,Description,LastModTime,LastModBy FROM PaymentType")
// 	Errcheck(err)
// 	defer rows.Close()

// 	for rows.Next() {
// 		var a PaymentType
// 		ReadPaymentTypes(rows, &a)
// 		t[a.PMTID] = a
// 	}
// 	Errcheck(rows.Err())
// 	return t
// }

// GetPaymentType reads a PaymentType structure based on the supplied bid and na
func GetPaymentType(id int64, a *PaymentType) {
	ReadPaymentType(RRdb.Prepstmt.GetPaymentType.QueryRow(id), a)
}

// GetPaymentTypeByName reads a PaymentType structure based on the supplied bid and na
func GetPaymentTypeByName(bid int64, name string, a *PaymentType) {
	ReadPaymentType(RRdb.Prepstmt.GetPaymentTypeByName.QueryRow(bid, name), a)
}

// GetPaymentTypesByBusiness returns a slice of payment types indexed by the PMTID for the supplied Business
func GetPaymentTypesByBusiness(bid int64) map[int64]PaymentType {
	var t map[int64]PaymentType
	t = make(map[int64]PaymentType)
	rows, err := RRdb.Prepstmt.GetPaymentTypesByBusiness.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		var a PaymentType
		ReadPaymentTypes(rows, &a)
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
	r := GetReceiptNoAllocations(rcptid)
	GetReceiptAllocations(rcptid, &r)
	return r
}

// GetReceiptAllocation returns a ReceiptAllocation structure for the supplied RCPTID
func GetReceiptAllocation(id int64) ReceiptAllocation {
	var r ReceiptAllocation
	row := RRdb.Prepstmt.GetReceiptAllocation.QueryRow(id)
	ReadReceiptAllocation(row, &r)
	return r
}

// GetReceiptNoAllocations returns a Receipt structure for the supplied RCPTID.
// It does not get the receipt allocations
func GetReceiptNoAllocations(rcptid int64) Receipt {
	var r Receipt
	row := RRdb.Prepstmt.GetReceipt.QueryRow(rcptid)
	ReadReceipt(row, &r)
	return r
}

// GetReceiptDuplicate returns a Receipt structure for the supplied RCPTID
func GetReceiptDuplicate(dt *time.Time, amt float64, docno string) Receipt {
	var r Receipt
	row := RRdb.Prepstmt.GetReceiptDuplicate.QueryRow(dt, amt, docno)
	ReadReceipt(row, &r)
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
		ReadReceiptAllocations(rows, &a)
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

// GetReceiptAllocationsInRAIDDateRange for the supplied RentalAgreement in date range [d1 - d2).
// To do this we select all the ReceiptAllocations that occurred during d1-d2 that involved
// raid.
func GetReceiptAllocationsInRAIDDateRange(bid, raid int64, d1, d2 *time.Time) []ReceiptAllocation {
	rows, err := RRdb.Prepstmt.GetReceiptAllocationsInRAIDDateRange.Query(bid, raid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t = []ReceiptAllocation{}
	for rows.Next() {
		var r ReceiptAllocation
		ReadReceiptAllocations(rows, &r)
		t = append(t, r)
	}
	return t
}

// GetReceiptAllocationsByASMID returns any payment allocation on targeted at the supplied ASMID.
// This call is used primarily to determine how much payment is left to make on a partially paid
// assessment.
func GetReceiptAllocationsByASMID(bid, asmid int64) []ReceiptAllocation {
	rows, err := RRdb.Prepstmt.GetReceiptAllocationsByASMID.Query(bid, asmid)
	Errcheck(err)
	defer rows.Close()
	var t = []ReceiptAllocation{}
	for rows.Next() {
		var r ReceiptAllocation
		ReadReceiptAllocations(rows, &r)
		t = append(t, r)
	}
	return t
}

// GetUnallocatedReceiptsByPayor returns the receipts paid by the supplied payor tcid that
// have not yet been fully allocated.
func GetUnallocatedReceiptsByPayor(bid, tcid int64) []Receipt {
	rows, err := RRdb.Prepstmt.GetUnallocatedReceiptsByPayor.Query(bid, tcid)
	Errcheck(err)
	defer rows.Close()
	var t = []Receipt{}
	for rows.Next() {
		var r Receipt
		ReadReceipts(rows, &r)
		r.RA = make([]ReceiptAllocation, 0) // the receipt may be partially allocated
		GetReceiptAllocations(r.RCPTID, &r)
		t = append(t, r)
	}
	return t
}

// GetPayorUnallocatedReceiptsCount returns a count of unallocated receipts for the supplied bid & tcid
func GetPayorUnallocatedReceiptsCount(bid, tcid int64) int {
	var i int
	row := RRdb.Prepstmt.GetPayorUnallocatedReceiptsCount.QueryRow(bid, tcid)
	row.Scan(&i)
	return i
}

//=======================================================
//  R E N T A B L E
//=======================================================

// GetRentableByID reads a Rentable structure based on the supplied Rentable id
func GetRentableByID(rid int64, r *Rentable) {
	row := RRdb.Prepstmt.GetRentable.QueryRow(rid)
	Errcheck(ReadRentable(row, r))
}

// GetRentable reads and returns a Rentable structure based on the supplied Rentable id
func GetRentable(rid int64) Rentable {
	var r Rentable
	GetRentableByID(rid, &r)
	return r
}

// GetRentableByName reads and returns a Rentable structure based on the supplied Rentable id
func GetRentableByName(name string, bid int64) (Rentable, error) {
	var r Rentable
	row := RRdb.Prepstmt.GetRentableByName.QueryRow(name, bid)
	err := ReadRentable(row, &r)
	return r, err
}

// GetRentableTypeDown returns the values needed for typedown controls:
// input:   bid - business
//            s - string or substring to search for
//        limit - return no more than this many matches
// return a slice of RentableTypeDowns and an error.
func GetRentableTypeDown(bid int64, s string, limit int) ([]RentableTypeDown, error) {
	var m []RentableTypeDown
	s = "%" + s + "%"
	rows, err := RRdb.Prepstmt.GetRentableTypeDown.Query(bid, s, limit)
	if err != nil {
		return m, err
	}
	defer rows.Close()
	for rows.Next() {
		var t RentableTypeDown
		ReadRentableTypeDown(rows, &t)
		m = append(m, t)
	}
	return m, nil
}

// GetXRentable reads an XRentable structure based on the RID.
func GetXRentable(rid int64, x *XRentable) {
	if x.R.RID == 0 && rid > 0 {
		GetRentableByID(rid, &x.R)
	}
	x.S = GetAllRentableSpecialtyRefs(x.R.BID, x.R.RID)
}

// GetRentableUser returns a Rentable User record with the supplied RUID
func GetRentableUser(ruid int64) (RentableUser, error) {
	row := RRdb.Prepstmt.GetRentableUser.QueryRow(ruid)
	var r RentableUser
	err := ReadRentableUser(row, &r)
	return r, err
}

// GetRentableUserByRBT returns a Rentable User record matching the supplied
// RID, BID, TCID
func GetRentableUserByRBT(rid, bid, tcid int64) (RentableUser, error) {
	row := RRdb.Prepstmt.GetRentableUserByRBT.QueryRow(rid, bid, tcid)
	var r RentableUser
	err := ReadRentableUser(row, &r)
	return r, err
}

// GetRentableSpecialtyTypeByName returns a list of specialties associated with the supplied Rentable
func GetRentableSpecialtyTypeByName(bid int64, name string) RentableSpecialty {
	var rsp RentableSpecialty
	err := RRdb.Prepstmt.GetRentableSpecialtyTypeByName.QueryRow(bid, name).Scan(&rsp.RSPID, &rsp.BID, &rsp.Name, &rsp.Fee, &rsp.Description, &rsp.CreateTS, &rsp.CreateBy)
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
	err := RRdb.Prepstmt.GetRentableSpecialtyType.QueryRow(rspid).Scan(&rs.RSPID, &rs.BID, &rs.Name, &rs.Fee, &rs.Description, &rs.CreateTS, &rs.CreateBy)
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
		Errcheck(rows.Scan(&a.BID, &a.RID, &a.RSPID, &a.DtStart, &a.DtStop, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy))
		rs = append(rs, a)
	}
	Errcheck(rows.Err())
	return rs
}

// GetRentableTypeRef gets RentableTypeRef record for given RTRID -- RentableTypeRef ID (unique ID)
func GetRentableTypeRef(rtrid int64) (RentableTypeRef, error) {
	var rtr RentableTypeRef
	row := RRdb.Prepstmt.GetRentableTypeRef.QueryRow(rtrid)
	err := ReadRentableTypeRef(row, &rtr)
	return rtr, err
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

// GetRTRefs performs the query over the supplied rows and returns a slice of result records
func GetRTRefs(rows *sql.Rows) []RentableTypeRef {
	var rs []RentableTypeRef
	defer rows.Close()
	for rows.Next() {
		var a RentableTypeRef
		Errcheck(ReadRentableTypeRefs(rows, &a))
		rs = append(rs, a)
	}
	Errcheck(rows.Err())
	return rs
}

// GetRentableTypeRefsByRange loads all the RentableTypeRef records that overlap the supplied time range
// and returns them in an array
func GetRentableTypeRefsByRange(RID int64, d1, d2 *time.Time) []RentableTypeRef {
	rows, err := RRdb.Prepstmt.GetRentableTypeRefsByRange.Query(RID, d1, d2)
	Errcheck(err)
	return GetRTRefs(rows)
}

// GetRentableTypeRefs loads all the RentableTypeRef records for a particular
func GetRentableTypeRefs(RID int64) []RentableTypeRef {
	rows, err := RRdb.Prepstmt.GetRentableTypeRefs.Query(RID)
	Errcheck(err)
	return GetRTRefs(rows)
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

// GetRentableStatus gets RentableStatus record for given RSID -- RentableStatus ID (unique ID)
func GetRentableStatus(rsid int64) (RentableStatus, error) {
	var rs RentableStatus
	row := RRdb.Prepstmt.GetRentableStatus.QueryRow(rsid)
	err := ReadRentableStatus(row, &rs)
	return rs, err
}

// GetRentableStatusRows loads all the RentableStatus records for rows
func GetRentableStatusRows(rows *sql.Rows) []RentableStatus {
	var rs []RentableStatus
	defer rows.Close()
	for rows.Next() {
		var a RentableStatus
		ReadRentableStatuss(rows, &a)
		rs = append(rs, a)
	}
	Errcheck(rows.Err())
	return rs
}

// GetRentableStatusByRange loads all the RentableStatus records that overlap the supplied time range
func GetRentableStatusByRange(RID int64, d1, d2 *time.Time) []RentableStatus {
	rows, err := RRdb.Prepstmt.GetRentableStatusByRange.Query(RID, d1, d2)
	Errcheck(err)
	return GetRentableStatusRows(rows)
}

// GetAllRentableStatus loads all the RentableStatus records that overlap the supplied time range
func GetAllRentableStatus(RID int64) []RentableStatus {
	rows, err := RRdb.Prepstmt.GetAllRentableStatus.Query(RID)
	Errcheck(err)
	return GetRentableStatusRows(rows)
}

//=======================================================
//  R E N T A B L E   T Y P E
//=======================================================

// GetRentableType returns characteristics of the Rentable
func GetRentableType(rtid int64, rt *RentableType) error {
	err := RRdb.Prepstmt.GetRentableType.QueryRow(rtid).Scan(&rt.RTID, &rt.BID, &rt.Style, &rt.Name, &rt.RentCycle,
		&rt.Proration, &rt.GSRPC, &rt.ManageToBudget, &rt.CreateTS, &rt.CreateBy, &rt.LastModTime, &rt.LastModBy)
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
	err := RRdb.Prepstmt.GetRentableTypeByStyle.QueryRow(name, bid).Scan(&rt.RTID, &rt.BID, &rt.Style, &rt.Name,
		&rt.RentCycle, &rt.Proration, &rt.GSRPC, &rt.ManageToBudget, &rt.CreateTS, &rt.CreateBy, &rt.LastModTime, &rt.LastModBy)
	return rt, err
}

// GetBusinessRentableTypes returns a slice of RentableType indexed by the RTID
func GetBusinessRentableTypes(bid int64) map[int64]RentableType {
	var t map[int64]RentableType
	t = make(map[int64]RentableType)
	rows, err := RRdb.Prepstmt.GetAllBusinessRentableTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var a RentableType
		Errcheck(rows.Scan(&a.RTID, &a.BID, &a.Style, &a.Name, &a.RentCycle, &a.Proration, &a.GSRPC,
			&a.ManageToBudget, &a.CreateTS, &a.CreateBy, &a.LastModTime, &a.LastModBy))
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
		// Errcheck(rows.Scan(&a.RTID, &a.MarketRate, &a.DtStart, &a.DtStop))
		Errcheck(ReadRentableMarketRates(rows, &a))
		if a.DtStart.After(LatestMRDTStart) {
			LatestMRDTStart = a.DtStart
			rt.MRCurrent = a.MarketRate
		}
		rt.MR = append(rt.MR, a)
	}
	Errcheck(rows.Err())
}

// GetRentableMarketRateInstance returns instance of rentableMarketRate for given RMRID
func GetRentableMarketRateInstance(rmrid int64) (RentableMarketRate, error) {
	var rmr RentableMarketRate
	row := RRdb.Prepstmt.GetRentableMarketRateInstance.QueryRow(rmrid)
	err := ReadRentableMarketRate(row, &rmr)
	return rmr, err
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

// GetRentableUsersInRange returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentableUsersInRange(rid int64, d1, d2 *time.Time) []RentableUser {
	rows, err := RRdb.Prepstmt.GetRentableUsersInRange.Query(rid, d1, d2)
	Errcheck(err)
	defer rows.Close()
	var t []RentableUser
	// t = make([]RentableUser, 0)
	for rows.Next() {
		var r RentableUser
		ReadRentableUsers(rows, &r)
		// Errcheck(rows.Scan(&r.RID, &r.TCID, &r.DtStart, &r.DtStop))
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

	m := GetRentalAgreementPayorsInRange(raid, d1, d2)
	r.P = make([]XPerson, 0)
	for i := 0; i < len(m); i++ {
		var xp XPerson
		GetXPerson(m[i].TCID, &xp)
		r.P = append(r.P, xp)
	}

	for j := 0; j < len(r.R); j++ {
		n := GetRentableUsersInRange(r.R[j].R.RID, d1, d2)
		r.T = make([]XPerson, 0)
		for i := 0; i < len(n); i++ {
			var xp XPerson
			GetXPerson(n[i].TCID, &xp)
			r.T = append(r.T, xp)
		}
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
		ReadRentalAgreementRentables(rows, &r)
		// Errcheck(rows.Scan(&r.RAID, &r.BID, &r.RID, &r.CLID, &r.ContractRent, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRARentableForDate gets the RentalAgreementRentable plus the associated rentables and payors for the
// time period specified
func GetRARentableForDate(raid int64, d1 *time.Time, rar *RentalAgreementRentable) error {
	row := RRdb.Prepstmt.GetRARentableForDate.QueryRow(raid, d1, d1)
	return ReadRentalAgreementRentable(row, rar)
}

// GetRentalAgreementRentable returns Rentable record matching the supplied RARID
func GetRentalAgreementRentable(rarid int64) (RentalAgreementRentable, error) {
	row := RRdb.Prepstmt.GetRentalAgreementRentable.QueryRow(rarid)
	var r RentalAgreementRentable
	err := ReadRentalAgreementRentable(row, &r)
	return r, err
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
		ReadRentalAgreementRentables(rows, &r)
		// Errcheck(rows.Scan(&r.RAID, &r.BID, &r.RID, &r.CLID, &r.ContractRent, &r.DtStart, &r.DtStop))
		t = append(t, r)
	}
	return t
}

// GetRentalAgreementPayorByRBT returns Rental Agreement Payor record matching the supplied
// RAID, BID, TCID
func GetRentalAgreementPayorByRBT(raid, bid, tcid int64) (RentalAgreementPayor, error) {
	row := RRdb.Prepstmt.GetRentalAgreementPayorByRBT.QueryRow(raid, bid, tcid)
	var r RentalAgreementPayor
	err := ReadRentalAgreementPayor(row, &r)
	return r, err
}

// GetRentalAgreementPayor returns Rental Agreement Payor record matching the supplied id
func GetRentalAgreementPayor(id int64) (RentalAgreementPayor, error) {
	row := RRdb.Prepstmt.GetRentalAgreementPayor.QueryRow(id)
	var r RentalAgreementPayor
	err := ReadRentalAgreementPayor(row, &r)
	return r, err
}

// GetRentalAgreementPayorsInRange returns an array of payors (in the form of payors) associated with the supplied RentalAgreement ID
// during the time range d1-d2
func GetRentalAgreementPayorsInRange(raid int64, d1, d2 *time.Time) []RentalAgreementPayor {
	rows, err := RRdb.Prepstmt.GetRentalAgreementPayorsInRange.Query(raid, d1, d2)
	Errcheck(err)
	return GetRentalAgreementPayorsByRows(rows)
}

// GetRentalAgreementsByPayor returns an array of RentalAgreementPayor where the supplied
// TCID is a payor on the specified date
func GetRentalAgreementsByPayor(bid, tcid int64, dt *time.Time) []RentalAgreementPayor {
	rows, err := RRdb.Prepstmt.GetRentalAgreementsByPayor.Query(bid, tcid, dt, dt)
	Errcheck(err)
	return GetRentalAgreementPayorsByRows(rows)
}

// GetRentalAgreementPayorsByRows returns an array of RentalAgreementPayor records
// that were matched by the supplied sql.Rows
func GetRentalAgreementPayorsByRows(rows *sql.Rows) []RentalAgreementPayor {
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

// GetTransactantTypeDown returns the values needed for typedown controls:
// input:   bid - business
//            s - string or substring to search for
//        limit - return no more than this many matches
// return a slice of TransactantTypeDowns and an error.
func GetTransactantTypeDown(bid int64, s string, limit int) ([]TransactantTypeDown, error) {
	var m []TransactantTypeDown
	s = "%" + s + "%"
	rows, err := RRdb.Prepstmt.GetTransactantTypeDown.Query(bid, s, s, s, s, limit)
	if err != nil {
		return m, err
	}
	defer rows.Close()
	for rows.Next() {
		var t TransactantTypeDown
		ReadTransactantTypeDowns(rows, &t)
		m = append(m, t)
	}
	return m, nil
}

// GetTCIDByNote used to get TCID from Note Comment
// originally to get it from people csv Notes field
func GetTCIDByNote(cmt string) int {
	var id int
	rows, err := RRdb.Prepstmt.FindTCIDByNote.Query(cmt)
	Errcheck(err)
	defer rows.Close()

	// just return first, in case of duplicate
	// TODO: need to verify
	for rows.Next() {
		ReadTCIDByNote(rows, &id)
		return id
	}
	return id
}

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
func GetTransactant(tid int64, t *Transactant) error {
	return ReadTransactant(RRdb.Prepstmt.GetTransactant.QueryRow(tid), t)
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

// func GetRentalAgreementGridInfo(raid int64, d1, d2 *time.Time) []RentalAgreementGrid {
// 	var m []RentalAgreementGrid
// 	rows, err := RRdb.Prepstmt.UIRAGrid(raid, d1, d2)
// 	Errcheck(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var t RentalAgreementGrid
// 		ReadRentalAgreementGrids(rows, &t)
// 		m = append(m, &t)
// 	}
// 	return m
// }

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

// GetCountBusinessCustomAttrRefs get total count for CustomAttrRefs
// with particular associated business
func GetCountBusinessCustomAttrRefs(bid int64) int {
	var id int
	rows, err := RRdb.Prepstmt.CountBusinessCustomAttrRefs.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		ReadCountBusinessCustomAttrRefs(rows, &id)
		return id
	}
	return id
}

// GetCountBusinessCustomAttributes get total count for CustomAttributes
// with particular associated business
func GetCountBusinessCustomAttributes(bid int64) int {
	var id int
	rows, err := RRdb.Prepstmt.CountBusinessCustomAttributes.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		ReadCountBusinessCustomAttributes(rows, &id)
		return id
	}
	return id
}

// GetCountBusinessRentableTypes get total count for RentableTypes
// with particular associated business
func GetCountBusinessRentableTypes(bid int64) int {
	var id int
	rows, err := RRdb.Prepstmt.CountBusinessRentableTypes.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		ReadCountBusinessRentableTypes(rows, &id)
		return id
	}
	return id
}

// GetCountBusinessTransactants get total count for Transactants
// with particular associated business
func GetCountBusinessTransactants(bid int64) int {
	var id int
	rows, err := RRdb.Prepstmt.CountBusinessTransactants.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		ReadCountBusinessTransactants(rows, &id)
		return id
	}
	return id
}

// GetCountBusinessRentables get total count for Rentables
// with particular associated business
func GetCountBusinessRentables(bid int64) int {
	var id int
	rows, err := RRdb.Prepstmt.CountBusinessRentables.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		ReadCountBusinessRentables(rows, &id)
		return id
	}
	return id
}

// GetCountBusinessRentalAgreements get total count for RentalAgreements
// with particular associated business
func GetCountBusinessRentalAgreements(bid int64) int {
	var id int
	rows, err := RRdb.Prepstmt.CountBusinessRentalAgreements.Query(bid)
	Errcheck(err)
	defer rows.Close()

	for rows.Next() {
		ReadCountBusinessRentalAgreements(rows, &id)
		return id
	}
	return id
}
