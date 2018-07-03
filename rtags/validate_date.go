package rtags

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

//
var earliestDate = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)

// DateValidator for date validation
type DateValidator struct {
}

// Validate method for DateValidator
func (v DateValidator) Validate(val interface{}) error {

	fmt.Println("**************")
	fmt.Println("DateValidator")
	fmt.Println("**************")
	s, ok := val.(string)
	if !ok {
		return fmt.Errorf("should be type of string")
	}

	x, err := rlib.StringToDate(s)
	if err != nil {
		return fmt.Errorf("date is not in a valid format")
	}
	if x.Before(earliestDate) {
		x = earliestDate
		return fmt.Errorf("it should be earliest date %s", x)
	}

	return nil
}

// getDateValidatorFromTagValues returns instantiated `DateValidator`
// from passed tag value options
func getDateValidatorFromTagValues(tagOptions, fieldName string) DateValidator {
	return DateValidator{}
}
