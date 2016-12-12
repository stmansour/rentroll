package core

import (
	"reflect"
)

// BuildFieldMap build a map of the supplied struct pointer so that
// you can index the map by a field name and the associated value is
// the field's index within p
func BuildFieldMap(p interface{}) (map[string]int, bool) {
	fmap, ok := map[string]int{}, false

	v := reflect.ValueOf(p).Elem()

	// if p is not type of struct then return
	if v.Kind() != reflect.Struct {
		return fmap, ok
	}

	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name
		fmap[n] = j
	}
	ok = true
	return fmap, ok
}

// GetStructFields return the fields of struct pointer
// in order which they are defined
func GetStructFields(p interface{}) ([]string, bool) {
	fields, ok := []string{}, false

	v := reflect.ValueOf(p).Elem()

	// if p is not type of struct then return
	if v.Kind() != reflect.Struct {
		return fields, ok
	}

	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name
		fields = append(fields, n)
	}
	ok = true
	return fields, ok
}
