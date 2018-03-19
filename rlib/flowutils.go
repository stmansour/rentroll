package rlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// ErrFlowInvalidJSONData etc.. all are error constants used for flows
var ErrFlowInvalidJSONData = errors.New("Invalid JSON data")

// getFlowID will return unique ID with combination of unix nano, userid
func getFlowID(UserID int64) string {
	u := uint32(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%x-%x", u, UserID)
}

// IsFlowDataJSON checks that passed flow data is valid json or not
// This data will be inserted/updated to data with json type column
func IsFlowDataJSON(b []byte) bool {
	var raw json.RawMessage
	return json.Unmarshal(b, &raw) == nil
}
