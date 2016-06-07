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
	RRdb.Prepstmt.InsertAgreementPayor, err = RRdb.dbrr.Prepare("INSERT INTO agreementpayors (RAID,PID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  AgreementPet
	//===============================
	PETflds := "PETID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAgreementPet, err = RRdb.dbrr.Prepare("SELECT " + PETflds + " FROM agreementpets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAgreementPets, err = RRdb.dbrr.Prepare("SELECT " + PETflds + " FROM agreementpets WHERE RAID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PETflds)
	RRdb.Prepstmt.InsertAgreementPet, err = RRdb.dbrr.Prepare("INSERT INTO agreementpets (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateAgreementPet, err = RRdb.dbrr.Prepare("UPDATE agreementpets SET " + s3 + " WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAgreementPet, err = RRdb.dbrr.Prepare("DELETE FROM agreementpets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAllAgreementPets, err = RRdb.dbrr.Prepare("DELETE FROM agreementpets WHERE RAID=?")
	Errcheck(err)

	//===============================
	//  Agreement Rentable
	//===============================
	RRdb.Prepstmt.InsertAgreementRentable, err = RRdb.dbrr.Prepare("INSERT INTO agreementrentables (RAID,RID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.FindAgreementByRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables where RID=? and DtStop>=? and DtStart<=?")
	Errcheck(err)

	//===============================
	//  Agreement Renter
	//===============================
	RRdb.Prepstmt.InsertAgreementRenter, err = RRdb.dbrr.Prepare("INSERT INTO agreementrenters (RAID,RENTERID,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Assessments
	//===============================
	RRdb.Prepstmt.GetAssessment, err = RRdb.dbrr.Prepare("SELECT ASMID, BID, RID, ASMTID, RAID, Amount, Start, Stop, RentalPeriod, ProrationMethod, AcctRule,Comment, LastModTime, LastModBy from assessments WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentalPeriod,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE BID=? and Start<? and Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentalPeriod,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE RID=? and Stop >= ? and Start < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessment, err = RRdb.dbrr.Prepare("INSERT INTO assessments (ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentalPeriod,ProrationMethod,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  AssessmentType
	//===============================
	RRdb.Prepstmt.GetAssessmentType, err = RRdb.dbrr.Prepare("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM assessmenttypes WHERE ASMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentTypeByName, err = RRdb.dbrr.Prepare("SELECT ASMTID,RARequired,Name,Description,LastModTime,LastModBy FROM assessmenttypes WHERE Name=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessmentType, err = RRdb.dbrr.Prepare("INSERT INTO assessmenttypes (RARequired,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Building
	//===============================
	RRdb.Prepstmt.InsertBuilding, err = RRdb.dbrr.Prepare("INSERT INTO building (BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertBuildingWithID, err = RRdb.dbrr.Prepare("INSERT INTO building (BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetBuilding, err = RRdb.dbrr.Prepare("SELECT BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM building WHERE BLDGID=?")
	Errcheck(err)

	//==========================================
	// Business
	//==========================================
	RRdb.Prepstmt.GetAllBusinesses, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM business")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE DES=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessSpecialtyTypes, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertBusiness, err = RRdb.dbrr.Prepare("INSERT INTO business (DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)

	//==========================================
	// Custom Attribute
	//==========================================
	RRdb.Prepstmt.InsertCustomAttribute, err = RRdb.dbrr.Prepare("INSERT INTO customattr (Type,Name,Value,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttribute, err = RRdb.dbrr.Prepare("SELECT CID,Type,Name,Value,LastModTime,LastModBy FROM customattr where CID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttribute, err = RRdb.dbrr.Prepare("DELETE FROM customattr where CID=?")
	Errcheck(err)

	//==========================================
	// Custom Attribute Ref
	//==========================================
	RRdb.Prepstmt.InsertCustomAttributeRef, err = RRdb.dbrr.Prepare("INSERT INTO customattrref (ElementType,ID,CID) VALUES(?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeRefs, err = RRdb.dbrr.Prepare("SELECT CID FROM customattrref where ElementType=? and ID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttributeRef, err = RRdb.dbrr.Prepare("DELETE FROM customattrref where CID=? and ElementType=? and ID=?")
	Errcheck(err)

	//==========================================
	// JOURNAL
	//==========================================
	RRdb.Prepstmt.GetJournal, err = RRdb.dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker WHERE JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournal, err = RRdb.dbrr.Prepare("INSERT INTO journal (BID,RAID,Dt,Amount,Type,ID,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocation, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from journalallocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocations, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from journalallocation WHERE JID=?")
	Errcheck(err)

	RRdb.Prepstmt.InsertJournalAllocation, err = RRdb.dbrr.Prepare("INSERT INTO journalallocation (JID,RID,Amount,ASMID,AcctRule) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertJournalMarker, err = RRdb.dbrr.Prepare("INSERT INTO journalmarker (BID,State,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.DeleteJournalAllocations, err = RRdb.dbrr.Prepare("DELETE FROM journalallocation WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalEntry, err = RRdb.dbrr.Prepare("DELETE FROM journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalMarker, err = RRdb.dbrr.Prepare("DELETE FROM journalmarker WHERE JMID=?")
	Errcheck(err)

	//==========================================
	// LEDGER
	//==========================================
	LDGRfields := "LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLedgerByGLNo, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? and GLNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerByType, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? and Type=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedger, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerList, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetDefaultLedgers, err = RRdb.dbrr.Prepare("SELECT " + LDGRfields + " FROM ledger WHERE BID=? and Type>=10 ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedger, err = RRdb.dbrr.Prepare("INSERT INTO ledger (BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedger, err = RRdb.dbrr.Prepare("DELETE FROM ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedger, err = RRdb.dbrr.Prepare("UPDATE ledger SET BID=?,RAID=?,GLNumber=?,Status=?,Type=?,Name=?,AcctType=?,RAAssociated=?,LastModBy=? WHERE LID=?")
	Errcheck(err)

	//==========================================
	// LEDGER ENTRY
	//==========================================
	LEfields := "LEID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy"
	RRdb.Prepstmt.GetAllLedgerEntriesInRange, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " from ledgerentry WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " from ledgerentry WHERE BID=? and GLNumber=? and ?<=Dt and Dt<? ORDER BY JAID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntry, err = RRdb.dbrr.Prepare("SELECT " + LEfields + " FROM ledgerentry where LEID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerEntry, err = RRdb.dbrr.Prepare("INSERT INTO ledgerentry (BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerEntry, err = RRdb.dbrr.Prepare("DELETE FROM ledgerentry WHERE LEID=?")
	Errcheck(err)

	//==========================================
	// LEDGER MARKER
	//==========================================
	LMfields := "LMID,LID,BID,DtStart,DtStop,Balance,State,LastModTime,LastModBy"
	RRdb.Prepstmt.GetLatestLedgerMarkerByLID, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? and LID=? ORDER BY DtStop DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerByDateRange, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? and LID=? and DtStop>? and DtStart<? ORDER BY LID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkers, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? ORDER BY LMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerMarkersInRange, err = RRdb.dbrr.Prepare("SELECT " + LMfields + " FROM ledgermarker WHERE BID=? and DtStop>? and DtStart<=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerMarker, err = RRdb.dbrr.Prepare("DELETE FROM ledgermarker WHERE LMID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerMarker, err = RRdb.dbrr.Prepare("INSERT INTO ledgermarker (LID,BID,DtStart,DtStop,Balance,State,LastModBy) VALUES(?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedgerMarker, err = RRdb.dbrr.Prepare("UPDATE ledgermarker SET LMID=?,LID=?,BID=?,DtStart=?,DtStop=?,Balance=?,State=?,LastModBy=? WHERE LMID=?")
	Errcheck(err)

	//==========================================
	// PAYMENT TYPES
	//==========================================
	RRdb.Prepstmt.InsertPaymentType, err = RRdb.dbrr.Prepare("INSERT INTO paymenttypes (BID,Name,Description,LastModBy) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetPaymentTypesByBusiness, err = RRdb.dbrr.Prepare("SELECT PMTID,BID,Name,Description,LastModTime,LastModBy FROM paymenttypes WHERE BID=?")
	Errcheck(err)

	//==========================================
	// PAYOR
	//==========================================
	PAYORfields := "PID,TCID,CreditLimit,TaxpayorID,AccountRep,EligibleFuturePayor,LastModTime,LastModBy"
	RRdb.Prepstmt.GetPayor, err = RRdb.dbrr.Prepare("SELECT " + PAYORfields + " FROM payor where PID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PAYORfields)
	RRdb.Prepstmt.InsertPayor, err = RRdb.dbrr.Prepare("INSERT INTO payor (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// PROSPECT
	//==========================================
	PRSPfields := "PRSPID,TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModTime,LastModBy"
	RRdb.Prepstmt.GetProspect, err = RRdb.dbrr.Prepare("SELECT " + PRSPfields + " FROM prospect where PRSPID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(PRSPfields)
	RRdb.Prepstmt.InsertProspect, err = RRdb.dbrr.Prepare("INSERT INTO prospect (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// RECEIPT
	//==========================================
	RRdb.Prepstmt.GetReceipt, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy FROM receipt WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModTime,LastModBy from receipt WHERE BID=? and Dt >= ? and DT < ?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceipt, err = RRdb.dbrr.Prepare("INSERT INTO receipt (RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceipt, err = RRdb.dbrr.Prepare("DELETE FROM receipt WHERE RCPTID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT ALLOCATION
	//==========================================
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from receiptallocation WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertReceiptAllocation, err = RRdb.dbrr.Prepare("INSERT INTO receiptallocation (RCPTID,Amount,ASMID,AcctRule) VALUES(?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceiptAllocations, err = RRdb.dbrr.Prepare("DELETE FROM receiptallocation WHERE RCPTID=?")
	Errcheck(err)

	//===============================
	//  Rentable
	//===============================
	RNTfields := "RID,RTID,BID,Name,AssignmentTime,RentalPeriodDefault,RentalPeriod,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentable, err = RRdb.dbrr.Prepare("SELECT " + RNTfields + " FROM rentable WHERE RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableByName, err = RRdb.dbrr.Prepare("SELECT " + RNTfields + " FROM rentable WHERE Name=? AND BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.dbrr.Prepare("SELECT " + RNTfields + " FROM rentable WHERE BID=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RNTfields)
	RRdb.Prepstmt.InsertRentable, err = RRdb.dbrr.Prepare("INSERT INTO rentable (" + s1 + ") VALUES(" + s2 + ")")

	//===============================
	//  Rental Agreement
	//===============================
	RAfields := "RAID,RATID,BID,RentalStart,RentalStop,PossessionStart,PossessionStop,Renewal,SpecialProvisions,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentalAgreementByBusiness, err = RRdb.dbrr.Prepare("SELECT " + RAfields + " from rentalagreement where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.dbrr.Prepare("SELECT " + RAfields + " from rentalagreement WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentalAgreement, err = RRdb.dbrr.Prepare("INSERT INTO rentalagreement (RATID,BID,RentalStart,RentalStop,PossessionStart,PossessionStop,Renewal,SpecialProvisions,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreements, err = RRdb.dbrr.Prepare("SELECT RAID from rentalagreement WHERE BID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement Template
	//===============================
	RRdb.Prepstmt.GetAllRentalAgreementTemplates, err = RRdb.dbrr.Prepare("SELECT RATID,RentalTemplateNumber,RentalAgreementType,LastModTime,LastModBy FROM rentalagreementtemplate")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplate, err = RRdb.dbrr.Prepare("SELECT RATID,RentalTemplateNumber,RentalAgreementType,LastModTime,LastModBy FROM rentalagreementtemplate WHERE RATID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplateByRefNum, err = RRdb.dbrr.Prepare("SELECT RATID,RentalTemplateNumber,RentalAgreementType,LastModTime,LastModBy FROM rentalagreementtemplate WHERE RentalTemplateNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentalAgreementTemplate, err = RRdb.dbrr.Prepare("INSERT INTO rentalagreementtemplate (RentalTemplateNumber,RentalAgreementType,LastModBy) VALUES(?,?,?)")
	Errcheck(err)

	//===============================
	//  RentableSpecialty
	//===============================
	RRdb.Prepstmt.InsertRentableSpecialtyType, err = RRdb.dbrr.Prepare("INSERT INTO rentablespecialtytypes (RSPID,BID,Name,Fee,Description) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetSpecialtyByName, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE BID=? and Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialties, err = RRdb.dbrr.Prepare("SELECT RSPID FROM rentablespecialties WHERE BID=? and RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialty, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE RSPID=?")
	Errcheck(err)

	//===============================
	//  RentableStatus
	//===============================
	RNTSTATUSflds := "RID,DtStart,DtStop,Status,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableStatusByRange, err = RRdb.dbrr.Prepare("SELECT " + RNTSTATUSflds + " FROM rentablestatus WHERE RID=? and DtStop>? and DtStart<?")
	Errcheck(err)

	RRdb.Prepstmt.InsertRentableStatus, err = RRdb.dbrr.Prepare("INSERT INTO rentablestatus (RID,DtStart,DtStop,Status,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableStatus, err = RRdb.dbrr.Prepare("UPDATE rentablestatus SET Status=?,LastModBy=? WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableStatus, err = RRdb.dbrr.Prepare("DELETE from rentablestatus WHERE RID=? and DtStart=? and DtStop=?")
	Errcheck(err)

	//===============================
	//  Rentable Type
	//===============================
	RTYfields := "RTID,BID,Style,Name,RentalPeriod,Proration,ManageToBudget,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRentableType, err = RRdb.dbrr.Prepare("SELECT " + RTYfields + " FROM rentabletypes WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByStyle, err = RRdb.dbrr.Prepare("SELECT " + RTYfields + " FROM rentabletypes WHERE Style=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableTypes, err = RRdb.dbrr.Prepare("SELECT " + RTYfields + " FROM rentabletypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableType, err = RRdb.dbrr.Prepare("INSERT INTO rentabletypes (RTID,BID,Style,Name,RentalPeriod,Proration,ManageToBudget,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  RentableMarketRates
	//===============================
	RRdb.Prepstmt.GetRentableMarketRates, err = RRdb.dbrr.Prepare("SELECT RTID,MarketRate,DtStart,DtStop from rentablemarketrate WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableMarketRates, err = RRdb.dbrr.Prepare("INSERT INTO rentablemarketrate (RTID,MarketRate,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.GetAgreementRentables, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementPayors, err = RRdb.dbrr.Prepare("SELECT RAID,PID,DtStart,DtStop from agreementpayors WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementRenters, err = RRdb.dbrr.Prepare("SELECT RAID,RENTERID,DtStart,DtStop from agreementrenters WHERE RAID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)

	//==========================================
	// RENTER
	//==========================================
	RENTERflds := "RENTERID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureRenter,Industry,Source,LastModTime,LastModBy"
	RRdb.Prepstmt.GetRenter, err = RRdb.dbrr.Prepare("SELECT " + RENTERflds + " FROM renter where RENTERID=?")
	Errcheck(err)
	s1, s2, s3 = GenSQLInsertAndUpdateStrings(RENTERflds)
	RRdb.Prepstmt.InsertRenter, err = RRdb.dbrr.Prepare("INSERT INTO renter (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//==========================================
	// TRANSACTANT
	//==========================================
	RRdb.Prepstmt.GetTransactant, err = RRdb.dbrr.Prepare("SELECT " + TRNSfields + " FROM transactant WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllTransactants, err = RRdb.dbrr.Prepare("SELECT " + TRNSfields + " FROM transactant")
	Errcheck(err)
	RRdb.Prepstmt.FindTransactantByPhoneOrEmail, err = RRdb.dbrr.Prepare("SELECT " + TRNSfields + " FROM transactant where WorkPhone=? OR CellPhone=? or PrimaryEmail=? or SecondaryEmail=?")
	Errcheck(err)

	s1, s2, s3 = GenSQLInsertAndUpdateStrings(TRNSfields)
	RRdb.Prepstmt.InsertTransactant, err = RRdb.dbrr.Prepare("INSERT INTO transactant (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTransactant, err = RRdb.dbrr.Prepare("UPDATE transactant SET " + s3 + " WHERE TCID=?")
	Errcheck(err)

}
