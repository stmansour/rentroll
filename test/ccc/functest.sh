#!/bin/bash
TESTNAME="CCC - Setup CCC and its properties"
TESTSUMMARY="Setup and run CCC company and its properties"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
BUD="CCC"

source ../share/base.sh


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
# docsvtest "f" "-l strlists.csv -L 25,${BUD}" "StringLists"
docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
# docsvtest "i" "-u custom.csv -L 14,${BUD}" "CustomAttributes"
# docsvtest "j" "-U assigncustom.csv -L 15,${BUD}" "AssignCustomAttributes"
docsvtest "k" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
docsvtest "l" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
# docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"

logcheck
