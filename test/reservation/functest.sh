#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="Reservation"
TESTSUMMARY="Test Reservation search, create, mod"
DBGENDIR=${SRCTOP}/tools/dbgen
CREATENEWDB=0
RRBIN="${SRCTOP}/tmp/rentroll"
CATRML="${SRCTOP}/tools/catrml/catrml"

#SINGLETEST=""  # This runs all the tests

source ${TESTHOME}/share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 10/24/2018"

#------------------------------------------------------------------------------
#  TEST a
#
#  Determine the availability of a rentable type during a particular timeframe
#
#  Scenario:
#  find the availability of a RTID = 5 and BID=1 from 2019-02-14 to 2019-02-15.
#
#  Expected Results:
#   1.  There should be one rentable available at that time.
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    # # Send the command to change the flow to Active:
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%228%2F31%2F2018%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22DtStart%22%3A%222019-02-14%22%2C%22DtStop%22%3A%222019-02-15%22%2C%22RLID%22%3A0%2C%22RTRID%22%3A0%2C%22RTID%22%3A5%2C%22RID%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1" "request" "a0"  "reservation-searchAvailable"
    #
    # # Generate an assessment report from Aug 1 to Oct 1. The security deposit
    # # assessment for RAID 1 should no longer be present
    # docsvtest "a1" "-G ${BUD} -g 8/1/18,10/1/18 -L 11,${BUD}" "Assessments-2018-AUG"
    #
    # # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # # info.
    # echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%228%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    # dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "a2"  "PayorStatement--StmtInfo"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
