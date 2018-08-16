#!/bin/bash

TESTNAME="Validate RAFlow"
TESTSUMMARY="Test for validating RAFlow business check and basic check"
DBGEN=../../tools/dbgen
CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < validateraflow.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  Validate newly initiated RAFlow
#
#  Expected Results:
#  1. Error: must be at least one parent rentable exist
#  2. Error: must be at least one occupant exist
#------------------------------------------------------------------------------

# Send the command to validate initiated flow:
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "a0"  "Validate RAFlow -- Initiated RAFlow"

#------------------------------------------------------------------------------
#  TEST a1
#  Validate raflow which have no error
#
#  Expected Results:
#  1. Error count must be 0
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A2%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "a1"  "Validate RAFlow -- Error free RAFlow"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
