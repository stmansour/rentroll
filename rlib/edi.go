package rlib

// modifyInterfaceEDI is similar to printArg but starts with a reflect value, not an interface{} value.
import (
	"reflect"
	"time"
)

// dateRangeFieldsMap contains map of stop to start date.
// The logic needs to be ensure that end date modification will
// not prior to start date by looking into this mapping for a end date.
var dateRangeFieldsMap = map[string]string{
	"DtStop":              "DtStart",              // the default one which exists in most struct
	"AgreementStop":       "AgreementStart",       // RentalAgreement, Rentroll view
	"PossessionStop":      "PossessionStart",      // RentalAgreement, Rentroll view
	"RentStop":            "RentStart",            // RentalAgreement, Rentroll view
	"RentalAgreementStop": "RentalAgreementStart", // Rentables{grid, form}
	"Stop":                "Start",                // AsssessmentGrid
	"RARDtStop":           "RARDtStart",           // RentalAgreementRentables
	"DtMRStop":            "DtMRStart",            // MRHistory in rlib
}

// known struct list to allow date conversion.
var ediKnownStructMap = map[string]bool{
	"rlib.RAAcctBal":               true,
	"rlib.RentRollStaticInfo":      true,
	"rlib.VacancyMarker":           true,
	"rlib.AssessmentType":          true,
	"rlib.RentalAgreementGrid":     true,
	"rlib.RatePlanRef":             true,
	"rlib.RentalAgreement":         true,
	"rlib.RentalAgreementRentable": true,
	"rlib.RentalAgreementTax":      true,
	"rlib.RentalAgreementPayor":    true,
	"rlib.RentableUser":            true,
	"rlib.Pet":                     true,
	"rlib.Vehicle":                 true,
	"rlib.Assessment":              true,
	"rlib.AR":                      true,
	"rlib.RentableMarketRate":      true,
	"rlib.RentableTypeTax":         true,
	"rlib.MRHistory":               true,
	"rlib.RentableTypeRef":         true,
	"rlib.RentCycleRef":            true,
	"rlib.RentableSpecialtyRef":    true,
	"rlib.RentableUseStatus":       true,
	"rlib.XRentable":               true,
	"rlib.JournalMarker":           true,
	"rlib.RAFlowJSONData":          true,
	"ws.ARSendForm":                true,
	"ws.PrARGrid":                  true,
	"ws.AssessmentSendForm":        true,
	"ws.AssessmentGrid":            true,
	"ws.RentalAgr":                 true,
	"ws.RentalAgrForm":             true,
	"ws.RAPayor":                   true,
	"ws.RAPeople":                  true,
	"ws.RAR":                       true,
	"ws.PrRentableOther":           true,
	"ws.RentableUseStatusGridRec":  true,
	"ws.RentableTypeRefGridRec":    true,
	"ws.RentableMarketRateGridRec": true,
	"ws.StatementInfoGridRecord":   true,
	"ws.PayorHistory":              true,
	"ws.StmtGrid":                  true,
}

// DateMode are etc. all constants used for end date inclusion condition
const (
	layout = "01/02/2006"
)

// getField gets the i'th field of the struct value.
// If the field is itself is an interface, return a value for
// the thing inside the interface, not the interface itself.
func getField(v reflect.Value, i int) reflect.Value {
	val := v.Field(i)
	if val.Kind() == reflect.Interface && !val.IsNil() {
		val = val.Elem()
	}
	return val
}

// modifyInterfaceEDI will change elemStopDate field
// if its valid and can set and falls in the list of
// defined switch cases of types, related to time.Time type.
func modifyInterfaceEDI(elemStopDate, elemStartDate reflect.Value) {

	// this applies to only date ranges so make sure that both start and end date elements are valid
	// if target field is being able to set then
	if elemStartDate.IsValid() && elemStopDate.IsValid() && elemStopDate.CanSet() {

		var dtStop, dtStart time.Time
		switch elemStopDate.Type().String() { // type().string() gives the full path
		case "rlib.JSONDate":
			dtStop = (time.Time)(elemStopDate.Interface().(JSONDate))
			dtStart = (time.Time)(elemStartDate.Interface().(JSONDate))
		case "rlib.JSONDateTime":
			dtStop = (time.Time)(elemStopDate.Interface().(JSONDateTime))
			dtStart = (time.Time)(elemStartDate.Interface().(JSONDateTime))
		case "rlib.NullDate":
			// TODO(Sudip): what if it contains invalid date, i.e., valid flag = false?
			nd := elemStopDate.Interface().(NullDate)
			dtStop = nd.Time
			ns := elemStartDate.Interface().(NullDate)
			dtStart = ns.Time
		/*case "rlib.NullDateTime":
		// TODO(Sudip): what if it contains invalid date, i.e., valid flag = false?
		nd := elemStopDate.Interface().(NullDateTime)
		dtStop = nd.Time
		ns := elemStartDate.Interface().(NullDateTime)
		dtStart = ns.Time*/
		case "time.Time":
			dtStop = elemStopDate.Interface().(time.Time)
			dtStart = elemStartDate.Interface().(time.Time)
		default: // TODO(Sudip): better handle situation here
			return
		}

		// move one day back
		modStopDate := dtStop.AddDate(0, 0, -1)

		// TODO(Sudip): might want to consider some exceptional cases, confirm with Steve

		// if `modStopDate` is not prior or equals to start date
		// then only set the modified end dates value in its place
		// otherwise just return, don't proceed further
		if dtStart.After(modStopDate) || dtStart.Equal(modStopDate) {
			return
		}

		// convert back from time.Time to original defined type
		var v reflect.Value
		switch elemStopDate.Type().String() {
		case "rlib.JSONDate":
			v = reflect.ValueOf((JSONDate)(modStopDate))
		case "rlib.JSONDateTime":
			v = reflect.ValueOf((JSONDateTime)(modStopDate))
		case "rlib.NullDate":
			nd := elemStopDate.Interface().(NullDate)
			nd.Time = modStopDate
			v = reflect.ValueOf(nd)
		/*case "rlib.NullDateTime":
		nd := elemStopDate.Interface().(NullDateTime)
		nd.Time = modStopDate
		v = reflect.ValueOf(nd)*/
		default:
			v = reflect.ValueOf(modStopDate)
		}

		// set packed modified end date reflect Value in original struct
		elemStopDate.Set(v)
	}
}

func lookForInterfaceStopDate(value reflect.Value, depth int) {

	// Reference: https://golang.org/src/fmt/print.go#L687 |> modifyInterfaceEDI = printValue <|

	switch f := value; value.Kind() {
	case reflect.Struct:
		// This applies to all structs except type of time.Time itself.
		// otherwise it will apply modification two times.

		/*if depth == 0 { // what if value itself is type of time.Time
			modifyInterfaceEDI(value)
		}*/

		// is it known struct and enabled for looking and end date modification
		// Console("Struct type with package: %s\n", f.Type().String())
		enabledForLook, ok := ediKnownStructMap[f.Type().String()]
		if ok && enabledForLook {
			// operate on list of end dates
			// dateRangeFieldMap <=> {endDate: startDate}
			for edFieldName, sdFieldName := range dateRangeFieldsMap {

				// get field from ed string
				elemStopDate := f.FieldByName(edFieldName)

				// get field from sd string
				elemStartDate := f.FieldByName(sdFieldName)

				// modify field stop date value
				modifyInterfaceEDI(elemStopDate, elemStartDate)
			}
		}

		// Now, target on any field, composed internal structs via different kind of
		// elements
		for i := 0; i < f.NumField(); i++ {
			switch f.Field(i).Kind() {
			case reflect.Struct,
				reflect.Interface,
				reflect.Map,
				reflect.Array, reflect.Slice,
				reflect.Ptr:
				lookForInterfaceStopDate(getField(f, i), depth+1)
			}
		}
	case reflect.Interface:
		value := f.Elem()
		if value.IsValid() {
			lookForInterfaceStopDate(value, depth+1)
		}
	case reflect.Map:
		// TODO(Sudip): what if key itself contains a struct
		keys := f.MapKeys()
		for _, key := range keys {
			lookForInterfaceStopDate(f.MapIndex(key), depth+1)
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < f.Len(); i++ {
			lookForInterfaceStopDate(f.Index(i), depth+1)
		}
	case reflect.Ptr:
		// pointer to array or slice or struct?  ok at top level
		// but not embedded (avoid loops)
		switch a := f.Elem(); a.Kind() {
		case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
			lookForInterfaceStopDate(a, depth+1)
			return
		}
		/*fallthrough
		case reflect.Chan, reflect.Func, reflect.UnsafePointer:
			p.fmtPointer(f, verb)
			fmt.Println("Chan, Func, UnsafePointer")*/
	}
}

// EDIEnabledForBID checks whether EDI for business with BID is enabled or not
func EDIEnabledForBID(BID int64) bool {
	var ediEnabled bool

	// look, if FLAGS is set or not
	if bizCache, ok := RRdb.BizCache[BID]; ok {
		ediEnabled = bizCache.FLAGS&1 > 0
	}

	return ediEnabled
}

// HandleInterfaceEDI handles the end date inclusion situation
// based on application DateMode setting for the p interface
// which is having dateranges fields
// right now it supports only "struct" kind of element
// -----------------------------------------------------------------------------
func HandleInterfaceEDI(p interface{}, BID int64) {

	// if end date inclusion enabled
	if EDIEnabledForBID(BID) {

		/*Console("Elem kind: %s\n", reflect.ValueOf(p).Elem().Kind())
		Console("Struct with Package: %s\n", reflect.ValueOf(p).Elem().Type().String())*/

		// 1. make sure interface should be kind of ptr, otherwise field value will not be changed
		// 2. underlying element should be kind of struct
		if reflect.TypeOf(p).Kind() != reflect.Ptr && reflect.ValueOf(p).Elem().Kind() != reflect.Struct {
			return
		}

		// send the reflect.Value of given interface p
		lookForInterfaceStopDate(reflect.ValueOf(p), 0) // depth=0
	}
}

// EDIHandleIncomingDateRange will modify front end dates coming from web
// service/command line app. If edi is enabled for BID then it will modify
// stopDate incrementing by one day.
// -----------------------------------------------------------------------------
func EDIHandleIncomingDateRange(BID int64, dtStart, dtStop *time.Time) {
	if EDIEnabledForBID(BID) {
		*dtStop = dtStop.AddDate(0, 0, 1)
	}

	//----------------------------------------------------------------------
	// it is ok for start/stop dates to be equal -- that's what happens on
	// non-recurring assessments.  But it is not ok for the stop time to be
	// before the start time.
	//----------------------------------------------------------------------
	if dtStop.Before(*dtStart) {
		*dtStop = *dtStart
	}
}

// EDIHandleIncomingJSONDateRange will modify front end dates coming from web
// service/command line app. If edi is enabled for BID then it will modify
// stopDate incrementing by one day.
// -----------------------------------------------------------------------------
func EDIHandleIncomingJSONDateRange(BID int64, dtStart, dtStop *JSONDate) {
	d1 := time.Time(*dtStart)
	d2 := time.Time(*dtStop)
	EDIHandleIncomingDateRange(BID, &d1, &d2)
	*dtStop = JSONDate(d2)
}

// EDIHandleOutgoingDateRange will modify give stopDate if EDI is enabled for
// the given business
// -----------------------------------------------------------------------------
func EDIHandleOutgoingDateRange(BID int64, d1, stopDate *time.Time) {
	if EDIEnabledForBID(BID) {
		*stopDate = stopDate.AddDate(0, 0, -1)
		if stopDate.Before(*d1) {
			*stopDate = *d1
		}
	}
}

// EDIHandleNDOutgoingDateRange is like EDIHandleOutgoingDateRange but handles
// NullDates
// -----------------------------------------------------------------------------
func EDIHandleNDOutgoingDateRange(BID int64, d1, stopDate *NullDate) {
	if !d1.Valid || !stopDate.Valid {
		return
	}
	EDIHandleOutgoingDateRange(BID, &d1.Time, &stopDate.Time)
}

// EDIHandleOutgoingJSONDateTimeRange will modify give stopDate if EDI is enabled for
// the given business
// -----------------------------------------------------------------------------
func EDIHandleOutgoingJSONDateTimeRange(BID int64, DtStart, DtStop *JSONDateTime) {
	var d1 = time.Time(*DtStart)
	var d2 = time.Time(*DtStop)

	if EDIEnabledForBID(BID) {
		d2 = d2.AddDate(0, 0, -1)
		*DtStop = JSONDateTime(d2)
		if d2.Before(d1) {
			*DtStop = *DtStart
		}
	}
}

// EDIHandleOutgoingJSONDateRange will modify give stopDate if EDI is enabled for
// the given business
// -----------------------------------------------------------------------------
func EDIHandleOutgoingJSONDateRange(BID int64, DtStart, DtStop *JSONDate) {
	var d1 = time.Time(*DtStart)
	var d2 = time.Time(*DtStop)

	if EDIEnabledForBID(BID) {
		d2 = d2.AddDate(0, 0, -1)
		*DtStop = JSONDate(d2)
		if d2.Before(d1) {
			*DtStop = *DtStart
		}
	}
}
