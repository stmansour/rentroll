#!/bin/bash

TESTNAME="RAID Account Balance and Expenses"
TESTSUMMARY="Test rentroll RA Acct Balance and Expenses"

CREATENEWDB=0
BUD=OKC

source ../share/base.sh

echo "Create new database..."
pushd ../importers/onesite/okc;if [ ! -f iso.sql ]; then ./functest.sh ; fi; popd
mysql --no-defaults rentroll < ../importers/onesite/okc/iso.sql

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST 00
#  Add some receipts
#
#  Scenario:
#		A payment is made for each payment type
#
#  Expected Results:
#	1.	A receipt log report that shows all the receipts, the types, and
#       expected total.
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226333-123%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A2%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%221234%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A123%2C%22Comment%22%3A%22hello%22%2C%22OtherPayorName%22%3A%22Bill+Blatt%22%2C%22RentableName%22%3A%226333-123%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a00"  "Receipts-CreateReceipt-AMEX"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226309-033%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A5%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A5%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22887766%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A88.77%2C%22Comment%22%3A%22Blotto%22%2C%22OtherPayorName%22%3A%22Hugh+Jass%22%2C%22RentableName%22%3A%226309-033%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a01"  "Receipts-CreateReceipt-Discover"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%227004-200%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A9%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A9%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22132456789123%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A200%2C%22Comment%22%3A%22200%22%2C%22OtherPayorName%22%3A%22Mr.+Winky%22%2C%22RentableName%22%3A%227004-200%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a02"  "Receipts-CreateReceipt-MoneyOrder"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226313-054%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A8%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A8%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22876543098%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A543%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Hans+Moleman%22%2C%22RentableName%22%3A%226313-054%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a03"  "Receipts-CreateReceipt-VISA"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226305-018%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A1%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A1%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%224884883773%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A488%2C%22Comment%22%3A%22laugher%22%2C%22OtherPayorName%22%3A%22Sideshow+Bob%22%2C%22RentableName%22%3A%226305-018%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a04"  "Receipts-CreateReceipt-ACH"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226337-139%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A7%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A7%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22291716%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A39%2C%22Comment%22%3A%22Abner!%22%2C%22OtherPayorName%22%3A%22Edna+Kravitz%22%2C%22RentableName%22%3A%226337-139%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a05"  "Receipts-CreateReceipt-MasterCard"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226303-014%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A3%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A3%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22cash%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A1000%2C%22Comment%22%3A%22excellent%22%2C%22OtherPayorName%22%3A%22Montgomery+Burns%22%2C%22RentableName%22%3A%226303-014%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a06"  "Receipts-CreateReceipt-Cash"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226365-249%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A4%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A4%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22549%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A5.49%2C%22Comment%22%3A%22Brandine%22%2C%22OtherPayorName%22%3A%22Cletus+Yokel%22%2C%22RentableName%22%3A%226365-249%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a07"  "Receipts-CreateReceipt-Check"
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226600-001%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A6%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A6%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22000000001%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A1000%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Edward+Snowden%22%2C%22RentableName%22%3A%226600-001%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "a08"  "Receipts-CreateReceipt-EFT"

CSVLOADRANGE="-G OKC -g 1/1/18,2/1/18"
docsvtest "a10" "-L 13,OKC ${CSVLOADRANGE}" "ReceiptList"
docsvtest "a11" "-L 30,OKC,9 ${CSVLOADRANGE}" "Receipt"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
