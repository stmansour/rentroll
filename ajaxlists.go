package main

import (
	"fmt"
	"io"
	"net/http"
	"rentroll/rlib"
)

// SvcUILists returns JSON for the Javascript lists needed for the UI
func SvcUILists(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcUILists\n")

	response := `yesNoList = [ 'no', 'yes' ];
assignmentTimeList = [ 'unset', 'Pre-Assign', 'Commencement'];
`

	io.WriteString(w, response)

	s := "businesses = ["
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

	s = "companyOrPerson = ["
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
	s = "stateAbbr = ["
	l = len(stateAbbr)
	for j := 0; j < l; j++ {
		s += "'" + stateAbbr[j] + "'"
		if j+1 < l {
			s += ","
		}
	}
	s += "];\n"
	io.WriteString(w, s)
}
