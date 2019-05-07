package main

// This program was written by one of the contractors at Aubergine Solutions.
// I have gone through and commented it (as well as re-organized it and in some
// places refactored it) as I am debugging an issue and need to learn how this
// program works.
//
// -sman
// 5/7/2019

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
//-----------------------------------------------------------------------------
type FlowResponse struct {
	Record  ws.RAFlowResponse `json:"record,omitempty"`
	Message string            `json:"message,omitempty"`
	Status  string            `json:"status"`
}

// JSONPayloadDesc is a struct that stores request payload and short
// description about the test
//-----------------------------------------------------------------------------
type JSONPayloadDesc struct {
	ReqData     Payload
	Description string
}

// Payload contains information regarding status of the transition through the
// states of a Rental Agreement.
// THESE STRUCTURES ARE DEFINED IN payload.json FILES IN EACH TEST DIRECTORY
//-----------------------------------------------------------------------------
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
	TerminationDate   string
	NoticeToMoveDate  string
}

var now = time.Now()
var nowPlus5days = now.AddDate(0, 0, 5)
var nowPlus1month = now.AddDate(0, 1, 0)
var dtNowPlus5days = nowPlus5days.Format(rlib.RRDATEFMT4)
var dtNowPlus1month = nowPlus1month.Format(rlib.RRDATEFMT4)
var dtTerminate = time.Date(2019, time.May, 7, 14, 12, 0, 0, time.UTC)
var dtTerminateDate = dtTerminate.Format(rlib.RRDATEFMT4)
var folderName string

// readCommandLineArgs
//-----------------------------------------------------------------------------
func readCommandLineArgs() {
	pFolder := flag.String("f", "", "test directory")

	flag.Parse()

	folderName = *pFolder
}

// getPayloadsFromJSON builds an array of Payload structs from the data
// in payload.json found in the directory below this executable that was
// supplied in the cmdline -f option.
//
// INPUTS
//   payloads - array of JSONPayloadDesc to fill
//
// RETURNS
//   any error encountered.
//-----------------------------------------------------------------------------
func getPayloadsFromJSON(payloads *[]JSONPayloadDesc) error {
	var err error
	var folderPath, payloadFilePath string
	var jsonData []byte

	if folderPath, err = osext.ExecutableFolder(); err != nil {
		return err
	}

	//---------------------------------------------------------------------
	// The payload.json file contains payloads and descriptions about the
	// tests to run...
	//---------------------------------------------------------------------
	payloadFilePath = path.Join(folderPath, folderName, "payload.json")
	// rlib.Console("getPayloadsFromJSON: %s\n", payloadFilePath)
	if jsonData, err = ioutil.ReadFile(payloadFilePath); err != nil {
		return err
	}

	return json.Unmarshal(jsonData, payloads)
}

// main
//--------------------------------------------------------------------------
func main() {
	var err error
	var testPayloads []JSONPayloadDesc

	//------------------------------------------------------------
	// Get the directory of the payloads file, from -f <dir>
	//------------------------------------------------------------
	readCommandLineArgs()

	//------------------------------------------------------------
	// Read the commands we need to execute
	//------------------------------------------------------------
	if err = getPayloadsFromJSON(&testPayloads); err != nil {
		fmt.Printf("Error from getPayloadsFromJSON: %s\n", err)
		os.Exit(1)
	}

	//------------------------------------------------------------
	// Read the commands we need to execute
	//------------------------------------------------------------
	for key, payload := range testPayloads {
		var err error
		var req *http.Request
		var respBody []byte
		var respRecord []byte
		var apiResponse FlowResponse
		testNo := key + 1

		// rlib.Console("payload[%d]: RAID = %d, Action = %d, RefNo = %s\n", testNo, payload.ReqData.RAID, payload.ReqData.Action, payload.ReqData.UserRefNo)

		//------------------------------------------------------------
		// if payload contains DocumentDate as "99/99/9999"
		// that means that we need to set DocumentDate
		//------------------------------------------------------------
		if payload.ReqData.DocumentDate == "99/99/9999" {
			payload.ReqData.DocumentDate = dtNowPlus5days
		}

		//------------------------------------------------------------
		// if payload contains NoticeToMoveDate as "88/88/8888"
		// that means that we need to set NoticeToMoveDate
		//------------------------------------------------------------
		if payload.ReqData.NoticeToMoveDate == "88/88/8888" {
			payload.ReqData.NoticeToMoveDate = dtNowPlus1month
		}

		//------------------------------------------------------------
		// if the command is terminate, set terminate date to May 7, 2019
		// (the day this issue was discovered)
		//------------------------------------------------------------
		if payload.ReqData.Action == rlib.RAActionTerminate {
			payload.ReqData.TerminationDate = dtTerminateDate
		}

		//------------------------------------------------------------
		// Create the request, send it, and read the response
		//------------------------------------------------------------
		url := "http://localhost:8270/v1/raactions/1/"
		if req, err = buildRequest(url, payload.ReqData); err != nil {
			fmt.Println("Error from buildRequest: ", err)
			os.Exit(1)
		}
		if respBody, err = makeRequestAndReadResponseBody(req); err != nil {
			fmt.Println("Error: from makeRequestAndReadResponseBody", err)
			os.Exit(1)
		}
		if err = json.Unmarshal(respBody, &apiResponse); err != nil {
			fmt.Println("Error unmarshaling response: ", err)
			os.Exit(1)
		}

		var respRAID int64
		var respUserRefNo string
		if apiResponse.Status == statusSuccess {
			//------------------------------------------------------------------
			// UserRefNo gets generated each time the test is run. It is a guid.
			// It is overwritten here presumably to ensure that the gold-files
			// pass the comparison check.
			//------------------------------------------------------------------
			if folderName == "raid_version" && apiResponse.Record.Flow.UserRefNo != "" {
				apiResponse.Record.Flow.UserRefNo = "OVERRIDE1234567890"
			}
			respRAID = apiResponse.Record.Flow.ID
			respUserRefNo = apiResponse.Record.Flow.UserRefNo
		}

		//-------------------------------------------------------------
		// format the server response into something readable...
		//-------------------------------------------------------------
		if respRecord, err = json.MarshalIndent(apiResponse, "", "    "); err != nil {
			fmt.Println("Internal Error: ", err)
			os.Exit(1)
		}

		testInfoString := fmt.Sprintf("Test %d: %s \n", testNo, payload.Description)
		testInfoString += fmt.Sprintf("Request( RAID: %d, UserRefNo: %s )\n", payload.ReqData.RAID, payload.ReqData.UserRefNo)
		testInfoString += fmt.Sprintf("Response( RAID: %d, UserRefNo: %s )\n", respRAID, respUserRefNo)
		if err = dumpResponseInFile(testNo, testInfoString, respRecord); err != nil {
			fmt.Printf("Error writing response file: %s\n", err)
			os.Exit(1)
		}
	}
}

// buildRequest creates and returns an http.Request pointer for the request
// defined by the payload param.
//
// INPUTS
//       url - where to send the request
//   payload - definition of what to send
//
// RETURNS
//   ptr to the http.Request object
//   any error encountered
//-----------------------------------------------------------------------------
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

// makeRequestAndReadResponseBody calls the server with the request and
// reads back the server reply
//
// INPUTS
//   req - the server request
//
// RETURNS
//   the server reply
//   any errors encountered
//----------------------------------------------------------------------------
func makeRequestAndReadResponseBody(req *http.Request) ([]byte, error) {
	var resp *http.Response
	var err error
	var respBody []byte

	client := &http.Client{}
	if resp, err = client.Do(req); err != nil {
		err = fmt.Errorf("client do err: %s", err)
		return respBody, err
	}
	defer resp.Body.Close()

	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("read response body err: %s", err)
		return respBody, err
	}
	return respBody, nil
}

// dumpResponseInFile creates the text file containing the server's reply to
// a request.
//
// INPUTS
//   testNo         - test number, an index number
//   testInfoString - summary information about the request and response
//   respRecord     - http response
//
// RETURNS
//   any errors encountered
//----------------------------------------------------------------------------
func dumpResponseInFile(testNo int, testInfoString string, respRecord []byte) error {
	var err error
	var folderPath string

	testInfoString += string(respRecord)
	data := []byte(testInfoString)
	filename := "a" + strconv.Itoa(testNo)
	if folderPath, err = osext.ExecutableFolder(); err != nil {
		return err
	}
	filePath := path.Join(folderPath, folderName, filename)
	return ioutil.WriteFile(filePath, data, 0644)
}
