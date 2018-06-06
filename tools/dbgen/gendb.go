package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

type tableMakerFunc func(context.Context, *GenDBConf) error

type tableMaker struct {
	Name    string
	Handler tableMakerFunc
}

var iRID = int64(1)

var handlers = []tableMaker{
	{"People", createTransactants},
	{"Rentable Types and Rentables", createRentableTypesAndRentables},
	{"Rental Agreements", createRentalAgreements},
	{"Receipts", createReceipts},
	{"ApplyReceipts", applyReceipts},
	{"Deposits", CreateDeposits},
	{"TaskLists", CreateTaskLists},
}

// GenerateDB is the RentRoll Database generator. It creates a
// database for testing based on parameters in the supplied configuration
// context dbConf.
//
// The current implementation adds to the existing database. Typically a
// database is created with the following information already in it:
//
//		* Business
//		* GLAccounts (Chart of Accounts)
//		* AR (Account Rules)
//		* Payment Types
// 		* Depositories
// 		* Deposit Methods
//		* Rental Agreement Templates
//
// A database like this is stored in empty.sql and can be used or replaced
// with any other starting point database.
//
//
// INPUTS:
//  ctx    - database ctx
//  dbConf - conf; the configuration data
//
// RETURNS:
//  any errors encountered
//-----------------------------------------------------------------------------
func GenerateDB(ctx context.Context, dbConf *GenDBConf) error {
	var (
		ar  rlib.AR
		err error
	)

	BID := dbConf.BIZ[0].BID
	err = rlib.InitBizInternals(BID, &dbConf.xbiz) // used by handlers
	if err != nil {
		return err
	}

	// These are the account rules the program needs
	var ars = []struct {
		name string
		ar   *int64
	}{
		{"Rent Non-Taxable", &dbConf.ARIDrent},
		{"Security Deposit Assessment", &dbConf.ARIDsecdep},
		{"Receive a Payment", &dbConf.ARIDCheckPayment},
	}
	//---------------------------------
	// Load the account rules needed...
	//---------------------------------
	for i := 0; i < len(ars); i++ {
		ar, err = rlib.GetARByName(ctx, BID, ars[i].name)
		if err != nil {
			fmt.Printf("Error getting Account Rule %s: %s\n", ars[i].name, err.Error())
			os.Exit(1)
		}
		if ar.ARID == 0 {
			fmt.Printf("err: account rule %q is missing\n", ars[i].name)
			os.Exit(1)
		}
		*(ars[i].ar) = ar.ARID
	}

	if dbConf.OpDepository == 0 {
		d, err := rlib.GetDepositoryByName(ctx, BID, dbConf.OpDepositoryName)
		rlib.Errcheck(err)
		if d.DEPID == 0 {
			fmt.Printf("Creating Depository:  %s", dbConf.OpDepositoryName)
			d = rlib.Depository{}
		}
		dbConf.OpDepository = d.DEPID
	}

	if dbConf.SecDepDepository == 0 {
		d, err := rlib.GetDepositoryByName(ctx, BID, dbConf.SecDepDepositoryName)
		rlib.Errcheck(err)
		if d.DEPID == 0 {
			return fmt.Errorf("Could not find Depository named %q", dbConf.SecDepDepositoryName)
		}
		dbConf.SecDepDepository = d.DEPID
	}
	if dbConf.PTypeCheck == 0 {
		var pt rlib.PaymentType
		err = rlib.GetPaymentTypeByName(ctx, BID, dbConf.PTypeCheckName, &pt)
		rlib.Errcheck(err)
		if pt.PMTID == 0 {
			return fmt.Errorf("Could not find Payment Type with name %q", dbConf.PTypeCheckName)
		}
		dbConf.PTypeCheck = pt.PMTID
	}

	//---------------------------------------
	// Now spin through all the handlers...
	//---------------------------------------
	for i := 0; i < len(handlers); i++ {
		if err := handlers[i].Handler(ctx, dbConf); err != nil {
			return err
		}
	}

	return nil
}

// createTransactants
//-----------------------------------------------------------------------------
func createTransactants(ctx context.Context, dbConf *GenDBConf) error {
	for i := 0; i < dbConf.PeopleCount; i++ {
		var t rlib.Transactant
		t.BID = dbConf.BIZ[0].BID
		if dbConf.RandNames {
			t.FirstName = GenerateRandomFirstName()
			t.MiddleName = GenerateRandomFirstName()
			t.LastName = GenerateRandomLastName()
			t.PreferredName = GenerateRandomFirstName()
			t.CellPhone = GenerateRandomPhoneNumber()
		} else {
			t.FirstName = fmt.Sprintf("John%04d", i)
			t.MiddleName = "Q"
			t.LastName = fmt.Sprintf("Doe%04d", i)
			t.PreferredName = fmt.Sprintf("J%04d", i)
			t.CellPhone = GenerateRandomPhoneNumber()
		}

		t.Address = GenerateRandomAddress()
		t.City = GenerateRandomCity()
		t.State = GenerateRandomState()
		t.Country = "USA"
		t.PostalCode = fmt.Sprintf("%05d", rand.Intn(100000))
		t.PrimaryEmail = GenerateRandomEmail(t.LastName, t.FirstName)
		t.SecondaryEmail = GenerateRandomEmail(t.LastName, t.FirstName)
		t.CompanyName = GenerateRandomCompany()
		t.WorkPhone = GenerateRandomPhoneNumber()

		_, err := rlib.InsertTransactant(ctx, &t)
		if err != nil {
			return err
		}
	}
	return nil
}

// createRentableTypesAndRentables
//-----------------------------------------------------------------------------
func createRentableTypesAndRentables(ctx context.Context, dbConf *GenDBConf) error {
	var err error
	for i := 0; i < len(dbConf.RT); i++ {

		var name = dbConf.RT[i].Name
		if len(name) == 0 {
			name = fmt.Sprintf("RType%03d", i)
		}
		var style = dbConf.RT[i].Style
		if len(style) == 0 {
			style = fmt.Sprintf("ST%03d", i)
		}

		//-------------------------------
		// Default rent AccountRule for
		// this RentableType
		//-------------------------------
		var ar rlib.AR
		ar.BID = dbConf.BIZ[0].BID
		ar.Name = fmt.Sprintf("Rent %s", style)
		ar.Description = fmt.Sprintf("Default rent assessment for rentable type %s", name)
		ar.ARType = 0                              // Assessment
		ar.DebitLID = 9                            // Acct# 12001 - RentRoll Receivables
		ar.CreditLID = 18                          // Acct# 41001 - Gross Scheduled Rent non-taxable
		ar.DefaultAmount = dbConf.RT[i].MarketRate // default rent amount
		ar.FLAGS = (1 << 2) | (1 << 4)             // RAID rqd, is Rent
		ar.DtStart = rlib.TIME0                    // make this rule "forever"
		ar.DtStop = rlib.ENDOFTIME                 // make this rule "forever"
		_, err = rlib.InsertAR(ctx, &ar)
		if err != nil {
			return err
		}

		//-------------------------------
		// RentableType...
		//-------------------------------
		var rt rlib.RentableType
		rt.BID = dbConf.BIZ[0].BID
		rt.Style = style
		rt.Name = name
		rt.RentCycle = dbConf.RT[i].RentCycle
		rt.Proration = dbConf.RT[i].ProrateCycle
		rt.GSRPC = dbConf.RT[i].ProrateCycle
		rt.FLAGS |= 0x4
		rt.ARID = ar.ARID
		_, err = rlib.InsertRentableType(ctx, &rt)
		if err != nil {
			return err
		}

		//-------------------------------
		// RentableMarketRate...
		//-------------------------------
		var mr rlib.RentableMarketRate
		mr.DtStart = dbConf.DtBOT
		mr.DtStop = dbConf.DtEOT
		mr.MarketRate = dbConf.RT[i].MarketRate
		mr.RTID = rt.RTID
		_, err = rlib.InsertRentableMarketRates(ctx, &mr)
		if err != nil {
			return err
		}

		if err = createRentables(ctx, dbConf, &dbConf.RT[i], &mr, rt.RTID); err != nil {
			return err
		}

		//-------------------------------
		// Custom Attributes...
		//-------------------------------
		if dbConf.RT[i].SQFT > 0 {
			// rlib.Console("Found custom attribute SQFT = %d\n", dbConf.RT[i].SQFT)
			var c rlib.CustomAttribute
			c.BID = rt.BID
			c.Name = "Square Feet"
			c.Type = rlib.CUSTINT
			c.Units = "sqft"
			c.Value = fmt.Sprintf("%d", dbConf.RT[i].SQFT)
			_, err = rlib.InsertCustomAttribute(ctx, &c)
			if err != nil {
				return err
			}

			var cr rlib.CustomAttributeRef
			cr.BID = rt.BID                        // this business
			cr.ElementType = rlib.ELEMRENTABLETYPE // the id is that of a RentableType
			cr.ID = rt.RTID                        // this is the RTID
			cr.CID = c.CID                         // this is the associated custom attribute
			_, err = rlib.InsertCustomAttributeRef(ctx, &cr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// createRentables
//-----------------------------------------------------------------------------
func createRentables(ctx context.Context, dbConf *GenDBConf, rt *RType, mr *rlib.RentableMarketRate, RTID int64) error {
	for i := 0; i < rt.Count; i++ {
		var r rlib.Rentable
		var err error

		r.RID = iRID
		r.BID = dbConf.BIZ[0].BID
		r.RentableName = fmt.Sprintf("Rentable%03d", iRID)
		errlist := bizlogic.InsertRentable(ctx, &r)
		if errlist != nil {
			return bizlogic.BizErrorListToError(errlist)
		}

		var rtr rlib.RentableTypeRef
		rtr.DtStart = dbConf.DtBOT
		rtr.DtStop = dbConf.DtEOT
		rtr.BID = dbConf.BIZ[0].BID
		rtr.RTID = RTID
		rtr.RID = r.RID
		_, err = rlib.InsertRentableTypeRef(ctx, &rtr)
		if err != nil {
			return err
		}

		var rs rlib.RentableStatus
		rs.DtStart = dbConf.DtBOT
		rs.DtStop = dbConf.DtEOT
		rs.BID = dbConf.BIZ[0].BID
		rs.RID = r.RID
		rs.LeaseStatus = rlib.LEASESTATUSleased
		rs.UseStatus = rlib.USESTATUSinService
		_, err = rlib.InsertRentableStatus(ctx, &rs)
		if err != nil {
			return err
		}
		iRID++
	}
	return nil
}

// createReceipts reads all assessments and creates a separate receipt for
// each one.
//-----------------------------------------------------------------------------
func createReceipts(ctx context.Context, dbConf *GenDBConf) error {
	qry := fmt.Sprintf("SELECT %s FROM Assessments WHERE BID=%d AND (PASMID=0 OR RentCycle=0)", rlib.RRdb.DBFields["Assessments"], dbConf.BIZ[0].BID)
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		if dbConf.RandomizePayments && dbConf.RRand.Intn(100) < dbConf.RandMissPayment {
			continue // some randomness - a missed payment
		}
		var a rlib.Assessment
		err = rlib.ReadAssessments(rows, &a)
		if err != nil {
			return err
		}

		if !((a.RentCycle > rlib.RECURNONE && a.PASMID > 0) || a.RentCycle == rlib.RECURNONE) {
			continue
		}
		depid := dbConf.OpDepository
		if a.ARID == dbConf.ARIDsecdep {
			depid = dbConf.SecDepDepository
		}

		var rcpt rlib.Receipt
		rcpt.ARID = dbConf.ARIDCheckPayment
		rcpt.BID = dbConf.BIZ[0].BID
		rcpt.PMTID = dbConf.PTypeCheck
		rcpt.DEPID = depid
		rcpt.RAID = a.RAID
		rcpt.Dt = a.Start
		rcpt.DocNo = fmt.Sprintf("%d", rand.Int63n(int64(1000000)))
		rcpt.Amount = a.Amount
		rcpt.ARID = dbConf.ARIDCheckPayment
		pa, _ := rlib.GetRentalAgreementPayorsInRange(ctx, a.RAID, &rlib.TIME0, &rlib.ENDOFTIME)
		if len(pa) > 0 {
			rcpt.TCID = pa[0].TCID
		}

		err = bizlogic.InsertReceipt(ctx, &rcpt)
		if err != nil {
			return err
		}
	}
	return rows.Err()
}

// applyReceipts reads all transactants and applies all their unallocated
// funds to unpaid Assessments
//-----------------------------------------------------------------------------
func applyReceipts(ctx context.Context, dbConf *GenDBConf) error {
	// rlib.Console("Entered applyReceipts\n")

	rows, err := rlib.RRdb.Prepstmt.GetUnallocatedReceipts.Query(dbConf.BIZ[0].BID)
	if err != nil {
		return err
	}
	defer rows.Close()

	// We need a list of payors.  Build a map indexed by TCID, that points
	// to the total number of receipts for that payor which are unallocated.
	var u = map[int64]int{}
	for rows.Next() {
		var r rlib.Receipt
		err = rlib.ReadReceipts(rows, &r)
		if err != nil {
			return err
		}
		// rlib.Console("Unallocated Receipt:  RCPTID = %d, Amount = %8.2f, Payor = %d\n", r.RCPTID, r.Amount, r.TCID)
		i, ok := u[r.TCID]
		if ok {
			u[r.TCID] = i + 1
		} else {
			u[r.TCID] = 1
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	// rlib.Console("Payors with unallocated receipts:\n")
	for k := range u {
		if dbConf.RandomizePayments && dbConf.RRand.Intn(100) < dbConf.RandMissApply {
			continue // some randomness - don't apply this payment
		}
		// rlib.Console("Payor with TCID=%d has %d unallocated receipts\n", k, v)
		dt := dbConf.DtStart
		bizlogic.AutoAllocatePayorReceipts(ctx, k, &dt)
	}

	return nil
}

// createRentalAgreements
//-----------------------------------------------------------------------------
func createRentalAgreements(ctx context.Context, dbConf *GenDBConf) error {
	BID := dbConf.BIZ[0].BID
	err := rlib.GetXBusiness(ctx, BID, &dbConf.xbiz)
	if err != nil {
		return err
	}
	d1 := time.Date(dbConf.DtStart.Year(), dbConf.DtStart.Month(), dbConf.DtStart.Day(), 0, 0, 0, 0, time.UTC)
	epoch := time.Date(dbConf.DtStart.Year(), dbConf.DtStart.Month(), 1, 0, 0, 0, 0, time.UTC)
	if dbConf.DtStart.Day() > 1 {
		epoch = epoch.AddDate(0, 1, 0)

	}
	d2 := epoch.AddDate(2, 0, 0)
	if d2.Day() != 1 {
		d2 = time.Date(d2.Year(), d2.Month(), 1, 0, 0, 0, 0, time.UTC)
	}
	rentableC, err := rlib.GetCountByTableName(ctx, "Rentable", BID)
	if err != nil {
		return err
	}
	MaxRID := int64(rentableC)

	tC, err := rlib.GetCountByTableName(ctx, "Transactant", BID)
	if err != nil {
		return err
	}
	MaxTCID := int64(tC)

	RID := int64(1)

	for i := 0; i < dbConf.RACount; i++ {
		var ra rlib.RentalAgreement
		ra.RATID = 1
		ra.BID = BID
		ra.AgreementStart = d1
		ra.AgreementStop = d2
		ra.PossessionStart = d1
		ra.PossessionStop = d2
		ra.RentStart = d1
		ra.RentStop = d2
		ra.RentCycleEpoch = epoch
		ra.UnspecifiedAdults = rand.Int63n(4)
		ra.UnspecifiedChildren = rand.Int63n(3)
		ra.Renewal = 2
		_, err := rlib.InsertRentalAgreement(ctx, &ra)
		if err != nil {
			return err
		}
		//-------------------------------------------------------
		// Create the LedgerMarker for this Rental Agreement
		// 2 weeks prior to the contract commencement
		// just in case some preliminary accounting is ever
		// required...
		//-------------------------------------------------------
		var lm rlib.LedgerMarker
		lm.RAID = ra.RAID
		lm.Dt = d1.AddDate(0, 0, -14)
		_, err = rlib.InsertLedgerMarker(ctx, &lm)
		if err != nil {
			return err
		}

		RIDMktRate, err := rlib.GetRentableMarketRate(ctx, &dbConf.xbiz, RID, &d1, &d2)
		if err != nil {
			return err
		}

		//-------------------------------------
		// Assign Rentable
		//-------------------------------------
		var rar rlib.RentalAgreementRentable
		if RID > MaxRID {
			continue
		}
		rtr, err := rlib.GetRentableTypeRefForDate(ctx, RID, &d1)
		if err != nil {
			return err
		}
		rar.BID = BID
		rar.RAID = ra.RAID
		rar.RARDtStart = d1
		rar.RARDtStop = d2
		rar.RID = RID
		rar.ContractRent = RIDMktRate
		_, err = rlib.InsertRentalAgreementRentable(ctx, &rar)
		if err != nil {
			return err
		}
		//----------------------------------------------------------
		// Create the LedgerMarker for this RID, RAID combination
		//----------------------------------------------------------
		lm.RID = RID
		_, err = rlib.InsertLedgerMarker(ctx, &lm)
		if err != nil {
			return err
		}

		//-------------------------------------
		// Assign Payor
		//-------------------------------------
		TCID := int64(1) + int64(i)%MaxTCID // wrap around as needed
		var rap rlib.RentalAgreementPayor
		rap.BID = BID
		rap.DtStart = d1
		rap.DtStop = d2
		rap.RAID = ra.RAID
		rap.TCID = TCID
		_, err = rlib.InsertRentalAgreementPayor(ctx, &rap)
		if err != nil {
			return err
		}

		//-------------------------------------
		// Assign User
		//-------------------------------------
		var rau rlib.RentableUser
		rau.BID = BID
		rau.RID = RID
		rau.DtStart = d1
		rau.DtStop = d2
		rau.TCID = TCID
		_, err = rlib.InsertRentableUser(ctx, &rau)
		if err != nil {
			return err
		}

		//-------------------------------------
		// Generate Rent Assessments
		//-------------------------------------
		var asmRent rlib.Assessment
		var asmSecDep rlib.Assessment
		asmRent.BID = BID
		asmRent.RID = RID
		asmRent.RAID = ra.RAID
		asmRent.Amount = RIDMktRate
		asmRent.RentCycle = dbConf.xbiz.RT[rtr.RTID].RentCycle
		asmRent.ProrationCycle = dbConf.xbiz.RT[rtr.RTID].Proration
		asmRent.Start = epoch
		asmRent.Stop = d2
		asmRent.ARID = dbConf.ARIDrent
		be := bizlogic.InsertAssessment(ctx, &asmRent, 1)
		if be != nil {
			return bizlogic.BizErrorListToError(be)
		}

		//----------------------------------------------------------
		// Add prorated rent for initial month if start date is not
		// the epoch date.
		//----------------------------------------------------------
		// rlib.Console("d1.Day() = %d, epoch.Day() = %d\n", d1.Day(), epoch.Day())
		if d1.Day() > epoch.Day() {
			var a rlib.Assessment
			td2 := time.Date(d1.Year(), d1.Month(), epoch.Day(), d1.Hour(), d1.Minute(), d1.Second(), d1.Nanosecond(), d1.Location())
			td2 = rlib.NextPeriod(&td2, asmRent.RentCycle)
			a.BID = BID
			a.RID = RID
			a.RAID = ra.RAID
			tot, np, tp := rlib.SimpleProrateAmount(RIDMktRate, asmRent.RentCycle, asmRent.ProrationCycle, &d1, &td2, &epoch)
			a.Amount = tot
			if a.Amount < RIDMktRate {
				a.Comment = fmt.Sprintf("prorated for %d of %d %s", np, tp, rlib.ProrationUnits(asmRent.ProrationCycle))
			}
			a.RentCycle = rlib.RECURNONE
			a.ProrationCycle = rlib.RECURNONE
			a.Start = d1
			a.Stop = d1
			a.ARID = dbConf.ARIDrent
			be = bizlogic.InsertAssessment(ctx, &a, 1)
			if be != nil {
				return bizlogic.BizErrorListToError(be)
			}
		}

		//-------------------------------------
		// Generate SecDep Assessments
		//-------------------------------------
		asmSecDep.BID = BID
		asmSecDep.RID = RID
		asmSecDep.RAID = ra.RAID
		asmSecDep.Amount = RIDMktRate * float64(2.0)
		asmSecDep.RentCycle = rlib.RECURNONE
		asmSecDep.ProrationCycle = rlib.RECURNONE
		asmSecDep.Start = d1
		asmSecDep.Stop = d1
		asmSecDep.ARID = dbConf.ARIDsecdep
		be = bizlogic.InsertAssessment(ctx, &asmSecDep, 1)
		if be != nil {
			return bizlogic.BizErrorListToError(be)
		}

		RID++
		if i+1 < dbConf.RACount && RID > MaxRID {
			fmt.Printf("Halting Rental Agreement creation at RAID = %d because all Rentables are rented\n", ra.RAID)
			break
		}
	}
	return nil
}

// CreateDeposits generates a deposits for the receipts
//-----------------------------------------------------------------------------
func CreateDeposits(ctx context.Context, dbConf *GenDBConf) error {
	// rlib.Console("Entered: CreateDeposits\n")
	var SecDeps = []int64{}
	var OpDeps = []int64{}
	bid := dbConf.BIZ[0].BID
	qry := fmt.Sprintf("SELECT %s FROM Receipt WHERE BID=%d AND DID=0", rlib.RRdb.DBFields["Receipt"], bid)
	// rlib.Console("query = %q\n", qry)
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()
	//------------------------------------------------------------------------
	// Collect the payments, separate security deposits from other payments
	//------------------------------------------------------------------------
	SecDepAmt := float64(0)
	OpDepAmt := float64(0)
	for i := 0; rows.Next(); i++ {
		var a rlib.Receipt
		if err = rlib.ReadReceipts(rows, &a); err != nil {
			return err
		}
		// rlib.Console("Receipt: %d  DEPID = %d\n", a.RCPTID, a.DEPID)
		if a.DEPID == dbConf.OpDepository {
			SecDeps = append(SecDeps, a.RCPTID)
			SecDepAmt += a.Amount
		} else {
			OpDeps = append(OpDeps, a.RCPTID)
			OpDepAmt += a.Amount
		}
	}
	if rows.Err() != nil {
		return rows.Err()
	}

	//----------------------
	// make the deposits
	//----------------------
	if SecDepAmt > float64(0) {
		var b = rlib.Deposit{
			BID:    bid,
			DEPID:  dbConf.SecDepDepository,
			DPMID:  int64(1),
			Dt:     dbConf.DtStart,
			Amount: SecDepAmt,
		}
		// rlib.Console("Security Deposit amount = %8.2f\n", SecDepAmt)

		if e := bizlogic.SaveDeposit(ctx, &b, SecDeps); len(e) > 0 {
			bizlogic.PrintBizErrorList(e)
			return bizlogic.BizErrorListToError(e)
		}
	}

	if OpDepAmt > float64(0) {
		var c = rlib.Deposit{
			BID:    bid,
			DEPID:  dbConf.OpDepository,
			DPMID:  int64(1),
			Dt:     dbConf.DtStart,
			Amount: OpDepAmt,
		}
		// rlib.Console("Op amount = %8.2f\n", OpDepAmt)
		if e := bizlogic.SaveDeposit(ctx, &c, OpDeps); len(e) > 0 {
			bizlogic.PrintBizErrorList(e)
			return bizlogic.BizErrorListToError(e)
		}
	}

	return nil
}

// CreateTaskLists creates the initial task list, associates it with the business
// as the ClosePeriod TaskList, and creates all past instances
//-----------------------------------------------------------------------------
func CreateTaskLists(ctx context.Context, dbConf *GenDBConf) error {
	rlib.Console("Entered: CreateTaskLists\n")
	TLDID := int64(1)
	BID := TLDID

	//----------------------------------------------------------
	// We will need the descriptor to determine the epoch...
	//----------------------------------------------------------
	tld, err := rlib.GetTaskListDefinition(ctx, TLDID)
	if err != nil {
		return err
	}

	//----------------------------------------------------------
	// Pivot day is the config file's start date...
	//----------------------------------------------------------
	pivot := time.Date(dbConf.DtStart.Year(), dbConf.DtStart.Month(), dbConf.DtStart.Day(),
		tld.Epoch.Hour(), tld.Epoch.Minute(), 0, 0, time.UTC)

	//----------------------------------------------------------
	// Create the first instance.
	//----------------------------------------------------------
	tl, err := rlib.CreateTaskListInstance(ctx, TLDID, 0, &pivot)
	if err != nil {
		return err
	}

	//----------------------------------------------------------
	// Update the business to use this tasklist as the company
	// ClosePeriod TaskList
	//----------------------------------------------------------
	var biz rlib.Business
	if err = rlib.GetBusiness(ctx, BID, &biz); err != nil {
		return err
	}
	biz.ClosePeriodTLID = tl.TLID
	if err = rlib.UpdateBusiness(ctx, &biz); err != nil {
		return err
	}

	//----------------------------------------------------------
	// Create any past instances
	//----------------------------------------------------------
	dtNext := rlib.NextInstance(&pivot, tl.Cycle)
	now := time.Now()
	ptlid := tl.TLID
	for {
		pivot = dtNext
		rlib.Console("loop:  pivot = %s\n", pivot.Format(rlib.RRDATETIMERPTFMT))
		tl, err := rlib.CreateTaskListInstance(ctx, TLDID, ptlid, &pivot)
		if err != nil {
			return err
		}
		dtNext := rlib.NextInstance(&pivot, tl.Cycle)
		rlib.Console("loop: rlib.NextInstance(&pivot, tl.Cycle) --> dtNext = %s\n", dtNext.Format(rlib.RRDATETIMERPTFMT))
		rlib.Console("loop: dtNext = %s\n", dtNext.Format(rlib.RRDATETIMERPTFMT))
		if dtNext.After(now) {
			break
		}
	}

	return nil
}
