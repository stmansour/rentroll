package rlib

// GetCompanyByDesignation returns a Company struct for the Phonebook company with the
// supplied designation. If no such company exists, c.CoCode will be 0
func GetCompanyByDesignation(des string) (Company, error) {
	var c Company
	err := RRdb.PBsql.GetCompanyByDesignation.QueryRow(des).Scan(&c.CoCode, &c.LegalName, &c.CommonName,
		&c.Address, &c.Address2, &c.City, &c.State, &c.PostalCode, &c.Country, &c.Phone,
		&c.Fax, &c.Email, &c.Designation, &c.Active, &c.EmploysPersonnel, &c.LastModTime,
		&c.LastModBy)
	return c, err
}
