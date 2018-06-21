package rlib

import
// "fmt"

"strings"

// Define all the SQL prepared statements.

// June 3, 2016 -- As the params change, it's easy to forget to update all the statements with the correct
// field names and the proper number of replacement characters.  I'm starting a convention where the SELECT
// fields are set into a variable and used on all the SELECT statements for that table.  The fields and
// replacement variables for INSERT and UPDATE are derived from the SELECT string.

var mySQLRpl = string("?")
var myRpl = mySQLRpl

// TRNSfields defined fields for Transactant, used in at least one other function
var TRNSfields = string("TCID,BID,NLID,FirstName,MiddleName,LastName,PreferredName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,Website,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy")

// ASMTflds defined fields for AssessmentTypes, used in at least one other function
// var ASMTflds = string("ASMTID,RARequired,ManageToBudget,Name,Description,LastModTime,LastModBy")

// GenSQLInsertAndUpdateStrings generates a string suitable for SQL INSERT and UPDATE statements given the fields as used in SELECT statements.
//
//  example:
//	given this string:      "LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,LastModTime,LastModBy"
//  we return these five strings:
//  1)  "BID,RAID,GLNumber,Status,Type,Name,AcctType,LastModBy"                 -- use for SELECT
//  2)  "?,?,?,?,?,?,?,?"  														-- use for INSERT
//  3)  "BID=?RAID=?,GLNumber=?,Status=?,Type=?,Name=?,AcctType=?,LastModBy=?"  -- use for UPDATE
//  4)  "LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,LastModBy", 			-- use for INSERT (no PRIMARYKEY), add "WHERE LID=?"
//  5)  "?,?,?,?,?,?,?,?,?"  													-- use for INSERT (no PRIMARYKEY)
//
// Note that in this convention, we remove LastModTime from insert and update statements (the db is set up to update them by default) and
// we remove the initial ID as that number is AUTOINCREMENT on INSERTs and is not updated on UPDATE.
func GenSQLInsertAndUpdateStrings(s string) (string, string, string, string, string) {
	fields := strings.Split(s, ",")

	// mostly 0th element is ID, but it is not necessary
	s0 := fields[0]
	s2 := fields[1:] // skip the ID

	insertFields := []string{} // fields which are allowed while INSERT
	updateFields := []string{} // fields which are allowed while while UPDATE

	// remove fields which value automatically handled by database while insert and update op.
	for _, fld := range s2 {
		fld = strings.TrimSpace(fld)
		if fld == "" { // if nothing then continue
			continue
		}
		// INSERT FIELDS Inclusion
		if fld != "LastModTime" && fld != "CreateTS" { // remove these fields for INSERT
			insertFields = append(insertFields, fld)
		}
		// UPDATE FIELDS Inclusion
		if fld != "LastModTime" && fld != "CreateTS" && fld != "CreateBy" { // remove these fields for UPDATE
			updateFields = append(updateFields, fld)
		}
	}

	var s3, s4 string
	for i := range insertFields {
		if i == len(insertFields)-1 {
			s3 += myRpl
		} else {
			s3 += myRpl + ","
		}
	}

	for i, uFld := range updateFields {
		if i == len(updateFields)-1 {
			s4 += uFld + "=" + myRpl
		} else {
			s4 += uFld + "=" + myRpl + ","
		}
	}

	// list down insert fields with comma separation
	s = strings.Join(insertFields, ",")

	s5 := s0 + "," + s     // for INSERT where first val is not AUTOINCREMENT
	s6 := s3 + "," + myRpl // for INSERT where first val is not AUTOINCREMENT
	return s, s3, s4, s5, s6
}

func buildPreparedStatements() {
	var err error
	var s1, s2, s3, s4, s5, flds string

	//===============================
	//  AccountRule
	//  AR
	//===============================
	flds = "ARID,BID,Name,ARType,DebitLID,CreditLID,Description,RARequired,DtStart,DtStop,FLAGS,DefaultAmount,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["AR"] = flds
	RRdb.Prepstmt.GetAR, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM AR WHERE ARID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetARByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM AR WHERE BID=? AND Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetARsByType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM AR WHERE BID=? AND ARType=?")
	Errcheck(err)
	RRdb.Prepstmt.GetARsByFLAGS, err = RRdb.Dbrr.Prepare("SELECT DISTINCT " + flds + " FROM AR WHERE BID=? AND (CASE WHEN ? > 0 THEN FLAGS&? ELSE FLAGS=0 END)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllARs, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM AR WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertAR, err = RRdb.Dbrr.Prepare("INSERT INTO AR (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateAR, err = RRdb.Dbrr.Prepare("UPDATE AR SET " + s3 + " WHERE ARID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAR, err = RRdb.Dbrr.Prepare("DELETE FROM AR WHERE ARID=?")
	Errcheck(err)

	//===============================
	//  Assessments
	//===============================
	flds = "ASMID,PASMID,RPASMID,AGRCPTID,BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,InvoiceNo,AcctRule,ARID,FLAGS,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Assessments"] = flds
	RRdb.Prepstmt.GetAssessment, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentInstance, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE Start=? AND PASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentDuplicate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE Start=? AND Amount=? AND PASMID=? AND RID=? AND RAID=? AND ATypeLID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllAssessmentsByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE BID=? AND (PASMID=0 OR RentCycle=0) AND Start<? AND Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRecurringAssessmentsByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE BID=? AND PASMID=0 AND RentCycle>0 AND Start<? AND Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetEpochAssessmentsByRentalAgreement, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE RAID=? AND (PASMID=0 or RentCycle=0)")
	Errcheck(err)
	RRdb.Prepstmt.GetAllSingleInstanceAssessments, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE BID=? AND (PASMID!=0 OR RentCycle=0) AND Start<? AND Stop>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentsByRAIDRID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE PASMID = 0 AND BID=? AND RAID=? AND RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentsByRAIDRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE (RentCycle=0  OR (RentCycle>0 AND PASMID>0)) AND RAID=? AND Stop>=? AND Start<?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentsByRARRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE (RentCycle=0  OR (RentCycle>0 AND PASMID>0)) AND RAID=? AND RID=? AND Stop>=? AND Start<?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableAssessments, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE RID=? AND Stop >= ? AND Start < ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentInstancesByParent, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE PASMID=? AND Stop >= ? AND Start < ?")
	Errcheck(err)
	RRdb.Prepstmt.GetAssessmentFirstInstance, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE PASMID=? ORDER BY Start LIMIT 1")
	Errcheck(err)
	// FLAGS bits 0-1 mean: 0 = unpaid, 1 = partially paid, 2 = fully paid.
	// So, FLAGS & 3 gives us the values of bits 0-1.  if the value is 0 or 1 then the assessment is not yet paid.
	// So (FLAGS & 3) < 2 means that the assessment is not yet paid
	// Note that if FLAGS & 0x3 == 3 then the assessment is an offset and should not be considered for payment
	RRdb.Prepstmt.GetUnpaidAssessmentsByRAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Assessments WHERE RAID=? AND (FLAGS & 3)<2 AND (FLAGS & 4)=0 AND (PASMID!=0 OR RentCycle=0) ORDER BY Start ASC")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertAssessment, err = RRdb.Dbrr.Prepare("INSERT INTO Assessments (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateAssessment, err = RRdb.Dbrr.Prepare("UPDATE Assessments SET " + s3 + " WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAssessment, err = RRdb.Dbrr.Prepare("DELETE from Assessments WHERE ASMID=?")
	Errcheck(err)

	//===============================
	//  Building
	//===============================
	flds = "BLDGID,BID,Address,Address2,City,State,PostalCode,Country,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Building"] = flds
	RRdb.Prepstmt.GetBuilding, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Building WHERE BLDGID=?")
	Errcheck(err)
	s1, s2, _, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertBuilding, err = RRdb.Dbrr.Prepare("INSERT INTO Building (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.InsertBuildingWithID, err = RRdb.Dbrr.Prepare("INSERT INTO Building (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)

	//==========================================
	// Business
	//==========================================
	flds = "BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,ClosePeriodTLID,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Business"] = flds
	RRdb.Prepstmt.GetAllBusinesses, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Business ORDER BY Name ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetBusiness, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Business WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessByDesignation, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Business WHERE BUD=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertBusiness, err = RRdb.Dbrr.Prepare("INSERT INTO Business (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateBusiness, err = RRdb.Dbrr.Prepare("UPDATE Business SET " + s3 + " WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableSpecialtyTypes, err = RRdb.Dbrr.Prepare("SELECT RSPID,BID,Name,Fee,Description,CreateTS,CreateBy,LastModTime,LastModBy FROM RentableSpecialty WHERE BID=?")
	Errcheck(err)

	//==========================================
	// Business Properties
	//==========================================
	flds = "BPID,BID,Name,Data,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["BusinessProperties"] = flds
	RRdb.Prepstmt.GetBusinessProperties, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM BusinessProperties where BPID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetBusinessPropertiesByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM BusinessProperties WHERE Name=? AND BID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertBusinessProperties, err = RRdb.Dbrr.Prepare("INSERT INTO BusinessProperties (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateBusinessPropertiesData, err = RRdb.Dbrr.Prepare("UPDATE BusinessProperties SET Data = JSON_REPLACE(Data, CONCAT('$.', ?), CAST(? AS JSON)) where BPID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteBusinessProperties, err = RRdb.Dbrr.Prepare("DELETE from BusinessProperties WHERE BPID=?")
	Errcheck(err)

	//==========================================
	// Close Period
	//==========================================
	flds = "CPID,BID,TLID,Dt,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["ClosePeriod"] = flds
	RRdb.Prepstmt.GetClosePeriod, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ClosePeriod WHERE CPID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLastClosePeriod, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ClosePeriod WHERE BID=? ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertClosePeriod, err = RRdb.Dbrr.Prepare("INSERT INTO ClosePeriod (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateClosePeriod, err = RRdb.Dbrr.Prepare("UPDATE ClosePeriod SET " + s3 + " WHERE CPID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteClosePeriod, err = RRdb.Dbrr.Prepare("DELETE FROM ClosePeriod WHERE CPID=?")
	Errcheck(err)

	//==========================================
	// Custom Attribute
	//==========================================
	flds = "CID,BID,Type,Name,Value,Units,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["CustomAttr"] = flds
	RRdb.Prepstmt.CountBusinessCustomAttributes, err = RRdb.Dbrr.Prepare("SELECT COUNT(CID) FROM CustomAttr WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttribute, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM CustomAttr WHERE CID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeByVals, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM CustomAttr WHERE Type=? AND Name=? AND Value=? AND Units=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllCustomAttributes, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM CustomAttr")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertCustomAttribute, err = RRdb.Dbrr.Prepare("INSERT INTO CustomAttr (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateCustomAttribute, err = RRdb.Dbrr.Prepare("UPDATE CustomAttr SET " + s3 + " WHERE CID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttribute, err = RRdb.Dbrr.Prepare("DELETE FROM CustomAttr WHERE CID=?")
	Errcheck(err)

	//==========================================
	// Custom Attribute Ref
	//==========================================
	flds = "ElementType,BID,ID,CID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["CustomAttrRef"] = flds
	RRdb.Prepstmt.CountBusinessCustomAttrRefs, err = RRdb.Dbrr.Prepare("SELECT COUNT(CID) FROM CustomAttrRef WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeRefs, err = RRdb.Dbrr.Prepare("SELECT CID FROM CustomAttrRef WHERE ElementType=? and ID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetCustomAttributeRef, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM CustomAttrRef WHERE ElementType=? and ID=? and CID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllCustomAttributeRefs, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM CustomAttrRef")
	Errcheck(err)

	_, _, _, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertCustomAttributeRef, err = RRdb.Dbrr.Prepare("INSERT INTO CustomAttrRef (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteCustomAttributeRef, err = RRdb.Dbrr.Prepare("DELETE FROM CustomAttrRef WHERE CID=? and ElementType=? and ID=?")
	Errcheck(err)

	//==========================================
	// DEPOSIT
	//==========================================
	flds = "DID,BID,DEPID,DPMID,Dt,Amount,ClearedAmount,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Deposit"] = flds
	RRdb.Prepstmt.GetDeposit, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Deposit WHERE DID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllDepositsInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Deposit WHERE BID=? AND ?<=Dt AND Dt<?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertDeposit, err = RRdb.Dbrr.Prepare("INSERT INTO Deposit (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteDeposit, err = RRdb.Dbrr.Prepare("DELETE FROM Deposit WHERE DID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateDeposit, err = RRdb.Dbrr.Prepare("UPDATE Deposit SET " + s3 + " WHERE DID=?")
	Errcheck(err)

	//==========================================
	// DEPOSIT METHOD
	//==========================================
	flds = "DPMID,BID,Method,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["DepositMethod"] = flds
	RRdb.Prepstmt.GetDepositMethod, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DepositMethod WHERE DPMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetDepositMethodByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DepositMethod WHERE BID=? and Method=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllDepositMethods, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DepositMethod WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertDepositMethod, err = RRdb.Dbrr.Prepare("INSERT INTO DepositMethod (" + s1 + ") VALUES (" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateDepositMethod, err = RRdb.Dbrr.Prepare("UPDATE DepositMethod SET " + s3 + " WHERE DPMID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteDepositMethod, err = RRdb.Dbrr.Prepare("DELETE FROM DepositMethod WHERE DPMID=?")
	Errcheck(err)

	//==========================================
	// DEPOSIT PART
	//==========================================
	flds = "DPID,DID,BID,RCPTID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["DepositPart"] = flds
	RRdb.Prepstmt.GetDepositParts, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DepositPart WHERE DID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertDepositPart, err = RRdb.Dbrr.Prepare("INSERT INTO DepositPart (" + s1 + ") VALUES (" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateDepositPart, err = RRdb.Dbrr.Prepare("UPDATE DepositPart SET " + s3 + " WHERE DPID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteDepositPart, err = RRdb.Dbrr.Prepare("DELETE FROM DepositPart WHERE DPID=?")
	Errcheck(err)

	//==========================================
	// DEPOSITORY
	//==========================================
	flds = "DEPID,BID,LID,Name,AccountNo,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Depository"] = flds
	RRdb.Prepstmt.GetDepository, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Depository WHERE DEPID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetDepositoryByAccount, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Depository WHERE BID=? AND AccountNo=?")
	Errcheck(err)
	RRdb.Prepstmt.GetDepositoryByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Depository WHERE BID=? AND Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetDepositoryByLID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Depository WHERE BID=? AND LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllDepositories, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Depository WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertDepository, err = RRdb.Dbrr.Prepare("INSERT INTO Depository (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteDepository, err = RRdb.Dbrr.Prepare("DELETE FROM Depository WHERE DEPID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateDepository, err = RRdb.Dbrr.Prepare("UPDATE Depository SET " + s3 + " WHERE DEPID=?")
	Errcheck(err)

	//==========================================
	// EXPENSE
	//==========================================
	flds = "EXPID,RPEXPID,BID,RID,RAID,Amount,Dt,AcctRule,ARID,FLAGS,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Expense"] = flds

	RRdb.Prepstmt.GetExpense, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Expense WHERE EXPID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertExpense, err = RRdb.Dbrr.Prepare("INSERT INTO Expense (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteExpense, err = RRdb.Dbrr.Prepare("DELETE FROM Expense WHERE EXPID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateExpense, err = RRdb.Dbrr.Prepare("UPDATE Expense SET " + s3 + " WHERE EXPID=?")
	Errcheck(err)

	//==========================================
	// Flow
	//==========================================
	flds = "FlowID,BID,UserRefNo,FlowType,ID,Data,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Flow"] = flds
	RRdb.Prepstmt.GetFlowMetaDataInRange, err = RRdb.Dbrr.Prepare("SELECT FlowID,BID,UserRefNo,FlowType,CreateTS,CreateBy,LastModTime,LastModBy FROM Flow WHERE ? <= CreateTS AND CreateTS < ?")
	Errcheck(err)
	RRdb.Prepstmt.GetFlow, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Flow where FlowID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetFlowsByFlowType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Flow where FlowType=?")
	Errcheck(err)
	RRdb.Prepstmt.GetFlowIDsByUser, err = RRdb.Dbrr.Prepare("SELECT DISTINCT FlowID FROM Flow where CreateBy=?")
	Errcheck(err)
	RRdb.Prepstmt.GetFlowForRAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Flow WHERE FlowType=? AND ID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertFlow, err = RRdb.Dbrr.Prepare("INSERT INTO Flow (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateFlowData, err = RRdb.Dbrr.Prepare("UPDATE Flow SET Data = JSON_REPLACE(Data, CONCAT('$.', ?), CAST(? AS JSON)) where FlowID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteFlow, err = RRdb.Dbrr.Prepare("DELETE from Flow WHERE FlowID=?")
	Errcheck(err)

	//==========================================
	// INVOICE
	//==========================================
	flds = "InvoiceNo,BID,Dt,DtDue,Amount,DeliveredBy,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Invoice"] = flds
	RRdb.Prepstmt.GetInvoice, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Invoice WHERE InvoiceNo=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllInvoicesInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Invoice WHERE BID=? AND ?>=Dt AND DtDue<=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertInvoice, err = RRdb.Dbrr.Prepare("INSERT INTO Invoice (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteInvoice, err = RRdb.Dbrr.Prepare("DELETE FROM Invoice WHERE InvoiceNo=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateInvoice, err = RRdb.Dbrr.Prepare("UPDATE Invoice SET " + s3 + " WHERE InvoiceNo=?")
	Errcheck(err)

	//==========================================
	// INVOICE PART
	//==========================================
	flds = "InvoiceNo,BID,ASMID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["InvoiceAssessment"] = flds
	RRdb.Prepstmt.GetInvoiceAssessments, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM InvoiceAssessment WHERE InvoiceNo=?")
	Errcheck(err)
	_, _, _, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertInvoiceAssessment, err = RRdb.Dbrr.Prepare("INSERT INTO InvoiceAssessment (" + s4 + ") VALUES (" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteInvoiceAssessments, err = RRdb.Dbrr.Prepare("DELETE FROM InvoiceAssessment WHERE InvoiceNo=?")
	Errcheck(err)

	//==========================================
	// INVOICE PAYOR
	//==========================================
	flds = "InvoiceNo,BID,PID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["InvoicePayor"] = flds
	RRdb.Prepstmt.GetInvoicePayors, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM InvoicePayor WHERE InvoiceNo=?")
	Errcheck(err)
	_, _, _, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertInvoicePayor, err = RRdb.Dbrr.Prepare("INSERT INTO InvoicePayor (" + s4 + ") VALUES (" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteInvoicePayors, err = RRdb.Dbrr.Prepare("DELETE FROM InvoicePayor WHERE InvoiceNo=?")
	Errcheck(err)

	//==========================================
	// JOURNAL
	//==========================================
	flds = "JID,BID,Dt,Amount,Type,ID,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Journal"] = flds
	RRdb.Prepstmt.GetJournal, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from Journal WHERE JID=?")
	Errcheck(err)
	// RRdb.Prepstmt.GetJournalInstance, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from Journal WHERE Type=0 AND Raid=0 AND ID=? AND ?<=Dt AND Dt<?")
	// Errcheck(err)
	RRdb.Prepstmt.GetJournalVacancy, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from Journal WHERE Type=0 AND ID=? AND ?<=Dt AND Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalByReceiptID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from Journal WHERE Type=2 AND ID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllJournalsInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from Journal WHERE BID=? AND ?<=Dt AND Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalByTypeAndID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from Journal WHERE Type=? AND ID=?")
	Errcheck(err)

	s1, s2, _, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertJournal, err = RRdb.Dbrr.Prepare("INSERT INTO Journal (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	RRdb.Prepstmt.DeleteJournal, err = RRdb.Dbrr.Prepare("DELETE FROM Journal WHERE JID=?")
	Errcheck(err)

	//==========================================
	// Journal Allocation
	//==========================================
	flds = "JAID,BID,JID,RID,RAID,TCID,RCPTID,Amount,ASMID,EXPID,AcctRule,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["JournalAllocation"] = flds
	RRdb.Prepstmt.GetJournalAllocation, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from JournalAllocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocations, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from JournalAllocation WHERE JID=? ORDER BY Amount DESC, RAID ASC, ASMID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocationsByASMID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from JournalAllocation WHERE ASMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalAllocationsByASMandRCPTID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from JournalAllocation WHERE ASMID>0 AND RCPTID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertJournalAllocation, err = RRdb.Dbrr.Prepare("INSERT INTO JournalAllocation (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateJournalAllocation, err = RRdb.Dbrr.Prepare("UPDATE JournalAllocation SET " + s3 + " WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalAllocation, err = RRdb.Dbrr.Prepare("DELETE FROM JournalAllocation WHERE JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalAllocations, err = RRdb.Dbrr.Prepare("DELETE FROM JournalAllocation WHERE JID=?")
	Errcheck(err)

	//==========================================
	// Journal Markers
	//==========================================
	flds = "JMID,BID,State,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["JournalMarker"] = flds
	RRdb.Prepstmt.GetJournalMarker, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from JournalMarker WHERE JMID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetJournalMarkers, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from JournalMarker ORDER BY JMID DESC LIMIT ?")
	Errcheck(err)

	s1, s2, _, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertJournalMarker, err = RRdb.Dbrr.Prepare("INSERT INTO JournalMarker (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteJournalMarker, err = RRdb.Dbrr.Prepare("DELETE FROM JournalMarker WHERE JMID=?")
	Errcheck(err)

	//==========================================
	// LEDGER-->  GLAccount
	//==========================================
	flds = "LID,PLID,BID,RAID,TCID,GLNumber,Name,AcctType,AllowPost,FLAGS,Description,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["GLAccount"] = flds
	RRdb.Prepstmt.GetLedgerByGLNo, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? AND GLNumber=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgersForGrid, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? ORDER BY GLNumber LIMIT ? OFFSET ?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? AND Name=?")
	Errcheck(err)
	// RRdb.Prepstmt.GetLedgerByType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? AND Type=?")
	// Errcheck(err)
	// RRdb.Prepstmt.GetRABalanceLedger, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? AND Type=1 AND RAID=?")
	// Errcheck(err)
	// RRdb.Prepstmt.GetSecDepBalanceLedger, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? AND Type=2 AND RAID=?")
	// Errcheck(err)
	RRdb.Prepstmt.GetLedger, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerList, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? ORDER BY GLNumber ASC, Name ASC")
	Errcheck(err)
	// RRdb.Prepstmt.GetDefaultLedgers, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM GLAccount WHERE BID=? AND Type>=10 ORDER BY GLNumber ASC")
	// Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)

	// fmt.Println("INSERT INTO GLAccount (" + s1 + ") VALUES(" + s2 + ")")
	RRdb.Prepstmt.InsertLedger, err = RRdb.Dbrr.Prepare("INSERT INTO GLAccount (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedger, err = RRdb.Dbrr.Prepare("UPDATE GLAccount SET " + s3 + " WHERE LID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedger, err = RRdb.Dbrr.Prepare("DELETE FROM GLAccount WHERE LID=?")
	Errcheck(err)

	//==========================================
	// LEDGER ENTRY
	//==========================================
	flds = "LEID,BID,JID,JAID,LID,RAID,RID,TCID,Dt,Amount,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["LedgerEntry"] = flds
	RRdb.Prepstmt.GetAllLedgerEntriesInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND ?<=Dt AND Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND LID=? AND RAID=? AND ?<=Dt AND Dt<?")
	Errcheck(err)
	RRdb.Prepstmt.CountLedgerEntries, err = RRdb.Dbrr.Prepare("SELECT COUNT(LEID) FROM LedgerEntry WHERE LID=? AND BID=?")
	Errcheck(err)
	// RRdb.Prepstmt.GetLedgerEntriesInRangeByGLNo, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND GLNo=? AND ?<=Dt AND Dt<? ORDER BY JAID ASC")
	// Errcheck(err)
	// RRdb.Prepstmt.GetLedgerEntriesInRangeByLID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND LID=? AND ?<=Dt AND Dt<? ORDER BY Amount DESC, Dt ASC")
	RRdb.Prepstmt.GetLedgerEntriesInRangeByLID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND LID=? AND ?<=Dt AND Dt<? ORDER BY Dt ASC, Amount DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntryByJAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND LID=? AND JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesByJAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from LedgerEntry WHERE BID=? AND JAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesForRAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerEntry WHERE ?<=Dt AND Dt<? AND RAID=? AND LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntriesForRentable, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerEntry WHERE ?<=Dt AND Dt<? AND RID=? AND LID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerEntriesForRAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerEntry WHERE ?<=Dt AND Dt<? AND RAID=? ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetAllLedgerEntriesForRID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerEntry WHERE ?<=Dt AND Dt<? AND RID=? ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerEntry, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerEntry where LEID=?")
	Errcheck(err)

	s1, s2, _, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertLedgerEntry, err = RRdb.Dbrr.Prepare("INSERT INTO LedgerEntry (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerEntry, err = RRdb.Dbrr.Prepare("DELETE FROM LedgerEntry WHERE LEID=?")
	Errcheck(err)

	//==========================================
	// LEDGER MARKER
	//==========================================
	flds = "LMID,LID,BID,RAID,RID,TCID,Dt,Balance,State,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["LedgerMarker"] = flds
	RRdb.Prepstmt.GetLatestLedgerMarkerByLID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE BID=? AND LID=? AND RAID=0 AND RID=0 AND TCID=0 ORDER BY Dt DESC")
	Errcheck(err)
	RRdb.Prepstmt.GetInitialLedgerMarkerByRAID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE RAID=? AND State=3")
	Errcheck(err)
	RRdb.Prepstmt.GetInitialLedgerMarkerByRID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE RAID=? AND State=3")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerByDateRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE BID=? AND LID=? AND RAID=0 AND RID=0 AND TCID=0 AND Dt>?  ORDER BY LID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkers, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE BID=? AND RAID=0 AND RID=0 AND TCID=0 ORDER BY LMID DESC LIMIT ?")
	Errcheck(err)

	RRdb.Prepstmt.GetRARentableLedgerMarkerOnOrBefore, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE RAID=? AND RID=? AND LID=0 AND TCID=0 AND Dt<=? ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetLedgerMarkerOnOrBefore, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE BID=? AND RAID=0 AND RID=0 AND LID=? AND TCID=0 AND Dt<=? ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetRALedgerMarkerOnOrBefore, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE RAID=? AND Dt<=? ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetRALedgerMarkerOnOrAfter, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE RAID=? AND Dt>=? ORDER BY Dt ASC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetTCLedgerMarkerOnOrBefore, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE TCID=? AND Dt<=? ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetTCLedgerMarkerOnOrAfter, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE TCID=? AND Dt>=? ORDER BY Dt ASC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetRALedgerMarkerOnOrBeforeDeprecated, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE BID=? AND LID=? AND RAID=? AND Dt<=?  ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableLedgerMarkerOnOrBefore, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM LedgerMarker WHERE BID=? AND LID=? AND RID=? and Dt<=?  ORDER BY Dt DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.DeleteLedgerMarker, err = RRdb.Dbrr.Prepare("DELETE FROM LedgerMarker WHERE LMID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertLedgerMarker, err = RRdb.Dbrr.Prepare("INSERT INTO LedgerMarker (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateLedgerMarker, err = RRdb.Dbrr.Prepare("UPDATE LedgerMarker SET " + s3 + " WHERE LMID=?")
	Errcheck(err)

	//==========================================
	// NOTES
	//==========================================
	flds = "NID,BID,NLID,PNID,NTID,RID,RAID,TCID,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Notes"] = flds
	RRdb.Prepstmt.GetNote, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Notes WHERE NID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetNoteAndChildNotes, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Notes WHERE PNID=? ORDER BY LastModTime ASC")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertNote, err = RRdb.Dbrr.Prepare("INSERT INTO Notes (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateNote, err = RRdb.Dbrr.Prepare("UPDATE Notes SET " + s3 + " WHERE NID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteNote, err = RRdb.Dbrr.Prepare("DELETE FROM Notes WHERE NID=?")
	Errcheck(err)

	//==========================================
	// NOTELIST
	//==========================================
	flds = "NLID,BID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["NoteList"] = flds
	RRdb.Prepstmt.GetNoteList, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM NoteList WHERE NLID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetNoteListMembers, err = RRdb.Dbrr.Prepare("SELECT NID FROM Notes WHERE NLID=? and PNID=0")
	Errcheck(err)
	s1, s2, _, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertNoteList, err = RRdb.Dbrr.Prepare("INSERT INTO NoteList (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteNoteList, err = RRdb.Dbrr.Prepare("DELETE FROM NoteList WHERE NLID=?")
	Errcheck(err)

	//==========================================
	// NOTETYPE
	//==========================================
	flds = "NTID,BID,Name,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["NoteType"] = flds
	RRdb.Prepstmt.GetNoteType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM NoteType WHERE NTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllNoteTypes, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM NoteType WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertNoteType, err = RRdb.Dbrr.Prepare("INSERT INTO NoteType (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateNoteType, err = RRdb.Dbrr.Prepare("UPDATE NoteType SET " + s3 + " WHERE NTID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteNoteType, err = RRdb.Dbrr.Prepare("DELETE FROM NoteType WHERE NTID=?")
	Errcheck(err)

	//==========================================
	// PAYMENT TYPES
	//==========================================
	flds = "PMTID,BID,Name,Description,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["PaymentType"] = flds
	RRdb.Prepstmt.GetPaymentType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM PaymentType WHERE PMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetPaymentTypeByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM PaymentType WHERE BID=? AND Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetPaymentTypesByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM PaymentType WHERE BID=? ORDER BY Name ASC")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertPaymentType, err = RRdb.Dbrr.Prepare("INSERT INTO PaymentType (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdatePaymentType, err = RRdb.Dbrr.Prepare("UPDATE PaymentType SET " + s3 + " WHERE PMTID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeletePaymentType, err = RRdb.Dbrr.Prepare("DELETE FROM PaymentType WHERE PMTID=?")
	Errcheck(err)

	//==========================================
	// PAYOR
	//==========================================
	flds = "TCID,BID,CreditLimit,TaxpayorID,ThirdPartySource,EligibleFuturePayor,FLAGS,SSN,DriversLicense,GrossIncome,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Payor"] = flds
	RRdb.Prepstmt.GetPayor, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Payor where TCID=?")
	Errcheck(err)
	_, _, s3, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertPayor, err = RRdb.Dbrr.Prepare("INSERT INTO Payor (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdatePayor, err = RRdb.Dbrr.Prepare("UPDATE Payor SET " + s3 + " WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeletePayor, err = RRdb.Dbrr.Prepare("DELETE from Payor WHERE TCID=?")
	Errcheck(err)

	//==========================================
	// PROSPECT
	//==========================================
	flds = "TCID,BID,CompanyAddress,CompanyCity,CompanyState,CompanyPostalCode,CompanyEmail,CompanyPhone,Occupation,DesiredUsageStartDate,RentableTypePreference,FLAGS,EvictedDes,ConvictedDes,BankruptcyDes,Approver,DeclineReasonSLSID,OtherPreferences,FollowUpDate,CSAgent,OutcomeSLSID,CurrentAddress,CurrentLandLordName,CurrentLandLordPhoneNo,CurrentReasonForMoving,CurrentLengthOfResidency,PriorAddress,PriorLandLordName,PriorLandLordPhoneNo,PriorReasonForMoving,PriorLengthOfResidency,CommissionableThirdParty,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Prospect"] = flds
	RRdb.Prepstmt.GetProspect, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Prospect where TCID=?")
	Errcheck(err)
	_, _, s3, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertProspect, err = RRdb.Dbrr.Prepare("INSERT INTO Prospect (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateProspect, err = RRdb.Dbrr.Prepare("UPDATE Prospect SET " + s3 + " WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteProspect, err = RRdb.Dbrr.Prepare("DELETE from Prospect WHERE TCID=?")
	Errcheck(err)

	//==========================================
	// User
	//==========================================
	flds = "TCID,BID,Points,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyContactEmail,AlternateAddress,EligibleFutureUser,FLAGS,Industry,SourceSLSID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["User"] = flds
	RRdb.Prepstmt.GetUser, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM User where TCID=?")
	Errcheck(err)

	_, _, s3, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertUser, err = RRdb.Dbrr.Prepare("INSERT INTO User (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateUser, err = RRdb.Dbrr.Prepare("UPDATE User SET " + s3 + " WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteUser, err = RRdb.Dbrr.Prepare("DELETE from User WHERE TCID=?")
	Errcheck(err)

	//==========================================
	// RATE PLAN
	//==========================================
	flds = "RPID,BID,Name,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RatePlan"] = flds
	RRdb.Prepstmt.GetRatePlan, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlan WHERE RPID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRatePlanByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlan WHERE BID=? AND Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRatePlans, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlan WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRatePlan, err = RRdb.Dbrr.Prepare("INSERT INTO RatePlan (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRatePlan, err = RRdb.Dbrr.Prepare("UPDATE RatePlan SET " + s3 + " WHERE RPID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRatePlan, err = RRdb.Dbrr.Prepare("DELETE FROM RatePlan WHERE RPID=?")
	Errcheck(err)

	//==========================================
	// RATE PLAN REF
	//==========================================
	flds = "RPRID,BID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RatePlanRef"] = flds
	RRdb.Prepstmt.GetRatePlanRef, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlanRef WHERE RPRID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRatePlanRefsInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RatePlanRef WHERE RPID=? and ?>=DtStart and ?<DtStop")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRatePlanRefsInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RatePlanRef WHERE ?>=DtStart and ?<DtStop")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRatePlanRef, err = RRdb.Dbrr.Prepare("INSERT INTO RatePlanRef (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRatePlanRef, err = RRdb.Dbrr.Prepare("UPDATE RatePlanRef SET " + s3 + " WHERE RPRID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRatePlanRef, err = RRdb.Dbrr.Prepare("DELETE FROM RatePlanRef WHERE RPRID=?")
	Errcheck(err)

	//==========================================
	// RATE PLAN REF RPRID's Rate Info for Rentable Type
	//==========================================
	flds = "RPRID,BID,RTID,FLAGS,Val,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RatePlanRefRTRate"] = flds
	RRdb.Prepstmt.GetRatePlanRefRTRate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlanRefRTRate WHERE RPRID=? and RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRatePlanRefRTRates, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlanRefRTRate WHERE RPRID=?")
	Errcheck(err)

	_, _, _, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRatePlanRefRTRate, err = RRdb.Dbrr.Prepare("INSERT INTO RatePlanRefRTRate (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRatePlanRefRTRate, err = RRdb.Dbrr.Prepare("UPDATE  RatePlanRefRTRate SET BID=?,FLAGS=?,Val=?,LastModBy=? WHERE RPRID=? and RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRatePlanRefRTRate, err = RRdb.Dbrr.Prepare("DELETE FROM  RatePlanRefRTRate WHERE RPRID=? and RTID=?")
	Errcheck(err)

	//==========================================
	// RATE PLAN Ref RPRID's Rate Info for Specialties
	//==========================================
	flds = "RPRID,BID,RTID,RSPID,FLAGS,Val,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RatePlanRefSPRate"] = flds
	RRdb.Prepstmt.GetRatePlanRefSPRate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlanRefSPRate WHERE RPRID=? and RTID=? and RSPID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRatePlanRefSPRates, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RatePlanRefSPRate WHERE RPRID=? and RTID=?")
	Errcheck(err)

	_, _, _, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRatePlanRefSPRate, err = RRdb.Dbrr.Prepare("INSERT INTO RatePlanRefSPRate (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRatePlanRefSPRate, err = RRdb.Dbrr.Prepare("UPDATE RatePlanRefSPRate SET FLAGS=?,Val=?,LastModBy=? WHERE RPRID=? and RTID=? and RSPID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRatePlanRefSPRate, err = RRdb.Dbrr.Prepare("DELETE FROM RatePlanRefSPRate WHERE RPRID=? and RSPID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT
	//==========================================
	flds = "RCPTID,PRCPTID,BID,TCID,PMTID,DEPID,DID,RAID,Dt,DocNo,Amount,AcctRuleReceive,ARID,AcctRuleApply,FLAGS,Comment,OtherPayorName,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Receipt"] = flds
	RRdb.Prepstmt.GetReceipt, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Receipt WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptDuplicate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Receipt WHERE Dt=? AND Amount=? AND DocNo=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptsInDateRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Receipt WHERE BID=? AND Dt >= ? AND Dt < ?")
	Errcheck(err)

	//  FLAGS bits 0-1:  0 unallocated, 1 = partially allocated, 2 = fully allocated,  bit 2: voided entry
	RRdb.Prepstmt.GetUnallocatedReceipts, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Receipt WHERE BID=? AND (FLAGS & 7)<2 ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetUnallocatedReceiptsByPayor, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Receipt WHERE BID=? AND TCID=? AND (FLAGS & 3)<2 AND 0=(FLAGS & 4) ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetPayorUnallocatedReceiptsCount, err = RRdb.Dbrr.Prepare("SELECT COUNT(*) FROM Receipt WHERE BID=? AND TCID=? AND (FLAGS & 3)<2 AND 0=(FLAGS & 4)")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertReceipt, err = RRdb.Dbrr.Prepare("INSERT INTO Receipt (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceipt, err = RRdb.Dbrr.Prepare("DELETE FROM Receipt WHERE RCPTID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateReceipt, err = RRdb.Dbrr.Prepare("UPDATE Receipt SET " + s3 + " WHERE RCPTID=?")
	Errcheck(err)

	//==========================================
	// RECEIPT ALLOCATION
	//==========================================
	flds = "RCPAID,RCPTID,BID,RAID,Dt,Amount,ASMID,FLAGS,AcctRule,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["ReceiptAllocation"] = flds
	RRdb.Prepstmt.GetReceiptAllocation, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ReceiptAllocation WHERE RCPAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptAllocations, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ReceiptAllocation WHERE RCPTID=? ORDER BY Amount DESC, RAID ASC, ASMID ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptAllocationsThroughDate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ReceiptAllocation WHERE ASMID>0 AND RCPTID=? AND Dt <= ? ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetASMReceiptAllocationsInRAIDDateRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ReceiptAllocation WHERE ASMID>0 AND RAID=? AND Dt >= ? and Dt < ? ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetASMReceiptAllocationsInRARDateRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ReceiptAllocation WHERE ASMID>0 AND RAID=? AND ASMID=? AND ?<=Dt AND Dt<? ORDER BY Dt ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetReceiptAllocationsByASMID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM ReceiptAllocation WHERE BID=? AND ASMID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertReceiptAllocation, err = RRdb.Dbrr.Prepare("INSERT INTO ReceiptAllocation (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceiptAllocation, err = RRdb.Dbrr.Prepare("DELETE FROM ReceiptAllocation WHERE RCPAID=?") // delete just this allocation
	Errcheck(err)
	RRdb.Prepstmt.DeleteReceiptAllocations, err = RRdb.Dbrr.Prepare("DELETE FROM ReceiptAllocation WHERE RCPTID=?") // delete all associated with the receipt
	Errcheck(err)
	RRdb.Prepstmt.UpdateReceiptAllocation, err = RRdb.Dbrr.Prepare("UPDATE ReceiptAllocation SET " + s3 + " WHERE RCPAID=?")
	Errcheck(err)

	//===============================
	//  Rentable
	//===============================
	flds = "RID,BID,RentableName,AssignmentTime,MRStatus,DtMRStart,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Rentable"] = flds
	RRdb.Prepstmt.CountBusinessRentables, err = RRdb.Dbrr.Prepare("SELECT COUNT(RID) FROM Rentable WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeDown, err = RRdb.Dbrr.Prepare("SELECT RID,RentableName FROM Rentable WHERE BID=? AND (RentableName LIKE ?) LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentable, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Rentable WHERE RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Rentable WHERE RentableName=? AND BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentablesByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Rentable WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentable, err = RRdb.Dbrr.Prepare("INSERT INTO Rentable (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentable, err = RRdb.Dbrr.Prepare("UPDATE Rentable SET " + s3 + " WHERE RID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement
	//===============================
	flds = "RAID,RATID,BID,NLID,AgreementStart,AgreementStop,PossessionStart,PossessionStop,RentStart,RentStop,RentCycleEpoch,UnspecifiedAdults,UnspecifiedChildren,Renewal,SpecialProvisions,LeaseType,ExpenseAdjustmentType,ExpensesStop,ExpenseStopCalculation,BaseYearEnd,ExpenseAdjustment,EstimatedCharges,RateChange,NextRateChange,PermittedUses,ExclusiveUses,ExtensionOption,ExtensionOptionNotice,ExpansionOption,ExpansionOptionNotice,RightOfFirstRefusal,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentalAgreement"] = flds
	RRdb.Prepstmt.CountBusinessRentalAgreements, err = RRdb.Dbrr.Prepare("SELECT COUNT(RAID) FROM RentalAgreement WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementByBusiness, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreement where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreement, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreement WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreementsByRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreement WHERE BID=? AND ?<=AgreementStop AND ?>AgreementStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreements, err = RRdb.Dbrr.Prepare("SELECT RAID from RentalAgreement WHERE BID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentalAgreement, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreement (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentalAgreement, err = RRdb.Dbrr.Prepare("UPDATE RentalAgreement SET " + s3 + " WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentalAgreement, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreement WHERE RAID=?")
	Errcheck(err)

	RRdb.Prepstmt.GetRentalAgreementTypeDown, err = RRdb.Dbrr.Prepare("SELECT Transactant.TCID,Transactant.FirstName,Transactant.MiddleName,Transactant.LastName,Transactant.CompanyName,Transactant.IsCompany,RentalAgreementPayors.RAID FROM Transactant LEFT JOIN RentalAgreementPayors ON RentalAgreementPayors.TCID=Transactant.TCID WHERE Transactant.BID=? AND RentalAgreementPayors.RAID>0 AND (Transactant.FirstName LIKE ? OR Transactant.LastName LIKE ? OR Transactant.CompanyName LIKE ?) GROUP BY RentalAgreementPayors.RAID ORDER BY RentalAgreementPayors.DtStart DESC, RentalAgreementPayors.RAPID ASC LIMIT ?")
	Errcheck(err)

	//====================================================
	//  Rental Agreement Rentable
	//====================================================
	flds = "RARID,RAID,BID,RID,CLID,ContractRent,RARDtStart,RARDtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentalAgreementRentables"] = flds
	RRdb.Prepstmt.GetRARentableForDate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreementRentables WHERE RAID=? AND ?>=RARDtStart AND ?<RARDtStop")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementRentables, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreementRentables WHERE RAID=? and ?<RARDtStop and ?>=RARDtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreementRentables, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreementRentables WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementsForRentable, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreementRentables WHERE RID=? and ?<RARDtStop and ?>RARDtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetAgreementsForRentable, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreementRentables WHERE RID=? and ?<RARDtStop and ?>RARDtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementRentable, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentalAgreementRentables WHERE RARID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentalAgreementRentable, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementRentables (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentalAgreementRentable, err = RRdb.Dbrr.Prepare("UPDATE RentalAgreementRentables SET " + s3 + " WHERE RARID=?")
	Errcheck(err)
	RRdb.Prepstmt.FindAgreementByRentable, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementRentables WHERE RID=? AND RARDtStop>? AND RARDtStart<=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentalAgreementRentable, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementRentables WHERE RARID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAllRentalAgreementRentables, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementRentables WHERE RAID=?")
	Errcheck(err)

	//====================================================
	//  Rental Agreement Users
	//====================================================
	flds = "RUID,RID,BID,TCID,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableUsers"] = flds
	RRdb.Prepstmt.GetRentableUser, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentableUsers WHERE RUID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableUsersInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentableUsers WHERE RID=? and ?<DtStop and ?>=DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableUserByRBT, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentableUsers WHERE RID=? AND BID=? AND TCID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.UpdateRentableUser, err = RRdb.Dbrr.Prepare("UPDATE RentableUsers SET " + s3 + " WHERE RUID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableUserByRBT, err = RRdb.Dbrr.Prepare("UPDATE RentableUsers SET DtStart=?,DtStop=?,LastModBy=? WHERE RID=? AND BID=? AND TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertRentableUser, err = RRdb.Dbrr.Prepare("INSERT INTO RentableUsers (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableUserByRBT, err = RRdb.Dbrr.Prepare("DELETE from RentableUsers WHERE RID=? AND BID=? AND TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableUser, err = RRdb.Dbrr.Prepare("DELETE FROM RentableUsers WHERE RUID=?")
	Errcheck(err)

	//====================================================
	//  Rental Agreement Payors
	//====================================================
	flds = "RAPID,RAID,BID,TCID,DtStart,DtStop,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentalAgreementPayors"] = flds
	RRdb.Prepstmt.GetRentalAgreementPayor, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPayors WHERE RAPID=?")
	Errcheck(err)
	// RRdb.Prepstmt.GetRentalAgreementPayorsInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPayors WHERE RAID=? AND ((?<DtStop AND ?>DtStart) OR (DtStop=DtStart AND (?=DtStart || ?=DtStop)))")
	RRdb.Prepstmt.GetRentalAgreementPayorsInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPayors WHERE RAID=? AND ?<DtStop AND ?>DtStart")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementPayorByRBT, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPayors WHERE RAID=? AND BID=? AND TCID=?")
	Errcheck(err)
	// RRdb.Prepstmt.GetRentalAgreementsByPayor, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPayors WHERE BID=? AND TCID=? AND DtStart<=? AND ?<DtStop")
	RRdb.Prepstmt.GetRentalAgreementsByPayor, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPayors WHERE BID=? AND TCID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentalAgreementPayor, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementPayors (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentalAgreementPayor, err = RRdb.Dbrr.Prepare("UPDATE RentalAgreementPayors SET " + s3 + " WHERE RAPID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentalAgreementPayorByRBT, err = RRdb.Dbrr.Prepare("UPDATE RentalAgreementPayors SET DtStart=?,DtStop=?,FLAGS=?,LastModBy=? WHERE RAID=? AND BID=? AND TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentalAgreementPayorByRBT, err = RRdb.Dbrr.Prepare("DELETE from RentalAgreementPayors WHERE RAID=? AND BID=? AND TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentalAgreementPayor, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementPayors WHERE RAPID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAllRentalAgreementPayors, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementPayors WHERE RAID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement Pets
	//===============================
	flds = "PETID,BID,RAID,TCID,Type,Breed,Color,Weight,Name,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentalAgreementPets"] = flds
	RRdb.Prepstmt.GetRentalAgreementPet, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentalAgreementPets, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPets WHERE RAID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetPetsByTransactant, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementPets WHERE TCID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentalAgreementPet, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementPets (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentalAgreementPet, err = RRdb.Dbrr.Prepare("UPDATE RentalAgreementPets SET " + s3 + " WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentalAgreementPet, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementPets WHERE PETID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteAllRentalAgreementPets, err = RRdb.Dbrr.Prepare("DELETE FROM RentalAgreementPets WHERE RAID=?")
	Errcheck(err)

	//===============================
	//  Rental Agreement Template
	//===============================
	flds = "RATID,BID,RATemplateName,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentalAgreementTemplate"] = flds
	RRdb.Prepstmt.GetAllRentalAgreementTemplates, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementTemplate")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementTemplate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementTemplate WHERE RATID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentalAgreementByRATemplateName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentalAgreementTemplate WHERE RATemplateName=?")
	Errcheck(err)
	s1, s2, _, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentalAgreementTemplate, err = RRdb.Dbrr.Prepare("INSERT INTO RentalAgreementTemplate (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//===============================
	//  RentableTypeRef
	//===============================
	flds = "RTRID,RID,BID,RTID,OverrideRentCycle,OverrideProrationCycle,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableTypeRef"] = flds
	RRdb.Prepstmt.GetRentableTypeRef, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypeRef WHERE RTRID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeRefsByRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypeRef WHERE RID=? AND DtStop>? AND DtStart<? ORDER BY DtStart ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeRefs, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypeRef WHERE RID=? ORDER BY DtStart ASC")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentableTypeRef, err = RRdb.Dbrr.Prepare("INSERT INTO RentableTypeRef (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableTypeRef, err = RRdb.Dbrr.Prepare("UPDATE RentableTypeRef SET " + s3 + " WHERE RTRID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableTypeRef, err = RRdb.Dbrr.Prepare("DELETE from RentableTypeRef WHERE RTRID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableTypeRefWithRTID, err = RRdb.Dbrr.Prepare("DELETE from RentableTypeRef WHERE RTID=?")
	Errcheck(err)

	//===============================
	//  RentableSpecialtyRef
	//===============================
	flds = "BID,RID,RSPID,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableSpecialtyRef"] = flds
	RRdb.Prepstmt.GetRentableSpecialtyRefs, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableSpecialtyRef WHERE BID=? AND RID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialtyRefsByRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableSpecialtyRef WHERE BID=? AND RID=? AND DtStop>? AND DtStart<? ORDER BY DtStart ASC")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableSpecialtyRefs, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableSpecialtyRef WHERE BID=? ORDER BY DtStart ASC")
	Errcheck(err)

	_, _, s3, s4, s5 = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentableSpecialtyRef, err = RRdb.Dbrr.Prepare("INSERT INTO RentableSpecialtyRef (" + s4 + ") VALUES(" + s5 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableSpecialtyRef, err = RRdb.Dbrr.Prepare("UPDATE RentableSpecialtyRef SET " + s3 + " WHERE RID=? AND DtStart=? AND DtStop=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableSpecialtyRef, err = RRdb.Dbrr.Prepare("DELETE from RentableSpecialtyRef WHERE RID=? AND DtStart=? AND DtStop=?")
	Errcheck(err)

	//===============================
	//  RentableSpecialty
	//===============================
	flds = "RSPID,BID,Name,Fee,Description,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableSpecialty"] = flds
	RRdb.Prepstmt.GetRentableSpecialtyTypeByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableSpecialty WHERE BID=? AND Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableSpecialtyType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableSpecialty WHERE RSPID=?")
	Errcheck(err)
	s1, s2, _, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentableSpecialtyType, err = RRdb.Dbrr.Prepare("INSERT INTO RentableSpecialty (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)

	//===============================
	//  RentableStatus
	//===============================
	flds = "RSID,RID,BID,DtStart,DtStop,DtNoticeToVacate,UseStatus,LeaseStatus,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableStatus"] = flds
	RRdb.Prepstmt.GetRentableStatus, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableStatus WHERE RSID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableStatusByRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableStatus WHERE RID=? AND DtStop>? AND DtStart<?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableStatusOnOrAfter, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableStatus WHERE RID=? AND DtStart>=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllRentableStatus, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableStatus WHERE RID=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentableStatus, err = RRdb.Dbrr.Prepare("INSERT INTO RentableStatus (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableStatus, err = RRdb.Dbrr.Prepare("UPDATE RentableStatus SET " + s3 + " WHERE RSID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableStatus, err = RRdb.Dbrr.Prepare("DELETE from RentableStatus WHERE RSID=?")
	Errcheck(err)

	//===============================
	//  Rentable Type
	//===============================
	flds = "RTID,BID,Style,Name,RentCycle,Proration,GSRPC,ARID,FLAGS,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableTypes"] = flds
	RRdb.Prepstmt.CountBusinessRentableTypes, err = RRdb.Dbrr.Prepare("SELECT COUNT(RTID) FROM RentableTypes WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableType, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypes WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByStyle, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypes WHERE Style=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableTypeByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypes WHERE Name=? and BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllBusinessRentableTypes, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM RentableTypes WHERE BID=? ORDER BY RTID ASC")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentableType, err = RRdb.Dbrr.Prepare("INSERT INTO RentableTypes (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableType, err = RRdb.Dbrr.Prepare("UPDATE RentableTypes SET " + s3 + " WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableTypeToActive, err = RRdb.Dbrr.Prepare("UPDATE RentableTypes SET FLAGS=FLAGS&(~(1<<0)),LastModBy=? WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableTypeToInactive, err = RRdb.Dbrr.Prepare("UPDATE RentableTypes SET FLAGS=FLAGS|(1<<0),LastModBy=? WHERE RTID=?")
	Errcheck(err)

	//===============================
	//  RentableMarketRates
	//===============================
	flds = "RMRID,RTID,BID,MarketRate,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["RentableMarketRate"] = flds
	RRdb.Prepstmt.GetRentableMarketRates, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentableMarketRate WHERE RTID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetRentableMarketRateInstance, err = RRdb.Dbrr.Prepare("SELECT " + flds + " from RentableMarketRate WHERE RMRID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertRentableMarketRates, err = RRdb.Dbrr.Prepare("INSERT INTO RentableMarketRate (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateRentableMarketRateInstance, err = RRdb.Dbrr.Prepare("UPDATE RentableMarketRate SET " + s3 + " WHERE RMRID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteRentableMarketRateInstance, err = RRdb.Dbrr.Prepare("DELETE from RentableMarketRate WHERE RMRID=?")
	Errcheck(err)

	//==========================================
	// SOURCE
	//==========================================
	flds = "SourceSLSID,BID,Name,Industry,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["DemandSource"] = flds
	RRdb.Prepstmt.GetDemandSource, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DemandSource WHERE SourceSLSID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetDemandSourceByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DemandSource WHERE BID=? and Name=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllDemandSources, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM DemandSource WHERE BID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertDemandSource, err = RRdb.Dbrr.Prepare("INSERT INTO DemandSource (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateDemandSource, err = RRdb.Dbrr.Prepare("UPDATE DemandSource SET " + s3 + " WHERE SourceSLSID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteDemandSource, err = RRdb.Dbrr.Prepare("DELETE from DemandSource WHERE SourceSLSID=?")
	Errcheck(err)

	//==========================================
	// STRING LIST
	//==========================================
	flds = "SLID,BID,Name,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["StringList"] = flds
	RRdb.Prepstmt.GetStringList, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM StringList WHERE SLID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllStringLists, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM StringList WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetStringListByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM StringList WHERE BID=? AND Name=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertStringList, err = RRdb.Dbrr.Prepare("INSERT INTO StringList (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateStringList, err = RRdb.Dbrr.Prepare("UPDATE StringList SET " + s3 + " WHERE SLID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteStringList, err = RRdb.Dbrr.Prepare("DELETE from StringList WHERE SLID=?")
	Errcheck(err)

	//==========================================
	// SLString
	//==========================================
	flds = "SLSID,BID,SLID,Value,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["SLString"] = flds
	RRdb.Prepstmt.GetSLString, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM SLString WHERE SLSID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetSLStrings, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM SLString WHERE SLID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertSLString, err = RRdb.Dbrr.Prepare("INSERT INTO SLString (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateSLString, err = RRdb.Dbrr.Prepare("UPDATE SLString SET " + s3 + " WHERE SLSID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteSLString, err = RRdb.Dbrr.Prepare("DELETE from SLString WHERE SLSID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteSLStrings, err = RRdb.Dbrr.Prepare("DELETE from SLString WHERE SLID=?")
	Errcheck(err)

	//==========================================
	// SubAR
	//==========================================
	flds = "SARID,ARID,SubARID,BID,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["SubAR"] = flds
	RRdb.Prepstmt.GetSubAR, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM SubAR WHERE SARID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetSubARs, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM SubAR WHERE ARID=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertSubAR, err = RRdb.Dbrr.Prepare("INSERT INTO SubAR (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateSubAR, err = RRdb.Dbrr.Prepare("UPDATE SubAR SET " + s3 + " WHERE SARID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteSubAR, err = RRdb.Dbrr.Prepare("DELETE from SubAR WHERE SARID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteSubARs, err = RRdb.Dbrr.Prepare("DELETE from SubAR WHERE ARID=?")
	Errcheck(err)

	//==========================================
	// TASK
	//==========================================
	flds = "TID,BID,TLID,Name,Worker,DtDue,DtPreDue,DtDone,DtPreDone,FLAGS,DoneUID,PreDoneUID,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Task"] = flds
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.GetTask, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Task WHERE TID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTasks, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Task WHERE TLID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertTask, err = RRdb.Dbrr.Prepare("INSERT INTO Task (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTask, err = RRdb.Dbrr.Prepare("UPDATE Task SET " + s3 + " WHERE TID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteTask, err = RRdb.Dbrr.Prepare("DELETE from Task WHERE TID=?")
	Errcheck(err)

	//==========================================
	// TASKLIST
	//==========================================
	//      1    2   3    4     5     6        7      8          9    10      11         12        13           14             15       16       17,         18
	flds = "TLID,BID,PTLID,TLDID,Name,Cycle,DtDue,DtPreDue,DtDone,DtPreDone,FLAGS,DoneUID,PreDoneUID,EmailList,DtLastNotify,DurWait,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["TaskList"] = flds
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.GetTaskList, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskList WHERE TLID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllParentTaskLists, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskList WHERE PTLID=0")
	Errcheck(err)
	RRdb.Prepstmt.GetLatestCompletedTaskList, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskList WHERE (FLAGS & 16 > 0) AND ((PTLID = 0 AND TLID = ?) OR PTLID=?) ORDER BY DtDone DESC LIMIT 1")
	Errcheck(err)
	RRdb.Prepstmt.GetTaskListInstanceInRange, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskList WHERE (PTLID=? OR (PTLID=0 AND TLID=?)) AND DtDue >= ? AND DtDue < ?")
	Errcheck(err)

	where := `WHERE
    -- the TaskList is enabled
    (FLAGS & 1 = 0)
    AND
	(
		(
			-- no notifications have been made
			(FLAGS & 32 = 0)
			OR
			(
				-- notification has been made
				(FLAGS & 32 > 0)
				AND
				-- wait period after last notify has passed
				(DATE_ADD(DtLastNotify, interval DurWait/1000 microsecond) < ?)
			)
		)
		AND
		(
			-- PreDone check needed  No Due Date         due rqd                Done not set    DueDate passed   DueDate not passed   PreDone not set    PreDueDate has passed
			((FLAGS & 2) > 0  AND  ((FLAGS & 4) = 0 OR ((FLAGS & 4) > 0 AND ( ((FLAGS & 16 = 0) AND ? > DtDue) OR ? < DtDue) ) ) AND ((FLAGS & 8 = 0) AND ? > DtPreDue))
			OR
			--  Done check needed  Done is not set AND Due date has passed
			((FLAGS & 4) > 0  AND  (FLAGS & 16 = 0) AND ? > DtDue)
		)
	);`
	RRdb.Prepstmt.GetDueTaskLists, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskList " + where)
	Errcheck(err)

	RRdb.Prepstmt.CheckForTLDInstances, err = RRdb.Dbrr.Prepare("SELECT COUNT(*) FROM TaskList WHERE TLDID=?")
	Errcheck(err)

	RRdb.Prepstmt.CheckForTLDInstances, err = RRdb.Dbrr.Prepare("SELECT COUNT(*) FROM TaskList WHERE TLDID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertTaskList, err = RRdb.Dbrr.Prepare("INSERT INTO TaskList (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTaskList, err = RRdb.Dbrr.Prepare("UPDATE TaskList SET " + s3 + " WHERE TLID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteTaskList, err = RRdb.Dbrr.Prepare("DELETE from TaskList WHERE TLID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteTaskListTasks, err = RRdb.Dbrr.Prepare("DELETE from Task WHERE TLID=?")
	Errcheck(err)

	//==========================================
	// TASKDESCRIPTOR
	//==========================================
	flds = "TDID,BID,TLDID,Name,Worker,EpochDue,EpochPreDue,FLAGS,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["TaskDescriptor"] = flds
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.GetTaskDescriptor, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskDescriptor WHERE TDID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTaskListDescriptors, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskDescriptor WHERE TLDID=?")
	Errcheck(err)
	RRdb.Prepstmt.InsertTaskDescriptor, err = RRdb.Dbrr.Prepare("INSERT INTO TaskDescriptor (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTaskDescriptor, err = RRdb.Dbrr.Prepare("UPDATE TaskDescriptor SET " + s3 + " WHERE TDID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteTaskDescriptor, err = RRdb.Dbrr.Prepare("DELETE from TaskDescriptor WHERE TDID=?")
	Errcheck(err)

	//==========================================
	// TASK LIST DEFINITION
	//==========================================
	flds = "TLDID,BID,Name,Cycle,Epoch,EpochDue,EpochPreDue,FLAGS,EmailList,DurWait,Comment,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["TaskListDefinition"] = flds
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.GetAllTaskListDefinitions, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskListDefinition WHERE BID=? AND FLAGS & 1 = 0")
	Errcheck(err)
	RRdb.Prepstmt.GetTaskListDefinition, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskListDefinition WHERE TLDID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTaskListDefinitionByName, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM TaskListDefinition WHERE BID=? AND Name=?")
	Errcheck(err)
	// qry := "INSERT INTO TaskListDefinition (" + s1 + ") VALUES(" + s2 + ")"
	// Console("qry = %s\n", qry)
	RRdb.Prepstmt.InsertTaskListDefinition, err = RRdb.Dbrr.Prepare("INSERT INTO TaskListDefinition (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTaskListDefinition, err = RRdb.Dbrr.Prepare("UPDATE TaskListDefinition SET " + s3 + " WHERE TLDID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteTaskListDefinition, err = RRdb.Dbrr.Prepare("DELETE from TaskListDefinition WHERE TLDID=?")
	Errcheck(err)

	//==========================================
	// TRANSACTANT
	//==========================================
	RRdb.DBFields["Transactant"] = TRNSfields
	RRdb.Prepstmt.GetTransactantTypeDown, err = RRdb.Dbrr.Prepare("SELECT TCID,FirstName,MiddleName,LastName,CompanyName,IsCompany FROM Transactant WHERE BID=? AND (FirstName LIKE ? OR MiddleName LIKE ? OR LastName LIKE ? OR CompanyName LIKE ?) LIMIT ?")
	Errcheck(err)
	RRdb.Prepstmt.CountBusinessTransactants, err = RRdb.Dbrr.Prepare("SELECT COUNT(TCID) FROM Transactant WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetTransactant, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetAllTransactants, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant")
	Errcheck(err)
	RRdb.Prepstmt.GetAllTransactantsForBID, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant WHERE BID=?")
	Errcheck(err)
	RRdb.Prepstmt.FindTransactantByPhoneOrEmail, err = RRdb.Dbrr.Prepare("SELECT " + TRNSfields + " FROM Transactant where WorkPhone=? OR CellPhone=? or PrimaryEmail=? or SecondaryEmail=?")
	Errcheck(err)
	RRdb.Prepstmt.FindTCIDByNote, err = RRdb.Dbrr.Prepare("SELECT t.TCID FROM Transactant t, Notes n WHERE t.NLID = n.NLID AND n.Comment=?")
	Errcheck(err)

	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(TRNSfields)
	RRdb.Prepstmt.InsertTransactant, err = RRdb.Dbrr.Prepare("INSERT INTO Transactant (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateTransactant, err = RRdb.Dbrr.Prepare("UPDATE Transactant SET " + s3 + " WHERE TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteTransactant, err = RRdb.Dbrr.Prepare("DELETE from Transactant WHERE TCID=?")
	Errcheck(err)

	//==========================================
	// UIGrid
	//==========================================
	RRdb.Prepstmt.UIRAGrid, err = RRdb.Dbrr.Prepare("SELECT ra.RAID,rap.TCID,ra.AgreementStart,ra.AgreementStop FROM RentalAgreement ra INNER JOIN RentalAgreementPayors rap ON ra.RAID=rap.RAID AND rap.DtStart<=? AND ?<rap.DtStop WHERE ra.BID=?")
	Errcheck(err)

	//==========================================
	// Vehicle
	//==========================================
	flds = "VID,TCID,BID,VehicleType,VehicleMake,VehicleModel,VehicleColor,VehicleYear,VIN,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DtStart,DtStop,CreateTS,CreateBy,LastModTime,LastModBy"
	RRdb.DBFields["Vehicle"] = flds
	RRdb.Prepstmt.GetVehicle, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Vehicle where VID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetVehiclesByTransactant, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Vehicle where TCID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetVehiclesByBID, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Vehicle where BID=?")
	Errcheck(err)
	RRdb.Prepstmt.GetVehiclesByLicensePlate, err = RRdb.Dbrr.Prepare("SELECT " + flds + " FROM Vehicle where LicensePlateNumber=?")
	Errcheck(err)
	s1, s2, s3, _, _ = GenSQLInsertAndUpdateStrings(flds)
	RRdb.Prepstmt.InsertVehicle, err = RRdb.Dbrr.Prepare("INSERT INTO Vehicle (" + s1 + ") VALUES(" + s2 + ")")
	Errcheck(err)
	RRdb.Prepstmt.UpdateVehicle, err = RRdb.Dbrr.Prepare("UPDATE Vehicle SET " + s3 + " WHERE VID=?")
	Errcheck(err)
	RRdb.Prepstmt.DeleteVehicle, err = RRdb.Dbrr.Prepare("DELETE from Vehicle WHERE VID=?")
	Errcheck(err)

}
