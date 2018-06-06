#!/bin/bash
TESTNAME="CLOSE PERIOD test"
TESTSUMMARY="Close a period and write ledger markers"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

source ../share/base.sh

echo "BEGIN CLOSE PERIOD FUNCTIONAL TEST" >>${LOGFILE}

logcheck
