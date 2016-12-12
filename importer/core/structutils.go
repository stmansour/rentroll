package core

import (
	"fmt"
	"reflect"
)

// func BuildFieldMap(p interface{}) map[string]int {
// 	var fmap = map[string]int{}
// 	v := reflect.ValueOf(p).Elem()
// 	for j := 0; j < v.NumField(); j++ {
// 		n := v.Type().Field(j).Name
// 		fmap[n] = j
// 	}
// 	return fmap
// }

// Attributes accepts any type of interface
// returns map of field name with type
// source: http://merbist.com/2011/06/27/golang-reflection-exampl/
func Attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}
	return attrs
}

// GetAttribVal used for getting value of given field
// in provided interface
func GetAttribVal(m interface{}, attribute string) map[string]reflect.Value {
	rs := make(map[string]reflect.Value)
	v := reflect.ValueOf(m).Elem()
	t := v.FieldByName(attribute).Type().String()
	val := reflect.ValueOf(v.FieldByName(attribute).Interface())
	rs[t] = val
	return rs
}

// func GetAttrs(m interface{}) string {
// 	rs := ""
// 	t := reflect.TypeOf(m)
// 	rs += fmt.Sprintf("%v", t)
// 	// Get the type and kind of our user variable
// 	rs += fmt.Sprintf("Type: %s", t.Name())
// 	rs += fmt.Sprintf("Kind: %s", t.Kind())
// 	return rs
// }

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
//--------------------------------------------------------------------
func MigrateStructVals(pa interface{}, pb interface{}) {
	m := BuildFieldMap(pb) // we'll map pb's fields, then process pa one field at a time

	ar := reflect.ValueOf(pa).Elem()
	for i := 0; i < ar.NumField(); i++ {
		afield := ar.Field(i)
		afldname := ar.Type().Field(i).Name
		if !afield.IsValid() { // skip fields in an invalid state, nil pointers, zero-valued lists, ...
			continue
		}
		bdx, ok := m[afldname]
		if !ok { // if pb doesn't have this field, move on
			continue
		}
		br := reflect.ValueOf(pb).Elem()
		bfield := br.Field(bdx)
		if !bfield.CanSet() { // if it cannot be set then just move on
			continue
		}
		switch afield.Type().String() { // we need to copy things differently, depending on the data type
		case "int", "int64", "float64", "time.Time", "[]int":
			bfield.Set(reflect.ValueOf(afield.Interface()))
		case "string":
			bfield.Set(reflect.ValueOf(afield.Interface()))
		default:
			fmt.Printf("Unhandled data type: %s\n", afield.Type().String())
		}
	}
}
