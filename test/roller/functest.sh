#!/bin/bash

TESTNAME="RAID Account Balance and Expenses"
TESTSUMMARY="Test rentroll RA Acct Balance and Expenses"

CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < prodrr.sql

source ../share/base.sh

docsvtest "a" "-b business.csv -L 3" "Business"

#------------------------------------------------------------------
#  This approach is unique to this test. Since we need the same
#  chart of accounts in multiple businesses, we will create a
#  copy of the chart of accounts for REX and change rex to the
#  name of each business, save it to a temp file, and import it.
#  That way we only need to keep a single file of accounts.
#------------------------------------------------------------------

declare -a arr=("BRO" "CCC" "OL2" "PAC" "SUM")
for i in "${arr[@]}"
do
	sed "s/REX/${i}/" coa.csv > tmpcoa
	sed "s/REX/${i}/" ar.csv > tmpar
	docsvtest "b${i}" "-c tmpcoa -L 10,${i}" "ChartOfAccounts${i}"
	docsvtest "c${i}" "-ar tmpar -L 29,${i}" "AccountRules{i}"
done


logcheck

exit 0
