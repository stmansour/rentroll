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
	"rentroll/rlib"
	"rentroll/worker"
	"time"
)

// App is the global application structure
var App struct {
	dbdir  *sql.DB   // phonebook db
	dbtws  *sql.DB   // tws db
	DBDir  string    // phonebook database
	DBtws  string    // name of TWS database
	DBUser string    // user for all databases
	Action string    // action to perform
	Dt     time.Time // call worker with this date time
}

var testOwner = string("TWS Basic Tester1")

func readCommandLineArgs() {
	var err error
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbtwsPtr := flag.String("M", "tws", "database name (tws)")
	aptr := flag.String("a", "add", "add, wait, reschedule, or complete a work item")
	dtptr := flag.String("dt", "2018-03-01", "run assessment instance worker with this date")
	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBtws = *dbtwsPtr
	App.DBUser = *dbuPtr
	App.Action = *aptr
	App.Dt, err = rlib.StringToDate(*dtptr)
	if err != nil {
		rlib.LogAndPrintError("readCommandLineArgs", err)
		os.Exit(1)
	}
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
	App.dbtws, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbtws.Close()
	err = App.dbtws.Ping()
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

	rlib.InitDBHelpers(App.dbtws, App.dbdir)
	rlib.SessionInit(15)

	//-----------------------------------------
	// Generate the assessment for March...
	//-----------------------------------------
	ctx := context.Background()
	worker.CreateAsmInstCore(ctx, &App.Dt)
}
