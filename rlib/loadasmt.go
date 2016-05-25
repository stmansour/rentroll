package rlib

import (
	"fmt"
	"strings"
)

// CreateAssessmentType reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentType(sa []string, lineno int) {
	funcname := "CreateAssessmentType"
	des := strings.TrimSpace(sa[0])
	if strings.ToLower(des) == "name" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Check to see if this assessment type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		a1, _ := GetAssessmentTypeByName(des)
		if len(a1.Name) > 0 {
			Ulog("%s: AssessmentType named %s already exists\n", funcname, des)
			return
		}
	}

	var a AssessmentType
	a.Name = strings.TrimSpace(sa[0])
	if len(a.Name) == 0 {
		Ulog("%s: line %d - Name cannot be empty\n", funcname, lineno)
		return
	}

	//-------------------------------------------------------------------
	// OccupancyRqd
	//-------------------------------------------------------------------
	a.OccupancyRqd, _ = IntFromString(sa[1], "OccupancyRqd value is invalid")
	if a.OccupancyRqd < 0 || a.OccupancyRqd > 1 {
		fmt.Printf("%s: line %d - OccupancyRqd must be 0 or 1.  Found: %s\n", funcname, lineno, sa[1])
		return
	}

	a.Description = sa[2]
	Errlog(InsertAssessmentType(&a))
}

// LoadAssessmentTypesCSV loads a csv file with assessment types and processes each one
func LoadAssessmentTypesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateAssessmentType(t[i], i+1)
	}
}
