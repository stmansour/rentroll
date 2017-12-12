#!/bin/bash

TESTNAME="RAID Account Balance and Expenses"
TESTSUMMARY="Test rentroll RA Acct Balance and Expenses"

CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < baltest.sql

source ../share/base.sh

./acctbal > z

genericlogcheck "z"  ""  "AcctBal-Checks"

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

# get Statement
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222016-11-01%22%2C%22searchDtStop%22%3A%222016-12-01%22%7D" > request
dojsonPOST "http://localhost:8270/v1/stmtdetail/1/5" "request" "a1"  "WebService--StatementDetail"

# get Expenses
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222016-11-01%22%2C%22searchDtStop%22%3A%222017-08-01%22%7D" > request
dojsonPOST "http://localhost:8270/v1/expense/1" "request" "b1"  "WebService--ExpensesSearch"

# get a particular Expense
echo "%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmInstForm%22%7D" > request
dojsonPOST "http://localhost:8270/v1/expense/1/2" "request" "c1"  "WebService--GetExpense"

# save a new Expense
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22expenseForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22EXPID%22%3A0%2C%22ARID%22%3A27%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Dt%22%3A%228%2F11%2F2017%22%2C%22Amount%22%3A12%2C%22AcctRule%22%3A%22%22%2C%22RName%22%3A%22309+S+Rexford%22%2C%22Comment%22%3A%22test%22%2C%22FLAGS%22%3A0%2C%22Mode%22%3A0%2C%22PREXPID%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/expense/1/0" "request" "d1"  "WebService--SaveNewExpense"

# update the comment on the Expense we just saved
echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22expenseForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22EXPID%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22RAID%22%3A1%2C%22Amount%22%3A12%2C%22Dt%22%3A%228%2F11%2F2017%22%2C%22ARID%22%3A27%2C%22ARName%22%3A%22%22%2C%22RName%22%3A%22309+S+Rexford%22%2C%22FLAGS%22%3A0%2C%22Comment%22%3A%22big+time+comment%22%2C%22PREXPID%22%3A%22%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/expense/1/0" "request" "e1"  "WebService--UpdateExpense"

# get StatementInfo
echo "%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22stmtDetailForm%22%7D" > request
dojsonPOST "http://localhost:8270/v1/stmtinfo/1/5" "request" "f1"  "WebService--StatementInfo"


stopRentRollServer
echo "RENTROLL SERVER STOPPED"


logcheck

exit 0
