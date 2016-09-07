package rlib

import (
	"fmt"
	"sort"
	"time"
)

//-------------------------------------------------
//  ASSESSMENT
//-------------------------------------------------

// IDtoString is the generic ID prenter. Provide it a prefix and an id
// and it will return the formatted id as a string.
func IDtoString(pre string, id int64) string {
	if 0 == id {
		return "0"
	}
	return fmt.Sprintf("%s%08d", pre, id)
}

// IDtoString is the method to produce a consistent printable id string
func (t *Assessment) IDtoString() string {
	return IDtoString("ASM", t.ASMID)
}

//-------------------------------------------------
//  BUSINESS
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (t *Business) IDtoString() string {
	return IDtoString("B", t.BID)
}

//-------------------------------------------------
//  CUSTOM ATTRIBUTE
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (t *CustomAttribute) IDtoString() string {
	return IDtoString("C", t.CID)
}

// TypeToString returns a string describing the data type of the cell.
func (t *CustomAttribute) TypeToString() string {
	switch t.Type {
	case CUSTSTRING:
		return "string"
	case CUSTINT:
		return "int"
	case CUSTUINT:
		return "uint"
	case CUSTFLOAT:
		return "float"
	case CUSTDATE:
		return "date"
	}
	return "unknown"
}

// IDtoString is the method to produce a consistent printable id string
func (t *Deposit) IDtoString() string {
	return IDtoString("DEP", t.DEPID)
}

// IDtoString is the method to produce a consistent printable id string
func (t *DepositMethod) IDtoString() string {
	return IDtoString("DPM", t.DPMID)
}

// IDtoString is the method to produce a consistent printable id string
func (t *Depository) IDtoString() string {
	return IDtoString("DEP", t.DEPID)
}

// IDtoString is the method to produce a consistent printable id string
func (t *GLAccount) IDtoString() string {
	return IDtoString("L", t.LID)
}

//-------------------------------------------------
//  INVOICE
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Invoice) IDtoString() string {
	return IDtoString("IN", a.InvoiceNo)
}

// IDtoString is the method to produce a consistent printable id string
func (a *LedgerMarker) IDtoString() string {
	return IDtoString("LM", a.LMID)
}

// IDtoString is the method to produce a consistent printable id string
func (a *PaymentType) IDtoString() string {
	return IDtoString("PMT", a.PMTID)
}

//-------------------------------------------------
//  RECEIPT
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Receipt) IDtoString() string {
	return IDtoString("RCPT", a.RCPTID)
}

//-------------------------------------------------
//  RATE PLAN
//-------------------------------------------------

// IDtoString for RatePlans returns a unique identifier string.
func (t *RatePlan) IDtoString() string {
	return IDtoString("RP", t.RPID)
}

// IDtoString for RatePlanRefs returns a unique identifier string.
func (t *RatePlanRef) IDtoString() string {
	return IDtoString("RPR", t.RPRID)
}

// IDtoString for Rentables returns a unique identifier string.
func (t *Rentable) IDtoString() string {
	return IDtoString("R", t.RID)
}

//-------------------------------------------------
//  RENTAL AGREEMENT
//-------------------------------------------------

// IDtoString for RentalAgreements returns a unique identifier string.
func (t *RentalAgreement) IDtoString() string {
	return IDtoString("RA", t.RAID)
}

// IDtoString for RentalAgreementTemplate returns a unique identifier string.
func (t *RentalAgreementTemplate) IDtoString() string {
	return IDtoString("RAT", t.RATID)
}

//-------------------------------------------------
//  RENTABLE SPECIALTY
//-------------------------------------------------

// IDtoString for RentableSpecialty returns a unique identifier string.
func (t *RentableSpecialty) IDtoString() string {
	return IDtoString("RSP", t.RSPID)
}

// IDtoString for RentableType returns a unique identifier string.
func (t *RentableType) IDtoString() string {
	return IDtoString("RT", t.RTID)
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

// GetUserNameList loops through all the rentables associated with this rental agreement. It returns an array
// of strings with all the User names associated with each Rentable in the Rental Agreement for the supplied
// time range
func (t *RentalAgreement) GetUserNameList(d1, d2 *time.Time) []string {
	var m []string
	c := make(map[string]int)
	rl := GetRentalAgreementRentables(t.RAID, d1, d2)

	for i := 0; i < len(rl); i++ {
		r := GetRentable(rl[i].RID)
		n := r.GetUserNameList(d1, d2)
		for j := 0; j < len(n); j++ { // loop through, but only add unique names
			_, ok := c[n[j]] // have we seen this name?
			if !ok {         // if not, then...
				m = append(m, n[j]) // ...add it
				c[n[j]] = 1         // but mark that we've seen it
			}
		}
	}
	return m
}

//-------------------------------------------------
//  TRANSACTANT
//-------------------------------------------------

// IDtoString for XPerson returns a unique identifier string.
func (t *XPerson) IDtoString() string {
	return IDtoString("TC", t.Trn.TCID)
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
