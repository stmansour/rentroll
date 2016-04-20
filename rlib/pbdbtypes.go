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
