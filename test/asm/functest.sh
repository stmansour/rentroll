#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableLeaseStatus"
TESTSUMMARY="Test Rentable Lease Status code"
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
#  Validate that assessments cannot be saved if the date values are in a
#  Closed period.
#
#  Scenario:
#  In the test db, we have closed periods of 7/2018, 8/2018, 9/2018.  We
#  will attempt to save assessments dated in the month of 7/2018
#
#
#  Expected Results:
#   1.  The attempts to save the assessment should result in an error
#------------------------------------------------------------------------------
TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    #----------------------------------------------------------
    # try to create a new assessment in a closed period
    #----------------------------------------------------------
    encodeRequest '{"cmd":"save","recid":0,"name":"asmEpochForm","record":{"ARID":19,"recid":0,"RID":1,"ASMID":0,"PASMID":0,"ATypeLID":0,"InvoiceNo":0,"RAID":1,"BID":1,"BUD":"REX","Start":"7/18/2018","Stop":"7/18/2018","RentCycle":0,"ProrationCycle":0,"TCID":0,"Amount":50,"Rentable":"Rentable001","AcctRule":"","Comment":"","ExpandPastInst":false,"FLAGS":0,"Mode":0,"LastModByUser":"","CreateByUser":""}}'
    dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-1a"

    #----------------------------------------------------------
    # try to update an existing assessment in a closed period
    #----------------------------------------------------------
    encodeRequest '{"cmd":"save","recid":0,"name":"asmInstForm","record":{"recid":0,"ASMID":2,"BID":1,"BUD":"REX","PASMID":1,"RID":1,"Rentable":"Rentable001","RAID":1,"Amount":1000,"Start":"7/1/2018","Stop":"7/1/2018","RentCycle":6,"ProrationCycle":4,"InvoiceNo":0,"ARID":26,"Comment":"Try to change in closed period","LastModByUser":"UID-0","CreateByUser":"UID-0","ExpandPastInst":false,"FLAGS":2,"Mode":0}}'
    dojsonPOST "http://localhost:8270/v1/asm/1/2" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-1a"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
