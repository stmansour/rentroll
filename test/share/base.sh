#############################################################################
# RentRoll test infrastructure base
#
# Include this script in your test script to get the base testing capabilities.
#
# Your script can override the following values:
#
# 	BUD="REX"    - name of the business to use for your tests
# 	RRDATERANGE  - default RentRoll date range. You can override this value with a different range.
#############################################################################

#############################################################################
# Set default values
#############################################################################
SECONDS=0
ERRFILE="err.txt"
UNAME=$(uname)
LOGFILE="log"
MYSQLOPTS=""
MYSQL=$(which mysql)
TESTCOUNT=0			## this is an internal counter, your external script should not touch it
SCRIPTPATH=$(pwd -P)

if [ "x${RRPORT}" = "x" ]; then
	RRPORT="8270"
fi

if [ "x${RRBIN}" = "x" ]; then
	RRBIN="../../tmp/rentroll"
else
	echo "RBIN was pre-set to:  \"${RRBIN}\""
fi
TREPORT="${RRBIN}/../../test/testreport.txt"

RENTROLL="${RRBIN}/rentroll -A"
CSVLOAD="${RRBIN}/rrloadcsv"
GOLD="./gold"

SKIPCOMPARE=0
FORCEGOOD=0
TESTCOUNT=0

if [ "x" == "x${RRDATERANGE}" ]; then
	echo "RRDATERANGE not set.  Setting to default = -j 2016-01-01 -k 2016-02-01"
	RRDATERANGE="-j 2016-03-01 -k 2016-04-01"
fi

if [ "x" == "${BUD}x" ]; then
	echo "BUD not set.  Setting to default = REX"
	BUD="REX"
fi

if [ "x" == "x${CSVDATERANGE}" ]; then
	echo "CSVDATERANGE not set.  Setting to default = -g 1/1/16,2/1/16"
	CSVDATERANGE="-G ${BUD} -g 1/1/16,2/1/16"
fi


#############################################################################
#  This code ensures that mysql does not touch production databases.
#  The way identity is kept, default usage of mysql or mysqldump often
#  goes straight to the production databases.
#############################################################################
if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi

pause() {
	echo
	echo
	read -p "Press [Enter] to continue,  Q or X to quit..." x
	x=$(echo "${x}" | tr "[:upper:]" "[:lower:]")
	if [ ${x} == "q" -o ${x} == "x" ]; then
		exit 0
	fi
}

csvload() {
	echo "command is:  ${CSVLOAD} ${1}"
	echo
	${CSVLOAD} ${1}
}

app() {
	echo "command is:  ${RENTROLL} ${RRDATERANGE} -b ${BUD} ${1}"
	${RENTROLL} ${RRDATERANGE} -b ${BUD} ${1}
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
DE)  Deposits
DM)  Deposit Methods
DY)  Depositories 
G)   GSR
I)   Invoice
IR)  Invoice Report
J)   Journal
L)   Ledger
LA)  Ledger Activity
LB)  Ledger Balance
MR)  Market Rate for Rentable
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
SA)  Specialty Assignments
SO)  Sources
ST)  Statements
T)   Rental Agreement Templates
U)   Custom Attribute Assignments


X) Exit

input is case insensitive
EOF

	read -p "Enter choice: " choice
	choice=$(echo "${choice}" | tr "[:upper:]" "[:lower:]")
	case ${choice} in
		  d) app "-b ${BUD} -r 14,2015-05-25" ;;
		 ir) app "-b ${BUD} -r 9,IN00001" ;;
		  j) app "-b ${BUD} -r 1" ;;
		  l) app "-b ${BUD} -r 2" ;;
		 la) app "-b ${BUD} -r 10" ;;
		 lb) app "-b ${BUD} -r 17" ;;
		 mr) app "-b ${BUD} -r 20,R001" ;;
		  a) csvload "-L 11,${BUD}" ;;
		  b) csvload "-L 3" ;;
		  c) csvload "-L 10,${BUD}" ;;
		 ca) csvload "-L 14" ;;
		 de) csvload "${CSVDATERANGE} -L 19,${BUD}" ;;
		 dm) csvload "-L 23,${BUD}" ;;
		 dy) csvload "-L 18,${BUD}" ;;
		  g) app "-b ${BUD} -r 11" ;;
		  i) csvload "-L 20,${BUD}" ;;
		 nt) csvload "-L 17,${BUD}" ;;
		  p) csvload "-L 7,${BUD}" ;;
		 pe) csvload "-L 16,RA0002" ;;
		 pt) csvload "-L 12,${BUD}" ;;
		  q) exit 0 ;;
		  r) csvload "-L 13,${BUD}" ;;
		 ra) csvload "-L 9,${BUD}" ;;
		rab) app "-b ${BUD} -r 12,11,RA001,2016-07-04"; app "-b ${BUD} -r 12,9,RA001,2016-07-04" ;;
		 rc) app "-b ${BUD} -r 7" ;;
		 re) csvload "-L 6,${BUD}" ;;
		 rp) csvload "-L 26,${BUD}" ;;
		rpr) csvload "-L 27,${BUD}" ;;
		 rr) app "-b ${BUD} -r 4" ;;
		 rs) csvload "-L 22,${BUD}" ;;
		 rt) csvload "-L 5,${BUD}" ;;
		  s) csvload "-L 21,${BUD}" ;;
		 sa) csvload "-L 22,${BUD}" ;;
		 so) csvload "-L 24,${BUD}" ;;
		 st) app "-b ${BUD} -r 8" ;;
		  t) csvload "-L 8,${BUD}" ;;
		  u) csvload "-L 15" ;;
		  x)	exit 0 ;;
		  *)	echo "Unknown report: ${choice}"
	esac
	pause
done
}

usage() {
	cat <<EOF

SYNOPSIS
	$0 [-f -o -r]

	Rentroll test script. Compare the output of each step to its associated
	.gold known-good output. If they miscompare, fail and stop the script.
	If they match, keep going until all tasks are completed.

OPTIONS
	-f  Executes all the steps of the test but does not compare the output
	    to the known-good files. This is useful when making a slight change
	    to something just to see how it will work.

	-o  Regenerate the .gold files based on the output from this run. Only
	    use this option if you're sure the output is correct. This option
	    can be a huge time saver, but use it with caution. All .gold files
	    are maintained in the ./${GOLD}/ directory.

	-r  Run the script in interactive REPORT mode. A menu of report options
	    is displayed. Type in the letter(s) for the the report you want and
	    it will run. When the report completes, the script will pause for you
	    to review the output. Pressing Return will go back to the menu of
	    reports.
EOF
}

##########################################################################
# elapsedtime()
# Shows the number of seconds that was needed to run this script
##########################################################################
elapsedtime() {
	duration=$SECONDS
	msg="ElapsedTime: $(($duration / 60)) min $(($duration % 60)) sec"
	echo "${msg}" >>${LOGFILE}
	echo "${msg}"

}

passmsg() {
	printf "PASSED  %-10s  %-40.40s  %6d  \n" "${TESTDIR}" "${TESTNAME}" ${TESTCOUNT} >> ${TREPORT}
}

failmsg() {
	printf "FAILED  %-10s  %-40.40s  %6d  \n" "${TESTDIR}" "${TESTNAME}" ${TESTCOUNT} >> ${TREPORT}
}

forcemsg() {
	printf "FORCED  %-10s  %-40.40s  %6d  \n" "${TESTDIR}" "${TESTNAME}" ${TESTCOUNT} >> ${TREPORT}
}

tdir() {
	local IFS=/
	local p n m
	p=( ${SCRIPTPATH} )
	n=${#p[@]}
	m=$(( n-1 ))
	TESTDIR=${p[$m]}
}


#############################################################################
# docsvtest()  
#    The purpose of this routine is to call rrloadcsv with the 
#     parameters supplied in $2 and send its output to a file
#     named $1. After trrloadcsv completes, the output in $1 will
#     be compared with the output in gold/$1.gold.  If there are
#     no diffs, then the test passes.  If there are diffs, then 
#     it terminates execution of the script after doing 
#     the following:
#
#        (a) Displays the diffs
#        (b) Displays the mv command to use if the newly generated
#            output is now correct and the gold/$1.gold file needs
#            to be updated.  You can just copy the command and paste
#            it into your command line.  Very handy
#        (c) Displays the full command it used to generate the output
#            in $1. This is very handy for reproducing a problem.
#
#     Additionally, there are some Environment Variables that cause
#     it to perform several functions that are very handy:
#
#        SKIPCOMPARE - ${SKIPCOMPARE} defaults to 0. As long as its
#            value is 0 the output in $1 will be compared to
#            gold/$1.gold .  However, there may be times where
#            you want the script to run to completion even if the
#            output miscompares with what is in gold/*  By convention,
#            all of my "functest.sh" scripts use the -f option to
#            set this value.
#
#        FORCEGOOD - ${FORCEGOOD} is set to 0 by default. If it is set
#            set to 1 it means that the output generated and stored in
#            $1 during this run is known to be "correct", even though
#            it may be different than what is in gold/$1.gold.  It will
#            automatically copy $1 to gold/$1.go. This is
#            extremely handy if a change was made to the table output
#            generator, or if any new fields were added to the database
#            and you've validated in some other way that everything is
#            working after such a change.  By convention, all of my
#            "function.sh" scripts use the -o option to set FORCEGOOD
#            to 1.    
#                  
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title for reporting.  No spaces
#############################################################################
docsvtest () {
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %3s  %s... " ${TESTCOUNT} $1 $3

	if [ "x${2}" != "x" ]; then
		${CSVLOAD} $2 >${1} 2>&1
	fi

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${1} ${GOLD}/${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${GOLD}/${1}.gold ]; then
			echo "UNSET CONTENT" > ${GOLD}/${1}.gold
			echo "Created a default ${GOLD}/$1.gold for you. Update this file with known-good output."
		fi
		UDIFFS=$(diff ${1} ${GOLD}/${1}.gold | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${1} ${GOLD}/${1}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${CSVLOAD} ${2}" >> ${ERRFILE}
			echo "Differences in ${1} are as follows:" >> ${ERRFILE}
			diff ${GOLD}/${1}.gold ${1} >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg
			exit 1
		fi
	else
		echo 
	fi
}

#############################################################################
# doOnesiteTest()  
#	just like docsvtest only for Onesite
#                 
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title for reporting.  No spaces
#############################################################################
doOnesiteTest () {
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %3s  %s... " ${TESTCOUNT} $1 $3

	if [ "x${2}" != "x" ]; then
		${RRBIN}/onesiteload $2 >${1} 2>&1
	fi

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${1} ${GOLD}/${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${GOLD}/${1}.gold ]; then
			echo "UNSET CONTENT" > ${GOLD}/${1}.gold
			echo "Created a default ${GOLD}/$1.gold for you. Update this file with known-good output."
		fi
		UDIFFS=$(diff ${1} ${GOLD}/${1}.gold | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${1} ${GOLD}/${1}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${CSVLOAD} ${2}" >> ${ERRFILE}
			echo "Differences in ${1} are as follows:" >> ${ERRFILE}
			diff ${GOLD}/${1}.gold ${1} >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg
			exit 1
		fi
	else
		echo 
	fi
}


########################################
# mysqlverify()
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title
#       $4 = mysql validation query
########################################
mysqlverify () {
# Generate the mysql commands needed to validate...
cat >xxqq <<EOF
use rentroll;
${4}
EOF
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %3s  %s... " ${TESTCOUNT} $1 $3
	${CSVLOAD} $2 >>${LOGFILE} 2>&1
	mysql --no-defaults <xxqq >${1}

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${1} ${GOLD}/${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${GOLD}/${1}.gold ]; then
			echo "UNSET CONTENT" > ${GOLD}/${1}.gold
			echo "Created a default $1.gold for you. Update this file with known-good output."
		fi
		UDIFFS=$(diff ${1} ${GOLD}/${1}.gold | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${1} ${GOLD}/${1}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${CSVLOAD} ${2}" >> ${ERRFILE}
			echo "Differences in ${1} are as follows:" >> ${ERRFILE}
			diff ${GOLD}/${1}.gold ${1} >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg
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
	printf "PHASE %2s  %3s  %s... " ${TESTCOUNT} $1 $3
	${RENTROLL} ${2} >${1} 2>&1

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${1} ${GOLD}/${1}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${GOLD}/${1}.gold ]; then
			echo "UNSET CONTENT" > ${GOLD}/${1}.gold
			echo "Created a default ${GOLD}/$1.gold for you. Update this file with known-good output."
		fi
		UDIFFS=$(diff ${1} ${GOLD}/${1}.gold | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${1} ${GOLD}/${1}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${RENTROLL} ${2}" >> ${ERRFILE}
			echo "Differences in ${1} are as follows:" >> ${ERRFILE}
			diff ${GOLD}/${1}.gold ${1} >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg
			exit 1
		fi
	else
		echo 
	fi
}

##########################################################################
# logcheck()
#   Compares log to log.gold 
#   Date related fields are detected with a regular expression and changed
#   to "current time".  More filters may be needed depending on what goes
#   into the logfile.
#	Parameters:
#		none at this time
##########################################################################
logcheck() {
	echo -n "Test completed: " >> ${LOGFILE}
	date >> ${LOGFILE}
	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${LOGFILE} ${GOLD}/${LOGFILE}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		echo -n "PHASE x: Log file check...  "
		if [ ! -f ${GOLD}/${LOGFILE}.gold -o ! -f ${LOGFILE} ]; then
			echo "Missing file -- Required files for this check: log.gold and log"
			failmsg
			exit 1
		fi
		declare -a out_filters=(
			's/^Date\/Time:.*/current time/'
			's/^Test completed:.*/current time/'
			's/(20[1-4][0-9]\/[0-1][0-9]\/[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
			's/(20[1-4][0-9]\/[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
			's/(20[1-4][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
		)
		cp ${GOLD}/${LOGFILE}.gold ll.g
		cp log llog
		for f in "${out_filters[@]}"
		do
			perl -pe "$f" ll.g > llx1; mv llx1 ll.g
			perl -pe "$f" llog > lly1; mv lly1 llog
		done
		UDIFFS=$(diff llog ll.g | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
			passmsg
			rm -f ll.g llog
		else
			echo "FAILED:  differences are as follows:" >> ${ERRFILE}
			diff ll.g llog >> ${ERRFILE}
			echo >> ${ERRFILE}
			echo "If the new output is correct:  mv ${LOGFILE} ${GOLD}/${LOGFILE}.gold" >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg
			exit 1
		fi
	else
		echo "FINISHED...  but did not check output"
	fi
	elapsedtime
}

#########################################################
# startRentRollServer()
#	Kills any currently running instances of the server
#   then starts it up again.  The port is set to the
#   default port of 8270.  If you set RRPORT prior 
#   to including base.sh to override the port number
#########################################################
startRentRollServer () {
	stopRentRollServer
	${RRBIN}/rentroll -p ${RRPORT} > ${RRBIN}/rrlog 2>&1 &
	sleep 1
}

#########################################################
# stopRentRollServer()
#	Kills any currently running instances of the server
#########################################################
stopRentRollServer () {
	killall rentroll > /dev/null 2>&1
	sleep 1
}

########################################
# dojsonPOST()
#   Simulate a POST command to the server and use
#   the supplied file name as the json data
#	Parameters:
# 		$1 = url
#       $2 = json file
# 		$3 = base file name
#		$4 = title
########################################
dojsonPOST () {
	TESTCOUNT=$((TESTCOUNT + 1))
	printf "PHASE %2s  %3s  %s... " ${TESTCOUNT} $3 $4
	CMD="curl -s -X POST ${1} -H \"Content-Type: application/json\" -d @${2}"
	${CMD} | python -m json.tool >${3} 2>>${LOGFILE}

	if [ "${FORCEGOOD}" = "1" ]; then
		cp ${3} ${GOLD}/${3}.gold
		echo "DONE"
	elif [ "${SKIPCOMPARE}" = "0" ]; then
		if [ ! -f ${GOLD}/${3}.gold ]; then
			echo "UNSET CONTENT" > ${GOLD}/${3}.gold
			echo "Created a default ${GOLD}/$1.gold for you. Update this file with known-good output."
		fi

		#--------------------------------------------------------------------
		# The actual data has timestamp information that changes every run.
		# The timestamp can be filtered out for purposes of testing whether
		# or not the web service could be called and can return the expected
		# data.
		#--------------------------------------------------------------------
		declare -a out_filters=(
			's/(^[ \t]+"LastModTime":).*/$1 TIMESTAMP/'
		)
		cp gold/${3}.gold qqx
		cp ${3} qqy
		for f in "${out_filters[@]}"
		do
			perl -pe "$f" qqx > qqx1; mv qqx1 qqx
			perl -pe "$f" qqy > qqy1; mv qqy1 qqy
		done

		UDIFFS=$(diff qqx qqy | wc -l)
		if [ ${UDIFFS} -eq 0 ]; then
			echo "PASSED"
		else
			echo "FAILED...   if correct:  mv ${3} ${GOLD}/${3}.gold" >> ${ERRFILE}
			echo "Command to reproduce:  ${CMD}" >> ${ERRFILE}
			echo "Differences in ${3} are as follows:" >> ${ERRFILE}
			diff qqx qqy >> ${ERRFILE}
			cat ${ERRFILE}
			failmsg
			exit 1
		fi
	else
		echo 
	fi
	rm -f qqx qqy
}

#--------------------------------------------------------------------------
#  Handle command line options...
#--------------------------------------------------------------------------
tdir
while getopts "forR:" o; do
	echo "o = ${o}"
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
echo    "Test Name:    ${TESTNAME}" > ${LOGFILE}
echo    "Test Purpose: ${TESTSUMMARY}" >> ${LOGFILE}
echo -n "Date/Time:    " >>${LOGFILE}
date >> ${LOGFILE}
echo >>${LOGFILE}

echo -n "Create new database... " >> ${LOGFILE} 2>&1
${RRBIN}/rrnewdb
if [ $? -eq 0 ]; then
	echo " successful" >> ${LOGFILE} 2>&1
else
	echo " ERROR" >> ${LOGFILE} 2>&1
	echo "Failed to create new database" > ${ERRFILE}
	cat ${ERRFILE}
	failmsg
	exit 1
fi

 