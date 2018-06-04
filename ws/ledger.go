package ws

import (
	"context"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"time"
)

//-------------------------------------------------------------------
//                        **** SEARCH ****
//-------------------------------------------------------------------

// no search area here because there is no main grid

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

//-------------------------------------------------------------------
//                         **** SAVE ****
//-------------------------------------------------------------------

// SearchLedgersResponse is a response string to the search request for receipts
type SearchLedgersResponse struct {
	Status  string       `json:"status"`
	Total   int64        `json:"total"`
	Records []LedgerGrid `json:"records"`
}

//-------------------------------------------------------------------
//                         **** GET ****
//-------------------------------------------------------------------

// GetLedgerResponse is the response to a GetAR request
type GetLedgerResponse struct {
	Status string     `json:"status"`
	Record ARSendForm `json:"record"`
}

// SvcLedgerHandler generates a report of all ARs defined business d.BID
// wsdoc {
//  @Title  Search Account Rules
//	@URL /v1/ledgers/:BUI
//  @Method  POST
//	@Synopsis Search Account Rules
//  @Description  Search all ARs and return those that match the Search Logic.
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchLedgersResponse
// wsdoc }
func SvcLedgerHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcLedgerHandler"
	fmt.Printf("Entered %s\n", funcname)

	if d.Service == "ledgers" {
		switch d.wsSearchReq.Cmd {
		case "get":
			searchLedgers(w, r, d)
			break
		default:
			err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
			SvcErrorReturn(w, err, funcname)
			return
		}
	}
	if d.Service == "ledger" {
		switch d.wsSearchReq.Cmd {
		case "get":
		case "save":
		case "delete":
		default:
		}
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

// searchLedgers returns a list of Ledgers
// wsdoc {
//  @Title  list ARs
//	@URL /v1/ledgers/:BUI
//  @Method  GET
//	@Synopsis Get Account Rules
//  @Description  Get all ARs associated with BID
//  @Desc By default, the search is made for receipts from "today" to 31 days prior.
//	@Input WebGridSearchRequest
//  @Response SearchLedgersResponse
// wsdoc }
func searchLedgers(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "searchLedgers"
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

	if err = rows.Err(); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	g.Total = int64(len(g.Records))
	SvcWriteResponse(d.BID, &g, w)
}
