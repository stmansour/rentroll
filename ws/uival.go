package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// BIDToBUD returns bud from cached list (rlib.RRdb.BUDlist)
func BIDToBUD(businessID int64) (string, error) {
	for bud, bid := range rlib.RRdb.BUDlist {
		if businessID == bid {
			return bud, nil
		}
	}
	rlib.Console("*** ERROR *** Could not find business for bid: %d\n", businessID)
	rlib.Console("RRdb.BUDlist = %#v\n", rlib.RRdb.BUDlist)
	return "", fmt.Errorf("Could not find business for bid: %d", businessID)
}

func getListTypes(ctx context.Context, bid int64, s string, t int) (map[string][]IDTextMap, error) {
	list := []IDTextMap{{ID: 0, Text: s}}   // initialize list with 0-id value
	appData := make(map[string][]IDTextMap) // json response data
	bud, err := BIDToBUD(bid)
	if err != nil {
		return appData, err
	}
	m, _ := rlib.GetARsByType(ctx, bid, t) // get records and append in IDTextMap list
	for i := 0; i < len(m); i++ {
		list = append(list, IDTextMap{ID: m[i].ARID, Text: m[i].Name})
	}
	appData[bud] = list
	return appData, nil
}

// GetAssessmentList returns all assessments for the supplied business
//-----------------------------------------------------------------------------
func GetAssessmentList(ctx context.Context, bid int64) (map[string][]IDTextMap, error) {
	return getListTypes(ctx, bid, " -- Select Account Rule -- ", rlib.ARASSESSMENT)
}

// GetExpenseList returns all assessments for the supplied business
//-----------------------------------------------------------------------------
func GetExpenseList(ctx context.Context, bid int64) (map[string][]IDTextMap, error) {
	return getListTypes(ctx, bid, " -- Select Expense Rule -- ", rlib.AREXPENSE)
}

// GetReceiptList returns all assessments for the supplied business
//-----------------------------------------------------------------------------
func GetReceiptList(ctx context.Context, bid int64) (map[string][]IDTextMap, error) {
	return getListTypes(ctx, bid, " -- Select Receipt Rule -- ", rlib.ARRECEIPT)
}

// GetDepositoryList returns all assessments for the supplied business
//-----------------------------------------------------------------------------
func GetDepositoryList(ctx context.Context, bid int64) (map[string][]IDTextMap, error) {

	// initialize list with 0-id value
	list := []IDTextMap{{ID: 0, Text: " -- Select Depository -- "}}

	// json response data
	appData := make(map[string][]IDTextMap)

	bud, err := BIDToBUD(bid)
	if err != nil {
		return appData, err
	}

	m, _ := rlib.GetAllDepositories(ctx, bid)
	for i := 0; i < len(m); i++ {
		list = append(list, IDTextMap{ID: m[i].DEPID, Text: m[i].Name})
	}
	appData[bud] = list

	return appData, nil
}

// GetTLDs returns a map of TaskListDefinitions for the UI
//-----------------------------------------------------------------------------
func GetTLDs(ctx context.Context, bid int64) (map[string][]IDTextMap, error) {
	rlib.Console("Entered GetTLDS\n")
	// initialize list with 0-id value
	list := []IDTextMap{{ID: 0, Text: " -- Select TaskList -- "}}
	appData := make(map[string][]IDTextMap)
	bud, err := BIDToBUD(bid)
	if err != nil {
		return appData, err
	}

	m, err := rlib.GetAllTaskListDefinitions(ctx, bid)
	if err != nil {
		return appData, err
	}
	rlib.Console("GetAllTaskListDefinitions return m, len = %d\n", len(m))

	for i := 0; i < len(m); i++ {
		list = append(list, IDTextMap{ID: m[i].TLDID, Text: m[i].Name})
	}
	appData[bud] = list

	return appData, nil

}

// SvcUIErrAndVarResponse encapsulates a lot of lines that would need to appear
// in each case of a switch.  This just makes things a lot more readable and
// it bottlenecks the handling so it is easy to extend or modify.
//-----------------------------------------------------------------------------
func SvcUIErrAndVarResponse(w http.ResponseWriter, funcname string, err error, x interface{}) {
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(x)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	SvcWrite(w, b)
	// if err := json.NewEncoder(w).Encode(x); err != nil {

	// }
}

// SvcUIVal returns the requested variable in JSON form
//
// wsdoc {
//  @Title  Get UI Value
//	@URL /v1/uival/:BID/varname
//  @Method  GET
//	@Synopsis Return JSON representing the UI Value
//  @Desc Return data can be parsed to create the string lists used in the UI.
//	@Input
//  @Response JSONResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcUIVal(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcUIVal"
	rlib.Console("Entered %s\n", funcname)
	switch d.DetVal {
	case "app.AssessmentRules":
		asmData, err := GetAssessmentList(r.Context(), d.BID)
		SvcUIErrAndVarResponse(w, funcname, err, asmData)
	case "app.ReceiptRules":
		rcptData, err := GetReceiptList(r.Context(), d.BID)
		SvcUIErrAndVarResponse(w, funcname, err, rcptData)
	case "app.ExpenseRules":
		expData, err := GetExpenseList(r.Context(), d.BID)
		SvcUIErrAndVarResponse(w, funcname, err, expData)
	case "app.Depositories":
		data, err := GetDepositoryList(r.Context(), d.BID)
		SvcUIErrAndVarResponse(w, funcname, err, data)
	case "app.DepMethods":
		depmeth := GetJSDepositMethods(r.Context())
		SvcUIErrAndVarResponse(w, funcname, nil, depmeth)
	case "app.TaskListDefinitions":
		tlds, err := GetTLDs(r.Context(), d.BID)
		SvcUIErrAndVarResponse(w, funcname, err, tlds)
	case "app.Applicants":
		var a []rlib.StringList
		a, err := rlib.GetAllStringLists(r.Context(), d.BID)
		SvcUIErrAndVarResponse(w, funcname, err, a)
	default:
		e := fmt.Errorf("Unknown variable requested: %s", d.DetVal)
		SvcErrorReturn(w, e, funcname)
		return
	}
}
