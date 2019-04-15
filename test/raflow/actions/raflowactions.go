package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"rentroll/rlib"
	"rentroll/ws"
	"strconv"
	"time"

	"github.com/kardianos/osext"
)

const (
	statusError   = "error"
	statusSuccess = "success"
)

// FlowResponse defines the information sent as a reply to server reqeuests
type FlowResponse struct {
	Record  ws.RAFlowResponse `json:"record,omitempty"`
	Message string            `json:"message,omitempty"`
	Status  string            `json:"status"`
}

// JSONPayloadDesc is a struct that stores request payload and short description about the test
type JSONPayloadDesc struct {
	ReqData     Payload
	Description string
}

// Payload contains information regarding status of the transition through the
// states of a Rental Agreement
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
var noticeToMoveDate = afterOneMonth.Format(rlib.RRDATEFMT4)

var folderName string

func readCommandLineArgs() {
	pFolder := flag.String("f", "", "test directory")

	flag.Parse()

	folderName = *pFolder
}

func getPayloadsFromJSON(payloads *[]JSONPayloadDesc) error {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	// read json file which contains payloads and description about tests
	payloadFilePath := path.Join(folderPath, folderName, "payload.json")

	jsonData, err := ioutil.ReadFile(payloadFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, payloads)
	if err != nil {
		return err
	}
	return err
}

func main() {
	var err error
	var testPayloads []JSONPayloadDesc

	readCommandLineArgs()

	err = getPayloadsFromJSON(&testPayloads)
	if err != nil {
		fmt.Println("Internal Error: ", err)
		return
	}

	// rlib.Console("A\n")

	for key, payload := range testPayloads {
		var req *http.Request
		var respBody []byte
		var respRecord []byte

		var apiResponse FlowResponse

		testNo := key + 1

		// rlib.Console("B\n")
		// if payload contains DocumentDate as "99/99/9999"
		// that means that we need to set DocumentDate
		if payload.ReqData.DocumentDate == "99/99/9999" {
			payload.ReqData.DocumentDate = documentDate
		}

		// if payload contains NoticeToMoveDate as "88/88/8888"
		// that means that we need to set NoticeToMoveDate
		if payload.ReqData.NoticeToMoveDate == "88/88/8888" {
			payload.ReqData.NoticeToMoveDate = noticeToMoveDate
		}

		// rlib.Console("C\n")
		url := "http://localhost:8270/v1/raactions/1/"

		req, err = buildRequest(url, payload.ReqData)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}
		// rlib.Console("C1\n")

		// rlib.Console("req = %#v\n", req)
		respBody, err = makeRequestAndReadResponseBody(req)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}
		// rlib.Console("C2\n")

		err = getDataFromResponseBody(respBody, &apiResponse)
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}
		// rlib.Console("D\n")

		var respRAID int64
		var respUserRefNo string

		if apiResponse.Status == statusSuccess {

			err = json.Unmarshal(respBody, &apiResponse)
			if err != nil {
				fmt.Println("Internal Error: ", err)
				return
			}
			// rlib.Console("E\n")

			// here we override UserRefNo
			// because in raid_version tests userRefNo generated will be random everytime
			// hence to pass the comparison with gold files we set UserRefNo as "OVERRIDEEN1234567890"
			if folderName == "raid_version" && apiResponse.Record.Flow.UserRefNo != "" {
				apiResponse.Record.Flow.UserRefNo = "OVERRIDEEN1234567890"
			}
			respRAID = apiResponse.Record.Flow.ID
			respUserRefNo = apiResponse.Record.Flow.UserRefNo
			// rlib.Console("F\n")

		}

		// rlib.Console("G\n")

		respRecord, err = json.MarshalIndent(apiResponse, "", "    ")
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		// rlib.Console("H\n")

		testInfoString := fmt.Sprintf("Test %d: %s \n", testNo, payload.Description)
		testInfoString += fmt.Sprintf("Request( RAID: %d, UserRefNo: %s )\n", payload.ReqData.RAID, payload.ReqData.UserRefNo)
		testInfoString += fmt.Sprintf("Response( RAID: %d, UserRefNo: %s )\n", respRAID, respUserRefNo)
		dumpResponseInFile(testNo, testInfoString, respRecord)
		// rlib.Console("I\n")
	}
	// rlib.Console("J\n")
}

func buildRequest(url string, payload Payload) (*http.Request, error) {
	var req *http.Request
	var err error

	// rlib.Console("url = %s\n", url)
	// rlib.Console("payload = %#v\n", payload)
	b, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("marshall payload err: %s", err)
		return req, err
	}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		err = fmt.Errorf("new request err: %s", err)
		return req, err
	}
	req.Header.Set("Content-Type", "application/json")
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
	return respBody, nil
}

func getDataFromResponseBody(respBody []byte, apiResponse *FlowResponse) error {
	var err error
	err = json.Unmarshal(respBody, apiResponse)
	if err != nil {
		err = fmt.Errorf("unmarshal api response err: %s", err)
		return err
	}
	return nil
}

func dumpResponseInFile(testNo int, testInfoString string, respRecord []byte) error {
	var err error

	testInfoString += string(respRecord)

	data := []byte(testInfoString)

	filename := "a" + strconv.Itoa(testNo)

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}

	// read json file which contains payloads
	filePath := path.Join(folderPath, folderName, filename)

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		err = fmt.Errorf("write file err: %s", err)
		return err
	}
	return nil
}
