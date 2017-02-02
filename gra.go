package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// this is a structure specifically for the UI. It will be
// automatically populated from an rlib.RentalAgreement struct
type gxrentalagr struct {
	Recid                  int64         `json:"recid"` // this is to support the w2ui form
	RAID                   int64         // internal unique id
	RATID                  int64         // reference to Occupancy Master Agreement
	BID                    rlib.XJSONBud // Business (so that we can process by Business)
	NLID                   int64         // Note ID
	AgreementStart         rlib.JSONTime // start date for rental agreement contract
	AgreementStop          rlib.JSONTime // stop date for rental agreement contract
	PossessionStart        rlib.JSONTime // start date for Occupancy
	PossessionStop         rlib.JSONTime // stop date for Occupancy
	RentStart              rlib.JSONTime // start date for Rent
	RentStop               rlib.JSONTime // stop date for Rent
	RentCycleEpoch         rlib.JSONTime // Date on which rent cycle recurs. Start date for the recurring rent assessment
	UnspecifiedAdults      int64         // adults who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels
	UnspecifiedChildren    int64         // children who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels.
	Renewal                int64         // 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
	SpecialProvisions      string        // free-form text
	LeaseType              int64         // Full Service Gross, Gross, ModifiedGross, Tripple Net
	ExpenseAdjustmentType  int64         // Base Year, No Base Year, Pass Through
	ExpensesStop           float64       // cap on the amount of oexpenses that can be passed through to the tenant
	ExpenseStopCalculation string        // note on how to determine the expense stop
	BaseYearEnd            rlib.JSONTime // last day of the base year
	ExpenseAdjustment      rlib.JSONTime // the next date on which an expense adjustment is due
	EstimatedCharges       float64       // a periodic fee charged to the tenant to reimburse LL for anticipated expenses
	RateChange             float64       // predetermined amount of rent increase, expressed as a percentage
	NextRateChange         rlib.JSONTime // he next date on which a RateChange will occur
	PermittedUses          string        // indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
	ExclusiveUses          string        // those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
	ExtensionOption        string        // the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
	ExtensionOptionNotice  rlib.JSONTime // the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string        // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  rlib.JSONTime // the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string        // Tenant may have the right to purchase their premises if LL chooses to sell
	LastModTime            rlib.JSONTime // when was this record last written
	LastModBy              int64         // employee UID (from phonebook) that modified it
}

type gxrentalagrForm struct {
	Recid                  int64         `json:"recid"` // this is to support the w2ui form
	RAID                   int64         // internal unique id
	RATID                  int64         // reference to Occupancy Master Agreement
	NLID                   int64         // Note ID
	AgreementStart         rlib.JSONTime // start date for rental agreement contract
	AgreementStop          rlib.JSONTime // stop date for rental agreement contract
	PossessionStart        rlib.JSONTime // start date for Occupancy
	PossessionStop         rlib.JSONTime // stop date for Occupancy
	RentStart              rlib.JSONTime // start date for Rent
	RentStop               rlib.JSONTime // stop date for Rent
	RentCycleEpoch         rlib.JSONTime // Date on which rent cycle recurs. Start date for the recurring rent assessment
	UnspecifiedAdults      int64         // adults who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels
	UnspecifiedChildren    int64         // children who are not accounted for in RentalAgreementPayor or RentableUser structs.  Used mostly by hotels.
	Renewal                int64         // 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
	SpecialProvisions      string        // free-form text
	LeaseType              int64         // Full Service Gross, Gross, ModifiedGross, Tripple Net
	ExpenseAdjustmentType  int64         // Base Year, No Base Year, Pass Through
	ExpensesStop           float64       // cap on the amount of oexpenses that can be passed through to the tenant
	ExpenseStopCalculation string        // note on how to determine the expense stop
	BaseYearEnd            rlib.JSONTime // last day of the base year
	ExpenseAdjustment      rlib.JSONTime // the next date on which an expense adjustment is due
	EstimatedCharges       float64       // a periodic fee charged to the tenant to reimburse LL for anticipated expenses
	RateChange             float64       // predetermined amount of rent increase, expressed as a percentage
	NextRateChange         rlib.JSONTime // he next date on which a RateChange will occur
	PermittedUses          string        // indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
	ExclusiveUses          string        // those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
	ExtensionOption        string        // the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
	ExtensionOptionNotice  rlib.JSONTime // the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
	ExpansionOption        string        // the right to expand to certanin spaces that are typically contiguous to their primary space
	ExpansionOptionNotice  rlib.JSONTime // the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
	RightOfFirstRefusal    string        // Tenant may have the right to purchase their premises if LL chooses to sell
	LastModTime            rlib.JSONTime // when was this record last written
	LastModBy              int64         // employee UID (from phonebook) that modified it
}

type gxrentalagrOther struct {
	BID rlib.W2uiHTMLSelect // Business (so that we can process by Business)
}

// SvcSearchHandlerRentalAgr generates a report of all RentalAgreements defined business d.BID
func SvcSearchHandlerRentalAgr(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcSearchHandlerRentalAgr\n")
	var p rlib.RentalAgreement
	var err error
	var g struct {
		Status  string        `json:"status"`
		Total   int64         `json:"total"`
		Records []gxrentalagr `json:"records"`
	}
	t := time.Now()
	srch := fmt.Sprintf("BID=%d AND AgreementStop>%q", d.BID, t.Format(rlib.RRDATEINPFMT)) // default WHERE clause
	order := "RAID ASC"                                                                    // default ORDER
	q, qw := gridBuildQuery("RentalAgreement", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("RentalAgreement", qw)
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("db query = %s\n", q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := int64(d.webreq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.RentalAgreement
		var q gxrentalagr
		rlib.ReadRentalAgreements(rows, &p)
		p.Recid = i
		rlib.MigrateStructVals(&p, &q)
		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.webreq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}
	fmt.Printf("g.Total = %d\n", g.Total)
	rlib.Errcheck(rows.Err())
	SvcWriteResponse(&g, w)

}

// SvcFormHandlerRentalAgreement formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the RAID as follows:
//       0    1          2    3
// 		/gsvc/rentalagrs/BID/RAID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcFormHandlerRentalAgreement\n")
	var err error

	if d.RAID, err = SvcExtractIDFromURI(r.RequestURI, "RAID", 3, w); err != nil {
		return
	}

	fmt.Printf("Requester UID = %d, BID = %d,  RAID = %d\n", d.UID, d.BID, d.RAID)

	switch d.webreq.Cmd {
	case "get":
		getRentalAgreement(w, r, d)
		break
	case "save":
		saveRentalAgreement(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.webreq.Cmd)
		SvcGridErrorReturn(w, err)
		return
	}
}

func saveRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "saveRentalAgreement"
	target := `"record":`
	fmt.Printf("SvcFormHandlerRentalAgreement save\n")
	fmt.Printf("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	fmt.Printf("record is at index = %d\n", i)
	if i < 0 {
		e := fmt.Errorf("saveRentalAgreement: cannot find %s in form json", target)
		SvcGridErrorReturn(w, e)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	fmt.Printf("data to unmarshal is:  %s\n", s)

	//===============================================================
	//------------------------------
	// Handle all the non-list data
	//------------------------------
	var foo gxrentalagrForm

	err := json.Unmarshal([]byte(s), &foo)
	if err != nil {
		e := fmt.Errorf("Error with json.Unmarshal:  %s", err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	// migrate the variables that transfer without needing special handling...
	var a rlib.RentalAgreement
	rlib.MigrateStructVals(&foo, &a)

	//---------------------------
	//  Handle all the list data
	//---------------------------
	var bar gxrentalagrOther
	err = json.Unmarshal([]byte(s), &bar)
	if err != nil {
		fmt.Printf("Data unmarshal error: %s\n", err.Error())
		e := fmt.Errorf("%s: Error with json.Unmarshal:  %s", funcname, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}

	var ok bool
	a.BID, ok = rlib.RRdb.BUDlist[bar.BID.ID]
	if !ok {
		e := fmt.Errorf("Could not map BID value: %s", bar.BID.ID)
		rlib.Ulog("%s", e.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	//===============================================================

	fmt.Printf("Update complete:  RA = %#v\n", a)

	// Now just update the database
	err = rlib.UpdateRentalAgreement(&a)
	if err != nil {
		e := fmt.Errorf("Error updating Rental Agreement RAID = %d: %s", a.RAID, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	SvcWriteSuccessResponse(w)
}

func getRentalAgreement(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	var g struct {
		Status string      `json:"status"`
		Record gxrentalagr `json:"record"`
	}
	a, err := rlib.GetRentalAgreement(d.RAID)
	if err != nil {
		e := fmt.Errorf("getRentalAgreement: cannot read RentalAgreement RAID = %d, err = %s", d.RAID, err.Error())
		SvcGridErrorReturn(w, e)
		return
	}
	if a.RAID > 0 {
		var gg gxrentalagr
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
