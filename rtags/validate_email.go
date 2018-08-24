package rtags

import (
	"fmt"
	"regexp"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// Regular expression to validate email address.
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// StringMin struct for string minimum length validation
type EmailStringMin struct {
	Value int
	Skip  bool
}

// StringMax struct for string maximum length validation
type EmailStringMax struct {
	Value int
	Skip  bool
}

// EmailValidator for email validation
type EmailValidator struct {
	Min EmailStringMin
	Max EmailStringMax
}

// Validate method for EmailValidator
func (v EmailValidator) Validate(val interface{}) error {
	s, ok := val.(string)
	if !ok {
		return fmt.Errorf("must be type of string")
	}

	// get length of string
	sl := len(s)

	// min length check
	if !v.Min.Skip && sl < v.Min.Value {
		return fmt.Errorf("must be at least %d chars long", v.Min.Value)
	}

	// max length check, max should be >= min
	if !v.Max.Skip && v.Max.Value >= v.Min.Value && sl > v.Max.Value {
		return fmt.Errorf("must be less than %d chars", v.Max.Value)
	}

	// it does match regex pattern then raise an error
	if !mailRe.MatchString(s) {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}

// getEmailValidatorFromTagValues returns instantiated `EmailValidator`
// from passed tag value options
func getEmailValidatorFromTagValues(tagOptions, fieldName string) EmailValidator {
	validator := EmailValidator{
		Min: EmailStringMin{Value: 0, Skip: false},
		Max: EmailStringMax{Value: 0, Skip: false},
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
