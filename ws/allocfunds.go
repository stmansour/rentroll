package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

// UnallocatedReceiptsPayors individual records to be shown on grid
type UnallocatedReceiptsPayors struct {
	Recid int `json:"recid"`
	TCID  int64
	Name  string
	BID   int64
}

// SearchAllocFundsResponse used to send down the list of records on allocated funds grid
type SearchAllocFundsResponse struct {
	Status  string                      `json:"status"`
	Total   int64                       `json:"total"`
	Records []UnallocatedReceiptsPayors `json:"records"`
}

// UnpaidAsm unpaid assessments of a payor
type UnpaidAsm struct {
	Recid      int              `json:"recid"`
	DtStart    rlib.JSONTime    `json:"Date"`
	ASMID      int64            `json:"ASMID"`
	ARID       int64            `json:"ARID"`
	Name       string           `json:"Assessment"`
	Amount     float64          `json:"Amount"`
	AmountPaid float64          `json:"AmountPaid"`
	Allocate   rlib.NullFloat64 `json:"Allocate"`
}

// PayorUnpaidAsmsResponse used to send down the list of unpaid assessments
type PayorUnpaidAsmsResponse struct {
	Status  string      `json:"status"`
	Total   int64       `json:"total"`
	Records []UnpaidAsm `json:"records"`
	// Time    rlib.JSONTime `json:"time"`
}

// PayorFund is used to get total unallocated fund for a payor
type PayorFund struct {
	Fund float64 `json:"fund"`
}

// PayorFundResponse response of payor fund request
type PayorFundResponse struct {
	Status string    `json:"status"`
	Record PayorFund `json:"record"`
}

// AllocFundSaveRequest used to allocate fund to unpaid assessments
type AllocFundSaveRequest struct {
	// Time    rlib.JSONTime `json:"time"`
	TCID    int64
	BID     int64
	Records []UnpaidAsm `json:"records"`
}

// SvcSearchHandlerAllocFunds formats a complete data record for a alloc funds suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the TCID as follows:
//           0    1     2   3
// uri      /v1/allocfunds/BUI/TCID
// The server command can be:
//      get
//      save
//-----------------------------------------------------------------------------------
func SvcSearchHandlerAllocFunds(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcSearchHandlerAllocFunds"
	)

	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d\n", d.wsSearchReq.Cmd, d.BID)

	switch d.wsSearchReq.Cmd {
	case "get":
		getUnallocFundPayors(w, r, d)
		break
	case "save":
		allocatePayorFund(w, r, d)
		break
	default:
		err := fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcGridErrorReturn(w, err, funcname)
		return
	}
}

// getUnallocFundPayors generates a list of all payors who has unallocated receipts
// wsdoc {
//  @Title  Search Payors with unallocated receipts
//  @URL /v1/allocfunds/:BUI
//  @Method  GET
//  @Synopsis Return a list of payors with Unallocated receipts
//  @Description This service returns a list of Payors with unallocated receipts
//  @Input WebGridSearchRequest
//  @Response SearchAllocFundsResponse
// wsdoc }
func getUnallocFundPayors(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "getUnallocFundPayors"
		err      error
		g        SearchAllocFundsResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	rows, err := rlib.RRdb.Prepstmt.GetUnallocatedReceipts.Query(d.BID)
	defer rows.Close()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	i := d.wsSearchReq.Offset
	count := 0

	u := map[int64]rlib.Transactant{}

	// get the transactants list
	for rows.Next() {
		// get receipts record
		var r rlib.Receipt
		rlib.ReadReceipts(rows, &r)

		// get Transactant records
		var t rlib.Transactant
		err = rlib.GetTransactant(r.TCID, &t)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		u[r.TCID] = t
	}
	err = rows.Err()
	if err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	now := time.Now()
	for _, t := range u {
		var q UnallocatedReceiptsPayors

		// We know that the payor has funds in one or more receipts that has
		// not yet been allocated. However, there may not be any assessments
		// toward which we can apply the funds.  If this is the case, then just
		// skip to the next payor
		m := bizlogic.GetAllUnpaidAssessmentsForPayor(t.BID, t.TCID, &now)
		if len(m) == 0 { // no assessments to pay? ...
			continue // then move on to the next payor
		}

		q.Recid = i
		q.TCID = t.TCID
		q.BID = t.BID
		q.Name = t.FirstName + " " + t.LastName

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}

	g.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&g, w)
}

// allocatePayorFund allocates user fund to unpaid assessments
// wsdoc {
//  @Title  Allocate Payor fund
//  @URL /v1/allocfunds/:BUI
//  @Method  POST
//  @Synopsis Allocate fund of payor to unpaid assessments
//  @Description This service save the unallocate fund to unpaid assessments
//  @Input AllocPayorFundRequest
//  @Response SvcStatusResponse
// wsdoc }
func allocatePayorFund(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var (
		funcname = "allocatePayorFund"
		err      error
		foo      AllocFundSaveRequest
	)

	fmt.Printf("Entered %s\n", funcname)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		SvcGridErrorReturn(w, err, funcname)
		return
	}

	fmt.Printf("Began to allocate funds for TCID=%d, BID=%d\n", foo.TCID, foo.BID)

	// Need to init some internals for Business
	var xbiz rlib.XBusiness
	rlib.InitBizInternals(foo.BID, &xbiz)

	dt := time.Now()

	// get receipts for payor
	n := rlib.GetUnallocatedReceiptsByPayor(foo.BID, foo.TCID)
	fmt.Printf("number of unallocated receipts: %d\n", len(n))

	for _, asmRec := range foo.Records {

		// This is how much the user wanted to allocate for this assessment...
		amt := asmRec.Allocate.Float64

		// The user may have decided not to pay anything here. If so, skip to the next assessment.
		if amt == float64(0) {
			continue
		}

		asm, err := rlib.GetAssessment(asmRec.ASMID)
		if err != nil {
			SvcGridErrorReturn(w, err, funcname)
			return
		}

		needed := bizlogic.AssessmentUnpaidPortion(&asm)
		fmt.Printf("ASMID = %d, Requested Amount = %.2f, AR = %d\n", asm.ASMID, amt, asm.ARID)

		for j := 0; j < len(n); j++ {
			fmt.Printf("Entered for Receipt: %d\n", n[j].RCPTID)
			if n[j].FLAGS&3 == 2 { // if there are no funds left in this receipt...
				continue // move on to the next receipt
			}

			err := bizlogic.PayAssessment(&asm, &n[j], &needed, &amt, &dt)
			fmt.Printf("Applied %.2f to ASMID: %d.  Amount still owed: %.2f\n", amt, asm.ASMID, needed)
			if err != nil {
				SvcGridErrorReturn(w, err, funcname)
				return
			}
			if amt < bizlogic.ROUNDINGERR { // if we've applied the requested amount...
				break // ... then break out of the loop; we're done
			}
		}
	}

	SvcWriteSuccessResponse(w)
}

// SvcHandlerGetUnpaidAsms generates a list of all unpaid assessments for a payor
// wsdoc {
//  @Title  Get Unpaid Assessments for a payor
//  @URL /v1/unpaidasms/:BUI/TCID
//  @Method  GET
//  @Synopsis Return a list of unpaid assessments of a payor
//  @Description This service returns a list of unpaid assessments of a payor
//  @Input PayorUnpaidAsmsRequest
//  @Response PayorUnpaidAsmsResponse
// wsdoc }
func SvcHandlerGetUnpaidAsms(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcHandlerGetUnpaidAsms"
		res      PayorUnpaidAsmsResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	TCID := d.ID
	dt := time.Now()
	// res.Time = rlib.JSONTime(dt) // let's keep it here as of now
	m := bizlogic.GetAllUnpaidAssessmentsForPayor(d.BID, TCID, &dt)

	for i, asm := range m {
		var rec UnpaidAsm
		rec.Recid = i
		rec.DtStart = rlib.JSONTime(asm.Start)
		rec.Amount = asm.Amount
		rec.AmountPaid = asm.Amount - bizlogic.AssessmentUnpaidPortion(&m[i])
		ar, err := rlib.GetAR(asm.ARID)
		if err != nil {
			fmt.Printf("%s: Error while getting AR (ARID=%d) for Assessment: %d, error=<%s>\n", funcname, asm.ARID, asm.ASMID, err.Error())
			SvcGridErrorReturn(w, err, funcname)
			return
		}
		rec.Name = ar.Name
		rec.ASMID = asm.ASMID
		rec.ARID = asm.ARID
		res.Records = append(res.Records, rec)
		res.Total++
	}

	res.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&res, w)
}

// SvcHandlerTotalUnallocFund generates a total unallocated fund for a payor from receipts
// wsdoc {
//  @Title  Get total unallocated fund
//  @URL /v1/payorfund/:BUI/TCID
//  @Method  GET
//  @Synopsis Return a total amount of unallocated fund
//  @Description This service returns the total amount of unallocated fund
//  @Input GridWebSearchRequest
//  @Response PayorFundResponse
// wsdoc }
func SvcHandlerTotalUnallocFund(w http.ResponseWriter, r *http.Request, d *ServiceData) {

	var (
		funcname = "SvcHandlerGetUnpaidAsms"
		res      PayorFundResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	TCID := d.ID
	m := rlib.GetUnallocatedReceiptsByPayor(d.BID, TCID)

	for _, rcpt := range m {
		if rcpt.FLAGS&3 == 2 { // if there are no funds left in this receipt...
			continue // move on to the next receipt
		}

		amt := bizlogic.RemainingReceiptFunds(&rcpt)
		res.Record.Fund += amt
	}

	res.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(&res, w)
}
