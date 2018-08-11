#!/bin/bash

TESTNAME="Flow2RA"
TESTSUMMARY="Test Flow data to permanent tables"
DBGEN=../../tools/dbgen
CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  An existing rental agreement (RAID=1) is being amended. This test
#  verifies that all assessments from RAID=1 to the new RAID (2) are correct.
#
#  Scenario:
#	RAID 1 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#       RAID 2 - AgreementStart = 8/8/2018,   AgreementStop = 3/1/2020
#       The flow used to create RAID 2 has no links between its fees and
#                the assessments in RAID 1. So, the handling tests how
#                "unlinked" assessments are handled when amending a rental
#                agreement.
#
#  Expected Results:
#   1.  RAID 1 and all its assessments must be stopped on 8/8/2018
#   2.	RAID 1 recurring assessments that affect period 8/8/2018 - 3/1/2020
#       must be migrated to RAID 2
#   3.  Prorated payments must be made to cover the partial months for
#       both RAID 1 and RAID 2
#   4.  Unlinked non-recurring assessments are reversed from the old RAID
#       if they are not paid and/or if they are not already reversed.
#------------------------------------------------------------------------------

# Send the command to change the flow to Active:
echo "%7B%22UserRefNo%22%3A%22G4OT34LK1266DWUQ765I%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a0"  "WebService--Action-setTo-ACTIVE"

docsvtest "a1" "-G ${BUD} -g 8/1/18,10/1/18 -L 11,${BUD}" "Assessments-2018-AUG"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
