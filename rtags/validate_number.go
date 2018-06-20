package rtags

import (
	"fmt"
	"rentroll/rlib"
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
func getNumberValidatorFromTagValues(tagOptions []string, fieldName string) NumberValidator {
	validator := NumberValidator{}

	// loop over the list of options with values
	for _, opt := range tagOptions {
		switch {
		case strings.Contains(opt, "min"): // min option
			mn := NumberMin{Value: 0, Skip: false}

			_, err := fmt.Sscanf(opt, "min=%d", &mn.Value)
			if err != nil {
				// NOTE: this should be marked as critical level
				rlib.Console("Field: `%s`, Option `min` value does not match `min=NUMBER` format in `number` validation type: %s\n", fieldName, err.Error())

				// mark skip as true in case of error
				mn.Skip = true
			}

			// assign struct
			validator.Min = mn

		case strings.Contains(opt, "max"): // max option
			mx := NumberMax{Value: 0, Skip: false}

			_, err := fmt.Sscanf(opt, "max=%d", &mx.Value)
			if err != nil {
				// NOTE: this should be marked as critical level
				rlib.Console("Field: `%s`, Option `max` value does not match `max=NUMBER` format in `number` validation type: %s\n", fieldName, err.Error())

				// mark skip as true in case of error
				mx.Skip = true
			}

			// assign struct
			validator.Max = mx
		}
	}

	return validator
}
