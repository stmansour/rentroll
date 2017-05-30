#!/bin/bash
TESTNAME="GSR - Test Account Rules"
TESTSUMMARY="Test AR by processing assessments and a recipt"

RRDATERANGE="-j 2017-01-01 -k 2017-02-01"

source ../share/base.sh

echo "BEGIN GSR FUNCTIONAL TEST" >>${LOGFILE}
#========================================================================================
# INITIALIZE THE BUSINESS
#   This section has the 1-time tasks to set up the business and get the accounts to
#   their correct starting values.
#========================================================================================
docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "ar" "-ar ar.csv" "AccountRules"
docsvtest "c" "-m depmeth.csv -L 23,${BUD}" "DepositMethods"
docsvtest "d" "-d depository.csv -L 18,${BUD}" "Depositories"
docsvtest "e" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "ga" "-V vehicle.csv -L 28,${BUD}" "Vehicles"
docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "i" "-u custom.csv -L 14,${BUD}" "CustomAttributes"
docsvtest "j" "-U assigncustom.csv -L 15,${BUD}" "AssignCustomAttributes"
docsvtest "k" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
docsvtest "l" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"


#========================================================================================
# JANUARY 2017
#========================================================================================
RRDATERANGE="-j 2017-01-01 -k 2017-02-01"
CSVLOADRANGE="-G ${BUD} -g 1/1/17,2/1/17"

# 1.  Load new assessments for this period.  For this test, we start the rent assessments now.
docsvtest "b1" "-A asm2017-01.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2017-JAN"

# 2.  Generate recurring assessment instances  -  Note: will be done by server automatically by the TimedWorkScheduler  (18 processes journal only)
dorrtest "a1" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-JOURNALS-2017-JAN"

# 3.  Force the ledger entries to be made for this  -- Note: will be done by server automatically by the TimedWorkScheduler  (19 processes ledgers only)
dorrtest "a1a" "${RRDATERANGE} -x -b ${BUD} -r 19" "Process-LEDGERS-2017-JAN"

#========================================================================================
# FEBRUARY 2017
#========================================================================================
RRDATERANGE="-j 2017-02-01 -k 2017-03-01"
CSVLOADRANGE="-G ${BUD} -g 2/1/17,3/1/17"
docsvtest "b2" "-A asm2017-02.csv ${CSVLOADRANGE} -L 11,${BUD}" "Assessments-2017-FEB"
dorrtest "a2" "${RRDATERANGE} -x -b ${BUD} -r 18" "Process-JOURNALS-2017-FEB"
dorrtest "a2a" "${RRDATERANGE} -x -b ${BUD} -r 19" "Process-LEDGERS-2017-FEB"

# Run the checker...
./gsr -db

logcheck
