package rtags

import (
	"fmt"
	"reflect"
	"strings"
)

// getValidatorFromTag returns validator struct corresponding to validation type
// validation type should be defined in first argument of tag string value
// Example:
//      // For number validation type
//      Field int `validate:"number,min=2.max=3"`
//
//      // for string validation type
//      Field int `validate:"string,min=2.max=3"`
//
func getValidatorFromTag(tagValue, fieldName string) Validator {

	// replace any misplaced whitespace in tag value
	tagValue = strings.Replace(tagValue, " ", "", -1)

	// get args of tagValue by splitting ","
	args := strings.Split(tagValue, ",")

	// switch case of numerous validation type
	switch args[0] {
	case "string":
		return getStringValidatorFromTagValues(strings.Join(args[1:], ","), fieldName)
	case "email":
		return getEmailValidatorFromTagValues(strings.Join(args[1:], ","), fieldName)
	case "number":
		return getNumberValidatorFromTagValues(strings.Join(args[1:], ","), fieldName)
	default:
		return DefaultValidator{}
	}
}

// ValidateStructFromTagRules validate struct fields from tag rules
// followed by `validate` keyword.
//
// As a special case, if the field tag is "-", the field is always omitted.
//
// Struct Example to define validate rules
//      type User struct {
//          Id    int    `validate:"-"`
//          Name  string `validate:"string,min=2,max=10"`
//          Email string `validate:"string,min=3,max=32"`
//      }
//
func ValidateStructFromTagRules(s interface{}) []error {
	errs := []error{}

	// get reflected value
	v := reflect.ValueOf(s)

	// if v is pointer then
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// loop over list of fields
	for i := 0; i < v.NumField(); i++ {
		// get tag by `validate` keyword
		tag := v.Type().Field(i).Tag.Get(tagName)

		// skip if tag not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// get validator from the tag string value
		validator := getValidatorFromTag(tag, v.Type().Field(i).Name)

		// perform validation on field interface value
		err := validator.Validate(v.Field(i).Interface())

		// append error to list
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}
