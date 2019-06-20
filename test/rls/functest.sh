#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="RentableLeaseStatus"
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

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-SaveWithError"
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
#   1.  In the database, the RentableLeaseStatus records for RID 1 are:
#		1/1/2019 - 1/3/2019
#		1/3/2019 - 3/1/2020
#		3/1/2020 - 3/5/2020
#
#       An attempt to save a new record with this date range:
#		3/4/2020 - 12/31/9999
#       This will change the 3rd region above to 3/1/2020 - 3/4/2020
#       and add a new record from 3/4/2020 to 12/31/9999
#
#   1.  Next we attempt to save a new record with this date range
#		3/5/2020 - 12/31/9999
#       and this should work.
#------------------------------------------------------------------------------
TFILES="b"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22RLID%22%3A0%2C%22LeaseStatus%22%3A0%2C%22DtStart%22%3A%223%2F4%2F2020%22%2C%22DtStop%22%3A%2212%2F1%2F9999%22%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save"

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22RLID%22%3A0%2C%22LeaseStatus%22%3A0%2C%22DtStart%22%3A%223%2F5%2F2020%22%2C%22DtStop%22%3A%2212%2F1%2F9999%22%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save"
fi

#------------------------------------------------------------------------------
#  TEST c
#
#  Validate search service virtual scroll support discovered in the UI.
#  Also, test the delete web service
#
#  Scenario:
#  The RentableStatusRecords for a rentable are greater than 100 (default
#  request size). This test will validate the return values for successive
#  calls from the virtual control list.
#
#  Expected Results:
#   1.  First batch has OFFSET = 0, LIMIT = 100.
#		The count will be > 100, but the returned solution set will contain
#		100 entries.
#
#   1.  Next we attempt to save a new record with this date range
#		3/5/2020 - 12/30/9999
#       and this should work.
#
#   3.  Delete 3 RLID records in one call (254,255,171)
#       After the delete, a fetch over date range 2/16/2022 - 12/31/2022
#       should result in only one RLID (172)
#------------------------------------------------------------------------------
TFILES="c"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Get"

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A100%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}${STEP}"  "RentableTypeRefs-GetOffset"

    # Delete 254,255,171
    echo "%7B%22cmd%22%3A%22delete%22%2C%22RLIDList%22%3A%5B254%2C255%2C171%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}${STEP}"  "RentableTypeRefs-GetOffset"
    # Read back time range 2/16/2022 - 12/31/2022.  We should only find 1 entry (RLID=172)
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F16%2F2022%22%2C%22searchDtStop%22%3A%2212%2F31%2F2022%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}${STEP}"  "RentableTypeRefs-GetOffset"

fi

#------------------------------------------------------------------------------
#  TEST d
#
#  Validate the update of existing RentableLeaseStatus records.  The test is
#  the result of a bug where a new record is inserted that has the same end
#  date as the record after it.
#
#  Scenario:
#  The results of updates to RentableLeaseStatus records should be that
#  the start to end time ranged formed by all records has no overlaps.
#  The first test starts with the database as follows:
#
#       7/ 1/2019 - 12/31/9999  Reserved
#       7/ 1/2018 -  6/30/2019  Leased
#       1/ 1/2018 -  6/30/2018  Not Leased
#
#  Expected Results:
#
#   1.  The test will change the Leased range from 7/1/2018 - 6/21/2019.
#       The result should be:
#
#       6/21/2019 - 12/31/9999  Reserved
#       7/ 1/2018 -  7/01/2019  Leased      // no change - it's setting a time span to the same status it's already set to
#       1/ 1/2018 -  6/30/2018  Not Leased
#
#------------------------------------------------------------------------------
TFILES="d"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    # change to Leased = 7/1/2018 - 6/21/2019  (note: xd.sql was already in that
    # LeaseStatus state. But it should not add a new record)
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":3,"BID":1,"BUD":"REX","RID":1,"LeaseStatus":1,"DtStart":"7/1/2018","DtStop":"6/20/2019","Comment":"","CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}' > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Save"

    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Get"

    # set  6/21/2019 - 12/31/9999 to Reserved
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":0,"RLID":4,"BID":1,"BUD":"REX","RID":1,"LeaseStatus":2,"DtStart":"6/21/2019","DtStop":"12/30/9999","Comment":"","CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Save"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Get"


    # do both of the above actions in a single webcall
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":3,"BID":1,"BUD":"REX","RID":1,"LeaseStatus":1,"DtStart":"7/1/2018","DtStop":"6/20/2019","Comment":"","CreateBy":0,"LastModBy":0,"w2ui":{}},{"recid":0,"RLID":4,"BID":1,"BUD":"REX","RID":1,"LeaseStatus":2,"DtStart":"6/21/2019","DtStop":"12/30/9999","Comment":"","CreateBy":0,"LastModBy":0,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-GetOffset"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Get"

fi

#------------------------------------------------------------------------------
#  TEST e
#
#  This test covers an issue found in the Rental Agreement testing
#
#  Scenario:
#  Start with these Rentable Lease Status records
#
#       3/01/2020 - 12/31/9999  Reserved
#       2/13/2018 -  3/01/2020  Leased
#       1/01/2017 -  2/13/2018  Not Leased
#
#  Then do a SetRentableLeaseStatus for range 2/13/2018 - 3/1/2020
#
#  Expected Results:
#
#   1.  The test will change the Leased range from 2/13/2018 - 3/1/2020.
#
#       3/01/2020 - 12/31/9999  Reserved
#       2/13/2018 -  3/01/2020  Leased
#       1/01/2017 -  2/13/2018  Not Leased
#
#------------------------------------------------------------------------------
TFILES="e"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    # change to Leased = 2/13/2018 - 3/1/2020
    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A1%2C%22RLID%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22LeaseStatus%22%3A1%2C%22DtStart%22%3A%222%2F13%2F2018%22%2C%22DtStop%22%3A%222%2F29%2F2020%22%2C%22Comment%22%3A%22%22%2C%22CreateBy%22%3A0%2C%22LastModBy%22%3A0%2C%22w2ui%22%3A%7B%7D%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Save"

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableTypeRefs-Get"

fi

#------------------------------------------------------------------------------
#  TEST f
#
#  Error that came up in UI testing.  Overlapping of same type should be merged.
#
#  Scenario:
#
#  test all known cases of SetRentableLeaseStatus
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   08/01/2019 - 12/31/9999
    #   2   04/01/2019 - 08/01/2019
    #   1   03/01/2018 - 04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 4
    #-----------------------------------

    #--------------------------------------------------
    # SetRentableLeaseStatus - Case 1a
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # SetStatus  2 (reserved) 4/1/2019 - 9/1/2019
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   04/01/2019 - 12/31/9999
    #   1   03/01/2019   04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 3
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"BID":1,"BUD":"REX","RID":1,"LeaseStatus":2,"DtStart":"4/1/2019","DtStop":"8/31/2019","Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-1a"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #--------------------------------------------------
    # SetRentableLeaseStatus - Case 1c
    # SetStatus  0 (not leased) 4/1/2019 - 9/1/2019
    # Note: EDI in effect, DtStop expressed as "through 8/31/2019"
    # Result needs to be:
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/01/2019 - 12/31/9999
    #   0   04/01/2019 - 09/01/2019
    #   1   03/01/2019   04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 4
    #--------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"BID":1,"BUD":"REX","RID":1,"LeaseStatus":0,"DtStart":"4/1/2019","DtStop":"8/31/2019","Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-1c"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------
    # SetRentableLeaseStatus - Case 1b
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/22/2019 - 12/31/9999
    #   1   09/15/2019 - 09/22/2019
    #   2   09/01/2019 - 09/15/2019
    #   0   04/01/2019 - 09/01/2019
    #   1   03/01/2019   04/01/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 6
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"LeaseStatus":1,"DtStart":"9/15/2019","DtStop":"9/21/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-1b"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------
    # SetRentableLeaseStatus - Case 1d
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/22/2019 - 12/31/9999
    #   1   09/15/2019 - 09/22/2019
    #   2   09/01/2019 - 09/15/2019
    #   0   03/15/2019 - 09/01/2019
    #   1   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 6
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":0,"LeaseStatus":0,"DtStart":"3/15/2019","DtStop":"8/31/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-1d"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------
    # SetRentableLeaseStatus - Case 2b
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/22/2019 - 12/31/9999
    #   1   09/15/2019 - 09/22/2019
    #   2   09/01/2019 - 09/15/2019
    #   1   08/01/2019 - 09/07/2019
    #   0   03/15/2019 - 08/01/2019
    #   1   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"LeaseStatus":1,"DtStart":"8/1/2019","DtStop":"9/6/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-2b"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------
    # SetRentableLeaseStatus - Case 2c
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/22/2019 - 12/31/9999
    #   1   09/15/2019 - 09/22/2019
    #   2   09/01/2019 - 09/15/2019
    #   1   08/07/2019 - 09/07/2019
    #   0   03/15/2019 - 08/07/2019
    #   1   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"LeaseStatus":0,"DtStart":"7/1/2019","DtStop":"8/6/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-2c"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------
    # SetRentableLeaseStatus - Case 2d
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/22/2019 - 12/31/9999
    #   1   09/15/2019 - 09/22/2019
    #   2   09/01/2019 - 09/15/2019
    #   1   08/01/2019 - 09/07/2019
    #   0   03/15/2019 - 08/01/2019
    #   1   03/01/2018   03/15/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"LeaseStatus":1,"DtStart":"8/1/2019","DtStop":"8/10/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-2d"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

    #-------------------------------------------------------
    # SetRentableLeaseStatus - Case 2a
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
    #  Use  DtStart      DtStop
    #  ----------------------------
    #   2   09/22/2019 - 12/31/9999
    #   1   09/15/2019 - 09/22/2019
    #   2   09/07/2019 - 09/15/2019
    #   1   03/01/2018   09/07/2019
    #   0   01/01/2018   03/01/2019
    # Total Records: 7
    #-------------------------------------------------------
    encodeRequest '{"cmd":"save","selected":[],"limit":0,"offset":0,"changes":[{"recid":1,"RLID":13,"LeaseStatus":1,"DtStart":"3/7/2019","DtStop":"8/5/2019","BID":1,"BUD":"REX","RID":1,"Comment":"","CreateBy":211,"LastModBy":211,"w2ui":{}}],"RID":1}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Save-2a"
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0}'
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}${STEP}"  "RentableLeaseStatus-Search"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
