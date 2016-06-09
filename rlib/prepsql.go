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
var TRNSfields = string("TCID,RENTERID,PID,PRSPID,FirstName,MiddleName,LastName,PreferredName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,Website,Notes,LastModTime,LastModBy")

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
	RRdb.Prepstmt.InsertAgreementPayor, err = RRdb.dbrr.Prepare("INSERT INTO AgreementPayors (RAID,PID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  AgreementPet
	//===============================
	PETflds := "PETID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAgreementPet, err = RRdb.dbrr.Prepare("SELECT " + PETflds + " FROM AgreementPets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAgreementPets, err = RRdb.dbrr.Prepare("SELECT " + PETflds + " FROM AgreementPets WHERE RAID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PETflds)
	RRdb.Prepstmt.InsertAgreementPet, err = RRdb.dbrr.Prepare("INSERT INTO AgreementPets (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateAgreementPet, err = RRdb.dbrr.Prepare("UPDATE AgreementPets SET " + s3 + " WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAgreementPet, err = RRdb.dbrr.Prepare("DELETE FROM AgreementPets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAllAgreementPets, err = RRdb.dbrr.Prepare("DELETE FROM AgreementPets WHERE RAID=?")
	Errcheck(err)

	//===============================
	//  Agreement Rentable
	//===============================
	RRdb.Prepstmt.InsertAgreementRentable, err = RRdb.dbrr.Prepare("INSERT INTO AgreementRentables (RAID,RID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.FindAgreementByRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from AgreementRentables where RID=? and DtStop>=? and DtStart<=?")
	Errcheck(err)

	//===============================
	//  Agreement Renter
	//===============================
	RRdb.Prepstmt.InsertAgreementRenter, err = RRdb.dbrr.Prepare("INSERT INTO AgreementRenters (RAID,RENTERID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Assessments
	//===============================
	RRdb.Prepstmt.GetAssessment, err = RRdb.dbrr.Prepare("SELECT ASMID, BID, RID, ASMTID, RAID, Amount, Start, Stop, RentCycle, ProrationMethod, AcctRule,Comment, LastModTime, LastModBy from Assessments WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentCycle,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM Assessments WHERE BID=? and Start<? and Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentCycle,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM Assessments WHERE RID=? and Stop >= ? and Start < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessment, err = RRdb.dbrr.Prepare("INSERT INTO Assessments (ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentCycle,ProrationMethod,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  AssessmentType
	//===============================
	RRdb.Prepstmt.GetAssessmentType, err = RRdb.dbrr.Prepare("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM AssessmentTypes WHERE ASMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentTypeByName, err = RRdb.dbrr.Prepare("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM AssessmentTypes WHERE Name=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessmentType, err = RRdb.dbrr.Prepare("INSERT INTO AssessmentTypes (RARequired,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Building
	//===============================
	RRdb.Prepstmt.InsertBuilding, err = RRdb.dbrr.Prepare("INSERT INTO Building (BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertBuildingWithID, err = RRdb.dbrr.Prepare("INSERT INTO Building (BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetBuilding, err = RRdb.dbrr.Prepare("SELECT BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM Building WHERE BLDGID=?")
	Errcheck(err)

	//==========================================
	// Business
	//==========================================
	RRdb.Prepstmt.GetAllBusinesses, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM Business")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM Business WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM Business WHERE DES=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessSpecialtyTypes, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM RentableSpecialtyTypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertBusiness, err = RRdb.dbrr.Prepare("INSERT INTO Business (DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)

	//==========================================
	// Custom Attribute
	//==========================================
	RRdb.Prepstmt.InsertCustomAttribute, err = RRdb.dbrr.Prepare("INSERT INTO CustomAttr (Type,Name,Value,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttribute, err = RRdb.dbrr.Prepare("SELECT CID,Type,Name,Value,LastModTime,LastModBy FROM CustomAttr where CID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttribute, err = RRdb.dbrr.Prepare("DELETE FROM CustomAttr where CID=?")
	Errcheck(err)

	//==========================================
	// Custom Attribute Ref
	//==========================================
	RRdb.Prepstmt.InsertCustomAttributeRef, err = RRdb.dbrr.Prepare("INSERT INTO CustomAttrRef (ElementType,ID,CID) VALUES(?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeRefs, err = RRdb.dbrr.Prepare("SELECT CID FROM CustomAttrRef where ElementType=? and ID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttributeRef, err = RRdb.dbrr.Prepare("DELETE FROM CustomAttrRef where CID=? and ElementType=? and ID=?")
	Errcheck(err)

	//==========================================
	// JOURNAL
	//==========================================
	RRdb.Prepstmt.GetJournal, err = RRdb.dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from Journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from JournalMarker WHERE JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from JournalMarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournal, err = RRdb.dbrr.Prepare("INSERT INTO Journal (BID,RAID,Dt,Amount,Type,ID,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from Journal WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocation, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from JournalAllocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocations, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from JournalAllocation WHERE JID=?")
	Errcheck(err)

	RRdb.Prepstmt.InsertJournalAllocation, err = RRdb.dbrr.Prepare("INSERT INTO JournalAllocation (JID,RID,Amount,ASMID,AcctRule) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournalMarker, err = RRdb.dbrr.Prepare("INSERT INTO JournalMarker (BID,State,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.DeleteJournalAllocations, err = RRdb.dbrr.Prepare("DELETE FROM JournalAllocation WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalEntry, err = RRdb.dbrr.Prepare("DELETE FROM Journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalMarker, err = RRdb.dbrr.Prepare("DELETE FROM JournalMarker WHERE JMID=?")
	Errcheck(err)

	//==========================================
	// LEDGER
	//==========================================
	LDGRfields := "LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLedgerByGLNo, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM Ledger WHERE BID=? and GLNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerByType, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM Ledger WHERE BID=? and Type=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedger, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM Ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerList, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM Ledger WHERE BID=? ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetDefaultLedgers, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM Ledger WHERE BID=? and Type>=10 ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedger, err = RRdb.dbrr.Prepare("INSERT INTO Ledger (BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedger, err = RRdb.dbrr.Prepare("DELETE FROM Ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedger, err = RRdb.dbrr.Prepare("UPDATE Ledger SET BID=?,RAID=?,GLNumber=?,Status=?,Type=?,Name=?,AcctType=?,RAAssociated=?,LastModBy=? WHERE LID=?")
	Errcheck(err)

	//==========================================
	// LEDGER ENTRY
	//==========================================
	LEfields := "LEID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAllLedgerEntriesInRange, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " from LedgerEntry WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " from LedgerEntry WHERE BID=? and GLNumber=? and ?<=Dt and Dt<? ORDER BY JAID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntry, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " FROM LedgerEntry where LEID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerEntry, err = RRdb.dbrr.Prepare("INSERT INTO LedgerEntry (BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerEntry, err = RRdb.dbrr.Prepare("DELETE FROM LedgerEntry WHERE LEID=?")
	Errcheck(err)

	//==========================================
	// LEDGER MARKER
	//==========================================
	LMfields := "LMID,LID,BID,DtStart,DtStop,Balance,State,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLatestLedgerMarkerByLID, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? and LID=? ORDER BY DtStop DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerByDateRange, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? and LID=? and DtStop>? and DtStart<? ORDER BY LID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkers, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? ORDER BY LMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerMarkersInRange, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM LedgerMarker WHERE BID=? and DtStop>? and DtStart<=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerMarker, err = RRdb.dbrr.Prepare("DELETE FROM LedgerMarker WHERE LMID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerMarker, err = RRdb.dbrr.Prepare("INSERT INTO LedgerMarker (LID,BID,DtStart,DtStop,Balance,State,LastModBy) VALUES(?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedgerMarker, err = RRdb.dbrr.Prepare("UPDATE LedgerMarker SET LMID=?,LID=?,BID=?,DtStart=?,DtStop=?,Balance=?,State=?,LastModBy=? WHERE LMID=?")
	Errcheck(err)

	//==========================================
	// PAYMENT TYPES
	//==========================================
	RRdb.Prepstmt.InsertPaymentType, err = RRdb.dbrr.Prepare("INSERT INTO PaymentTypes (BID,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetPaymentTypesByBusiness, err = RRdb.dbrr.Prepare("SELECT PMTID,BID,Name,Description,LastModTime,LastModBy FROM PaymentTypes WHERE BID=?")
	Errcheck(err)

	//==========================================
	// PAYOR
	//==========================================
	PAYORfields := "PID,TCID,CreditLimit,TaxpayorID,AccountRep,EligibleFuturePayor,LastModTime,LastModBy"
	RRdb.Prepstmt.GetPayor, err = RRdb.dbrr.Prepare("SELECT " + PAYORfields + " FROM Payor where PID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PAYORfields)
	RRdb.Prepstmt.InsertPayor, err = RRdb.dbrr.Prepare("INSERT INTO Payor (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// PROSPECT
	//==========================================
	PRSPfields := "PRSPID,TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModTime,LastModBy"
	RRdb.Prepstmt.GetProspect, err = RRdb.dbrr.Prepare("SELECT " + PRSPfields + " FROM Prospect where PRSPID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PRSPfields)
	RRdb.Prepstmt.InsertProspect, err = RRdb.dbrr.Prepare("INSERT INTO Prospect (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// RECEIPT
	//==========================================
	RRdb.Prepstmt.GetReceipt, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy FROM Receipt WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy from Receipt WHERE BID=? and Dt >= ? and DT < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceipt, err = RRdb.dbrr.Prepare("INSERT INTO Receipt (RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceipt, err = RRdb.dbrr.Prepare("DELETE FROM Receipt WHERE RCPTID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT ALLOCATION
	//==========================================
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from ReceiptAllocation WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceiptAllocation, err = RRdb.dbrr.Prepare("INSERT INTO ReceiptAllocation (RCPTID,Amount,ASMID,AcctRule) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceiptAllocations, err = RRdb.dbrr.Prepare("DELETE FROM ReceiptAllocation WHERE RCPTID=?")
	Errcheck(err)

	//===============================
	//  Rentable
	//===============================
	RNTfields := "RID,RTID,BID,Name,AssignmentTime,RentalPeriodDefault,RentCycle,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentable, err = RRdb.dbrr.Prepare("SELECT " + RNTfields + " FROM Rentable WHERE RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableByName, err = RRdb.dbrr.Prepare("SELECT " + RNTfields + " FROM Rentable WHERE Name=? AND BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.dbrr.Prepare("SELECT " + RNTfields + " FROM Rentable WHERE BID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RNTfields)
	RRdb.Prepstmt.InsertRentable, err = RRdb.dbrr.Prepare("INSERT INTO Rentable (" + s1 + ") VALUES(" + s2 + ")")

	//===============================
	//  Rental Agreement
	//===============================
	RAfields := "RAID,RATID,BID,RentalStart,RentalStop,PossessionStart,PossessionStop,Renewal,SpecialProvisions,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentalAgreementByBusiness, err = RRdb.dbrr.Prepare("SELECT " + RAfields + " from RentalAgreement where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.dbrr.Prepare("SELECT " + RAfields + " from RentalAgreement WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentalAgreement, err = RRdb.dbrr.Prepare("INSERT INTO RentalAgreement (RATID,BID,RentalStart,RentalStop,PossessionStart,PossessionStop,Renewal,SpecialProvisions,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreements, err = RRdb.dbrr.Prepare("SELECT RAID from RentalAgreement WHERE BID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement Template
	//===============================
	RRdb.Prepstmt.GetAllRentalAgreementTemplates, err = RRdb.dbrr.Prepare("SELECT RATID,RentalTemplateNumber,RentalAgreementType,LastModTime,LastModBy FROM RentalAgreementTemplate")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplate, err = RRdb.dbrr.Prepare("SELECT RATID,RentalTemplateNumber,RentalAgreementType,LastModTime,LastModBy FROM RentalAgreementTemplate WHERE RATID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplateByRefNum, err = RRdb.dbrr.Prepare("SELECT RATID,RentalTemplateNumber,RentalAgreementType,LastModTime,LastModBy FROM RentalAgreementTemplate WHERE RentalTemplateNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentalAgreementTemplate, err = RRdb.dbrr.Prepare("INSERT INTO RentalAgreementTemplate (RentalTemplateNumber,RentalAgreementType,LastModBy) VALUES(?,?,?)")
	Errcheck(err)

	//===============================
	//  RentableRTID
	//===============================
	RRTIDflds := "RID,RTID,DtStart,DtStop,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableRTIDsByRange, err = RRdb.dbrr.Prepare("SELECT " + RRTIDflds + " FROM RentableRTID WHERE RID=? and DtStop>? and DtStart<? ORDER BY DtStart ASC")
	Errcheck(err)

	RRdb.Prepstmt.InsertRentableRTID, err = RRdb.dbrr.Prepare("INSERT INTO RentableRTID (RID,RTID,DtStart,DtStop,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableRTID, err = RRdb.dbrr.Prepare("UPDATE RentableRTID SET RTID=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableRTID, err = RRdb.dbrr.Prepare("DELETE from RentableRTID WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)

	//===============================
	//  RentableSpecialty
	//===============================
	RRdb.Prepstmt.InsertRentableSpecialtyType, err = RRdb.dbrr.Prepare("INSERT INTO RentableSpecialtyTypes (RSPID,BID,Name,Fee,Description) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetSpecialtyByName, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM RentableSpecialtyTypes WHERE BID=? and Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialties, err = RRdb.dbrr.Prepare("SELECT RSPID FROM RentableSpecialties WHERE BID=? and RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialty, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM RentableSpecialtyTypes WHERE RSPID=?")
	Errcheck(err)

	//===============================
	//  RentableStatus
	//===============================
	RNTSTATUSflds := "RID,DtStart,DtStop,Status,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableStatusByRange, err = RRdb.dbrr.Prepare("SELECT " + RNTSTATUSflds + " FROM RentableStatus WHERE RID=? and DtStop>? and DtStart<?")
	Errcheck(err)

	RRdb.Prepstmt.InsertRentableStatus, err = RRdb.dbrr.Prepare("INSERT INTO RentableStatus (RID,DtStart,DtStop,Status,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableStatus, err = RRdb.dbrr.Prepare("UPDATE RentableStatus SET Status=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableStatus, err = RRdb.dbrr.Prepare("DELETE from RentableStatus WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)

	//===============================
	//  Rentable Type
	//===============================
	RTYfields := "RTID,BID,Style,Name,RentCycle,Proration,ManageToBudget,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableType, err = RRdb.dbrr.Prepare("SELECT " + RTYfields + " FROM RentableTypes WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByStyle, err = RRdb.dbrr.Prepare("SELECT " + RTYfields + " FROM RentableTypes WHERE Style=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableTypes, err = RRdb.dbrr.Prepare("SELECT " + RTYfields + " FROM RentableTypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableType, err = RRdb.dbrr.Prepare("INSERT INTO RentableTypes (RTID,BID,Style,Name,RentCycle,Proration,ManageToBudget,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  RentableMarketRates
	//===============================
	RRdb.Prepstmt.GetRentableMarketRates, err = RRdb.dbrr.Prepare("SELECT RTID,MarketRate,DtStart,DtStop from RentableMarketrate WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableMarketRates, err = RRdb.dbrr.Prepare("INSERT INTO RentableMarketrate (RTID,MarketRate,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.GetAgreementRentables, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from AgreementRentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementPayors, err = RRdb.dbrr.Prepare("SELECT RAID,PID,DtStart,DtStop from AgreementPayors WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementRenters, err = RRdb.dbrr.Prepare("SELECT RAID,RENTERID,DtStart,DtStop from AgreementRenters WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from AgreementRentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)

	//==========================================
	// RENTER
	//==========================================
	RENTERflds := "RENTERID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureRenter,Industry,Source,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRenter, err = RRdb.dbrr.Prepare("SELECT " + RENTERflds + " FROM Renter where RENTERID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RENTERflds)
	RRdb.Prepstmt.InsertRenter, err = RRdb.dbrr.Prepare("INSERT INTO Renter (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// TRANSACTANT
	//==========================================
	RRdb.Prepstmt.GetTransactant, err = RRdb.dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllTransactants, err = RRdb.dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant")
	Errcheck(err)
	RRdb.Prepstmt.FindTransactantByPhoneOrEmail, err = RRdb.dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant where WorkPhone=? OR CellPhone=? or PrimaryEmail=? or SecondaryEmail=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(TRNSfields)
	RRdb.Prepstmt.InsertTransactant, err = RRdb.dbrr.Prepare("INSERT INTO Transactant (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTransactant, err = RRdb.dbrr.Prepare("UPDATE Transactant SET " + s3 + " WHERE TCID=?")
	Errcheck(err)

}
