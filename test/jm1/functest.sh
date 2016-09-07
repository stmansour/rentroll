#!/bin/bash
TESTNAME="JM1"
TESTSUMMARY="Setup and run JM1 company and the Rexford Properties"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "c" "-m depmeth.csv -L 23,${BUD}" "DepositMethods"
docsvtest "d" "-d depository.csv -L 18,${BUD}" "Depositories"
docsvtest "e" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "f" "-l strlists.csv -L 25,${BUD}" "StringLists"
docsvtest "g" "-p people.csv  -L 7" "People"
docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "i" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "j" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
docsvtest "k" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "l" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
#dorrtest "l" "-A -b ${BUD} -L 22,asm20151201.csv" "Assessments-2015-DEC"
docsvtest "n" "-A asm2015Dec.csv -G ${BUD} -g 12/1/15,1/1/16 -L 11,${BUD}" "Assessments-2015-DEC"
docsvtest "o" "-e rcpt2015Dec.csv -G ${BUD} -g 12/1/15,1/1/16 -L 13,${BUD}" "Receipts-2015-DEC"
dorrtest "p" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 11" "GSR"

#  INITIALIZE database with deposit information and verify Accounts
dorrtest "q" "-j 2015-12-01 -k 2016-01-01 -x -b ${BUD}" "Process-2015-DEC"
dorrtest "r" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 1" "Journal"
dorrtest "s" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 2" "Ledgers"
dorrtest "t" "-r 12,1,RA001,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA01"
dorrtest "u" "-r 12,7,RA001,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA01"
dorrtest "v" "-r 12,1,RA002,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA02"
dorrtest "x" "-r 12,7,RA002,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA02"
dorrtest "y" "-r 12,1,RA003,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA-03"
dorrtest "z" "-r 12,7,RA003,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA-03"


# JANUARY 2016
RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
docsvtest "a1" "-A asm2016Jan.csv -G ${BUD} -g 1/1/16,2/1/16 -L 11,${BUD}" "Assessments-2016-JAN"
docsvtest "b1" "-e rcpt2016Jan.csv -G ${BUD} -g 1/1/16,2/1/16 -L 13,${BUD}" "Receipts-2016-JAN"
dorrtest "c1" "${RRDATERANGE} -x -b ${BUD}" "Process-2016-JAN"
dorrtest "d1" "${RRDATERANGE} -b ${BUD} -r 15" "Vacancy-2016-JAN"
dorrtest "e1" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "f1" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "g1" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest "h1" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest "i1" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"
docsvtest "j1" "-i invoice-2016Jan-Read.csv -G ${BUD} -g 1/1/16,2/1/16 -L 20,REX" "Invoice-2016JAN-Read"
docsvtest "k1" "-i invoice-2016Jan-Costea.csv -G ${BUD} -g 1/1/16,2/1/16 -L 20,REX" "invoice-2016Jan-Costea"
docsvtest "l1" "-i invoice-2016Jan-Haroutunian.csv -G ${BUD} -g 1/1/16,2/1/16 -L 20,REX" "invoice-2016Jan-Haroutunian"
dorrtest "m1" "${RRDATERANGE} -b ${BUD} -r 9,IN001" "InvoiceReport-2016Jan-Read"
dorrtest "n1" "${RRDATERANGE} -b ${BUD} -r 9,2" "InvoiceReport-2016Jan-Costea"
dorrtest "o1" "${RRDATERANGE} -b ${BUD} -r 9,3" "InvoiceReport-2016Jan-Haroutunian"
docsvtest "p1" "-y deposit.csv -G ${BUD} -g 1/1/16,2/1/16 -L 19,${BUD}" "Deposits-2016-JAN"


# FEBRUARY 2016
RRDATERANGE="-j 2016-02-01 -k 2016-03-01"
dorrtest "q1" "-j 2016-02-01 -k 2016-02-02 -x -b ${BUD}" "Process-2016-FEB-recurringAssessments"
docsvtest "r1" "-A asm2016Feb.csv -G ${BUD} -g 2/1/16,3/1/16 -L 11,${BUD}" "Assessments-2016-FEB"
docsvtest "s1" "-e rcpt2016Feb.csv -G ${BUD} -g 2/1/16,3/1/16 -L 13,${BUD}" "Receipts-2016-FEB"
dorrtest "t1" "${RRDATERANGE} -b ${BUD}" "Process-2016-FEB"
dorrtest "u1" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "v1" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "w1" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest "x1" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest "y1" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"


logcheck
