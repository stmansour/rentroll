#!/bin/bash

TESTNAME="Flow2RA"
TESTSUMMARY="Test Flow data to permanent tables"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  An existing rental agreement (RAID=1) is being amended. This test
#  verifies that all assessments from RAID=1 to the new RAID (2) are correct.
#  It also verifies that payments are properly filtered. For example,
#  I the original agreement there is a Security Deposit request in
#  September. This security deposit is not in the fees for the amended
#  Rental Agreement, so it should be reversed in the old rental agreement
#
#  Scenario:
#  RAID 1 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#  RAID 2 - AgreementStart = 8/8/2018,   AgreementStop = 3/1/2020
#           The flow used to create RAID 2 has no links between its fees and
#           the assessments in RAID 1. So, the handling tests how "unlinked"
#           assessments are handled when amending a rental agreement.
#
#  Expected Results:
#   1.  All RAID 1 recurring assessment definitions that overlap the period
#       8/8/2018 - 3/1/2020 must have their stop date set to 8/8/201
#   2.  The RAID 1 rent assessment has already occured, and it has been paid.
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  There is a Security Deposit assessment (ASMID=20) due on 9/14 in the
#       old rental agreement. It is not in the fees list for the RefNo, so it
#       should be reversed
#------------------------------------------------------------------------------

# Send the command to change the flow to Active:
#echo "%7B%22UserRefNo%22%3A%22G4OT34LK1266DWUQ765I%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request

echo "%7B%22UserRefNo%22%3A%22K3GO9UEJE0UJ010F7382%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a0"  "WebService--Action-setTo-ACTIVE"

docsvtest "a1" "-G ${BUD} -g 8/1/18,10/1/18 -L 11,${BUD}" "Assessments-2018-AUG"

echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%228%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "a2"  "PayorStatement--StmtInfo"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
