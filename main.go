package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"phonebook/lib"
	"rentroll/rlib"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

// CmdRUNBOOKS and the rest are command numbers used by the Dispatch function.
const (
	CmdRUNBOOKS = 1 // Run journal and ledgers over a defined period
)

// DispatchCtx is a type of struct needed for the Dispatch function. It defines
// everything needed to run a particular command. It is the responsibility of the
// caller to fill out all the needed ctx information.
type DispatchCtx struct {
	Cmd     int
	DtStart time.Time
	DtStop  time.Time
	Report  int64
}

// App is the global data structure for this app
var App struct {
	dbdir     *sql.DB  // phonebook db
	dbrr      *sql.DB  //rentroll db
	DBDir     string   // phonebook database
	DBRR      string   //rentroll database
	DBUser    string   // user for all databases
	Report    int64    // if testing engine, which report/action to perform
	bizfile   string   // TEMPORARY - tests loading bizcsv
	LogFile   *os.File // where to log messages
	sStart    string   //start time
	sStop     string   //stop time
	AsmtTypes map[int64]rlib.AssessmentType
	PmtTypes  map[int64]rlib.PaymentType
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pStart := flag.String("j", "2015-11-01", "Accounting Period start time")
	pStop := flag.String("k", "2015-12-01", "Accounting Period end time")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	bizPtr := flag.String("b", "b.csv", "add business via csv file")
	rptPtr := flag.Int64("r", 0, "report: 0 = generate journal records, 1 = journal, 2 = rentable")

	flag.Parse()
	if *verPtr {
		fmt.Printf("Version:    %s\nBuild Time: %s\n", getVersionNo(), getBuildTime())
		os.Exit(0)
	}
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.Report = *rptPtr
	App.bizfile = *bizPtr
	App.sStart = *pStart
	App.sStop = *pStop
}

func intTest(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	fmt.Printf("INTERNAL TEST\n")
	m := rlib.ParseAcctRule(xbiz, 1, d1, d2, "d ${DFLTGENRCV} 1000.0, c 40001 ${UMR}, d 41004 ${UMR} ${aval(${DFLTGENRCV})} -", float64(1000), float64(8)/float64(30))

	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d] = %#v\n", i, m[i])
	}
	fmt.Printf("DONE\n")
}

func main() {
	readCommandLineArgs()
	//==============================================
	// Open the logfile and begin logging...
	//==============================================
	var err error
	App.LogFile, err = os.OpenFile("rentroll.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	lib.Errcheck(err)
	defer App.LogFile.Close()
	log.SetOutput(App.LogFile)
	rlib.Ulog("*** Accord RENTROLL ***\n")

	//----------------------------
	// Open RentRoll database
	//----------------------------
	// s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", DBUser, DBRR)
	s := rlib.RRGetSQLOpenString(App.DBUser, App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = lib.GetSQLOpenString(App.DBUser, App.DBDir)
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
	initRentRoll()

	ctx := createStartupCtx()
	Dispatch(&ctx)
}
