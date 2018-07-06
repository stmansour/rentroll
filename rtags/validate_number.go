package rtags

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// IntNumberMin struct for minimum value of a number validation
type IntNumberMin struct {
	Value int64
	Skip  bool
}

// IntNumberMax struct for maximum value of a number validation
type IntNumberMax struct {
	Value int64
	Skip  bool
}

// FloatNumberMin struct for minimum value of a number validation
type FloatNumberMin struct {
	Value float64
	Skip  bool
}

// FloatNumberMax struct for maximum value of a number validation
type FloatNumberMax struct {
	Value float64
	Skip  bool
}

// IntegerNumberValidator for string validation
type IntegerNumberValidator struct {
	Min IntNumberMin
	Max IntNumberMax
}

// FloatNumberValidator for string validation
type FloatNumberValidator struct {
	Min FloatNumberMin
	Max FloatNumberMax
}

// Validate method for IntegerNumberValidator
func (v IntegerNumberValidator) Validate(val interface{}) error {
	// don't panic
	n, ok := val.(int) // TODO(Akshay): Convert it to int64 and than check its type
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
		return fmt.Errorf("should be less than %d", v.Max.Value)
	}

	return nil
}

// Validate method for FloatNumberValidator
func (v FloatNumberValidator) Validate(val interface{}) error {
	// don't panic
	n, ok := val.(float64)
	if !ok {
		return fmt.Errorf("should be a decimal number")
	}

	// get float64 compatible number
	num := float64(n)

	// min length check
	if !v.Min.Skip && num < v.Min.Value {
		return fmt.Errorf("should be greater than %.2f", v.Min.Value)
	}

	// max length check
	if !v.Max.Skip && v.Max.Value >= v.Min.Value && num > v.Max.Value {
		return fmt.Errorf("should be less than %.2f", v.Max.Value)
	}

	return nil
}

// getIntegerNumberValidatorFromTagValues returns instantiated `IntegerNumberValidator`
// from passed tag value options
func getIntegerNumberValidatorFromTagValues(tagOptions, fieldName string) IntegerNumberValidator {
	validator := IntegerNumberValidator{
		Min: IntNumberMin{Value: 0, Skip: false},
		Max: IntNumberMax{Value: 0, Skip: false},
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

// getFloatNumberValidatorFromTagValues returns instantiated `FloatNumberValidator`
// from passed tag value options
func getFloatNumberValidatorFromTagValues(tagOptions, fieldName string) FloatNumberValidator {
	validator := FloatNumberValidator{
		Min: FloatNumberMin{Value: 0.00, Skip: false},
		Max: FloatNumberMax{Value: 0.00, Skip: false},
	}

	var err error

	// min option
	minStr := strings.Split(tagOptions, "min=")
	if len(minStr) > 1 {
		nStr := strings.Split(minStr[1], ",")[0]
		validator.Min.Value, err = strconv.ParseFloat(nStr, 64)
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
		validator.Max.Value, err = strconv.ParseFloat(nStr, 64)
		if err != nil {
			rlib.Console("Field: %s, Option `max` value parsing error: %s", fieldName, err.Error())
			validator.Max.Skip = true
		}
	} else { // in case not found then
		validator.Max.Skip = true
	}

	return validator
}
