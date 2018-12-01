#!/bin/bash

TESTNAME="raflow Actions"
TESTSUMMARY="Test RAID version by taking different actions on RA"
DBGEN=../../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../../tmp/rentroll"
ACTIONSBIN="../"
THISDIR="raid_version"

echo "Create new database..."
mysql --no-defaults rentroll < raidVersion.sql

source ../../../share/base.sh
source ../actions_base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 10/15/2018"
startRentRollServer

execute ${THISDIR}

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

outputCheck "a1"  ""  "action_\"recieved_notice_to_move\"_and_provide_NoticeToMoveDate_on_RA_with_state_Active"
outputCheck "a2"  ""  "action_\"terminate\"_on_RA_with_state_Active"
outputCheck "a3"  ""  "action_\"complete_move_in\"_on_RA_with_state_NoticeToMove"
outputCheck "a4"  ""  "action_\"terminate\"_on_RA_with_state_NoticeToMove"
outputCheck "a5"  ""  "action_\"complete_move_in\"_on_RA_with_state_Terminated"
outputCheck "a6"  ""  "action_\"application_being_completed\"_on_RA_with_state_Active"
outputCheck "a7"  ""  "action_\"set_pending_first_approval\"_on_RA_with_state_Active"
outputCheck "a8"  ""  "action_\"application_being_completed\"_on_RA_with_state_NoticeToMove"
outputCheck "a9"  ""  "action_\"application_being_completed\"_on_RA_with_state_Terminated"

logcheck

exit 0
