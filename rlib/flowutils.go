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

// IsFlowDataValidJSON checks that passed flow data is valid json or not
// This data will be inserted/updated to data with json type column
func IsFlowDataValidJSON(raw json.RawMessage) bool {
	_, err := json.Marshal(&raw)
	return err == nil
}

// Alphabet contains caps of the alphabet
var Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Digits contains characters for 0 - 9
var Digits = "0123456789"

// GenerateUserRefNo generate a unique identifier that can be given to users
// to refer to the Flow.  Given the sensitive data, we cannot use the FlowID
// which is monotonically increasing and too easy to guess, or mistype resulting
// in a valid FlowID that would give a user access to information that they
// should not have. This generates an id that is highly unique so if the user
// mistypes it, there is almost no chance that the result will be a valid ID
//
// INPUTS:
//
// RETURNS:
//     the id string
//-----------------------------------------------------------------------------
func GenerateUserRefNo() string {
	var l []byte

	// Generate 5 random digits and 5 random letters
	for i := 0; i < 5; i++ {
		l = append(l, Alphabet[RRdb.Rand.Intn(26)])
	}
	for i := 0; i < 5; i++ {
		l = append(l, Digits[RRdb.Rand.Intn(10)])
	}
	// move them around some random number of times
	for i := 0; i < 4+RRdb.Rand.Intn(4); i++ {
		j := RRdb.Rand.Intn(7)
		k := RRdb.Rand.Intn(7)
		l[k], l[j] = l[j], l[k]
	}
	return string(l)
}
