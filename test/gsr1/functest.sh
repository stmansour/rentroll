#!/bin/bash
RRBIN="../../tmp/rentroll"
CVSLOAD="${RRBIN}/rrloadcsv"
UNAME=$(uname)
RENTROLL="${RRBIN}/rentroll -A"
LOGFILE="log"
MYSQLOPTS=""

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
	${CVSLOAD} $2 >> ${LOGFILE} 2>&1
	echo >>${LOGFILE}
}

echo "CSV IMPORT TEST" > ${LOGFILE}
echo -n "Date/Time: " >>${LOGFILE}
date >> ${LOGFILE}
echo >>${LOGFILE}

echo "CREATE NEW DATABASE" >> ${LOGFILE} 2>&1
${RRBIN}/rrnewdb

dotest "DEFINE BUSINESS" "-b business.csv -L 3"
# dotest "DEFINE ASSESSMENT TYPES" "-a asmtypes.csv -L 4"
dotest "DEFINE RENTABLE TYPES" "-R rentabletypes.csv -L 5,DHR"
dotest "DEFINE PEOPLE" "-p people.csv  -L 7"
dotest "DEFINE RENTABLES" "-r rentable.csv -L 6,DHR"
dotest "DEFINE RENTAL AGREEMENT TEMPLATES" "-T ratemplates.csv  -L 8"
dotest "DEFINE RENTAL AGREEMENTS" "-C ra.csv -L 9,DHR"
dotest "DEFINE PETS" "-E pets.csv -L 16,RA0001"
dotest "DEFINE CHART OF ACCOUNTS" "-c coa.csv -L 10,DHR"
dotest "DEFINE SPECIALTIES" "-s specialties.csv"
dotest "DEFINE RENTABLE SPECIALTY REFERENCES" "-F rspref.csv"
dotest "DEFINE ASSESSMENTS" "-A asmt.csv -L 11,DHR"
dotest "DEFINE PAYMENT TYPES" "-P pmt.csv -L 12,DHR"
dotest "DEFINE RECEIPTS" "-e rcpt.csv -L 13,DHR"
dotest "DEFINE CUSTOM ATTRIBUTES" "-u custom.csv -L 14"
dotest "DEFINE ASSIGN CUSTOM ATTRIBUTES" "-U assigncustom.csv -L 15"

echo "process payments and receipts"
echo "PROCESS PAYMENTS AND RECEIPTS" >> ${LOGFILE} 2>&1

${RENTROLL} -j "2016-03-01" -k "2016-04-01" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 1 >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 2  >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-04-01" -k "2016-05-01" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-05-01" -k "2016-06-01" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-05-01" -k "2016-06-01" -r 1 >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-05-01" -k "2016-06-01" -r 2  >> ${LOGFILE} 2>&1

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
	echo "FAILED...  if correct:   mv log log.gold"
	echo "Differences are as follows:"
	diff ll.g llog
	exit 1
fi
