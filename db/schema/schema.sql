-- Conventions used:
--     Table names are all lower case
--     Field names are camel case
--     Money values are all stored as DECIMAL(19,4)

-- ********************************
-- *********  UNIQUE IDS  *********
-- ********************************
--    PRID = property id
--    UTID = unit type id
--   USPID = unit specialty id
--   OFSID = offset id
--  ASMTID = assessment type id
--   PMTID = payment type id
-- AVAILID = availability id
--  BLDGID = building id
--  UNITID = unit id
--    TCID = transactant id
--     TID = tenant id
--     PID = payor id
--   OATID = occupancy agreement template id
--    OAID = occupancy agreement
--  RCPTID = receipt id
--  DISBID = disbursement id
--     LID = ledger id

DROP DATABASE IF EXISTS rentroll;
CREATE DATABASE rentroll;
USE rentroll;
GRANT ALL PRIVILEGES ON rentroll TO 'ec2-user'@'localhost';
GRANT ALL PRIVILEGES ON rentroll.* TO 'ec2-user'@'localhost';


-- **************************************
-- ****                              ****
-- ****      OCCUPANCY AGREEMENT     ****
-- ****                              ****
-- **************************************

CREATE TABLE occupancyagreementtemplate (
    OATID INT NOT NULL AUTO_INCREMENT,                        -- internal unique id
    ReferenceNumber VARCHAR(35) DEFAULT '',                   -- Occupancy Agreement Reference Number
    RentalAgreementType SMALLINT NOT NULL DEFAULT 0,          -- 1=leasehold, 2=month-to-month, 3=hotel
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (OATID)     
);      
        
CREATE TABLE occupancyagreement (        
    OAID INT NOT NULL AUTO_INCREMENT,                         -- internal unique id
    OATID INT NOT NULL DEFAULT 0,                             -- reference to Occupancy Master Agreement
    PRID INT NOT NULL DEFAULT 0,                              -- property (so that we can process by property)
    UNITID INT NOT NULL DEFAULT 0,                            -- associated unit
    PID INT NOT NULL DEFAULT 0,                               -- who is the payor for this agreement
    PrimaryTenant INT NOT NULL DEFAULT 0,                     -- TID of primary tenant.  
    OccupancyStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',
    OccupancyStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',
    Renewal SMALLINT NOT NULL DEFAULT 0,                      -- month to month automatic renewal, lease extension options, none.
    ProrationMethod SMALLINT NOT NULL DEFAULT 0,              -- daily, monthly based on actual, monthly based on 30 days
    SecurityDepositAmount DECIMAL(19,4),
    SpecialProvisions VARCHAR(1024) NOT NULL DEFAULT '',  
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (OAID)
);

-- query this table for rows where OAID=(the occupancy agreement for the unit)
-- the return list will be the TIDs of all tenants in that unit
CREATE TABLE unittenants (
    OAID INT NOT NULL DEFAULT 0,                              -- the unit's occupancy agreement
    TID INT NOT NULL DEFAULT 0                                -- the tenant
);



-- **************************************
-- ****                              ****
-- ****          PROPERTY            ****
-- ****                              ****
-- **************************************
-- Unit Types - associated with a property are stored
-- in the unittypes table and have the property PID
-- Occupancy Type List - hardcoded

CREATE TABLE property (
    PRID INT NOT NULL AUTO_INCREMENT,
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
    PRIMARY KEY (PRID)
);

-- unit types are associated with a particular property
-- There is no "global" unit type since they are all different enough where
-- it does not make sense to try to share them across properties
CREATE TABLE unittypes (
    UTID INT NOT NULL AUTO_INCREMENT,
    PRID INT NOT NULL,                                      -- associated property id
    Style CHAR(15) NOT NULL DEFAULT '',
    Name VARCHAR(256) NOT NULL DEFAULT '',
    SqFt MEDIUMINT NOT NULL DEFAULT 0,
    MarketRate Decimal(19,4) NOT NULL DEFAULT 0.0,          -- market rate for this unit
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (UTID)
);

-- a collection of unit specialties that are available.
-- different units may be more or less desirable based upon special characteristics
-- of the unit, such as Lake View, Courtyard View, Washer Dryer Connections, 
-- Washer Dryer provided, close to parking, better views, fireplaces, special 
-- remodeling or finishes, etc.  This is where those special characteristics are defined
CREATE TABLE unitspecialtytypes (
    USPID INT NOT NULL AUTO_INCREMENT,
    PRID INT NOT NULL,
    Name VARCHAR(25) NOT NULL DEFAULT '',
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    Description VARCHAR(256) NOT NULL DEFAULT '',
    PRIMARY KEY (USPID)
);


-- **************************************
-- ****                              ****
-- ****       COMMON TYPES           ****
-- ****                              ****
-- **************************************
-- this table list all the pre-defined assessments
-- this will include offsets and disbursements
CREATE TABLE assessmenttypes (
    ASMTID INT NOT NULL AUTO_INCREMENT,                        -- what type of assessment
    Name VARCHAR(35) NOT NULL DEFAULT '',                      -- name for the assessment
    LastModTime TIMESTAMP,                                     -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMTID)
);

CREATE TABLE paymenttypes (
    PMTID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Name VARCHAR(25),
    Description VARCHAR(256) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PMTID)
);

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
-- This describes the world of assessments for a particular property
-- Query this table for a particular PRID, the solution set is the list
-- of assessments for that particular property.

-- applicable assessments for a specific property
CREATE TABLE propertyassessments (
    PRID INT NOT NULL DEFAULT 0,
    ASMTID INT NOT NULL DEFAULT 0
);

-- applicable assessments for a specific property
CREATE TABLE propertypaymenttypes (
    PRID INT NOT NULL DEFAULT 0,
    PMTID MEDIUMINT NOT NULL DEFAULT 0
);

-- **************************************
-- ****                              ****
-- ****          BUILDING            ****
-- ****                              ****
-- **************************************
CREATE TABLE building (
    BLDGID INT NOT NULL AUTO_INCREMENT,           -- unique id for this building
    PRID INT NOT NULL DEFAULT 0,                  -- which property it belongs to
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
-- ****           UNITS              ****
-- ****                              ****
-- **************************************
CREATE TABLE unit (
    UNITID INT NOT NULL AUTO_INCREMENT,                 -- unique id for this unit -- it is unique across all properties and buildings
    PRID INT NOT NULL DEFAULT 0,                        -- which property it belongs to
    BLDGID INT NOT NULL DEFAULT 0,                      -- which building
    UTID INT NOT NULL DEFAULT 0,                        -- which unit type
    LID INT NOT NULL DEFAULT 0,                         -- which ledger keeps track of what's owed on this unit
    AVAILID INT NOT NULL DEFAULT 0,                     -- how is the unit made available
    Payor VARCHAR(35) NOT NULL DEFAULT '',              -- who is responsible for paying
    Name VARCHAR(10) NOT NULL DEFAULT '',               -- ex: 101, 1023B, 107A          
    -- Abbreviation VARCHAR(20),                        -- unit abbreviation  -- REMOVED - it's part of unittype
    DefaultOccType SMALLINT NOT NULL DEFAULT 0,         -- unset, short term, longterm
    OccType SMALLINT NOT NULL DEFAULT 0,                -- unset, short term, longterm
    ScheduledRent DECIMAL(19,4) NOT NULL DEFAULT 0.0,   -- budgeted rent for this unit -  The MARKET RATE
    Assignment SMALLINT NOT NULL DEFAULT 0,             -- Pre-assign or assign at occupy commencement
    Report SMALLINT NOT NULL DEFAULT 1,                 -- 1 = apply to rentroll, 0 = skip on rentroll
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (UNITID)
);

-- charges associated with a unit
-- offsets are entered here as negative values
-- disbursements go here too
-- Query by UNITID where RunStart < Stop 
CREATE TABLE assessments (
    ASMID INT NOT NULL AUTO_INCREMENT,
    UNITID INT NOT NULL DEFAULT 0,                          -- unit associated with this assessment
    ASMTID INT NOT NULL DEFAULT 0,                          -- what type of assessment (ex: Rent, SecurityDeposit, ...)
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- assessment amount
    Start DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- epoch date for the assessment - recurrences are based on this date
    Stop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00',   -- stop date - when the tenant moves out or when the charge is no longer applicable
    Frequency SMALLINT NOT NULL DEFAULT 0,                  -- 0 = one time only, 1 = daily, 2 = weekly, 3 = monthly,   4 = yearly
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMID)
);


-- For each unit, what specialties does it have...
-- this is simply a list of USPIDs.
-- Selecting all entries where the unit == UNITID
-- will be the list of all the unit specialties for that unit.
CREATE TABLE unitspecialties (
    UNITID INT NOT NULL DEFAULT 0,                    -- unique id of unit
    USPID INT NOT NULL DEFAULT 0                      -- unique id of specialty (see Table unitspecialties)
);

-- **************************************
-- ****                              ****
-- ****           PEOPLE             ****
-- ****                              ****
-- **************************************

-- transactant - fields common to all people and
-- ids of prospect/tenant/payor as appropriate
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

CREATE TABLE prospect (
    PRSPID INT NOT NULL AUTO_INCREMENT,                 -- unique id of this prospect
    TCID INT NOT NULL DEFAULT 0,                        -- associated transactant (has Name and all contact info)
    ApplicationFee DECIMAL(19,4) NOT NULL DEFAULT 0.0,  -- if non-zero this prospect is an applicant
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PRSPID)
);

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

CREATE TABLE payor  (
    PID INT NOT NULL AUTO_INCREMENT,                          -- unique id of this payor
    TCID INT NOT NULL,                                        -- associated transactant
    CreditLimit DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    --
    --  JOE: there must be more credit information???
    --
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
    PID INT NOT NULL DEFAULT 0,
    OAID INT NOT NULL DEFAULT 0,
    -- SecurityDepositReceivable also a calculated value
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ApplyToGeneralReceivable DECIMAL(19,4),                   -- Breakdown is in receiptallocation table
    ApplyToSecurityDeposit DECIMAL(19,4),                     -- Can we just handle this as part of receipt allocation
    PRIMARY KEY (RCPTID)
);

CREATE TABLE receiptallocation (
    RCTPID INT NOT NULL DEFAULT 0,
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AccountNo VARCHAR(10) 
);

-- **************************************
-- ****                              ****
-- ****           LEDGERS            ****
-- ****                              ****
-- **************************************
CREATE TABLE ledger (
    LID INT NOT NULL AUTO_INCREMENT,                          -- unique id for this Ledger
    AccountNo VARCHAR(10) NOT NULL DEFAULT '',                -- if not '' then it's a link a QB account
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- balance date and time
    Balance DECIMAL(19,4) NOT NULL DEFAULT 0.0,               -- balance amount
    Deposit DECIMAL(14,4) NOT NULL DEFAULT 0.0,               -- deposit balance
    PRIMARY KEY (LID)        
);


