-- Conventions used:
--     Table names are all lower case
--     Field names are camel case
--     Money values are all stored as DECIMAL(19,4)

-- ********************************
-- *********  UNIQUE IDS  *********
-- ********************************
-- ASMID = Assessment id
-- ASMTID = assessment type id
-- AVAILID = availability id
-- BID = business id
-- BLDGID = building id
-- CID = custom attribute id
-- DISBID = disbursement id
-- JAID = journal allocation id
-- JID = journal id
-- JMID = journal marker id
-- LID = ledger id
-- LMID = ledger marker id
-- OFSID = offset id
-- PID = payor id
-- PMTID = payment type id
-- PRSPID = Prospect id
-- RAID = rental agreement / occupancy agreement
-- RATID = rental agreement template id
-- RCPTID = receipt id
-- RID = rentable id
-- RSPID = unit specialty id
-- RTID = rentable type id
-- TCID = transactant id
-- TCID = transactant id
-- TID = tenant id

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
    RATID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal unique id
    ReferenceNumber VARCHAR(100) DEFAULT '',                   -- Occupancy Agreement Reference Number
    RentalAgreementType SMALLINT NOT NULL DEFAULT 0,          -- 1=leasehold, 2=month-to-month, 3=hotel
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RATID)     
);      
        
-- ===========================================
--   RENTAL AGREEMENT
-- ===========================================
CREATE TABLE rentalagreement (        
    RAID BIGINT NOT NULL AUTO_INCREMENT,                      -- internal unique id
    RATID BIGINT NOT NULL DEFAULT 0,                          -- reference to Rental Template (Occupancy Master Agreement)
    BID BIGINT NOT NULL DEFAULT 0,                            -- business (so that we can process by business)
    PrimaryTenant BIGINT NOT NULL DEFAULT 0,                  -- TID of primary tenant.  
    RentalStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- date when rental starts
    RentalStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- date when rental stops
    OccStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',     -- date when Occupancy starts
    OccStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when Occupancy stops
    Renewal SMALLINT NOT NULL DEFAULT 0,                      -- 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
    SpecialProvisions VARCHAR(1024) NOT NULL DEFAULT '',      -- free-form text
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RAID)
);

CREATE TABLE agreementrentables (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    RID BIGINT NOT NULL DEFAULT 0,                            -- rentable id
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this rentable was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this rentable was no longer being billed to this agreement
);

CREATE TABLE agreementpayors (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    PID BIGINT NOT NULL DEFAULT 0,                            -- who is the payor for this agreement
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this payor was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this payor was no longer being billed to this agreement
);

CREATE TABLE agreementtenants (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- the unit's occupancy agreement
    TID BIGINT NOT NULL DEFAULT 0,                            -- the tenant
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this tenant was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this tenant was no longer being billed to this agreement
);



-- **************************************
-- ****                              ****
-- ****    USER DEFINED ATTRIBUTES   ****
-- ****                              ****
-- ************************************** 
CREATE TABLE customattr (
    CID BIGINT NOT NULL AUTO_INCREMENT,        -- unique identifer for this custom attribute
    Type SMALLINT NOT NULL DEFAULT 0,          -- 0 = string, 1 = int64, 2 = float64
    Name VARCHAR (100) NOT NULL DEFAULT '',    -- a name
    Value VARCHAR (256) NOT NULL DEFAULT '',   -- its value in string form
    LastModTime TIMESTAMP,                     -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,    -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (CID)
);

CREATE TABLE customattrref (
    ElementType BIGINT NOT NULL,               -- for what type of object is this a ref:  1=Person, 2=Company, 3=Business-Unit, 4=executable service, 5=RentableType
    ID          BIGINT NOT NULL,               -- the UID of the object type. That is, if ObjectType == 5, the ID is the RTID (rentable type id)
    CID         BIGINT NOT NULL                -- uid of the custom attribute
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
    BID BIGINT NOT NULL AUTO_INCREMENT,
    DES VARCHAR(100) NOT NULL DEFAULT '',               -- this is the link to phonebook
    Name VARCHAR(100) NOT NULL DEFAULT '',
    DefaultAccrual SMALLINT NOT NULL DEFAULT 0,         -- default for every unit in the building: 0=unset, 1=hourly, 2=daily, 3=weekly, 4=monthly, 5=quarterly, 6=yearly
    ParkingPermitInUse SMALLINT NOT NULL DEFAULT 0,     -- yes/no  0 = no, 1 = yes
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (BID)
);



-- ===========================================
--   RENTABLE TYPES 
-- ===========================================
CREATE TABLE rentabletypes (
    RTID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                          -- associated business id
    Style CHAR(15) NOT NULL DEFAULT '',                     -- need not be unique
    Name VARCHAR(256) NOT NULL DEFAULT '',                  -- must be unique
    Accrual BIGINT NOT NULL DEFAULT 0,                    -- price accrual frequency
    Proration BIGINT NOT NULL DEFAULT 0,                    --  prorate frequency
    Report SMALLINT NOT NULL DEFAULT 0,
    ManageToBudget SMALLINT NOT NULL DEFAULT 0,             -- 0 do not manage this category of rentable to budget, 1 = manage to budget defined by MarketRate
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RTID)
);

CREATE TABLE rentablemarketrate (
    RTID BIGINT NOT NULL DEFAULT 0,                           -- associated rentable type
    MarketRate DECIMAL(19,4) NOT NULL DEFAULT 0.0,            -- market rate for the time range
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DtStop DATETIME NOT NULL DEFAULT '9999-12-31 23:59:59'    -- assume it's unbounded. if an updated Market rate is added, set this to the stop date
);


-- ===========================================
--   RENTABLE SPECIALTY TYPES
-- ===========================================
-- a collection of unit specialties that are available.
-- different units may be more or less desirable based upon special characteristics
-- of the unit, such as Lake View, Courtyard View, Washer Dryer Connections, 
-- Washer Dryer provided, close to parking, better views, fireplaces, special 
-- remodeling or finishes, etc.  This is where those special characteristics are defined
CREATE TABLE rentablespecialtytypes (
    RSPID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL,
    Name VARCHAR(100) NOT NULL DEFAULT '',
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    Description VARCHAR(256) NOT NULL DEFAULT '',
    PRIMARY KEY (RSPID)
);

-- ===========================================
--   ASSESSMENT TYPES
-- ===========================================
-- this table list all the pre-defined assessments
-- this will include offsets and disbursements
CREATE TABLE assessmenttypes (
    ASMTID BIGINT NOT NULL AUTO_INCREMENT,          -- what type of assessment
    OccupancyRqd SMALLINT NOT NULL DEFAULT 0,       -- 0 = Valid anytime, 1 = valid only during occupancy
    Name VARCHAR(100) NOT NULL DEFAULT '',           -- name for the assessment
    Description VARCHAR(1024) NOT NULL DEFAULT '',   -- describe the assessment
    -- TODO: Type needs to be removed
    -- Type SMALLINT NOT NULL DEFAULT 0,            -- normal case, positive number is: 0 = DEBIT, 1 = CREDIT
    LastModTime TIMESTAMP,                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,         -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMTID)
);

-- ===========================================
--   PAYMENT TYPES
-- ===========================================
CREATE TABLE paymenttypes (
    PMTID MEDIUMINT NOT NULL AUTO_INCREMENT,
    BID MEDIUMINT NOT NULL DEFAULT 0,
    Name VARCHAR(100) NOT NULL DEFAULT '',
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
    AVAILID BIGINT NOT NULL AUTO_INCREMENT,
    Name VARCHAR(100) NOT NULL DEFAULT '',
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
    BID BIGINT NOT NULL DEFAULT 0,
    ASMTID BIGINT NOT NULL DEFAULT 0
);

-- applicable assessments for a specific business
CREATE TABLE businesspaymenttypes (
    BID BIGINT NOT NULL DEFAULT 0,
    PMTID MEDIUMINT NOT NULL DEFAULT 0
);

-- **************************************
-- ****                              ****
-- ****          BUILDING            ****
-- ****                              ****
-- **************************************
CREATE TABLE building (
    BLDGID BIGINT NOT NULL AUTO_INCREMENT,        -- unique id for this building
    BID BIGINT NOT NULL DEFAULT 0,                -- which business it belongs to
    Address VARCHAR(100) NOT NULL DEFAULT '',      -- building address
    Address2 VARCHAR(100) NOT NULL DEFAULT '',       
    City VARCHAR(100) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(100) NOT NULL DEFAULT '',
    Country VARCHAR(100) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                        -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,       -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (BLDGID)
);

-- **************************************
-- ****                              ****
-- ****          RENTABLE            ****
-- ****                              ****
-- **************************************
CREATE TABLE rentable (
    RID BIGINT NOT NULL AUTO_INCREMENT,                            -- unique identifier for this rentable
    RTID BIGINT NOT NULL DEFAULT 0,                                -- what sort of a rentable is this?
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business associated with this rentable
    Name VARCHAR(100) NOT NULL DEFAULT '',                          -- must be unique, name for this instance, "101" for a room number, CP744 carport number, etc 
    Assignment SMALLINT NOT NULL DEFAULT 0,                        -- Unknown = 0, Pre-assign = 1, assign at occupy commencement = 2
    Report SMALLINT NOT NULL DEFAULT 1,                            -- 1 = apply to rentroll, 0 = skip on rentroll
    DefaultOccType SMALLINT NOT NULL DEFAULT 0,                    -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
    OccType SMALLINT NOT NULL DEFAULT 0,                           -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
    State SMALLINT NOT NULL DEFAULT 0,                             -- 0 = online, 1 = administrative unit, 2 = owner occupied, 3 = offline
    -- ManageToBudget SMALLINT NOT NULL DEFAULT 0,                 -- 0 = do not manage to budget, 1 = manage to MarketRate set in RentableType
    LastModTime TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RID)
);


-- **************************************
-- ****                              ****
-- ****            UNIT              ****
-- ****                              ****
-- **************************************
-- Fields unique to an apartment or hotel room 
CREATE TABLE unit (
    UNITID BIGINT NOT NULL AUTO_INCREMENT,              -- unique id for this unit -- it is unique across all properties and buildings
    RID BIGINT NOT NULL DEFAULT 0,                      -- associated rentable
    BLDGID BIGINT NOT NULL DEFAULT 0,                   -- which building
    RTID BIGINT NOT NULL DEFAULT 0,                     -- which rentable type
    AVAILID BIGINT NOT NULL DEFAULT 0,                  -- how is the unit made available
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (UNITID)
    -- Abbreviation VARCHAR(100),                       -- unit abbreviation  -- REMOVED - it's part of unittype
);

-- ===========================================
--   RENTABLE SPECIALTIES
-- ===========================================
-- For each unit, what specialties does it have...
-- this is simply a list of RSPIDs.
-- will be the list of all the unit specialties for that unit.
CREATE TABLE rentablespecialties (
    BID BIGINT NOT NULL DEFAULT 0,                      -- the business
    RID BIGINT NOT NULL DEFAULT 0,                      -- unique id of unit
    RSPID BIGINT NOT NULL DEFAULT 0                     -- unique id of specialty (see Table rentablespecialties)
);


-- **************************************
-- ****                              ****
-- ****        ASSESSMENTS           ****
-- ****                              ****
-- **************************************
-- charges associated with a rentable
CREATE TABLE assessments (
    ASMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business id
    RID BIGINT NOT NULL DEFAULT 0,                          -- rental id
    ASMTID BIGINT NOT NULL DEFAULT 0,                       -- what type of assessment (ex: Rent, SecurityDeposit, ...)
    RAID BIGINT NOT NULL DEFAULT 0,                         -- Associated Rental Agreement ID
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- Assessment amount
    Start DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- epoch date for the assessment - recurrences are based on this date
    Stop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00',   -- stop date - when the tenant moves out or when the charge is no longer applicable
    Accrual SMALLINT NOT NULL DEFAULT 0,                  -- 0 = one time only, 1 = daily, 2 = weekly, 3 = monthly,   4 = yearly
    ProrationMethod SMALLINT NOT NULL DEFAULT 0,            -- 
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',              -- Accounting rule - which acct debited, which credited
    Comment VARCHAR(256) NOT NULL DEFAULT '',               -- for comments such as "Prior period adjustment"
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
    TCID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique id of unit
    TID BIGINT NOT NULL DEFAULT 0,                         -- associated tenant id
    PID BIGINT NOT NULL DEFAULT 0,                         -- associated payor id
    PRSPID BIGINT NOT NULL DEFAULT 0,                      -- associated prospect id
    FirstName VARCHAR(100) NOT NULL DEFAULT '',
    MiddleName VARCHAR(100) NOT NULL DEFAULT '',
    LastName VARCHAR(100) NOT NULL DEFAULT '',
    CompanyName VARCHAR(100) NOT NULL DEFAULT '',
    IsCompany SMALLINT NOT NULL DEFAULT 0,                  -- 0 == this is a person,  1 == this is a company 
    PrimaryEmail VARCHAR(100) NOT NULL DEFAULT '',
    SecondaryEmail VARCHAR(100) NOT NULL DEFAULT '',
    WorkPhone VARCHAR(100) NOT NULL DEFAULT '',
    CellPhone VARCHAR(100) NOT NULL DEFAULT '',
    Address VARCHAR(100) NOT NULL DEFAULT '',            -- person address
    Address2 VARCHAR(100) NOT NULL DEFAULT '',       
    City VARCHAR(100) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(100) NOT NULL DEFAULT '',
    Country VARCHAR(100) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

-- ===========================================
--   PROSPECT
-- ===========================================
CREATE TABLE prospect (
    PRSPID BIGINT NOT NULL AUTO_INCREMENT,                 -- unique id of this prospect
    TCID BIGINT NOT NULL DEFAULT 0,                        -- associated transactant (has Name and all contact info)
    ApplicationFee DECIMAL(19,4) NOT NULL DEFAULT 0.0,  -- if non-zero this prospect is an applicant
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PRSPID)
);

-- ===========================================
--   TENANT
-- ===========================================
CREATE TABLE tenant (
    TID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id of this tenant
    TCID BIGINT NOT NULL,                                  -- associated transactant
    Points BIGINT NOT NULL DEFAULT 0,                      -- bonus points for this tenant
    CarMake VARCHAR(100) NOT NULL DEFAULT '',
    CarModel VARCHAR(100) NOT NULL DEFAULT '',
    CarColor VARCHAR(100) NOT NULL DEFAULT '',
    CarYear BIGINT NOT NULL DEFAULT 0,
    LicensePlateState VARCHAR(100) NOT NULL DEFAULT '',
    LicensePlateNumber VARCHAR(100) NOT NULL DEFAULT '',
    ParkingPermitNumber VARCHAR(100) NOT NULL DEFAULT '',
    AccountRep BIGINT NOT NULL DEFAULT 0,                              -- Phonebook UID of account rep
    DateofBirth DATE NOT NULL DEFAULT '1970-01-01T00:00:00',
    EmergencyContactName VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactAddress VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactTelephone VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyEmail VARCHAR(100) NOT NULL DEFAULT '',
    AlternateAddress VARCHAR(100) NOT NULL DEFAULT '',
    ElibigleForFutureOccupancy SMALLINT NOT NULL DEFAULT 1,         -- yes/no
    Industry VARCHAR(100) NOT NULL DEFAULT '',                       -- (e.g., construction, retail, banking etc.)
    Source  VARCHAR(100) NOT NULL DEFAULT '',                        -- (e.g., resident referral, newspaper, radio, post card, expedia, travelocity, etc.)
    InvoicingCustomerNumber VARCHAR(100) NOT NULL DEFAULT '',        -- [drawn from the invoicing section] [only applies if invoicing authorization has been provide
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TID)
);

-- ===========================================
--   PAYOR
-- ===========================================
CREATE TABLE payor  (
    PID BIGINT NOT NULL AUTO_INCREMENT,                          -- unique id of this payor
    TCID BIGINT NOT NULL,                                        -- associated transactant
    CreditLimit DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    EmployerName  VARCHAR(100) NOT NULL DEFAULT '',
    EmployerStreetAddress VARCHAR(100) NOT NULL DEFAULT '',
    EmployerCity VARCHAR(100) NOT NULL DEFAULT '',
    EmployerState VARCHAR(100) NOT NULL DEFAULT '',
    EmployerPostalCode VARCHAR(100) NOT NULL DEFAULT '',
    EmployerEmail VARCHAR(100) NOT NULL DEFAULT '',
    EmployerPhone VARCHAR(100) NOT NULL DEFAULT '',
    Occupation VARCHAR(100) NOT NULL DEFAULT '',
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
    RCPTID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for this receipt
    BID BIGINT NOT NULL DEFAULT 0,
    RAID BIGINT NOT NULL DEFAULT 0,
    PMTID BIGINT NOT NULL DEFAULT 0,
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AcctRule VARCHAR(1500) NOT NULL DEFAULT '',
    Comment VARCHAR(256) NOT NULL DEFAULT '',                   -- for comments like "Prior Period Adjustment"
    -- ApplyToGeneralReceivable DECIMAL(19,4),                  -- Breakdown is in receiptallocation table
    -- ApplyToSecurityDeposit DECIMAL(19,4),                    -- Can we just handle this as part of receipt allocation
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RCPTID)
);

CREATE TABLE receiptallocation (
    RCPTID BIGINT NOT NULL DEFAULT 0,                              -- sum of all amounts in this table with RCPTID must equal the receipt with RCPTID in receipt table
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID BIGINT NOT NULL DEFAULT 0,                               -- the id of the assessment that caused this payment
    AcctRule VARCHAR(150)
);  

-- **************************************
-- ****                              ****
-- ****           JOURNAL            ****
-- ****                              ****
-- **************************************
CREATE TABLE journal (
    JID BIGINT NOT NULL AUTO_INCREMENT,                            -- a journal entry
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                                -- associated rental agreement
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',            -- date when it occurred
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                     -- how much
    Type SMALLINT NOT NULL DEFAULT 0,                              -- 0 = unknown, 1 = assessment, 2 = payment/receipt
    ID BIGINT NOT NULL DEFAULT 0,                                  -- if Type == 1 then it is the ASMID that caused this entry, if Type ==2 then it is the RCPTID
    -- no last mod by, etc., this is all handled in the journalaudit table
    Comment VARCHAR(256) NOT NULL DEFAULT '',                 -- for notes like "prior period adjustment"
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (JID)
);

CREATE TABLE journalallocation (
    JAID BIGINT NOT NULL AUTO_INCREMENT,
    JID BIGINT NOT NULL DEFAULT 0,                                 -- sum of all amounts in this table with RCPTID must equal the receipt with RCPTID in receipt table
    RID BIGINT NOT NULL DEFAULT 0,                                 -- associated rentable
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID BIGINT NOT NULL DEFAULT 0,                               -- may not be present if assessment records have been backed up and removed.
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',
    PRIMARY KEY (JAID)
);  

CREATE TABLE journalmarker (
    JMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business id
    State SMALLINT NOT NULL DEFAULT 0,                             -- 0 = unknown, 1 = Closed, 2 = Locked
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (JMID)
);

CREATE TABLE journalaudit (
    JID BIGINT NOT NULL DEFAULT 0,          -- what JID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,       -- UID of person making the change
    ModTime TIMESTAMP                       -- timestamp of change    
);

CREATE TABLE journalmarkeraudit (
    JMID BIGINT NOT NULL DEFAULT 0,         -- what JMID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,       -- UID of person making the change
    ModTime TIMESTAMP                       -- timestamp of change
);

-- **************************************
-- ****                              ****
-- ****           LEDGERS            ****
-- ****                              ****
-- **************************************
CREATE TABLE ledgerentry (
    LID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for this Ledger
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    JID BIGINT NOT NULL DEFAULT 0,                            -- journal entry giving rise to this
    JAID BIGINT NOT NULL DEFAULT 0,                           -- the allocation giving rise to this ledger entry
    GLNumber VARCHAR(100) NOT NULL DEFAULT '',                 -- if not '' then it's a link a QB  GeneralLedger (GL)account
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- balance date and time
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                -- balance amount since last close
    Comment VARCHAR(256) NOT NULL DEFAULT '',                 -- for notes like "prior period adjustment"
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LID)        
);

CREATE TABLE ledgermarker (
    LMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                            -- payor id, only valid if TYPE is
    GLNumber VARCHAR(100) NOT NULL DEFAULT '',                 -- if not '' then it's a link a QB  GeneralLedger (GL)account
    Status SMALLINT NOT NULL DEFAULT 0,                       -- Whether a GL Account is currently unknown=0, inactive=1, active=2 
    State SMALLINT NOT NULL DEFAULT 0,                        -- 0 = unknown, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- period start
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- period end
    Balance DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    Type SMALLINT NOT NULL DEFAULT 0,                         -- flag: 0 = not a default account, 1 = Payor Account , 
    --                                                                 10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC, 15 sec dep receivable, 16 sec dep assessment
    Name VARCHAR(100) NOT NULL DEFAULT '',
    AcctType VARCHAR(100) NOT NULL DEFAULT '',                -- Quickbooks Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, 
                                                              --    Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, 
                                                              --    Cost of Goods Sold, Other Income, Other Expense
    RAAssociated SMALLINT NOT NULL DEFAULT 0,                 -- 1 = Unassociated with RentalAgreement, 2 = Associated with Rental Agreement, 0 = unknown
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LMID)
);

CREATE TABLE ledgeraudit (
    LID BIGINT NOT NULL DEFAULT 0,              -- what LID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,           -- UID of person making the change
    ModTime TIMESTAMP                           -- timestamp of change    
);

CREATE TABLE ledgermarkeraudit (
    LMID BIGINT NOT NULL DEFAULT 0,             -- what LMID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,           -- UID of person making the change
    ModTime TIMESTAMP                           -- timestamp of change    
);


-- **************************************
-- ****                              ****
-- ****        INITIALIZATION        ****
-- ****                              ****
-- **************************************
-- ----------------------------------------------------------------------------------------
--    LEDGER MARKERS - These define the required ledgers
-- ----------------------------------------------------------------------------------------
INSERT INTO ledgermarker (BID,RAID,GLNumber,Status,State,DtStart,DtStop,Balance,Type,Name) VALUES
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,10,"Bank Account"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,11,"General Accounts Receivable"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,12,"Gross Scheduled Rent"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,13,"Loss to Lease"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,14,"Vacancy"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,15,"Security Deposit Receivable"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,16,"Security Deposit Assessment"),
    (1,0,"",2,3,"2015-10-01","2015-10-31",0.0,17,"Owner Equity");
