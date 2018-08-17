#!/bin/bash

TESTNAME="raflow Actions"
TESTSUMMARY="Test Different Actions taken on Flow"
DBGEN=../../tools/dbgen
CREATENEWDB=0

# echo "Create new database..."
mysql --no-defaults rentroll < raflowactions.sql

source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a0
#  Action setFirstApporval for flow with incomplete data of existing RA
#
#  Expected Results:
#  1. Error: must have at least one occupant
#------------------------------------------------------------------------------

# hit api to take "Set Pending First Approval" for flow with incomplete data of Existing RA 
echo "%7B%0A%20%20%20%20%22UserRefNo%22%3A%20%22VJFC558GW9MM625CT176%22%2C%0A%20%20%20%20%22RAID%22%3A%202%2C%0A%20%20%20%20%22Version%22%3A%20%22refno%22%2C%0A%20%20%20%20%22Action%22%3A%201%2C%0A%20%20%20%20%22Mode%22%3A%20%22Action%22%0A%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a0"  "RAFlowActions -- Invalid_Flow_Of_Existing_RA_Submit"

#------------------------------------------------------------------------------
#  TEST a1
#  Action setFirstApporval for flow with incomplete data of new RA
#
#  Expected Results:
#  1. Error: must have at least one parent rentable
#  2. Error: must have at least one occupant
#------------------------------------------------------------------------------

# hit api to take "Set Pending First Approval" for flow with incomplete data of Existing RA
echo "%7B%22UserRefNo%22%3A%22YCE20N8G44N45TIW1M95%22%2C%22RAID%22%3A0%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A1%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a1" "RAFlowActions -- Invalid_Flow_Of_New_RA_Submit"

#------------------------------------------------------------------------------
#  TEST a2
#  Action setFirstApporval for flow with proper data of Existing RA
#------------------------------------------------------------------------------

# hit api to take "Set Pending First Approval" for flow of Existing RA
echo "%7B%22UserRefNo%22%3A%22FU1T222ATL6HWFS61388%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A1%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a2" "RAFlowActions -- Flow_Of_Existing_RA_State_0_To_1_via_Action"

#------------------------------------------------------------------------------
#  TEST a3
#  State pendingFirstApproval to pendingSecondApporval of flow with proper data of Existing RA
#------------------------------------------------------------------------------

# hit api to set state to "Pending Second Approval" for flow of Existing RA
echo "%7B%22UserRefNo%22%3A%22FU1T222ATL6HWFS61388%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Mode%22%3A%22State%22%2C%22Decision1%22%3A1%2C%22DeclineReason1%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a3" "RAFlowActions -- Flow_Of_Existing_RA_State_1_To_2_via_State"

#------------------------------------------------------------------------------
#  TEST a4
#  State pendingSecondApporval to MoveIn of flow with proper data of Existing RA
#------------------------------------------------------------------------------

# hit api to set state to "Move In" for flow of Existing RA
echo "%7B%22UserRefNo%22%3A%22FU1T222ATL6HWFS61388%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Mode%22%3A%22State%22%2C%22Decision2%22%3A1%2C%22DeclineReason2%22%3A0%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a4" "RAFlowActions -- Flow_Of_Existing_RA_State_2_To_3_via_State"

# echo "" > request
# dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a#" "RAFlowActions -- "

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
