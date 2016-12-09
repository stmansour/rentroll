package rlib

import "reflect"

// BuildFieldMap build a map of the supplied struct pointer so that
// you can index the map by a field name and the associated value is
// the field's index within p
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
	m := BuildFieldMap(pb) // we'll map pb's fields, then process pa one field at a time

	ar := reflect.ValueOf(pa).Elem()
	for i := 0; i < ar.NumField(); i++ {
		fa := ar.Field(i)
		afldname := ar.Type().Field(i).Name
		if !fa.IsValid() { // skip fields in an invalid state, nil pointers, zero-valued lists, ...
			continue
		}
		bdx, ok := m[afldname]
		if !ok { // if pb doesn't have this field, move on
			continue
		}
		br := reflect.ValueOf(pb).Elem()
		fb := br.Field(bdx)
		if !fb.CanSet() { // if it cannot be set then just move on
			continue
		}
		if fa.Type() == fb.Type() {
			fb.Set(reflect.ValueOf(fa.Interface()))
		} else {
			val := reflect.ValueOf(fa.Interface())
			fb.Set(val.Convert(fb.Type()))
		}
	}
	return nil
}
