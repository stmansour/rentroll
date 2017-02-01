package core

import (
	"reflect"
)

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
