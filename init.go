package main

import "rentroll/rlib"

func initRentRoll() {
	initLists()
	initJFmt()
	initTFmt()
	rpnInit()
	loadDefaultCashAccts()
}

func initLists() {
	App.AsmtTypes = rlib.GetAssessmentTypes()
	App.PmtTypes = rlib.GetPaymentTypes()
}

// Basically this is turned off for now. We'll get to default cash accounts at some point.
func loadDefaultCashAccts() {
	// App.DefaultCash = make(map[int64]LedgerMarker, 0)
	// s := "SELECT BID,Address,Address2,City,State,PostalCode,Country,Phone,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	s := "SELECT BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModTime,LastModBy from business"
	rows, err := App.dbrr.Query(s)
	rlib.Errcheck(err)
	defer rows.Close()
	for rows.Next() {
		var xprop rlib.XBusiness
		rlib.Errcheck(rows.Scan(&xprop.P.BID, &xprop.P.Designation,
			/* &xprop.P.Address, &xprop.P.Address2, &xprop.P.City, &xprop.P.State,
			&xprop.P.PostalCode, &xprop.P.Country, &xprop.P.Phone,  */
			&xprop.P.Name, &xprop.P.DefaultOccupancyType,
			&xprop.P.ParkingPermitInUse, &xprop.P.LastModTime, &xprop.P.LastModBy))

		// All we really needed was the BID...
		// App.DefaultCash[xprop.P.BID] = GetDefaultCashLedgerMarker(xprop.P.BID)
		// if App.DefaultCash[xprop.P.BID].LMID == 0 {
		// 	fmt.Printf("No default cash account was found for business %d, %s\n", xprop.P.BID, xprop.P.Name)
		// }
	}
}
