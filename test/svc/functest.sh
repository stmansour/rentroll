#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Web Services"
TESTSUMMARY="Test Web Services"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..." 
echo "Cmd =  ${BASHDIR}/getdb.sh"
${TOP}/tools/bashtools/getdb.sh

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
startRentRollServer

echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/accounts/1" "request" "s"  "WebService--ChartOfAccounts"
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22LastName%22%2C%22direction%22%3A%22asc%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/transactants/1" "request" "t"  "WebService--GetTransactants"
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/rentables/1" "request" "u"  "WebService--GetRentables"
echo "request%3d%7b%22cmd%22%3a%22get%22%2c%22selected%22%3a%5b%5d%2c%22limit%22%3a100%2c%22offset%22%3a0%7d" > request
dojsonPOST "http://localhost:8270/v1/receipts/1" "request" "v"  "WebService--GetReceipts"
echo "request%3d%7b%22cmd%22%3a%22get%22%2c%22selected%22%3a%5b%5d%2c%22limit%22%3a100%2c%22offset%22%3a0%7d" > request
dojsonPOST "http://localhost:8270/v1/asm/1" "request" "w"  "WebService--GetAssessments"


stopRentRollServer
echo "RENTROLL SERVER STOPPED" 

logcheck
