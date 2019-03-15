#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableMarketRate"
TESTSUMMARY="Test Rentable Market Rate code"
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
#  Validate changes to the Rentable Market Rate are handled properly
#
#  Scenario:
#  End dates are in the requests as the actual date - 1 day because the last day
#  is inclusive when EDI is enabled.
#
#
#  Expected Results:
#    Results are detailed in each test case.  The date ranges listed are
#    as expected in the database.  The UI displays end dates as actual day - 1
#    when EDI is enabled.
#------------------------------------------------------------------------------
TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer
    #-----------------------------------
    # INITIAL RENTABLE LEASE STATUS
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  08/01/2019 - 12/31/9999
    # 1200  04/01/2019 - 08/01/2019
    # 1100  03/01/2018 - 04/01/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 4
    #-----------------------------------

    #--------------------------------------------------
    # SetRentableMarketRate - Case 1a
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # MarketRate:  1200  4/1/2019 - 9/1/2019
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  04/01/2019 - 12/31/9999
    # 1100  03/01/2019   04/01/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 3
    # 0,1
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"BID":1,"BUD":"REX","RTID":1,"MarketRate":1200,"DtStart":"4/1/2019","DtStop":"8/31/2019","w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-1a"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #--------------------------------------------------
    # SetRentableMarketRate - Case 1c
    # MarketRate:  1000  4/1/2019 - 9/1/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/01/2019 - 12/31/9999
    # 1000  04/01/2019 - 09/01/2019
    # 1100  03/01/2019   04/01/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 4
    # 2,3
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"BID":1,"BUD":"REX","RTID":1,"MarketRate":1000,"DtStart":"4/1/2019","DtStop":"8/31/2019","w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-1c"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #-------------------------------------------------------
    # SetRentableMarketRate - Case 1b
    #-----------------------------------------------
    # CASE 1a -  rus contains b[0], match == false
    #-----------------------------------------------
    #     b[0]: @@@@@@@@@@@@@@@@@@@@@
    #      rus:      ############
    #   Result: @@@@@############@@@@
    #----------------------------------------------------
    # MarketRate:  1100  9/15/2019 - 9/22/2019
    # Note: EDI in effect, DtStop expressed as "through 9/21/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/22/2019 - 12/31/9999
    # 1100  09/15/2019 - 09/22/2019
    # 1200  09/01/2019 - 09/15/2019
    # 1000  04/01/2019 - 09/01/2019
    # 1100  03/01/2019   04/01/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 6
    # 4,5
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"MarketRate":1100,"DtStart":"9/15/2019","DtStop":"9/21/2019","BID":1,"BUD":"REX","RTID":1,"w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-1b"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #-------------------------------------------------------
    # SetRentableMarketRate - Case 1d
    #-----------------------------------------------
    # CASE 1d -  rus prior to b[0], match == false
    #-----------------------------------------------
    #      rus:     @@@@@@@@@@@@
    #     b[0]: ##########
    #   Result: ####@@@@@@@@@@@@
    #-----------------------------------------------
    # MarketRate: 1100  3/15/2019 - 9/01/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/22/2019 - 12/31/9999
    # 1100  09/15/2019 - 09/22/2019
    # 1200  09/01/2019 - 09/15/2019
    # 1000  03/15/2019 - 09/01/2019
    # 1100  03/01/2018   03/15/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 6
    # 6,7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":0,"MarketRate":1000,"DtStart":"3/15/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","RTID":1,"w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-1d"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #-------------------------------------------------------
    # SetRentableMarketRate - Case 2b
    #-----------------------------------------------
    #  Case 2b
    #  neither match. Update both b[0] and b[1], add new rus
    #   b[0:1]   @@@@@@@@@@************
    #   rus            #######
    #   Result   @@@@@@#######*********
    #-----------------------------------------------
    # MarketRate:  1100  8/1/2019 - 9/7/2019
    # Note: EDI in effect, DtStop expressed as "through 9/6/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/22/2019 - 12/31/9999
    # 1100  09/15/2019 - 09/22/2019
    # 1200  09/07/2019 - 09/15/2019
    # 1100  08/01/2019 - 09/07/2019
    # 1000  03/15/2019 - 08/01/2019
    # 1100  03/01/2018   03/15/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 7
    # 8,9
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"MarketRate":1100,"DtStart":"8/1/2019","DtStop":"9/6/2019","BID":1,"BUD":"REX","RTID":1,"w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-2b"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #-------------------------------------------------------
    # SetRentableMarketRate - Case 2c
    #-----------------------------------------------
    #  Case 2c
    #  merge rus and b[0], update b[1]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            @@@@@@@
    #   Result   @@@@@@@@@@@@@*********
    #-----------------------------------------------
    # MarketRate:  1000  7/1/2019 - 8/7/2019
    # Note: EDI in effect, DtStop expressed as "through 8/6/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/22/2019 - 12/31/9999
    # 1100  09/15/2019 - 09/22/2019
    # 1200  09/01/2019 - 09/15/2019
    # 1100  08/07/2019 - 09/07/2019
    # 1000  03/15/2019 - 08/07/2019
    # 1100  03/01/2018   03/15/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 7
    # 10,11
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"MarketRate":1000,"DtStart":"7/1/2019","DtStop":"8/6/2019","BID":1,"BUD":"REX","RTID":1,"w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-2c"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #-------------------------------------------------------
    # SetRentableMarketRate - Case 2d
    #-----------------------------------------------
    #  Case 2d
    #  merge rus and b[1], update b[0]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            *******
    #   Result   @@@@@@****************
    #-----------------------------------------------
    # MarketRate:  1100  8/1/2019 - 8/10/2019
    # Note: EDI in effect, DtStop expressed as "through 8/9/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/22/2019 - 12/31/9999
    # 1100  09/15/2019 - 09/22/2019
    # 1200  09/01/2019 - 09/15/2019
    # 1100  08/01/2019 - 09/07/2019
    # 1000  03/15/2019 - 08/01/2019
    # 1100  03/01/2018   03/15/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 7
    # 12,13
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"MarketRate":1100,"DtStart":"8/1/2019","DtStop":"8/10/2019","BID":1,"BUD":"REX","RTID":1,"w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-2d"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

    #-------------------------------------------------------
    # SetRentableMarketRate - Case 2a
    #-----------------------------------------------
    #  Case 2a
    #  all are the same, merge them all into b[0], delete b[1]
    #   b[0:1]   ********* ************
    #   rus            *******
    #   Result   **********************
    #-----------------------------------------------
    # MarketRate:  1100  3/7/2019 - 8/6/2019
    # Note: EDI in effect, DtStop expressed as "through 8/5/2019"
    # Result needs to be:
    #  MR$  DtStart      DtStop
    #  ----------------------------
    # 1200  09/22/2019 - 12/31/9999
    # 1100  09/15/2019 - 09/22/2019
    # 1200  09/07/2019 - 09/15/2019
    # 1100  03/01/2018   09/07/2019
    # 1000  01/01/2018   03/01/2019
    # Total Records: 7
    # 14,15
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RMRID":12,"MarketRate":1100,"DtStart":"3/7/2019","DtStop":"8/5/2019","BID":1,"BUD":"REX","RTID":1,"w2ui":{}}],"RTID":1}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Save-2a"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rmr/1/1" "request" "${TFILES}${STEP}"  "RentableMarketRate-Search"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
