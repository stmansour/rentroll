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

#------------------------------------------------------------------------------
#  Simple routine to parse the server response for specific variables.
#  Update as needed for your tests
#------------------------------------------------------------------------------
parseServerReply() {
    RLID=$(cat serverreply | sed 's/^.*:\([0-9][0-9]*\)}/\1/')
    # echo "RLID = ${RLID}"
}

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
RENTROLLSERVERNOW="-testDtNow 5/1/2019"

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

    #-------------------------------------------------------------------------
    # Create a Reservation
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}2"  "reservation-saveReservation"

    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}3"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Delete the reservation and keep the deposit on account
    #-------------------------------------------------------------------------
    PART1='{"cmd":"delete","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"RAID":827,"TCID":0,"Amount":250,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FLAGS":0,"FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}4"  "reservation-cancelReservation"

    #-------------------------------------------------------------------------
    # Create a Reservation
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}5"  "reservation-saveReservation"

    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}6"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Delete the reservation and refund the deposit
    #-------------------------------------------------------------------------
    PART1='{"cmd":"delete","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"RAID":827,"TCID":0,"Amount":250,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FLAGS":1,"FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}7"  "reservation-cancelReservation"

    #-------------------------------------------------------------------------
    # Create a Reservation
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}8"  "reservation-saveReservation"

    parseServerReply
    echo "RLID = ${RLID}"

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}9"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Delete the reservation and book the deposit as income (payor forfeits deposit)
    #-------------------------------------------------------------------------
    PART1='{"cmd":"delete","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"RAID":827,"TCID":0,"Amount":250,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FLAGS":2,"FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}10"  "reservation-cancelReservation"
fi

#------------------------------------------------------------------------------
#  TEST e
#
#  Create a reservation then update the number of unspecified adults. This
#  should do a simple update to the RentalAgreement -- it should not cancel
#  anything
#
#  Scenario:
#  see individual calls below
#
#  Expected Results:
#   see individual commands below
#------------------------------------------------------------------------------
TFILES="e"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    stopRentRollServer
    mysql --no-defaults rentroll < xd.sql
    startRentRollServer

    #-------------------------------------------------------------------------
    # Create a Reservation
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}0"  "reservation-saveReservation"

    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}1"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Update the reservation and keep the deposit on account. In this case,
    # we're just going to add an unspecified adult to the reservation
    #-------------------------------------------------------------------------
    PART1='{"cmd":"save","record":{"recid":0,"rdBID":0,"TCID":834,"RAID":827,"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"Rate":0,"DBAmount":85,"Amount":85,"DBDeposit":10,"Deposit":10,"DepASMID":844,"DBDepASMID":844,"Discount":0,"LeaseStatus":2,"Nights":3,"UnspecifiedAdults":7,"UnspecifiedChildren":0,"RentableName":"Rentable007","IsCompany":false,"CompanyName":"","FirstName":"Billy Bob","MiddleName":"","LastName":"Thorton","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","FLAGS":0,"CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"JZD4AQVGLTVK2P4YOGAH","Comment":"","BUD":""}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}2"  "reservation-updateReservation"
    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}3"  "reservation-saveReservation" "ConfirmationCode"
fi

#------------------------------------------------------------------------------
#  TEST f
#
#  Decrease the deposit on a reservation, then update only the deposit amount
#
#  Scenario:
#  see individual calls below
#
#  Expected Results:
#   see individual commands below
#------------------------------------------------------------------------------
TFILES="f"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    stopRentRollServer
    mysql --no-defaults rentroll < xd.sql
    startRentRollServer

    #-------------------------------------------------------------------------
    # Create a Reservation.  Set initial deposit to $25
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":25,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}0"  "reservation-saveReservation"

    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}1"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Update the reservation: change the deposit to $10, keep overpayment on account
    # The result should be no change to the rental agreement, but the assessment
    # for the deposit should be reversed and replaced with a new one.
    #-------------------------------------------------------------------------
    PART1='{"cmd":"save","record":{"recid":0,"rdBID":0,"TCID":834,"RAID":827,"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"Rate":0,"DBAmount":85,"Amount":85,"DBDeposit":10,"Deposit":10,"DepASMID":844,"DBDepASMID":844,"Discount":0,"LeaseStatus":2,"Nights":3,"UnspecifiedAdults":7,"UnspecifiedChildren":0,"RentableName":"Rentable007","IsCompany":false,"CompanyName":"","FirstName":"Billy Bob","MiddleName":"","LastName":"Thorton","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","FLAGS":0,"CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"JZD4AQVGLTVK2P4YOGAH","Comment":"","BUD":""}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}2"  "reservation-updateReservation"
    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the ASMID which should be 846
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}3"  "reservation-saveReservation" "ConfirmationCode"

fi

#------------------------------------------------------------------------------
#  TEST g
#
#  Increase the deposit. Make sure a charge for the delta is made
#
#  Scenario:
#  see individual calls below
#
#  Expected Results:
#   see individual commands below
#------------------------------------------------------------------------------
TFILES="g"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    stopRentRollServer
    mysql --no-defaults rentroll < xd.sql
    startRentRollServer

    #-------------------------------------------------------------------------
    # Create a Reservation.  Set initial deposit to $10
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","Nights":3,"RLID":0,"RTRID":0,"rdRTID":4,"RID":7,"RAID":0,"TCID":0,"Amounmt":0,"Deposit":10,"LeaseStatus":2,"RentableName":"Rentable007","FirstName":"Billy Bob","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"","Comment":"","RTID":4,"Amount":25}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}0"  "reservation-saveReservation"

    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}1"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Update the reservation: change the deposit to $25
    # The result should be no change to the rental agreement, but a new
    # assessment for the difference in the deposit should be added. And the
    # credit card should be billed $15
    #-------------------------------------------------------------------------
    PART1='{"cmd":"save","record":{"recid":0,"rdBID":0,"TCID":834,"RAID":827,"DtStart":"Tue, 18 Jun 2019 00:00:00 GMT","DtStop":"Fri, 21 Jun 2019 00:00:00 GMT","RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"Rate":0,"DBAmount":85,"Amount":85,"DBDeposit":10,"Deposit":25,"DepASMID":844,"DBDepASMID":844,"Discount":0,"LeaseStatus":2,"Nights":3,"UnspecifiedAdults":7,"UnspecifiedChildren":0,"RentableName":"Rentable007","IsCompany":false,"CompanyName":"","FirstName":"Billy Bob","MiddleName":"","LastName":"Thorton","Email":"bbt@boozer.com","Phone":"1234567890","Street":"123 Elm","City":"Murfreesboro","Country":"","State":"AK","PostalCode":"12345","FLAGS":0,"CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"JZD4AQVGLTVK2P4YOGAH","Comment":"","BUD":""}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}2"  "reservation-updateReservation"
    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the ASMID which should be 846
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}3"  "reservation-saveReservation" "ConfirmationCode"

fi
#------------------------------------------------------------------------------
#  TEST h
#
#  Change the Dates associated with a rexervation. In this case the Rental
#  Agreement will change and the Deposit Assessment will be reversed and a
#  new Assessment will replace it
#
#  Scenario:
#  see individual calls below
#
#  Expected Results:
#   see individual commands below
#------------------------------------------------------------------------------
TFILES="h"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    stopRentRollServer
    mysql --no-defaults rentroll < x${TFILES}.sql
    startRentRollServer

    #-------------------------------------------------------------------------
    # Create a Reservation.  Set initial deposit to $10
    #-------------------------------------------------------------------------
    encodeRequest '{"cmd":"save","record":{"rdBID":1,"BUD":{"id":"REX","text":"REX"},"DtStart":"Wed, 04 Sep 2019 00:00:00 GMT","DtStop":"Fri, 06 Sep 2019 00:00:00 GMT","Nights":2,"RLID":0,"RTRID":0,"rdRTID":3,"RID":6,"RAID":0,"TCID":1089,"Amount":250,"Deposit":10,"DepASMID":0,"LeaseStatus":2,"RentableName":"Rentable006","FirstName":"William","UnspecifiedAdults":0,"UnspecifiedChildren":0,"LastName":"Thorton","IsCompany":false,"CompanyName":"Hicks R Us","Email":"bb@backwoods.com","Phone":"123-890-7654","Street":"123 Hayseed St","City":"Broken Pine","State":"AK","PostalCode":"64549","CCName":"BILLYBOB THORTON","CCType":"VISA","CCNumber":"1234567890","CCExpMonth":"2","CCExpYear":"2022","ConfirmationCode":"","Comment":"","PGName":[{"TCID":1089,"BID":1,"FirstName":"William","MiddleName":"Robert","LastName":"Thorton","CompanyName":"Hicks R Us","IsCompany":false,"PrimaryEmail":"bb@backwoods.com","SecondaryEmail":"","WorkPhone":"123-456-7890","CellPhone":"123-890-7654","Address":"123 Hayseed St","Address2":"","City":"Broken Pine","State":"AK","PostalCode":"64549","recid":64}],"RTID":3}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/0" "request" "${TFILES}0"  "reservation-saveReservation"

    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the RAID
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}1"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # At this point, the reservation has the following associated records:
    # RAID:  837
    # ASMID: 847
    #
    # Update the reservation:
    #   1. change the deposit to $25
    #   2. change dates from Sep 4 - 6 to Sep 7 2019 - Sep 9 2019
    #   3. change RID from 6 to 7
    #   4. change RTID to 8
    #
    # This should result in a new RentalAgreement
    # The old RentableLeaseStatus for RID 6 should be set to NOTLEASED
    # ASMID 847 should be reversed
    # A new assessment must be created for $25 and with a comment that the
    #       cc was charged an additional $15
    #
    #-------------------------------------------------------------------------
    echo "Updating RLID = ${RLID}"
    PART1='{"cmd":"save","record":{"recid":0,"rdBID":0,"TCID":1089,"RAID":837,"DtStart":"Sat, 07 Sep 2019 00:00:00 GMT","DtStop":"Mon, 09 Sep 2019 00:00:00 GMT","RLID":'
    PART2=',"RTRID":0,"rdRTID":4,"RID":7,"Rate":0,"DBAmount":75,"Amount":75,"DBDeposit":10,"Deposit":25,"DepASMID":847,"DBDepASMID":847,"Discount":0,"LeaseStatus":2,"Nights":2,"UnspecifiedAdults":0,"UnspecifiedChildren":0,"RentableName":"Rentable007","IsCompany":false,"CompanyName":"Hicks R Us","FirstName":"William","MiddleName":"","LastName":"Thorton","Email":"bb@backwoods.com","Phone":"123-890-7654","Street":"123 Hayseed St","City":"Broken Pine","Country":"","State":"AK","PostalCode":"64549","FLAGS":0,"CCName":"","CCType":"","CCNumber":"","CCExpMonth":"","CCExpYear":"","ConfirmationCode":"JZT17T1CHB1MODFXMEAQ","Comment":"","BUD":"","RTID":4}}'
    CMD="${PART1}${RLID}${PART2}"

    encodeRequest "${CMD}" > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}2"  "reservation-updateReservation"
    parseServerReply

    #-------------------------------------------------------------------------
    # Read this Reservation back to determine the ASMID which should be 846
    # Note the optional 5th parameter on dojsonPOST.  This instructs
    #     dojsonPOST to ignore the value for the supplied property name.
    #     Since every CONFCODE is different, we do not compare to known good.
    #-------------------------------------------------------------------------
    encodeRequest 'request={"cmd":"get","recid":0,"name":"resUpdateForm"}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/${RLID}" "request" "${TFILES}3"  "reservation-saveReservation" "ConfirmationCode"

    #-------------------------------------------------------------------------
    # Make sure that the new deposit assessment is correct and that the
    # credit card charges are correct.
    #-------------------------------------------------------------------------
    ASMID=$(python -m json.tool serverreply | grep DepASMID | grep -v DB | sed 's/^.*: *\([0-9][0-9][0-9]\).*/\1/')
    echo "ASMID = ${ASMID}"
    encodeRequest '{"cmd":"get","record":{}}'
    dojsonPOST "http://localhost:8270/v1/asm/1/${ASMID}" "request" "${TFILES}4"  "readDepositAssessment"
fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
