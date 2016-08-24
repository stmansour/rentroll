#!/bin/bash
ERRFILE="err.txt"
RRBIN="../../tmp/rentroll"
MYSQLOPTS=""
UNAME=$(uname)
CSVLOAD="${RRBIN}/rrloadcsv"
RENTROLL="${RRBIN}/rentroll -A"
LOGFILE="log"
BUD="REX"

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi

#############################################################################
# doLogTest()
#   Description:
#		This routine runs a test, validates the results, and prints results
#
#	Parameters:
# 		$1 = base file name.  Expects to find $1.txt and $1.gold
#		$2 = app options
# 		$3 = title
#############################################################################
doLogTest () {
	echo "${3}"
	echo "${3}" >> ${1} 2>&1
	${CSVLOAD} ${2} >> ${1} 2>&1
	echo >>${1}
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
		echo "file $1.gold not found. Creating an empty $1.gold file for you"
		touch $1.gold
	fi
	UDIFFS=$(diff $1.gold $1.txt | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else >> ${ERRFILE}
		echo "FAILED..." >> ${ERRFILE}
		echo "    if correct:    mv $1.txt $1.gold" >> ${ERRFILE}
		echo "    to reproduce:  ${CSVLOAD} $2" >> ${ERRFILE}
		echo "Differences are as follows:" >> ${ERRFILE}
		diff $1.gold $1.txt >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
}
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
	echo "${3}" > ${1}.txt 2>&1
	${RENTROLL} $2 >> ${1}.txt 2>&1

	if [ ! -f $1.gold ]; then
		echo "file $1.gold not found. Creating an empty $1.gold file for you"
		touch $1.gold
	fi
	UDIFFS=$(diff $1.gold $1.txt | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else
		echo "FAILED..." >> ${ERRFILE}
		echo "    if correct:    mv $1.txt $1.gold" >> ${ERRFILE}
		echo "    to reproduce:  ${RENTROLL} $2" >> ${ERRFILE}
		echo "Differences are as follows:" >> ${ERRFILE}
		diff $1.gold $1.txt >> ${ERRFILE}
		cat ${ERRFILE}
		exit 1
	fi
}

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
	echo "command is:  ${RENTROLL} -j 2016-03-01 -k 2016-04-01 ${1}"
	${RENTROLL} -j "2016-03-01" -k "2016-04-01" ${1}
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
ST)  Statements
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
		 st) app "-r 8" ;;
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
func.sh - test script and report utility
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

#--------------------------------------------------------------------------
#  On with the test! Initialize the db, run the app, generate the reports
#--------------------------------------------------------------------------
rm -f ${ERRFILE}
echo "CSV IMPORT TEST" > ${LOGFILE}
echo -n "Date/Time: " >> ${LOGFILE}
date >> ${LOGFILE}
echo >> ${LOGFILE}

echo "CREATE NEW DATABASE" >> ${LOGFILE} 2>&1
${RRBIN}/rrnewdb

doLogTest ${LOGFILE} "-b business.csv -L 3" "DEFINE BUSINESS"
doLogTest ${LOGFILE} "-d depository.csv -L 18,REX" "DEFINE DEPOSITORIES"
doLogTest ${LOGFILE} "-m dm.csv -L 23,REX" "DEFINE DEPOSIT METHODS"
doLogTest ${LOGFILE} "-S sources.csv -L 24,REX" "DEFINE SOURCES"
doLogTest ${LOGFILE} "-R rentabletypes.csv -L 5,REX" "DEFINE RENTABLE TYPES"
doLogTest ${LOGFILE} "-p people.csv  -L 7" "DEFINE PEOPLE"
doLogTest ${LOGFILE} "-r rentable.csv -L 6,REX" "DEFINE RENTABLES"
doLogTest ${LOGFILE} "-T ratemplates.csv  -L 8" "DEFINE RENTAL AGREEMENT TEMPLATES"
doLogTest ${LOGFILE} "-C ra.csv -L 9,REX" "DEFINE RENTAL AGREEMENTS"
doLogTest ${LOGFILE} "-E pets.csv -L 16,RA0001" "DEFINE PETS"
doLogTest ${LOGFILE} "-c coa.csv -L 10,REX" "DEFINE CHART OF ACCOUNTS"
doLogTest ${LOGFILE} "-A asmt.csv -L 11,REX" "DEFINE ASSESSMENTS"
doLogTest ${LOGFILE} "-P pmt.csv -L 12,REX" "DEFINE PAYMENT TYPES"
doLogTest ${LOGFILE} "-e rcpt.csv -L 13,REX" "DEFINE RECEIPTS"
doLogTest ${LOGFILE} "-u custom.csv -L 14" "DEFINE CUSTOM ATTRIBUTES"
doLogTest ${LOGFILE} "-U assigncustom.csv -L 15" "DEFINE ASSIGN CUSTOM ATTRIBUTES"

echo "process Statements, Assessments, and Payments"
echo "PROCESS STATEMENTS, ASSESSMENTS, AND PAYMENTS" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 1 >> ${LOGFILE} 2>&1		# Journals
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 2  >> ${LOGFILE} 2>&1	# Ledgers
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 8  >> ${LOGFILE} 2>&1	# Statements

doCSVtest "i" "-i invoice.csv -L 20,REX" "CREATE INVOICE"
dotest "k" "-j 2016-03-01 -k 2016-04-01 -r 9,IN0001" "INVOICE REPORT"

echo >>${LOGFILE}

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
	echo "FAILED...  if correct:   mv log log.gold" >> ${ERRFILE}
	echo "Differences are as follows:" >> ${ERRFILE}
	diff ll.g llog >> ${ERRFILE}
	cat ${ERRFILE}
	exit 1
fi
