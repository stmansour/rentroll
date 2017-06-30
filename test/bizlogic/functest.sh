#!/bin/bash

TESTNAME="Bizlogic Tester"
TESTSUMMARY="Test rentroll business logic"

RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < rex.sql

source ../share/base.sh

./bizlogic > z

genericlogcheck "z"  ""  "Accts-Bizlogic-Checks"

logcheck

exit 0