package rlib

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

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
	// Urole    Role               // user's role for permissions
	// Breadcrumbs  []Crumb        // where is the user in the screen hierarchy
	// Pp           map[string]int // quick way to reference person permissions based on field name
	// Pco          map[string]int // quick way to reference company permissions based on field name
	// Pcl          map[string]int // quick way to reference class permissions based on field name
	// Ppr          map[string]int
	// UIDorig      int            // original uid (for use with method sessionBecome())
	// UsernameOrig string         // original username
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

// sessionCookieName is the name of the Roller cookie where the session
// token is stored.
var sessionCookieName = string("airoller")

// GetSessionCookieName simply returns a string containing the session
// cookie name. We want this to be a private / unchangable name.
//-----------------------------------------------------------------------------
func GetSessionCookieName() string {
	return sessionCookieName
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

func dumpSessions() {
	i := 0
	for _, v := range sessions {
		fmt.Printf("%2d. %s\n", i, v.ToString())
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
func SessionNew(token, username, name string, uid int64, rid int64) *Session {
	s := new(Session)
	s.Token = token
	s.Username = username
	s.Name = name
	s.UID = uid

	switch AppConfig.AuthNType {
	case "Accord Directory":
		s.ImageURL = fmt.Sprintf("%spictures/%d.png", AppConfig.AuthNHost, s.UID)
	}

	s.Expire = time.Now().Add(SessionTimeout)

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
func CreateSession(uid int64, w http.ResponseWriter, r *http.Request) (*Session, error) {
	expiration := time.Now().Add(SessionTimeout)

	//----------------------------------------------
	// TODO: lookup username in address book data
	//----------------------------------------------
	dp, err := GetDirectoryPerson(r.Context(), uid)
	if err != nil {
		var bad Session
		return &bad, err
	}
	Console("DIR PERSON UserName = %s\n", dp.UserName)
	Console("dp = %#v\n", dp)
	RoleID := int64(0)

	//=================================================================================
	// There could be multiple sessions from the same user on different browsers.
	// These could be on the same or separate machines. We need the IP and the browser
	// to ensure uniqueness...
	//
	// The cookie is:   username + User-Agent + remote ip address
	//=================================================================================
	key := dp.UserName + r.Header.Get("User-Agent") + r.RemoteAddr
	token := fmt.Sprintf("%x", md5.Sum([]byte(key)))
	name := dp.FirstName
	if len(dp.PreferredName) > 0 {
		name = dp.PreferredName
	}
	Console("dp = %#v\n", dp)
	s := SessionNew(token, dp.UserName, name, uid, RoleID)
	Console("session = %#v\n", s)
	cookie := http.Cookie{Name: sessionCookieName, Value: s.Token, Expires: expiration}
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
	r.AddCookie(&cookie) // need this so that the redirect to search finds the cookie
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

// GetSession returns the session based on the cookie in the supplied
// HTTP connection.  It does NOT refresh the cookie. If you want it refreshed
// you can simply call the Refresh method on the returned pointer.
//
// INPUT
//  r - the request where w should look for the cookie
//
// RETURNS
//  session - pointer to the new session
//  error   - any error encountered
//-----------------------------------------------------------------------------
func GetSession(w http.ResponseWriter, r *http.Request) (*Session, error) {
	var ok bool
	Console("GetSession 1\n")
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		Console("GetSession 2\n")
		if strings.Contains(err.Error(), "cookie not present") {
			Console("GetSession 3\n")
			return nil, nil
		}
		Console("GetSession 4\n")
		return nil, err
	}
	Console("GetSession 5\n")
	sess, ok := sessions[cookie.Value]
	if !ok || sess == nil {
		Console("GetSession 6\n")
		cookie.Expires = time.Now()
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		err := fmt.Errorf("There is a problem with your session: %s.  Please Sign In again", UnrecognizedCookie)
		return nil, err // cookie had a value, but not found in our session table
	}
	Console("GetSession 7\n")
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
	cookie, err := r.Cookie(sessionCookieName)
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
	cookie, err := r.Cookie(sessionCookieName)
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
//  session - pointer to the new session
//-----------------------------------------------------------------------------
func SessionDelete(s *Session, w http.ResponseWriter, r *http.Request) {
	// Console("Session being deleted: %s\n", s.ToString())
	// Console("sessions before delete:\n")
	// dumpSessions()

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
	// Console("sessions after delete:\n")
	// dumpSessions()
}

// ErrSessionRequired session required error
var ErrSessionRequired = errors.New("Session Required, Please Login")
