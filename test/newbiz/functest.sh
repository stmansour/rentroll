#!/bin/bash
pushd ../../db/schema;make newdb;popd
./newbiz -b nb.csv -a asmt.csv -R rt.csv -s specialties.csv >log 

cat >xxqq <<EOF
use rentroll;
select BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModBy from business;
EOF
mysql --no-defaults <xxqq >x
echo "PHASE 1: New Businesses"
if [ ! -f x.gold -o ! -f x ]; then
	echo "Missing file: two files are required for checking this phase: x.gold and x"
	exit 1
fi
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
echo "PHASE 2: Assessment Types"
if [ ! -f y.gold -o ! -f y ]; then
	echo "Missing file: two files are required for checking this phase: y.gold and y"
	exit 1
fi
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
select RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModBy from rentabletypes;
EOF
mysql --no-defaults <xxqq >z
echo "PHASE 3: Rentable Types"
if [ ! -f z.gold -o ! -f z ]; then
	echo "Missing file: two files are required for checking this phase: z.gold and z"
	exit 1
fi
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
echo "PHASE 4: Rentable Market Rates"
if [ ! -f w.gold -o ! -f w ]; then
	echo "Missing file: two files are required for checking this phase: w.gold and w"
	exit 1
fi
UDIFFS=$(diff w w.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 4: PASSED"
else
	echo "PHASE 4: FAILED:  differences are as follows:"
	diff w.gold w
	exit 1
fi

cat >xxqq <<EOF
use rentroll;
select * from rentablespecialtytypes;
EOF
mysql --no-defaults <xxqq >v
echo "PHASE 5: Rentable Specialty Types"
if [ ! -f v.gold -o ! -f v ]; then
	echo "Missing file: two files are required for checking this phase: v.gold and v"
	exit 1
fi
UDIFFS=$(diff v v.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PHASE 5: PASSED"
else
	echo "PHASE 5: FAILED:  differences are as follows:"
	diff v.gold v
	exit 1
fi


echo "FINAL PHASE: Log file check"
echo "PHASE 5: Rentable Specialty Types"
if [ ! -f log.gold -o ! -f log ]; then
	echo "Missing file -- Required files for this check: log.gold and log"
	exit 1
fi
UDIFFS=$(diff log log.gold | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "FINAL PHASE: PASSED"
else
	echo "FINAL PHASE: FAILED:  differences are as follows:"
	diff log.gold log
	exit 1
fi

echo "NEWBIZ TESTS PASSED"
exit 0