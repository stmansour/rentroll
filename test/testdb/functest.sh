#!/bin/bash
TESTNAME="CCC"
TESTSUMMARY="Setup a database with multiple businesses for testing purposes"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
BUD="OKC"

source ../share/base.sh

pushd ../jm1;./functest.sh;popd
pushd ../ccc;./functest.sh -n -f;popd
pushd ../importers/onesite/onesite_exported_2;./functest.sh -n -f;popd
pushd ../importers/roomkey/roomkey_exported_guest;./functest.sh -n -f;popd

docsvtest "a" "-E cccpets.csv" "Pets"
