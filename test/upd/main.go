// The purpose of this test is to validate the db update routines.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"rentroll/rlib"
	"strings"
	"time"

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

	updatePerson(&biz)
	updateCustomAttr(&biz)
	updateReceipt(&biz)
	updateRAPayor(&biz)
	updateRUser(&biz)
}

func updateRUser(biz *rlib.Business) {
	tcid := int64(14)
	rid := int64(1)
	_, err := rlib.GetRentableUserByRBT(rid, biz.BID, tcid)
	if err == nil {
		fmt.Printf("The database is messed up.  There should not be any RentalAgreementPayors\n")
		os.Exit(1)
	}
	if !strings.Contains(err.Error(), "no rows") {
		fmt.Printf("The database is messed up.  There should not be any RentalAgreementPayors\n")
		os.Exit(1)
	}
	now := time.Now()
	nextYear := now.AddDate(1, 0, 0)
	rap := rlib.RentableUser{RID: rid, BID: biz.BID, TCID: tcid, DtStart: now, DtStop: nextYear}
	err = rlib.InsertRentableUser(&rap)
	if err != nil {
		fmt.Printf("Error inserting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	nextYear = nextYear.AddDate(0, 11, 0)
	rap.DtStop = nextYear
	if err = rlib.UpdateRentableUserByRBT(&rap); err != nil {
		fmt.Printf("Error updating RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	r1, err := rlib.GetRentableUserByRBT(rid, biz.BID, tcid)
	if err != nil {
		fmt.Printf("Error getting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	if r1.DtStop.Equal(nextYear) {
		fmt.Printf("Error expected time = %s, found time = %s\n", r1.DtStop.Format(rlib.RRDATEFMT4), nextYear.Format(rlib.RRDATEFMT4))
		os.Exit(1)
	}
	fmt.Printf("UpdateRentableUserByRBT: successful\n")
}

func updateRAPayor(biz *rlib.Business) {
	tcid := int64(14)
	raid := int64(1)
	_, err := rlib.GetRentalAgreementPayor(raid, biz.BID, tcid)
	if err == nil {
		fmt.Printf("The database is messed up.  There should not be any RentalAgreementPayors\n")
		os.Exit(1)
	}
	if !strings.Contains(err.Error(), "no rows") {
		fmt.Printf("The database is messed up.  There should not be any RentalAgreementPayors\n")
		os.Exit(1)
	}
	now := time.Now()
	nextYear := now.AddDate(1, 0, 0)
	rap := rlib.RentalAgreementPayor{RAID: raid, BID: biz.BID, TCID: tcid, DtStart: now, DtStop: nextYear, FLAGS: uint64(0)}
	_, err = rlib.InsertRentalAgreementPayor(&rap)
	if err != nil {
		fmt.Printf("Error inserting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	nextYear = nextYear.AddDate(0, 11, 0)
	rap.DtStop = nextYear
	if err = rlib.UpdateRentalAgreementPayorByRBT(&rap); err != nil {
		fmt.Printf("Error updating RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	r1, err := rlib.GetRentalAgreementPayor(raid, biz.BID, tcid)
	if err != nil {
		fmt.Printf("Error getting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	if r1.DtStop.Equal(nextYear) {
		fmt.Printf("Error expected time = %s, found time = %s\n", r1.DtStop.Format(rlib.RRDATEFMT4), nextYear.Format(rlib.RRDATEFMT4))
		os.Exit(1)
	}
	fmt.Printf("UpdateRentalAgreementPayorByRBT: successful\n")
}

func updateReceipt(biz *rlib.Business) {
	var r rlib.Receipt
	r.RAID = 1
	r.BID = biz.BID
	r.Amount = float64(42.17)
	r.Dt = time.Date(2017, time.February, 14, 0, 0, 0, 0, time.UTC)
	r.DocNo = "12345"
	r.PMTID = 1
	_, err := rlib.InsertReceipt(&r)
	if err != nil {
		fmt.Printf("Error inserting Receipt: %s\n", err.Error())
		os.Exit(1)
	}
	r.Amount = 4217000.00
	err = rlib.UpdateReceipt(&r)
	if err != nil {
		fmt.Printf("Error updating Receipt: %s\n", err.Error())
		os.Exit(1)
	}
	r1 := rlib.GetReceiptNoAllocations(r.RCPTID)
	if r1.Amount != r.Amount {
		if err != nil {
			fmt.Printf("Updated Receipt (%d) amount error. Expected %12.2f, found %12.2f\n", r.RCPTID, r.Amount, r1.Amount)
			os.Exit(1)
		}
	}
	fmt.Printf("UpdateReceipt: successful\n")
}

func updateCustomAttr(biz *rlib.Business) {
	ca := rlib.GetCustomAttribute(1)
	ca.Value = "5000"
	err := rlib.UpdateCustomAttribute(&ca)
	if err != nil {
		fmt.Printf("Error updating CustomAttribute: %s\n", err.Error())
		os.Exit(1)
	}
	ca1 := rlib.GetCustomAttribute(1)
	if ca.Value != ca1.Value {
		fmt.Printf("CustomAttribute update failed.  Expected %s, found %s\n", ca.Value, ca1.Value)
	}
	fmt.Print("CustomAttribute updates successful\n")
}

func updatePerson(biz *rlib.Business) {
	// Update a person...
	//----------------------------------------------------
	var xp rlib.XPerson
	var err error
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
