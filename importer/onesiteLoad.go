package main

import (
	"flag"
	"fmt"
	"rentroll/importer/onesite"
	"time"
)

var oneSiteCSVFile string

func readCommandLineArgs() {
	fp := flag.String("i", "onesite/onesite.csv", "the name of the onesite CSV file to import")
	flag.Parse()
	oneSiteCSVFile = *fp
}

func main() {
	fmt.Printf("%s\n", time.Now())
	readCommandLineArgs()
	onesite.Init()
	// try to open sample.csv file
	rs := onesite.LoadOneSiteCSV(oneSiteCSVFile)
	fmt.Printf("%s", rs)
}
