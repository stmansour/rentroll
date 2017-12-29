package rlib

import "context"

// GetCompanyByDesignation returns a Company struct for the Phonebook company with the
// supplied designation. If no such company exists, c.CoCode will be 0
func GetCompanyByDesignation(ctx context.Context, des string) (Company, error) {
	var c Company

	// session... context
	if !RRdb.noAuth {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return c, ErrSessionRequired
		}
	}

	err := RRdb.PBsql.GetCompanyByDesignation.QueryRow(des).Scan(&c.CoCode, &c.LegalName, &c.CommonName,
		&c.Address, &c.Address2, &c.City, &c.State, &c.PostalCode, &c.Country, &c.Phone,
		&c.Fax, &c.Email, &c.Designation, &c.Active, &c.EmploysPersonnel, &c.LastModTime,
		&c.LastModBy)
	return c, err
}

// GetCompany returns a Company struct for the Phonebook company with the
// supplied designation. If no such company exists, c.CoCode will be 0
func GetCompany(ctx context.Context, n int64) (Company, error) {
	var c Company

	// session... context
	if !RRdb.noAuth {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return c, ErrSessionRequired
		}
	}

	err := RRdb.PBsql.GetCompany.QueryRow(n).Scan(&c.CoCode, &c.LegalName, &c.CommonName,
		&c.Address, &c.Address2, &c.City, &c.State, &c.PostalCode, &c.Country, &c.Phone,
		&c.Fax, &c.Email, &c.Designation, &c.Active, &c.EmploysPersonnel, &c.LastModTime,
		&c.LastModBy)
	return c, err
}

// GetBusinessUnitByDesignation returns a Class (BusinessUnit) struct for the Phonebook class with the
// supplied designation. If no such class exists, c.CoCode will be 0
func GetBusinessUnitByDesignation(ctx context.Context, des string) (BusinessUnit, error) {
	var c BusinessUnit

	// session... context
	if !RRdb.noAuth {
		_, ok := SessionFromContext(ctx)
		if !ok {
			return c, ErrSessionRequired
		}
	}

	// err := RRdb.PBsql.GetBusinessUnitByDesignation.QueryRow(des).Scan(&c.ClassCode, &c.CoCode, &c.Name, &c.Designation, &c.Description, &c.LastModTime, &c.LastModBy)
	err := RRdb.Dbdir.QueryRow("SELECT ClassCode,CoCode,Name,Designation,Description,LastModTime,LastModBy FROM classes WHERE Designation=?", des).Scan(&c.ClassCode, &c.CoCode, &c.Name, &c.Designation, &c.Description, &c.LastModTime, &c.LastModBy)
	return c, err
}
