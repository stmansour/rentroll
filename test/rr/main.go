// The purpose of this test is to validate the db update routines.
package main

import (
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"os"
	"rentroll/rlib"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir  *sql.DB        // phonebook db
	dbrr   *sql.DB        //rentroll db
	DBDir  string         // phonebook database
	DBRR   string         //rentroll database
	DBUser string         // user for all databases
	PortRR int            // rentroll port
	Bud    string         // Biz Unit Descriptor
	Xbiz   rlib.XBusiness // lots of info about this biz
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
}

func main() {
	var err error
	readCommandLineArgs()
	rlib.RRReadConfig()

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

	rlib.RpnInit()
	rlib.InitDBHelpers(App.dbrr, App.dbdir)

	biz := rlib.GetBusinessByDesignation(App.Bud)
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}
	rlib.InitBizInternals(biz.BID, &App.Xbiz)

	DoTest()
}

// DoTest checks Security Deposit Balances
func DoTest() {
	funcname := "DoTest"
	// RentRoll report dates
	dtStart := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	dtStop := time.Date(2017, time.February, 1, 0, 0, 0, 0, time.UTC)

	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := dtStart
	raids := []int64{1, 3, 2, 0, 4, 5, 0, 0}

	lm := rlib.GetRARentableLedgerMarkerOnOrBefore(1, 1, &dtStart)
	rlib.Console("raid=1, rid=1, Dt=%s, lm = %#v\n", dtStart.Format(rlib.RRDATEFMT3), lm)

	// set the limits for which RA(s) we want to process
	iStart := int64(2)
	iStop := int64(3)

	for rid := iStart; rid < iStop; rid++ {
		if int64(0) == raids[rid-1] {
			continue
		}
		x, err := rlib.GetSecDepBalance(App.Xbiz.P.BID, raids[rid-1], rid, &d1, &d2)
		if err != nil {
			fmt.Printf("err = %s\n", err.Error())
			os.Exit(1)
		}
		rlib.Console("SecDep Opening balance on %s  =  %.2f\n\n", dtStart.Format(rlib.RRDATEFMTSQL), x)
		x, err = rlib.GetSecDepBalance(App.Xbiz.P.BID, raids[rid-1], rid, &dtStart, &dtStop)
		if err != nil {
			fmt.Printf("err = %s\n", err.Error())
			os.Exit(1)
		}
		rlib.Console("SecDep Activity between %s and %s  =  %.2f\n",
			dtStart.Format(rlib.RRDATEFMTSQL), dtStop.Format(rlib.RRDATEFMTSQL), x)

		rlib.Console("before rlib.GetBeginEndRARBalance:  dtStart = %s, dtStop = %s\n", dtStart.Format(rlib.RRDATEFMT3), dtStop.Format(rlib.RRDATEFMT3))
		openingBal, closingBal, err := rlib.GetBeginEndRARBalance(rid, raids[rid-1], &dtStart, &dtStop)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			os.Exit(1)
		}
		rlib.Console("rid=%d, raid=%d, %s - %s:   openingBal = %.2f, closingBal = %.2f\n\n\n",
			rid, raids[rid-1], d1.Format(rlib.RRDATEFMT3), d2.Format(rlib.RRDATEFMT3), openingBal, closingBal)
	}
}
