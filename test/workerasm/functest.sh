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
res=$(grep "Cannot add new assessment instance on Feb 1, 2018 after RentalAgreement (RA-2) stop date Feb 1, 2018" ${WLOG} | wc -l)
if (( res != 1 )); then
	echo "ERROR: assessment instance for 2/1/2018 was not suppressed."
	exit 2
fi
dorrtest "wk0" "${RRDATERANGE} -b ${BUD} -r 24" "Assessments-Feb-2018"

logcheck

exit 0
