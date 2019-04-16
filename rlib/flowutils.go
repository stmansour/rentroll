package rlib

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// ErrFlowInvalidJSONData etc.. all are error constants used for flows
var ErrFlowInvalidJSONData = errors.New("Invalid JSON data")

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

// RandCharSet is the list of characters that we'll use for the IDs we generate.
// They have been chosen to minimize the chance of error for human beings who
// may need to type them or communicate them by voice.
var RandCharSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// IDBase is the base number of the code we're generating. Essentially, we
// are converting the random number to base x where x is the number of
// characters in RandCharSet.
var IDBase = int64(len(RandCharSet))

// Timestamp of the last code generated, to help avoid collisions
var lastTimeMs int64

// We generate 72-bits of randomness which get turned into 12 characters and appended to the
// timestamp to prevent collisions with other clients.  We store the last characters we
// generated because in the event of a collision, we'll use those same characters except
// "incremented" by one.
var lastID [12]int64
var mu sync.Mutex

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
// func GenerateUserRefNo() string {
// 	var l []byte
//
// 	// Generate half letters and half digits
// 	l1 := UserRefNoLength / 2
// 	l2 := UserRefNoLength - l1
// 	for i := 0; i < l1; i++ {
// 		l = append(l, Alphabet[RRdb.Rand.Intn(26)])
// 	}
// 	for i := 0; i < l2; i++ {
// 		l = append(l, Digits[RRdb.Rand.Intn(10)])
// 	}
// 	// move them around some random number of times
// 	swaps := 5 + RRdb.Rand.Intn(10)
// 	for i := 0; i < swaps; i++ {
// 		j := RRdb.Rand.Intn(10)
// 		k := 10 + RRdb.Rand.Intn(10)
// 		l[k], l[j] = l[j], l[k]
// 	}
// 	return string(l)
// }
func GenerateUserRefNo() string {
	l := len(lastID)
	var guid = make([]byte, 20)
	mu.Lock()
	timeMs := time.Now().UTC().UnixNano() / 1e6
	if timeMs == lastTimeMs {
		// increment lastRandChars
		for i := 0; i < l; i++ {
			lastID[i]++
			if lastID[i] < IDBase {
				break
			}
			lastID[i] = 0 // set this to 0 and inc the next char
		}
	} else {
		for i := 0; i < l; i++ {
			lastID[i] = int64(RRdb.Rand.Intn(int(IDBase)))
		}
	}
	lastTimeMs = timeMs // save for comparison next time
	// put random as the second part
	for i := 0; i < l; i++ {
		guid[19-i] = RandCharSet[lastID[i]]
	}
	mu.Unlock()

	// put current time at the beginning
	for i := 7; i >= 0; i-- {
		n := int(timeMs % IDBase)
		guid[i] = RandCharSet[n]
		timeMs = timeMs / IDBase
	}
	return string(guid[:])
}
