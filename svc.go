package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
)

// ServiceData is the generalized data gatherer for svcHandler. It allows all the common data
// to be centrally parsed and passed to a specific handler, which may need to parse further
// to get its unique data.
type ServiceData struct {
	BID    int64
	Cmd    string
	Limit  int
	Offset int
}

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

// Svcs is the table of all service handlers
var Svcs = []ServiceHandler{
	{"transactants", SvcTransactants},
	{"GLAccounts", SvcGLAccounts},
}

// SvcGridErrorReturn formats an error return to the grid widget and sends it
func SvcGridErrorReturn(w http.ResponseWriter, err error) {
	var e SvcGridError
	e.Status = "error"
	e.Message = fmt.Sprintf("Error converting g:  %s\n", err.Error())
	b, _ := json.Marshal(e)
	w.Write(b)
}

// svcHandler is the main dispatch point for service requests
func svcHandler(w http.ResponseWriter, r *http.Request) {
	var t W2uiGridRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cmd    = %s\n", t.Cmd)
	fmt.Printf("Limit  = %d\n", t.Limit)
	fmt.Printf("Offset = %d\n", t.Offset)

	path := "/svc/"                     // this is the part of the URL that got us into this handler
	cmdinfo := r.RequestURI[len(path):] // this pulls off the specific request
	fmt.Printf("svcHandler must process \"%s\"\n", cmdinfo)

	var d ServiceData
	d.BID = int64(1)
	d.Cmd = r.FormValue("cmd")
	d.Limit, _ = strconv.Atoi(r.FormValue("limit"))
	d.Offset, _ = strconv.Atoi(r.FormValue("offset"))

	for i := 0; i < len(Svcs); i++ {
		if Svcs[i].Cmd == cmdinfo {
			Svcs[i].Handler(w, r, &d)
			break
		}
	}
	fmt.Printf("\n-------------------------------------\n\n")
}

// SvcTransactants generates a report of all Businesses defined in the database.
func SvcTransactants(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	rows, err := rlib.RRdb.Prepstmt.GetAllTransactantsForBID.Query(d.BID)
	rlib.Errcheck(err)
	defer rows.Close()

	var g struct {
		Status  string             `json:"status"`
		Total   int                `json:"total"`
		Records []rlib.Transactant `json:"records"`
	}

	i := 0
	for rows.Next() {
		var p rlib.Transactant
		rlib.ReadTransactants(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		i++
	}
	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Total = len(g.Records) // this is also the value of i, but what we're doing should be clear
	g.Status = "success"
	b, err := json.Marshal(g)
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	w.Write(b)
}

// SvcGLAccounts generates a report of all Businesses defined in the database.
func SvcGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	rows, err := rlib.RRdb.Prepstmt.GetLedgerList.Query(d.BID)
	rlib.Errcheck(err)
	defer rows.Close()

	var g struct {
		Status  string           `json:"status"`
		Total   int              `json:"total"`
		Records []rlib.GLAccount `json:"records"`
	}

	i := 0
	for rows.Next() {
		var p rlib.GLAccount
		rlib.ReadGLAccounts(rows, &p)
		p.Recid = i
		g.Records = append(g.Records, p)
		i++
	}

	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Total = len(g.Records) // this is also the value of i, but what we're doing should be clear
	g.Status = "success"
	b, err := json.Marshal(g)
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	w.Write(b)
}
