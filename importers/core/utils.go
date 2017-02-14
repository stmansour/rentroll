package core

import (
	"regexp"
	"rentroll/rlib"
)

// StringInSlice used to check whether string a
// is present or not in slice list
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IntegerInSlice used to check whether int a
// is present or not in slice list
func IntegerInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsValidEmail used to check valid email or not
func IsValidEmail(email string) bool {
	// TODO: confirm which regex to use
	// Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+$`)
	return Re.MatchString(email)
}

// GetImportedCount get map of summaryCount as an argument
// then it hit db to get imported count for each type
func GetImportedCount(summaryCount map[int]map[string]int, BID int64) {
	for dbType := range summaryCount {
		switch dbType {
		case DBCustomAttrRef:
			summaryCount[DBCustomAttrRef]["imported"] += rlib.GetCountBusinessCustomAttrRefs(BID)
			break
		case DBCustomAttr:
			summaryCount[DBCustomAttr]["imported"] += rlib.GetCountBusinessCustomAttributes(BID)
			break
		case DBRentableType:
			summaryCount[DBRentableType]["imported"] += rlib.GetCountBusinessRentableTypes(BID)
			break
		case DBPeople:
			summaryCount[DBPeople]["imported"] += rlib.GetCountBusinessTransactants(BID)
			break
		case DBRentable:
			summaryCount[DBRentable]["imported"] += rlib.GetCountBusinessRentables(BID)
			break
		case DBRentalAgreement:
			summaryCount[DBRentalAgreement]["imported"] += rlib.GetCountBusinessRentalAgreements(BID)
			break
		}
	}
}
