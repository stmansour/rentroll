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
	// RARequired
	//-------------------------------------------------------------------
	a.RARequired, _ = IntFromString(sa[1], "RARequired value is invalid")
	if a.RARequired < RARQDINRANGE || a.RARequired > RARQDANY {
		fmt.Printf("%s: line %d - RARequired must be in the range %d to %d.  Found: %s\n", funcname, lineno, RARQDINRANGE, RARQDANY, sa[1])
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
