#!/bin/bash

TESTNAME="Onesite Import"
TESTSUMMARY="Tests initizing RentRoll DB from importing OneSite rentroll report."

RRBIN="../../../tmp/rentroll"
TEMPCSVSTORE="../../../importers/onesite/tempCSVs"

source ../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"

rm -f ${TEMPCSVSTORE}/rentabletypes_*.csv ./rentabletypes_*.csv
${RRBIN}/onesiteLoad -csv ./onesite.csv -bud ${BUD} >c 2>&1
cp ${TEMPCSVSTORE}/*.csv .
mv rentabletypes_* rt.csv

docsvtest "f" "-R rt.csv -L 5,${BUD}" "RentableTypes"

# docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
# docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
# docsvtest "j" "-C ra.csv -L 9,${BUD}" "RentalAgreements"

logcheck
