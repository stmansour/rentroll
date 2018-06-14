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
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

# Create a non-recurring assessment
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A4%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%227%2F1%2F2017%22%2C%22Stop%22%3A%227%2F1%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A30%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309%2BS%2BRexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a0"  "WebService--Create_NonRecurring_Assessment"

# Create a recurring rent assessment
#    -- set a known endpoint prior to date of creation so
#       that the recurrences will not change as time goes on.
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A2%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22BID%22%3A1%2C%22Start%22%3A%221%2F1%2F2017%22%2C%22Stop%22%3A%227%2F31%2F2017%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22TCID%22%3A0%2C%22Amount%22%3A3750%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309%2BS%2BRexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a1"  "WebService--Create_Recurring_Assessment"

# Create a non-recurring assessment to reverse...
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A5%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22BID%22%3A1%2C%22Start%22%3A%227%2F19%2F2017%22%2C%22Stop%22%3A%227%2F19%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A40%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309%2BS%2BRexford%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a2"  "WebService--Create_nonrecurring_Assessment_to_reverse"

# Do the reversal...
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22asmInstForm%22%2C%22ASMID%22%3A10%2C%22ReverseMode%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/10" "request" "a3"  "WebService--Reverse_nonrecurring_assessment"

# Create a recurring assessment to reverse...
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A1%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%221%2F1%2F2017%22%2C%22Stop%22%3A%227%2F31%2F2017%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22TCID%22%3A0%2C%22Amount%22%3A4000%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%227%2F19%2F2017%22%2C%22LastModBy%22%3A0%2C%22ExpandPastInst%22%3A1%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22Rentable%22%3A%22309%2BS%2BRexford%22%7D%7D" > request
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
$(curl -s "http://localhost:8270/v1/report/1?r=RPTla&dtstart=2017-01-01&dtstop=2017-08-01" >a10)
doValidateFile "a10" "WebService--Ledger_Activity_Report"

# Reverse receipt #1
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22receiptForm%22%2C%22RCPTID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/1" "request" "a11"  "WebService--Reverse_Receipt_1"

# Look at a LedgerBalance after reversing the receipt to make sure all the accounts are correct...
$(curl -s "http://localhost:8270/v1/report/1?r=RPTla&dtstart=2017-01-01&dtstop=2017-08-01" >a12)
doValidateFile "a12" "WebService--Ledger_Activity_Report"

# Force an error on Account Update. Try to make an account (Accounts Receivable - 11000) a summary
# account when it is called out by an AccountRule...
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%2C%22record%22%3A%7B%22LID%22%3A5%2C%22PLID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%2211000%22%2C%22Status%22%3A2%2C%22Type%22%3A0%2C%22Name%22%3A%22Accounts%2BReceivable%22%2C%22AcctType%22%3A%22Cash%22%2C%22AllowPost%22%3Afalse%2C%22Description%22%3A%22update%2Bby%2Bfunctional%2Btest%22%2C%22LastModTime%22%3A%222017-07-19T15%3A58%3A00Z%22%2C%22LastModBy%22%3A0%2C%22recid%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/account/1/5" "request" "a13"  "WebService--ERROR-Set_Incorrect_AllowPost"

# Read it back and make sure that AllowPosts is true
echo "%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%7D" > request
dojsonPOST "http://localhost:8270/v1/account/1/5" "request" "a14"  "WebService--ERROR-VRFY-1"

# Force an error on Account Update. Try to make a summary account (Cash - 10000) AllowPosts = true
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%2C%22record%22%3A%7B%22LID%22%3A1%2C%22PLID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%2210000%22%2C%22Status%22%3A2%2C%22Type%22%3A0%2C%22Name%22%3A%22Cash%22%2C%22AcctType%22%3A%22Cash%22%2C%22AllowPost%22%3Atrue%2C%22Description%22%3A%22%22%2C%22LastModTime%22%3A%222017-07-04T17%3A41%3A00Z%22%2C%22LastModBy%22%3A0%2C%22CreateTS%22%3A%222017-07-04T17%3A41%3A00Z%22%2C%22CreateBy%22%3A0%2C%22recid%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/account/1/1" "request" "a15"  "WebService--ERROR-Set_Incorrect_AllowPost2"

# Read it back and make sure that AllowPosts is false
echo "%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%7D" > request
dojsonPOST "http://localhost:8270/v1/account/1/1" "request" "a16"  "WebService--ERROR-VRFY-2"

# Read Deposit Methods
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/depmeth/1" "request" "a17"  "WebService--DepositMethods"

# Ensure that we cannot delete an account that is in use...
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22accountForm%22%2C%22LID%22%3A5%7D" > request
dojsonPOST "http://localhost:8270/v1/account/1/5" "request" "a18"  "WebService--ERROR-VRFY-3"


#------------------------------------------------------------------------------
#  TEST a19 - a20
#  Add a GLAccount of the same name to multiple businesses.  This test
#  was added as a result of a bug found in production where the Lockbox
#  account could not be created on PAC because it had been created in REX
#  and the db query was missing the BID, so it searched all Accounts rather
#  than just the accounts in the business being edited.
#
#  Scenario:
#		Add an account named "Lockbox" to REX and to PAC
#
#  Expected Results:
#	1.	There should be no error in adding the account to either business.
#       Even though the account is named "Lockbox" in both cases, it is not
#       a problem because they are in different businesses.
#------------------------------------------------------------------------------

# Add Lockbox account to business 1 (REX)
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22LID%22%3A0%2C%22PLID%22%3A9%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%2210102%22%2C%22Status%22%3A2%2C%22Name%22%3A%22FRB%2B92844%2BLockbox%22%2C%22AcctType%22%3A%22Cash%22%2C%22AllowPost%22%3Atrue%2C%22FLAGS%22%3A0%2C%22OffsetAccount%22%3A0%2C%22Description%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/account/1/0" "request" "a19"  "WebService--Add-LockboxAcct-toREX"

#--------------------------------------------------------------------------------------------
# Add Lockbox account to business 2 (PAC) with same account name and same account number
# This should be allowed because it is going to a different business
#--------------------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22LID%22%3A0%2C%22PLID%22%3A0%2C%22BID%22%3A2%2C%22BUD%22%3A%22PAC%22%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%2210102%22%2C%22Status%22%3A2%2C%22Name%22%3A%22FRB%2B92844%2BLockbox%22%2C%22AcctType%22%3A%22Cash%22%2C%22AllowPost%22%3Atrue%2C%22FLAGS%22%3A0%2C%22OffsetAccount%22%3A0%2C%22Description%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/account/2/0" "request" "a20"  "WebService--Add-LockboxAcct-toPAC"

#--------------------------------------------------------------------------------------------
# TEST a21
# Validate that an assessment cannot be made against a non-existent Rental Agreement.
#
# Scenario:
# 		Attempt to add an assessment to RAID 1000, where RAID 1000 does not exist
#
# Expected Results:
#		Bizlogic should catch the issue and return an appropriate error message.
#--------------------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A5%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1000%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%227%2F20%2F2017%22%2C%22Stop%22%3A%227%2F20%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A50%2C%22Rentable%22%3A%22309+S+Rexford%22%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22Validate+RAID+checking%22%2C%22ExpandPastInst%22%3A0%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22LastModByUser%22%3A%22%22%2C%22CreateByUser%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/2/0" "request" "a21"  "WebService--ValidateRAIDonAssessmentAdd"

#--------------------------------------------------------------------------------------------
# TEST a22
# Validate that an assessment cannot cross business boundaries
#
# Scenario:
# 		Attempt to add an assessment to RAID 4, where RAID 4 exists but is part of
#		a different business.
#
# Expected Results:
#		Bizlogic should catch the issue and return an appropriate error message.
#--------------------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A5%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%222%2F28%2F2018%22%2C%22Stop%22%3A%222%2F28%2F2018%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A50%2C%22Rentable%22%3A%22309+S+Rexford%22%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22ExpandPastInst%22%3A0%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22LastModByUser%22%3A%22%22%2C%22CreateByUser%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/2/0" "request" "a22"  "WebService--ValidateRAIDdoesNotCrossBizBoundary"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck
