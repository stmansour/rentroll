#!/bin/bash

TESTNAME="raflow Actions"
TESTSUMMARY="Test Improper Sequence of Flow"
DBGEN=../../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../../tmp/rentroll"
ACTIONSBIN="../"
THISDIR="improper_sequence_flow"

echo "Create new database..."
mysql --no-defaults rentroll < improperSequence.sql

source ../../../share/base.sh
source ../actions_base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 10/15/2018"
startRentRollServer

execute ${THISDIR}

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

outputCheck "a1"  ""  "action_\"set_pending_first_approval\"_on_flow_with_valid_data"
outputCheck "a2"  ""  "action_\"set_to_move-in\"_on_flow_with_valid_data"
outputCheck "a3"  ""  "set_document_date_of_flow_with_valid_data"
outputCheck "a4"  ""  "take_action_of_\"complete_move_in\"_on_flow_with_valid_data"
outputCheck "a5"  ""  "action_\"set_to_move-in\"_on_brand_new_flow1_with_valid_data"
outputCheck "a6"  ""  "set_document_date_of_brand_new_flow1_with_valid_data"
outputCheck "a7"  ""  "take_action_of_\"complete_move_in\"_on_brand_new_flow1_with_valid_data"
outputCheck "a8"  ""  "action_\"set_pending_second_approval\"_on_brand_new_flow2_with_valid_data"
outputCheck "a9"  ""  "approve_and_set_\"move-in_/_execute_modification\"_on_brand_new_flow2_with_valid_data"
outputCheck "a10"  ""  "set_document_date_of_brand_new_flow2_with_valid_data"
outputCheck "a11"  ""  "take_action_of_\"complete_move_in\"_on_brand_new_flow2_with_valid_data"

logcheck

exit 0
