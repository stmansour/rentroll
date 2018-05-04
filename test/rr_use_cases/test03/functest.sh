#!/bin/bash

TESTNAME="Test03 - Multiple Rental Agreements for same user with vacancy"
TESTSUMMARY="Generates separate databases for multiple rental agreements for same user"

RRBIN="../../../tmp/rentroll"
RR_DB_CORE="../"

source ../../share/base.sh
source ../rr_base.sh

# ------------------------------------------------------------------------------
#  TEST 3
#  Multiple Rental Agreements for same user

# Scanario:
#   "Lauren Beck" has an agreement from 1 Jan, 2016 to 1 Jul, 2016
#   for "311, Rexford" Rentable with "Rex-311" style,
#   agrees to pay the rent at 3000$ contract rate (MR: 3000$) amount at monthly.
#
#   "Lauren Beck" reconsider the agreement, continue to stay at the same
#   Rentable for another 5 months from 1 Aug, 2016.
#
#   Just for the sake of testing purpose, we're not making any payments
#   but we're calculating assessment for two different RAs.
#
# Expactation:
#   - Rentroll report should show two RA entries for the Rentable "311, Rexford"
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
docsvtest "i3" "-R rt3.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j3" "-r r3.csv -L 6,${BUD}" "Rentables"
docsvtest "k3" "-C ra3.csv -L 9,${BUD}" "RentalAgreements"

# create Assessment from 1st Jan 2016 to 1st Jun 2016 with all instances, with contract Rent at 3000$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22Amount%22%3A3000%2C%22Start%22%3A%221%2F1%2F2016%22%2C%22Stop%22%3A%227%2F1%2F2016%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws301"  "WebService--AddAssessments-01Jan2016-01Jan2017"

# create Assessment from 1st Aug 2016 to 1st Jan 2017 with all instances, with contract Rent at 3000$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A2%2C%22Amount%22%3A3000%2C%22Start%22%3A%228%2F1%2F2016%22%2C%22Stop%22%3A%221%2F1%2F2017%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws302"  "WebService--AddAssessments-01Aug2016-01Jan2017"

# Do a text version of the Journal to make sure all the funds are properly transferred
dorrtest "l3" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"

# generate the rentroll report from 1st of April to 3rd of June
dorrtest "m3" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

# dump the database with structures
mysqldump --no-defaults rentroll > rrDumpTest03.sql

stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
