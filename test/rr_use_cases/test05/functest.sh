#!/bin/bash

TESTNAME="Test05 - New Payor/Users within existing RentalAgreement"
TESTSUMMARY="Generates separate databases for new Payors/Users on existing RA"

RRBIN="../../../tmp/rentroll"
RR_DB_CORE="../"

source ../../share/base.sh
source ../rr_base.sh

# ------------------------------------------------------------------------------
#  TEST 5
#  Adding new Payors and Users on existing RentalAgreement

# Scanario:
#   "Rita Costea" has an agreement from 1 Jan, 2016 to 01 Jan, 2017
#   for "309 1/2, Rexford" Rentable with "Rex-309-1/2" style,
#   agrees to pay the rent at 3500$ contract rate amount at monthly.
#
#   On 15th Jun, 2016 one of relative, "Daniel Costea" with one child joins
#   same Rentable with same RA signed by "Rita Costea"
#
#   Just for the sake of testing purpose, let's assume that June month rent
#   paid by Rita and July month rent is paid by Daniel.
#
# Expactation:
#   - from the July, report should show two entries in Payors, Users columns
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
docsvtest "i5" "-R rt5.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j5" "-r r5.csv -L 6,${BUD}" "Rentables"
docsvtest "k5" "-C ra5.csv -L 9,${BUD}" "RentalAgreements"

# now add Daniel Costea in current RA occupied by Rita with web service as Payor
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22tcidRAPayorPicker%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22TCID%22%3A4%2C%22RAID%22%3A1%2C%22FirstName%22%3A%22Daniel%22%2C%22MiddleName%22%3A%22%22%2C%22LastName%22%3A%22Costea%22%2C%22IsCompany%22%3Afalse%2C%22CompanyName%22%3A%22%22%2C%22DtStart%22%3A%226%2F15%2F2016%22%2C%22DtStop%22%3A%221%2F1%2F2017%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/rapayor/1/1" "request" "ws501" "WebService--AddPayor-DanielCostea-15Jun2016-01Jan2017"

# now add Daniel Costea in current RA occupied by Rita with web service as User
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22tcidRUserPicker%22%2C%22record%22%3A%7B%22recid%22%3A1%2C%22BID%22%3A1%2C%22TCID%22%3A4%2C%22RID%22%3A1%2C%22RentableName%22%3A%22309%201%2F2%20Rexford%22%2C%22FirstName%22%3A%22Daniel%22%2C%22MiddleName%22%3A%22%22%2C%22LastName%22%3A%22Costea%22%2C%22IsCompany%22%3Afalse%2C%22CompanyName%22%3A%22%22%2C%22DtStart%22%3A%226%2F15%2F2016%22%2C%22DtStop%22%3A%221%2F1%2F2017%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/ruser/1/1" "request" "ws502" "WebService--AddUser-DanielCostea-15Jun2016-01Jan2017"

# now add Child1 Costea in current RA occupied by Rita with web service as User
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22tcidRUserPicker%22%2C%22record%22%3A%7B%22recid%22%3A1%2C%22BID%22%3A1%2C%22TCID%22%3A5%2C%22RID%22%3A1%2C%22RentableName%22%3A%22309%201%2F2%20Rexford%22%2C%22FirstName%22%3A%22Child1%22%2C%22MiddleName%22%3A%22%22%2C%22LastName%22%3A%22Costea%22%2C%22IsCompany%22%3Afalse%2C%22CompanyName%22%3A%22%22%2C%22DtStart%22%3A%226%2F15%2F2016%22%2C%22DtStop%22%3A%221%2F1%2F2017%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/ruser/1/1" "request" "ws503" "WebService--AddUser-Child1Costea-15Jun2016-01Jan2017"

# create Assessment from 1st Jan 2016 to 1st Jan 2017 with all instances, with contract Rent at 3500$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22ARID%22%3A26%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22Amount%22%3A3500%2C%22Start%22%3A%221%2F1%2F2016%22%2C%22Stop%22%3A%221%2F1%2F2017%22%2C%22InvoiceNo%22%3A0%2C%22Comment%22%3A%22%22%2C%22ReverseMode%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws504"  "WebService--AddAssessments-01Jan2016-01Jan2017"

# ---------------------------
#       JUNE PAYMENT - RITA
# ---------------------------
# create a receipt for the rent assessment of Jun 2016, paid on 1st of Jun with Amount of 3500$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%226%2F1%2F2016%22%2C%22DocNo%22%3A%22234234234_1%22%2C%22Payor%22%3A%22Rita%20Costea%20(TCID%3A%203)%22%2C%22TCID%22%3A3%2C%22Amount%22%3A3500%2C%22Comment%22%3A%22Receipt%20Paid%20by%20Rita%20Costea%20for%20Jun%2C%202016%20Rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "ws505"  "WebService--AddReceipt-01Jun2016"

# deposit the amount from the receipt created in abobe step, on 1st of Jun
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%226%2F1%2F2016%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A3500%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "ws506"  "WebService--Deposit-01Jun2016"

# Now, apply the funds towards Jun month Assessment on 2nd of Jun (was assessed on 1st Jun)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A3%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%226%2F2%2F2016%22%2C%22ASMID%22%3A7%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Rent%20Non-Taxable%22%2C%22Amount%22%3A3500%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3500%2C%22Dt%22%3A%226%2F2%2F2016%22%2C%22Allocate%22%3A3500%2C%22Date_%22%3A%222016-06-02T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222016-06-02T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "ws507"  "WebService--ApplyThePayment-02Jun2016"
# ---------------------------

# ---------------------------
#      JULY PAYMENT - DANIEL
# ---------------------------
# create a receipt for the rent assessment of Jul 2016, paid on 1st of Jul with Amount of 3500$
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%227%2F1%2F2016%22%2C%22DocNo%22%3A%22234234234_2%22%2C%22Payor%22%3A%22Daniel%20Costea%20(TCID%3A%204)%22%2C%22TCID%22%3A4%2C%22Amount%22%3A3500%2C%22Comment%22%3A%22Receipt%20Paid%20by%20Daniel%20Costea%20for%20Jul%2C%202016%20Rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "ws508"  "WebService--AddReceipt-01Jul2016"

# deposit the amount from the receipt created in abobe step, on 1st of Jul
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B2%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%227%2F1%2F2016%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A3500%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "ws509"  "WebService--Deposit-01Jul2016"

# Now, apply the funds towards Jul month Assessment on 2nd of Jul (was assessed on 1st Jul)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A4%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%227%2F2%2F2016%22%2C%22ASMID%22%3A8%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Rent%20Non-Taxable%22%2C%22Amount%22%3A3500%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3500%2C%22Dt%22%3A%227%2F2%2F2016%22%2C%22Allocate%22%3A3500%2C%22Date_%22%3A%222016-07-02T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222016-07-02T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "ws510"  "WebService--ApplyThePayment-02Jul2016"
# ---------------------------

# Do a text version of the Journal to make sure all the funds are properly transferred
dorrtest "l5" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"

# generate the rentroll report from 1st of April to 3rd of June
dorrtest "m5" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

# dump the database with structures
mysqldump --no-defaults rentroll > rrDumpTest05.sql

stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
