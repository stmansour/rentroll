package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
)

// StatementDetail is a structure to fill the statement detail grid
type StatementDetail struct {
	Recid        int64         `json:"recid"` // this is to support the w2ui form
	RAID         int64         // internal unique id
	BID          int64         // Business (so that we can process by Business)
	BUD          rlib.XJSONBud // which business
	ID           string        // rcpt or asmt id
	Dt           rlib.JSONDate // date of the assessment or payment
	Descr        string        // about the assessment/receipt
	AsmtAmount   float64       // amount of assessment
	RcptAmount   float64       // amount of receipt
	RentableName string        // associated rentable name
	Balance      float64       // sum
}

// StmtDetailResponse is the response data for a Rental Agreement Search
type StmtDetailResponse struct {
	Status  string            `json:"status"`
	Total   int64             `json:"total"`
	Records []StatementDetail `json:"records"`
}

// SvcStatementDetail is the response data for a Stmt Grid search
// wsdoc {
//  @Title  Statement Detail
//	@URL /v1/stmtDetail/:BUI/:RAID
//  @Method  POST
//	@Synopsis Returns account details for the supplied RAID in the date range
//  @Description  Returns the assessments and receipts for the time range
//	@Input WebGridSearchRequest
//  @Response StmtDetailResponse
// wsdoc }
func SvcStatementDetail(w http.ResponseWriter, r *http.Request, sd *ServiceData) {
	funcname := "SvcStatementDetails"
	rlib.Console("Entered %s\n", funcname)
	var g StmtDetailResponse
	var xbiz rlib.XBusiness

	bud, err := bidToBud(sd.BID)
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	//
	// UGH!
	//=======================================================================
	rlib.InitBizInternals(sd.BID, &xbiz)
	rlib.Console("sd.BID = %d\n", sd.BID)
	_, ok := rlib.RRdb.BizTypes[sd.BID]
	if !ok {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d]", sd.BID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	if len(rlib.RRdb.BizTypes[sd.BID].GLAccounts) == 0 {
		e := fmt.Errorf("nothing exists in rlib.RRdb.BizTypes[%d].GLAccounts", sd.BID)
		SvcGridErrorReturn(w, e, funcname)
		return
	}
	//=======================================================================
	// UGH!

	//--------------------------------------------
	// Get the statement data...
	//--------------------------------------------
	d1 := sd.wsSearchReq.SearchDtStart
	d2 := sd.wsSearchReq.SearchDtStop
	m, err := rlib.GetRAIDStatementInfo(sd.ID, &d1, &d2)
	if err != nil {
		// e := fmt.Errorf("GetRAIDAccountBalance returned error: %s", err.Error())
		// SvcGridErrorReturn(w, e, funcname)
		g.Total = 0
		g.Status = "success"
		SvcWriteResponse(&g, w)
		return
	}

	//--------------------------------------------
	// Set the opening balance.
	//--------------------------------------------
	var b, c, d float64
	var a = StatementDetail{
		BID:     sd.BID,
		BUD:     rlib.XJSONBud(bud),
		RAID:    sd.ID,
		Dt:      rlib.JSONDate(m.DtStart),
		Descr:   "Opening Balance",
		Balance: m.OpeningBal,
	}
	g.Records = append(g.Records, a)
	b = m.OpeningBal
	count := 0
	for i := sd.wsSearchReq.Offset; i < len(m.Stmt); i++ {

		var a = StatementDetail{
			BID:  sd.BID,
			BUD:  rlib.XJSONBud(bud),
			RAID: sd.ID,
			Dt:   rlib.JSONDate(m.Stmt[i].Dt),
		}

		descr := ""
		if m.Stmt[i].T == 1 || m.Stmt[i].T == 2 {
			if m.Stmt[i].A.ARID > 0 {
				descr = rlib.RRdb.BizTypes[sd.BID].AR[m.Stmt[i].A.ARID].Name
			} else {
				descr = rlib.RRdb.BizTypes[sd.BID].GLAccounts[m.Stmt[i].A.ATypeLID].Name
			}
		}
		switch m.Stmt[i].T {
		case 1: // assessments
			amt := m.Stmt[i].Amt
			c -= amt
			b -= amt
			a.Dt = rlib.JSONDate(m.Stmt[i].Dt)
			a.ID = rlib.IDtoShortString("ASM", m.Stmt[i].A.ASMID)
			a.Descr = descr
			a.AsmtAmount = amt
		case 2: // receipts
			amt := m.Stmt[i].Amt
			d += amt
			b += amt
			if m.Stmt[i].A.ASMID > 0 {
				descr = fmt.Sprintf("%s (%s)", descr, rlib.IDtoShortString("ASM", m.Stmt[i].A.ASMID))
			}
			a.ID = rlib.IDtoShortString("RCPT", m.Stmt[i].R.RCPTID)
			a.Descr = descr
			a.RcptAmount = amt
		}
		a.RentableName = m.Stmt[i].RNT.RentableName
		a.Balance = b
		a.Recid = int64(i)
		g.Records = append(g.Records, a)
		count++
		if count >= sd.wsSearchReq.Limit {
			break
		}
	}

	a = StatementDetail{
		BID:        sd.BID,
		BUD:        rlib.XJSONBud(bud),
		RAID:       sd.ID,
		Dt:         rlib.JSONDate(m.DtStop.AddDate(0, 0, -1)),
		Descr:      "Closing Balance",
		Balance:    m.ClosingBal,
		AsmtAmount: c,
		RcptAmount: d,
	}
	g.Records = append(g.Records, a)

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	g.Total = int64(len(g.Records))
	SvcWriteResponse(&g, w)

}
