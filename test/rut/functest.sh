#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableUseType"
TESTSUMMARY="Test Rentable Lease Status code"
DBGENDIR=${SRCTOP}/tools/dbgen
CREATENEWDB=0
RRBIN="${SRCTOP}/tmp/rentroll"
CATRML="${SRCTOP}/tools/catrml/catrml"

#SINGLETEST=""  # This runs all the tests

source ${TESTHOME}/share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
# RENTROLLSERVERNOW="-testDtNow 10/24/2018"

#------------------------------------------------------------------------------
#  TEST a
#
#  Validate that the dates are properly EDI handled
#
#  Scenario:
#  End dates are listed as the actual date - 1day because the last day is
#  inclusive
#
#  Expected Results:
#   1.  In the database, the key date ranges are set as follows:
#		1/1/2019 - 1/3/2019
#		1/3/2019 - 3/1/2020
#		3/1/2020 - 12/31/9999
#
#       Since the business has the EDI flag set, the UI must send
#       the data with the following date ranges:
#		1/1/2019 - 1/2/2019
#		1/3/2019 - 2/29/2020
#		3/1/2020 - 12/30/9999
#------------------------------------------------------------------------------
TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}' > request
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-SaveWithError"
fi

#------------------------------------------------------------------------------
#  TEST b
#
#  Validate that save biz logic catches overlap propblems. This was from a bug
#  discovered in the UI.
#
#  Scenario:
#  A new RentableStatusRecord overlaps with an existing record
#
#  Expected Results:
#   1.  In the database, the RentableUseType records for RID 1 are:
#       102 3/1/2020 - 3/5/2020
#       101 1/3/2019 - 3/1/2020
#		100 1/1/2019 - 1/3/2019
#
#       An attempt to save a new record with this date range:
#		3/4/2020 - 12/31/9999  UseType 102
#       This will change the 3rd region above to 3/1/2020 - 3/4/2020
#       and add a new record from 3/4/2020 to 12/31/9999
#
#   2.  Next we attempt to save a new record with this date range
#		3/5/2020 - 12/31/9999 UseType 100 and this should change the records to:
#
#       UT   DRange
#       ---  ---------------------
#       100  03/05/2020 - 12/31/9999
#       102  03/01/2020 - 03/05/2020
#       101  01/03/2019 - 03/01/2020
#       100  01/01/2019 - 01/03/2019
#------------------------------------------------------------------------------
TFILES="b"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":3,"BID":1,"BUD":"REX","RID":1,"UTID":0,"UseType":102,"DtStart":"3/4/2020","DtStop":"12/30/9999"}],"RID":1}' > request
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save"

    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":3,"BID":1,"BUD":"REX","RID":1,"UTID":0,"UseType":100,"DtStart":"3/5/2020","DtStop":"12/30/9999"}],"RID":1}' > request
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save"
fi

#------------------------------------------------------------------------------
#  TEST f
#
#  Error that came up in UI testing.  Overlapping of same type should be merged.
#
#  Scenario:
#
#  test all known cases of SetRentableUseType
#
#  Expected Results:
#   see detailed comments below.  Each case refers to an area in the source
#   code that it should hit.  If there's anything wrong, we'll know right
#   where to go in the source to fix it.
#
#------------------------------------------------------------------------------
TFILES="f"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer
    #-----------------------------------
    # INITIAL RENTABLE LEASE STATUS
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   08/01/2019 - 12/31/9999
    #  102   04/01/2019 - 08/01/2019
    #  101   03/01/2018 - 04/01/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 4
    #-----------------------------------

    #--------------------------------------------------
    # SetRentableUseType - Case 1a
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # SetStatus  2 (reserved) 4/1/2019 - 9/1/2019
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   04/01/2019 - 12/31/9999
    #  101   03/01/2019   04/01/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 3
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"BID":1,"BUD":"REX","RID":1,"UseType":102,"DtStart":"4/1/2019","DtStop":"8/31/2019","Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-1a"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #--------------------------------------------------
    # SetRentableUseType - Case 1c
    # SetStatus  0 (not leased) 4/1/2019 - 9/1/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/01/2019 - 12/31/9999
    #  100   04/01/2019 - 09/01/2019
    #  101   03/01/2019   04/01/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 4
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"BID":1,"BUD":"REX","RID":1,"UseType":100,"DtStart":"4/1/2019","DtStop":"8/31/2019","Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-1c"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #-------------------------------------------------------
    # SetRentableUseType - Case 1b
    #-----------------------------------------------
    # CASE 1a -  rus contains b[0], match == false
    #-----------------------------------------------
    #     b[0]: @@@@@@@@@@@@@@@@@@@@@
    #      rus:      ############
    #   Result: @@@@@############@@@@
    #----------------------------------------------------
    # SetStatus  1 (leased) 9/15/2019 - 9/22/2019
    # Note: EDI in effect, DtStop expressed as "through 9/21/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/22/2019 - 12/31/9999
    #  101   09/15/2019 - 09/22/2019
    #  102   09/01/2019 - 09/15/2019
    #  100   04/01/2019 - 09/01/2019
    #  101   03/01/2019   04/01/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 6
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"UseType":101,"DtStart":"9/15/2019","DtStop":"9/21/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-1b"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #-------------------------------------------------------
    # SetRentableUseType - Case 1d
    #-----------------------------------------------
    # CASE 1d -  rus prior to b[0], match == false
    #-----------------------------------------------
    #      rus:     @@@@@@@@@@@@
    #     b[0]: ##########
    #   Result: ####@@@@@@@@@@@@
    #-----------------------------------------------
    # SetStatus 1 (leased) 3/15/2019 - 9/01/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/22/2019 - 12/31/9999
    #  101   09/15/2019 - 09/22/2019
    #  102   09/01/2019 - 09/15/2019
    #  100   03/15/2019 - 09/01/2019
    #  101   03/01/2018   03/15/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 6
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"UseType":100,"DtStart":"3/15/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-1d"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #-------------------------------------------------------
    # SetRentableUseType - Case 2b
    #-----------------------------------------------
    #  Case 2b
    #  neither match. Update both b[0] and b[1], add new rus
    #   b[0:1]   @@@@@@@@@@************
    #   rus            #######
    #   Result   @@@@@@#######*********
    #-----------------------------------------------
    # SetStatus  1 (leased) 8/1/2019 - 9/7/2019
    # Note: EDI in effect, DtStop expressed as "through 9/6/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/22/2019 - 12/31/9999
    #  101   09/15/2019 - 09/22/2019
    #  102   09/07/2019 - 09/15/2019
    #  101   08/01/2019 - 09/07/2019
    #  100   03/15/2019 - 08/01/2019
    #  101   03/01/2018   03/15/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"UseType":101,"DtStart":"8/1/2019","DtStop":"9/6/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-2b"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #-------------------------------------------------------
    # SetRentableUseType - Case 2c
    #-----------------------------------------------
    #  Case 2c
    #  merge rus and b[0], update b[1]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            @@@@@@@
    #   Result   @@@@@@@@@@@@@*********
    #-----------------------------------------------
    # SetStatus  0 (not leased) 7/1/2019 - 8/7/2019
    # Note: EDI in effect, DtStop expressed as "through 8/6/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/22/2019 - 12/31/9999
    #  101   09/15/2019 - 09/22/2019
    #  102   09/07/2019 - 09/15/2019
    #  101   08/07/2019 - 09/07/2019
    #  100   03/15/2019 - 08/07/2019
    #  101   03/01/2018   03/15/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"UseType":100,"DtStart":"7/1/2019","DtStop":"8/6/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-2c"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #-------------------------------------------------------
    # SetRentableUseType - Case 2d
    #-----------------------------------------------
    #  Case 2d
    #  merge rus and b[1], update b[0]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            *******
    #   Result   @@@@@@****************
    #-----------------------------------------------
    # SetStatus  1 (leased) 8/1/2019 - 8/10/2019
    # Note: EDI in effect, DtStop expressed as "through 8/9/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/22/2019 - 12/31/9999
    #  101   09/15/2019 - 09/22/2019
    #  102   09/07/2019 - 09/15/2019
    #  101   08/01/2019 - 09/07/2019
    #  100   03/15/2019 - 08/01/2019
    #  101   03/01/2018   03/15/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"UseType":101,"DtStart":"8/1/2019","DtStop":"8/10/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-2d"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

    #-------------------------------------------------------
    # SetRentableUseType - Case 2a
    #-----------------------------------------------
    #  Case 2a
    #  all are the same, merge them all into b[0], delete b[1]
    #   b[0:1]   ********* ************
    #   rus            *******
    #   Result   **********************
    #-----------------------------------------------
    # SetStatus  1 (leased) 3/7/2019 - 8/6/2019
    # Note: EDI in effect, DtStop expressed as "through 8/5/2019"
    # Result needs to be:
    #   UT  DtStart      DtStop
    #  ----------------------------
    #  102   09/22/2019 - 12/31/9999
    #  101   09/15/2019 - 09/22/2019
    #  102   09/07/2019 - 09/15/2019
    #  101   03/01/2018   09/07/2019
    #  100   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"UTID":0,"UseType":101,"DtStart":"3/7/2019","DtStop":"8/5/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Save-2a"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1" "request" "${TFILES}${STEP}"  "RentableUseType-Search"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
