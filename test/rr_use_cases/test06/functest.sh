#!/bin/bash

TESTNAME="Test06 - Different Users/Payors for a RentalAgreement"
TESTSUMMARY="Generates separate databases for different Payors/Users on RA"

RRBIN="../../../tmp/rentroll"
RR_DB_CORE="../"

source ../../share/base.sh
source ../rr_base.sh

# ------------------------------------------------------------------------------
#  TEST 6
#  Payor is paying rent for their children.

# Scanario:
#   "Kevin Mills" has an agreement from 1 Jan, 2016 to 01 Jan, 2017
#   for "312 Rexford" Rentable with "Rex-312" style,
#   agrees to pay the rent at 2800$ contract rate amount at monthly.
#   Kevin Mills is Payor on this RA but actually paying rent for her
#   children (Child1, Child2 Mills).
#
#   Just for the sake of testing purpose, we're considering only two
#   Receipts.
#
# Expactation:
#   - Report should show two users, "Kevin Mills" as Payor
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
docsvtest "i6" "-R rt6.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j6" "-r r6.csv -L 6,${BUD}" "Rentables"
docsvtest "k6" "-C ra6.csv -L 9,${BUD}" "RentalAgreements"

# create Assessment from 1st Jan 2016 to 1st Jan 2017 with all instances, with contract Rent at 2800$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22Amount%22%3A2800%2C%22Start%22%3A%221%2F1%2F2016%22%2C%22Stop%22%3A%221%2F1%2F2017%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws601"  "WebService--AddAssessments-01Jan2016-01Jan2017"

# ---------------------------
#       JANUARY PAYMENT
# ---------------------------
# create a receipt for the rent assessment of Jan 2016, paid on 1st of Jan with Amount of 2800$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F1%2F2016%22%2C%22DocNo%22%3A%22234234234_1%22%2C%22Payor%22%3A%22Kevin%20Mills%20(TCID%3A%208)%22%2C%22TCID%22%3A8%2C%22Amount%22%3A2800%2C%22Comment%22%3A%22Receipt%20Paid%20by%20Kevin%20Mills%20for%20Jan%2C%202016%20Rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "ws602"  "WebService--AddReceipt-01Jan2016"

# deposit the amount from the receipt created in abobe step, on 1st of Jan
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%221%2F1%2F2016%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A2800%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "ws603"  "WebService--Deposit-01Jan2016"

# Now, apply the funds towards Jan month Assessment on 2nd of Jan (was assessed on 1st Jan)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A8%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%221%2F2%2F2016%22%2C%22ASMID%22%3A2%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Rent%20Non-Taxable%22%2C%22Amount%22%3A2800%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A2800%2C%22Dt%22%3A%221%2F2%2F2016%22%2C%22Allocate%22%3A2800%2C%22Date_%22%3A%222016-01-02T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222016-01-02T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "ws604"  "WebService--ApplyThePayment-02Jan2016"
# ---------------------------

# ---------------------------
#      FEBRUARY PAYMENT
# ---------------------------
# create a receipt for the rent assessment of Feb 2016, paid on 1st of Feb with Amount of 3000$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%222%2F1%2F2016%22%2C%22DocNo%22%3A%22234234234_2%22%2C%22Payor%22%3A%22Kevin%20Mills%20(TCID%3A%208)%22%2C%22TCID%22%3A8%2C%22Amount%22%3A2800%2C%22Comment%22%3A%22Receipt%20Paid%20by%20Kevin%20Mills%20for%20Feb%2C%202016%20Rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "ws605"  "WebService--AddReceipt-01Feb2016"

# deposit the amount from the receipt created in abobe step, on 1st of Feb
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B2%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%222%2F1%2F2016%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A2800%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "ws606"  "WebService--Deposit-01Feb2016"

# Now, apply the funds towards Feb month Assessment on 2nd of Feb (was assessed on 1st Feb)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A8%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%222%2F2%2F2016%22%2C%22ASMID%22%3A3%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Rent%20Non-Taxable%22%2C%22Amount%22%3A2800%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A2800%2C%22Dt%22%3A%222%2F2%2F2016%22%2C%22Allocate%22%3A2800%2C%22Date_%22%3A%222016-02-02T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222016-02-02T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "ws607"  "WebService--ApplyThePayment-02Feb2016"
# ---------------------------

# Do a text version of the Journal to make sure all the funds are properly transferred
dorrtest "l6" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"

# generate the rentroll report from 1st of April to 3rd of June
dorrtest "m6" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

# dump the database with structures
mysqldump --no-defaults rentroll > rrDumpTest06.sql

stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
