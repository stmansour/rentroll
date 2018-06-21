#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..
BINDIR=${TOP}/tmp/rentroll

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

cp config.json ${BINDIR}/config.json

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
#   a0 - The db we read has 1 pre-defined TaskListDefinition
#   a1 - Write a new TaskListDefinition
#   a2 - Update a TaskListDefinition
#   a3 - This search should return 2 matches
#   a4 - Delete it - which means set the FLAGS to make it inactive
#   a5 - Read back TLDID 2, ensure that Name and FLAGS were updated
#   a6 - This search should only return 1 match because TLDID 2 was made inactive
#   a7 - Read back the predefined TaskListDefinition
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tlds/1" "request" "a0"  "WebService--Search_TaskListDefinitions"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22Cycle%22%3A6%2C%22Epoch%22%3A%221%2F1%2F2018%22%2C%22EpochDue%22%3A%221%2F31%2F2018%22%2C%22EpochPreDue%22%3A%221%2F20%2F2018%22%2C%22DurWait%22%3A86400000000000%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Period%20Close%22%2C%22TLDID%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/0" "request" "a1"  "WebService--Insert_TaskListDefinitions"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22TLDID%22%3A2%2C%22BID%22%3A1%2C%22Cycle%22%3A6%2C%22Epoch%22%3A%221%2F1%2F2018%22%2C%22EpochDue%22%3A%221%2F31%2F2018%22%2C%22EpochPreDue%22%3A%221%2F20%2F2018%22%2C%22DurWait%22%3A86400000000000%2C%22EmailList%22%3A%22bounce%40simulator.amazonses.com%22%2C%22FLAGS%22%3A0%2C%22Name%22%3A%22Tucasa%20Apts%20Period%20Close%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/2" "request" "a2"  "WebService--Update_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tlds/1" "request" "a3"  "WebService--Search_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22delete%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/2" "request" "a4"  "WebService--Delete_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/2" "request" "a5"  "WebService--Read_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tlds/1" "request" "a6"  "WebService--Search_TaskListDefinitions"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tld/1/1" "request" "a7"  "WebService--Get_TaskListDefinition_1"

#------------------------------------------------------------------------------
#  TEST b
#  Basic TaskDescriptor Test Suite
#
#  Scenario:
#		Try all the readers, inserters, updaters and deleters.
#
#  Expected Results:
#   b0 - The db we read has 3 pre-defined TaskListDefinition
#   b1 - Insert a new TaskDescriptor (4)
#   b2 - Read back the newly inserted task descriptor (4)
#   b3 - Modify Task Descriptor 4
#   b4 - Read back the updated task descriptor to make sure it worked
#   b5 - Delete Descriptor 4
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tds/1/1" "request" "b0"  "WebService--Search_TaskDescriptors"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22TDName%22%3A%22Validate%20Washing%20Room%20Totals%22%2C%22TDID%22%3A0%2C%22TLDID%22%3A1%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A2%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/0" "request" "b1"  "WebService--Insert_TaskDescriptor"
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%223%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F31%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/td/1/4" "request" "b2"  "WebService--Read_TaskDescriptor"
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A0%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22TDName%22%3A%22Validate%20Laundry%20Totals%22%2C%22TDID%22%3A4%2C%22TLDID%22%3A1%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A2%7D%7D" > request
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
#	Break all the bizlogic rules and make sure it gets caught.
#
#  Expected Results:
#   c0 - Attempt to assign a Descriptor to a non-existent TaskListDefinition
#   c1 - Attempt to assign a Descriptor to an invalid business
#   c2 - Attempt to save a Descriptor with no name
#------------------------------------------------------------------------------
echo "%7B%22recid%22%3A0%2C%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A0%2C%22EpochDue%22%3A%222018-01-31T20%3A00%3A00Z%22%2C%22EpochPreDue%22%3A%222018-01-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22TDName%22%3A%22break%20bizlogic%22%2C%22TDID%22%3A0%2C%22TLDID%22%3A7981%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A0%7D%7D" > request
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
#	Break all the bizlogic rules and make sure it gets caught.
#
#  Expected Results:
#   d0 - Insert a new instance of TLD 1
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
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%223%2F1%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tls/1/0" "request" "d3"  "WebService--Search_TaskList"
echo "%7B%22cmd%22%3A%22delete%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tl/1/2" "request" "d4"  "WebService--Delete_TaskList"

#------------------------------------------------------------------------------
#  TEST e
#  TaskList Test Suite
#
#  Scenario:
#	Break all the bizlogic rules and make sure it gets caught.
#
#  Expected Results:
#   e0 - Get the task list for TaskList 2
#   e1 - Get a specific task
#   e2 - Update a task
#   e3 - delete a task
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%7D" > request
dojsonPOST "http://localhost:8270/v1/tasks/1/1" "request" "e0"  "WebService--Get_Tasks"
echo "%7B%22cmd%22%3A%22get%22%7D" > request
dojsonPOST "http://localhost:8270/v1/task/1/5" "request" "e1"  "WebService--Get_Task"
echo "%7B%22cmd%22%3A%22save%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22Comment%22%3A%22This%20is%20an%20update%22%2C%22CreateBy%22%3A0%2C%22CreateTS%22%3A%222018-03-21T20%3A07%3A00Z%22%2C%22DoneUID%22%3A0%2C%22DtDone%22%3A%222018-02-28T17%3A00%3A00Z%22%2C%22DtDue%22%3A%222018-02-28T20%3A00%3A00Z%22%2C%22DtPreDone%22%3A%222018-02-19T23%3A12%3A00Z%22%2C%22DtPreDue%22%3A%222018-02-20T20%3A00%3A00Z%22%2C%22FLAGS%22%3A0%2C%22LastModBy%22%3A0%2C%22LastModTime%22%3A%222018-03-21T20%3A07%3A00Z%22%2C%22Name%22%3A%22Walk%20the%20Units%22%2C%22PreDoneUID%22%3A0%2C%22TID%22%3A5%2C%22TLID%22%3A2%2C%22Worker%22%3A%22Manual%22%2C%22recid%22%3A5%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/task/1/5" "request" "e2"  "WebService--Update_Task"
echo "%7B%22cmd%22%3A%22delete%22%7D" > request
dojsonPOST "http://localhost:8270/v1/task/1/4" "request" "e3"  "WebService--Delete_Task"

#------------------------------------------------------------------------------
#  TEST f
#  TaskList report
#
#  Scenario:
#      Print a tasklist report
#
#  Expected Results:
#      f0 - Get the task list for TaskList 2
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%7D" > request
dorrtest "f0" "-b ${BUD} -r 25,1" "Tasklist"

#------------------------------------------------------------------------------
#  TEST g
#  TaskList close period
#
#  Scenario:
#		Call close period based on task list 1
#
#  Expected Results:
#	f0 - Get the task list for TaskList 2
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22get%22%7D" > request
dojsonPOST "http://localhost:8270/v1/closeperiod/1" "request" "g0"  "WebService--Get_Tasks"


#------------------------------------------------------------------------------
#  TEST x
#  Save a new user with some Prospect information that will be encrypted.
#  This kind of a test is a little out of place here, but it doesn't hurt
#  anything, and the database used had a small number of Transactants, so
#  debugging was easy.
#
#  Scenario:
#		Save a Transactant with lots of information in Prospect, User,
#       and Payor parts. Read back info. Make sure the encrypted / decrypted
#       values (SSN, DriversLicense) work correctly.
#
#  Expected Results:
#	x0 - Save the user
#   x1 - Update the user
#   x2 - Read the user 
#------------------------------------------------------------------------------

# INSERT
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22transactantForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22TCID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22%22%2C%22NLID%22%3A0%2C%22FirstName%22%3A%22Grazyna%22%2C%22MiddleName%22%3A%22Carol%22%2C%22LastName%22%3A%22Nieves%22%2C%22PreferredName%22%3A%22Rocco%22%2C%22CompanyName%22%3A%22Jack%2BIn%2BThe%2BBox%2BInc.%22%2C%22IsCompany%22%3Afalse%2C%22PrimaryEmail%22%3A%22GNieves628%40yahoo.com%22%2C%22SecondaryEmail%22%3A%22GNieves7765%40yahoo.com%22%2C%22WorkPhone%22%3A%22(735)%2B303-9714%22%2C%22CellPhone%22%3A%22(384)%2B789-3137%22%2C%22Address%22%3A%2241344%2BCedar%22%2C%22Address2%22%3A%22%22%2C%22City%22%3A%22Normal%22%2C%22State%22%3A%22LA%22%2C%22PostalCode%22%3A%2298081%22%2C%22Country%22%3A%22USA%22%2C%22CompanyAddress%22%3A%2283853%2BJackson%22%2C%22CompanyCity%22%3A%22Glendale%22%2C%22CompanyState%22%3A%22AL%22%2C%22CompanyPostalCode%22%3A%2284059%22%2C%22CompanyEmail%22%3A%22JackInTheBoxIncG157%40aol.com%22%2C%22CompanyPhone%22%3A%22(477)%2B124-9172%22%2C%22Website%22%3A%22%22%2C%22Occupation%22%3A%22%22%2C%22DesiredUsageStartDate%22%3A%226%2F15%2F2018%22%2C%22RentableTypePreference%22%3A0%2C%22FLAGS%22%3A0%2C%22Approver%22%3A152%2C%22DeclineReasonSLSID%22%3A0%2C%22OtherPreferences%22%3A%22%22%2C%22FollowUpDate%22%3A%226%2F17%2F2018%22%2C%22CSAgent%22%3A208%2C%22OutcomeSLSID%22%3A0%2C%22FloatingDeposit%22%3A0%2C%22RAID%22%3A0%2C%22Points%22%3A0%2C%22DateofBirth%22%3A%223%2F24%2F1977%22%2C%22EmergencyContactName%22%3A%22Justa%2BBolton%22%2C%22EmergencyContactAddress%22%3A%2284978%2B12th%2CStamford%2CNE%2B27887%22%2C%22EmergencyContactTelephone%22%3A%22(488)%2B376-5373%22%2C%22EmergencyContactEmail%22%3A%22JustaB9853%40gmail.com%22%2C%22AlternateAddress%22%3A%221358%2BThirteenth%2CKaneohe%2CRI%2B31847%22%2C%22EligibleFutureUser%22%3Afalse%2C%22Industry%22%3A%22%22%2C%22SourceSLSID%22%3A0%2C%22CreditLimit%22%3A5693%2C%22TaxpayorID%22%3A%2206018360%22%2C%22ThirdPartySource%22%3A99%2C%22EligibleFuturePayor%22%3Atrue%2C%22SSN%22%3A%22123-45-6789%22%2C%22DriversLicense%22%3A%22U1234567%22%2C%22GrossIncome%22%3A532.64%2C%22CommissionableThirdParty%22%3A%22RexFord%20Third%20Party%22%2C%22CurrentAddress%22%3A%2255732%20Hickory%2C%20Pasadena%2C%20AR%2002081%22%2C%22CurrentLandLordName%22%3A%22Rebecca%20Chambers%22%2C%22CurrentLandLordPhoneNo%22%3A%22(888)%20957-9596%22%2C%22CurrentLengthOfResidency%22%3A%222%20years%208%20months%22%2C%22CurrentReasonForMoving%22%3A125%2C%22PriorAddress%22%3A%2296276%20Eighth%2C%20Marysville%2C%20FL%2041318%22%2C%22PriorLandLordName%22%3A%22Rae%20Nielsen%22%2C%22PriorLandLordPhoneNo%22%3A%22(785)%20628-7712%22%2C%22PriorLengthOfResidency%22%3A%222%20years%202%20months%22%2C%22PriorReasonForMoving%22%3A128%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/person/1/0" "request" "x0" "WebService--NewTransactant"

# UPDATE
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22transactantForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22TCID%22%3A16%2C%22BID%22%3A1%2C%22BUD%22%3A%22%22%2C%22NLID%22%3A0%2C%22FirstName%22%3A%22Grazyna%22%2C%22MiddleName%22%3A%22Carol%22%2C%22LastName%22%3A%22Nieves%22%2C%22PreferredName%22%3A%22Rocco%22%2C%22CompanyName%22%3A%22Jack%2BIn%2BThe%2BBox%2BInc.%22%2C%22IsCompany%22%3Afalse%2C%22PrimaryEmail%22%3A%22GNieves628%40yahoo.com%22%2C%22SecondaryEmail%22%3A%22GNieves7765%40yahoo.com%22%2C%22WorkPhone%22%3A%22(735)%2B303-9714%22%2C%22CellPhone%22%3A%22(384)%2B789-3137%22%2C%22Address%22%3A%2241344%2BCedar%22%2C%22Address2%22%3A%22%22%2C%22City%22%3A%22Normal%22%2C%22State%22%3A%22LA%22%2C%22PostalCode%22%3A%2298081%22%2C%22Country%22%3A%22USA%22%2C%22CompanyAddress%22%3A%2283853%2BJackson%22%2C%22CompanyCity%22%3A%22Glendale%22%2C%22CompanyState%22%3A%22AL%22%2C%22CompanyPostalCode%22%3A%2284059%22%2C%22CompanyEmail%22%3A%22JackInTheBoxIncG157%40aol.com%22%2C%22CompanyPhone%22%3A%22(477)%2B124-9172%22%2C%22Website%22%3A%22%22%2C%22Occupation%22%3A%22%22%2C%22DesiredUsageStartDate%22%3A%226%2F15%2F2018%22%2C%22RentableTypePreference%22%3A0%2C%22FLAGS%22%3A0%2C%22Approver%22%3A152%2C%22DeclineReasonSLSID%22%3A0%2C%22OtherPreferences%22%3A%22%22%2C%22FollowUpDate%22%3A%226%2F17%2F2018%22%2C%22CSAgent%22%3A208%2C%22OutcomeSLSID%22%3A0%2C%22FloatingDeposit%22%3A0%2C%22RAID%22%3A0%2C%22Points%22%3A0%2C%22DateofBirth%22%3A%223%2F24%2F1977%22%2C%22EmergencyContactName%22%3A%22Justa%2BBolton%22%2C%22EmergencyContactAddress%22%3A%2284978%2B12th%2CStamford%2CNE%2B27887%22%2C%22EmergencyContactTelephone%22%3A%22(488)%2B376-5373%22%2C%22EmergencyContactEmail%22%3A%22JustaB9853%40gmail.com%22%2C%22AlternateAddress%22%3A%221358%2BThirteenth%2CKaneohe%2CRI%2B31847%22%2C%22EligibleFutureUser%22%3Afalse%2C%22Industry%22%3A%22%22%2C%22SourceSLSID%22%3A0%2C%22CreditLimit%22%3A5693%2C%22TaxpayorID%22%3A%2206018360%22%2C%22ThirdPartySource%22%3A99%2C%22EligibleFuturePayor%22%3Atrue%2C%22SSN%22%3A%22123-45-6789%22%2C%22DriversLicense%22%3A%22U1234567%22%2C%22GrossIncome%22%3A324.56%2C%22CommissionableThirdParty%22%3A%22Islo%20bella%20Third%20Party%22%2C%22CurrentAddress%22%3A%2212232%20South%20Dakota%2C%20Yonkers%2C%20FL%2011528%22%2C%22CurrentLandLordName%22%3A%22Tyrone%20Clemons%22%2C%22CurrentLandLordPhoneNo%22%3A%22(466)%20290-7809%22%2C%22CurrentLengthOfResidency%22%3A%2210%20years%207%20months%22%2C%22CurrentReasonForMoving%22%3A117%2C%22PriorAddress%22%3A%227856%20N%20400%2C%20Vallejo%2C%20AK%2086258%22%2C%22PriorLandLordName%22%3A%22Errol%20Foley%22%2C%22PriorLandLordPhoneNo%22%3A%22(985)%20102-3688%22%2C%22PriorLengthOfResidency%22%3A%222%20years%203%20months%22%2C%22PriorReasonForMoving%22%3A127%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/person/1/16" "request" "x1" "WebService--UpdateTransactant"

# GET
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22TCID%22%3A16%7D" > request
dojsonPOST "http://localhost:8270/v1/person/1/16" "request" "x2" "WebService--GetTransactant"

#------------------------------------------------------------------------------
#  FINISH
#------------------------------------------------------------------------------
echo "RENTROLL SERVER STOPPED"
stopRentRollServer
logcheck