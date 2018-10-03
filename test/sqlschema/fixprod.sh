#!/bin/bash
MODFILE="dbqqqmods.sql"
# Last updated on Aug 3, 2018
cat > ${MODFILE} <<EOF
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
