package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mojo/util"
	"net/http"
	"rentroll/rlib"
)

// ResetPWData is the struct with the username and password
// used for authentication
type ResetPWData struct {
	Username string `json:"username"`
}

// SvcResetPW handles authentication requests from clients.
//
// wsdoc {
//  @Title ResetPW
//  @URL /v1/resetpw
//  @Method  POST
//  @Synopsis Reset a user's password
//  @Descr Given the username, this routine will regenerate the password
//  @Descr for that user and send an email to their account.  The domain
//  @Descr of the email address must be one of the supported domains:
//  @Descr accordinterests.com, l-objet.com, and myisolabella.com are the
//  @Descr supported domains at the time this routine was written. Others
//  @Descr may be added and removed over time.
//  @Input AuthenticateData
//  @Response SvcStatus
// wsdoc }
//-----------------------------------------------------------------------------
func SvcResetPW(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var funcname = "SvcResetPW"
	var a ResetPWData

	rlib.Console("Entered %s\n", funcname)

	if err := json.Unmarshal([]byte(d.data), &a); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	if len(a.Username) == 0 {
		e := fmt.Errorf("%s: Username must not be blank", funcname)
		SvcErrorReturn(w, e, funcname)
		return
	}

	b, err := json.Marshal(&a)
	if err != nil {
		e := fmt.Errorf("Error marshaling json data: %s", err.Error())
		util.Ulog("%s: %s\n", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	url := rlib.AppConfig.AuthNHost + "v1/resetpw"
	rlib.Console("posting request to: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	defer resp.Body.Close()

	rlib.Console("response Status: %s\n", resp.Status)
	rlib.Console("response Headers: %s\n", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	rlib.Console("response Body: %s\n", string(body))

	var foo rlib.AIRAuthenticateResponse
	if err := json.Unmarshal([]byte(body), &foo); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}

	switch foo.Status {
	case "success":
		rlib.Console("ResetPW succeeded\n")
		SvcWriteSuccessResponse(d.BID, w)
		return
	case "error":
		e := fmt.Errorf("%s", foo.Message)
		SvcErrorReturn(w, e, funcname)
		return
	default:
		e := fmt.Errorf("%s: Unexpected response from authentication service:  %s", funcname, foo.Status)
		SvcErrorReturn(w, e, funcname)
		return
	}
}
