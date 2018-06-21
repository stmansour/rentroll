package rtags

// tagname to parse validation options from struct field tag
const (
	tagName = "validate"
)

// Validator interface which only have Validate method
type Validator interface {
	Validate(interface{}) error
}

// DefaultValidator which always passed in validation
type DefaultValidator struct {
}

// Validate method for DefaultValidator
func (v DefaultValidator) Validate(val interface{}) error {
	return nil
}
