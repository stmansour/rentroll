#!/bin/bash

TESTNAME="Multi-RAID Payments"
TESTSUMMARY="Test Single Payment - multi-RAID"

RRDATERANGE="-j 2017-02-01 -k 2017-03-01"
CSVDATERANGE="-g 2/1/17,3/1/17"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "c" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "d" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "e" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "f" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
docsvtest "g" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "h" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "i" "-A asmt.csv -G ${BUD} ${CSVDATERANGE} -L 11,${BUD}" "Assessments"

docsvtest "j" "-e rcpt.csv -G ${BUD} ${CSVDATERANGE} -L 13,${BUD}" "Receipts"

# FEB 2017  -  MAR 2017
dorrtest "k" "${RRDATERANGE} -b ${BUD}" "ProcessRecurringAssessments"

dorrtest "l" "${RRDATERANGE} -b ${BUD} -r 11" "GSR"
dorrtest "m" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "o" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"

dorrtest "r" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

logcheck
