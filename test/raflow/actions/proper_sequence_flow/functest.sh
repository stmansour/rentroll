#!/bin/bash

TESTNAME="raflow Actions"
TESTSUMMARY="Test Proper Sequence(0=>1=>2=>3=>4) of Flow"
DBGEN=../../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../../tmp/rentroll"
ACTIONSBIN="../"
THISDIR="proper_sequence_flow"

echo "Create new database..."
mysql --no-defaults rentroll < properSequence.sql

source ../../../share/base.sh
source ../actions_base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

execute ${THISDIR}

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

outputCheck "a1"  ""  "action_\"set_pending_first_approval\"_on_flow_with_valid_data"
outputCheck "a2"  ""  "approve_and_set_\"pending_second_approval\"_on_flow_with_valid_data"
outputCheck "a3"  ""  "approve_and_set_\"move-in_/_execute_modification\"_on_flow_with_valid_data"
outputCheck "a4"  ""  "set_document_date_of_flow_with_valid_data"
outputCheck "a5"  ""  "take_action_of_\"complete_move_in\"_on_flow_with_valid_data"
outputCheck "a6"  ""  "action_\"set_pending_first_approval\"_on_brand_new_flow_with_valid_data"
outputCheck "a7"  ""  "approve_and_set_\"pending_second_approval\"_on_brand_new_flow_with_valid_data"
outputCheck "a8"  ""  "approve_and_set_\"move-in_/_execute_modification\"_on_brand_new_flow_with_valid_data"
outputCheck "a9"  ""  "set_document_date_of_brand_new_flow_with_valid_data"
outputCheck "a10"  ""  "take_action_of_\"complete_move_in\"_on_brand_new_flow_with_valid_data"


logcheck

exit 0