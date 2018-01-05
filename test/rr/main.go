// The purpose of this test is to validate the db update routines.
package main

import (
	"context"
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
	NoAuth bool
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
	App.NoAuth = *noauth
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
	rlib.SetAuthFlag(App.NoAuth)

	// create background context
	ctx := context.Background()

	biz, err := rlib.GetBusinessByDesignation(ctx, App.Bud)
	rlib.Errcheck(err)
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}

	err = rlib.InitBizInternals(biz.BID, &App.Xbiz)
	rlib.Errcheck(err)

	DoTest(ctx)
}

// DoTest checks Security Deposit Balances
func DoTest(ctx context.Context) {
	test1(ctx)
	test2(ctx)
}

func test2(ctx context.Context) {
	var ds = []struct {
		sd1, sd2 string
	}{
		{"2016-10-01", "2016-10-08"},
		{"2016-10-15", "2016-10-25"},
		{"2016-10-21", "2016-10-29"},
	}
	dtStart, _ := rlib.StringToDate("2016-10-01")
	dtStop, _ := rlib.StringToDate("2016-11-01")
	var d []rlib.Period
	for i := 0; i < len(ds); i++ {
		var p rlib.Period
		p.D1, _ = rlib.StringToDate(ds[i].sd1)
		p.D2, _ = rlib.StringToDate(ds[i].sd2)
		d = append(d, p)
	}
	// dd := rlib.AggregatePeriods(&dtStart, &dtStop, d)
	// for i := 0; i < len(dd); i++ {
	// 	fmt.Printf("dd[%d] = %s - %s\n", i, dd[i].D1.Format("2006-01-02"), dd[i].D2.Format("2006-01-02"))
	// }
	ddd := rlib.FindGaps(&dtStart, &dtStop, d)
	for i := 0; i < len(ddd); i++ {
		fmt.Printf("ddd[%d] = %s - %s\n", i, ddd[i].D1.Format("2006-01-02"), ddd[i].D2.Format("2006-01-02"))
	}

}

func test1(ctx context.Context) {
	const funcname = "DoTest"
	var err error
	// RentRoll report dates
	dtStart := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	dtStop := time.Date(2017, time.February, 1, 0, 0, 0, 0, time.UTC)

	d1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := dtStart
	raids := []int64{1, 3, 2, 0, 4, 5, 0, 0}

	lm, _ := rlib.GetRARentableLedgerMarkerOnOrBefore(ctx, 1, 1, &dtStart)
	rlib.Errcheck(err)
	rlib.Console("raid=1, rid=1, Dt=%s, lm = %#v\n", dtStart.Format(rlib.RRDATEFMT3), lm)

	// set the limits for which RA(s) we want to process
	iStart := int64(2)
	iStop := int64(3)

	for rid := iStart; rid < iStop; rid++ {
		if int64(0) == raids[rid-1] {
			continue
		}
		x, err := rlib.GetSecDepBalance(ctx, App.Xbiz.P.BID, raids[rid-1], rid, &d1, &d2)
		if err != nil {
			fmt.Printf("err = %s\n", err.Error())
			os.Exit(1)
		}
		rlib.Console("SecDep Opening balance on %s  =  %.2f\n\n", dtStart.Format(rlib.RRDATEFMTSQL), x)
		x, err = rlib.GetSecDepBalance(ctx, App.Xbiz.P.BID, raids[rid-1], rid, &dtStart, &dtStop)
		if err != nil {
			fmt.Printf("err = %s\n", err.Error())
			os.Exit(1)
		}
		rlib.Console("SecDep Activity between %s and %s  =  %.2f\n",
			dtStart.Format(rlib.RRDATEFMTSQL), dtStop.Format(rlib.RRDATEFMTSQL), x)

		rlib.Console("before rlib.GetBeginEndRARBalance:  dtStart = %s, dtStop = %s\n", dtStart.Format(rlib.RRDATEFMT3), dtStop.Format(rlib.RRDATEFMT3))
		openingBal, closingBal, err := rlib.GetBeginEndRARBalance(ctx, App.Xbiz.P.BID, rid, raids[rid-1], &dtStart, &dtStop)
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			os.Exit(1)
		}
		rlib.Console("rid=%d, raid=%d, %s - %s:   openingBal = %.2f, closingBal = %.2f\n\n\n",
			rid, raids[rid-1], d1.Format(rlib.RRDATEFMT3), d2.Format(rlib.RRDATEFMT3), openingBal, closingBal)
	}
}
