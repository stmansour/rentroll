#!/bin/bash

TESTNAME="RR View Use Case Databases"
TESTSUMMARY="Generates separate databases for multiple use cases"

# CREATENEWDB=0

source ../share/base.sh

function dbcore() {
	docsvtest "a" "-b business.csv -L 3" "Business"
	docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
	docsvtest "c" "-ar ar.csv" "AccountRules"
	docsvtest "d" "-m depmeth.csv -L 23,${BUD}" "DepositMethods"
	docsvtest "e" "-d depository.csv -L 18,${BUD}" "Depositories"
	docsvtest "f" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
	docsvtest "t" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
	docsvtest "h" "-p people.csv  -L 7,${BUD}" "People"
}

function stop() {
	pkill rentroll;exit 1
}

#------------------------------------------------------------------------------
#  TEST 00
#  Simple Asmt/Rcpt -  Non recurring assessment, a receipt, apply payments.
#
#  Scenario:  
#		Assess $100 Electric Base Fee, receive a receipt for $250. Apply
#       the funds toward the $100 assessment.
#
#  Expected Results:
#	1.	$150 should carry forward to the next period.
#------------------------------------------------------------------------------
echo "STARTING RENTROLL SERVER"
startRentRollServer
dbcore
docsvtest "i" "-R rt1.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j" "-r r1.csv -L 6,${BUD}" "Rentables"
docsvtest "k" "-C ra1.csv -L 9,${BUD}" "RentalAgreements"

mysqldump --no-defaults rentroll >empty1.sql

# Create a non-recurring Assessment
echo "*** TEST 00 ***"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A11%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%2211%2F3%2F2017%22%2C%22Stop%22%3A%2211%2F3%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A100%2C%22Rentable%22%3A%22309+Rexford%22%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22ExpandPastInst%22%3A0%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "a00"  "WebService--CreateAssessment"

# Receive a Receipt of $250 
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%2211%2F3%2F2017%22%2C%22DocNo%22%3A%22234234234%22%2C%22Payor%22%3A%22Aaron%2BRead%2B(TCID%3A%2B1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A250%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "a01"  "WebService--ReceiveReceipt"

# Create a Deposit
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%2211%2F4%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A250%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "a02"  "WebService--CreateDeposit"

# Apply the payment
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A1%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%2211%2F3%2F2017%22%2C%22ASMID%22%3A1%2C%22ARID%22%3A11%2C%22Assessment%22%3A%22Electric%20Base%20Fee%22%2C%22Amount%22%3A100%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A100%2C%22Dt%22%3A%2211%2F4%2F2017%22%2C%22Allocate%22%3A100%2C%22Date_%22%3A%222017-11-03T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-11-04T07%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "a03"  "WebService--ApplyThePayment"

# Do a text version of the Journal and LedgerActivity to make sure all the
# funds are properly transferred
RRDATERANGE="-j 2017-11-01 -k 2017-12-01"
dorrtest "a04" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "a05" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"

echo "Saving test00.sql"
mysqldump --no-defaults rentroll >test00.sql


#------------------------------------------------------------------------------
#  TEST 01
#  Floating Deposit -  Test Receipt where RAID is required.
#
#  Scenario: 
#      In this example, Receipt.RAID will be non-zero.  
#      A $1000 floating deposit is made in October 2017, and $500 more is
#      added to the floating deposit in November. A rentroll report is made
#      for the month of November.
#
#  Expected Results:
#	1.	The Beginning Security Deposit should be $1000. The Ending Security
#      Deposit amount for November should be $1500.
#
#  Notes: 
#	1.	Since we're using floating deposits, there is no associated rentable.
#	2.	TBD: should we group entries by RAID in the "no rentable" section?
#       should we provide a totals row? 
#               If we do not, then there is nothing that shows the
#               accumulation of multiple floating deposits.
#------------------------------------------------------------------------------
echo "*** TEST 01 ***"
newDB
dbcore
docsvtest "i" "-R rt1.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j" "-r r1.csv -L 6,${BUD}" "Rentables"
docsvtest "lf" "-C raPlusFloat.csv -L 9,${BUD}" "RentalAgreements"

# Create the October Receipt
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A15%2C%22PMTID%22%3A2%2C%22RAID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%2210%2F2%2F2017%22%2C%22DocNo%22%3A%2212345%22%2C%22Payor%22%3A%22Kevin+Mills+(TCID%3A+8)%22%2C%22TCID%22%3A8%2C%22Amount%22%3A1000%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b00"  "WebService--AddFloatingDeposit1"

# Make the October bank deposit
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A2%2C%22DEPName%22%3A2%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%2211%2F5%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A1000%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "b01"  "WebService--CreateBankDeposit"


# Create the November Receipt
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A15%2C%22PMTID%22%3A2%2C%22RAID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%2211%2F2%2F2017%22%2C%22DocNo%22%3A%222345%22%2C%22Payor%22%3A%22Kevin+Mills+(TCID%3A+8)%22%2C%22TCID%22%3A8%2C%22Amount%22%3A500%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b02"  "WebService--AddFloatingDeposit2"

# Make the November bank deposit
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B2%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A2%2C%22DEPName%22%3A2%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%2211%2F4%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A500%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "b03"  "WebService--CreateBankDeposit"

# Do a text version of the Journal to make sure all the funds are properly transferred
RRDATERANGE="-j 2017-10-01 -k 2017-12-01"
dorrtest "b04" "${RRDATERANGE} -b ${BUD} -r 1" "JournalReport"

# Do a text version of the Rentroll report for November
RRDATERANGE="-j 2017-11-01 -k 2017-12-01"
dorrtest "b05" "${RRDATERANGE} -b ${BUD} -r 4" "RentrollReport"

mysqldump --no-defaults rentroll >rrFloatingDep.sql

#------------------------------------------------------------------------------
#  TEST 02
#  Reverse Floating Deposit
#  Scenario: This scenario uses the database created in
#      TEST 1 and simply reverses the first receipt.
#
#  Expected Results:
#	1.	The result should be that the $1000 is removed, but the $500 remains
#		on the books and should be seen in the Security Deposits totals.
#------------------------------------------------------------------------------
echo "*** TEST 02 ***"

# Make the reversal entry
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22receiptForm%22%2C%22RCPTID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/1" "request" "c01"  "WebService--ReverseDeposit"

# Rentroll report for November should not have a Beginning Balance on
# for Security Deposits
RRDATERANGE="-j 2017-11-01 -k 2017-12-01"
dorrtest "c02" "${RRDATERANGE} -b ${BUD} -r 4" "RentrollReport"
mysqldump --no-defaults rentroll > test02.sql

#------------------------------------------------------------------------------
#  TEST 03
#  Delete A Deposit
#
#  Scenario:
#		Here we simply delete the Deposit that was created in TEST 01.
#       This could happen when you plan to deposit a check into one account
#       but after creating the deposit in Roller you realize that you makde the
#       Deposit to the wrong Depository. So, you delete the Deposit you just
#       created, which means the $500 Receipt is now available to be included
#       in another deposit. Then you create a new deposit to for correct
#       Depository.
#
#  Expected Results:
#	1.	The $500 receipt is available to be deposited after deleting the
#		first deposit.
#	2.	We can successfully create a second deposit.
#------------------------------------------------------------------------------
echo "*** TEST 03 ***"
mysql --no-defaults rentroll <rrFloatingDep.sql

# Delete the deposit from the wrong depository
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22depositFormBtns%22%2C%22DID%22%3A2%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/2" "request" "c00"  "WebService--DeleteDeposit"

# Create a new Deposit with the Receipt from the deleted deposit
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B2%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%2211%2F6%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A500%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "c01"  "WebService--CreateDeposit"
mysqldump --no-defaults rentroll > test03.sql

#------------------------------------------------------------------------------
#  TEST 04
#  Transfer funds from one account to another using an Expense
#
#  Scenario:
#		A payort provides $1025 payment that covers $1000 floating deposit
#		and $25 Application Fee. The check is deposited into the operating
#		account. The $1000 must be deposited in a separate bank account, the
#		security deposits account, using out-of-band capabilities from the
#		bank. We must reflect the change in Roller. We transfer the $1000
#		from the operating account to the security deposit account using the
#		Expense capability.
#
#  Expected Results:
#	1.	We see the $25 in the the operating account and $1000 in the
#		security deposit account.
#------------------------------------------------------------------------------
echo "*** TEST 04 ***"
mysql --no-defaults rentroll <empty1.sql

# $1000 Deposit
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A28%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%2211%2F7%2F2017%22%2C%22Stop%22%3A%2211%2F7%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A1000%2C%22Rentable%22%3A%22309+Rexford%22%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22ExpandPastInst%22%3A0%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "d00"  "WebService--CreateAssessment-1000"

# $25 Application Fee
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A1%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%2211%2F7%2F2017%22%2C%22Stop%22%3A%2211%2F7%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A25%2C%22Rentable%22%3A%22309+Rexford%22%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22ExpandPastInst%22%3A0%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "d01"  "WebService--CreateAssessment-25"

# $1025 Receipt for both
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A2%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%2211%2F7%2F2017%22%2C%22DocNo%22%3A%2223423%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A1025%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "d02"  "WebService--ReceiveReceipt"

# Deposit to Operating account
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%2211%2F7%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A1025%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "d03"  "WebService--CreateDeposit"

# Allocate Funds
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A1%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%2211%2F7%2F2017%22%2C%22ASMID%22%3A1%2C%22ARID%22%3A28%2C%22Assessment%22%3A%22Security%20Deposit%20Assessment%22%2C%22Amount%22%3A1000%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A1000%2C%22Dt%22%3A%2211%2F7%2F2017%22%2C%22Allocate%22%3A1000%2C%22Date_%22%3A%222017-11-07T08%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-11-07T08%3A00%3A00.000Z%22%7D%2C%7B%22recid%22%3A1%2C%22Date%22%3A%2211%2F7%2F2017%22%2C%22ASMID%22%3A2%2C%22ARID%22%3A1%2C%22Assessment%22%3A%22Application%20Fee%22%2C%22Amount%22%3A25%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A25%2C%22Dt%22%3A%2211%2F7%2F2017%22%2C%22Allocate%22%3A25%2C%22Date_%22%3A%222017-11-07T08%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-11-07T08%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1" "request" "d04"  "WebService--allocatefunds"

# $1000 moved from Operating Account to SecurityDeposit Account
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22expenseForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22EXPID%22%3A0%2C%22ARID%22%3A38%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Dt%22%3A%2211%2F7%2F2017%22%2C%22Amount%22%3A1000%2C%22AcctRule%22%3A%22%22%2C%22RName%22%3A%22309+Rexford%22%2C%22Comment%22%3A%22bank+transfer+SEC+DEP+to+proper+account%22%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22PREXPID%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/expense/1/0" "request" "d04"  "WebService--EXPENSE-Transfer"

mysqldump --no-defaults rentroll > test04.sql

#------------------------------------------------------------------------------
#  TEST 05
#  Rentable Type Change.
#
#  Scenario:
#		If the MarketRate for a Rentable Type changes during
#     	the period of a RentRoll view or report, the PeriodGSR must reflect the
#     	change. This RentRoll view/report is for November 2017. The GSR is $3500
#     	per month from Jan 1 to Nov 15, then it changes to $4000. So, in Nov,
#		we have 14 days at $3500 and 16 days at $4000.
#
#  Expected Result:
#	1.	The prorated GSR works out to $3766.67 for the period 11/1/2017 to 
#		12/1/2017.
#------------------------------------------------------------------------------
createDB
dbcore
docsvtest "ii" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "jj" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "kk" "-C ra.csv -L 9,${BUD}" "RentalAgreements"

# Add a market rate change in the middle of November
echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B1%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RMRID%22%3A1%2C%22RTID%22%3A1%2C%22MarketRate%22%3A3500%2C%22DtStart%22%3A%221%2F1%2F2014%22%2C%22DtStop%22%3A%2211%2F15%2F2017%22%2C%22w2ui%22%3A%7B%7D%7D%2C%7B%22recid%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RTID%22%3A1%2C%22RMRID%22%3A0%2C%22MarketRate%22%3A4000%2C%22DtStart%22%3A%2211%2F15%2F2017%22%2C%22DtStop%22%3A%2212%2F31%2F9999%22%2C%22w2ui%22%3A%7B%7D%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "e00"  "WebService--ChangeGSR"

mysqldump --no-defaults rentroll > test05.sql

#------------------------------------------------------------------------------
#  TEST 06
#  Scenario: Assessment is made in October. The assessment is paid during the
#  		November. The report range is set to the month of November.
#
#  Expected Result:
#	1.	We should see the payment as a row in the RentRoll view/report.
#	2.	We should see the Assessment amount due reflected in the Beginning
#       Receivables subtotal line
#------------------------------------------------------------------------------
echo "*** TEST 06 ***"
newDB
dbcore
docsvtest "i" "-R rt1.csv -L 5,${BUD}" "RentableTypes"
docsvtest "j" "-r r1.csv -L 6,${BUD}" "Rentables"
docsvtest "lf" "-C raPlusFloat.csv -L 9,${BUD}" "RentalAgreements"

# Create a non-recurring Assessment in October
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A7%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%2210%2F12%2F2017%22%2C%22Stop%22%3A%2210%2F12%2F2017%22%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22TCID%22%3A0%2C%22Amount%22%3A50%2C%22Rentable%22%3A%22309+Rexford%22%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22ExpandPastInst%22%3A0%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "f00"  "WebService--CreateAssessment"

# Create the November Receipt
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A25%2C%22PMTID%22%3A2%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%2211%2F15%2F2017%22%2C%22DocNo%22%3A%221234%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A50%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "f01"  "WebService--ReceiveReceipt"

# Apply the payment
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A1%2C%22BID%22%3A1%2C%22records%22%3A%5B%7B%22recid%22%3A0%2C%22Date%22%3A%2210%2F12%2F2017%22%2C%22ASMID%22%3A1%2C%22ARID%22%3A7%2C%22Assessment%22%3A%22Broken%20Window%20charge%22%2C%22Amount%22%3A50%2C%22AmountPaid%22%3A0%2C%22AmountOwed%22%3A50%2C%22Dt%22%3A%2211%2F15%2F2017%22%2C%22Allocate%22%3A50%2C%22Date_%22%3A%222017-10-12T07%3A00%3A00.000Z%22%2C%22Dt_%22%3A%222017-11-15T08%3A00%3A00.000Z%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/allocfunds/1/" "request" "f02"  "WebService--ApplyThePayment"

RRDATERANGE="-j 2017-11-01 -k 2017-12-01"
dorrtest "f03" "${RRDATERANGE} -b ${BUD} -r 4" "Rentroll"


mysqldump --no-defaults rentroll > test06.sql



stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
