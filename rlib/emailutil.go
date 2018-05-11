package rlib

import "regexp"

// Here are a few email addresses that caused problems before. The regexp
// has been adjusted so that these are all correctly filtered out.
//
//    ryan.t.baldwain@.com
//    mp_sargechuck@yahoo.

var validemail = regexp.MustCompile(`^([a-zA-Z0-9][-_.a-zA-Z0-9]*)@([-_.a-zA-Z0-9]+)\.([a-zA-Z]+)$`)

// ValidEmailAddress parses s and ensures that it conforms to the regexp above,
// which to the best of my ability represents an email address.
//
// INPUT
// s - the email address to verify
//
// RETURNS
// bool - true = the address is valid, otherwise false
//-----------------------------------------------------------------------------
func ValidEmailAddress(s string) bool {
	return validemail.MatchString(s)
}
