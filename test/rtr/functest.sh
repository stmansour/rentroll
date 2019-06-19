#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableTypeRefs"
TESTSUMMARY="Test Reservation search, create, mod"
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
#  Validate a save.
#
#  Scenario:
#  Change the end date of an RTR but provide erroneous data. Make sure we
#  get error message response.
#
#
#  Expected Results:
#   2.  Error message response
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A0%2C%22RLID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22LeaseStatus%22%3A0%2C%22DtStart%22%3A%221%2F1%2F2019%22%2C%22DtStop%22%3A%221%2F3%2F2019%22%2C%22Comment%22%3A%22%22%2C%22CreateBy%22%3A0%2C%22LastModBy%22%3A0%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}0"  "RentableTypeRefs-SaveWithError"
fi

#------------------------------------------------------------------------------
#  TEST b
#
#  Validate a save with correct data.
#
#  Scenario:
#  Change the end date of an RTR.
#
#
#  Expected Results:
#   2.  the end date of RTID 1 should change to 2/28/2019
#------------------------------------------------------------------------------
TFILES="b"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A0%2C%22RTRID%22%3A1%2C%22RTID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22OverrideRentCycle%22%3A0%2C%22OverrideProrationCycle%22%3A0%2C%22DtStart%22%3A%221%2F1%2F2019%22%2C%22DtStop%22%3A%222%2F28%2F2019%22%2C%22CreateBy%22%3A0%2C%22LastModBy%22%3A0%2C%22w2ui%22%3A%7B%7D%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}1"  "RentableTypeRefs-Save"
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
    # INITIAL RENTABLE TYPE REF
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   08/01/2019 - 12/31/9999
    #   5   04/01/2019 - 08/01/2019
    #   2   03/01/2019   04/01/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 4
    # I know the initial state is in error because there are two
    # consecutive records where
    #-----------------------------------

    #--------------------------------------------------
    # SetRentableTypeRef - Case 1a
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # SetTypeRef  5 4/1/2019 - 9/1/2019
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   04/01/2019 - 12/31/9999
    #   2   03/01/2019   04/01/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 3
    # 0 1
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":5,"DtStart":"4/1/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #--------------------------------------------------
    # SetRentableTypeRef - Case 1c
    #-----------------------------------------------
    # CASE 1c -  rus prior to b[0], match == false
    #-----------------------------------------------
    #      rus: @@@@@@@@@@@@
    #     b[0]:       ##########
    #   Result: @@@@@@@@@@@@####
    #-----------------------------------------------
    # SetTypeRef  7 - 4/1/2019 - 9/1/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   09/01/2019 - 12/31/9999
    #   7   04/01/2019 - 09/01/2019
    #   2   03/01/2019   04/01/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 4
    # 2 3
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":7,"DtStart":"4/1/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #-------------------------------------------------------
    # SetRentableTypeRef - Case 1b
    #-----------------------------------------------
    # CASE 1a -  rus contains b[0], match == false
    #-----------------------------------------------
    #     b[0]: @@@@@@@@@@@@@@@@@@@@@
    #      rus:      ############
    #   Result: @@@@@############@@@@
    #----------------------------------------------------
    # SetTypeRef  6 - 9/15/2019 - 9/22/2019
    # Note: EDI in effect, DtStop expressed as "through 9/21/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop     RSID
    #  ---------------------------- ----
    #   5   09/22/2019 - 12/31/9999  15
    #   6   09/15/2019 - 09/22/2019  16
    #   5   09/01/2019 - 09/15/2019  10
    #   7   04/01/2019 - 09/01/2019  14
    #   2   03/01/2019   04/01/2019  11
    #   1   01/01/2018   03/01/2019   5
    # Total Records: 6
    # 4 5
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":6,"DtStart":"9/15/2019","DtStop":"9/21/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #-------------------------------------------------------
    # SetRentableTypeRef - Case 1d
    #-----------------------------------------------
    # CASE 1d -  rus after to b[0], match == false
    #-----------------------------------------------
    #      rus:     @@@@@@@@@@@@
    #     b[0]: ##########
    #   Result: ####@@@@@@@@@@@@
    #-----------------------------------------------
    # SetTypeRef  6 - 3/15/2019 - 9/7/2019
    # START IN THE MIDDLE OF EXISTING, STOP AT THE BEGINNING OF ANOTHER EXISTING
    # Note: EDI in effect, DtStop expressed as "through 9/6/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   09/22/2019 - 12/31/9999
    #   6   09/15/2019 - 09/22/2019
    #   5   09/01/2019 - 09/15/2019
    #   6   03/15/2019 - 09/01/2019
    #   2   03/01/2018   03/15/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 6
    # 6 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":6,"DtStart":"3/15/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #-------------------------------------------------------
    # SetRentableTypeRef - Case 2b
    #-----------------------------------------------
    #  Case 2b
    #  neither match. Update both b[0] and b[1], add new rus
    #   b[0:1]   @@@@@@@@@@************
    #   rus            #######
    #   Result   @@@@@@#######*********
    #-----------------------------------------------
    # SetTypeRef  3 - 8/1/2019 - 9/7/2019
    # Note: EDI in effect, DtStop expressed as "through 9/6/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   09/22/2019 - 12/31/9999
    #   6   09/15/2019 - 09/22/2019
    #   5   09/07/2019 - 09/15/2019
    #   3   08/01/2019 - 09/07/2019
    #   6   03/15/2019 - 08/01/2019
    #   2   03/01/2018   03/15/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 7
    # 8 9
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":3,"DtStart":"8/1/2019","DtStop":"9/6/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #-------------------------------------------------------
    # SetRentableTypeRef - Case 2c
    #-----------------------------------------------
    #  Case 2c
    #  merge rus and b[0], update b[1]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            @@@@@@@
    #   Result   @@@@@@@@@@@@@*********
    #-----------------------------------------------
    # SetTypeRef  5 - 7/1/2019 - 8/7/2019
    # Note: EDI in effect, DtStop expressed as "through 8/6/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   09/22/2019 - 12/31/9999
    #   6   09/15/2019 - 09/22/2019
    #   5   09/07/2019 - 09/15/2019
    #   3   08/07/2019 - 09/07/2019
    #   6   03/15/2019 - 08/07/2019
    #   2   03/01/2018   03/15/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 7
    # 10 11
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":6,"DtStart":"7/1/2019","DtStop":"8/6/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #-------------------------------------------------------
    # SetRentableTypeRef - Case 2d
    #-----------------------------------------------
    #  Case 2d
    #  merge rus and b[1], update b[0]
    #   b[0:1]   @@@@@@@@@@************
    #   rus            *******
    #   Result   @@@@@@****************
    #-----------------------------------------------
    # SetTypeRef  3 - 8/1/2019 - 8/10/2019
    # Note: EDI in effect, DtStop expressed as "through 8/9/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   09/22/2019 - 12/31/9999
    #   6   09/15/2019 - 09/22/2019
    #   5   09/07/2019 - 09/15/2019
    #   3   08/01/2019 - 09/07/2019
    #   6   03/15/2019 - 08/01/2019
    #   2   03/01/2018   03/15/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 7
    # 12 13
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":3,"DtStart":"8/1/2019","DtStop":"8/9/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"

    #-------------------------------------------------------
    # SetRentableTypeRef - Case 2a
    #-----------------------------------------------
    #  Case 2a
    #  all are the same, merge them all into b[0], delete b[1]
    #   b[0:1]   ********* ************
    #   rus            *******
    #   Result   **********************
    #-----------------------------------------------
    # SetTypeRef  6 - 7/1/2019 - 9/20/2019
    # Note: EDI in effect, DtStop expressed as "through 9/19/2019"
    # Result needs to be:
    # RTID  DtStart      DtStop
    #  ----------------------------
    #   5   09/22/2019 - 12/31/9999
    #   6   03/15/2019 - 09/22/2019
    #   2   03/01/2018   03/15/2019
    #   1   01/01/2018   03/01/2019
    # Total Records: 7
    # 14 15
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"RTID":6,"DtStart":"7/1/2019","DtStop":"9/19/2019","BID":1,"BUD":"REX","recid":0,"RTRID":23,"RID":1,"OverrideRentCycle":0,"OverrideProrationCycle":0,"CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRef-Search"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
