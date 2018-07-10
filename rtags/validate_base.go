package rtags

import "reflect"

// tagname to parse validation options from struct field tag
const (
	tagName = "validate"
)

// Validator interface which only have Validate method
type Validator interface {
	Validate(interface{}) error
}

// DefaultValidator which always passed in validation
type DefaultValidator struct {
}

// Validate method for DefaultValidator
func (v DefaultValidator) Validate(val interface{}) error {
	return nil
}

// isEmptyValue to check value is blank, nil, zero
// Reference Link: https://golang.org/src/encoding/json/encode.go
func isEmptyValue(v reflect.Value) bool {

	switch v.Kind() {

	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0

	case reflect.Bool:
		return !v.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Interface, reflect.Ptr:
		return v.IsNil()

	}

	return false
}
