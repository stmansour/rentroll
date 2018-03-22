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
#  Basic TaskListDefinition Test Suite
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

#------------------------------------------------------------------------------
#  TEST b
#  Basic TaskDescriptor Test Suite
#
#  Scenario:
#		Try all the readers, inserters, updaters and deleters.
#
#  Expected Results:
#	b0 - The db we read has 3 pre-defined TaskListDefinition
#   b1 - Insert a new TaskDescriptor (4)
#   b2 - Read back the newly inserted task descriptor (4)
#   b3 - Modify Task Descriptor 4
#   b4 - Read back the updated task descriptor to make sure it worked
#   b5 - Delete Descriptor 4
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tds/1/1" "request" "b0"  "WebService--Search_TaskDescriptors"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A0%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Validate%20Washing%20Room%20Totals%22%2C%22TDID%22%3A0%2C%22TLDID%22%3A1%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A2%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/0" "request" "b1"  "WebService--Insert_TaskDescriptor"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/4" "request" "b2"  "WebService--Read_TaskDescriptor"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A0%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Validate%20Laundry%20Totals%22%2C%22TDID%22%3A4%2C%22TLDID%22%3A1%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A2%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/4" "request" "b3"  "WebService--Update_TaskDescriptor"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/4" "request" "b4"  "WebService--Read_TaskDescriptor"
echo "%7B%22cmd%22%3A%22delete%22%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/4" "request" "b5"  "WebService--Delete_TaskDescriptor"

#------------------------------------------------------------------------------
#  TEST c
#  TaskDescriptor Business Logic Test Suite
#
#  Scenario:
#		Break all the bizlogic rules and make sure it gets caught.
#
#  Expected Results:
#	c0 - Attempt to assign a Descriptor to a non-existent TaskListDefinition
#   c1 - Attempt to assign a Descriptor to an invalid business
#   c2 - Attempt to save a Descriptor with no name
#------------------------------------------------------------------------------
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A0%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22break%20bizlogic%22%2C%22TDID%22%3A0%2C%22TLDID%22%3A7981%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/0" "request" "c0"  "WebService--Insert_TaskDescriptor"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A337%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22break%20bizlogic%22%2C%22TDID%22%3A0%2C%22TLDID%22%3A1%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/td/337/0" "request" "c1"  "WebService--Insert_TaskDescriptor"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22%22%2C%22TDID%22%3A0%2C%22TLDID%22%3A1%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/0" "request" "c2"  "WebService--Insert_TaskDescriptor"

#------------------------------------------------------------------------------
#  TEST d
#  TaskList Test Suite
#
#  Scenario:
#		Break all the bizlogic rules and make sure it gets caught.
#
#  Expected Results:
#	d0 - Insert a new instance of TLD 1
#   d1 - Insert another instance of TLD 2
#   d2 - Read instance 1 (TLID = 1)
#   d3 - update instnce 2
#   d4 - read back the update
#   d5 - delete instance 2
#------------------------------------------------------------------------------
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22recid%22%3A3%2C%22Cycle%22%3A6%2C%22DtDue%22%3A%221%2F31%2F2018%22%2C%22DtPreDue%22%3A%221%2F20%2F2018%22%2C%22Pivot%22%3A%222%2F14%2F2018%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Period%20Close%22%2C%22TLID%22%3A0%2C%22TLDID%22%3A1%2C%22DoneUID%22%3A0%2C%22PreDoneUID%22%3A0%2C%22Comment%22%3A%22An%20instance%20of%20TLDID%201%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tl/1/0" "request" "d0"  "WebService--Insert_TaskList"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22recid%22%3A3%2C%22Cycle%22%3A6%2C%22DtDue%22%3A%221%2F31%2F2018%22%2C%22DtPreDue%22%3A%221%2F20%2F2018%22%2C%22Pivot%22%3A%222%2F14%2F2018%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Period%20Close%22%2C%22TLID%22%3A0%2C%22TLDID%22%3A1%2C%22DoneUID%22%3A0%2C%22PreDoneUID%22%3A0%2C%22Comment%22%3A%22Another%20instance%20of%20TLDID%201%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tl/1/0" "request" "d1"  "WebService--Insert_TaskList"
echo "%7B%22recid%22%3A1%2C%22cmd%22%3A%22get%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22recid%22%3A3%2C%22Cycle%22%3A6%2C%22DtDue%22%3A%221%2F31%2F2018%22%2C%22DtPreDue%22%3A%221%2F20%2F2018%22%2C%22Pivot%22%3A%223%2F3%2F2018%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Period%20Close%22%2C%22TLID%22%3A1%2C%22TLDID%22%3A1%2C%22DoneUID%22%3A0%2C%22PreDoneUID%22%3A0%2C%22Comment%22%3A%22An%20instance%20of%20TLDID%201%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tl/1/1" "request" "d2"  "WebService--Read_TaskList"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%222%2F28%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tls/1/0" "request" "d3"  "WebService--Search_TaskList"
echo "%7B%22cmd%22%3A%22delete%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tl/1/2" "request" "d4"  "WebService--Delete_TaskList"

#------------------------------------------------------------------------------
#  TEST e
#  TaskList Test Suite
#
#  Scenario:
#		Break all the bizlogic rules and make sure it gets caught.
#
#  Expected Results:
#	e0 - Get the task list for TaskList 2
#   e1 - Get a specific task
#   e2 - 
#   e3 - 
#   e4 - 
#   e5 - 
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tasks/1/2" "request" "e0"  "WebService--Get_Tasks"
echo "%7B%22cmd%22%3A%22get%22%7D" > request
dojsonPOST "http://localhost:8270/v1/task/1/5" "request" "e1"  "WebService--Get_Task"
echo "%7B%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22Comment%22%3A%22This%20is%20an%20update%22%2C%22CreateBy%22%3A0%2C%22CreateTS%22%3A%222018-03-21T20%3A07%3A48Z%22%2C%22DoneUID%22%3A0%2C%22DtDone%22%3A%222018-02-28T17%3A00%3A00Z%22%2C%22DtDue%22%3A%222018-02-28T20%3A00%3A00Z%22%2C%22DtPreDone%22%3A%222018-02-19T23%3A12%3A07Z%22%2C%22DtPreDue%22%3A%222018-02-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22LastModBy%22%3A0%2C%22LastModTime%22%3A%222018-03-21T20%3A07%3A48Z%22%2C%22Name%22%3A%22Walk%20the%20Units%22%2C%22PreDoneUID%22%3A0%2C%22TID%22%3A5%2C%22TLID%22%3A2%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A5%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/task/1/5" "request" "e2"  "WebService--Update_Task"
echo "%7B%22cmd%22%3A%22delete%22%7D" > request
dojsonPOST "http://localhost:8270/v1/task/1/4" "request" "e3"  "WebService--Delete_Task"


echo "RENTROLL SERVER STOPPED"
stopRentRollServer
logcheck
