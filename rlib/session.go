package rlib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ValidateCookie describes what the auth server wants to
// validate the cookie value
type ValidateCookie struct {
	Status    string `json:"status"`
	CookieVal string `json:"cookieval"`
	IP        string `json:"ip"`
	UserAgent string `json:"useragent"`
	FLAGS     uint64 `json:"flags"`
}

// Session is the structure definition for  a user session
// with this program.
type Session struct {
	Token    string    // this is the md5 hash, unique id
	Username string    // associated username
	Name     string    // user's preferred name if it exists, otherwise the FirstName
	UID      int64     // user's db uid
	CoCode   int64     // logged in user's company (from Accord Directory)
	ImageURL string    // user's picture
	Expire   time.Time // when does the cookie expire
	RoleID   int64     // security role id
}

// UnrecognizedCookie is the error string associated with an
// unrecognized value for the airoller cookie
var UnrecognizedCookie = string("unrecognized cookie")

// ReqSessionMem is the channel int used to request write permission to the session list
var ReqSessionMem chan int

// ReqSessionMemAck is the channel int used to handshake for access to the session list
var ReqSessionMemAck chan int

// SessionCleanupTime defines the time interval between the routine that removes
// expired sessions.
var SessionCleanupTime time.Duration

// sessions is the session list managed by this code
var sessions map[string]*Session

// SessionTimeout defines how long a session can remain idle before it expires.
var SessionTimeout time.Duration // in minutes

// SessionCookieName is the name of the Roller cookie where the session
// token is stored.
var SessionCookieName = string("air")

// GetSessionCookieName simply returns a string containing the session
// cookie name. We want this to be a private / unchangeable name.
//-----------------------------------------------------------------------------
func GetSessionCookieName() string {
	return SessionCookieName
}

// SessionDispatcher is a Go routine that controls access to shared memory.
//-----------------------------------------------------------------------------
func SessionDispatcher() {
	for {
		select {
		case <-ReqSessionMem:
			ReqSessionMemAck <- 1 // tell caller go ahead
			<-ReqSessionMemAck    // block until caller is done with mem
		}
	}
}

// SessionCleanup a Go routine to periodically spin through the session list
// and remove any sessions which have timed out.
//-----------------------------------------------------------------------------
func SessionCleanup() {
	for {
		select {
		case <-time.After(SessionCleanupTime * time.Minute):
			ReqSessionMem <- 1                 // ask to access the shared mem, blocks until granted
			<-ReqSessionMemAck                 // make sure we got it
			ss := make(map[string]*Session, 0) // here's the new Session list
			n := 0                             // total number removed
			now := time.Now()                  // this is the timestamp we'll compare against
			// Console("Cleanup time: %s\n", now.Format(RRDATETIMEINPFMT))
			for k, v := range sessions { // look at every Session
				// Console("Found session: %s, expire time: %s\n", v.Name, v.Expire.Format(RRDATETIMEINPFMT))
				if now.After(v.Expire) { // if it's still active...
					n++ // removed another
				} else {
					ss[k] = v // ...copy it to the new list
				}
			}
			sessions = ss         // set the new list
			ReqSessionMemAck <- 1 // tell SessionDispatcher we're done with the data
			// Console("SessionCleanup completed. %d removed. Current Session list size = %d\n", n, len(sessions))
		}
	}
}

// SessionInit must be called prior to using the session subsystem. It
// initializes structures and starts the dispatcher
//
// INPUT
//  timeout - the number of minutes before a session times out
//
// RETURNS
//  nothing at this time
//-----------------------------------------------------------------------------
func SessionInit(timeout int) {
	sessions = make(map[string]*Session)
	ReqSessionMem = make(chan int)
	ReqSessionMemAck = make(chan int)
	SessionCleanupTime = time.Duration(1)
	SessionTimeout = time.Duration(timeout) * time.Minute
	go SessionDispatcher()
	go SessionCleanup()
}

// SessionGet returns the session associated with the supplied token, if it
// exists. It may no longer exist because it timed out.
//
// INPUT
//  token -  the index into the session table for this session. This is the
//           value that is stored in a web session cookie
//
// RETURNS
//  session - pointer to the session if the bool is true
//  bool    - true if the session was found, false otherwise
//-----------------------------------------------------------------------------
func SessionGet(token string) (*Session, bool) {
	s, ok := sessions[token]
	return s, ok
}

// ToString is the stringer for sessions
//
// RETURNS
//  a string representation of the session entry
//-----------------------------------------------------------------------------
func (s *Session) ToString() string {
	if nil == s {
		return "nil"
	}
	return fmt.Sprintf("User(%s) Name(%s) UID(%d) Token(%s)",
		s.Username, s.Name, s.UID, s.Token)
}

// DumpSessions sends the contents of the session table to the consol.
//
// RETURNS
//  a string representation of the session entry
//-----------------------------------------------------------------------------
func DumpSessions() {
	i := 0
	for _, v := range sessions {
		Console("%2d. %s\n", i, v.ToString())
		i++
	}
}

// SessionNew creates a new session and adds it to the session list
//
// INPUT
//  token    - the unique token string. This will be used to index the session
//             list
//  username - the username from the authentication service
//  name     - the name to use in the session
//  uid      - the userid associated with username. From the auth server.
//  rid      - security role id
//
// RETURNS
//  session - pointer to the new session
//-----------------------------------------------------------------------------
func SessionNew(token, username, name string, uid int64, imgurl string, rid int64, expire *time.Time) *Session {
	s := new(Session)
	s.Token = token
	s.Username = username
	s.Name = name
	s.UID = uid
	s.Expire = *expire

	switch AppConfig.AuthNType {
	case "Accord Directory":
		s.ImageURL = imgurl
	}

	ReqSessionMem <- 1 // ask to access the shared mem, blocks until granted
	<-ReqSessionMemAck // make sure we got it
	sessions[token] = s
	ReqSessionMemAck <- 1 // tell SessionDispatcher we're done with the data

	return s
}

// CreateSession creates an HTTP Cookie with the token for this session
//
// INPUT
//  w    - where to write the set cookie
//  r - the request where w should look for the cookie
//
// RETURNS
//  session - pointer to the new session
//-----------------------------------------------------------------------------
func CreateSession(ctx context.Context, a *AIRAuthenticateResponse) (*Session, error) {
	expiration := time.Time(a.Expire)

	// expiration := time.Now().Add(15 * time.Minute)

	//----------------------------------------------
	// TODO: lookup username in address book data
	//----------------------------------------------
	Console("Calling GetDirectoryPerson with UID = %d\n", a.UID)
	var c DirectoryPerson
	err := RRdb.PBsql.GetDirectoryPerson.QueryRow(a.UID).Scan(&c.UID, &c.UserName, &c.LastName, &c.MiddleName, &c.FirstName, &c.PreferredName, &c.PreferredName, &c.OfficePhone, &c.CellPhone)
	SkipSQLNoRowsError(&err)
	if err != nil {
		var bad Session
		Console("*** ERROR *** GetDirectoryPerson - %s\n", err.Error())
		return &bad, err
	}
	Console("Successfully read info from directory for UID = %d\n", c.UID)

	RoleID := int64(0) // we haven't yet implemented Role
	name := c.FirstName
	if len(c.PreferredName) > 0 {
		name = c.PreferredName
	}
	s := SessionNew(a.Token, c.UserName, name, a.UID, a.ImageURL, RoleID, &expiration)
	return s, nil
}

// IsUnrecognizedCookieError returns true if the error is UnrecognizedCookie.
//
// INPUT
//  err - the error to check
//
// RETURNS
//  bool - true means it is an UnrecognizedCookie error
//         false means it is not
//-----------------------------------------------------------------------------
func IsUnrecognizedCookieError(err error) bool {
	return strings.Contains(err.Error(), UnrecognizedCookie)
}

// ValidateSessionCookie is used to ensure that the session is still valid.
// Even if the session is found in our internal table, the 'air' is cookie used
// other applications in the suite. Someone may have logged out from a
// different app. If the cookie is not validated, then destroy the session
//
// INPUTS
//  r  - pointer to the http request, which may be updated after we add the
//       context value to it.
//  d  - our service data struct
//
// RETURNS
//  cookie - the http cookie or nil if it doesn't exist
//-----------------------------------------------------------------------------
func ValidateSessionCookie(r *http.Request) (ValidateCookie, error) {
	funcname := "ValidateSessionCookie"
	Console("Entered %s\n", funcname)
	var vc ValidateCookie
	c, err := r.Cookie(SessionCookieName)
	if err != nil {
		if strings.Contains(err.Error(), "no air cookie in request headers") {
			return vc, nil
		}
		return vc, nil
	}
	vc.CookieVal = c.Value
	vc.FLAGS = 1 << 1 // validate AND reset the expire time. This says 15 min (or whatever it is) from now is the new expire time.

	pbr, err := json.Marshal(&vc)
	if err != nil {
		return vc, fmt.Errorf("Error marshaling json data: %s", err.Error())
	}

	//-----------------------------------------------------------------------
	// Send to the authenication server
	//-----------------------------------------------------------------------
	url := AppConfig.AuthNHost + "v1/validatecookie"
	Console("posting request to: %s\n", url)
	Console("              data: %s\n", string(pbr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pbr))
	req.Header.Set("Content-Type", "application/json")
	Console("\n*** req = %#v\n\n", req)
	client := &http.Client{}
	Console("\n*** client = %#v\n\n", client)
	resp, err := client.Do(req)
	if err != nil {
		return vc, fmt.Errorf("%s: failed to execute client.Do:  %s", funcname, err.Error())
	}
	defer resp.Body.Close()

	Console("Response status = %s, status code = %d\n", resp.Status, resp.StatusCode)

	body, _ := ioutil.ReadAll(resp.Body)
	Console("*** Directory Service *** response Body: %s\n", string(body))

	if err := json.Unmarshal([]byte(body), &vc); err != nil {
		return vc, fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
	}
	Console("Successfully unmarshaled response: %s\n", string(body))
	if vc.Status != "success" {
		vc.CookieVal = ""
	}
	return vc, nil
}

// GetSession returns the session based on the cookie in the supplied
// HTTP connection.  It does NOT refresh the cookie. If you want it refreshed
// you can simply call the Refresh method on the returned pointer.
//
// INPUT
//  r - the request where we look for the cookie
//
// RETURNS
//  session - pointer to the new session
//  error   - any error encountered
//-----------------------------------------------------------------------------
func GetSession(ctx context.Context, w http.ResponseWriter, r *http.Request) (*Session, error) {
	funcname := "GetSession"
	var b AIRAuthenticateResponse
	var ok bool

	// Console("GetSession 1\n")
	// Console("\nSession Table:\n")
	// DumpSessions()
	// Console("\n")
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		// Console("GetSession 2\n")
		if strings.Contains(err.Error(), "cookie not present") {
			// Console("GetSession 3\n")
			return nil, nil
		}
		// Console("GetSession 4\n")
		return nil, err
	}
	// Console("GetSession 5\n")
	sess, ok := sessions[cookie.Value]
	if !ok || sess == nil {
		// Console("GetSession 6\n")
		//--------------------------------------------------------
		// We have a cookie, but we don't have it in our
		// session table. So, the user could have logged in from
		// another AIR app.  Validate the cookie with the AUTH
		// server.
		//--------------------------------------------------------
		var a = ValidateCookie{
			CookieVal: cookie.Value,
			IP:        r.RemoteAddr,
			UserAgent: r.UserAgent(),
		}
		//-----------------------------------------------------------------------
		// Marshal together a new request buffer...
		//-----------------------------------------------------------------------
		pbr, err := json.Marshal(&a)
		if err != nil {
			e := fmt.Errorf("Error marshaling json data: %s", err.Error())
			Ulog("%s: %s\n", funcname, err.Error())
			return nil, e
		}
		// Console("Request to auth server:  %s\n", string(pbr))

		//-----------------------------------------------------------------------
		// Send to the authenication server
		//-----------------------------------------------------------------------
		url := AppConfig.AuthNHost + "v1/validatecookie"
		Console("%s: posting request to: %s\n", funcname, url)
		Console("\tbody: %s\n", string(pbr))
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(pbr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			e := fmt.Errorf("%s: failed to execute client.Do:  %s", funcname, err.Error())
			return nil, e
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		Console("response Body: %s\n", string(body))

		if err := json.Unmarshal([]byte(body), &b); err != nil {
			e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
			return nil, e
		}
		Console("Successfully unmarshaled response: %s\n", string(body))
		//-------------------------------------
		// build a session from this data...
		//-------------------------------------
		switch b.Status {
		case "success":
			Console("Authentication succeeded\n")
		case "failure":
			Console("Cookie was not found. Could be logged off by another app.\n")
			e := fmt.Errorf("%s", b.Message)
			return nil, e
		default:
			e := fmt.Errorf("%s", b.Message)
			return nil, e
		}
		Console("Directory Service Expire time = %s\n", time.Time(b.Expire).Format(RRDATETIMEINPFMT))
		s, err := CreateSession(ctx, &b)
		if err != nil {
			return nil, err
		}
		cookie := http.Cookie{Name: SessionCookieName, Value: b.Token, Expires: s.Expire, Path: "/"}
		http.SetCookie(w, &cookie) // a cookie cannot be set after writing anything to a response writer
		return s, nil
	}
	// Console("GetSession 7\n")
	Console("sess.Username = %s\n", sess.Username)
	Console("sess.UID = %d\n", sess.UID)
	Console("sess.Token = %s\n", sess.Token)
	Console("sess.Name = %s\n", sess.Name)
	Console("sess.Expires = %s\n", sess.Expire.Format(RRDATETIMEINPFMT))
	return sess, nil
}

// Refresh updates the cookie and Session with a new expire time.
//
// INPUT
//  w - where to write the set cookie
//  r - the request where w should look for the cookie
//
// RETURNS
//  session - pointer to the new session
//-----------------------------------------------------------------------------
func (s *Session) Refresh(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie(SessionCookieName)
	if nil != cookie && err == nil {
		cookie.Expires = time.Now().Add(SessionTimeout)
		ReqSessionMem <- 1        // ask to access the shared mem, blocks until granted
		<-ReqSessionMemAck        // make sure we got it
		s.Expire = cookie.Expires // update the Session information
		ReqSessionMemAck <- 1     // tell SessionDispatcher we're done with the data
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		return 0
	}
	return 1
}

// ExpireCookie expires the cookie associated with this session now
//
// INPUT
//  w - where to write the set cookie
//  r - the request where w should look for the cookie
//
// RETURNS
//  nothing at this time
//-----------------------------------------------------------------------------
func (s *Session) ExpireCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SessionCookieName)
	if nil != cookie && err == nil {
		cookie.Expires = time.Now()
		cookie.Path = "/"
		http.SetCookie(w, cookie)
	}
}

// SessionDelete removes the supplied Session.
// if there is a better idiomatic way to do this, please let me know.
// INPUT
//  Session  - pointer to the session to tdelete
//             list
//  w        - where to write the set cookie
//  r        - the request where w should look for the cookie
//
// RETURNS
//  session  - pointer to the new session
//-----------------------------------------------------------------------------
func SessionDelete(s *Session, w http.ResponseWriter, r *http.Request) {
	if nil == s {
		Console("SessionDelete: supplied session is nil\n")
		return
	}
	Console("Session being deleted: %s\n", s.ToString())
	Console("sessions before delete:\n")
	DumpSessions()

	ss := make(map[string]*Session, 0)

	ReqSessionMem <- 1 // ask to access the shared mem, blocks until granted
	<-ReqSessionMemAck // make sure we got it
	for k, v := range sessions {
		if s.Token != k {
			ss[k] = v
		}
	}
	sessions = ss
	ReqSessionMemAck <- 1 // tell SessionDispatcher we're done with the data
	s.ExpireCookie(w, r)
	Console("sessions after delete:\n")
	DumpSessions()
}

// ErrSessionRequired session required error
var ErrSessionRequired = errors.New("Session Required, Please Login")
