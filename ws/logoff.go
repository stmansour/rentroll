package ws

import (
	"net/http"
	"rentroll/rlib"
)

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

	//-----------------------------------------------------------------
	// If we get this far, it means that we do have a session (d.sess)
	// Just delete the session.  This will also expire the cookie
	//-----------------------------------------------------------------
	rlib.SessionDelete(d.sess, w, r)
	SvcWriteSuccessResponse(w)
}
