// The purpose of this test is to validate the time gap finder
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"rentroll/rlib"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir  *sql.DB        // phonebook db
	dbrr   *sql.DB        //rentroll db
	DBDir  string         // phonebook database
	DBRR   string         //rentroll database
	DBUser string         // user for all databases
	PortRR int            // rentroll port
	Bud    string         // Biz Unit Descriptor
	Xbiz   rlib.XBusiness // lots of info about this biz
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
}

func main() {
	// var err error
	readCommandLineArgs()
	doTest()
}

func doTest() {
	var ds = []struct {
		sd1, sd2 string
	}{
		{"2016-10-01", "2016-10-07"},
		{"2016-10-15", "2016-10-25"},
		{"2016-10-21", "2016-10-29"},
		{"2016-10-21", "2016-10-29"},
	}
	dtStart, _ := rlib.StringToDate("2016-10-01")
	dtStop, _ := rlib.StringToDate("9999-11-01")
	// dtStop, _ := rlib.StringToDate("2016-11-01")
	var d []rlib.Period
	for i := 0; i < len(ds); i++ {
		var p rlib.Period
		p.D1, _ = rlib.StringToDate(ds[i].sd1)
		p.D2, _ = rlib.StringToDate(ds[i].sd2)
		d = append(d, p)
	}
	aggrlist := rlib.TimeListAggregate(d)
	fmt.Printf("TimeListAggregate returned aggrlist.  len(aggrlist) = %d\n", len(aggrlist))
	for i := 0; i < len(aggrlist); i++ {
		fmt.Printf("aggrlist[%d] = %s - %s\n", i, aggrlist[i].D1.Format("2006-01-02"), aggrlist[i].D2.Format("2006-01-02"))
	}

	gaps := rlib.FindGaps(&dtStart, &dtStop, d)
	fmt.Print("Gaps\n")
	for i := 0; i < len(gaps); i++ {
		fmt.Printf("gaps[%d] = %s - %s\n", i, gaps[i].D1.Format("2006-01-02"), gaps[i].D2.Format("2006-01-02"))
	}
}
