package core

import (
	"strconv"
)

// constants for csv types
const (
	// RENTABLETYPECSV NO
	RENTABLETYPECSV = iota
	// CUSTOMATTRIUTESCSV NO
	CUSTOMATTRIUTESCSV = iota
	// PEOPLECSV NO
	PEOPLECSV = iota
	// RENTABLECSV NO
	RENTABLECSV = iota
	// RENTALAGREEMENTCSV NO
	RENTALAGREEMENTCSV = iota
)

// const for db types
const (
	DBCustomAttr      = iota
	DBRentableType    = iota
	DBCustomAttrRef   = iota
	DBPeople          = iota
	DBRentable        = iota
	DBRentalAgreement = iota
)

// DBTypeMapStrings holds dbtype int to string format
var DBTypeMapStrings = map[int]string{
	DBCustomAttr:      strconv.Itoa(DBCustomAttr),
	DBRentableType:    strconv.Itoa(DBRentableType),
	DBCustomAttrRef:   strconv.Itoa(DBCustomAttrRef),
	DBPeople:          strconv.Itoa(DBPeople),
	DBRentable:        strconv.Itoa(DBRentable),
	DBRentalAgreement: strconv.Itoa(DBRentalAgreement),
	-1:                "",
}

// DBTypeMap holds db type name to count
var DBTypeMap = map[int]string{
	DBCustomAttr:      "Custom Attributes",
	DBRentableType:    "Rentable Types",
	DBCustomAttrRef:   "Custom Attribute References",
	DBPeople:          "Transactants",
	DBRentable:        "Rentables",
	DBRentalAgreement: "Rental Agreements",
}
