#!/bin/bash

TESTNAME="Flow2RA"
TESTSUMMARY="Test Flow data to permanent tables"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

SINGLE=""  # This runs all the tests
# SINGLE="b"   # Run just test b

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
#  RAID  2 - AgreementStart = 6/8/2018,  AgreementStop = 3/1/2020
#
#  Expected Results:
#   1.  All RAID 1 recurring assessment definitions that overlap the period
#       6/8/2018 - 3/1/2020 must have their stop date set to 6/8/201
#   2.  The RAID 1 rent assessment has already occured, and it has been paid.
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  All rent assessment instances for the period containing 6/8/2018 and
#       all periods after 6/8/2018 must be Reversed.
#   4.  Rent assessments for the month of June and all months afterwards
#       up to the present must have an instance in the database tied to the
#       new rental agreement
#   5.  Rent for the first period of the change (June 1, 2018) will have
#       a prorated assessment for RAID 1 covering June 1 to 8, and another
#       prorated assessment covering June 8 - 30.
#   6.  Recurring fees will need to be created for the new RA (2). A rent
#       assessment must be added for June, July, and August. The transition
#       month's rent in this case will need to be  prorated to account for
#       days June 8 thru June 30.
#------------------------------------------------------------------------------
if [ "${SINGLE}a" = "a" -o "${SINGLE}a" = "aa" ]; then
    RAID1REFNO="UJF64M3Y28US5BHW5400"
    RAIDAMENDEDID="2"

    echo "Create new database... x0.sql"
    mysql --no-defaults rentroll < x0.sql

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
fi

#------------------------------------------------------------------------------
#  TEST b
#  This is just like test a except that periods from Feb through July are
#  closed. This means that the reversal entries will need to be made on
#  Aug 1.
#
#  Scenario:
#  RAID  1 - AgreementStart = 2/13/2018,  AgreementStop = 6/13/2020
#  RAID 25 - AgreementStart = 6/13/2018,  AgreementStop = 3/1/2020
#            Verify that correcting entries are made on Aug 1.
#
#  Expected Results:
#   1.  WIP
#
#
#   2.
#
#
#   3.
#
#
#------------------------------------------------------------------------------
if [ "${SINGLE}b" = "b" -o "${SINGLE}b" = "bb" ]; then
    echo "Create new database... x1.sql"
    mysql --no-defaults rentroll < x1.sql

    RAIDREFNO="5R6I7HQM1M1922LD35HH"
    RAIDAMENDEDID="2"

    # Send the command to change the RefNo to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "b0"  "WebService--Backdated-RA-Amendment"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "b1"  "PayorStatement--StmtInfo"
fi


#------------------------------------------------------------------------------

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
