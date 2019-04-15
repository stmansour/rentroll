package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"rentroll/rlib"
	"time"
)

// RnTy holds config information about a rentable type
type RnTy struct {
	Count        int `json:"Count"`
	MarketRate   int `json:"MarketRate"`
	RentCycle    int `json:"RentCycle"`
	ProrateCycle int `json:"ProrateCycle"`
}

// DBGenType holds config information for the db that will be generated
type DBGenType struct {
	RSeed                int64   `json:"RSeed"`
	DtStart              string  `json:"DtStart"`
	DtStop               string  `json:"DtStop"`
	RT                   []RnTy  `json:"RT"`
	OpDepositoryName     string  `json:"OpDepositoryName"`
	SecDepDepositoryName string  `json:"SecDepDepositoryName"`
	PTypeCheckName       string  `json:"PTypeCheckName"`
	PeopleCount          int     `json:"PeopleCount"`
	RACount              int     `json:"RACount"`
	Carports             int     `json:"Carports"`
	CPMarketRate         float64 `json:"CPMarketRate"`
	CPRentCycle          int     `json:"CPRentCycle"`
	CPProrateCycle       int     `json:"CPProrateCycle"`
	RandNames            bool    `json:"RandNames"`
}

// App is the global application structure
var App struct {
	OutputType int // 0 = dbgen json, 1 = amendment RA date range
}

func readCommandLineArgs() {
	otPtr := flag.Int("outype", 0, "0 = dbgen json, 1 = escaped RA date range, 2 RA date range")

	flag.Parse()

	App.OutputType = *otPtr
}

func main() {
	readCommandLineArgs()
	now := time.Now()                                              // 9/18/2018
	d2 := now.AddDate(0, 3, 0)                                     // 12/18/2018
	d2 = time.Date(d2.Year(), d2.Month(), 1, 0, 0, 0, 0, time.UTC) // 12/1/2018
	d1 := d2.AddDate(-1, 1, 0)                                     // 1/1/2018
	d1 = time.Date(d1.Year(), d1.Month(), 1, 0, 0, 0, 0, time.UTC)
	//rlib.Console("Rental Agreement time ranges = %s\n", rlib.ConsoleDRange(&d1, &d2))

	switch App.OutputType {
	case 0:
		dbgenJSON(d1, d2)
	case 1:
		newRARangeEscaped(d1, d2)
	case 2:
		newRARange(d1, d2)
	}

}
func newRARangeEscaped(d1, d2 time.Time) {
	dtStop := d2.AddDate(1, 0, 0)
	fmt.Printf("DTSTART: %s\n", url.QueryEscape(d2.Format(rlib.RRDATEFMT3)))
	fmt.Printf("DTSTOP: %s\n", url.QueryEscape(dtStop.Format(rlib.RRDATEFMT3)))
}

func newRARange(d1, d2 time.Time) {
	dtStop := d2.AddDate(1, 0, 0)
	fmt.Printf("DTSTART: %s\n", d2.Format(rlib.RRDATEFMT3))
	fmt.Printf("DTSTOP: %s\n", dtStop.Format(rlib.RRDATEFMT3))
}

func dbgenJSON(d1, d2 time.Time) {
	var a = DBGenType{
		RSeed:   int64(1537208713575508370),
		DtStart: d1.Format(rlib.RRDATEFMT3),
		DtStop:  d2.Format(rlib.RRDATEFMT3),
		RT: []RnTy{
			{
				Count:        1,
				MarketRate:   1000,
				RentCycle:    rlib.RECURMONTHLY,
				ProrateCycle: rlib.RECURDAILY,
			},
		},
		OpDepositoryName:     "Wells Fargo",
		SecDepDepositoryName: "Bank Of America",
		PTypeCheckName:       "Check",
		PeopleCount:          1,
		RACount:              1,
		Carports:             0,
		CPMarketRate:         float64(35),
		CPRentCycle:          rlib.RECURMONTHLY,
		CPProrateCycle:       rlib.RECURDAILY,
		RandNames:            true,
	}

	b, err := json.Marshal(&a)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	fmt.Printf("%s\n", string(b))
}
