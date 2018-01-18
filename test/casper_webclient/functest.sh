#!/bin/bash

TOP=../..
RRBIN=${TOP}/tmp/rentroll
RSD="-rsd ${RRBIN}"

TESTNAME="CasperJS UI Test"
TESTSUMMARY="UI Testing with casperjs"

# do not create new db
CREATENEWDB=0

source ../share/base.sh

#--------------------------------------------------------------------
#  Use the testdb for these tests... (dbgen with db4.json, as of now)
#--------------------------------------------------------------------

# server with noauth
RENTROLLSERVERAUTH="-noauth"

echo "Running casper UI test cases..."
# casperjs test index.js
# casperjs test index.js --log-level=debug --verbose

# call loader
doCasperUITest "a" "./index.js" "CasperUITesting"

logcheck
