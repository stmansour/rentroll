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

// IDtoShortString is the generic ID prenter. Provide it a prefix and an id
// and it will return the formatted id as a string.
func IDtoShortString(pre string, id int64) string {
	if 0 == id {
		return "0"
	}
	return fmt.Sprintf("%s-%d", pre, id)
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
func (t *Expense) IDtoString() string {
	return IDtoString("EXP", t.EXPID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (t *Expense) IDtoShortString() string {
	return IDtoShortString("EXP", t.EXPID)
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
func (a *Journal) IDtoString() string {
	return IDtoString("J", a.JID)
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

// IDtoString for Vehicle returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *Vehicle) IDtoString() string {
	return IDtoString("V", t.VID)
}

// GetUserNameList returns an array of strings with all the User names
// associated with the Rentable. the strings are sorted alphabetically
//-----------------------------------------------------------------------------
func (t *Rentable) GetUserNameList(d1, d2 *time.Time) []string {
	var m []string
	users := GetRentableUsersInRange(t.RID, d1, d2) // get all defined renters for this period
	for i := 0; i < len(users); i++ {
		var tr Transactant
		GetTransactant(users[i].TCID, &tr)
		m = append(m, tr.GetUserName())
	}
	sort.Strings(m)
	return m
}

//-------------------------------------------------
// RENTAL AGREEMENT
//-------------------------------------------------

// GetPayorLastNames returns an array of strings that contains the last names
// of every Payor responsible for this Rental Agreement during the timespan d1,d2.
func (t *RentalAgreement) GetPayorLastNames(d1, d2 *time.Time) []string {
	var sa []string
	for i := 0; i < len(t.P); i++ {
		if d1.Before(t.AgreementStop) && d2.After(t.AgreementStart) {
			sa = append(sa, t.P[i].Trn.LastName)
		}
	}
	return sa
}

// GetPayorNameList returns an array of strings with all the Payor names
// associated with the Rental Agreement
//-----------------------------------------------------------------------------
func (t *RentalAgreement) GetPayorNameList(d1, d2 *time.Time) []string {
	var m []string
	payors := GetRentalAgreementPayorsInRange(t.RAID, d1, d2) // get all defined renters for this period
	for i := 0; i < len(payors); i++ {
		var tr Transactant
		GetTransactant(payors[i].TCID, &tr)
		m = append(m, tr.GetUserName())
	}
	return m
}

// GetUserNameList loops through all the rentables associated with this rental
// agreement. It returns an array of strings with all the User names
// associated with each Rentable in the Rental Agreement for the supplied time
// range.
//-----------------------------------------------------------------------------
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

// GetTheRentableName is used to get the name of the highest priced rentable
// (based on the rentable's type).  There is an assumption that all rentables
// belonging to this rental agreement have the same rent cycle.
//-----------------------------------------------------------------------------
func (t *RentalAgreement) GetTheRentableName(d1, d2 *time.Time) string {
	var xbiz XBusiness
	GetXBusiness(t.BID, &xbiz)
	max := float64(0)
	var theRentable Rentable

	// Console("Entered: RentalAgreement.GetTheRentableName\n")
	rl := GetRentalAgreementRentables(t.RAID, d1, d2)
	for i := 0; i < len(rl); i++ {
		r := GetRentable(rl[i].RID)
		amt := GetRentableMarketRate(&xbiz, &r, d1, d2)
		// Console("Rentable = %d, MarketRate = %.2f\n", r.RID, amt)
		if amt > max {
			theRentable = r
		}
	}
	return theRentable.RentableName
}

//-------------------------------------------------
//  TRANSACTANT
//-------------------------------------------------

// GetTransactantLastName returns the Last name of the user if the Transactant
// is a person or the CompanyName if it is a company
//-----------------------------------------------------------------------------
func (t *Transactant) GetTransactantLastName() string {
	if t.IsCompany > 0 {
		return t.CompanyName
	}
	return t.LastName
}

// GetPayorName returns an array of strings that contains the last names
// of every Payor responsible for this Rental Agreement during the timespan d1,d2.
func (t *Transactant) GetPayorName() string {
	if t.IsCompany > 0 {
		return t.CompanyName
	}
	return t.FirstName + " " + t.LastName
}

// GetFullTransactantName returns the full user name if the Transactant is a
// person or the CompanyName if it is a company
//-----------------------------------------------------------------------------
func (t *Transactant) GetFullTransactantName() string {
	if t.IsCompany > 0 {
		return t.CompanyName
	}
	s := t.FirstName
	if len(t.MiddleName) > 0 {
		s += " " + t.MiddleName
	}
	return s + " " + t.LastName
}

// IDtoString for XPerson returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *XPerson) IDtoString() string {
	return IDtoString("TC", t.Trn.TCID)
}

// IDtoString for XPerson returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *Transactant) IDtoString() string {
	return IDtoString("TC", t.TCID)
}

// GetUserName returns a string with the user's first, middle, and last name
//-----------------------------------------------------------------------------
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

// SingleLineAddress returns the transactant's address
//-----------------------------------------------------------------------------
func (t *Transactant) SingleLineAddress() string {
	a := t.Address
	if len(t.Address2) > 0 {
		a += ", " + t.Address2
	}
	if len(t.City) > 0 {
		a += ", " + t.City
	}
	if len(t.State) > 0 {
		a += ", " + t.State
	}
	if len(t.PostalCode) > 0 {
		a += " " + t.PostalCode
	}
	if len(t.Country) > 0 {
		a += ", " + t.Country
	}
	return a
}
