#!/bin/bash

TESTNAME="Business"
TESTSUMMARY="Test Business create and update"
DBGENDIR=../../tools/dbgen
CREATENEWDB=0
RRBIN="../../tmp/rentroll"

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"

# RENTROLLSERVERNOW="-testDtNow 10/24/2018"
# SINGLETEST=""  # This runs all the tests
# echo "SINGLETEST = ${SINGLETEST}"

#------------------------------------------------------------------------------
#  TEST a
#  Validate the basic readers and writers for Businesses.
#
#  Scenario
#  Read all the businesses in the database
#
#  Expected Results:
#  1. first form includes a blank BID.  Should read the biz just fine
#  2. Second form include a 0 BID, Should read the biz just fine
#  3. Request a specific business
#  4. Create a new business
#  5. Read the businesses again to make sure the new business was created
#  6. Validate that the typedown query functions properly
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database... db${TFILES}.sql"

    mysql --no-defaults rentroll < db${TFILES}.sql
    startRentRollServer

    #----------------------------------------------------------------
    # Requests business with url  v1/business/
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22limit%22%3A100%7D" > request
    dojsonPOST "http://localhost:8270/v1/business/" "request" "${TFILES}0"  "WebService--GetBusinesses"

    #----------------------------------------------------------------
    # Requests business with url  v1/business/0   this should have
    # exactly the same result as the previous read
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22limit%22%3A100%7D" > request
    dojsonPOST "http://localhost:8270/v1/business/0" "request" "${TFILES}0"  "WebService--GetBusinesses"

    #----------------------------------------------------------------
    # request a specific business
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22limit%22%3A100%7D" > request
    dojsonPOST "http://localhost:8270/v1/business/2" "request" "${TFILES}1"  "WebService--GetBusiness"

    #----------------------------------------------------------------
    # Create a new business
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22Name%22%3A%22Billy%20Bob's%20Bungee%20Emporium%22%2C%22BUD%22%3A%22BBBE%22%2C%22BID%22%3A0%2C%22DefaultRentCycle%22%3A6%2C%22DefaultProrationCycle%22%3A4%2C%22DefaultGSRPC%22%3A4%2C%22FLAGS%22%3A1%2C%22EDIenabled%22%3Atrue%2C%22AllowBackdatedRA%22%3Atrue%2C%22Disabled%22%3Afalse%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/business/0" "request" "${TFILES}2"  "WebService--CreateBusiness"

    #----------------------------------------------------------------
    # Requests business with url  v1/business/
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22limit%22%3A100%7D" > request
    dojsonPOST "http://localhost:8270/v1/business/" "request" "${TFILES}3"  "WebService--GetBusinesses"

    #----------------------------------------------------------------
    # Validate typedown
    #----------------------------------------------------------------
    dojsonGET "http://localhost:8270/v1/tltd/REX?request=%7B%22search%22%3A%22m%22%2C%22max%22%3A250%7D" "${TFILES}4" "WebService--TaskListTypeDown"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
