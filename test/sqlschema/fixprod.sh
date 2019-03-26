#!/bin/bash
MODFILE="dbqqqmods.sql"

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

cat > ${MODFILE} <<EOF

# March 18, 2019
# CREATE TABLE RentableUseType (
#     UTID BIGINT NOT NULL AUTO_INCREMENT,                            -- unique id for Rentable Use Type
#     RID BIGINT NOT NULL DEFAULT 0,                                  -- associated Rentable
#     BID BIGINT NOT NULL DEFAULT 0,                                  -- Business
#     UseType SMALLINT NOT NULL DEFAULT 0,                            -- 100 = Standard, 101=Administrative, 102=Employee, 103=OwnerOccupied
#     DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',        -- start time for this state
#     DtStop DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- stop time for this state
#     Comment VARCHAR(2048) NOT NULL DEFAULT '',                      -- company notes for this person
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                            -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,          -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                             -- employee UID (from phonebook) that created this record
#     PRIMARY KEY (UTID)
# );
EOF

#==========================================================================
#  This script tests the mods placed in dbmod.sh and verifies that they
#  can be applied, en-masse, to the production schema and that the result
#  passes the schema check.
#==========================================================================

MYSQL="mysql --no-defaults"
MYSQLDUMP="mysqldump --no-defaults"
DBNAME="rentroll"
FIXED="fixrr.sql"

f="rr.sql"

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
