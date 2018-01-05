package rlib

import (
	"context"
	"time"
)

func temporaryGetLTLAR(bid int64) string {
	// "c ${GLGSRENT} _,d ${GLVAC} _"
	// rlib.RRdb.
	// QueryRow("SELECT ClassCode,CoCode,Name,Designation,Description,LastModTime,LastModBy FROM classes WHERE Designation=?",
	var a, b GLAccount
	q := "SELECT " + RRdb.DBFields["GLAccount"] + " FROM GLAccount WHERE Name LIKE \"%rent-not taxable%\""
	row := RRdb.Dbrr.QueryRow(q)
	ReadGLAccount(row, &a)
	if a.LID == 0 {
		q := "SELECT " + RRdb.DBFields["GLAccount"] + " FROM GLAccount WHERE Name LIKE \"%rent-not taxable%\""
		row := RRdb.Dbrr.QueryRow(q)
		ReadGLAccount(row, &a)
		if a.LID == 0 {
			a.GLNumber = "41000" // yes this is a total hack until I can get rid of some old infrastructure, it only applies to old tests though
		}
	}
	q = "SELECT " + RRdb.DBFields["GLAccount"] + " FROM GLAccount WHERE Name LIKE \"%vacancy%\""
	row = RRdb.Dbrr.QueryRow(q)
	ReadGLAccount(row, &b)
	if b.LID == 0 {
		b.GLNumber = "41101" // yes this is a total hack until I can get rid of some old infrastructure, it only applies to old tests though
	}
	return "c " + a.GLNumber + " _,d " + b.GLNumber + " _"
}

// ProcessRentable looks for any time period for which the Rentable has
// no rent assessment during the supplied time range. If it finds any
// vacant time periods, it generates vacancy Journal entries for that period.
// The approach here is to look for all rental agreements that apply to this
// Rentable over the supplied time period. Spin through each of them and
// count the number of days that are covered.  Compare this to the total
// number of days in the period. Then generate a vacancy entry for the time
// period for which no rent was assessed.
// The return value is the number of vacancy records added
//============================================================================================
func ProcessRentable(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time, r *Rentable) (int, error) {
	const funcname = "ProcessRentable"

	var (
		nr  = 0
		err error
	)

	m, err := VacancyDetect(ctx, xbiz, d1, d2, r.RID)
	if err != nil {
		return nr, err
	}

	// fmt.Printf("ProcessRentable: r = %s (%d), period=(%s - %s) len(m) = %d\n", r.Name, r.RID, d1.Format("Jan 2"), d2.Format("Jan 2"), len(m))
	for i := 0; i < len(m); i++ {
		// the umr rate is in cost/accrualDuration. The duration of the VacancyMarkers
		// are in integral multiples of rangeIncDur.  We need to prorate the amount of
		// each entry accordingly
		var j Journal
		j.BID = xbiz.P.BID
		j.Amount = RoundToCent(m[i].Amount)
		// TODO: fix the next line
		j.Dt = m[i].DtStop.AddDate(0, 0, -1) // associated date is period end - 1 proration cycle (or 1 sec if no proration)
		j.Type = JNLTYPEUNAS                 // this is an unassociated entry
		// j.RAID = 0                           // we really mean it, it is unassociated
		j.ID = r.RID             // mark the associated Rentable
		j.Comment = m[i].Comment // this will note consecutive days for vacancy
		// fmt.Printf("ProcessRentable: insert journal entry: %s - %s, %8.2f\n", j.Dt.Format(RRDATEINPFMT), j.Comment, j.Amount)

		// TODO: this check must be more thorough. Examine the RentCycle, query over the time of the rent cycle
		//       to see if other records were generated and handle.
		//       By convention, the period for Vacancy detection is:  supplied range, date/time of journal entry =
		//       rentcycle - 1 prorationcycle (or one second whichever is larger)
		// These entries must be idempotent. Make sure it does not already exist.
		jv, err := GetJournalVacancy(ctx, r.RID, &j.Dt, &m[i].DtStop)
		if err != nil {
			return nr, err
		}

		if jv.JID != 0 { // if the JID >0 ..
			continue // then this entry was already generated, keep going
		}

		jid, err := InsertJournal(ctx, &j)
		if err != nil {
			Errlog(err)
			return nr, err
		}

		nr++
		if jid > 0 {
			var ja JournalAllocation
			ja.JID = jid
			ja.Amount = j.Amount
			ja.ASMID = 0 // it's unassociated

			ja.AcctRule = temporaryGetLTLAR(xbiz.P.BID)
			ja.RID = r.RID
			ja.BID = r.BID
			// fmt.Printf("VACANCY: inserting journalAllocation entry: %#v\n", ja)
			_, err := InsertJournalAllocationEntry(ctx, &ja)
			if err != nil {
				Ulog("%s: Error while inserting journalAllocation entry: %s\n", funcname, err.Error())
				// TODO(Steve): ignore error?
			}
			j.JA = append(j.JA, ja)
		}
		InitLedgerCache()
		_, err = GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
		if err != nil {
			return nr, err
		}
	}

	return nr, err
}

// GenVacancyJournals creates Journal entries that cover vacancy for
// every Rentable where the Rentable type is being managed to budget
//===============================================================================================
func GenVacancyJournals(ctx context.Context, xbiz *XBusiness, d1, d2 *time.Time) (int, error) {

	var (
		err error
		nr  = 0
	)

	rows, err := RRdb.Prepstmt.GetAllRentablesByBusiness.Query(xbiz.P.BID)
	if err != nil {
		return nr, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Rentable
		err = ReadRentables(rows, &r)
		if err != nil {
			return nr, err
		}

		b, err := ProcessRentable(ctx, xbiz, d1, d2, &r)
		if err != nil {
			return nr, err
		}

		nr += b
	}

	return nr, rows.Err()
}
