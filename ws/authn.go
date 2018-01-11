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
	User string `json:"user"`
	Pass string `json:"pass"`
}

// AuthenticateResponse is the reply structure from Accord Directory
type AuthenticateResponse struct {
	Status   string `json:"status"`   // success or error
	UID      int64  `json:"uid"`      // only present when Status = "success"
	Username string `json:"username"` // user's first or preferred name
	Name     string `json:"name"`     // user's first or preferred name
	ImageURL string `json:"imageurl"` // url to user's image
	Message  string `json:"message"`  // only present when Status = "error"
}

// SvcAuthenticate handles authentication requests from clients.
//
// wsdoc {
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
		SvcErrorReturn(w, e, funcname)
		return
	}

	var b AuthenticateResponse
	if SvcCtx.NoAuth {
		b.Status = "success"
		b.Username = "noauth"
		b.Name = "NoAuth"
		b.UID = 0
	} else {
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
			e := fmt.Errorf("%s: failed to execute client.Do:  %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		defer resp.Body.Close()

		// fmt.Println("response Status:", resp.Status)
		// fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		rlib.Console("response Body: %s\n", string(body))

		if err := json.Unmarshal([]byte(body), &b); err != nil {
			e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	}

	switch b.Status {
	case "success":
		rlib.Console("Authentication succeeded\n")
	case "error":
		e := fmt.Errorf("%s", b.Message)
		SvcErrorReturn(w, e, funcname)
		return
	default:
		e := fmt.Errorf("%s: Unexpected response from authentication service:  %s", funcname, b.Status)
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("b.Username = %s, b.UID = %d, b.Name = %s\n", b.Username, b.UID, b.Name)
	w.Header().Set("Content-Type", "application/json")
	// rlib.Console("Creating session\n")
	s, err := rlib.CreateSession(b.UID, w, r)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	b.ImageURL = s.ImageURL
	b.Username = s.Username
	rlib.Console("Created session: %#v\n", s)
	rlib.Console("Created response: %#v\n", b)
	SvcWriteResponse(&b, w)
}
