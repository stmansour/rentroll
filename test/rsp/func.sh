#!/bin/bash
TESTNAME="RentalSpecialtyRefs, Assessment Check"
TESTSUMMARY="Tests RentalSpecialtyRefs associated with Rentables. Based on test/engine, init.sql specialty code was replaced with csvloader."
BUD="SRC"
RRDATERANGE="-j 2015-11-01 -k 2015-12-01"

source ../share/base.sh

mysql ${MYSQLOPTS} <init.sql
if [ $? -eq 0 ]; then
	echo "Init was successful"
else
	echo "INIT HAD ERRORS"
	exit 1
fi

docsvtest "a" "-F rspref.csv -L 21,${BUD}" "Specialties"

dorrtest "b" "${RRDATERANGE} -b ${BUD} " "Process"
dorrtest "c" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "d" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "e" "${RRDATERANGE} -b ${BUD} -r 5" "AssessmentCheck"

logcheck
