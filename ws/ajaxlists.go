package ws

import (
	"fmt"
	"io"
	"net/http"
	"rentroll/rlib"
)

// String2Int64MapToJSList generates a string of JS code that assigns
// all the map strings in m to an array.  Suitable for a JS eval call.
func String2Int64MapToJSList(name string, m *rlib.Str2Int64Map) string {
	s := name + "=["
	l := len(*m)
	i := 0
	for k := range *m {
		s += "'" + k + "'"
		if i+1 < l {
			s += ","
		}
		i++
	}
	s += "];\n"
	return s
}

// SvcUILists returns JSON for the Javascript lists needed for the UI
// wsdoc {
//  @Title  Get UI Lists
//	@URL /v1/accounts/:BUI
//  @Method  GET, POST
//	@Synopsis Return string lists that are used in the UI
//  @Description Return data can be processed by eval() to create the string lists used in the UI
//	@Input WebRequest
//  @Response string
// wsdoc }
func SvcUILists(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcUILists\n")

	response := `yesNoList = [ 'no', 'yes' ];
assignmentTimeList = [ 'unset', 'Pre-Assign', 'Commencement'];
`
	io.WriteString(w, response)

	s := "app.businesses = ["
	l := len(rlib.RRdb.BUDlist)
	i := 0
	for k := range rlib.RRdb.BUDlist {
		s += "'" + k + "'"
		if i+1 < l {
			s += ","
		}
		i++
	}
	s += "];\n"
	io.WriteString(w, s)

	s = "app.companyOrPerson = ["
	l = len(rlib.CompanyOrPersonMap)
	i = 0
	for k := range rlib.CompanyOrPersonMap {
		s += "'" + k + "'"
		if i+1 < l {
			s += ","
		}
		i++
	}
	s += "];\n"
	io.WriteString(w, s)

	var stateAbbr = []string{"AK", "AL", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}
	s = "app.stateAbbr = ["
	l = len(stateAbbr)
	for j := 0; j < l; j++ {
		s += "'" + stateAbbr[j] + "'"
		if j+1 < l {
			s += ","
		}
	}
	s += "];\n"
	io.WriteString(w, s)

	io.WriteString(w, String2Int64MapToJSList("app.renewalMap", &rlib.RenewalMap))
}
