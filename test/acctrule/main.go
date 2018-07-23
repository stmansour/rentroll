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
	rlib.SessionInit(10) // must be called before calling InitBizInternals

	// Create background context
	ctx := context.Background()

	biz, err := rlib.GetBusinessByDesignation(ctx, App.Bud)
	rlib.Errcheck(err)
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}
	err = rlib.GetXBusiness(ctx, biz.BID, &App.Xbiz)
	rlib.Errcheck(err)

	// based on work from: https://play.golang.org/p/p842UZpQaK
	var m []rlib.AcctRule
	var xbiz rlib.XBusiness
	now := time.Now()
	rpnCtx := rlib.RpnCreateCtx(&xbiz, 1, &now, &now, &m, float64(-2000), float64(1.0))

	rlib.RpnInit()
	var expr = []string{
		"-2000.00",
		"3.0 4.0 -",
		"2000.00",
	}

	for i := 0; i < len(expr); i++ {
		// TODO(Steve): should we ignore error?
		x, _ := rlib.RpnCalculateEquation(ctx, &rpnCtx, expr[i])
		fmt.Printf("%d. x = %.2f\n", i, x)
	}
}
