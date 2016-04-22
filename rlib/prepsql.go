package rlib

import "fmt"

func buildPreparedStatements() {
	var err error
	// Prepare("select deduction from deductions where uid=?")
	// Prepare("select type from compensation where uid=?")
	// Prepare("INSERT INTO compensation (uid,type) VALUES(?,?)")
	// Prepare("DELETE FROM compensation WHERE UID=?")
	// Prepare("update classes set Name=?,Designation=?,Description=?,lastmodby=? where ClassCode=?")
	// Errcheck(err)

	RRdb.Prepstmt.GetRentalAgreementByBusiness, err = RRdb.dbrr.Prepare("SELECT RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetUnit, err = RRdb.dbrr.Prepare("SELECT UNITID,BLDGID,UTID,RID,AVAILID,LastModTime,LastModBy FROM unit where UNITID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTransactant, err = RRdb.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTenant, err = RRdb.dbrr.Prepare("SELECT TID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyAddressEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber FROM tenant where TID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentable, err = RRdb.dbrr.Prepare("SELECT RID,LID,RTID,BID,UNITID,Name,Assignment,Report,DefaultOccType,OccType,LastModTime,LastModBy FROM rentable where RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetProspect, err = RRdb.dbrr.Prepare("SELECT PRSPID,TCID,ApplicationFee FROM prospect where PRSPID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetPayor, err = RRdb.dbrr.Prepare("SELECT PID,TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerZipcode,Occupation,LastModTime,LastModBy FROM payor where PID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetUnitSpecialties, err = RRdb.dbrr.Prepare("SELECT USPID FROM unitspecialties where BID=? and UNITID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetUnitSpecialtyType, err = RRdb.dbrr.Prepare("SELECT USPID,BID,Name,Fee,Description FROM unitspecialtytypes where USPID=?")
	Errcheck(err)

	//===============================
	//  Rentable Type
	//===============================
	RRdb.Prepstmt.GetRentableType, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes where RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByName, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes where Name=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableTypes, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableType, err = RRdb.dbrr.Prepare("INSERT INTO rentabletypes (RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)

	//===============================
	//  Receipt
	//===============================
	RRdb.Prepstmt.GetUnitReceipts, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment FROM receipt WHERE RAID=? and Dt>=? and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceipt, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment FROM receipt WHERE RCPTID=?")
	Errcheck(err)

	//===============================
	//  Assessments
	//===============================
	RRdb.Prepstmt.GetUnitAssessments, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE UNITID=? and Stop >= ? and Start < ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE RID=? and Stop >= ? and Start < ?")
	Errcheck(err)

	//===============================
	//  AssessmentType
	//===============================
	RRdb.Prepstmt.GetAssessmentType, err = RRdb.dbrr.Prepare("SELECT ASMTID,Name,Description,LastModTime,LastModBy FROM assessmenttypes WHERE ASMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentTypeByName, err = RRdb.dbrr.Prepare("SELECT ASMTID,Name,Description,LastModTime,LastModBy FROM assessmenttypes WHERE Name=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertAssessmentType, err = RRdb.dbrr.Prepare("INSERT INTO assessmenttypes (Name,Description,LastModBy) VALUES(?,?,?)")
	Errcheck(err)

	s := fmt.Sprintf("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE (ASMTID=%d or ASMTID=%d) and UNITID=?", SECURITYDEPOSIT, SECURITYDEPOSITASSESSMENT)
	RRdb.Prepstmt.GetSecurityDepositAssessment, err = RRdb.dbrr.Prepare(s)
	Errcheck(err)
	// RRdb.Prepstmt.GetUnitRentalAgreements, err = RRdb.dbrr.Prepare("SELECT RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where unitid=? and RentalStop > ? and RentalStart < ?")
	// Errcheck(err)
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.dbrr.Prepare("SELECT RID,LID,RTID,BID,UNITID,Name,Assignment,Report,LastModTime,LastModBy FROM rentable WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE BID=?")
	/* Address,Address2,City,State,PostalCode,Country,Phone, */
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE DES=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessSpecialtyTypes, err = RRdb.dbrr.Prepare("SELECT USPID,BID,Name,Fee,Description FROM unitspecialtytypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,UNITID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE BID=? and Start<? and Stop>?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.dbrr.Prepare("SELECT RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement where RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment from receipt where BID=? and Dt >= ? and DT < ?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from receiptallocation where RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocation, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from journalallocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocations, err = RRdb.dbrr.Prepare("SELECT JAID,JID,RID,Amount,ASMID,AcctRule from journalallocation WHERE JID=?")
	Errcheck(err)

	//===============================
	//  RentableMarketRates
	//===============================
	RRdb.Prepstmt.GetRentableMarketRates, err = RRdb.dbrr.Prepare("SELECT RTID,MarketRate,DtStart,DtStop from rentablemarketrate WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableMarketRates, err = RRdb.dbrr.Prepare("INSERT INTO rentablemarketrate (RTID,MarketRate,DtStart,DtStop) VALUES(?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.GetAssessment, err = RRdb.dbrr.Prepare("SELECT ASMID, BID, RID, UNITID, ASMTID, RAID, Amount, Start, Stop, Frequency, ProrationMethod, AcctRule,Comment, LastModTime, LastModBy from assessments where ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker where JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournal, err = RRdb.dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal where JID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementRentables, err = RRdb.dbrr.Prepare("SELECT RAID,RID,UNITID,DtStart,DtStop from agreementrentables where RID=? and ?<DtStop and ?>DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementPayors, err = RRdb.dbrr.Prepare("SELECT RAID,PID,DtStart,DtStop from agreementpayors where RAID=? and ?<DtStop and ?>DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,UNITID,DtStart,DtStop from agreementrentables where RID=? and ?<DtStop and ?>DtStart")
	Errcheck(err)

	RRdb.Prepstmt.InsertJournal, err = RRdb.dbrr.Prepare("INSERT INTO journal (BID,RAID,Dt,Amount,Type,ID,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
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

	RRdb.Prepstmt.GetLedgerMarkerByGLNo, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name FROM ledgermarker WHERE BID=? and GLNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLatestLedgerMarkerByGLNo, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name FROM ledgermarker WHERE BID=? and GLNumber=? ORDER BY DtStop DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerByGLNoDateRange, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name FROM ledgermarker WHERE BID=? and GLNumber=? and DtStop>? and DtStart<? ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkers, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name from ledgermarker WHERE BID=? ORDER BY LMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerMarkersInRange, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name from ledgermarker WHERE BID=? and DtStop>? and DtStart<?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerInitList, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name from ledgermarker WHERE BID=? and State=3 ORDER BY GLNumber ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetDefaultLedgerMarkers, err = RRdb.dbrr.Prepare("SELECT LMID,BID,PID,GLNumber,State,DtStart,DtStop,Balance,Type,Name FROM ledgermarker WHERE BID=? and Type>=10 ORDER BY DtStop DESC")
	Errcheck(err)

	RRdb.Prepstmt.GetAllLedgersInRange, err = RRdb.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy from ledger WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerInRangeByGLNo, err = RRdb.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy from ledger WHERE BID=? and GLNumber=? and ?<=Dt and Dt<? ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedger, err = RRdb.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy FROM ledger where LID=?")
	Errcheck(err)

	RRdb.Prepstmt.DeleteLedgerEntry, err = RRdb.dbrr.Prepare("DELETE FROM ledger WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerMarker, err = RRdb.dbrr.Prepare("DELETE FROM ledgermarker WHERE LMID=?")
	Errcheck(err)

	RRdb.Prepstmt.InsertLedger, err = RRdb.dbrr.Prepare("INSERT INTO ledger (BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertLedgerMarker, err = RRdb.dbrr.Prepare("INSERT INTO ledgermarker (BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name) VALUES(?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)

	RRdb.Prepstmt.InsertBusiness, err = RRdb.dbrr.Prepare("INSERT INTO business (DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModBy) VALUES(?,?,?,?,?)")
	Errcheck(err)

	//==========================================
	// PHONEBOOK
	//==========================================
	RRdb.PBsql.GetCompanyByDesignation, err = RRdb.dbdir.Prepare("SELECT CoCode,LegalName,CommonName,Address,Address2,City,State,PostalCode,Country,Phone,Fax,Email,Designation,Active,EmploysPersonnel,LastModTime,LastModBy FROM companies WHERE Designation=?")
	Errcheck(err)
}
