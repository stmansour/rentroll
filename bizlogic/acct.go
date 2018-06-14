package bizlogic

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"sort"
)

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
//
// @return - a list of accounts that can be parent accounts
//-----------------------------------------------------------------------------
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
//-----------------------------------------------------------------------------
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

// OKToDelete makes checks to determine whether or not it is OK to
// delete an account.
//
// INPUTS
//    l  =  ledger (gl account) in question
//
// RETURNS
//    bool        - true means it's ok to delete, false means do not delete
//    [] BizError - List of errors encountered, explains why the account
//                  cannot be deleted
//-----------------------------------------------------------------------------
func OKToDelete(ctx context.Context, l *rlib.GLAccount) (bool, []BizError) {
	var errlist []BizError
	ok := true // start out positive, change it if problems are encountered

	//----------------------------------------------------------
	// Are there any Ledger Entries?  If so, we cannot allow
	// the ledger to be deleted.
	//----------------------------------------------------------
	n, err := rlib.GetCountLedgerEntries(ctx, l.LID, l.BID)
	if err != nil {
		return false, bizErrSys(&err)
	}

	if n > 0 {
		errlist = append(errlist, BizErrors[AcctHasLedgerEntries])
		ok = false
	}

	//----------------------------------------------------------
	// Is it referenced by any account rule?
	//----------------------------------------------------------
	var sa []string
	for _, v := range rlib.RRdb.BizTypes[l.BID].AR {
		if l.LID == v.DebitLID || l.LID == v.CreditLID {
			sa = append(sa, v.Name)
		}
	}
	if len(sa) > 0 {
		ok = false
		sort.Strings(sa)
		rulenames := " ("
		for i := 0; i < len(sa); i++ {
			if i > 0 {
				rulenames += ", "
			}
			rulenames += sa[i]
		}
		b := BizErrors[AcctRefInRule]
		b.Message += rulenames + ")"
		errlist = append(errlist, b)
	}

	return ok, errlist
}

// xxx - It is possible that at time0 account X is not a parent, so it
// was used in an AccountRule.  Then, at a later time, a user updates
// the Chart of Accounts and attempts to use account X as a parent

// SaveGLAccount saves or updates the supplied ledger. It performs multiple
// checks to ensure that the values are valid and will not let invalid
// states or values be saved.
//
// INPUTS
//    l  =  ledger (gl account) to save
//-----------------------------------------------------------------------------
func SaveGLAccount(ctx context.Context, l *rlib.GLAccount) []BizError {
	//	p1 := int64(0)
	var (
		err     error
		errlist []BizError
	)

	accts, err := rlib.GetGLAccountMap(ctx, l.BID)
	if err != nil {
		return bizErrSys(&err)
	}

	rules, err := rlib.GetARMap(ctx, l.BID)
	if err != nil {
		return bizErrSys(&err)
	}

	//-------------------------------------------------------------------------
	// First, ensure that AllowPosts is in the correct state:
	//     * If any account rule refers to this account it MUST  1. In this
	//       case, no other account can call this account as its parent.
	//     * If it is 0, it means that it is the parent to at least one other
	//       account.  If this is NOT the case then we will quietly change the
	//       AllowPost value to 1. This represents how people will approach
	//       the problem... they will remove all the parent relationships from
	//       the accounts first -- then they would expect the account to be
	//       usable in an account rule.
	//
	// If the input ledger is a brand new and unsaved (i.e., its LID == 0) then
	// it's OK to allow posts with no further checking because it is not
	// possible that any other account has listed it as a parent.
	//-------------------------------------------------------------------------
	if l.LID > 0 { //  is this an existing ledger?
		if l.AllowPost { // if so, is it allowing posts?
			//-------------------------------------------------------
			// Posts are allowed as long as no account refers to this
			// account as its parent. Make sure that's the case.
			//-------------------------------------------------------
			found := false
			for _, v := range accts {
				found = (v.PLID == l.LID)
				if found {
					break
				}
			}
			if found {
				errlist = append(errlist, BizErrors[PostToSummaryAcct])
				return errlist
			}
		}
		if !l.AllowPost { // is this a summary account
			//-------------------------------------------------------
			// AllowPost can be 0 unless there is already an Account
			// Rule that uses it for debit or credit.  Make sure no
			// Rule is using this account...
			//-------------------------------------------------------
			found := false
			for _, v := range rules {
				found = v.DebitLID == l.LID || v.CreditLID == l.LID
				if found {
					break
				}
			}
			if found {
				errlist = append(errlist, BizErrors[RuleUsesAcct])
				return errlist
			}
		}
	}
	//---------------------------------------------------------------
	// Even though no account rule currently uses this account, the
	// user may be trying to remove all the parent-child relation-
	// ships so that they CAN use it as to post to. So, if this
	// account is not a Parent to any other account then it is OK to
	// allow posts and we should set the value to 1.
	//---------------------------------------------------------------
	if !l.AllowPost {
		if l.LID == 0 { // is this a new account?
			l.AllowPost = true // if so, then we can allow posts until we find relationships that don't allow it
		} else {
			found := false // assume this is not a parent to any account
			for _, v := range accts {
				found = (v.PLID == l.LID) // is l.PID a parent?
				if found {                // yes?
					break // we're done
				}
			}
			if !found { // is l.PID free of all Parent-child relationships?
				l.AllowPost = true // if so, then we can allow posts
			}
		}
	}

	//-----------------------------------------------------------------
	// Only update the FLAGS that a client is allowed to change...
	//-----------------------------------------------------------------
	// var lcur rlib.GLAccount // if it is an existing account, load current version
	// if l.LID > 0 {
	// 	lcur = rlib.GetLedger(l.LID)
	// }
	curflags := l.FLAGS // start with what the client sent
	// DO ANY FLAG CHECKING NEEDED
	// at the moment there are no flags to check
	curflags = 0 // change this as needed if any flags are defined
	l.FLAGS = curflags

	//-----------------------------------------------------------------
	// OK, we've made all the checks we know about.  Now we can save it
	//-----------------------------------------------------------------
	if l.LID == 0 {
		_, err = rlib.InsertLedger(ctx, l)
		if err != nil {
			e := fmt.Errorf("Error saving Account %s, Error:= %s", l.Name, err.Error())
			return bizErrSys(&e)
		}
	} else {
		err = rlib.UpdateLedger(ctx, l)
		if err != nil {
			e := fmt.Errorf("Error updating account %s, Error:= %s", l.Name, err.Error())
			return bizErrSys(&e)
		}
	}

	return nil
}
