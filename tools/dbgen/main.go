// The purpose of this test is to validate the db update routines.
package main

import (
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"os"
	"rentroll/bizlogic"
	"rentroll/rlib"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir        *sql.DB        // phonebook db
	dbrr         *sql.DB        //rentroll db
	DBDir        string         // phonebook database
	DBRR         string         //rentroll database
	DBUser       string         // user for all databases
	Bud          string         // Biz Unit Descriptor
	Xbiz         rlib.XBusiness // lots of info about this biz
	ConfFileName string         // configuration file
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	pFile := flag.String("f", "dbconf.json", "settings that define DB generation")

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.Bud = *pBud
	App.ConfFileName = *pFile
}

func main() {
	//----------------------------------------------------------------
	// Initialize the empty database. It should contain things like:
	//   Chart of Accounts
	//   Account Rules
	//   Depositories
	//   Payment Types
	//   Deposit Methods
	//----------------------------------------------------------------
	rc := InitEmptyDB()
	if rc != 0 {
		os.Exit(rc)
	}

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
	bizlogic.InitBizLogic()

	//----------------------------
	// Read initialization
	//----------------------------
	ctx, err := ReadConfig(App.ConfFileName)
	if err != nil {
		fmt.Printf("Error loading %s: %s\n", App.ConfFileName, err.Error())
		os.Exit(1)
	}

	ctx.BIZ, err = rlib.GetAllBusinesses()
	if err != nil {
		fmt.Printf("Error loading businesses: %s\n", err.Error())
		os.Exit(1)
	}
	if len(ctx.BIZ) == 0 {
		fmt.Printf("Error: database contains no businesses\n")
		os.Exit(1)
	}

	//----------------------------
	// Build the database
	//----------------------------
	err = GenerateDB(&ctx)
	if err != nil {
		fmt.Printf("Error generating database: %s\n", err.Error())
		os.Exit(1)
	}

}
