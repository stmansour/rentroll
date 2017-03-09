package ws

// These are general utilty routines to support w2ui grid components.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"rentroll/rlib"
	"strings"
	"time"
)

// SvcGridError is the generalized error structure to return errors to the grid widget
type SvcGridError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// SvcStatusResponse is the response to return status when no other data
// needs to be returned
type SvcStatusResponse struct {
	Status string `json:"status"` // typically "success"
	Recid  int64  `json:"recid"`  // set to id of newly inserted record
}

// ServiceHandler describes the handler for all services
type ServiceHandler struct {
	Cmd     string
	Handler func(http.ResponseWriter, *http.Request, *ServiceData)
	NeedBiz bool
}

// GenSearch describes a search condition
type GenSearch struct {
	Field    string `json:"field"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}

// ColSort is what the UI uses to indicate how the return values should be sorted
type ColSort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

// WebGridSearchRequest is a struct suitable for describing a webservice operation.
type WebGridSearchRequest struct {
	Cmd         string      `json:"cmd"`         // get, save, delete
	Limit       int         `json:"limit"`       // max number to return
	Offset      int         `json:"offset"`      // solution set offset
	Selected    []int       `json:"selected"`    // selected rows
	SearchLogic string      `json:"searchLogic"` // OR | AND
	Search      []GenSearch `json:"search"`      // what fields and what values
	Sort        []ColSort   `json:"sort"`        // sort criteria
}

// WebFormRequest is a struct suitable for describing a webservice operation.
type WebFormRequest struct {
	Cmd      string      `json:"cmd"`    // get, save, delete
	Recid    int         `json:"recid"`  // max number to return
	FormName string      `json:"name"`   // solution set offset
	Record   interface{} `json:"record"` // selected rows
}

// WebTypeDownRequest is a search call made by a client while the user is
// typing in something to search for and the expecation is that the solution
// set will be sent back in realtime to aid the user.  Search is a string
// to search for -- it's what the user types in.  Max is the maximum number
// of matches to return.
type WebTypeDownRequest struct {
	Search string `json:"search"`
	Max    int    `json:"max"`
}

// ServiceData is the generalized data gatherer for svcHandler. It allows all
// the common data to be centrally parsed and passed to a handler, which may
// need to parse further to get its unique data.  It includes fields for
// common data elements in web svc requests
type ServiceData struct {
	UID           int64                // user id of requester
	BID           int64                // which business
	TCID          int64                // TCID if supplied
	RAID          int64                // RAID if supplied
	RID           int64                // RAID if supplied
	RCPTID        int64                // RCPTID if supplied
	ASMID         int64                // ASMID if supplied
	Dt            time.Time            // for cmds that need a single date
	D1            time.Time            // start of date range
	D2            time.Time            // end of date range
	wsSearchReq   WebGridSearchRequest // what did the search requester ask for
	wsTypeDownReq WebTypeDownRequest
	data          string // the raw unparsed data
}

// Svcs is the table of all service handlers
var Svcs = []ServiceHandler{
	{"transactants", SvcSearchHandlerTransactants, true},
	{"transactantstd", SvcTransactantTypeDown, true},
	{"accounts", SvcSearchHandlerGLAccounts, true},
	{"asms", SvcSearchHandlerAssessments, true},
	{"asm", SvcFormHandlerAssessment, true},
	{"rar", SvcRARentables, true},
	{"receipts", SvcSearchHandlerReceipts, true},
	{"receipt", SvcFormHandlerReceipt, true},
	{"rentables", SvcSearchHandlerRentables, true},
	{"rentalagr", SvcFormHandlerRentalAgreement, true},
	{"rentalagrs", SvcSearchHandlerRentalAgr, true},
	{"person", SvcFormHandlerXPerson, true},
	{"rapeople", SvcRAPeople, true},
	{"rapayor", SvcRAPeople, true},
	{"ruser", SvcRAPeople, true},
	{"rapets", SvcRAPets, true},
	{"rentable", SvcFormHandlerRentable, true},
	{"uilists", SvcUILists, false},
}

// SvcGridErrorReturn formats an error return to the grid widget and sends it
func SvcGridErrorReturn(w http.ResponseWriter, err error) {
	var e SvcGridError
	e.Status = "error"
	e.Message = fmt.Sprintf("Error: %s\n", err.Error())
	b, _ := json.Marshal(e)
	SvcWrite(w, b)
}

// SvcGetInt64 tries to read an int64 value from the supplied string.
// If it fails for any reason, it sends writes an error message back
// to the caller and returns the error.  Otherwise, it returns an
// int64 and returns nil
func SvcGetInt64(s, errmsg string, w http.ResponseWriter) (int64, error) {
	i, err := rlib.IntFromString(s, "not an integer number")
	if err != nil {
		err = fmt.Errorf("%s: %s", errmsg, err.Error())
		SvcGridErrorReturn(w, err)
		return i, err
	}
	return i, nil
}

// SvcExtractIDFromURI extracts an int64 id value from position pos of the supplied uri.
// The URI is of the form returned by http.Request.RequestURI .  In particular:
//
//	pos:     0    1      2  3
//  uri:    /v1/rentable/34/421
//
// So, in the example uri above, a call where pos = 3 would return int64(421). errmsg
// is a string that will be used in the error message if the requested position had an
// error during conversion to int64. So in the example above, pos 3 is the RID, so
// errmsg would probably be set to "RID"
func SvcExtractIDFromURI(uri, errmsg string, pos int, w http.ResponseWriter) (int64, error) {
	var ID = int64(0)
	var err error

	sa := strings.Split(uri[1:], "/")
	// fmt.Printf("uri parts:  %v\n", sa)
	if len(sa) < pos+1 {
		err = fmt.Errorf("Expecting at least %d elements in URI: %s, but found only %d", pos+1, uri, len(sa))
		// fmt.Printf("err = %s\n", err)
		SvcGridErrorReturn(w, err)
		return ID, err
	}
	// fmt.Printf("sa[pos] = %s\n", sa[pos])
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
	if len(htmlData) == 0 {
		d.wsSearchReq.Cmd = "?"
		return nil
	}
	u, err := url.QueryUnescape(string(htmlData))
	if err != nil {
		e := fmt.Errorf("%s: Error with QueryUnescape: %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return e
	}
	fmt.Printf("Unescaped htmlData = %s\n", u)

	u = strings.TrimPrefix(u, "request=") // strip off "request=" if it is present (w2ui sends this string)
	d.data = u
	err = json.Unmarshal([]byte(u), &d.wsSearchReq)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return e
	}
	return err
}

func getGETdata(w http.ResponseWriter, r *http.Request, d *ServiceData) error {
	funcname := "getGETdata"
	s, err := url.QueryUnescape(strings.TrimSpace(r.URL.String()))
	if err != nil {
		e := fmt.Errorf("%s: Error with url.QueryUnescape:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return e
	}
	fmt.Printf("Unescaped query = %s\n", s)
	w2uiPrefix := "request="
	n := strings.Index(s, w2uiPrefix)
	fmt.Printf("n = %d\n", n)
	if n > 0 {
		fmt.Printf("Will process as Typedown\n")
		d.data = s[n+len(w2uiPrefix):]
		fmt.Printf("%s: will unmarshal: %s\n", funcname, d.data)
		if err = json.Unmarshal([]byte(d.data), &d.wsTypeDownReq); err != nil {
			e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
			SvcGridErrorReturn(w, e)
			return e
		}
		d.wsSearchReq.Cmd = "typedown"
	} else {
		fmt.Printf("Will process as web search command\n")
		d.wsSearchReq.Cmd = r.URL.Query().Get("cmd")
	}
	return nil
}

func showRequestHeaders(r *http.Request) {
	fmt.Printf("Request Headers\n")
	fmt.Printf("-----------------------\n")
	for k, v := range r.Header {
		fmt.Printf("%s: ", k)
		for i := 0; i < len(v); i++ {
			fmt.Printf("%q  ", v[i])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-----------------------\n")
}

func showWebRequest(d *ServiceData) {
	if d.wsSearchReq.Cmd == "typedown" {
		fmt.Printf("TYPEDOWN REQUEST\n")
		fmt.Printf("Search  = %q\n", d.wsTypeDownReq.Search)
		fmt.Printf("Max     = %d\n", d.wsTypeDownReq.Max)
	} else {
		fmt.Printf("Cmd         = %s\n", d.wsSearchReq.Cmd)
		fmt.Printf("Limit       = %d\n", d.wsSearchReq.Limit)
		fmt.Printf("Offset      = %d\n", d.wsSearchReq.Offset)
		fmt.Printf("searchLogic = %s\n", d.wsSearchReq.SearchLogic)
		for i := 0; i < len(d.wsSearchReq.Search); i++ {
			fmt.Printf("search[%d] - Field = %s,  Type = %s,  Value = %s,  Operator = %s\n", i, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Type, d.wsSearchReq.Search[i].Value, d.wsSearchReq.Search[i].Operator)
		}
		for i := 0; i < len(d.wsSearchReq.Sort); i++ {
			fmt.Printf("sort[%d] - Field = %s,  Direction = %s\n", i, d.wsSearchReq.Sort[i].Field, d.wsSearchReq.Sort[i].Direction)
		}
	}
}

// V1ServiceHandler is the main dispatch point for WEB SERVICE requests
//
// The expected input is of the form:
//		request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D
// This is exactly what the w2ui grid sends as a request.
//
// Decoded, this message looks something like this:
//		request={"cmd":"get","selected":[],"limit":100,"offset":0}
//
// The leading "request=" is optional. This routine parses the basic information, then contacts an appropriate
// handler for more detailed processing.  It will set the Cmd member variable.
//
// W2UI sometimes sends requests that look like this: request=%7B%22search%22%3A%22s%22%2C%22max%22%3A250%7D
// using HTTP GET (rather than its more typical POST).  The command decodes to this: request={"search":"s","max":250}
//
//-----------------------------------------------------------------------------------------------------------
func V1ServiceHandler(w http.ResponseWriter, r *http.Request) {
	funcname := "V1ServiceHandler"
	fmt.Printf("==========================================================================================\n")
	fmt.Printf("Entered %s. r.Method = %s, URL = %s \n", funcname, r.Method, r.URL.String())
	var err error
	var d ServiceData

	showRequestHeaders(r)
	switch r.Method {
	case "POST":
		if nil != getPOSTdata(w, r, &d) {
			return
		}
	case "GET":
		if nil != getGETdata(w, r, &d) {
			return
		}
	}
	showWebRequest(&d)

	//-----------------------------------------------------------------------
	// General form is /v1/{subservice}/{BID}/{ID}
	//-----------------------------------------------------------------------
	path := "/v1/"                         // this is the part of the URL that got us into this handler
	subURLPath := r.RequestURI[len(path):] // this pulls off the specific request

	//-----------------------------------------------------------------------
	// pathElements: 0         1     2
	// Break up {subservice}/{BUI}/{ID} into an array of strings
	// BID is common to nearly all commands
	//-----------------------------------------------------------------------
	ss := strings.Split(subURLPath, "?") // it could be GET command
	pathElements := strings.Split(ss[0], "/")
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
	var abud []string
	if len(pathElements) >= 2 {
		abud = strings.Split(pathElements[1], "?")                            // in this array, abud[0] will always be what want to parse
		d.BID, err = rlib.IntFromString(abud[0], "bad request integer value") // assume it's a BID
		if err != nil {
			var ok bool // OK, let's see if it's a BUD
			d.BID, ok = rlib.RRdb.BUDlist[abud[0]]
			if !ok {
				d.BID = 0
			}
		}
	}

	fmt.Printf("d.BID = %d\n", d.BID)

	//-----------------------------------------------------------------------
	//  Now call the appropriate handler to do the rest
	//-----------------------------------------------------------------------
	found := false
	for i := 0; i < len(Svcs); i++ {
		if Svcs[i].Cmd == requestedSvc {
			if Svcs[i].NeedBiz && d.BID == 0 {
				e := fmt.Errorf("Could not identify business: %s", abud[0])
				fmt.Printf("***ERROR IN URL***  %s", e.Error())
				SvcGridErrorReturn(w, err)
			}
			Svcs[i].Handler(w, r, &d)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("**** YIPES! **** %s - Handler not found\n", subURLPath)
		e := fmt.Errorf("Service not recognized: %s", requestedSvc)
		fmt.Printf("***ERROR IN URL***  %s", e.Error())
		SvcGridErrorReturn(w, e)
	}
	fmt.Printf("\n-------------------------------------\n\n")
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
	q := "SELECT * FROM " + table + " WHERE"
	return gridBuildQueryWhereClause(q, table, srch, order, d, p)
}

func gridBuildQueryWhereClause(q, table, srch, order string, d *ServiceData, p interface{}) (string, string) {
	qw := ""
	if len(d.wsSearchReq.Search) > 0 {
		val := reflect.ValueOf(p).Elem() // reflect value of input p
		count := 0
		for i := 0; i < len(d.wsSearchReq.Search); i++ {
			if d.wsSearchReq.Search[i].Field == "recid" || len(d.wsSearchReq.Search[i].Value) == 0 {
				continue
			}
			// look for this field in p
			for j := 0; j < val.NumField(); j++ {
				field := val.Field(j)                   // this is field[j] of p
				n := val.Type().Field(j).Name           // variable name for field(i)
				if n != d.wsSearchReq.Search[i].Field { // is this the field we're looking for?
					continue
				}
				t := field.Type().String() // Is it a type we can handle?
				if t != "string" {         // TODO: handle all data types
					continue
				}
				switch d.wsSearchReq.Search[i].Operator {
				case "begins":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s like '%s%%'", &count)
				case "ends":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s like '%%%s'", &count)
				case "is":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s='%s'", &count)
				case "between":
					qw = gridHandleField(qw, d.wsSearchReq.SearchLogic, d.wsSearchReq.Search[i].Field, d.wsSearchReq.Search[i].Value, " %s like '%%%s%%'", &count)
				default:
					fmt.Printf("Unhandled search operator: %s\n", d.wsSearchReq.Search[i].Operator)
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
	if len(d.wsSearchReq.Sort) > 0 {
		for i := 0; i < len(d.wsSearchReq.Sort); i++ {
			if i > 0 {
				q += ","
			}
			q += d.wsSearchReq.Sort[i].Field + " " + d.wsSearchReq.Sort[i].Direction
		}
	} else {
		q += order
	}

	// now set up the offset and limit
	q += fmt.Sprintf(" LIMIT %d OFFSET %d", d.wsSearchReq.Limit, d.wsSearchReq.Offset)
	return q, qw
}

// GetRowCount returns the number of database rows in the supplied table with the supplied where clause
func GetRowCount(table, where string) (int64, error) {
	count := int64(0)
	var err error
	s := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", table, where)
	de := rlib.RRdb.Dbrr.QueryRow(s).Scan(&count)
	if de != nil {
		err = fmt.Errorf("GetRowCount: query=\"%s\"    err = %s", s, de.Error())
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
	SvcWrite(w, b)
}

// SvcWrite is a general write routine for service calls... it is a bottleneck
// where we can place debug statements as needed.
func SvcWrite(w http.ResponseWriter, b []byte) {
	fmt.Printf("first 200 chars of response: %-200.200s\n", string(b))
	// fmt.Printf("\nResponse Data:  %s\n\n", string(b))
	w.Write(b)
}

// SvcWriteSuccessResponse is used to complete a successful write operation on w2ui form save requests.
func SvcWriteSuccessResponse(w http.ResponseWriter) {
	var g = SvcStatusResponse{Status: "success"}
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// SvcWriteSuccessResponseWithID is used to complete a successful write operation on w2ui form save requests.
func SvcWriteSuccessResponseWithID(w http.ResponseWriter, id int64) {
	var g = SvcStatusResponse{Status: "success", Recid: id}
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}
