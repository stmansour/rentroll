#!/bin/bash

TESTNAME="Onesite Import"
TESTSUMMARY="Tests initizing RentRoll DB from importing OneSite rentroll report."

RRBIN="../../../tmp/rentroll"
TEMPCSVSTORE="../../../importers/onesite/tempCSVs"

source ../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"

# remove all csv files from temp store
rm -f ${TEMPCSVSTORE}/rentableTypes_*.csv ./rentableTypes_*.csv
rm -f ${TEMPCSVSTORE}/people_*.csv ./people_*.csv
rm -f ${TEMPCSVSTORE}/rentable_*.csv ./rentable_*.csv
rm -f ${TEMPCSVSTORE}/rentalAgreement_*.csv ./rentalAgreement_*.csv

# call loader
${RRBIN}/onesiteLoad -csv ./onesite.csv -bud ${BUD} >c 2>&1
cp ${TEMPCSVSTORE}/*.csv .

mv rentableTypes_* rt.csv
mv people_* people.csv
mv rentable_* rentable.csv
mv rentalAgreement_* ra.csv

docsvtest "f" "-R rt.csv -L 5,${BUD}" "RentableTypes"
docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "j" "-C ra.csv -L 9,${BUD}" "RentalAgreements"

# clear files after work done
# rm rt.csv people.csv rentable.csv ra.csv

logcheck
