package ws

import (
	"fmt"
	"net/http"
	"rentroll/rlib"
	"sort"
)

// RentableTypeRecord is struct to list down individual rentable type
type RentableTypeRecord struct {
	RTID int64  `json:"id"`
	Name string `json:"text"`
}

// GetRentableTypesResponse is the response to a GetRentable request
type GetRentableTypesResponse struct {
	Status  string               `json:"status"`
	Total   int64                `json:"total"`
	Records []RentableTypeRecord `json:"records"`
}

// SvcRentableTypesList generates a report of all Rentables defined business d.BID
// wsdoc {
//  @Title  Rentable Type List
//  @URL /v1/rtlist/:BUI
//  @Method  GET
//  @Synopsis Get Rentable Types
//  @Description Get all rentable types list for a requested business
//  @Input WebGridSearchRequest
//  @Response GetRentableTypesResponse
// wsdoc }
func SvcRentableTypesList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	fmt.Println("Entered service handler for SvcRentableTypesList")

	var (
		g GetRentableTypesResponse
	)

	// get rentable types for a business
	m := rlib.GetBusinessRentableTypes(d.BID)

	// sort keys
	var keys rlib.Int64Range
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	// append records in ascending order
	var rentableTypesList []RentableTypeRecord
	for _, rtid := range keys {
		rentableTypesList = append(rentableTypesList, RentableTypeRecord{RTID: m[rtid].RTID, Name: m[rtid].Name})
	}
	g.Records = rentableTypesList
	fmt.Printf("GetBusinessRentableTypes returned %d records\n", len(g.Records))
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(&g, w)
}
