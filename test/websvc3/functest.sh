#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Tasks, Tasklists"
TESTSUMMARY="Test Tasks, Tasklists, TaskDescriptors, TaskListDefinitions"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < tasks.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer


#------------------------------------------------------------------------------
#  TEST a
#  Basic Tests
#
#  Scenario:
#		Try all the readers, inserters, updaters and deleters.
#
#  Expected Results:
#	a0 - The db we read has 1 pre-defined TaskListDefinition
#   a1 - Write a new TaskListDefinition
#   a2 - Update a TaskListDefinition
#   a3 - This search should return 2 matches
#   a4 - Delete it - which means set the FLAGS to make it inactive
#   a5 - Read back TLDID 2, ensure that Name and FLAGS were updated
#   a6 - This search should only return 1 match because TLDID 2 was made inactive
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tlds/1" "request" "a0"  "WebService--Search_TaskListDefinitions"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22Cycle%22%3A6%2C%22Epoch%22%3A%221%2F1%2F2018%22%2C%22EpochDue%22%3A%221%2F31%2F2018%22%2C%22EpochPreDue%22%3A%221%2F20%2F2018%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Period%20Close%22%2C%22TLDID%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/0" "request" "a1"  "WebService--Insert_TaskListDefinitions"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22TLDID%22%3A2%2C%22BID%22%3A1%2C%22Cycle%22%3A6%2C%22Epoch%22%3A%221%2F1%2F2018%22%2C%22EpochDue%22%3A%221%2F31%2F2018%22%2C%22EpochPreDue%22%3A%221%2F20%2F2018%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Apts%20Period%20Close%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/2" "request" "a2"  "WebService--Update_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tlds/1" "request" "a3"  "WebService--Search_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22delete%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/2" "request" "a4"  "WebService--Delete_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/2" "request" "a5"  "WebService--Read_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tlds/1" "request" "a6"  "WebService--Search_TaskListDefinitions"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck
