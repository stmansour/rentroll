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

# INIITIATE A NEW RAFLOW WITH SECTIONS BASIC DATA
echo "%7B%22cmd%22%3A%22init%22%2C%22FlowType%22%3A%22RA%22%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/0/" "request" "a0"  "RAFlow - initiate brand new raflow"

# CHANGE ALL DATES(TERM, RENT, POSSESSION) START: 11 JAN, 2018 | STOP: 1 JAN, 2020, CSAGENT TO 1
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22dates%22%2C%22BID%22%3A1%2C%22Data%22%3A%7B%22CSAgent%22%3A1%2C%22RentStop%22%3A%221%2F1%2F2020%22%2C%22RentStart%22%3A%2211%2F1%2F2018%22%2C%22AgreementStop%22%3A%221%2F1%2F2020%22%2C%22AgreementStart%22%3A%2211%2F1%2F2018%22%2C%22PossessionStop%22%3A%221%2F1%2F2020%22%2C%22PossessionStart%22%3A%2211%2F1%2F2018%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a1"  "RAFlow - modify dates and cs agent"

# ADD EXISTING PERSON WITH TCID 3 (HAVING ONE PET & ONE VEHICLE)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A3%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-person/1/1/" "request" "a2"  "RAFlow - add person (TCID:1)"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
