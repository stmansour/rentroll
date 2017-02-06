package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroll/rlib"
	"strings"
	"time"
)

// RARList is the struct containing the JSON return values for this web service
type RARList struct {
	Status  string
	Records []rlib.RentalAgreementRentable
}

// This command returns the rentables associated with the supplied RAID.  If no dates are supplied
// then the current date is assumed.

// SvcRARentables returns the Rentables associated with the RAID supplied
//  Called with URL:
//       0    1       2    3
// 		/gsvc/xperson/BID/RAID?dt=2017-01-03
func SvcRARentables(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Printf("entered SvcRARentables\n")
	s := r.URL.String()
	fmt.Printf("s = %s\n", s)
	s1 := strings.Split(s, "?")
	fmt.Printf("s1 = %#v\n", s1)
	ss := strings.Split(s1[0][1:], "/")
	fmt.Printf("ss = %#v\n", ss)
	raid, err := rlib.IntFromString(ss[3], "bad request integer value")
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	now := time.Now()
	dt := now
	if len(s1) > 1 && len(s1[1]) > 0 {
		sd := strings.Split(s1[1], "=")
		fmt.Printf("dt = %s\n", sd[1])
		dt, err = rlib.StringToDate(sd[1])
		if err != nil {
			SvcGridErrorReturn(w, fmt.Errorf("invalid date:  %s", sd[1]))
			return
		}
	}
	var rar RARList
	rar.Records = rlib.GetRentalAgreementRentables(raid, &dt, &dt)
	fmt.Printf("len(rar.Records) = %d\n", len(rar.Records))
	for i := 0; i < len(rar.Records); i++ {
		fmt.Printf("%d. RID = %d, ContractRent = %8.2f\n", i, rar.Records[i].RID, rar.Records[i].ContractRent)
	}
	rar.Status = "success"
	fmt.Printf("rar = %#v\n", rar)
	b, err := json.Marshal(&rar)
	if err != nil {
		SvcGridErrorReturn(w, fmt.Errorf("cannot marshal records:  %s", err.Error()))
		return
	}
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("b = %s\n", string(b))
	w.Write(b)
}
