#!/bin/bash

TESTNAME="Assessment Instance Generator Worker"
TESTSUMMARY="Test the generation of asm instances"

CREATENEWDB=0
RRDATERANGE="-j 2018-03-01 -k 2018-04-01"
BUD="REX"

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

source ../share/base.sh

WLOG=workasm.log

#------------------------------------------------------------------------------
#  TEST 00
#  Scenario: The database we load has Rental Agreement 2 with a stop date
#       2/1/2018.  Assessment ASMID=3 charges to Rental Agreement 2. It
#       recurs and has a Stop Date of 3/1/2018 (the RA's original stop date).
#       RentalAgreement 2 was modified to end on 2/1/2018. Its assessments were
#       not adjusted.  This test makes sure that if a recurring assessment
#       continues past its associated RentalAgreement stop date, the worker
#       that creates instances of recurring assessments does not create 
#       new assessment instances that go beyond the RentalAgreement's end
#       date.
#
#       When the worker code is called, simulating the date as 2/1/2018. It
#       should NOT add a new assessment instance. 
#
#  Expected Result:
#	1.	When the worker code is called, simulating the date as 2/1/2018. It
#       should NOT add a new assessment instance. So, the Assessment report
#       should NOT show a $3750 instance for Feb.
#------------------------------------------------------------------------------
RRDATERANGE="-j 2018-02-01 -k 2018-03-01"
./workerasm -dt "2/1/2018" > ${WLOG}
res=$(grep "Cannot add new assessment instance on Feb 1, 2018 after RentalAgreement (RA-2) stop date Feb 1, 2018" workerasm.log | wc -l)
if (( res != 1 )); then
	echo "ERROR: assessment instance for 2/1/2018 was not suppressed."
	exit 2
fi
dorrtest "wk0" "${RRDATERANGE} -b ${BUD} -r 24" "Assessments-Feb-2018"

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer


#------------------------------------------------------------------------------
#  TEST 01
#  Scenario: The database we load has Rental Agreement 2 with a stop date
#       2/1/2018.  Assessment ASMID=3 charges to Rental Agreement 2. It
#       recurs and has a Stop Date of 3/1/2018 (the RA's original stop date).
#       In this test we will modify the stop date of RentalAgreement 2 to
#       Jan 15, 2018. This should cause the Assessment 3's Stop date to
#       become Jan 15, 2018 as well.
#
#  Expected Result:
#	1.	After the edit, Assessment 3 stop date should be 1/15/2018
#------------------------------------------------------------------------------

#
# Execute web service command to save RentalAgreement with end date 1/15/2018
#
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22rentalagrForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RAID%22%3A2%2C%22RATID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22NLID%22%3A0%2C%22AgreementStart%22%3A%223%2F1%2F2014%22%2C%22AgreementStop%22%3A%221%2F15%2F2018%22%2C%22PossessionStart%22%3A%223%2F1%2F2014%22%2C%22PossessionStop%22%3A%221%2F15%2F2018%22%2C%22RentStart%22%3A%223%2F1%2F2014%22%2C%22RentStop%22%3A%221%2F15%2F2018%22%2C%22RentCycleEpoch%22%3A%223%2F1%2F2014%22%2C%22UnspecifiedAdults%22%3A0%2C%22UnspecifiedChildren%22%3A0%2C%22Renewal%22%3A%22lease+extension+option%22%2C%22SpecialProvisions%22%3A%22%22%2C%22LeaseType%22%3A0%2C%22ExpenseAdjustmentType%22%3A0%2C%22ExpensesStop%22%3A0%2C%22ExpenseStopCalculation%22%3A%22%22%2C%22BaseYearEnd%22%3A%221%2F1%2F1900%22%2C%22ExpenseAdjustment%22%3A%221%2F1%2F1900%22%2C%22EstimatedCharges%22%3A0%2C%22RateChange%22%3A0%2C%22NextRateChange%22%3A%221%2F1%2F1900%22%2C%22PermittedUses%22%3A%22%22%2C%22ExclusiveUses%22%3A%22%22%2C%22ExtensionOption%22%3A%22%22%2C%22ExtensionOptionNotice%22%3A%221%2F1%2F1900%22%2C%22ExpansionOption%22%3A%22%22%2C%22ExpansionOptionNotice%22%3A%221%2F1%2F1900%22%2C%22RightOfFirstRefusal%22%3A%22%22%2C%22Payors%22%3Anull%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/rentalagr/1/2" "request" "wk1"  "WebService--UpdateRentalAgreement"

#
# Validate RentalAgreement 2's end date is 1/15/2018
#
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%221%2F1%2F2018%22%2C%22searchDtStop%22%3A%222%2F1%2F2018%22%7D" > request
dojsonPOST "http://localhost:8270/v1/rentalagrs/1" "request" "wk2"  "WebService--GetRentalAgreement"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
