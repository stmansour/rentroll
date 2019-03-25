#!/bin/bash
MODFILE="dbqqqmods.sql"
# Last updated on Mar 22, 2019
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
