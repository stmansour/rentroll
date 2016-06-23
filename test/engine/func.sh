#!/bin/bash
# This is a quick functional test for the rentroll enging
# It uses the values initialized in directory ../ledger1, generates
# Journal and Ledger records, generates the reports, and validates that
# the reports are what we expect
RRBIN="../../tmp/rentroll"
SCRIPTLOG="f.log"
APP="${RRBIN}/rentroll -A -j 2015-11-01 -k 2015-12-01"
MYSQLOPTS=""
UNAME=$(uname)

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
 	MYSQLOPTS="--no-defaults"
fi


#############################################################################
# dotest()
#   Description:
#		This routine runs a test, validates the results, and prints results
#
#	Parameters:
# 		$1 = base file name.  Expects to find $1.txt and $1.gold
#		$2 = app options
# 		$3 = title
#############################################################################
dotest () {
	echo -n "${3}... "
	${APP} $2 >$1.txt 2>&1
	UDIFFS=$(diff $1.gold $1.txt | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else
		echo "FAILED..."
		echo "    if correct:    mv $1.txt $1.gold"
		echo "    to reproduce:  ${APP} $2"
		echo "Differences are as follows:"
		diff $1.gold $1.txt
		exit 1
	fi
}


#--------------------------------------------------------------------------
#  On with the test! Initialize the db, run the app, generate the reports
#--------------------------------------------------------------------------
echo -n "Test Run " >log 2>&1
date >>log

${RRBIN}/rrnewdb
mysql ${MYSQLOPTS} <init.sql
if [ $? -eq 0 ]; then
	echo "Init was successful"
else
	echo "INIT HAD ERRORS"
	exit 1
fi
rm -f w x y z

dotest "gjl" " " "Generate Journal and Ledgers"
dotest "j" "-r 1" "Journal Report"
dotest "l" "-r 2" "Ledger Report"
dotest "c" "-r 5" "Assessment Checker"
dotest "lb" "-r 6" "Ledger Balances"

echo "RENTROLL ENGINE TESTS PASSED"
exit 0