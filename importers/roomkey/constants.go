package roomkey

import (
	"rentroll/rcsv"
)

// splittedCSVStoreName holds the name of csvstore folder
var splittedCSVStoreName = "temp_CSVs"

// prefixCSVFile is a map which holds the prefix of csv files
// so that temporarily program can create csv files with this
var prefixCSVFile = map[string]string{
	"rentable_types":   "rentableTypes_",
	"people":           "people_",
	"rental_agreement": "rentalAgreement_",
	"rentable":         "rentable_",
}

// define column fields with order
const (
	Empty1         = iota
	Guest          = iota
	Description    = iota
	ResID          = iota
	DateRes        = iota
	DateIn         = iota
	Empty3         = iota
	DateOut        = iota
	Adults         = iota
	Child          = iota
	Room           = iota
	RoomType       = iota
	Rate           = iota
	RateName       = iota
	GroupCorporate = iota
)

// defined csv columns
var csvCols = []rcsv.CSVColumn{
	{Name: "", Index: Empty1},
	{Name: "Guest", Index: Guest},
	{Name: "", Index: Description},
	{Name: "Res. ID", Index: ResID},
	{Name: "DateRes", Index: DateRes},
	{Name: "DateIn", Index: DateIn},
	{Name: "", Index: Empty3},
	{Name: "DateOut", Index: DateOut},
	{Name: "Adults", Index: Adults},
	{Name: "Child.", Index: Child},
	{Name: "Room", Index: Room},
	{Name: "RoomType", Index: RoomType},
	{Name: "Rate", Index: Rate},
	{Name: "RateName", Index: RateName},
	{Name: "Group/Corporate Name", Index: GroupCorporate},
}

// minimum & maximum subsequent blank columns (no. of commas) in roomkey CSV headers
var minBlankColumns = 3
var maxBlankColumns = 5

// dummy column name to be replaced for blank column headers
var dummyBlankColumnName = "_BLANK_"
