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
}

const (
	DateMode             = true
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

func modifyInterfaceEDI(value reflect.Value) {

	// Reference: https://golang.org/src/fmt/print.go#L687
	// modifyInterfaceEDI = printValue

	switch f := value; value.Kind() {
	case reflect.Struct:
		/*
		   As we're in struct type, look for the stopDate(s) inside this struct
		*/

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

		// Now, target on any field, composed structs via different kind of
		// elements
		for i := 0; i < f.NumField(); i++ {
			switch f.Field(i).Kind() {
			case reflect.Struct,
				reflect.Interface,
				reflect.Map,
				reflect.Array, reflect.Slice,
				reflect.Ptr:
				modifyInterfaceEDI(getField(f, i))
			}
		}
	case reflect.Interface:
		value := f.Elem()
		if value.IsValid() {
			modifyInterfaceEDI(value)
		}
	case reflect.Map:
		// TODO(Sudip): what if key itself contains a struct
		keys := f.MapKeys()
		for _, key := range keys {
			modifyInterfaceEDI(f.MapIndex(key))
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < f.Len(); i++ {
			modifyInterfaceEDI(f.Index(i))
		}
	case reflect.Ptr:
		// pointer to array or slice or struct?  ok at top level
		// but not embedded (avoid loops)
		switch a := f.Elem(); a.Kind() {
		case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
			modifyInterfaceEDI(a)
			return
		}
		/*fallthrough
		case reflect.Chan, reflect.Func, reflect.UnsafePointer:
			p.fmtPointer(f, verb)
			fmt.Println("Chan, Func, UnsafePointer")*/
	}
}

// HandleInterfaceEDI handles the end date inclusion situation
// based on application DateMode setting for the p interface
// which is having dateranges fields
func HandleInterfaceEDI(p interface{}) {
	if DateMode {

		if reflect.TypeOf(p).Kind() != reflect.Ptr {
			return
		}

		// send the reflect.Value of given interface p
		modifyInterfaceEDI(reflect.ValueOf(p))
	}
}
