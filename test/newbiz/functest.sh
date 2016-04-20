#!/bin/bash
pushd ../../db/schema;make newdb;popd
./newbiz

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
echo "NEWBIZ TESTS PASSED"
exit 0