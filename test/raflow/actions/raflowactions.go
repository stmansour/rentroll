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

type FlowResponse struct {
	Record  ws.RAFlowResponse `json:"record,omitempty"`
	Message string            `json:"message,omitempty"`
	Status  string            `json:"status"`
}

type JSONPayloadDesc struct {
	ReqData     Payload
	Description string
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

	// read json file which contains payloads
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

	// fmt.Println(testPayloads)
	// fmt.Println()

	for key, payload := range testPayloads {
		var req *http.Request
		var respBody []byte
		var respRecord []byte

		var apiResponse FlowResponse

		testNo := key + 1

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

		url := "http://localhost:8270/v1/raactions/1/"

		req, err = buildRequest(url, payload.ReqData)
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

			// here we override UserRefNo
			// because in raid_version tests userRefNo generated will be random everytime
			// hence to pass the comparision with gold files we set UserRefNo as "OVERRIDEEN1234567890"
			if folderName == "raid_version" && apiResponse.Record.Flow.UserRefNo != "" {

				/*flowId := apiResponse.Record.Flow.FlowID
				newUrl := fmt.Sprintf("http://localhost:8270/v1/flow/1/%d/", flowId)

				var newReq *http.Request
				var newRespBody []byte
				var newApiResponse FlowResponse

				var flowJSONData rlib.RAFlowJSONData

				err = json.Unmarshal(apiResponse.Record.Flow.Data, &flowJSONData)
				if err != nil {
					err = fmt.Errorf("unmarshal api response err: %s", err)
					fmt.Println("Internal Error: ", err)
					return
				}

				temp := struct {
					CSAgent         int64
					RentStop        string
					RentStart       string
					AgreementStop   string
					AgreementStart  string
					PossessionStop  string
					PossessionStart string
				}{
					flowJSONData.Dates.CSAgent,
					"3/1/2020",
					"3/13/2018",
					"3/1/2020",
					"3/13/2018",
					"3/1/2020",
					"3/13/2018",
				}

				newPayload := struct {
					Cmd         string
					FlowType    string
					FlowId      int64
					FlowPartKey string
					BID         int64
					Data        struct {
						CSAgent         int64
						RentStop        string
						RentStart       string
						AgreementStop   string
						AgreementStart  string
						PossessionStop  string
						PossessionStart string
					}
				}{
					"save",
					"RA",
					flowId,
					"dates",
					int64(1),
					temp,
				}

				var b []byte
				b, err = json.Marshal(newPayload)
				if err != nil {
					err = fmt.Errorf("marshall payload err: %s", err)
					fmt.Println("Internal Error: ", err)
					return
				}

				newReq, err = http.NewRequest("POST", newUrl, bytes.NewBuffer(b))
				if err != nil {
					err = fmt.Errorf("new request err: %s", err)
					fmt.Println("Internal Error: ", err)
					return
				}
				newReq.Header.Set("Content-Type", "application/json")

				newRespBody, err = makeRequestAndReadResponseBody(newReq)
				if err != nil {
					fmt.Println("Internal Error: ", err)
					return
				}

				err = getDataFromResponseBody(newRespBody, &newApiResponse)
				if err != nil {
					fmt.Println("Internal Error: ", err)
					return
				}
				apiResponse.Record.Flow = newApiResponse.Record.Flow*/

				apiResponse.Record.Flow.UserRefNo = "OVERRIDEEN1234567890"
			}
			respRAID = apiResponse.Record.Flow.ID
			respUserRefNo = apiResponse.Record.Flow.UserRefNo

		}

		respRecord, err = json.MarshalIndent(apiResponse, "", "    ")
		if err != nil {
			fmt.Println("Internal Error: ", err)
			return
		}

		testInfoString := fmt.Sprintf("Test %d: %s \n", testNo, payload.Description)
		testInfoString += fmt.Sprintf("Request( RAID: %d, UserRefNo: %s )\n", payload.ReqData.RAID, payload.ReqData.UserRefNo)
		testInfoString += fmt.Sprintf("Response( RAID: %d, UserRefNo: %s )\n", respRAID, respUserRefNo)
		dumpResponseInFile(testNo, testInfoString, respRecord)
	}
}

func buildRequest(url string, payload Payload) (*http.Request, error) {
	var req *http.Request
	var err error

	b, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("marshall payload err: %s", err)
		return req, err
	}

	// url := "http://localhost:8270/v1/raactions/1/"
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
