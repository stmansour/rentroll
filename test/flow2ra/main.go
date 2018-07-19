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
	"rentroll/ws"
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
	FlowID  int64          // the flow to migrate to permanent tables
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	idPtr := flag.Int("flowid", 0, "FlowID to migrate. If not specified then RAID 1 is migrated.")
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
	App.FlowID = int64(*idPtr)
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
	var flowID = int64(App.FlowID)
	var err error

	if flowID == 0 {
		rlib.Console("Retrieving Rental Agreement\n")
		//--------------------------------------
		// Make a Flow out of one of the RAs
		//--------------------------------------
		raid := int64(1)
		ra, err := rlib.GetRentalAgreement(ctx, raid)
		if err != nil {
			fmt.Printf("Could not read RentalAgreement: %s\n", err.Error())
			return
		}

		//--------------------------------------
		// Insert new flow
		//--------------------------------------
		rlib.Console("Creating a Flow for RAID %d\n", ra.RAID)
		flowID, err = ws.GetRA2FlowCore(ctx, &ra, s.UID)
		if err != nil {
			rlib.Console("DoTest - CB.err\n")
			fmt.Printf("Could not get Flow for RAID = %d: %s\n", ra.RAID, err.Error())
			return
		}

		fmt.Printf("Successfully created FlowID = %d\n", flowID)
	} else {
		fmt.Printf("FlowID set to %d\n", flowID)
	}

	//------------------------------------------------------------
	// Insert it back into the permanent db tables as an updated
	// Rental Agreement linked to its predecessor
	//------------------------------------------------------------
	tx, tctx, err := rlib.NewTransactionWithContext(ctx)
	if err != nil {
		fmt.Printf("Could not create transaction context: %s\n", err.Error())
		return
	}
	nraid, err := ws.Flow2RA(tctx, flowID)
	if err != nil {
		tx.Rollback()
		rlib.Console("Flow2RA error\n")
		fmt.Printf("Could not write Flow back to db: %s\n", err.Error())
		return
	}
	if err = tx.Commit(); err != nil {
		fmt.Printf("Error committing transaction: %s\n", err.Error())
		return
	}
	rlib.Console("Successfully created new Rental Agreement, RAID = %d\n", nraid)
}
