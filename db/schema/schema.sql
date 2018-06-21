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
-- USERID = User id

DROP DATABASE IF EXISTS rentroll;
CREATE DATABASE rentroll;
USE rentroll;
GRANT ALL PRIVILEGES ON rentroll.* TO 'ec2-user'@'localhost';
set GLOBAL sql_mode='ALLOW_INVALID_DATES';

-- ===========================================
--   ID COUNTERS
--  may not need this
-- ===========================================
-- CREATE TABLE IDCounters (
--     InvoiceNo BIGINT NOT NULL DEFAULT 0                     -- unique number for invoices
-- );

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
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY(TAXID)
);

CREATE TABLE TaxRate (
    TAXID BIGINT NOT NULL DEFAULT 0,                        -- reference to which tax this table represents
    BID BIGINT NOT NULL DEFAULT 0,                          -- what business is this tax associated with
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',    -- date when this tax rate goes into effect
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',     -- date when this tax rate is no longer applicable
    Rate DECIMAL(19,4) NOT NULL DEFAULT 0,                  -- floating point number representing the rate. Set to 0 if not applicable.
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0,                   -- floating point number.  Set to 0 if not applicable.
    Formula VARCHAR(256) NOT NULL DEFAULT '',               -- RPN calculator notation of formula, Set to '' if not needed
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                      -- employee UID (from phonebook) that created this record
);

CREATE TABLE StringList (
    SLID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id for this stringlist
    BID BIGINT NOT NULL DEFAULT 0,                          -- the business to which this stringlist belongs
    Name VARCHAR(50) NOT NULL DEFAULT '',                   -- stringlist name
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY(SLID)
);

CREATE TABLE SLString (
    SLSID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique id for this string
    BID BIGINT NOT NULL DEFAULT 0,                          -- the business to which this stringlist belongs
    SLID BIGINT NOT NULL DEFAULT 0,                         -- to which stringlist does this string belong?
    Value VARCHAR(256) NOT NULL DEFAULT '',                 -- value of this string
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY(SLSID)
);

CREATE TABLE NoteType (
    NTID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id of this note type
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business associated with this NoteType
    Name VARCHAR(128) NOT NULL DEFAULT '',                  -- General, Payment, Receipt, Contact History ...
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that created this record
    PRIMARY KEY (NTID)
);

CREATE TABLE Notes (
    NID BIGINT NOT NULL AUTO_INCREMENT,                     -- ID for this note
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business associated with this NoteType
    NLID BIGINT NOT NULL DEFAULT 0,                         -- note list containing this note
    PNID BIGINT NOT NULL DEFAULT 0,                         -- NID of parent note
    NTID BIGINT NOT NULL DEFAULT 0,                         -- What type of note is this
    RID BIGINT NOT NULL DEFAULT 0,                          -- Meta-tag - this note is related to Rentable RID
    RAID BIGINT NOT NULL DEFAULT 0,                         -- Meta-tag - this note is related to Rentable Agreement RAID
    TCID BIGINT NOT NULL DEFAULT 0,                         -- Meta-tag - this note is related to Transactant TCID
    Comment VARCHAR(1024) NOT NULL DEFAULT '',              -- the actual note
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that created this record
    PRIMARY KEY (NID)
);

CREATE TABLE NoteList (
    NLID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id for this notelist
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business associated with this NoteType
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that created this record
    PRIMARY KEY (NLID)
);

-- **************************************
-- ****                              ****
-- ****    USER DEFINED ATTRIBUTES   ****
-- ****                              ****
-- **************************************
CREATE TABLE CustomAttr (
    CID BIGINT NOT NULL AUTO_INCREMENT,                     -- unique identifer for this custom attribute
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business associated with this NoteType
    Type SMALLINT NOT NULL DEFAULT 0,                       -- 0 = string, 1 = int64, 2 = float64
    Name VARCHAR (100) NOT NULL DEFAULT '',                 -- a name
    Value VARCHAR (256) NOT NULL DEFAULT '',                -- its value in string form
    Units VARCHAR (256) NOT NULL DEFAULT '',                -- optional units value
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (CID)
);

CREATE TABLE CustomAttrRef (
    CARID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique identifer for this custom attribute Reference
    ElementType BIGINT NOT NULL,                            -- for what type of object is this a ref:  1=Person, 2=Company, 3=Business-Unit, 4=executable service, 5=RentableType
    BID         BIGINT NOT NULL DEFAULT 0,                  -- Business associated with this NoteType
    ID          BIGINT NOT NULL,                            -- the UID of the object type. That is, if ObjectType == 5, the ID is the RTID (Rentable type id)
    CID         BIGINT NOT NULL,                            -- uid of the custom attribute
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (CARID)
);

-- ===========================================
--   RENTAL AGREEMENT TEMPLATE
-- ===========================================
CREATE TABLE RentalAgreementTemplate (
    RATID BIGINT NOT NULL AUTO_INCREMENT,                   -- internal unique id
    BID BIGINT NOT NULL DEFAULT 0,                          -- BizUnit Reference
    RATemplateName VARCHAR(100) DEFAULT '',                 -- Rental Agreement Template Name
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RATID)
);

-- ===========================================
--   RENTAL AGREEMENT
-- ===========================================
CREATE TABLE RentalAgreement (
    RAID BIGINT NOT NULL AUTO_INCREMENT,                                -- internal unique id
    RATID BIGINT NOT NULL DEFAULT 0,                                    -- reference to Rental Template (Occupancy Master Agreement)
    BID BIGINT NOT NULL DEFAULT 0,                                      -- Business (so that we can process by Business)
    NLID BIGINT NOT NULL DEFAULT 0,                                     -- NoteList ID
    AgreementStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',         -- date when rental starts (may be blank if RA initiated for floating deposit)
    AgreementStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',          -- date when rental stops  (may be blank if RA initiated for floating deposit)
    PossessionStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',        -- date when usage starts  (may be blank if RA initiated for floating deposit)
    PossessionStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',         -- date when usage stops   (may be blank if RA initiated for floating deposit)
    RentStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',              -- date when Rent starts   (may be blank if RA initiated for floating deposit)
    RentStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',               -- date when Rent stops    (may be blank if RA initiated for floating deposit)
    RentCycleEpoch DATE NOT NULL DEFAULT '1970-01-01 00:00:00',         -- Date on which rent cycle recurs. Start date for the recurring rent assessment
    -- FloatingDepositAssessment DATE NOT NULL DEFAULT '1970-01-01 00:00:00'  -- Date on which floating deposit was assessed.
    UnspecifiedAdults SMALLINT NOT NULL DEFAULT 0,                      -- # of Adults who are NOT accounted for in RentalAgreementPayor and RentableUser entries. Useful in hotels
    UnspecifiedChildren SMALLINT NOT NULL DEFAULT 0,                    -- # of Children who are NOT transactants that will participate in the possession of the rentable
    Renewal SMALLINT NOT NULL DEFAULT 0,                                -- 0 = not set, 1 = month to month automatic renewal, 2 = lease extension options
    SpecialProvisions VARCHAR(1024) NOT NULL DEFAULT '',                -- free-form text
    LeaseType BIGINT NOT NULL DEFAULT 0,                                -- Full Service Gross, Gross, ModifiedGross, Tripple Net
    ExpenseAdjustmentType BIGINT NOT NULL DEFAULT 0,                    -- Base Year, No Base Year, Pass Through
    ExpensesStop DECIMAL(19,4) NOT NULL DEFAULT 0,                      -- cap on the amount of oexpenses that can be passed through to the tenant
    ExpenseStopCalculation VARCHAR(128) NOT NULL DEFAULT '',            -- note on how to determine the expense stop
    BaseYearEnd DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- last day of the base year
    ExpenseAdjustment DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- the next date on which an expense adjustment is due
    EstimatedCharges DECIMAL(19,4) NOT NULL DEFAULT 0,                  -- a periodic fee charged to the tenant to reimburse LL for anticipated expenses
    RateChange DECIMAL(19,4) NOT NULL DEFAULT 0,                        -- predetermined amount of rent increase, expressed as a percentage
    CSAgent BIGINT NOT NULL DEFAULT 0,                                  -- Accord Directory UserID - for the CSAgent
    NextRateChange DATE NOT NULL DEFAULT '1970-01-01 00:00:00',         -- the next date on which a RateChange will occur
    PermittedUses VARCHAR(128) NOT NULL DEFAULT '',                     -- indicates primary use of the space, ex: doctor's office, or warehouse/distribution, etc.
    ExclusiveUses VARCHAR(128) NOT NULL DEFAULT '',                     -- those uses to which the tenant has the exclusive rights within a complex, ex: Trader Joe's may have the exclusive right to sell groceries
    ExtensionOption VARCHAR(128) NOT NULL DEFAULT '',                   -- the right to extend the term of lease by giving notice to LL, ex: 2 options to extend for 5 years each
    ExtensionOptionNotice DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- the last date by which a Tenant can give notice of their intention to exercise the right to an extension option period
    ExpansionOption VARCHAR(128) NOT NULL DEFAULT '',                   -- the right to expand to certanin spaces that are typically contiguous to their primary space
    ExpansionOptionNotice DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- the last date by which a Tenant can give notice of their intention to exercise the right to an Expansion Option
    RightOfFirstRefusal VARCHAR(128) NOT NULL DEFAULT '',               -- Tenant may have the right to purchase their premises if LL chooses to sell
    FLAGS BIGINT NOT NULL DEFAULT 0,                                    -- 1<<0 - is application pending?
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                                -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,              -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                                 -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RAID)
);

CREATE TABLE RentalAgreementRentables (
    RARID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal unique id
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    RID BIGINT NOT NULL DEFAULT 0,                            -- Rentable ID
    PRID BIGINT NOT NULL DEFAULT 0,                           -- Parent Rentable ID
    CLID BIGINT NOT NULL DEFAULT 0,                           -- Commission Ledger (for outside salespeople to get a commission)
    ContractRent DECIMAL(19,4) NOT NULL DEFAULT 0.0,          -- The contract rent for this rentable
    RARDtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- date when this Rentable was added to the agreement
    RARDtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',    -- date when this Rentable was no longer being billed to this agreement
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                 -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RARID)
);

CREATE TABLE RentalAgreementPayors (
    RAPID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal unique id
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    TCID BIGINT NOT NULL DEFAULT 0,                           -- who is the Payor for this agreement
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this Payor was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',       -- date when this Payor was no longer being billed to this agreement
    FLAGS BIGINT NOT NULL DEFAULT 0,                          -- 1 << 0 is the bit that indicates this payor is a 'guarantor'
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RAPID)
);

CREATE TABLE RentableUsers (
    RUID BIGINT NOT NULL AUTO_INCREMENT,                      -- internal unique id
    RID BIGINT NOT NULL DEFAULT 0,                            -- the associated Rentable
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    TCID BIGINT NOT NULL DEFAULT 0,                           -- the Users of the rentable
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this User was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00' ,      -- date when this User was no longer being billed to this agreement
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RUID)
);

CREATE TABLE RentalAgreementTax (
    RAID BIGINT NOT NULL DEFAULT 0,                           -- Rental Agreement id
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this flag went into effect
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',       -- date when this flag was no longer in effect
    FLAGS BIGINT NOT NULL DEFAULT 0,                          -- 1 << 0 is the bit that indicates whether or not the rental agreement is taxable
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                        -- employee UID (from phonebook) that created this record
);

CREATE TABLE RentalAgreementPets (
    PETID BIGINT NOT NULL AUTO_INCREMENT,                     -- internal id for this pet
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    RAID BIGINT NOT NULL DEFAULT 0,                           -- the unit's occupancy agreement
    TCID BIGINT NOT NULL DEFAULT 0,                           -- Contact person for this pet
    Type VARCHAR(100) NOT NULL DEFAULT '',                    -- type of animal, ex: dog, cat, ...
    Breed VARCHAR(100) NOT NULL DEFAULT '',                   -- breed.  example Beagle, German Shephard, Siamese, etc.
    Color VARCHAR(100) NOT NULL DEFAULT '',                   -- fur or other color
    Weight DECIMAL(19,4) NOT NULL DEFAULT 0,                  -- in pounds
    Name VARCHAR(100) NOT NULL DEFAULT '',                    -- the pet's name
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',      -- date when this User was added to the agreement
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',       -- date when this User was no longer being billed to this agreement
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (PETID)
);

-- Referenced by a Rentable associated with a RentalAgreement  (RentalAgreementRentables)
CREATE TABLE CommissionLedger (
    CLID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this Commission Ledger
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business (so that we can process by Business)
    RAID BIGINT NOT NULL DEFAULT 0,                           -- associated with this RAID
    RID BIGINT NOT NULL DEFAULT 0,                            -- associated with this rentable??????
    Salesperson  VARCHAR(100) NOT NULL DEFAULT '',            -- who referred
    Percent DECIMAL(19,4) NOT NULL DEFAULT 0,                 -- what percent are we paying them. If 0 then we're paying a specific Amount
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0,                  -- what amount are we paying them. If 0 then we're paying a percentage
    PaymentDueDate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',  -- enterer will fill it out
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY(CLID)
);

-- **************************************
-- ****                              ****
-- ****          RATE PLAN           ****
-- ****                              ****
-- **************************************
CREATE TABLE RatePlan (
    RPID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business
    Name VARCHAR(100) NOT NULL DEFAULT '',                    -- The name of this RatePlan
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY(RPID)
);
    -- these flags are Property-Industry-specific.  IMPLEMENT THESE
    -- FLAGS BIGINT NOT NULL DEFAULT 0,
            -- 1<<1   GDSAvailable, if 1 then this rate plan can be made available on GDS
            -- 1<<2   SaberAvailable, if 1 then this rate plan canb be made available on Saber

-- RatePlanRef contains the time sensitive attributes of a RatePlan
CREATE TABLE RatePlanRef (
    RPRID BIGINT NOT NULL AUTO_INCREMENT,                     -- unique id for this rate plan
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business
    RPID BIGINT NOT NULL DEFAULT 0,                           -- which rateplan
    DtStart DATE NULL DEFAULT '1970-01-01 00:00:00',          -- when does it go into effect
    DtStop DATE NULL DEFAULT '1970-01-01 00:00:00',           -- when does it stop
    FeeAppliesAge SMALLINT NOT NULL DEFAULT 0,                -- the age at which a user is counted when determining extra user fees or eligibility for rental
    MaxNoFeeUsers SMALLINT NOT NULL DEFAULT 0,                -- maximum number of users for no fees. Greater than this number means fee applies
    AdditionalUserFee DECIMAL(19,4) NOT NULL DEFAULT 0,       -- extra fee per user when exceeding MaxNoFeeUsers
    PromoCode VARCHAR(100),                                   -- just a string
    CancellationFee DECIMAL(19,4) NOT NULL DEFAULT 0,         -- charge for cancellation
    FLAGS BIGINT NOT NULL DEFAULT 0,                          -- 1<<0 -- HideRate
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY(RPRID)
);

-- RatePlanRefRTRate is RatePlan RPRID's rate information for the RentableType (RTID)
CREATE TABLE RatePlanRefRTRate (
    RPRRTRateID BIGINT NOT NULL AUTO_INCREMENT,                 -- unique id for this rate plan ref RT Rate
    RPRID BIGINT NOT NULL DEFAULT 0,                            -- which RatePlanRef is this
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business
    RTID BIGINT NOT NULL DEFAULT 0,                             -- which RentableType
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate,
    Val DECIMAL(19,4) NOT NULL DEFAULT 0,                       -- Val
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RPRRTRateID)
);

-- RatePlanRefSPRate is RatePlan RPRID's rate information for the Specialties
CREATE TABLE RatePlanRefSPRate (
    RPRSPRateID BIGINT NOT NULL AUTO_INCREMENT,                -- unique id for this rate plan ref SP Rate
    RPRID BIGINT NOT NULL DEFAULT 0,                            -- which RatePlanRef is this
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business
    RTID BIGINT NOT NULL DEFAULT 0,                             -- which RentableType
    RSPID BIGINT NOT NULL DEFAULT 0,                            -- which Specialty
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 = percent flag 0 = Val is an absolute price, 1 = percent of MarketRate,
    Val DECIMAL(19,4) NOT NULL DEFAULT 0,                       -- Val
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,       -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RPRSPRateID)
);

-- Rate plans can have other deliverables. These can be things like 2 tickets to SeaWorld, free meal vouchers, etc.
-- A RatePlan can refer to multiple OtherDeliverables.  SELECT * FROM RatePlanOD WHERE RPID=MyRatePlan will return all
-- the OtherDeliverables associated with MyRatePlan
CREATE TABLE RatePlanOD (
    RPRID BIGINT NOT NULL DEFAULT 0,                            -- with which RatePlanRef is this OtherDeliverable associated?
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business
    ODID BIGINT NOT NULL DEFAULT 0,                             -- points to an OtherDeliverables
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that created this record
);

-- These are for promotions - like 2 Seaworld tickets, etc.  Referenced by Rate Plan Refs
-- Multiple rate plans can refer to the same OtherDeliverables.
CREATE TABLE OtherDeliverables (
    ODID BIGINT NOT NULL AUTO_INCREMENT,                      -- Unique ID for this OtherDeliverables
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business
    Name VARCHAR(256),                                        -- Description of the other deliverables. Ex: 2 Seaworld tickets
    Active TINYINT(1) NOT NULL DEFAULT 0,                       -- Flag: Is this list still active?  0 = not active, 1 = active
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
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
    BUD VARCHAR(100) NOT NULL DEFAULT '',                       -- Business Unit Designation
    Name VARCHAR(100) NOT NULL DEFAULT '',                      -- Business Full Name
    DefaultRentCycle SMALLINT NOT NULL DEFAULT 0,               -- default for every rentable type - useful to initialize UI
    DefaultProrationCycle SMALLINT NOT NULL DEFAULT 0,          -- default for every rentable type - useful to initialize UI
    DefaultGSRPC SMALLINT NOT NULL DEFAULT 0,                   -- default for every rentable type - useful to initialize UI
    ClosePeriodTLID BIGINT NOT NULL DEFAULT 0,                  -- The tasklist needed for closing a period
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- last bit =0(EDI disabled), =1(EDI enabled)
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (BID)
);
--    ParkingPermitInUse SMALLINT NOT NULL DEFAULT 0,           -- yes/no  0 = no, 1 = yes

-- ===========================================
--   Business Properties
-- ===========================================
CREATE TABLE BusinessProperties (
    BPID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business
    Name VARCHAR(100) NOT NULL DEFAULT '',                      -- Property Name
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- last bit =0(EDI disabled), =1(EDI enabled)
    Data JSON DEFAULT NULL,                                     -- JSON Data for this property
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (BPID)
);

-- ===========================================
--   TASKLIST AND TASK
-- ===========================================
CREATE TABLE Task (
    TID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,
    TLID BIGINT NOT NULL DEFAULT 0,                             -- the TaskList to which this task belongs
    Name VARCHAR(256) NOT NULL DEFAULT '',                      -- Task text
    Worker VARCHAR(80) NOT NULL DEFAULT '',                     -- Name of the associated work function
    DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',      -- Task Due Date
    DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- Pre Completion due date
    DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',     -- Task completion Date
    DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- Task Pre Completion Date
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 pre-completion required (if 0 then there is no pre-completion required)
                                                                -- 1<<1 PreCompletion done (if 0 it is not yet done)
                                                                -- 1<<2 Completion done (if 0 it is not yet done)
    DoneUID BIGINT NOT NULL DEFAULT 0,                          -- user who marked this task done
    PreDoneUID BIGINT NOT NULL DEFAULT 0,                       -- user who marked this task predone
    Comment VARCHAR(2048) NOT NULL DEFAULT '',                  -- any user comments
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY(TID)
);

CREATE TABLE TaskList (
    TLID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,
    PTLID BIGINT NOT NULL DEFAULT 0,                            -- Parent TLID or 0 if this is the parent (first) of a repeating set
    TLDID BIGINT NOT NULL DEFAULT 0,                            -- the TaskListDefinition that describes this tasklist
    Name VARCHAR(256) NOT NULL DEFAULT '',                      -- TaskList name
    Cycle BIGINT NOT NULL DEFAULT 0,                            -- recurrence frequency (not editable)
    DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',      -- All tasks in task list are due on this date
    DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- All tasks in task list pre-completion date
    DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',     -- Task completion Date
    DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- Task Pre Completion Date
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 : 0 = active, 1 = inactive
                                                                -- 1<<1 : 0 = task list definition does not have a PreDueDate, 1 = has a PreDueDate
                                                                -- 1<<1 : 0 = task list definition does not have a DueDate, 1 = has a DueDate
                                                                -- 1<<3 : 0 = DtPreDue has not been set, 1 = DtPreDue has been set
                                                                -- 1<<4 : 0 = DtDue has not been set, 1 = DtDue has been set
                                                                -- 1<<5 : 0 = no notification has been sent, 1 = Notification sent on DtLastNotify
    DoneUID BIGINT NOT NULL DEFAULT 0,                          -- user who marked this task done
    PreDoneUID BIGINT NOT NULL DEFAULT 0,                       -- user who marked this task predone
    EmailList VARCHAR(2048) NOT NULL DEFAULT '',                -- list of email addresses for when due date arrives
    DtLastNotify DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00', -- timestamp of last notification
    DurWait BIGINT NOT NULL DEFAULT 86400000000000,             -- how long to wait after failure notification for next check (default: 1 day)
    Comment VARCHAR(2048) NOT NULL DEFAULT '',                  -- any user comments
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY(TLID)
);

CREATE TABLE TaskDescriptor (
    TDID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,
    TLDID BIGINT NOT NULL DEFAULT 0,                            -- the TaskListDefinition to which this taskDescr belongs
    Name VARCHAR(256) NOT NULL DEFAULT '',                      -- Task text
    Worker VARCHAR(80) NOT NULL DEFAULT '',                     -- Name of the associated work function
    EpochDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- Task Due Date
    EpochPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00', -- Pre Completion due date
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 pre-completion required (if 0 then there is no pre-completion required)
    Comment VARCHAR(2048) NOT NULL DEFAULT '',                  -- any user comments
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY(TDID)
);

CREATE TABLE TaskListDefinition (
    TLDID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,
    Name VARCHAR(256) NOT NULL DEFAULT '',                      -- TaskList name
    Cycle BIGINT NOT NULL DEFAULT 0,                            -- recurrence frequency (editable)
    Epoch DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',      -- TaskList Start Date - day on which the instance is initiated
    EpochDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- Task Due Date
    EpochPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',-- Pre Completion due date
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 : 0 = active, 1 = inactive
    EmailList VARCHAR(2048) NOT NULL DEFAULT '',                -- list of email addresses for when due date arrives - will apply to all TaskList instances
    DurWait BIGINT NOT NULL DEFAULT 86400000000000,             -- how long to wait after failure notification for next check (default: 1 day)
    Comment VARCHAR(2048) NOT NULL DEFAULT '',                  -- any user comments
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY(TLDID)
);


-- ===========================================
--   RENTABLE TYPES
-- ===========================================
CREATE TABLE RentableTypes (
    RTID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                              -- associated Business id
    Style CHAR(255) NOT NULL DEFAULT '',                        -- need not be unique
    Name VARCHAR(256) NOT NULL DEFAULT '',                      -- must be unique
    RentCycle BIGINT NOT NULL DEFAULT 0,                        -- rent accrual frequency
    Proration BIGINT NOT NULL DEFAULT 0,                        -- prorate frequency
    GSRPC BIGINT NOT NULL DEFAULT 0,                            -- Increments in which GSR is calculated to account for rate changes
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0:  0=active, 1=inactive
                                                                -- 1<<1:  0=cannot be a child rentable, 1 = can be a child
                                                                -- 1<<2:  0=No(do not manage this category of Rentable to budget)
                                                                --        1=Yes(manage to budget defined by MarketRate & MRs are required)
    ARID BIGINT NOT NULL DEFAULT 0,                             -- ARID reference, for default rent amount for this rentable types
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RTID)
);

CREATE TABLE RentableMarketRate (
    RMRID BIGINT NOT NULL AUTO_INCREMENT,
    RTID BIGINT NOT NULL DEFAULT 0,                             -- associated Rentable type
    BID BIGINT NOT NULL DEFAULT 0,                              -- associated Business id
    MarketRate DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- market rate for the time range
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DtStop DATETIME NOT NULL DEFAULT '9999-12-31 23:59:59',     -- assume it's unbounded. if an updated Market rate is added, set this to the stop date
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RMRID)
);

-- RentableType RTID needs to have tax TAXID applied to rental assessments.
-- There can be as many of these records as needed per rentable type.
CREATE TABLE RentableTypeTax (
    RTID BIGINT NOT NULL DEFAULT 0,                             -- associated Rentable type
    BID BIGINT NOT NULL DEFAULT 0,                              -- associated Business id
    TAXID BIGINT NOT NULL DEFAULT 0,                            -- which tax
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DtStop DATETIME NOT NULL DEFAULT '9999-12-31 23:59:59',     -- assume it's unbounded. if an updated Market rate is added, set this to the stop date
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that created this record
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
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RSPID)
);

-- ===========================================
--   PAYMENT TYPE
-- ===========================================
CREATE TABLE PaymentType (
    PMTID MEDIUMINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL,
    Name VARCHAR(100) NOT NULL DEFAULT '',
    Description VARCHAR(256) NOT NULL DEFAULT '',
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
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
    BID BIGINT NOT NULL,
    Name VARCHAR(100) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
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
    ATypeLID BIGINT NOT NULL DEFAULT 0,
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                              -- employee UID (from phonebook) that created this record
);

-- applicable Assessments for a specific Business
CREATE TABLE BusinessPaymentTypes (
    BID BIGINT NOT NULL DEFAULT 0,
    PMTID MEDIUMINT NOT NULL DEFAULT 0,
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                              -- employee UID (from phonebook) that created this record
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
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (BLDGID)
);

-- **************************************
-- ****                              ****
-- ****          RENTABLE            ****
-- ****                              ****
-- **************************************
CREATE TABLE Rentable (
    RID BIGINT NOT NULL AUTO_INCREMENT,                             -- unique identifier for this Rentable
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business associated with this Rentable
    RentableName VARCHAR(100) NOT NULL DEFAULT '',                  -- must be unique, name for this instance, "101" for a room number, CP744 carport number, etc
    AssignmentTime SMALLINT NOT NULL DEFAULT 0,                     -- Unknown = 0, OK to pre-assign = 1, assign at occupancy commencement = 2
    MRStatus SMALLINT NOT NULL DEFAULT 0,                           -- Make Ready Status - current value as of DtMR, when this value changes it goes into a MRHistory record
    DtMRStart TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,         -- Time that MRStatus was set
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    Comment VARCHAR(2048) NOT NULL DEFAULT '',                      -- For notes such as Alarm codes, and other things
    PRIMARY KEY (RID)
    -- RentalPeriodDefault SMALLINT NOT NULL DEFAULT 0,             -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
    -- RentCycle SMALLINT NOT NULL DEFAULT 0,                       -- 0 = one time only, 1 = secondly, 2 = minutely, 3 = hourly, 4 = daily, 5 = weekly, 6 = monthly, 7 = quarterly, 8 = yearly
);

CREATE TABLE MRHistory (
    MRHID BIGINT NOT NULL AUTO_INCREMENT,                           -- unique id for MakeReady History
    MRStatus SMALLINT NOT NULL DEFAULT 0,                           -- see definition in Rentable table field
    DtMRStart TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,         -- when the rentable went into this status
    DtMRStop TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when the rentable changed to a different status
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (MRHID)
);

CREATE TABLE RentableStatus (
    RSID BIGINT NOT NULL AUTO_INCREMENT,                            -- unique id for Rentable Status
    RID BIGINT NOT NULL DEFAULT 0,                                  -- associated Rentable
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business
    UseStatus SMALLINT NOT NULL DEFAULT 0,                          -- 1=InService, 2=Administrative, 3=Employee, 4=OwnerOccupied, 5=OfflineRennovation, 6=OfflineMaintenance, 7=Model, 8=Inactive
    LeaseStatus SMALLINT NOT NULL DEFAULT 0,                        -- 1=Vacant-rented, 2=VacantNotRented, 3=OnNoticePreLeased, 4=OnNoticeAvailable, 5=Leased, 6=Unavailable, 7=Inactive
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',        -- start time for this state
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- stop time for this state
    DtNoticeToVacate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- user has indicated they will vacate on this date
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RSID)
);

CREATE TABLE RentableTypeRef (
    RTRID BIGINT NOT NULL AUTO_INCREMENT,                           -- unique id for Rentable Type Reference
    RID BIGINT NOT NULL DEFAULT 0,                                  -- the Rentable this record belongs to
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business
    RTID BIGINT NOT NULL DEFAULT 0,                                 -- the Rentable type for this period
    OverrideRentCycle BIGINT NOT NULL DEFAULT 0,                    -- RentCycle override. 0 = unset (use RentableType.RentCycle), > 0 means the override frequency
    OverrideProrationCycle BIGINT NOT NULL DEFAULT 0,               -- Proration override. 0 = unset (use RentableType.Proration), > 0 means the override proration
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',        -- start time for this state
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- stop time for this state
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RTRID)
);

CREATE TABLE RentableSpecialtyRef (
    RSPRefID BIGINT NOT NULL AUTO_INCREMENT,                        -- unique id for Rentable specialty Reference
    BID BIGINT NOT NULL DEFAULT 0,                                  -- the Business
    RID BIGINT NOT NULL DEFAULT 0,                                  -- unique id of unit
    RSPID BIGINT NOT NULL DEFAULT 0,                                -- unique id of specialty (see Table RentableSpecialties)
    DtStart DATE NOT NULL DEFAULT '1970-01-01 00:00:00',            -- start time for this state
    DtStop DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- stop time for this state
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RSPRefID)
);


-- **************************************
-- ****                              ****
-- ****           PEOPLE             ****
-- ****                              ****
-- **************************************
-- This is DemandSource  referenced by RentalAgreement
CREATE TABLE DemandSource (
    SourceSLSID BIGINT NOT NULL AUTO_INCREMENT,             -- DemandSource ID - unique id for this source
    BID BIGINT NOT NULL DEFAULT 0,                          -- What business is this
    Name VARCHAR(100),                                      -- Name of the source
    Industry VARCHAR(100),                                  -- What industry -- THIS BECOMES A REFERENCE TO "Industry" StringList
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (SourceSLSID)
);

CREATE TABLE LeadSource (
    LSID BIGINT NOT NULL AUTO_INCREMENT,                    -- DemandSource ID - unique id for this source
    BID BIGINT NOT NULL DEFAULT 0,                          -- What business is this
    Name VARCHAR(100),                                      -- Name of the source
    IndustrySLID BIGINT NOT NULL DEFAULT 0,                 -- What industry -- THIS BECOMES A REFERENCE TO "Industry" StringList
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (LSID)
);

-- ===========================================
--   TRANSACTANT
--   fields common to all people and businesses
-- ===========================================
CREATE TABLE Transactant (
    TCID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id of unit
    BID BIGINT NOT NULL DEFAULT 0,                          -- which business
    NLID BIGINT NOT NULL DEFAULT 0,                         -- notes associated with this transactant
    FirstName VARCHAR(100) NOT NULL DEFAULT '',
    MiddleName VARCHAR(100) NOT NULL DEFAULT '',
    LastName VARCHAR(100) NOT NULL DEFAULT '',
    PreferredName VARCHAR(100) NOT NULL DEFAULT '',
    CompanyName VARCHAR(100) NOT NULL DEFAULT '',
    IsCompany TINYINT(1) NOT NULL DEFAULT 0,                  -- 0 == this is a person,  1 == this is a company
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
    FLAGS BIGINT NOT NULL DEFAULT 0,                        /* 1<<0 OptIntoMarketingCampaign -- Does the user want to receive mkting info
                                                               1<<1 AcceptGeneralEmail       -- Will user accept email
                                                               1<<2 VIP                      -- Is this person a VIP
                                                            */
    Comment VARCHAR(2048) NOT NULL DEFAULT '',              -- company notes for this person
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (TCID)
);

--    UseCount BIGINT NOT NULL DEFAULT 0,               -- This count is incremented each time a transactant enters into a RentalAgreement.  Count > 1 means it's a ReturnUser
--    Flags BIGINT NOT NULL DEFAULT 0,                  -- For flags as described below:

-- ===========================================
--   PROSPECT
-- ===========================================

CREATE TABLE Prospect (
    TCID BIGINT NOT NULL,                                        -- associated Transactant (has Name and all contact info)
    BID BIGINT NOT NULL DEFAULT 0,                               -- which business
    CompanyAddress VARCHAR(100) NOT NULL DEFAULT '',
    CompanyCity VARCHAR(100) NOT NULL DEFAULT '',
    CompanyState VARCHAR(100) NOT NULL DEFAULT '',
    CompanyPostalCode VARCHAR(100) NOT NULL DEFAULT '',
    CompanyEmail VARCHAR(100) NOT NULL DEFAULT '',
    CompanyPhone VARCHAR(100) NOT NULL DEFAULT '',
    Occupation VARCHAR(100) NOT NULL DEFAULT '',
    DesiredUsageStartDate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',   -- User's initial indication of move in date, actual move in date is in Rental Agreement
    RentableTypePreference BIGINT NOT NULL DEFAULT 0,            -- This would be "model" preference  (Rentable Type name) for room or residence, but could apply to all rentables
    --  -----------  ----------------------------------------------------------------
    --  (FLAGS & 3)  Meaning
    --  -----------  ----------------------------------------------------------------
    --       0       Renters / Users have not completely filled out the application.
    --       1       Application has been filled out. It is being reviewed
    --       2       Application was approved by Approver on ApplicationDecisionDate
    --       3       Application was declined. Reason is in DeclineReasonSLSID
    -- ------------------------------------------------------------------------------
    FLAGS BIGINT NOT NULL DEFAULT 0,                             /* 1<<0 - 0 = application in progress, 1 application is filled out completely
                                                                    1<<1 - approved/not approved
                                                                    1<<2 - Previously Evicted: 0 = no, 1 = yes
                                                                    1<<3 - Previously Convicted of a felony: 0 = no, 1 = yes
                                                                    1<<4 - Previously declared bankruptcy: 0 = no, 1 = yes
    */
    EvictedDes VARCHAR(2048) NOT NULL DEFAULT '',                -- explanation when FLAGS & (1<<2) > 0
    ConvictedDes VARCHAR(2048) NOT NULL DEFAULT '',              -- explanation when FLAGS & (1<<3) > 0
    BankruptcyDes VARCHAR(2048) NOT NULL DEFAULT '',             -- explanation when FLAGS & (1<<4) > 0
    Approver BIGINT NOT NULL DEFAULT 0,                          -- who approved or declined
    -- ApplicationDecisionDate DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- datetime when application decision was made
    DeclineReasonSLSID BIGINT NOT NULL DEFAULT 0,                -- ID to string in list of choices, Melissa will provide the list.
    OtherPreferences VARCHAR(1024) NOT NULL DEFAULT '',          -- Arbitrary text, anything else they might request
    FollowUpDate DATE NOT NULL DEFAULT '1970-01-01 00:00:00',    -- automatically fill out this date to sysdate + 24hrs
    CSAgent BIGINT NOT NULL DEFAULT 0,                           -- Accord Directory UserID - for the CSAgent
    OutcomeSLSID BIGINT NOT NULL DEFAULT 0,                      -- id of string from a list of outcomes.
    CurrentAddress VARCHAR(200) NOT NULL DEFAULT '',             -- address of residence at the time this rental application was filled out
    CurrentLandLordName VARCHAR(100) NOT NULL DEFAULT '',        -- landlord            "
    CurrentLandLordPhoneNo VARCHAR(20) NOT NULL DEFAULT '',      -- phone number        ""
    CurrentReasonForMoving BIGINT NOT NULL DEFAULT 0,            -- string list id
    CurrentLengthOfResidency VARCHAR(100) NOT NULL DEFAULT '',   -- length of stay is just a string
    PriorAddress VARCHAR(200) NOT NULL DEFAULT '',               -- address of residence prior to "current residence"
    PriorLandLordName VARCHAR(100) NOT NULL DEFAULT '',          -- landlord            "
    PriorLandLordPhoneNo VARCHAR(20) NOT NULL DEFAULT '',        -- phone number        ""
    PriorReasonForMoving BIGINT NOT NULL DEFAULT 0,              -- string list id
    PriorLengthOfResidency VARCHAR(100) NOT NULL DEFAULT '',     -- length of stay is just a string
    CommissionableThirdParty TEXT NOT NULL DEFAULT '',           -- Sometimes bookings come into Isola Bella from 3rd parties and they get a commission
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,       -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                          -- employee UID (from phonebook) that created this record
    PRIMARY KEY (TCID)
);

-- --new  Custom Fields
-- NumberBedrooms -- SMALLINT NOT NULL DEFAULT 0,  This is unique to a room or residence. bedroom count
-- NumberOfPets   -- SMALLINT NOT NULL DEFAULT 0,  This is unique to a room or residence. may just add to formal pet schema
-- NumberOfPeople -- SMALLINT NOT NULL DEFAULT 0,  This is unique to a room or residence. count of people who will be living in the unit
-- --new
--

-- ===========================================
--   USER
-- ===========================================
CREATE TABLE User (
    TCID BIGINT NOT NULL,                                        -- associated Transactant
    BID BIGINT NOT NULL DEFAULT 0,                               -- which business
    Points BIGINT NOT NULL DEFAULT 0,                            -- bonus points for this User
    DateofBirth DATE NOT NULL DEFAULT '1970-01-01T00:00:00',
    EmergencyContactName VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactAddress VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactTelephone VARCHAR(100) NOT NULL DEFAULT '',
    EmergencyContactEmail VARCHAR(100) NOT NULL DEFAULT '',
    AlternateAddress VARCHAR(100) NOT NULL DEFAULT '',
    EligibleFutureUser TINYINT(1) NOT NULL DEFAULT 1,            -- yes/no
    FLAGS BIGINT NOT NULL DEFAULT 0,                             /*
                                                                  */
    Industry VARCHAR(100) NOT NULL DEFAULT '',                   -- (e.g., construction, retail, banking etc.)
    SourceSLSID BIGINT NOT NULL DEFAULT 0,                       -- (e.g., resident referral, newspaper, radio, post card, expedia, travelocity, etc.)
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,       -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                          -- employee UID (from phonebook) that created this record
    PRIMARY KEY (TCID)
);

-- ===========================================
--   PAYOR
-- ===========================================
CREATE TABLE Payor (
    TCID BIGINT NOT NULL,                                        -- associated Transactant
    BID BIGINT NOT NULL DEFAULT 0,                               -- which business
    TaxpayorID VARCHAR(25) NOT NULL DEFAULT '',     
    CreditLimit DECIMAL(19,4) NOT NULL DEFAULT 0.0,     
    ThirdPartySource BIGINT NOT NULL DEFAULT 0,                        -- Accord (renting company) Phonebook UID of account rep
    EligibleFuturePayor TINYINT(1) NOT NULL DEFAULT 1,           -- yes/no
    FLAGS BIGINT NOT NULL DEFAULT 0,                             /*
                                                                  */
    SSN CHAR(128) NOT NULL DEFAULT '',                           -- ssn - encrypted
    DriversLicense CHAR(128) NOT NULL DEFAULT '',                -- drivers license number - encrypted
    GrossIncome DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- gross wages
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,       -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                          -- employee UID (from phonebook) that created this record
    PRIMARY KEY (TCID)
);

CREATE TABLE Vehicle (
    VID BIGINT NOT NULL AUTO_INCREMENT,                          -- Unique identifier for vehicle
    TCID BIGINT NOT NULL DEFAULT 0,                              -- Transactant ID of vehicle owner
    BID BIGINT NOT NULL DEFAULT 0,
    VehicleType VARCHAR(80) NOT NULL DEFAULT '',
    VehicleMake VARCHAR(80) NOT NULL DEFAULT '',
    VehicleModel VARCHAR(80) NOT NULL DEFAULT '',
    VehicleColor VARCHAR(80) NOT NULL DEFAULT '',
    VehicleYear BIGINT NOT NULL DEFAULT 0,
    VIN VARCHAR(20) NOT NULL DEFAULT '',
    LicensePlateState VARCHAR(80) NOT NULL DEFAULT '',
    LicensePlateNumber VARCHAR(80) NOT NULL DEFAULT '',
    ParkingPermitNumber VARCHAR(80) NOT NULL DEFAULT '',
    DtStart DATE NOT NULL DEFAULT '1970-01-01T00:00:00',
    DtStop DATE NOT NULL DEFAULT '1970-01-01T00:00:00',
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (VID)
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
    RPASMID BIGINT NOT NULL DEFAULT 0,                      -- reversal parent Assessment, if it is non-zero, then the assessment has been reversed.
    AGRCPTID BIGINT NOT NULL DEFAULT 0,                     -- Auto-generator RCTPID is >0 when this assessment was autogenerated due to RCPTID's SubARs
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
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',              -- Accounting rule override- which acct debited, which credited
    ARID BIGINT NOT NULL DEFAULT 0,                         -- The accounting rule to apply
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- Bits 0-1:  0 = unpaid, 1 = partially paid, 2 = fully paid, 3 = not-defined at this time
                                                            -- Bit 2: 1 = this assmt has been reversed.
    Comment VARCHAR(256) NOT NULL DEFAULT '',               -- for comments such as "Prior period adjustment"
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (ASMID)
);

-- the actual tax rate or fee will be read from the TaxRate table based on the instance date of the assessment
CREATE TABLE AssessmentTax (
    ASMID BIGINT NOT NULL DEFAULT 0,                        -- the assessment to which this tax is bound
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business id
    TAXID BIGINT NOT NULL DEFAULT 0,                        -- what type of tax.
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- bit 0 = override this tax -- do not apply, bit 1 - override and use OverrideAmount
    OverrideTaxApprover MEDIUMINT NOT NULL DEFAULT 0,       -- if tax is overridden, who approved it
    OverrideAmount DECIMAL(19,4) NOT NULL DEFAULT 0,        -- Don't calculate. Use this amount. OverrideApprover required.  0 if not applicable.
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                      -- employee UID (from phonebook) that created this record
);

-- **************************************
-- ****                              ****
-- ****          EXPENSE             ****
-- ****                              ****
-- **************************************
-- charges associated with a Rentable
CREATE TABLE Expense (
    EXPID BIGINT NOT NULL AUTO_INCREMENT,                   -- unique id for expense
    RPEXPID BIGINT NOT NULL DEFAULT 0,                      -- reversal parent Expense, if it is non-zero, then the expense has been reversed.
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business id
    RID BIGINT NOT NULL DEFAULT 0,                          -- Associated rentable id
    RAID BIGINT NOT NULL DEFAULT 0,                         -- Associated Rental Agreement ID
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,              -- Expense amount
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',     -- epoch date for recurring expenses; the date/time of the expense for instances
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',              -- Accounting rule override- which acct debited, which credited
    ARID BIGINT NOT NULL DEFAULT 0,                         -- The accounting rule to apply
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- bit 2 = Reversed
    Comment VARCHAR(256) NOT NULL DEFAULT '',               -- for comments such as "Prior period adjustment"
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (EXPID)
);

-- **************************************
-- ****                              ****
-- ****     AccountRule              ****
-- ****                              ****
-- **************************************
CREATE TABLE AR (
    ARID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                          -- Business id
    Name VARCHAR(100) NOT NULL DEFAULT '',
    SubARID BIGINT NOT NULL DEFAULT 0,                      --
    ARType SMALLINT NOT NULL DEFAULT 0,                     -- Assessment = 0, Receipt = 1, Expense = 2
    RARequired SMALLINT NOT NULL DEFAULT 0,                 -- 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
    DebitLID BIGINT NOT NULL DEFAULT 0,                     -- Ledger ID of debit part
    CreditLID BIGINT NOT NULL DEFAULT 0,                    -- Ledger ID of crdit part
    Description VARCHAR(1024) NOT NULL DEFAULT '',
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',-- epoch date for recurring assessments; the date/time of the assessment for instances
    DtStop DATETIME NOT NULL DEFAULT '9999-12-31 00:00:00', -- stop date for recurrent assessments; the date/time of the assessment for instances
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- 1<<0 = apply funds to Receive accts,  (that is allocate it immediately)
                                                            -- 1<<1 - Auto Populate on New Rental Agreement,
                                                            -- 1<<2 = RAID required,
                                                            -- 1<<3 = subARIDs apply (i.e., there are other ar rules that apply to this AR Rule)
                                                            -- 1<<4 = Is Rent Assessment
                                                            -- 1<<5 = Is Security Deposit Assessment
                                                            -- 1<<6 = Is NonRecur charge
    DefaultAmount DECIMAL(19,4) NOT NULL DEFAULT 0.0,       -- amount to initialize interface with
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY(ARID)
);

-- **************************************
-- ****                              ****
-- ****     SubAccountRule           ****
-- ****                              ****
-- **************************************
CREATE TABLE SubAR (
    SARID BIGINT NOT NULL AUTO_INCREMENT,
    ARID BIGINT NOT NULL DEFAULT 0,                         -- Which ARID
    SubARID BIGINT NOT NULL DEFAULT 0,                      -- SubARID
    BID BIGINT NOT NULL DEFAULT 0,
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY(SARID)
);

-- **************************************
-- ****                              ****
-- ****           RECEIPTS           ****
-- ****                              ****
-- **************************************
CREATE TABLE Receipt (
    RCPTID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this Receipt
    PRCPTID BIGINT NOT NULL DEFAULT 0,                          -- Parent RCPT, if non-zero then it is the RCPTID of a receipt being reversed
    BID BIGINT NOT NULL DEFAULT 0,
    TCID BIGINT NOT NULL DEFAULT 0,                             -- Payor, even if OtherPayorName is present this field must have the payor for whom the OtherPayorName is paying
    PMTID BIGINT NOT NULL DEFAULT 0,
    DEPID BIGINT NOT NULL DEFAULT 0,                            -- Depository for this payment
    DID BIGINT NOT NULL DEFAULT 0,                              -- Deposit id to which this receipt belongs
    RAID BIGINT NOT NULL DEFAULT 0,                             -- RAID - needed for special case receipts
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DocNo VARCHAR(50) NOT NULL DEFAULT '',                      -- Check Number, MoneyOrder number, etc., the traceback for the payment
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    AcctRuleReceive VARCHAR(215) NOT NULL DEFAULT '',           --
    ARID BIGINT NOT NULL DEFAULT 0,                             -- identifies the account rule used on Receipt
    AcctRuleApply VARCHAR(4096) NOT NULL DEFAULT '',            -- How the funds will be applied
    FLAGS BIGINT NOT NULL DEFAULT 0,                            /* bits 0-1 : 0 unallocated, 1 = partially allocated, 2 = fully allocated,
                                                                 * bit 2 = VOID THIS RECEIPT
                                                                 */
    Comment VARCHAR(256) NOT NULL DEFAULT '',                   -- for comments like "Prior Period Adjustment"
    OtherPayorName VARCHAR(128) NOT NULL DEFAULT '',            -- If not '' then Payment was made by a payor who is not on the RA, and may not be in our system at all
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RCPTID)
);

CREATE TABLE ReceiptAllocation (
    RCPAID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this allocation
    RCPTID BIGINT NOT NULL DEFAULT 0,                           -- sum of all amounts in this table with RCPTID must equal the Receipt with RCPTID in Receipt table
    BID BIGINT NOT NULL DEFAULT 0,
    RAID BIGINT NOT NULL DEFAULT 0,
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    ASMID BIGINT NOT NULL DEFAULT 0,                            -- the id of the assessment that caused this payment
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- bit 2:  VOID THIS RECEIPT-ALLOCATION
    AcctRule VARCHAR(150) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (RCPAID)
);

-- **************************************
-- ****                              ****
-- ****          DEPOSIT             ****
-- ****                              ****
-- **************************************
CREATE TABLE DepositMethod (
    DPMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                              -- which business
    Method VARCHAR(50) NOT NULL DEFAULT '',                     -- 0 = not specified, 1 = Hand Delivery, Scanned Batch, US Mail
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (DPMID)
);

CREATE TABLE Depository (
    DEPID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for a depository
    BID BIGINT NOT NULL DEFAULT 0,                              -- business id
    LID BIGINT NOT NULL DEFAULT 0,                              -- the GL account that represents this depository
    Name VARCHAR(256),                                          -- Name of Depository: First Data, Nyax, Oklahoma Fidelity
    AccountNo VARCHAR(256),                                     -- account number at this Depository
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (DEPID)
);

CREATE TABLE Deposit (
    DID BIGINT NOT NULL AUTO_INCREMENT,                         -- UniqueID for this deposit
    BID BIGINT NOT NULL DEFAULT 0,                              -- business id
    DEPID BIGINT NOT NULL DEFAULT 0,                            -- DepositoryID where the Deposit was made
    DPMID BIGINT NOT NULL DEFAULT 0,                            -- Deposit Method
    Dt DATE NOT NULL DEFAULT '1970-01-01 00:00:00',             -- Date of deposit
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                  -- total amount of all Receipts in this deposit
    ClearedAmount DECIMAL(19,4) NOT NULL DEFAULT 0.0,           -- Amount cleared by the bank
    FLAGS BIGINT NOT NULL DEFAULT 0,                            -- bitflags
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (DID)
);

CREATE TABLE DepositPart (
    DPID BIGINT NOT NULL AUTO_INCREMENT,
    DID BIGINT NOT NULL DEFAULT 0,
    BID BIGINT NOT NULL DEFAULT 0,                              -- business id
    RCPTID BIGINT NOT NULL DEFAULT 0,                           -- the receipt
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (DPID)
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
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (InvoiceNo)
);

CREATE TABLE InvoiceAssessment (
    InvoiceASMID BIGINT NOT NULL AUTO_INCREMENT,                -- Unique id for this invoice Assessment
    InvoiceNo BIGINT NOT NULL DEFAULT 0,                        -- which invoice
    BID BIGINT NOT NULL DEFAULT 0,                              -- bid
    ASMID BIGINT NOT NULL DEFAULT 0,                            -- assessment id
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (InvoiceASMID)
);

CREATE TABLE InvoicePayor (
    InvoicePayorID BIGINT NOT NULL AUTO_INCREMENT,              -- Unique id for this invoice Payor
    InvoiceNo BIGINT NOT NULL DEFAULT 0,                        -- which invoice
    BID BIGINT NOT NULL DEFAULT 0,                              -- bid
    PID BIGINT NOT NULL DEFAULT 0,                              -- Payor id
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (InvoicePayorID)
);


-- **************************************
-- ****                              ****
-- ****           JOURNAL            ****
-- ****                              ****
-- **************************************
CREATE TABLE Journal (
    JID BIGINT NOT NULL AUTO_INCREMENT,                             -- a Journal entry
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business id
    -- RAID BIGINT NOT NULL DEFAULT 0,                                 -- associated rental agreement
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',             -- date when it occurred
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                      -- how much
    Type SMALLINT NOT NULL DEFAULT 0,                               -- 0 = unassociated with RA, 1 = assessment, 2 = payment/Receipt, 3 = Expense
    ID BIGINT NOT NULL DEFAULT 0,                                   -- if Type == 0 then it is the RentableID,
                                                                    -- if Type == 1 then it is the ASMID that caused this entry,
                                                                    -- if Type == 2 then it is the RCPTID
                                                                    -- if Type == 3 then it is the EXPID
    Comment VARCHAR(256) NOT NULL DEFAULT '',                       -- for notes like "prior period adjustment"
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (JID)
);

CREATE TABLE JournalAllocation (
    JAID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business id
    JID BIGINT NOT NULL DEFAULT 0,                                  -- sum of all amounts in this table with RCPTID must equal the Receipt with RCPTID in Receipt table
    RID BIGINT NOT NULL DEFAULT 0,                                  -- associated Rentable
    RAID BIGINT NOT NULL DEFAULT 0,                                 -- associated Rental Agreement
    TCID BIGINT NOT NULL DEFAULT 0,                                 -- if > 0 this is the payor who made the payment - important if RID and RAID == 0 -- means it's unallocated funds
    RCPTID BIGINT NOT NULL DEFAULT 0,                               -- associated receipt if TCID > 0. If both ASMID and RCPTID are > 0, then the assessment was generated as
                                                                    -- as part of a SubAR rule that binds the two
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                      -- Amount transacted
    ASMID BIGINT NOT NULL DEFAULT 0,                                -- may not be present if assessment records have been backed up and removed.
    EXPID BIGINT NOT NULL DEFAULT 0,                                -- the associated expense.
    AcctRule VARCHAR(200) NOT NULL DEFAULT '',
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
    PRIMARY KEY (JAID)
);

CREATE TABLE JournalMarker (
    JMID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                                 -- Business id
    State SMALLINT NOT NULL DEFAULT 0,                             -- 0 = unknown, 1 = Closed, 2 = Locked
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                           -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,         -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that created this record
    PRIMARY KEY (JMID)
);

CREATE TABLE JournalAudit (
    JID BIGINT NOT NULL DEFAULT 0,                                  -- what JID was affected
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business id
    UID MEDIUMINT NOT NULL DEFAULT 0,                               -- UID of person making the change
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                              -- employee UID (from phonebook) that created this record
);

CREATE TABLE JournalMarkerAudit (
    JMID BIGINT NOT NULL DEFAULT 0,                                 -- what JMID was affected
    BID BIGINT NOT NULL DEFAULT 0,                                  -- Business id
    UID MEDIUMINT NOT NULL DEFAULT 0,                               -- UID of person making the change
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                              -- employee UID (from phonebook) that created this record
);

-- **************************************
-- ****                              ****
-- ****           LEDGERS            ****
-- ****                              ****
-- **************************************
CREATE TABLE LedgerEntry (
    LEID BIGINT NOT NULL AUTO_INCREMENT,                      -- unique id for this LedgerEntry
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    JID BIGINT NOT NULL DEFAULT 0,                            -- Journal entry giving rise to this
    JAID BIGINT NOT NULL DEFAULT 0,                           -- the allocation giving rise to this LedgerEntry
    LID BIGINT NOT NULL DEFAULT 0,                            -- associated GLAccount
    RAID BIGINT NOT NULL DEFAULT 0,                           -- associated Rental Agreement
    RID BIGINT NOT NULL DEFAULT 0,                            -- associated Rentable
    TCID BIGINT NOT NULL DEFAULT 0,                           -- Payor associated with this entry
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- balance date and time
    Amount DECIMAL(19,4) NOT NULL DEFAULT 0.0,                -- balance amount since last close
    Comment VARCHAR(256) NOT NULL DEFAULT '',                 -- for notes like "prior period adjustment"
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (LEID)
);

CREATE TABLE LedgerMarker (
    LMID BIGINT NOT NULL AUTO_INCREMENT,
    LID BIGINT NOT NULL DEFAULT 0,                            -- associated GLAccount
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                           -- 0 means it's either a marker for a Rentable or the balance for the whole account;  > 0 AND LID=0 means it's the amount associated with rental agreement RAID
    RID BIGINT NOT NULL DEFAULT 0,                            -- 0 means it's either a marker for a RentalAgreement or the balance for a whole account; >0 means it's the amount associated with Rentable RID
    TCID BIGINT NOT NULL DEFAULT 0,                           -- if 0 then LM for whole acct, if > 0 then it's the amount for this payor; TCID
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',       -- Balance is valid as of this time
    Balance DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    State SMALLINT NOT NULL DEFAULT 0,                        -- 0 = Open, 1 = Closed, 2 = Locked, 3 = InitialMarker (no records prior)
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (LMID)
);

-- GL Account
CREATE TABLE GLAccount (
    LID BIGINT NOT NULL AUTO_INCREMENT,                       -- unique id for this GLAccount
    PLID BIGINT NOT NULL DEFAULT 0,                           -- Parent ID for this GLAccount.  0 if no parent.
    BID BIGINT NOT NULL DEFAULT 0,                            -- Business id
    RAID BIGINT NOT NULL DEFAULT 0,                           -- rental agreement account, only valid if TYPE is 1
    TCID BIGINT NOT NULL DEFAULT 0,                           -- Payor, only valid if TYPE is 2
    GLNumber VARCHAR(100) NOT NULL DEFAULT '',                -- if not '' then it's a link a QB  GeneralLedger (GL)account
    -- Status SMALLINT NOT NULL DEFAULT 0,                       -- Whether a GL Account is currently unknown=0, inactive=1, active=2
    Name VARCHAR(100) NOT NULL DEFAULT '',
    AcctType VARCHAR(100) NOT NULL DEFAULT '',                -- Quickbooks Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable,
                                                              --    Other Current Asset, Other Asset, Accounts Payable, Other Current Liability,
                                                              --    Cost of Goods Sold, Other Income, Other Expense
    AllowPost TINYINT(1) NOT NULL DEFAULT 0,                  -- 0 - do not allow posts to this ledger. 1 = allow posts
    -- RARequired SMALLINT NOT NULL DEFAULT 0,                -- 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
    FLAGS BIGINT NOT NULL DEFAULT 0,                          -- 1<<0 - inactive:  0 = active, 1 = inactive  (this replaces the old Status field)
    Description VARCHAR(1024) NOT NULL DEFAULT '',            -- describe the assessment
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                      -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                       -- employee UID (from phonebook) that created this record
    PRIMARY KEY (LID)
);


CREATE TABLE LedgerAudit (
    LEID BIGINT NOT NULL DEFAULT 0,                             -- what LEID was affected
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business id
    UID MEDIUMINT NOT NULL DEFAULT 0,                           -- UID of person making the change
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that created this record
);

CREATE TABLE LedgerMarkerAudit (
    LMID BIGINT NOT NULL DEFAULT 0,                             -- what LMID was affected
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business id
    UID MEDIUMINT NOT NULL DEFAULT 0,                           -- UID of person making the change
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0                          -- employee UID (from phonebook) that created this record
);

CREATE TABLE ClosePeriod (
    CPID BIGINT NOT NULL AUTO_INCREMENT,                        -- Close Period ID
    BID BIGINT NOT NULL DEFAULT 0,                              -- Business id
    TLID BIGINT NOT NULL DEFAULT 0,                             -- Task List that was used for close
    Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- Date/Time of close
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
    PRIMARY KEY (CPID)
);


-- **************************************
-- ****                              ****
-- ****            FLOW              ****
-- ****                              ****
-- **************************************
CREATE TABLE Flow (
    FlowID BIGINT NOT NULL AUTO_INCREMENT,
    BID BIGINT NOT NULL DEFAULT 0,                                                         -- Business id
    UserRefNo VARCHAR(50) NOT NULL DEFAULT '',                                             -- reference id to share with the user(s)
    FlowType VARCHAR(50) NOT NULL DEFAULT '',                                              -- for which flow we're storing data ("RA=Rental Agreement Flow")
    ID BIGINT NOT NULL DEFAULT 0,                                                          -- ID associated with flow type, typically a permanent table ID, RAID for flow "RA"
    Data JSON DEFAULT NULL,                                                                -- JSON Data for each flow type
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was it last updated
    LastModBy BIGINT NOT NULL DEFAULT 0,                                                   -- who modified it last
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,                                 -- when was it created
    CreateBy BIGINT NOT NULL DEFAULT 0,                                                    -- who created it
    PRIMARY KEY(FlowID)
);
