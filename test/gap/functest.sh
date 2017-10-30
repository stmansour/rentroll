#!/bin/bash

TESTNAME="Gap Finder"
TESTSUMMARY="Test the Time Gap locating algorithm"

RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

source ../share/base.sh

./gap > z
genericlogcheck "z"  ""  "DBUpdaterChecks"

logcheck

exit 0
