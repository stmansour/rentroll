#!/bin/bash

# This script is used to copy tables from one database to another.
# 
# Usage:
#
#   1. Make sure the database you want to copy files FROM is loaded as
#      the current database. 
#   2. Update this script with the names of the tables you want to copy
#   3. Update the TARGETDB with the mysqldump xxx.sql file you want to
#      to copy the tables TO
#   4. Run the script


DATABASE="rentroll"
TARGETDB="empty.sql"
MYSQLDUMP="mysqldump --no-defaults"
MYSQL="mysql --no-defaults"
TMPDIR="xxqqdd"

#=====================================================
#  Put dir/sqlfilename in the list below
#=====================================================
declare -a tables=(
	TaskDescriptor
	TaskListDefinition
)


rm -rf ${TMPDIR}
mkdir ${TMPDIR}
for t in "${tables[@]}"
do
	echo "copy table ${t}"
	${MYSQLDUMP} ${DATABASE} ${t} > "${TMPDIR}/${t}"
done

if [ ! -f ${TARGETDB} ]; then
	echo "file not found: ${t}"
	exit 1
fi

echo "${MYSQL} ${DATABASE} < ${TARGETDB}"
${MYSQL} ${DATABASE} < ${TARGETDB}
for f in "${tables[@]}"
do
	${MYSQL} ${DATABASE} < "${TMPDIR}/${f}"
	echo "paste table ${f}"
done

echo "done!"