package ws

import (
	"html/template"
	"rentroll/rlib"
)

// RRfuncMap is a map of functions passed to each html page that can be referenced
// as needed to produce the page
var RRfuncMap map[string]interface{}

// InitReports initializes the reports subsystem. Historically it
// did more than it does today. Currently, it initializes the map
// of functions that can be used by a startup web page such as home.html
// or rhome.html
//-----------------------------------------------------------------------------
func InitReports() {
	RRfuncMap = template.FuncMap{
		"DateToString": rlib.DateToString,
		"GetVersionNo": rlib.GetVersionNo,
		"getBuildTime": rlib.GetBuildTime,
		"RRCommaf":     rlib.RRCommaf,
	}
}
