package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strconv"
	"strings"
)

// DepositListGrid contains the data from Deposit that is targeted to the UI Grid that displays
// a list of Deposit structs
type DepositListGrid struct {
	Recid         int64 `json:"recid"`
	DPID          int64
	DID           int64
	BID           int64
	BUD           rlib.XJSONBud
	Dt            rlib.JSONDate
	RCPTID        int64
	DocNo         string
	TCID          int64
	PMTID         int64
	Payors        string
	Amount        float64
	ClearedAmount float64
	FLAGS         uint64
	PMTName       string
	Check         bool
	LastModTime   rlib.JSONDateTime
	LastModBy     int64
	CreateTS      rlib.JSONDateTime
	CreateBy      int64
}

// DepositListSearchResponse is a response string to the search request for Deposit records
type DepositListSearchResponse struct {
	Status  string            `json:"status"`
	Total   int64             `json:"total"`
	Records []DepositListGrid `json:"records"`
}

var depositListSearchFieldMap = rlib.SelectQueryFieldMap{
	"DPID":    {"DepositPart.DPID"},
	"DID":     {"DepositPart.DID"},
	"BID":     {"DepositPart.BID"},
	"RCPTID":  {"DepositPart.RCPTID"},
	"Dt":      {"Receipt.Dt"},
	"Amount":  {"Receipt.Amount"},
	"TCID":    {"Receipt.TCID"},
	"DocNo":   {"Receipt.DocNo"},
	"PMTID":   {"Receipt.PMTID"},
	"PMTName": {"PaymentType.Name"},
	//	"Payors":        {"Transactant.FirstName", "Transactant.LastName", "Transactant.CompanyName"},
	"ClearedAmount": {"Deposit.ClearedAmount"},
	"FLAGS":         {"Deposit.FLAGS"},
	"LastModTime":   {"Deposit.LastModTime"},
	"LastModBy":     {"Deposit.LastModBy"},
	"CreateTS":      {"Deposit.CreateTS"},
	"CreateBy":      {"Deposit.CreateBy"},
}

// which fields needs to be fetch to satisfy the struct
var qfields = rlib.SelectQueryFields{
	"DepositPart.DPID",
	"DepositPart.DID",
	"DepositPart.BID",
	"DepositPart.RCPTID",
	"Receipt.Dt",
	"Receipt.Amount",
	"Receipt.TCID",
	"Receipt.DocNo",
	"Receipt.PMTID",
	"PaymentType.Name",
	//	"GROUP_CONCAT(DISTINCT CASE WHEN Transactant.IsCompany > 0 THEN Transactant.CompanyName ELSE CONCAT(Transactant.FirstName, ' ', Transactant.LastName) END SEPARATOR ', ') AS Payors",
	"Deposit.ClearedAmount",
	"Deposit.FLAGS",
	"Deposit.LastModTime",
	"Deposit.LastModBy",
	"Deposit.CreateTS",
	"Deposit.CreateBy",
}

// SvcHandlerDepositList handle web service calls for depositlist...
//-----------------------------------------------------------------------------------
func SvcHandlerDepositList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcHandlerDepositList"
	var err error
	rlib.Console("Entered %s\n", funcname)
	rlib.Console("Request: %s:  BID = %d,  DID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID > 0 && d.wsSearchReq.Limit > 0 {
			SvcDepositReceiptsAndUndeposited(w, r, d) // it is a query for the grid.
			return
		}
		if d.ID <= 0 && d.wsSearchReq.Limit > 0 {
			SvcUndepositedReceiptList(w, r, d) // it is a query for the grid.
			return
		}
	// case "save":
	// 	saveDeposit(w, r, d)
	// 	break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// // depositListGridRowScan scans a result from sql row and dump it in a PrARGrid struct
// func depositListGridRowScan(rows *sql.Rows, a *DepositListGrid) error {
// 	err := rows.Scan(&a.DPID, &a.DID, &a.BID, &a.RCPTID, &a.Dt, &a.Amount, &a.TCID, &a.DocNo, &a.PMTID, &a.PMTName /*&a.Payors,*/, &a.ClearedAmount, &a.FLAGS, &a.LastModTime, &a.LastModBy, &a.CreateTS, &a.CreateBy)
// 	return err
// }

// SvcUndepositedReceiptList returns the list of possible receipts that can be included
// as part of a new deposit.
// wsdoc {
//  @Title  Undeposited Receipt List
//	@URL /v1/depositlist/:BUI[/0]
//  @Method  POST
//	@Synopsis Return a list of Receipts that have not been deposited
//  @Descr  If d.ID == 0 or is -1 because no ID was supplied then
//  @Descr  the return list will be all the undeposited receipts.
//	@Input WebGridSearchRequest
//  @Response DepositListSearchResponse
// wsdoc }
func SvcUndepositedReceiptList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcUndepositedReceiptList"
		g        DepositListSearchResponse
		err      error
	)

	rlib.Console("Entered %s\n", funcname)

	whr := `Receipt.BID=%d AND Receipt.DID=0 AND Receipt.FLAGS & 4 = 0 AND Receipt.Dt >= %q and Receipt.Dt < %q`
	whr = fmt.Sprintf(whr, d.BID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := "Receipt.Dt ASC, Receipt.RCPTID ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, receiptsFieldsMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	receiptsQuery := `
	SELECT {{.SelectClause}}
	FROM Receipt
	LEFT JOIN Transactant ON Receipt.TCID=Transactant.TCID
	LEFT JOIN AR ON Receipt.ARID=AR.ARID
	LEFT JOIN PaymentType ON Receipt.PMTID=PaymentType.PMTID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(receiptsQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(receiptsQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		fmt.Printf("Error from rlib.GetQueryCount: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	receiptsQueryWithLimit := receiptsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(receiptsQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		fmt.Printf("Error from DB Query: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q PrReceiptGrid
		var dlg DepositListGrid

		q, err = receiptsGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		rlib.MigrateStructVals(&q, &dlg)
		dlg.Recid = i
		dlg.PMTName = q.PmtTypeName
		if q.Payor.Valid {
			dlg.Payors = q.Payor.String
		}

		g.Records = append(g.Records, dlg)

		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}

	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)

}

// SvcDepositReceiptsAndUndeposited returns the list of receipts associated with a deposit. If
// the DID
// wsdoc {
//  @Title  Deposit Receipts
//	@URL /v1/depositlist/:BUI/DID
//  @Method  POST
//	@Synopsis Return a Deposit's Receipts and all undeposited receipts
//  @Descr  If d.ID > 0 then the return list will be the list of receipts
//  @Descr  associated with deposit DID.  If DID == 0 or not supplied then
//  @Descr  the call should be made to SvcUndepositedReceiptList.
//	@Input WebGridSearchRequest
//  @Response DepositListSearchResponse
// wsdoc }
func SvcDepositReceiptsAndUndeposited(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "SvcDepositReceiptsAndUndeposited"
		g        DepositListSearchResponse
		err      error
		// order    = "Receipt.DID DESC,Receipt.Dt ASC" // default ORDER
		// whr      = fmt.Sprintf("DepositPart.BID=%d AND Receipt.BID=%d AND Receipt.FLAGS & 4=0 AND (DepositPart.DID=%d OR Receipt.DID=0)", d.BID, d.BID, d.ID)
		//whr      = fmt.Sprintf("DepositPart.DID=%d", d.ID)
	)

	rlib.Console("Entered %s\n", funcname)

	whr := `Receipt.BID=%d AND (Receipt.DID=%d OR (Receipt.DID=0 AND Receipt.FLAGS & 4 = 0 AND Receipt.Dt >= %q and Receipt.Dt < %q))`
	whr = fmt.Sprintf(whr, d.BID, d.ID, d.wsSearchReq.SearchDtStart.Format(rlib.RRDATEFMTSQL), d.wsSearchReq.SearchDtStop.Format(rlib.RRDATEFMTSQL))
	order := "Receipt.DID DESC,Receipt.Dt ASC" // default ORDER

	// get where clause and order clause for sql query
	whereClause, orderClause := GetSearchAndSortSQL(d, receiptsFieldsMap)
	if len(whereClause) > 0 {
		whr += " AND (" + whereClause + ")"
	}
	if len(orderClause) > 0 {
		order = orderClause
	}

	receiptsQuery := `
	SELECT {{.SelectClause}}
	FROM Receipt
	LEFT JOIN Transactant ON Receipt.TCID=Transactant.TCID
	LEFT JOIN AR ON Receipt.ARID=AR.ARID
	LEFT JOIN PaymentType ON Receipt.PMTID=PaymentType.PMTID
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}}`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(receiptsQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get TOTAL COUNT First
	countQuery := rlib.RenderSQLQuery(receiptsQuery, qc)
	g.Total, err = rlib.GetQueryCount(countQuery)
	if err != nil {
		fmt.Printf("Error from rlib.GetQueryCount: %s\n", err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("g.Total = %d\n", g.Total)

	// FETCH the records WITH LIMIT AND OFFSET
	// limit the records to fetch from server, page by page
	limitAndOffsetClause := `
	LIMIT {{.LimitClause}}
	OFFSET {{.OffsetClause}};`

	// build query with limit and offset clause
	// if query ends with ';' then remove it
	receiptsQueryWithLimit := receiptsQuery + limitAndOffsetClause

	// Add limit and offset value
	qc["LimitClause"] = strconv.Itoa(d.wsSearchReq.Limit)
	qc["OffsetClause"] = strconv.Itoa(d.wsSearchReq.Offset)

	// get formatted query with substitution of select, where, order clause
	qry := rlib.RenderSQLQuery(receiptsQueryWithLimit, qc)
	rlib.Console("db query = %s\n", qry)

	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		rlib.Console("%s: Error from DB Query: %s\n", funcname, err.Error())
		SvcErrorReturn(w, err, funcname)
		return
	}
	defer rows.Close()

	i := int64(d.wsSearchReq.Offset)
	count := 0
	for rows.Next() {
		var q PrReceiptGrid
		var dlg DepositListGrid

		q, err = receiptsGridRowScan(rows, q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		rlib.MigrateStructVals(&q, &dlg)
		dlg.Recid = i
		dlg.BID = d.BID
		dlg.BUD = rlib.GetBUDFromBIDList(q.BID)
		dlg.Check = q.DID == d.ID
		g.Records = append(g.Records, dlg)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}

	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
}
