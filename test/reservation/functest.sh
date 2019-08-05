#!/bin/bash
TESTHOME=..
SRCTOP=${TESTHOME}/..
TESTNAME="Reservations"
TESTSUMMARY="Test Reservation search, create, mod"
DBGENDIR=${SRCTOP}/tools/dbgen
CREATENEWDB=0
RRBIN="${SRCTOP}/tmp/rentroll"
CATRML="${SRCTOP}/tools/catrml/catrml"

#SINGLETEST=""  # This runs all the tests

source ${TESTHOME}/share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 10/24/2017"

#------------------------------------------------------------------------------
#  TEST a
#
#  Validate the LeaseStatus insertion service: reservation cmd: save
#
#  Scenario:
#  Walk through the use cases in https://docs.google.com/presentation/d/1v3eEvATppP501MVM6vjv4VoQBDgZo-gq_wPPqUhblV4/edit#slide=id.g4f52e75848_0_0
#
#
#  Expected Results:
#   1.  There should be one rentable available at that time.
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    # Case 1
    #       2/6     2/12
    #        d1      d2                  d1         d2
    #        |       |                   |          |
    #    +---------------+         +----------------------+
    #    |       A       |   ==>   |  A  |    new   |  A1 |
    #    +---------------+         +----------------------+
    #        |         |                 |          |
    #
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22bookResForm%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22BUD%22%3A%22%22%2C%22RTID%22%3A1%2C%22RID%22%3A1%2C%22Nights%22%3A1%2C%22LeaseStatus%22%3A1%2C%22DtStart%22%3A%222%2F6%2F2019%22%2C%22DtStop%22%3A%222%2F16%2F2019%22%2C%22RLID%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}0"  "WebService--SaveRentableLeaseStatus-Rentable(1)"

    # Case 2
    #       2/1     2/6
    #        d1      d2                  d1         d2
    #        |       |                   |          |
    #    +-----------+             +-----------------
    #    |       A   |       ==>   |  A  |    new   |
    #    +-----------+             +-----------------
    #        |       |                   |          |
    #
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22bookResForm%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22BUD%22%3A%22%22%2C%22RTID%22%3A1%2C%22RID%22%3A1%2C%22Nights%22%3A1%2C%22LeaseStatus%22%3A1%2C%22DtStart%22%3A%222%2F1%2F2019%22%2C%22DtStop%22%3A%222%2F6%2F2019%22%2C%22RLID%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}1"  "WebService--SaveRentableLeaseStatus-Rentable(1)"

    # Case 3
    #   2/6     2/10
    #    d1      d2                d1      d2
    #    |       |                 |       |
    #    +-----------+             +-----------------
    #    |       A   |       ==>   | new   |   A    |
    #    +-----------+             +-----------------
    #    |       |                 |       |
    #
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22bookResForm%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22BUD%22%3A%22%22%2C%22RTID%22%3A1%2C%22RID%22%3A1%2C%22Nights%22%3A1%2C%22LeaseStatus%22%3A1%2C%22DtStart%22%3A%222%2F6%2F2019%22%2C%22DtStop%22%3A%222%2F10%2F2019%22%2C%22RLID%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}2"  "WebService--SaveRentableLeaseStatus-Rentable(1)"

    # Case 4 & 8
    #       2/5     2/11
    #        d1      d2                  d1         d2
    #        |       |                   |          |
    #    +---------------+         +----------------------+
    #    | A   | B |  C  |   ==>   |  A  |    new   |  C |
    #    +---------------+         +----------------------+
    #        |       |                   |          |
    #
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22bookResForm%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22BUD%22%3A%22%22%2C%22RTID%22%3A1%2C%22RID%22%3A1%2C%22Nights%22%3A1%2C%22LeaseStatus%22%3A1%2C%22DtStart%22%3A%222%2F5%2F2019%22%2C%22DtStop%22%3A%222%2F11%2F2019%22%2C%22RLID%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}3"  "WebService--SaveRentableLeaseStatus-Rentable(1)"

    # echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222%2F1%2F2019%22%2C%22searchDtStop%22%3A%222%2F28%2F2019%22%2C%22record%22%3A%7B%22BID%22%3A1%2C%22DtStart%22%3A%222%2F1%2F2019%22%2C%22DtStop%22%3A%222%2F28%2F2019%22%2C%22RLID%22%3A0%2C%22RTRID%22%3A0%2C%22RTID%22%3A1%2C%22RID%22%3A0%7D%7D" > request
    # dojsonPOST "http://localhost:8270/v1/reservation/1" "request" "${TFILES}z"  "reservation-searchAvailable"
fi

#------------------------------------------------------------------------------
#  TEST b
#
#  Determine the availability of a rentable type during a particular timeframe
#
#  Scenario:
#  find the availability of a RTID = 5 and BID=1 from 2019-02-14 to 2019-02-15.
#
#  Expected Results:
#   1.  There should be one rentable available at that time.
#------------------------------------------------------------------------------
TFILES="b"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    # search for availability
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"8/1/2018","searchDtStop":"8/31/2018","record":{"BID":1,"DtStart":"2019-02-14","DtStop":"2019-02-15","RLID":0,"RTRID":0,"RTID":5,"RID":0}}' > request
    dojsonPOST "http://localhost:8270/v1/available/1" "request" "${TFILES}0"  "reservation-searchAvailable"
fi

#------------------------------------------------------------------------------
#  TEST c
#
#  Test the ability to search and list existing reservations
#
#  Scenario:
#  Search reservations from 4/20/2018 to 4/23/2018
#
#  Expected Results:
#   1.  There should be six reservations.
#------------------------------------------------------------------------------
TFILES="c"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    #--------------------------------------------------------------------------
    # search for reservations in a time range. There should be six
    #--------------------------------------------------------------------------
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"4/20/2018","searchDtStop":"4/23/2018"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1" "request" "${TFILES}0"  "reservation-searchReservations"

    # get a particular Reservation
    encodeRequest '{"cmd":"get"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/167" "request" "${TFILES}1"  "reservation-getReservation"
fi

#------------------------------------------------------------------------------
#  TEST d
#
#  Test the ability search for available rentables...
#
#  Scenario:
#  see individual calls below
#
#  Expected Results:
#   see individual commands below
#------------------------------------------------------------------------------
TFILES="d"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then

    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    #-------------------------------------------------------------------------
    # search all rentable types for availability, this search should produce
    # no results
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"record":{"recid":0,"BID":1,"BUD":"REX","RTID":0,"Nights":3,"DtStart":"Thu, 20 Jun 2019 07:00:00 GMT","DtStop":"Sun, 23 Jun 2019 07:00:00 GMT"}}' > request
    dojsonPOST "http://localhost:8270/v1/available/1" "request" "${TFILES}0"  "reservation-searchForAvailableRooms"

    #-------------------------------------------------------------------------
    # search all rentable types for availability, this search should produce
    # 4 results
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"record":{"recid":0,"BID":1,"BUD":"REX","RTID":0,"Nights":3,"DtStart":"Tue, 17 Sep 2019 07:00:00 GMT","DtStop":"Fri, 20 Sep 2019 07:00:00 GMT"}}' > request
    dojsonPOST "http://localhost:8270/v1/available/1" "request" "${TFILES}1"  "reservation-searchForAvailableRooms"

    # Create a Reservation
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}2"  "reservation-saveReservation"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
