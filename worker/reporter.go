package worker

import (
	"context"
	"database/sql"
	"rentroll/rlib"
	"strings"
	"time"
	"tws"
)

// CreateTLReporterInstances is a worker that is called by TWS periodically to
// check for recurring assessments that have instances needing to be created.
// When their instance date arrives, this routine will generate the new instance.
// After generating all instances whose time has arrived it will reschedule itself
// to be called again the next day.
//-----------------------------------------------------------------------------
func CreateTLReporterInstances(item *tws.Item) {
	tws.ItemWorking(item)
	now := time.Now()
	ctx := context.Background()
	TLReporterCore(ctx)

	// reschedule for midnight tomorrow...
	resched := now.Add(1 * time.Minute)
	tws.RescheduleItem(item, resched)
}

// TLReporterCore provides a more testable calling routine for processing
// Task Lists.  This routine checks all active task lists and emails a
// TaskList report to the defined email addresses if any of the following
// conditions exist:
//
// 1. The TaskList has a PreDue date and that date has past
// 2. The TaskList has a Due date and that date has past
//
//-----------------------------------------------------------------------------
func TLReporterCore(ctx context.Context) error {
	var err error
	var rows *sql.Rows
	now := time.Now()

	if tx, ok := rlib.DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(rlib.RRdb.Prepstmt.GetDueTaskLists)
		defer stmt.Close()
		rows, err = stmt.Query(now, now)
	} else {
		rows, err = rlib.RRdb.Prepstmt.GetDueTaskLists.Query(now, now)
	}

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var a rlib.TaskList
		if err = rlib.ReadTaskLists(rows, &a); err != nil {
			return err
		}
		rlib.Console("Found: %s,  TLID = %d\n", a.Name, a.TLID)
		s := strings.TrimSpace(a.EmailList)
		if len(s) == 0 {
			rlib.Console("EmailList is blank. Skipping.\n")
			continue
		}
		sa := strings.Split(s, ",")
		for i := 0; i < len(sa); i++ {
			to := strings.TrimSpace(sa[i])
			TLReporterSendEmail(to,a.TLID)
		}
	}

	return rows.Err()
}

// TLReporterSendEmail sends an email message to recipient e with
// an attachment containing the PDF of a tasklist report for
// the TaskList with TLID = tlid
//
// INPUTS:
//  e    - email address
//  tlid - the TLID of the TaskList for the report
//
// RETURNS:
//  any errors encountered
//-----------------------------------------------------------------------------
func TLReporterSendEmail(e string, tlid int64) error {
	rlib.Console("send email to: %s\n",e)
	return nil
}