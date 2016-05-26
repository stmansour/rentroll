package rlib

import "time"

// Company is the structure of company attributes
type Company struct {
	CoCode           int
	LegalName        string
	CommonName       string
	Address          string
	Address2         string
	City             string
	State            string
	PostalCode       string
	Country          string
	Phone            string
	Fax              string
	Email            string
	Designation      string
	Active           int
	EmploysPersonnel int
	LastModTime      time.Time
	LastModBy        int
}

// BusinessUnit is the structure of bizunit attributes
// The historical name for BusinessUnit in Phonebook was "class" after the QuickBooks equivalent
type BusinessUnit struct {
	ClassCode   int
	CoCode      int
	Name        string
	Designation string
	Description string
	LastModTime time.Time
	LastModBy   int
	C           Company // parent company
}
