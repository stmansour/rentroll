#!/bin/bash
RRBIN="../../tmp/rentroll"
MYSQLOPTS=""
UNAME=$(uname)
CVSLOAD="${RRBIN}/rrloadcsv"
RENTROLL="${RRBIN}/rentroll -A"
LOGFILE="log"

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi

echo "CSV IMPORT TEST" > ${LOGFILE}
echo -n "Date/Time: " >>${LOGFILE}
date >> ${LOGFILE}
echo >>${LOGFILE}

echo "CREATE NEW DATABASE" >> ${LOGFILE} 2>&1
${RRBIN}/rrnewdb

echo "import Business"
echo "DEFINE BUSINESS" >> ${LOGFILE} 2>&1
${CVSLOAD} -b business.csv -L 3 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import assessment types"
echo "DEFINE ASSESSMENT TYPES" >> ${LOGFILE} 2>&1
${CVSLOAD} -a asmtypes.csv -L 4 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import Rentable types"
echo "DEFINE RENTABLE TYPES" >> ${LOGFILE} 2>&1
${CVSLOAD} -R rentabletypes.csv -L 5,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import people"
echo "DEFINE PEOPLE" >> ${LOGFILE} 2>&1
${CVSLOAD} -p people.csv  -L 7 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import rentables"
echo "DEFINE RENTABLES" >> ${LOGFILE} 2>&1
${CVSLOAD} -r rentable.csv -L 6,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import rental agreement templates"
echo "DEFINE RENTAL AGREEMENT TEMPLATES" >> ${LOGFILE} 2>&1
${CVSLOAD} -T ratemplates.csv  -L 8 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import rental agreements"
echo "DEFINE RENTAL AGREEMENTS" >> ${LOGFILE} 2>&1
${CVSLOAD} -C ra.csv -L 9,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import pets"
echo "DEFINE PETS" >> ${LOGFILE} 2>&1
# CMD="${CVSLOAD} -E pets.csv -L 16,RA0001"
# echo ${CMD}
${CVSLOAD} -E pets.csv -L 16,RA0001 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import chart of accounts"
echo "DEFINE CHART OF ACCOUNTS" >> ${LOGFILE} 2>&1
${CVSLOAD} -c coa.csv -L 10,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}


echo "import rental specialties"
echo "DEFINE SPECIALTIES" >> ${LOGFILE} 2>&1
${CVSLOAD} -s specialties.csv >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import rentable specialty references"
echo "DEFINE RENTABLE SPECIALTY REFERENCES" >> ${LOGFILE} 2>&1
${CVSLOAD} -F rspref.csv >> ${LOGFILE} 2>&1
echo >>${LOGFILE}


echo "import Assessments"
echo "DEFINE ASSESSMENTS" >> ${LOGFILE} 2>&1
${CVSLOAD} -A asmt.csv -L 11,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import payment types"
echo "DEFINE PAYMENT TYPES" >> ${LOGFILE} 2>&1
${CVSLOAD} -P pmt.csv -L 12,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import receipts"
echo "DEFINE RECEIPTS" >> ${LOGFILE} 2>&1
${CVSLOAD} -e rcpt.csv -L 13,DHR >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import custom attributes"
echo "DEFINE CUSTOM ATTRIBUTES" >> ${LOGFILE} 2>&1
${CVSLOAD} -u custom.csv -L 14 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "import assign custom attributes"
echo "DEFINE ASSIGN CUSTOM ATTRIBUTES" >> ${LOGFILE} 2>&1
${CVSLOAD} -U assigncustom.csv -L 15 >> ${LOGFILE} 2>&1
echo >>${LOGFILE}

echo "process payments and receipts"
echo "PROCESS PAYMENTS AND RECEIPTS" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 1 >> ${LOGFILE} 2>&1
${RENTROLL} -j "2016-03-01" -k "2016-04-01" -r 2  >> ${LOGFILE} 2>&1

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
