package rlib

import (
	"fmt"
	"time"
)

//-------------------------------------------------
//  ASSESSMENT
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (t *Assessment) IDtoString() string {
	return fmt.Sprintf("ASM%08d", t.ASMID)
}

//-------------------------------------------------
//  BUSINESS
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (t *Business) IDtoString() string {
	return fmt.Sprintf("B%08d", t.BID)
}

//-------------------------------------------------
//  INVOICE
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Invoice) IDtoString() string {
	return fmt.Sprintf("IN%08d", a.InvoiceNo)
}

//-------------------------------------------------
//  RECEIPT
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Receipt) IDtoString() string {
	return fmt.Sprintf("RCPT%08d", a.RCPTID)
}

//-------------------------------------------------
//  RENTAL AGREEMENT
//-------------------------------------------------

// IDtoString for RentalAgreements returns a unique identifier string.
func (t *RentalAgreement) IDtoString() string {
	return fmt.Sprintf("RA%08d", t.RAID)
}

//-------------------------------------------------
//  RENTABLE SPECIALTY
//-------------------------------------------------

// IDtoString for RentalAgreements returns a unique identifier string.
func (t *RentableSpecialtyType) IDtoString() string {
	return fmt.Sprintf("RSP%08d", t.RSPID)
}

// GetUserNameList returns an array of strings with all the User names associated with the Rentable
func (t *RentalAgreement) GetUserNameList(d1, d2 *time.Time) []string {
	var m []string
	users := GetRentableUsers(t.RAID, d1, d2) // get all defined renters for this period
	for i := 0; i < len(users); i++ {
		var tr Transactant
		GetTransactant(users[i].USERID, &tr)
		m = append(m, tr.GetUserName())
	}
	return m
}

// GetPayorNameList returns an array of strings with all the Payor names associated with the Rental Agreement
func (t *RentalAgreement) GetPayorNameList(d1, d2 *time.Time) []string {
	var m []string
	payors := GetRentalAgreementPayors(t.RAID, d1, d2) // get all defined renters for this period
	for i := 0; i < len(payors); i++ {
		var tr Transactant
		GetTransactant(payors[i].PID, &tr)
		m = append(m, tr.GetUserName())
	}
	return m
}

//-------------------------------------------------
//  TRANSACTANT
//-------------------------------------------------

// GetUserName returns a string with the user's first, middle, and last name
func (t *Transactant) GetUserName() string {
	s := t.FirstName + " "
	if len(t.MiddleName) > 0 {
		s += t.MiddleName + " "
	}
	s += t.LastName
	return s
}
