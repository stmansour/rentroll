-- Initialize Rentroll Database with EXAMPLE 1 data 
--   This example revolves around a fictional business:

-- Business:
-- 	Springfield Retirement Castle
-- 	2001 Creaking Oak Drive
-- 	Springfield, MO 65803
-- 	USA

-- Renters:
-- 	Homer Simpson
-- 	Edna Krabappel

-- EXAMPLE 1  -  UNIT 101
-- 	Assessment period:  Nov 1, 2015 – Nov 30, 2015
-- 	Property: #1 
-- 	Unit 101
-- 	Monthly Rent: from Jul 1, 2012 to Oct 31, 2015:  $1000
--  			  beginning Nov 1 2015:  $1200
-- 	Unit Specialty: Lake View ($50)
-- 	Unit Specialty: Fireplace ($20)

-- 	Deposit currently held:  $1000
-- 	Deposit for next renter: $1500

-- 	Renter 1: (Edna Krabappel) vacates unit 101
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

-- 	Renter 2 (Homer Simpson) rents unit 101
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


USE rentroll

-- define the business
INSERT INTO business (Name,Address,Address2,City,State,PostalCode,Country,Phone,DefaultOccupancyType,ParkingPermitInUse) VALUES
	("Springfield Retirement Castle","2001 Creaking Oak Drive","","Springfield","MO","65803","USA","939-555-1000",3,0);

-- =======================================================================
--  RENTABLE TYPES
-- =======================================================================
INSERT INTO rentabletypes (BID,Name,Frequency,Proration,Report,ManageToBudget) VALUES
	(1,"Residential", 6,4,1,1),				-- 1  
	(1,"Office",      6,4,1,1),				-- 2  
	(1,"Industrial",  6,4,1,1),				-- 3  
	(1,"Unimproved",  6,4,1,1),				-- 4  
	(1,"Vehicle",     3,0,1,1), 			-- 5  Car
	(1,"Carport",     6,4,1,0);		 		-- 6  Carport

INSERT INTO rentablemarketrate (RTID,MarketRate,DtStart,DtStop) VALUES
	(1,   0.0, "1970-01-01 00:00:00", "1970-01-01 00:00:00"),   -- 1 ignore
	(2,   0.0, "1970-01-01 00:00:00", "1970-01-01 00:00:00"),	-- 2  ""
	(3,   0.0, "1970-01-01 00:00:00", "1970-01-01 00:00:00"),	-- 3  ""
	(4,   0.0, "1970-01-01 00:00:00", "1970-01-01 00:00:00"),	-- 4  ""
	(5,  10.0, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),	-- 5  Car
	(6,  35.0, "1970-01-01 00:00:00", "9999-12-31 00:00:00");	-- 6  Carport

-- =======================================================================
--  UNIT TYPES
-- =======================================================================
INSERT INTO unittypes (BID,Style,Name,SqFt,MarketRate,Frequency,Proration) VALUES
	(1,"GM","Geezer Miser",385,1100.0,6,4),				-- 1  rented monthly, prorate daily
	(1,"FS","Flat Studio",726,1500.0,6,4),				-- 2     "
	(1,"SBL","SB Loft",770,1750.0,6,4),					-- 3     "
	(1,"KDS","KD Suite",1123,2000.0,6,4);				-- 4     "

INSERT INTO unitmarketrate (UTID,MarketRate,DtStart,DtStop) VALUES
	(1, 1100.00, "1970-01-01 00:00:00", "2015-10-31 00:00:00"),   	-- 1:  GM, Geezer Miser 
	(2, 1500.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 2:  FS, Flat Studio
	(3, 1750.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 3: SBL, SB Loft
	(4, 2000.00, "1970-01-01 00:00:00", "9999-12-31 00:00:00"),		-- 4: KDS, Krusty Deluxe Suite
	(1, 1200.00, "2015-10-01 00:00:00", "9999-12-31 00:00:00");   	-- 1:  GM, Geezer Miser  ** RAISED THE RENT **

-- define unit specialties

-- unitspecialtytype -> rentalspecialtytype
INSERT INTO unitspecialtytypes (BID,Name,Fee,Description) VALUES
	(1,"Lake View",50.0,"Overlooks the lake"),						-- assmt 59
	(1,"Courtyard View",50.0,"Rear windows view the courtyard"),	-- assmt 60
	(1,"Top Floor",100.0,"Penthouse"),								-- assmt 61
	(1,"Fireplace",20.0,"Wood burning, gas fireplace");				-- assmt 62

-- define the assessments
INSERT INTO businessassessments (BID,ASMTID) VALUES
	(1, 1),		-- Rent
	(1, 2),		-- Security Deposit
	(1, 3),		-- Security Deposit Forfeiture
	(1, 4);		-- Application Fees

-- define the building
INSERT INTO building (BID,Address,Address2,City,State,PostalCode,Country) VALUES
	(1,"2001 Creaking Oak Drive","","Springfield","MO","65803","USA");


-- Rental agreement templates
INSERT INTO rentalagreementtemplate (ReferenceNumber, RentalAgreementType) VALUES
	("RAT001", 2),
	("RAT002", 2),	-- port
	("RAT003", 2),	-- rental unit
	("RAT004", 2);

-- =======================================================================
--  RENTAL UNITS
-- =======================================================================
INSERT INTO rentable (LID,RTID,BID,UNITID,Name,Assignment) VALUES
	(1,1,1,  1,"101",1),  -- monthly rent for unit 1, recurs on the first of the month
  	(2,1,1,  2,"102",1),
  	(3,1,1,  3,"103",1),
  	(4,1,1,  4,"104",1),
  	(5,1,1,  5,"105",1),
  	(6,1,1,  6,"106",1),
  	(7,1,1,  7,"107",1);

-- =======================================================================
--  carports
-- =======================================================================
INSERT INTO rentable (LID,RTID,BID,UNITID,Name,Assignment,DefaultOccType,OccType) VALUES
	( 8,2,1,  1,"CP001",1,2,2),		-- carport  Krabappel, then Simpson
	( 9,2,1,  1,"CP002",1,2,2),		-- carport  Simpson
	(10,2,1,  1,"CP003",1,2,2),		-- carport
	(11,2,1,  1,"CP004",1,2,2),		-- carport
	(12,2,1,  1,"CP005",1,2,2),		-- carport
	(13,2,1,  1,"CP006",1,2,2),		-- carport
	(14,2,1,  1,"CP007",1,2,2),		-- carport
	(15,2,1,  1,"CP008",1,2,2),		-- carport
	(16,2,1,  1,"CP009",1,2,2),		-- carport
	(17,2,1,  1,"CP010",1,2,2),		-- carport
	(18,2,1,  1,"CP011",1,2,2);		-- carport


-- define the units
--     OccType == occupancy type -- 0 = unset, 1 = leasehold, 2 = month-to-month, 3 = hotel, 4 = hourly rental
INSERT INTO unit (RID,BLDGID,UTID) VALUES
	 (1,1,1),
	 (2,1,2),
	 (3,1,3),
	 (4,1,4),
	 (5,1,1),
	 (6,1,2),
	 (7,1,3);

-- Define unit specialties...
INSERT INTO unitspecialties (BID,UNITID,USPID) VALUES
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

-- define the tenants.  First as transactants, second as tenants, 3rd as payors
INSERT INTO transactant (FirstName,LastName) VALUES
	("Edna", "Krabappel"),			-- 1
	("Ned", "Flanders"),			-- 2
	("Moe", "Szyslak"),				-- 3
	("Montgomery", "Burns"),		-- 4
	("Nelson", "Muntz"),			-- 5
	("Milhouse", "Van Houten"),		-- 6
	("Clancey", "Wiggum"),			-- 7
	("Homer", "Simpson");			-- 8

-- define the tenants.
INSERT INTO tenant (TCID) VALUES
	  (1),  (2),  (3),  (4),  (5),  (6),  (7),  (8);

-- define the payors.
INSERT INTO payor (TCID) VALUES
	  (1),  (2),  (3),  (4),  (5),  (6),  (7),  (8);

-- =======================================================================
--  RENTAL AGREEMENTS
--    These are initially generated when the rentor changes from
--    an applicant to a tenant (or payor as the case may be)
--    RATID - rental agreement template
-- =======================================================================
INSERT INTO rentalagreement (RATID,BID,RID,UNITID,PID,LID,PrimaryTenant,RentalStart,RentalStop,Renewal) VALUES
	(6,1, 1, 1, 1, 1, 1,"2004-01-01","2015-11-08",1),	--  1 Krabappel
	(6,1, 2, 2, 2, 2, 2,"2004-01-01","2017-07-04",1),	--  2 Flanders
	(6,1, 3, 3, 3, 3, 3,"2004-01-01","2017-07-04",1),	--  3 Szyslak
	(6,1, 4, 4, 4, 4, 4,"2004-01-01","2017-07-04",1),	--  4 Burns
	(6,1, 5, 5, 5, 5, 5,"2004-01-01","2017-07-04",1),	--  5 Muntz
	(6,1, 6, 6, 6, 6, 6,"2004-01-01","2017-07-04",1),	--  6 Van Houten
	(6,1, 7, 7, 7, 7, 7,"2004-01-01","2017-07-04",1),	--  7 Wiggum
	(6,1, 1, 1, 8, 8, 8,"2015-11-21","2016-11-21",1);	--  8 Simpson

-- =======================================================================
--  CONTRACT RENT ASSESSMENTS
--    These are initially generated when the rentor changes from
--    an applicant to a tenant (or payor as the case may be)
-- =======================================================================
INSERT INTO assessments (UNITID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod, AcctRule) VALUES
	(1, 1, 1, 1, 1,1000.00,"2014-07-01","2015-11-08", 6, 4, "d ${DFLTGENRCV} 1000.0, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} ${aval(${DFLTGENRCV})} -"),		-- #1  Krabappel - Rent
	(1, 1, 1, 1, 8,1200.00,"2015-11-21","2016-11-21", 6, 4, "d ${DFLTGENRCV} 1200.0, c ${DFLTGSRENT} ${UMR}, d ${DFLTLTL} ${UMR} ${aval(${DFLTGENRCV})} -");		-- #2  Simpson rent
	-- (1, 1, 1, 1, 1,1000.00,"2014-07-01","2015-11-08", 6, 4, "d ${DFLTGENRCV} 1000.0, c ${DFLTGSRENT} 1000.0"),		-- #1  Krabappel - Rent
	-- (1, 1, 1, 1, 8,1000.00,"2015-11-21","2016-11-21", 6, 4, "d ${DFLTGENRCV} 1000.0, c ${DFLTGSRENT} 1000.0");		-- #2  Simpson rent
	-- (2, 1, 2, 1, 2,1050.00,"2011-04-01","2016-04-30", 6, 4, "d ${DFLTGENRCV} 1050.00, c ${DFLTGSRENT} 1050.00"),		
	-- (3, 1, 3, 1, 3,1095.00,"2015-04-01","2016-03-31", 6, 4, "d ${DFLTGENRCV} 1095.00, c ${DFLTGSRENT} 1095.00"),
	-- (4, 1, 4, 1, 4,1075.00,"2013-10-01","2016-03-31", 6, 4, "d ${DFLTGENRCV} 1075.00, c ${DFLTGSRENT} 1075.00"),
	-- (5, 1, 5, 1, 5, 950.00,"2015-04-01","2016-03-31", 6, 4, "d ${DFLTGENRCV}  950.00, c ${DFLTGSRENT}  950.00"),
	-- (6, 1, 6, 1, 6,1095.00,"2015-10-01","2015-10-31", 6, 4, "d ${DFLTGENRCV} 1095.00, c ${DFLTGSRENT} 1095.00"),
	-- (7, 1, 7, 1, 7,1045.00,"2001-11-01","2016-05-31", 6, 4, "d ${DFLTGENRCV} 1045.00, c ${DFLTGSRENT} 1045.00");

-- =======================================================================
--  UNIT SPECIALTY ASSESSMENTS
-- =======================================================================
INSERT INTO assessments (UNITID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod, AcctRule) VALUES
	(1, 1, 1, 59, 1,50.00,"2014-07-01","2015-11-08", 6, 4, "d ${DFLTGENRCV} 50.00, c ${DFLTGSRENT} ${aval(${DFLTGENRCV})}"),		-- #3 Lake view  Krabappel
	(1, 1, 1, 62, 1,20.00,"2014-07-01","2015-11-08", 6, 4, "d ${DFLTGENRCV} 20.00, c ${DFLTGSRENT} ${aval(${DFLTGENRCV})}"),		-- #4 Fireplace  Krabappel
	(1, 1, 1, 59, 8,50.00,"2015-11-21","2016-11-21", 6, 4, "d ${DFLTGENRCV} 50.00, c ${DFLTGSRENT} ${aval(${DFLTGENRCV})}"),		-- #5 Lake view  Simpson
	(1, 1, 1, 62, 8,20.00,"2015-11-21","2016-11-21", 6, 4, "d ${DFLTGENRCV} 20.00, c ${DFLTGSRENT} ${aval(${DFLTGENRCV})}");		-- #6 Fireplace  Simpson

-- =======================================================================
--  CONTRACT SECURITY DEPOSIT
--    These are initially generated when the rentor changes from
--    an applicant to a tenant (or payor as the case may be)
-- =======================================================================
INSERT INTO assessments (UNITID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod, AcctRule) VALUES
	(  1, 1, 1, 2, 1,1000.00,"2014-07-01", "2014-07-01", 0, 0, "d ${DFLTSECDEPRCV} 1000.00, c ${DFLTSECDEPASMT} ${aval(${DFLTSECDEPRCV})}"),		-- #7 Krabappel deposit
	(  1, 1, 1, 2, 8,1500.00,"2015-11-21", "2015-11-21", 0, 0, "d ${DFLTSECDEPRCV} 1500.00, c ${DFLTSECDEPASMT} ${aval(${DFLTSECDEPRCV})}");		-- #8 Simpson deposit

-- =======================================================================
--  CARPORT ASSESSMENTS
--    These can be generated at any time. Typically they will be
--    created along with the rental agreement
-- =======================================================================
INSERT INTO assessments (UNITID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod, AcctRule) VALUES
	(1, 1, 8, 28, 1,35.00,"2014-07-01","2015-11-08", 6, 4, "d ${DFLTGENRCV} 35.00, c 42007 35.00"),		-- #9  Krabappel, ends Nov 10
	(1, 1, 8, 28, 8,35.00,"2015-11-21","2016-11-10", 6, 4, "d ${DFLTGENRCV} 35.00, c 42007 35.00"),		-- #10 Simpson, starts Nov 21
	(1, 1, 9, 28, 8,35.00,"2015-11-21","2016-11-10", 6, 4, "d ${DFLTGENRCV} 35.00, c 42007 35.00");		-- #11 Simpson, starts Nov 21

-- =======================================================================
--  DAMAGE ASSESSMENTS
-- =======================================================================
INSERT INTO assessments (UNITID,BID,RID,ASMTID,RAID,Amount,Start,Stop,Frequency,ProrationMethod, AcctRule) VALUES
	(1, 1, 1, 53, 1,250.00,"2015-11-08","2015-11-08", 0, 0, "d ${DFLTSECDEPASMT} 250.00, c 42006 250.00"),	-- #12  Krabappel, $250 damages
	(1, 1, 1, 55, 1,750.00,"2015-11-08","2015-11-08", 0, 0, "d ${DFLTSECDEPASMT} 750.00, c 10001 750.00");

-- =======================================================================
--  OTHER ASSESSMENTS
-- =======================================================================
-- INSERT INTO assessments (UNITID,BID,RID,ASMTID,Amount,Start,Stop,Frequency, AcctRule) VALUES
-- 	(  1,1,1,10, 50.00,"2015-10-01", "2016-12-31", 6, "d ${DFLTGENRCV}, c 42002"),	-- Water (utility) reimbursement
-- 	(  1,1,1,21,150.00,"2015-12-01", "2015-12-01", 0, "d ${DFLTGENRCV}, c 42002"),	-- Furniture rental
-- 	(  2,1,2,10,100.00,"2015-10-01", "2016-12-21", 6, "d ${DFLTGENRCV}, c 42002"),	-- Water (utility) reimbursement
-- 	(  2,1,2,47, 90.00,"2015-12-01", "2015-12-01", 0, "d ${DFLTGENRCV}, c 42003"),	-- Late payment fee
-- 	(  3,1,3,47,155.00,"2015-12-01", "2015-12-01", 0, "d ${DFLTGENRCV}, c 42003"),	-- Late payment fee
-- 	(  4,1,4,10,100.00,"2015-10-01", "2016-12-31", 6, "d ${DFLTGENRCV}, c 42002"),	-- Water (utility) reimbursement
-- 	(  4,1,4,28,105.00,"2001-10-01", "2016-05-31", 6, "d ${DFLTGENRCV}, c 42007"),	-- carport fee
-- 	(  5,1,5,10,100.00,"2015-10-01", "2016-12-31", 6, "d ${DFLTGENRCV}, c 42002"),	-- Water (utility) reimbursement
-- 	(  5,1,5,28, 35.00,"2015-10-01", "2016-12-01", 6, "d ${DFLTGENRCV}, c 42007"),	-- carport fee
-- 	(  6,1,6,10,100.00,"2015-10-01", "2016-12-31", 6, "d ${DFLTGENRCV}, c 42002"),	-- Water (utility) reimbursement
-- 	(  7,1,7,47, 90.00,"2015-12-01", "2015-12-01", 0, "d ${DFLTGENRCV}, c 42003"),	-- Late payment fee
-- 	(  7,1,7,10,100.00,"2015-10-01", "2016-12-31", 6, "d ${DFLTGENRCV}, c 42002"),	-- Water (utility) reimbursement
-- 	(  7,1,7,11, 17.00,"2015-12-01", "2015-12-31", 0, "d ${DFLTGENRCV}, c 42002");	-- Water (utility) Overage

-- =======================================================================
--  RECEIPTS
-- =======================================================================

-- TODO:  ADD ACCTRULE TO RECEIPTS...

INSERT INTO receipt (BID,RAID,PMTID,Dt,Amount,AcctRule) VALUES
	(1,1,2,"2004-01-01", 1000.00, "d ${DFLTCASH} 1000.00, c 11002 1000.00");			-- 1  Krabappel's initial security deposit
INSERT INTO receiptallocation (RCPTID,Amount,ASMID,AcctRule) VALUES
	(1,1000.00,7, "d ${DFLTCASH} 1000.00, c 11002 1000.00");		

INSERT INTO receipt (BID,RAID,PMTID,Dt,Amount,AcctRule) VALUES
	(1,1,1,"2015-11-21",  294.66, "c ${DFLTGENRCV} 266.67, d ${DFLTCASH} 266.67, c ${DFLTGENRCV} 13.33, d ${DFLTCASH} 13.33, c ${DFLTGENRCV} 5.33, d ${DFLTCASH} 5.33, c ${DFLTGENRCV} 9.33,d ${DFLTCASH} 9.33"); 			-- 2   Krabappel pays her fees in full
INSERT INTO receiptallocation (RCPTID,Amount,ASMID,AcctRule) VALUES
	(2,266.67,1,"c ${DFLTGENRCV} 266.67, d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),	-- rent
	(2, 13.33,3,"c ${DFLTGENRCV} 13.33, d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),		-- Lake View
	(2,  5.33,4,"c ${DFLTGENRCV} 5.33, d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),		-- Fireplace
	(2,  9.33,9,"c ${DFLTGENRCV} 9.33, d ${DFLTCASH} ${aval(${DFLTGENRCV})}");			-- CP001

INSERT INTO receipt (BID,RAID,PMTID,Dt,Amount,AcctRule) VALUES
	(1,8,1,"2015-11-15",1946.68, "d ${DFLTCASH} 1500.00, c 11002 1500.00, c ${DFLTGENRCV} 400.00, d ${DFLTCASH} 400.00, c ${DFLTGENRCV} 16.67, d ${DFLTCASH} 16.67, c ${DFLTGENRCV} 6.67, d ${DFLTCASH} 6.67, c ${DFLTGENRCV} 11.67,d ${DFLTCASH} 11.67, c ${DFLTGENRCV} 11.67,d ${DFLTCASH} 11.67");  			-- 3   Simpson pays his fees in full
INSERT INTO receiptallocation (RCPTID,Amount,ASMID,AcctRule) VALUES
	(3,1500.00, 8,"d ${DFLTCASH} 1500.00,c ${DFLTSECDEPRCV} ${aval(${DFLTCASH})}"),	--  security deposit
	(3, 400.00, 2,"c ${DFLTGENRCV}  400.00,d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),		--  rent
	(3,  16.67, 5,"c ${DFLTGENRCV} 16.67,d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),	--  Lake View
	(3,   6.67, 6,"c ${DFLTGENRCV}  6.67,d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),		--  Fireplace
	(3,  11.67,10,"c ${DFLTGENRCV} 11.67,d ${DFLTCASH} ${aval(${DFLTGENRCV})}"),		--  CP001
	(3,  11.67,11,"c ${DFLTGENRCV} 11.67,d ${DFLTCASH} ${aval(${DFLTGENRCV})}");		--  CP002

-- INSERT INTO receipt (BID,PID,RAID,PMTID,Dt,Amount) VALUES
-- 	(1,1,1,55,"2015-11-11", 750.00);  			-- 4   Security Deposit refuncd to Krabappel
-- INSERT INTO receiptallocation (RCPTID,Amount,ASMID) VALUES
-- 	(4,750.00,7);		-- security deposit return
	-- (1,1,1,"2015-12-06",1100.00),
	-- (2,2,1,"2015-12-15", 805.00), 
	-- (3,3,1,"2015-12-15",1060.00),
	-- (4,4,4,"2015-12-07", 200.00),	
	-- (4,4,4,"2015-12-09", 350.00),	
	-- (6,6,1,"2015-12-15", 995.00),	
	-- (7,7,1,"2015-12-03", 950.00),	
	-- (7,7,1,"2015-12-10", 950.00);

-- =======================================================================
--  JOURNAL MARKERS
-- =======================================================================
INSERT INTO journalmarker (BID,State,DtStart,DtStop) VALUES
	(1, 3, "2015-10-31", "2015-10-31");


-- =======================================================================
--  LEDGERS MARKERS
-- =======================================================================
INSERT INTO ledgermarker (BID,PID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name) VALUES
	(1,1,"RA-1", 2,3,"2015-10-01","2015-10-31",0.0,2,"Krabappel"),						--  1 Krabappel
	(1,2,"RA-2", 2,3,"2015-10-01","2015-10-31",0.0,2,"Flanders"),						--  2 Flanders
	(1,3,"RA-3", 2,3,"2015-10-01","2015-10-31",0.0,2,"Szyslak"),						--  3 Szyslak
	(1,4,"RA-4", 2,3,"2015-10-01","2015-10-31",0.0,2,"Burns"),							--  4 Burns
	(1,5,"RA-5", 2,3,"2015-10-01","2015-10-31",0.0,2,"Muntz"),							--  5 Muntz
	(1,6,"RA-6", 2,3,"2015-10-01","2015-10-31",0.0,2,"Van Houten"),						--  6 Van Houten
	(1,7,"RA-7", 2,3,"2015-10-01","2015-10-31",0.0,2,"Wiggum"),							--  7 Wiggum
	(1,8,"RA-8", 2,3,"2015-10-01","2015-10-31",0.0,2,"Simpson"),						--  8 Simpson
	(1,0,"10001",2,3,"2015-10-01","2015-10-31",0.0,10,"Bank Account FRB 2332352"),
	(1,0,"11001",2,3,"2015-10-01","2015-10-31",0.0,11,"General Accounts Receivable"),
	(1,0,"11002",2,3,"2015-10-01","2015-10-31",0.0,15,"Security Deposit Receivable"),
	(1,0,"11003",2,3,"2015-10-01","2015-10-31",0.0,0,"American Express 93892335 In Process"),
	(1,0,"10004",2,3,"2015-10-01","2015-10-31",0.0,0,"MasterCard/VISA 38992355 in Process"),
	(1,0,"10005",2,3,"2015-10-01","2015-10-31",0.0,0,"Discover 883523553 In Process"),
	(1,0,"23000",2,3,"2015-10-01","2015-10-31",-1000.0,16,"Security Deposit Assessment"),
	(1,0,"40001",2,3,"2015-10-01","2015-10-31",0.0,12,"Gross Schedule Rent"),
	(1,0,"40002",2,3,"2015-10-01","2015-10-31",0.0,0,"Carport Rental"),
	(1,0,"41001",2,3,"2015-10-01","2015-10-31",0.0,14,"Vancancy"),
	(1,0,"41002",2,3,"2015-10-01","2015-10-31",0.0,0,"Administrative Concession"),
	(1,0,"41003",2,3,"2015-10-01","2015-10-31",0.0,0,"Off Line Unit"),
	(1,0,"41004",2,3,"2015-10-01","2015-10-31",0.0,13,"Loss to Lease"),
	(1,0,"41005",2,3,"2015-10-01","2015-10-31",0.0,0,"Employee Concession"),
	(1,0,"41006",2,3,"2015-10-01","2015-10-31",0.0,0,"Payor Concession"),
	(1,0,"41007",2,3,"2015-10-01","2015-10-31",0.0,0,"Bad Debt Write-Off"),
	(1,0,"41999",2,3,"2015-10-01","2015-10-31",0.0,0,"Other Offsets"),
	(1,0,"42001",2,3,"2015-10-01","2015-10-31",0.0,0,"Pet Fees"),
	(1,0,"42002",2,3,"2015-10-01","2015-10-31",0.0,0,"Utility Reimbursement"),
	(1,0,"42003",2,3,"2015-10-01","2015-10-31",0.0,0,"Late Fees"),
	(1,0,"42004",2,3,"2015-10-01","2015-10-31",0.0,0,"Damages"),
	(1,0,"42005",2,3,"2015-10-01","2015-10-31",0.0,0,"Pest Control Reimbursements"),
	(1,0,"42006",2,3,"2015-10-01","2015-10-31",0.0,0,"Security Deposit Forfeiture"),
	(1,0,"42007",2,3,"2015-10-01","2015-10-31",0.0,0,"Carports"),
	(1,0,"42008",2,3,"2015-10-01","2015-10-31",0.0,0,"Utility overage"),
	(1,0,"42999",2,3,"2015-10-01","2015-10-31",0.0,0,"Other Income");
         

