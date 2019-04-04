#!/bin/bash

TESTNAME="Server Functions for Receipts Client"
TESTSUMMARY="Tests receipts and printouts for Receipts client"

CREATENEWDB=0
BUD=OKC

source ../share/base.sh

echo "Create new database..."
mysql --no-defaults rentroll < iso.sql

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  Add some receipts
#
#  Scenario:
#		A payment is made for each payment type
#
#  Expected Results:
#	1.	A receipt log report that shows all the receipts, the types, and
#       expected total.
#   2.  An individual receipt for RCPT-9
#   3.  A receipt log with the total amount collected
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
docsvtest "a12" "-L 31,OKC,9 ${CSVLOADRANGE}" "HotelReceipt"

#------------------------------------------------------------------------------
#  TEST b
#  Reverse a receipt
#
#  Scenario:
#		Add a payment. Then reverse it
#
#  Expected Results:
#	1.	Ensure that the reversal function works
#   2.  Ensure that the total in the receipt list log is correct after reversal
#------------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226325-100%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A2%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A2%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456789%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A10000%2C%22Comment%22%3A%22Reverse+Me!%22%2C%22OtherPayorName%22%3A%22Sven+Fjordsen%22%2C%22RentableName%22%3A%226325-100%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "b00"  "Receipts-ReceiptToReverse"
echo "%7B%22cmd%22%3A%22delete%22%2C%22formname%22%3A%22receiptForm%22%2C%22RCPTID%22%3A10%2C%22client%22%3A%22receipts%22%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/10" "request" "b01"  "Receipts-ReceiptToReverse"
CSVLOADRANGE="-G OKC -g 1/1/18,2/1/18"
docsvtest "b02" "-L 13,OKC ${CSVLOADRANGE}" "ReceiptList"
docsvtest "b03" "-L 30,OKC,10 ${CSVLOADRANGE}" "Receipt"
docsvtest "b04" "-L 30,OKC,11 ${CSVLOADRANGE}" "Receipt"



#------------------------------------------------------------------------------
#  TEST c
#  Validate reversals occur and are correct for all circumstances.  Also
#  validate that updates work correctly when a reversal is not required.
#
#  Scenario:
#		Create a new receipt that we will use to edit each field, one at
#       a time and validate that the updates are correct. Some updates
#       will result in a reversal.
#
#  Expected Results:
#	1.	Issuing the reverse command simply reverses a recei
#       expected total.
#   2.  An individual receipt for RCPT-9
#------------------------------------------------------------------------------
#--------------------------------------------------
# Add a $500 receipt from Tutti McTutu on 1/18/2018
#--------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226301-002%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A0%2C%22PMTID%22%3A5%2C%22RAID%22%3A0%2C%22PmtTypeName%22%3A5%2C%22BID%22%3A2%2C%22BUD%22%3A%22OKC%22%2C%22DID%22%3A0%2C%22Dt%22%3A%221%2F18%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A500%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Tutti+McTutu%22%2C%22RentableName%22%3A%226301-002%22%2C%22FLAGS%22%3A0%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/0" "request" "c00"  "Receipts-Add_a_Receipt"

#-------------------------------------------------------------
# Change the date to 2017. This should cause a reversal.
#-------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226301-002%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A12%2C%22PRCPTID%22%3A0%2C%22BID%22%3A2%2C%22DID%22%3A0%2C%22BUD%22%3A%22OKC%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A5%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Amount%22%3A500%2C%22ARID%22%3A0%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Tutti+McTutu%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%226301-002%22%2C%22PmtTypeName%22%3A5%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/12" "request" "c01"  "Receipts-Change_the_date"

#------------------------------------------------------------------------
# Change the Payment Type to VISA, this should not cause a reversal
#------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226301-002%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A14%2C%22PRCPTID%22%3A0%2C%22BID%22%3A2%2C%22DID%22%3A0%2C%22BUD%22%3A%22OKC%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A8%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Amount%22%3A500%2C%22ARID%22%3A0%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Tutti+McTutu%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%226301-002%22%2C%22PmtTypeName%22%3A8%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/14" "request" "c02"  "Receipts-Change_the_payment_type"
#------------------------------------------------------------------------
# Change the amount to $50, this should cause a reversal
#------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226301-002%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A14%2C%22PRCPTID%22%3A0%2C%22BID%22%3A2%2C%22DID%22%3A0%2C%22BUD%22%3A%22OKC%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A8%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Amount%22%3A50%2C%22ARID%22%3A0%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Tutti+McTutu%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%226301-002%22%2C%22PmtTypeName%22%3A8%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/14" "request" "c03"  "Receipts-Change_the_amount"

#------------------------------------------------------------------------
# Change the address, this should cause a reversal
#------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226315-055%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A16%2C%22PRCPTID%22%3A0%2C%22BID%22%3A2%2C%22DID%22%3A0%2C%22BUD%22%3A%22OKC%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A8%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Amount%22%3A50%2C%22ARID%22%3A0%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Tutti+McTutu%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%226315-055%22%2C%22PmtTypeName%22%3A8%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/16" "request" "c04"  "Receipts-Change_the_address"

#------------------------------------------------------------------------
# Change the Payor, this should cause a reversal
#------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226315-055%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A18%2C%22PRCPTID%22%3A0%2C%22BID%22%3A2%2C%22DID%22%3A0%2C%22BUD%22%3A%22OKC%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A8%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Amount%22%3A50%2C%22ARID%22%3A0%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22Tutti+de+Taco%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%226315-055%22%2C%22PmtTypeName%22%3A8%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/18" "request" "c05"  "Receipts-Change_the_payor"

#------------------------------------------------------------------------
# Change the comment, this should NOT cause a reversal
#------------------------------------------------------------------------
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22client%22%3A%22receipts%22%2C%22RentableName%22%3A%226315-055%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A20%2C%22PRCPTID%22%3A0%2C%22BID%22%3A2%2C%22DID%22%3A0%2C%22BUD%22%3A%22OKC%22%2C%22RAID%22%3A0%2C%22PMTID%22%3A8%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Dt%22%3A%221%2F17%2F2018%22%2C%22DocNo%22%3A%22123456782%22%2C%22Amount%22%3A50%2C%22ARID%22%3A0%2C%22Comment%22%3A%22hooray!%22%2C%22OtherPayorName%22%3A%22Tutti+de+Taco%22%2C%22FLAGS%22%3A0%2C%22RentableName%22%3A%226315-055%22%2C%22PmtTypeName%22%3A8%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/receipt/2/20" "request" "c06"  "Receipts-Change_the_comment"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
