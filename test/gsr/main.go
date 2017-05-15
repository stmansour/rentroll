package main

import (
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"os"
	"rentroll/bizlogic"
	"rentroll/rcsv"
	"rentroll/rlib"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir   *sql.DB        // phonebook db
	dbrr    *sql.DB        //rentroll db
	AcctDep string         // account depository
	DBDir   string         // phonebook database
	DBRR    string         //rentroll database
	DBUser  string         // user for all databases
	DtStart time.Time      // range start time
	DtStop  time.Time      // range stop time
	BUD     string         // business unit designator
	Xbiz    rlib.XBusiness // xbusiness associated with -G  (BUD)
}

func bizErrCheck(sa []string) {
	if len(sa) < 2 {
		fmt.Printf("Company Designation is required to list Rental Agreements\n")
		os.Exit(1)
	}
}

func loaderGetBiz(s string) int64 {
	bid := rcsv.GetBusinessBID(s)
	if bid == 0 {
		fmt.Printf("unrecognized Business designator: %s\n", s)
		os.Exit(1)
	}
	return bid
}
func readCommandLineArgs() {
	pDates := flag.String("g", "", "Date Range.  Example: 1/1/16,2/1/16")
	pBUD := flag.String("G", "", "BUD - business unit designator")

	flag.Parse()
	App.BUD = strings.TrimSpace(*pBUD)
	var err error
	s := *pDates
	if len(s) > 0 {
		ss := strings.Split(s, ",")
		App.DtStart, err = rlib.StringToDate(ss[0])
		if err != nil {
			fmt.Printf("Invalid start date:  %s\n", ss[0])
			os.Exit(1)
		}
		App.DtStop, err = rlib.StringToDate(ss[1])
		if err != nil {
			fmt.Printf("Invalid stop date:  %s\n", ss[1])
			os.Exit(1)
		}
	}
}

type csvimporter struct {
	Name    string
	CmdOpt  string
	Handler func(string) []error
}

func rrDoLoad(fname string, handler func(string) []error) {
	// fmt.Printf("calling handler for: %q\n", fname)
	m := handler(fname)
	fmt.Print(rcsv.ErrlistToString(&m))
}

func main() {
	readCommandLineArgs()
	rlib.RRReadConfig()

	var err error

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

	if len(App.BUD) > 0 {
		b2 := rlib.GetBusinessByDesignation(App.BUD)
		if b2.BID == 0 {
			fmt.Printf("Could not find Business Unit named %s\n", App.BUD)
			os.Exit(1)
		}
		rlib.GetXBusiness(b2.BID, &App.Xbiz)
	}
	if App.Xbiz.P.BID > 0 {
		rcsv.InitRCSV(&App.DtStart, &App.DtStop, &App.Xbiz)
		rlib.InitBizInternals(App.Xbiz.P.BID, &App.Xbiz)
	}

	//----------------------------------------------------
	// For the receipt, we just create an entry by hand.
	// Everything else is derived
	//----------------------------------------------------
	var r rlib.Receipt
	r.BID = 1
	r.Dt = time.Now()
	r.TCID = 2
	r.Amount = float64(7500.00)
	r.DocNo = "9843"
	r.PMTID = 2
	r.ARID = 2

	// create the receipt
	_, err = rlib.InsertReceipt(&r)
	if err != nil {
		rlib.Ulog("Error inserting receipt: %s\n", err.Error())
		return
	}

	// get the AR for this receipt...
	ar, err := rlib.GetAR(r.ARID)
	if err != nil {
		rlib.Ulog("Error getting AR: %s\n", err.Error())
		return
	}

	// get GL Account Info for
	d := rlib.RRdb.BizTypes[r.BID].GLAccounts[ar.DebitLID]
	c := rlib.RRdb.BizTypes[r.BID].GLAccounts[ar.CreditLID]

	//----------------------------------------------------------
	// Add a journal entry for it.  Note that at this stage
	// there is no RID or RAID associated with the transaction.
	// We simply move the funds into a bank and credit the
	// available funds account.
	//----------------------------------------------------------
	var j rlib.Journal
	j.BID = r.BID
	j.Dt = r.Dt
	j.Amount = r.Amount
	j.Type = rlib.JNLTYPERCPT
	j.ID = r.RCPTID // this is the receipt just created
	_, err = rlib.InsertJournal(&j)
	if err != nil {
		rlib.Ulog("Error inserting journal: %s\n", err.Error())
		return
	}

	var ja rlib.JournalAllocation
	ja.JID = j.JID
	ja.BID = j.BID
	ja.Amount = j.Amount
	ja.AcctRule = fmt.Sprintf("d %s _, c %s _", d.GLNumber, c.GLNumber)

	// update the ledger

	// process the receipt and put the payment into unapplied funds

	bizlogic.AutoProcessReceipt(&r)

}
