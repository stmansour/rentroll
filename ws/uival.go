package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// returns bud from cached list (rlib.RRdb.BUDlist)
func bidToBud(businessID int64) (string, error) {
	for bud, bid := range rlib.RRdb.BUDlist {
		if businessID == bid {
			return bud, nil
		}
	}
	return "", fmt.Errorf("Could not find business for bid: %d", businessID)
}

// GetAssessmentList returns all assessments for the supplied business
func GetAssessmentList(bid int64) (map[string][]IDTextMap, error) {

	// initialize list with 0-id value
	list := []IDTextMap{{ID: 0, Text: "Select Assessment Rule"}}

	// json response data
	appData := make(map[string][]IDTextMap)

	bud, err := bidToBud(bid)
	if err != nil {
		return appData, err
	}

	// get records and append in IDTextMap list
	m := rlib.GetARsByType(bid, rlib.ARASSESSMENT)
	for i := 0; i < len(m); i++ {
		list = append(list, IDTextMap{ID: m[i].ARID, Text: m[i].Name})
	}
	appData[bud] = list

	return appData, nil
}

// GetReceiptList returns all assessments for the supplied business
func GetReceiptList(bid int64) (map[string][]IDTextMap, error) {

	// initialize list with 0-id value
	list := []IDTextMap{{ID: 0, Text: "Select Receipt Rule"}}

	// json response data
	appData := make(map[string][]IDTextMap)

	bud, err := bidToBud(bid)
	if err != nil {
		return appData, err
	}

	// get records and append in IDTextMap list
	m := rlib.GetARsByType(bid, rlib.ARRECEIPT)
	for i := 0; i < len(m); i++ {
		list = append(list, IDTextMap{ID: m[i].ARID, Text: m[i].Name})
	}
	appData[bud] = list

	return appData, nil
}

// SvcUIVal returns the requested variable in JSON form
//
// wsdoc {
//  @Title  Get UI Value
//	@URL /v1/uival/:BID/varname
//  @Method  GET
//	@Synopsis Return JSON representing the UI Value
//  @Desc Return data can be processed by eval() to create the string lists used in the UI.
//	@Input
//  @Response JSONResponse
// wsdoc }
func SvcUIVal(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcUIVar"
	fmt.Printf("Entered %s\n", funcname)
	switch d.DetVal {
	case "app.Assessments":
		// get assessments data
		asmData, err := GetAssessmentList(d.BID)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// send down then json stuff
		if err := json.NewEncoder(w).Encode(asmData); err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
	case "app.Receipts":
		// get receipts data
		rcptData, err := GetReceiptList(d.BID)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		// send down then json stuff
		if err := json.NewEncoder(w).Encode(rcptData); err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}
	default:
		e := fmt.Errorf("Unknown variable requested: %s", d.DetVal)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
}
