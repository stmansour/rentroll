package bizlogic

import "rentroll/rlib"

// PossibleParentAccounts returns the list of possible Parent Accounts.
// The only accounts that can be shown as possible Parent accounts
// in the Account edit form are those not used in an Account Rule.  To
// produce this list, we need to start with the full list of accounts and
// remove all accounts called out in the Assessment/Receipt Rules.  We
// cannot use the Summary Account attribute because that attribute is
// programmatically given to an account when another Account calls it
// out as its parent.  This creates a Chicken-and-Egg problem.  A newly
// created account that is intended to be a Summary Account does not yet
// have the Summary Account attribute, because no other Account lists it
// as its parent.  And if we use the Summary Account attribute to
// determine what accounts show up in the list of possible Parent accounts,
// our newly created account will not appear in that list so no account can
// ever add it as a parent.
// @return - a list of accounts that can be parent accounts
func PossibleParentAccounts(bid int64) []rlib.GLAccount {
	var m = map[int64]int{}

	// init the map with a 0 count for each account
	for _, v := range rlib.RRdb.BizTypes[bid].GLAccounts {
		m[v.LID] = 0
	}

	// Increment the value associated with each account for every time
	// it is called out in any account rule
	for _, v := range rlib.RRdb.BizTypes[bid].AR {
		m[v.DebitLID]++
		m[v.CreditLID]++
	}

	// return a list of acocun
	var n []rlib.GLAccount
	for _, v := range rlib.RRdb.BizTypes[bid].GLAccounts {
		if 0 == m[v.LID] {
			n = append(n, v)
		}
	}
	return n
}

// PossiblePostAccounts returns a list of accounts that are permissible to
// use in Assessment/Receipt Rules.  Essentially, it excludes any account
// that is used as a parent account.
//
// TODO - reimplement after fixes have been made to SummaryAccount flag
//        it will be much faster
//
// @return - a list of GLAccounts that can be used for posting
func PossiblePostAccounts(bid int64) []rlib.GLAccount {
	var m = map[int64]int{}

	// init the map with a 0 count for each account
	for _, v := range rlib.RRdb.BizTypes[bid].GLAccounts {
		m[v.LID] = 0
	}

	// For each account LID, increment m[LID] if it is used as
	// a parent
	for _, p := range rlib.RRdb.BizTypes[bid].GLAccounts {
		for _, v := range rlib.RRdb.BizTypes[bid].GLAccounts {
			if p.LID == v.PLID {
				m[p.LID]++
				break // no need to look any further
			}
		}
	}

	var n []rlib.GLAccount
	for _, v := range rlib.RRdb.BizTypes[bid].GLAccounts {
		if 0 == m[v.LID] { // if there were no references to this LID as a parent...
			n = append(n, v) // ... then it is OK to post to it
		}
	}
	return n
}

// SaveGLAccount saves or updates the supplied ledger.
// It loads the existing ledger prior to the saving, if it exists.
// If the new parent ledger is different than the old parent, it scans
// Since Summary Accounts cannot have entries posted to them, we need
// to filter the accounts available in the debit / credit dropdown
// lists in the Assessment/Receipt Rules form so that only accounts
// that are NOT Summary Accounts are available. We only post to
// accounts that are called out in Assessment/Receipt rules.
//
func SaveGLAccount(l *rlib.GLAccount) {
	//	p1 := int64(0)

}
