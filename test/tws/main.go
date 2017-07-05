package main

//=============================================================================
// Basic test program for tws
//=============================================================================

import (
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"os"
	"rentroll/rlib"
	"time"
	"tws"
)

// App is the global application structure
var App struct {
	dbdir  *sql.DB // phonebook db
	dbtws  *sql.DB // tws db
	DBDir  string  // phonebook database
	DBtws  string  // name of TWS database
	DBUser string  // user for all databases
	Action string  // action to perform
}

var testOwner = string("TWS Basic Tester1")

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbtwsPtr := flag.String("M", "tws", "database name (tws)")
	aptr := flag.String("a", "add", "add, wait, reschedule, or complete a work item")
	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBtws = *dbtwsPtr
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

	rlib.RpnInit()
	tws.Init(App.dbtws, App.dbdir)
	fmt.Printf("Successfully opened databases\n")
	fmt.Printf("Node = %s\n", tws.TWSctx.Node)
	tws.RegisterWorker("WorkHandler", WorkHandler)

	doWork()
}

// WorkHandler is a test work handling function
func WorkHandler(item *tws.Item) {
	tws.ItemWorking(item)
	fmt.Printf("I am work handler %s\n", item.Owner)
	os.Exit(0)
}

func doWork() {
	fmt.Printf("Entered doWork()\n")
	switch App.Action {
	case "add":
		fmt.Printf("\tdoWork->add\n")
		item := tws.Item{
			Owner:        testOwner,
			OwnerData:    "I have no data",
			WorkerName:   "WorkHandler",
			ActivateTime: time.Now(),
		}
		tws.InsertItem(&item)
	case "wait":
		fmt.Printf("\tdoWork->wait\n")
		select {
		case <-time.After(15 * time.Second):
			fmt.Printf("15 seconds have elapsed and WorkHandler was not called\n")
			os.Exit(1)
		}
	case "reschedule":
		fmt.Printf("\tdoWork->reschedule\n")
		item := findItem(testOwner)
		fmt.Printf("Rescheduling for 5 seconds from now\n")
		tws.RescheduleItem(&item, time.Now().Add(5*time.Second))
	case "complete":
		fmt.Printf("\tdoWork->complete\n")
		item := findItem(testOwner)
		tws.CompleteItem(&item)
	}
}

func findItem(o string) tws.Item {
	fmt.Printf("Entered findItem(%s)\n", o)
	m, err := tws.FindItem(o)
	if err != nil {
		fmt.Printf("Error with FindItem = %s\n", err.Error())
		os.Exit(1)
	}
	if len(m) == 0 {
		fmt.Printf("Could not find any items for Owner = %s\n", o)
		os.Exit(1)
	}
	fmt.Printf("Found Item: Owner = %s, OwnerData = %s, WorkerName = %s\n", m[0].Owner, m[0].OwnerData, m[0].WorkerName)
	return m[0]
}
