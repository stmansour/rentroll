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
#  The flow for RAID 1 is updated to the active state causing an amended
#  Rental Agreement (RAID=24) to be created.
#  RAID  1 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#  RAID 24 - AgreementStart = 8/20/2018,  AgreementStop = 3/1/2020
#            The flow used to create RAID 24 has no links between its fees and
#            the assessments in RAID 1. So, the handling tests how "unlinked"
#            assessments are handled when amending a rental agreement.
#
#  Expected Results:
#   1.  All RAID 1 recurring assessment definitions that overlap the period
#       8/8/2018 - 3/1/2020 must have their stop date set to 8/8/201
#   2.  The RAID 1 rent assessment has already occured, and it has been paid.
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  There is a Security Deposit assessment for RAID 1 due on 9/20 in the
#       old rental agreement. It is not in the fees list for the RefNo, so it
#       should be reversed in RAID 1 and not present in RAID 24
#------------------------------------------------------------------------------
RAID1REFNO="T7LYN5K18Z7F756KE64C"
RAIDAMENDEDID="24"

# Send the command to change the flow to Active:
echo "%7B%22UserRefNo%22%3A%22${RAID1REFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a0"  "WebService--Action-setTo-ACTIVE"

# Generate an assessment report from Aug 1 to Oct 1. The security deposit
# assessment for RAID 1 should no longer be present
docsvtest "a1" "-G ${BUD} -g 8/1/18,10/1/18 -L 11,${BUD}" "Assessments-2018-AUG"

# Generate a payor statement -- ensure that 2 RAs are there and have correct
# info.
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%228%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "a2"  "PayorStatement--StmtInfo"

#------------------------------------------------------------------------------
#  TEST b
#  This is just like test a except that the $4500 security deposit assessment
#  from the origin RA (RAID 2) is kept.  Since its time frame falls into that
#  of the amended Rental Agreement, it becomes part of that rental agreement.
#
#  Scenario:
#  RAID  2 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#  RAID 25 - AgreementStart = 8/20/2018,  AgreementStop = 3/1/2020
#            Verify that the Security Deposit on 9/20 is linked to the new
#            rental agreement.
#
#  Expected Results:
#   1.  ASMID 402 (which was charged to RAID 1) should be reversed and a new
#       one should be created (ASMID 412) associated with the amended RAID (25)
#       8/8/2018 - 3/1/2020 must have their stop date set to 8/8/201
#   2.  The RAID 1 rent assessment has already occured, and it has been paid.
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  There is a Security Deposit assessment (ASMID=402) due on 9/20 in the
#       old rental agreement. It is not in the fees list for the RefNo, so it
#       should be reversed
#------------------------------------------------------------------------------
RAID1REFNO="NZXY8FS6NHJ34N383950"
RAIDAMENDEDID="24"

# Send the command to change the RefNo to Active:
echo "%7B%22UserRefNo%22%3A%22${RAID1REFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "b0"  "WebService--Action-setTo-ACTIVE"

# Generate a payor statement -- ensure that 2 RAs are there and have correct
# info.
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "b1"  "PayorStatement--StmtInfo"


# import rr.sql again to test update existing RA
echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

#----------------------------------------------------------------------------------
# TEST z
# This test is for verifying update RAflow of existing RAFlow
#
# Scenario:
# Edit one existing RA Application from its view mode only.
# Update Pet/Vehicle/People information fields value
# After updating information, move RAApplication to state 'Complete Move-In'
#
# Expected Result:
# Check same RA Application flow's data. It must be match with the updated information
#---------------------------------------------------------------------------------

# send command to Edit existing Rental Agreement with RAID: 3
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowType%22%3A%22RA%22%2C%22RAID%22%3A3%2C%22UserRefNo%22%3Anull%2C%22Version%22%3A%22refno%22%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/0" "request" "z0" "Rental Agreement--Edit--RAID:3"

echo "" > request
dojsonPOST "http://localhost:8270/v1/flow/1/0" "request" "z1" "Rental Agreement--RAID:3--Update Flow Information"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
