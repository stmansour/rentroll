#!/bin/bash
MODFILE="dbqqqmods.sql"
# Last updated on Aug 3, 2018
cat > ${MODFILE} <<EOF
RENAME TABLE RentalAgreementPets TO Pets;
CREATE TABLE TBind (
    TBID BIGINT NOT NULL AUTO_INCREMENT,                    -- unique id
    SourceElemType BIGINT NOT NULL DEFAULT 0,               -- Source element type, example: 14 = Pet, 15 = Vehicle. Values defined in dbtypes.go
    SourceElemID BIGINT NOT NULL DEFAULT 0,                 -- ID of the Source Element for the Associated Element.  Ex. if SourceElemType = 14, then SourceElemID is the PETID
    AssocElemType BIGINT NOT NULL DEFAULT 0,                -- Associated element type, example: 14 = Pet, 15 = Vehicle. Values defined in dbtypes.go
    AssocElemID BIGINT NOT NULL DEFAULT 0,                  -- ID for the Associated Element.  Ex. if AssocElemType = 14, then AssocElemID is the PETID
    DtStart DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',-- epoch date for recurring assessments; the date/time of the assessment for instances
    DtStop DATETIME NOT NULL DEFAULT '2066-01-01 00:00:00', -- stop date for recurrent assessments; the date/time of the assessment for instances
    FLAGS BIGINT NOT NULL DEFAULT 0,                        -- Bits 0-1:  0 = unpaid, 1 = partially paid, 2 = fully paid, 3 = not-defined at this time
    LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
    LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
    CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- when was this record created
    CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
    PRIMARY KEY (TBID)
);
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
