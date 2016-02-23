package main

import "rentroll/rlib"

func buildPreparedStatements() {
	var err error
	// Prepare("select deduction from deductions where uid=?")
	// Prepare("select type from compensation where uid=?")
	// Prepare("INSERT INTO compensation (uid,type) VALUES(?,?)")
	// Prepare("DELETE FROM compensation WHERE UID=?")
	// Prepare("update classes set Name=?,Designation=?,Description=?,lastmodby=? where ClassCode=?")
	// rlib.Errcheck(err)

	App.prepstmt.occAgrByProperty, err = App.dbrr.Prepare("SELECT RAID,OATID,PRID,UNITID,PID,PrimaryTenant,RentalStart,RentalStop,Renewal,ProrationMethod,ScheduledRent,Frequency,SecurityDepositAmount,SpecialProvisions,LastModTime,LastModBy from rentalagreement where PRID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnit, err = App.dbrr.Prepare("SELECT UNITID,BLDGID,UTID,RID,AVAILID,DefaultOccType,OccType,LastModTime,LastModBy FROM unit where UNITID=?")
	rlib.Errcheck(err)
	App.prepstmt.getLedger, err = App.dbrr.Prepare("SELECT LID,AccountNo,Dt,Balance,Deposit FROM ledger where LID=?")
	rlib.Errcheck(err)
	App.prepstmt.getTransactant, err = App.dbrr.Prepare("SELECT TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModTime,LastModBy FROM transactant WHERE TCID=?")
	rlib.Errcheck(err)
	App.prepstmt.getTenant, err = App.dbrr.Prepare("SELECT TID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyAddressEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber FROM tenant where TID=?")
	rlib.Errcheck(err)
	App.prepstmt.getRentable, err = App.dbrr.Prepare("SELECT RID,LID,RTID,PRID,PID,RAID,UNITID,Name,ScheduledRent,Frequency,Assignment,Report,LastModTime,LastModBy FROM rentable where RID=?")
	rlib.Errcheck(err)
	App.prepstmt.getProspect, err = App.dbrr.Prepare("SELECT PRSPID,TCID,ApplicationFee FROM prospect where PRSPID=?")
	rlib.Errcheck(err)
	App.prepstmt.getPayor, err = App.dbrr.Prepare("SELECT PID,TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerZipcode,Occupation,LastModTime,LastModBy FROM payor where PID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitSpecialties, err = App.dbrr.Prepare("SELECT USPID FROM unitspecialties where UNITID=?")
	rlib.Errcheck(err)
	App.prepstmt.getUnitSpecialtyType, err = App.dbrr.Prepare("SELECT USPID,PRID,Name,Fee,Description FROM unitspecialtytypes where USPID=?")
	rlib.Errcheck(err)
}
