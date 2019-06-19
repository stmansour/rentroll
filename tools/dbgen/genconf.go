package main

import (
	"context"
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
	Reserve      bool    // indicates whether rentable should be reserved when RA ends
}

// GenDBConf provides attribute information for what is used by the code in
// the generation of the database.
type GenDBConf struct {
	DtStart              time.Time           // default start time for all start time attributes
	DtStop               time.Time           // default stop time for all stop time attributes
	PeopleCount          int                 // defines the number of Transactants
	RACount              int                 // defines the number of Rental Agreements to create
	RT                   []RType             // defines the rentable types and the count of Rentables
	Carports             int                 // count of carports, style CPnnn, name "Car Port nnn"
	CPMarketRate         float64             // market rate for Carports
	CPRentCycle          int64               // 0 = nonrecur, 1 = secondly, 2 ... as defined in ./rlib/dbtypes
	CPProrateCycle       int64               // just like RentCycle
	DtBOT                time.Time           // Beginning of Time
	DtEOT                time.Time           // End of Time
	BIZ                  []rlib.Business     // all businesses in db
	ARIDrent             int64               // Acct Rule for rent assessments
	ARIDsecdep           int64               // Acct rule for security deposit assessments
	ARIDCheckPayment     int64               // Acct Rule to use for check payments
	PTypeCheck           int64               // pmtid for checks
	PTypeCheckName       string              // name of pmtid for checks
	OpDepository         int64               // the operational bank depository
	SecDepDepository     int64               // the security deposit depository
	OpDepositoryName     string              // the operational bank depository
	SecDepDepositoryName string              // the security deposit depository
	xbiz                 rlib.XBusiness      // business we're working on
	RandomizePayments    bool                // if true skip payments and allocation by percentages below
	RandMissPayment      int                 // if RandomizePayments is true, skip payments on this percent (0-99)
	RandMissApply        int                 // if RandomizePayments is true, skip payment application on this percent (0-99)
	RSeed                int64               // to reproduce the same database
	RSource              rand.Source         // for creating random numbers
	RRand                *rand.Rand          // our base for generating random numbers
	RandNames            bool                // if true create random names instead of predictables names
	PetFees              []rlib.AR           // account rules applied to new Rental Agreements for pets
	VehicleFees          []rlib.AR           // account rules applied to new Rental Agreements for vehicles
	Epochs               rlib.BizPropsEpochs // default epochs (triggers)
	ResDepARID           int64               // ARID of rule to use for reservation deposits
	HotelReserveDtStart  time.Time           // start of date range for reservations.  Default is DtStop + 1 day
	HotelReserveDtStop   time.Time           // stop of date range for reservations.  Default is DtStart + 1 year
	HotelReservePct      float64             // the percent of hotel rooms that are reserved from (HotelReserveDtStart,HotelReserveDtStop)
}

// GenDBRead is struct that gets loaded from the -f json file specified when
// the program starts.  The data is transferred to a GenDBConf structure
type GenDBRead struct {
	DtStart              string   `json:"DtStart"`              // default start time for all start time attributes
	DtStop               string   `json:"DtStop"`               // default stop time for all stop time attributes
	PeopleCount          int      `json:"PeopleCount"`          // defines the number of Transactants
	RACount              int      `json:"RACount"`              // defines the number of Rental Agreements to create
	OpDepositoryName     string   `json:"OpDepositoryName"`     // the operational bank depository
	SecDepDepositoryName string   `json:"SecDepDepositoryName"` // the security deposit depository
	RSeed                int64    `json:"RSeed"`                // if specified it will seed the random number generator
	RandomizePayments    bool     `json:"RandomizePayments"`    // if non-zero then skip payments and allocation by percentages below
	RandNames            bool     `json:"RandNames"`            // if true then create real names rather than numeric predictable names
	RandMissPayment      int      `json:"RandMissPayment"`      // if RandomizePayments is true, skip payments on this percent (0-99)
	RandMissApply        int      `json:"RandMissApply"`        // if RandomizePayments is true, skip payment application on this percent (0-99)
	Carports             int      `json:"Carports"`             // number of carports -- they can be child rentables
	CPMarketRate         float64  `json:"CPMarketRate"`         // market rate for Carports
	CPRentCycle          int64    // 0 = nonrecur, 1 = secondly, 2 ... as defined in ./rlib/dbtypes
	CPProrateCycle       int64    // just like RentCycle
	PTypeCheckName       string   // name of ptid for checks
	RT                   []RType  `json:"RT"` // defines the rentable types and the count of Rentables
	PetFees              []string // array of Account Rule names for pet fees on a new Rental Agreement
	VehicleFees          []string // array of Account Rule names for vehicle fees on a new Rental Agreement
	HotelReserveDtStart  string   `json:"HotelReserveDtStart"` // start of date range for reservations.  Default is DtStop + 1 day
	HotelReserveDtStop   string   `json:"HotelReserveDtStop"`  // stop of date range for reservations.  Default is DtStart + 1 year
	HotelReservePct      float64  // the percent of hotel rooms that are reserved from (HotelReserveDtStart,HotelReserveDtStop)
}

// ReadConfig will read the configuration file  if
// it exists in the current directory
//=======================================================================================
func ReadConfig(fname string) (GenDBConf, error) {
	var a GenDBRead
	var b GenDBConf
	var err error
	var BID = int64(1)

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

	b.PeopleCount = a.PeopleCount
	b.RACount = a.RACount
	b.OpDepositoryName = a.OpDepositoryName
	b.SecDepDepositoryName = a.SecDepDepositoryName
	b.Carports = a.Carports
	b.CPMarketRate = a.CPMarketRate
	b.RandNames = a.RandNames
	b.PTypeCheckName = a.PTypeCheckName
	b.CPRentCycle = a.CPRentCycle
	b.CPProrateCycle = a.CPProrateCycle

	b.RandomizePayments = a.RandomizePayments
	b.RT = a.RT
	b.DtStart, err = rlib.StringToDate(a.DtStart)
	if err != nil {
		return b, fmt.Errorf("Error converting date string %s: %s", a.DtStart, err.Error())
	}
	b.DtStop, err = rlib.StringToDate(a.DtStop)
	if err != nil {
		return b, fmt.Errorf("Error converting date string %s: %s", a.DtStop, err.Error())
	}
	if len(a.HotelReserveDtStart) > 0 {
		b.HotelReserveDtStart, err = rlib.StringToDate(a.HotelReserveDtStart)
		if err != nil {
			return b, fmt.Errorf("Error converting date string %s: %s", a.HotelReserveDtStart, err.Error())
		}
	} else {
		b.HotelReserveDtStart = b.DtStart.Add(24 * time.Hour)
	}
	if len(a.HotelReserveDtStop) > 0 {
		b.HotelReserveDtStop, err = rlib.StringToDate(a.HotelReserveDtStop)
		if err != nil {
			return b, fmt.Errorf("Error converting date string %s: %s", a.HotelReserveDtStop, err.Error())
		}
	} else {
		b.HotelReserveDtStop = b.DtStart.AddDate(1, 0, 0)
	}
	if a.HotelReservePct == float64(0) {
		b.HotelReservePct = float64(0.5)
	}
	b.DtBOT = time.Date(b.DtStart.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	b.DtEOT = rlib.ENDOFTIME
	if b.RandomizePayments {
		b.RSeed = a.RSeed
		b.RandMissApply = a.RandMissApply
		b.RandMissPayment = a.RandMissPayment
		rlib.Console("*** RandomizePayments is in effect ***\n")
		rlib.Console("Seed = %d, MissPayments = %d%%, MissApply = %d%%\n", b.RSeed, b.RandMissPayment, b.RandMissApply)
	}

	ctx := context.Background()

	//-----------------------------------------------
	// Get the reservation deposit Account Rule...
	//-----------------------------------------------
	var ar rlib.AR
	if ar, err = rlib.GetARByName(ctx, 1, "Reservation Deposit"); err != nil {
		return b, fmt.Errorf("Could not get Reservation Deposit account rule.  err = %s", err.Error())
	}

	//--------------------------------
	// BUSINESS PROPERTIES
	//--------------------------------
	var bp = rlib.BizProps{
		PetFees:     []string{},
		VehicleFees: []string{},
		ResDepARID:  ar.ARID,
	}
	bp.PetFees = a.PetFees
	bp.VehicleFees = a.VehicleFees

	epoch := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	bp.Epochs.Daily = epoch
	bp.Epochs.Weekly = epoch
	bp.Epochs.Monthly = epoch
	bp.Epochs.Quarterly = epoch
	bp.Epochs.Yearly = epoch
	b.Epochs = bp.Epochs
	b.ResDepARID = bp.ResDepARID

	data, err := json.Marshal(&bp)
	if err != nil {
		return b, err
	}
	var props = rlib.BusinessProperties{
		BID:   BID,
		Name:  "general",
		FLAGS: 0,
		Data:  data,
	}
	if _, err = rlib.InsertBusinessProperties(ctx, &props); err != nil {
		return b, err
	}

	if err = InitBizProps(ctx, BID, a.PetFees, &b.PetFees, "PetFees"); err != nil {
		return b, err
	}

	if err = InitBizProps(ctx, BID, a.VehicleFees, &b.VehicleFees, "VehicleFees"); err != nil {
		return b, err
	}

	//-------------------------------------
	// RANDOM NUMBER GENERATOR
	//-------------------------------------
	if a.RSeed == int64(0) {
		a.RSeed = time.Now().UnixNano()
	}
	b.RSource = rand.NewSource(a.RSeed)
	b.RSeed = a.RSeed
	b.RRand = rand.New(b.RSource)

	// rlib.Console("Date Range:  %s\n\n", rlib.ConsoleDRange(&b.DtStart, &b.DtStop))
	rlib.Console("Rental Agreement range: %s\n", rlib.ConsoleDRange(&b.DtStart, &b.DtStop))
	rlib.Console("          HotelReserve: %s\n", rlib.ConsoleDRange(&b.HotelReserveDtStart, &b.HotelReserveDtStop))
	rlib.Console(" Hotel Reserve Percent: %2.1f%%\n", float64(100)*b.HotelReservePct)
	rlib.Console("                 RSeed: %d\n", a.RSeed)

	return b, nil
}

// InitBizProps initializes the business properties with the supplied info and
// updates the GenDBConf struct with the associated list of account rules.
//
// INPUTS:
//
func InitBizProps(ctx context.Context, bid int64, a []string, b *[]rlib.AR, s string) error {

	for i := 0; i < len(a); i++ {
		ar, err := rlib.GetARByName(ctx, bid, a[i])
		if err != nil {
			return err
		}
		*b = append(*b, ar)
	}
	return nil
}
