package rlib

import "strings"

// CreateAssessmentType reads an assessment type string array and creates a database record for the assessment type
func CreateAssessmentType(sa []string) {
	des := strings.TrimSpace(sa[0])
	if des == "assessment type" || des == "assessmenttype" {
		return // this is just the column heading
	}

	//-------------------------------------------------------------------
	// Check to see if this assessment type is already in the database
	//-------------------------------------------------------------------
	if len(des) > 0 {
		a1, _ := GetAssessmentTypeByName(des)
		if len(a1.Name) > 0 {
			Ulog("CreateAssessmentType: AssessmentType named %s already exists\n", des)
			return
		}
	}

	var a AssessmentType
	a.Name = strings.TrimSpace(sa[0])
	a.Description = sa[1]
	Errlog(InsertAssessmentType(&a))
}

// LoadAsessmentTypesCSV loads a csv file with assessment types and processes each one
func LoadAsessmentTypesCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreateAssessmentType(t[i])
	}
}
