#!/bin/bash
pushd tools/schemacmp
./schcmp.sh >w 2>&1
e=$(grep TABLE report.txt | wc -l)
echo "*************************************************************************************"
if [ ${e} -gt 0 ]; then
    echo "****"
    echo "****  NOTICE  ***  Differences Count between local schema and production: ${e}"
    echo "****"
    echo "See ./tools/schemacmp/report.txt for details"
else
    echo "***   Schema on local machine and Roller production are the same."
    echo "***   The build bundle will be publised to the repo."
fi
echo "*************************************************************************************"
rm -f w
popd
