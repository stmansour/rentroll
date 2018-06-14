package rlib

import (
	"context"
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

// IDtoShortString is the method to produce a consistent printable id string
func (t *Assessment) IDtoShortString() string {
	return IDtoShortString("ASM", t.ASMID)
}

//-------------------------------------------------
//  BUSINESS
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (t *Business) IDtoString() string {
	return IDtoString("B", t.BID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (t *Business) IDtoShortString() string {
	return IDtoShortString("B", t.BID)
}

//-------------------------------------------------
//  CUSTOM ATTRIBUTE
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (t *CustomAttribute) IDtoString() string {
	return IDtoString("C", t.CID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (t *CustomAttribute) IDtoShortString() string {
	return IDtoShortString("C", t.CID)
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

// IDtoShortString is the method to produce a consistent printable id string
func (t *Deposit) IDtoShortString() string {
	return IDtoShortString("DEP", t.DEPID)
}

// IDtoString is the method to produce a consistent printable id string
func (t *DepositMethod) IDtoString() string {
	return IDtoString("DPM", t.DPMID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (t *DepositMethod) IDtoShortString() string {
	return IDtoShortString("DPM", t.DPMID)
}

// IDtoString is the method to produce a consistent printable id string
func (t *Depository) IDtoString() string {
	return IDtoString("DEP", t.DEPID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (t *Depository) IDtoShortString() string {
	return IDtoShortString("DEP", t.DEPID)
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

// IDtoShortString is the method to produce a consistent printable id string
func (t *GLAccount) IDtoShortString() string {
	return IDtoShortString("L", t.LID)
}

//-------------------------------------------------
//  INVOICE
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Invoice) IDtoString() string {
	return IDtoString("IN", a.InvoiceNo)
}

// IDtoShortString is the method to produce a consistent printable id string
func (a *Invoice) IDtoShortString() string {
	return IDtoShortString("IN", a.InvoiceNo)
}

// IDtoString is the method to produce a consistent printable id string
func (a *Journal) IDtoString() string {
	return IDtoString("J", a.JID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (a *Journal) IDtoShortString() string {
	return IDtoShortString("J", a.JID)
}

// IDtoString is the method to produce a consistent printable id string
func (a *LedgerMarker) IDtoString() string {
	return IDtoString("LM", a.LMID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (a *LedgerMarker) IDtoShortString() string {
	return IDtoShortString("LM", a.LMID)
}

// IDtoString is the method to produce a consistent printable id string
func (a *PaymentType) IDtoString() string {
	return IDtoString("PMT", a.PMTID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (a *PaymentType) IDtoShortString() string {
	return IDtoShortString("PMT", a.PMTID)
}

//-------------------------------------------------
//  RECEIPT
//-------------------------------------------------

// IDtoString is the method to produce a consistent printable id string
func (a *Receipt) IDtoString() string {
	return IDtoString("RCPT", a.RCPTID)
}

// IDtoShortString is the method to produce a consistent printable id string
func (a *Receipt) IDtoShortString() string {
	return IDtoShortString("RCPT", a.RCPTID)
}

//-------------------------------------------------
//  RATE PLAN
//-------------------------------------------------

// IDtoString for RatePlans returns a unique identifier string.
func (t *RatePlan) IDtoString() string {
	return IDtoString("RP", t.RPID)
}

// IDtoShortString for RatePlans returns a unique identifier string.
func (t *RatePlan) IDtoShortString() string {
	return IDtoShortString("RP", t.RPID)
}

// IDtoString for RatePlanRefs returns a unique identifier string.
func (t *RatePlanRef) IDtoString() string {
	return IDtoString("RPR", t.RPRID)
}

// IDtoShortString for RatePlanRefs returns a unique identifier string.
func (t *RatePlanRef) IDtoShortString() string {
	return IDtoShortString("RPR", t.RPRID)
}

// IDtoString for Rentables returns a unique identifier string.
func (t *Rentable) IDtoString() string {
	return IDtoString("R", t.RID)
}

// IDtoShortString for Rentables returns a unique identifier string.
func (t *Rentable) IDtoShortString() string {
	return IDtoShortString("R", t.RID)
}

//-------------------------------------------------
//  RENTAL AGREEMENT
//-------------------------------------------------

// IDtoString for RentalAgreements returns a unique identifier string.
func (t *RentalAgreement) IDtoString() string {
	return IDtoString("RA", t.RAID)
}

// IDtoShortString for RentalAgreements returns a unique identifier string.
func (t *RentalAgreement) IDtoShortString() string {
	return IDtoShortString("RA", t.RAID)
}

// IDtoString for RentalAgreementTemplate returns a unique identifier string.
func (t *RentalAgreementTemplate) IDtoString() string {
	return IDtoString("RAT", t.RATID)
}

// IDtoShortString for RentalAgreementTemplate returns a unique identifier string.
func (t *RentalAgreementTemplate) IDtoShortString() string {
	return IDtoShortString("RAT", t.RATID)
}

//-------------------------------------------------
//  RENTABLE SPECIALTY
//-------------------------------------------------

// IDtoString for RentableSpecialty returns a unique identifier string.
func (t *RentableSpecialty) IDtoString() string {
	return IDtoString("RSP", t.RSPID)
}

// IDtoShortString for RentableSpecialty returns a unique identifier string.
func (t *RentableSpecialty) IDtoShortString() string {
	return IDtoShortString("RSP", t.RSPID)
}

// IDtoString for RentableType returns a unique identifier string.
func (t *RentableType) IDtoString() string {
	return IDtoString("RT", t.RTID)
}

// IDtoShortString for RentableType returns a unique identifier string.
func (t *RentableType) IDtoShortString() string {
	return IDtoShortString("RT", t.RTID)
}

// IDtoString for Vehicle returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *Vehicle) IDtoString() string {
	return IDtoString("V", t.VID)
}

// IDtoShortString for Vehicle returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *Vehicle) IDtoShortString() string {
	return IDtoShortString("V", t.VID)
}

// GetUserNameList returns an array of strings with all the User names
// associated with the Rentable. the strings are sorted alphabetically
//-----------------------------------------------------------------------------
func (t *Rentable) GetUserNameList(ctx context.Context, d1, d2 *time.Time) ([]string, error) {
	var m []string

	// TODO(Steve): ignore error?
	users, err := GetRentableUsersInRange(ctx, t.RID, d1, d2) // get all defined renters for this period
	if err != nil {
		return m, err
	}

	for i := 0; i < len(users); i++ {
		var tr Transactant
		err = GetTransactant(ctx, users[i].TCID, &tr)
		if err != nil {
			return m, err
		}
		m = append(m, tr.GetUserName())
	}
	sort.Strings(m)
	return m, err
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
func (t *RentalAgreement) GetPayorNameList(ctx context.Context, d1, d2 *time.Time) ([]string, error) {
	var m []string
	payors, err := GetRentalAgreementPayorsInRange(ctx, t.RAID, d1, d2) // get all defined renters for this period
	if err != nil {
		return m, err
	}
	for i := 0; i < len(payors); i++ {
		var tr Transactant
		err := GetTransactant(ctx, payors[i].TCID, &tr)
		if err != nil {
			return m, err
		}
		m = append(m, tr.GetUserName())
	}
	return m, err
}

// GetUserNameList loops through all the rentables associated with this rental
// agreement. It returns an array of strings with all the User names
// associated with each Rentable in the Rental Agreement for the supplied time
// range.
//-----------------------------------------------------------------------------
func (t *RentalAgreement) GetUserNameList(ctx context.Context, d1, d2 *time.Time) ([]string, error) {
	var m []string
	c := make(map[string]int)
	rl, err := GetRentalAgreementRentables(ctx, t.RAID, d1, d2)
	if err != nil {
		return m, err
	}

	for i := 0; i < len(rl); i++ {
		r, err := GetRentable(ctx, rl[i].RID)
		if err != nil {
			return m, err
		}
		n, err := r.GetUserNameList(ctx, d1, d2)
		if err != nil {
			return m, err
		}
		for j := 0; j < len(n); j++ { // loop through, but only add unique names
			_, ok := c[n[j]] // have we seen this name?
			if !ok {         // if not, then...
				m = append(m, n[j]) // ...add it
				c[n[j]] = 1         // but mark that we've seen it
			}
		}
	}
	return m, err
}

// GetTheRentableName is used to get the name of the highest priced rentable
// (based on the rentable's type).  There is an assumption that all rentables
// belonging to this rental agreement have the same rent cycle.
//-----------------------------------------------------------------------------
func (t *RentalAgreement) GetTheRentableName(ctx context.Context, d1, d2 *time.Time) (string, error) {
	var (
		xbiz        XBusiness
		theRentable Rentable
	)
	err := GetXBusiness(ctx, t.BID, &xbiz)
	if err != nil {
		return theRentable.RentableName, err
	}
	max := float64(0)

	// Console("Entered: RentalAgreement.GetTheRentableName\n")
	rl, err := GetRentalAgreementRentables(ctx, t.RAID, d1, d2)
	if err != nil {
		return theRentable.RentableName, err
	}
	for i := 0; i < len(rl); i++ {
		r, err := GetRentable(ctx, rl[i].RID)
		if err != nil {
			return theRentable.RentableName, err
		}
		amt, err := GetRentableMarketRate(ctx, &xbiz, r.RID, d1, d2)
		if err != nil {
			return theRentable.RentableName, err
		}
		// Console("Rentable = %d, MarketRate = %.2f\n", r.RID, amt)
		if amt > max {
			theRentable = r
		}
	}
	return theRentable.RentableName, err
}

//-------------------------------------------------
//  TRANSACTANT
//-------------------------------------------------

// GetTransactantLastName returns the Last name of the user if the Transactant
// is a person or the CompanyName if it is a company
//-----------------------------------------------------------------------------
func (t *Transactant) GetTransactantLastName() string {
	if t.IsCompany {
		return t.CompanyName
	}
	return t.LastName
}

// GetPayorName returns an array of strings that contains the last names
// of every Payor responsible for this Rental Agreement during the timespan d1,d2.
func (t *Transactant) GetPayorName() string {
	if t.IsCompany {
		return t.CompanyName
	}
	return t.FirstName + " " + t.LastName
}

// GetFullTransactantName returns the full user name if the Transactant is a
// person or the CompanyName if it is a company
//-----------------------------------------------------------------------------
func (t *Transactant) GetFullTransactantName() string {
	if t.IsCompany {
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

// IDtoShortString for XPerson returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *XPerson) IDtoShortString() string {
	return IDtoShortString("TC", t.Trn.TCID)
}

// IDtoString for XPerson returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *Transactant) IDtoString() string {
	return IDtoString("TC", t.TCID)
}

// IDtoShortString for XPerson returns a unique identifier string.
//-----------------------------------------------------------------------------
func (t *Transactant) IDtoShortString() string {
	return IDtoShortString("TC", t.TCID)
}

// GetUserName returns a string with the user's first, middle, and last name
//-----------------------------------------------------------------------------
func (t *Transactant) GetUserName() string {
	if t.IsCompany {
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
