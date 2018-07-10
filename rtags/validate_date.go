package rtags

import (
	"fmt"
	"rentroll/rlib"
	"time"
)

// DateValidator for date validation
type DateValidator struct {
}

// Validate method for DateValidator
func (v DateValidator) Validate(val interface{}) error {

	// Date: Jan 1, 2000 00:00:00 UTC
	var earliestDate = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

	// don't panic
	s, ok := val.(rlib.JSONDate)
	if !ok {
		return fmt.Errorf("should be type of date")
	}

	// blank check
	// It done while marshalling/unmarshalling rlib.JSONDate fields

	// get the date timestamp
	ts := time.Time(s)

	// Dates must be Jan 1, 2000 00:00:00 UTC or later
	if ts.Before(earliestDate) {
		return fmt.Errorf("Dates must be Jan 1, 2000 00:00:00 UTC or later")
	}

	return nil
}

// getDateValidatorFromTagValues returns instantiated `DateValidator`
// from passed tag value options
func getDateValidatorFromTagValues(tagOptions, fieldName string) DateValidator {
	// Here don't require extra thing to do for date validator
	return DateValidator{}
}
