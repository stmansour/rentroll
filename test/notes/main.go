package main

import (
	"database/sql"
	"extres"
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

func printNote(n *rlib.Note, indent int) {
	ind := ""
	for i := 0; i < indent; i++ {
		ind += fmt.Sprintf(" ")
	}

	fmt.Printf("%sN%03d  NL%03d  PN%03d  U%04d  %s\n",
		ind, n.NID, n.NLID, n.PNID, n.LastModBy,
		n.LastModTime.Format(rlib.RRDATETIMEINPFMT))
	fmt.Printf("%s%s\n\n", ind, n.Comment)
}

func testNotes() {
	// funcname := "testNotes"

	// Create a notelist with a couple of child notes
	var nl rlib.NoteList
	var err error

	nl.LastModBy = 128
	nl.NLID, err = rlib.InsertNoteList(&nl)

	// add some notes to the NoteList

	var n, n1, n2, n3 rlib.Note
	n.Comment = "I am the parent note. I have much to say. Bla bla"
	n.LastModBy = 211
	n.NTID = 1 // first comment type
	n.NLID = nl.NLID
	nid, err := rlib.InsertNote(&n)
	rlib.Errcheck(err)

	n1.PNID = nid
	n1.NLID = nl.NLID
	n1.Comment = "I am a note that was added to the parent note"
	n1.LastModBy = 198
	n1.NTID = 2 // second comment type
	_, err = rlib.InsertNote(&n1)
	rlib.Errcheck(err)

	n2.PNID = nid
	n2.NLID = nl.NLID
	n2.Comment = "I am also a note that was added to the parent note. I'm the last note"
	n2.LastModBy = 207
	n2.NTID = 3 // third comment type
	_, err = rlib.InsertNote(&n2)
	rlib.Errcheck(err)

	// this one is the second not in the nlist
	n3.NLID = nl.NLID
	n3.Comment = "I should be the second note in the NoteList"
	n3.LastModBy = 211
	n3.NTID = 1 // third comment type
	_, err = rlib.InsertNote(&n3)
	if err != nil {
		fmt.Printf("Error inserting note: %s\n", err.Error())
	}

	// Now readback the notelist
	d := rlib.GetNoteList(nl.NLID)
	fmt.Printf("NoteList: %d  created by uid %d\n", d.NLID, d.LastModBy)
	for i := 0; i < len(d.N); i++ {
		printNote(&d.N[i], 0)
		for j := 0; j < len(d.N[i].CN); j++ {
			printNote(&d.N[i].CN[j], 4)
		}
	}
}

func main() {
	var err error
	readCommandLineArgs()

	//----------------------------
	// Open RentRoll database
	//----------------------------
	if err = rlib.RRReadConfig(); err != nil {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	s := extres.GetSQLOpenString(rlib.AppConfig.RRDbname, &rlib.AppConfig)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = extres.GetSQLOpenString(rlib.AppConfig.Dbname, &rlib.AppConfig)
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
