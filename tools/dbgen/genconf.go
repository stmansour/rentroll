package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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
	SQFT         int64   // square feed for this rentable
	Name         string  //
	Style        string  // very short but functionally descriptive name
}

// GenDBConf provides attribute information for what is created in the database
// Rent Assessments are created based on the Market Rate by default. A future update
// may enable varying Contract Rent amounts.
type GenDBConf struct {
	DtStart              time.Time       // default start time for all start time attributes
	DtStop               time.Time       // default stop time for all stop time attributes
	PeopleCount          int             // defines the number of Transactants
	RACount              int             // defines the number of Rental Agreements to create
	RT                   []RType         // defines the rentable types and the count of Rentables
	DtBOT                time.Time       // Beginning of Time
	DtEOT                time.Time       // End of Time
	BIZ                  []rlib.Business // all businesses in db
	ARIDrent             int64           // Acct Rule for rent assessments
	ARIDsecdep           int64           // Acct rule for security deposit assessments
	ARIDCheckPayment     int64           // Acct Rule to use for check payments
	PTypeCheck           int64           // pmtid for checks
	PTypeCheckName       string          // name of pmtid for checks
	OpDepository         int64           // the operational bank depository
	SecDepDepository     int64           // the security deposit depository
	OpDepositoryName     string          // the operational bank depository
	SecDepDepositoryName string          // the security deposit depository
	xbiz                 rlib.XBusiness  // business we're working on
	RandomizePayments    bool            // if true skip payments and allocation by percentages below
	RandMissPayment      int             // if RandomizePayments is true, skip payments on this percent (0-99)
	RandMissApply        int             // if RandomizePayments is true, skip payment application on this percent (0-99)
	RSeed                int64           // to reproduce the same database
	RSource              rand.Source     // for creating random numbers
	RRand                *rand.Rand      // our base for generating random numbers
	RandNames            bool            // if true create random names instead of predictables names
}

// GenDBRead is struct that gets loaded from the -f json file specified when
// the program starts.  The data is transferred to a GenDBConf structure
type GenDBRead struct {
	DtStart              string  `json:"DtStart"`              // default start time for all start time attributes
	DtStop               string  `json:"DtStop"`               // default stop time for all stop time attributes
	PeopleCount          int     `json:"PeopleCount"`          // defines the number of Transactants
	RACount              int     `json:"RACount"`              // defines the number of Rental Agreements to create
	OpDepositoryName     string  `json:"OpDepositoryName"`     // the operational bank depository
	SecDepDepositoryName string  `json:"SecDepDepositoryName"` // the security deposit depository
	RSeed                int64   `json:"RSeed"`                // if specified it will seed the random number generator
	RandomizePayments    int     `json:"RandomizePayments"`    // if non-zero then skip payments and allocation by percentages below
	RandNames            bool    `json:"RandNames"`            // if true then create real names rather than numeric predictable names
	RandMissPayment      int     `json:"RandMissPayment"`      // if RandomizePayments is true, skip payments on this percent (0-99)
	RandMissApply        int     `json:"RandMissApply"`        // if RandomizePayments is true, skip payment application on this percent (0-99)
	PTypeCheckName       string  // name of pmtid for checks
	RT                   []RType `json:"RT"` // defines the rentable types and the count of Rentables
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
	if err != nil {
		fmt.Printf("Error unmarshaling json file: %s\n", err.Error())
		os.Exit(1)
	}
	// rlib.Console("after unmarshal, a = %#v\n", a)
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
	b.DtEOT = time.Date(3001, time.January, 1, 0, 0, 0, 0, time.UTC)
	b.OpDepositoryName = a.OpDepositoryName
	b.SecDepDepositoryName = a.SecDepDepositoryName
	b.PTypeCheckName = a.PTypeCheckName
	b.RandomizePayments = a.RandomizePayments != 0
	b.RandNames = a.RandNames
	if b.RandomizePayments {
		b.RSeed = a.RSeed
		b.RandMissApply = a.RandMissApply
		b.RandMissPayment = a.RandMissPayment
		rlib.Console("*** RandomizePayments is in effect ***\n")
		rlib.Console("Seed = %d, MissPayments = %d%%, MissApply = %d%%\n", b.RSeed, b.RandMissPayment, b.RandMissApply)
	}

	rlib.Console("DtStart = %s, DtStop = %s\n", b.DtStart.Format(rlib.RRDATEFMT4), b.DtStop.Format(rlib.RRDATEFMT4))

	if a.RSeed == int64(0) {
		a.RSeed = time.Now().UnixNano()
	}
	rlib.Console("RSeed = %d\n", a.RSeed)
	b.RSource = rand.NewSource(a.RSeed)
	b.RSeed = a.RSeed
	b.RRand = rand.New(b.RSource)

	return b, nil
}
