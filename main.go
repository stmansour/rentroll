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
	AsmtTypes map[int64]rlib.AssessmentType
	PmtTypes  map[int64]rlib.PaymentType
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
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
}

func intTest(xbiz *rlib.XBusiness, d1, d2 *time.Time) {
	fmt.Printf("INTERNAL TEST\n")
	m := parseAcctRule(xbiz, 1, d1, d2, "d ${DFLTGENRCV} 1000.0, c 40001 ${UMR}, d 41004 ${UMR} ${aval(${DFLTGENRCV})} -", float64(1000), float64(8)/float64(30))

	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d] = %#v\n", i, m[i])
	}
	fmt.Printf("DONE\n")
}

// Dispatch generates the supplied report for all properties
func Dispatch(d1, d2 time.Time, report int64) {
	// s := "SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	s := "SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	// For every business
	for rows.Next() {
		var xbiz rlib.XBusiness
		// read its definition
		rlib.Errcheck(rows.Scan(&xbiz.P.BID, &xbiz.P.Designation,
			/* &xbiz.P.Address, &xbiz.P.Address2, &xbiz.P.City, &xbiz.P.State,
			&xbiz.P.PostalCode, &xbiz.P.Country, &xbiz.P.Phone, */
			&xbiz.P.Name, &xbiz.P.DefaultOccupancyType,
			&xbiz.P.ParkingPermitInUse, &xbiz.P.LastModTime, &xbiz.P.LastModBy))
		// get its info
		rlib.GetXBusiness(xbiz.P.BID, &xbiz)
		if nil == rlib.RRdb.BizTypes[xbiz.P.BID] {
			bt := rlib.BusinessTypes{
				BID:          xbiz.P.BID,
				AsmtTypes:    make(map[int64]*rlib.AssessmentType),
				PmtTypes:     make(map[int64]*rlib.PaymentType),
				DefaultAccts: make(map[int64]*rlib.LedgerMarker),
			}
			rlib.RRdb.BizTypes[xbiz.P.BID] = &bt
		}
		// Gather its chart of accounts
		rlib.GetDefaultLedgerMarkers(xbiz.P.BID)
		// fmt.Printf("Dispatch: After call to GetDefaultLedgerMarkers: rlib.RRdb.BizTypes[%d].DefaultAccts = %#v\n", xbiz.P.BID, rlib.RRdb.BizTypes[xbiz.P.BID].DefaultAccts)

		// and generate the requested report...
		switch report {
		case 1:
			JournalReportText(&xbiz, &d1, &d2)
		case 2:
			LedgerReportText(&xbiz, &d1, &d2)
		case 3:
			intTest(&xbiz, &d1, &d2)
		case 4:
			fmt.Printf("biz csv = %s\n", App.bizfile)
			// AddNewBusiness(*bizptr)
		default:
			// fmt.Printf("Generating Journal Records for %s through %s\n", d1.Format(RRDATEFMT), d2.AddDate(0, 0, -1).Format(RRDATEFMT))
			GenerateJournalRecords(&xbiz, &d1, &d2)
			GenerateLedgerRecords(&xbiz, &d1, &d2)
		}
	}
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

	//  func Date(year int64 , month Month, day, hour, min, sec, nsec int64 , loc *Location) Time
	start := time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	Dispatch(start, stop, App.Report)
}
