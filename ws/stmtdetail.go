package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"rentroll/rrpt"
	"time"
)

// StatementDetail is a structure to fill the statement detail grid
type StatementDetail struct {
	Recid      int64         `json:"recid"` // this is to support the w2ui form
	RAID       int64         // internal unique id
	BID        int64         // Business (so that we can process by Business)
	BUD        rlib.XJSONBud // which business
	ID         string        // rcpt or asmt id
	Dt         rlib.JSONDate // date of the assessment or payment
	Descr      string        // about the assessment/receipt
	AsmtAmount float64       // amount of assessment
	RcptAmount float64       // amount of receipt
	Balance    float64       // sum
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
	var g StmtDetailResponse
	d1 := time.Now()
	d2 := d1.AddDate(0, 1, 0)
	m := rrpt.GetStatementData(sd.BID, sd.ID, &d1, &d2)
	var b = rlib.RoundToCent(m[0].Amt) // element 0 is always the account balance
	var c = float64(0)                 // credit
	var d = float64(0)                 // debit

	for i := 0; i < len(m); i++ {
		bud, err := bidToBud(sd.BID)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		var a = StatementDetail{
			BID:  sd.BID,
			BUD:  rlib.XJSONBud(bud),
			RAID: sd.ID,
			Dt:   rlib.JSONDate(m[i].Dt),
		}

		descr := ""
		if m[i].T == 1 || m[i].T == 2 {
			if m[i].A.ARID > 0 {
				descr = rlib.RRdb.BizTypes[sd.BID].AR[m[i].A.ARID].Name
			} else {
				descr = rlib.RRdb.BizTypes[sd.BID].GLAccounts[m[i].A.ATypeLID].Name
			}
		}
		switch m[i].T {
		case 1: // assessments
			amt := rlib.RoundToCent(m[i].Amt)
			c += amt
			b += amt
			a.Dt = rlib.JSONDate(m[i].Dt)
			a.ID = rlib.IDtoString("ASM", m[i].ID)
			a.Descr = descr
			a.AsmtAmount = amt
		case 2: // receipts
			amt := rlib.RoundToCent(m[i].Amt)
			d += amt
			b += amt
			if m[i].A.ASMID > 0 {
				descr = fmt.Sprintf("%s (%s)", descr, m[i].A.IDtoString())
			}
			a.ID = rlib.IDtoString("RCPT", m[i].ID)
			a.Descr = descr
			a.RcptAmount = amt
		case 3: // opening balance
			a.Descr = "Opening Balance"
		}
		a.Balance = b
		g.Records = append(g.Records, a)
	}
	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	g.Total = int64(len(g.Records))
	SvcWriteResponse(&g, w)

}
