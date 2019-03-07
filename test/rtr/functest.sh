#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableTypeRefs"
TESTSUMMARY="Test Reservation search, create, mod"
DBGENDIR=${SRCTOP}/tools/dbgen
CREATENEWDB=0
RRBIN="${SRCTOP}/tmp/rentroll"
CATRML="${SRCTOP}/tools/catrml/catrml"

#SINGLETEST=""  # This runs all the tests

source ${TESTHOME}/share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
# RENTROLLSERVERNOW="-testDtNow 10/24/2018"

#------------------------------------------------------------------------------
#  TEST a
#
#  Validate a save.
#
#  Scenario:
#  Change the end date of an RTR but provide erroneous data. Make sure we
#  get error message response.
#
#
#  Expected Results:
#   1.  Error message response
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A0%2C%22RLID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22LeaseStatus%22%3A0%2C%22DtStart%22%3A%221%2F1%2F2019%22%2C%22DtStop%22%3A%221%2F3%2F2019%22%2C%22Comment%22%3A%22%22%2C%22CreateBy%22%3A0%2C%22LastModBy%22%3A0%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}0"  "RentableTypeRefs-SaveWithError"
fi

#------------------------------------------------------------------------------
#  TEST b
#
#  Validate a save with correct data.
#
#  Scenario:
#  Change the end date of an RTR.
#
#
#  Expected Results:
#   1.  the end date of RTID 1 should change to 2/28/2019
#------------------------------------------------------------------------------
TFILES="b"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A0%2C%22RTRID%22%3A1%2C%22RTID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22OverrideRentCycle%22%3A0%2C%22OverrideProrationCycle%22%3A0%2C%22DtStart%22%3A%221%2F1%2F2019%22%2C%22DtStop%22%3A%222%2F28%2F2019%22%2C%22CreateBy%22%3A0%2C%22LastModBy%22%3A0%2C%22w2ui%22%3A%7B%7D%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}1"  "RentableTypeRefs-Save"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
