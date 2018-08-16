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
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "a1"  "Validate RAFlow -- NoError"

# People section related functional tests
#------------------------------------------------------------------------------
#  TEST p0
#  Validate raflow which have no renter
#
#  Expected Results:
#  1. Error: must one renter exists
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A3%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p0"  "Validate RAFlow -- Renter"

#------------------------------------------------------------------------------
#  TEST p1
#  Validate raflow which have one transanctant with not sufficient required information to get approval
#
#  Expected Results:
#  1. Error: Primary Email, Workphone, Occupation, Current Information, Taxpayer ID, Gross Income, Drivers Lic,
#  Emergency contact, Source field have respective error message.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A4%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p1"  "Validate RAFlow -- Basic Detail"

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
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p2"  "Validate RAFlow -- CompanyName"

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
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p3"  "Validate RAFlow -- FirstName and LastName"

#------------------------------------------------------------------------------
#  TEST p4
#  Validate raflow : People section
#
#  Scenario:
#  If role is set to Renter or guarantor than it must have mentioned GrossIncome
#
#  Expected Results:
#  1. Error: Gross income must have error
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A7%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p4"  "Validate RAFlow -- Gross Income"

#------------------------------------------------------------------------------
#  TEST p5
#  Validate raflow : People section
#
#  Scenario:
#  Do not provide workphone and cellphone
#
#  Expected Results:
#  1. Error: Either Workphone or CellPhone is compulsory.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A8%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p5"  "Validate RAFlow -- Workphone or Cell phone"

#------------------------------------------------------------------------------
#  TEST p6
#  Validate raflow : People section
#
#  Scenario:
#  Do not provide EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail.
#
#  Expected Results:
#  1. Error: EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail are required when IsCompany flag is false.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A9%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p6"  "Validate RAFlow -- Emergency contact information"

#------------------------------------------------------------------------------
#  TEST p7
#  Validate raflow : People section
#
#  Scenario:
#  There are two transanctant. One have role: Renter, User Another have role: Gurantor
#  Do not provide sourceSLSID for each.
#
#  Expected Results:
#  1. Error: SourceSLSID must be greater than 0 when role is set to Renter, User
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A10%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p7"  "Validate RAFlow -- SourceSLSID"

#------------------------------------------------------------------------------
#  TEST p8
#  Validate raflow : People section
#
#  Scenario:
#  Brand new RA application must have have current address.
#
#  Expected Results:
#  1. Error: "current" address related information required
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A11%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p8"  "Validate RAFlow -- Current Address"

#------------------------------------------------------------------------------
#  TEST p9
#  Validate raflow : People section
#
#  Scenario:
#  TaxpayorID is only require when role is set to Renter or Guarantor.
#  Here, 3 transanct of each role. Each role have TaxpayorID blank.
#
#  Expected Results:
#  1. Error: TaxpayorID must be require for the transanctant who have role renter or gurantor
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A12%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "p9"  "Validate RAFlow -- TaxpayorID"

#------------------------------------------------------------------------------
#  TEST d0
#  Validate raflow : Date section
#
#  Scenario:
#  Stop dates are not prior to Start dates
#
#  Expected Results:
#  1. Error: Start dates must be prior to stop dates
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A13%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/" "request" "d0"  "Validate RAFlow -- Dates"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
