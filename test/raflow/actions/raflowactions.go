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
)

// FlowResponse is the response of returning updated flow with status
type FlowResponse struct {
	Record ws.RAFlowResponse `json:"record"`
	Status string            `json:"status"`
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

func main() {

	var testPayloads []Payload

	// take action of "set pending first approval" on flow with invalid data
	payload := Payload{
		UserRefNo: "VJFC558GW9MM625CT176",
		RAID:      2,
		Version:   "refno",
		Action:    1,
		Mode:      "Action",
	}
	testPayloads = append(testPayloads, payload)

	// take action of "set pending first approval" on flow with valid data
	payload = Payload{
		UserRefNo: "FU1T222ATL6HWFS61388",
		RAID:      1,
		Version:   "refno",
		Action:    1,
		Mode:      "Action",
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
			err = fmt.Errorf("test%d failed: %s", testNo, err)
		}

		respBody, err = makeRequestAndReadResponseBody(req)
		if err != nil {
			err = fmt.Errorf("test%d failed: %s", testNo, err)
		}

		err = getDateFromResponseBody(respBody, &apiResponse, &raFlowData)
		if err != nil {
			err = fmt.Errorf("test%d failed: %s", testNo, err)
		}

		var issues []string

		switch testNo {
		case 1:
			issues = checkTestCase1(&apiResponse, &raFlowData)
		case 2:
			issues = checkTestCase2(&apiResponse, &raFlowData)
		default:
			fmt.Println("invalid testNo: ", testNo)
		}

		if len(issues) > 0 {
			fmt.Printf("test%d failed:\n", testNo)
			for issueNo, issueString := range issues {
				fmt.Printf("\t%d - %s\n", (issueNo + 1), issueString)
			}
		} else {
			fmt.Printf("test%d passed\n", testNo)
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

	// fmt.Printf("\nRESPONSE Body: %s\n\n", readd)
	return respBody, nil
}

func getDateFromResponseBody(respBody []byte, apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) error {

	var err error
	err = json.Unmarshal(respBody, apiResponse)
	if err != nil {
		err = fmt.Errorf("unmarshal api response err: %s", err)
		return err
	}

	// get raflow data from API response into struct
	err = json.Unmarshal(apiResponse.Record.Flow.Data, raFlowData)
	if err != nil {
		err = fmt.Errorf("unmarshal flow data response err: %s", err)
		return err
	}

	// fmt.Printf("\nRESPONSE FLOW DATA: %+v\n\n", raFlowData.Meta)
	return nil
}

func checkTestCase1(apiResponse *FlowResponse, raFlowData *rlib.RAFlowJSONData) []string {
	var issues []string

	meta := raFlowData.Meta
	validationCheck := apiResponse.Record.ValidationCheck
	dataFullFilled := apiResponse.Record.DataFulfilled

	currentState := meta.RAFLAGS & uint64(0xF)

	if currentState != rlib.RASTATEAppEdit {
		issueString := fmt.Sprintf("state is: %s, should be: %s", rlib.RAStates[currentState], rlib.RAStates[rlib.RASTATEAppEdit])
		issues = append(issues, issueString)
	}

	if validationCheck.Total != 1 {
		issueString := fmt.Sprintf("error count is: %d, should be: %d", validationCheck.Total, 1)
		issues = append(issues, issueString)
	}

	// dataFullFilled struct for this perticular test
	// this struct acts as gold struct with which we compare dataFullFilled of API response
	var goldTestDataFullFilled = rlib.RADataFulfilled{
		Dates:       true,
		People:      false,
		Pets:        true,
		Vehicles:    true,
		Rentables:   true,
		ParentChild: true,
		Tie:         false,
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
