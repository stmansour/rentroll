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

#--------------------------------------------------------------------
#  Use the testdb for these tests... (dbgen with db4.json, as of now)
#--------------------------------------------------------------------

# server with noauth
RENTROLLSERVERAUTH="-noauth"

# specific file that needs to be tested
CYPRESS_SPEC="./cypress/integration/roller_spec.js"

pushd ${DBGENDIR}
./dbgen -f db4.json -noauth
popd

if [ "${IAMJENKINS}" == "jenkins" ]; then
    # if build machine then record the activity
    doCypressUITest "a" "--spec ${CYPRESS_SPEC} --record" "CypressUITesting"
else
    # run cypress test with only roller_spec.js with videoRecording false as of now
    doCypressUITest "a" "--config videoRecording=false --spec ${CYPRESS_SPEC}" "CypressUITesting"
fi

logcheck
