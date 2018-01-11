package rlib

import "database/sql"

// PBprepSQL is the structure of prepared sql statements for the Phonebook db
type PBprepSQL struct {
	GetCompanyByDesignation      *sql.Stmt
	GetCompany                   *sql.Stmt
	GetBusinessUnitByDesignation *sql.Stmt
	GetDirectoryPerson           *sql.Stmt
}

func buildPBPreparedStatements() {
	var err error
	var flds string

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

	//==========================================
	// GetDirectoryPerson
	//==========================================
	flds = "UID,UserName,LastName,MiddleName,FirstName,PreferredName,PreferredName,OfficePhone,CellPhone"
	RRdb.PBsql.GetDirectoryPerson, err = RRdb.Dbdir.Prepare("SELECT " + flds + " FROM people WHERE UID=?")
	Errcheck(err)

}
