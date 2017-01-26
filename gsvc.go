package main

// These are general utilty routines to support w2ui grid components.

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"rentroll/rlib"
	"strings"
)

// SvcGridError is the generalized error structure to return errors to the grid widget
type SvcGridError struct {
	Status  string
	Message string
}

// ServiceHandler describes the handler for all services
type ServiceHandler struct {
	Cmd     string
	Handler func(http.ResponseWriter, *http.Request, *ServiceData)
}

// WebRequest is a struct suitable for describing a webservice operation.
type WebRequest struct {
	Cmd         string `json:"cmd"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	Selected    []int  `json:"selected"`
	SearchLogic string `json:"searchLogic"`
	Search      []struct {
		Field    string `json:"field"`
		Type     string `json:"type"`
		Value    string `json:"value"`
		Operator string `json:"operator"`
	} `json:"search"`
	Sort []struct {
		Field     string `json:"field"`
		Direction string `json:"direction"`
	} `json:"sort"`
}

// ServiceData is the generalized data gatherer for svcHandler. It allows all the common data
// to be centrally parsed and passed to a specific handler, which may need to parse further
// to get its unique data.
type ServiceData struct {
	UID    int64      // user id of requester
	BID    int64      // which business
	TCID   int64      // TCID if supplied
	RAID   int64      // RAID if supplied
	RID    int64      // RAID if supplied
	webreq WebRequest // what did the grid ask for
	data   string     // the raw unparsed data
}

// Svcs is the table of all service handlers
var Svcs = []ServiceHandler{
	{"transactants", SvcSearchHandlerTransactants},
	{"accounts", SvcSearchHandlerGLAccounts},
	{"rentables", SvcSearchHandlerRentables},
	{"rentalagrs", SvcSearchHandlerRentalAgr},
	{"xperson", SvcFormHandlerXPerson},
	{"xrentable", SvcFormHandlerRentable},
	{"xrentalagr", SvcFormHandlerRentalAgreement},
	{"uilists", SvcUILists},
}

// SvcGridErrorReturn formats an error return to the grid widget and sends it
func SvcGridErrorReturn(w http.ResponseWriter, err error) {
	var e SvcGridError
	e.Status = "error"
	e.Message = fmt.Sprintf("Error converting g:  %s\n", err.Error())
	b, _ := json.Marshal(e)
	w.Write(b)
}

// SvcGetInt64 tries to read an int64 value from the supplied string.
// If it fails for any reason, it sends writes an error message back
// to the caller and returns the error.  Otherwise, it returns an
// int64 and returns nil
func SvcGetInt64(s, errmsg string, w http.ResponseWriter) (int64, error) {
	i, err := rlib.IntFromString(s, "not an integer number")
	if err != nil {
		err = fmt.Errorf("%s: %s\n", errmsg, err.Error())
		SvcGridErrorReturn(w, err)
		return i, err
	}
	return i, nil
}

// SvcExtractIDFromURI extracts an int64 id value from position pos of the supplied uri.
// The URI is of the form returned by http.Request.RequestURI .  In particular:
//
//	pos:     0    1         2  3
//  uri:    /gsvc/xrentable/34/421
//
// So, in the example uri above, a call where pos = 3 would return int64(421). errmsg
// is a string that will be used in the error message if the requested position had an
// error during conversion to int64. So in the example above, pos 3 is the RID, so
// errmsg would probably be set to "RID"
func SvcExtractIDFromURI(uri, errmsg string, pos int, w http.ResponseWriter) (int64, error) {
	var ID = int64(0)
	var err error

	sa := strings.Split(uri[1:], "/")
	if len(sa) < pos+1 {
		err = fmt.Errorf("Expecting at least %d elements in URI: %s, but found only %d\n", pos+1, uri, len(sa))
		SvcGridErrorReturn(w, err)
		return ID, err
	}
	ID, err = SvcGetInt64(sa[pos], errmsg, w)
	return ID, err
}

func getPOSTdata(w http.ResponseWriter, r *http.Request, d *ServiceData) error {
	funcname := "getPOSTdata"
	var err error
	htmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := fmt.Errorf("%s: Error reading message Body: %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return e
	}
	fmt.Printf("htmlData = %s\n", htmlData)
	u, err := url.QueryUnescape(string(htmlData))
	if err != nil {
		e := fmt.Errorf("%s: Error with QueryUnescape: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return e
	}

	requestHeader := "request=" // this is what w2ui starts all its grid requests with
	i := strings.Index(u, requestHeader)
	if i >= 0 {
		u = u[i+len(requestHeader):]
		d.data = u
	}
	err = json.Unmarshal([]byte(u), &d.webreq)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return e
	}
	return err
}

func showRequestHeaders(r *http.Request) {
	fmt.Printf("Request Headers\n")
	fmt.Printf("-----------------------------------------------------------------------------------------------\n")
	for k, v := range r.Header {
		fmt.Printf("%s: ", k)
		for i := 0; i < len(v); i++ {
			fmt.Printf("%q  ", v[i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-----------------------------------------------------------------------------------------------\n")
}

func showWebRequest(d *ServiceData) {
	fmt.Printf("Cmd         = %s\n", d.webreq.Cmd)
	fmt.Printf("Limit       = %d\n", d.webreq.Limit)
	fmt.Printf("Offset      = %d\n", d.webreq.Offset)
	fmt.Printf("searchLogic = %s\n", d.webreq.SearchLogic)
	for i := 0; i < len(d.webreq.Search); i++ {
		fmt.Printf("search[%d] - Field = %s,  Type = %s,  Value = %s,  Operator = %s\n", i, d.webreq.Search[i].Field, d.webreq.Search[i].Type, d.webreq.Search[i].Value, d.webreq.Search[i].Operator)
	}
	for i := 0; i < len(d.webreq.Sort); i++ {
		fmt.Printf("sort[%d] - Field = %s,  Direction = %s\n", i, d.webreq.Sort[i].Field, d.webreq.Sort[i].Direction)
	}
}

// gridServiceHandler is the main dispatch point for w2ui grid service requests
//
// The expected input is of the form:
//		request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D
// This is exactly what the w2ui grid sends as a request.
//
// Decoded, this message looks like this:
//		request={"cmd":"get","selected":[],"limit":100,"offset":0}
//
// Some routines need more information than what is encoded in the request. In
// these cases the extra information is passed in the request URI.  This information
// must be parsed by each function, as the data needed will be different from
// from function to function.
//-----------------------------------------------------------------------------------------------------------
func gridServiceHandler(w http.ResponseWriter, r *http.Request) {
	funcname := "gridServiceHandler"
	fmt.Printf("Entered %s.  r.Method = %s\n", funcname, r.Method)
	var err error
	var d ServiceData

	showRequestHeaders(r)
	switch r.Method {
	case "POST":
		if nil != getPOSTdata(w, r, &d) {
			return
		}
	case "GET":
		d.webreq.Cmd = r.URL.Query().Get("cmd")
	}
	showWebRequest(&d)

	//-----------------------------------------------------------------------
	// General form is /gsvc/{subservice}/{BID}/{ID}
	//-----------------------------------------------------------------------
	path := "/gsvc/"                       // this is the part of the URL that got us into this handler
	subURLPath := r.RequestURI[len(path):] // this pulls off the specific request

	//-----------------------------------------------------------------------
	// pathElements: 0         1     2
	// Break up {subservice}/{BID}/{ID} into an array of strings
	// BID is common to nearly all commands
	//-----------------------------------------------------------------------
	pathElements := strings.Split(subURLPath, "/")
	requestedSvc := pathElements[0]

	//-----------------------------------------------------------------------
	//  DEBUGGING INFORMATION
	//-----------------------------------------------------------------------
	fmt.Printf("Command specific info in URL:  %s\n", subURLPath) // print before we strip it off
	for i := 0; i < len(pathElements); i++ {
		fmt.Printf("%d. %s\n", i, pathElements[i])
	}

	//-----------------------------------------------------------------------
	// There are commands that have associated ID (UILists, for example), and
	// if we're processing such a GET request there may be "?param1=val1&..."
	// on the end of this string.  So, before we try to parse the BID/BUD,
	// we'll split it at the '?' character and
	// process the first bit for the BID...
	//-----------------------------------------------------------------------
	abud := strings.Split(pathElements[2], "?")                           // in this array, abud[0] will always be what want to parse
	d.BID, err = rlib.IntFromString(abud[0], "bad request integer value") // assume it's a BID
	if err != nil {
		var ok bool // OK, let's see if it's a BUD
		d.BID, ok = rlib.RRdb.BUDlist[abud[0]]
		if !ok {
			e := fmt.Errorf("Could not identify business: %s\n", abud[0])
			fmt.Printf("***ERROR IN URL***  %s", e.Error())
			SvcGridErrorReturn(w, err)
		}
	}

	//-----------------------------------------------------------------------
	//  Now call the appropriate handler to do the rest
	//-----------------------------------------------------------------------
	found := false
	for i := 0; i < len(Svcs); i++ {
		if Svcs[i].Cmd == requestedSvc {
			fmt.Printf("%s - Handler found\n", subURLPath)
			Svcs[i].Handler(w, r, &d)
			found = true
			break
		}
	}
	if !found {
		e := fmt.Errorf("Service not recognized: %s\n", requestedSvc)
		fmt.Printf("***ERROR IN URL***  %s", e.Error())
		SvcGridErrorReturn(w, err)
	}
	fmt.Printf("\n-------------------------------------\n\n")
}

// SvcUILists returns JSON for the Javascript lists needed for the UI
func SvcUILists(w http.ResponseWriter, r *http.Request, d *ServiceData) {
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
}

func gridHandleField(q, logic, field, value, format string, count *int) string {
	if *count > 0 {
		q += " " + logic
	}
	q += fmt.Sprintf(format, field, value)
	*count++
	return q
}

// gridBuildQuery builds a query from the supplied base and the sort / search parameters
// in the supplied w2ui grid structure.  To play with this routine in isolation
// use this:
//				https://play.golang.org/p/HOkP77h0Ts
//
// Parameters:
// 		table - the name of the table to query
// 		srch  - the default where clause. Used if the Search info is empty. Does not require
//              the keyword "WHERE".  That is, flter == "BID=1" when you want the where clause
//              to be "WHERE BID=1"
// 		order - default sorting clause. Used when Sort is empty
//      p     - pointer to the struct associated with the database table. It is used to match
//              the fields passed in by the UI.  We need to determine what type of fields
//              they are in order to properly construct the WHERE clause
//
// Returns:
//     string - the full query
//     string - the WHERE clause suitable for a COUNT(*) query
//----------------------------------------------------------------------------------------------
func gridBuildQuery(table, srch, order string, d *ServiceData, p interface{}) (string, string) {
	// Handle Search
	// q := fmt.Sprintf("SELECT * FROM "+table+" WHERE BID=%d AND (", d.BID)
	q := "SELECT * FROM " + table + " WHERE"
	qw := ""
	if len(d.webreq.Search) > 0 {
		val := reflect.ValueOf(p).Elem() // reflect value of input p
		count := 0
		for i := 0; i < len(d.webreq.Search); i++ {
			if d.webreq.Search[i].Field == "recid" || len(d.webreq.Search[i].Value) == 0 {
				continue
			}
			// look for this field in p
			for j := 0; j < val.NumField(); j++ {
				field := val.Field(j)              // this is field[j] of p
				n := val.Type().Field(j).Name      // variable name for field(i)
				if n != d.webreq.Search[i].Field { // is this the field we're looking for?
					continue
				}
				t := field.Type().String() // Is it a type we can handle?
				if t != "string" {         // TODO: handle all data types
					continue
				}
				switch d.webreq.Search[i].Operator {
				case "begins":
					qw = gridHandleField(qw, d.webreq.SearchLogic, d.webreq.Search[i].Field, d.webreq.Search[i].Value, " %s like '%s%%'", &count)
				case "ends":
					qw = gridHandleField(qw, d.webreq.SearchLogic, d.webreq.Search[i].Field, d.webreq.Search[i].Value, " %s like '%%%s'", &count)
				case "is":
					qw = gridHandleField(qw, d.webreq.SearchLogic, d.webreq.Search[i].Field, d.webreq.Search[i].Value, " %s='%s'", &count)
				case "between":
					qw = gridHandleField(qw, d.webreq.SearchLogic, d.webreq.Search[i].Field, d.webreq.Search[i].Value, " %s like '%%%s%%'", &count)
				default:
					fmt.Printf("Unhandled search operator: %s\n", d.webreq.Search[i].Operator)
				}
			}
		}
		if len(qw) > 0 {
			qw = fmt.Sprintf(" BID=%d AND (%s)", d.BID, qw)
		}
		q += qw         // add the WHERE information
		if count == 0 { // if we didn't match any of the search criteria...
			q += " " + srch // then revert to the default search clause
			qw = srch
		}
	} else {
		q += " " + srch // no search info supplied, use the default
		qw = srch
	}

	// Handle any Sorting requests
	q += " ORDER BY "
	if len(d.webreq.Sort) > 0 {
		for i := 0; i < len(d.webreq.Sort); i++ {
			if i > 0 {
				q += ","
			}
			q += d.webreq.Sort[i].Field + " " + d.webreq.Sort[i].Direction
		}
	} else {
		q += order
	}

	// now set up the offset and limit
	q += fmt.Sprintf(" LIMIT %d OFFSET %d", d.webreq.Limit, d.webreq.Offset)
	return q, qw
}

// GetRowCount returns the number of database rows in the supplied table with the supplied where clause
func GetRowCount(table, where string) (int64, error) {
	count := int64(0)
	var err error
	s := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", table, where)
	de := rlib.RRdb.Dbrr.QueryRow(s).Scan(&count)
	if de != nil {
		err = fmt.Errorf("GetRowCount: query=\"%s\"    err = %s\n", s, de.Error())
	}
	return count, err
}

// SvcWriteResponse finishes the transaction with the W2UI client
func SvcWriteResponse(g interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(g)
	if err != nil {
		e := fmt.Errorf("Error marshaling json data: %s", err.Error())
		rlib.Ulog("SvcWriteResponse: %s\n", err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	fmt.Printf("first 200 chars of response: %-200.200s\n", string(b))
	// fmt.Printf("\nResponse Data:  %s\n\n", string(b))
	w.Write(b)
}

// SvcWriteSuccessResponse is used to complete a successful write operation on w2ui form save requests.
func SvcWriteSuccessResponse(w http.ResponseWriter) {
	var g struct {
		Status string `json:"status"`
	}
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}
