package main

import (
	"fmt"
	"rentroll/importer/onesite"
	"time"
)

func main() {
	fmt.Printf("%s\n", time.Now())
	onesite.Init()
	oneSiteCSVFile := "onesite/onesite.csv"
	// try to open sample.csv file
	rs := onesite.LoadOneSiteCSV(oneSiteCSVFile)
	fmt.Printf("%s", rs)
}
