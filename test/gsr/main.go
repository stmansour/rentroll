package main

import (
	"context"
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
	dbdir     *sql.DB        // phonebook db
	dbrr      *sql.DB        // rentroll db
	AcctDep   string         // account depository
	DBDir     string         // phonebook database
	DBRR      string         // rentroll database
	DBUser    string         // user for all databases
	DtStart   time.Time      // range start time
	DtStop    time.Time      // range stop time
	Bal       int            // if < 0 make the total funds less than what is needed, == 0 means equal to what is needed, > 0 means more than what is needed
	Chk2      float64        // amount of check2
	BUD       string         // business unit designator
	GenDbOnly bool           // if true, just set up the db with unallocated funds and exit
	Xbiz      rlib.XBusiness // xbusiness associated with -G  (BUD)
	NoAuth    bool
}

func bizErrCheck(sa []string) {
	if len(sa) < 2 {
		fmt.Printf("Company Designation is required to list Rental Agreements\n")
		os.Exit(1)
	}
}

func loaderGetBiz(ctx context.Context, s string) int64 {
	bid, _ := rcsv.GetBusinessBID(ctx, s)
	if bid == 0 {
		fmt.Printf("unrecognized Business designator: %s\n", s)
		os.Exit(1)
	}
	return bid
}
func readCommandLineArgs() {
	pDates := flag.String("g", "", "Date Range.  Example: 1/1/16,2/1/16")
	pBUD := flag.String("G", "", "BUD - business unit designator")
	pBal := flag.String("funds", "eq", "less: less funds than needed, eq: exactly what is needed,  more: than what is needed")
	pDB := flag.Bool("db", false, "Just generate the db with unallocated funds and exit. Do not apply unallocated funds.")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")

	flag.Parse()
	App.BUD = strings.TrimSpace(*pBUD)
	App.GenDbOnly = *pDB
	App.NoAuth = *noauth

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

	switch strings.ToLower(*pBal) {
	case "eq":
		App.Chk2 = float64(3100)
	case "less":
		App.Chk2 = float64(2500)
	case "more":
		App.Chk2 = float64(3500)
	default:
		fmt.Printf("Unexpected funds value: %s, expecting one of { eq | less | more }\n", *pBal)
		fmt.Printf("Proceeding with default value of \"eq\"\n")
		App.Chk2 = float64(3100)
	}
}

type csvImporteFunc func(context.Context, string) []error

type csvimporter struct {
	Name    string
	CmdOpt  string
	Handler csvImporteFunc
}

func rrDoLoad(ctx context.Context, fname string, handler csvImporteFunc) {
	// fmt.Printf("calling handler for: %q\n", fname)
	m := handler(ctx, fname)
	fmt.Print(rcsv.ErrlistToString(&m))
}

func createReceipt(ctx context.Context, bid int64, amt float64, docno string, dt *time.Time) rlib.Receipt {
	var (
		err error
		r   = rlib.Receipt{
			BID:    bid,
			Dt:     *dt,
			TCID:   2, // for test purposes, this is the payor for all receipts
			Amount: amt,
			DocNo:  docno,
			PMTID:  2,
		}
	)

	arname := "Payment By Check"
	arule, err := rlib.GetARByName(ctx, bid, arname)
	if err != nil {
		rlib.Ulog("Error getting Account Rule by name(%s): %s\n", arname, err.Error())
		return r
	}
	r.ARID = arule.ARID

	// create the receipt
	_, err = rlib.InsertReceipt(ctx, &r)
	if err != nil {
		rlib.Ulog("Error inserting receipt: %s\n", err.Error())
		return r
	}

	// get the AR for this receipt...
	ar := rlib.RRdb.BizTypes[r.BID].AR[r.ARID]

	// get GL Account Info for
	d := rlib.RRdb.BizTypes[r.BID].GLAccounts[ar.DebitLID]
	c := rlib.RRdb.BizTypes[r.BID].GLAccounts[ar.CreditLID]

	// create the receipt allocation
	var ra rlib.ReceiptAllocation
	ra.RCPTID = r.RCPTID
	ra.Amount = r.Amount
	ra.AcctRule = fmt.Sprintf("d %s _, c %s _", d.GLNumber, c.GLNumber)
	ra.BID = r.BID
	ra.Dt = r.Dt
	_, err = rlib.InsertReceiptAllocation(ctx, &ra)
	if err != nil {
		rlib.Ulog("Error inserting receipt: %s\n", err.Error())
		return r
	}
	r.RA = append(r.RA, ra)

	return r
}

func createJournalAndLedgerEntries(ctx context.Context, xbiz *rlib.XBusiness, r *rlib.Receipt, d1, d2, dt1, dt2 *time.Time) error {
	//----------------------------------------------------------
	// Add a journal entry for it.  Note that at this stage
	// we simply move the funds into a bank and credit the
	// available funds account.
	//----------------------------------------------------------
	j, err := rlib.ProcessNewReceipt(ctx, xbiz, d1, d2, r)
	if err != nil {
		rlib.Ulog("Error from rlib.ProcessNewReceipt: %s\n", err.Error())
		return err
	}
	//--------------------------------------------------------------
	// update the ledgers
	//--------------------------------------------------------------
	fmt.Printf("GENERATING LEDGER ENTRIES...\n")
	_, err = rlib.GenerateLedgerEntriesFromJournal(ctx, xbiz, &j, d1, d2)
	if err != nil {
		return err
	}

	//----------------------------------------------
	// force the LedgerMarkers to be generated...
	//----------------------------------------------
	return rlib.GenerateLedgerMarkers(ctx, xbiz, d2)
}

func addUnallocatedReceipts(ctx context.Context, xbiz *rlib.XBusiness, bid int64) {
	err := rlib.InitBizInternals(1, xbiz)
	rlib.Errcheck(err)
	rlib.InitLedgerCache()
	dt2 := time.Now()
	dt1 := dt2.AddDate(0, 0, -6)
	month := dt2.Month()
	d1 := time.Date(dt2.Year(), month, 1, 0, 0, 0, 0, time.UTC) // beginning of this month
	if month == time.December {
		month = time.January
	} else {
		month++
	}
	d2 := time.Date(dt2.Year(), month, 1, 0, 0, 0, 0, time.UTC) // up to but not including beginning of next month

	//----------------------------------------------------
	// We'll create 2 receipts; for $4000 and $3500
	//----------------------------------------------------
	r1 := createReceipt(ctx, bid, float64(4000), "9846", &dt1)
	r2 := createReceipt(ctx, bid, App.Chk2, "9859", &dt2)
	if r1.RCPTID == 0 || r2.RCPTID == 0 {
		fmt.Printf("Could not create receipts\n")
		return
	}

	if nil != createJournalAndLedgerEntries(ctx, xbiz, &r1, &d1, &d2, &dt1, &dt2) {
		return
	}
	if nil != createJournalAndLedgerEntries(ctx, xbiz, &r2, &d1, &d2, &dt1, &dt2) {
		return
	}
}

func doTest(ctx context.Context) {
	// INITIALIZE...
	var xbiz rlib.XBusiness
	bid := int64(1) // Business ID = 1 = REX
	addUnallocatedReceipts(ctx, &xbiz, bid)
	if App.GenDbOnly {
		return
	}

	//--------------------------------------------------------------
	// Now we are ready to start the test. We have 2 unallocated
	// receipts for a user.
	// ...
	// Work through the list of payors that have unallocated funds
	//--------------------------------------------------------------
	rows, err := rlib.RRdb.Prepstmt.GetUnallocatedReceipts.Query(bid)
	rlib.Errcheck(err)
	defer rows.Close()

	// all we need is the list of payors.  For this loop we just
	// put them in a map indexed by TCID, the value at the index is
	// the total number of receipts for that payor which are unallocated
	var u = map[int64]int{}
	for rows.Next() {
		var r rlib.Receipt
		err = rlib.ReadReceipts(rows, &r)
		rlib.Errcheck(err)
		// fmt.Printf("Unallocated Receipt:  RCPTID = %d, Amount = %8.2f, Payor = %d\n", r.RCPTID, r.Amount, r.TCID)
		i, ok := u[r.TCID]
		if ok {
			u[r.TCID] = i + 1
		} else {
			u[r.TCID] = 1
		}
	}
	rlib.Errcheck(rows.Err())

	// Display the list of payors found with Unallocated receipts
	fmt.Printf("Payors with unallocated receipts:\n")
	for k, v := range u {
		fmt.Printf("Payor with TCID=%d has %d unallocated receipts\n", k, v)
	}

	// We assume the user chose to work on Payor with TCID = 2
	tcid := int64(2)
	dt := time.Now()
	err = bizlogic.AutoAllocatePayorReceipts(ctx, tcid, &dt)
	rlib.Errcheck(err)

	// print remaining unpaid assessments, and remaining receipts with unallocated funds
	m := bizlogic.GetAllUnpaidAssessmentsForPayor(ctx, bid, tcid, &dt)
	fmt.Printf("\n\nRemaining unpaid assessments for payor %d:  %d\n", tcid, len(m))
	for i := 0; i < len(m); i++ {
		fmt.Printf("%d. Assessment %d, amount still owed: %.2f\n", i, m[i].ASMID, bizlogic.AssessmentUnpaidPortion(ctx, &m[i]))
	}
	n, err := rlib.GetUnallocatedReceiptsByPayor(ctx, bid, tcid)
	rlib.Errcheck(err)
	fmt.Printf("\nRemaining unallocated funds for payor %d:  %d\n", tcid, len(n))
	for i := 0; i < len(n); i++ {
		fmt.Printf("%d. Receipt %d, amount remaining: %.2f\n", i, n[i].RCPTID, bizlogic.RemainingReceiptFunds(ctx, &n[i]))
	}
	fmt.Printf("-------------------------------------------------------------\n")
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
	rlib.SetAuthFlag(App.NoAuth)
	rlib.SessionInit(10) // must be called before calling InitBizInternals

	// create background context
	ctx := context.Background()

	if len(App.BUD) > 0 {
		b2, err := rlib.GetBusinessByDesignation(ctx, App.BUD)
		rlib.Errcheck(err)
		if b2.BID == 0 {
			fmt.Printf("Could not find Business Unit named %s\n", App.BUD)
			os.Exit(1)
		}
		rlib.GetXBusiness(ctx, b2.BID, &App.Xbiz)
	}
	if App.Xbiz.P.BID > 0 {
		rcsv.InitRCSV(&App.DtStart, &App.DtStop, &App.Xbiz)
		err = rlib.InitBizInternals(App.Xbiz.P.BID, &App.Xbiz)
		rlib.Errcheck(err)
	}

	doTest(ctx)
}
