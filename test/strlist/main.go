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
	dbdir   *sql.DB        // phonebook db
	dbrr    *sql.DB        //rentroll db
	DBDir   string         // phonebook database
	DBRR    string         //rentroll database
	DBUser  string         // user for all databases
	PortRR  int            // rentroll port
	Bud     string         // Biz Unit Descriptor
	Xbiz    rlib.XBusiness // lots of info about this biz
	NoAuth  bool           //
	Verbose bool           //
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")
	verb := flag.Bool("v", false, "verbose output - shows the ciphertext")

	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
	App.NoAuth = *noauth
	App.Verbose = *verb
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
	rlib.SessionInit(10) // must be called before calling InitBizInternals

	//--------------------------------------
	// create background context
	//--------------------------------------
	rlib.SessionInit(15) // must be called first, creates channels
	now := time.Now()
	expire := now.Add(1 * time.Minute)
	ssn := rlib.SessionNew("Flow2RATester", "Flow2RATester", "Flow2RATester", -99999, "", 0, &expire)
	ctx := context.Background()
	ctx = rlib.SetSessionContextKey(ctx, ssn)
	DoTest(ctx, ssn)
}

// DoTest does all the useful and interesting work
func DoTest(ctx context.Context, s *rlib.Session) {
	var xbiz rlib.XBusiness
	var bid = int64(1)
	var err error
	if err = rlib.InitBizInternals(bid, &xbiz); err != nil {
		log.Printf("InitBizInternals error: %s\n", err.Error())
		return
	}

	fmt.Printf("StringList RollerMsgs MSGRAUPDATED:\n")
	fmt.Printf("    Value = %s\n", rlib.RRdb.BizTypes[bid].Msgs.S[rlib.MSGRAUPDATED].Value)
	fmt.Printf("    SLSID = %d\n", rlib.RRdb.BizTypes[bid].Msgs.S[rlib.MSGRAUPDATED].SLSID)
	fmt.Printf("    SLID  = %d\n", rlib.RRdb.BizTypes[bid].Msgs.S[rlib.MSGRAUPDATED].SLID)

}
