package ws

import (
	"context"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

// LedgerGrid is a structure specifically for the UI Grid.
type LedgerGrid struct {
	Recid     int64 `json:"recid"` // this is to support the w2ui form
	LID       int64
	GLNumber  string
	Name      string
	Active    string
	AllowPost string
	Balance   float64
	LMDate    string
	LMAmount  float64
	LMState   string
}

// SearchLedgersResponse is a response string to the search request for receipts
type SearchLedgersResponse struct {
	Status  string       `json:"status"`
	Total   int64        `json:"total"`
	Records []LedgerGrid `json:"records"`
}

// GetLedgerResponse is the response to a GetAR request
type GetLedgerResponse struct {
	Status string     `json:"status"`
	Record ARSendForm `json:"record"`
}

// SvcSearchHandlerLedger generates a report of all ARs defined business d.BID
// wsdoc {
//  @Title  Search Account Rules
//	@URL /v1/ars/:BUI
//  @Method  POST
//	@Synopsis Search Account Rules
//  @Description  Search all ARs and return those that match the Search Logic.
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchLedgersResponse
// wsdoc }
func SvcSearchHandlerLedger(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerLedger"
	fmt.Printf("Entered %s\n", funcname)

	switch d.wsSearchReq.Cmd {
	case "get":
		getLedgerGrid(w, r, d)
		break
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// GetAccountBalance returns the balance of the account at time dt
//
func GetAccountBalance(ctx context.Context, bid, lid int64, dt *time.Time) (float64, rlib.LedgerMarker) {
	var bal float64
	lm, err := rlib.GetRALedgerMarkerOnOrBeforeDeprecated(ctx, bid, lid, 0, dt) // find nearest ledgermarker, use it as a starting point
	if err != nil {
		return bal, lm
	}

	bal, _ = rlib.GetAccountActivity(ctx, bid, lid, &lm.Dt, dt)
	return bal, lm
}

// LMStates is an array of strings describing the meaning of the states a Ledger Marker can have.
var LMStates = []string{
	"open", "closed", "locked", "initial",
}

// getLedgerGrid returns a list of ARs for w2ui grid
// wsdoc {
//  @Title  list ARs
//	@URL /v1/ars/:BUI
//  @Method  GET
//	@Synopsis Get Account Rules
//  @Description  Get all ARs associated with BID
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchLedgersResponse
// wsdoc }
func getLedgerGrid(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getLedgerGrid"
	var (
		err error
		g   SearchLedgersResponse
	)

	rows, err := rlib.RRdb.Prepstmt.GetLedgersForGrid.Query(d.BID, d.wsSearchReq.Limit, d.wsSearchReq.Offset)
	if err != nil {
		fmt.Printf("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	dt := time.Time(d.wsSearchReq.SearchDtStart)
	i := int64(d.wsSearchReq.Offset)
	for rows.Next() {
		var acct rlib.GLAccount
		err = rlib.ReadGLAccounts(rows, &acct)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		active := "active"
		if 1 == acct.Status {
			active = "inactive"
		}
		posts := "yes"
		if acct.AllowPost == 0 {
			posts = "no"
		}

		bal, lm := GetAccountBalance(r.Context(), acct.BID, acct.LID, &dt)

		state := "??"
		j := int(lm.State)
		if 0 <= j && j <= 3 {
			state = LMStates[j]
		}

		var lg = LedgerGrid{
			Recid:     i,
			LID:       acct.LID,
			GLNumber:  acct.GLNumber,
			Name:      acct.Name,
			Active:    active,
			AllowPost: posts,
			Balance:   bal,
			LMDate:    lm.Dt.In(rlib.RRdb.Zone).Format("Jan _2, 2006 15:04:05 MST"),
			LMAmount:  lm.Balance,
			LMState:   state,
		}

		g.Records = append(g.Records, lg)
		i++
	}

	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	g.Total = int64(len(g.Records))
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}

// // SvcFormHandlerAR formats a complete data record for a person suitable for use with the w2ui Form
// // For this call, we expect the URI to contain the BID and the ARID as follows:
// //           0    1     2   3
// // uri 		/v1/receipt/BUI/ARID
// // The server command can be:
// //      get
// //      save
// //      delete
// //-----------------------------------------------------------------------------------
// func SvcFormHandlerAR(w http.ResponseWriter, r *http.Request, d *ServiceData) {
// 	var (
// 		funcname = "SvcFormHandlerAR"
// 		err      error
// 	)
// 	fmt.Printf("Entered %s\n", funcname)
// 	if d.ARID, err = SvcExtractIDFromURI(r.RequestURI, "ARID", 3, w); err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	fmt.Printf("Request: %s:  BID = %d,  ID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ARID)

// 	switch d.wsSearchReq.Cmd {
// 	case "get":
// 		getARForm(w, r, d)
// 		break
// 	case "save":
// 		saveARForm(w, r, d)
// 		break
// 	case "delete":
// 		deleteARForm(w, r, d)
// 		break
// 	default:
// 		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}
// }

// // saveARForm returns the requested receipt
// // wsdoc {
// //  @Title  Save AR
// //	@URL /v1/ars/:BUI/:ARID
// //  @Method  GET
// //	@Synopsis Save a AR
// //  @Desc  This service saves a AR.  If :ARID exists, it will
// //  @Desc  be updated with the information supplied. All fields must
// //  @Desc  be supplied. If ARID is 0, then a new receipt is created.
// //	@Input SaveARInput
// //  @Response SvcStatusResponse
// // wsdoc }
// func saveARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {

// 	var (
// 		funcname = "saveARForm"
// 		foo      SaveARInput
// 		bar      SaveAROther
// 		a        rlib.AR
// 		err      error
// 	)

// 	fmt.Printf("Entered %s\n", funcname)
// 	fmt.Printf("record data = %s\n", d.data)

// 	// get data
// 	data := []byte(d.data)

// 	if err := json.Unmarshal(data, &foo); err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	if err := json.Unmarshal(data, &bar); err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	// migrate foo.Record data to a struct's fields
// 	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
// 	fmt.Printf("saveAR - first migrate: a = %#v\n", a)

// 	var ok bool
// 	a.BID, ok = rlib.RRdb.BUDlist[bar.Record.BID.ID]
// 	if !ok {
// 		e := fmt.Errorf("%s: Could not map BID value: %s", funcname, bar.Record.BID.ID)
// 		SvcErrorReturn(w, e, funcname)
// 		return
// 	}

// 	a.CreditLID, ok = rlib.StringToInt64(bar.Record.CreditLID.ID) // CreditLID has drop list
// 	if !ok {
// 		e := fmt.Errorf("%s: invalid CreditLID value: %s", funcname, bar.Record.CreditLID.ID)
// 		SvcErrorReturn(w, e, funcname)
// 		return
// 	}

// 	a.DebitLID, ok = rlib.StringToInt64(bar.Record.DebitLID.ID) // DebitLID has drop list
// 	if !ok {
// 		e := fmt.Errorf("%s: invalid DebitLID value: %s", funcname, bar.Record.DebitLID.ID)
// 		SvcErrorReturn(w, e, funcname)
// 		return
// 	}

// 	a.ARType, ok = rlib.StringToInt64(bar.Record.ARType.ID) // ArType has drop list
// 	if !ok {
// 		e := fmt.Errorf("%s: Invalid ARType value: %s", funcname, bar.Record.ARType.ID)
// 		SvcErrorReturn(w, e, funcname)
// 		return
// 	}
// 	fmt.Printf("saveAR - second migrate: a = %#v\n", a)

// 	// get PriorToRAStart and PriorToRAStop values and accordingly get RARequired field value
// 	formBoolMap := [2]bool{foo.Record.PriorToRAStart, foo.Record.PriorToRAStop}
// 	for raReq, boolMap := range raRequiredMap {
// 		if boolMap == formBoolMap {
// 			a.RARequired = int64(raReq)
// 			break
// 		}
// 	}

// 	// save or update
// 	if a.ARID == 0 && d.ARID == 0 {
// 		// This is a new AR
// 		fmt.Printf(">>>> NEW RECEIPT IS BEING ADDED\n")
// 		_, err = rlib.InsertAR(&a)
// 	} else {
// 		// update existing record
// 		fmt.Printf("Updating existing AR: %d\n", a.ARID)
// 		err = rlib.UpdateAR(&a)
// 	}
// 	if err != nil {
// 		e := fmt.Errorf("Error saving receipt (ARID=%d\n: %s", d.ARID, err.Error())
// 		SvcErrorReturn(w, e, funcname)
// 		return
// 	}

// 	SvcWriteSuccessResponseWithID(d.BID,w, a.ARID)
// }

// // which fields needs to be fetched for SQL query for receipts grid
// var getARQuerySelectFields = rlib.SelectQueryFields{
// 	"AR.ARID",
// 	"AR.Name",
// 	"AR.ARType",
// 	"AR.DebitLID",
// 	"debitQuery.Name as DebitLedgerName",
// 	"AR.CreditLID",
// 	"creditQuery.Name as CreditLedgerName",
// 	"AR.Description",
// 	"AR.DtStart",
// 	"AR.DtStop",
// 	"AR.RARequired",
// }

// // for what RARequired value, prior and after value are
// var raRequiredMap = map[int][2]bool{
// 	0: {false, false}, // during
// 	1: {true, false},  // prior or during
// 	2: {false, true},  // after or during
// 	3: {true, true},   // after or during or prior
// }

// // getARForm returns the requested ars
// // wsdoc {
// //  @Title  Get AR
// //	@URL /v1/ars/:BUI/:ARID
// //  @Method  GET
// //	@Synopsis Get information on a AR
// //  @Description  Return all fields for ars :ARID
// //	@Input WebGridSearchRequest
// //  @Response GetLedgerResponse
// // wsdoc }
// func getARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {

// 	var (
// 		funcname = "getARForm"
// 		g        GetLedgerResponse
// 		err      error
// 	)
// 	fmt.Printf("entered %s\n", funcname)

// 	arQuery := `
// 	SELECT
// 		{{.SelectClause}}
// 	FROM AR
// 	INNER JOIN GLAccount as debitQuery on AR.DebitLID=debitQuery.LID
// 	INNER JOIN GLAccount as creditQuery on AR.CreditLID=creditQuery.LID
// 	WHERE {{.WhereClause}};`

// 	qc := rlib.QueryClause{
// 		"SelectClause": strings.Join(getARQuerySelectFields, ","),
// 		"WhereClause":  fmt.Sprintf("AR.BID=%d AND AR.ARID=%d", d.BID, d.ARID),
// 	}

// 	// get formatted query with substitution of select, where, order clause
// 	q := rlib.RenderSQLQuery(arQuery, qc)
// 	fmt.Printf("db query = %s\n", q)

// 	// execute the query
// 	rows, err := rlib.RRdb.Dbrr.Query(q)
// 	if err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var gg ARSendForm

// 		gg.BID = rlib.GetBUDFromBIDList(d.BID)

// 		err = rows.Scan(&gg.ARID, &gg.Name, &gg.ARType, &gg.DebitLID, &gg.DebitLedgerName, &gg.CreditLID, &gg.CreditLedgerName, &gg.Description, &gg.DtStart, &gg.DtStop, &gg.raRequired)
// 		if err != nil {
// 			SvcErrorReturn(w, err, funcname)
// 			return
// 		}

// 		// according to RARequired map, fill out PriorToRAStart, PriorToRAStop values
// 		raReqMappedVal := raRequiredMap[gg.raRequired]
// 		gg.PriorToRAStart = raReqMappedVal[0]
// 		gg.PriorToRAStop = raReqMappedVal[1]
// 		g.Record = gg
// 	}

// 	// error check
// 	err = rows.Err()
// 	if err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	g.Status = "success"
// 	w.Header().Set("Content-Type", "application/json")
// 	SvcWriteResponse(d.BID,&g, w)
// }

// // deleteAR request delete AR from database
// // wsdoc {
// //  @Title  Delete AR
// //	@URL /v1/ars/:BUI/:ARID
// //  @Method  DELETE
// //	@Synopsis Delete record for a AR
// //  @Description  Delete record from database ars :ARID
// //	@Input WebGridSearchRequest
// //  @Response SvcWriteSuccessResponse
// // wsdoc }
// func deleteARForm(w http.ResponseWriter, r *http.Request, d *ServiceData) {
// 	var (
// 		funcname = "deleteARForm"
// 		del      DeleteARForm
// 	)

// 	fmt.Printf("Entered %s\n", funcname)
// 	fmt.Printf("record data = %s\n", d.data)

// 	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	if err := rlib.DeleteAR(del.ARID); err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	SvcWriteSuccessResponse(d.BID,w)
// }
