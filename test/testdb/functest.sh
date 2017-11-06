#!/bin/bash
TESTNAME="TestDB"
TESTSUMMARY="Setup a database with multiple businesses for testing purposes"

RRDATERANGE="-j 2016-01-01 -k 2016-02-01"
BUD="OKC"

source ../share/base.sh

if [ -f rex.sql ]; then
	mysql --no-defaults rentroll < rex.sql
else
	pushd ../jm1;./functest.sh ;popd
fi

pushd ../ccc;./functest.sh  -n -f;popd
pushd ../importers/onesite/onesite_exported_2;./functest.sh -n -f; ../../../../tmp/rentroll/rrloadcsv -c coa.csv -ar ar.csv; popd
pushd ../importers/roomkey/roomkey_exported_guest;./functest.sh -n -f; ../../../../tmp/rentroll/rrloadcsv -c coa.csv -ar ar.csv; popd

# Import some other things that several of the Businesses are missing

docsvtest "a" "-E cccpets.csv" "Pets"
docsvtest "b" "-P pmt.csv" "PaymentTypes"

cat >xyz <<EOF
UPDATE RentalAgreementPayors SET DtStop='9999-10-31' WHERE RAPID='21';
UPDATE RentalAgreement SET AgreementStop='9999-10-31', PossessionStop='9999-10-31', RentStop='9999-10-31' WHERE RAID='16';
UPDATE RentalAgreementPayors SET DtStop='9999-10-31' WHERE RAPID='21';
EOF

mysql --no-defaults rentroll < xyz
