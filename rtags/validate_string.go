package rtags

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// StringMin struct for string minimum length validation
type StringMin struct {
	Value int
	Skip  bool
}

// StringMax struct for string maximum length validation
type StringMax struct {
	Value int
	Skip  bool
}

// StringValidator for string validation
type StringValidator struct {
	Min StringMin
	Max StringMax
}

// Validate method for StringValidator
func (v StringValidator) Validate(val interface{}) error {
	// don't panic
	s, ok := val.(string)
	if !ok {
		return fmt.Errorf("should be type of string")
	}

	// get length of string
	sl := len(s)

	// blank check
	if sl == 0 {
		return fmt.Errorf("cannot be blank")
	}

	// min length check
	if !v.Min.Skip && sl < v.Min.Value {
		return fmt.Errorf("should be at least %d chars long", v.Min.Value)
	}

	// max length check, max should be >= min
	if !v.Max.Skip && v.Max.Value >= v.Min.Value && sl > v.Max.Value {
		return fmt.Errorf("should be less than %d chars", v.Max.Value)
	}

	return nil
}

// getStringValidatorFromTagValues returns instantiated `StringValidator`
// from passed tag value options
func getStringValidatorFromTagValues(tagOptions, fieldName string) StringValidator {
	validator := StringValidator{
		Min: StringMin{Value: 0, Skip: false},
		Max: StringMax{Value: 0, Skip: false},
	}

	var err error

	// min option
	minStr := strings.Split(tagOptions, "min=")
	if len(minStr) > 1 {
		nStr := strings.Split(minStr[1], ",")[0]
		validator.Min.Value, err = strconv.Atoi(nStr)
		if err != nil {
			rlib.Console("Field: %s, Option `min` value parsing error: %s", fieldName, err.Error())
			validator.Min.Skip = true
		}
	} else { // in case not found then
		validator.Min.Skip = true
	}

	// max option
	maxStr := strings.Split(tagOptions, "max=")
	if len(maxStr) > 1 {
		nStr := strings.Split(maxStr[1], ",")[0]
		validator.Max.Value, err = strconv.Atoi(nStr)
		if err != nil {
			rlib.Console("Field: %s, Option `min` value parsing error: %s", fieldName, err.Error())
			validator.Max.Skip = true
		}
	} else { // in case not found then
		validator.Max.Skip = true
	}

	return validator
}
