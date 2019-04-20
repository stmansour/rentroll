#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableUseStatus"
TESTSUMMARY="Test Rentable Use Status code"
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
#  Validate that the search query returns the proper data
#
#  Scenario:
#
#  The database in xa should have many use status fields.   Triplets of
#  the form housekeeping, ready, in-service.
#
#
#  Expected Results:
#   1.  Search for UseStatus on rentable 4 and make sure that the data
#       returned matches the patterns we expect.
#------------------------------------------------------------------------------
TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/4" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"
fi

#------------------------------------------------------------------------------
#  TEST c
#
#  Error that came up in UI testing.  Overlapping of same type should be merged.
#
#  Scenario:
#
#  Existing db has a use status 4/1/2019 - 7/31/2019 in Ready State. It also
#  has a record from 7/31/2019 to 12/31/9999 in Ready state.  If we extend
#  the latter record 1 or more days forward it should merge the two records.
#  Similarly if we extend the former 1 day or more earlier, it should merge the
#  two.
#
#  Expected Results:
#   see detailed comments below.  Each case refers to an area in the source
#   code that it should hit.  If there's anything wrong, we'll know right
#   where to go in the source to fix it.
#
#------------------------------------------------------------------------------
TFILES="c"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer
    #-----------------------------------
    # INITIAL RENTABLE USE STATUS
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   08/01/2019 - 12/31/9999
    #   0   04/01/2019 - 08/01/2019
    #   4   03/01/2019   04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 4
    #-----------------------------------

    #--------------------------------------------------
    # SetRentableUseStatus - Case 1a
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # SetStatus  0 4/1/2019 - 9/1/2019
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   04/01/2019 - 12/31/9999
    #   4   03/01/2019   04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 3
    # c0,c1
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"BID":1,"BUD":"REX","RID":1,"UseStatus":0,"DtStart":"4/1/2019","DtStop":"8/31/2019","Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #--------------------------------------------------
    # SetRentableUseStatus - Case 1c
    # SetStatus  3  4/1/2019 - 9/1/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   09/01/2019 - 12/31/9999
    #   3   04/01/2019 - 09/01/2019
    #   4   03/01/2019   04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 4
    # c3,c4
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"BID":1,"BUD":"REX","RID":1,"UseStatus":3,"DtStart":"4/1/2019","DtStop":"8/31/2019","Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #-------------------------------------------------------
    # SetRentableUseStatus - Case 1b
    #-----------------------------------------------
    # CASE 1a -  rus contains b[0], match == false
    #-----------------------------------------------
    #     b[0]: @@@@@@@@@@@@@@@@@@@@@
    #      rus:      ############
    #   Result: @@@@@############@@@@
    #----------------------------------------------------
    # SetStatus  2 9/15/2019 - 9/22/2019
    # Note: EDI in effect, DtStop expressed as "through 9/21/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop     RSID
    #  ---------------------------- ----
    #   0   09/22/2019 - 12/31/9999  15
    #   2   09/15/2019 - 09/22/2019  16
    #   0   09/01/2019 - 09/15/2019  10
    #   3   04/01/2019 - 09/01/2019  14
    #   4   03/01/2019   04/01/2019  11
    #   0   01/01/2018   03/01/2019   5
    # Total Records: 6
    # c5,c6
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"UseStatus":2,"DtStart":"9/15/2019","DtStop":"9/21/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #-------------------------------------------------------
    # SetRentableUseStatus - Case 1d
    #-----------------------------------------------
    # CASE 1d -  rus prior to b[0], match == false
    #-----------------------------------------------
    #      rus:     @@@@@@@@@@@@
    #     b[0]: ##########
    #   Result: ####@@@@@@@@@@@@
    #-----------------------------------------------
    # SetStatus  1 (repair) 3/15/2019 - 9/1/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   09/22/2019 - 12/31/9999
    #   2   09/15/2019 - 09/22/2019
    #   0   09/01/2019 - 09/15/2019
    #   1   03/15/2019 - 09/01/2019
    #   4   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 6
    # c7,c8
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":0,"UseStatus":1,"DtStart":"3/15/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #-------------------------------------------------------
    # SetRentableUseStatus - Case 2b
    #-----------------------------------------------
    #  Case 2b
    #  neither match. Update both b[0] and b[1], add new rus
    #   b[0:1]   @@@@@@@@@@************
    #   rus            #######
    #   Result   @@@@@@#######*********
    #-----------------------------------------------
    # SetStatus  3 8/1/2019 - 9/7/2019
    # Note: EDI in effect, DtStop expressed as "through 9/6/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   09/22/2019 - 12/31/9999
    #   2   09/15/2019 - 09/22/2019
    #   0   09/01/2019 - 09/15/2019
    #   3   08/01/2019 - 09/07/2019
    #   1   03/15/2019 - 08/01/2019
    #   4   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    # c9,c10
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"UseStatus":3,"DtStart":"8/1/2019","DtStop":"9/6/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #-------------------------------------------------------
    # SetRentableUseStatus - Case 2c
    #-----------------------------------------------
    #  Case 2c
    #  merge rus and b[0], update b[1]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            @@@@@@@
    #   Result   @@@@@@@@@@@@@*********
    #-----------------------------------------------
    # SetStatus  1 7/1/2019 - 8/7/2019
    # Note: EDI in effect, DtStop expressed as "through 8/6/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   09/22/2019 - 12/31/9999
    #   2   09/15/2019 - 09/22/2019
    #   0   09/01/2019 - 09/15/2019
    #   3   08/07/2019 - 09/07/2019
    #   1   03/15/2019 - 08/07/2019
    #   4   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    # c11,c12
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"UseStatus":1,"DtStart":"7/1/2019","DtStop":"8/6/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #-------------------------------------------------------
    # SetRentableUseStatus - Case 2d
    #-----------------------------------------------
    #  Case 2d
    #  merge rus and b[1], update b[0]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            *******
    #   Result   @@@@@@****************
    #-----------------------------------------------
    # SetStatus  3 (employee) 8/1/2019 - 8/10/2019
    # Note: EDI in effect, DtStop expressed as "through 8/9/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   09/22/2019 - 12/31/9999
    #   2   09/15/2019 - 09/22/2019
    #   0   09/01/2019 - 09/15/2019
    #   3   08/01/2019 - 09/07/2019
    #   1   03/15/2019 - 08/01/2019
    #   4   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    # c13,c14
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"UseStatus":3,"DtStart":"8/1/2019","DtStop":"8/10/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

    #-------------------------------------------------------
    # SetRentableUseStatus - Case 2a
    #-----------------------------------------------
    #  Case 2a
    #  all are the same, merge them all into b[0], delete b[1]
    #   b[0:1]   ********* ************
    #   rus            *******
    #   Result   **********************
    #-----------------------------------------------
    # SetStatus  0 (ready) 9/7/2019 - 9/30/2019
    # Note: EDI in effect, DtStop expressed as "through 9/29/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   0   09/01/2019 - 12/31/9999
    #   3   08/01/2019 - 09/07/2019
    #   1   03/15/2019 - 08/01/2019
    #   4   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RSID":13,"UseStatus":0,"DtStart":"9/7/2019","DtStop":"9/29/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1" "request" "${TFILES}${STEP}"  "RentableUseStatus-Search"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
