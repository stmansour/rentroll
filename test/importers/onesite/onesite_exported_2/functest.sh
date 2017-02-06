#!/bin/bash

TESTNAME="Onesite Import (Exported CSV no. 2)"
TESTSUMMARY="Tests initizing RentRoll DB from importing OneSite rentroll report."

RRBIN="../../../../tmp/rentroll"
TEMPCSVSTORE="${RRBIN}/importers/onesite/temp_CSVs"
BUD=ISO
source ../../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"

# remove all csv files from temp store
rm -f ${TEMPCSVSTORE}/rentableTypes_*.csv ./rentableTypes_*.csv
rm -f ${TEMPCSVSTORE}/people_*.csv ./people_*.csv
rm -f ${TEMPCSVSTORE}/rentable_*.csv ./rentable_*.csv
rm -f ${TEMPCSVSTORE}/rentalAgreement_*.csv ./rentalAgreement_*.csv
rm -f ${TEMPCSVSTORE}/customAttribute_*.csv ./customAttribute_*.csv

# call loader
doOnesiteTest "c" "-csv ./onesite_2.csv -bud ${BUD} -testmode 1" "OnesiteRentrollCSV"

# Print out All the different data types for validation
docsvtest "d" "-L 5,${BUD}" "RentableTypes"
docsvtest "e" "-L 7,${BUD}" "People"
docsvtest "f" "-L 6,${BUD}" "Rentables"
docsvtest "h" "-L 9,${BUD}" "RentalAgreements"
docsvtest "k" "-L 14,${BUD}" "CustomAttribute"
docsvtest "l" "-L 15,${BUD}" "CustomAttributeRef"

logcheck

