package main

import (
	"encoding/json"
	"fmt"
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

// W2uiGridRequest is a struct suitable for holding the json data
// posted to a web service by the W2ui Grid.
type W2uiGridRequest struct {
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
	UID  int64           // user id of requester
	BID  int64           // which business
	TCID int64           // TCID if supplied
	RAID int64           // RAID if supplied
	RID  int64           // RAID if supplied
	greq W2uiGridRequest // what did the grid ask for
	data string          // the raw unparsed data
}

// Svcs is the table of all service handlers
var Svcs = []ServiceHandler{
	{"transactants", SvcTransactants},
	{"xperson", SvcXPerson},
	{"accounts", SvcGLAccounts},
	{"rentables", SvcRentables},
	{"xrentable", SvcRentable},
}

// SvcGridErrorReturn formats an error return to the grid widget and sends it
func SvcGridErrorReturn(w http.ResponseWriter, err error) {
	var e SvcGridError
	e.Status = "error"
	e.Message = fmt.Sprintf("Error converting g:  %s\n", err.Error())
	b, _ := json.Marshal(e)
	w.Write(b)
}

// gridServiceHandler is the main dispatch point for w2ui grid service requests
//
// The expected input is of the form:
//		request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D
// This is exactly what the w2ui grid sends as a request.
//
// Some routines need more information than what is encoded in the request. In
// these cases the extra information is passed in the request URI.  This information
// must be parsed by each function, as the data needed will be different from
// from function to function.
//-----------------------------------------------------------------------------------------------------------
func gridServiceHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Entered gridServiceHandler\n")

	funcname := "gridServiceHandler"
	var d ServiceData
	htmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e := fmt.Errorf("%s: Error reading message Body: %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	u, err := url.QueryUnescape(string(htmlData))
	if err != nil {
		e := fmt.Errorf("%s: Error with QueryUnescape: %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	requestHeader := "request=" // this is what w2ui starts all its grid requests with
	i := strings.Index(u, requestHeader)
	if i < 0 {
		e := fmt.Errorf("%s: Bad request format.  Looking for \"request=\"...  found: %s\n", funcname, u)
		SvcGridErrorReturn(w, e)
		return
	}
	u = u[i+len(requestHeader):]
	d.data = u
	err = json.Unmarshal([]byte(u), &d.greq)
	if err != nil {
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	fmt.Printf("Cmd         = %s\n", d.greq.Cmd)
	fmt.Printf("Limit       = %d\n", d.greq.Limit)
	fmt.Printf("Offset      = %d\n", d.greq.Offset)
	fmt.Printf("searchLogic = %s\n", d.greq.SearchLogic)
	for i := 0; i < len(d.greq.Search); i++ {
		fmt.Printf("search[%d] - Field = %s,  Type = %s,  Value = %s,  Operator = %s\n", i, d.greq.Search[i].Field, d.greq.Search[i].Type, d.greq.Search[i].Value, d.greq.Search[i].Operator)
	}
	for i := 0; i < len(d.greq.Sort); i++ {
		fmt.Printf("sort[%d] - Field = %s,  Direction = %s\n", i, d.greq.Sort[i].Field, d.greq.Sort[i].Direction)
	}

	fmt.Printf("Full URI:  %s\n", r.RequestURI) // print before we strip it off
	path := "/gsvc/"                            // this is the part of the URL that got us into this handler
	cmdinfo := r.RequestURI[len(path):]         // this pulls off the specific request
	fmt.Printf("URI info:  %s\n", cmdinfo)      // print before we strip it off

	sa := strings.Split(cmdinfo, "/")
	for i := 0; i < len(sa); i++ {
		fmt.Printf("%d. %s\n", i, sa[i])
	}
	cmdinfo = sa[0]
	d.UID, err = rlib.IntFromString(sa[1], "bad request integer value")
	if err != nil {
		SvcGridErrorReturn(w, err)
	}
	d.BID, err = rlib.IntFromString(sa[2], "bad request integer value")
	if err != nil {
		SvcGridErrorReturn(w, err)
	}

	for i := 0; i < len(Svcs); i++ {
		if Svcs[i].Cmd == cmdinfo {
			Svcs[i].Handler(w, r, &d)
			break
		}
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
	// q := fmt.Sprintf("SELECT * FROM "+table+" WHERE BID=%d AND (", d.BID)
	q := "SELECT * FROM " + table + " WHERE"
	qw := ""
	if len(d.greq.Search) > 0 {
		val := reflect.ValueOf(p).Elem() // reflect value of input p
		count := 0
		for i := 0; i < len(d.greq.Search); i++ {
			if d.greq.Search[i].Field == "recid" || len(d.greq.Search[i].Value) == 0 {
				continue
			}
			// look for this field in p
			for j := 0; j < val.NumField(); j++ {
				field := val.Field(j)            // this is field[j] of p
				n := val.Type().Field(j).Name    // variable name for field(i)
				if n != d.greq.Search[i].Field { // is this the field we're looking for?
					continue
				}
				t := field.Type().String() // Is it a type we can handle?
				if t != "string" {         // TODO: handle all data types
					continue
				}
				switch d.greq.Search[i].Operator {
				case "begins":
					qw = gridHandleField(qw, d.greq.SearchLogic, d.greq.Search[i].Field, d.greq.Search[i].Value, " %s like '%s%%'", &count)
				case "ends":
					qw = gridHandleField(qw, d.greq.SearchLogic, d.greq.Search[i].Field, d.greq.Search[i].Value, " %s like '%%%s'", &count)
				case "is":
					qw = gridHandleField(qw, d.greq.SearchLogic, d.greq.Search[i].Field, d.greq.Search[i].Value, " %s='%s'", &count)
				case "between":
					qw = gridHandleField(qw, d.greq.SearchLogic, d.greq.Search[i].Field, d.greq.Search[i].Value, " %s like '%%%s%%'", &count)
				default:
					fmt.Printf("Unhandled search operator: %s\n", d.greq.Search[i].Operator)
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
	if len(d.greq.Sort) > 0 {
		for i := 0; i < len(d.greq.Sort); i++ {
			if i > 0 {
				q += ","
			}
			q += d.greq.Sort[i].Field + " " + d.greq.Sort[i].Direction
		}
	} else {
		q += order
	}

	// now set up the offset and limit
	q += fmt.Sprintf(" LIMIT %d OFFSET %d", d.greq.Limit, d.greq.Offset)
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
	fmt.Printf("first 100 chars of response: %100.100s\n", string(b))
	// fmt.Printf("\nResponse Data:  %s\n\n", string(b))
	w.Write(b)
}

// SvcGLAccounts generates a report of all GLAccounts for a the business unit
// called out in d.BID
func SvcGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcGLAccounts\n")
	var p rlib.GLAccount
	var err error
	var g struct {
		Status  string           `json:"status"`
		Total   int64            `json:"total"`
		Records []rlib.GLAccount `json:"records"`
	}

	srch := fmt.Sprintf("BID=%d", d.BID) // default WHERE clause
	order := "GLNumber ASC, Name ASC"    // default ORDER
	q, qw := gridBuildQuery("GLAccount", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("GLAccount", qw)
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("db query = %s\n", q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := d.greq.Offset
	count := 0
	for rows.Next() {
		var p rlib.GLAccount
		rlib.ReadGLAccounts(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		count++ // update the count only after adding the record
		if count >= d.greq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}

	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}

// SvcTransactants generates a report of all Transactants defined business d.BID
func SvcTransactants(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcTransactants")
	var p rlib.Transactant
	var err error
	var g struct {
		Status  string             `json:"status"`
		Total   int64              `json:"total"`
		Records []rlib.Transactant `json:"records"`
	}

	srch := fmt.Sprintf("BID=%d", d.BID)   // default WHERE clause
	order := "LastName ASC, FirstName ASC" // default ORDER
	q, qw := gridBuildQuery("Transactant", srch, order, d, &p)
	fmt.Printf("db query = %s\n", q)

	g.Total, err = GetRowCount("Transactant", qw) // total number of rows that match the criteria
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := int64(d.greq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.Transactant
		rlib.ReadTransactants(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		count++ // update the count only after adding the record
		if count >= d.greq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}
	fmt.Printf("Loaded %d transactants\n", len(g.Records))
	fmt.Printf("g.Total = %d\n", g.Total)
	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
