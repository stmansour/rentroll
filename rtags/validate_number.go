package rtags

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// NumberMin struct for minimum value of a number validation
type NumberMin struct {
	Value int64
	Skip  bool
}

// NumberMax struct for maximum value of a number validation
type NumberMax struct {
	Value int64
	Skip  bool
}

// NumberValidator for string validation
type NumberValidator struct {
	Min NumberMin
	Max NumberMax
}

// Validate method for NumberValidator
func (v NumberValidator) Validate(val interface{}) error {
	// don't panic
	n, ok := val.(int)
	if !ok {
		return fmt.Errorf("should be type of number")
	}

	// get int64 compatible number
	num := int64(n)

	// min length check
	if !v.Min.Skip && num < v.Min.Value {
		return fmt.Errorf("should be greater than %d", v.Min.Value)
	}

	// max length check
	if !v.Max.Skip && v.Max.Value >= v.Min.Value && num > v.Max.Value {
		return fmt.Errorf("should be less than %d chars", v.Max.Value)
	}

	return nil
}

// getNumberValidatorFromTagValues returns instantiated `NumberValidator`
// from passed tag value options
func getNumberValidatorFromTagValues(tagOptions, fieldName string) NumberValidator {
	validator := NumberValidator{
		Min: NumberMin{Value: 0, Skip: false},
		Max: NumberMax{Value: 0, Skip: false},
	}

	var err error

	// min option
	minStr := strings.Split(tagOptions, "min=")
	if len(minStr) > 1 {
		nStr := strings.Split(minStr[1], ",")[0]
		validator.Min.Value, err = strconv.ParseInt(nStr, 10, 64)
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
		validator.Max.Value, err = strconv.ParseInt(nStr, 10, 64)
		if err != nil {
			rlib.Console("Field: %s, Option `max` value parsing error: %s", fieldName, err.Error())
			validator.Max.Skip = true
		}
	} else { // in case not found then
		validator.Max.Skip = true
	}

	return validator
}
