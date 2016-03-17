package main

import (
	"fmt"
	"rentroll/rlib"
)

func initLists() {
	App.AsmtTypes = GetAssessmentTypes()
	App.PmtTypes = GetPaymentTypes()
}

func loadDefaultCashAccts() {
	App.DefaultCash = make(map[int]LedgerMarker, 0)
	s := "SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var xprop XBusiness
		rlib.Errcheck(rows.Scan(&xprop.P.BID, &xprop.P.Address, &xprop.P.Address2, &xprop.P.City, &xprop.P.State,
			&xprop.P.PostalCode, &xprop.P.Country, &xprop.P.Phone, &xprop.P.Name, &xprop.P.DefaultOccupancyType,
			&xprop.P.ParkingPermitInUse, &xprop.P.LastModTime, &xprop.P.LastModBy))

		// All we really needed was the BID...
		App.DefaultCash[xprop.P.BID] = GetDefaultCashLedgerMarker(xprop.P.BID)
		if App.DefaultCash[xprop.P.BID].LMID == 0 {
			fmt.Printf("No default cash account was found for business %d, %s\n", xprop.P.BID, xprop.P.Name)
		}
	}
}
