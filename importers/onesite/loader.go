package onesite

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
	"path"
	"reflect"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// constants used for onesite importer
const (
	SplittedCSVStore = "/tmp/onesite"
)

// GetOneSiteMapping reads json file and loads
// field mapping structure in go for further usage
func GetOneSiteMapping(OneSiteFieldMap *OneSiteMap) error {

	// Caller returns program counter, filename, line no, ok
	_, filename, _, ok := runtime.Caller(1)
	if ok == false {
		panic("Unable to get current filename")
	}

	// read json file which contains mapping of onesite fields
	mapperFilePath := path.Join(path.Dir(filename), "mapper.json")
	fieldmap, err := ioutil.ReadFile(mapperFilePath)
	if err != nil {
		// fmt.Errorf("File error: %v\n", err)
		panic(err)
		// return err
		// ???
	}
	err = json.Unmarshal(fieldmap, OneSiteFieldMap)
	if err != nil {
		// fmt.Errorf("%s", err)
		panic(err)
	}
	return err
}

// LoadOneSiteCSV loads the values from the supplied csv file and creates rlib.Business records
// as needed.
func LoadOneSiteCSV(fname string) string {
	rs := ""

	// this count used to skip number of rows from the very top of csv
	var skipRowsCount int

	// TODO: dynamically detect to skip no of rows by checking headers value

	// define column fields with order
	const (
		Unit            = iota
		FloorPlan       = iota
		UnitDesignation = iota
		Sqft            = iota
		UnitLeaseStatus = iota
		Name            = iota
		PhoneNumber     = iota
		Email           = iota
		MoveIn          = iota
		NoticeForDate   = iota
		MoveOut         = iota
		LeaseStart      = iota
		LeaseEnd        = iota
		MarketAddl      = iota
		DepOnHand       = iota
		Balance         = iota
		TotalCharges    = iota
		Rent            = iota
		WaterReImb      = iota
		Corp            = iota
		Discount        = iota
		Platinum        = iota
		Tax             = iota
		ElectricReImb   = iota
		Fire            = iota
		ConcSpecl       = iota
		WashDry         = iota
		EmplCred        = iota
		Short           = iota
		PetFee          = iota
		TrashReImb      = iota
		TermFee         = iota
		LakeView        = iota
		Utility         = iota
		Furn            = iota
		Mtom            = iota
		Referral        = iota
	)

	// defined csv columns
	var csvCols = []rcsv.CSVColumn{
		{Name: "Unit", Index: Unit},
		{Name: "FloorPlan", Index: FloorPlan},
		{Name: "UnitDesignation", Index: UnitDesignation},
		{Name: "SQFT", Index: Sqft},
		{Name: "Unit/LeaseStatus", Index: UnitLeaseStatus},
		{Name: "Name", Index: Name},
		{Name: "PhoneNumber", Index: PhoneNumber},
		{Name: "Email", Index: Email},
		{Name: "Move-In", Index: MoveIn},
		{Name: "NoticeForDate", Index: NoticeForDate},
		{Name: "Move-Out", Index: MoveOut},
		{Name: "LeaseStart", Index: LeaseStart},
		{Name: "LeaseEnd", Index: LeaseEnd},
		{Name: "Market+Addl.", Index: MarketAddl},
		{Name: "DepOnHand", Index: DepOnHand},
		{Name: "Balance", Index: Balance},
		{Name: "TotalCharges", Index: TotalCharges},
		{Name: "RENT", Index: Rent},
		{Name: "WATERREIMB", Index: WaterReImb},
		{Name: "CORP", Index: Corp},
		{Name: "DISCOUNT", Index: Discount},
		{Name: "Platinum", Index: Platinum},
		{Name: "TAX", Index: Tax},
		{Name: "ELECTRICREIMB", Index: ElectricReImb},
		{Name: "Fire", Index: Fire},
		{Name: "CONC/SPECL", Index: ConcSpecl},
		{Name: "WASH/DRY", Index: WashDry},
		{Name: "EMPLCRED", Index: EmplCred},
		{Name: "SHORT", Index: Short},
		{Name: "PETFEE", Index: PetFee},
		{Name: "TRASHREIMB", Index: TrashReImb},
		{Name: "TERMFEE", Index: TermFee},
		{Name: "Lakeview", Index: LakeView},
		{Name: "UTILITY", Index: Utility},
		{Name: "FURN", Index: Furn},
		{Name: "MTOM", Index: Mtom},
		{Name: "REFERRAL", Index: Referral},
	}

	// Onesite csv headers slice and load it from csvCols
	OneSiteCSVHeaders := []string{}
	for _, header := range csvCols {
		OneSiteCSVHeaders = append(OneSiteCSVHeaders, header.Name)
	}
	OneSiteColumnLength := len(OneSiteCSVHeaders)

	// load csv file and get data from csv
	t := rlib.LoadCSV(fname)

	// ================================
	// First loop for validation on csv
	// ================================
	funcname := "LoadOneSiteCSV"

	for i := 0; i < len(t); i++ {

		// Calculate SkipRowsCount in first loop
		// if found then assign value in it and for rest of the rows
		// do validate csv columns

		CSVRowDataString := strings.Replace(
			strings.Join(t[i][:OneSiteColumnLength], ","),
			" ", "", -1)
		CSVHeaderString := strings.Replace(
			strings.Join(OneSiteCSVHeaders[:OneSiteColumnLength], ","),
			" ", "", -1)

		if CSVRowDataString == CSVHeaderString {
			skipRowsCount = i
		}

		// if skipRowsCount found then do validation over csv rows
		if skipRowsCount > 0 {
			x, vrs := rcsv.ValidateCSVColumns(csvCols, t[i][:OneSiteColumnLength], funcname, i)
			if x > 0 {
				// return vrs, 1
				fmt.Println(vrs)
				return vrs.Error()
			}
		}
	}

	// ================================
	// Second loop for splitting data of csv
	// Create csv files required for rentroll
	// ================================
	// in second round do split

	// TODO: create all csv file here(rentable, custom attribute, people, rentable aggrement)

	// get current timestamp used for creating csv files unique way
	currentTime := time.Now()

	// RFC3339Nano is const format defined in time package
	// <FORMAT> = <SAMPLE>
	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// it is helpful while creating unique files
	currentTimeFormat := currentTime.Format(time.RFC3339Nano)

	// get onesite mapping
	var OneSiteFieldMap OneSiteMap
	err := GetOneSiteMapping(&OneSiteFieldMap)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err.Error()
	}
	fmt.Printf("%v", OneSiteFieldMap)

	// ======================================
	// create rentabletypes file with current time
	// ======================================
	rentableTypeCSVFilePath := SplittedCSVStore + "/rentabletypes_" + currentTimeFormat + ".csv"
	rentableTypeCSVFile, err := os.Create(rentableTypeCSVFilePath)
	if err != nil {
		panic(fmt.Sprintf("Error while creating file %s with error (%s)", rentableTypeCSVFilePath, err))
	}
	defer rentableTypeCSVFile.Close()

	// create csv writer
	rentableTypeCSVWriter := csv.NewWriter(rentableTypeCSVFile)

	// TODO: write headers from struct rather than provides hard coded values
	rentableTypeCSVHeaders := []string{
		"BUD", "Style", "Name",
		"RentCycle", "Proration",
		"GSRPC", "ManageToBudget",
		"MarketRate", "DtStart", "DtStop",
	}
	rentableTypeCSVWriter.Write(rentableTypeCSVHeaders)

	// avoidDuplicateRentableTypeData used to keep track of rentableTypeData with Style field
	// so that duplicate entries can be avoided while creating rentableType csv file
	avoidDuplicateRentableTypeData := []string{}

	for i := skipRowsCount + 1; i < len(t); i++ {
		rowLoaded, csvRow := LoadOneSiteCSVRow(csvCols, t[i][:OneSiteColumnLength])

		// NOTE: might need to change logic, if t[i] contains blank data that we should
		// stop the loop as we have to skip rest of the rows (please look at onesite.csv)
		if !rowLoaded {
			fmt.Println("\nNo more data to parse")
			break
		}

		checkRentableTypeStyle := csvRow.FloorPlan
		Stylefound := core.StringInSlice(checkRentableTypeStyle, avoidDuplicateRentableTypeData)
		if Stylefound {
			// jump to next row
			// TODO: remove this jump as in future we need to include
			// other type of parsing to create csv files
			continue
		} else {
			avoidDuplicateRentableTypeData = append(avoidDuplicateRentableTypeData, checkRentableTypeStyle)
		}

		currentYear, _, _ := currentTime.Date()
		DtStart := "1/1/" + strconv.Itoa(currentYear)
		DtStop := "1/1/" + strconv.Itoa(currentYear+1)
		// DtStart := time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentTime.Location())
		// DtStop := time.Date(currentYear+1, 1, 1, 0, 0, 0, 0, currentTime.Location())

		// Create rentabletype csv file
		// takes csvRow data, rentableType mapping between fields, current time

		// TODO: take values from user which have been supplied and for no values
		// which have not been supplied take default values
		userSuppliedValues := map[string]string{}
		userSuppliedValues["BUD"] = "REX"
		userSuppliedValues["RentCycle"] = "6"
		userSuppliedValues["Proration"] = "4"
		userSuppliedValues["GSRPC"] = "4"
		userSuppliedValues["ManageToBudget"] = "1"
		userSuppliedValues["DtStart"] = DtStart
		userSuppliedValues["DtStop"] = DtStop
		ok, rentableTypeInstance, rentableTypeData := GetRentableTypeCSVRow(
			&csvRow, &OneSiteFieldMap.RentableTypeCSV,
			currentTimeFormat, userSuppliedValues,
		)
		fmt.Println(ok)
		fmt.Println(rentableTypeInstance)
		fmt.Println(rentableTypeData)
		rentableTypeCSVWriter.Write(rentableTypeData)
	}
	rentableTypeCSVWriter.Flush()
	return rs
}

// RollBackSplitOperation func used to clear out the things
// that created by program temporarily while loading onesite data
//  and if any error occurs
func RollBackSplitOperation(timestamp string) {
	panic("RollBackSplitOperation")
}

// GetRentableTypeCSVRow used to create rentabletype
// csv from onesite csv data to dump data via rcsv routine
func GetRentableTypeCSVRow(
	oneSiteRow *OneSiteCSVRow,
	fieldMap *core.RentableTypeCSV,
	timestamp string,
	userSuppliedValues map[string]string,
) (bool, *core.RentableTypeCSV, []string) {

	// take initial variable
	var rentableType core.RentableTypeCSV
	ok := false

	// ======================================
	// Load rentableType's data from onesiterow data
	// ======================================
	reflectedRentableType := reflect.ValueOf(&rentableType).Elem()
	reflectedOneSiteRow := reflect.ValueOf(oneSiteRow).Elem()
	reflectedFieldMap := reflect.ValueOf(fieldMap).Elem()

	// length of reflectedRentableType
	rRTLength := reflectedRentableType.NumField()

	// return data array
	dataMap := make(map[int]string)

	for i := 0; i < rRTLength; i++ {
		// get rentableType field
		rentableTypeField := reflectedRentableType.Type().Field(i)

		// if rentableTypeField value exist in userSuppliedValues map
		// then set it first
		suppliedValue, ok := userSuppliedValues[rentableTypeField.Name]
		if ok {
			reflectedRentableType.Field(i).Set(reflect.ValueOf(suppliedValue))
			dataMap[i] = suppliedValue
		}

		// get mapping field if not found then panic error
		MappedFieldName, ok := reflectedFieldMap.FieldByName(rentableTypeField.Name).Interface().(string)
		if !ok {
			panic("coudln't get mapping field")
		}

		// if has not value then continue
		if !reflectedOneSiteRow.FieldByName(MappedFieldName).IsValid() {
			continue
		}

		// get field by mapping field name and then value
		OneSiteFieldValue := reflectedOneSiteRow.FieldByName(MappedFieldName).Interface()

		dataMap[i] = OneSiteFieldValue.(string)

		// set the value if possible
		if reflectedRentableType.FieldByName(rentableTypeField.Name).CanSet() {
			reflectedRentableType.FieldByName(rentableTypeField.Name).Set(
				reflect.ValueOf(OneSiteFieldValue))
		}
		fmt.Println("=======================")
		fmt.Printf("RentableTypeFieldName: %s\t", rentableTypeField.Name)
		fmt.Printf("MappedFieldName: %s\t", MappedFieldName)
		fmt.Printf("Value: %s\n", OneSiteFieldValue)
	}

	dataArray := []string{}

	for i := 0; i < rRTLength; i++ {
		dataArray = append(dataArray, dataMap[i])
	}

	return ok, &rentableType, dataArray
}

// Init configure required settings
func Init() {
	if _, err := os.Stat(SplittedCSVStore); os.IsNotExist(err) {
		os.MkdirAll(SplittedCSVStore, 0700)
	}
}
