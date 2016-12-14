#!/bin/bash

TESTNAME="Onesite Import"
TESTSUMMARY="Tests initizing RentRoll DB from importing OneSite rentroll report."

RRBIN="../../../tmp/rentroll"
source ../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"

# Sudip this is a quick hack to geth things into a known state
# rework this code as needed so it is more production-ready
rm -f /tmp/onesite/rentabletypes_*.csv ./rentabletypes_*.csv

${RRBIN}/onesiteLoad -csv ./onesite.csv -bud ${BUD} >c 2>&1

# Sudip  -- please update this code as needed.
cp /tmp/onesite/*.csv .

# Sudip this is a quick hack to validate the file you're producing. Please
# rework this code as needed so it is more production-ready
mv rentabletypes_* rt.csv
docsvtest "f" "-R rt.csv -L 5,${BUD}" "RentableTypes"

# docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
# docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
# docsvtest "j" "-C ra.csv -L 9,${BUD}" "RentalAgreements"

logcheck
