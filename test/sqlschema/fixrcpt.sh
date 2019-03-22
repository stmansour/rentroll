#!/bin/bash
MODFILE="dbqqqmods.sql"
# Last updated on Mar 22, 2019
cat > ${MODFILE} <<EOF
# RENAME TABLE RentalAgreementPets TO Pets;
# CREATE TABLE TBind (
#     TBID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id
#     SourceElemType BIGINT NOT NULL DEFAULT 0,               -- Source element type, example: 14 = Pet, 15 = Vehicle. Values defined in dbtypes.go
#     SourceElemID BIGINT NOT NULL DEFAULT 0,                 -- ID of the Source Element for the Associated Element.  Ex. if SourceElemType = 14, then SourceElemID is the PETID
#     AssocElemType BIGINT NOT NULL DEFAULT 0,                -- Associated element type, example: 14 = Pet, 15 = Vehicle. Values defined in dbtypes.go
#     AssocElemID BIGINT NOT NULL DEFAULT 0,                  -- ID for the Associated Element.  Ex. if AssocElemType = 14, then AssocElemID is the PETID
#     DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',-- epoch date for recurring assessments; the date/time of the assessment for instances
#     DtStop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00', -- stop date for recurrent assessments; the date/time of the assessment for instances
#     FLAGS BIGINT NOT NULL DEFAULT 0,                        -- Bits 0-1:  0 = unpaid, 1 = partially paid, 2 = fully paid, 3 = not-defined at this time
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
#     PRIMARY KEY (TBID)
# );
# ALTER TABLE TBind ADD COLUMN BID BIGINT NOT NULL DEFAULT 0 AFTER TBID;
#
# RENAME TABLE RentableStatus TO RentableUseStatus;
# ALTER TABLE RentableUseStatus ADD Comment VARCHAR(2048) NOT NULL DEFAULT '' AFTER DtStop;
# ALTER TABLE RentableUseStatus DROP Column DtNoticeToVacate;
# CREATE TABLE RentableLeaseStatus (
#     RLID BIGINT NOT NULL AUTO_INCREMENT,                            -- unique id for Rentable Status
#     RID BIGINT NOT NULL DEFAULT 0,                                  -- associated Rentable
#     BID BIGINT NOT NULL DEFAULT 0,                                  -- Business
#     LeaseStatus SMALLINT NOT NULL DEFAULT 0,                        -- 0 = Not Leased, 1 = Leased, 2 = Reserved
#     DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',        -- start time for this state
#     DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- stop time for this state
#     Comment VARCHAR(2048) NOT NULL DEFAULT '',                      -- company notes for this person
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
#     PRIMARY KEY (RLID)
# );
# ALTER TABLE RentableUseStatus DROP LeaseStatus;
# ALTER TABLE RentableLeaseStatus ADD FirstName VARCHAR(50) NOT NULL DEFAULT '' AFTER Comment;
# ALTER TABLE RentableLeaseStatus ADD LastName VARCHAR(50) NOT NULL DEFAULT '' AFTER FirstName;
# ALTER TABLE RentableLeaseStatus ADD Email VARCHAR(100) NOT NULL DEFAULT '' AFTER LastName;
# ALTER TABLE RentableLeaseStatus ADD Phone VARCHAR(100) NOT NULL DEFAULT '' AFTER Email;
# ALTER TABLE RentableLeaseStatus ADD Address VARCHAR(100) NOT NULL DEFAULT '' AFTER Phone;
# ALTER TABLE RentableLeaseStatus ADD Address2 VARCHAR(100) NOT NULL DEFAULT '' AFTER Address;
# ALTER TABLE RentableLeaseStatus ADD City VARCHAR(100) NOT NULL DEFAULT '' AFTER Address2;
# ALTER TABLE RentableLeaseStatus ADD State CHAR(25) NOT NULL DEFAULT '' AFTER City;
# ALTER TABLE RentableLeaseStatus ADD PostalCode VARCHAR(100) NOT NULL DEFAULT '' AFTER State;
# ALTER TABLE RentableLeaseStatus ADD Country VARCHAR(100) NOT NULL DEFAULT '' AFTER PostalCode;
# ALTER TABLE RentableLeaseStatus ADD CCName VARCHAR(100) NOT NULL DEFAULT '' AFTER Country;
# ALTER TABLE RentableLeaseStatus ADD CCType VARCHAR(100) NOT NULL DEFAULT '' AFTER CCName;
# ALTER TABLE RentableLeaseStatus ADD CCNumber VARCHAR(100) NOT NULL DEFAULT '' AFTER CCType;
# ALTER TABLE RentableLeaseStatus ADD CCExpMonth VARCHAR(100) NOT NULL DEFAULT '' AFTER CCNumber;
# ALTER TABLE RentableLeaseStatus ADD CCExpYear VARCHAR(100) NOT NULL DEFAULT '' AFTER CCExpMonth;


EOF

#==========================================================================
#  This script tests the mods placed in dbmod.sh and verifies that they
#  can be applied, en-masse, to the production schema and that the result
#  passes the schema check.
#==========================================================================

MYSQL="mysql --no-defaults"
MYSQLDUMP="mysqldump --no-defaults"
DBNAME="receipts"
FIXED="fixrcpts.sql"

f="receipts.sql"

rm -f ${FIXED}
echo "DROP DATABASE IF EXISTS ${DBNAME}; create database ${DBNAME}" | ${MYSQL}
echo -n "${f}: loading... "
${MYSQL} ${DBNAME} < ${f}
if [ $? -ne 0 ]; then
    exit 2
fi

echo -n "updating... "
${MYSQL} ${DBNAME} < ${MODFILE}
if [ $? -ne 0 ]; then
    exit 2
fi

echo -n "saving... "
${MYSQLDUMP} ${DBNAME} > ${FIXED}
if [ $? -ne 0 ]; then
    exit 2
fi

echo "done"
