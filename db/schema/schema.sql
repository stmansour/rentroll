-- Conventions used:
--     Table names are all lower case
--     Field names are camel case
--     Money values are all stored as DECIMAL(19,4)

-- ********************************
-- *********  UNIQUE IDS  *********
-- ********************************
--   ASMID = Assessment id
--  ASMTID = assessment type id
-- AVAILID = availability id
--  BLDGID = building id
--     BID = business id
--  DISBID = disbursement id
--     LID = ledger id
--   OFSID = offset id
--     PID = payor id
--     RID = rentable id
--   PMTID = payment type id
--  PRSPID = Prospect id
--    RAID = rental agreement / occupancy agreement
--   RATID = occupancy agreement template id
--    TCID = transactant id
--  RCPTID = receipt id
--    RTID = rentable type id
--     TID = tenant id
--  UNITID = unit id
--    UTID = unit type id
--   USPID = unit specialty id

DROP DATABASE IF EXISTS rentroll;
CREATE DATABASE rentroll;
USE rentroll;
GRANT ALL PRIVILEGES ON rentroll TO 'ec2-user'@'localhost';
GRANT ALL PRIVILEGES ON rentroll.* TO 'ec2-user'@'localhost';


-- **************************************
-- ****                              ****
-- ****       RENTAL AGREEMENT       ****
-- ****                              ****
-- **************************************

-- ===========================================
--   RENTAL AGREEMENT TEMPLATE
-- ===========================================
CREATE TABLE rentalagreementtemplate (
    RATID INT NOT NULL AUTO_INCREMENT,                        -- internal unique id
    ReferenceNumber VARCHAR(35) DEFAULT '',                   -- Occupancy Agreement Reference Number
    RentalAgreementType SMALLINT NOT NULL DEFAULT 0,          -- 1=leasehold, 2=month-to-month, 3=hotel
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RATID)     
);      
        
-- ===========================================
--   RENTAL AGREEMENT
-- ===========================================
CREATE TABLE rentalagreement (        
    RAID INT NOT NULL AUTO_INCREMENT,                         -- internal unique id
    RATID INT NOT NULL DEFAULT 0,                             -- reference to Occupancy Master Agreement
    BID INT NOT NULL DEFAULT 0,                               -- business (so that we can process by business)
    RID INT NOT NULL DEFAULT 0,                               -- rentable id
    UNITID INT NOT NULL DEFAULT 0,                            -- associated unit
    PID INT NOT NULL DEFAULT 0,                               -- who is the payor for this agreement
    PrimaryTenant INT NOT NULL DEFAULT 0,                     -- TID of primary tenant.  
    RentalStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',
    RentalStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',
    Renewal SMALLINT NOT NULL DEFAULT 0,                      -- month to month automatic renewal, lease extension options, none.
    SpecialProvisions VARCHAR(1024) NOT NULL DEFAULT '',  
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RAID)
);

-- ===========================================
--   UNIT TENANTS
-- ===========================================
-- query this table for rows where RAID=(the rental agreement for the unit)
-- the return list will be the TIDs of all tenants in that unit
CREATE TABLE unittenants (
    RAID INT NOT NULL DEFAULT 0,                              -- the unit's occupancy agreement
    TID INT NOT NULL DEFAULT 0                                -- the tenant
);


-- **************************************
-- ****                              ****
-- ****          BUSINESS            ****
-- ****                              ****
-- **************************************
-- Unit Types - associated with a business are stored
-- in the rentabletypes table and have the business PID
-- Occupancy Type List - hardcoded

CREATE TABLE business (
    BID INT NOT NULL AUTO_INCREMENT,
    Address VARCHAR(35) NOT NULL DEFAULT '',
    Address2 VARCHAR(35) NOT NULL DEFAULT '',
    City VARCHAR(25) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(11) NOT NULL DEFAULT '',
    Country VARCHAR(25) NOT NULL DEFAULT '',
    Phone VARCHAR(25) NOT NULL DEFAULT '',
    Name VARCHAR(50) NOT NULL DEFAULT '',
    DefaultOccupancyType SMALLINT NOT NULL DEFAULT 0,       -- default for every unit in the building: 0=unset, 1=daily, 2=weekly, 3=monthly, 4=quarterly, 5=yearly
    ParkingPermitInUse SMALLINT NOT NULL DEFAULT 0,         -- yes/no  0 = no, 1 = yes
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (BID)
);

-- ===========================================
--   RENTABLE TYPES 
-- ===========================================
CREATE TABLE rentabletypes (
    RTID INT NOT NULL AUTO_INCREMENT,
    BID INT NOT NULL DEFAULT 0,                            -- associated business id
    Name VARCHAR(256) NOT NULL DEFAULT '',
    Amount Decimal(19,4) NOT NULL DEFAULT 0.0,              -- rental price
    Frequency INT NOT NULL DEFAULT 0,                       -- price accrual frequency
    Proration INT NOT NULL DEFAULT 0,                       --  prorate frequency
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RTID)
);

-- ===========================================
--   UNIT TYPES 
-- ===========================================
--  unit types are associated with a particular business
--  There is no "global" unit type since they are all different enough where
--  it does not make sense to try to share them across properties.
--  Offset=Debit=positive
--  Assessment=Credit=negative
CREATE TABLE unittypes (
    UTID INT NOT NULL AUTO_INCREMENT,
    BID INT NOT NULL DEFAULT 0,                            -- associated business id
    Style CHAR(15) NOT NULL DEFAULT '',
    Name VARCHAR(256) NOT NULL DEFAULT '',
    SqFt MEDIUMINT NOT NULL DEFAULT 0,
    MarketRate Decimal(19,4) NOT NULL DEFAULT 0.0,          -- market rate for this unit
    Frequency INT NOT NULL DEFAULT 0,                       -- MarketRate accrual frequency
    Proration INT NOT NULL DEFAULT 0,                       --  prorate frequency
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (UTID)
);


-- ===========================================
--   UNIT SPECIALTY TYPES
-- ===========================================
-- a collection of unit specialties that are available.
-- different units may be more or less desirable based upon special characteristics
-- of the unit, such as Lake View, Courtyard View, Washer Dryer Connections, 
-- Washer Dryer provided, close to parking, better views, fireplaces, special 
-- remodeling or finishes, etc.  This is where those special characteristics are defined
CREATE TABLE unitspecialtytypes (
    USPID INT NOT NULL AUTO_INCREMENT,
    BID INT NOT NULL,
    Name VARCHAR(25) NOT NULL DEFAULT '',
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    Description VARCHAR(256) NOT NULL DEFAULT '',
    PRIMARY KEY (USPID)
);

-- **************************************
-- ****                              ****
-- ****       FINANCIAL TYPES        ****
-- ****                              ****
-- **************************************


-- ===========================================
--   ASSESSMENT TYPES
-- ===========================================
-- this table list all the pre-defined assessments
-- this will include offsets and disbursements
CREATE TABLE assessmenttypes (
    ASMTID INT NOT NULL AUTO_INCREMENT,             -- what type of assessment
    Name VARCHAR(35) NOT NULL DEFAULT '',           -- name for the assessment

    -- TODO: Type needs to be removed
    Type SMALLINT NOT NULL DEFAULT 0,               -- normal case, positive number is: 0 = DEBIT, 1 = CREDIT
    LastModTime TIMESTAMP,                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,         -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMTID)
);

-- ===========================================
--   PAYMENT TYPES
-- ===========================================
CREATE TABLE paymenttypes (
    PMTID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Name VARCHAR(25) NOT NULL DEFAULT '',
    Description VARCHAR(256) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PMTID)
);

-- ===========================================
--   AVAILABILITY TYPES
-- ===========================================
-- Examples: Occupied, Offline, Administrative, Vacant - Not Ready, 
--           Vacant - Made Ready, Vacant - Inspected, plus custom values
-- Custom values will be added with their own uniq AVAILID
CREATE TABLE availabilitytypes (
    AVAILID INT NOT NULL AUTO_INCREMENT,
    Name VARCHAR(35) NOT NULL DEFAULT '',
    PRIMARY KEY (AVAILID)
);

-- ***************************************
-- ****                               ****
-- ****       COMMON TYPES            ****
-- **** SCOPED TO A SPECIFIC PROPERTY ****
-- ****                               ****
-- ***************************************

-- This describes the world of assessments for a particular business
-- Query this table for a particular BID, the solution set is the list
-- of assessments for that particular business.

-- applicable assessments for a specific business
CREATE TABLE businessassessments (
    BID INT NOT NULL DEFAULT 0,
    ASMTID INT NOT NULL DEFAULT 0
);

-- applicable assessments for a specific business
CREATE TABLE businesspaymenttypes (
    BID INT NOT NULL DEFAULT 0,
    PMTID MEDIUMINT NOT NULL DEFAULT 0
);

-- **************************************
-- ****                              ****
-- ****          BUILDING            ****
-- ****                              ****
-- **************************************
CREATE TABLE building (
    BLDGID INT NOT NULL AUTO_INCREMENT,           -- unique id for this building
    BID INT NOT NULL DEFAULT 0,                  -- which business it belongs to
    Address VARCHAR(35) NOT NULL DEFAULT '',      -- building address
    Address2 VARCHAR(35) NOT NULL DEFAULT '',       
    City VARCHAR(25) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(10) NOT NULL DEFAULT '',
    Country VARCHAR(25) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                        -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,       -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (BLDGID)
);

-- **************************************
-- ****                              ****
-- ****          RENTABLE            ****
-- ****                              ****
-- **************************************
-- Any item that can be rented to a Payor
-- It will be tracked, it will have a Ledger
-- It will be included in the rentroll processing
CREATE TABLE rentable (
    RID INT NOT NULL AUTO_INCREMENT,
    LID INT NOT NULL DEFAULT 0,                             -- which ledger keeps track of what's owed on this rentable
    RTID INT NOT NULL DEFAULT 0,                            -- what sort of a rentable is this?
    BID INT NOT NULL DEFAULT 0,                            -- Property associated with this rentable
    -- PID INT NOT NULL DEFAULT 0,                             -- who is responsible for paying
    -- RAID INT NOT NULL DEFAULT 0,                            -- rental agreement
    UNITID INT NOT NULL DEFAULT 0,                          -- unit (if applicable)
    Name VARCHAR(10) NOT NULL DEFAULT '',                   -- name unique to the instance "101" for a room number 744 carport number, etc 
    Assignment SMALLINT NOT NULL DEFAULT 0,                 -- Pre-assign or assign at occupy commencement
    Report SMALLINT NOT NULL DEFAULT 1,                     -- 1 = apply to rentroll, 0 = skip on rentroll
    DefaultOccType SMALLINT NOT NULL DEFAULT 0,         -- unset, short term, longterm
    OccType SMALLINT NOT NULL DEFAULT 0,                -- unset, short term, longterm
    LastModTime TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RID)
);


-- **************************************
-- ****                              ****
-- ****            UNIT              ****
-- ****                              ****
-- **************************************
-- Fields unique to an apartment or hotel room 
CREATE TABLE unit (
    UNITID INT NOT NULL AUTO_INCREMENT,                 -- unique id for this unit -- it is unique across all properties and buildings
    RID INT NOT NULL DEFAULT 0,                         -- associated rentable
    BLDGID INT NOT NULL DEFAULT 0,                      -- which building
    UTID INT NOT NULL DEFAULT 0,                        -- which unit type
    AVAILID INT NOT NULL DEFAULT 0,                     -- how is the unit made available
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (UNITID)
    -- Abbreviation VARCHAR(20),                        -- unit abbreviation  -- REMOVED - it's part of unittype
);

-- ===========================================
--   UNIT SPECIALTIES
-- ===========================================
-- For each unit, what specialties does it have...
-- this is simply a list of USPIDs.
-- Selecting all entries where the unit == UNITID
-- will be the list of all the unit specialties for that unit.
CREATE TABLE unitspecialties (
    BID INT NOT NULL DEFAULT 0,                         -- the business
    UNITID INT NOT NULL DEFAULT 0,                      -- unique id of unit
    USPID INT NOT NULL DEFAULT 0                        -- unique id of specialty (see Table unitspecialties)
);


-- **************************************
-- ****                              ****
-- ****        ASSESSMENTS           ****
-- ****                              ****
-- **************************************
-- charges associated with a rentable
CREATE TABLE assessments (
    ASMID INT NOT NULL AUTO_INCREMENT,
    BID INT NOT NULL DEFAULT 0,                             -- Business id
    RID INT NOT NULL DEFAULT 0,                             -- rental id
    UNITID INT NOT NULL DEFAULT 0,                          -- unit associated with this assessment (could be "subid")
    ASMTID INT NOT NULL DEFAULT 0,                          -- what type of assessment (ex: Rent, SecurityDeposit, ...)
    RAID INT NOT NULL DEFAULT 0,                            -- Associated Rental Agreement ID
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- Assessment amount
    Start DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- epoch date for the assessment - recurrences are based on this date
    Stop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00',   -- stop date - when the tenant moves out or when the charge is no longer applicable
    Frequency SMALLINT NOT NULL DEFAULT 0,                  -- 0 = one time only, 1 = daily, 2 = weekly, 3 = monthly,   4 = yearly
    ProrationMethod SMALLINT NOT NULL DEFAULT 0,            -- 
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',              -- Accounting rule - which acct debited, which credited
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMID)
);


-- **************************************
-- ****                              ****
-- ****           PEOPLE             ****
-- ****                              ****
-- **************************************

-- ===========================================
--   TRANSACTANT
-- ===========================================
-- transactant - fields common to all people and
CREATE TABLE transactant (
    TCID INT NOT NULL AUTO_INCREMENT,                   -- unique id of unit
    TID INT NOT NULL DEFAULT 0,                         -- associated tenant id
    PID INT NOT NULL DEFAULT 0,                         -- associated payor id
    PRSPID INT NOT NULL DEFAULT 0,                      -- associated prospect id
    FirstName VARCHAR(35) NOT NULL DEFAULT '',
    MiddleName VARCHAR(35) NOT NULL DEFAULT '',
    LastName VARCHAR(35) NOT NULL DEFAULT '', 
    PrimaryEmail VARCHAR(35) NOT NULL DEFAULT '',
    SecondaryEmail VARCHAR(35) NOT NULL DEFAULT '',
    WorkPhone VARCHAR(25) NOT NULL DEFAULT '',
    CellPhone VARCHAR(25) NOT NULL DEFAULT '',
    Address VARCHAR(35) NOT NULL DEFAULT '',            -- person address
    Address2 VARCHAR(35) NOT NULL DEFAULT '',       
    City VARCHAR(25) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(10) NOT NULL DEFAULT '',
    Country VARCHAR(25) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

-- ===========================================
--   PROSPECT
-- ===========================================
CREATE TABLE prospect (
    PRSPID INT NOT NULL AUTO_INCREMENT,                 -- unique id of this prospect
    TCID INT NOT NULL DEFAULT 0,                        -- associated transactant (has Name and all contact info)
    ApplicationFee DECIMAL(19,4) NOT NULL DEFAULT 0.0,  -- if non-zero this prospect is an applicant
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PRSPID)
);

-- ===========================================
--   TENANT
-- ===========================================
CREATE TABLE tenant (
    TID INT NOT NULL AUTO_INCREMENT,                    -- unique id of this tenant
    TCID INT NOT NULL,                                  -- associated transactant
    Points INT NOT NULL DEFAULT 0,                      -- bonus points for this tenant
    CarMake VARCHAR(25) NOT NULL DEFAULT '',
    CarModel VARCHAR(25) NOT NULL DEFAULT '',
    CarColor VARCHAR(25) NOT NULL DEFAULT '',
    CarYear INT NOT NULL DEFAULT 0,
    LicensePlateState VARCHAR(35) NOT NULL DEFAULT '',
    LicensePlateNumber VARCHAR(35) NOT NULL DEFAULT '',
    ParkingPermitNumber VARCHAR(35) NOT NULL DEFAULT '',
    AccountRep INT NOT NULL DEFAULT 0,                              -- Phonebook UID of account rep
    DateofBirth DATE NOT NULL DEFAULT '1970-01-01T00:00:00',
    EmergencyContactName VARCHAR(35) NOT NULL DEFAULT '',
    EmergencyContactAddress VARCHAR(35) NOT NULL DEFAULT '',
    EmergencyContactTelephone VARCHAR(25) NOT NULL DEFAULT '',
    EmergencyAddressEmail VARCHAR(35) NOT NULL DEFAULT '',
    AlternateAddress VARCHAR(35) NOT NULL DEFAULT '',
    ElibigleForFutureOccupancy SMALLINT NOT NULL DEFAULT 1,         -- yes/no
    Industry VARCHAR(35) NOT NULL DEFAULT '',                       -- (e.g., construction, retail, banking etc.)
    Source  VARCHAR(35) NOT NULL DEFAULT '',                        -- (e.g., resident referral, newspaper, radio, post card, expedia, travelocity, etc.)
    InvoicingCustomerNumber VARCHAR(35) NOT NULL DEFAULT '',        -- [drawn from the invoicing section] [only applies if invoicing authorization has been provide
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TID)
);

-- ===========================================
--   PAYOR
-- ===========================================
CREATE TABLE payor  (
    PID INT NOT NULL AUTO_INCREMENT,                          -- unique id of this payor
    TCID INT NOT NULL,                                        -- associated transactant
    CreditLimit DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    EmployerName  VARCHAR(35) NOT NULL DEFAULT '',
    EmployerStreetAddress VARCHAR(35) NOT NULL DEFAULT '',
    EmployerCity VARCHAR(35) NOT NULL DEFAULT '',
    EmployerState VARCHAR(35) NOT NULL DEFAULT '',
    EmployerZipcode VARCHAR(35) NOT NULL DEFAULT '',
    Occupation VARCHAR(35) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PID)
);

-- **************************************
-- ****                              ****
-- ****           RECEIPTS           ****
-- ****                              ****
-- **************************************
CREATE TABLE receipt (
    RCPTID INT NOT NULL AUTO_INCREMENT,                       -- unique id for this receipt
    BID INT NOT NULL DEFAULT 0,
    PID INT NOT NULL DEFAULT 0,
    RAID INT NOT NULL DEFAULT 0,
    PMTID INT NOT NULL DEFAULT 0,
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    -- ApplyToGeneralReceivable DECIMAL(19,4),                   -- Breakdown is in receiptallocation table
    -- ApplyToSecurityDeposit DECIMAL(19,4),                     -- Can we just handle this as part of receipt allocation
    PRIMARY KEY (RCPTID)
);

CREATE TABLE receiptallocation (
    RCPTID INT NOT NULL DEFAULT 0,                              -- sum of all amounts in this table with RCPTID must equal the receipt with RCPTID in receipt table
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID INT NOT NULL DEFAULT 0                                -- the id of the assessment that caused this payment, if null then credit payors acct this amount,
                                                                -- if > than owed, then credit the overage into the payor's account
);  

-- **************************************
-- ****                              ****
-- ****           JOURNAL            ****
-- ****                              ****
-- **************************************
CREATE TABLE journal (
    JID INT NOT NULL AUTO_INCREMENT,                            -- a journal entry
    BID INT NOT NULL DEFAULT 0,                                 -- Business id
    RAID INT NOT NULL DEFAULT 0,                                -- associated rental agreement
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- date when it occurred
    Amount DECIMAL(19.4) NOT NULL DEFAULT 0.0,                  -- how much
    Type SMALLINT NOT NULL DEFAULT 0,                           -- 0 = unknown, 1 = assessment, 2 = payment/receipt
    ID INT NOT NULL DEFAULT 0,                                  -- if Type == 1 then it is the ASMID that caused this entry, of Type ==2 then it is the RCPTID
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (JID)
);

CREATE TABLE journalallocation (
    JID INT NOT NULL DEFAULT 0,                                 -- sum of all amounts in this table with RCPTID must equal the receipt with RCPTID in receipt table
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID INT NOT NULL DEFAULT 0 
);  

CREATE TABLE journalmarker (
    JMID INT NOT NULL AUTO_INCREMENT,
     State SMALLINT NOT NULL DEFAULT 0,                        -- 0 = unknown, 1 = Available, 2 = closed
   LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (JMID)
);

-- **************************************
-- ****                              ****
-- ****           LEDGERS            ****
-- ****                              ****
-- **************************************
CREATE TABLE ledger (
    LID INT NOT NULL AUTO_INCREMENT,                          -- unique id for this Ledger
    GLNumber VARCHAR(10) NOT NULL DEFAULT '',                 -- if not '' then it's a link a QB  GeneralLedger (GL)account
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- balance date and time
    Status SMALLINT NOT NULL DEFAULT 0,                       -- Whether a GL Account is currently active or inactive
    Type SMALLINT NOT NULL DEFAULT 0,                         -- Classification of a GL Account as one of the following:  bank, accounts receivable, liability, 
    Balance DECIMAL(19,4) NOT NULL DEFAULT 0.0,               -- balance amount
    Name VARCHAR(50) NOT NULL DEFAULT '',                     -- 
    PRIMARY KEY (LID)        
);

CREATE TABLE ledgermarker (
    JMID INT NOT NULL AUTO_INCREMENT,
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    State SMALLINT NOT NULL DEFAULT 0,                        -- 0 = unknown, 1 = Available, 2 = closed
    PRIMARY KEY (JMID)
);

