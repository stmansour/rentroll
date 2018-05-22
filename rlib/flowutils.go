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

// IsFlowDataValidJSON checks that passed flow data is valid json or not
// This data will be inserted/updated to data with json type column
func IsFlowDataValidJSON(raw json.RawMessage) bool {
	_, err := json.Marshal(&raw)
	return err == nil
}
