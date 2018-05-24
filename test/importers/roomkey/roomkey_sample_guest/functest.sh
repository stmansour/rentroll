#!/bin/bash

TESTNAME="RoomKey Import with Guest Data Export"
TESTSUMMARY="Tests initizing RentRoll DB from importing RoomKey rentroll report."
BUD="OKC"

RRBIN="../../../../tmp/rentroll"
TEMPCSVSTORE="${RRBIN}/importers/roomkey/temp_CSVs"

source ../../../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"

# remove all csv files from temp store
rm -f ${TEMPCSVSTORE}/rentableTypes_*.csv ./rentableTypes_*.csv
rm -f ${TEMPCSVSTORE}/people_*.csv ./people_*.csv
rm -f ${TEMPCSVSTORE}/rentable_*.csv ./rentable_*.csv
rm -f ${TEMPCSVSTORE}/rentalAgreement_*.csv ./rentalAgreement_*.csv

# call loader
doRoomKeyTest "b" "-csv ./roomkey.csv -bud ${BUD} -guestinfo ./guestdataexport.csv -testmode 1 -debug 1" "RoomKeyRentrollCSV"

# Print out All the different data types for validation
docsvtest "c" "-L 5,${BUD}" "RentableTypes"
docsvtest "d" "-L 7,${BUD}" "People"
docsvtest "e" "-L 6,${BUD}" "Rentables"
docsvtest "f" "-L 9,${BUD}" "RentalAgreements"

logcheck
