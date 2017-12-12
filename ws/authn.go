package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentroll/rlib"
)

// AuthenticateData is the struct with the username and password
// used for authentication
type AuthenticateData struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	ResetPwd bool   `json:"resetPwd"`
}

// AuthenticateResponse is the reply structure from Accord Directory
type AuthenticateResponse struct {
	Status  string `json:"status"`  // success or error
	UID     int64  `json:"uid"`     // only present when Status = "success"
	Message string `json:"message"` // only present when Status = "error"
}

// SvcAuthenticate handles authentication requests from clients.
//
//  @Title Authenticate
//  @URL /v1/authn
//  @Method  POST
//  @Synopsis Authenticate a user
//  @Descr It contacts Accord Directory server to authenticate users. If successful,
//  @Descr it creates a session for the user and sends a response with Status
//  @Descr set to "success".  If it is not successful, it sends  response
//  @Descr with Status set to "error" and provides the reason as given by
//  @Descr the Accord Directory server.
//  @Input AuthenticateData
//  @Response SvcStatus
// wsdoc }
//-----------------------------------------------------------------------------
func SvcAuthenticate(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var funcname = "SvcAuthenticate"
	var a AuthenticateData

	rlib.Console("Entered %s\n", funcname)
	// rlib.Console("record data = %s\n", d.data)  // this has user's password, so try not to print (you may forget to remove it)

	if err := json.Unmarshal([]byte(d.data), &a); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	//rlib.Console("User = %s, Pass = %s\n", a.User, a.Pass)

	//-----------------------------------------------------------------------
	// TODO: Implmentate to handle reset/forgot password
	// if `resetPwd` is true, then user has requested to reset password
	// OR maybe user forgots their password
	//-----------------------------------------------------------------------
	if a.ResetPwd {
		e := fmt.Errorf("%s: feature is not available at the moment", funcname)
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	//-----------------------------------------------------------------------
	// There's no need to Marshal the data into JSON format. We already have
	// it in d.data.  Just pass it along to the authenication server
	//-----------------------------------------------------------------------
	url := rlib.AppConfig.AuthNHost + "v1/authenticate"
	rlib.Console("posting request to: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(d.data)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	rlib.Console("response Body: %s\n", string(body))

	var b AuthenticateResponse
	if err := json.Unmarshal([]byte(body), &b); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e, funcname)
		return
	}

	switch b.Status {
	case "success":
		rlib.Console("Authentication succeeded\n")
	case "error":
		e := fmt.Errorf("%s", b.Message)
		SvcGridErrorReturn(w, e, funcname)
		return
	default:
		e := fmt.Errorf("%s: Unexpected response from authentication service:  %s", funcname, b.Status)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	rlib.Console("Creating session\n")
	s, err := rlib.CreateSession(a.User, w, r)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}
	rlib.Console("Created session: %s\n", s.Token)
	SvcWriteResponse(&b, w)
}
