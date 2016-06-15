package rlib

func buildPBPreparedStatements() {
	var err error

	//==========================================
	// BUSINESS UNIT
	//==========================================
	RRdb.PBsql.GetBusinessUnitByDesignation, err = RRdb.Dbdir.Prepare("SELECT ClassCode,CoCode,Name,Designation,Description,LastModTime,LastModBy FROM classes WHERE Designation=?")
	Errcheck(err)

	//==========================================
	// COMPANY
	//==========================================
	RRdb.PBsql.GetCompany, err = RRdb.Dbdir.Prepare("SELECT CoCode,LegalName,CommonName,Address,Address2,City,State,PostalCode,Country,Phone,Fax,Email,Designation,Active,EmploysPersonnel,LastModTime,LastModBy FROM companies WHERE CoCode=?")
	Errcheck(err)
	RRdb.PBsql.GetCompanyByDesignation, err = RRdb.Dbdir.Prepare("SELECT CoCode,LegalName,CommonName,Address,Address2,City,State,PostalCode,Country,Phone,Fax,Email,Designation,Active,EmploysPersonnel,LastModTime,LastModBy FROM companies WHERE Designation=?")
	Errcheck(err)

}
