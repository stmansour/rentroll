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

#------------------------------------------------------------------------------
#  TEST 1
#  Floating Deposit -  Receipt where RAID is required. In this
#      example, Receipt.RAID will be non-zero.  In this scenario
#      a $1000 floating deposit is made in October 2017, and $500
#      more is added to the floating deposit in November. The Ending
#      Security Deposit amount for November should be $1500
#------------------------------------------------------------------------------
echo "STARTING RENTROLL SERVER"
startRentRollServer
dbcore
docsvtest "g1" "-R rt1.csv -L 5,${BUD}" "RentableTypes"
docsvtest "i1" "-r r1.csv -L 6,${BUD}" "Rentables"
docsvtest "j1" "-C ra1.csv -L 9,${BUD}" "RentalAgreements"

# Create the Receipt
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A15%2C%22PMTID%22%3A4%2C%22RAID%22%3A1%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%2210%2F2%2F2017%22%2C%22DocNo%22%3A%22234234234%22%2C%22Payor%22%3A%22Aaron%2BRead%2B(TCID%3A%2B1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A1000%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "k1"  "WebService--AddFloatingDeposit"

# Make the bank deposit
# echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B1%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%2210%2F3%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A1000%2C%22ClearedAmount%22%3A0%7D%7D" > request
# dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "l1"  "WebService--CreateBankDeposit"

mysqldump --no-defaults rentroll >rrFloatingDep.sql

#----------------------------------------------------
#  TEST 2
#  Rentable Type Change during vacancy.
#----------------------------------------------------
#createDB
#dbcore


stopRentRollServer
echo "RENTROLL SERVER STOPPED"
logcheck

exit 0
