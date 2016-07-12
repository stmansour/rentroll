#!/bin/bash
RRBIN="../../tmp/rentroll"
SCRIPTLOG="f.log"
APP="${RRBIN}/rentroll -A"
MYSQLOPTS=""
UNAME=$(uname)

TESTCOUNT=0

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi

echo "CREATE NEW DATABASE"
${RRBIN}/rrnewdb
${RRBIN}/rrloadcsv -b nb.csv -O nt.csv
./notes > log

echo -n "Log file check...  "
if [ ! -f log.gold -o ! -f log ]; then
	echo "Missing file -- Required files for this check: log.gold and log"
	exit 1
fi
declare -a out_filters=(
	's/(20[1-4][0-9]\/[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
	's/(20[1-4][0-9]-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
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
	echo "FAILED...   if correct:   mv log log.gold"
	echo "Differences are as follows:"
	diff ll.g llog
	exit 1
fi

echo "NOTES TEST PASSED"
exit 0
