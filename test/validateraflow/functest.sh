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

# People section related functional tests
#------------------------------------------------------------------------------
#  TEST p0
#  Validate raflow which have no renter
#
#  Expected Results:
#  1. Error: must one renter exists
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A3%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p0"  "Validate RAFlow -- One renter must exists"

#------------------------------------------------------------------------------
#  TEST p1
#  Validate raflow which have one transanctant with not sufficient required information to get approval
#
#  Expected Results:
#  1. Error: Primary Email, Workphone, Occupation, Current Information, Taxpayer ID, Gross Income, Drivers Lic,
#  Emergency contact, Source field have respective error message.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A4%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p1"  "Validate RAFlow -- Transactant must fill basic detail"

#------------------------------------------------------------------------------
#  TEST p2
#  Validate raflow : People section
#
#  Scenario:
#  isCompany flag is true but Company Name isn't provided
#
#  Expected Results:
#  1. Error: Company name is must require.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A5%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p2"  "Validate RAFlow -- Must require company name"

#------------------------------------------------------------------------------
#  TEST p3
#  Validate raflow : People section
#
#  Scenario:
#  isCompany flag is false but Firstname and lastname isn't provided
#
#  Expected Results:
#  1. Error: Firstname and lastnanme are must require.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A6%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p3"  "Validate RAFlow -- Must require firstname and lastname"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
