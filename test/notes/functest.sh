#!/bin/bash
TESTNAME="Statement Report"
TESTSUMMARY="Exercise csv import. Generate books for 3 months. Test Statement report."

source ../share/base.sh

${RRBIN}/rrloadcsv -b nb.csv -O nt.csv
./notes > ${LOGFILE}

logcheck