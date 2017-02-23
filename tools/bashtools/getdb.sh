#!/bin/bash
GETFILE="/usr/local/accord/bin/getfile.sh"
DATABASENAME="rentroll"
DBUSER="ec2-user"

rm -rf ${DATABASENAME}db*
${GETFILE} accord/db/${DATABASENAME}db.sql.gz
gunzip ${DATABASENAME}db.sql
echo "DROP DATABASE IF EXISTS ${DATABASENAME}; CREATE DATABASE ${DATABASENAME}; USE ${DATABASENAME};" > restore.sql
echo "source ${DATABASENAME}db.sql" >> restore.sql
echo "GRANT ALL PRIVILEGES ON ${DATABASENAME} TO 'ec2-user'@'localhost' WITH GRANT OPTION;" >> restore.sql
mysql --no-defaults < restore.sql
echo "Done."

