#!/bin/bash

TESTNAME="Test01 - Rent Assessment Payment for June month"
TESTSUMMARY="Generates separate databases for Rent Assessment"

RRBIN="../../../tmp/rentroll"
RR_DB_CORE="../"

source ../../share/base.sh
source ../rr_base.sh

# ------------------------------------------------------------------------------
# TEST 1:
#   A simple straight-forward test case for rent receivables with all
#   assessments and two paid receipts.
#
# Scanario:
#   "Aaron Read" has an agreement from 1 Jan, 2016 to 1 Jan, 2017
#   for "309, Rexford" Rentable with "Rex1" style,
#   living with "Kirsten Read" and one child,
#   agrees to pay the rent at 3500$ contract rate (MR: 3500$) amount at monthly.
#
#   Just for the sake of testing purpose, we're only considering the payment for
#   first TWO months Rent Assessment (assessed at 1st day of a month).
#   Fund is applied on 2nd of a month after money deposited in bank.
#
# Expactation:
#   - Rentroll report show $42000 in GSR.
#   - Rentroll report Grand total (Ending Receivable should be $35000 as payment done
#     by two receipts with total amount of $7000)
# ------------------------------------------------------------------------------

# look for the report for the entire agreement period
RRDATERANGE="-j 2016-01-01 -k 2017-01-01" # yyyy-mm-dd
# CSVLOADRANGE="-G ${BUD} -g 6/1/17,7/1/17" # mm/d/yy, US format

# This test requires UTC timezone, so ensure that no timezone is specified in config.json
cp config.json ${RRBIN}/config.json

# start the web server
echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

# setup some ready-to-go basic infrastructure
dbcore

# now load rentable, rentabletypes, rentalAgreements
docsvtest "i1" "-R rt1.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j1" "-r r1.csv -L 6,${BUD}" "Rentables"
docsvtest "k1" "-C ra1.csv -L 9,${BUD}" "RentalAgreements"

# create an assessment from 1st Jan 2016 to 1st Jan 2017 with all instances with contract rent of 3500$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22Amount%22%3A3500%2C%22Start%22%3A%221%2F1%2F2016%22%2C%22Stop%22%3A%221%2F1%2F2017%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws101"  "WebService--AddAssessments-01Jan2016-01Jan2017"

# ---------------------------
#       JANUARY PAYMENT
# ---------------------------
# create a receipt for the rent assessment of Jan 2016, paid on 1st of Jan
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F1%2F2016%22%2C%22DocNo%22%3A%22234234234%22%2C%22Payor%22%3A%22Aaron%20Read%20(TCID%3A%201)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A3500%2C%22Comment%22%3A%22Receipt%20Paid%20by%20Aaron%20Read%20for%20Jan%2C%202016%20Rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "ws102"  "WebService--AddReceipt-01Jan2016"

# deposit the amount from the receipt created in abobe step, on 1st of Jan
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%221%2F1%2F2016%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A3500%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "ws103"  "WebService--Deposit-01Jan2016"

# Now, apply the funds towards Jan month Assessment on 2nd of Jan (was assessed on 1st Jan)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A1%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%221%2F2%2F2016%22%2C%22ASMID%22%3A2%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Rent%20Non-Taxable%22%2C%22Amount%22%3A3500%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3500%2C%22Dt%22%3A%221%2F2%2F2016%22%2C%22Allocate%22%3A3500%2C%22Date_%22%3A%222016-01-02T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222016-01-02T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "ws104"  "WebService--ApplyThePayment-02Jan2016"
# ---------------------------

# ---------------------------
#      FEBRUARY PAYMENT
# ---------------------------
# create a receipt for the rent assessment of Feb 2016, paid on 1st of Feb
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%222%2F1%2F2016%22%2C%22DocNo%22%3A%22234234234_2%22%2C%22Payor%22%3A%22Aaron%20Read%20(TCID%3A%201)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A3500%2C%22Comment%22%3A%22Receipt%20Paid%20by%20Aaron%20Read%20for%20Feb%2C%202016%20Rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "ws105"  "WebService--AddReceipt-01Feb2016"

# deposit the amount from the receipt created in abobe step, on 1st of Feb
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B2%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%222%2F1%2F2016%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A3500%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "ws106"  "WebService--Deposit-01Feb2016"

# Now, apply the funds towards Feb month Assessment on 2nd of Feb (was assessed on 1st Feb)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A1%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%222%2F2%2F2016%22%2C%22ASMID%22%3A3%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Rent%20Non-Taxable%22%2C%22Amount%22%3A3500%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3500%2C%22Dt%22%3A%222%2F2%2F2016%22%2C%22Allocate%22%3A3500%2C%22Date_%22%3A%222016-02-02T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222016-02-02T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "ws107"  "WebService--ApplyThePayment-02Feb2016"
# ---------------------------

# Do a text version of the Journal to make sure all the funds are properly transferred
dorrtest "l1" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"

# generate the rentroll report from 1st of June to 3rd of July
dorrtest "m1" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

# dump the database with structures
mysqldump --no-defaults rentroll > rrDumpTest01.sql

stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
