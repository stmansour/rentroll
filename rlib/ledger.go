package rlib

import (
	"context"
	"time"
)

// LedgerMarkers are records that contain the balance for the associated
// account at a particulart date/time. These speed up the calculation of
// an account balance at a particular date/time.
//
// LedgerMarkers can be written to the database at any point in time, but
// by convention, we write them when a financial period is closed.
//
// Types of LedgerMarkers:
//   * Account Balance - there will be a ledger for every account in
//     the chart of accounts
//   * Rental Agreement Balance - the balance associated with a rental
//     agreement for the specific account
//   * Rentable Balance - the balance associated with a particular
//     rentable for the specific account
//

// RemoveLedgerEntries clears out the records in the supplied range provided the range is not closed by a LedgerMarker
func RemoveLedgerEntries(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) error {
	// Remove the LedgerEntries and the ledgerallocation entries
	rows, err := RRdb.Prepstmt.GetAllLedgerEntriesInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var l LedgerEntry
		ReadLedgerEntries(rows, &l)
		DeleteLedgerEntry(ctx, l.LEID)
	}
	return err
}

// ledgerCache is a mapping of glNames to ledger structs
var ledgerCache map[string]GLAccount

// InitLedgerCache starts a new ledger cache
func InitLedgerCache() {
	ledgerCache = make(map[string]GLAccount)
}

// GetCachedLedgerByGL checks the cache with index string s. If there is an
// entry there and the BID matches the requested BID we return the ledger
// struct immediately. Otherwise, the ledger is loaded from the database and
// stored in the cache at index s.  If no ledger is found with GLNumber s,
// then a ledger with LID = 0 is returned.
//-----------------------------------------------------------------------------
func GetCachedLedgerByGL(ctx context.Context, bid int64, s string) (GLAccount, error) {
	var (
		l   GLAccount
		ok  bool
		err error
	)

	l, ok = ledgerCache[s]
	if ok {
		if l.BID == bid {
			return l, err
		}
	}

	l, err = GetLedgerByGLNo(ctx, bid, s)
	if err != nil || 0 == l.LID {
		// if you want to log the error then separate the above if clause condition
		Ulog("GetCachedLedgerByGL: error getting ledger %s from business %d. \n", s, bid)
		l.LID = 0
	} else {
		ledgerCache[s] = l
	}
	return l, err
}

// GenerateLedgerEntriesFromJournal creates all the LedgerEntries necessary
// to describe the Journal entry provided. The number of LedgerEntries
// inserted is returned
//-----------------------------------------------------------------------------
func GenerateLedgerEntriesFromJournal(ctx context.Context, xbiz *XBusiness, j *Journal, d1, d2 *time.Time) (int, error) {
	var (
		nr  int
		err error
	)

	for i := 0; i < len(j.JA); i++ {
		m, err := ParseAcctRule(ctx, xbiz, j.JA[i].RID, d1, d2, j.JA[i].AcctRule, j.JA[i].Amount, 1.0)
		if err != nil {
			return nr, err
		}

		for k := 0; k < len(m); k++ {
			var l LedgerEntry
			l.BID = xbiz.P.BID
			l.JID = j.JID
			l.RID = j.JA[i].RID
			l.JAID = j.JA[i].JAID
			l.RAID = j.JA[i].RAID
			l.TCID = j.JA[i].TCID
			l.Dt = j.Dt
			l.Amount = RoundToCent(m[k].Amount)
			if m[k].Action == "c" {
				l.Amount = -l.Amount
			}
			ledger, err := GetCachedLedgerByGL(ctx, l.BID, m[k].Account)
			if err != nil {
				return nr, err
			}

			l.LID = ledger.LID
			if l.Amount >= float64(0.005) || l.Amount < float64(-0.005) { // ignore rounding errors
				dup, err := GetLedgerEntryByJAID(ctx, l.BID, l.LID, l.JAID)
				if err != nil {
					Ulog("GetLedgerEntryByJAID error: %s", err.Error())
					return nr, err
				}

				if dup.LEID == 0 {
					_, err = InsertLedgerEntry(ctx, &l)
					// TODO(Steve): should we return error from here?
					if err != nil {
						Ulog("InsertLedgerEntry error: %s", err.Error())
					}
					nr++
				}
			}
		}
	}
	return nr, err
}

// UpdateRentableLedgerMarkers keeps track of the balance associated with a
// particular Rentable.
//
// INPUTS
//		bid   - business id
//		dt    - save the balance for the subledger on this date
//
// RETURNS
//      error - any error encountered
//-----------------------------------------------------------------------------
func UpdateRentableLedgerMarkers(bid int64, dt *time.Time) error {
	// //----------------------------------------------------------
	// // For each Rentable
	// //----------------------------------------------------------
	// q := fmt.Sprintf("SELECT Rentable.ID, FROM Rentable WHERE BID=%d", RRdb.DBFields["Rentable"], bid)
	// rows, err := RRdb.Dbrr.Query(q)
	// if err != nil {
	// 	return err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var a Rentable
	// 	ReadRentables(rows, &a)

	// 	//----------------------------------------------------------
	// 	// What RAIDs does this Rentable belong to during this
	// 	// period?
	// 	//----------------------------------------------------------

	// 	//----------------------------------------------------------
	// 	// Find prev LedgerMarker for this Rentable.  Create one
	// 	// if one does not already exist.
	// 	//----------------------------------------------------------
	// 	lastLM := GetRARentableLedgerMarkerOnOrBefore(bid, raid, a.RID, dt)
	// 	if lastLM.LMID == 0 {

	// 	}

	// 	//----------------------------------------------------------
	// 	// find all Rental Agreements that this Rentable was a part
	// 	// of since lastLM.Dt
	// 	//----------------------------------------------------------

	// 	//----------------------------------------------------------
	// 	// foreach Rental Agreement...
	// 	//----------------------------------------------------------

	// 	//----------------------------------------------------------
	// 	// Sum all the activity on this Rentable since the last
	// 	// the last LedgerMarker up-to-but-not-including dt.
	// 	//----------------------------------------------------------

	// 	//----------------------------------------------------------
	// 	// Save it and move on...
	// 	//----------------------------------------------------------
	// }
	return nil

}

// UpdateSubLedgerMarkers keeps track of totals per Rental
// Agreement at each LedgerMarker. This was necessary in order to determine
// exactly what each RentalAgreement did with respect to a specific ledger
// account.  The RAID is saved in the LedgerEntry. However, if we don't save
// a total in a LedgerMarker, then we would need to go back to the beginning
// of time and search all LedgerEntries for those that matched a particular
// Rental Agreement.  Instead, we will simply add a LedgerMarker for each
// Rental Agreement that affected a particular account with the total equal to
// its previous balance (if it exists) plus the activity during this period.
//
// If no LedgerMarker is found on or before d1, then one will be created.
//
// A new LedgerMarker will be created at dt with the new balance.
//
// INPUTS
//		bid   - business id
//		dt    - compute the balance for the subledger on this date
//-----------------------------------------------------------------------------
func UpdateSubLedgerMarkers(ctx context.Context, bid int64, d2 *time.Time) error {
	funcname := "UpdateSubLedgerMarkers"
	var lmacct LedgerMarker

	//--------------------------------------------------------------------
	// find the nearest previous ledger marker for any account
	// Its date will be d1, the start time. We'll need to compute all
	// activity that has taken place since that time in order to produce
	// the balance for each ledger marker
	//--------------------------------------------------------------------
	for k := range RRdb.BizTypes[bid].GLAccounts {
		lm, err := GetLedgerMarkerOnOrBefore(ctx, bid, k, d2)
		if err != nil {
			Ulog("%s: Error from GetLedgerMarkerOnOrBefore: %s\n", err.Error())
			continue
		}

		if lm.LID == 0 {
			continue
		}
		lmacct = lm
		break
	}

	lm, _ := GetLedgerMarkerOnOrBefore(ctx, bid, lmacct.LID, d2)
	d1 := &lm.Dt

	//-------------------------------
	// For each Rental Agreement
	//-------------------------------
	rows, err := RRdb.Prepstmt.GetRentalAgreementByBusiness.Query(bid)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ra RentalAgreement
		err = ReadRentalAgreements(rows, &ra)
		if err != nil {
			Ulog("%s: error reading RentalAgreement: %s\n", funcname, err.Error())
			return err
		}

		// Console("%s\n", Tline(80))
		// Console("Processing Rental Agreement RA%08d\n", ra.RAID)

		//---------------------------------------------------------------------
		// get all the ledger activity between d1 and d2 involving the current
		// RentalAgreement
		//---------------------------------------------------------------------
		m, err := GetAllLedgerEntriesForRAID(ctx, d1, d2, ra.RAID)
		if err != nil {
			Ulog("%s: GetLedgerEntriesForRAID returned error: %s\n", funcname, err.Error())
			return err
		}

		// Console("LedgerEntries for RAID = %d between %s - %s:  %d\n", ra.RAID, d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4), len(m))

		LIDprocessed := make(map[int64]int)

		//---------------------------------------------------------------------
		// Spin through all the transactions for this RAID...
		//---------------------------------------------------------------------
		for i := 0; i < len(m); i++ {
			_, processed := LIDprocessed[m[i].LID] // check this ledger for previous processing
			if processed {                         // did we process it?
				continue // yes: move on to the next one
			}
			if m[i].Amount == float64(0) {
				continue // sometimes an entry slips in with a 0 amount, ignore it
			}

			//-----------------------------------------------------------------
			// find the previous LedgerMarker for the GLAccount.  Create one
			// if none exist...
			//-----------------------------------------------------------------
			lm, err := LoadRALedgerMarker(ctx, bid, m[i].LID, m[i].RAID, d1)
			if err != nil {
				return err
			}

			// Console("%s\n", Tline(20))
			// Console("Processing L%08d\n", m[i].LID)
			// Console("LedgerMarker: LM%08d - %10s  Balance: %8.2f\n", lm.LMID, lm.Dt.Format(RRDATEFMT4), lm.Balance)

			//-----------------------------------------------------------------
			// Spin through the rest of the transactions involving m[i].LID
			// and compute the total
			//-----------------------------------------------------------------
			tot := m[i].Amount
			for j := i + 1; j < len(m); j++ {
				if m[j].LID == m[i].LID {
					tot += m[j].Amount
					// Console("\tLE%08d  -  %8.2f\n", m[j].LEID, m[j].Amount)
				}
			}
			LIDprocessed[m[i].LID] = 1 // mark that we've processed this ledger

			//-----------------------------------------------------------------
			// Create a new ledger marker on d2 with the updated total...
			//-----------------------------------------------------------------
			var lm2 LedgerMarker
			lm2.BID = lm.BID
			lm2.LID = lm.LID
			lm2.RAID = lm.RAID
			lm2.Dt = *d2
			lm2.Balance = lm.Balance + tot
			_, err = InsertLedgerMarker(ctx, &lm2) // lm2.LMID is updated if no error
			if err != nil {
				Ulog("%s: InsertLedgerMarker error: %s\n", funcname, err.Error())
				return err
			}
			// Console("LedgerMarker: RAID = %d, Balance = %8.2f\n", lm2.RAID, lm2.Balance)
		}
	}

	return rows.Err()
}

func closeLedgerPeriod(ctx context.Context, xbiz *XBusiness, li *GLAccount, lm *LedgerMarker, dt *time.Time, state int64) error {
	const funcname = "closeLedgerPeriod"
	bal, err := GetRAAccountBalance(ctx, li.BID, li.LID, 0, dt)
	if err != nil {
		return err
	}

	var nlm LedgerMarker
	nlm = *lm
	nlm.Balance = bal
	nlm.Dt = *dt
	nlm.State = state
	nlm.BID = xbiz.P.BID

	// Console("%s : InsertLedgerMarker nlm.State = %d\n",funcname, nlm.State)

	_, err = InsertLedgerMarker(ctx, &nlm) // this is a period close
	if err != nil {
		Ulog("%s: Error while inserting LedgerMarker: %s\n", funcname, err.Error())
		return err
	}
	// Console("%s : After insert: nlm.LMID = %d, nlm.State = %d\n",funcname, nlm.LMID, nlm.State)

	return err
}

// GenerateLedgerMarkers creates all ledgermarkers at d2
func GenerateLedgerMarkers(ctx context.Context, xbiz *XBusiness, d2 *time.Time) error {
	funcname := "GenerateLedgerMarkers"
	// Console("Entered: %s\n", funcname)
	//----------------------------------------------------------------------------------
	// Spin through all ledgers and update the LedgerMarkers with the ending balance...
	//----------------------------------------------------------------------------------
	t, err := GetLedgerList(ctx, xbiz.P.BID) // this list contains the list of all GLAccount numbers
	if err != nil {
		return err
	}

	for i := 0; i < len(t); i++ {
		lm, err := GetLedgerMarkerOnOrBefore(ctx, xbiz.P.BID, t[i].LID, d2)
		if err != nil {
			LogAndPrint("%s: Could not get GLAccount %d (%s) in business %d\n", funcname, t[i].LID, t[i].GLNumber, xbiz.P.BID)
			continue
		}
		if lm.LMID == 0 {
			LogAndPrint("%s: Could not get GLAccount %d (%s) in business %d\n", funcname, t[i].LID, t[i].GLNumber, xbiz.P.BID)
			continue
		}
		// Console("Calling closeLedgerPeriod for LMID = %d, d2 = %s with LMCLOSED\n", lm.LMID, d2.Format(RRDATETIMERPTFMT))
		err = closeLedgerPeriod(ctx, xbiz, &t[i], &lm, d2, LMCLOSED)
		if err != nil {
			return err
		}
	}

	//----------------------------------------------------------------------------------
	// Now we need to update the ledger markers for RAIDs and RIDs and TCIDs
	//----------------------------------------------------------------------------------
	return UpdateSubLedgerMarkers(ctx, xbiz.P.BID, d2)
	//UpdatePayorSubLedgers(xbiz.P.BID, d1, d2)
}

// GenerateLedgerEntries creates ledgers records based on the Journal records over the supplied time range.
func GenerateLedgerEntries(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) (int, error) {
	// Console("Generate Ledger Records: BID=%d, d1 = %s, d2 = %s\n", xbiz.P.BID, d1.Format(RRDATEFMT4), d2.Format(RRDATEFMT4))
	// funcname := "GenerateLedgerEntries"

	nr := 0
	if err := RemoveLedgerEntries(ctx, xbiz, d1, d2); err != nil {
		Ulog("Could not remove existing LedgerEntries from %s to %s. err = %v\n", d1.Format(RRDATEFMT), d2.Format(RRDATEFMT), err)
		return nr, err
	}
	InitLedgerCache()

	//----------------------------------------------------------------------------------
	// Loop through the Journal records for this time period, update all ledgers...
	//----------------------------------------------------------------------------------
	rows, err := RRdb.Prepstmt.GetAllJournalsInRange.Query(xbiz.P.BID, d1, d2)
	if err != nil {
		return nr, err
	}
	defer rows.Close()

	for rows.Next() {
		var j Journal
		if err = ReadJournals(rows, &j); err != nil {
			return nr, err
		}
		if err = GetJournalAllocations(ctx, &j); err != nil {
			return nr, err
		}
		n, err := GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
		if err != nil {
			return nr, err
		}
		nr += n
	}

	if err = rows.Err(); err != nil {
		return nr, err
	}
	if err = GenerateLedgerMarkers(ctx, xbiz, d2); err != nil {
		return nr, err
	}
	return nr, err
}
