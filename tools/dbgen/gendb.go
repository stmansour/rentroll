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
	{"HotelBookings", HotelBookings},
}

var noClose = rlib.ClosePeriod{
	Dt:           rlib.TIME0,
	OpenPeriodDt: rlib.TIME0,
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
			return fmt.Errorf("could not find Payment Type with name %q", dbConf.PTypeCheckName)
		}
		dbConf.PTypeCheck = pt.PMTID
	}

	//---------------------------------------
	// Now spin through all the handlers...
	//---------------------------------------
	for i := 0; i < len(handlers); i++ {
		rlib.Console("%d. %s\n", i, handlers[i].Name)
		if err := handlers[i].Handler(ctx, dbConf); err != nil {
			return err
		}
	}

	return nil
}

// createRandomCar returns a Vehicle struct filled out with some random
// car information
//-----------------------------------------------------------------------------
func createRandomCar(t *rlib.Transactant, dbConf *GenDBConf) rlib.Vehicle {
	var v rlib.Vehicle
	v.TCID = t.TCID
	v.BID = t.BID
	v.VehicleType = "car"
	j := IG.Rand.Intn(len(IG.Cars))
	v.VehicleMake = IG.Cars[j].Make
	v.VehicleModel = IG.Cars[j].Model
	v.VehicleYear = int64(IG.Cars[j].Year)
	v.VIN = GenerateRandomVIN()
	v.VehicleColor = GenerateRandomCarColor()
	v.LicensePlateState = GenerateRandomState()
	v.LicensePlateNumber = GenerateRandomLicensePlate()
	v.ParkingPermitNumber = fmt.Sprintf("%07d", IG.Rand.Intn(10000000))
	v.DtStart = dbConf.DtStart
	v.DtStop = dbConf.DtStop
	return v
}
func createRandomDog(t *rlib.Transactant, dbConf *GenDBConf) rlib.Pet {
	var p rlib.Pet
	p.TCID = t.TCID
	p.BID = t.BID
	p.Type = "dog"
	p.Breed = GenerateRandomDog()
	p.Color = GenerateRandomDogColor()
	p.Weight = float64(10 + IG.Rand.Intn(50))
	p.Name = GenerateRandomDogName()
	p.DtStart = dbConf.DtStart
	p.DtStop = dbConf.DtStop
	return p
}

func createRandomCat(t *rlib.Transactant, dbConf *GenDBConf) rlib.Pet {
	var p rlib.Pet
	p.TCID = t.TCID
	p.BID = t.BID
	p.Type = "cat"
	p.Breed = GenerateRandomCat()
	p.Color = GenerateRandomCatColor()
	p.Weight = float64(5 + IG.Rand.Intn(20))
	p.Name = GenerateRandomCatName()
	p.DtStart = dbConf.DtStart
	p.DtStop = dbConf.DtStop
	return p
}

func createRandomPet(t *rlib.Transactant, dbConf *GenDBConf) {
}

// createTransactants
//-----------------------------------------------------------------------------
func createTransactants(ctx context.Context, dbConf *GenDBConf) error {
	funcname := "createTransactants"
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

		//-------------------------------------
		// TRANSACTION...
		//-------------------------------------
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

		//-------------------------------------
		// USER...
		//-------------------------------------
		now := time.Now()
		ecfirst := GenerateRandomFirstName()
		eclast := GenerateRandomLastName()
		var u = rlib.User{
			TCID: t.TCID,
			BID:  t.BID,
			// Points:                  int64(IG.Rand.Intn(5000)),
			DateofBirth:               now.AddDate(-20-IG.Rand.Intn(45), 0, -IG.Rand.Intn(365)),
			EmergencyContactName:      ecfirst + " " + eclast,
			EmergencyContactAddress:   GenerateRandomAddress() + "," + GenerateRandomCity() + "," + GenerateRandomState() + " " + fmt.Sprintf("%05d", rand.Intn(100000)),
			EmergencyContactTelephone: GenerateRandomPhoneNumber(),
			EmergencyContactEmail:     GenerateRandomEmail(eclast, ecfirst),
			AlternateEmailAddress:     GenerateRandomAddress() + "," + GenerateRandomCity() + "," + GenerateRandomState() + " " + fmt.Sprintf("%05d", rand.Intn(100000)),
			EligibleFutureUser:        IG.Rand.Intn(2) > 0,
			Industry:                  IG.Industries.S[IG.Rand.Intn(len(IG.Industries.S))].SLSID,
			SourceSLSID:               IG.HowFound.S[IG.Rand.Intn(len(IG.HowFound.S))].SLSID,
		}

		_, err = rlib.InsertUser(ctx, &u)
		if err != nil {
			return err
		}

		//-------------------------------------
		// PAYOR...
		//-------------------------------------
		var p = rlib.Payor{
			TCID:                t.TCID,
			BID:                 t.BID,
			CreditLimit:         float64(IG.Rand.Intn(30000)),
			TaxpayorID:          fmt.Sprintf("%08d", IG.Rand.Intn(10000000)),
			EligibleFuturePayor: true,
			DriversLicense:      GenerateRandomDriversLicense(),
			GrossIncome:         float64(10000 + IG.Rand.Intn(140000)),
		}
		_, err = rlib.InsertPayor(ctx, &p)
		if err != nil {
			return err
		}

		//-------------------------------------
		// PROSPECT...
		//-------------------------------------
		ec := rlib.Stripchars(GenerateRandomCity(), ".@ ")
		cmp := rlib.Stripchars(t.CompanyName, ".@ ")
		var pr = rlib.Prospect{
			TCID:                     t.TCID,
			BID:                      t.BID,
			CompanyAddress:           GenerateRandomAddress(),
			CompanyCity:              ec,
			CompanyState:             GenerateRandomState(),
			CompanyPostalCode:        fmt.Sprintf("%05d", rand.Intn(100000)),
			CompanyEmail:             GenerateRandomEmail(ec, cmp),
			CompanyPhone:             GenerateRandomPhoneNumber(),
			Occupation:               GenerateRandomOccupation(),
			FLAGS:                    0,
			CurrentAddress:           GenerateRandomOneLineAddress(),
			CurrentLandLordName:      GenerateRandomName(),
			CurrentLandLordPhoneNo:   GenerateRandomPhoneNumber(),
			CurrentReasonForMoving:   IG.WhyLeaving.S[IG.Rand.Intn(len(IG.WhyLeaving.S))].SLSID,
			CurrentLengthOfResidency: GenerateRandomDurationString(),
			PriorAddress:             GenerateRandomOneLineAddress(),
			PriorLandLordName:        GenerateRandomName(),
			PriorLandLordPhoneNo:     GenerateRandomPhoneNumber(),
			PriorReasonForMoving:     IG.WhyLeaving.S[IG.Rand.Intn(len(IG.WhyLeaving.S))].SLSID,
			PriorLengthOfResidency:   GenerateRandomDurationString(),
			ThirdPartySource:         GenerateRandomName(),
		}
		_, err = rlib.InsertProspect(ctx, &pr)
		if err != nil {
			return err
		}

		//-----------------------------------------
		// create vehicles.
		// X% chance that there will be a vehicle
		// Y% chance that will be 2 vehicles
		//-----------------------------------------
		if IG.Rand.Intn(100) < 95 { // x%
			vcount := 1
			if IG.Rand.Intn(100) < 5 { // y%
				vcount++
			}
			for j := 0; j < vcount; j++ {
				v := createRandomCar(&t, dbConf)
				_, err = rlib.InsertVehicle(ctx, &v)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return err
				}
				// now create the TBind for this vehicl...
				var tb = rlib.TBind{
					BID:            t.BID,
					SourceElemType: rlib.ELEMPERSON,
					SourceElemID:   t.TCID,
					AssocElemType:  rlib.ELEMVEHICLE,
					AssocElemID:    v.VID,
					DtStart:        dbConf.DtStart,
					DtStop:         rlib.ENDOFTIME,
					FLAGS:          0,
				}
				_, err = rlib.InsertTBind(ctx, &tb)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return err
				}
			}
		}

		//--------------------------
		// Create Pets...
		//--------------------------
		if IG.Rand.Intn(100) < 68 { // x%
			vcount := 1
			if IG.Rand.Intn(100) < 5 { // y%
				vcount++
			}
			for j := 0; j < vcount; j++ {
				var p rlib.Pet
				if IG.Rand.Intn(70) < 40 {
					p = createRandomDog(&t, dbConf)
				} else {
					p = createRandomCat(&t, dbConf)
				}
				_, err = rlib.InsertPet(ctx, &p)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return err
				}
				// now create the TBind for this pet...
				var tb = rlib.TBind{
					BID:            t.BID,
					SourceElemType: rlib.ELEMPERSON,
					SourceElemID:   t.TCID,
					AssocElemType:  rlib.ELEMPET,
					AssocElemID:    p.PETID,
					DtStart:        dbConf.DtStart,
					DtStop:         rlib.ENDOFTIME,
					FLAGS:          0,
				}
				_, err = rlib.InsertTBind(ctx, &tb)
				if err != nil {
					rlib.LogAndPrintError(funcname, err)
					return err
				}
			}
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
		ar.DefaultRentCycle = rlib.RECURMONTHLY    // monthly rent cycle
		ar.DefaultProrationCycle = rlib.RECURDAILY // daily proration
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
		rt.FLAGS |= 0x4 /*manage to budget*/
		if dbConf.RT[i].Reserve {
			rt.FLAGS |= 1 << 3 // if 1 then Reserve after RA.  set to 0 for hotels
		}
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
		_, err = rlib.InsertRentableMarketRate(ctx, &mr)
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
	err = createChildRentableTypes(ctx, dbConf)
	return err
}

func createChildRentableTypes(ctx context.Context, dbConf *GenDBConf) error {

	var err error
	var rt rlib.RentableType

	style := "CP000"
	name := "Car Port 000"

	//-------------------------------
	// Default rent AccountRule for
	// this RentableType
	//-------------------------------
	var ar rlib.AR
	ar.BID = dbConf.BIZ[0].BID
	ar.Name = fmt.Sprintf("Rent %s", style)
	ar.Description = fmt.Sprintf("Default rent assessment for rentable type %s", name)
	ar.ARType = 0                          // Assessment
	ar.DebitLID = 9                        // Acct# 12001 - RentRoll Receivables
	ar.CreditLID = 18                      // Acct# 41001 - Gross Scheduled Rent non-taxable
	ar.DefaultAmount = dbConf.CPMarketRate // default rent amount
	ar.FLAGS = (1 << 2) | (1 << 4)         // RAID rqd, is Rent
	ar.DtStart = rlib.TIME0                // make this rule "forever"
	ar.DtStop = rlib.ENDOFTIME             // make this rule "forever"
	_, err = rlib.InsertAR(ctx, &ar)
	if err != nil {
		return err
	}

	//-----------------------------
	// RENTABLE TYPE
	//-----------------------------
	rt.BID = dbConf.BIZ[0].BID
	rt.Style = style
	rt.Name = name
	rt.RentCycle = dbConf.CPRentCycle
	rt.Proration = dbConf.CPProrateCycle
	rt.GSRPC = dbConf.CPProrateCycle
	rt.FLAGS |= 0x2 /*child*/ | 0x4 /*manage to budget*/
	rt.ARID = ar.ARID
	_, err = rlib.InsertRentableType(ctx, &rt)
	if err != nil {
		return err
	}

	//-----------------------------
	// RENTABLE MARKET RATE
	//-----------------------------
	var mr rlib.RentableMarketRate
	mr.DtStart = dbConf.DtBOT
	mr.DtStop = dbConf.DtEOT
	mr.MarketRate = dbConf.CPMarketRate
	mr.RTID = rt.RTID
	_, err = rlib.InsertRentableMarketRate(ctx, &mr)
	if err != nil {
		return err
	}

	for i := 0; i < dbConf.Carports; i++ {
		var r rlib.Rentable

		//-----------------------------
		// RENTABLE
		//-----------------------------
		r.BID = rt.BID
		r.RentableName = fmt.Sprintf("CP%03d", i+1)
		errlist := bizlogic.InsertRentable(ctx, &r)
		if errlist != nil {
			return bizlogic.BizErrorListToError(errlist)
		}

		//-----------------------------
		// RENTABLE TYPE REF
		//-----------------------------
		var rtr rlib.RentableTypeRef
		rtr.DtStart = dbConf.DtBOT
		rtr.DtStop = dbConf.DtEOT
		rtr.BID = rt.BID
		rtr.RTID = rt.RTID
		rtr.RID = r.RID
		_, err = rlib.InsertRentableTypeRef(ctx, &rtr)
		if err != nil {
			return err
		}

		//-----------------------------
		// RENTABLE USE STATUS
		//-----------------------------
		var rs rlib.RentableUseStatus
		rs.DtStart = dbConf.DtBOT
		rs.DtStop = dbConf.DtEOT
		rs.BID = dbConf.BIZ[0].BID
		rs.RID = r.RID
		rs.UseStatus = rlib.USESTATUSready
		_, err = rlib.InsertRentableUseStatus(ctx, &rs)
		if err != nil {
			return err
		}
		//-----------------------------
		// RENTABLE USE TYPE
		//-----------------------------
		var rut rlib.RentableUseType
		rut.DtStart = dbConf.DtBOT
		rut.DtStop = dbConf.DtEOT
		rut.BID = dbConf.BIZ[0].BID
		rut.RID = r.RID
		rut.UseType = rlib.USETYPEstandard
		_, err = rlib.InsertRentableUseType(ctx, &rut)
		if err != nil {
			return err
		}
		//-----------------------------
		// RENTABLE LEASE STATUS
		//-----------------------------
		var rl rlib.RentableLeaseStatus
		rl.DtStart = dbConf.DtBOT
		rl.DtStop = dbConf.DtEOT
		rl.BID = dbConf.BIZ[0].BID
		rl.RID = r.RID
		rl.LeaseStatus = rlib.LEASESTATUSnotleased
		_, err = rlib.InsertRentableLeaseStatus(ctx, &rl)
		if err != nil {
			return err
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

		var rs rlib.RentableUseStatus
		rs.DtStart = dbConf.DtBOT
		rs.DtStop = dbConf.DtEOT
		rs.BID = dbConf.BIZ[0].BID
		rs.RID = r.RID
		rs.UseStatus = rlib.USESTATUSready
		_, err = rlib.InsertRentableUseStatus(ctx, &rs)
		if err != nil {
			return err
		}

		var rut rlib.RentableUseType
		rut.DtStart = dbConf.DtBOT
		rut.DtStop = dbConf.DtEOT
		rut.BID = dbConf.BIZ[0].BID
		rut.RID = r.RID
		rut.UseType = rlib.USETYPEstandard
		_, err = rlib.InsertRentableUseType(ctx, &rut)
		if err != nil {
			return err
		}

		var rl rlib.RentableLeaseStatus
		rl.DtStart = dbConf.DtBOT
		rl.DtStop = dbConf.DtEOT
		rl.BID = dbConf.BIZ[0].BID
		rl.RID = r.RID
		rl.LeaseStatus = rlib.LEASESTATUSnotleased
		_, err = rlib.InsertRentableLeaseStatus(ctx, &rl)
		if err != nil {
			return err
		}
		// rlib.Console("Just saved LeasedStatus\n")

		iRID++
	}
	return nil
}

// createReceipts reads all assessments and creates a separate receipt for
// each one.
//-----------------------------------------------------------------------------
func createReceipts(ctx context.Context, dbConf *GenDBConf) error {
	//                                                                it's an epoch but nonrecur     it's an instance
	qry := fmt.Sprintf("SELECT %s FROM Assessments WHERE BID=%d AND ((PASMID=0 AND RentCycle=0) OR PASMID != 0)",
		rlib.RRdb.DBFields["Assessments"], dbConf.BIZ[0].BID)
	rows, err := rlib.RRdb.Dbrr.Query(qry)
	if err != nil {
		return err
	}
	defer rows.Close()

	//--------------------------------------
	// Spin through all the assessments
	//--------------------------------------
	for i := 0; rows.Next(); i++ {
		var a rlib.Assessment
		err = rlib.ReadAssessments(rows, &a)
		if err != nil {
			return err
		}

		//---------------------------------------------------------------------
		// If we have been asked to miss some percentage of payments, roll
		// the weighted dice and see if we pay or not...
		//---------------------------------------------------------------------
		if dbConf.RandomizePayments && dbConf.RRand.Intn(100) < dbConf.RandMissPayment {
			continue // some randomness - a missed payment
		}
		// so if we get to this point, we pay

		//---------------------------------------------------------------------
		// Identify the depository where this check will be deposited...
		//---------------------------------------------------------------------
		depid := dbConf.OpDepository
		if a.ARID == dbConf.ARIDsecdep {
			depid = dbConf.SecDepDepository
		}

		//---------------------------------------------------------------------
		// Create the payment...
		//---------------------------------------------------------------------
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
		rcpt.Comment = fmt.Sprintf("payment for %s", a.IDtoShortString())
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

func getStartDateFromRange(dbConf *GenDBConf) time.Time {
	return dbConf.DtStart
}

func getStopDateFromRange(dbConf *GenDBConf, dtmpStart *time.Time) time.Time {
	dt := dbConf.DtStop
	// validate that dt is after dtmpStart
	if !dt.After(*dtmpStart) {
		fmt.Printf("FATAL ERROR: in getStopDateFromRange - stop date must be after start date\n")
		fmt.Printf("             dtmpStart = %s,   stop = %s\n", rlib.ConsoleDate(dtmpStart), rlib.ConsoleDate(&dt))
		panic("FATAL")
	}
	return dt
}

// createRentalAgreements - this creates long term rental agreements -- for
// apartment renters.
//-----------------------------------------------------------------------------
func createRentalAgreements(ctx context.Context, dbConf *GenDBConf) error {
	// rlib.Console("entered: createRentalAgreements\n")
	BID := dbConf.BIZ[0].BID
	err := rlib.GetXBusiness(ctx, BID, &dbConf.xbiz)
	if err != nil {
		return err
	}
	dtmpStart := getStartDateFromRange(dbConf)
	dtmpStop := getStopDateFromRange(dbConf, &dtmpStart)
	d1 := time.Date(dtmpStart.Year(), dtmpStart.Month(), dtmpStart.Day(), 0, 0, 0, 0, time.UTC)
	epoch := time.Date(dtmpStart.Year(), dtmpStart.Month(), 1, 0, 0, 0, 0, time.UTC)
	if dtmpStart.Day() > 1 {
		epoch = epoch.AddDate(0, 1, 0)

	}
	d2 := dtmpStop
	if d2.Day() != 1 {
		d2 = time.Date(d2.Year(), d2.Month(), d2.Day(), 0, 0, 0, 0, time.UTC)
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
	now := time.Now()
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
		ra.FLAGS = 0x4 /*status = approved*/ | (1 << 4) /* approver1 */ | (1 << 5) /*approver2*/
		ra.ApplicationReadyUID = GenerateValidUID()
		ra.ApplicationReadyDate = dbConf.DtStart
		ra.Approver1 = GenerateValidUID()
		ra.DeclineReason1 = 0
		ra.DecisionDate1 = dbConf.DtStart
		ra.Approver2 = GenerateValidUID()
		ra.DeclineReason2 = 0
		ra.DecisionDate2 = dbConf.DtStart
		ra.MoveInUID = GenerateValidUID()
		ra.MoveInDate = dbConf.DtStart
		ra.DocumentDate = dbConf.DtStart
		ra.ActiveUID = GenerateValidUID()
		ra.ActiveDate = dbConf.DtStart
		ra.OtherPreferences = ""
		ra.FollowUpDate = now.AddDate(0, 0, 2)
		ra.CSAgent = GenerateValidUID()
		ra.Outcome = 0

		// rlib.Console("createRentalAgreements: A\n")
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
		lm.BID = ra.BID
		lm.RAID = ra.RAID
		lm.State = rlib.LMINITIAL
		lm.Dt = d1.AddDate(0, 0, -14)
		_, err = rlib.InsertLedgerMarker(ctx, &lm)
		if err != nil {
			return err
		}

	tryRID:
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
		//-----------------------------------------------------------
		// If the rentable type is for a hotel, we don't want to
		// use it here.
		//-----------------------------------------------------------
		// rlib.Console("createRentalAgreements: B, rtr.RTID = %d\n", rtr.RTID)
		if dbConf.xbiz.RT[rtr.RTID].FLAGS&(1<<3) == 0 {
			RID++
			if RID > MaxRID {
				fmt.Printf("Halting Rental Agreement creation at RAID = %d because no more Rentables of RentableType where FLAGS (1<<3) is set.\n", ra.RAID)
				break
			}
			goto tryRID
		}

		// rlib.Console("createRentalAgreements: C\n")

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
		// set its use and lease status...
		//----------------------------------------------------------
		// rlib.Console("createRentalAgreements: Prepare use status:  BID=%d, RID=%d, rs=%d, %s\n", BID, RID, rlib.USESTATUSinService, rlib.ConsoleDRange(&d1, &d2))
		if err = rlib.SetRentableUseStatusAbbr(ctx, BID, RID, rlib.USESTATUSinService, &d1, &d2); err != nil {
			return err
		}

		// rlib.Console("createRentalAgreements: Prepare lease status:  BID=%d, RID=%d, rs=%d, %s\n", BID, RID, rlib.LEASESTATUSleased, rlib.ConsoleDRange(&d1, &d2))
		if err = rlib.SetRentableLeaseStatusAbbr(ctx, BID, RID, rlib.LEASESTATUSleased, &d1, &d2, true); err != nil {
			return err
		}
		// now reserve this rentable forever past the Rental Agreement end date...
		if err = rlib.SetRentableLeaseStatusAbbr(ctx, BID, RID, rlib.LEASESTATUSreserved, &d2, &rlib.ENDOFTIME, false); err != nil {
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
		// Assign Users...
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

		//-------------------------------------------------------
		// Assign all Pets associated with this TCID to
		// this RAID
		//-------------------------------------------------------
		pl, err := rlib.GetPetsByTransactant(ctx, TCID)
		if err != nil {
			return err
		}
		for j := 0; j < len(pl); j++ {
			pl[j].RAID = ra.RAID
			if err = rlib.UpdatePet(ctx, &pl[j]); err != nil {
				return err
			}
		}

		//--------------------------------------------------------
		// Generate assessments for Rent and Security Deposit...
		//--------------------------------------------------------
		var asmRent rlib.Assessment
		asmRent.BID = BID
		asmRent.RID = RID
		asmRent.RAID = ra.RAID
		asmRent.Amount = RIDMktRate
		asmRent.RentCycle = dbConf.xbiz.RT[rtr.RTID].RentCycle
		asmRent.ProrationCycle = dbConf.xbiz.RT[rtr.RTID].Proration
		asmRent.Start = epoch
		asmRent.Stop = d2
		asmRent.ARID = dbConf.ARIDrent
		noClose.ExpandAsmDtStart = asmRent.Start // this is necessary
		noClose.ExpandAsmDtStop = asmRent.Stop   // this is necessary
		be := bizlogic.InsertAssessment(ctx, &asmRent, 1, &noClose)
		if be != nil {
			return bizlogic.BizErrorListToError(be)
		}
		//-------------------------------------
		// Generate SecDep Assessments
		//-------------------------------------
		var asmSecDep rlib.Assessment
		asmSecDep.BID = BID
		asmSecDep.RID = RID
		asmSecDep.RAID = ra.RAID
		asmSecDep.Amount = RIDMktRate * float64(2.0)
		asmSecDep.RentCycle = rlib.RECURNONE
		asmSecDep.ProrationCycle = rlib.RECURNONE
		asmSecDep.Start = d1
		asmSecDep.Stop = d1
		asmSecDep.ARID = dbConf.ARIDsecdep
		be = bizlogic.InsertAssessment(ctx, &asmSecDep, 1, &noClose)
		if be != nil {
			return bizlogic.BizErrorListToError(be)
		}

		//--------------------------------------------------------------------
		// Pet Assessments - Here we handle the Recurring Rent assessments
		// (if any)
		//--------------------------------------------------------------------
		petList, err := rlib.GetPetsByTransactant(ctx, TCID)
		if err != nil {
			return err
		}
		for k := 0; k < len(petList); k++ { // generate the assessment(s) for each pet
			for j := 0; j < len(dbConf.PetFees); j++ {
				//---------------------------------------------------------------------
				// Since we're creating the rental / recurring assessments here, skip
				// this fee if it is a one-time only fee...
				//---------------------------------------------------------------------
				flags := rlib.RRdb.BizTypes[BID].AR[dbConf.PetFees[j].ARID].FLAGS
				//   non-recurring   OR       not rent
				if flags&(1<<6) != 0 || flags&(1<<4) == 0 {
					continue // skip if non-recurring or not rent
				}
				var asm = rlib.Assessment{
					BID:            BID,
					RID:            RID,
					RAID:           ra.RAID,
					Amount:         dbConf.PetFees[j].DefaultAmount,
					RentCycle:      dbConf.xbiz.RT[rtr.RTID].RentCycle,
					ProrationCycle: dbConf.xbiz.RT[rtr.RTID].Proration,
					FLAGS:          1 << 3, // PETID required
					AssocElemType:  rlib.ELEMPET,
					AssocElemID:    petList[k].PETID,
					Start:          epoch,
					Stop:           d2,
					ARID:           dbConf.PetFees[j].ARID,
				}
				noClose.ExpandAsmDtStart = asm.Start                    // this is necessary
				noClose.ExpandAsmDtStop = asm.Stop                      // this is necessary
				be := bizlogic.InsertAssessment(ctx, &asm, 1, &noClose) // bizlogic will not expand it if it is a single instanced assessment
				if be != nil {
					return bizlogic.BizErrorListToError(be)
				}
			}
		}

		//-------------------------------------
		// Recurring Vehicle Assessments
		//-------------------------------------
		vList, err := rlib.GetVehiclesByTransactant(ctx, TCID)
		if err != nil {
			return err
		}
		for k := 0; k < len(vList); k++ { // generate recurring vehicle fees for each vehicle
			for j := 0; j < len(dbConf.VehicleFees); j++ {
				//---------------------------------------------------------------------
				// Since we're creating the rental / recurring assessments here, skip
				// this fee if it is a one-time only fee...
				//
				//     1<<4 -  Is Rent Assessment
				//     1<<6 -  Is non-recur charge
				//     1<<7 -  PETID Required in Assessment
				//---------------------------------------------------------------------
				flags := rlib.RRdb.BizTypes[BID].AR[dbConf.VehicleFees[j].ARID].FLAGS
				//    non-recurring  OR        not rent
				if flags&(1<<6) != 0 || flags&(1<<4) == 0 {
					continue // skip if non-recurring or not rent
				}
				var asm = rlib.Assessment{
					BID:            BID,
					RID:            RID,
					RAID:           ra.RAID,
					Amount:         dbConf.VehicleFees[j].DefaultAmount,
					RentCycle:      dbConf.xbiz.RT[rtr.RTID].RentCycle,
					ProrationCycle: dbConf.xbiz.RT[rtr.RTID].Proration,
					FLAGS:          1 << 4, // VID required
					AssocElemType:  rlib.ELEMVEHICLE,
					AssocElemID:    vList[k].VID,
					Start:          epoch,
					Stop:           d2,
					ARID:           dbConf.VehicleFees[j].ARID,
				}
				noClose.ExpandAsmDtStart = asm.Start                    // this is necessary
				noClose.ExpandAsmDtStop = asm.Stop                      // this is necessary
				be := bizlogic.InsertAssessment(ctx, &asm, 1, &noClose) // bizlogic will not expand it if it is a single instanced assessment
				if be != nil {
					return bizlogic.BizErrorListToError(be)
				}
			}
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
				// a.Comment = fmt.Sprintf("prorated for %d of %d %s", np, tp, rlib.ProrationUnits(asmRent.ProrationCycle))
				a.Comment = rlib.ProrateComment(np, tp, asmRent.ProrationCycle)
			}
			a.RentCycle = rlib.RECURNONE
			a.ProrationCycle = rlib.RECURNONE
			a.Start = d1
			a.Stop = d1
			a.ARID = dbConf.ARIDrent
			noClose.ExpandAsmDtStart = a.Start // this is necessary
			noClose.ExpandAsmDtStop = a.Stop   // this is necessary
			be = bizlogic.InsertAssessment(ctx, &a, 1, &noClose)
			if be != nil {
				return bizlogic.BizErrorListToError(be)
			}

			//--------------------------------------------------------------
			// Handle prorated pet fees and non-recurring fees.
			// Here are the AR.FLAGS bits:
			//
			//     1<<4 -  Is Rent Assessment
			//     1<<6 -  Is non-recur charge
			//     1<<7 -  PETID required
			//     1<<8 -  VID required
			//--------------------------------------------------------------
			for k := 0; k < len(petList); k++ {
				for j := 0; j < len(dbConf.PetFees); j++ {
					//-----------------------------------------------------------------------
					// Check for:
					//    *  a rent assessment AND we're not beginning the on the epoch,
					//    *  a one-time charge
					//-----------------------------------------------------------------------
					flags := rlib.RRdb.BizTypes[BID].AR[dbConf.PetFees[j].ARID].FLAGS
					needProratedRent := flags&(1<<4) != 0 && dbConf.DtStart.Day() != dbConf.Epochs.Monthly.Day()
					oneTimeCharge := flags&(1<<6) != 0
					//-----------------------------------------------------------------------
					// If it's not a prorated Rent charge or a one-time-charge, then skip it
					//-----------------------------------------------------------------------
					if !(needProratedRent || oneTimeCharge) {
						continue // skip if it's rent or it's not a one-time charge
					}
					cmt := ""
					rc := int64(rlib.RECURNONE)
					if needProratedRent {
						rc = dbConf.xbiz.RT[rtr.RTID].RentCycle
					}
					tot, np, tp := rlib.SimpleProrateAmount(dbConf.PetFees[j].DefaultAmount, rc, dbConf.xbiz.RT[rtr.RTID].Proration, &d1, &td2, &epoch)
					if tot < dbConf.PetFees[j].DefaultAmount {
						// cmt = fmt.Sprintf("prorated for %d of %d %s", np, tp, rlib.ProrationUnits(dbConf.xbiz.RT[rtr.RTID].Proration))
						cmt = rlib.ProrateComment(np, tp, dbConf.xbiz.RT[rtr.RTID].Proration)

					}
					var asm = rlib.Assessment{
						BID:            BID,
						RID:            RID,
						RAID:           ra.RAID,
						Amount:         tot,
						RentCycle:      rlib.RECURNONE,
						ProrationCycle: rlib.RECURNONE,
						AssocElemType:  rlib.ELEMPET,
						AssocElemID:    petList[k].PETID,
						Start:          d1,
						Stop:           d1,
						ARID:           dbConf.PetFees[j].ARID,
						Comment:        cmt,
					}
					be := bizlogic.InsertAssessment(ctx, &asm, 1, &noClose)
					if be != nil {
						return bizlogic.BizErrorListToError(be)
					}
				}
			}

			//------------------------------------------------------------------
			// Handle prorated vehicle and non-recurring fees. Relevant AR flags:
			//     1<<4 -  Is Rent Assessment
			//     1<<6 -  Is non-recur charge
			//     1<<7 -  PETID required
			//     1<<8 -  VID required
			//------------------------------------------------------------------
			for k := 0; k < len(vList); k++ {
				for j := 0; j < len(dbConf.VehicleFees); j++ {
					//-----------------------------------------------------------------------
					// Check for:
					//    *  a rent assessment AND we're not beginning the on the epoch,
					//    *  a one-time charge
					//-----------------------------------------------------------------------
					flags := rlib.RRdb.BizTypes[BID].AR[dbConf.VehicleFees[j].ARID].FLAGS
					needProratedRent := flags&(1<<4) != 0 && dbConf.DtStart.Day() != dbConf.Epochs.Monthly.Day()
					oneTimeCharge := flags&(1<<6) != 0
					//-----------------------------------------------------------------------
					// If it's not a prorated Rent charge or a one-time-charge, then skip it
					//-----------------------------------------------------------------------
					if !(needProratedRent || oneTimeCharge) {
						continue // skip if it's rent or it's not a one-time charge
					}
					cmt := ""
					rc := int64(rlib.RECURNONE)
					if needProratedRent {
						rc = dbConf.xbiz.RT[rtr.RTID].RentCycle
					}
					tot, np, tp := rlib.SimpleProrateAmount(dbConf.VehicleFees[j].DefaultAmount, rc, dbConf.xbiz.RT[rtr.RTID].Proration, &d1, &td2, &epoch)
					if tot < dbConf.VehicleFees[j].DefaultAmount {
						// cmt = fmt.Sprintf("prorated for %d of %d %s", np, tp, rlib.ProrationUnits(dbConf.xbiz.RT[rtr.RTID].Proration))
						cmt = rlib.ProrateComment(np, tp, dbConf.xbiz.RT[rtr.RTID].Proration)
					}
					var asm = rlib.Assessment{
						BID:            BID,
						RID:            RID,
						RAID:           ra.RAID,
						Amount:         tot,
						RentCycle:      rlib.RECURNONE,
						ProrationCycle: rlib.RECURNONE,
						AssocElemType:  rlib.ELEMVEHICLE,
						AssocElemID:    vList[k].VID,
						Start:          d1,
						Stop:           d1,
						ARID:           dbConf.VehicleFees[j].ARID,
						Comment:        cmt,
					}
					be := bizlogic.InsertAssessment(ctx, &asm, 1, &noClose)
					if be != nil {
						return bizlogic.BizErrorListToError(be)
					}
				}
			}
		}

		RID++
		if i+1 < dbConf.RACount && RID > MaxRID {
			fmt.Printf("Halting Rental Agreement creation at RAID = %d because all Rentables are rented\n", ra.RAID)
			break
		}
	}
	return nil
}

// makeDeposits is an intermediary function that makes daily deposits for receipts.
//
// INPUTS
//     ctx       - context for db transactions
//     dbConf    - module configuration
//     SecDepAmt - amount being deposited to the security deposit account
//                 depository
//     OpDepAmt  - amount being deposited to the operational account
//                 depository
//     SecDeps   - pointer to a slice of RCPTIDs that are being deposited in
//                 the in the security deposit account
//     OpDeps   - pointer to a slice of RCPTIDs that are being deposited in
//                 the operational account
//-----------------------------------------------------------------------------
func makeDeposits(ctx context.Context, dbConf *GenDBConf, SecDepAmt, OpDepAmt float64, dt *time.Time, SecDeps, OpDeps *[]int64) error {
	if SecDepAmt > float64(0) {
		var b = rlib.Deposit{
			BID:    dbConf.BIZ[0].BID,
			DEPID:  dbConf.SecDepDepository,
			DPMID:  int64(1),
			Dt:     *dt,
			Amount: SecDepAmt,
		}
		if e := bizlogic.SaveDeposit(ctx, &b, *SecDeps); len(e) > 0 {
			bizlogic.PrintBizErrorList(e)
			return bizlogic.BizErrorListToError(e)
		}
	}
	if OpDepAmt > float64(0) {
		var c = rlib.Deposit{
			BID:    dbConf.BIZ[0].BID,
			DEPID:  dbConf.OpDepository,
			DPMID:  int64(1),
			Dt:     *dt,
			Amount: OpDepAmt,
		}
		if e := bizlogic.SaveDeposit(ctx, &c, *OpDeps); len(e) > 0 {
			bizlogic.PrintBizErrorList(e)
			return bizlogic.BizErrorListToError(e)
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
	qry := fmt.Sprintf("SELECT %s FROM Receipt WHERE BID=%d AND DID=0 ORDER BY Dt ASC", rlib.RRdb.DBFields["Receipt"], bid)
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
	lastdep := rlib.TIME0
	lastrcpt := rlib.TIME0
	lastdepNotInitialized := true
	for i := 0; rows.Next(); i++ {
		var a rlib.Receipt
		if err = rlib.ReadReceipts(rows, &a); err != nil {
			return err
		}

		//------------------------------------------
		// initialize to date of first receipt
		//------------------------------------------
		if lastdepNotInitialized {
			lastdep = time.Date(a.Dt.Year(), a.Dt.Month(), a.Dt.Day(), 0, 0, 0, 0, time.UTC)
			lastdepNotInitialized = false // it is now initialized
		}

		//---------------------------------------------------------------
		// Deposits are made daily.  If the day has changed then make a
		// deposit for what we have collected already before processing
		// the new receipt...
		//---------------------------------------------------------------
		dt := time.Date(a.Dt.Year(), a.Dt.Month(), a.Dt.Day(), 0, 0, 0, 0, time.UTC) // snap dt of deposit
		if dt.After(lastdep) {                                                       // is receipt  date AFTER the last receipt receipt we processed
			err = makeDeposits(ctx, dbConf, SecDepAmt, OpDepAmt, &lastrcpt, &SecDeps, &OpDeps) // if so then deposit what we have
			if err != nil {
				return err
			}
			SecDepAmt = float64(0)
			OpDepAmt = float64(0)
			SecDeps = []int64{}
			OpDeps = []int64{}
			lastdep = dt
		}
		if a.DEPID == dbConf.OpDepository {
			SecDeps = append(SecDeps, a.RCPTID)
			SecDepAmt += a.Amount
		} else {
			OpDeps = append(OpDeps, a.RCPTID)
			OpDepAmt += a.Amount
		}
		lastrcpt = a.Dt // date of the last receipt we processed
	}
	if rows.Err() != nil {
		return rows.Err()
	}

	//-------------------------------------------------------
	// Deposit anything that has not yet been deposited...
	//-------------------------------------------------------
	makeDeposits(ctx, dbConf, SecDepAmt, OpDepAmt, &lastrcpt, &SecDeps, &OpDeps) // if so then deposit what we have
	if err != nil {
		return err
	}
	return nil
}

// CreateTaskLists creates the initial task list, associates it with the business
// as the ClosePeriod TaskList, and creates all past instances
//-----------------------------------------------------------------------------
func CreateTaskLists(ctx context.Context, dbConf *GenDBConf) error {
	// rlib.Console("Entered: CreateTaskLists\n")
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
	now := time.Now()
	return rlib.TaskListExpandPastInstances(ctx, &tl, &tld, &pivot, &now)
}
