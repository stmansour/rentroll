#!/bin/bash
TESTNAME="JM1 - 12 months of books on the Rexford Properties"
TESTSUMMARY="Setup and run JM1 company and the Rexford Properties for 1 year"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"

source ../share/base.sh

echo "BEGIN JM1 FUNCTIONAL TEST" >>${LOGFILE}
#========================================================================================
# INITIALIZE THE BUSINESS
#   This section has the 1-time tasks to set up the business and get the accounts to
#   their correct starting values.
#========================================================================================
docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "c" "-m depmeth.csv -L 23,${BUD}" "DepositMethods"
docsvtest "d" "-d depository.csv -L 18,${BUD}" "Depositories"
docsvtest "e" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "f" "-l strlists.csv -L 25,${BUD}" "StringLists"
docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "ga" "-V vehicle.csv -L 28,${BUD}" "Vehicles"
docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "i" "-u custom.csv -L 14,${BUD}" "CustomAttributes"
docsvtest "j" "-U assigncustom.csv -L 15,${BUD}" "AssignCustomAttributes"
docsvtest "k" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
docsvtest "l" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"

# get the deposits on the books
docsvtest "n" "-A asm2015-12.csv -G ${BUD} -g 12/1/15,1/1/16 -L 11,${BUD}" "Assessments-2015-DEC"
docsvtest "o" "-e rcpt2015-12.csv -G ${BUD} -g 12/1/15,1/1/16 -L 13,${BUD}" "Receipts-2015-DEC"

# validate GSR
dorrtest "p" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 11" "GSR"

#  INITIALIZE database with deposit information and verify Accounts
dorrtest "q" "-j 2015-12-01 -k 2016-01-01 -x -b ${BUD}" "Process-2015-DEC"
dorrtest "r" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 1" "Journal"
dorrtest "s" "-j 2015-12-01 -k 2016-01-01 -b ${BUD} -r 2" "Ledgers"
dorrtest "t" "-r 12,1,RA001,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA01"
dorrtest "u" "-r 12,7,RA001,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA01"
dorrtest "v" "-r 12,1,RA002,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA02"
dorrtest "x" "-r 12,7,RA002,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA02"
dorrtest "y" "-r 12,1,RA003,2016-01-01 -b ${BUD}" "AccountBalance-GeneralAccountsReceivable-RA-03"
dorrtest "z" "-r 12,7,RA003,2016-01-01 -b ${BUD}" "AccountBalance-SecurityDeposits-RA-03"

#========================================================================================
# JANUARY 2016
#    Normal month
#========================================================================================
RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
CSVLOADRANGE="-G ${BUD} -g 1/1/16,2/1/16"
# 1.  Generate recurring assessment instances  -  Note: will be done by server automatically  (18 processes journal only)
dorrtest "a1" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-JAN"

# 2.  Load new assessments for this period.  For this test, we start the rent assessments now.
docsvtest "b1" "-A asm2016-01.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-JAN"

# 3.  Create Invoices for each tenant
docsvtest "c1" "-i invoice-2016-01-Read.csv ${CSVLOADRANGE} -L 20,REX" "Invoice-2016Jan-Read"
docsvtest "d1" "-i invoice-2016-01-Costea.csv ${CSVLOADRANGE} -L 20,REX" "invoice-2016Jan-Costea"
docsvtest "e1" "-i invoice-2016-01-Haroutunian.csv ${CSVLOADRANGE} -L 20,REX" "invoice-2016Jan-Haroutunian"
dorrtest "f1" "${RRDATERANGE} -b ${BUD} -r 9,IN001" "InvoiceReport-2016Jan-Read"
dorrtest "g1" "${RRDATERANGE} -b ${BUD} -r 9,2" "InvoiceReport-2016Jan-Costea"
dorrtest "h1" "${RRDATERANGE} -b ${BUD} -r 9,3" "InvoiceReport-2016Jan-Haroutunian"

# 4. Enter any receipts (and assessments if any) since Jan1 - end of the month
docsvtest "i1" "-e rcpt2016-01.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-JAN"

# 5. Create deposits for all receipts
docsvtest "j1" "-y deposit-2016-01.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-JAN"

# 6. Process anything that was just added
dorrtest "k3" "${RRDATERANGE} -b ${BUD}" "Finish-2016-JAN"

# 7. Generate final reports for the month
dorrtest "l1" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "m1" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest "n1" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest "o1" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest "p1" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest "q1" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# FEBRUARY 2016
#    Haroutunian moves out on Feb 8
#========================================================================================
RRDATERANGE="-j 2016-02-01 -k 2016-03-01"
CSVLOADRANGE="-G ${BUD} -g 2/1/16,3/1/16"
dorrtest  "a2" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-FEB"
docsvtest "b2" "-A asm2016-02.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-FEB"
docsvtest "c2" "-i invoice-2016-02-Read.csv ${CSVLOADRANGE} -L 20,REX" "Invoice-2016-02-Read"
docsvtest "d2" "-i invoice-2016-02-Costea.csv ${CSVLOADRANGE} -L 20,REX" "invoice-2016-02-Costea"
docsvtest "e2" "-i invoice-2016-02-Haroutunian.csv ${CSVLOADRANGE} -L 20,REX" "invoice-2016-02-Haroutunian"
dorrtest  "f2" "${RRDATERANGE} -b ${BUD} -r 9,IN001" "InvoiceReport-2016-02-Read"
dorrtest  "g2" "${RRDATERANGE} -b ${BUD} -r 9,2" "InvoiceReport-2016-02-Costea"
dorrtest  "h2" "${RRDATERANGE} -b ${BUD} -r 9,3" "InvoiceReport-2016-02-Haroutunian"
docsvtest "i2" "-e rcpt2016-02.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-FEB"
docsvtest "j2" "-y deposit-2016-02.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-FEB"
dorrtest  "k2" "${RRDATERANGE} -b ${BUD}" "Finish-2016-FEB"
dorrtest  "l2" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m2" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n2" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o2" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p2" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q2" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# MARCH 2016
#    GSR and Contract rent change to 3750 for 309 Rexford
#    Haroutunian receives 865.29 Deposit return, forfeits the rest
#========================================================================================

##-----------------------------------------------------
##  1. Update end date on RentalAgreement 1 to 3/1/18
##  2. Update ContractRent to $3750/month
##  3. Update MarketRate to $3750/month
##-----------------------------------------------------
cat >xxyyzz <<EOF
use rentroll
update RentalAgreement SET AgreementStop="2018-03-01",PossessionStop="2018-03-01",RentStop="2018-03-01" WHERE RAID=1;
INSERT INTO RentalAgreementRentables (RAID,BID,RID,CLID,ContractRent,RARDtStart,RARDtStop) VALUES(1,1,1,0,3750,"2016-03-01 00:00:00","2018-03-01 00:00:00");
INSERT INTO RentableMarketRate (RTID,BID,MarketRate,DtStart,DtStop) VALUES(1,1,3750,"2016-03-01 00:00:00","2018-03-01 00:00:00");
INSERT INTO RentalAgreementPayors (RAID,BID,TCID,DtStart,DtStop) VALUES(1,1,1,"2016-03-01 00:00:00","2018-03-01 00:00:00");
INSERT INTO RentalAgreementPayors (RAID,BID,TCID,DtStart,DtStop) VALUES(1,1,2,"2016-03-01 00:00:00","2018-03-01 00:00:00");
EOF
${MYSQL} --no-defaults <xxyyzz
rm -f xxyyzz

RRDATERANGE="-j 2016-03-01 -k 2016-04-01"
CSVLOADRANGE="-G ${BUD} -g 3/1/16,4/1/16"
docsvtest "b3" "-A asm2016-03.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Mar"
dorrtest  "a3" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-Mar"
docsvtest "i3" "-e rcpt2016-03.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Mar"
docsvtest "j3" "-y deposit-2016-03.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Mar"
dorrtest  "k3" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Mar"
dorrtest  "l3" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m3" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n3" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o3" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p3" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q3" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# APRIL 2016
#    GSR and Contract rent change to 4150 for 311 Rexford
#========================================================================================

##----------------------------------------------------------
##  1. Update MarketRate for RentableType 3 to $4150/month
##----------------------------------------------------------
cat >xxyyzz <<EOF
use rentroll
INSERT INTO RentableMarketRate (RTID,MarketRate,DtStart,DtStop) VALUES(3,4150,"2016-04-01 00:00:00","2018-04-01 00:00:00");
EOF
${MYSQL} --no-defaults <xxyyzz
rm -f xxyyzz
dorrtest  "z3" "-j 2016-01-01 -k 2016-06-01 -b ${BUD} -r 20,R003" "MarketRateValidation"

##----------------------------------------------------------
##  2. Process the rent checks and generate reports
##----------------------------------------------------------
RRDATERANGE="-j 2016-04-01 -k 2016-05-01"
CSVLOADRANGE="-G ${BUD} -g 4/1/16,5/1/16"
# docsvtest "b4" "-A asm2016Apr.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Apr"  		## no new assessments this month
dorrtest  "a4" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-Apr"
docsvtest "i4" "-e rcpt2016-04.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Apr"
docsvtest "j4" "-y deposit-2016-04.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Apr"
dorrtest  "k4" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Apr"
dorrtest  "l4" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m4" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n4" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o4" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p4" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q4" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"
#========================================================================================
# MAY 2016
#    GSR and Contract rent change to 3800 for 309.5 Rexford
#========================================================================================

##----------------------------------------------------------
##  1. Update MarketRate for RentableType 2 to $3800/month
##     Update ContractRent for Rentable 2 to $3800/month
##----------------------------------------------------------
cat >xxyyzz <<EOF
use rentroll
INSERT INTO RentableMarketRate (RTID,BID,MarketRate,DtStart,DtStop) VALUES(2,1,3800,"2016-05-01 00:00:00","2016-10-01 00:00:00");
INSERT INTO RentalAgreementRentables (RAID,BID,RID,CLID,ContractRent,RARDtStart,RARDtStop) VALUES(2,1,2,0,3800,"2016-05-01 00:00:00","2016-08-28 00:00:00");
UPDATE RentalAgreementRentables SET RARDtStop="2016-05-01" WHERE ContractRent=3550 AND RID=2;
EOF
${MYSQL} --no-defaults <xxyyzz
rm -f xxyyzz
dorrtest  "z4" "-j 2016-01-01 -k 2016-09-01 -b ${BUD} -r 20,R002" "MarketRateValidation"

##----------------------------------------------------------
##  2. Process the rent checks and generate reports
##----------------------------------------------------------
RRDATERANGE="-j 2016-05-01 -k 2016-06-01"
CSVLOADRANGE="-G ${BUD} -g 5/1/16,6/1/16"
docsvtest "b5" "-A asm2016-05.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-May"
dorrtest  "a5" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-May"
docsvtest "i5" "-e rcpt2016-05.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-May"
docsvtest "j5" "-y deposit-2016-05.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-May"
dorrtest  "k5" "${RRDATERANGE} -b ${BUD}" "Finish-2016-May"
dorrtest  "l5" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m5" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n5" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o5" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p5" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q5" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# JUNE 2016
#========================================================================================
RRDATERANGE="-j 2016-06-01 -k 2016-07-01"
CSVLOADRANGE="-G ${BUD} -g 6/1/16,7/1/16"
# docsvtest "b6" "-A asm2016-06.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Jun"  		## no new assessments this month
dorrtest  "a6" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-Jun"
docsvtest "i6" "-e rcpt2016-06.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Jun"
docsvtest "j6" "-y deposit-2016-06.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Jun"
dorrtest  "k6" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Jun"
dorrtest  "l6" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m6" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n6" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o6" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p6" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q6" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# JULY 2016
#    Add assessments for rent and security deposit for 311 Rexford
#========================================================================================
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"
CSVLOADRANGE="-G ${BUD} -g 7/1/16,8/1/16"
docsvtest "b7" "-A asm2016-07.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Jul"
dorrtest  "a7" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-Jul"
docsvtest "i7" "-e rcpt2016-07.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Jul"
docsvtest "j7" "-y deposit-2016-07.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Jul"
dorrtest  "k7" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Jul"
dorrtest  "l7" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m7" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n7" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o7" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p7" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q7" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# AUGUST 2016
#========================================================================================
RRDATERANGE="-j 2016-08-01 -k 2016-09-01"
CSVLOADRANGE="-G ${BUD} -g 8/1/16,9/1/16"
# docsvtest "b8" "-A asm2016-08.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Aug"  ## no new assessments
dorrtest  "a8" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-2016-Aug"
docsvtest "i8" "-e rcpt2016-08.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Aug"
docsvtest "j8" "-y deposit-2016-08.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Aug"
dorrtest  "k8" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Aug"
dorrtest  "l8" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m8" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n8" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o8" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p8" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q8" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# SEPTEMBER 2016
#========================================================================================
RRDATERANGE="-j 2016-09-01 -k 2016-10-01"
CSVLOADRANGE="-G ${BUD} -g 9/1/16,10/1/16"
dorrtest  "a9" "${RRDATERANGE} -b ${BUD}" "Process-2016-Sep"
docsvtest "i9" "-e rcpt2016-09.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Sep"
docsvtest "j9" "-y deposit-2016-09.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Sep"
dorrtest  "k9" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Sep"
dorrtest  "l9" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m9" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n9" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o9" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p9" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q9" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

##----------------------------------------------------------
##  1. Update MarketRate for RentableType 2 to $4000/month
##  2. Set contract rent for Rentable 2 to $4000/month
##----------------------------------------------------------
cat >xxyyzz <<EOF
use rentroll
INSERT INTO RentableMarketRate (RTID,BID,MarketRate,DtStart,DtStop) VALUES(2,1,4000,"2016-10-01 00:00:00","2018-01-01 00:00:00");
INSERT INTO RentalAgreementRentables (RAID,BID,RID,CLID,ContractRent,RARDtStart,RARDtStop) VALUES(2,1,2,0,4000,"2016-10-01 00:00:00","2018-01-01 00:00:00");
EOF
${MYSQL} --no-defaults <xxyyzz
rm -f xxyyzz
#========================================================================================
# OCTOBER 2016
#========================================================================================
RRDATERANGE="-j 2016-10-01 -k 2016-11-01"
CSVLOADRANGE="-G ${BUD} -g 10/1/16,11/1/16"
docsvtest "b10" "-A asm2016-10.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Oct"
dorrtest  "a10" "${RRDATERANGE} -b ${BUD}" "Process-2016-Oct"
docsvtest "i10" "-e rcpt2016-10.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Oct"
docsvtest "j10" "-y deposit-2016-10.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Oct"
dorrtest  "k10" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Oct"
dorrtest  "l10" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m10" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n10" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o10" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p10" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q10" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# NOVEMBER 2016
#========================================================================================
RRDATERANGE="-j 2016-11-01 -k 2016-12-01"
CSVLOADRANGE="-G ${BUD} -g 11/1/16,12/1/16"
docsvtest "b11" "-A asm2016-11.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2016-Nov"
dorrtest  "a11" "${RRDATERANGE} -b ${BUD}" "Process-2016-Nov"
docsvtest "i11" "-e rcpt2016-11.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Nov"
docsvtest "j11" "-y deposit-2016-11.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Nov"
dorrtest  "k11" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Nov"
dorrtest  "l11" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m11" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n11" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o11" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p11" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q11" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

#========================================================================================
# DECEMBER 2016
#========================================================================================
RRDATERANGE="-j 2016-12-01 -k 2017-01-01"
CSVLOADRANGE="-G ${BUD} -g 12/1/16,1/1/17"
# docsvtest "b12" "-A asm2016-12.csv ${CSVLOADRANGE} -L 12,${BUD}" "Assessments-2016-Dec"
dorrtest  "a12" "${RRDATERANGE} -b ${BUD}" "Process-2016-Dec"
docsvtest "i12" "-e rcpt2016-12.csv ${CSVLOADRANGE} -L 13,${BUD}" "Receipts-2016-Dec"
docsvtest "j12" "-y deposit-2016-12.csv ${CSVLOADRANGE} -L 19,${BUD}" "Deposits-2016-Dec"
dorrtest  "k12" "${RRDATERANGE} -b ${BUD}" "Finish-2016-Dec"
dorrtest  "l12" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest  "m12" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"
dorrtest  "n12" "${RRDATERANGE} -b ${BUD} -r 10" "LedgerActivity"
dorrtest  "o12" "${RRDATERANGE} -b ${BUD} -r 17" "LedgerBalance"
dorrtest  "p12" "${RRDATERANGE} -b ${BUD} -r 8" "Statements"
dorrtest  "q12" "${RRDATERANGE} -b ${BUD} -r 4" "RentRoll"

logcheck
