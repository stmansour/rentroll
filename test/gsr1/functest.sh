#!/bin/bash
ERRFILE="err.txt"
LOGFILE="log"
MYSQLOPTS=""
UNAME=$(uname)

SKIPCOMPARE=0
FORCEGOOD=0
TESTCOUNT=0

RRBIN="../../tmp/rentroll"
CSVLOAD="${RRBIN}/rrloadcsv"
RENTROLL="${RRBIN}/rentroll -A"

BUD="DHR"

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi


#############################################################################
# dotest()
#   Description:
#		This routine runs a test, validates the results, and prints results
#
#	Parameters:
# 		$1 = title
#		$2 = app options
#############################################################################
dotest () {
	echo "${1}"
	echo "${1}" >> ${LOGFILE} 2>&1
	${CSVLOAD} $2 >> ${LOGFILE} 2>&1
	echo >>${LOGFILE}
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
D)   Delinquency
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
		  d) app "-r 14" ;;
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
		 rp) csvload "-L 26,DHR" ;;
		rpr) csvload "-L 27,DHR" ;;
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

rm -f ${ERRFILE}
echo "CSV IMPORT TEST" > ${LOGFILE}
echo -n "Date/Time: " >>${LOGFILE}
date >> ${LOGFILE}
echo >>${LOGFILE}

echo "CREATE NEW DATABASE" >> ${LOGFILE} 2>&1
${RRBIN}/rrnewdb

# dotest "DEFINE BUSINESS" "-b business.csv -L 3"
# # dotest "DEFINE ASSESSMENT TYPES" "-a asmtypes.csv -L 4"
# dotest "DEFINE RENTABLE TYPES" "-R rentabletypes.csv -L 5,${BUD}"
# dotest "DEFINE PEOPLE" "-p people.csv  -L 7"
# dotest "DEFINE RENTABLES" "-r rentable.csv -L 6,${BUD}"
# dotest "DEFINE RENTAL AGREEMENT TEMPLATES" "-T ratemplates.csv  -L 8"
# dotest "DEFINE RENTAL AGREEMENTS" "-C ra.csv -L 9,${BUD}"
# dotest "DEFINE PETS" "-E pets.csv -L 16,RA0001"
# dotest "DEFINE CHART OF ACCOUNTS" "-c coa.csv -L 10,${BUD}"
# dotest "DEFINE SPECIALTIES" "-s specialties.csv"
# dotest "DEFINE RENTABLE SPECIALTY REFERENCES" "-F rspref.csv"
# dotest "DEFINE ASSESSMENTS" "-A asmt.csv -L 11,${BUD}"
# dotest "DEFINE PAYMENT TYPES" "-P pmt.csv -L 12,${BUD}"
# dotest "DEFINE RECEIPTS" "-e rcpt.csv -L 13,${BUD}"
# dotest "DEFINE CUSTOM ATTRIBUTES" "-u custom.csv -L 14"
# dotest "DEFINE ASSIGN CUSTOM ATTRIBUTES" "-U assigncustom.csv -L 15"

# echo "process payments and receipts"
# echo "PROCESS PAYMENTS AND RECEIPTS" >> ${LOGFILE} 2>&1

# ${RENTROLL} -j "2016-03-01" -k "2016-04-01" >> ${LOGFILE} 2>&1
# ${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 1 >> ${LOGFILE} 2>&1
# ${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 2  >> ${LOGFILE} 2>&1
# ${RENTROLL} -j "2016-04-01" -k "2016-05-01" >> ${LOGFILE} 2>&1
# ${RENTROLL} -j "2016-05-01" -k "2016-06-01" >> ${LOGFILE} 2>&1
# ${RENTROLL} -j "2016-05-01" -k "2016-06-01" -r 1 >> ${LOGFILE} 2>&1
# ${RENTROLL} -j "2016-05-01" -k "2016-06-01" -r 2  >> ${LOGFILE} 2>&1

# echo >>${LOGFILE}

# echo -n "PHASE x: Log file check...  "
# if [ ! -f log.gold -o ! -f log ]; then
# 	echo "Missing file -- Required files for this check: log.gold and log"
# 	exit 1
# fi
# declare -a out_filters=(
# 	's/^Date\/Time:.*/current time/'
# 	's/(20[1-4][0-9]\/[0-1][0-9]\/[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
# )
# cp log.gold ll.g
# cp log llog
# for f in "${out_filters[@]}"
# do
# 	perl -pe "$f" ll.g > x1; mv x1 ll.g
# 	perl -pe "$f" llog > y1; mv y1 llog
# done
# UDIFFS=$(diff llog ll.g | wc -l)
# if [ ${UDIFFS} -eq 0 ]; then
# 	echo "PASSED"
# 	rm -f ll.g llog
# else
# 	echo "FAILED"
# 	echo "FAILED" >> ${ERRFILE}
# 	echo "    if correct:   mv log log.gold" >> ${ERRFILE}
# 	echo "Differences are as follows:" >> ${ERRFILE}
# 	diff ll.g llog >> ${ERRFILE}
# 	cat ${ERRFILE}
# 	exit 1
# fi

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "d" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "c" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "e" "-p people.csv  -L 7" "People"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "g" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "h" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "i" "-E pets.csv -L 16,RA0001" "Pets"
docsvtest "j" "-s specialties.csv" "Specialties"
docsvtest "k" "-F rspref.csv" "SpecialtyRefs"
docsvtest "l" "-A asmt.csv -L 11,${BUD}" "Assessments"
docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "n" "-e rcpt.csv -L 13,${BUD}" "Receipts"
docsvtest "o" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "p" "-U assigncustom.csv -L 15" "AssignCustomAttributes"

dorrtest "q" "-j 2016-03-01 -k 2016-04-01" "Process"
dorrtest "r" "-j 2016-03-01 -k 2016-04-01 -r 1" "March-Journal"
dorrtest "s" "-j 2016-03-01 -k 2016-04-01 -r 2" "March-Ledger"
dorrtest "t" "-j 2016-04-01 -k 2016-05-01" "April-Process"
dorrtest "u" "-j 2016-05-01 -k 2016-06-01" "May-Process"
dorrtest "v" "-j 2016-05-01 -k 2016-06-01 -r 1" "May-Journal"
dorrtest "w" "-j 2016-05-01 -k 2016-06-01 -r 2" "May-Ledger"
dorrtest "x" "-j 2016-05-01 -k 2016-06-01 -r 8" "May-Statement"


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
