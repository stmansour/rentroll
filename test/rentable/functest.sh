#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="Rentables"
TESTSUMMARY="Test Rentables"
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
#  Validate that DeleteRentable works correctly
#
#  Scenario:
#  A rentable can be delete if it:
#	1. has no associated assessments
#	2. is not called out in any Rental Agreement
#	3. has no reservations in the future.
#
#
#  Expected Results:
#  Deletes go through if the criteria above are met and that they are
#  inhibited if any of the criteria are met.
#------------------------------------------------------------------------------
TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    #-----------------------------------------------------
    # This delete is on RID=16, carport #2  (CP002)
    #-----------------------------------------------------
    encodeRequest '{"cmd":"delete","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentable/1/16" "request" "${TFILES}${STEP}"  "Rentable-delete"

    #-------------------------------------------------------------------
    # This delete is on RID=23, cannot be allowed as it is called out in
    # a rental agreement and appears in assessments.
    #-------------------------------------------------------------------
    encodeRequest '{"cmd":"delete","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentable/1/23" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------------------
    # This delete is on RID=9, cannot be allowed as it is called out in
    # a rental agreement and appears in assessments.
    #-------------------------------------------------------------------
    encodeRequest '{"cmd":"delete","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentable/1/9" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
