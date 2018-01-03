// The purpose of this test is to validate the db update routines.
package main

import (
	"context"
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"log"
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

	flag.Parse()

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
	if /*biz.BID == 0*/ err != nil {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}

	err = rlib.InitBizInternals(biz.BID, &App.Xbiz)
	if err != nil {
		fmt.Printf("Error in InitBizInternals: %s\n", err.Error())
		os.Exit(1)
	}

	DoTest(ctx)
}

// DoTest does the account balance checks for Rental Agreements
func DoTest(ctx context.Context) {
	raid := int64(5)
	d1 := time.Date(2017, time.May, 1, 0, 0, 0, 0, rlib.RRdb.Zone)
	d2 := d1.AddDate(0, 1, 0)

	fmt.Printf("RA statement info for period %s - %s\n", d1.Format(rlib.RRDATEREPORTFMT), d2.Format(rlib.RRDATEREPORTFMT))
	m, err := rlib.GetRAIDStatementInfo(ctx, raid, &d1, &d2)
	if err != nil {
		log.Fatalf("*** ERROR *** GetRAIDAccountBalance returned error: %s\n", err.Error())
	}
	fmt.Printf("m.OpeningBal -->  %8.2f\n", m.OpeningBal)

	newbal := m.LmStart.Balance
	for i := 0; i < len(m.Gap); i++ {
		switch m.Gap[i].T {
		case 1: // Assessment
			newbal -= m.Gap[i].Amt
			fmt.Printf("date = %s, asmt = %8.2f,  bal = %8.2f\n", m.Gap[i].A.Start.Format(rlib.RRDATEREPORTFMT), -m.Gap[i].Amt, newbal)
		case 2: // Receipt Allocation
			newbal += m.Gap[i].Amt
			fmt.Printf("date = %s, RCPT = %8.2f,  bal = %8.2f\n", m.Gap[i].R.Dt.Format(rlib.RRDATEREPORTFMT), m.Gap[i].Amt, newbal)
		}
	}
	fmt.Printf("OpeningBal = %8.2f,   newbal = %8.2f\n", m.OpeningBal, newbal)

	fmt.Printf("\nSTATEMENT\n")
	fmt.Printf("%s Opening Balance: %8.2f\n", m.DtStart.Format(rlib.RRDATEREPORTFMT), m.OpeningBal)
	newbal = m.OpeningBal
	for i := 0; i < len(m.Stmt); i++ {
		switch m.Stmt[i].T {
		case 1: // Assessment
			newbal -= m.Stmt[i].Amt
			fmt.Printf("%s, asmt = %8.2f,  bal = %8.2f\n", m.Stmt[i].A.Start.Format(rlib.RRDATEREPORTFMT), -m.Stmt[i].Amt, newbal)
		case 2: // Receipt Allocation
			newbal += m.Stmt[i].Amt
			fmt.Printf("%s, RCPT = %8.2f,  bal = %8.2f\n", m.Stmt[i].R.Dt.Format(rlib.RRDATEREPORTFMT), m.Stmt[i].Amt, newbal)
		}
	}
	fmt.Printf("%s ClosingBal = %8.2f,   newbal = %8.2f\n", m.DtStop.AddDate(0, 0, -1).Format(rlib.RRDATEREPORTFMT), m.ClosingBal, newbal)
}
