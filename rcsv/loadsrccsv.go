package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// CSV FIELDS FOR THIS MODULE
//    0    1     2
//    BUD, Name, Industry
//    REX, FAA,  Aviation
//    REX, IRS,  Excessive Rules

// CreateSourceCSV reads an assessment type string array and creates a database record for the assessment type
func CreateSourceCSV(sa []string, lineno int) (string, int) {
	funcname := "CreateSourceCSV"
	var a rlib.DemandSource
	var err error

	const (
		BUD     = 0
		Name    = iota
		Industy = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"Name", Name},
		{"Industry", Industy},
	}

	rs, x := ValidateCSVColumns(csvCols, sa, funcname, lineno)
	if x > 0 {
		return rs, 1
	}
	if lineno == 1 {
		return rs, 0
	}

	des := strings.ToLower(strings.TrimSpace(sa[BUD]))

	//-------------------------------------------------------------------
	// Business
	//-------------------------------------------------------------------
	var b rlib.Business
	if len(des) > 0 {
		b = rlib.GetBusinessByDesignation(des)
		if b.BID < 1 {
			rlib.Ulog("CreateRentalSpecialtyType: rlib.Business named %s not found\n", sa[BUD])
			return rs, CsvErrorSensitivity
		}
	}
	a.BID = b.BID

	//-------------------------------------------------------------------
	// Name
	//-------------------------------------------------------------------
	s := strings.TrimSpace(sa[Name])
	if len(s) > 0 {
		var src rlib.DemandSource
		rlib.GetDemandSourceByName(b.BID, s, &src)
		if len(src.Name) > 0 {
			fmt.Printf("%s: line %d - DemandSource named %s already exists.\n", funcname, lineno, s)
			return rs, CsvErrorSensitivity
		}
	}
	a.Name = s

	//-------------------------------------------------------------------
	// Industry
	//-------------------------------------------------------------------
	a.Industry = strings.TrimSpace(sa[Industy])

	_, err = rlib.InsertDemandSource(&a)
	if err != nil {
		fmt.Printf("%s: line %d - error inserting DemandSource: %v\n", funcname, lineno, err)
		return rs, CsvErrorSensitivity
	}

	return rs, 0
}

// LoadSourcesCSV loads a csv file with a chart of accounts and creates rlib.GLAccount markers for each
func LoadSourcesCSV(fname string) string {
	rs := ""
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		s, err := CreateSourceCSV(t[i], i+1)
		rs += s
		if err > 0 {
			break
		}
	}
	return rs
}
