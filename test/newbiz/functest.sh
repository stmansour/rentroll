#!/bin/bash
pushd ../../db/schema;make newdb;popd
./newbiz -b nb.csv -a asmt.csv

cat >xxqq <<EOF
use rentroll;
select BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModBy from business;
EOF
mysql --no-defaults <xxqq >x
UDIFFS=$(diff x x.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 1: PASSED"
else
	echo "PHASE 1: FAILED:  differences are as follows:"
	diff x.gold x
	exit 1
fi
cat >xxqq <<EOF
use rentroll;
select Name,Description,LastModBy from assessmenttypes;
EOF
mysql --no-defaults <xxqq >y
UDIFFS=$(diff y y.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 2: PASSED"
else
	echo "PHASE 2: FAILED:  differences are as follows:"
	diff y.gold y
	exit 1
fi
echo "NEWBIZ TESTS PASSED"
exit 0