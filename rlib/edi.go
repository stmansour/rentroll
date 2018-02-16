package rlib

// modifyInterfaceEDI is similar to printArg but starts with a reflect value, not an interface{} value.
import (
	"reflect"
	"time"
)

var exceptionalStructStopDateMap = map[string][]string{
	"RentRollStaticInfo": []string{"RentStop", "AgreementStop", "PossessionStop"},
	"RentalAgreement":    []string{"RentStop", "AgreementStop", "PossessionStop"},
	"AssessmentGrid":     []string{"Stop"},
	"PrRentableOther":    []string{"RentalAgreementStop"},
}

// DateMode are etc. all constants used for end date inclusion condition
const (
	// DateMode             = true
	layout               = "01/02/2006"
	defaultStopDateField = "DtStop"
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

func modifyInterfaceEDI(elemStopDate reflect.Value) {

	// if target field is being able to set then
	if elemStopDate.IsValid() && elemStopDate.CanSet() {

		var dtStop time.Time
		switch elemStopDate.Type().String() { // type().string() gives the full path
		case "rlib.JSONDate":
			dtStop = (time.Time)(elemStopDate.Interface().(JSONDate))
		case "rlib.JSONDateTime":
			dtStop = (time.Time)(elemStopDate.Interface().(JSONDateTime))
		case "rlib.NullDate":
			// TODO(Sudip): what if it contains invalid date, i.e., valid flag = false?
			nd := elemStopDate.Interface().(NullDate)
			dtStop = nd.Time
		/*case "rlib.NullDateTime":
		// TODO(Sudip): what if it contains invalid date, i.e., valid flag = false?
		nd := elemStopDate.Interface().(NullDateTime)
		dtStop = nd.Time*/
		case "time.Time":
			dtStop = elemStopDate.Interface().(time.Time)
		default: // TODO(Sudip): better handle situation here

			return
		}

		// move one day back
		dtStop = dtStop.AddDate(0, 0, -1)

		// TODO: make sure stopDate should not be prior to start date
		//       Here it needs the proper map of start and stop dates

		// convert back from time.Time to original defined type
		var v reflect.Value
		switch elemStopDate.Type().String() {
		case "rlib.JSONDate":
			v = reflect.ValueOf((JSONDate)(dtStop))
		case "rlib.JSONDateTime":
			v = reflect.ValueOf((JSONDateTime)(dtStop))
		case "rlib.NullDate":
			nd := elemStopDate.Interface().(NullDate)
			nd.Time = dtStop
			v = reflect.ValueOf(nd)
		/*case "rlib.NullDateTime":
		nd := elemStopDate.Interface().(NullDateTime)
		nd.Time = dtStop
		v = reflect.ValueOf(nd)*/
		default:
			v = reflect.ValueOf(dtStop)
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

		// get exceptional dates from the map if not found then take
		// default one
		exceptionalDates, ok := exceptionalStructStopDateMap[f.Type().Name()]
		if !ok { // process on default field -> "DtStop"
			exceptionalDates = append(exceptionalDates, defaultStopDateField)
		}

		// operate on list of end dates
		for _, edFieldName := range exceptionalDates {

			// get field from ed string
			elemStopDate := f.FieldByName(edFieldName)

			// modify field stop date value
			modifyInterfaceEDI(elemStopDate)
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
	// TODO(Sudip): it should coming from cache, not by hitting db everytime
	// if b, ok := RRdb.BizTypes[BID]; ok {
	// 	ediEnabled = b.FLAGS&1 > 0 // see if bit 1 is set or not
	// }
	var xbiz XBusiness
	err := GetXBiz(BID, &xbiz)
	if err == nil {
		ediEnabled = xbiz.P.FLAGS&1 > 0
	}
	return ediEnabled
}

// HandleInterfaceEDI handles the end date inclusion situation
// based on application DateMode setting for the p interface
// which is having dateranges fields
func HandleInterfaceEDI(p interface{}, BID int64) {

	// if end date inclusion enabled
	if EDIEnabledForBID(BID) {

		// make sure interface should be kind of ptr, otherwise field value will not be changed
		if reflect.TypeOf(p).Kind() != reflect.Ptr {
			return
		}

		// send the reflect.Value of given interface p
		lookForInterfaceStopDate(reflect.ValueOf(p), 0) // depth=0
	}
}
