#!/bin/bash
TESTNAME="3 Months Books, then Delinquency Report"
TESTSUMMARY="Verify Delinquency Reports"

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"
docsvtest "d" "-R rentabletypes.csv -L 5,${BUD}" "RentableTypes"
docsvtest "c" "-c coa.csv -L 10,${BUD}" "ChartOfAccounts"
docsvtest "e" "-p people.csv  -L 7" "People"
docsvtest "f" "-r rentable.csv -L 6,${BUD}" "Rentables"
docsvtest "g" "-T ratemplates.csv  -L 8" "RentalAgreementTemplates"
docsvtest "h" "-C ra.csv -L 9,${BUD}" "RentalAgreements"
docsvtest "i" "-E pets.csv -L 16,RA0001" "Pets"
docsvtest "j" "-s specialties.csv" "Specialties"
docsvtest "k" "-F rspref.csv" "SpecialtyRefs"
docsvtest "l" "-A asmt.csv -L 11,${BUD}" "Assessments"
docsvtest "m" "-P pmt.csv -L 12,${BUD}" "PaymentTypes"
docsvtest "n" "-e rcpt.csv -L 13,${BUD}" "Receipts"
docsvtest "o" "-u custom.csv -L 14" "CustomAttributes"
docsvtest "p" "-U assigncustom.csv -L 15" "AssignCustomAttributes"

dorrtest "q" "-j 2016-03-01 -k 2016-04-01" "Process"
dorrtest "r" "-j 2016-03-01 -k 2016-04-01 -r 1" "March-Journal"
dorrtest "s" "-j 2016-03-01 -k 2016-04-01 -r 2" "March-Ledger"
dorrtest "t" "-j 2016-04-01 -k 2016-05-01" "April-Process"
dorrtest "u" "-j 2016-05-01 -k 2016-06-01" "May-Process"
dorrtest "v" "-j 2016-05-01 -k 2016-06-01 -r 1" "May-Journal"
dorrtest "w" "-j 2016-05-01 -k 2016-06-01 -r 2" "May-Ledger"
dorrtest "x" "-j 2016-05-01 -k 2016-06-01 -r 8" "May-Statement"
dorrtest "y" "-j 2016-05-01 -k 2016-06-01 -r 14,2016-05-23" "May-23-Delinquency"
dorrtest "z" "-j 2016-05-01 -k 2016-06-01 -r 14,2016-07-01" "Jul-01-Delinquency"

logcheck
