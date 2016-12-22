package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"rentroll/importers/onesite"
	"rentroll/rlib"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir  *sql.DB // phonebook db
	dbrr   *sql.DB //rentroll db
	DBDir  string  // phonebook database
	DBRR   string  //rentroll database
	DBUser string  // user for all databases
}

var userSuppliedValues = make(map[string]string)

// GetOneSiteFieldDefaultValues used to return map[string]string
// with field values, which is default to values defined here
func GetOneSiteFieldDefaultValues() map[string]string {
	defaults := map[string]string{}
	defaults["ManageToBudget"] = "1"
	defaults["RentCycle"] = "6"
	defaults["Proration"] = "4"
	defaults["GSRPC"] = "4"
	defaults["AssignmentTime"] = "1"
	defaults["Renewal"] = "2"
	return defaults
}

// MergeSuppliedAndDefaultValues used to merge
// override values from userSuppliedValues map into matched
// field of Defaults
func MergeSuppliedAndDefaultValues() {
	defaults := GetOneSiteFieldDefaultValues()

	// override default values to userSuppliedValues map
	// if not passed
	for k := range userSuppliedValues {
		if userSuppliedValues[k] == "" {
			if defaultVal, ok := defaults[k]; ok {
				userSuppliedValues[k] = defaultVal
			}
		}
	}

	// append also defaults fields in userSuppliedValues
	// if it does not exist in map
	for k, v := range defaults {
		if _, ok := userSuppliedValues[k]; !ok {
			userSuppliedValues[k] = v
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
	fp := flag.String("csv", "", "the name of the onesite CSV file to import")
	// a bud must be passed
	bud := flag.String("bud", "", "A business unit designation")
	// frequency should default to monthly
	frequency := flag.String("frequency", "", "Rent Cycle")
	// proration should default to daily
	proration := flag.String("proration", "", "Proration Cycle")
	// gsrpc should default to daily
	gsrpc := flag.String("gsrpc", "", "GSRPC")
	// is it for testing purpose
	testmode := flag.String("testmode", "0", "testing")

	// parse db options
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")

	// parse the values from command line
	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr

	// ================================
	// check for values which must be required
	// ================================
	if *fp == "" {
		ok = false
		errors = append(errors, "Please, pass onesite csv input file")
	}

	if *bud == "" {
		ok = false
		errors = append(errors, "Please, pass business unit designation")
	}

	// if not ok then return with errors, otherwise fill up values in map
	if !ok {
		return ok, errors
	}

	userSuppliedValues["OneSiteCSV"] = *fp
	userSuppliedValues["BUD"] = *bud
	userSuppliedValues["RentCycle"] = *frequency
	userSuppliedValues["Proration"] = *proration
	userSuppliedValues["GSRPC"] = *gsrpc
	userSuppliedValues["testmode"] = *testmode

	return ok, errors
}

func main() {
	// read command line argument first
	ok, inputErrors := readCommandLineArgs()
	if !ok {
		fmt.Printf("%v\n", inputErrors)
		return
	}

	// db initialization
	rlib.RRReadConfig()
	var err error

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

	// call onesite loader
	errors, msg := onesite.CSVHandler(userSuppliedValues)
	fmt.Printf("\nONESITE ERRORS:= %v", errors)
	fmt.Printf("\nONESITE MSG:= %v", msg)
}
