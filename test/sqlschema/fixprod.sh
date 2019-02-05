#!/bin/bash
MODFILE="dbqqqmods.sql"

cat > ${MODFILE} <<EOF
ALTER TABLE RentableLeaseStatus ADD FirstName VARCHAR(50) NOT NULL DEFAULT '' AFTER Comment;
ALTER TABLE RentableLeaseStatus ADD LastName VARCHAR(50) NOT NULL DEFAULT '' AFTER FirstName;
ALTER TABLE RentableLeaseStatus ADD Email VARCHAR(100) NOT NULL DEFAULT '' AFTER LastName;
ALTER TABLE RentableLeaseStatus ADD Phone VARCHAR(100) NOT NULL DEFAULT '' AFTER Email;
ALTER TABLE RentableLeaseStatus ADD Address VARCHAR(100) NOT NULL DEFAULT '' AFTER Phone;
ALTER TABLE RentableLeaseStatus ADD Address2 VARCHAR(100) NOT NULL DEFAULT '' AFTER Address;
ALTER TABLE RentableLeaseStatus ADD City VARCHAR(100) NOT NULL DEFAULT '' AFTER Address2;
ALTER TABLE RentableLeaseStatus ADD State CHAR(25) NOT NULL DEFAULT '' AFTER City;
ALTER TABLE RentableLeaseStatus ADD PostalCode VARCHAR(100) NOT NULL DEFAULT '' AFTER State;
ALTER TABLE RentableLeaseStatus ADD Country VARCHAR(100) NOT NULL DEFAULT '' AFTER PostalCode;
ALTER TABLE RentableLeaseStatus ADD CCName VARCHAR(100) NOT NULL DEFAULT '' AFTER Country;
ALTER TABLE RentableLeaseStatus ADD CCType VARCHAR(100) NOT NULL DEFAULT '' AFTER CCName;
ALTER TABLE RentableLeaseStatus ADD CCNumber VARCHAR(100) NOT NULL DEFAULT '' AFTER CCType;
ALTER TABLE RentableLeaseStatus ADD CCExpMonth VARCHAR(100) NOT NULL DEFAULT '' AFTER CCNumber;
ALTER TABLE RentableLeaseStatus ADD CCExpYear VARCHAR(100) NOT NULL DEFAULT '' AFTER CCExpMonth;

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
