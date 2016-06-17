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
-- BID = Business id
-- BLDGID = Building id
-- CID = custom attribute id
-- DISBID = disbursement id
-- JAID = Journal allocation id
-- JID = Journal id
-- JMID = Journal marker id
-- LEID = Ledger id
-- LMID = Ledger marker id
-- OFSID = offset id
-- PID = Payor id
-- PMTID = payment type id
-- PRSPID = Prospect id
-- RAID = rental agreement / occupancy agreement
-- RATID = rental agreement template id
-- RCPTID = Receipt id
-- RID = Rentable id
-- RSPID = unit specialty id
-- RTID = Rentable type id
-- TCID = Transactant id
-- TCID = Transactant id
-- USERID = User id

DROP DATABASE IF EXISTS rentroll;
CREATE DATABASE rentroll;
USE rentroll;
GRANT ALL PRIVILEGES ON rentroll TO 'ec2-user'@'localhost';
GRANT ALL PRIVILEGES ON rentroll.* TO 'ec2-user'@'localhost';

-- ===========================================
--   RENTAL AGREEMENT TEMPLATE
-- ===========================================
CREATE TABLE RentalAgreementTemplate (
    RATID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal unique id
    BID BIGINT NOT NULL DEFAULT 0,                            -- BizUnit Reference
    RentalTemplateNumber VARCHAR(100) DEFAULT '',             -- Occupancy Agreement Reference Number
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RATID)     
);      
        
-- ===========================================
--   RENTAL AGREEMENT
-- ===========================================
CREATE TABLE RentalAgreement (        
    RAID BIGINT NOT NULL AUTO_INCREMENT,                      -- internal unique id
    RATID BIGINT NOT NULL DEFAULT 0,                          -- reference to Rental Template (Occupancy Master Agreement)
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    RentalStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- date when rental starts
    RentalStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- date when rental stops
    PossessionStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',     -- date when Occupancy starts
    PossessionStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when Occupancy stops
    Renewal SMALLINT NOT NULL DEFAULT 0,                      -- 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
    SpecialProvisions VARCHAR(1024) NOT NULL DEFAULT '',      -- free-form text
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RAID)
);

CREATE TABLE RentalAgreementRentables (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    RID BIGINT NOT NULL DEFAULT 0,                            -- Rentable id
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this Rentable was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this Rentable was no longer being billed to this agreement
);

CREATE TABLE RentalAgreementPayors (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    PID BIGINT NOT NULL DEFAULT 0,                            -- who is the Payor for this agreement
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this Payor was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this Payor was no longer being billed to this agreement
);

CREATE TABLE RentableUsers (
    RID BIGINT NOT NULL DEFAULT 0,                            -- the associated Rentable
    USERID BIGINT NOT NULL DEFAULT 0,                         -- the Users of the rentable
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this User was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this User was no longer being billed to this agreement
);

CREATE TABLE RentalAgreementPets (
    PETID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal id for this pet
    RAID BIGINT NOT NULL DEFAULT 0,                           -- the unit's occupancy agreement
    Type VARCHAR(100) NOT NULL DEFAULT '',                    --  type of animal, ex: dog, cat, ...
    Breed VARCHAR(100) NOT NULL DEFAULT '',                   --  breed.  example Beagle, German Shephard, Siamese, etc.
    Color VARCHAR(100) NOT NULL DEFAULT '',                   --
    Weight DECIMAL(19,4) NOT NULL DEFAULT 0,                  --  in pounds
    Name VARCHAR(100) NOT NULL DEFAULT '',                    --
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this User was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',       -- date when this User was no longer being billed to this agreement
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PETID)
);

-- **************************************
-- ****                              ****
-- ****    USER DEFINED ATTRIBUTES   ****
-- ****                              ****
-- ************************************** 
CREATE TABLE CustomAttr (
    CID BIGINT NOT NULL AUTO_INCREMENT,        -- unique identifer for this custom attribute
    Type SMALLINT NOT NULL DEFAULT 0,          -- 0 = string, 1 = int64, 2 = float64
    Name VARCHAR (100) NOT NULL DEFAULT '',    -- a name
    Value VARCHAR (256) NOT NULL DEFAULT '',   -- its value in string form
    LastModTime TIMESTAMP,                     -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,    -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (CID)
);

CREATE TABLE CustomAttrRef (
    ElementType BIGINT NOT NULL,               -- for what type of object is this a ref:  1=Person, 2=Company, 3=Business-Unit, 4=executable service, 5=RentableType
    ID          BIGINT NOT NULL,               -- the UID of the object type. That is, if ObjectType == 5, the ID is the RTID (Rentable type id)
    CID         BIGINT NOT NULL                -- uid of the custom attribute
);

-- **************************************
-- ****                              ****
-- ****          BUSINESS            ****
-- ****                              ****
-- **************************************
-- Unit Types - associated with a Business are stored
-- in the RentableTypes table and have the Business PID
-- Occupancy Type List - hardcoded

CREATE TABLE Business (
    BID BIGINT NOT NULL AUTO_INCREMENT,
    BUD VARCHAR(100) NOT NULL DEFAULT '',               -- Business Unit Designation
    Name VARCHAR(100) NOT NULL DEFAULT '',
    DefaultRentalPeriod SMALLINT NOT NULL DEFAULT 0,    -- default for every unit in the Building: 0=unset, 1=hourly, 2=daily, 3=weekly, 4=monthly, 5=quarterly, 6=yearly
    ParkingPermitInUse SMALLINT NOT NULL DEFAULT 0,     -- yes/no  0 = no, 1 = yes
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (BID)
);



-- ===========================================
--   RENTABLE TYPES 
-- ===========================================
CREATE TABLE RentableTypes (
    RTID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                          -- associated Business id
    Style CHAR(15) NOT NULL DEFAULT '',                     -- need not be unique
    Name VARCHAR(256) NOT NULL DEFAULT '',                  -- must be unique
    RentCycle BIGINT NOT NULL DEFAULT 0,                    -- rent accrual frequency
    Proration BIGINT NOT NULL DEFAULT 0,                    -- prorate frequency
    GSPRC BIGINT NOT NULL DEFAULT 0,                        -- Increments in which GSR is calculated to account for rate changes
    ManageToBudget SMALLINT NOT NULL DEFAULT 0,             -- 0 do not manage this category of Rentable to budget, 1 = manage to budget defined by MarketRate
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RTID)
);

CREATE TABLE RentableMarketrate (
    RTID BIGINT NOT NULL DEFAULT 0,                             -- associated Rentable type
    MarketRate DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- market rate for the time range
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',    
    DtStop DATETIME NOT NULL DEFAULT '9999-12-31 23:59:59'      -- assume it's unbounded. if an updated Market rate is added, set this to the stop date
);


-- ===========================================
--   RENTABLE SPECIALTY TYPES
-- ===========================================
-- a collection of unit specialties that are available.
-- different units may be more or less desirable based upon special characteristics
-- of the unit, such as Lake View, Courtyard View, Washer Dryer Connections, 
-- Washer Dryer provided, close to parking, better views, fireplaces, special 
-- remodeling or finishes, etc.  This is where those special characteristics are defined
CREATE TABLE RentableSpecialtyType (
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
-- this table list all the pre-defined Assessments
-- this will include offsets and disbursements
CREATE TABLE AssessmentTypes (
    ASMTID BIGINT NOT NULL AUTO_INCREMENT,                      -- what type of assessment
    RARequired SMALLINT NOT NULL DEFAULT 0,                     -- 0 = Valid anytime, 1 = valid only during occupancy
    Name VARCHAR(100) NOT NULL DEFAULT '',                      -- name for the assessment
    Description VARCHAR(1024) NOT NULL DEFAULT '',              -- describe the assessment
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMTID)    
);

-- ===========================================
--   PAYMENT TYPES
-- ===========================================
CREATE TABLE PaymentTypes (
    PMTID MEDIUMINT NOT NULL AUTO_INCREMENT,
    BID MEDIUMINT NOT NULL DEFAULT 0,
    Name VARCHAR(100) NOT NULL DEFAULT '',
    Description VARCHAR(256) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PMTID)
);

-- ===========================================
--   AVAILABILITY TYPES
-- ===========================================
-- Examples: Occupied, Offline, Administrative, Vacant - Not Ready, 
--           Vacant - Made Ready, Vacant - Inspected, plus custom values
-- Custom values will be added with their own uniq AVAILID
CREATE TABLE AvailabilityTypes (
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
-- This describes the world of Assessments for a particular Business
-- Query this table for a particular BID, the solution set is the list
-- of Assessments for that particular Business.
-- applicable Assessments for a specific Business
CREATE TABLE BusinessAssessments (
    BID BIGINT NOT NULL DEFAULT 0,
    ASMTID BIGINT NOT NULL DEFAULT 0
);

-- applicable Assessments for a specific Business
CREATE TABLE BusinessPaymentTypes (
    BID BIGINT NOT NULL DEFAULT 0,
    PMTID MEDIUMINT NOT NULL DEFAULT 0
);

-- **************************************
-- ****                              ****
-- ****          BUILDING            ****
-- ****                              ****
-- **************************************
CREATE TABLE Building (
    BLDGID BIGINT NOT NULL AUTO_INCREMENT,                          -- unique id for this Building
    BID BIGINT NOT NULL DEFAULT 0,                                  -- which Business it belongs to
    Address VARCHAR(100) NOT NULL DEFAULT '',                       -- Building address
    Address2 VARCHAR(100) NOT NULL DEFAULT '',                         
    City VARCHAR(100) NOT NULL DEFAULT '',                  
    State CHAR(25) NOT NULL DEFAULT '',                 
    PostalCode VARCHAR(100) NOT NULL DEFAULT '',                    
    Country VARCHAR(100) NOT NULL DEFAULT '',                   
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (BLDGID)
);

-- **************************************
-- ****                              ****
-- ****          RENTABLE            ****
-- ****                              ****
-- **************************************
CREATE TABLE Rentable (
    RID BIGINT NOT NULL AUTO_INCREMENT,                            -- unique identifier for this Rentable
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business associated with this Rentable
    Name VARCHAR(100) NOT NULL DEFAULT '',                         -- must be unique, name for this instance, "101" for a room number, CP744 carport number, etc 
    AssignmentTime SMALLINT NOT NULL DEFAULT 0,                    -- Unknown = 0, Pre-assign = 1, assign at occupy commencement = 2
    LastModTime TIMESTAMP,                                         -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RID)
    -- RentalPeriodDefault SMALLINT NOT NULL DEFAULT 0,               -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
    -- RentCycle SMALLINT NOT NULL DEFAULT 0,                         -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
);

CREATE TABLE RentableTypeRef (
    RID BIGINT NOT NULL DEFAULT 0,                                  -- the Rentable this record belongs to
    RTID BIGINT NOT NULL DEFAULT 0,                                 -- the Rentable type for this period
    RentCycle BIGINT NOT NULL DEFAULT 0,                            -- RentCycle override. 0 = unset, > 0 means the frequency
    ProrationCycle BIGINT NOT NULL DEFAULT 0,                      -- Proration override. 0 = unset, > 0 means the override proration
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- start time for this state
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- stop time for this state
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that modified it 
);

CREATE TABLE RentableStatus (
    RID BIGINT NOT NULL DEFAULT 0,                                  -- associated Rentable
    Status SMALLINT NOT NULL DEFAULT 0,                             -- 0 = UNKNOWN -- 1 = ONLINE, 2 = ADMIN, 3 = EMPLOYEE, 4 = OWNEROCC, 5 = OFFLINE
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- start time for this state
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- stop time for this state
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that modified it 
);

CREATE TABLE RentableSpecialtyRef (
    BID BIGINT NOT NULL DEFAULT 0,                                  -- the Business
    RID BIGINT NOT NULL DEFAULT 0,                                  -- unique id of unit
    RSPID BIGINT NOT NULL DEFAULT 0,                                -- unique id of specialty (see Table RentableSpecialties)
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- start time for this state
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- stop time for this state
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that modified it 
);


-- **************************************
-- ****                              ****
-- ****        ASSESSMENTS           ****
-- ****                              ****
-- **************************************
-- charges associated with a Rentable
CREATE TABLE Assessments (
    ASMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business id
    RID BIGINT NOT NULL DEFAULT 0,                          -- rental id
    ASMTID BIGINT NOT NULL DEFAULT 0,                       -- what type of assessment (ex: Rent, SecurityDeposit, ...)
    RAID BIGINT NOT NULL DEFAULT 0,                         -- Associated Rental Agreement ID
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- Assessment amount
    Start DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- epoch date for the assessment - recurrences are based on this date
    Stop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00',   -- stop date - when the User moves out or when the charge is no longer applicable
    RentCycle SMALLINT NOT NULL DEFAULT 0,                  -- 0 = one time only, 1 = daily, 2 = weekly, 3 = monthly,   4 = yearly
    ProrationCycle SMALLINT NOT NULL DEFAULT 0,            -- 
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
-- Transactant - fields common to all people and
CREATE TABLE Transactant (
    TCID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique id of unit
    USERID BIGINT NOT NULL DEFAULT 0,                         -- associated User id
    PID BIGINT NOT NULL DEFAULT 0,                         -- associated Payor id
    PRSPID BIGINT NOT NULL DEFAULT 0,                      -- associated Prospect id
    FirstName VARCHAR(100) NOT NULL DEFAULT '',
    MiddleName VARCHAR(100) NOT NULL DEFAULT '',
    LastName VARCHAR(100) NOT NULL DEFAULT '',
    PreferredName VARCHAR(100) NOT NULL DEFAULT '',
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
    Website VARCHAR(100) NOT NULL DEFAULT '',
    Notes VARCHAR(256) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

-- ===========================================
--   PROSPECT
-- ===========================================
CREATE TABLE Prospect (
    PRSPID BIGINT NOT NULL AUTO_INCREMENT,                 -- unique id of this Prospect
    TCID BIGINT NOT NULL DEFAULT 0,                        -- associated Transactant (has Name and all contact info)
    EmployerName  VARCHAR(100) NOT NULL DEFAULT '',
    EmployerStreetAddress VARCHAR(100) NOT NULL DEFAULT '',
    EmployerCity VARCHAR(100) NOT NULL DEFAULT '',
    EmployerState VARCHAR(100) NOT NULL DEFAULT '',
    EmployerPostalCode VARCHAR(100) NOT NULL DEFAULT '',
    EmployerEmail VARCHAR(100) NOT NULL DEFAULT '',
    EmployerPhone VARCHAR(100) NOT NULL DEFAULT '',
    Occupation VARCHAR(100) NOT NULL DEFAULT '',
    ApplicationFee DECIMAL(19,4) NOT NULL DEFAULT 0.0,  -- if non-zero this Prospect is an applicant
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PRSPID)
);

-- ===========================================
--   TENANT
-- ===========================================
CREATE TABLE User (
    USERID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id of this User
    TCID BIGINT NOT NULL,                                  -- associated Transactant
    Points BIGINT NOT NULL DEFAULT 0,                      -- bonus points for this User
    CarMake VARCHAR(100) NOT NULL DEFAULT '',
    CarModel VARCHAR(100) NOT NULL DEFAULT '',
    CarColor VARCHAR(100) NOT NULL DEFAULT '',
    CarYear BIGINT NOT NULL DEFAULT 0,
    LicensePlateState VARCHAR(100) NOT NULL DEFAULT '',
    LicensePlateNumber VARCHAR(100) NOT NULL DEFAULT '',
    ParkingPermitNumber VARCHAR(100) NOT NULL DEFAULT '',
    DateofBirth DATE NOT NULL DEFAULT '1970-01-01T00:00:00',
    EmergencyContactName VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactAddress VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactTelephone VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyEmail VARCHAR(100) NOT NULL DEFAULT '',
    AlternateAddress VARCHAR(100) NOT NULL DEFAULT '',
    EligibleFutureUser SMALLINT NOT NULL DEFAULT 1,               -- yes/no
    Industry VARCHAR(100) NOT NULL DEFAULT '',                      -- (e.g., construction, retail, banking etc.)
    Source  VARCHAR(100) NOT NULL DEFAULT '',                       -- (e.g., resident referral, newspaper, radio, post card, expedia, travelocity, etc.)
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (USERID)
);

-- ===========================================
--   PAYOR
-- ===========================================
CREATE TABLE Payor  (
    PID BIGINT NOT NULL AUTO_INCREMENT,                         -- unique id of this Payor
    TCID BIGINT NOT NULL,                                       -- associated Transactant
    TaxpayorID VARCHAR(25) NOT NULL DEFAULT '',
    CreditLimit DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AccountRep BIGINT NOT NULL DEFAULT 0,                       -- Phonebook UID of account rep
    EligibleFuturePayor SMALLINT NOT NULL DEFAULT 1,            -- yes/no
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PID)
);

-- **************************************
-- ****                              ****
-- ****           RECEIPTS           ****
-- ****                              ****
-- **************************************
CREATE TABLE Receipt (
    RCPTID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for this Receipt
    BID BIGINT NOT NULL DEFAULT 0,
    RAID BIGINT NOT NULL DEFAULT 0,  -- THIS IS AN ISSUE... It can go away -- ReceiptAllocation has an associated Assessment, which has the RAID
    PMTID BIGINT NOT NULL DEFAULT 0,
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AcctRule VARCHAR(1500) NOT NULL DEFAULT '',
    Comment VARCHAR(256) NOT NULL DEFAULT '',                   -- for comments like "Prior Period Adjustment"
    -- ApplyToGeneralReceivable DECIMAL(19,4),                  -- Breakdown is in ReceiptAllocation table
    -- ApplyToSecurityDeposit DECIMAL(19,4),                    -- Can we just handle this as part of Receipt allocation
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RCPTID)
);

CREATE TABLE ReceiptAllocation (
    RCPTID BIGINT NOT NULL DEFAULT 0,                              -- sum of all amounts in this table with RCPTID must equal the Receipt with RCPTID in Receipt table
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID BIGINT NOT NULL DEFAULT 0,                               -- the id of the assessment that caused this payment
    AcctRule VARCHAR(150)
);  

-- **************************************
-- ****                              ****
-- ****           JOURNAL            ****
-- ****                              ****
-- **************************************
CREATE TABLE Journal (
    JID BIGINT NOT NULL AUTO_INCREMENT,                            -- a Journal entry
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                                -- associated rental agreement
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',            -- date when it occurred
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                     -- how much
    Type SMALLINT NOT NULL DEFAULT 0,                              -- 0 = unknown, 1 = assessment, 2 = payment/Receipt
    ID BIGINT NOT NULL DEFAULT 0,                                  -- if Type == 1 then it is the ASMID that caused this entry, if Type ==2 then it is the RCPTID
    -- no last mod by, etc., this is all handled in the JournalAudit table
    Comment VARCHAR(256) NOT NULL DEFAULT '',                 -- for notes like "prior period adjustment"
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (JID)
);

CREATE TABLE JournalAllocation (
    JAID BIGINT NOT NULL AUTO_INCREMENT,
    JID BIGINT NOT NULL DEFAULT 0,                                 -- sum of all amounts in this table with RCPTID must equal the Receipt with RCPTID in Receipt table
    RID BIGINT NOT NULL DEFAULT 0,                                 -- associated Rentable
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID BIGINT NOT NULL DEFAULT 0,                               -- may not be present if assessment records have been backed up and removed.
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',
    PRIMARY KEY (JAID)
);  

CREATE TABLE JournalMarker (
    JMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business id
    State SMALLINT NOT NULL DEFAULT 0,                             -- 0 = unknown, 1 = Closed, 2 = Locked
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (JMID)
);

CREATE TABLE JournalAudit (
    JID BIGINT NOT NULL DEFAULT 0,          -- what JID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,       -- UID of person making the change
    ModTime TIMESTAMP                       -- timestamp of change    
);

CREATE TABLE JournalMarkerAudit (
    JMID BIGINT NOT NULL DEFAULT 0,         -- what JMID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,       -- UID of person making the change
    ModTime TIMESTAMP                       -- timestamp of change
);

-- **************************************
-- ****                              ****
-- ****           LEDGERS            ****
-- ****                              ****
-- **************************************
-- RENAME to Ledger
CREATE TABLE LedgerEntry (
    LEID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this Ledger
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    JID BIGINT NOT NULL DEFAULT 0,                            -- Journal entry giving rise to this
    JAID BIGINT NOT NULL DEFAULT 0,                           -- the allocation giving rise to this Ledger entry
    LID BIGINT NOT NULL DEFAULT 0,                            -- associated Ledger
    RAID BIGINT NOT NULL DEFAULT 0,                           -- associated Rental Agreement
    -- GLNo VARCHAR(100) NOT NULL DEFAULT '',                    -- if not '' then it's a link a QB  GeneralLedger (GL)account
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- balance date and time
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                -- balance amount since last close
    Comment VARCHAR(256) NOT NULL DEFAULT '',                 -- for notes like "prior period adjustment"
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LEID)        
);

CREATE TABLE LedgerMarker (
    LMID BIGINT NOT NULL AUTO_INCREMENT,
    LID BIGINT NOT NULL DEFAULT 0,                            -- associated Ledger
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- period start
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- period end
    Balance DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    State SMALLINT NOT NULL DEFAULT 0,                        -- 0 = Open, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LMID)
);

-- GL Account
CREATE TABLE Ledger (
    LID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for this Ledger
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                           -- rental agreement account, only valid if TYPE is 1
    GLNumber VARCHAR(100) NOT NULL DEFAULT '',                -- if not '' then it's a link a QB  GeneralLedger (GL)account
    Status SMALLINT NOT NULL DEFAULT 0,                       -- Whether a GL Account is currently unknown=0, inactive=1, active=2 
    Type SMALLINT NOT NULL DEFAULT 0,                         -- flag: 0 = not a special account of any kind, 1 = RentalAgreement Balance, 
    --                                                                 10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC, 15 sec dep receivable, 16 sec dep assessment
    Name VARCHAR(100) NOT NULL DEFAULT '',
    AcctType VARCHAR(100) NOT NULL DEFAULT '',                -- Quickbooks Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, 
                                                              --    Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, 
                                                              --    Cost of Goods Sold, Other Income, Other Expense
    RAAssociated SMALLINT NOT NULL DEFAULT 0,                 -- 1 = Unassociated with RentalAgreement, 2 = Associated with Rental Agreement, 0 = unknown
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LID)
);

CREATE TABLE LedgerAudit (
    LEID BIGINT NOT NULL DEFAULT 0,             -- what LEID was affected
    UID MEDIUMINT NOT NULL DEFAULT 0,           -- UID of person making the change
    ModTime TIMESTAMP                           -- timestamp of change    
);

CREATE TABLE LedgerMarkerAudit (
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
--    LEDGERs  - These define the required ledgers
-- ----------------------------------------------------------------------------------------
INSERT INTO Ledger (BID,RAID,GLNumber,Status,Type,Name) VALUES
    (1,0,"",2,10,"Bank Account"),                   -- 1
    (1,0,"",2,11,"General Accounts Receivable"),    -- 2
    (1,0,"",2,12,"Gross Scheduled Rent"),           -- 3
    (1,0,"",2,13,"Loss to Lease"),                  -- 4
    (1,0,"",2,14,"Vacancy"),                        -- 5
    (1,0,"",2,15,"Security Deposit Receivable"),    -- 6
    (1,0,"",2,16,"Security Deposit Assessment"),    -- 7
    (1,0,"",2,17,"Owner Equity");                   -- 8
-- ----------------------------------------------------------------------------------------
--    LEDGERs MARKERS - These define the required ledgers
-- ----------------------------------------------------------------------------------------
INSERT INTO LedgerMarker (BID,LID,State,DtStart,DtStop,Balance) VALUES
    (1,1,3,"2015-10-01","2015-10-31",0.0),
    (1,2,3,"2015-10-01","2015-10-31",0.0),
    (1,3,3,"2015-10-01","2015-10-31",0.0),
    (1,4,3,"2015-10-01","2015-10-31",0.0),
    (1,5,3,"2015-10-01","2015-10-31",0.0),
    (1,6,3,"2015-10-01","2015-10-31",0.0),
    (1,7,3,"2015-10-01","2015-10-31",0.0),
    (1,8,3,"2015-10-01","2015-10-31",0.0);
