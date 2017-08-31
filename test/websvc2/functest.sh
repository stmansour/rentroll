#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Deposits"
TESTSUMMARY="Test Desposits web services"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < ../ws/restore.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
startRentRollServer

# echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22DID%22%2C%22direction%22%3A%22asc%22%7D%5D%2C%22searchDtStart%22%3A%221%2F1%2F2016%22%2C%22searchDtStop%22%3A%225%2F1%2F2016%22%7D" > request
# dojsonPOST "http://localhost:8270/v1/deposit/1" "request" "a0"  "WebService--Read_Deposits"

# echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22DID%22%2C%22direction%22%3A%22asc%22%7D%5D%2C%22searchDtStart%22%3A%221%2F1%2F2016%22%2C%22searchDtStop%22%3A%225%2F1%2F2016%22%7D" > request
# dojsonPOST "http://localhost:8270/v1/deposit/1/3" "request" "a1"  "WebService--Read_Deposit_3"

# # Read the list of receipts associated with a deposit
# echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%5D%2C%22searchDtStart%22%3A%221%2F1%2F2016%22%2C%22searchDtStop%22%3A%225%2F1%2F2016%22%7D" > request
# dojsonPOST "http://localhost:8270/v1/depositlist/1/2" "request" "a2"  "WebService--Read_Receipts_for_deposit_2"

# Add 3 new receipts
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221001%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A100%2C%22Comment%22%3A%22I+am+check+1001%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b00"  "WebService--Add_Receipt_1"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221002%22%2C%22Payor%22%3A%22Kirsten+Read+(TCID%3A+2)%22%2C%22TCID%22%3A2%2C%22Amount%22%3A200%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b01"  "WebService--Add_Receipt_2"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221003%22%2C%22Payor%22%3A%22Alex+Vahabzadeh+(TCID%3A+11)%22%2C%22TCID%22%3A11%2C%22Amount%22%3A300%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b02"  "WebService--Add_Receipt_3"

# Create a deposit from the 3 receipts just added
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%2C36%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A600%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "b03"  "WebService--Create_Deposit"

# Remove the $300 receipt from this deposit
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%5D%2C%22record%22%3A%7B%22recid%22%3A13%2C%22DID%22%3A13%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22Amount%22%3A300%2C%22ClearedAmount%22%3A0%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/13" "request" "b04"  "WebService--Remove_a_receipt_from_Deposit_13"

# Now add it back
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%2C36%5D%2C%22record%22%3A%7B%22recid%22%3A13%2C%22DID%22%3A13%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22Amount%22%3A600%2C%22ClearedAmount%22%3A0%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/13" "request" "b05"  "WebService--Add_the_receipt_back_to_Deposit_13"

# Now modify the receipt so that it gets auto-reversed and a new receipt is added
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A36%2C%22PRCPTID%22%3A0%2C%22BID%22%3A1%2C%22DID%22%3A13%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A2%2C%22Payor%22%3A%22Alex+Vahabzadeh+(TCID%3A+11)%22%2C%22TCID%22%3A11%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221003%22%2C%22Amount%22%3A250%2C%22ARID%22%3A3%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%2C%22PmtTypeName%22%3A2%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/36" "request" "b06"  "WebService--Handle_autoReverseCorrections"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck
