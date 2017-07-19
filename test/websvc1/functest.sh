#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Assessments, Receipts, and Reversals"
TESTSUMMARY="Test Web Services"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < asmtest.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
startRentRollServer

# Create a non-recurring assessment
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A%7B%22id%22%3A4%2C%22text%22%3A%22Electric+Base+Fee+Assessment%22%7D%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A%7B%22id%22%3A%22REX%22%2C%22text%22%3A%22REX%22%7D%2C%22Start%22%3A%227%2F1%2F2017%22%2C%22Stop%22%3A%227%2F1%2F2017%22%2C%22RentCycle%22%3A%7B%22id%22%3A%22Norecur%22%2C%22text%22%3A%22Norecur%22%7D%2C%22ProrationCycle%22%3A%7B%22id%22%3A%22Norecur%22%2C%22text%22%3A%22Norecur%22%7D%2C%22TCID%22%3A0%2C%22Amount%22%3A30%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309+S+Rexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a0"  "WebService--Create_NonRecurring_Assessment"

# Create a recurring rent assessment
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A%7B%22id%22%3A2%2C%22text%22%3A%22Non-Taxable+Rent+Assessment%22%7D%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A%7B%22id%22%3A%22REX%22%2C%22text%22%3A%22REX%22%7D%2C%22Start%22%3A%221%2F1%2F2017%22%2C%22Stop%22%3A%221%2F1%2F2018%22%2C%22RentCycle%22%3A%7B%22id%22%3A%22Monthly%22%2C%22text%22%3A%22Monthly%22%7D%2C%22ProrationCycle%22%3A%7B%22id%22%3A%22Daily%22%2C%22text%22%3A%22Daily%22%7D%2C%22TCID%22%3A0%2C%22Amount%22%3A3750%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309+S+Rexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a1"  "WebService--Create_Recurring_Assessment"

# Create a non-recurring assessment to reverse...
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A%7B%22id%22%3A5%2C%22text%22%3A%22Water+and+Sewer+Base+Fee%22%7D%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A%7B%22id%22%3A%22REX%22%2C%22text%22%3A%22REX%22%7D%2C%22Start%22%3A%227%2F19%2F2017%22%2C%22Stop%22%3A%227%2F19%2F2017%22%2C%22RentCycle%22%3A%7B%22id%22%3A%22Norecur%22%2C%22text%22%3A%22Norecur%22%7D%2C%22ProrationCycle%22%3A%7B%22id%22%3A%22Norecur%22%2C%22text%22%3A%22Norecur%22%7D%2C%22TCID%22%3A0%2C%22Amount%22%3A40%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309+S+Rexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a2"  "WebService--Create_nonrecurring_Assessment_to_reverse"

# Do the reversal...
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22asmInstForm%22%2C%22ASMID%22%3A10%2C%22ReverseMode%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/10" "request" "a3"  "WebService--Reverse_nonrecurring_assessment"

# Create a recurring assessment to reverse...
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A%7B%22id%22%3A1%2C%22text%22%3A%22Taxable+Rent+Assessment%22%7D%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A%7B%22id%22%3A%22REX%22%2C%22text%22%3A%22REX%22%7D%2C%22Start%22%3A%221%2F1%2F2017%22%2C%22Stop%22%3A%221%2F1%2F2018%22%2C%22RentCycle%22%3A%7B%22id%22%3A%22Monthly%22%2C%22text%22%3A%22Monthly%22%7D%2C%22ProrationCycle%22%3A%7B%22id%22%3A%22Daily%22%2C%22text%22%3A%22Daily%22%7D%2C%22TCID%22%3A0%2C%22Amount%22%3A4000%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309+S+Rexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a4"  "WebService--Create_recurring_Assessment_to_reverse"

# Do the reversal...
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22asmEpochForm%22%2C%22ASMID%22%3A12%2C%22ReverseMode%22%3A2%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/12" "request" "a5"  "WebService--Reverse_recurring_assessment"

# Add a $100 check...
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A19%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A1%2C%22PmtTypeName%22%3A1%2C%22Dt%22%3A%227%2F19%2F2017%22%2C%22DocNo%22%3A%224321%22%2C%22Payor%22%3A%22Aaron%20Read%20(TCID%3A%201)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A100%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "a6"  "WebService--100_USD_check"

# Add a $1000 wire transfer...
echo "%7B%22cmd%22%3A%22save%22%2C%22formname%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A19%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A3%2C%22PmtTypeName%22%3A3%2C%22Dt%22%3A%227%2F19%2F2017%22%2C%22DocNo%22%3A%2248762876342%22%2C%22Payor%22%3A%22Aaron%20Read%20(TCID%3A%201)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A1000%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "a7"  "WebService--1000_USD_wire"

# Add a $25,180 ACH...
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A19%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22Dt%22%3A%227%2F19%2F2017%22%2C%22DocNo%22%3A%224645727272%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A25180%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "a8"  "WebService--25180_USD_ACH"

# Allocate these funds to the assessments...
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A1%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%221%2F1%2F2017%22%2C%22ASMID%22%3A3%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%221%2F1%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-01-01T08%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%22changes%22%3A%7B%22Dt%22%3A%221%2F1%2F2017%22%7D%7D%7D%2C%7B%22recid%22%3A1%2C%22Date%22%3A%222%2F1%2F2017%22%2C%22ASMID%22%3A4%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%222%2F1%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-02-01T08%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%22changes%22%3A%7B%22Dt%22%3A%222%2F1%2F2017%22%7D%7D%7D%2C%7B%22recid%22%3A2%2C%22Date%22%3A%223%2F1%2F2017%22%2C%22ASMID%22%3A5%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%223%2F1%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-03-01T08%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%22changes%22%3A%7B%22Dt%22%3A%223%2F1%2F2017%22%7D%7D%7D%2C%7B%22recid%22%3A3%2C%22Date%22%3A%224%2F1%2F2017%22%2C%22ASMID%22%3A6%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%224%2F4%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-04-01T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%22changes%22%3A%7B%22Dt%22%3A%224%2F4%2F2017%22%7D%7D%7D%2C%7B%22recid%22%3A4%2C%22Date%22%3A%225%2F1%2F2017%22%2C%22ASMID%22%3A7%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%225%2F1%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-05-01T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%22changes%22%3A%7B%22Dt%22%3A%225%2F1%2F2017%22%7D%7D%7D%2C%7B%22recid%22%3A5%2C%22Date%22%3A%226%2F1%2F2017%22%2C%22ASMID%22%3A8%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%226%2F1%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-06-01T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%22changes%22%3A%7B%22Dt%22%3A%226%2F1%2F2017%22%7D%7D%7D%2C%7B%22recid%22%3A6%2C%22Date%22%3A%227%2F1%2F2017%22%2C%22ASMID%22%3A1%2C%22ARID%22%3A4%2C%22Assessment%22%3A%22Electric%20Base%20Fee%20Assessment%22%2C%22Amount%22%3A30%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A30%2C%22Dt%22%3A%227%2F1%2F2017%22%2C%22Allocate%22%3A30%2C%22Date_%22%3A%222017-07-01T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%7D%7D%2C%7B%22recid%22%3A7%2C%22Date%22%3A%227%2F1%2F2017%22%2C%22ASMID%22%3A9%2C%22ARID%22%3A2%2C%22Assessment%22%3A%22Non-Taxable%20Rent%20Assessment%22%2C%22Amount%22%3A3750%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A3750%2C%22Dt%22%3A%227%2F1%2F2017%22%2C%22Allocate%22%3A3750%2C%22Date_%22%3A%222017-07-01T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-07-19T07%3A00%3A00.000Z%22%2C%22w2ui%22%3A%7B%7D%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1/" "request" "a9"  "WebService--Allocate_Funds"

# Look at a LedgerBalance report to make sure all the accounts are correct...
$(curl -s "http://localhost:8270/wsvc/211/1?r=RPTla&dtstart=2017-01-01&dtstop=2017-08-01" >a10)
doValidateFile "a10" "WebService--Ledger_Activity_Report"

# Reverse receipt #1
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22receiptForm%22%2C%22RCPTID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/1" "request" "a11"  "WebService--Reverse_Receipt_1"

# Look at a LedgerBalance after reversing the receipt to make sure all the accounts are correct...
$(curl -s "http://localhost:8270/wsvc/211/1?r=RPTla&dtstart=2017-01-01&dtstop=2017-08-01" >a12)
doValidateFile "a12" "WebService--Ledger_Activity_Report"



stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck
