package main

//=============================================================================
// Basic test program for tws
//=============================================================================

import (
	"context"
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"os"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"rentroll/worker"
	"rentroll/ws"
	"time"
)

// App is the global application structure
var App struct {
	dbdir  *sql.DB  // phonebook db
	dbrr   *sql.DB  // tws db
	DBDir  string   // phonebook database
	DBtws  string   // name of TWS database
	DBUser string   // user for all databases
	Action string   // action to perform
	NoAuth bool     // if true then skip authentication
	Idx    int      // which test index
	Comm   chan int //
	Email  string   // comma separated list of email addresses
}

var testOwner = string("TWS Basic Tester1")

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "tws", "database name (tws)")
	email := flag.String("email", "", "comma separated list of email addresses")
	aptr := flag.String("a", "add", "add, wait, reschedule, or complete a work item")
	noauth := flag.Bool("noauth", false, "run server in no-auth mode")
	flag.Parse()

	App.DBDir = *dbnmPtr
	App.NoAuth = *noauth
	App.DBtws = *dbrrPtr
	App.DBUser = *dbuPtr
	App.Action = *aptr
	App.Email = *email
}

func main() {
	var err error
	readCommandLineArgs()
	// App.NoAuth = true // for now, let's just always do noauth
	rlib.RRReadConfig()

	//----------------------------
	// Open database
	//----------------------------
	if err = rlib.RRReadConfig(); err != nil {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	s := extres.GetSQLOpenString(rlib.AppConfig.RRDbname, &rlib.AppConfig)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBtws.Ping for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = extres.GetSQLOpenString(rlib.AppConfig.Dbname, &rlib.AppConfig)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open: Error = %v\n", err)
		os.Exit(1)
	}
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("dbdir.Ping: Error = %v\n", err)
		os.Exit(1)
	}

	rlib.InitDBHelpers(App.dbrr, App.dbdir)
	rlib.RpnInit()
	bizlogic.InitBizLogic()
	ws.InitReports()
	rlib.SetAuthFlag(App.NoAuth)
	ws.SvcInit(App.NoAuth) // currently needed for testing
	// tws.Init(App.dbrr, App.dbdir) //
	// worker.Init()              // don't init these, it introduces randomness
	rlib.SessionInit(15) //

	if len(App.Email) > 0 {
		bl, err := rlib.GetAllBiz()
		if err != nil {
			fmt.Printf("Error getting businesses: %s\n", err.Error())
			os.Exit(1)
		}
		if len(bl) == 0 {
			fmt.Printf("No businesses in the database!\n")
			os.Exit(1)
		}
		TaskEmail(bl[0].BID, App.Email)
		return // this main function is now complete
	}

	rlib.Console("calling doWork()\n")
	doWork()
}

func doWork() {
	now := time.Now()
	expire := now.Add(1 * time.Minute)
	s := rlib.SessionNew("BotToken-"+worker.TLReportBotDes, worker.TLReportBotDes, worker.TLReportBotDes, worker.TLReportBot, "", -1, &expire)
	ctx := context.Background()
	ctx = rlib.SetSessionContextKey(ctx, s)

	//---------------------------------------------------
	// Create instances for June 1, 2018 midnight PDT
	//---------------------------------------------------
	dt, _ := rlib.StringToDate("2018-06-01 07:00:00 UTC")
	rlib.Console("Calling TLInstanceBotCore with date = %s\n", dt.Format(rlib.RRDATETIMERPTFMT))
	rlib.Console("                               date = %s\n", dt.In(rlib.RRdb.Zone).Format(rlib.RRDATETIMERPTFMT))
	err := worker.TLInstanceBotCore(ctx, &dt)
	if err != nil {
		rlib.Console("TLInstanceBotCore returns error = %s\n", err.Error())
	}

	//----------------------------------------------------------
	// call it a second time to ensure that it is idempotent
	//----------------------------------------------------------
	rlib.Console("Calling TLInstanceBotCore a second time with the same info to check idempotence\n")
	err = worker.TLInstanceBotCore(ctx, &dt)
	if err != nil {
		rlib.Console("TLInstanceBotCore returns error = %s\n", err.Error())
	}

	rlib.Console("Finished!\n")
}

// TaskEmail is used to test the email capability of a late task.  This
// routine will create a non-recurring tasklist that is due in 2 mins.
// when its time arrives it will generate the mail and send it to the
// supplied email address string.
//
// INPUTS:
//     bid - which business
//      ea - comma separated list of email addresses
//-----------------------------------------------------------------------------
func TaskEmail(bid int64, ea string) {
	funcname := "TaskEmail"
	now := time.Now()
	expire := now.Add(1 * time.Minute)
	s := rlib.SessionNew("BotToken-"+worker.TLReportBotDes, worker.TLReportBotDes, worker.TLReportBotDes, worker.TLReportBot, "", -1, &expire)
	ctx := context.Background()
	ctx = rlib.SetSessionContextKey(ctx, s)

	var tldef rlib.TaskListDefinition
	tldef.BID = bid
	tldef.Cycle = rlib.RECURNONE
	tldef.Name = "Test TaskList"
	tldef.EpochDue = now.Add(-1 * time.Minute)
	tldef.EpochPreDue = now.Add(-1 * time.Hour)
	tldef.EmailList = ea
	tldef.FLAGS = 0x2 | 0x4 // this tasklist has both a due date and a pre-due date

	err := rlib.InsertTaskListDefinition(ctx, &tldef)
	if err != nil {
		rlib.LogAndPrint("rlib.InsertTaskListDefinition: error = %s\n", err.Error())
		return
	}

	var t = []rlib.TaskDescriptor{
		{Name: "Make A TaskList", Worker: "Manual", EpochDue: tldef.EpochDue, EpochPreDue: tldef.EpochPreDue},
		{Name: "Mark 'Make A Tasklist' compete", Worker: "Manual", EpochDue: tldef.EpochDue, EpochPreDue: tldef.EpochPreDue},
		{Name: "Mark the second task completed", Worker: "Manual", EpochDue: tldef.EpochDue, EpochPreDue: tldef.EpochPreDue},
		{Name: "Mark entire TaskList complete", Worker: "Manual", EpochDue: tldef.EpochDue, EpochPreDue: tldef.EpochPreDue},
	}

	for i := 0; i < len(t); i++ {
		t[i].TLDID = tldef.TLDID
		err := rlib.InsertTaskDescriptor(ctx, &t[i])
		if err != nil {
			rlib.LogAndPrint("rlib.InsertTaskDescriptor: error = %s\n", err.Error())
			return
		}
	}

	//----------------------------------------------
	// Now, create an instance of this task list.
	//----------------------------------------------
	tlid, err := rlib.CreateTaskListInstance(ctx, tldef.TLDID, 0, &now)
	if err != nil {
		rlib.LogAndPrint("CreateTaskListInstance:  error = %s\n", err.Error())
		return
	}

	//----------------------------------------------
	// Now, create an instance of this task list.
	//----------------------------------------------
	tl, err := rlib.GetTaskList(ctx, tlid)
	if err != nil {
		rlib.LogAndPrint("rlib.InsertTaskDescriptor: error = %s\n", err.Error())
		return
	}
	rlib.Console("Created tasklist %d\n", tl.TLID)
	m, err := rlib.GetTasks(ctx, tl.TLID)
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return
	}
	m[0].DtPreDone = now.Add(-30 * time.Minute)
	m[0].DtDone = now.Add(1 * time.Minute)
	err = rlib.UpdateTask(ctx, &m[0])
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return
	}

	m[1].DtPreDone = now.Add(-15 * time.Minute)
	m[1].DtDone = now.Add(-1 * time.Minute)
	err = rlib.UpdateTask(ctx, &m[1])
	if err != nil {
		rlib.LogAndPrintError(funcname, err)
		return
	}

}
