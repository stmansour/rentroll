#!/bin/bash

TOP=../..
RRBIN=${TOP}/tmp/rentroll
RSD="-rsd ${RRBIN}"

TESTNAME="CasperJS UI Test"
TESTSUMMARY="UI Testing with casperjs"

CREATENEWDB=0

source ../share/base.sh

# TODO: feed some known data in db server, so that expacted results can be match for each UI testcases
# use `rrloadcsv` program

###################
# TEMP BASED DB
###################
# cp ../ws/restore.sql .

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
# echo "Create new database..."
# mysql --no-defaults rentroll < restore.sql
###################

echo "STARTING RENTROLL SERVER"
startRentRollServer

echo "Running casper UI test cases..."
# casperjs test index.js
# casperjs test index.js --log-level=debug --verbose

# call loader
doCasperUITest "a" "./index.js" "CasperUITesting"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

