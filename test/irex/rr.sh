#!/bin/bash
TESTNAME="IREX - UI TestDB for REX"
TESTSUMMARY="minimal db for UI testing on REX"

RRDATERANGE="-j 2017-01-01 -k 2017-02-01"

source ../share/base.sh

echo "BEGIN IREX FUNCTIONAL TEST" >>${LOGFILE}
#========================================================================================
# INITIALIZE THE BUSINESS
#   This section has the 1-time tasks to set up the business and get the accounts to
#   their correct starting values.
#========================================================================================
docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "Business"
docsvtest "e" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"

logcheck

