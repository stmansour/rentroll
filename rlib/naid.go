package rlib

import (
	"fmt"
	"sort"
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

// IDtoString is the method to produce a consistent printable id string
func (t *GLAccount) IDtoString() string {
	return fmt.Sprintf("L%08d", t.LID)
}

//-------------------------------------------------
//  INVOICE
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Invoice) IDtoString() string {
	return fmt.Sprintf("IN%08d", a.InvoiceNo)
}

// IDtoString is the method to produce a consistent printable id string
func (a *LedgerMarker) IDtoString() string {
	return fmt.Sprintf("LM%08d", a.LMID)
}

//-------------------------------------------------
//  RECEIPT
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Receipt) IDtoString() string {
	return fmt.Sprintf("RCPT%08d", a.RCPTID)
}

//-------------------------------------------------
//  RATE PLAN
//-------------------------------------------------

// IDtoString for RatePlans returns a unique identifier string.
func (t *RatePlan) IDtoString() string {
	return fmt.Sprintf("RP%08d", t.RPID)
}

// IDtoString for RatePlanRefs returns a unique identifier string.
func (t *RatePlanRef) IDtoString() string {
	return fmt.Sprintf("RPR%08d", t.RPRID)
}

// IDtoString for Rentables returns a unique identifier string.
func (t *Rentable) IDtoString() string {
	return fmt.Sprintf("R%08d", t.RID)
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

// IDtoString for RentableSpecialty returns a unique identifier string.
func (t *RentableSpecialty) IDtoString() string {
	return fmt.Sprintf("RSP%08d", t.RSPID)
}

// IDtoString for RentableType returns a unique identifier string.
func (t *RentableType) IDtoString() string {
	return fmt.Sprintf("RT%08d", t.RTID)
}

// GetUserNameList returns an array of strings with all the User names associated with the Rentable. the strings are sorted alphabetically
func (t *Rentable) GetUserNameList(d1, d2 *time.Time) []string {
	var m []string
	users := GetRentableUsers(t.RID, d1, d2) // get all defined renters for this period
	for i := 0; i < len(users); i++ {
		var tr Transactant
		GetTransactant(users[i].TCID, &tr)
		m = append(m, tr.GetUserName())
	}
	sort.Strings(m)
	return m
}

// GetPayorNameList returns an array of strings with all the Payor names associated with the Rental Agreement
func (t *RentalAgreement) GetPayorNameList(d1, d2 *time.Time) []string {
	var m []string
	payors := GetRentalAgreementPayors(t.RAID, d1, d2) // get all defined renters for this period
	for i := 0; i < len(payors); i++ {
		var tr Transactant
		GetTransactant(payors[i].TCID, &tr)
		m = append(m, tr.GetUserName())
	}
	return m
}

//-------------------------------------------------
//  TRANSACTANT
//-------------------------------------------------

// IDtoString for XPerson returns a unique identifier string.
func (t *XPerson) IDtoString() string {
	return fmt.Sprintf("TC%08d", t.Trn.TCID)
}

// GetUserName returns a string with the user's first, middle, and last name
func (t *Transactant) GetUserName() string {
	if t.IsCompany > 0 {
		return t.CompanyName
	}
	s := t.FirstName + " "
	if len(t.MiddleName) > 0 {
		s += t.MiddleName + " "
	}
	s += t.LastName
	return s
}
