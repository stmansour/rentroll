package rlib

import (
	"encoding/json"
	"errors"
)

// ErrFlowInvalidJSONData etc.. all are error constants used for flows
var ErrFlowInvalidJSONData = errors.New("Invalid JSON data")

/*// GetFlowID will return unique ID with combination of unix nano, userid
func GetFlowID() string {
	u := uint32(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%x", u)
}*/

// IsByteDataValidJSON checks that passed bytes data is valid json or not
func IsByteDataValidJSON(b []byte) bool {
	var raw json.RawMessage
	return json.Unmarshal(b, &raw) == nil
}

// IsValidJSONConversion checks that passed flow data is valid json or not
// This data will be inserted/updated to data with json type column
func IsValidJSONConversion(raw json.RawMessage) bool {
	_, err := json.Marshal(&raw)
	return err == nil
}

// Alphabet contains caps of the alphabet
var Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Digits contains characters for 0 - 9
var Digits = "0123456789"

// UserRefNoLength specifies the number of characters needed in the RefNo
// in order to be "safe"
const UserRefNoLength = 20

// GenerateUserRefNo generate a unique identifier that can be given to users
// to refer to the Flow.
//
// It can generate a settable number of characters.  Initialized to 20.
//
// Given the sensitive data, we cannot use the FlowID
// which is monotonically increasing and too easy to guess, or mistype resulting
// in a valid FlowID that would give a user access to information that they
// should not have. This generates an id that is highly unique so if the user
// mistypes it, there is almost no chance that the result will be a valid ID
//
// https://play.golang.org/p/h6icSimZK2M
//
// INPUTS:
//
// RETURNS:
//     the id string
//-----------------------------------------------------------------------------
func GenerateUserRefNo() string {
	var l []byte

	// Generate half letters and half digits
	l1 := UserRefNoLength / 2
	l2 := UserRefNoLength - l1
	for i := 0; i < l1; i++ {
		l = append(l, Alphabet[RRdb.Rand.Intn(26)])
	}
	for i := 0; i < l2; i++ {
		l = append(l, Digits[RRdb.Rand.Intn(10)])
	}
	// move them around some random number of times
	swaps := 5 + RRdb.Rand.Intn(10)
	for i := 0; i < swaps; i++ {
		j := RRdb.Rand.Intn(10)
		k := 10 + RRdb.Rand.Intn(10)
		l[k], l[j] = l[j], l[k]
	}
	return string(l)
}
