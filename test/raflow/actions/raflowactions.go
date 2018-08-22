package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"rentroll/rlib"
	"rentroll/ws"
	"time"
)

const (
	statusError = "error"

	ExistingRAValidFlow   = "FU1T222ATL6HWFS61388"
	ExistingRAValidFlow1  = "ON5SY742BDK19L0D9M34"
	ExistingRAValidFlow2  = "NDSD46NIKR363005L3I8"
	ExistingRAInValidFlow = "VJFC558GW9MM625CT176"
	BrandNewValidFlow     = "HL76CY7PA47W6H9U8K38"
	BrandNewInValidFlow   = "YCE20N8G44N45TIW1M95"

	VersionRefNo = "refno"
	VersionRAID  = "raid"

	ModeAction = "Action"
	ModeState  = "State"

	DecisionAccept  = 1
	DecisionDecline = 2
)

var testNames = map[int]string{
	1:  "action \"set pending first approval\" on flow with invalid data",
	2:  "action \"set pending first approval\" on flow with valid data",
	3:  "approve and set \"pending second approval\" on flow with valid data",
	4:  "approve and set \"move-in / execute modification\" on flow with valid data",
	5:  "set document date of flow with valid data",
	6:  "take action of \"complete move in\" on flow with valid data",
	7:  "action \"set pending first approval\" on brand new flow with invalid data",
	8:  "action \"set pending first approval\" on brand new flow with valid data",
	9:  "approve and set \"pending second approval\" on brand new flow with valid data",
	10: "approve and set \"move-in / execute modification\" on brand new flow with valid data",
	11: "set document date of brand new flow with valid data",
	12: "take action of \"complete move in\" on brand new flow with valid data",
	13: "decline at \"pending first approval\" on flow with valid data",
	14: "decline at \"pending second approval\" on flow with valid data",
}

// FlowResponse is the response of returning updated flow with status
type FlowResponse struct {
	Record  ws.RAFlowResponse `json:"record"`
	Message string            `json:"message"`
	Status  string            `json:"status"`
}

type Payload struct {
	UserRefNo         string
	RAID              int64
	Version           string
	Action            int64
	Mode              string
	Decision1         int64
	DeclineReason1    int64
	Decision2         int64
	DeclineReason2    int64
	DocumentDate      string
	TerminationReason int64
	NoticeToMoveDate  string
}

// goldDataFullFilled struct acts as gold struct with which we compare dataFullFilled of API response
var goldDataFullFilled = rlib.RADataFulfilled{
	Dates:       true,
	People:      true,
	Pets:        true,
	Vehicles:    true,
	Rentables:   true,
	ParentChild: true,
	Tie:         true,
}

var today = time.Now()
var afterFiveDays = today.AddDate(0, 0, 5)
var afterOneMonth = today.AddDate(0, 1, 0)

var documentDate = afterFiveDays.Format(rlib.RRDATEFMT4)
var notiveToMoveDate = afterOneMonth.Format(rlib.RRDATEFMT4)

var goldDocumentDate = afterFiveDays.Format(rlib.RRDATEINPFMT) + " 00:00:00 UTC"
var goldAfterOneMonth = afterOneMonth.Format(rlib.RRDATEINPFMT) + " 00:00:00 UTC"

func main() {

	var testPayloads []Payload

	var payload Payload

	// InVALID SCENARIO
	// take action of "set pending first approval" on flow with invalid data
	payload = Payload{
		UserRefNo: ExistingRAInValidFlow,
		RAID:      2,
		Version:   VersionRefNo,
		Action:    rlib.RAActionSetToFirstApproval,
		Mode:      ModeAction,
	}
	testPayloads = append(testPayloads, payload)

	// VALID SCENARIO
	// take action of "set pending first approval" on flow with valid data
	payload = Payload{
		UserRefNo: ExistingRAValidFlow,
		RAID:      1,
		Version:   VersionRefNo,
		Action:    rlib.RAActionSetToFirstApproval,
		Mode:      ModeAction,
	}
	testPayloads = append(testPayloads, payload)

	// approve "pending first approval" by accepting and set "pending second approval" on flow with valid data
	payload = Payload{
		UserRefNo:      ExistingRAValidFlow,
		RAID:           1,
		Version:        VersionRefNo,
		Mode:           ModeState,
		Decision1:      DecisionAccept,
		DeclineReason1: 0,
	}
	testPayloads = append(testPayloads, payload)

	// approve "pending second approval" by accepting and set "move-in / execute modification" on flow with valid data
	payload = Payload{
		UserRefNo:      ExistingRAValidFlow,
		RAID:           1,
		Version:        VersionRefNo,
		Mode:           ModeState,
		Decision2:      DecisionAccept,
		DeclineReason2: 0,
	}
	testPayloads = append(testPayloads, payload)

	// set document date of flow with valid data
	payload = Payload{
		UserRefNo:    ExistingRAValidFlow,
		RAID:         1,
		Version:      VersionRefNo,
		Mode:         ModeState,
		DocumentDate: documentDate,
	}
	testPayloads = append(testPayloads, payload)

	// take action of "complete move in" on flow with valid data
	payload = Payload{
		UserRefNo: ExistingRAValidFlow,
		RAID:      1,
		Version:   VersionRefNo,
		Action:    rlib.RAActionCompleteMoveIn,
		Mode:      ModeAction,
	}
	testPayloads = append(testPayloads, payload)

	// INVALID SCENARIO
	// take action of "set pending first approval" on brand new flow with invalid data
	payload = Payload{
		UserRefNo: BrandNewInValidFlow,
		RAID:      0,
		Version:   VersionRefNo,
		Action:    rlib.RAActionSetToFirstApproval,
		Mode:      ModeAction,
	}
	testPayloads = append(testPayloads, payload)

	// VALID SCENARIO
	// take action of "set pending first approval" on brand new flow with valid data
	payload = Payload{
		UserRefNo: BrandNewValidFlow,
		RAID:      0,
		Version:   VersionRefNo,
		Action:    rlib.RAActionSetToFirstApproval,
		Mode:      ModeAction,
	}
	testPayloads = append(testPayloads, payload)

	// approve "pending first approval" by accepting and set "pending second approval" on brand new flow with valid data
	payload = Payload{
		UserRefNo:      BrandNewValidFlow,
		RAID:           0,
		Version:        VersionRefNo,
		Mode:           ModeState,
		Decision1:      DecisionAccept,
		DeclineReason1: 0,
	}
	testPayloads = append(testPayloads, payload)

	// approve "pending second approval" by accepting and set "move-in / execute modification" on brand new flow with valid data
	payload = Payload{
		UserRefNo:      BrandNewValidFlow,
		RAID:           0,
		Version:        VersionRefNo,
		Mode:           ModeState,
		Decision2:      DecisionAccept,
		DeclineReason2: 0,
	}
	testPayloads = append(testPayloads, payload)

	// set document date of brand new flow with valid data
	payload = Payload{
		UserRefNo:    BrandNewValidFlow,
		RAID:         0,
		Version:      VersionRefNo,
		Mode:         ModeState,
		DocumentDate: documentDate,
	}
	testPayloads = append(testPayloads, payload)

	// take action of "complete move in" on brand new flow with valid data
	payload = Payload{
		UserRefNo: BrandNewValidFlow,
		RAID:      0,
		Version:   VersionRefNo,
		Action:    rlib.RAActionCompleteMoveIn,
		Mode:      ModeAction,
	}
	testPayloads = append(testPayloads, payload)

	// VALID SCENARIO
	// decline at "pending first approval" on flow with valid data
	payload = Payload{
		UserRefNo:      ExistingRAValidFlow1,
		RAID:           3,
		Version:        VersionRefNo,
		Mode:           ModeState,
		Decision1:      DecisionDecline,
		DeclineReason1: 75, //Criminal Background
	}
	testPayloads = append(testPayloads, payload)

	// decline at "pending second approval" on flow with valid data
	payload = Payload{
		UserRefNo:      ExistingRAValidFlow2,
		RAID:           4,
		Version:        VersionRefNo,
		Mode:           ModeState,
		Decision2:      DecisionDecline,
		DeclineReason2: 75, //Criminal Background
	}
	testPayloads = append(testPayloads, payload)

	// fmt.Println(testPayloads)
	// fmt.Println()

	for key, value := range testPayloads {
		var req *http.Request
		var respBody []byte
		var err error
		var apiResponse FlowResponse
		var raFlowData rlib.RAFlowJSONData

		testNo := key + 1

		req, err = buildRequest(value)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		respBody, err = makeRequestAndReadResponseBody(req)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		err = getDataFromResponseBody(respBody, &apiResponse, &raFlowData)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		var issues []string

		switch testNo {
		// SET Pending First Approval on INVALID EXISTING RA FLOW
		case 1:
			issues = checkTestCase1(&apiResponse, &raFlowData, value.UserRefNo)

		// Proper chain from state 0 to 4 for VALID EXISTING RA FLOW
		case 2:
			issues = checkTestCase2(&apiResponse, &raFlowData)
		case 3:
			issues = checkTestCase3(&apiResponse, &raFlowData)
		case 4:
			issues = checkTestCase4(&apiResponse, &raFlowData)
		case 5:
			issues = checkTestCase5(&apiResponse, &raFlowData)
		case 6:
			issues = checkTestCase6(&apiResponse, &raFlowData, value.UserRefNo)

		// SET Pending First Approval on INVALID BRAND NEW FLOW
		case 7:
			issues = checkTestCase1(&apiResponse, &raFlowData, value.UserRefNo)

		// Proper chain from state 0 to 4 for VALID BRAND NEW FLOW
		case 8:
			issues = checkTestCase2(&apiResponse, &raFlowData)
		case 9:
			issues = checkTestCase3(&apiResponse, &raFlowData)
		case 10:
			issues = checkTestCase4(&apiResponse, &raFlowData)
		case 11:
			issues = checkTestCase5(&apiResponse, &raFlowData)
		case 12:
			issues = checkTestCase6(&apiResponse, &raFlowData, value.UserRefNo)

		// Decline at Pending First Approval on VALID EXISTING RA FLOW
		case 13:
			issues = checkTestCase7(&apiResponse, &raFlowData)
		// Decline at Pending Second Approval on VALID EXISTING RA FLOW
		case 14:
			issues = checkTestCase8(&apiResponse, &raFlowData)

		default:
			fmt.Println("invalid testNo: ", testNo)
		}

		if len(issues) > 0 {
			fmt.Printf("\ntest%d: %s ....FAILED:\n", testNo, testNames[testNo])
			for issueNo, issueString := range issues {
				fmt.Printf("\t%d - %s\n", (issueNo + 1), issueString)
			}
		} else {
			fmt.Printf("test%d: %s ....PASSED\n", testNo, testNames[testNo])
		}
	}
}

func buildRequest(payload Payload) (*http.Request, error) {
	var req *http.Request
	var err error

	b, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("marshall payload err: %s", err)
		return req, err
	}

	url := "http://localhost:8270/v1/raactions/1/"
	// fmt.Println("\nURL: ", url)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		err = fmt.Errorf("new request err: %s", err)
		return req, err
	}
	req.Header.Set("Content-Type", "application/json")
	// fmt.Printf("\nRequest: %+v\n\n", req)

	return req, nil
}

func makeRequestAndReadResponseBody(req *http.Request) ([]byte, error) {
	var resp *http.Response
	var err error
	var respBody []byte

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		err = fmt.Errorf("client do err: %s", err)
		return respBody, err
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read response body err: %s", err)
		return respBody, err
	}

	// fmt.Printf("\nRESPONSE Body: %s\n\n", respBody)
	return respBody, nil
}

func getDataFromResponseBody(respBody []byte, apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) error {

	var err error
	err = json.Unmarshal(respBody, apiResponse)
	if err != nil {
		err = fmt.Errorf("unmarshal api response err: %s", err)
		return err
	}

	// If status of response is error, then we we simply return nil
	if apiResponse.Status == statusError {
		return nil
	}

	// get raflow data from API response into struct
	err = json.Unmarshal(apiResponse.Record.Flow.Data, raFlowData)
	if err != nil {
		err = fmt.Errorf("unmarshal flow data response err: %s", err)
		return err
	}

	// fmt.Printf("\nRESPONSE FLOW DATA: %+v\n\n", apiResponse)
	return nil
}

func checkTestCase1(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData, UserRefNo string) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	if currentState != rlib.RASTATEAppEdit {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEAppEdit])
		issues = append(issues, issueString)
	}

	var goldErrCount int

	// dataFullFilled struct for this perticular test
	// this struct acts as gold struct with which we compare dataFullFilled of API response
	var goldTestDataFullFilled rlib.RADataFulfilled

	switch UserRefNo {
	case ExistingRAInValidFlow:
		goldErrCount = 1
		goldTestDataFullFilled = rlib.RADataFulfilled{
			Dates:       true,
			People:      false,
			Pets:        true,
			Vehicles:    true,
			Rentables:   true,
			ParentChild: true,
			Tie:         false,
		}
	case BrandNewInValidFlow:
		goldErrCount = 2
		goldTestDataFullFilled = rlib.RADataFulfilled{
			Dates:       true,
			People:      false,
			Pets:        true,
			Vehicles:    true,
			Rentables:   false,
			ParentChild: true,
			Tie:         false,
		}
	}
	if validationCheck.Total != goldErrCount {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, goldErrCount)
		issues = append(issues, issueString)
	}

	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldTestDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldTestDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase2(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	if currentState != rlib.RASTATEPendingApproval1 {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEPendingApproval1])
		issues = append(issues, issueString)
	}

	if meta.ApplicationReadyUID != int64(-99999) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(-99999))
		issues = append(issues, issueString)
	}

	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase3(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	if currentState != rlib.RASTATEPendingApproval2 {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEPendingApproval2])
		issues = append(issues, issueString)
	}

	if meta.ApplicationReadyUID != int64(-99999) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(-99999))
		issues = append(issues, issueString)
	}

	if meta.Approver1 != int64(-99999) {
		issueString := fmt.Sprintf("Approver1 is: %d, should be: %d", meta.Approver1, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 4th bit of flag
	decision1 := uint64((meta.RAFLAGS >> 4) & 1)
	if decision1 != uint64(1) {
		issueString := fmt.Sprintf("Decision1 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason1 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason1 is: %d, should be: %d", meta.DeclineReason1, int64(0))
		issues = append(issues, issueString)
	}

	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase4(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	// Check State
	if currentState != rlib.RASTATEMoveIn {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEMoveIn])
		issues = append(issues, issueString)
	}

	// Check info related to state 0
	if meta.ApplicationReadyUID != int64(-99999) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(-99999))
		issues = append(issues, issueString)
	}

	// Check info related to state 1
	if meta.Approver1 != int64(-99999) {
		issueString := fmt.Sprintf("Approver1 is: %d, should be: %d", meta.Approver1, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 4th bit of flag
	decision1 := uint64((meta.RAFLAGS >> 4) & 1)
	if decision1 != uint64(1) {
		issueString := fmt.Sprintf("Decision1 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason1 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason1 is: %d, should be: %d", meta.DeclineReason1, int64(0))
		issues = append(issues, issueString)
	}

	// Check info related to state 2
	if meta.Approver2 != int64(-99999) {
		issueString := fmt.Sprintf("Approver2 is: %d, should be: %d", meta.Approver2, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 5th bit of flag
	decision2 := uint64((meta.RAFLAGS >> 5) & 1)
	if decision2 != uint64(1) {
		issueString := fmt.Sprintf("Decision2 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason2 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason2 is: %d, should be: %d", meta.DeclineReason2, int64(0))
		issues = append(issues, issueString)
	}

	// Check Validation Error count
	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	// Check data fullfilled or not
	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase5(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	// Check State
	if currentState != rlib.RASTATEMoveIn {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEMoveIn])
		issues = append(issues, issueString)
	}

	// Check info related to state 0
	if meta.ApplicationReadyUID != int64(-99999) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(-99999))
		issues = append(issues, issueString)
	}

	// Check info related to state 1
	if meta.Approver1 != int64(-99999) {
		issueString := fmt.Sprintf("Approver1 is: %d, should be: %d", meta.Approver1, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 4th bit of flag
	decision1 := uint64((meta.RAFLAGS >> 4) & 1)
	if decision1 != uint64(1) {
		issueString := fmt.Sprintf("Decision1 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason1 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason1 is: %d, should be: %d", meta.DeclineReason1, int64(0))
		issues = append(issues, issueString)
	}

	// Check info related to state 2
	if meta.Approver2 != int64(-99999) {
		issueString := fmt.Sprintf("Approver2 is: %d, should be: %d", meta.Approver2, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 5th bit of flag
	decision2 := uint64((meta.RAFLAGS >> 5) & 1)
	if decision2 != uint64(1) {
		issueString := fmt.Sprintf("Decision2 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason2 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason2 is: %d, should be: %d", meta.DeclineReason2, int64(0))
		issues = append(issues, issueString)
	}

	documentDateInMeta := time.Time(meta.DocumentDate).Format(rlib.RRDATETIMEINPFMT)

	// Check Document Date
	if documentDateInMeta != goldDocumentDate {
		issueString := fmt.Sprintf("DocumentDate is: %s, should be: %s", documentDateInMeta, goldDocumentDate)
		issues = append(issues, issueString)
	}

	// Check Validation Error count
	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	// Check data fullfilled or not
	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase6(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData, UserRefNo string) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	// Check State
	if currentState != rlib.RASTATEActive {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEActive])
		issues = append(issues, issueString)
	}

	// Check info related to state 0
	if meta.ApplicationReadyUID != int64(-99999) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(-99999))
		issues = append(issues, issueString)
	}

	// Check info related to state 1
	if meta.Approver1 != int64(-99999) {
		issueString := fmt.Sprintf("Approver1 is: %d, should be: %d", meta.Approver1, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 4th bit of flag
	decision1 := uint64((meta.RAFLAGS >> 4) & 1)
	if decision1 != uint64(1) {
		issueString := fmt.Sprintf("Decision1 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason1 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason1 is: %d, should be: %d", meta.DeclineReason1, int64(0))
		issues = append(issues, issueString)
	}

	// Check info related to state 2
	if meta.Approver2 != int64(-99999) {
		issueString := fmt.Sprintf("Approver2 is: %d, should be: %d", meta.Approver2, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 5th bit of flag
	decision2 := uint64((meta.RAFLAGS >> 5) & 1)
	if decision2 != uint64(1) {
		issueString := fmt.Sprintf("Decision2 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason2 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason2 is: %d, should be: %d", meta.DeclineReason2, int64(0))
		issues = append(issues, issueString)
	}

	documentDateInMeta := time.Time(meta.DocumentDate).Format(rlib.RRDATETIMEINPFMT)

	// Check Document Date
	if documentDateInMeta != goldDocumentDate {
		issueString := fmt.Sprintf("DocumentDate is: %s, should be: %s", documentDateInMeta, goldDocumentDate)
		issues = append(issues, issueString)
	}

	var goldNewRAID int64
	switch UserRefNo {
	case ExistingRAValidFlow:
		goldNewRAID = 5
	case BrandNewValidFlow:
		goldNewRAID = 6
	}

	// check info related to state 4
	if meta.RAID != goldNewRAID {
		issueString := fmt.Sprintf("new RAID is: %d, should be: %d", meta.RAID, goldNewRAID)
		issues = append(issues, issueString)
	}

	// Check Validation Error count
	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	// Check data fullfilled or not
	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase7(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	if currentState != rlib.RASTATETerminated {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATETerminated])
		issues = append(issues, issueString)
	}

	if meta.ApplicationReadyUID != int64(269) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(269))
		issues = append(issues, issueString)
	}

	if meta.Approver1 != int64(-99999) {
		issueString := fmt.Sprintf("Approver1 is: %d, should be: %d", meta.Approver1, int64(-99999))
		issues = append(issues, issueString)
	}

	// check decision1 from 4th bit of flag
	decision1 := uint64((meta.RAFLAGS >> 4) & 1)
	if decision1 != uint64(0) {
		issueString := fmt.Sprintf("Decision1 is: Approved, should be: Declined")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason1 != int64(75) {
		issueString := fmt.Sprintf("DeclineReason1 is: %d, should be: %d", meta.DeclineReason1, int64(75))
		issues = append(issues, issueString)
	}

	if meta.TerminatorUID != int64(-99999) {
		issueString := fmt.Sprintf("TerminatorUID is: %d, should be: %d", meta.TerminatorUID, int64(-99999))
		issues = append(issues, issueString)
	}

	if meta.LeaseTerminationReason != int64(171) {
		issueString := fmt.Sprintf("LeaseTerminationReason is: %d, should be: %d", meta.LeaseTerminationReason, int64(171))
		issues = append(issues, issueString)
	}

	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}

func checkTestCase8(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	// if server returns error than return from here
	// setting server error message as issue
	if apiResponse.Status == statusError {
		issues = append(issues, apiResponse.Message)
		return issues
	}

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	if currentState != rlib.RASTATETerminated {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATETerminated])
		issues = append(issues, issueString)
	}

	if meta.ApplicationReadyUID != int64(269) {
		issueString := fmt.Sprintf("ApplicationReadyUID is: %d, should be: %d", meta.ApplicationReadyUID, int64(269))
		issues = append(issues, issueString)
	}

	if meta.Approver1 != int64(269) {
		issueString := fmt.Sprintf("Approver1 is: %d, should be: %d", meta.Approver1, int64(269))
		issues = append(issues, issueString)
	}

	// check decision1 from 4th bit of flag
	decision1 := uint64((meta.RAFLAGS >> 4) & 1)
	if decision1 != uint64(1) {
		issueString := fmt.Sprintf("Decision1 is: Declined, should be: Approved")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason1 != int64(0) {
		issueString := fmt.Sprintf("DeclineReason1 is: %d, should be: %d", meta.DeclineReason1, int64(0))
		issues = append(issues, issueString)
	}

	// check decision2 from 5th bit of flag
	decision2 := uint64((meta.RAFLAGS >> 5) & 1)
	if decision2 != uint64(0) {
		issueString := fmt.Sprintf("Decision2 is: Approved, should be: Declined")
		issues = append(issues, issueString)
	}

	if meta.DeclineReason2 != int64(75) {
		issueString := fmt.Sprintf("DeclineReason2 is: %d, should be: %d", meta.DeclineReason2, int64(75))
		issues = append(issues, issueString)
	}

	if meta.TerminatorUID != int64(-99999) {
		issueString := fmt.Sprintf("TerminatorUID is: %d, should be: %d", meta.TerminatorUID, int64(-99999))
		issues = append(issues, issueString)
	}

	if meta.LeaseTerminationReason != int64(171) {
		issueString := fmt.Sprintf("LeaseTerminationReason is: %d, should be: %d", meta.LeaseTerminationReason, int64(171))
		issues = append(issues, issueString)
	}

	if validationCheck.Total > 0 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 0)
		issues = append(issues, issueString)
	}

	dataFullFilledCheck := false
	dataFullFilledCheck = reflect.DeepEqual(dataFullFilled, goldDataFullFilled)

	if !dataFullFilledCheck {
		issueString := fmt.Sprintf("dataFullFilled is: %+v, should be: %+v", dataFullFilled, goldDataFullFilled)
		issues = append(issues, issueString)
	}

	return issues
}
