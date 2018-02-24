#!/bin/bash
pushd tools/schemacmp
./schcmp.sh >w 2>&1
e=$(grep TABLE report.txt | wc -l)
if [ ${e} -gt 0 ]; then
    echo "*** NOTICE ***  Differences Count between local schema and production: ${e}"
    echo "See ./tools/schemacmp/report.txt for details"
else
    echo "Schema on local machine and production are the same.  PASS"
fi
rm -f w
popd
