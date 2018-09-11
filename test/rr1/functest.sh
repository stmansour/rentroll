#!/bin/bash

TESTNAME="Renter Changing During RentCycle"
TESTSUMMARY="Tests changing Rental Agreements multiple times during a rent cycle"

RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "b" "-u custom.csv -L 14,${BUD}" "CustomAttributes"
docsvtest "c" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "d" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "e" "-p people.csv  -L 7,${BUD}" "People"
docsvtest "ea" "-V vehicle.csv" "Vehicles"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "g" "-T ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
docsvtest "h" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "i" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "j" "-U assigncustom.csv -L 15,${BUD}" "AssignCustomAttributes"

#
# Test k
# This loads 3 assessments involving rentable 4.  Rentable 4 goes through three
# separate rent changes during the month of July.  Here's how it breaks down:
#
#     Remember that ranges are specified with the
#     end date meaning up-to-but-not-including.
#     Mathmatically, [start,end)
#
#     Dates                      Market Rate
#     ------------------------------------------
#     11/01/2013 - 07/05/2016    $3100 / month
#     07/05/2016 - 07/09/2016    $4000 / month
#     07/11/2016 - ENDOFTIME     $3200 / month
#
# The assessments are:
#
#  1. 07/01/2016 - 07/04/2016    3 days @ $3100/31 -> $100/day     = $ 300.00
#  2. 07/05/2016 - 07/09/2016    4 days @ $4000/31 -> $129.032/day = $ 516.13
#  3. 07/11/2016 - 08/01/2016   21 days @ $3200/31 -> $103.226/day = $2167.74
#
# So the expected amounts in the report are: $300, $516.13, and $2167.74
#-------------------------------------------------------------------------------
docsvtest "k" "-A asmt.csv -G ${BUD} -g 7/1/16,8/1/16 -L 11,${BUD}" "Assessments"
docsvtest "l" "-e rcpt.csv -G ${BUD} -g 7/1/16,8/1/16 -L 13,${BUD}" "Receipts"

# process payments and receipts
dorrtest "m" "${RRDATERANGE} -b ${BUD}" "Process"
dorrtest "n" "${RRDATERANGE} -b ${BUD} -r 1" "Journal"
dorrtest "o" "${RRDATERANGE} -b ${BUD} -r 2" "Ledgers"

logcheck
