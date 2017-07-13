-- Initialize Rentroll Database with EXAMPLE 1 data 
--   This example revolves around a fictional Business:

-- Business:
-- 	Springfield Retirement Castle
-- 	2001 Creaking Oak Drive
-- 	Springfield, MO 65803
-- 	USA

-- Users:
-- 	Homer Simpson
-- 	Edna Krabappel

-- EXAMPLE 1  -  UNIT 101
-- 	Assessment period:  Nov 1, 2015 – Nov 30, 2015
-- 	Property: #1 
-- 	Unit 101
-- 	Monthly Rent: from Jul 1, 2012 to Oct 01, 2015:  $1000
--  			  beginning Oct 1 2015:  $1200
-- 	Unit Specialty: Lake View ($50)
-- 	Unit Specialty: Fireplace ($20)

-- 	Deposit currently held:  $1000
-- 	Deposit for next User: $1500

-- 	User 1: (Edna Krabappel) vacates unit 101
-- 		She occupies the unit from Nov 1 – Nov 8 (8 days)
-- 		RentalAgreement #1 has set rent of $1000/month
-- 		$250 of her $1000 deposit was forfeited to cover damages assessed
-- 		The remaining $750 was returned to Krabappel
-- 		Krabappel also rents a carport @ $35/month
-- 		Her rental charges for the 8 days she stays in November: 
-- 			Rent:      8/30 of $1000 = $266.67 
-- 			Fireplace: 8/30 of   $20 =   $5.33
-- 			Lake View: 8/30 of   $50 =  $13.33
-- 			Carport:   8/30 of   $35 =   $9.33
-- 			----------------------------------
-- 			Rent for November:         $294.66					

-- 	Unit was vacant from Nov 9 – Nov 20 (12 days)
-- 			Unit loss to vacancy:      $428.00

-- 	User 2 (Homer Simpson) rents unit 101
-- 		RentalAgreement #8
-- 		He signs a 1 year rental agreement for unit 101:  
-- 			Rent:     $1270/month   ($1200 + $50 + $20)
-- 			Deposit:  $1500
-- 			Carports: 2   @$35/month each
-- 		He occupies unit 101 from Nov 21 – Nov 30  (10 days)
-- 		He sends in a check for 1880.00 which breaks down as follows
-- 			Rent:      10/30 of $1200 =  $400.00 
-- 			Fireplace: 10/30 of   $20 =    $6.67
-- 			Lake View: 10/30 of   $50 =   $16.67
-- 			Carport 1: 10/30 of   $35 =    11.67
-- 			Carport 2: 10/30 of   $35 =    11.67
-- 			-------------------------------------
-- 			Rent for November:           $446.68
-- 			Deposit:                    $1500.00  (since there is no proration on security deposit)
-- 			-------------------------------------
-- 			Total:                      $1946.68	

-- UPDATE:  We change the MarketRate carport fee to $40/month on Nov 15, 2015.
--			This will affect the vacancy calculation for vacant days beginning
--			Nov 15. In this test, we also introduce an inconsistency between
--          the ContractRent for CP001 and the Assessment will differ. This 
--			should cause the Checker report to list the discrepancy.
--          It is also a good check for the vacancy calculator, which must process
--			the rate change.

USE rentroll


-- ----------------------------------------------------------------------------------------
--     PAYMENT TYPES
-- ----------------------------------------------------------------------------------------
INSERT INTO PaymentTypes (BID, Name,Description) VALUES
	(1,"Check","Personal check from Payor"),
	(1,"VISA","Credit card charge"),
	(1,"AMEX", "American Express credit card"),
	(1,"Cash","Cash");


-- ----------------------------------------------------------------------------------------
--     AVAILABILITY TYPES
-- ----------------------------------------------------------------------------------------
INSERT INTO AvailabilityTypes (Name) VALUES
	("Occupied"),
	("Offline"),
	("Administrative"),
	("Vacant"),
	("Not Ready"),
	("Vacant - Made Ready"),
	("Vacant - Inspected");



-- =======================================================================
--  BUSINESS
-- =======================================================================
INSERT INTO Business (BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC) VALUES
	("REX", "Springfield Retirement Castle",6,4,4);

-- =======================================================================
--  RENTABLE TYPES
-- =======================================================================
INSERT INTO RentableTypes (BID,Style, Name,RentCycle,Proration,GSRPC,ManageToBudget) VALUES
	(1,"GM","Geezer Miser", 6,4,4,1),				-- 1  From TUC "Efficiency"  385 sqft
	(1,"FS","Flat Studio",  6,4,4,1),				-- 2  From TUC "1/1 Studio"  726 sqft
	(1,"SBL","SB Loft",     6,4,4,1),				-- 3  From TUC "2/2 Flat"    770 sqft
	(1,"KDS","KD Suite",    6,4,4,1),				-- 4  From TUC "2/2 Studio" 1123 sqft
	(1,"CAR","Vehicle",     3,0,4,1), 				-- 5  Car
	(1,"CPT","Carport",     6,4,4,1);		 		-- 6  Carport

INSERT INTO RentableMarketRate (RTID,MarketRate,DtStart,DtStop) VALUES
	(1, 1000.00, "1970-01-01 00:00:00", "2015-10-01 00:00:00"),   	-- 1: GM, Geezer Miser 
	(2, 1500.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 2: FS, Flat Studio
	(3, 1750.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 3: SBL, SB Loft
	(4, 2000.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 4: KDS, Krusty Deluxe Suite
	(5,   10.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 5  Car
	(6,   35.00, "1970-01-01 00:00:00", "2015-11-15 00:00:00"),		-- 6  Carport
	(6,   40.00, "2015-11-15 00:00:00", "9999-12-31 00:00:00"),		-- 6  Carport  ** RAISED THE RENT **
	(1, 1200.00, "2015-10-01 00:00:00", "9999-12-31 00:00:00");   	-- 1: GM, Geezer Miser  ** RAISED THE RENT **


-- define unit specialties

-- rentablespecialtytype
INSERT INTO RentableSpecialty (BID,Name,Fee,Description) VALUES
	(1,"Lake View",50.0,"Overlooks the lake"),						-- assmt 59
	(1,"Courtyard View",50.0,"Rear windows view the courtyard"),	-- assmt 60
	(1,"Top Floor",100.0,"Penthouse"),								-- assmt 61
	(1,"Fireplace",20.0,"Wood burning, gas fireplace");				-- assmt 62

-- define the Assessments
INSERT INTO BusinessAssessments (BID,ATypeLID) VALUES
	(1, 1),		-- Rent
	(1, 2),		-- Security Deposit
	(1, 3),		-- Security Deposit Forfeiture
	(1, 4);		-- Application Fees

-- define the Building
INSERT INTO Building (BID,Address,Address2,City,State,PostalCode,Country) VALUES
	(1,"2001 Creaking Oak Drive","","Springfield","MO","65803","USA");


-- Rental agreement templates
INSERT INTO RentalAgreementTemplate (RATemplateName, BID) VALUES
	("RAT001", 1),
	("RAT002", 1),	-- port
	("RAT003", 1),	-- rental unit
	("RAT004", 1);

-- =======================================================================
--  RENTABLES
-- =======================================================================
INSERT INTO Rentable (BID,Name,AssignmentTime) VALUES
	(1,"101",1),  -- RID 1
  	(1,"102",1),	-- RID 2
  	(1,"103",1),	-- RID 3
  	(1,"104",1),	-- RID 4
  	(1,"105",1),	-- RID 5
  	(1,"106",1),	-- RID 6
  	(1,"107",1);	-- RID 7


-- =======================================================================
--  carports
-- =======================================================================
INSERT INTO Rentable (BID,Name,AssignmentTime) VALUES
	(1,"CP001",1),		-- RID 8  Krabappel, then Simpson
	(1,"CP002",1);		-- RID 9  Simpson

-- =======================================================================
--  RentableState - All Rentables
-- =======================================================================
INSERT INTO RentableStatus (RID,DtStart,DtStop,Status) VALUES
	(1,"2014-01-01","9999-01-01",1),  -- RID 1
	(2,"2014-01-01","9999-01-01",1),  -- RID 2
	(3,"2014-01-01","9999-01-01",1),  -- RID 3
	(4,"2014-01-01","9999-01-01",1),  -- RID 4
	(5,"2014-01-01","9999-01-01",1),  -- RID 5
	(6,"2014-01-01","9999-01-01",1),  -- RID 6
	(7,"2014-01-01","9999-01-01",1),  -- RID 7
	(8,"2014-01-01","9999-01-01",1),  -- RID 8
	(9,"2014-01-01","9999-01-01",1);  -- RID 9

-- =======================================================================
--  RentableTypeRef - All Rentables
-- =======================================================================
INSERT INTO RentableTypeRef (RID,RTID,DtStart,DtStop) VALUES
	(1,1,"2014-01-01","9999-01-01"),  -- RID 1
	(2,2,"2014-01-01","9999-01-01"),  -- RID 2
	(3,3,"2014-01-01","9999-01-01"),  -- RID 3
	(4,4,"2014-01-01","9999-01-01"),  -- RID 4
	(5,1,"2014-01-01","9999-01-01"),  -- RID 5
	(6,2,"2014-01-01","9999-01-01"),  -- RID 6
	(7,3,"2014-01-01","9999-01-01"),  -- RID 7
	(8,6,"2014-01-01","9999-01-01"),  -- RID 8
	(9,6,"2014-01-01","9999-01-01");  -- RID 9


-- =======================================================================
--  UNIT SPECIALTIES
-- =======================================================================
INSERT INTO RentableSpecialtyRef (BID,RID,RSPID) VALUES
	(1,1,1),
	(1,1,4),
	(1,2,2),
	(1,2,3),
	(1,2,4),
	(1,3,1),
	(1,4,2),
	(1,4,4),
	(1,6,3),
	(1,7,1),
	(1,7,4),
	(1,7,3);

-- =======================================================================
--  TRANSACTANTS
-- =======================================================================
-- define the renters.  First as transactants, second as renters, 3rd as payors
INSERT INTO Transactant (BID,FirstName,LastName,PrimaryEmail,Address,City,State,PostalCode,Country) VALUES
	(1,"Edna", "Krabappel","edna@springfield.com","","Springfield","MO","64055","USA"),			-- 1
	(1,"Ned", "Flanders","flanman@springfield.com","","Springfield","MO","64055","USA"),			-- 2
	(1,"Moe", "Szyslak","moe@springfield.com","","Springfield","MO","64055","USA"),				-- 3
	(1,"Montgomery", "Burns","burnsie@springfield.com","","Springfield","MO","64055","USA"),		-- 4
	(1,"Nelson", "Muntz","nelson@springfield.com","","Springfield","MO","64055","USA"),			-- 5
	(1,"Milhouse", "Van Houten","milhouse@springfield.com","","Springfield","MO","64055","USA"),	-- 6
	(1,"Clancey", "Wiggum","wiggum@springfield.com","","Springfield","MO","64055","USA"),			-- 7
	(1,"Homer", "Simpson","homer@springfield.com","744 Evergreen Terrace","Springfield","MO","64055","USA"),			-- 8
	(1,"Marge", "Simpson","marge@springfield.com","744 Evergreen Terrace","Springfield","MO","64055","USA");			-- 9


-- define the renters.
INSERT INTO User (TCID) VALUES
	  (1),  (2),  (3),  (4),  (5),  (6),  (7),  (8), (9);

-- define the payors.
INSERT INTO Payor (TCID) VALUES
	  (1),  (2),  (3),  (4),  (5),  (6),  (7),  (8), (9);

-- =======================================================================
--  RENTAL AGREEMENTS
--    These are initially generated when the rentor changes from
--    an applicant to a User (or Payor as the case may be)
--    RATID - rental agreement template
-- =======================================================================
INSERT INTO RentalAgreement (RATID,BID,AgreementStart,AgreementStop,PossessionStart,PossessionStop,RentStart,RentStop,Renewal) VALUES
	(6,1, "2004-01-01","2015-11-09","2010-01-01","2015-11-09","2010-01-01","2015-11-09",1),	--  1 Krabappel
	(6,1, "2004-01-01","2017-07-04","2004-01-01","2017-07-04","2004-01-01","2017-07-04",1),	--  2 Flanders
	(6,1, "2004-01-01","2017-07-04","2004-01-01","2017-07-04","2004-01-01","2017-07-04",1),	--  3 Szyslak
	(6,1, "2004-01-01","2017-07-04","2004-01-01","2017-07-04","2004-01-01","2017-07-04",1),	--  4 Burns
	(6,1, "2004-01-01","2017-07-04","2004-01-01","2017-07-04","2004-01-01","2017-07-04",1),	--  5 Muntz
	(6,1, "2004-01-01","2017-07-04","2004-01-01","2017-07-04","2004-01-01","2017-07-04",1),	--  6 Van Houten
	(6,1, "2004-01-01","2017-07-04","2004-01-01","2017-07-04","2004-01-01","2017-07-04",1),	--  7 Wiggum
	(6,1, "2015-11-21","2016-11-21","2015-11-21","2017-11-21","2015-11-21","2017-11-21",1);	--  8 Simpson

INSERT INTO RentalAgreementRentables (RAID,RID,ContractRent,DtStart,DtStop) VALUES
	(1,1,1000,"2004-01-01","2015-11-09"),		-- Krabappel - apartment
	(1,8,  35,"2004-01-01","2015-11-09"),		-- Krabappel - carport
	(8,1,1200,"2015-11-21","2016-11-21"),		-- Simpson - apartment
	(8,8,  40,"2015-11-21","2016-11-21"),		-- Simpson - carport 1
	(8,9,  40,"2015-11-21","2016-11-21");		-- Simpson - carport 2

INSERT INTO RentalAgreementPayors (RAID,TCID,DtStart,DtStop) VALUES
	(1,1,"2004-01-01","2015-11-09"),		-- Krabappel is Payor for rental agreement 1
	(8,8,"2015-11-21","2016-11-21");		-- Homer is Payor for rental agreements 8

INSERT INTO RentableUsers (RID,TCID,DtStart,DtStop) VALUES
	(1,1,"2004-01-01","2015-11-09"),		-- Krabappel is User for rentable 1
	(8,8,"2015-11-21","2016-11-21"),		-- Homer is a User for Rentable 1
	(8,9,"2015-11-21","2016-11-21");		-- Marge is a User for Rentable 1

-- =======================================================================
--  CONTRACT RENT ASSESSMENTS
--    These are initially generated when the rentor changes from
--    an applicant to a User (or Payor as the case may be)
-- =======================================================================
INSERT INTO Assessments (BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle, AcctRule) VALUES
	(1, 1, 3, 1,1000.00,"2014-07-01","2015-11-09", 6, 4, "d ${GLGENRCV} _, c ${GLGSRENT} ${UMR}, d ${GLLTL} ${UMR} _ -"),		-- #1  Krabappel - Rent
	(1, 1, 3, 8,1200.00,"2015-11-21","2016-11-21", 6, 4, "d ${GLGENRCV} _, c ${GLGSRENT} ${UMR}, d ${GLLTL} ${UMR} ${aval(${GLGENRCV})} -");		-- #2  Simpson rent

-- =======================================================================
--  UNIT SPECIALTY ASSESSMENTS
-- =======================================================================
INSERT INTO Assessments (BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle, AcctRule) VALUES
	(1, 1, 28, 1,50.00,"2014-07-01","2015-11-09", 6, 4, "d ${GLGENRCV} _, c ${GLGSRENT} _"),		-- #3 Lake view  Krabappel
	(1, 1, 29, 1,20.00,"2014-07-01","2015-11-09", 6, 4, "d ${GLGENRCV} _, c ${GLGSRENT} _"),		-- #4 Fireplace  Krabappel
	(1, 1, 28, 8,50.00,"2015-11-21","2016-11-21", 6, 4, "d ${GLGENRCV} _, c ${GLGSRENT} _"),		-- #5 Lake view  Simpson
	(1, 1, 29, 8,20.00,"2015-11-21","2016-11-21", 6, 4, "d ${GLGENRCV} _, c ${GLGSRENT} _");		-- #6 Fireplace  Simpson

-- =======================================================================
--  CONTRACT SECURITY DEPOSIT
--    These are initially generated when the rentor changes from
--    an applicant to a User (or Payor as the case may be)
-- =======================================================================
INSERT INTO Assessments (BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle, AcctRule) VALUES
	(1, 1, 2, 1,1000.00,"2014-07-01", "2014-07-01", 0, 0, "d ${GLGENRCV} _, c ${GLSECDEP} _"),		-- #7 Krabappel deposit
	(1, 1, 31, 8,1500.00,"2015-11-21", "2015-12-02", 6, 0, "d ${GLGENRCV} _, c ${GLSECDEP} _");		-- #8 Simpson deposit: (recurring is a hack to make it process)

-- =======================================================================
--  CARPORT ASSESSMENTS
--    These can be generated at any time. Typically they will be
--    created along with the rental agreement
-- =======================================================================
INSERT INTO Assessments (BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle, AcctRule) VALUES
	(1, 8, 30, 1,35.00,"2014-07-01","2015-11-09", 6, 4, "d ${GLGENRCV} _, c 42007 _"),		-- #9  Krabappel, ends Nov 10
	(1, 8, 30, 8,35.00,"2015-11-21","2016-11-10", 6, 4, "d ${GLGENRCV} _, c 42007 _"),		-- #10 Simpson, starts Nov 21
	(1, 9, 30, 8,35.00,"2015-11-21","2016-11-10", 6, 4, "d ${GLGENRCV} _, c 42007 _");		-- #11 Simpson, starts Nov 21

-- =======================================================================
--  DAMAGE ASSESSMENTS
-- =======================================================================
INSERT INTO Assessments (BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle, AcctRule) VALUES
	(1, 1, 24, 1,250.00,"2015-11-08","2015-12-05", 6, 0, "d ${GLSECDEP} _, c 42006 _"),	-- #12  Krabappel, $250 damages
	(1, 1, 31, 1,750.00,"2015-11-08","2015-12-05", 6, 0, "d ${GLSECDEP} _, c 10001 _");

-- =======================================================================
--  RECEIPTS
-- =======================================================================

-- TODO:  ADD ACCTRULE TO RECEIPTS...

INSERT INTO Receipt (BID,RAID,PMTID,Dt,DocNo,Amount,AcctRule) VALUES
	(1,1,2,"2004-01-01", "2301", 1000.00, "d ${GLSECDEP} _, c ${GLGENRCV} _");			-- 1  Krabappel's initial security deposit
INSERT INTO ReceiptAllocation (RCPTID,Amount,ASMID,AcctRule) VALUES
	(1,1000.00,7, "d ${GLSECDEP} 1000.00, c ${GLGENRCV} 1000.00");		

-- INSERT INTO Receipt (BID,RAID,PMTID,Dt,DocNo,Amount,AcctRule) VALUES
-- 	(1,1,1,"2015-11-21","5794",  294.66, "ASM(1) c ${GLGENRCV} 266.67, ASM(1) d ${GLCASH} 266.67, ASM(3) c ${GLGENRCV} 13.33, ASM(3) d ${GLCASH} 13.33, ASM(4) c ${GLGENRCV} 5.33, ASM(4) d ${GLCASH} 5.33, ASM(9) c ${GLGENRCV} 9.33,ASM(9) d ${GLCASH} 9.33"); 			-- 2   Krabappel pays her fees in full
-- INSERT INTO ReceiptAllocation (RCPTID,Amount,ASMID,AcctRule) VALUES
-- 	(2,266.67,1,"c ${GLGENRCV} _, d ${GLCASH} _"),	-- rent
-- 	(2, 13.33,3,"c ${GLGENRCV} _, d ${GLCASH} _"),	-- Lake View
-- 	(2,  5.33,4,"c ${GLGENRCV} _, d ${GLCASH} _"),	-- Fireplace
-- 	(2,  9.33,9,"c ${GLGENRCV} _, d ${GLCASH} _");	-- CP001

-- INSERT INTO Receipt (BID,RAID,PMTID,Dt,DocNo,Amount,AcctRule) VALUES
-- 	(1,8,1,"2015-11-30","8330",1946.68, "d ${GLSECDEP} 1500.00, c 11002 1500.00, c ${GLGENRCV} 400.00, d ${GLCASH} 400.00, c ${GLGENRCV} 16.67, d ${GLCASH} 16.67, c ${GLGENRCV} 6.67, d ${GLCASH} 6.67, c ${GLGENRCV} 11.67,d ${GLCASH} 11.67, c ${GLGENRCV} 11.67,d ${GLCASH} 11.67");  			-- 3   Simpson pays his fees in full
-- INSERT INTO ReceiptAllocation (RCPTID,Amount,ASMID,AcctRule) VALUES
-- 	(3,1500.00, 8,"d ${GLSECDEP}   _,c ${GLGENRCV} _"),	--  security deposit
-- 	(3, 400.00, 2,"c ${GLGENRCV} _,d ${GLCASH}      _"),	--  rent
-- 	(3,  16.67, 5,"c ${GLGENRCV} _,d ${GLCASH}      _"),	--  Lake View
-- 	(3,   6.67, 6,"c ${GLGENRCV} _,d ${GLCASH}      _"),	--  Fireplace
-- 	(3,  11.67,10,"c ${GLGENRCV} _,d ${GLCASH}      _"),	--  CP001
-- 	(3,  11.67,11,"c ${GLGENRCV} _,d ${GLCASH}      _");	--  CP002


-- =======================================================================
--  JOURNAL MARKERS
-- =======================================================================
INSERT INTO JournalMarker (BID,State,DtStart,DtStop) VALUES
	(1, 3, "2015-10-31", "2015-10-31");


-- =======================================================================
--  LEDGERS MARKERS
-- =======================================================================
INSERT INTO GLAccount (BID,RAID,GLNumber,Status,Type,Name) VALUES
	(1,0,"11003",2,0,"American Express 93892335 In Process"),
	(1,0,"10004",2,0,"MasterCard/VISA 38992355 in Process"),
	(1,0,"10005",2,0,"Discover 883523553 In Process"),
	(1,0,"40002",2,0,"Carport Rental"),
	(1,0,"41002",2,0,"Administrative Concession"),
	(1,0,"41003",2,0,"Off Line Unit"),
	(1,0,"41005",2,0,"Employee Concession"),
	(1,0,"41006",2,0,"Payor Concession"),
	(1,0,"41007",2,0,"Bad Debt Write-Off"),
	(1,0,"41999",2,0,"Other Offsets"),
	(1,0,"42001",2,0,"Pet Fees"),
	(1,0,"42002",2,0,"Utility Reimbursement"),
	(1,0,"42003",2,0,"Late Fees"),
	(1,0,"42004",2,0,"Damages"),
	(1,0,"42005",2,0,"Pest Control Reimbursements"),
	(1,0,"42006",2,0,"Security Deposit Forfeiture"),
	(1,0,"42007",2,0,"Carports"),
	(1,0,"42008",2,0,"Utility overage"),
	(1,0,"42999",2,0,"Other Income"),
	(1,0,"43001",2,0,"Lake View"),
	(1,0,"43002",2,0,"Fireplace"),
	(1,0,"43003",2,0,"Carport"),
	(1,0,"43004",2,0,"Security Deposit");
         
-- =======================================================================
--  LEDGERS MARKERS
-- =======================================================================
INSERT INTO LedgerMarker (BID,LID,State,Dt,Balance) VALUES
	( 1,9,3,"2015-10-31",0.0),
	(1,10,3,"2015-10-31",0.0),
	(1,11,3,"2015-10-31",0.0),
	(1,12,3,"2015-10-31",0.0),
	(1,13,3,"2015-10-31",0.0),
	(1,14,3,"2015-10-31",0.0),
	(1,15,3,"2015-10-31",0.0),
	(1,16,3,"2015-10-31",0.0),
	(1,17,3,"2015-10-31",0.0),
	(1,18,3,"2015-10-31",0.0),
	(1,19,3,"2015-10-31",0.0),
	(1,20,3,"2015-10-31",0.0),
	(1,21,3,"2015-10-31",0.0),
	(1,22,3,"2015-10-31",0.0),
	(1,23,3,"2015-10-31",0.0),
	(1,24,3,"2015-10-31",0.0),
	(1,25,3,"2015-10-31",0.0),
	(1,26,3,"2015-10-31",0.0),
	(1,27,3,"2015-10-31",0.0),
	(1,28,3,"2015-10-31",0.0),
	(1,29,3,"2015-10-31",0.0),
	(1,30,3,"2015-10-31",0.0),
	(1,31,3,"2015-10-31",0.0),
	(1,32,3,"2015-10-31",0.0),
	(1,33,3,"2015-10-31",0.0),
	(1,34,3,"2015-10-31",0.0),
	(1,35,3,"2015-10-31",0.0),
	(1,36,3,"2015-10-31",0.0),
	(1,37,3,"2015-10-31",0.0),
	(1,38,3,"2015-10-31",0.0),
	(1,39,3,"2015-10-31",0.0);
         
UPDATE GLAccount SET GLNumber="10001" WHERE Name = "Bank Account";
UPDATE GLAccount SET GLNumber="11001" WHERE Name = "General Accounts Receivable";
UPDATE GLAccount SET GLNumber="40001" WHERE Name = "Gross Scheduled Rent";
UPDATE GLAccount SET GLNumber="41004" WHERE Name = "Loss to Lease";
UPDATE GLAccount SET GLNumber="41001" WHERE Name = "Vacancy";
UPDATE GLAccount SET GLNumber="11002" WHERE Name = "Purge Me Acct";
UPDATE GLAccount SET GLNumber="23000" WHERE Name = "Security Deposit Assessment";
UPDATE GLAccount SET GLNumber="20001" WHERE Name = "Owner Equity";

-- UPDATE LedgerMarker SET Balance=-1000.00 WHERE LID=7;
UPDATE GLAccount SET Name="Bank Account FRB 2332352" WHERE GLNumber = "10001";
