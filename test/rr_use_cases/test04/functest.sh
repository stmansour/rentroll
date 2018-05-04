#!/bin/bash

TESTNAME="Test04 - Multiple Rental Agreements for same Rentable for different user"
TESTSUMMARY="Generates separate databases for multiple rental agreements for different user"

RRBIN="../../../tmp/rentroll"
RR_DB_CORE="../"

source ../../share/base.sh
source ../rr_base.sh

# ------------------------------------------------------------------------------
#  TEST 4
#  Multiple Rental Agreements for same rentable for different users

# Scanario:
#   "Alex Vahabzadeh" has an agreement from 1 Jan, 2016 to 30 Jun, 2016
#   for "311, Rexford" Rentable with "Rex-311" style,
#   agrees to pay the rent at 3000$ contract rate amount at monthly.
#
#   "Chris Depp" signs the agreement for the above same rentable,
#   from 1st Jul, 2016 to 1 Jan, 2017 at contract rate 3000$ at monthly.
#
# Expactation:
#   - Rentroll report should show two RA entries with different users/payors
# ------------------------------------------------------------------------------

# look for the report for the entire agreement period
RRDATERANGE="-j 2016-01-01 -k 2017-01-01" # yyyy-mm-dd

# This test requires UTC timezone, so ensure that no timezone is specified in config.json
cp config.json ${RRBIN}/config.json

# start the web server
echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

# setup some ready-to-go basic infrastructure
dbcore

# now load rentable, rentabletypes, rentalAgreements
docsvtest "i4" "-R rt4.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j4" "-r r4.csv -L 6,${BUD}" "Rentables"
docsvtest "k4" "-C ra4.csv -L 9,${BUD}" "RentalAgreements"

# create Assessment from 1st Jan 2016 to 1st Jun 2016 with all instances, with contract Rent at 3000$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22Amount%22%3A3000%2C%22Start%22%3A%221%2F1%2F2016%22%2C%22Stop%22%3A%227%2F1%2F2016%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws401"  "WebService--AddAssessments-01Jan2016-01Jun2017"

# create Assessment from 1st Aug 2016 to 1st Jan 2017 with all instances, with contract Rent at 3000$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A2%2C%22Amount%22%3A3000%2C%22Start%22%3A%227%2F1%2F2016%22%2C%22Stop%22%3A%221%2F1%2F2017%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws402"  "WebService--AddAssessments-01Aug2016-01Jan2017"

# Do a text version of the Journal to make sure all the funds are properly transferred
dorrtest "l4" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"

# generate the rentroll report from 1st of April to 3rd of June
dorrtest "m4" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

# dump the database with structures
mysqldump --no-defaults rentroll > rrDumpTest04.sql

stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
