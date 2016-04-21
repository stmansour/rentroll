#!/bin/bash
pushd ../../db/schema;make newdb;popd
./newbiz -b nb.csv -a asmt.csv -R rt.csv

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

cat >xxqq <<EOF
use rentroll;
select RTID,BID,Name,Frequency,Proration,Report,ManageToBudget,LastModBy from rentabletypes;
EOF
mysql --no-defaults <xxqq >z
UDIFFS=$(diff z z.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 3: PASSED"
else
	echo "PHASE 3: FAILED:  differences are as follows:"
	diff z.gold z
	exit 1
fi

cat >xxqq <<EOF
use rentroll;
select RTID,MarketRate,DtStop from rentablemarketrate;
EOF
mysql --no-defaults <xxqq >w
UDIFFS=$(diff w w.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 4: PASSED"
else
	echo "PHASE 4: FAILED:  differences are as follows:"
	diff w.gold w
	exit 1
fi
echo "NEWBIZ TESTS PASSED"
exit 0