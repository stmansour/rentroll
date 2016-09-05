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
#dorrtest "l" "-A -b ${BUD} -L 22,asm20151201.csv" "Assessments-2015-DEC"
docsvtest "l" "-A asm20151201.csv -G ${BUD} -g 12/1/15,1/1/16 -L 11,${BUD}" "Assessments-2015-DEC"
docsvtest "m" "-e rcpt20151201.csv -G ${BUD} -g 12/1/15,1/1/16 -L 13,${BUD}" "Receipts-2015-DEC"

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


# JANUARY 2016
RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
docsvtest "a1" "-A asm20160101.csv -G ${BUD} -g 1/1/16,2/1/16 -L 11,${BUD}" "Assessments-2016-JAN"
docsvtest "b1" "-e rcpt20160101.csv -G ${BUD} -g 1/1/16,2/1/16 -L 13,${BUD}" "Receipts-2016-JAN"
dorrtest "c1" "${RRDATERANGE} -x -b ${BUD}" "Process-2016-JAN"
dorrtest "d1" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "e1" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "f1" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest "g1" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest "h1" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

# FEBRUARY 2016
RRDATERANGE="-j 2016-02-01 -k 2016-03-01"
dorrtest "i1" "-j 2016-02-01 -k 2016-02-02 -x -b ${BUD}" "Process-2016-FEB-recurringAssessments"
docsvtest "j1" "-A asm20160201.csv -G ${BUD} -g 2/1/16,3/1/16 -L 11,${BUD}" "Assessments-2016-FEB"
docsvtest "k1" "-e rcpt20160201.csv -G ${BUD} -g 2/1/16,3/1/16 -L 13,${BUD}" "Receipts-2016-FEB"
dorrtest "l1" "${RRDATERANGE} -b ${BUD}" "Process-2016-FEB"
dorrtest "m1" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "n1" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "o1" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest "p1" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest "q1" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"


#dorrtest "n" "-j 2014-12-01 -k 2015-01-01 -b ${BUD}" "ProcessDeposits"
#dorrtest "o" "${RRDATERANGE} -b ${BUD} -r 11" "GSR"

logcheck
