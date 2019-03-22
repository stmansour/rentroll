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

#------------------------------------------------------------------------------
#  TEST b
#
#  Validate that changes cannot be made to assessments, receipts, expenses,
#  and deposits during a closed period.  These tests will attempt to add,
#  change, and delete or reverse existing elements in a closed period. The
#  server must reject all these changes.
#
#  Scenario:
#  The comments below will describe what each step of the test is doing
#
#  Expected Results:
#
#   The server should return an error for each of the requests.
#------------------------------------------------------------------------------
TFILES="b"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    #------------------
    #  DEPOSITS
    #------------------

    # attempt to add a deposit in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"depositForm","Receipts":[95],"record":{"recid":0,"check":0,"DID":0,"BID":1,"BUD":"REX","DEPID":2,"DEPName":2,"DPMID":1,"DPMName":1,"Dt":"3/20/2018","FLAGS":0,"Amount":150,"ClearedAmount":0}}' > request
    dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "${TFILES}${STEP}"  "Deposit-SaveAttemptInClosedPeriod"

    # attempt to change a deposit in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"depositForm","Receipts":[76,15,77,16,78,44,45,93,46,94,47,75],"record":{"recid":1,"DID":1,"BID":1,"BUD":"REX","DEPID":2,"DEPName":2,"DPMID":1,"DPMName":1,"Dt":"1/3/2018","Amount":4836.12,"ClearedAmount":4800,"FLAGS":0}}' > request
    dojsonPOST "http://localhost:8270/v1/deposit/1/1" "request" "${TFILES}${STEP}"  "Deposit-ChangeAttemptInClosedPeriod"

    # attempt to delete a deposit in a closed period
    encodeRequest '{"cmd":"delete","formname":"depositFormBtns","DID":1}' > request
    dojsonPOST "http://localhost:8270/v1/deposit/1/1" "request" "${TFILES}${STEP}"  "Deposit-DeleteAttemptInClosedPeriod"

    # attempt to change the date of a deposit in an open period to a date in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"depositForm","Receipts":[13,29,43,60,74,91],"record":{"recid":15,"DID":15,"BID":1,"BUD":"REX","DEPID":2,"DEPName":2,"DPMID":1,"DPMName":1,"Dt":"2/1/2018","Amount":5020,"ClearedAmount":0,"FLAGS":0}}' > request
    dojsonPOST "http://localhost:8270/v1/receipt/1/15" "request" "${TFILES}${STEP}"  "Deposit-ChangeAttemptInClosedPeriod"

    #------------------
    #  EXPENSES
    #------------------

    # attempt to add an expense in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"expenseForm","record":{"recid":0,"EXPID":0,"ARID":4,"RID":2,"RAID":2,"BID":1,"BUD":"REX","Dt":"3/20/2018","Amount":50,"AcctRule":"","RName":"Rentable002","Comment":"","FLAGS":0,"Mode":0,"PREXPID":"","DtLastClose":"","LastModByUser":"","CreateByUser":""}}' > request
    dojsonPOST "http://localhost:8270/v1/expense/1/0" "request" "${TFILES}${STEP}"  "Expense-SaveAttemptInClosedPeriod"

    # attempt to change an expense in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"expenseForm","record":{"recid":0,"EXPID":1,"BID":1,"BUD":"REX","RID":3,"RAID":3,"Amount":150,"Dt":"4/1/2018","ARID":4,"ARName":"","RName":"Rentable003","FLAGS":0,"Comment":"test","DtLastClose":"2018-04-30 17:00:00 UTC","LastModByUser":"Steve Mansour","CreateByUser":"Steve Mansour","PREXPID":""}}' > request
    dojsonPOST "http://localhost:8270/v1/expense/1/1" "request" "${TFILES}${STEP}"  "Expense-ChangeAttemptInClosedPeriod"

    # attempt to Reverse an expense in a closed period
    encodeRequest '{"cmd":"delete","formname":"expenseForm","ID":1}' > request
    dojsonPOST "http://localhost:8270/v1/expense/1/1" "request" "${TFILES}${STEP}"  "Expense-ReverseAttemptInClosedPeriod"

    # attempt to change the date of an expense in an open period to a date in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"expenseForm","record":{"recid":0,"EXPID":2,"BID":1,"BUD":"REX","RID":2,"RAID":2,"Amount":20,"Dt":"3/20/2018","ARID":5,"ARName":"","RName":"Rentable002","FLAGS":0,"Comment":"","DtLastClose":"2018-04-30 17:00:00 UTC","LastModByUser":"Steve Mansour","CreateByUser":"Steve Mansour","PREXPID":""}}' > request
    dojsonPOST "http://localhost:8270/v1/receipt/1/2" "request" "${TFILES}${STEP}"  "Expense-ChangeAttemptInClosedPeriod"

    #------------------
    #  RECEIPTS
    #------------------

    # attempt to add a receipt in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"receiptForm","client":"roller","record":{"recid":0,"RCPTID":16,"PRCPTID":0,"BID":1,"DID":1,"BUD":"REX","RAID":1,"PMTID":2,"Payor":"Katlyn Wilford Paul (TCID: 1)","TCID":1,"Dt":"1/3/2018","DocNo":"659062","Amount":110,"ARID":25,"Comment":"payment for ASM-17","OtherPayorName":"","LastModByUser":"UID-0","CreateByUser":"UID-0","FLAGS":2,"RentableName":"","DtLastClose":"2018-04-30 17:00:00 UTC","PmtTypeName":2}}' > request
    dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "${TFILES}${STEP}"  "Receipt-SaveAttemptInClosedPeriod"

    # attempt to change a receipt in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"receiptForm","record":{"recid":0,"EXPID":1,"BID":1,"BUD":"REX","RID":3,"RAID":3,"Amount":150,"Dt":"4/1/2018","ARID":4,"ARName":"","RName":"Rentable003","FLAGS":0,"Comment":"test","DtLastClose":"2018-04-30 17:00:00 UTC","LastModByUser":"Steve Mansour","CreateByUser":"Steve Mansour","PREXPID":""}}' > request
    dojsonPOST "http://localhost:8270/v1/receipt/1/16" "request" "${TFILES}${STEP}"  "Receipt-ChangeAttemptInClosedPeriod"

    # attempt to Reverse a receipt in a closed period
    encodeRequest '{"cmd":"delete","formname":"receiptForm","RCPTID":16}' > request
    dojsonPOST "http://localhost:8270/v1/receipt/1/16" "request" "${TFILES}${STEP}"  "Receipt-ReverseAttemptInClosedPeriod"

    # attempt to change the date of a receipt in an open period to a date in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"receiptForm","client":"roller","record":{"recid":0,"RCPTID":91,"PRCPTID":0,"BID":1,"DID":15,"BUD":"REX","RAID":4,"PMTID":2,"Payor":"Elizebeth Azucena Santiago (TCID: 4)","TCID":4,"Dt":"2/1/2018","DocNo":"188752","Amount":1500,"ARID":25,"Comment":"payment for ASM-97","OtherPayorName":"","LastModByUser":"UID-0","CreateByUser":"UID-0","FLAGS":2,"RentableName":"","DtLastClose":"2018-04-30 17:00:00 UTC","PmtTypeName":2}}' > request
    dojsonPOST "http://localhost:8270/v1/receipt/1/91" "request" "${TFILES}${STEP}"  "Receipt-ChangeAttemptInClosedPeriod"

    #------------------
    #  ASSESSMENTS
    #------------------

    # attempt to add an assessment in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"asmEpochForm","record":{"ARID":8,"recid":0,"RID":1,"ASMID":0,"PASMID":0,"ATypeLID":0,"InvoiceNo":0,"RAID":1,"BID":1,"BUD":"REX","Start":"3/20/2018","Stop":"3/20/2020","RentCycle":6,"ProrationCycle":4,"TCID":0,"Amount":50,"Rentable":"Rentable001","AcctRule":"","Comment":"","ExpandPastInst":true,"FLAGS":0,"Mode":0,"LastModByUser":"","CreateByUser":""}}' > request
    dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "${TFILES}${STEP}"  "Assessment-SaveAttemptInClosedPeriod"

    # attempt to change an assessment in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"asmInstForm","record":{"recid":0,"ASMID":17,"BID":1,"BUD":"REX","PASMID":0,"RID":1,"Rentable":"Rentable001","RAID":1,"Amount":50,"Start":"1/3/2018","Stop":"1/3/2018","RentCycle":0,"ProrationCycle":0,"InvoiceNo":0,"ARID":39,"Comment":"","DtLastClose":"2018-04-30 17:00:00 UTC","LastModByUser":"UID-0","CreateByUser":"UID-0","ExpandPastInst":false,"FLAGS":2,"Mode":0}}' > request
    dojsonPOST "http://localhost:8270/v1/asm/1/17" "request" "${TFILES}${STEP}"  "Assessment-ChangeAttemptInClosedPeriod"

    # attempt to Reverse an assessment in a closed period
    encodeRequest '{"cmd":"delete","formname":"asmInstForm","ASMID":17,"ReverseMode":0}' > request
    dojsonPOST "http://localhost:8270/v1/asm/1/17" "request" "${TFILES}${STEP}"  "Assessment-ReverseAttemptInClosedPeriod"

    # attempt to change the date of an assessment in an open period to a date in a closed period
    encodeRequest '{"cmd":"save","recid":0,"name":"asmInstForm","record":{"recid":0,"ASMID":79,"BID":1,"BUD":"REX","PASMID":66,"RID":3,"Rentable":"Rentable003","RAID":3,"Amount":10,"Start":"2/1/2018","Stop":"2/1/2019","RentCycle":6,"ProrationCycle":4,"InvoiceNo":0,"ARID":24,"Comment":"","DtLastClose":"2018-04-30 17:00:00 UTC","LastModByUser":"UID-0","CreateByUser":"UID-0","ExpandPastInst":true,"FLAGS":2,"Mode":0}}' > request
    dojsonPOST "http://localhost:8270/v1/asm/1/79" "request" "${TFILES}${STEP}"  "Assessment-ChangeAttemptInClosedPeriod"
fi

#------------------------------------------------------------------------------
#  TEST c
#
#  Validate that we can reopen a closed period.  This is done by deleting
#  the last closed period.
#
#  Scenario:
#  Try to delete a ClosedPeriod prior to the last closed period - this must
#  result in an error.  Then delete the last closed period to validate that
#  it works correctly.
#
#  Expected Results:
#
#   The server should return an error for any attempt to delete a close
#   period that is not the LAST close period.
#
#   It should allow the deletion of the last closed period.
#------------------------------------------------------------------------------
TFILES="c"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    mysql --no-defaults rentroll < x${TFILES}.sql

    # Attempt to REOPEN A CLOSED PERIOD but not the last closed period.  This should fail
    encodeRequest '{"cmd":"delete","record":{"BID":1,"CPID":6,"TLID":1,"TLName":"Monthly Close","LastDtDone":"1900-01-01 00:00:00 UTC","LastDtClose":"1900-01-01 00:00:00 UTC","LastLedgerMarker":"1900-01-01 00:00:00 UTC","CloseTarget":"2018-01-31 17:00:00 UTC","TLIDTarget":1,"TLNameTarget":"Monthly Close","DtDueTarget":"2018-01-31 17:00:00 UTC","DtDoneTarget":"2018-09-24 18:46:00 UTC","DtDone":"1900-01-01 00:00:00 UTC"}}' > request
    dojsonPOST "http://localhost:8270/v1/closeperiod/1/5" "request" "${TFILES}${STEP}"  "CloseInfo-AttemptToDeleteClosePeriod"

    # REOPEN A CLOSED PERIOD (by deleting a Closed Period)
    encodeRequest '{"cmd":"delete","record":{"BID":1,"CPID":6,"TLID":1,"TLName":"Monthly Close","LastDtDone":"1900-01-01 00:00:00 UTC","LastDtClose":"1900-01-01 00:00:00 UTC","LastLedgerMarker":"1900-01-01 00:00:00 UTC","CloseTarget":"2018-01-31 17:00:00 UTC","TLIDTarget":1,"TLNameTarget":"Monthly Close","DtDueTarget":"2018-01-31 17:00:00 UTC","DtDoneTarget":"2018-09-24 18:46:00 UTC","DtDone":"1900-01-01 00:00:00 UTC"}}' > request
    dojsonPOST "http://localhost:8270/v1/closeperiod/1/6" "request" "${TFILES}${STEP}"  "CloseInfo-DeleteLastClosePeriod"


fi


stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck
