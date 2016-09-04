#!/bin/bash

TESTNAME="Rexford March Books"
TESTSUMMARY="Tested values for Rexford properties March 2016. Also tests: DepositMethods, Invoice, Statements, Sources"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "d" "-d depository.csv -L 18,${BUD}" "Depositories"
docsvtest "c" "-m dm.csv -L 23,${BUD}" "DepositMethods"
docsvtest "e" "-S sources.csv -L 24,${BUD}" "Sources"
docsvtest "f" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "g" "-p people.csv  -L 7" "People"
docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "i" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "j" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "k" "-E pets.csv -L 16,RA0001" "Pets"
docsvtest "l" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "m" "-A asmt.csv -L 11,${BUD}" "Assessments"
docsvtest "n" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "o" "-e rcpt.csv -L 13,${BUD}" "Receipts"
docsvtest "p" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "q" "-U assigncustom.csv -L 15" "AssignCustomAttributes"

dorrtest "r" "-j 2016-03-01 -k 2016-04-01 -b ${BUD}" "Process"
dorrtest "s" "-j 2016-03-01 -k 2016-04-01 -b ${BUD} -r 1" "March-Journal"	# Journals
dorrtest "t" "-j 2016-03-01 -k 2016-04-01 -b ${BUD} -r 2" "March-Ledger"	# Ledgers
dorrtest "u" "-j 2016-03-01 -k 2016-04-01 -b ${BUD} -r 8" "Statements"	# Statements
docsvtest "v" "-i invoice.csv -L 20,REX" "CreateInvoice"
dorrtest  "w" "-j 2016-03-01 -k 2016-04-01 -b ${BUD} -r 9,IN0001" "Invoice"

logcheck
