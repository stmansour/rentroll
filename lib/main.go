package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"runtime/debug"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

//==========================================
//    PRID = property id
//    UTID = unit type id
//   USPID = unit specialty id
//   OFSID = offset id
//  ASMTID = assessment type id
//   PMTID = payment type id
// AVAILID = availability id
//  BLDGID = building id
//  UNITID = unit id
//    TCID = transactant id
//     TID = tenant id
//     PID = payor id
//   OATID = occupancy agreement template id
//    OAID = occupancy agreement
//  RCPTID = receipt id
//  DISBID = disbursement id
//     LID = ledger id
//==========================================
type occupancyAgreement struct {
	OAID                  int
	OATID                 int
	PRID                  int
	UNITID                int
	PID                   int
	PrimaryTenant         int
	OccupancyStart        time.Time
	OccupancyStop         time.Time
	Renewal               int
	ProrationMethod       int
	SecurityDepositAmount float32
	SpecialProvisions     string
}

type assessment struct {
	ASMID     int
	UNITID    int
	ASMTID    int
	Amount    float32
	Start     time.Time
	Stop      time.Time
	Frequency int
}

// collection of prepared sql statements
type prepSQL struct {
	occAgrByProperty *sql.Stmt
}

// App is the global data structure for this app
var App struct {
	dbdir    *sql.DB
	dbrr     *sql.DB
	DBDir    string
	DBRR     string
	DBUser   string
	prepstmt prepSQL
}

func errcheck(err error) {
	if err != nil {
		fmt.Printf("error = %v\n", err)
		debug.PrintStack()
		log.Fatal(err)
	}
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	flag.Parse()
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
}

func main() {
	readCommandLineArgs()

	var err error
	s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", App.DBUser, App.DBDir)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBDir, App.DBUser, err)
	}
	defer App.dbdir.Close()
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("App.DBDir.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBDir, App.DBUser, err)
	}

	s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", App.DBUser, App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("App.DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
	}
	buildPreparedStatements()

	//  func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	start := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2015, time.December, 31, 23, 59, 59, 0, time.UTC)
	doPropertyAssessments(&start, &stop, 1)

}
