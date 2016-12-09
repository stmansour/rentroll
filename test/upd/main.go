// The purpose of this test is to validate the db update routines.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
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
	readCommandLineArgs()
	rlib.RRReadConfig()
	var err error

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

	biz := rlib.GetBusinessByDesignation(App.Bud)
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}
	rlib.GetXBusiness(biz.BID, &App.Xbiz)

	// Update a person...
	//----------------------------------------------------
	var xp rlib.XPerson
	TCID := int64(1)
	rlib.GetXPerson(TCID, &xp)

	if len(xp.Trn.PreferredName) > 0 ||
		len(xp.Trn.MiddleName) > 0 ||
		len(xp.Trn.SecondaryEmail) > 0 {
		fmt.Printf("Initial database is not in the expected state. Re-run ./functest.sh\n")
		os.Exit(1)
	}

	pn := "Billy Bob"
	mn := "Cudworth"
	se := "quintilian@nethersole.uk"
	xp.Trn.PreferredName = pn
	xp.Trn.MiddleName = mn
	xp.Trn.SecondaryEmail = se
	err = rlib.UpdateTransactant(&xp.Trn)
	if err != nil {
		fmt.Printf("Error updating Transactant: %s\n", err.Error())
		os.Exit(1)
	}

	rlib.GetTransactant(TCID, &xp.Trn)
	if xp.Trn.PreferredName != pn || xp.Trn.MiddleName != mn || xp.Trn.SecondaryEmail != se {
		fmt.Printf("Transactant update failed\n")
		os.Exit(1)
	}

	// Update a user...
	//----------------------------------------------------
	if len(xp.Usr.EmergencyContactName) > 0 ||
		len(xp.Usr.EmergencyContactAddress) > 0 ||
		len(xp.Usr.EmergencyContactTelephone) > 0 {
		fmt.Printf("Initial database is not in the expected state. Re-run ./functest.sh\n")
		os.Exit(1)
	}

	ecn := "Howard Hughes"
	eca := "Danvers State Mental Hospital, Massachusetts"
	ecp := "BR549"
	xp.Usr.EmergencyContactName = ecn
	xp.Usr.EmergencyContactAddress = eca
	xp.Usr.EmergencyContactTelephone = ecp
	err = rlib.UpdateUser(&xp.Usr)
	if err != nil {
		fmt.Printf("Error updating Transactant: %s\n", err.Error())
		os.Exit(1)
	}

	rlib.GetTransactant(TCID, &xp.Trn)
	if xp.Usr.EmergencyContactName != ecn || xp.Usr.EmergencyContactAddress != eca || xp.Usr.EmergencyContactTelephone != ecp {
		fmt.Printf("User update failed\n")
		os.Exit(1)
	}

	fmt.Printf("Successfully updated XPerson\n")
}
