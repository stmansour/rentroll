#!/bin/bash
TESTNAME="CCC"
TESTSUMMARY="Setup a database with multiple businesses for testing purposes"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
BUD="OKC"

source ../share/base.sh

pushd ../jm1;./functest.sh;popd

docsvtest "a" "-b Business.csv -L 3" "Business"
docsvtest "b" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "b1" "-c ccccoa.csv -L 10,CCC" "ChartOfAccounts"
docsvtest "c" "-m depmeth.csv -L 23,${BUD}" "DepositMethods"
docsvtest "d" "-d depository.csv -L 18,${BUD}" "Depositories"
# docsvtest "e" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "e1" "-R cccrt.csv -L 5,CCC" "RentableTypes"
# docsvtest "f" "-l strlists.csv -L 25,${BUD}" "StringLists"
docsvtest "g1" "-p cccpeeps.csv  -L 7,CCC" "People"
#docsvtest "g" "-p people.csv  -L 7,${BUD}" "People"
#docsvtest "h" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "h1" "-r cccrentable.csv -L 6,CCC" "Rentables"

doOnesiteTest "g" "-csv ./onesite-RentRoll-Isola.csv -bud ${BUD} -testmode 1" "OnesiteRentrollCSV"
#doOnesiteTest "g" "-csv ./onesite-2017FEB03.csv -bud ${BUD} -testmode 1" "OnesiteRentrollCSV"

# docsvtest "i" "-u custom.csv -L 14" "CustomAttributes"
# docsvtest "j" "-U assigncustom.csv -L 15" "AssignCustomAttributes"
# docsvtest "k" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
# docsvtest "l" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
# docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
