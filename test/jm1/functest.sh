#!/bin/bash
TESTNAME="JM1"
TESTSUMMARY="Setup and run JM1 company and the Rexford Properties"

RRDATERANGE="-j 2015-12-01 -k 2016-01-01"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "c" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "d" "-l strlists.csv -L 25,${BUD}" "StringLists"
docsvtest "e" "-p people.csv  -L 7" "People"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "g" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "h" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
docsvtest "i" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "j" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "k" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"

docsvtest "l" "-A asm20151201.csv -L 11,${BUD}" "Assessments-2015-DEC"
docsvtest "m" "-e rcpt20151201.csv -L 13,${BUD}" "Receipts-2015-DEC"

#  INITIALIZE database with deposit information and verify Accounts
dorrtest "p" "-j 2015-12-01 -k 2016-01-01 -x -b ${BUD}" "Process-2015-DEC"
dorrtest "q" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 1" "Journal"
dorrtest "r" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 2" "Ledgers"
dorrtest "s" "-r 12,1,RA001,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA01"
dorrtest "t" "-r 12,7,RA001,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA01"
dorrtest "u" "-r 12,1,RA002,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA02"
dorrtest "v" "-r 12,7,RA002,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA02"
dorrtest "x" "-r 12,1,RA003,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA-03"
dorrtest "y" "-r 12,7,RA003,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA-03"


# Start the recurring rent assessments on Jan 1
dorrtest "a1" "-j 2016-01-01 -k 2016-02-01 -b ${BUD} -L 22,asm20160101.csv" "Assessments-2016-JAN"

exit 1
dorrtest "c1" "-j 2016-01-01 -k 2016-02-01 -x -b ${BUD}" "Process-2016-JAN"
docsvtest "b1" "-e rcpt20160101.csv -L 13,${BUD}" "Receipts-2016-JAN"
dorrtest "d1" "-j 2016-01-01 -k 2016-02-01 -b ${BUD} -r 1" "Journal"
dorrtest "e1" "-j 2016-01-01 -k 2016-02-01 -b ${BUD} -r 2" "Ledgers"


#dorrtest "n" "-j 2014-12-01 -k 2015-01-01 -b ${BUD}" "ProcessDeposits"
#dorrtest "o" "${RRDATERANGE} -b ${BUD} -r 11" "GSR"
#dorrtest "u" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

logcheck
