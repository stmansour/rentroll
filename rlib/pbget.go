package rlib

import (
	"context"
	"extres"
	"fmt"
)

// GetCompanyByDesignation returns a Company struct for the Phonebook company with the
// supplied designation. If no such company exists, c.CoCode will be 0
func GetCompanyByDesignation(ctx context.Context, des string) (Company, error) {
	var c Company

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return c, ErrSessionRequired
		}
	}

	err := RRdb.PBsql.GetCompanyByDesignation.QueryRow(des).Scan(&c.CoCode, &c.LegalName, &c.CommonName,
		&c.Address, &c.Address2, &c.City, &c.State, &c.PostalCode, &c.Country, &c.Phone,
		&c.Fax, &c.Email, &c.Designation, &c.Active, &c.EmploysPersonnel, &c.LastModTime,
		&c.LastModBy)
	SkipSQLNoRowsError(&err)
	return c, err
}

// GetCompany returns a Company struct for the Phonebook company with the
// supplied designation. If no such company exists, c.CoCode will be 0
func GetCompany(ctx context.Context, n int64) (Company, error) {
	var c Company

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return c, ErrSessionRequired
		}
	}

	err := RRdb.PBsql.GetCompany.QueryRow(n).Scan(&c.CoCode, &c.LegalName, &c.CommonName,
		&c.Address, &c.Address2, &c.City, &c.State, &c.PostalCode, &c.Country, &c.Phone,
		&c.Fax, &c.Email, &c.Designation, &c.Active, &c.EmploysPersonnel, &c.LastModTime,
		&c.LastModBy)
	SkipSQLNoRowsError(&err)
	return c, err
}

// GetBusinessUnitByDesignation returns a Class (BusinessUnit) struct for the Phonebook class with the
// supplied designation. If no such class exists, c.CoCode will be 0
func GetBusinessUnitByDesignation(ctx context.Context, des string) (BusinessUnit, error) {
	var c BusinessUnit

	// session... context
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return c, ErrSessionRequired
		}
	}

	// err := RRdb.PBsql.GetBusinessUnitByDesignation.QueryRow(des).Scan(&c.ClassCode, &c.CoCode, &c.Name, &c.Designation, &c.Description, &c.LastModTime, &c.LastModBy)
	err := RRdb.Dbdir.QueryRow("SELECT ClassCode,CoCode,Name,Designation,Description,LastModTime,LastModBy FROM classes WHERE Designation=?", des).Scan(&c.ClassCode, &c.CoCode, &c.Name, &c.Designation, &c.Description, &c.LastModTime, &c.LastModBy)
	SkipSQLNoRowsError(&err)
	return c, err
}

// GetDirectoryPerson reads the public fields for a person in Accord Directory
// based on the supplied UID.
func GetDirectoryPerson(ctx context.Context, uid int64) (DirectoryPerson, error) {
	var c DirectoryPerson

	// Console("RRdb.noAuth = %t, AppConfig.Env = %d\n", RRdb.noAuth, AppConfig.Env)
	if !(RRdb.noAuth && AppConfig.Env != extres.APPENVPROD) || !RRdb.noAuth {
		_, ok := SessionFromContext(ctx)
		if !ok {
			// Console("GetDirectoryPerson -- returning empty DirectoryPerson\n")
			return c, ErrSessionRequired
		}
	}

	err := RRdb.PBsql.GetDirectoryPerson.QueryRow(uid).Scan(&c.UID, &c.UserName, &c.LastName, &c.MiddleName, &c.FirstName, &c.PreferredName, &c.PreferredName, &c.OfficePhone, &c.CellPhone)
	SkipSQLNoRowsError(&err)
	// Console("GetDirectoryPerson -- read directory person. c.UserName = %s\n", c.UserName)
	return c, err
}

// DisplayName returns the name for use on Roller's interface
// for this user
func (t *DirectoryPerson) DisplayName() string {
	var name string
	name = t.PreferredName
	if len(name) == 0 {
		name = t.FirstName
	}
	if len(t.LastName) > 0 {
		if len(name) > 0 {
			name += " "
		}
		name += t.LastName
	}
	if len(name) == 0 {
		name = fmt.Sprintf("UID-%d", t.UID)
	}
	return name
}

// GetNameForUID reads the directory entry for the supplied uid and returns
// a formatted name string.  If a database error occurs, it will return
// a string of the format UID-nnn  where nnn is the UID as a string. So,
// for example, if you supply 0 for uid, the return string will be
// "UID-0".  The errors will be logged.
//
//
// INPUTS:
//  ctx - database context
//  uid - the directory UID of the person
//
// RETURNS:
//  formatted name string
//-----------------------------------------------------------------------------
func GetNameForUID(ctx context.Context, uid int64) string {
	funcname := "GetNameForUID"
	c, err := GetDirectoryPerson(ctx, uid)
	if err != nil {
		Ulog("%s: err = %s\n", funcname, err.Error())
		return fmt.Sprintf("UID-%d", uid)
	}
	return c.DisplayName()
}
