#!/bin/bash

TESTNAME="TWS - Scheduled Work tester"
TESTSUMMARY="Validate prescheduled work execution"
TOP="../.."
BINDIR="${TOP}/tmp/rentroll"

CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

if [ ! -f bizerr.csv ]; then
	ln -s ${BINDIR}/bizerr.csv
fi

source ../share/base.sh

./tws -noauth > z

genericlogcheck "z"  ""  "Validations"

# echo "STARTING RENTROLL SERVER"
# RENTROLLSERVERAUTH="-noauth"
# startRentRollServer

# # get Statement
# # echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222016-11-01%22%2C%22searchDtStop%22%3A%222016-12-01%22%7D" > request
# dojsonPOST "http://localhost:8270/v1/stmtdetail/1/5" "request" "a1"  "WebService--StatementDetail"

# stopRentRollServer
# echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
