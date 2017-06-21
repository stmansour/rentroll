#!/bin/bash
TESTNAME="JM1 - 12 months of books on the Rexford Properties"
TESTSUMMARY="Setup and run JM1 company and the Rexford Properties for 1 year"

RRDATERANGE="-j 2016-03-01 -k 2016-04-01"
CREATENEWDB=0

source ../share/base.sh

echo "Generate Instances" >>${LOGFILE}

dorrtest "a" "-j 2016-03-01 -k 2016-04-01 -x -b ${BUD}" "Process-2016-MAR"
dorrtest "a" "-j 2016-04-01 -k 2016-05-01 -x -b ${BUD}" "Process-2016-APR"
dorrtest "a" "-j 2016-05-01 -k 2016-06-01 -x -b ${BUD}" "Process-2016-MAY"
dorrtest "a" "-j 2016-06-01 -k 2016-07-01 -x -b ${BUD}" "Process-2016-JUN"
dorrtest "a" "-j 2016-07-01 -k 2016-08-01 -x -b ${BUD}" "Process-2016-JUL"
dorrtest "a" "-j 2016-08-01 -k 2016-09-01 -x -b ${BUD}" "Process-2016-AUG"
dorrtest "a" "-j 2016-09-01 -k 2016-10-01 -x -b ${BUD}" "Process-2016-SEP"
dorrtest "a" "-j 2016-10-01 -k 2016-11-01 -x -b ${BUD}" "Process-2016-OCT"
dorrtest "a" "-j 2016-11-01 -k 2016-12-01 -x -b ${BUD}" "Process-2016-NOV"
dorrtest "a" "-j 2016-12-01 -k 2017-01-01 -x -b ${BUD}" "Process-2016-DEC"
dorrtest "a" "-j 2017-01-01 -k 2017-02-01 -x -b ${BUD}" "Process-2017-JAN"
dorrtest "a" "-j 2017-02-01 -k 2017-03-01 -x -b ${BUD}" "Process-2017-FEB"
dorrtest "a" "-j 2017-03-01 -k 2017-04-01 -x -b ${BUD}" "Process-2017-MAR"
dorrtest "a" "-j 2017-04-01 -k 2017-05-01 -x -b ${BUD}" "Process-2017-APR"
dorrtest "a" "-j 2017-05-01 -k 2017-06-01 -x -b ${BUD}" "Process-2017-MAY"
dorrtest "a" "-j 2017-06-01 -k 2017-07-01 -x -b ${BUD}" "Process-2017-JUN"
dorrtest "a" "-j 2017-07-01 -k 2017-08-01 -x -b ${BUD}" "Process-2017-JUL"
