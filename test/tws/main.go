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
	"tws"
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
}

var testOwner = string("TWS Basic Tester1")

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "tws", "database name (tws)")
	aptr := flag.String("a", "add", "add, wait, reschedule, or complete a work item")
	noauth := flag.Bool("noauth", false, "run server in no-auth mode")
	flag.Parse()

	App.DBDir = *dbnmPtr
	App.NoAuth = *noauth
	App.DBtws = *dbrrPtr
	App.DBUser = *dbuPtr
	App.Action = *aptr
}

func main() {
	var err error
	readCommandLineArgs()
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
	ws.SvcInit(App.NoAuth)        // currently needed for testing
	tws.Init(App.dbrr, App.dbdir) //
	// worker.Init()              // don't init these, it introduces randomness
	rlib.SessionInit(15) //
	rlib.Console("calling doWork()\n")
	doWork()
}

var wrk = map[string]worker.Worker{
	"TWSAsmtBotChecker": {"TWSAsmtBotChecker", "TWS Test Worker", -9999, uint64(0), TWSAsmtBotChecker},
}

var checkTimes = []time.Time{
	time.Date(2018, time.March, 1, 0, 0, 0, 0, time.UTC),
	time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC),
}

func doWork() {
	rlib.Console("Entered doWork\n")
	//------------------------------------------------
	// 1. Register a new handler
	//------------------------------------------------
	rlib.Console("Calling InitCore\n")
	worker.InitCore(wrk)

	//------------------------------------------------
	// 2. Make a channel for the worker thread to
	//    communicate when it is finished
	//------------------------------------------------
	App.Comm = make(chan int)

	//------------------------------------------------
	// 3. wait for us to get notification of completion
	//------------------------------------------------
	rlib.Console("Waiting for signal from workers...\n")
	done := <-App.Comm
	rlib.Console("Received completion signal. done = %d\n", done)
}

// TWSAsmtBotChecker is a checker function.
//-----------------------------------------------------------------------------
func TWSAsmtBotChecker(item *tws.Item) {
	rlib.Console("\n\n------------------------------------------------\n")
	rlib.Console("Entered TWSAsmtBotChecker.  App.Idx = %d, len(checkTimes) = %d\n", App.Idx, len(checkTimes))
	tws.ItemWorking(item)
	now := time.Now().In(rlib.RRdb.Zone)
	ctx := context.Background()
	rlib.Console("Created ctx\n")

	if App.Idx < len(checkTimes) {
		rlib.Console("Calling CreateAsmInstCore\n")
		worker.CreateAsmInstCore(ctx, &checkTimes[App.Idx])
	}

	App.Idx++
	rlib.Console("Incr App.Idx.  new value: %d\n", App.Idx)
	if App.Idx < len(checkTimes) {
		rlib.Console("Rescheduled another pass\n")
		tws.RescheduleItem(item, now)
	} else {
		rlib.Console("Done, signaling completion\n")
		App.Comm <- 1 // signal completion
	}
}
