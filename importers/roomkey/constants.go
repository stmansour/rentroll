package roomkey

import (
	"rentroll/rcsv"
)

// TempCSVStoreName holds the name of csvstore folder
var TempCSVStoreName = "temp_CSVs"

// TempCSVStore is used to store temporary csv files
var TempCSVStore string

// FieldDefaultValues is used to overwrite if user has not passed to values for these fields
var FieldDefaultValues = map[string]string{
	"ManageToBudget": "1", // always take to default this one
	"RentCycle":      "6", // maybe overridden by user supplied value
	"Proration":      "4", // maybe overridden by user supplied value
	"GSRPC":          "4", // maybe overridden by user supplied value
	"AssignmentTime": "1", // always take to default this one
	"Renewal":        "2", // always take to default this one
}

// prefixCSVFile is a map which holds the prefix of csv files
// so that temporarily program can create csv files with this
var prefixCSVFile = map[string]string{
	"rentable_types":   "rentableTypes_",
	"people":           "people_",
	"rental_agreement": "rentalAgreement_",
	"rentable":         "rentable_",
}

// RRRentableStatus is status for rentable in rentroll system
var RRRentableStatus = map[string]string{
	"unknown":        "0",
	"online":         "1",
	"admin":          "2",
	"employee":       "3",
	"owner occupied": "4",
	"offline":        "5",
}

// RentableStatusCSV is mapping for rentable status between roomkey and rentroll
var RentableStatusCSV = map[string]string{
	"vacant":   "online",
	"occupied": "online",
	"model":    "admin",
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

// define guest export column fields with order
const (
	GuestID            = iota
	CreateDate         = iota
	Title              = iota
	GuestName          = iota
	FirstName          = iota
	LastName           = iota
	Address            = iota
	City               = iota
	StateProvince      = iota
	Country            = iota
	CountryNationality = iota
	ZipPostalCode      = iota
	Email              = iota
	MainPhone          = iota
	Mobile             = iota
	PhoneOther         = iota
	RentalStatus       = iota
	VIPStatus          = iota
	VIPDescription     = iota
	GuestNote          = iota
	Air                = iota
	Auto               = iota
	Loyalty            = iota
	Language           = iota
	VehicleMake        = iota
	VehicleModel       = iota
	VehicleColor       = iota
	LicensePlate       = iota
	EmailMarketing     = iota
	EmailGeneral       = iota
	Address2Name       = iota
	Address2           = iota
	City2              = iota
	StateProvince2     = iota
	Country2           = iota
	ZipPostalCode2     = iota
	Email2             = iota
	MainPhone2         = iota
	Mobile2            = iota
	PhoneOther2        = iota
	RoomNights         = iota
	Stays              = iota
	AvgStay            = iota
	Revenue            = iota
	Reservations       = iota
	Cancellations      = iota
)

// defined csv columns
var guestCSVCols = []rcsv.CSVColumn{
	{Name: "Guest ID", Index: GuestID},
	{Name: "Create Date", Index: CreateDate},
	{Name: "Title", Index: Title},
	{Name: "Guest Name", Index: GuestName},
	{Name: "First Name", Index: FirstName},
	{Name: "Last Name", Index: LastName},
	{Name: "Address", Index: Address},
	{Name: "City", Index: City},
	{Name: "State/Province", Index: StateProvince},
	{Name: "Country", Index: Country},
	{Name: "Country of Nationality", Index: CountryNationality},
	{Name: "Zip/Postal Code", Index: ZipPostalCode},
	{Name: "Email", Index: Email},
	{Name: "Main Phone", Index: MainPhone},
	{Name: "Mobile", Index: Mobile},
	{Name: "Phone Other", Index: PhoneOther},
	{Name: "Rental Status", Index: RentalStatus},
	{Name: "VIP Status", Index: VIPStatus},
	{Name: "VIP Description", Index: VIPDescription},
	{Name: "Guest Note", Index: GuestNote},
	{Name: "Air #", Index: Air},
	{Name: "Auto #", Index: Auto},
	{Name: "Loyalty #", Index: Loyalty},
	{Name: "Language", Index: Language},
	{Name: "Vehicle Make", Index: VehicleMake},
	{Name: "Vehicle Model", Index: VehicleModel},
	{Name: "Vehicle Color", Index: VehicleColor},
	{Name: "License Plate #", Index: LicensePlate},
	{Name: "Email Marketing", Index: EmailMarketing},
	{Name: "Email General", Index: EmailGeneral},
	{Name: "Address 2 Name", Index: Address2Name},
	{Name: "Address 2", Index: Address2},
	{Name: "City 2", Index: City2},
	{Name: "State/Province 2", Index: StateProvince2},
	{Name: "Country 2", Index: Country2},
	{Name: "Zip/Postal 2", Index: ZipPostalCode2},
	{Name: "Email 2", Index: Email2},
	{Name: "Main Phone 2", Index: MainPhone2},
	{Name: "Mobile 2", Index: Mobile2},
	{Name: "Phone Other 2", Index: PhoneOther2},
	{Name: "Room Nights", Index: RoomNights},
	{Name: "Stays", Index: Stays},
	{Name: "Avg. Stay", Index: AvgStay},
	{Name: "Revenue", Index: Revenue},
	{Name: "Reservations", Index: Reservations},
	{Name: "Cancellations", Index: Cancellations},
}

// CSVLoadHandler struct is for routines that want to table-ize their loading.
type csvLoadHandler struct {
	Fname        string
	Handler      func(string) []error
	TraceDataMap string
	DBType       int
}

// minimum & maximum subsequent blank columns (no. of commas) in roomkey CSV headers
var minBlankColumns = 3
var maxBlankColumns = 5

// dummy column name to be replaced for blank column headers
var dummyBlankColumnName = "_BLANK_"

var dupTransactantWithPrimaryEmail = "PrimaryEmail"
var dupTransactantWithCellPhone = "CellPhone"

// will be used exact before rowIndex to format Notes in people csv "roomkey:<rowIndex>"
const (
	roomkeyNotesPrefix = "roomkey:"
	tcidPrefix         = "TC000"
)

// this slice contains list of strings which should be discarded
// used in csvRecordsToSkip function
var csvRecordsSkipList = []string{
	rcsv.DupTransactant,
	rcsv.DupRentableType,
	rcsv.DupCustomAttribute,
	rcsv.DupRentable,
	rcsv.RentableAlreadyRented,
}
