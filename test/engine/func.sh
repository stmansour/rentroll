#!/bin/bash
# This is a quick functional test for the rentroll enging
# It uses the values initialized in directory ../ledger1, generates
# journal and ledger records, generates the reports, and validates that
# the reports are what we expect
RRBIN="../../tmp/rentroll"
SCRIPTLOG="f.log"
APP="${RRBIN}/rentroll -A"
MYSQLOPTS=""
UNAME=$(uname)

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
 	MYSQLOPTS="--no-defaults"
fi


#---------------------------------------------------------------------
#  Initialize the db, run the app, generate the reports
#---------------------------------------------------------------------
# pushd ../ledger1 ; make ; popd
${RRBIN}/rrnewdb
mysql ${MYSQLOPTS} <init.sql

rm -f w x y z

${APP}
${APP} -r 1 >j.txt
${APP} -r 2 >l.txt

echo "BEGIN ANALYSIS..."
cp j.gold w
cp j.txt x

UDIFFS=$(diff w x | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 1: PASSED"
else
	echo "PHASE 1: FAILED:  differences are as follows:"
	diff w x
	exit 1
fi

cp l.gold y
cp l.txt z

UDIFFS=$(diff y z | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 2: PASSED"
else
	echo "PHASE 2: FAILED:  differences are as follows:"
	diff y z
	exit 1
fi

echo "RENTROLL ENGINE TESTS PASSED"
exit 0