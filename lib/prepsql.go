package main

func buildPreparedStatements() {
	var err error
	// Prepare("select deduction from deductions where uid=?")
	// Prepare("select type from compensation where uid=?")
	// Prepare("INSERT INTO compensation (uid,type) VALUES(?,?)")
	// Prepare("DELETE FROM compensation WHERE UID=?")
	// Prepare("update classes set Name=?,Designation=?,Description=?,lastmodby=? where ClassCode=?")
	// errcheck(err)

	App.prepstmt.occAgrByProperty, err = App.dbrr.Prepare("SELECT OAID,OATID,PRID,UNITID,PID,PrimaryTenant,OccupancyStart,OccupancyStop,Renewal,ProrationMethod,SecurityDepositAmount from occupancyagreement where PRID=?")
	errcheck(err)
}
