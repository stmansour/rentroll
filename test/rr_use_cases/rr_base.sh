#!/bin/bash

TESTNAME="RR View Use Case Databases"
TESTSUMMARY="Generates separate databases for multiple use cases"

if [ "x${RR_DB_CORE}" = "x" ]; then
    RR_DB_CORE="./"
else
    echo "RR_DB_CORE was pre-set to:  \"${RR_DB_CORE}\""
fi

function dbcore() {
    docsvtest "a" "-b ${RR_DB_CORE}business.csv -L 3" "Business"
    docsvtest "b" "-c ${RR_DB_CORE}coa.csv -L 10,${BUD}" "ChartOfAccounts"
    docsvtest "c" "-ar ${RR_DB_CORE}ar.csv" "AccountRules"
    docsvtest "d" "-m ${RR_DB_CORE}depmeth.csv -L 23,${BUD}" "DepositMethods"
    docsvtest "e" "-d ${RR_DB_CORE}depository.csv -L 18,${BUD}" "Depositories"
    docsvtest "f" "-P ${RR_DB_CORE}pmt.csv -L 12,${BUD}" "PaymentTypes"
    docsvtest "t" "-T ${RR_DB_CORE}ratemplates.csv  -L 8,${BUD}" "RentalAgreementTemplates"
    docsvtest "h" "-p ${RR_DB_CORE}people.csv  -L 7,${BUD}" "People"
    docsvtest "g" "-P ${RR_DB_CORE}pmt.csv -L 12,${BUD}" "PaymentTypes"
}
