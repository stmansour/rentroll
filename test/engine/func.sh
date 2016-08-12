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
ERRFILE="err.txt"
CSVLOAD="${RRBIN}/rrloadcsv"
BUD="REX"

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
	if [ ! -f $1.gold ]; then
		echo "file $1.gold not found. Please create $1.gold then rerun test." >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
	UDIFFS=$(diff $1.gold $1.txt | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else
		echo "FAILED..."
		echo "${3} FAILED"  > ${ERRFILE}
		echo "    if correct:    mv $1.txt $1.gold" >> ${ERRFILE}
		echo "    to reproduce:  ${APP} $2" >> ${ERRFILE}
		echo "Differences are as follows:" >> ${ERRFILE}
		diff $1.gold $1.txt >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
}

#############################################################################
# doCSVtest()
#   Description:
#		This routine runs a test, validates the results, and prints results
#
#	Parameters:
# 		$1 = base file name.  Expects to find $1.txt and $1.gold
#		$2 = app options
# 		$3 = title
#############################################################################
doCSVtest () {
	echo -n "${3}... "
	echo "${3}" > ${1}.txt 2>&1
	${CSVLOAD} $2 >> ${1}.txt 2>&1

	if [ ! -f $1.gold ]; then
		echo "file $1.gold not found. Please create $1.gold then rerun test." >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
	UDIFFS=$(diff $1.gold $1.txt | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else
		echo "FAILED..."
		echo "${3} FAILED" > ${ERRFILE}
	cat ${ERRFILE}
		echo "    if correct:    mv $1.txt $1.gold" >> ${ERRFILE}
		echo "    to reproduce:  ${CSVLOAD} $2" >> ${ERRFILE}
		echo "Differences are as follows:" >> ${ERRFILE}
		diff $1.gold $1.txt >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
}


###   BEGIN ---  MENU DRIVEN REPORTS   
#############################################################################
# pause()
#   Description:
#		Wait for user to press a key before continuing
#############################################################################
pause() {
	read -p "Press [Enter] to continue,  Q or X to quit..." x
	x=$(echo "${x}" | tr "[:upper:]" "[:lower:]")
	if [ ${x} == "q" -o ${x} == "x" ]; then
		exit 0
	fi
}

csvload() {
	echo "command is:  ${CSVLOAD} ${1}"
	${CSVLOAD} ${1}
}

app() {
	echo "command is:  ${APP} ${1}"
	${APP} ${1}
}
#############################################################################
# doReport()
#   Description:
#		Run database reports based on user selection
#############################################################################
doReport () {
while :
do
	clear
	cat <<EOF
-----------------------------------------
   R E N T R O L L  --  R E P O R T S
-----------------------------------------
A)   Assessments
B)   Business
C)   Chart of Accounts
CA)  Custom Attributes
DY)  Depositories 
I)   Invoice
IR)  Invoice Report
J)   Journal
L)   Ledger
LA)  Ledger Activity
LB)  Ledger Balance
NT)  Note Types
P)   People
PE)  Pets
PT)  Payment Types
R)   Receipts
RA)  Rental Agreements
RC)  Rentable Count by Rentable Type
RE)  Rentables
RP)  RatePlans
RR)  RentRoll
RPR) RatePlanRef
RS)  Rentable Specialty Assignments
RT)  Rentable Types
S)   Rentable Specialties
T)   Rental Agreement Templates
U)   Custom Attribute Assignments


X) Exit

input is case insensitive
EOF

	read -p "Enter choice: " choice
	choice=$(echo "${choice}" | tr "[:upper:]" "[:lower:]")
	case ${choice} in
		 ir) app "-r 9,IN00001" ;;
		  j) app "-r 1" ;;
		  l) app "-r 2" ;;
		 la) app "-r 10" ;;
		 lb) app "-r 6" ;;
		  a) csvload "-L 11,${BUD}" ;;
		  b) csvload "-L 3" ;;
		  c) csvload "-L 10,${BUD}" ;;
		 ca) csvload "-L 14" ;;
		 dy) csvload "-L 18,${BUD}" ;;
		  i) csvload "-L 20,${BUD}" ;;
		 nt) csvload "-L 17,${BUD}" ;;
		  p) csvload "-L 7,${BUD}" ;;
		 pe) csvload "-L 16,RA0002" ;;
		 pt) csvload "-L 12,${BUD}" ;;
		  r) csvload "-L 13,${BUD}" ;;
		 ra) csvload "-L 9,${BUD}" ;;
		 rc) app "-r 7" ;;
		 re) csvload "-L 6,${BUD}" ;;
		 rp) csvload "-L 26,REX" ;;
		rpr) csvload "-L 27,REX" ;;
		 rr) app "-r 4" ;;
		 rs) csvload "-L 22,${BUD}" ;;
		 rt) csvload "-L 5,${BUD}" ;;
		  s) csvload "-L 21,${BUD}" ;;
		  t) csvload "-L 8" ;;
		  u) csvload "-L 15" ;;
		  x)	exit 0 ;;
		  *)	echo "Unknown report: ${choice}"
	esac
	pause
done
}


usage() {
	cat <<EOF
functest.sh - test script and report utility
	run this command with no options to perform the test
	run this command with -r or -R to bring up the report interface
EOF
}

#--------------------------------------------------------------------------
#  Look at the command line options first
#--------------------------------------------------------------------------
while getopts "rR:" o; do
	case "${o}" in
		r | R)
			doReport
			exit 0
			;;
		*) 	usage
			exit 1
			;;
	esac
done
shift $((OPTIND-1))
###  END --- MENU DRIVEN REPORTS


#--------------------------------------------------------------------------
#  On with the test! Initialize the db, run the app, generate the reports
#--------------------------------------------------------------------------
rm -f ${ERRFILE}
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

doCSVtest "ca" "-u custom.csv -L 14" "Custom Attributes" 
doCSVtest "ac" "-U assigncustom.csv -L 15" "Assign Custom Attributes"
doCSVtest "dp" "-d depository.csv -y deposit.csv -L 19,REX" "Deposits"

dotest "k" "-r 7" "Count of Rentables by Type"
# dotest "s" "-r 8" "Statements"
doCSVtest "i1" "-i invoice.csv -L 20,REX" "CREATE INVOICE"
dotest "i2" "-r 9,IN00001" "INVOICE REPORT"

echo "RENTROLL ENGINE TESTS PASSED"
exit 0