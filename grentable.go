package main

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
)

// this is a structure specifically for the UI. It will be
// automatically populated from an rlib.Rentable struct
type gxrentable struct {
	Recid          int64 `json:"recid"` // this is to support the w2ui form
	RID            int64
	BID            int64
	Name           string
	AssignmentTime rlib.XJSONAssignmentTime
	LastModTime    rlib.JSONTime
	LastModBy      int64
}

// SvcRentables generates a report of all Rentables defined business d.BID
func SvcRentables(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcRentables\n")
	var p rlib.Rentable
	var err error
	var g struct {
		Status  string       `json:"status"`
		Total   int64        `json:"total"`
		Records []gxrentable `json:"records"`
	}

	srch := fmt.Sprintf("BID=%d", d.BID) // default WHERE clause
	order := "Name ASC"                  // default ORDER
	q, qw := gridBuildQuery("Rentable", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("Rentable", qw)
	if err != nil {
		fmt.Printf("Error from GetRowCount: %s\n", err.Error())
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("db query = %s\n", q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	rlib.Errcheck(err)
	defer rows.Close()

	i := int64(d.greq.Offset)
	count := 0
	for rows.Next() {
		var p rlib.Rentable
		var q gxrentable
		rlib.ReadRentables(rows, &p)
		p.Recid = i
		rlib.MigrateStructVals(&p, &q)
		g.Records = append(g.Records, q)
		count++ // update the count only after adding the record
		if count >= d.greq.Limit {
			break // if we've added the max number requested, then exit
		}
		i++
	}
	fmt.Printf("g.Total = %d\n", g.Total)
	rlib.Errcheck(rows.Err())
	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
	SvcWriteResponse(&g, w)
}

// SvcRentable formats a complete data record for a person suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the TCID as follows:
// 		/gsvc/xperson/UID/BID/TCID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("Entered SvcRentable\n")
	var err error

	path := "/gsvc/"                // this is the part of the URL that got us into this handler
	uri := r.RequestURI[len(path):] // this pulls off the specific request
	sa := strings.Split(uri, "/")
	if len(sa) < 3 {
		e := fmt.Errorf("Error in URI, expecting /gsv/xperson/USRID/BID/RID but found: %s", uri)
		SvcGridErrorReturn(w, e)
		return
	}
	d.UID, err = rlib.IntFromString(sa[1], "not an integer number")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	d.BID, err = rlib.IntFromString(sa[2], "not an integer number")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	d.RID, err = rlib.IntFromString(sa[3], "not an integer number")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}

	fmt.Printf("Requester UID = %d, BID = %d,  RIDa = %d\n", d.UID, d.BID, d.RID)

	switch d.greq.Cmd {
	case "get":
		getRentable(w, r, d)
		break
	case "save":
		saveRentable(w, r, d)
		break
	}
}

func saveRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	// funcname := "saveRentable"
	target := `"record":`
	fmt.Printf("SvcRentable save\n")
	fmt.Printf("record data = %s\n", d.data)
	i := strings.Index(d.data, target)
	fmt.Printf("record is at index = %d\n", i)
	if i < 0 {
		e := fmt.Errorf("saveRentable: cannot find %s in form json", target)
		SvcGridErrorReturn(w, e)
		return
	}
	s := d.data[i+len(target):]
	s = s[:len(s)-1]
	fmt.Printf("data to unmarshal is:  %s\n", s)

	// var gxp gxperson
	// err := json.Unmarshal([]byte(s), &gxp)
	// if err != nil {
	// 	fmt.Printf("Data unmarshal error: %s\n", err.Error())
	// 	e := fmt.Errorf("%s: Error with json.Unmarshal:  %s\n", funcname, err.Error())
	// 	SvcGridErrorReturn(w, e)
	// 	return
	// }

	// fmt.Printf("Begin struct data migration\n")
	// var xp rlib.Rentable
	// rlib.MigrateStructVals(&gxp, &xp.Trn)
	// rlib.MigrateStructVals(&gxp, &xp.Usr)
	// rlib.MigrateStructVals(&gxp, &xp.Psp)
	// rlib.MigrateStructVals(&gxp, &xp.Pay)
	// fmt.Printf("end migration\n")

	// err = rlib.UpdateTransactant(&xp.Trn)
	// if err != nil {
	// 	e := fmt.Errorf("%s: UpdateTransactant error:  %s\n", funcname, err.Error())
	// 	SvcGridErrorReturn(w, e)
	// 	return
	// }
	// err = rlib.UpdateUser(&xp.Usr)
	// if err != nil {
	// 	e := fmt.Errorf("%s: UpdateUser error:  %s\n", funcname, err.Error())
	// 	SvcGridErrorReturn(w, e)
	// 	return
	// }

	// err = rlib.UpdateProspect(&xp.Psp)
	// if err != nil {
	// 	e := fmt.Errorf("%s: UpdateProspect error:  %s\n", funcname, err.Error())
	// 	SvcGridErrorReturn(w, e)
	// 	return
	// }

	// err = rlib.UpdatePayor(&xp.Pay)
	// if err != nil {
	// 	e := fmt.Errorf("%s: UpdatePayor err.Pay %s\n", funcname, err.Error())
	// 	SvcGridErrorReturn(w, e)
	// 	return
	// }

	// var g struct {
	// 	Status string `json:"status"`
	// }
	// g.Status = "success"
	// SvcWriteResponse(&g, w)
}

func getRentable(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered getRentable\n")
	var g struct {
		Status string     `json:"status"`
		Record gxrentable `json:"record"`
	}
	fmt.Printf("GetRentable( RID = %d )\n", d.RID)
	a := rlib.GetRentable(d.RID)
	fmt.Printf("Begin migration to form struct\n")
	if a.RID > 0 {
		var gg gxrentable
		rlib.MigrateStructVals(&a, &gg)
		g.Record = gg
	}
	fmt.Printf("End migration.  g.Record = %#v\n", g.Record)
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
