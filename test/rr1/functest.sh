#!/bin/bash

TESTNAME="Renter Changing During RentCycle"
TESTSUMMARY="Tests changing Rental Agreements multiple times during a rent cycle"

RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "c" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "d" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "e" "-p people.csv  -L 7" "People"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "g" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "h" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "i" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "j" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
docsvtest "k" "-A asmt.csv -L 11,${BUD}" "Assessments"
docsvtest "l" "-e rcpt.csv -L 13,${BUD}" "Receipts"

# process payments and receipts
dorrtest "m" "${RRDATERANGE} -b ${BUD}" "Process"
dorrtest "n" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "o" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"

logcheck
