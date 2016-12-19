package rlib

import (
	"fmt"
	"reflect"
	"strings"
)

var xjson = string("XJSON")

// XJSONprocess attempts to map a to b. If no converter can befound
// a message will be printed, then it will panic!
func XJSONprocess(a, b *reflect.Value) error {
	at := (*a).Type().String()
	bt := (*b).Type().String()
	for i := 0; i < len(assignmap); i++ {
		if strings.Index(at, assignmap[i].a) >= 0 && strings.Index(bt, assignmap[i].b) >= 0 {
			assignmap[i].mapper(a, b)
			return nil
		}
	}
	return fmt.Errorf("XJSONmap - no conversion between: %s and %s\n", at, bt)

	// fmt.Printf(s)
	// panic(s)
}

// Str2Int64Map is a generic type for mapping strings and int64s
type Str2Int64Map map[string]int64

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
// names for whatever pa is pointing to matches the field name in pb
// There is a basic assumption that the data will either copy directly
// or convert cleanly from one struct to another.
//--------------------------------------------------------------------
func MigrateStructVals(pa interface{}, pb interface{}) error {
	m := BuildFieldMap(pb)
	ar := reflect.ValueOf(pa).Elem()
	for i := 0; i < ar.NumField(); i++ {
		fa := ar.Field(i)
		afldname := ar.Type().Field(i).Name
		if !fa.IsValid() {
			continue
		}
		bdx, ok := m[afldname]
		if !ok {
			continue
		}
		br := reflect.ValueOf(pb).Elem()
		fb := br.Field(bdx)
		if !fb.CanSet() {
			continue
		}
		if fa.Type() == fb.Type() {
			fb.Set(reflect.ValueOf(fa.Interface()))
		} else {
			// if strings.Index(ta, xjson) >= 0 || strings.Index(tb, xjson) >= 0 {
			err := XJSONprocess(&fa, &fb)
			if err != nil {
				val := reflect.ValueOf(fa.Interface())
				fb.Set(val.Convert(fb.Type()))
			}
		}
	}
	return nil
}
