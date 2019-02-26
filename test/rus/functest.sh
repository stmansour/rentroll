#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableUsseStatus"
TESTSUMMARY="Test Rentable Use Status code"
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
#  Validate that the search query returns the proper data
#
#  Scenario:
#  
#  The database in xa should have many use status fields.   Triplets of
#  the form housekeeping, ready, in-service.
#
#
#  Expected Results:
#   1.  Search for UseStatus on rentable 4 and make sure that the data
#       returned matches the patterns we expect.
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/4" "request" "${TFILES}0"  "RentableUseStatus-Search"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
