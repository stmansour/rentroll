#!/bin/bash

TESTNAME="Init RA"
TESTSUMMARY="Test to init raflow and migrate to permanent tables RA"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

echo "Initialization of New RAFlow..."

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
