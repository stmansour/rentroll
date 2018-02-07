#!/bin/bash

TOP=../..
RRBIN=${TOP}/tmp/rentroll
RSD="-rsd ${RRBIN}"
DBGENDIR=${TOP}/tools/dbgen

TESTNAME="Cypress UI Test"
TESTSUMMARY="UI Testing with cypress"

# do not create new db
CREATENEWDB=0

source ../share/base.sh

# server with noauth
RENTROLLSERVERAUTH="-noauth"

# specific file that needs to be tested
CYPRESS_SPEC="./cypress/integration/roller_spec.js"

#--------------------------------------------------------------------
#  Use custom dumped .sql file for the webclient UI tests
#--------------------------------------------------------------------
echo "*** loading data from webclientTest.sql into rentroll db ***"
mysql --no-defaults rentroll < webclientTest.sql
if [[ $? == 0 ]]; then
    echo "*** data has been loaded from webclientTest.sql in rentroll db ***"
else
    exit 1
fi

if [ "${IAMJENKINS}" == "jenkins" ]; then
    # if build machine then record the activity
    doCypressUITest "a" "--spec ${CYPRESS_SPEC} --record" "CypressUITesting"
else
    # run cypress test with only roller_spec.js with videoRecording false as of now
    doCypressUITest "a" "--config videoRecording=false --spec ${CYPRESS_SPEC}" "CypressUITesting"
fi

# logcheck
