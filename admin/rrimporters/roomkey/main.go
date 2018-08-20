package main

import (
	"context"
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"phonebook/lib"
	"rentroll/importers/core"
	"rentroll/importers/roomkey"
	"rentroll/rlib"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kardianos/osext"
)

// App is the global application structure used for roomkey csv importer
var App struct {
	dbdir        *sql.DB  // phonebook db
	dbrr         *sql.DB  // rentroll db
	DBDir        string   // phonebook database
	DBRR         string   // rentroll database
	DBUser       string   // user for all databases
	LogFile      *os.File // where to log messages
	TestMode     int      // used for test purpose?
	CSV          string   // csv filename that needs to be load
	GuestInfoCSV string   // csv filename containing guest info
	debug        int      // debug records
	NoAuth       bool     // noauth flag
}

// userRRValues holds the values passed by user for rentroll attributes
var userRRValues = make(map[string]string)

// MergeSuppliedAndDefaultValues used to merge
// override values from userRRValues map into matched
// field of Defaults
func MergeSuppliedAndDefaultValues() {

	// override default values to userRRValues map
	// if not passed
	for k := range userRRValues {
		if userRRValues[k] == "" {
			if defaultVal, ok := roomkey.FieldDefaultValues[k]; ok {
				userRRValues[k] = defaultVal
			}
		}
	}

	// append also roomkey fields in userRRValues
	// if it does not exist in map
	for k, v := range roomkey.FieldDefaultValues {
		if _, ok := userRRValues[k]; !ok {
			userRRValues[k] = v
		}
	}
}

func readCommandLineArgs() []string {
	inputErrors := []string{}
	// a csv file must be passed
	fp := flag.String("csv", "", "Path of the roomkey CSV file to import")
	// a csv file must be passed
	guestInfoFp := flag.String("guestinfo", "", "Path of CSV file containing guest info (Guest Export)")
	// a bud must be passed
	bud := flag.String("bud", "", "A business unit designation")
	// frequency should default to monthly
	frequency := flag.String("frequency", "", "Rent Cycle")
	// proration should default to daily
	proration := flag.String("proration", "", "Proration Cycle")
	// gsrpc should default to daily
	gsrpc := flag.String("gsrpc", "", "GSRPC")
	// is it for testing purpose
	testmode := flag.Int("testmode", 0, "testing")
	// is it for debug purpose
	debug := flag.Int("debug", 0, "debug Records")
	// parse db options
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")

	// ================================
	// check for values which must be required
	// ================================

	// parse the values from command line
	flag.Parse()

	if *fp == "" {
		inputErrors = append(inputErrors, "Please, pass roomkey csv input file")
	}

	if *bud == "" {
		inputErrors = append(inputErrors, "Please, pass business unit designation")
	}

	// above inputs must required from users
	// so put condition here
	if len(inputErrors) > 0 {
		return inputErrors
	}

	// App structure values
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.TestMode = *testmode
	App.CSV = *fp
	App.GuestInfoCSV = *guestInfoFp
	App.debug = *debug
	App.NoAuth = *noauth

	// get user values
	userRRValues["RentCycle"] = *frequency
	userRRValues["Proration"] = *proration
	userRRValues["GSRPC"] = *gsrpc
	userRRValues["BUD"] = *bud

	return inputErrors
}

func main() {

	// ================================
	// COMMAND LINE OPTIONS VALIDATION
	// ================================
	inputErrors := readCommandLineArgs()
	if len(inputErrors) > 0 {
		for _, errText := range inputErrors {
			fmt.Println(errText)
		}
		os.Exit(1)
	}

	// ==========================================================
	// INITIAL SETUP: CSV TEMP STORAGE, DATABASE CONNECTION, LOG FILE
	// ==========================================================

	// error variable
	var err error

	// LOGFILE SETUP
	App.LogFile, err = os.OpenFile("roomkey.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	lib.Errcheck(err)
	defer App.LogFile.Close()
	log.SetOutput(App.LogFile)
	rlib.Ulog("*********** ROOMKEY IMPORTER HAS BEEN STARTED *********** \n")

	// CSV STORE CHECK
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		rlib.Ulog("INTERNAL ERROR <INITIALIZATION>: %s", err.Error())
		os.Exit(1)
	}

	// get path of splitted csv store
	roomkey.TempCSVStore = path.Join(folderPath, roomkey.TempCSVStoreName)

	// if tempCSVStore not exist then create it
	if _, err := os.Stat(roomkey.TempCSVStore); os.IsNotExist(err) {
		os.MkdirAll(roomkey.TempCSVStore, 0700)
	}
	if err != nil {
		rlib.Ulog("INTERNAL ERROR <INITIALIZATION>: %s", err.Error())
		os.Exit(1)
	}

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
	rlib.SetNoAuthFlag(App.NoAuth) // currently needed for testing
	rlib.SessionInit(10)           // must be called before calling InitBizInternals

	// create background context
	ctx := context.Background()

	// ==================================
	// AFTER DB SETUP DO VALIDATION OVER
	// USER SUPPLIED VALUES WITH DB VALUES
	// ==================================

	// merge user supplied values with default one
	MergeSuppliedAndDefaultValues()

	// now validation on user supplied values
	validateErrs, business := roomkey.ValidateUserSuppliedValues(ctx, userRRValues)
	if len(validateErrs) > 0 {
		for _, err := range validateErrs {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}

	// =======================
	// CALL ONSITE CSV HANDLER
	// =======================

	// call roomkey loader
	report, internalErr, done := roomkey.CSVHandler(
		ctx,
		App.CSV,
		App.GuestInfoCSV,
		App.TestMode,
		userRRValues,
		business,
		App.debug,
	)

	if internalErr {
		var roomKeyErrText string
		roomKeyErrText = core.ErrInternal.Error()
		fmt.Println(roomKeyErrText)
		os.Exit(1)
	}

	if !done {
		fmt.Printf("RoomKey CSV did not import properly. Please look out at the report.\n\n")
		fmt.Println(report)
	} else {
		// SUCCESS THEN REPORT IT
		fmt.Println(report)
	}
}
