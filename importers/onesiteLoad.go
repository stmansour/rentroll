package main

import (
	"flag"
	"fmt"
	"rentroll/importers/onesite"
)

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
	return defaults
}

// MergeSuppliedAndDefaultValues used to merge
// override values from userSuppliedValues map into matched
// field of Defaults
func MergeSuppliedAndDefaultValues() {
	defaults := GetOneSiteFieldDefaultValues()

	// override default values to userSuppliedValues map
	// if not passed
	for k, _ := range userSuppliedValues {
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

	// parse the values from command line
	flag.Parse()

	// ================================
	// check for values which must be required
	// ================================
	if *fp == "" {
		ok = false
	}
	errors = append(errors, "Please, pass onesite csv input file")

	if *bud == "" {
		ok = false
	}
	errors = append(errors, "Please, pass business unit designation")

	// if not ok then return with errors, otherwise fill up values in map
	if !ok {
		return ok, errors
	} else {
		userSuppliedValues["OneSiteCSV"] = *fp
		userSuppliedValues["BUD"] = *bud
		userSuppliedValues["RentCycle"] = *frequency
		userSuppliedValues["Proration"] = *proration
		userSuppliedValues["GSRPC"] = *gsrpc
	}
	return ok, errors
}

func main() {
	// read command line argument first
	ok, inputErrors := readCommandLineArgs()
	if !ok {
		fmt.Printf("%v\n", inputErrors)
		return
	}

	// merge user supplied values with default one
	MergeSuppliedAndDefaultValues()

	// call onesite loader
	errors, msg := onesite.CSVHandler(userSuppliedValues)
	fmt.Println(errors)
	fmt.Println(msg)
}
