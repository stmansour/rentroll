package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"rentroll/rlib"
	"time"
)

// RType defines the common attributes of the Rentable Types
type RType struct {
	Count        int     //number of Rentables of this type
	MarketRate   float64 // amount to charge for rent
	RentCycle    int64   // 0 = nonrecur, 1 = secondly, 2 ... as defined in ./rlib/dbtypes
	ProrateCycle int64   // just like RentCycle
}

// GenDBConf provides attribute information for what is created in the database
// Rent Assessments are created based on the Market Rate by default. A future update
// may enable varying Contract Rent amounts.
type GenDBConf struct {
	DtStart     time.Time // default start time for all start time attributes
	DtStop      time.Time // default stop time for all stop time attributes
	PeopleCount int       // defines the number of Transactants
	RACount     int       // defines the number of Rental Agreements to create
	RT          []RType   // defines the rentable types and the count of Rentables
	DtBOT       time.Time // Beginning of Time
	DtEOT       time.Time // End of Time
	BIZ         []rlib.Business
	ARIDrent    int64
	ARIDsecdep  int64
	xbiz        rlib.XBusiness
}

// GenDBRead is the preliminary loading point for db generation preferences.
type GenDBRead struct {
	DtStart     string  `json:"DtStart"`     // default start time for all start time attributes
	DtStop      string  `json:"DtStop"`      // default stop time for all stop time attributes
	PeopleCount int     `json:"PeopleCount"` // defines the number of Transactants
	RACount     int     `json:"RACount"`     // defines the number of Rental Agreements to create
	RT          []RType `json:"RT"`          // defines the rentable types and the count of Rentables
}

// ReadConfig will read the configuration file  if
// it exists in the current directory
//=======================================================================================
func ReadConfig(fname string) (GenDBConf, error) {
	var a GenDBRead
	var b GenDBConf
	var err error

	if _, err = os.Stat(fname); err != nil {
		return b, err
	}

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return b, err
	}
	err = json.Unmarshal(content, &a)
	b.RT = a.RT
	b.PeopleCount = a.PeopleCount
	b.RACount = a.RACount
	b.DtStart, err = rlib.StringToDate(a.DtStart)
	if err != nil {
		return b, fmt.Errorf("Error converting date string %s: %s", a.DtStart, err.Error())
	}
	b.DtStop, err = rlib.StringToDate(a.DtStop)
	if err != nil {
		return b, fmt.Errorf("Error converting date string %s: %s", a.DtStop, err.Error())
	}
	b.DtBOT = time.Date(b.DtStart.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	b.DtEOT = time.Date(3000, time.December, 31, 0, 0, 0, 0, time.UTC)

	rlib.Console("DtStart = %s, DtStop = %s\n", b.DtStart.Format(rlib.RRDATEFMT4), b.DtStop.Format(rlib.RRDATEFMT4))

	return b, nil
}
