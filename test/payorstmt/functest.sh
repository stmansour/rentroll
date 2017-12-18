#!/bin/bash

TESTNAME="Payor Statement Test"
TESTSUMMARY="Test Payor Statements"

CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < pstmt.sql

source ../share/base.sh

dorrtest "a" "-j 2017-01-01 -k 2017-02-01 -b ${BUD} -r 23,1" "PayorStatement-Bill-JAN"
dorrtest "b" "-j 2017-02-01 -k 2017-03-01 -b ${BUD} -r 23,1" "PayorStatement-Bill-FEB"
dorrtest "c" "-j 2017-03-01 -k 2017-04-01 -b ${BUD} -r 23,1" "PayorStatement-Bill-MAR"
dorrtest "d" "-j 2017-04-01 -k 2017-05-01 -b ${BUD} -r 23,1" "PayorStatement-Bill-APR"

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

# Request a payor detail statment for user 1
echo "# Create a non-recurring assessment"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%222%2F1%2F2017%22%7D" > request
dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "a0"  "PayorStatement--GridDetail"

echo "%7B%22cmd%22%3A%22get%22%7D" > request
dojsonPOST "http://localhost:8270/v1/payorstmtinfo/1/1" "request" "a1"  "PayorStatement--StmtInfo"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
