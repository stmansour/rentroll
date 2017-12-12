package rlib

import (
	"crypto/md5"
	"fmt"
	"net/http"
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

// SessionCleanup periodically spins through the list of sessions
// and removes any which have timed out.
//-----------------------------------------------------------------------------
func SessionCleanup() {
	for {
		select {
		case <-time.After(SessionCleanupTime * time.Minute):
			ReqSessionMem <- 1                 // ask to access the shared mem, blocks until granted
			<-ReqSessionMemAck                 // make sure we got it
			ss := make(map[string]*Session, 0) // here's the new Session list
			n := 0                             // total number removed
			for k, v := range sessions {       // look at every Session
				if time.Now().After(v.Expire) { // if it's still active...
					n++ // removed another
				} else {
					ss[k] = v // ...copy it to the new list
				}
			}
			sessions = ss         // set the new list
			ReqSessionMemAck <- 1 // tell SessionDispatcher we're done with the data
			//fmt.Printf("SessionCleanup completed. %d removed. Current Session list size = %d\n", n, len(sessions))
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
	// s.UIDorig = uid
	// s.ImageURL = getImageFilename(uid)
	// s.Breadcrumbs = make([]Crumb, 0)
	// getRoleInfo(rid, s)

	// var d personDetail
	// d.UID = uid
	//getSecurityList(&d)

	// err := .db.QueryRow(fmt.Sprintf("select cocode from people where uid=%d", uid)).Scan(&s.CoCode)
	// if nil != err {
	// 	ulog("Unable to read CoCode for userid=%d,  err = %v\n", uid, err)
	// }

	ReqSessionMem <- 1 // ask to access the shared mem, blocks until granted
	<-ReqSessionMemAck // make sure we got it
	sessions[token] = s
	ReqSessionMemAck <- 1 // tell SessionDispatcher we're done with the data

	// sulog("New Session: %s\n", s.ToString())
	// sulog("Session.Urole.perms = %+v\n", s.Urole.Perms)

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
func CreateSession(username string, w http.ResponseWriter, r *http.Request) (*Session, error) {
	expiration := time.Now().Add(SessionTimeout * time.Minute)

	//----------------------------------------------
	// TODO: lookup username in address book data
	//----------------------------------------------
	name := "Steve"
	uid := int64(211)
	RoleID := int64(0)

	//=================================================================================
	// There could be multiple sessions from the same user on different browsers.
	// These could be on the same or separate machines. We need the IP and the browser
	// to ensure uniqueness...
	//
	// The cookie is:   username + User-Agent + remote ip address
	//=================================================================================
	key := username + r.Header.Get("User-Agent") + r.RemoteAddr
	token := fmt.Sprintf("%x", md5.Sum([]byte(key)))

	s := SessionNew(token, username, name, uid, RoleID)
	cookie := http.Cookie{Name: "accord", Value: s.Token, Expires: expiration}
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
	r.AddCookie(&cookie) // need this so that the redirect to search finds the cookie
	return s, nil
}

// Refresh updates the cookie and Session with a new expire time.
//
// INPUT
//  w    - where to write the set cookie
//  r - the request where w should look for the cookie
//
// RETURNS
//  session - pointer to the new session
//-----------------------------------------------------------------------------
func (s *Session) Refresh(w http.ResponseWriter, r *http.Request) int {
	cookie, err := r.Cookie("accord")
	if nil != cookie && err == nil {
		cookie.Expires = time.Now().Add(SessionTimeout * time.Minute)
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

// SessionDelete removes the supplied Session.
// if there is a better idiomatic way to do this, please let me know.
// INPUT
//  Session  - pointer to the session to tdelete
//             list
//  username - the username from the authentication service
//  name     - the name to use in the session
//  uid      - the userid associated with username. From the auth server.
//  rid      - security role id
//
// RETURNS
//  session - pointer to the new session
//-----------------------------------------------------------------------------
func SessionDelete(s *Session) {
	// fmt.Printf("Session being deleted: %s\n", s.ToString())
	// fmt.Printf("sessions before delete:\n")
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
	// fmt.Printf("sessions after delete:\n")
	// dumpSessions()
}
