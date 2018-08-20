#!/bin/bash

TESTNAME="Init RA"
TESTSUMMARY="Test to init raflow and migrate to permanent tables RA"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

echo "Create new database..."
mysql --no-defaults rentroll < ra0.sql

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  Following test cases will be performed using only web service calls.
#
#  Test cases will be performed in a manner of sequences.
#  It will add/init data section by section as follows:
#       1. Dates/Agent
#       2. People
#       3. Pets
#       4. Vehicles
#       5. Rentables
#       6. Parent/Child
#       7. Tie
#------------------------------------------------------------------------------

# INIITIATE A NEW RAFLOW WITH BASE SECTIONS DATA
echo "%7B%22cmd%22%3A%22init%22%2C%22FlowType%22%3A%22RA%22%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/0/" "request" "a0"  "RAFlow - initiate brand new raflow"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
