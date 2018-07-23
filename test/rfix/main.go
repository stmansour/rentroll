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
	pBud := flag.String("b", "OKC", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")

	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
}

func doFix(ctx context.Context) {
	var m []rlib.Receipt
	funcname := "doFix"

	qry := fmt.Sprintf("SELECT %s FROM Receipt ORDER BY RCPTID ASC", rlib.RRdb.DBFields["Receipt"])
	rlib.Console("qry = %s\n", qry)
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("Error in query: %s\n", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r rlib.Receipt
		err = rlib.ReadReceipts(rows, &r)
		if err != nil {
			rlib.Console("Error in ReadReceipts: %s\n", err.Error())
			return
		}
		m = append(m, r)
	}

	lenm := len(m)
	lenm2 := lenm - 2
	fixes := 0

	for i := 0; i < lenm2; i++ {

		if m[i].RCPTID < int64(0) {
			continue
		}

		j := i + 1 // sometimes BID = 0 on this one
		k := j + 1 // sometimes an exact copy of m[i]

		// rlib.Console("5.  m[j].RCPTID = %d\n", m[j].RCPTID)
		// rlib.Console("    m[k].RCPTID = %d\n", m[k].RCPTID)
		//--------------------------------------------------------
		// If the one we're looking at has BID == 0, delete it
		//--------------------------------------------------------
		if m[i].BID == 0 {
			// rlib.Console("*** FIX ***  A\n")
			if err := rlib.DeleteReceipt(ctx, m[i].RCPTID); err != nil {
				rlib.Console("%s: Error deleting RCPTID %d: %s\n", funcname, m[i].RCPTID, err.Error())
			}
			m[i].RCPTID = int64(-1) // no chance of processing it further
			fixes++
			continue
		}

		riname, _ := rlib.ROCExtractRentableName(m[i].Comment)
		rkname, _ := rlib.ROCExtractRentableName(m[k].Comment)

		// rlib.Console("6.  %t: Name match:    %s : %s\n", riname == rkname, riname, rkname)
		// rlib.Console("    %t: Amounts match: %8.2f : %8.2f\n", m[k].Amount == m[i].Amount, m[k].Amount, m[i].Amount)
		// rlib.Console("    %t:                m[j].BID == 0\n", m[j].BID == 0)
		//--------------------------------------------------------
		// This is the main error pattern we're looking for...
		//--------------------------------------------------------
		if m[j].BID == 0 && m[k].Amount == m[i].Amount && m[k].OtherPayorName == m[i].OtherPayorName && riname == rkname {
			if err := rlib.DeleteReceipt(ctx, m[j].RCPTID); err != nil {
				rlib.Console("%s: Error deleting RCPTID %d: %s\n", funcname, m[j].RCPTID, err.Error())
			}
			m[j].RCPTID = int64(-1) // no chance of processing it further
			fixes++

			if err := rlib.DeleteReceipt(ctx, m[k].RCPTID); err != nil {
				rlib.Console("%s: Error deleting RCPTID %d: %s\n", funcname, m[k].RCPTID, err.Error())
			}
			m[k].RCPTID = int64(-1) // no chance of processing it further
			fixes++
		}
	}
	rlib.Console("Records processed: %d\n", lenm)
	rlib.Console("Fixes: %d\n", fixes)
}

func main() {
	var err error
	readCommandLineArgs()
	App.NoAuth = true
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

	// create background context
	ctx := context.Background()

	rlib.Console("ctx = %#v\n", ctx)
	rlib.Console("App.Bud = %s\n", App.Bud)
	biz, err := rlib.GetBusinessByDesignation(ctx, App.Bud)
	rlib.Errcheck(err)
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}
	err = rlib.GetXBusiness(ctx, biz.BID, &App.Xbiz)
	rlib.Errcheck(err)
	rlib.Console("BID = %d\n", App.Xbiz.P.BID)

	doFix(ctx)
}
