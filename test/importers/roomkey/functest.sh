#!/bin/bash

TESTNAME="RoomKey Import"
TESTSUMMARY="Tests initizing RentRoll DB from importing RoomKey rentroll report."
BUD="ISO"

RRBIN="../../../tmp/rentroll"
TEMPCSVSTORE="${RRBIN}/importers/roomkey/temp_CSVs"

source ../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"

# remove all csv files from temp store
rm -f ${TEMPCSVSTORE}/rentableTypes_*.csv ./rentableTypes_*.csv
rm -f ${TEMPCSVSTORE}/people_*.csv ./people_*.csv
rm -f ${TEMPCSVSTORE}/rentable_*.csv ./rentable_*.csv

# call loader
doRoomKeyTest "c" "-csv ./roomkey.csv -bud ${BUD} -testmode 1" "RoomKeyRentrollCSV"

# Print out All the different data types for validation
docsvtest "d" "-L 5,${BUD}" "RentableTypes"
docsvtest "e" "-L 7,${BUD}" "People"
docsvtest "f" "-L 6,${BUD}" "Rentables"

logcheck
