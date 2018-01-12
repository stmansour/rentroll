#!/bin/bash

TESTNAME="Onesite Import (Exported CSV no. 1) With Market Rent Value"
TESTSUMMARY="Tests initizing RentRoll DB from importing OneSite rentroll report."

RRBIN="../../../../tmp/rentroll"
TEMPCSVSTORE="${RRBIN}/importers/onesite/temp_CSVs"
BUD=OKC
source ../../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
# docsvtest "z" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"

# remove all csv files from temp store
rm -f ${TEMPCSVSTORE}/rentableTypes_*.csv ./rentableTypes_*.csv
rm -f ${TEMPCSVSTORE}/people_*.csv ./people_*.csv
rm -f ${TEMPCSVSTORE}/rentable_*.csv ./rentable_*.csv
rm -f ${TEMPCSVSTORE}/rentalAgreement_*.csv ./rentalAgreement_*.csv
rm -f ${TEMPCSVSTORE}/customAttribute_*.csv ./customAttribute_*.csv

# call loader
doOnesiteTest "b" "-csv ./onesite_mr_1.csv -bud ${BUD} -testmode 1" "OnesiteRentrollCSV"

# Print out All the different data types for validation
docsvIgnoreDatesTest "c" "-L 5,${BUD}" "RentableTypes"
docsvtest "d" "-L 7,${BUD}" "People"
docsvtest "e" "-L 6,${BUD}" "Rentables"
docsvtest "f" "-L 9,${BUD}" "RentalAgreements"
docsvtest "h" "-L 14,${BUD}" "CustomAttribute"
docsvtest "k" "-L 15,${BUD}" "CustomAttributeRef"
docsvtest "y" "-P pmt.csv -L 12,${BUD}" "Payment Types"

logcheck

mysqldump --no-defaults rentroll >iso.sql
