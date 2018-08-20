#!/bin/bash

TESTNAME="raflow Actions"
TESTSUMMARY="Test Different Actions taken on Flow"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

# echo "Create new database..."
mysql --no-defaults rentroll < raflowactions.sql

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

./actions > z

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

genericlogcheck "z"  ""  "raflowActions"

logcheck

exit 0
