#!/bin/bash

TESTNAME="RentRoll Report / Web Services"
TESTSUMMARY="Test the RentRoll Report and Web Services"

RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "c" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "d" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "e" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "ea" "-V vehicle.csv" "Vehicles"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "g" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
docsvtest "h" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "i" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "j" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
docsvtest "k" "-A asmt.csv -G ${BUD} -g 7/1/16,8/1/16 -L 11,${BUD}" "Assessments"
docsvtest "l" "-e rcpt.csv -G ${BUD} -g 7/1/16,8/1/16 -L 13,${BUD}" "Receipts"
# DEC 2014  -  JAN 2015
dorrtest "m" "-j 2014-12-01 -k 2015-01-01 -b ${BUD}" "ProcessDeposits"
# JUL 2016  -  AUG 2016
dorrtest "p" "${RRDATERANGE} -b ${BUD} -r 11" "GSR"
dorrtest "m" "${RRDATERANGE} -b ${BUD} " "Process"
dorrtest "n" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "o" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "q" "-r 12,11,RA001,2016-07-04 -b ${BUD}" "AccountBalance"
dorrtest "q1" "-r 12,9,RA001,2016-07-04 -b ${BUD}" "AccountBalance"
dorrtest "r" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"


echo "STARTING RENTROLL SERVER"
startRentRollServer
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/accounts/1" "request" "s"  "WebService--ChartOfAccounts"
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22LastName%22%2C%22direction%22%3A%22asc%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/transactants/1" "request" "t"  "WebService--GetTransactants"
echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/rentables/1" "request" "u"  "WebService--GetRentables"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck
