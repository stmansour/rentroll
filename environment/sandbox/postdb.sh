#!/bin/bash
TOP=../..
DEPLOY="/usr/local/accord/bin/deployfile.sh"
DB=sandboxdb.sql
JFROG=$(which jfrog)
DBGEN="${TOP}/tools/dbgen"
MYSQLDUMP="mysqldump --no-defaults"

if [ ! -f ${DEPLOY} ]; then
    DEPLOY=$(which deployfile.sh)
    if [ ! -f ${DEPLOY} ]; then
	echo "cannot find deployfile.sh"
	exit 1
    fi
fi

if [ ! -f ${JFROG} ]; then
    echo "cannot find jfrog"
    exit 1
fi

#--------------------------------------------------------------
# First, create a database and save in sandbox.sql.  We use
# dbgen to create as it always has the latest database updates
#--------------------------------------------------------------
pushd ${DBGEN}
./dbgen -f db25.json
${MYSQLDUMP} rentroll >${DB}
popd
mv ${DBGEN}/${DB} .
gzip ${DB}

#--------------------------------------------------------------
# Now deploy the file to both repos...
#--------------------------------------------------------------
${DEPLOY} ${DB}.gz accord/db
jfrog rt u ${DB}.gz accord/misc/${DB}.gz

