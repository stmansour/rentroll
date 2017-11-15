package main

import (
	"fmt"
	"math/rand"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"time"
)

type tableMaker struct {
	Name    string
	Handler func(*GenDBConf) error
}

var iRID = int64(1)

var handlers = []tableMaker{
	{"People", createTransactants},
	{"Rentable Types and Rentables", createRentableTypesAndRentables},
	{"Rental Agreements", createRentalAgreements},
}

// GenerateDB is the RentRoll Database generator. It creates a
// database for testing based on parameters in the supplied configuration
// context ctx.
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
//  ctx - context; the configuration data
//
// RETURNS:
//  any errors encountered
//-----------------------------------------------------------------------------
func GenerateDB(ctx *GenDBConf) error {
	var ar rlib.AR
	var err error
	BID := ctx.BIZ[0].BID
	rlib.InitBizInternals(BID, &ctx.xbiz) // used by handlers
	if ctx.ARIDrent == 0 {
		ar, err = rlib.GetARByName(BID, "Rent Non-Taxable")
		ctx.ARIDrent = ar.ARID
		if err != nil {
			return err
		}
	}
	if ctx.ARIDsecdep == 0 {
		ar, err = rlib.GetARByName(BID, "Security Deposit Assessment")
		ctx.ARIDsecdep = ar.ARID
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(handlers); i++ {
		if err := handlers[i].Handler(ctx); err != nil {
			return err
		}
	}
	return nil
}

func randomPhoneNumber() string {
	return fmt.Sprintf("(%d) %3d-%04d", 100+rand.Intn(899), 100+rand.Intn(899), rand.Intn(9999))
}

// createTransactants
//-----------------------------------------------------------------------------
func createTransactants(ctx *GenDBConf) error {
	for i := 0; i < ctx.PeopleCount; i++ {
		var t rlib.Transactant
		t.BID = ctx.BIZ[0].BID
		t.FirstName = fmt.Sprintf("John%04d", i)
		t.MiddleName = "Q"
		t.LastName = fmt.Sprintf("Doe%04d", i)
		t.PreferredName = fmt.Sprintf("J%04d", i)
		t.CellPhone = randomPhoneNumber()
		t.PrimaryEmail = fmt.Sprintf("jdoe%04d@example.com", i)
		_, err := rlib.InsertTransactant(&t)
		if err != nil {
			return err
		}
	}
	return nil
}

// createRentableTypesAndRentables
//-----------------------------------------------------------------------------
func createRentableTypesAndRentables(ctx *GenDBConf) error {
	var err error
	fmt.Printf("Rentable Types: %d\n", len(ctx.RT))
	for i := 0; i < len(ctx.RT); i++ {
		fmt.Printf("Rentable Type %d.  Rentable Count = %d\n", i, ctx.RT[i].Count)
		var rt rlib.RentableType
		rt.BID = 1
		rt.Style = fmt.Sprintf("ST%03d", i)
		rt.Name = fmt.Sprintf("RType%03d", i)
		rt.RentCycle = ctx.RT[i].RentCycle
		rt.Proration = ctx.RT[i].ProrateCycle
		rt.GSRPC = ctx.RT[i].ProrateCycle
		rt.ManageToBudget = 1
		_, err = rlib.InsertRentableType(&rt)
		if err != nil {
			return err
		}

		var mr rlib.RentableMarketRate
		mr.DtStart = ctx.DtBOT
		mr.DtStop = ctx.DtEOT
		mr.MarketRate = ctx.RT[i].MarketRate
		mr.RTID = rt.RTID
		if err = rlib.InsertRentableMarketRates(&mr); err != nil {
			return err
		}

		if err = createRentables(ctx, &ctx.RT[i], &mr, rt.RTID); err != nil {
			return err
		}
	}
	return nil
}

// createRentables
//-----------------------------------------------------------------------------
func createRentables(ctx *GenDBConf, rt *RType, mr *rlib.RentableMarketRate, RTID int64) error {
	for i := 0; i < rt.Count; i++ {
		var r rlib.Rentable
		r.RID = iRID
		r.BID = ctx.BIZ[0].BID
		r.RentableName = fmt.Sprintf("Rentable%03d", iRID)
		_, err := rlib.InsertRentable(&r)
		if err != nil {
			return err
		}

		var rtr rlib.RentableTypeRef
		rtr.DtStart = ctx.DtBOT
		rtr.DtStop = ctx.DtEOT
		rtr.BID = ctx.BIZ[0].BID
		rtr.RTID = RTID
		rtr.RID = r.RID
		if err = rlib.InsertRentableTypeRef(&rtr); err != nil {
			return err
		}

		var rs rlib.RentableStatus
		rs.DtStart = ctx.DtBOT
		rs.DtStop = ctx.DtEOT
		rs.BID = ctx.BIZ[0].BID
		rs.RID = r.RID
		rs.LeaseStatus = rlib.LEASESTATUSleased
		rs.UseStatus = rlib.USESTATUSinService
		if err = rlib.InsertRentableStatus(&rs); err != nil {
			return err
		}
		iRID++
	}
	return nil
}

// createRentalAgreements
//-----------------------------------------------------------------------------
func createRentalAgreements(ctx *GenDBConf) error {
	BID := ctx.BIZ[0].BID
	rlib.GetXBusiness(BID, &ctx.xbiz)
	d1 := time.Date(ctx.DtStart.Year(), ctx.DtStart.Month(), ctx.DtStart.Day(), 0, 0, 0, 0, time.UTC)
	d2 := d1.AddDate(2, 0, 0)
	epoch := time.Date(ctx.DtStart.Year(), ctx.DtStart.Month(), 1, 0, 0, 0, 0, time.UTC)
	if ctx.DtStart.Day() > 1 {
		epoch = epoch.AddDate(0, 1, 0)

	}
	MaxRID := int64(rlib.GetCountByTableName("Rentable", BID))
	MaxTCID := int64(rlib.GetCountByTableName("Transactant", BID))
	RID := int64(1)

	for i := 0; i < ctx.RACount; i++ {
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
		_, err := rlib.InsertRentalAgreement(&ra)
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
		err = rlib.InsertLedgerMarker(&lm)
		if err != nil {
			return err
		}

		RIDMktRate := rlib.GetRentableMarketRate(&ctx.xbiz, RID, &d1, &d2)

		//-------------------------------------
		// Assign Rentable
		//-------------------------------------
		var rar rlib.RentalAgreementRentable
		if RID > MaxRID {
			continue
		}
		rtr := rlib.GetRentableTypeRefForDate(RID, &d1)
		rar.BID = BID
		rar.RAID = ra.RAID
		rar.RARDtStart = d1
		rar.RARDtStop = d2
		rar.RID = RID
		rar.ContractRent = RIDMktRate
		_, err = rlib.InsertRentalAgreementRentable(&rar)

		//----------------------------------------------------------
		// Create the LedgerMarker for this RID, RAID combination
		//----------------------------------------------------------
		lm.RID = RID
		err = rlib.InsertLedgerMarker(&lm)
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
		_, err = rlib.InsertRentalAgreementPayor(&rap)

		//-------------------------------------
		// Assign User
		//-------------------------------------
		var rau rlib.RentableUser
		rau.BID = BID
		rau.RID = RID
		rau.DtStart = d1
		rau.DtStop = d2
		rau.TCID = TCID
		err = rlib.InsertRentableUser(&rau)

		//-------------------------------------
		// Generate Rent Assessments
		//-------------------------------------
		var asmRent rlib.Assessment
		var asmSecDep rlib.Assessment
		asmRent.BID = BID
		asmRent.RID = RID
		asmRent.RAID = ra.RAID
		asmRent.Amount = RIDMktRate
		asmRent.RentCycle = ctx.xbiz.RT[rtr.RTID].RentCycle
		asmRent.ProrationCycle = ctx.xbiz.RT[rtr.RTID].Proration
		asmRent.Start = epoch
		asmRent.Stop = d2
		asmRent.ARID = ctx.ARIDrent
		be := bizlogic.InsertAssessment(&asmRent, 1)
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
			a.ARID = ctx.ARIDrent
			be = bizlogic.InsertAssessment(&a, 1)
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
		asmSecDep.RentCycle = ctx.xbiz.RT[rtr.RTID].RentCycle
		asmSecDep.ProrationCycle = ctx.xbiz.RT[rtr.RTID].Proration
		asmSecDep.Start = epoch
		asmSecDep.Stop = d2
		asmSecDep.ARID = ctx.ARIDsecdep
		be = bizlogic.InsertAssessment(&asmSecDep, 1)
		if be != nil {
			return bizlogic.BizErrorListToError(be)
		}

		RID++
		if RID > MaxRID {
			fmt.Printf("Halting Rental Agreement creation at RAID = %d because all Rentables are rented\n", ra.RAID)
			break
		}
	}
	return nil
}
