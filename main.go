package main

import (
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"phonebook/lib"
	"rentroll/rcsv"
	"rentroll/rlib"
	"rentroll/worker"
	"rentroll/ws"
	"strings"
	"time"
	"tws"

	_ "github.com/go-sql-driver/mysql"
)

// DispatchCtx is a type of struct needed for the Dispatch function. It defines
// everything needed to run a particular command. It is the responsibility of the
// caller to fill out all the needed ctx information. Not all information is needed
// for all commands.
type DispatchCtx struct {
	Cmd          int                 // cmd to execute
	DtStart      time.Time           // period start time
	DtStop       time.Time           // period end
	xbiz         rlib.XBusiness      // BUD from cmd line
	OutputFormat int                 // how shall the output be formatted
	Report       int64               // which report to generate - this is used in batch mode operation
	Args         string              // full command line report string
	CSVLoadStr   string              // if loading CSV file, this will have index,filename
	w            http.ResponseWriter // for web responses
	r            *http.Request       // for web request introspection
}

// App is the global data structure for this app
var App struct {
	dbdir        *sql.DB  // phonebook db
	dbrr         *sql.DB  // rentroll db
	DBDir        string   // phonebook database
	PortRR       int      // port on which rentroll listens
	DBUser       string   // user for all databases
	Report       string   // if testing engine, which report/action to perform
	LogFile      *os.File // where to log messages
	BatchMode    bool     // if true, then don't start http, the command line request is for a batch process
	SkipVacCheck bool     // until the code is modified to process on each command entered, if set to false, this inibits batch processing to do vacancy calc.
	CSVLoad      string   // if loading csv, this string will have index,filename
	sStart       string   // start time
	sStop        string   // stop time
	Bud          string   // BUD from the command line
	CertFile     string   // public certificate
	KeyFile      string   //private key file
	//DBRR         string   // rentroll database
}

// RRfuncMap is a map of functions passed to each html page that can be referenced
// as needed to produce the page
var RRfuncMap map[string]interface{}

// Chttp is a server mux for handling unprocessed html page requests.
// For example, a .css file or an image file.
var Chttp = http.NewServeMux()

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	//dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pStart := flag.String("j", "2015-11-01", "Accounting Period start time")
	pStop := flag.String("k", "2015-12-01", "Accounting Period end time")
	pKey := flag.String("K", "localhost.key", "Private key file")
	pCert := flag.String("C", "localhost.crt", "Cert file")
	pBud := flag.String("b", "", "Business Unit Identifier (BUD)")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	rptPtr := flag.String("r", "0", "report: 0 = generate Journal records, 1 = Journal, 2 = Rentable, 4=Rentroll, 5=AssessmentCheck, 6=LedgerBalance, 7=RentableCountByType, 8=Statement, 9=Invoice, 10=LedgerActivity, 11=RentableGSR, 12-RALedgerBalanceOnDate,LID,RAID,Date, 13-RAAcctActivity,LID,RAID, 14,Date=delinqRpt")
	pLoad := flag.String("L", "", "CSV Load index,filename")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	bPtr := flag.Bool("A", false, "if specified run as a batch process, do not start http")
	xPtr := flag.Bool("x", false, "if specified, inhibit vacancy checking")

	flag.Parse()
	if *verPtr {
		fmt.Printf("Version:    %s\nBuild Time: %s\nBuild Machine: %s\n", ws.GetVersionNo(), ws.GetBuildTime(), ws.GetBuildMachine())
		os.Exit(0)
	}
	App.DBDir = *dbnmPtr
	//App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.Report = *rptPtr
	App.sStart = *pStart
	App.sStop = *pStop
	App.PortRR = *portPtr
	App.BatchMode = *bPtr
	App.Bud = *pBud
	App.SkipVacCheck = *xPtr
	App.CertFile = *pCert
	App.KeyFile = *pKey
	// fmt.Printf("*pLoad = %s\n", *pLoad)
	App.CSVLoad = *pLoad
}

func intTest(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	fmt.Printf("INTERNAL TEST\n")
	m := rlib.ParseAcctRule(xbiz, 1, d1, d2, "d ${GLGENRCV} 1000.0, c 40001 ${UMR}, d 41004 ${UMR} ${aval(${GLGENRCV})} -", float64(1000), float64(8)/float64(30))

	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d] = %#v\n", i, m[i])
	}
	fmt.Printf("DONE\n")
}

// HomeHandler serves static http content such as the .css files
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		Chttp.ServeHTTP(w, r)
	} else {
		http.Redirect(w, r, "/home/", http.StatusFound)
	}
}

func initHTTP() {
	Chttp.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/home/", HomeUIHandler)
	http.HandleFunc("/v1/", ws.V1ServiceHandler)
	http.HandleFunc("/wsvc/", webServiceHandler)
}

func main() {
	var err error
	readCommandLineArgs()
	// fmt.Printf("App.CSVLoad = %s\n", App.CSVLoad)
	//==============================================
	// Open the logfile and begin logging...
	//==============================================
	App.LogFile, err = os.OpenFile("rentroll.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	lib.Errcheck(err)
	defer App.LogFile.Close()
	log.SetOutput(App.LogFile)
	rlib.Ulog("*** Accord RENTROLL ***\n")

	//----------------------------
	// Open RentRoll database
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
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
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
	initRentRoll()

	if App.BatchMode {
		ctx := createStartupCtx()
		rcsv.InitRCSV(&ctx.DtStart, &ctx.DtStop, &ctx.xbiz)
		RunCommandLine(&ctx)
	} else {
		tws.Init(rlib.RRdb.Dbrr, rlib.RRdb.Dbdir) // starts the scheduler in a go routine. only initialize when we're in server mode
		worker.Init()                             // register Rentroll's TWS workers
		initHTTP()                                // identify the handlers
		rlib.Ulog("RentRoll initiating HTTP service on port %d and HTTPS on port %d\n", App.PortRR, App.PortRR+1)

		go http.ListenAndServeTLS(fmt.Sprintf(":%d", App.PortRR+1), App.CertFile, App.KeyFile, nil)
		err = http.ListenAndServe(fmt.Sprintf(":%d", App.PortRR), nil)
		if nil != err {
			fmt.Printf("*** Error on http.ListenAndServe: %v\n", err)
			rlib.Ulog("*** Error on http.ListenAndServe: %v\n", err)
			os.Exit(1)
		}
	}
}
