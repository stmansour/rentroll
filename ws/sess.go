package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// SvcSession represents an entry in the session table
type SvcSession struct {
	Recid    int64     `json:"recid"`
	Token    string    // this is the md5 hash, unique id
	Username string    // associated username
	Name     string    // user's preferred name if it exists, otherwise the FirstName
	UID      int64     // user's db uid
	CoCode   int64     // logged in user's company (from Accord Directory)
	ImageURL string    // user's picture
	Expire   time.Time // when does the cookie expire
	RoleID   int64     // security role id
}

// SvcSessionTable holds the struct for grids response
type SvcSessionTable struct {
	Status  string       `json:"status"`
	Total   int64        `json:"total"`
	Records []SvcSession `json:"records"`
}

// SvcHandlerSessions returns the list of active sessions
func SvcHandlerSessions(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcHandlerSessions"
	var err error

	rlib.Console("Entered %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "get":
		svcGetSessions(w, r, d) // it is a query for the grid.
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// svcGetSessions returns a representation of the session table.
//
// /v1/sessions/ will return the whole table provided you have the access.
//--------------------------------------------------------------------------
func svcGetSessions(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "svcGetSessions"
	var g SvcSessionTable

	m := rlib.GetSessionTable()
	i := int64(0)
	for _, v := range m {
		var ss SvcSession
		rlib.MigrateStructVals(v, &ss)
		ss.Recid = i
		g.Records = append(g.Records, ss)
		i++
	}

	g.Status = "success"
	g.Total = i
	SvcWriteResponse(d.BID, &g, w)
}
