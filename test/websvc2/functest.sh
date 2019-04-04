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
mysql --no-defaults rentroll < ../ws/wsdb.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

# Add 3 new receipts

# Add $100 by payor Aaron Read on 8/23/2017
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221001%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A100%2C%22Comment%22%3A%22I+am+check+1001%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b00"  "WebService--Add_Receipt_1"

# Add $200 by payor Kirsten Read on 8/23/2017
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221002%22%2C%22Payor%22%3A%22Kirsten+Read+(TCID%3A+2)%22%2C%22TCID%22%3A2%2C%22Amount%22%3A200%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b01"  "WebService--Add_Receipt_2"

# Add $300 by payor Alex Vahabzadeh on 8/23/2017
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DID%22%3A0%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221003%22%2C%22Payor%22%3A%22Alex+Vahabzadeh+(TCID%3A+11)%22%2C%22TCID%22%3A11%2C%22Amount%22%3A300%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/0" "request" "b02"  "WebService--Add_Receipt_3"

# Create a deposit from the 3 receipts just added
# Will create DID (deposit ID) 1 for BID 1 (REX)
# Deposit should total $600
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%2C36%5D%2C%22record%22%3A%7B%22recid%22%3A0%2C%22check%22%3A0%2C%22DID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22FLAGS%22%3A0%2C%22Amount%22%3A600%2C%22ClearedAmount%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/0" "request" "b03"  "WebService--Create_Deposit"

# Remove the $300 receipt from this deposit
# This will set the DID for RCPTID 36 to 0.
# Deposit should be $300 now
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%5D%2C%22record%22%3A%7B%22recid%22%3A1%2C%22DID%22%3A13%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22Amount%22%3A300%2C%22ClearedAmount%22%3A0%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/1" "request" "b04"  "WebService--Remove_a_receipt_from_Deposit_1"

# Now add it back into the deposit
# Deposit total is back to $600
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%2C36%5D%2C%22record%22%3A%7B%22recid%22%3A1%2C%22DID%22%3A13%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22Amount%22%3A600%2C%22ClearedAmount%22%3A0%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/deposit/1/1" "request" "b05"  "WebService--Add_the_receipt_back_to_Deposit_1"

# Now modify the amount in the receipt. Change the amount of RCPTID 3 from $300
# to $250.  This should cause the Deposit to reflect the new Receipt's RCPTID = 5,
# and 4, the reversal.
# Deposit total is now $550
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22roller%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A36%2C%22PRCPTID%22%3A0%2C%22BID%22%3A1%2C%22DID%22%3A13%2C%22BUD%22%3A%22REX%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A2%2C%22Payor%22%3A%22Alex%2BVahabzadeh%2B(TCID%3A%2B11)%22%2C%22TCID%22%3A11%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22DocNo%22%3A%221003%22%2C%22Amount%22%3A250%2C%22ARID%22%3A3%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22LastModByUser%22%3A%22UID-0%22%2C%22CreateByUser%22%3A%22UID-0%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%22%22%2C%22PmtTypeName%22%3A2%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/3" "request" "b06"  "WebService--Handle_ReverseCorrections"

# Now Reverse receipt 2. This will add the reversal to the Deposit.
# This should reduce the deposit amount to $350.
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22receiptForm%22%2C%22RCPTID%22%3A35%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/1/2" "request" "b07"  "WebService--Handle_Reverse"

# update the manual tally for the deposit to $350
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22depositForm%22%2C%22Receipts%22%3A%5B34%2C35%2C36%2C37%2C38%2C39%5D%2C%22record%22%3A%7B%22recid%22%3A1%2C%22DID%22%3A13%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22DEPID%22%3A1%2C%22DEPName%22%3A1%2C%22DPMID%22%3A1%2C%22DPMName%22%3A1%2C%22Dt%22%3A%228%2F23%2F2017%22%2C%22Amount%22%3A350%2C%22ClearedAmount%22%3A0%2C%22FLAGS%22%3A0%7D%7D" >request
dojsonPOST "http://localhost:8270/v1/deposit/1/1" "request" "b08"  "WebService--Handle_Reverse"



stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck
