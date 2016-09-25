#!/bin/bash
DOM=$(date +%d)
FULL=0
GET="/usr/local/accord/bin/getfile.sh"
RESTORE="/usr/local/accord/testtools/restoreMySQLdb.sh"
DATABASE="rentroll"
MYSQLOPTS="--no-defaults"
PROGNAME="rrlocalrestore"

##############################################################
#   USAGE
##############################################################
usage() {
    cat <<ZZEOF
ACCORD RentRoll database restore utility
Usage:   ${PROGNAME} [OPTIONS]

OPTIONS:
-d  day-of-month. Default is today's date. Note that if the daily
                  backup has not been performed yet, this would restore
                  last month's data.
-f 	full restore -- includes pictures. Default is data only
-h	print this help message
-n  force --no-defaults onto all mysql commands

Examples:
Command to restore data only
	bash$  ${PROGNAME}

Command to get help
    bash$  ${PROGNAME} -h

ZZEOF
	exit 0
}

##############################################################
#   RESTORE - DATA ONLY
##############################################################
restoreData() {
	# CWD=$(pwd)	
	# if [ ! -e ${CWD}/${BKUPNAM}.gz ]; then
	# 	echo "File does not exist: ${CWD}/${BKUPNAME}.gz"
	# 	exit 1
	# fi

	echo "Extracting data"
	gunzip -f bkup/${BKUPNAME}.gz

	echo "DROP DATABASE IF EXISTS ${DATABASE}; CREATE DATABASE ${DATABASE}; USE ${DATABASE};" > restore.sql
	echo "source bkup/${BKUPNAME}" >> restore.sql
	echo "GRANT ALL PRIVILEGES ON accord TO 'ec2-user'@'localhost' WITH GRANT OPTION;" >> restore.sql
	mysql ${MYSQLOPTS} < restore.sql
	echo "Done."

	echo "Cleaning up..."
	gzip -f bkup/${BKUPNAME}
	echo "Done."
}

##############################################################
#   MAIN ROUTINE
##############################################################
while getopts ":d:fhnN:" o; do
    case "${o}" in
        d)
            DOM=${OPTARG}
            if [ ${DOM} -gt 31 ]; then
            	echo "Largest value for DOM is 31."
            	exit 1
            fi
            if [ ${DOM} -lt 1 ]; then
            	echo "Small value for DOM is 1."
            	exit 1
            fi
            ;;
        h)
			usage
            ;;
        N)
			DATABASE=${OPTARG}
            ;;
        *)
			echo "UNRECOGNIZED OPTION:  ${o}"
            usage
            ;;
    esac
done
shift $((OPTIND-1))

BKUPNAME="${DATABASE}${DOM}.sql"
restoreData
