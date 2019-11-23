package rlib

import (
	"fmt"
	"reflect"
)

// ReverseMap takes a string-to-int64 map and does a search for the int64 val
// and returns the string. The return value is the string along with an error.
// The error is nil if the int64 was found, otherwise it indicates the problem.
func (t *Str2Int64Map) ReverseMap(m int64) (string, error) {
	for k, v := range *t {
		if m == v {
			return k, nil
		}
	}
	return "", fmt.Errorf("%d not found", m)
}

// Str2IntMap is a generic type for mapping strings and ints
type Str2IntMap map[string]int

// ReverseMap takes a Str2IntMap and does a search for the int val
// and returns the string and an error if not found.
func (t *Str2IntMap) ReverseMap(m int) (string, error) {
	for k, v := range *t {
		if m == v {
			return k, nil
		}
	}
	return "", fmt.Errorf("%d not found", m)
}

// BuildFieldMap creates a map so that we can find
// a field's index using its name as the map index
//--------------------------------------------------------------------
func BuildFieldMap(p interface{}) map[string]int {
	var fmap = map[string]int{}
	v := reflect.ValueOf(p).Elem()
	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name
		fmap[n] = j
	}
	return fmap
}

// MigrateStructVals copies values from pa to pb where the field
// names for the struct pa points to matches the field names in
// the struct pb points to.
// There is a basic assumption that the data will either copy directly
// or convert cleanly from one struct to another.  Where it does not
// it will call XJSONprocess to see if there is a known conversion.
//--------------------------------------------------------------------
func MigrateStructVals(pa interface{}, pb interface{}) error {
	m := BuildFieldMap(pb)
	ar := reflect.ValueOf(pa).Elem()
	for i := 0; i < ar.NumField(); i++ {
		fa := ar.Field(i)
		afldname := ar.Type().Field(i).Name
		// Console("MSV 1  field = %s\n", afldname)
		if !fa.IsValid() {
			continue
		}
		// Console("MSV 2\n")
		bdx, ok := m[afldname]
		if !ok {
			continue
		}
		// Console("MSV 3\n")
		br := reflect.ValueOf(pb).Elem()
		fb := br.Field(bdx)
		if !fb.CanSet() { // BEWARE: if a field name begins with a lowercase letter it cannot be set
			continue
		}
		// Console("MSV 4: Can set b field\n")
		if fa.Type() == fb.Type() {
			// Console("MSV 5\n")
			fb.Set(reflect.ValueOf(fa.Interface()))
			// Console("MSV 6\n")
		} else {
			// Console("MSV 7\n")
			err := XJSONprocess(&fa, &fb)
			if err != nil {
				// Console("MSV 8\n")
				val := reflect.ValueOf(fa.Interface()) // instantiate new pa value
				fb.Set(val.Convert(fb.Type()))         // set pb to the value of the new type value
			}
			// Console("MSV (8)\n")
		}
	}
	return nil
}
