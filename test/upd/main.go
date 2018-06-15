// The purpose of this test is to validate the db update routines.
package main

import (
	"context"
	"database/sql"
	"extres"
	"flag"
	"fmt"
	"os"
	"rentroll/rlib"
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
	NoAuth bool
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")

	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
	App.NoAuth = *noauth
}

func main() {
	var err error
	readCommandLineArgs()
	rlib.RRReadConfig()

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
	rlib.SetAuthFlag(App.NoAuth)

	// create background context
	ctx := context.Background()

	biz, err := rlib.GetBusinessByDesignation(ctx, App.Bud)
	rlib.Errcheck(err)
	if biz.BID == 0 {
		fmt.Printf("Could not find Business Unit named %s\n", App.Bud)
		os.Exit(1)
	}
	err = rlib.GetXBusiness(ctx, biz.BID, &App.Xbiz)
	rlib.Errcheck(err)

	updatePerson(ctx, &biz)
	updateCustomAttr(ctx, &biz)
	updateReceipt(ctx, &biz)
	updateRAPayor(ctx, &biz)
	updateRUser(ctx, &biz)
	updateRAR(ctx, &biz)
	Tasks(ctx, &biz)
	CheckBotNames()
}

func updateRAR(ctx context.Context, biz *rlib.Business) {
	var rar = rlib.RentalAgreementRentable{BID: 1, RAID: 2, RID: 3, ContractRent: float64(4500.00),
		RARDtStart: time.Date(2017, time.March, 7, 0, 0, 0, 0, time.UTC),
		RARDtStop:  time.Date(2018, time.March, 7, 0, 0, 0, 0, time.UTC)}
	rarid, err := rlib.InsertRentalAgreementRentable(ctx, &rar)
	if err != nil {
		fmt.Printf("Error inserting Rental Agreement Rentable: %s\n", err.Error())
		os.Exit(1)
	}
	rar.RARDtStop = rar.RARDtStop.AddDate(0, 4, 1)
	err = rlib.UpdateRentalAgreementRentable(ctx, &rar)
	if err != nil {
		fmt.Printf("Error inserting Rental Agreement Rentable: %s\n", err.Error())
		os.Exit(1)
	}
	rar2, err := rlib.GetRentalAgreementRentable(ctx, rarid)
	if err != nil {
		fmt.Printf("Error getting Rental Agreement Rentable: %s\n", err.Error())
		os.Exit(1)
	}
	if rar2.RARDtStart.Equal(rar.RARDtStart) && rar2.RARDtStop.Equal(rar.RARDtStop) &&
		rar.BID == rar2.BID && rar.RAID == rar2.RAID && rar.RID == rar2.RID &&
		rar.ContractRent == rar2.ContractRent && rar.CLID == rar2.CLID {
		fmt.Printf("UpdateRentalAgreementRentable: successful\n")
	} else {
		fmt.Printf("rar miscompared with rar2\n")
		os.Exit(1)
	}
}

func updateRUser(ctx context.Context, biz *rlib.Business) {
	tcid := int64(14)
	rid := int64(1)
	ru, err := rlib.GetRentableUserByRBT(ctx, rid, biz.BID, tcid)
	if err != nil {
		fmt.Printf("The database is messed up.  Error = %s\n", err.Error())
		os.Exit(1)
	}
	if ru.RUID > 0 {
		fmt.Printf("The database is messed up.  There should not be any RentalAgreementPayors\n")
		os.Exit(1)
	}
	now := time.Now()
	nextYear := now.AddDate(1, 0, 0)
	rap := rlib.RentableUser{RID: rid, BID: biz.BID, TCID: tcid, DtStart: now, DtStop: nextYear}
	_, err = rlib.InsertRentableUser(ctx, &rap)
	if err != nil {
		fmt.Printf("Error inserting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	nextYear = nextYear.AddDate(0, 11, 0)
	rap.DtStop = nextYear
	if err = rlib.UpdateRentableUserByRBT(ctx, &rap); err != nil {
		fmt.Printf("Error updating RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	r1, err := rlib.GetRentableUserByRBT(ctx, rid, biz.BID, tcid)
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

func updateRAPayor(ctx context.Context, biz *rlib.Business) {
	tcid := int64(14)
	raid := int64(1)
	rap1, err := rlib.GetRentalAgreementPayorByRBT(ctx, raid, biz.BID, tcid)
	if err != nil {
		fmt.Printf("A. The database is messed up. Error = %s\n", err.Error())
		os.Exit(1)
	}
	if rap1.RAPID > 0 {
		fmt.Printf("A. The database is messed up.  There should not be any RentalAgreementPayors\n")
		os.Exit(1)
	}
	now, _ := rlib.StringToDate("10/24/2016")
	next, _ := rlib.StringToDate("11/14/2016")
	rap := rlib.RentalAgreementPayor{RAID: raid, BID: biz.BID, TCID: tcid, DtStart: now, DtStop: next, FLAGS: uint64(0)}
	_, err = rlib.InsertRentalAgreementPayor(ctx, &rap)
	if err != nil {
		fmt.Printf("C. Error inserting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}

	rapid := rap.RAPID
	rap.DtStop, _ = rlib.StringToDate("1/14/2017")
	// fmt.Printf("Before update, rapid = %d,  rap.DtStop = %s\n", rapid, rap.DtStop.Format(rlib.RRDATEFMT4))
	if err = rlib.UpdateRentalAgreementPayor(ctx, &rap); err != nil {
		fmt.Printf("D. Error updating RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	r1, err := rlib.GetRentalAgreementPayor(ctx, rapid)
	if err != nil {
		fmt.Printf("E. Error getting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	// fmt.Printf("After update, rapid = %d,  r1.DtStop = %s\n", rapid, r1.DtStop.Format(rlib.RRDATEFMT4))
	if !r1.DtStop.Equal(rap.DtStop) {
		fmt.Printf("F. Error expected time = %s, found time = %s\n", rap.DtStop.Format(rlib.RRDATEFMT4), r1.DtStop.Format(rlib.RRDATEFMT4))
		os.Exit(1)
	}
	r1.DtStop, _ = rlib.StringToDate("2/14/2017")
	if err = rlib.UpdateRentalAgreementPayor(ctx, &r1); err != nil {
		fmt.Printf("G. Error updating RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	r2, err := rlib.GetRentalAgreementPayor(ctx, r1.RAPID)
	if err != nil {
		fmt.Printf("H. Error getting RentalAgreementPayor: %s\n", err.Error())
		os.Exit(1)
	}
	if !r2.DtStop.Equal(r1.DtStop) {
		fmt.Printf("I. Error expected time = %s, found time = %s\n", r1.DtStop.Format(rlib.RRDATEFMT4), r2.DtStop.Format(rlib.RRDATEFMT4))
		os.Exit(1)
	}
	fmt.Printf("UpdateRentalAgreementPayorByRBT: successful\n")
}

func updateReceipt(ctx context.Context, biz *rlib.Business) {
	var r rlib.Receipt
	r.BID = biz.BID
	r.Amount = float64(42.17)
	r.Dt = time.Date(2017, time.February, 14, 0, 0, 0, 0, time.UTC)
	r.DocNo = "12345"
	r.PMTID = 1
	_, err := rlib.InsertReceipt(ctx, &r)
	if err != nil {
		fmt.Printf("Error inserting Receipt: %s\n", err.Error())
		os.Exit(1)
	}
	r.Amount = 4217000.00
	err = rlib.UpdateReceipt(ctx, &r)
	if err != nil {
		fmt.Printf("Error updating Receipt: %s\n", err.Error())
		os.Exit(1)
	}
	r1, err := rlib.GetReceiptNoAllocations(ctx, r.RCPTID)
	rlib.Errcheck(err)
	if r1.Amount != r.Amount {
		if err != nil {
			fmt.Printf("Updated Receipt (%d) amount error. Expected %12.2f, found %12.2f\n", r.RCPTID, r.Amount, r1.Amount)
			os.Exit(1)
		}
	}
	fmt.Printf("UpdateReceipt: successful\n")
}

func updateCustomAttr(ctx context.Context, biz *rlib.Business) {
	ca, err := rlib.GetCustomAttribute(ctx, 1)
	rlib.Errcheck(err)
	ca.Value = "5000"
	err = rlib.UpdateCustomAttribute(ctx, &ca)
	if err != nil {
		fmt.Printf("Error updating CustomAttribute: %s\n", err.Error())
		os.Exit(1)
	}
	ca1, err := rlib.GetCustomAttribute(ctx, 1)
	rlib.Errcheck(err)
	if ca.Value != ca1.Value {
		fmt.Printf("CustomAttribute update failed.  Expected %s, found %s\n", ca.Value, ca1.Value)
	}
	fmt.Print("CustomAttribute updates successful\n")
}

func updatePerson(ctx context.Context, biz *rlib.Business) {
	// Update a person...
	//----------------------------------------------------
	var xp rlib.XPerson
	var err error
	TCID := int64(1)
	err = rlib.GetXPerson(ctx, TCID, &xp)
	rlib.Errcheck(err)

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
	err = rlib.UpdateTransactant(ctx, &xp.Trn)
	if err != nil {
		fmt.Printf("Error updating Transactant: %s\n", err.Error())
		os.Exit(1)
	}

	err = rlib.GetTransactant(ctx, TCID, &xp.Trn)
	rlib.Errcheck(err)
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
	err = rlib.UpdateUser(ctx, &xp.Usr)
	if err != nil {
		fmt.Printf("Error updating User: %s\n", err.Error())
		os.Exit(1)
	}

	err = rlib.GetTransactant(ctx, TCID, &xp.Trn)
	rlib.Errcheck(err)
	if xp.Usr.EmergencyContactName != ecn || xp.Usr.EmergencyContactAddress != eca || xp.Usr.EmergencyContactTelephone != ecp {
		fmt.Printf("User update failed\n")
		os.Exit(1)
	}

	fmt.Printf("Successfully updated XPerson\n")
}
