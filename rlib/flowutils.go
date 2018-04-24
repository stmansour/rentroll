package rlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// ErrFlowInvalidJSONData etc.. all are error constants used for flows
var ErrFlowInvalidJSONData = errors.New("Invalid JSON data")

// GetFlowID will return unique ID with combination of unix nano, userid
func GetFlowID() string {
	u := uint32(time.Now().UTC().UnixNano())
	return fmt.Sprintf("%x", u)
}

/*// IsByteDataValidJSON checks that passed bytes data is valid json or not
func IsByteDataValidJSON(b []byte) bool {
	var raw json.RawMessage
	return json.Unmarshal(b, &raw) == nil
}*/

// IsFlowDataValidJSON checks that passed flow data is valid json or not
// This data will be inserted/updated to data with json type column
func IsFlowDataValidJSON(raw json.RawMessage) bool {
	_, err := json.Marshal(&raw)
	return err == nil
}
