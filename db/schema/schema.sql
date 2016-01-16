-- *************************
-- *********  IDS  *********
-- *************************
--    PRID = property id
--    UTID = unit type id
--   USPID = unit specialty id
--   OFSID = offset id
--   ASMID = assessment id
--   PMTID = payment type id
-- AVAILID = availability id
--  BLDGID = building id
--  UNITID = unit id
--    TCID = transactant id


-- **************************************
-- ****                              ****
-- ****          PROPERTY            ****
-- ****                              ****
-- **************************************
-- Unit Types - associated with a property are stored
-- in the unittypes table and have the property PID
-- Occupancy Type List - hardcoded
CREATE TABLE property (
    PRID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Address VARCHAR(35) NOT NULL DEFAULT '',
    Address2 VARCHAR(35) NOT NULL DEFAULT '',
    City VARCHAR(25) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(10) NOT NULL DEFAULT '',
    Country VARCHAR(25) NOT NULL DEFAULT '',
    Name VARCHAR(50) NOT NULL DEFAULT '',
    OccupancyType SMALLINT NOT NULL DEFAULT 0,
    ParkingPermitInUse SMALLINT NOT NULL DEFAULT 0,
    MOAFID MEDIUMINT NOT NULL DEFAULT 0,
    LastModTime TIMESTAMP,
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,
    PRIMARY KEY (PRID)
);

CREATE TABLE unittypes (
    UTID MEDIUMINT NOT NULL,
    PRID MEDIUMINT NOT NULL,                            -- associated property id
    Designation CHAR(5) NOT NULL DEFAULT '',
    Description VARCHAR(256) NOT NULL DEFAULT '',
    PRIMARY KEY (UTID)
);

-- a collection of unit specialties that are available
-- for a specifi property
CREATE TABLE unitspecialties (
    USPID MEDIUMINT NOT NULL AUTO_INCREMENT,
    PID MEDIUMINT NOT NULL,
    Name VARCHAR(25) NOT NULL DEFAULT '',
    Fee DECIMAL(19,4) NOT NULL DEFAULT 0.0,
    Description VARCHAR(256) NOT NULL DEFAULT '',
    PRIMARY KEY (USPID)

);

CREATE TABLE offsets (
    OFSID MEDIUMINT NOT NULL,
    Description VARCHAR(256) NOT NULL DEFAULT ''
);

CREATE TABLE assessments (
    ASMID MEDIUMINT NOT NULL,
    Description VARCHAR(256) NOT NULL DEFAULT ''   
);

CREATE TABLE paymenttypes (
    PMTID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Description VARCHAR(256) NOT NULL DEFAULT ''
    PRIMARY KEY (PMTID)
);


-- examples: Occupied, Offline, Administrative, Vacant - Not Ready, 
--           Vacant - Made Ready, Vacant - Inspected, plus custom values
-- custom values will be added with their own uniq AVAILID
--
CREATE TABLE availabilitytype (
    AVAILID MEDIUMINT NOT NULL AUTO_INCREMENT,
    Description VARCHAR(35) NOT NULL DEFAULT '',
    PRIMARY KEY (AVAILID)
);

-- **************************************
-- ****                              ****
-- ****          BUILDING            ****
-- ****                              ****
-- **************************************
CREATE TABLE buildings (
    BLDGID MEDIUMINT NOT NULL AUTO_INCREMENT,           -- unique id for this building
    PRID MEDIUMINT NOT NULL DEFAULT 0,                  -- which property it belongs to
    Address VARCHAR(35) NOT NULL DEFAULT '',            -- building address
    Address2 VARCHAR(35) NOT NULL DEFAULT '',       
    City VARCHAR(25) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(10) NOT NULL DEFAULT '',
    Country VARCHAR(25) NOT NULL DEFAULT ''
);

-- **************************************
-- ****                              ****
-- ****           UNITS              ****
-- ****                              ****
-- **************************************
CREATE TABLE units (
    UNITID MEDIUMINT NOT NULL AUTO_INCREMENT,           -- unique id for this unit
    PRID MEDIUMINT NOT NULL DEFAULT 0,                  -- which property it belongs to
    BLDGID MEDIUMINT NOT NULL DEFAULT 0,                -- which building
    UTID MEDIUMINT NOT NULL DEFAULT 0,                  -- which unit type             
    -- Abbreviation VARCHAR(20),                           -- unit abbreviation  -- REMOVED - it's part of unittype
    DefaultOccType SMALLINT NOT NULL DEFAULT 0,         -- unset, short term, longterm
    ScheduledRent DECIMAL(19,4) NOT NULL DEFAULT 0.0,   -- budgeted rent for this unit
    Assignment SMALLINT NOT NULL DEFAULT 0,             -- Pre-assign or assign at occupy commencement
    Report SMALLINT NOT NULL DEFAULT 1,                 -- 1 = apply to rentroll, 0 = skip on rentroll
);

-- For each unit, what specialties does it have...
-- this is simply a list of USPIDs.  So selecting all entries where the unit == UNITID
-- will be the list of all the unit specialties.
CREATE TABLE unitspecialties (
    UNITID MEDIUMINT NOT NULL,                          -- unique id of unit
    USPID MEDIUMINT NOT NULL,                           -- unique id of specialty
);

-- **************************************
-- ****                              ****
-- ****           PEOPLE             ****
-- ****                              ****
-- **************************************

-- transactant - fields common to all people
CREATE TABLE transactant (
    TCID MEDIUMINT NOT NULL AUTO_INCREMENT,              -- unique id of unit
    FirstName VARCHAR(35) NOT NULL DEFAULT '',
    MiddleName VARCHAR(35) NOT NULL DEFAULT '',
    LastName VARCHAR(35) NOT NULL DEFAULT '', 
    Address VARCHAR(35) NOT NULL DEFAULT '',            -- person address
    Address2 VARCHAR(35) NOT NULL DEFAULT '',       
    City VARCHAR(25) NOT NULL DEFAULT '',
    State CHAR(25) NOT NULL DEFAULT '',
    PostalCode VARCHAR(10) NOT NULL DEFAULT '',
    Country VARCHAR(25) NOT NULL DEFAULT ''
    LastModTime TIMESTAMP,
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,
    PRIMARY KEY (TCID)
);

CREATE TABLE prospects (
);

CREATE TABLE tenant (
    TID MEDIUMINT NOT NULL AUTO_INCREMENT,              -- unique id of this tenant
    TCID MEDIUMINT NOT NULL,                             -- associated transactant
    CarColor VARCHAR(25) NOT NULL DEFAULT '',
    CarYear MEDIUMINT NOT NULL DEFAULT 0,
    AccountRep VARCHAR(35) NOT NULL DEFAULT '',
    LicensePlateState VARCHAR(35) NOT NULL DEFAULT '',
    LicensePlateNumber VARCHAR(35) NOT NULL DEFAULT '',
    ParkingPermitNumber VARCHAR(35) NOT NULL DEFAULT '',
    DateofBirth DATE NOT NULL DEFAULT '2000-01-01 00:00:00',
    EmergencyContactName VARCHAR(35) NOT NULL DEFAULT '',
    EmergencyContactAddress VARCHAR(35) NOT NULL DEFAULT '',
    EmergencyContactTelephone VARCHAR(35) NOT NULL DEFAULT '',
    EmergencyAddressEmail VARCHAR(35) NOT NULL DEFAULT '',
    AlternateAddress VARCHAR(35) NOT NULL DEFAULT '',
    ElibigleForFutureOccupancy SMALLINT NOT NULL DEFAULT 1,
    Industry VARCHAR(35) NOT NULL DEFAULT '',
    Source  VARCHAR(35) NOT NULL DEFAULT '',
    InvoicingCustomerNumber VARCHAR(35) NOT NULL DEFAULT ''
    LastModTime TIMESTAMP,
    LastModBy MEDIUMINT NOT NULL DEFAULT 0,
    PRIMARY KEY (TID)
);





CREATE TABLE payor  (
);



-- Add the Administrator as the first and only user
-- INSERT INTO people (UserName,FirstName,LastName) VALUES("administrator","Administrator","Administrator");
