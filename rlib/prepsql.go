package rlib

import "strings"

// Define all the SQL prepared statements.

// June 3, 2016 -- As the params change, it's easy to forget to update all the statements with the correct
// field names and the proper number of replacement characters.  I'm starting a convention where the SELECT
// fields are set into a variable and used on all the SELECT statements for that table.  The fields and
// replacement variables for INSERT and UPDATE are derived from the SELECT string.

var mySQLRpl = string("?")
var myRpl = mySQLRpl

// TRNSfields defined fields for Transactant, used in at least one other function
var TRNSfields = string("TCID,USERID,PID,PRSPID,FirstName,MiddleName,LastName,PreferredName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,Website,Notes,LastModTime,LastModBy")

// GenSQLInsertAndUpdateStrings generates a string suitable for SQL INSERT and UPDATE statements given the fields as used in SELECT statements.
//
//    example:  given this string:      "LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModTime,LastModBy"
//              we return these three:  "BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModBy", "?,?,?,?,?,?,?,?,?"  -- use for INSERT
//                                      "BID=?RAID=?,GLNumber=?,Status=?,Type=?,Name=?,AcctType=?,RAAssociated=?,LastModBy=?"      -- use for UPDATE
//
// Note that in this convention, we remove LastModTime from insert and update statements (the db is set up to update them by default) and
// we remove the initial ID as that number is AUTOINCREMENT on INSERTs and is not updated on UPDATE.
func GenSQLInsertAndUpdateStrings(s string) (string, string, string) {
	sa := strings.Split(s, ",")
	s2 := sa[1:]  // skip the ID
	l2 := len(s2) // how many fields
	if l2 > 2 {
		if s2[l2-2] == "LastModTime" { // if the last 2 values are "LastModTime" and "LastModBy"...
			s2[l2-2] = s2[l2-1] // ...move "LastModBy" to the previous slot...
			s2 = s2[:l2-1]      // ...and remove value .  We don't write LastModTime because it is set to automatically update
		}
	}
	s = strings.Join(s2, ",")
	l2 = len(s2) // may have changed

	// now s2 has the proper number of fields.  Produce a
	s3 := myRpl + ","               // start of the INSERT string
	s4 := s2[0] + "=" + myRpl + "," // start of the UPDATE string
	for i := 1; i < l2; i++ {
		s3 += myRpl               // for the INSERT string
		s4 += s2[i] + "=" + myRpl // for the UPDATE string
		if i < l2-1 {             // if there are more fields to come...
			s3 += "," // ...add a comma...
			s4 += "," // ...to both strings
		}
	}
	return s, s3, s4
}

func buildPreparedStatements() {
	var err error
	var s1, s2, s3 string
	// Prepare("select deduction from deductions where uid=?")
	// Prepare("select type from compensation where uid=?")
	// Prepare("INSERT INTO compensation (uid,type) VALUES(?,?)")
	// Prepare("DELETE FROM compensation WHERE UID=?")
	// Prepare("update classes set Name=?,Designation=?,Description=?,lastmodby=? where ClassCode=?")
	// Errcheck(err)

	//===============================
	//  Agreement Payor
	//===============================
	RRdb.Prepstmt.InsertRentalAgreementPayor, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementPayors (RAID,PID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  RentalAgreementPet
	//===============================
	PETflds := "PETID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentalAgreementPet, err = RRdb.Dbrr.Prepare("SELECT " + PETflds + " FROM RentalAgreementPets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreementPets, err = RRdb.Dbrr.Prepare("SELECT " + PETflds + " FROM RentalAgreementPets WHERE RAID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PETflds)
	RRdb.Prepstmt.InsertRentalAgreementPet, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementPets (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentalAgreementPet, err = RRdb.Dbrr.Prepare("UPDATE RentalAgreementPets SET " + s3 + " WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentalAgreementPet, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementPets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAllRentalAgreementPets, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementPets WHERE RAID=?")
	Errcheck(err)

	//===============================
	//  Agreement Rentable
	//===============================
	RRdb.Prepstmt.InsertRentalAgreementRentable, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementRentables (RAID,RID,ContractRent,DtStart,DtStop) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.FindAgreementByRentable, err = RRdb.Dbrr.Prepare("SELECT RAID,RID,ContractRent,DtStart,DtStop FROM RentalAgreementRentables WHERE RID=? AND DtStop>=? AND DtStart<=?")
	Errcheck(err)

	//===============================
	//  Rentable Users
	//===============================
	RRdb.Prepstmt.InsertRentableUser, err = RRdb.Dbrr.Prepare("INSERT INTO RentableUsers (RID,USERID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Assessments
	//===============================
	AsmFlds := "ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RecurCycle,ProrationCycle,AcctRule,Comment,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAssessment, err = RRdb.Dbrr.Prepare("SELECT " + AsmFlds + " FROM Assessments WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + AsmFlds + " FROM Assessments WHERE BID=? and Start<? and Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.Dbrr.Prepare("SELECT " + AsmFlds + " FROM Assessments WHERE RID=? and Stop >= ? and Start < ?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(AsmFlds)
	RRdb.Prepstmt.InsertAssessment, err = RRdb.Dbrr.Prepare("INSERT INTO Assessments (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//===============================
	//  AssessmentType
	//===============================
	RRdb.Prepstmt.GetAssessmentType, err = RRdb.Dbrr.Prepare("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM AssessmentTypes WHERE ASMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentTypeByName, err = RRdb.Dbrr.Prepare("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM AssessmentTypes WHERE Name=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessmentType, err = RRdb.Dbrr.Prepare("INSERT INTO AssessmentTypes (RARequired,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Building
	//===============================
	RRdb.Prepstmt.InsertBuilding, err = RRdb.Dbrr.Prepare("INSERT INTO Building (BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertBuildingWithID, err = RRdb.Dbrr.Prepare("INSERT INTO Building (BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetBuilding, err = RRdb.Dbrr.Prepare("SELECT BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM Building WHERE BLDGID=?")
	Errcheck(err)

	//==========================================
	// Business
	//==========================================
	RRdb.Prepstmt.GetAllBusinesses, err = RRdb.Dbrr.Prepare("SELECT BID,BUD,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM Business")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.Dbrr.Prepare("SELECT BID,BUD,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM Business WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.Dbrr.Prepare("SELECT BID,BUD,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM Business WHERE BUD=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessSpecialtyTypes, err = RRdb.Dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM RentableSpecialtyType WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertBusiness, err = RRdb.Dbrr.Prepare("INSERT INTO Business (BUD,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)

	//==========================================
	// Custom Attribute
	//==========================================
	RRdb.Prepstmt.InsertCustomAttribute, err = RRdb.Dbrr.Prepare("INSERT INTO CustomAttr (Type,Name,Value,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttribute, err = RRdb.Dbrr.Prepare("SELECT CID,Type,Name,Value,LastModTime,LastModBy FROM CustomAttr where CID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttribute, err = RRdb.Dbrr.Prepare("DELETE FROM CustomAttr where CID=?")
	Errcheck(err)

	//==========================================
	// Custom Attribute Ref
	//==========================================
	RRdb.Prepstmt.InsertCustomAttributeRef, err = RRdb.Dbrr.Prepare("INSERT INTO CustomAttrRef (ElementType,ID,CID) VALUES(?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeRefs, err = RRdb.Dbrr.Prepare("SELECT CID FROM CustomAttrRef where ElementType=? and ID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttributeRef, err = RRdb.Dbrr.Prepare("DELETE FROM CustomAttrRef where CID=? and ElementType=? and ID=?")
	Errcheck(err)

	//==========================================
	// JOURNAL
	//==========================================
	RRdb.Prepstmt.GetJournal, err = RRdb.Dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from Journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.Dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from JournalMarker WHERE JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.Dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from JournalMarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournal, err = RRdb.Dbrr.Prepare("INSERT INTO Journal (BID,RAID,Dt,Amount,Type,ID,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.Dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from Journal WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocation, err = RRdb.Dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from JournalAllocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocations, err = RRdb.Dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from JournalAllocation WHERE JID=?")
	Errcheck(err)

	RRdb.Prepstmt.InsertJournalAllocation, err = RRdb.Dbrr.Prepare("INSERT INTO JournalAllocation (JID,RID,Amount,ASMID,AcctRule) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournalMarker, err = RRdb.Dbrr.Prepare("INSERT INTO JournalMarker (BID,State,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.DeleteJournalAllocations, err = RRdb.Dbrr.Prepare("DELETE FROM JournalAllocation WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalEntry, err = RRdb.Dbrr.Prepare("DELETE FROM Journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalMarker, err = RRdb.Dbrr.Prepare("DELETE FROM JournalMarker WHERE JMID=?")
	Errcheck(err)

	//==========================================
	// LEDGER;  GLAccount
	//==========================================
	LDGRfields := "LID,PLID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,AllowPost,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLedgerByGLNo, err = RRdb.Dbrr.Prepare("SELECT " + LDGRfields + " FROM GLAccount WHERE BID=? AND GLNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerByType, err = RRdb.Dbrr.Prepare("SELECT " + LDGRfields + " FROM GLAccount WHERE BID=? AND Type=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRABalanceLedger, err = RRdb.Dbrr.Prepare("SELECT " + LDGRfields + " FROM GLAccount WHERE BID=? AND Type=1 AND RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedger, err = RRdb.Dbrr.Prepare("SELECT " + LDGRfields + " FROM GLAccount WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerList, err = RRdb.Dbrr.Prepare("SELECT " + LDGRfields + " FROM GLAccount WHERE BID=? ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetDefaultLedgers, err = RRdb.Dbrr.Prepare("SELECT " + LDGRfields + " FROM GLAccount WHERE BID=? AND Type>=10 ORDER BY GLNumber ASC")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(LDGRfields)

	RRdb.Prepstmt.InsertLedger, err = RRdb.Dbrr.Prepare("INSERT INTO GLAccount (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedger, err = RRdb.Dbrr.Prepare("DELETE FROM GLAccount WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedger, err = RRdb.Dbrr.Prepare("UPDATE GLAccount SET " + s3 + " WHERE LID=?")
	Errcheck(err)

	//==========================================
	// LEDGER ENTRY
	//==========================================
	LEfields := "LEID,BID,JID,JAID,LID,RAID,Dt,Amount,Comment,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAllLedgerEntriesInRange, err = RRdb.Dbrr.Prepare("SELECT " + LEfields + " from LedgerEntry WHERE BID=? AND ?<=Dt AND Dt<?")
	Errcheck(err)
	// RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo, err = RRdb.Dbrr.Prepare("SELECT " + LEfields + " from LedgerEntry WHERE BID=? AND GLNo=? AND ?<=Dt AND Dt<? ORDER BY JAID ASC")
	// Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesInRangeByLID, err = RRdb.Dbrr.Prepare("SELECT " + LEfields + " from LedgerEntry WHERE BID=? AND LID=? AND ?<=Dt AND Dt<? ORDER BY JAID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesForRAID, err = RRdb.Dbrr.Prepare("SELECT " + LEfields + " FROM LedgerEntry WHERE ?<=Dt AND Dt<? AND RAID=? AND LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntry, err = RRdb.Dbrr.Prepare("SELECT " + LEfields + " FROM LedgerEntry where LEID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(LEfields)
	RRdb.Prepstmt.InsertLedgerEntry, err = RRdb.Dbrr.Prepare("INSERT INTO LedgerEntry (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerEntry, err = RRdb.Dbrr.Prepare("DELETE FROM LedgerEntry WHERE LEID=?")
	Errcheck(err)

	//==========================================
	// LEDGER MARKER
	//==========================================
	LMfields := "LMID,LID,BID,DtStart,DtStop,Balance,State,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLatestLedgerMarkerByLID, err = RRdb.Dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? and LID=? ORDER BY DtStop DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerByDateRange, err = RRdb.Dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? and LID=? and DtStop>? and DtStart<? ORDER BY LID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkers, err = RRdb.Dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? ORDER BY LMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerMarkersInRange, err = RRdb.Dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? and DtStop>? and DtStart<=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerMarker, err = RRdb.Dbrr.Prepare("DELETE FROM LedgerMarker WHERE LMID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerMarker, err = RRdb.Dbrr.Prepare("INSERT INTO LedgerMarker (LID,BID,DtStart,DtStop,Balance,State,LastModBy) VALUES(?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedgerMarker, err = RRdb.Dbrr.Prepare("UPDATE LedgerMarker SET LMID=?,LID=?,BID=?,DtStart=?,DtStop=?,Balance=?,State=?,LastModBy=? WHERE LMID=?")
	Errcheck(err)

	//==========================================
	// PAYMENT TYPES
	//==========================================
	RRdb.Prepstmt.InsertPaymentType, err = RRdb.Dbrr.Prepare("INSERT INTO PaymentTypes (BID,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetPaymentTypesByBusiness, err = RRdb.Dbrr.Prepare("SELECT PMTID,BID,Name,Description,LastModTime,LastModBy FROM PaymentTypes WHERE BID=?")
	Errcheck(err)

	//==========================================
	// PAYOR
	//==========================================
	PAYORfields := "PID,TCID,CreditLimit,TaxpayorID,AccountRep,EligibleFuturePayor,LastModTime,LastModBy"
	RRdb.Prepstmt.GetPayor, err = RRdb.Dbrr.Prepare("SELECT " + PAYORfields + " FROM Payor where PID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PAYORfields)
	RRdb.Prepstmt.InsertPayor, err = RRdb.Dbrr.Prepare("INSERT INTO Payor (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// PROSPECT
	//==========================================
	PRSPfields := "PRSPID,TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModTime,LastModBy"
	RRdb.Prepstmt.GetProspect, err = RRdb.Dbrr.Prepare("SELECT " + PRSPfields + " FROM Prospect where PRSPID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PRSPfields)
	RRdb.Prepstmt.InsertProspect, err = RRdb.Dbrr.Prepare("INSERT INTO Prospect (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// RECEIPT
	//==========================================
	RRdb.Prepstmt.GetReceipt, err = RRdb.Dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy FROM Receipt WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.Dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy from Receipt WHERE BID=? and Dt >= ? and DT < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceipt, err = RRdb.Dbrr.Prepare("INSERT INTO Receipt (RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceipt, err = RRdb.Dbrr.Prepare("DELETE FROM Receipt WHERE RCPTID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT ALLOCATION
	//==========================================
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.Dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from ReceiptAllocation WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceiptAllocation, err = RRdb.Dbrr.Prepare("INSERT INTO ReceiptAllocation (RCPTID,Amount,ASMID,AcctRule) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceiptAllocations, err = RRdb.Dbrr.Prepare("DELETE FROM ReceiptAllocation WHERE RCPTID=?")
	Errcheck(err)

	//===============================
	//  Rentable
	//===============================
	RNTfields := "RID,BID,Name,AssignmentTime,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentable, err = RRdb.Dbrr.Prepare("SELECT " + RNTfields + " FROM Rentable WHERE RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableByName, err = RRdb.Dbrr.Prepare("SELECT " + RNTfields + " FROM Rentable WHERE Name=? AND BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + RNTfields + " FROM Rentable WHERE BID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RNTfields)
	RRdb.Prepstmt.InsertRentable, err = RRdb.Dbrr.Prepare("INSERT INTO Rentable (" + s1 + ") VALUES(" + s2 + ")")

	//===============================
	//  Rental Agreement
	//===============================
	RAfields := "RAID,RATID,BID,RentalStart,RentalStop,PossessionStart,PossessionStop,Renewal,SpecialProvisions,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentalAgreementByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + RAfields + " from RentalAgreement where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.Dbrr.Prepare("SELECT " + RAfields + " FROM RentalAgreement WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreementsByRange, err = RRdb.Dbrr.Prepare("SELECT " + RAfields + " FROM RentalAgreement WHERE BID=? AND ?<=RentalStop AND ?>RentalStart")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RAfields)
	RRdb.Prepstmt.InsertRentalAgreement, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreement (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	RRdb.Prepstmt.GetAllRentalAgreements, err = RRdb.Dbrr.Prepare("SELECT RAID from RentalAgreement WHERE BID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement Template
	//===============================
	RATflds := "RATID,BID,RentalTemplateNumber,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAllRentalAgreementTemplates, err = RRdb.Dbrr.Prepare("SELECT " + RATflds + " FROM RentalAgreementTemplate")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplate, err = RRdb.Dbrr.Prepare("SELECT " + RATflds + " FROM RentalAgreementTemplate WHERE RATID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementByRentalTemplateNumber, err = RRdb.Dbrr.Prepare("SELECT " + RATflds + " FROM RentalAgreementTemplate WHERE RentalTemplateNumber=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RATflds)
	RRdb.Prepstmt.InsertRentalAgreementTemplate, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementTemplate (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//===============================
	//  RentableTypeRef
	//===============================
	RRTIDflds := "RID,RTID,RentCycle,ProrationCycle,DtStart,DtStop,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableTypeRefsByRange, err = RRdb.Dbrr.Prepare("SELECT " + RRTIDflds + " FROM RentableTypeRef WHERE RID=? and DtStop>? and DtStart<? ORDER BY DtStart ASC")
	Errcheck(err)

	RRdb.Prepstmt.InsertRentableTypeRef, err = RRdb.Dbrr.Prepare("INSERT INTO RentableTypeRef (RID,RTID,RentCycle,ProrationCycle,DtStart,DtStop,LastModBy) VALUES(?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableTypeRef, err = RRdb.Dbrr.Prepare("UPDATE RentableTypeRef SET RTID=?,RentCycle=?,ProrationCycle=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableTypeRef, err = RRdb.Dbrr.Prepare("DELETE from RentableTypeRef WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)

	//===============================
	//  RentableSpecialtyRef
	//===============================
	RRdb.Prepstmt.GetRentableSpecialtyRefs, err = RRdb.Dbrr.Prepare("SELECT RSPID FROM RentableSpecialtyRef WHERE BID=? and RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialtyRefsByRange, err = RRdb.Dbrr.Prepare("SELECT BID,RID,RSPID,DtStart,DtStop,LastModTime,LastModBy FROM RentableSpecialtyRef WHERE BID=? and RID=? and DtStop>? and DtStart<? ORDER BY DtStart ASC")
	Errcheck(err)

	RRdb.Prepstmt.InsertRentableSpecialtyRef, err = RRdb.Dbrr.Prepare("INSERT INTO RentableSpecialtyRef (BID,RID,RSPID,DtStart,DtStop,LastModBy) VALUES(?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableSpecialtyRef, err = RRdb.Dbrr.Prepare("UPDATE RentableSpecialtyRef SET RSPID=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableSpecialtyRef, err = RRdb.Dbrr.Prepare("DELETE from RentableSpecialtyRef WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)

	//===============================
	//  RentableSpecialtyType
	//===============================
	RRdb.Prepstmt.InsertRentableSpecialtyType, err = RRdb.Dbrr.Prepare("INSERT INTO RentableSpecialtyType (RSPID,BID,Name,Fee,Description) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialtyTypeByName, err = RRdb.Dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM RentableSpecialtyType WHERE BID=? and Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialtyType, err = RRdb.Dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM RentableSpecialtyType WHERE RSPID=?")
	Errcheck(err)

	//===============================
	//  RentableStatus
	//===============================
	RNTSTATUSflds := "RID,DtStart,DtStop,Status,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableStatusByRange, err = RRdb.Dbrr.Prepare("SELECT " + RNTSTATUSflds + " FROM RentableStatus WHERE RID=? and DtStop>? and DtStart<?")
	Errcheck(err)

	RRdb.Prepstmt.InsertRentableStatus, err = RRdb.Dbrr.Prepare("INSERT INTO RentableStatus (RID,DtStart,DtStop,Status,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableStatus, err = RRdb.Dbrr.Prepare("UPDATE RentableStatus SET Status=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableStatus, err = RRdb.Dbrr.Prepare("DELETE from RentableStatus WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)

	//===============================
	//  Rentable Type
	//===============================
	RTYfields := "RTID,BID,Style,Name,RentCycle,Proration,GSPRC,ManageToBudget,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableType, err = RRdb.Dbrr.Prepare("SELECT " + RTYfields + " FROM RentableTypes WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByStyle, err = RRdb.Dbrr.Prepare("SELECT " + RTYfields + " FROM RentableTypes WHERE Style=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableTypes, err = RRdb.Dbrr.Prepare("SELECT " + RTYfields + " FROM RentableTypes WHERE BID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RTYfields)
	RRdb.Prepstmt.InsertRentableType, err = RRdb.Dbrr.Prepare("INSERT INTO RentableTypes (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//===============================
	//  RentableMarketRates
	//===============================
	RRdb.Prepstmt.GetRentableMarketRates, err = RRdb.Dbrr.Prepare("SELECT RTID,MarketRate,DtStart,DtStop from RentableMarketrate WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableMarketRates, err = RRdb.Dbrr.Prepare("INSERT INTO RentableMarketrate (RTID,MarketRate,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.GetRentalAgreementRentables, err = RRdb.Dbrr.Prepare("SELECT RAID,RID,ContractRent,DtStart,DtStop from RentalAgreementRentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementPayors, err = RRdb.Dbrr.Prepare("SELECT RAID,PID,DtStart,DtStop from RentalAgreementPayors WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableUsers, err = RRdb.Dbrr.Prepare("SELECT RID,USERID,DtStart,DtStop from RentableUsers WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.Dbrr.Prepare("SELECT RAID,RID,ContractRent,DtStart,DtStop from RentalAgreementRentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)

	//==========================================
	// RENTER
	//==========================================
	RENTERflds := "USERID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureUser,Industry,Source,LastModTime,LastModBy"
	RRdb.Prepstmt.GetUser, err = RRdb.Dbrr.Prepare("SELECT " + RENTERflds + " FROM User where USERID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RENTERflds)
	RRdb.Prepstmt.InsertUser, err = RRdb.Dbrr.Prepare("INSERT INTO User (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// TRANSACTANT
	//==========================================
	RRdb.Prepstmt.GetTransactant, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllTransactants, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant")
	Errcheck(err)
	RRdb.Prepstmt.FindTransactantByPhoneOrEmail, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant where WorkPhone=? OR CellPhone=? or PrimaryEmail=? or SecondaryEmail=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(TRNSfields)
	RRdb.Prepstmt.InsertTransactant, err = RRdb.Dbrr.Prepare("INSERT INTO Transactant (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTransactant, err = RRdb.Dbrr.Prepare("UPDATE Transactant SET " + s3 + " WHERE TCID=?")
	Errcheck(err)

}
