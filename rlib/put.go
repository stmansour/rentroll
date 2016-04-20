package rlib

import "fmt"

// InsertBusiness writes a new Business record.
// returns the new business ID and any associated error
func InsertBusiness(b *Business) (int64, error) {
	var bid = int64(0)
	res, err := RRdb.Prepstmt.InsertBusiness.Exec(b.Designation, b.Name, b.DefaultOccupancyType, b.ParkingPermitInUse, b.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			bid = int64(id)
		}
	}
	return bid, err
}

// InsertJournalEntry writes a new journal entry to the database
func InsertJournalEntry(j *Journal) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertJournal.Exec(j.BID, j.RAID, j.Dt, j.Amount, j.Type, j.ID, j.Comment, j.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	}
	return rid, err
}

// InsertJournalAllocationEntry writes a new journalallocation record to the database
func InsertJournalAllocationEntry(ja *JournalAllocation) error {
	_, err := RRdb.Prepstmt.InsertJournalAllocation.Exec(ja.JID, ja.RID, ja.Amount, ja.ASMID, ja.AcctRule)
	return err
}

// InsertJournalMarker writes a new journalmarker record to the database
func InsertJournalMarker(jm *JournalMarker) error {
	_, err := RRdb.Prepstmt.InsertJournalMarker.Exec(jm.BID, jm.State, jm.DtStart, jm.DtStop)
	return err
}

// InsertLedgerMarker writes a new journalmarker record to the database
func InsertLedgerMarker(l *LedgerMarker) error {
	_, err := RRdb.Prepstmt.InsertLedgerMarker.Exec(l.BID, l.PID, l.GLNumber, l.Status, l.State, l.DtStart, l.DtStop, l.Balance, l.Type, l.Name)
	if err != nil {
		fmt.Printf("InsertLedgerMarker: err = %#v\n", err)
	}
	return err
}

// InsertLedgerEntry writes a new journal entry to the database
func InsertLedgerEntry(l *Ledger) (int64, error) {
	var rid = int64(0)
	res, err := RRdb.Prepstmt.InsertLedger.Exec(l.BID, l.JID, l.JAID, l.GLNumber, l.Dt, l.Amount, l.Comment, l.LastModBy)
	if nil == err {
		id, err := res.LastInsertId()
		if err == nil {
			rid = int64(id)
		}
	} else {
		Ulog("Error inserting ledger:  %v\n", err)
	}
	return rid, err
}
