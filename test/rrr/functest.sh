#!/bin/bash
ERRFILE="err.txt"
MYSQLOPTS=""
UNAME=$(uname)
LOGFILE="log"
SKIPCOMPARE=0
FORCEGOOD=0
TESTCOUNT=0
BUD="REX"

RRBIN="../../tmp/rentroll"
CSVLOAD="${RRBIN}/rrloadcsv"
RENTROLL="${RRBIN}/rentroll -A -j 2016-07-01 -k 2016-08-01"


###   BEGIN ---  MENU DRIVEN REPORTS   
#############################################################################
# pause()
#   Description:
#		Wait for user to press a key before continuing
#############################################################################
pause() {
	read -p "Press [Enter] to continue, X to quit..." x
	if [ "$x" = "x" -o "$x" = "X" -o "$x" = "q" ]; then
		exit 0
	fi
}

csvload() {
	echo "command is:  ${CSVLOAD} ${1}"
	${CSVLOAD} ${1}
}

app() {
	echo "command is:  ${RENTROLL} ${1}"
	${RENTROLL} ${1}
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
G)   GSR
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
RAB) Rental Agreement Account Balance
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
		  g) app "-r 11" ;;
		  i) csvload "-L 20,${BUD}" ;;
		 nt) csvload "-L 17,${BUD}" ;;
		  p) csvload "-L 7,${BUD}" ;;
		 pe) csvload "-L 16,RA0002" ;;
		 pt) csvload "-L 12,${BUD}" ;;
		  q) exit 0 ;;
		  r) csvload "-L 13,${BUD}" ;;
		 ra) csvload "-L 9,${BUD}" ;;
		rab) app "-r 12,11,RA001,2016-07-04"; app "-r 12,9,RA001,2016-07-04" ;;
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
while getopts "forR:" o; do
	case "${o}" in
		r | R)
			doReport
			exit 0
			;;
		f)  SKIPCOMPARE=1
			echo "SKIPPING COMPARES..."
			;;
		o)	FORCEGOOD=1
			echo "OUTPUT OF THIS RUN IS SAVED AS *.GOLD"
			;;
		*) 	usage
			exit 1
			;;
	esac
done
shift $((OPTIND-1))
###  END --- MENU DRIVEN REPORTS

########################################
# docsvtest()
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title
########################################
docsvtest () {
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %s... " ${TESTCOUNT} $3
	${CSVLOAD} $2 >${1} 2>&1

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${1} ${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${1}.gold ]; then
			echo "UNSET CONTENT" > ${1}.gold
			echo "Created a default $1.gold for you. Update this file with known-good output."
		fi
		UDIFFS=$(diff ${1} ${1}.gold | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${1} ${1}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${CSVLOAD} ${2}" >> ${ERRFILE}
			echo "Differences in ${1} are as follows:" >> ${ERRFILE}
			diff ${1}.gold ${1} >> ${ERRFILE}
			cat ${ERRFILE}
			exit 1
		fi
	else
		echo 
	fi
}

########################################
# dorrtest()
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title
########################################
dorrtest () {
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %s... " ${TESTCOUNT} $3
	${RENTROLL} $2 >${1} 2>&1

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${1} ${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${1}.gold ]; then
			echo "UNSET CONTENT" > ${1}.gold
			echo "Created a default $1.gold for you. Update this file with known-good output."
		fi
		UDIFFS=$(diff ${1} ${1}.gold | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${1} ${1}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${RENTROLL} ${2}" >> ${ERRFILE}
			echo "Differences in ${1} are as follows:" >> ${ERRFILE}
			diff ${1}.gold ${1} >> ${ERRFILE}
			cat ${ERRFILE}
			exit 1
		fi
	else
		echo 
	fi
}


if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi

rm -f ${ERRFILE}
echo "CSV IMPORT TEST" > ${LOGFILE}
echo -n "Date/Time: " >>${LOGFILE}
date >> ${LOGFILE}
echo >>${LOGFILE}

echo "CREATE NEW DATABASE" >> ${LOGFILE} 2>&1
${RRBIN}/rrnewdb

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "c" "-c coa.csv -L 10,REX" "ChartOfAccounts"
docsvtest "d" "-R rentabletypes.csv -L 5,REX" "RentableTypes"
docsvtest "e" "-p people.csv  -L 7" "People"
docsvtest "f" "-r rentable.csv -L 6,REX" "Rentables"
docsvtest "g" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "h" "-C ra.csv -L 9,REX" "RentalAgreements"
docsvtest "i" "-P pmt.csv -L 12,REX" "PaymentTypes"
docsvtest "j" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
docsvtest "k" "-A asmt.csv -L 11,REX" "Assessments"
docsvtest "l" "-e rcpt.csv -L 13,REX" "Receipts"

# force the deposits to post, and the balances to be established
${RRBIN}/rentroll -A -j 2014-12-01 -k 2015-01-01

#pause

# process payments and receipts
dorrtest "p" "-r 11" "GSR"
dorrtest "m" "" "Process"

#pause

dorrtest "n" "-r 1" "Journal"
dorrtest "o" "-r 2" "Ledgers"

dorrtest "q" "-r 12,11,RA001,2016-07-04"
dorrtest "q1" "-r 12,9,RA001,2016-07-04"


echo >>${LOGFILE}

if [ "${SKIPCOMPARE}" = "0" ]; then
	echo -n "PHASE x: Log file check...  "
	if [ ! -f log.gold -o ! -f log ]; then
		echo "Missing file -- Required files for this check: log.gold and log"
		exit 1
	fi
	declare -a out_filters=(
		's/^Date\/Time:.*/current time/'
		's/(20[1-4][0-9]\/[0-1][0-9]\/[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
	)
	cp log.gold ll.g
	cp log llog
	for f in "${out_filters[@]}"
	do
		perl -pe "$f" ll.g > x1; mv x1 ll.g
		perl -pe "$f" llog > y1; mv y1 llog
	done
	UDIFFS=$(diff llog ll.g | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
		rm -f ll.g llog
	else
		echo "FAILED:  differences are as follows:" >> ${ERRFILE}
		diff ll.g llog >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
else
	echo "FINISHED...  but did not check output"
fi
