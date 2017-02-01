package core

import (
	"regexp"
	"strconv"
)

// StringInSlice used to check whether string a
// is present or not in slice list
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IntegerInSlice used to check whether int a
// is present or not in slice list
func IntegerInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsIntString used to check whether string a
// contains integer type value
func IsIntString(a string) bool {
	_, err := strconv.ParseInt(a, 10, 64)
	return err == nil
}

// IsUIntString used to check whether string a
// contains integer type value
func IsUIntString(a string) bool {
	_, err := strconv.ParseUint(a, 10, 64)
	return err == nil
}

// IsFloatString used to check whether string a
// contains float type value
func IsFloatString(a string) bool {
	_, err := strconv.ParseFloat(a, 64)
	return err == nil
}

// IsValidEmail used to check valid email or not
func IsValidEmail(email string) bool {
	// TODO: confirm which regex to use
	// Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+$`)
	return Re.MatchString(email)
}

// IsValidPhone used to check valid phone or not
func IsValidPhone(email string) bool {
	// Re := regexp.MustCompile(`^[0-9+\-\)\(]{1,14}$`)
	// return Re.MatchString(email)
	return true
}
