#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Deposits"
TESTSUMMARY="Test Desposits web services"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < ../ws/restore.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
startRentRollServer

echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22DID%22%2C%22direction%22%3A%22asc%22%7D%5D%2C%22searchDtStart%22%3A%221%2F1%2F2016%22%2C%22searchDtStop%22%3A%225%2F1%2F2016%22%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1" "request" "a0"  "WebService--Read_Deposits"

echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22DID%22%2C%22direction%22%3A%22asc%22%7D%5D%2C%22searchDtStart%22%3A%221%2F1%2F2016%22%2C%22searchDtStop%22%3A%225%2F1%2F2016%22%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/3" "request" "a1"  "WebService--Read_Deposit_3"

# Read the list of receipts associated with a deposit
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%5D%2C%22searchDtStart%22%3A%221%2F1%2F2016%22%2C%22searchDtStop%22%3A%225%2F1%2F2016%22%7D" > request
dojsonPOST "http://localhost:8270/v1/depositlist/1/2" "request" "a2"  "WebService--Read_Receipts_for_deposit_2"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck
