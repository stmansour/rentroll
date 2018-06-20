package rtags

import (
	"fmt"
	"regexp"
)

// Regular expression to validate email address.
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// EmailValidator for email validation
type EmailValidator struct {
}

// Validate method for EmailValidator
func (v EmailValidator) Validate(val interface{}) error {
	s, ok := val.(string)
	if !ok {
		return fmt.Errorf("should be type of string")
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
	return EmailValidator{}
}
