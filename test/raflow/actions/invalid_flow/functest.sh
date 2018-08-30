#!/bin/bash

TESTNAME="raflow Actions"
TESTSUMMARY="Test Action Set Pending First Approval to Flow with invalid data"
DBGEN=../../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../../tmp/rentroll"
ACTIONSBIN="../"
THISDIR="invalid_flow"

echo "Create new database..."
mysql --no-defaults rentroll < invalidFlow.sql

source ../../../share/base.sh
source ../actions_base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

execute ${THISDIR}

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

outputCheck "a1"  ""  "action_\"set_pending_first_approval\"_on_flow_with_invalid_data"
outputCheck "a2"  ""  "action_\"set_pending_first_approval\"_on_brand_new_flow_with_invalid_data"

logcheck

exit 0