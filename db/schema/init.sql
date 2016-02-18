-- Initialization data

-- ----------------------------------------------------------------------------------------
--    ASSESSMENT TYPES
-- ----------------------------------------------------------------------------------------
USE rentroll;
INSERT INTO assessmenttypes (Name) VALUES
	("Rent"),							-- 1 Rent: the recurring amount due under an Occupancy Agreement.  While most residential leases are one year or less, commecial leases may go on decades.  In those cases there is a formula for rent increases.  For example, our lease to HD Supply in Pacoima provides that we increase rent 7% on each 3rd year anniversary of Lease, and our Viacom lease provides for an annual fixed 2% increase.  Some leases provide for increases based upon CPI.   Rent is tied to a Unit and a Payor.
	("Security Deposit"),				-- 2 Security Deposit: We often assess an amount to secure performance by the tenant under their occupancy agreement.  When collected, this amount is a liability.  A security deposit is either returned (i.e., a negative assessment) or forfeited for Tenant’s non-performance (in which case it is assessed as Forfeited Security Deposit-an item of income).
	("Security Deposit Forfeiture"),	-- 3 Security Deposit Forfeiture: when we collect a security deposit, we record this as a liability (since we owe this money to the tenant in the absence of a breach).  When the tenant breaches, we may apply that deposit to their obligation, which constitutes income (and a corresponding decrease in the liability).  This is a non-recurring item, and will apply to a Payor and a Unit.
	("Application Fees"),				-- 4 Application Fees: the non-recurring fee charged for considering a rental application.  This fee will apply to a Unit and a Payor only if the applicant is accepted as a tenant.  I believe we should set up any applicant that pays an Application fee as a Payor with very limited information just so we have a record the party by whom payments are made and other payment data.  If they end up leasing, they will already be in the system.
	("Landlord Lien Sales"),			-- 5 Landlord Lien Sales: under some state laws, a landlord has a lien that arises by operation of law for personal property that remains in the Unit after a tenancy has terminated.  The landlord is allowed to sell the property, and apply the sales proceeds to the amount owed to the landlord.  This is a non-recurring assessment that will apply to a Unit and a Payor.
	("Pet Fees"),						-- 6 Pet Fees: some properties charge a one-time and/or monthly fee for a pet.  Thus, this may or may not be a recurring fee, and will apply to a Unit and a Payor.
	("Eviction Fees"),					-- 7 Eviction Fees: when we file an eviction on a tenant and the tenant reinstates by paying what is owed, we include a charge for the fees associated with filing the eviction.  This is not a recurring fee and will apply to a Unit and a Payor.
	("Electric Reimbursement"),			-- 8 Electric Reimbursement: when we pay the electic, we charge a fixed fee to the resident for useage up to a certain amount.  This is a recurring fee, and will apply to a Payor and a Unit.
	("Electric Overage"),				-- 9 Electric Overage: when we pay the electric and the resident uses an amount of electricity in excess of the maximum useage, we charge the resident for the overage.  This is calculated monthly and may or may not recur.  The charge will apply to a Payor and a Unit.
	("Water Reimbursement"),			-- 10 Water Reimbursement: when we pay the water, we charge a fixed fee to the resident for useage up to a certain amount.  This is a recurring fee, and will apply to a Payor and a Unit.
	("Water Overage"),					-- 11 Water Overage: when we pay the water and the resident uses an amount of water in excess of the maximum useage, we charge the resident for the overage.  This is calculated monthly and may or may not recur.  The charge will apply to a Payor and a Unit.
	("Trash Fee"),						-- 12 Trash Fee: sometimes we charge a fee for collection of trash.  The charge will apply to a Payor and a Unit.
	("Utility Fine"),					-- 13 Utility Fine: certain jurisdictions (CA mostly) will fine the landlord when water useage exceeds a prescribed amount.  So long as the landlord has taken certain measures for conservation, the landlord is allowed to pass these charges onto the tenants.  This is a non-recurring assessment that is associated with a Payor and a Unit.
	("NSF Fee"),						-- 14 NSF Fee: when a resident bounces a check, we charge an NSF fee.  This is a non-recurring assessment that is associated with a Payor and a Unit.
	("Maintenance Fee"),				-- 15 Maintenance Fee: when residents damage a Unit we may impose a fee for the repair.  We may also charge a fee to replace keys or do other similar items.  This is a non-recurring assessment that is associated with a Payor and a Unit.
	("Fines"),							-- 16 Fines: Certain acts by a resident may result in a fine being imposed, for example a towing fee for parking in a fire lane, or a fine for outside storage after a given number of warnings, etc.  These are non-recurring assessments that are associated with a Payor and a Unit.
	("Month to Month Fee"),				-- 17 Month to Month Fee: When a permanent resident chooses to a month-to-month occupancy agreement, we may charge an increased fee for this.  This will be a recurring fee that is associated with a Payor and a Unit.
	("Cancelation Fees"),				-- 18 Cancelation Fees: For our hotel guests, we may charge a no-show fee for those that booked a room, but never show up.  This is a non-recurring fee that is associated with a Payor, but not a Unit.
	("Housekeeping Fee"),				-- 19 Housekeeping Fee: Some guests and residents want a special cleaning from time to time, and this fee is charged.  This is a non-recurring fee that is associated with a Payor and a Unit.
	("Extra Person Charge"),			-- 20 Extra Person Charge: Some guests bring additional people and are charged a fee.  This is a recurring charge that is associated with a Payor and a Unit.
	("Furniture Rental"),				-- 21 Furniture Rental: Some residents prefer to lease furniture from us, and we charge a fee for this.  This will be a recurring fee that is associated with a Payor and a Unit.
	("Platinum Service Fee"),			-- 22 Platinum Service Fee: This is a fee that we charge for our highest level of service (housekeeping, food and linen service).  This will be a recurring fee that is associated with a Payor and a Unit.
	("Gold Service Fee"),				-- 23 Gold Service Fee: This is a fee that we charge for our intermediate level of service (housekeeping and food).  This will be a recurring fee that is associated with a Payor and a Unit.
	("Silver Service Fee"),				-- 24 Silver Service Fee: This is a fee that we charge for our basic level of service (housekeeping).  This will be a recurring fee that is associated with a Payor and a Unit.
	("Sales Tax"),						-- 25 Sales Tax: This is not a fee, but rather a liability.  Although the sales tax is owed by the purchaser, state law requires that the tax be collected and remitted by the Payee.  In this respect, this assessment is exactly like a security deposit.  It is collected by us and is a liability until further disposition.  (when we remit to the state sales tax agency, we eliminate the liability.)  Not all transactions are taxable.  Generally, hotel stays are taxable, along with furniture rental, and platinum, gold and silver service fees.
	("TOT Tax"),						-- 26 TOT Tax: This stands for Transient Occupancy Tax.  This tax is levied on hotel stays, and varies by jurisdiction.  If applicable, it will always be assessed, and will be a liability until remitted by us to the taxing authority.
	("Reletting Fees"),					-- 27 Reletting Fees: When a tenant moves early, we may charge a fee for reletting their apartment.  This is associated with a Payor and a Unit.
	("Carport Fees"),					-- 28 Carport Fees: I think we should set these up as a Unit, since the rental of a carport is the same as any other unit.
	("Garage Fees"),					-- 29 Garage Fees: I think we should set up a Garage as a special Unit for the same reason.
	("Reserved Parking Fees"),			-- 30 Reserved Parking Fees: Once again, we may be better off to treat these as a special Unit.
	("Transfer fees"),					-- 31 Transfer fees: when a resident moves from one Unit to another Unit, we often charge a fee.  This is non-recurring and will be associated with a Unit (the one from which the occupant moved) and a Payor.
	("Washer/Dryer Fee"),				-- 32 Washer/Dryer Fee: If we provide a washer/dryer, sometimes we charge a fee.  This will be recurring and will be associated with a Payor and a Unit.  (Note: sometimes we charge a washer/dryer connection fee—a fee that is assessed because the particular unit has a connection for washer/dryer.  This will be treated as a Unit Specialty, and not tracked separately as an assessable item.)
	("Association Dues Assessment"),	-- 33 Association Dues Assessment: Sometimes a property will have an owner’s association that charges dues, and these are billed to the tenant.  This will be recurring, and will be associated with a Unit and a Payor.
	("Insurance Reimbursement"),		-- 34 Insurance Reimbursement: Sometime the occupant pays for insurance.  This may or may not be recurring, and will be associated with a Unit and a Payor.
	("Tax Reimbursement "),				-- 35 Tax Reimbursement: Some tenants pay for the ad volarem property taxes associated with their unit.  This may or may not be recurring, and will be associated with a Unit and a Payor.


-- We will also need to create a way to accept other income for certain items.  
-- We should discuss this in a phone call.  These will be associated with a Payor, 
-- but it may be a special case Payor (such as “vending machines”).  
-- Here are the items of special income:

	("Special Event Fees"), 		--  36 Sometimes a guest or resident may use a common area or location and be charged a fee for this.  At the moment, this will cover meeting rooms (unless a particular meeting room is set up as a Unit), catering fees, set up fees, etc.  We may choose to further delineate these items in future versions.  For the moment, we need a place to record this income.  This is non-recurring.  This will be associated with a Payor, but not a Unit.
	("Convenience Store Sales"), 	--  37 We will not create the module right now to track inventory and sales by items, but we need to have a category for this to tie into our main system.  This will be non-recurring, and will be associated with a special Payor (“Convenience Store”) but not a Unit.
	("Courtesy Car Rental"), 		--  38 I would like for each of our cars to be a “Unit” for accounting purposes.  Renting a car is really no different than renting an apartment.  We should discuss this.
	("Vending Income"), 			--  39 This will be non-recurring (even though we collect daily, we do not know the amount).  These will be associated with a special Payor (“Vending”), but not a Unit.
	("Transportation"), 			--  40 This may or may not be recurring.  It is possible that we have a hotel guest that requires daily shuttles and we charge a fixed rate per day.  More often, it is a one-time fee.  This will be associated with a Payor, but not a Unit.
	("Food Sales"), 				--  41 Like Convenience Store Sales, we will need to create a module for this, but for the moment we can enter this as a lump sum.  This will be non-recurring, and may or may not be associated with a particular Payor.  (For example, if our guest brings someone with them, and we charge for their friend to eat, there is really no need to set up a Payor for that guest.  However, if we charge a separate fee to the Guest, it should appear on the Payor’s statement.  An example of this is room service.)
	("Liquor Sales"), 				--  42 
	("WashNFold Income"), 			--  43 We will ultimately need a module for this, but for the moment a lump sum is fine.  This will be non-recurring and will be associated with a particular Payor.
	("Spa Sales"), 					--  44 We will eventually need a module for this, but a lump sum is fine for now.  This will be a non-recurring charge, and may or may not be associated with a particular Payor.
	("Fitness Center"), 			--  45 Sales.  This may be for rental of the fitness center or for trainer fees, etc.  We will eventually need a module, but a lump sum is fine for now.  This will be a non-recurring charge, and may or may not be associate with a particular Payor.
	("CashOverShort"), 				--  46 When we count out cash receipts, we need to account for this.  This, again, is a case of a special Payor, and is not associate with any particular unit.

	("Late Payment Fee"),			--  47 Late payment of a fee
	("Security Deposit"),			--  48 Security Deposit
	("Bad Debt"),					--  49 Bad Debt
	("Off Line"),					--  50 Off Line
	("Concession"),					--  51 concession
	("Employee");					--  52 employee

-- ----------------------------------------------------------------------------------------
--     PAYMENT TYPES
-- ----------------------------------------------------------------------------------------
INSERT INTO paymenttypes (Name,Description) VALUES
	("Check","Personal check from payor"),
	("VISA","Credit card charge"),
	("AMEX", "American Express credit card"),
	("Cash","Cash");


-- ----------------------------------------------------------------------------------------
--     AVAILABILITY TYPES
-- ----------------------------------------------------------------------------------------
INSERT INTO availabilitytypes (Name) VALUES
	("Occupied"),
	("Offline"),
	("Administrative"),
	("Vacant"),
	("Not Ready"),
	("Vacant - Made Ready"),
	("Vacant - Inspected");

