package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentroll/rlib"
	"time"
)

// AuthenticateData is the struct with the username and password
// used for authentication
type AuthenticateData struct {
	User       string `json:"user"`
	Pass       string `json:"pass"`
	FLAGS      uint64 `json:"flags"`
	UserAgent  string `json:"useragent"`
	RemoteAddr string `json:"remoteaddr"`
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

	var b rlib.AIRAuthenticateResponse
	if SvcCtx.NoAuth { // test mode
		b.Status = "success"
		b.Username = "noauth"
		b.Name = "NoAuth"
		b.UID = 0
	} else {
		//-----------------------------------------------------------------------
		// fill in what the auth server needs...
		//-----------------------------------------------------------------------
		a.RemoteAddr = r.RemoteAddr // this needs to be the user's value, not our server's value
		a.UserAgent = r.UserAgent() // this needs to be the user's value, not our server's value
		fwdaddr := r.Header.Get("X-Forwarded-For")
		rlib.Console("Forwarded-For address: %q\n", fwdaddr)
		if len(fwdaddr) > 0 {
			a.RemoteAddr = fwdaddr
		}

		//-----------------------------------------------------------------------
		// Marshal together a new request buffer...
		//-----------------------------------------------------------------------
		pbr, err := json.Marshal(&a)
		if err != nil {
			e := fmt.Errorf("Error marshaling json data: %s", err.Error())
			rlib.Ulog("%s: %s\n", funcname, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
		rlib.Console("Request to auth server:  %s\n", string(pbr))

		//-----------------------------------------------------------------------
		// Send to the authenication server
		//-----------------------------------------------------------------------
		url := rlib.AppConfig.AuthNHost + "v1/authenticate"
		rlib.Console("posting request to: %s\n", url)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(pbr))
		req.Header.Set("Content-Type", "application/json")
		rlib.Console("\n*** req = %#v\n\n", req)
		client := &http.Client{}
		rlib.Console("\n*** client = %#v\n\n", client)
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
		rlib.Console("Successfully unmarshaled response: %s\n", string(body))
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
	rlib.Console("Directory Service Expire time = %s\n", time.Time(b.Expire).Format(rlib.RRDATETIMEINPFMT))
	s, err := rlib.CreateSession(r.Context(), &b)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	cookie := http.Cookie{Name: rlib.SessionCookieName, Value: b.Token, Expires: s.Expire, Path: "/"}
	http.SetCookie(w, &cookie) // a cookie cannot be set after writing anything to a response writer
	b.ImageURL = s.ImageURL
	b.Username = s.Username
	rlib.Ulog("user %s (%d) logged in\n", s.Username, s.UID)
	rlib.Console("Session Table:\n")
	rlib.DumpSessions()
	// rlib.Console("Created session: %#v\n", s)
	// rlib.Console("Created response: %#v\n", b)
	SvcWriteResponse(d.BID, &b, w)
}
