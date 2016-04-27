package rlib

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
	//===============================
	//  Assessments
	//===============================
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE RID=? and Stop >= ? and Start < ?")
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

	//===============================
	//  Building
	//===============================
	RRdb.Prepstmt.InsertBuilding, err = RRdb.dbrr.Prepare("INSERT INTO building (BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.InsertBuildingWithID, err = RRdb.dbrr.Prepare("INSERT INTO building (BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetBuilding, err = RRdb.dbrr.Prepare("SELECT BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM building WHERE BLDGID=?")
	Errcheck(err)

	//===============================
	//  Rentable
	//===============================
	RRdb.Prepstmt.InsertRentable, err = RRdb.dbrr.Prepare("INSERT INTO rentable (RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,LastModBy) VALUES(?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetRentable, err = RRdb.dbrr.Prepare("SELECT RID,RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,LastModTime,LastModBy FROM rentable WHERE RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableByName, err = RRdb.dbrr.Prepare("SELECT RID,RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,LastModTime,LastModBy FROM rentable WHERE Name=? AND BID=?")
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
	//  Rentable Type
	//===============================
	RRdb.Prepstmt.GetRentableType, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByStyle, err = RRdb.dbrr.Prepare("SELECT RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModTime,LastModBy FROM rentabletypes WHERE Style=? and BID=?")
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

	// RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.dbrr.Prepare("SELECT RID,LID,RTID,BID,Name,Assignment,Report,LastModTime,LastModBy FROM rentable WHERE BID=?")
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.dbrr.Prepare("SELECT RID,RTID,BID,Name,Assignment,Report,LastModTime,LastModBy FROM rentable WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.dbrr.Prepare("SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy FROM business WHERE DES=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessSpecialtyTypes, err = RRdb.dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description FROM rentablespecialtytypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.dbrr.Prepare("SELECT ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod,AcctRule,Comment,LastModTime,LastModBy FROM assessments WHERE BID=? and Start<? and Stop>?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.dbrr.Prepare("SELECT JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE BID=? and ?<=Dt and Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.dbrr.Prepare("SELECT RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModTime,LastModBy from rentalagreement WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.dbrr.Prepare("SELECT RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment from receipt WHERE BID=? and Dt >= ? and DT < ?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.dbrr.Prepare("SELECT RCPTID,Amount,ASMID,AcctRule from receiptallocation WHERE RCPTID=?")
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

	RRdb.Prepstmt.GetAssessment, err = RRdb.dbrr.Prepare("SELECT ASMID, BID, RID, ASMTID, RAID, Amount, Start, Stop, Frequency, ProrationMethod, AcctRule,Comment, LastModTime, LastModBy from assessments WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker WHERE JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.dbrr.Prepare("SELECT JMID,BID,State,DtStart,DtStop from journalmarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournal, err = RRdb.dbrr.Prepare("select JID,BID,RAID,Dt,Amount,Type,ID,Comment,LastModTime,LastModBy from journal WHERE JID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementRentables, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables WHERE RID=? and ?<DtStop and ?>DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementPayors, err = RRdb.dbrr.Prepare("SELECT RAID,PID,DtStart,DtStop from agreementpayors WHERE RAID=? and ?<DtStop and ?>DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.dbrr.Prepare("SELECT RAID,RID,DtStart,DtStop from agreementrentables WHERE RID=? and ?<DtStop and ?>DtStart")
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
	RRdb.Prepstmt.GetLedgerInRangeByGLNo, err = RRdb.dbrr.Prepare("SELECT LID,BID,JID,JAID,GLNumber,Dt,Amount,Comment,LastModTime,LastModBy from ledger WHERE BID=? and GLNumber=? and ?<=Dt and Dt<? ORDER BY JAID ASC")
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

	//==========================================
	// TRANSACTANT
	//==========================================
	RRdb.Prepstmt.InsertTransactant, err = RRdb.dbrr.Prepare("INSERT INTO transactant (TID,PID,PRSPID,FirstName,MiddleName,LastName,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTransactant, err = RRdb.dbrr.Prepare("UPDATE transactant SET TID=?,PID=?,PRSPID=?,FirstName=?,MiddleName=?,LastName=?,PrimaryEmail=?,SecondaryEmail=?,WorkPhone=?,CellPhone=?,Address=?,Address2=?,City=?,State=?,PostalCode=?,Country=?,LastModBy=? WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTransactant, err = RRdb.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant WHERE TCID=?")
	Errcheck(err)

	//==========================================
	// TENANT
	//==========================================
	RRdb.Prepstmt.InsertTenant, err = RRdb.dbrr.Prepare("INSERT INTO tenant (TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetTenant, err = RRdb.dbrr.Prepare("SELECT TID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber FROM tenant where TID=?")
	Errcheck(err)

	//==========================================
	// PAYOR
	//==========================================
	RRdb.Prepstmt.InsertPayor, err = RRdb.dbrr.Prepare("INSERT INTO payor (TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,LastModBy) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetPayor, err = RRdb.dbrr.Prepare("SELECT PID,TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,LastModTime,LastModBy FROM payor where PID=?")
	Errcheck(err)

	//==========================================
	// PROSPECT
	//==========================================
	RRdb.Prepstmt.InsertProspect, err = RRdb.dbrr.Prepare("INSERT INTO prospect (TCID,ApplicationFee,LastModBy) VALUES(?,?,?)")
	Errcheck(err)
	RRdb.Prepstmt.GetProspect, err = RRdb.dbrr.Prepare("SELECT PRSPID,TCID,ApplicationFee,LastModTime,LastModBy FROM prospect where PRSPID=?")
	Errcheck(err)

}
