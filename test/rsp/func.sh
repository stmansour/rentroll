#!/bin/bash
# This is a quick functional test for testing RentalSpecialtyRefs associated with Rentables.
# It is essentially a copy of test/engine/*  but the init.sql code that initialized rentable specialties
# was replaced by csv files, which invokes the routines in loadrsrefcsv 

ERRFILE="err.txt"
RRBIN="../../tmp/rentroll"
SCRIPTLOG="f.log"
APP="${RRBIN}/rentroll -A -j 2015-11-01 -k 2015-12-01"
MYSQLOPTS=""
UNAME=$(uname)

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
 	MYSQLOPTS="--no-defaults"
fi

echo -n "Test Run " >log 2>&1
date >>log

#---------------------------------------------------------------------
#  Initialize the db, run the app, generate the reports
#---------------------------------------------------------------------
rm -f ${ERRFILE}
${RRBIN}/rrnewdb
mysql ${MYSQLOPTS} <init.sql
if [ $? -eq 0 ]; then
	echo "Init was successful"
else
	echo "INIT HAD ERRORS"
	exit 1
fi

rm -f w x y z

${RRBIN}/rrloadcsv -F rspref.csv >rsplog 2>&1
UDIFFS=$(diff rsplog.gold rsplog | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "RENTAL SPECIALTIES CHECK: PASSED"
else
	echo "RENTAL SPECIALTIES CHECK: FAILED...  if correct:   mv rsplog rsplog.gold"
	echo "Differences are as follows:"
	diff rsplog.gold rsplog
	exit 1
fi

${APP} >>log 2>&1
${APP} -r 1 >j.txt 2>&1
${APP} -r 2 >l.txt 2>&1
${APP} -r 5 >c.txt 2>&1

echo "BEGIN ANALYSIS..."
cp j.gold w
cp j.txt x

UDIFFS=$(diff w x | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 1: PASSED"
else
	echo "PHASE 1: FAILED...  if correct:   mv j.txt j.gold" >> ${ERRFILE}
	echo "Differences are as follows:" >> ${ERRFILE}
	diff w x >> ${ERRFILE}
	cat ${ERRFILE}
	exit 1
fi

cp l.gold y
cp l.txt z

UDIFFS=$(diff y z | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 2: PASSED"
else
	echo "PHASE 2: FAILED...  if correct:   mv l.txt l.gold" >> ${ERRFILE}
	echo "Differences are as follows:" >> ${ERRFILE}
	diff y z >> ${ERRFILE}
	cat ${ERRFILE}
	exit 1
fi

cp c.gold c1
cp c.txt c2

UDIFFS=$(diff c1 c2 | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 3: PASSED"
else
	echo "PHASE 3: FAILED...  if correct:   mv c.txt c.gold" >> ${ERRFILE}
	echo "Differences are as follows:" >> ${ERRFILE}
	diff c1 c2 >> ${ERRFILE}
	cat ${ERRFILE}
	exit 1
fi

echo "RENTROLL ENGINE TESTS PASSED"
exit 0
