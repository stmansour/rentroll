-- Conventions used:
--     Table names are all lower case
--     Field names are camel case
--     Money values are all stored as DECIMAL(19,4)

-- ********************************
-- *********  UNIQUE IDS  *********
-- ********************************
-- ASMID = Assessment id
-- ATypeLID = assessment type id
-- AVAILID = availability id
-- BID = Business id
-- BLDGID = Building id
-- CID = custom attribute id
-- DISBID = disbursement id
-- JAID = Journal allocation id
-- JID = Journal id
-- JMID = Journal marker id
-- LEID = LedgerEntry id
-- LMID = LedgerMarker id
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
--   ID COUNTERS
--  may not need this
-- ===========================================
CREATE TABLE IDCounters (
    InvoiceNo BIGINT NOT NULL DEFAULT 0                     -- unique number for invoices
);

-- ===========================================
--   TAXES
-- ===========================================
CREATE TABLE Tax (
    TAXID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique identifier for this tax
    BID BIGINT NOT NULL DEFAULT 0,                          -- what business is this tax associated with
    Name VARCHAR(50),                                       -- a name for this tax
    TaxingAuthority VARCHAR(100),                           -- name of the Taxing Authority
    TaxingAuthorityAddress VARCHAR(256),                    -- where these taxes are sent
    FilingDate DATE NOT NULL DEFAULT '1970-01-01',          -- date on which taxes need to be filed
    FilingCycle BIGINT NOT NULL DEFAULT 0,                  -- epoch date for recurrence calculation
    Instructions VARCHAR(1024) NOT NULL DEFAULT '',         -- filing instructions
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY(TAXID)
);

CREATE TABLE TaxRate (
    TAXID BIGINT NOT NULL DEFAULT 0,                        -- reference to which tax this table represents
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',    -- date when this tax rate goes into effect
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',     -- date when this tax rate is no longer applicable
    Rate DECIMAL(19,4) NOT NULL DEFAULT 0,                  -- floating point number representing the rate. Set to 0 if not applicable.
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0,                   -- floating point number.  Set to 0 if not applicable.
    Formula VARCHAR(256) NOT NULL DEFAULT '',               -- RPN calculator notation of formula, Set to '' if not needed
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0                  -- employee UID (from phonebook) that modified it 
);



CREATE TABLE StringList (
    SLID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id for this stringlist
    BID BIGINT NOT NULL DEFAULT 0,                          -- the business to which this stringlist belongs
    Name VARCHAR(50) NOT NULL DEFAULT '',                   -- stringlist name
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY(SLID)
);

CREATE TABLE SLString (
    SLSID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique id for this string
    SLID BIGINT NOT NULL DEFAULT 0,                         -- to which stringlist does this string belong?
    Value VARCHAR(256) NOT NULL DEFAULT '',                 -- value of this string
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY(SLSID)
);

CREATE TABLE NoteType (
    NTID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id of this note type
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business associated with this NoteType
    Name VARCHAR(128) NOT NULL DEFAULT '',                  -- General, Payment, Receipt, ...
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (NTID)     
);

CREATE TABLE Notes (
    NID BIGINT NOT NULL AUTO_INCREMENT,                     -- ID for this note
    NLID BIGINT NOT NULL DEFAULT 0,                         -- note list containing this note
    PNID BIGINT NOT NULL DEFAULT 0,                         -- NID of parent not
    NTID BIGINT NOT NULL DEFAULT 0,                         -- What type of note is this
    RID BIGINT NOT NULL DEFAULT 0,                          -- Meta-tag - this note is related to Rentable RID
    RAID BIGINT NOT NULL DEFAULT 0,                         -- Meta-tag - this note is related to Rentable Agreement RAID
    TCID BIGINT NOT NULL DEFAULT 0,                         -- Meta-tag - this note is related to Transactant TCID
    Comment VARCHAR(1024) NOT NULL DEFAULT '',              -- the actual note
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (NID)     
);

CREATE TABLE NoteList (
    NLID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id for this notelist
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (NLID)     
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
    Units VARCHAR (256) NOT NULL DEFAULT '',   -- optional units value
    LastModTime TIMESTAMP,                     -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,    -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (CID)
);

CREATE TABLE CustomAttrRef (
    ElementType BIGINT NOT NULL,   -- for what type of object is this a ref:  1=Person, 2=Company, 3=Business-Unit, 4=executable service, 5=RentableType
    ID          BIGINT NOT NULL,   -- the UID of the object type. That is, if ObjectType == 5, the ID is the RTID (Rentable type id)
    CID         BIGINT NOT NULL    -- uid of the custom attribute
);

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
    RAID BIGINT NOT NULL AUTO_INCREMENT,                         -- internal unique id
    RATID BIGINT NOT NULL DEFAULT 0,                             -- reference to Rental Template (Occupancy Master Agreement)
    BID BIGINT NOT NULL DEFAULT 0,                               -- Business (so that we can process by Business)
    NLID BIGINT NOT NULL DEFAULT 0,                              -- NoteList ID
    AgreementStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- date when rental starts
    AgreementStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- date when rental stops
    PossessionStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00', -- date when Occupancy starts
    PossessionStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- date when Occupancy stops
    RentStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',       -- date when Rent starts
    RentStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',        -- date when Rent stops
    Renewal SMALLINT NOT NULL DEFAULT 0,                         -- 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
    SpecialProvisions VARCHAR(1024) NOT NULL DEFAULT '',         -- free-form text
    LastModTime TIMESTAMP,                                       -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RAID)
);

CREATE TABLE RentalAgreementRentables (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    RID BIGINT NOT NULL DEFAULT 0,                            -- Rentable id
    CLID BIGINT NOT NULL DEFAULT 0,                           -- Commission Ledger (for outside salespeople to get a commission)
    ContractRent DECIMAL(19,4) NOT NULL DEFAULT 0.0,          -- The contract rent for this rentable
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this Rentable was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this Rentable was no longer being billed to this agreement
);

CREATE TABLE RentalAgreementPayors (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    TCID BIGINT NOT NULL DEFAULT 0,                            -- who is the Payor for this agreement
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this Payor was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this Payor was no longer being billed to this agreement
);

CREATE TABLE RentableUsers (
    RID BIGINT NOT NULL DEFAULT 0,                            -- the associated Rentable
    -- USERID BIGINT NOT NULL DEFAULT 0,                      -- the Users of the rentable
    TCID BIGINT NOT NULL DEFAULT 0,                           -- the Users of the rentable
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this User was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00'        -- date when this User was no longer being billed to this agreement
);

CREATE TABLE RentalAgreementPets (
    PETID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal id for this pet
    RAID BIGINT NOT NULL DEFAULT 0,                           -- the unit's occupancy agreement
    Type VARCHAR(100) NOT NULL DEFAULT '',                    -- type of animal, ex: dog, cat, ...
    Breed VARCHAR(100) NOT NULL DEFAULT '',                   -- breed.  example Beagle, German Shephard, Siamese, etc.
    Color VARCHAR(100) NOT NULL DEFAULT '',                   -- fur or other color
    Weight DECIMAL(19,4) NOT NULL DEFAULT 0,                  -- in pounds
    Name VARCHAR(100) NOT NULL DEFAULT '',                    -- the pet's name
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this User was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',       -- date when this User was no longer being billed to this agreement
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (PETID)
);

-- Referenced by a Rentable associated with a RentalAgreement  (RentalAgreementRentables)
CREATE TABLE CommissionLedger (
    CLID BIGINT NOT NULL AUTO_INCREMENT,            -- unique id for this Commission Ledger
    RAID BIGINT NOT NULL DEFAULT 0,                 -- associated with this RAID
    RID BIGINT NOT NULL DEFAULT 0,                  -- associated with this rentable??????
    Salesperson  VARCHAR(100) NOT NULL DEFAULT '',  -- who referred
    Percent DECIMAL(19,4) NOT NULL DEFAULT 0,       -- what percent are we paying them. If 0 then we're paying a specific Amount
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0,        -- what amount are we paying them. If 0 then we're paying a percentage
    PaymentDueDate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',     -- enterer will fill it out
    PRIMARY KEY(CLID)
);

-- **************************************
-- ****                              ****
-- ****          RATE PLAN           ****
-- ****                              ****
-- ************************************** 
CREATE TABLE RatePlan (
    RPID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                      -- Business
    Name VARCHAR(100) NOT NULL DEFAULT '',              -- The name of this RatePlan
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY(RPID)
);
    -- these flags are Property-Industry-specific.  IMPLEMENT THESE
    -- FLAGS BIGINT NOT NULL DEFAULT 0,
            -- 1<<1   GDSAvailable, if 1 then this rate plan can be made available on GDS
            -- 1<<2   SaberAvailable, if 1 then this rate plan canb be made available on Saber

-- RatePlanRef contains the time sensitive attributes of a RatePlan
CREATE TABLE RatePlanRef (
    RPRID BIGINT NOT NULL AUTO_INCREMENT,               -- unique id for this rate plan
    RPID BIGINT NOT NULL DEFAULT 0,                     -- which rateplan
    DtStart DATE NULL DEFAULT '1970-01-01 00:00:00',    -- when does it go into effect
    DtStop DATE NULL DEFAULT '1970-01-01 00:00:00',     -- when does it stop
    FeeAppliesAge SMALLINT NOT NULL DEFAULT 0,          -- the age at which a user is counted when determining extra user fees or eligibility for rental
    MaxNoFeeUsers SMALLINT NOT NULL DEFAULT 0,          -- maximum number of users for no fees. Greater than this number means fee applies
    AdditionalUserFee DECIMAL(19,4) NOT NULL DEFAULT 0, -- extra fee per user when exceeding MaxNoFeeUsers
    PromoCode VARCHAR(100),                             -- just a string
    CancellationFee DECIMAL(19,4) NOT NULL DEFAULT 0,    -- charge for cancellation
    FLAGS BIGINT NOT NULL DEFAULT 0,                    -- 1<<0 -- HideRate
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY(RPRID)
);

-- RatePlanRefRTRate is RatePlan RPRID's rate information for the RentableType (RTID)
CREATE TABLE RatePlanRefRTRate (
    RPRID BIGINT NOT NULL DEFAULT 0,        -- which RatePlanRef is this
    RTID BIGINT NOT NULL DEFAULT 0,         -- which RentableType
    FLAGS BIGINT NOT NULL DEFAULT 0,        -- 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate, 
    Val DECIMAL(19,4) NOT NULL DEFAULT 0    -- Val 
);
-- RatePlanRefSPRate is RatePlan RPRID's rate information for the Specialties
CREATE TABLE RatePlanRefSPRate (
    RPRID BIGINT NOT NULL DEFAULT 0,        -- which RatePlanRef is this
    RTID BIGINT NOT NULL DEFAULT 0,         -- which RentableType
    RSPID BIGINT NOT NULL DEFAULT 0,        -- which Specialty
    FLAGS BIGINT NOT NULL DEFAULT 0,        -- 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate, 
    Val DECIMAL(19,4) NOT NULL DEFAULT 0    -- Val 
);

-- Rate plans can have other deliverables. These can be things like 2 tickets to SeaWorld, free meal vouchers, etc.
-- A RatePlan can refer to multiple OtherDeliverables.  SELECT * FROM RatePlanOD WHERE RPID=MyRatePlan will return all
-- the OtherDeliverables associated with MyRatePlan
CREATE TABLE RatePlanOD (
    RPRID BIGINT NOT NULL DEFAULT 0,        -- with which RatePlanRef is this OtherDeliverable associated?
    ODID BIGINT NOT NULL DEFAULT 0          -- points to an OtherDeliverables 
);

-- These are for promotions - like 2 Seaworld tickets, etc.  Referenced by Rate Plan Refs
-- Multiple rate plans can refer to the same OtherDeliverables. 
CREATE TABLE OtherDeliverables (
    ODID BIGINT NOT NULL AUTO_INCREMENT,    -- Unique ID for this OtherDeliverables
    Name VARCHAR(256),                      -- Description of the other deliverables. Ex: 2 Seaworld tickets 
    Active SMALLINT NOT NULL DEFAULT 0,     -- Flag: Is this list still active?  0 = not active, 1 = active
    PRIMARY KEY(ODID)
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
    GSRPC BIGINT NOT NULL DEFAULT 0,                        -- Increments in which GSR is calculated to account for rate changes
    ManageToBudget SMALLINT NOT NULL DEFAULT 0,             -- 0 do not manage this category of Rentable to budget, 1 = manage to budget defined by MarketRate
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (RTID)
);

CREATE TABLE RentableMarketRate (
    RTID BIGINT NOT NULL DEFAULT 0,                             -- associated Rentable type
    MarketRate DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- market rate for the time range
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',    
    DtStop DATETIME NOT NULL DEFAULT '9999-12-31 23:59:59'      -- assume it's unbounded. if an updated Market rate is added, set this to the stop date
);

-- RentableType RTID needs to have tax TAXID applied to rental assessments.
-- There can be as many of these records as needed per rentable type.
CREATE TABLE RentableTaxes (
    RTID BIGINT NOT NULL DEFAULT 0,                             -- associated Rentable type
    TAXID BIGINT NOT NULL DEFAULT 0                             -- which tax
);

-- ===========================================
--   RENTABLE SPECIALTY TYPES
-- ===========================================
-- a collection of unit specialties that are available.
-- different units may be more or less desirable based upon special characteristics
-- of the unit, such as Lake View, Courtyard View, Washer Dryer Connections, 
-- Washer Dryer provided, close to parking, better views, fireplaces, special 
-- remodeling or finishes, etc.  This is where those special characteristics are defined
CREATE TABLE RentableSpecialty (
    RSPID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL,
    Name VARCHAR(100) NOT NULL DEFAULT '',
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    Description VARCHAR(256) NOT NULL DEFAULT '',
    PRIMARY KEY (RSPID)
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
    ATypeLID BIGINT NOT NULL DEFAULT 0
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
    -- RentalPeriodDefault SMALLINT NOT NULL DEFAULT 0,            -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
    -- RentCycle SMALLINT NOT NULL DEFAULT 0,                      -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
);

CREATE TABLE RentableTypeRef (
    RID BIGINT NOT NULL DEFAULT 0,                                  -- the Rentable this record belongs to
    RTID BIGINT NOT NULL DEFAULT 0,                                 -- the Rentable type for this period
    RentCycle BIGINT NOT NULL DEFAULT 0,                            -- RentCycle override. 0 = unset (use RentableType.RentCycle), > 0 means the override frequency
    ProrationCycle BIGINT NOT NULL DEFAULT 0,                       -- Proration override. 0 = unset (use RentableType.Proration), > 0 means the override proration
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- start time for this state
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- stop time for this state
    LastModTime TIMESTAMP,                                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that modified it 
);

CREATE TABLE RentableStatus (
    RID BIGINT NOT NULL DEFAULT 0,                                  -- associated Rentable
    Status SMALLINT NOT NULL DEFAULT 0,                             -- 0 = UNKNOWN -- 1 = ONLINE, 2 = ADMIN, 3 = EMPLOYEE, 4 = OWNEROCC, 5 = OFFLINE, 
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- start time for this state
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- stop time for this state
    DtNoticeToVacate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- user has indicated they will vacate on this date
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
    ASMID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique id for assessment
    PASMID BIGINT NOT NULL DEFAULT 0,                       -- parent Assessment, if this is non-zero it means this assessment is an instance of the recurring assessment with id PASMID.
                                                            --     When non-zero DO NOT process as a recurring assessment, it is an instance
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business id
    RID BIGINT NOT NULL DEFAULT 0,                          -- rentable id
    ATypeLID BIGINT NOT NULL DEFAULT 0,                     -- Ledger ID describing the type of assessment (ex: Rent, SecurityDeposit, ...)
    RAID BIGINT NOT NULL DEFAULT 0,                         -- Associated Rental Agreement ID
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- Assessment amount
    Start DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- epoch date for recurring assessments; the date/time of the assessment for instances
    Stop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00',   -- stop date for recurrent assessments; the date/time of the assessment for instances
    RentCycle SMALLINT NOT NULL DEFAULT 0,                  -- 0 = non-recurring, 1 = secondly, 2 = minutely, 3=hourly, 4=daily, 5=weekly, 6=monthly, 7=quarterly, 8=yearly
    ProrationCycle SMALLINT NOT NULL DEFAULT 0,             -- 
    InvoiceNo BIGINT NOT NULL DEFAULT 0,                    -- DELETE THIS -- DON'T KEEP THE INVOICE REFERENCE IN THE ASSESSMENT... !!!! <<<<TODO
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',              -- Accounting rule - which acct debited, which credited
    Comment VARCHAR(256) NOT NULL DEFAULT '',               -- for comments such as "Prior period adjustment"
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (ASMID)
);

-- the actual tax rate or fee will be read from the TaxRate table based on the instance date of the assessment
CREATE TABLE AssessmentTax (
    ASMID BIGINT NOT NULL DEFAULT 0,                        -- the assessment to which this tax is bound
    TAXID BIGINT NOT NULL DEFAULT 0,                        -- what type of tax.
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- bit 0 = override this tax -- do not apply, bit 1 - override and use OverrideAmount
    OverrideTaxApprover MEDIUMINT NOT NULL DEFAULT 0,       -- if tax is overridden, who approved it
    OverrideAmount DECIMAL(19,4) NOT NULL DEFAULT 0,        -- Don't calculate. Use this amount. OverrideApprover required.  0 if not applicable.
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0                  -- employee UID (from phonebook) that modified it 
);

-- **************************************
-- ****                              ****
-- ****           PEOPLE             ****
-- ****                              ****
-- **************************************
-- This is DemandSource  referenced by RentalAgreement
CREATE TABLE DemandSource (
    DSID BIGINT NOT NULL AUTO_INCREMENT,                    -- DemandSource ID - unique id for this source
    BID BIGINT NOT NULL DEFAULT 0,                          -- What business is this
    Name VARCHAR(100),                                      -- Name of the source
    Industry VARCHAR(100),                                  -- What industry -- THIS BECOMES A REFERENCE TO "Industry" StringList
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (DSID)
);

CREATE TABLE LeadSource (
    LSID BIGINT NOT NULL AUTO_INCREMENT,                    -- DemandSource ID - unique id for this source
    BID BIGINT NOT NULL DEFAULT 0,                          -- What business is this
    Name VARCHAR(100),                                      -- Name of the source
    IndustrySLID BIGINT NOT NULL DEFAULT 0,                 -- What industry -- THIS BECOMES A REFERENCE TO "Industry" StringList
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LSID)
);

-- ===========================================
--   TRANSACTANT
-- ===========================================
-- Transactant - fields common to all people and
CREATE TABLE Transactant (
    TCID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id of unit
    BID BIGINT NOT NULL DEFAULT 0,                          -- which business
    NLID BIGINT NOT NULL DEFAULT 0,                         -- notes associated with this transactant
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
    Address VARCHAR(100) NOT NULL DEFAULT '',               -- person address
    Address2 VARCHAR(100) NOT NULL DEFAULT '',       
    City VARCHAR(100) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(100) NOT NULL DEFAULT '',
    Country VARCHAR(100) NOT NULL DEFAULT '',
    Website VARCHAR(100) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP,                              -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,             -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

--    UseCount BIGINT NOT NULL DEFAULT 0,            -- This count is incremented each time a transactant enters into a RentalAgreement.  Count > 1 means it's a ReturnUser 
--    Flags BIGINT NOT NULL DEFAULT 0,                -- For flags as described below:
--        --   1<<0 OptIntoMarketingCampaign          -- Does the user want to receive mkting info
--        --   1<<1 AcceptGeneralEmail                -- Will user accept email
--        --   1<<2 VIP                               -- Is this person a VIP 

-- ===========================================
--   PROSPECT
-- ===========================================
CREATE TABLE Prospect (
    -- PRSPID BIGINT NOT NULL DEFAULT 0,                 -- unique id of this Prospect
    TCID BIGINT NOT NULL DEFAULT 0,                        -- associated Transactant (has Name and all contact info)
    EmployerName  VARCHAR(100) NOT NULL DEFAULT '',
    EmployerStreetAddress VARCHAR(100) NOT NULL DEFAULT '',
    EmployerCity VARCHAR(100) NOT NULL DEFAULT '',
    EmployerState VARCHAR(100) NOT NULL DEFAULT '',
    EmployerPostalCode VARCHAR(100) NOT NULL DEFAULT '',
    EmployerEmail VARCHAR(100) NOT NULL DEFAULT '',
    EmployerPhone VARCHAR(100) NOT NULL DEFAULT '',
    Occupation VARCHAR(100) NOT NULL DEFAULT '',
    ApplicationFee DECIMAL(19,4) NOT NULL DEFAULT 0.0,      -- if non-zero this Prospect is an applicant
    DesiredMoveInDate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- User's initial indication of move in date, actual move in date is in Rental Agreement
    RentableTypePreference BIGINT NOT NULL DEFAULT 0,          -- This would be "model" preference  (Rentable Type name) for room or residence, but could apply to all rentables 
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- bit 0 - approved/not approved
    Approver BIGINT NOT NULL DEFAULT 0,                     -- who approved or declined
    DeclineReasonSLSID BIGINT NOT NULL DEFAULT 0,           -- ID to string in list of choices, Melissa will provide the list.
    OtherPreferences VARCHAR(1024) NOT NULL DEFAULT '',     -- Arbitrary text, anything else they might request
    FollowUpDate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- automatically fill out this date to sysdate + 24hrs
    CSAgent BIGINT NOT NULL DEFAULT 0,                      -- Accord Directory UserID - for the CSAgent 
    OutcomeSLSID BIGINT NOT NULL DEFAULT 0,                 -- id of string from a list of outcomes. Melissa to provide reasons
    FloatingDeposit DECIMAL (19,4) NOT NULL DEFAULT 0.0,    --  d $(GLCASH) _, c $(GLGENRCV) _; assign to a shell of a Rental Agreement 
    RAID BIGINT NOT NULL DEFAULT 0,                         -- created to hold On Account amount of Floating Deposit
    LastModTime TIMESTAMP,                                  -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

-- --new  Custom Fields
-- NumberBedrooms -- SMALLINT NOT NULL DEFAULT 0,  This is unique to a room or residence. bedroom count
-- NumberOfPets   -- SMALLINT NOT NULL DEFAULT 0,    This is unique to a room or residence. may just add to formal pet schema
-- NumberOfPeople -- SMALLINT NOT NULL DEFAULT 0,  This is unique to a room or residence. count of people who will be living in the unit
-- --new
-- 

-- ===========================================
--   USER
-- ===========================================
CREATE TABLE User (
    TCID BIGINT NOT NULL DEFAULT 0,                                       -- associated Transactant
    Points BIGINT NOT NULL DEFAULT 0,                           -- bonus points for this User
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
    EligibleFutureUser SMALLINT NOT NULL DEFAULT 1,              -- yes/no
    Industry VARCHAR(100) NOT NULL DEFAULT '',                   -- (e.g., construction, retail, banking etc.)
    DSID BIGINT NOT NULL DEFAULT 0,                               -- (e.g., resident referral, newspaper, radio, post card, expedia, travelocity, etc.)
    LastModTime TIMESTAMP,                                       -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

-- ===========================================
--   PAYOR
-- ===========================================
CREATE TABLE Payor  (
    -- PID BIGINT NOT NULL DEFAULT 0,                         -- unique id of this Payor 
    TCID BIGINT NOT NULL DEFAULT 0,                                       -- associated Transactant
    TaxpayorID VARCHAR(25) NOT NULL DEFAULT '',
    CreditLimit DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AccountRep BIGINT NOT NULL DEFAULT 0,                       -- Phonebook UID of account rep
    EligibleFuturePayor SMALLINT NOT NULL DEFAULT 1,            -- yes/no
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (TCID)
);

-- **************************************
-- ****                              ****
-- ****           RECEIPTS           ****
-- ****                              ****
-- **************************************
CREATE TABLE Receipt (
    RCPTID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this Receipt
    PRCPTID BIGINT NOT NULL DEFAULT 0,                          -- Parent RCPT, if non-zero then it is the RCPTID of a receipt with an error that we're correcting in this receipt
    BID BIGINT NOT NULL DEFAULT 0,
    RAID BIGINT NOT NULL DEFAULT 0,                             -- THIS IS AN ISSUE... It can go away -- ReceiptAllocation has an associated Assessment, which has the RAID
    PMTID BIGINT NOT NULL DEFAULT 0,
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DocNo VARCHAR(50) NOT NULL DEFAULT '',                      -- Check Number, MoneyOrder number, etc., the traceback for the payment
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AcctRule VARCHAR(1500) NOT NULL DEFAULT '',
    Comment VARCHAR(256) NOT NULL DEFAULT '',                   -- for comments like "Prior Period Adjustment"
    OtherPayorName VARCHAR(128) NOT NULL DEFAULT '',            -- If not '' then Payment was made by a payor who is not on the RA, and may not be in our system at all
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
-- ****          DEPOSIT             ****
-- ****                              ****
-- **************************************
CREATE TABLE DepositMethod (
    DPMID BIGINT NOT NULL AUTO_INCREMENT, 
    BID BIGINT NOT NULL DEFAULT 0,                              -- which business
    Name VARCHAR(50) NOT NULL DEFAULT '',                       -- 0 = not specified, 1 = Hand Delivery, Scanned Batch, CC Shift 4, CC NAYAX, ACH, US Mail
    PRIMARY KEY (DPMID)
);

CREATE TABLE Depository (
    DEPID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for a depository
    BID BIGINT NOT NULL DEFAULT 0,                              -- business id
    Name VARCHAR(256),                                          -- Name of Depository: First Data, Nyax, Oklahoma Fidelity
    AccountNo VARCHAR(256),                                     -- account number at this Depository
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (DEPID)
);

CREATE TABLE Deposit (
    DID BIGINT NOT NULL AUTO_INCREMENT,                         -- UniqueID for this deposit
    BID BIGINT NOT NULL DEFAULT 0,                              -- business id
    DEPID BIGINT NOT NULL DEFAULT 0,                            -- DepositoryID where the Deposit was made
    DPMID BIGINT NOT NULL DEFAULT 0,                            -- Deposit Method
    Dt DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- Date of deposit
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                  -- total amount of all Receipts in this deposit
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (DID)
);

CREATE TABLE DepositPart (
    DID BIGINT NOT NULL DEFAULT 0,
    RCPTID BIGINT NOT NULL DEFAULT 0
);

-- **************************************
-- ****                              ****
-- ****          INVOICE             ****
-- ****                              ****
-- **************************************

CREATE TABLE Invoice (
    InvoiceNo BIGINT NOT NULL AUTO_INCREMENT,                   -- Unique id for this invoice
    BID BIGINT NOT NULL DEFAULT 0,                              -- bid (remit to)
    Dt DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- Date of invoice
    DtDue DATE NOT NULL DEFAULT '1970-01-01 00:00:00',          -- Date when the invoice is due
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                  -- total amount of all assessments in this invoice
    DeliveredBy VARCHAR(256) NOT NULL DEFAULT '',               -- mail, FedEx, UPS, ...
    LastModTime TIMESTAMP,                                      -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (InvoiceNo)
);

CREATE TABLE InvoiceAssessment (
    InvoiceNo BIGINT NOT NULL DEFAULT 0,                        -- which invoice
    ASMID BIGINT NOT NULL DEFAULT 0                             -- assessment id
);

CREATE TABLE InvoicePayor (
    InvoiceNo BIGINT NOT NULL DEFAULT 0,                        -- which invoice
    PID BIGINT NOT NULL DEFAULT 0                               -- Payor id
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
    LEID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this LedgerEntry
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    JID BIGINT NOT NULL DEFAULT 0,                            -- Journal entry giving rise to this
    JAID BIGINT NOT NULL DEFAULT 0,                           -- the allocation giving rise to this LedgerEntry
    LID BIGINT NOT NULL DEFAULT 0,                            -- associated GLAccount
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
    LID BIGINT NOT NULL DEFAULT 0,                            -- associated GLAccount
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                           -- 0 means it's the balance for the whole account;  > 0 means it's the amount associated with rental agreement RAID
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- Balance is valid as of this time
    Balance DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    State SMALLINT NOT NULL DEFAULT 0,                        -- 0 = Open, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
    LastModTime TIMESTAMP,                                    -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,                   -- employee UID (from phonebook) that modified it 
    PRIMARY KEY (LMID)
);

-- GL Account
CREATE TABLE GLAccount (
    LID BIGINT NOT NULL AUTO_INCREMENT,             -- unique id for this GLAccount
    PLID BIGINT NOT NULL DEFAULT 0,                 -- Parent ID for this GLAccount.  0 if no parent.
    BID BIGINT NOT NULL DEFAULT 0,                  -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                 -- rental agreement account, only valid if TYPE is 1
    GLNumber VARCHAR(100) NOT NULL DEFAULT '',      -- if not '' then it's a link a QB  GeneralLedger (GL)account
    Status SMALLINT NOT NULL DEFAULT 0,             -- Whether a GL Account is currently unknown=0, inactive=1, active=2 
    Type SMALLINT NOT NULL DEFAULT 0,               -- flag: 0 = not a special account of any kind, 
    --                                                       1 = RentalAgreement Receivable Balance, 
    --                                                       2 = RentalAgreement Security Deposit Balance,
    --                                                       3 - 9 Reserved
    --                                                       10-default cash, 11-GENRCV, 12-GrossSchedRENT, 13-LTL, 14-VAC, 15 sec dep receivable, 16 sec dep assessment
    Name VARCHAR(100) NOT NULL DEFAULT '',
    AcctType VARCHAR(100) NOT NULL DEFAULT '',      -- Quickbooks Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, 
                                                    --    Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, 
                                                    --    Cost of Goods Sold, Other Income, Other Expense
    RAAssociated SMALLINT NOT NULL DEFAULT 0,       -- 1 = Unassociated with RentalAgreement, 2 = Associated with Rental Agreement, 0 = unknown
    AllowPost SMALLINT NOT NULL DEFAULT 0,          -- 0 - do not allow posts to this ledger. 1 = allow posts
    RARequired SMALLINT NOT NULL DEFAULT 0,         -- 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
    ManageToBudget SMALLINT NOT NULL DEFAULT 0,     -- 0 = do not manage to budget; no ContractRent amount required. 1 = Manage to budget, ContractRent required.
    Description VARCHAR(1024) NOT NULL DEFAULT '',  -- describe the assessment
    LastModTime TIMESTAMP,                          -- when was this record last written
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,         -- employee UID (from phonebook) that modified it 
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
INSERT INTO GLAccount (BID,RAID,GLNumber,Status,Type,Name) VALUES
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
INSERT INTO LedgerMarker (BID,LID,State,Dt,Balance) VALUES
    (1,1,3,"2015-10-31",0.0),
    (1,2,3,"2015-10-31",0.0),
    (1,3,3,"2015-10-31",0.0),
    (1,4,3,"2015-10-31",0.0),
    (1,5,3,"2015-10-31",0.0),
    (1,6,3,"2015-10-31",0.0),
    (1,7,3,"2015-10-31",0.0),
    (1,8,3,"2015-10-31",0.0);
