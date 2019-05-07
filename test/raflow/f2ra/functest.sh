#!/bin/bash

TESTNAME="Flow2RA"
TESTSUMMARY="Test Flow data to permanent tables"
DBGENDIR=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"
CATRML="../../../tools/catrml/catrml"

#SINGLETEST=""  # This runs all the tests

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 10/24/2018"

# echo "SINGLETEST = ${SINGLETEST}"


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
#   7.  Lease Status should be updated to reflect the new RA
#------------------------------------------------------------------------------
TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    RAID1REFNO="UJF64M3Y28US5BHW5400"
    RAIDAMENDEDID="2"

    echo "Create new database... x${TFILES}.sql"
    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    # Send the command to change the flow to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAID1REFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "${TFILES}${STEP}"  "WebService--Action-setTo-ACTIVE"

    # Generate an assessment report from Aug 1 to Oct 1. The security deposit
    # assessment for RAID 1 should no longer be present
    docsvtest "${TFILES}${STEP}" "-G ${BUD} -g 8/1/18,10/1/18 -L 11,${BUD}" "Assessments-2018-AUG"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%228%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # validate lease status
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"1/1/2017","searchDtStop":"1/1/2021","Bool1":false}' > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST b
#  This is just like test a except that periods from Feb through July are
#  closed. This means that the reversal entries will need to be made on
#  Aug 1.
#
#  Scenario:
#  RAID  1 - AgreementStart = 2/13/2018,  AgreementStop = 6/13/2020
#  RAID  2 - AgreementStart = 6/13/2018,  AgreementStop = 3/1/2020
#            Verify that correcting entries are made on Aug 1.
#
#  Graphically shown here: https://docs.google.com/presentation/d/1YO6DWzn_KFB9h2xjOoItrAxaxhokQ6cwgymEdkfMmJw/edit#slide=id.g408d3d1457_4_1
#
#  Expected Results:
#   1.  June rent reversed and broken into 2 separate new norecur assessments:
#       One assigned to RAID 1 for  6/1 thru 6/12 → snap to Aug 1 due to closed period
#       One assigned to RAID 2 for 6/12 thru 6/30 → snap to Aug 1 due to closed period
#
#   2.  July rent reversed and a new rent assessment created and assigned to
#       RAID 2 for July → snap to Aug 1 due to closed period
#
#   3.  August rent assessment created on Aug 1.
#
#   4.  September rent assessment created on Sep 1
#
#   5   Lease Status should be updated to reflect the new RA.
#------------------------------------------------------------------------------
TFILES="b"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer


    RAIDREFNO="5R6I7HQM1M1922LD35HH"
    RAIDAMENDEDID="2"

    # Send the command to change the RefNo to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "${TFILES}${STEP}"  "WebService--Backdated-RA-Amendment"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi


#------------------------------------------------------------------------------
#  TEST c
#  Further modifies the database created in TEST b.  Changes the rent to
#  $1100/month starting 7/21.
#
#  Scenario:
#  RAID  1 - AgreementStart = 2/13/2018,  AgreementStop = 6/13/2020
#  RAID  2 - AgreementStart = 6/13/2018,  AgreementStop = 7/21/2020
#  RAID  3 - AgreementStart = 7/21/2018,  AgreementStop = 3/1/2020
#            Verify that correcting entries are made on Aug 1.
#
#  Graphically shown here:
#
#  Expected Results:
#   1.
#   2.
#   3.
#   4.
#------------------------------------------------------------------------------
TFILES="c"
STEP=0
# if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
#     echo "Create new database... x${TFILES}.sql"
#
# stopRentRollServer
# mysql --no-defaults rentroll < x${TFILES}.sql
# startRentRollServer
#
#
#     RAIDREFNO="V91682OU9DNAST5K262A"
#     RAIDAMENDEDID="3"
#
#     # Send the command to change the RefNo to Active:
#     echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
#     dojsonPOST "http://localhost:8270/v1/raactions/1/2" "request" "${TFILES}${STEP}"  "WebService--Backdated-RA-Amendment-with-rent-change"
#
#     # Generate a payor statement -- ensure that 2 RAs are there and have correct
#     # info.
#     # echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
#     # dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"
# fi


#------------------------------------------------------------------------------
#  TEST d
#  Extends a rental agreement that is going to expire.  Also raises the rent
#  and handles the addition of a pet.
#
#  Scenario:
#  RAID  1 - AgreementStart = 2/13/2017,  AgreementStop = 1/1/2018
#  RAID  3 - AgreementStart = 1/1/2018,   AgreementStop = 1/1/2019
#  Add a cat.
#
#  Expected Results:
#   1.  New rent assessment for $1100/month
#   2.  Pet Fee assessment for $50
#   3.  Pet Rent assessment for $10/month
#   4.  RAID 1 is terminated
#   5.  The payor statement for 9/1 - 9/30 should show $10,040 as the Balance
#   6.  Lease Status should be updated to reflect the new RA.
#------------------------------------------------------------------------------
TFILES="d"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer


    RAIDREFNO="1RQTH0A0EO2JD003475M"
    RAIDAMENDEDID="3"

    # Send the command to change the RefNo to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--updated-rental-agreement"

    # Payor statement -- 2 RAs, Balance should be 0 for RA1, $10,040 for RA3
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%229%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST e
#  Normal Lease Extension
#
#  In this example, 2 months (or more) prior to AgreementStop, create an updated
#  amendment to renew the lease for another year.  When completed, the current
#  and the amended RA shound both be Active.
#
#  (ensuring that RentalAgreements that have expired have been changed to the
#  the Terminated state is a check that will be added to Close period)
#
#  Scenario:
#  Amend the existing rental agreement so that the amended RA starts immediately
#  after the old RA stops
#  Initial RA term:  1/1/2018 - 12/1/2018
#
#  Expected Results:
#   1. Walk through the entire approval flow...
#      a) changes the dates on the refno version
#      b) Add a rent assessment
#      c) go to first approval
#      d) approve to go to second approval
#      e) approve to set move in date
#      f) set the move in date
#      g) activate the new rental agreement
#
#   2. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="e"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer


    # #----------------------------------------------------------------
    # # put RA 1 into Edit mode...
    # #----------------------------------------------------------------
    # echo "%7B%22cmd%22%3A%22get%22%2C%22UserRefNo%22%3Anull%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22FlowType%22%3A%22RA%22%7D" > request
    # dojsonPOST "http://localhost:8270/v1/flow/1/0" "request" "${TFILES}${STEP}"  "WebService--edit-RA"
    # RAIDREFNO=$(cat ${TFILES}0 | grep UserRefNo | awk '{print $2}'|sed 's/"//g')

    DTSTART="12%2F1%2F2018"
    DTSTOP="11%2F30%2F2019"
    RAIDREFNO="DUO6X142Z3GC1597BHG0"

    #----------------------------------------------------------------
    # Send the command to change the Dates.
    # Note the use of ${DTSTART} and ${DTSTOP} in the echo command
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22dates%22%2C%22BID%22%3A1%2C%22Data%22%3A%7B%22CSAgent%22%3A209%2C%22RentStop%22%3A%22${DTSTOP}%22%2C%22RentStart%22%3A%22${DTSTART}%22%2C%22AgreementStop%22%3A%22${DTSTOP}%22%2C%22AgreementStart%22%3A%22${DTSTART}%22%2C%22PossessionStop%22%3A%22${DTSTOP}%22%2C%22PossessionStart%22%3A%22${DTSTART}%22%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/flow/1/1" "request" "${TFILES}${STEP}"  "WebService--update-dates"

    # for debugging, it's nice to have the database in this state
    mysqldump --no-defaults rentroll > "xxx${TFILES}.sql"

    #----------------------------------------------------------------
    # Send the command add a Rent assessment definition
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22rentables%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22RID%22%3A1%2C%22Fees%22%3A%5B%7B%22RentCycleText%22%3A%22Monthly%22%2C%22ProrationCycleText%22%3A%22Daily%22%2C%22recid%22%3A1%2C%22TMPASMID%22%3A0%2C%22ASMID%22%3A0%2C%22ARID%22%3A40%2C%22ARName%22%3A%22Rent%20ST000%22%2C%22ContractAmount%22%3A1100%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22Start%22%3A%22${DTSTART}%22%2C%22Stop%22%3A%22${DTSTOP}%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22%22%7D%5D%2C%22RTID%22%3A1%2C%22RTFLAGS%22%3A4%2C%22SalesTax%22%3A0%2C%22RentCycle%22%3A6%2C%22TransOccTax%22%3A0%2C%22RentableName%22%3A%22Rentable001%22%2C%22AtSigningPreTax%22%3A0%2C%22recid%22%3A1%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%5D%7D" > request
    dojsonPOST "http://localhost:8270/v1/flow/1/1" "request" "${TFILES}${STEP}"  "WebService--add-rent-assessment"

    #----------------------------------------------------------------
    # Validate the RA-Flow, which automatically puts the Flow into
    # PendingFirstApproval if successful
    #----------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/validate-raflow/1/1" "request" "${TFILES}${STEP}"  "WebService--validate"

    #----------------------------------------------------------------
    # First Approver approves...
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Mode%22%3A%22State%22%2C%22Decision1%22%3A1%2C%22DeclineReason1%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Approver1"

    #----------------------------------------------------------------
    # Second Approver approves...
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Mode%22%3A%22State%22%2C%22Decision2%22%3A1%2C%22DeclineReason2%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Approver2"

    #----------------------------------------------------------------
    # Set move-in date
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Mode%22%3A%22State%22%2C%22DocumentDate%22%3A%22${DTSTART}%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--MoveInDate"


    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    # validate lease status
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"1/1/2017","searchDtStop":"1/1/2021","Bool1":false}' > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"

fi

#------------------------------------------------------------------------------
#  TEST f
#  Lease Holdover
#
#  Adjust the end date of the RentStop forward in time. (optional: adjust
#  PossessionStop). This should only update the existing RA.  It should extend
#  any recurring assessment definition with a stop date that matched the
#  RA RentStop date.  That is, new assessments will net be created.
#
#  RAID 1 - 2/13/2017 12/31/2017
#  RAID 2 - 1/1/2018 - 2/28/2018  -  RentStop and PossessionStop Only
#
#  Expected Results:
#  1. The Rental Agreement will be amended. The amendment RentStop will be
#     extended out to the end of Feb 2018.
#
#  2. Recurring assessments end on the RentStop date
#  3. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="f"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer


    RAIDREFNO="C8VV07E4FJI983MC5Z76"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    #---------------------------------------------------------------------------
    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    #---------------------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST g
#  Move RentStop back in time
#
#  Scenario:
#  We start with a Rental Agreement from 1/1/2018 to 12/31/2018.  For some
#  reason we wish to keep the agreement end date at 12/31/2018 but we want
#  to stop all the rent payments after 8/31/2018.  To do this we simply
#  change the RentStop date on the RentalAgreement to 8/31/2018.
#  In this example, no periods have been closed.
#
#  In this example we also add a pet -- no particular reason for this.
#
#  For this example, the approach will be to set the rental agreement dates
#  as follows:
#     Agreement Start/Stop:  1/1/2018 - 12/31/2018
#     Rent Start/Stop:       1/1/2018 - 8/31/2018
#     Possession Start/Stop: 1/1/2018 - 8/31/2018
#
#  Expected Results:
#  1. The Rental Agreement will be amended. The amendment will have the dates
#     just as described above.
#  2. All assessments for RA 1 should be reversed and all active
#     assessment should be for RAID 2.
#  3. $9000 of rent assessments should be reversed on RAID 1, $8080 of
#     assessments should be charged to RAID 2. The extra $80 is due to
#     the addidion of the pet.
#  4. There should be $9000 of unallocated funds for RA 2
#  5. Recurring assessments end on the RentStop date
#  6. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="g"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="VYGPF6A5V81V9S66H032"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    #---------------------------------------------------------------------------
    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    #---------------------------------------------------------------------------
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST h
#  Move RentStop back in time -- into a closed period.
#  This is exactly like test g except that Jan - Aug are closed periods.
#
#  Scenario
#  On Sep 19, 2018 set the rental agreement dates as follows:
#     Agreement Start/Stop:  1/1/2018 - 12/31/2018
#     Rent Start/Stop:       1/1/2018 - 7/31/2018
#     Possession Start/Stop: 1/1/2018 - 7/31/2018
#
#  Expected Results:
#  1. The Rental Agreement will be amended. The amendment will have the dates
#     just as described above.
#  2. All assessments for RA 1 should be reversed and
#     all the RAs should be for RAID 2.
#  3. $9000 of rent should be reversed on RAID 1, $8080 should be charged
#     to RAID 2 due to the addidion of the pet.
#  4. Validate Lease Status
#
#------------------------------------------------------------------------------
TFILES="h"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"

    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="227ZM0HGT05MQC41XM89"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    #---------------------------------------------------------------------------
    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    #---------------------------------------------------------------------------
    # %7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2018%22%2C%22searchDtStop%22%3A%227%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST i
#  The RentalAgreement (RAID=1) for Pablo Pearson (TCID=1) has 1 Pet, 2 cars,
#  and 1 Rentable (RID=1).  It's term is 2/13/2018 - 2/29/2020.  The Assessments
#  for October 2018 are:
#      ASM 28   Rent Non-Taxable - $1000
#      ASM 29   Pet Rent
#
#  Scenario
#  Amend the rental agreement on 10/18/2018
#
#  Expected Results:
#  1. Rental Agreement 1
#      - Stop recurring assessments ASM 1,3 on 10/18/2018
#      - Reverse both ASM 28 and 29
#      - New prorated Rent NonTaxable for 17 of 31 days
#      - New prorated Pet Rent for 17 of 31 days
#  2. RA 6
#      - New prorated Rent NonTaxable for 14 of 31 days
#      - New prorated Pet Rent for 14 of 31 days
#
#  3. Validate Lease Status
#
#------------------------------------------------------------------------------
TFILES="i"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"

    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="FV9M2N2695EFPG795KQ4"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%2210%2F1%2F2018%22%2C%22searchDtStop%22%3A%2210%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "${TFILES}${STEP}"  "PayorStatement--StmtInfo"

    # 1.  Generate text payor statement
    RRDATERANGE="-j 2018-10-01 -k 2018-11-01"
    dorrtest "${TFILES}${STEP}" "${RRDATERANGE} -x -b ${BUD} -r 23,1" "PayorReport"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST j
#  RA3 is being extended. It is set to stop on 2/29/2020.  We extend it for
#  another year.  The extension term is 3/1/2020 to 2/28/2021.
#
#  Scenario
#  Execute changes and cause the new RA to be created
#
#  Expected Results:
#  1. Term dates on the new RA must be 3/1/2020 to 2/28/2021
#
#  2. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="j"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"

    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="K6HV8GB4E222RJX864F4"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/3" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    TESTCOUNT=$((TESTCOUNT + 1))
    echo -n "PHASE  ${TESTCOUNT}: Checking start date... "
    AGRSTART=$(cat serverreply | python -m json.tool | grep AgreementStart | sed 's/"//g' |sed 's/,//g' | awk '{print $2}')
    if [ ${AGRSTART} != "3/1/2020" ]; then
        failmsg
        if [ "${ASKBEFOREEXIT}" = "1" ]; then
            pause ${1}
        else
            exit 2
        fi
    fi
    echo "PASSED"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2022%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST k
#  Bug was found.  Original RA (1): Term is 1/1/2018 - 1/1/2019.  Amend it on
#  10/15/2018 -  Net result RA1 ended up with PRAID=2  And RA2 had PRAI1. !!!
#
#  Scenario
#  PRAID is updated on Original RA and incorrectly.
#
#  Expected Results:
#  1. Original RA should retain PRAID = 0
#  2. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="k"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="WK4UD23XJ9S41831YZ6Y"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/3" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    mysqlverify "${TFILES}${STEP}" "RentalAgreements"	"select RAID,RATID,BID,PRAID,ORIGIN,AgreementStart,AgreementStop from RentalAgreement;"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST l
#  Bug was found.  Original RA (1): Term is 1/1/2018 - 1/1/2019.  Amend it on
#  10/15/2018.  Amend again on 10/23/2018 --- the bug found: all Assessments
#  were reversed and we were left with no active assessments.  This test
#  validates the fixes necessary to make this scenario work correctly
#
#  Scenario
#  Update the RA with an amendment on 10/27/2018
#
#  Expected Results:
#  1. The month of October should have 3 Rental Agreements in the Payor
#     Statement.
#  2. Prorationed assessments should be identified for each agreement.
#  3. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="l"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database... x${TFILES}.sql"

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="T0LVI9R4NH73DBS17008"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/3" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    #----------------------------------------------------------------
    # Generate text payor statement
    #----------------------------------------------------------------
    RRDATERANGE="-j 2018-10-01 -k 2018-11-01"
    dorrtest "${TFILES}${STEP}" "${RRDATERANGE} -x -b ${BUD} -r 23,1" "PayorReport"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST m
#  Bug was found. Backdate rental agreement into closed period. When the
#  assessments were created during the Flow2RA conversion, the prorated rent
#  assessment for the month of October for the amendment was missed. And
#  the rent assessments for the month of November were missed.  This test is
#  used to verify that bug fixes to make this scenario work.
#
#  Scenario
#  RA 1 is from 7/1/2018 to 6/30/2019
#  Periods for 7/31, 8/31, 9/30, and 10/31 have been closed.
#  Amend RA 1 on Nov 6. backdate the amendment to 10/15/2018
#
#  Expected Results:
#  1. Ensure that Prorated assessments for October Rent and PetRent are
#     created October for both RA 1 and RA 2.
#  2. Ensure that the rent assessments for November are created for RA 2
#  3. Validate Lease Status
#------------------------------------------------------------------------------
TFILES="m"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database... x${TFILES}.sql"

    RENTROLLSERVERNOW="-testDtNow 11/6/2018"
    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="5R86U008G1WNM7MU3FQ7"

    #----------------------------------------------------------------
    # Make the updated RefNo an Active RA
    #----------------------------------------------------------------
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"

    #----------------------------------------------------------------
    # Generate text payor statement
    #----------------------------------------------------------------
    RRDATERANGE="-j 2018-10-01 -k 2018-12-01"
    dorrtest "${TFILES}${STEP}" "${RRDATERANGE} -x -b ${BUD} -r 23,1" "PayorReport"

    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"
fi

#------------------------------------------------------------------------------
#  TEST n
#  Ensure proper handling of an amended RA where the amendment occurs
#  disjoint and prior to the RA it amends.
#
#  Scenario
#  RA 1 is from 5/1/2019 to 8/1/2019
#  Amend this RA to 1/1/2019 - 1/3/2019
#
#  on a timeline:
#
#  Expected Results:
#   see below in specific tests
#------------------------------------------------------------------------------
TFILES="n"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database... x${TFILES}.sql"

    RENTROLLSERVERNOW="-testDtNow 4/25/2019"
    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="JUYIPZL6L4Q74Q3DI2N8"

    #----------------------------------------------------------------
    # Active RentalAgreement is 1/1/2019
    #----------------------------------------------------------------
    encodeRequest '{"UserRefNo":"JUYIPZL6L4Q74Q3DI2N8","RAID":1,"Version":"refno","Action":4,"Mode":"Action"}' > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"
    # validate lease status
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2017%22%2C%22searchDtStop%22%3A%221%2F1%2F2021%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"

fi

#------------------------------------------------------------------------------
#  TEST o
#  Ensure proper handling of an amended RA where the amendment occurs
#  disjoint and after to the RA it amends.
#
#  Scenario
#  RA 1 is from 5/1/2019 to 7/31/2019
#  RA 2 is from 12/1/2019 to 1/1/2020
#  New RA not yet active from 11/1/2019 to 12/1/2019
#  Amend RA 1 to 9/1/2019 to 11/1/2019
#  Then create a new RA from 11/1/2019 - 12/1/2019
#
#  on a timeline:
#
#  Expected Results:
#   There will be a gap between RA 1 and its amended version. RA 1 should
#   be cancelled, the lease status in the gap should be LEASESTATUSnotleased.
#   There will be a gap between the amendment (RA 3) and RA 2. This gap should
#   be LEASESTATUSreserved - 11/1/2019 - 12/1/2019.
#
#   Next, create the new RA from 11/1 - 12/1. The key check is to ensure that
#   the LeaseStatus record right after the one from 11/1 - 12/1 does not
#   get overwritten with either "not leased" or "reserved" status because
#   that record corresponds to RA2
#------------------------------------------------------------------------------
TFILES="o"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database... x${TFILES}.sql"

    RENTROLLSERVERNOW="-testDtNow 4/25/2019"
    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    RAIDREFNO="JV71FYOQUDPZZ3Q0DM3S"

    #----------------------------------------------------------------
    # Make the new rental agreement amendment (RAID = 3) active
    #----------------------------------------------------------------
    encodeRequest '{"UserRefNo":"JV71FYOQUDPZZ3Q0DM3S","RAID":1,"Version":"refno","Action":4,"Mode":"Action"}' > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/1" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"
    # validate lease status
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"1/1/2017","searchDtStop":"1/1/2021","Bool1":false}' > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"

    # FlowID 3  -->  UserRefNo = JV7KM0D9SHCOP754WC6O
    #----------------------------------------------------------------
    # Make the new rental agreement from 11/1 - 12/1 active
    #----------------------------------------------------------------
    encodeRequest '{"UserRefNo":"JV7KM0D9SHCOP754WC6O","RAID":0,"Version":"refno","Action":4,"Mode":"Action"}' > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/3" "request" "${TFILES}${STEP}"  "WebService--Activate-RefNo"
    # validate lease status
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"1/1/2017","searchDtStop":"1/1/2021","Bool1":false}' > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "LeaseStatus"

    #----------------------------------------------------------------
    # Terminate the rental agreement from 11/1 - 12/1 active
    #----------------------------------------------------------------
    # encodeRequest '{"UserRefNo":"JV7KM0D9SHCOP754WC6O","RAID":0,"Version":"refno","Action":6,"Mode":"Action"}' > request
    # dojsonPOST "http://localhost:8270/v1/raactions/1/3" "request" "${TFILES}${STEP}"  "WebService--Terminate-RefNo"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
