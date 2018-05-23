#!/bin/bash

TESTNAME="TWS - Scheduled Work tester"
TESTSUMMARY="Validate TaskList usage of TWS"
TOP="../.."
BINDIR="${TOP}/tmp/rentroll"

CREATENEWDB=0

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

if [ ! -f bizerr.csv ]; then
	ln -s ${BINDIR}/bizerr.csv
fi

source ../share/base.sh

./tws3 > z

genericlogcheck "z"  ""  "Validations"

logcheck

exit 0
