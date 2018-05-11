package worker

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"gotable"
	"html/template"
	"os"
	"rentroll/rlib"
	"rentroll/rrpt"
	"strings"
	"time"
	"tws"

	"gopkg.in/gomail.v2"
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

	//---------------------------------------------------
	// Set up email dialer to use for any overdue tasks
	//---------------------------------------------------
	d := gomail.NewDialer(rlib.AppConfig.SMTPHost, rlib.AppConfig.SMTPPort, rlib.AppConfig.SMTPLogin, rlib.AppConfig.SMTPPass)

	//---------------------------------------------------
	// Set up email dialer to use for any overdue tasks
	//---------------------------------------------------
	var ri = rrpt.ReporterInfo{
		OutputFormat: gotable.TABLEOUTPDF,
		Bid:          0,   // don't know this yet
		D1:           now, // don't really need this
		D2:           now, // don't really need this
		BlankLineAfterRptName: false,
	}

	if tx, ok := rlib.DBTxFromContext(ctx); ok { // if transaction is supplied
		stmt := tx.Stmt(rlib.RRdb.Prepstmt.GetDueTaskLists)
		defer stmt.Close()
		rows, err = stmt.Query(now, now)
	} else {
		rows, err = rlib.RRdb.Prepstmt.GetDueTaskLists.Query(now, now, now)
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

		//----------------------------------------------------
		// TODO:  really need to cache xbiz or see if we can
		//        simply not set it at all. It should not be
		//        necessary for this report.
		//----------------------------------------------------
		ri.Bid = a.BID
		ri.ID = a.TLID
		rlib.Console("ri.BID = %d, ri.ID = %d\n", ri.Bid, ri.ID)
		var xbiz rlib.XBusiness
		if err = rlib.GetXBiz(ri.Bid, &xbiz); err != nil {
			return err
		}
		ri.Xbiz = &xbiz

		sa := strings.Split(s, ",")
		for i := 0; i < len(sa); i++ {
			to := strings.TrimSpace(sa[i])
			if rlib.ValidEmailAddress(to) {
				TLReporterSendEmail(ctx, to, &a, d, &ri)
			}
		}
	}

	return rows.Err()
}

// TLReportEmail encapsulates the data needed to fill out the
// report email template.
type TLReportEmail struct {
	TLName       string
	TLID         int64
	DtDue        string
	DtPreDue     string
	DtDone       string
	DtPreDone    string
	DtNextNotify string
	BotName      string
}

// TLReporterSendEmail sends an email message to recipient e with
// an attachment containing the PDF of a tasklist report for
// the TaskList with a.TLID
//
// INPUTS:
//  ctx  - ctx that may hold db txn
//  e    - email address
//  a    - the database TaskList struct
//  d    - email dialer for use if email needs to be sent
//  ri   - info needed by report generator
//
// RETURNS:
//  any errors encountered
//-----------------------------------------------------------------------------
func TLReporterSendEmail(ctx context.Context, e string, a *rlib.TaskList, d *gomail.Dialer, ri *rrpt.ReporterInfo) error {
	rlib.Console("send email to: %s\n", e)
	now := time.Now()
	n := now.Add(a.DurWait)

	//----------------------------------
	// Create message...
	//----------------------------------
	data := TLReportEmail{
		TLName:       a.Name,
		TLID:         a.TLID,
		DtDue:        a.DtDue.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT),
		DtPreDue:     a.DtPreDue.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT),
		DtDone:       a.DtDone.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT),
		DtPreDone:    a.DtPreDone.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT),
		DtNextNotify: n.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT),
		BotName:      "ReportBot",
	}

	btmpl := `<html><body><p>TaskList {{.TLName}} was due on {{.DtDue}} and has not yet been 
marked as completed. See the attached report for further details.
</p><p>
Next check for this task list: {{.DtNextNotify}}
</p><p>
-{{.BotName}}</p></body></html>`

	b := &bytes.Buffer{}
	template.Must(template.New("").Parse(btmpl)).Execute(b, data)
	subj := fmt.Sprintf("Status:  %s tasks are not complete", a.Name)
	rlib.Console("Subject: %s\n", subj)
	rlib.Console("Message Body: %s\n", b.String())

	//----------------------------------
	// Generate report for attachment
	//----------------------------------
	tbl := rrpt.TaskListReportTable(ctx, ri)
	tsh := rrpt.SingleTableReportHandler{
		ReportNames:             []string{"RPTtl", "task list"},
		TableHandler:            rrpt.TaskListReportTable,
		PDFprops:                nil,
		HTMLTemplate:            "",
		NeedsCustomPDFDimension: true,
		NeedsPDFTitle:           true,
	}
	rctx := rrpt.ReportContext{
		PDFPageSizeUnit: "in",
		PDFPageWidth:    float64(8.5),
		PDFPageHeight:   float64(11),
	}
	rptFileName := "TaskListReport.pdf"

	file, err := os.Create(rptFileName)
	if err != nil {
		rlib.Ulog("Cannot create file: %s\n", err.Error())
		return err
	}
	defer file.Close()

	if err = rrpt.WritePDFReport(file, &tsh, ri, &rctx, &tbl); err != nil {
		return err
	}
	file.Close()

	//----------------------------------
	// Send message...
	//----------------------------------
	msg := gomail.NewMessage()
	msg.SetHeader("From", "sman@accordinterests.com")
	msg.SetHeader("To", e)
	msg.SetHeader("Subject", subj)
	msg.SetBody("text/html", b.String())
	msg.Attach(rptFileName)
	if err := d.DialAndSend(msg); err != nil {
		return err
	}

	//----------------------------------
	// Remove the report file...
	//----------------------------------
	if err = os.Remove(rptFileName); err != nil {
		return err
	}

	//----------------------------------
	// Update TaskList ...
	//----------------------------------
	a.DtLastNotify = time.Now()
	a.FLAGS |= 1 << 5

	return rlib.UpdateTaskList(ctx, a)
}
