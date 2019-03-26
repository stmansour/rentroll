package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rentroll/rlib"
)

// ValidateCookie describes the data sent by an AIR app to check
// whether or not a cookie value is valid.
type ValidateCookie struct {
	CookieVal string `json:"cookieval"`
	FLAGS     uint64 `json:"flags"`
}

// SvcLogoff handles authentication requests from clients.
//
//  @Title Logoff
//  @URL /v1/logoff
//  @Method  POST
//  @Synopsis Logoff a user
//  @Descr It removes the user's session from the session table if it exists
//  @Input n/a
//  @Response SvcStatus
// wsdoc }
//-----------------------------------------------------------------------------
func SvcLogoff(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var funcname = "SvcLogoff"
	rlib.Console("Entered %s\n", funcname)
	if d.sess == nil {
		rlib.Console("%s:  d.sess is nil\n", funcname)
		err := fmt.Errorf("%s: cannot delete nil session", funcname)
		SvcErrorReturn(w, err, funcname)
		return
	}

	//-----------------------------------------------------------------
	// If we get this far, it means that we do have a session (d.sess)
	// Just delete the session.  This will also expire the cookie
	//-----------------------------------------------------------------
	if nil != d.sess {
		rlib.SessionDelete(d.sess, w, r)
	}

	// The logoff command uses the same data struct as ValidateCooki
	var a = ValidateCookie{
		CookieVal: d.sess.Token, // this is the cookie val we want to delete
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
	url := rlib.AppConfig.AuthNHost + "v1/logoff"
	rlib.Console("posting request to: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pbr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e := fmt.Errorf("%s: failed to execute client.Do:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	rlib.Console("response Body: %s\n", string(body))

	var b SvcStatus
	if err := json.Unmarshal([]byte(body), &b); err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcErrorReturn(w, e, funcname)
		return
	}
	rlib.Console("Status response: %s\n", b.Status)
	SvcWriteSuccessResponse(d.BID, w)
	rlib.Ulog("user %s logged off\n", d.sess.Username)
}
