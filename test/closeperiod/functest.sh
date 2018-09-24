#!/bin/bash
TESTNAME="CLOSE PERIOD test"
TESTSUMMARY="Close a Period and complete a TaskList"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
CREATENEWDB=0

source ../share/base.sh

echo "BEGIN CLOSE PERIOD FUNCTIONAL TEST" >>${LOGFILE}
echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  Test the closeinfo service and the closeperiod service
#
#  Scenario:
#  The database is initialized to have 8 task list instances all of which are
#  not morked as completed. We will complete a few, one-by-one and make sure
#  that they are marked as completed when we get the closeinfo response.
#
#  Expected Results:
#   1.  First call to closeinfo returns that no periods are closed. This is
#       indicated by CPID = 0.
#   2.  Send a command to close the first tasklist - January 2018.
#   3.  Send a command to close the period for January
#   4.  Read back the closeinfo and make sure that we have CPID == 1 and that
#       the close date is Jan 31, 2018.
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Create new database... x${TFILES}.sql"
    mysql --no-defaults rentroll < x${TFILES}.sql

    # Get the initial info, before any tasklist has been completed
    echo "%7B%22cmd%22%3A%22get%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/closeinfo/1/" "request" "${TFILES}0"  "WebService--CloseInfo"

    # Send a command to close the January 2018 period.
    echo "%7B%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22TLID%22%3A1%2C%22PTLID%22%3A0%2C%22TLDID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%7B%22id%22%3A%22REX%22%2C%22text%22%3A%22REX%22%7D%2C%22Name%22%3A%22Monthly%20Close%22%2C%22Cycle%22%3A6%2C%22DtDone%22%3A%221%2F1%2F1970%22%2C%22DtDue%22%3A%22Wed%2C%2031%20Jan%202018%2017%3A00%3A00%20GMT%22%2C%22DtPreDue%22%3A%22Sat%2C%2020%20Jan%202018%2017%3A00%3A00%20GMT%22%2C%22DtPreDone%22%3A%221%2F1%2F1970%22%2C%22DtLastNotify%22%3A%221900-01-01%2000%3A00%3A00%20UTC%22%2C%22DurWait%22%3A86400000000000%2C%22ChkDtDone%22%3Atrue%2C%22ChkDtDue%22%3Atrue%2C%22ChkDtPreDue%22%3Atrue%2C%22ChkDtPreDone%22%3Afalse%2C%22FLAGS%22%3A6%2C%22DoneUID%22%3A0%2C%22DoneName%22%3A%22%22%2C%22PreDoneUID%22%3A0%2C%22PreDoneName%22%3A%22%22%2C%22EmailList%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22CreateTS%22%3A%222018-09-24%2010%3A51%3A00%20UTC%22%2C%22CreateBy%22%3A0%2C%22LastModTime%22%3A%222018-09-24%2010%3A51%3A00%20UTC%22%2C%22LastModBy%22%3A0%2C%22TZOffset%22%3A420%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/tl/1/1" "request" "${TFILES}1"  "CloseInfo-complete_first_tasklist"

    # Now close January
    echo "%7B%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22TLID%22%3A1%2C%22TLName%22%3A%22Monthly%20Close%22%2C%22LastDtDone%22%3A%221900-01-01%2000%3A00%3A00%20UTC%22%2C%22LastDtClose%22%3A%221900-01-01%2000%3A00%3A00%20UTC%22%2C%22LastLedgerMarker%22%3A%221900-01-01%2000%3A00%3A00%20UTC%22%2C%22CloseTarget%22%3A%222018-01-31%2017%3A00%3A00%20UTC%22%2C%22TLIDTarget%22%3A1%2C%22TLNameTarget%22%3A%22Monthly%20Close%22%2C%22DtDueTarget%22%3A%222018-01-31%2017%3A00%3A00%20UTC%22%2C%22DtDoneTarget%22%3A%222018-09-24%2018%3A46%3A00%20UTC%22%2C%22DtDone%22%3A%221900-01-01%2000%3A00%3A00%20UTC%22%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/closeperiod/1/" "request" "${TFILES}2"  "WebService--CloseInfo"

    # Send the command to change the flow to Active:
    echo "%7B%22cmd%22%3A%22get%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/closeinfo/1/" "request" "${TFILES}3"  "WebService--CloseInfo"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck
