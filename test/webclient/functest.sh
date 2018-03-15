#!/bin/bash

TOP=../..
RRBIN=${TOP}/tmp/rentroll
RSD="-rsd ${RRBIN}"
DBGENDIR=${TOP}/tools/dbgen

TESTNAME="Cypress UI Test"
TESTSUMMARY="UI Testing with cypress"

# do not create new db
CREATENEWDB=0

# DBNAME
DBNAME="accord"

# set config.json path of this current folder
CONFIGPATH="$(pwd)"

source ../share/base.sh

# specific file that needs to be tested
CYPRESS_SPEC="./cypress/integration/*"

#--------------------------------------------------------------------
#  Use custom dumped "rentroll" .sql file for the webclient UI tests
#--------------------------------------------------------------------

echo "*** loading data from webclientTest.sql into rentroll db ***"
mysql --no-defaults rentroll < webclientTest.sql
if [[ $? == 0 ]]; then
    echo "*** data has been loaded from webclientTest.sql in rentroll db ***"
else
    exit 1
fi

#--------------------------------------------------------------------
#  Use custom dumped "accord" .sql file for the webclient UI tests
#--------------------------------------------------------------------
echo "*** loading data from accord.sql into accord db ***"
mysql --no-defaults accord < accord.sql
if [[ $? == 0 ]]; then
    echo "*** data has been loaded from accord.sql in accord db ***"
else
    exit 1
fi

if [ "${IAMJENKINS}" == "jenkins" ]; then
    # if build machine then record the activity
    doCypressUITest "a" "--env configFile=build --spec ${CYPRESS_SPEC} --record" "CypressUITesting"
else
    # run cypress test with only receipt_2_spec.js with videoRecording false as of now
    doCypressUITest "a" "--env configFile=development --spec ${CYPRESS_SPEC}" "CypressUITesting"
fi

# logcheck
