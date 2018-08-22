package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentroll/rlib"
	"rentroll/ws"
	"strconv"
	"time"
)

const (
	statusError   = "error"
	statusSuccess = "success"

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

type FlowResponse struct {
	Record  ws.RAFlowResponse `json:"record,omitempty"`
	Message string            `json:"message,omitempty"`
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

var today = time.Now()
var afterFiveDays = today.AddDate(0, 0, 5)
var afterOneMonth = today.AddDate(0, 1, 0)

var documentDate = afterFiveDays.Format(rlib.RRDATEFMT4)
var notiveToMoveDate = afterOneMonth.Format(rlib.RRDATEFMT4)

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
		var respRecord []byte
		var err error
		var apiResponse FlowResponse

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

		err = getDataFromResponseBody(respBody, &apiResponse)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		var respRAID int64
		var respUserRefNo string

		if apiResponse.Status == statusSuccess {

			err = json.Unmarshal(respBody, &apiResponse)
			if err != nil {
				fmt.Println("Internal Error: ", err)
				return
			}
			respRAID = apiResponse.Record.Flow.ID
			respUserRefNo = apiResponse.Record.Flow.UserRefNo

		}

		respRecord, err = json.MarshalIndent(apiResponse, "", "    ")
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		testInfoString := fmt.Sprintf("Test %d: %s \n", testNo, testNames[testNo])
		testInfoString += fmt.Sprintf("Request( RAID: %d, UserRefNo: %s )\n", value.RAID, value.UserRefNo)
		testInfoString += fmt.Sprintf("Response( RAID: %d, UserRefNo: %s )\n", respRAID, respUserRefNo)
		dumpResponseInFile(testNo, testInfoString, respRecord)
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

func getDataFromResponseBody(respBody []byte, apiResponse *FlowResponse) error {

	var err error
	err = json.Unmarshal(respBody, apiResponse)
	if err != nil {
		err = fmt.Errorf("unmarshal api response err: %s", err)
		return err
	}

	// fmt.Printf("\nRESPONSE FLOW DATA: %+v\n\n", apiResponse)
	return nil
}

func dumpResponseInFile(testNo int, testInfoString string, respRecord []byte) error {
	var err error

	testInfoString += string(respRecord)

	data := []byte(testInfoString)
	fileName := "a" + strconv.Itoa(testNo)
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file err: %s", err)
		return err
	}
	return nil
}
