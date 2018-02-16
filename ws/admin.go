package ws

import (
	"net/http"
	"rentroll/rlib"
)

// SvcDisableConsole disables console messages from printing out
func SvcDisableConsole(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	rlib.DisableConsole()
	SvcWriteSuccessResponse(d.BID, w)
}

// SvcEnableConsole enables console messages to print out
func SvcEnableConsole(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	rlib.EnableConsole()
	SvcWriteSuccessResponse(d.BID, w)
}
