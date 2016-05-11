#!/bin/bash
CVSLOAD=../newbiz/newbiz

pushd ../../db/schema;make newdb;popd
${CVSLOAD} -b business.csv -L 3

