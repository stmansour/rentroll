package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"phonebook/lib"
	"rentroll/importers/roomkey"
	"rentroll/rlib"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure used for roomkey csv importer
var App struct {
	dbdir    *sql.DB  // phonebook db
	dbrr     *sql.DB  // rentroll db
	DBDir    string   // phonebook database
	DBRR     string   // rentroll database
	DBUser   string   // user for all databases
	LogFile  *os.File // where to log messages
	TestMode int      // used for test purpose?
	CSV      string   // csv filename that needs to be load
}

// userRRValues holds the values passed by user for rentroll attributes
var userRRValues = make(map[string]string)

// GetRoomKeyFieldDefaultValues used to return map[string]string
// with field values, which is default to values defined here
func GetRoomKeyFieldDefaultValues() map[string]string {
	defaults := map[string]string{}
	defaults["ManageToBudget"] = "1" // always take to default this one
	defaults["RentCycle"] = "6"      // maybe overridden by user supplied value
	defaults["Proration"] = "4"      // maybe overridden by user supplied value
	defaults["GSRPC"] = "4"          // maybe overridden by user supplied value
	defaults["AssignmentTime"] = "1" // always take to default this one
	defaults["Renewal"] = "2"        // always take to default this one
	return defaults
}

// MergeSuppliedAndDefaultValues used to merge
// override values from userRRValues map into matched
// field of Defaults
func MergeSuppliedAndDefaultValues() {
	defaults := GetRoomKeyFieldDefaultValues()

	// override default values to userRRValues map
	// if not passed
	for k := range userRRValues {
		if userRRValues[k] == "" {
			if defaultVal, ok := defaults[k]; ok {
				userRRValues[k] = defaultVal
			}
		}
	}

	// append also defaults fields in userRRValues
	// if it does not exist in map
	for k, v := range defaults {
		if _, ok := userRRValues[k]; !ok {
			userRRValues[k] = v
		}
	}
}

// TODO: remove this accrual rate later
// Rental accrual rate
// 0 = one time only
// 1 = secondly
// 2 = minutely
// 3 = hourly
// 4 = daily
// 5 = weekly
// 6 = monthly
// 7 = quarterly
// 8 = yearly

func readCommandLineArgs() (bool, []string) {
	ok, errors := true, []string{}
	// a csv file must be passed
	fp := flag.String("csv", "", "the name of the roomkey CSV file to import")
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

	// parse db options
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")

	// ================================
	// check for values which must be required
	// ================================

	// parse the values from command line
	flag.Parse()

	if *fp == "" {
		ok = false
		errors = append(errors, "Please, pass roomkey csv input file")
	}

	if *bud == "" {
		ok = false
		errors = append(errors, "Please, pass business unit designation")
	}

	// if not ok then return with errors, otherwise fill up values in map
	if !ok {
		return ok, errors
	}

	// App structure values
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.TestMode = *testmode
	App.CSV = *fp

	// get user values
	userRRValues["RentCycle"] = *frequency
	userRRValues["Proration"] = *proration
	userRRValues["GSRPC"] = *gsrpc
	userRRValues["BUD"] = *bud

	return ok, errors
}

func main() {
	var err error

	// setup log file
	App.LogFile, err = os.OpenFile("roomkey.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	lib.Errcheck(err)
	defer App.LogFile.Close()
	log.SetOutput(App.LogFile)
	rlib.Ulog("IMPORTERS started \n")

	// read command line argument first
	ok, inputErrors := readCommandLineArgs()
	if !ok {
		fmt.Printf("Input Errors: %v\n", inputErrors)
		os.Exit(1)
	}

	// db initialization
	rlib.RRReadConfig()

	//----------------------------
	// Open RentRoll database
	//----------------------------
	// s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", DBUser, DBRR)
	s := rlib.RRGetSQLOpenString(App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = rlib.RRGetSQLOpenString(App.DBDir)
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
	// ###########################################
	// DB INIT COMPLETE     ##
	// ###########################################

	// merge user supplied values with default one
	MergeSuppliedAndDefaultValues()

	// call roomkey loader
	done, ErrReport, roomKeyErr := roomkey.CSVHandler(
		App.CSV,
		App.TestMode,
		userRRValues,
	)

	var roomKeyErrText string
	if roomKeyErr != nil {
		roomKeyErrText = roomKeyErr.Error()
	}
	fmt.Printf("\n1. ROOMKEY IMPORTING SUCCESSFULLY DONE: %v", done)
	fmt.Printf("\n2. ROOMKEY ERRORS: %v", roomKeyErrText)
	fmt.Printf("\n3. ROOMKEY CSV ERROR REPORT:")
	fmt.Printf("\n%s", strings.Repeat("=", 65))
	fmt.Printf("\n%s", ErrReport)
}
