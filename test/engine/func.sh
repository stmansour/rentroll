#!/bin/bash
# This is a quick functional test for the rentroll engine
# It uses the values initialized in directory ../ledger1, generates
# Journal and Ledger records, generates the reports, and validates that
# the reports are what we expect

TESTNAME="Basic Engine Test"
TESTSUMMARY="Generate journals and ledgers for 1 month. Tests proration, Invoice and Deposit, CustomAttributes, AssessmentChecker"
RRDATERANGE="-j 2015-11-01 -k 2015-12-01"

source ../share/base.sh

mysql ${MYSQLOPTS} <init.sql
if [ $? -eq 0 ]; then
	echo "Init was successful"
else
	echo "INIT HAD ERRORS"
	exit 1
fi

dorrtest "a1" "-j 2015-11-01 -k 2015-11-02 -b ${BUD} -x" "Process"
dorrtest "a2" "-j 2015-11-21 -k 2015-11-22 -b ${BUD} -x" "Process"
docsvtest "a3" "-e rcpt201501.csv -G ${BUD} -g 11/1/15,12/1/15 -L 13,${BUD}" "Receipts-2015-NOV"
dorrtest "a" "${RRDATERANGE} -b ${BUD} -x" "Process"
dorrtest "b" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "c" "${RRDATERANGE} -b ${BUD} -r 2" "Ledger"
dorrtest "d" "${RRDATERANGE} -b ${BUD} -r 5" "AssessmentChecker"
dorrtest "e" "${RRDATERANGE} -b ${BUD} -r 6" "LedgerBalance"

docsvtest "f" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "g" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
docsvtest "h" "-d depository.csv -y deposit.csv -L 19,${BUD}" "Deposits"

dorrtest "i" "${RRDATERANGE} -b ${BUD} -r 7" "CountRentables"
docsvtest "j" "-i invoice.csv -L 20,${BUD}" "CreateInvoice"
dorrtest "k" "${RRDATERANGE} -b ${BUD} -r 9,IN00001" "InvoiceReport"

logcheck
