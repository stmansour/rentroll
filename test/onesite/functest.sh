#!/bin/bash

TESTNAME="Onesite Import"
TESTSUMMARY="Tests initizing RentRoll DB from importing OneSite rentroll report."

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"

# Sudip this is a quick hack to geth things into a known state
# rework this code as needed so it is more production-ready
rm -f /tmp/onesite/rentabletypes*.csv ./rentabletypes*.csv

${RRBIN}/onesiteLoad -i ./onesite.csv >c 2>&1

# Sudip  -- please update this code as needed.
cp /tmp/onesite/*.csv .

# Sudip -- i'm using rt2 as a manually fixed version of what your code produces
#          after you fix things up, replace rt2.csv with the file your code generates

# Sudip this is a quick hack to validate the file you're producing. Please
# rework this code as needed so it is more production-ready
mv rentabletypes_20* rt.csv
docsvtest "f" "-R rt.csv -L 5,${BUD}" "RentableTypes"

# docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
# docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
# docsvtest "j" "-C ra.csv -L 9,${BUD}" "RentalAgreements"

logcheck
