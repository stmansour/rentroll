// Package onesite contains this program where data actually
// being imported from csv to rentroll database.

// Main program call `CSVHandler` function to do the actual job.
// `CSVHandler` calls main function `loadOneSiteCSV` and
// then creates a report based on response of `loadOneSiteCSV` call.

// `loadOneSiteCSV` writes data in CSVs and loads with help of rcsv
// loaders and return the response to `CSVHandler`.

package onesite

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"rentroll/importers/core"
	"rentroll/rcsv"
	"rentroll/rlib"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kardianos/osext"
)

// getOneSiteMapping reads json file and loads
// field mapping structure in go for further usage
func getOneSiteMapping(oneSiteFieldMap *CSVFieldMap) error {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	// read json file which contains mapping of onesite fields
	mapperFilePath := path.Join(folderPath, "mapper.json")

	fieldmap, err := ioutil.ReadFile(mapperFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fieldmap, oneSiteFieldMap)
	return err
}

// loadOneSiteCSV loads the values from the supplied csv file and
// creates rlib.Business records as needed
//
// INPUTS
//    ctx
//    oneSiteCSV
//    testMode
//    userRRValues
//    business
//    currentTime
//    currentTimeFormat
//    summaryReport
//
// RETURNS unitmap, csvError list, internal error, csv loaded?
//--------------------------------------------------------------------------
func loadOneSiteCSV(
	ctx context.Context,
	oneSiteCSV string,
	testMode int,
	userRRValues map[string]string,
	business *rlib.Business,
	currentTime time.Time,
	currentTimeFormat string,
	summaryReport map[int]map[string]int,
) (map[int]string, map[int][]string, bool) {

	// returns unitmap, csvError list, internal error, csv loaded?

	// returned csv errors should be in format
	// {
	// 	"rowIndex": ["E:errors",....., "W:warnings",....]
	// }
	// E stands for Error string, W stands for Warning string
	// UnitName can be accessible via traceUnitMap

	// =========================
	// DATA STRUCTURES AND VARS
	// =========================

	internalErrFlag := true
	csvErrors := map[int][]string{}

	//------------------------------------------------------------------------
	// this count used to skip number of rows from the very top of csv
	//------------------------------------------------------------------------
	var skipRowsCount int
	var rowIndex int

	//------------------------------------------------------------------------
	// customAttributesRefData holds the data after customAttr insertion
	// to insert custom attribute ref in system for each rentableType
	// so we identify each element in this list with Style Key
	//------------------------------------------------------------------------
	customAttributesRefData := map[string]CARD{}

	//------------------------------------------------------------------------
	// this map is used to hold csvRow typed struct after data has been loaded in it
	// by iterating over csv data, re-usable for rentable, rental agreement data
	//------------------------------------------------------------------------
	csvRowDataMap := map[int]*CSVRow{}

	// --------------------------- trace data map ---------------------------- //
	// trace<TYPE>CSVMap used to hold records
	// by which we can traceout which records has been writtern to csv
	// with key of row index of <TARGET_TYPE> CSV, value of original's imported csv rowNumber
	//------------------------------------------------------------------------
	traceRentableTypeCSVMap := map[int]int{}
	traceRentableCSVMap := map[int]int{}
	tracePeopleCSVMap := map[int]int{}
	traceRentalAgreementCSVMap := map[int]int{}
	traceCustomAttributeCSVMap := map[int]int{}

	//------------------------------------------------------------------------
	// traceTCIDMap hold TCID for each people to be loaded via people csv
	// with reference of original onesite csv
	//------------------------------------------------------------------------
	traceTCIDMap := map[int]string{}

	// traceUnitMap holds records by which we can trace the unit with row index of csv
	// Unit would be unique in onesite imported csv
	// key: rowIndex of onesite csv, value: Unit value of each row of onesite csv
	traceUnitMap := map[int]string{}

	// traceDuplicatePeople holds records with unique string (name, email, phone)
	// with duplicant match at row
	// e.g.; {
	// 	"phone": {"9999999999": [2,4]},
	// 	"name": {"foo, bar": 3},
	// }
	//------------------------------------------------------------------------
	traceDuplicatePeople := map[string][]string{
		"name": {}, "phone": {},
	}

	// --------------------- avoid duplicate data structures -------------------- //
	// avoidDuplicateRentableTypeData used to keep track of rentableTypeData with Style field
	// so that duplicate entries can be avoided while creating rentableType csv file
	//------------------------------------------------------------------------
	avoidDuplicateRentableTypeData := []string{}

	//------------------------------------------------------------------------
	// avoidDuplicateCustomAttributeData is tricky map which holds the
	// duplicate data in slice for each field defined in customAttributeMap
	//------------------------------------------------------------------------
	avoidDuplicateCustomAttributeData := map[string][]string{}
	for k := range customAttributeMap {
		avoidDuplicateCustomAttributeData[k] = []string{}
	}

	// --------------------------- csv record count ----------------------------
	// <TYPE>CSVRecordCount used to hold records count inserted in csv
	// initialize with 1 because first row contains headers in target generated csv
	// these are POSSIBLE record count that going to be imported
	//------------------------------------------------------------------------
	RentableTypeCSVRecordCount := 0
	RentableCSVRecordCount := 0
	PeopleCSVRecordCount := 0
	RentalAgreementCSVRecordCount := 0
	CustomAttributeCSVRecordCount := 0
	CustomAttrRefRecordCount := 0

	// ================================================
	// LOAD FIELD MAP AND GET HEADERS, LENGTH OF HEADERS
	// ================================================

	// load onesite mapping
	var oneSiteFieldMap CSVFieldMap
	err := getOneSiteMapping(&oneSiteFieldMap)
	if err != nil {
		// rlib.Console("loadOneSiteCSV: error A\n")
		rlib.Ulog("INTERNAL ERROR <ONESITE FIELD MAPPING>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// ==============================
	// COUNT ROWS NEEDS TO BE SKIPPED
	// ==============================

	// load csv file and get data from csv
	t := rlib.LoadCSV(oneSiteCSV)

	csvHeadersIndex := getCSVHeadersIndexMap()

	//------------------------------------------------------------------
	// detect how many rows we need to skip first
	//------------------------------------------------------------------
	for rowIndex := 0; rowIndex < len(t); rowIndex++ {
		for colIndex := 0; colIndex < len(t[rowIndex]); colIndex++ {
			// remove all white space and make lower case
			cellTextValue := strings.ToLower(
				core.SpecialCharsReplacer.Replace(t[rowIndex][colIndex]))

			// ********************************
			// MARKET RENT OR MARKET ADDL
			// ********************************
			// if marketRent found then remove marketAddl header
			// and make an entry for "marketrent" in csvColumnFieldMap with -1
			// keep "MarketAddl" mapping to `marketrent` still, anyways `MarketAddl`
			// going to be put in `MarketRate` of Rentroll field
			//------------------------------------------------------------------
			if cellTextValue == marketRent {
				delete(csvColumnFieldMap, "marketaddl")
				csvColumnFieldMap[marketRent] = "MarketAddl"
			}

			//------------------------------------------------------------------
			// assume that cellTextValue is a header. Check to see if it exists
			// in the map. If so, set its column index
			//------------------------------------------------------------------
			if field, ok := csvColumnFieldMap[cellTextValue]; ok {
				csvHeadersIndex[field] = colIndex
			}
		}
		//------------------------------------------------------------------
		// check after row columns parsing that headers are found or not
		//------------------------------------------------------------------
		headersFound := true
		for _, v := range csvHeadersIndex {
			if v == -1 {
				headersFound = false
				break
			}
		}

		if headersFound {
			//------------------------------------------------------------------
			// update rowIndex by 1 because we're going to break here
			//------------------------------------------------------------------
			rowIndex++
			skipRowsCount = rowIndex
			break
		}
	}

	//------------------------------------------------------------------
	// if skipRowsCount is still 0 that means data could not be parsed from csv
	//------------------------------------------------------------------
	if skipRowsCount == 0 {
		missingHeaders := []string{}

		//------------------------------------------------------------------
		// make message of missing columns
		//------------------------------------------------------------------
		for missedH, v := range csvHeadersIndex {
			if v == -1 {
				missingHeaders = append(missingHeaders, missedH)
			}
		}

		headerError := "Required data column(s) missing: "
		headerError += strings.Join(missingHeaders, ", ")

		// ******** special entry ***********
		// -1 means there is no data
		//------------------------------------------------------------------
		internalErrFlag = false
		csvErrors[-1] = append(csvErrors[-1], headerError)
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// =================================
	// DELETE DATA RELATED TO BUSINESS ID
	// =================================
	// delete business related data before starting to import in database
	//------------------------------------------------------------------
	_, err = rlib.DeleteBusinessFromDB(ctx, business.BID)
	if err != nil {
		rlib.Ulog("INTERNAL ERROR <DELETE BUSINESS>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	bid, err := rlib.InsertBusiness(ctx, business)
	if err != nil {
		rlib.Ulog("INTERNAL ERROR <INSERT BUSINESS>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}
	//------------------------------------------------------------------
	// set new BID as we have deleted and inserted it again
	// TODO:  remove this step after sman's next push
	// in InsertBusiness it will be set automatically
	//------------------------------------------------------------------
	business.BID = bid

	// ========================================================
	// WRITE DATA FOR CUSTOM ATTRIBUTE, RENTABLE TYPE, PEOPLE CSV
	// ========================================================

	// get created customAttibutes csv and writer pointer
	customAttributeCSVFile, customAttributeCSVWriter, ok :=
		CreateCustomAttibutesCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.CustomAttributeCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <CUSTOM ATTRIUTE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// ----------------------- create files and get csv writer object -----------------------
	// get created rentabletype csv and writer pointer
	rentableTypeCSVFile, rentableTypeCSVWriter, ok :=
		CreateRentableTypeCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.RentableTypeCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTABLE TYPE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// get created people csv and writer pointer
	peopleCSVFile, peopleCSVWriter, ok :=
		CreatePeopleCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.PeopleCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <PEOPLE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	//------------------------------------------------------------------
	// if skipRowsCount found get next row and proceed on rest of the rows with loop
	//------------------------------------------------------------------
	for rowIndex = skipRowsCount + 1; rowIndex <= len(t); rowIndex++ {

		//------------------------------------------------------------------
		// if column order has been validated then only perform
		// data validation on value, type
		//------------------------------------------------------------------
		rowLoaded, csvRow := loadOneSiteCSVRow(csvHeadersIndex, t[rowIndex-1])

		// **************************************************************
		// NOTE: might need to change logic, if t[i] contains blank data that
		// we should stop the loop as we have to skip rest of the rows
		// (please look at onesite csv)
		// **************************************************************
		if !rowLoaded {

			// what IF, only headers are there
			if rowIndex == skipRowsCount {
				// ******** special entry ***********
				// -1 means there is no data
				internalErrFlag = false
				csvErrors[-1] = append(csvErrors[-1], "There are no data rows present")
				return traceUnitMap, csvErrors, internalErrFlag
			}

			// else break the loop as there are no more data
			break
		}

		// rowLoaded successfully then do rest of the operation

		// get rentable status from csv data
		csvRentableUseStatus := csvRow.UnitLeaseStatus

		// get unit from csv data
		csvUnit := csvRow.Unit

		//------------------------------------------------------------------
		// for rentable status exists in csvRow, get set of csv types which
		// can be allowed to perform write data for csv
		// need to call validation function as in to get the values
		//------------------------------------------------------------------
		_, rrUseType, _ := IsValidRentableUseType(csvRentableUseStatus)
		csvTypesSet := canWriteCSVStatusMap[rrUseType]
		var canWriteData bool

		// mark Unit value with row index value
		// even if it is blank
		traceUnitMap[rowIndex] = csvUnit

		// keep csv record in this
		csvRowDataMap[rowIndex] = &csvRow

		//------------------------------------------------------------------
		// check first that for this row's status rentableType data can be written
		//------------------------------------------------------------------
		canWriteData = core.IntegerInSlice(core.RENTABLETYPECSV, csvTypesSet)
		if canWriteData {
			//------------------------------------------------------------------
			// Write data to file of rentabletype
			//------------------------------------------------------------------
			WriteRentableTypeCSVData(
				&RentableTypeCSVRecordCount,
				rowIndex,
				traceRentableTypeCSVMap,
				rentableTypeCSVWriter,
				&csvRow,
				&avoidDuplicateRentableTypeData,
				currentTime,
				currentTimeFormat,
				userRRValues,
				&oneSiteFieldMap.RentableTypeCSV,
				customAttributesRefData,
				business,
			)
		}

		// check first that for this row's status custom attributes data can be written
		canWriteData = core.IntegerInSlice(core.CUSTOMATTRIUTESCSV, csvTypesSet)
		if canWriteData {
			// Write data to file of CustomAttribute
			WriteCustomAttributeData(
				&CustomAttributeCSVRecordCount,
				rowIndex,
				traceCustomAttributeCSVMap,
				customAttributeCSVWriter,
				&csvRow,
				avoidDuplicateCustomAttributeData,
				currentTimeFormat,
				userRRValues,
				&oneSiteFieldMap.CustomAttributeCSV,
			)
		}

		// check first that for this row's status people data can be written
		canWriteData = core.IntegerInSlice(core.PEOPLECSV, csvTypesSet)
		if canWriteData {

			// if people data can be writable then init TCIDMap index
			// with blank string value
			traceTCIDMap[rowIndex] = ""

			// Write data to file of people
			WritePeopleCSVData(
				&PeopleCSVRecordCount,
				rowIndex,
				tracePeopleCSVMap,
				peopleCSVWriter,
				&csvRow,
				traceDuplicatePeople,
				currentTimeFormat,
				userRRValues,
				&oneSiteFieldMap.PeopleCSV,
				csvErrors,
			)
		}

	}

	// Close all files as we are done here with writing data
	rentableTypeCSVFile.Close()
	peopleCSVFile.Close()
	customAttributeCSVFile.Close()

	// =======================
	// NESTED UTILITY FUNCTIONS
	// =======================

	// getTraceDataMap from string name
	getTraceDataMap := func(traceDataMapName string) map[int]int {
		switch traceDataMapName {
		case "traceCustomAttributeCSVMap":
			return traceCustomAttributeCSVMap
		case "traceRentableTypeCSVMap":
			return traceRentableTypeCSVMap
		case "tracePeopleCSVMap":
			return tracePeopleCSVMap
		case "traceRentableCSVMap":
			return traceRentableCSVMap
		case "traceRentalAgreementCSVMap":
			return traceRentalAgreementCSVMap
		default:
			return nil
		}
	}

	// getIndexAndUnit used to get index and unit value from trace<TYPE>CSVMap map
	getIndexAndUnit := func(traceDataMap map[int]int, index int) (int, string) {
		var onesiteIndex int
		var unit string
		if onesiteIndex, ok := traceDataMap[index]; ok {
			if unit, ok := traceUnitMap[onesiteIndex]; ok {
				return onesiteIndex, unit
			}
			return onesiteIndex, unit
		}
		return onesiteIndex, unit
	}

	// rrDoLoad is a nested function
	// used to load data from csv with help of rcsv loaders
	rrDoLoad := func(ctx context.Context, fname string, handler func(context.Context, string) []error, traceDataMapName string, dbType int) bool {
		Errs := handler(ctx, fname)

		for _, err := range Errs {
			// skip warnings about already existing records
			// if it's not kind of to skip then process it and count in error report
			errText := err.Error()

			if !csvRecordsToSkip(err) {
				lineNo, reason, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				onesiteIndex, _ := getIndexAndUnit(traceDataMap, lineNo)
				// generate new error
				csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
			} else {
				rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s\n", fname, errText)
			}
		}
		// return with success
		return true
	}

	// *****************************************************
	// rrPeopleDoLoad (SPECIAL METHOD TO LOAD PEOPLE)
	// *****************************************************
	rrPeopleDoLoad := func(ctx context.Context, fname string, handler func(context.Context, string) []error, traceDataMapName string, dbType int) bool {
		Errs := handler(ctx, fname)

		for _, err := range Errs {
			// handling for duplicant transactant
			if strings.Contains(err.Error(), dupTransactantWithPrimaryEmail) {
				lineNo, _, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				onesiteIndex, _ := getIndexAndUnit(traceDataMap, lineNo)
				// load csvRow from dataMap to get email
				csvRow := *csvRowDataMap[onesiteIndex]
				pEmail := csvRow.Email
				// get tcid from email
				t, tErr := rlib.GetTransactantByPhoneOrEmail(ctx, business.BID, pEmail)
				if tErr != nil {
					// unable to get TCID
					reason := "E:<" + core.DBTypeMapStrings[core.DBPeople] + ">:Unable to get people information" + err.Error()
					csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
				} else if t.TCID == 0 {
					// unable to get TCID
					reason := "E:<" + core.DBTypeMapStrings[core.DBPeople] + ">:Unable to get people information"
					csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
				} else {
					// if duplicate people found
					rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s", fname, err.Error())
					// map it in tcid map
					traceTCIDMap[onesiteIndex] = tcidPrefix + strconv.FormatInt(t.TCID, 10)
				}
			} else {
				lineNo, reason, ok := parseLineAndErrorFromRCSV(err, dbType)
				if !ok {
					// INTERNAL ERROR - RETURN FALSE
					return false
				}
				// get tracedatamap
				traceDataMap := getTraceDataMap(traceDataMapName)
				// now get the original row index of imported onesite csv and Unit value
				onesiteIndex, _ := getIndexAndUnit(traceDataMap, lineNo)
				// generate new error
				csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
			}

			// *****************************************************************
			// AS WE DON'T HAVE MAPPING OF PHONENUMBER TO CELLPHONE
			// WE JUST AVOID THIS CHECK, BUT KEEP THIS IN CASE MAPPING
			// OF PHONENUMBER CHANGED TO CELLPHONE
			// PLACE IT AFTER DUPLICATE EMAIL CHECK
			// *****************************************************************
			/*
				else if strings.Contains(errText, dupTransactantWithCellPhone) {
					lineNo, _, ok := parseLineAndErrorFromRCSV(err, dbType)
					if !ok {
						// INTERNAL ERROR - RETURN FALSE
						return false
					}
					// get tracedatamap
					traceDataMap := getTraceDataMap(traceDataMapName)
					// now get the original row index of imported onesite csv and Unit value
					onesiteIndex, unit := getIndexAndUnit(traceDataMap, lineNo)
					// load csvRow from dataMap to get email
					csvRow := *csvRowDataMap[onesiteIndex]
					pCellNo := csvRow.PhoneNumber
					// get tcid from cellphonenumber
					t, tErr := rlib.GetTransactantByPhoneOrEmail(ctx, business.BID, pCellNo)
					if tErr != nil {
						// unable to get TCID
						reason := "E:<" + core.DBTypeMapStrings[core.DBPeople] + ">:Unable to get people information" + err.Error()
						csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
					} else if t.TCID == 0{
						// unable to get TCID
						reason := "E:<" + core.DBTypeMapStrings[core.DBPeople] + ">:Unable to get people information"
						csvErrors[onesiteIndex] = append(csvErrors[onesiteIndex], reason)
					} else {
						// if duplicate people found
						rlib.Ulog("DUPLICATE RECORD ERROR <%s>: %s", fname, err.Error())
						// map it in tcid map
						traceTCIDMap[onesiteIndex] = tcidPrefix + strconv.FormatInt(t.TCID, 10)
					}
				}
			*/
			// *****************************************************

		}
		// return with success
		return true
	}

	// =========================================
	// LOAD CUSTOM ATTRIBUTE & RENTABLE TYPE CSV
	// =========================================
	var h = []csvLoadHandler{
		{Fname: customAttributeCSVFile.Name(), Handler: rcsv.LoadCustomAttributesCSV, TraceDataMap: "traceCustomAttributeCSVMap", DBType: core.DBCustomAttr},
		{Fname: rentableTypeCSVFile.Name(), Handler: rcsv.LoadRentableTypesCSV, TraceDataMap: "traceRentableTypeCSVMap", DBType: core.DBRentableType},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrDoLoad(ctx, h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return traceUnitMap, csvErrors, internalErrFlag
			}
		}
	}

	// =====================================
	// INSERT CUSTOM ATTRIBUTE REF MANUALLY
	// AFTER CUSTOM ATTRIB AND RENTABLE TYPE
	// LOADED SUCCESSFULLY
	// =====================================

	// always sort keys
	var customAttributesRefDataKeys []string
	for k := range customAttributesRefData {
		customAttributesRefDataKeys = append(customAttributesRefDataKeys, k)
	}
	sort.Strings(customAttributesRefDataKeys)

	for _, key := range customAttributesRefDataKeys {
		errPrefix := "E:<" + core.DBTypeMapStrings[core.DBCustomAttrRef] + ">:"
		// find rentableType
		refData := customAttributesRefData[key]
		rt, err := rlib.GetRentableTypeByStyle(ctx, refData.Style, refData.BID)
		if err != nil {
			rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", err.Error())
			csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
			continue
		}

		// for all custom attribute defined in custom_attrib.go
		// find custom attribute ID
		for _, customAttributeConfig := range customAttributeMap {
			t, _ := strconv.ParseInt(customAttributeConfig["ValueType"], 10, 64)
			n := customAttributeConfig["Name"]
			v := strconv.Itoa(int(refData.SqFt))
			u := customAttributeConfig["Units"]
			ca, err := rlib.GetCustomAttributeByVals(ctx, t, n, v, u)
			if err != nil {
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", "CUSTOM ATTRIBUTE NOT FOUND IN DB")
				csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
				continue
			}
			// if resource not found then continue
			if ca.CID == 0 {
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", "CUSTOM ATTRIBUTE NOT FOUND IN DB")
				csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
				continue
			}

			// count possible values
			CustomAttrRefRecordCount++

			// insert custom attribute ref in system
			var a rlib.CustomAttributeRef
			a.ElementType = rlib.ELEMRENTABLETYPE
			a.BID = business.BID
			a.ID = rt.RTID
			a.CID = ca.CID

			// check that record already exists, if yes then just continue
			// without accounting it as an error
			ref, err := rlib.GetCustomAttributeRef(ctx, a.ElementType, a.ID, a.CID)
			if err != nil {
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", err.Error())
				continue
			} else {
				if ref.ElementType == a.ElementType && ref.CID == a.CID && ref.ID == a.ID {
					unit, _ := traceUnitMap[refData.RowIndex]
					errText := fmt.Sprintf(
						"This reference already exists. No changes were made. at row \"%d\" with unit \"%s\"",
						refData.RowIndex, unit)
					rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", errText)
					continue
				}
			}

			_, err = rlib.InsertCustomAttributeRef(ctx, &a)
			if err != nil {
				rlib.Ulog("ERROR <CUSTOMREF INSERTION>: %s", err.Error())
				csvErrors[refData.RowIndex] = append(csvErrors[refData.RowIndex], errPrefix+"Unable to insert custom attribute")
				continue
			}
		}
	}

	// ================
	// LOAD PEOPLE CSV
	// ================
	h = []csvLoadHandler{
		{Fname: peopleCSVFile.Name(), Handler: rcsv.LoadPeopleCSV, TraceDataMap: "tracePeopleCSVMap", DBType: core.DBPeople},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrPeopleDoLoad(ctx, h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return traceUnitMap, csvErrors, internalErrFlag
			}
		}
	}

	// ========================================================
	// GET TCID FOR EACH ROW FROM PEOPLE CSV AND UPDATE TCID MAP
	// ========================================================

	for onesiteIndex := range traceTCIDMap {
		tcid, _ := rlib.GetTCIDByNote(ctx, getPeopleNoteString(onesiteIndex, currentTimeFormat))
		// for duplicant case, it won't be found so need check here
		if tcid != 0 {
			traceTCIDMap[onesiteIndex] = tcidPrefix + strconv.Itoa(int(tcid))
		}
	}

	// ==============================================================
	// AFTER POSSIBLE TCID FOUND, WRITE RENTABLE & RENTAL AGREEMENT CSV
	// ==============================================================

	// get created people csv and writer pointer
	rentableCSVFile, rentableCSVWriter, ok :=
		CreateRentableCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.RentableCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTABLE CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// get created rental agreement csv and writer pointer
	rentalAgreementCSVFile, rentalAgreementCSVWriter, ok :=
		CreateRentalAgreementCSV(
			TempCSVStore, currentTimeFormat,
			&oneSiteFieldMap.RentalAgreementCSV,
		)
	if !ok {
		rlib.Ulog("INTERNAL ERROR <RENTAL AGREEMENT CSV>: %s\n", err.Error())
		return traceUnitMap, csvErrors, internalErrFlag
	}

	// always sort keys to iterate over csv rows in proper manner (from top to bottom)
	var csvRowDataMapKeys []int
	for k := range csvRowDataMap {
		csvRowDataMapKeys = append(csvRowDataMapKeys, k)
	}
	sort.Ints(csvRowDataMapKeys)

	// iteration over csv row data structure and write data to csv
	for _, rowIndex := range csvRowDataMapKeys {

		// load csvRow from dataMap
		csvRow := *csvRowDataMap[rowIndex]

		//---------------------------------------------------------------------
		// for rentable status exists in csvRow, get set of csv types which
		// can be allowed to perform write data for csv
		// need to call validation function as in get values
		//---------------------------------------------------------------------
		// rlib.Console("rowIndex = %v\n", rowIndex)
		// rlib.Console("\t UnitLeaseStatus = %v\n", csvRow.UnitLeaseStatus)
		_, rrUseType, _ := IsValidRentableUseType(csvRow.UnitLeaseStatus)
		csvTypesSet := canWriteCSVStatusMap[rrUseType]
		var canWriteData bool

		// check first that for this row's status rentable data can be written
		canWriteData = core.IntegerInSlice(core.RENTABLECSV, csvTypesSet)
		// rlib.Console("\t csvTypesSet = %v\n", csvTypesSet)
		// rlib.Console("\t canWriteData = %v\n", canWriteData)
		if canWriteData {
			// Write data to file of rentable
			WriteRentableData(
				&RentableCSVRecordCount,
				rowIndex,
				traceRentableCSVMap,
				rentableCSVWriter,
				&csvRow,
				currentTime,
				currentTimeFormat,
				userRRValues,
				&oneSiteFieldMap.RentableCSV,
				traceTCIDMap,
				csvErrors,
				rrUseType,
			)
		}

		// check first that for this row's status rental agreement data can be written
		canWriteData = core.IntegerInSlice(core.RENTALAGREEMENTCSV, csvTypesSet)
		if canWriteData {
			// Write data to file of rentalAgreement
			WriteRentalAgreementData(
				&RentalAgreementCSVRecordCount,
				rowIndex,
				traceRentalAgreementCSVMap,
				rentalAgreementCSVWriter,
				&csvRow,
				currentTime,
				currentTimeFormat,
				userRRValues,
				&oneSiteFieldMap.RentalAgreementCSV,
				traceTCIDMap,
				csvErrors,
			)
		}
	}

	// closing files
	rentableCSVFile.Close()
	rentalAgreementCSVFile.Close()

	// =====================================
	// LOAD RENTABLE & RENTAL AGREEMENT CSV
	// =====================================
	h = []csvLoadHandler{
		{Fname: rentableCSVFile.Name(), Handler: rcsv.LoadRentablesCSV, TraceDataMap: "traceRentableCSVMap", DBType: core.DBRentable},
		{Fname: rentalAgreementCSVFile.Name(), Handler: rcsv.LoadRentalAgreementCSV, TraceDataMap: "traceRentalAgreementCSVMap", DBType: core.DBRentalAgreement},
	}

	for i := 0; i < len(h); i++ {
		if len(h[i].Fname) > 0 {
			if !rrDoLoad(ctx, h[i].Fname, h[i].Handler, h[i].TraceDataMap, h[i].DBType) {
				// INTERNAL ERROR
				return traceUnitMap, csvErrors, internalErrFlag
			}
		}
	}

	// ============================
	// CLEAR THE TEMPORARY CSV FILES
	// ============================
	// testmode is not enabled then only remove temp files
	if testMode != 1 {
		clearSplittedTempCSVFiles(currentTimeFormat)
	}

	// ===============================
	// EVALUATE SUMMARY REPORT COUNT
	// ===============================

	// count possible values
	summaryReport[core.DBRentable]["possible"] = RentableCSVRecordCount
	summaryReport[core.DBRentalAgreement]["possible"] = RentalAgreementCSVRecordCount
	summaryReport[core.DBRentableType]["possible"] = RentableTypeCSVRecordCount
	summaryReport[core.DBCustomAttr]["possible"] = CustomAttributeCSVRecordCount
	summaryReport[core.DBCustomAttrRef]["possible"] = CustomAttrRefRecordCount
	summaryReport[core.DBPeople]["possible"] = PeopleCSVRecordCount

	// =======
	// RETURN
	// =======
	// no internal error so make it false
	internalErrFlag = false
	return traceUnitMap, csvErrors, internalErrFlag
}

// rollBackImportOperation func used to clear out the things
// that created by program temporarily while loading onesite data
//  and if any error occurs
func rollBackImportOperation(timestamp string) {
	clearSplittedTempCSVFiles(timestamp)
}

// clearSplittedTempCSVFiles func used only to clear
// temporarily csv files created by program
func clearSplittedTempCSVFiles(timestamp string) {
	for _, filePrefix := range prefixCSVFile {
		fileName := filePrefix + timestamp + ".csv"
		filePath := path.Join(TempCSVStore, fileName)
		os.Remove(filePath)
	}
}

// CSVHandler is main function to handle user uploaded
// csv and extract information
//
// INPUTS
//   ctx      database context
//   csvPath  where to put csv files we create
//   testMode
//   userRRValues
//   business
//   debugMode
//
// RETURNS report, internal error flag, done (csv loaded or not)
//---------------------------------------------------------------------
func CSVHandler(ctx context.Context, csvPath string, testMode int, userRRValues map[string]string, business *rlib.Business, debugMode int) (string, bool, bool) {
	// rlib.Console("*** Entered CSVHandler ***\n")
	csvLoaded := true         // csv loaded successfully flag
	csvReport := ""           // report text
	currentTime := time.Now() // get current timestamp used for creating csv files unique way

	//-------------------------------------------------------------
	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// it is helpful while creating unique files
	//-------------------------------------------------------------
	currentTimeFormat := currentTime.Format(time.RFC3339Nano)

	//-------------------------------------------------------------
	// summaryReportCount contains each type csv as a key
	// with count of total imported, possible, issues in csv data
	//-------------------------------------------------------------
	summaryReportCount := map[int]map[string]int{
		core.DBCustomAttr:      {"imported": 0, "possible": 0, "issues": 0},
		core.DBRentableType:    {"imported": 0, "possible": 0, "issues": 0},
		core.DBCustomAttrRef:   {"imported": 0, "possible": 0, "issues": 0},
		core.DBPeople:          {"imported": 0, "possible": 0, "issues": 0},
		core.DBRentable:        {"imported": 0, "possible": 0, "issues": 0},
		core.DBRentalAgreement: {"imported": 0, "possible": 0, "issues": 0},
	}

	//-------------------------------------------------------------
	//  Call onesite loader
	//-------------------------------------------------------------
	unitMap, csvErrs, internalErr := loadOneSiteCSV(ctx, csvPath, testMode, userRRValues, business, currentTime, currentTimeFormat, summaryReportCount)
	if internalErr { // if internal error then just return from here, nothing to do
		return csvReport, internalErr, csvLoaded
	}

	// check if there any errors from onesite loader
	if len(csvErrs) > 0 {
		csvReport, csvLoaded = errorReporting(ctx, business, csvErrs, unitMap, summaryReportCount, csvPath, debugMode, currentTime)
		if testMode != 1 { // if not testmode then only do rollback
			rollBackImportOperation(currentTimeFormat)
		}

		return csvReport, internalErr, csvLoaded
	}

	//-------------------------------------------------------------
	// Generate Report
	//-------------------------------------------------------------
	csvReport = successReport(ctx, business, summaryReportCount, csvPath, debugMode, currentTime)

	// ===== 5. Return =====
	return csvReport, internalErr, csvLoaded
}
