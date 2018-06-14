package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"sort"
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
	DtStart    rlib.JSONDate    `json:"Date"`
	ASMID      int64            `json:"ASMID"`
	ARID       int64            `json:"ARID"`
	Name       string           `json:"Assessment"`
	Amount     float64          `json:"Amount"`
	AmountPaid float64          `json:"AmountPaid"`
	AmountOwed float64          `json:"AmountOwed"`
	Dt         rlib.JSONDate    `json:"Dt"`
	Allocate   rlib.NullFloat64 `json:"Allocate"`
}

// PayorUnpaidAsmsResponse used to send down the list of unpaid assessments
type PayorUnpaidAsmsResponse struct {
	Status  string      `json:"status"`
	Total   int64       `json:"total"`
	Records []UnpaidAsm `json:"records"`
	// Time    rlib.JSONDate `json:"time"`
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
	// Time    rlib.JSONDate `json:"time"`
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

	const funcname = "SvcSearchHandlerAllocFunds"

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
		SvcErrorReturn(w, err, funcname)
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
	const funcname = "getUnallocFundPayors"
	var (
		err error
		g   SearchAllocFundsResponse
	)

	rows, err := rlib.RRdb.Prepstmt.GetUnallocatedReceipts.Query(d.BID)
	defer rows.Close()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	i := d.wsSearchReq.Offset
	count := 0
	u := map[int64]rlib.Transactant{}

	var payorTCIDList rlib.Int64Range

	// get the transactants list
	for rows.Next() {
		// get receipts record
		var rcpt rlib.Receipt
		err = rlib.ReadReceipts(rows, &rcpt)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		// get Transactant records
		var t rlib.Transactant
		err = rlib.GetTransactant(r.Context(), rcpt.TCID, &t)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		u[rcpt.TCID] = t
		if !rlib.Int64InSlice(rcpt.TCID, payorTCIDList) {
			payorTCIDList = append(payorTCIDList, rcpt.TCID)
		}
	}
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// sort the order of payors TCID
	sort.Sort(payorTCIDList)

	// now := time.Now()
	for _, tcid := range payorTCIDList {
		t := u[tcid]
		var q UnallocatedReceiptsPayors

		q.Recid = i
		q.TCID = t.TCID
		q.BID = t.BID
		if t.IsCompany {
			q.Name = t.CompanyName
		} else {
			q.Name = t.FirstName + " " + t.LastName
		}

		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++ // update the index no matter what
	}

	g.Status = "success"
	g.Total = int64(len(g.Records))
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &g, w)
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
	const funcname = "allocatePayorFund"
	var (
		err error
		foo AllocFundSaveRequest
	)

	rlib.Console("Entered %s\n", funcname)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	rlib.Console("Began to allocate funds for TCID=%d, BID=%d\n", foo.TCID, foo.BID)

	// Need to init some internals for Business
	var xbiz rlib.XBusiness
	err = rlib.InitBizInternals(foo.BID, &xbiz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// dt := time.Now()

	// get receipts for payor
	n, _ := rlib.GetUnallocatedReceiptsByPayor(r.Context(), foo.BID, foo.TCID)
	rlib.Console("number of unallocated receipts: %d\n", len(n))

	for _, asmRec := range foo.Records {

		// This is how much the user wanted to allocate for this assessment...
		amt := asmRec.Allocate.Float64

		// The user may have decided not to pay anything here. If so, skip to the next assessment.
		if amt == float64(0) {
			continue
		}

		asm, err := rlib.GetAssessment(r.Context(), asmRec.ASMID)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}

		needed := bizlogic.AssessmentUnpaidPortion(r.Context(), &asm)
		rlib.Console("ASMID = %d, Requested Amount = %.2f, AR = %d\n", asm.ASMID, amt, asm.ARID)
		dt := time.Time(asmRec.Dt)
		rlib.Console("Allocation date: %s\n", dt.Format(rlib.RRDATEREPORTFMT))

		for j := 0; j < len(n); j++ {
			rlib.Console("*******************\nprocessing Receipt: %d\n", n[j].RCPTID)
			if n[j].FLAGS&3 == 2 { // if there are no funds left in this receipt...
				continue // move on to the next receipt
			}

			err := bizlogic.PayAssessment(r.Context(), &asm, &n[j], &needed, &amt, &dt)
			rlib.Console("amt = %.2f .  Amount still owed: %.2f\n", amt, needed)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			if amt < bizlogic.ROUNDINGERR { // if we've applied the requested amount...
				rlib.Console("ASMID %d is paid off, moving on to next record\n", asm.ASMID)
				break // ... then break out of the loop; we're done
			}
		}
	}

	SvcWriteSuccessResponse(d.BID, w)
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
	const funcname = "SvcHandlerGetUnpaidAsms"
	var (
		res PayorUnpaidAsmsResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	TCID := d.ID
	dt := time.Now()
	// res.Time = rlib.JSONDate(dt) // let's keep it here as of now
	m := bizlogic.GetAllUnpaidAssessmentsForPayor(r.Context(), d.BID, TCID, &dt)

	for i, asm := range m {
		var rec UnpaidAsm
		rec.Recid = i
		rec.DtStart = rlib.JSONDate(asm.Start)
		rec.Amount = asm.Amount
		rec.AmountOwed = bizlogic.AssessmentUnpaidPortion(r.Context(), &m[i])
		rec.AmountPaid = rec.Amount - rec.AmountOwed
		ar, err := rlib.GetAR(r.Context(), asm.ARID)
		if err != nil {
			fmt.Printf("%s: Error while getting AR (ARID=%d) for Assessment: %d, error=<%s>\n", funcname, asm.ARID, asm.ASMID, err.Error())
			SvcErrorReturn(w, err, funcname)
			return
		}
		rec.Name = ar.Name
		rec.ASMID = asm.ASMID
		rec.ARID = asm.ARID
		rec.Dt = rlib.JSONDate(time.Now())
		res.Records = append(res.Records, rec)
		res.Total++
	}

	res.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &res, w)
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
	const funcname = "SvcHandlerGetUnpaidAsms"
	var (
		res PayorFundResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	TCID := d.ID
	m, _ := rlib.GetUnallocatedReceiptsByPayor(r.Context(), d.BID, TCID)

	for _, rcpt := range m {
		if rcpt.FLAGS&3 == 2 { // if there are no funds left in this receipt...
			continue // move on to the next receipt
		}

		amt := bizlogic.RemainingReceiptFunds(r.Context(), &rcpt)
		res.Record.Fund += amt
	}

	res.Status = "success"
	w.Header().Set("Content-Type", "application/json")
	SvcWriteResponse(d.BID, &res, w)
}
