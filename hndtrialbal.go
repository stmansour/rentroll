package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"text/template"
	"time"

	"github.com/dustin/go-humanize"
)

// RRCommaf returns a floating point number formated with commas for every 3 orders of magnitude
// and 2 points after the decimal
func RRCommaf(x float64) string {
	return humanize.FormatFloat("#,###.##", x)
}

// LMSum takes an array of LedgerMarkers, sums the Balance value of each, and returns the sum
func LMSum(m *[]rlib.LedgerMarker) float64 {
	bal := float64(0)
	for _, v := range *m {
		bal += v.Balance
	}
	return bal
}

func hndTrialBalance(w http.ResponseWriter, r *http.Request) {
	var err error
	var ui RRuiSupport
	var L LMResults
	ui.L = &L

	funcname := "trialBalanceHandler"

	D1 := r.FormValue("DtStart")
	D2 := r.FormValue("DtStop")
	des := r.FormValue("Business")

	ui.DtStart, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(D1))
	if err != nil {
		fmt.Printf("%s: Invalid start date:  %s\n", funcname, D1)
	}
	ui.DtStop, err = time.Parse(rlib.RRDATEINPFMT, strings.TrimSpace(D2))
	if err != nil {
		fmt.Printf("%s: Invalid start date:  %s\n", funcname, D2)
	}

	if len(des) > 0 {
		b1, _ := rlib.GetBusinessByDesignation(des)
		if len(b1.Designation) == 0 {
			rlib.Ulog("%s: Business with designation %s does net exist\n", funcname, des)
			return
		}
		ui.L.biz = &b1
	}

	rows, err := rlib.RRdb.Prepstmt.GetAllLedgerMarkersInRange.Query(ui.L.biz.BID, ui.DtStart, ui.DtStop)
	rlib.Errcheck(err)
	defer rows.Close()
	ui.L.LM = make([]rlib.LedgerMarker, 0)
	for rows.Next() {
		var r rlib.LedgerMarker
		rlib.Errcheck(rows.Scan(&r.LMID, &r.LID, &r.BID, &r.DtStart, &r.DtStop, &r.Balance, &r.State, &r.LastModTime, &r.LastModBy))
		ui.L.LM = append(ui.L.LM, r)
	}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.New("trialbal.html").Funcs(RRfuncMap).ParseFiles("./html/trialbal.html")
	if nil != err {
		fmt.Printf("%s: error loading template: %v\n", funcname, err)
	}
	err = t.Execute(w, &ui)
	if nil != err {
		rlib.LogAndPrintError(funcname, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
