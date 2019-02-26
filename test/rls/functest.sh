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
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}0"  "RentableLeaseStatus-SaveWithError"
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
#		3/4/2020 - 12/30/9999
#       must result in an error.
#
#   2.  Next we attempt to save a new record with this date range
#		3/5/2020 - 12/30/9999
#       and this should work.
#------------------------------------------------------------------------------
TFILES="b"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22RLID%22%3A0%2C%22LeaseStatus%22%3A0%2C%22DtStart%22%3A%223%2F4%2F2020%22%2C%22DtStop%22%3A%2212%2F1%2F9999%22%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}0"  "RentableLeaseStatus-SaveWithError2"

    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A0%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1%2C%22RLID%22%3A0%2C%22LeaseStatus%22%3A0%2C%22DtStart%22%3A%223%2F5%2F2020%22%2C%22DtStop%22%3A%2212%2F1%2F9999%22%7D%5D%2C%22RID%22%3A1%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/1" "request" "${TFILES}1"  "RentableLeaseStatus-Save"
fi

#------------------------------------------------------------------------------
#  TEST c
#
#  Validate search service virtual scroll support.
#  discovered in the UI.
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
#   2.  Next we attempt to save a new record with this date range
#		3/5/2020 - 12/30/9999
#       and this should work.
#------------------------------------------------------------------------------
TFILES="c"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}0"  "RentableTypeRefs-Get"

    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A100%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableleasestatus/1/3" "request" "${TFILES}1"  "RentableTypeRefs-GetOffset"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
