package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"rentroll/rlib"
)
import _ "github.com/go-sql-driver/mysql"

// App is the global application structure
var App struct {
	dbdir  *sql.DB // phonebook db
	dbrr   *sql.DB //rentroll db
	DBDir  string  // phonebook database
	DBRR   string  //rentroll database
	DBUser string  // user for all databases
	BID    int64   //business
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")

	flag.Parse()
	App.DBUser = *dbuPtr
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
}

func initApp() {
	App.BID = 1 // known to be the only business
	rlib.InitBusinessFields(App.BID)
	rlib.RRdb.BizTypes[App.BID].NoteTypes = rlib.GetAllNoteTypes(App.BID) // we initialized to 1 business
}

func testNotes() {
	// var b rlib.Business
	// b.Designation = "REX"
	// b.DefaultRentalPeriod = rlib.ACCRUALMONTHLY
	// b.Name = "Rexford Apartments"
	// b.ParkingPermitInUse = 0
	// bid, err := rlib.InsertBusiness(&b)
	// if err != nil {
	// 	rlib.Ulog("testNotes: error inserting rlib.Business = %v\n", err)
	// }

	// funcname := "testNotes"
	// Create a note with a couple of child notes
	var n, n1, n2 rlib.Note
	n.Comment = "I am the parent note"
	n.LastModBy = 211
	n.NTID = 1 // first comment type
	nid, err := rlib.InsertNote(&n)
	rlib.Errcheck(err)

	n1.PNID = nid
	n1.Comment = "I am a note that was added to the parent note"
	n1.LastModBy = 198
	n1.NTID = 2 // second comment type
	_, err = rlib.InsertNote(&n1)
	rlib.Errcheck(err)

	n2.PNID = nid
	n2.Comment = "I am also a note that was added to the parent note. I'm the last note"
	n2.LastModBy = 207
	n2.NTID = 3 // third comment type
	_, err = rlib.InsertNote(&n2)
	rlib.Errcheck(err)

	// Now readback the notelist
	m := rlib.GetNoteList(nid)
	for i := 0; i < len(m); i++ {
		fmt.Printf("%s  UID = %3d  Type %s\n",
			m[i].LastModTime.Format(rlib.RRDATETIMEINPFMT),
			m[i].LastModBy,
			rlib.RRdb.BizTypes[App.BID].NoteTypes[i].Name)
		fmt.Printf("%s\n\n", m[i].Comment)
	}

}

func main() {
	readCommandLineArgs()
	rlib.RRReadConfig()

	var err error

	//----------------------------
	// Open RentRoll database
	//----------------------------
	// s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", DBUser, DBRR)
	s := rlib.RRGetSQLOpenString(App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = rlib.RRGetSQLOpenString(App.DBDir)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open: Error = %v\n", err)
		os.Exit(1)
	}
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("dbdir.Ping: Error = %v\n", err)
		os.Exit(1)
	}

	rlib.RpnInit()
	rlib.InitDBHelpers(App.dbrr, App.dbdir)
	initApp()

	testNotes()
}
