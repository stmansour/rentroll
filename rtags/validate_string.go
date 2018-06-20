package rtags

import (
	"fmt"
	"rentroll/rlib"
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
func getStringValidatorFromTagValues(tagOptions []string, fieldName string) StringValidator {
	validator := StringValidator{}

	// loop over the list of options with values
	for _, opt := range tagOptions {
		switch {
		case strings.Contains(opt, "min"): // min option
			mn := StringMin{Value: 0, Skip: false}

			_, err := fmt.Sscanf(opt, "min=%d", &mn.Value)
			if err != nil {
				// NOTE: this should be marked as critical level
				rlib.Console("Field: `%s`, Option `min` value does not match `min=NUMBER` format in `string` validation type: %s\n", fieldName, err.Error())

				// mark skip as true in case of error
				mn.Skip = true
			}

			// assign struct
			validator.Min = mn

		case strings.Contains(opt, "max"): // max option
			mx := StringMax{Value: 0, Skip: false}

			_, err := fmt.Sscanf(opt, "max=%d", &mx.Value)
			if err != nil {
				// NOTE: this should be marked as critical level
				rlib.Console("Field: `%s`, Option `max` value does not match `max=NUMBER` format in `string` validation type: %s\n", fieldName, err.Error())

				// mark skip as true in case of error
				mx.Skip = true
			}

			// assign struct
			validator.Max = mx
		}
	}

	return validator
}
