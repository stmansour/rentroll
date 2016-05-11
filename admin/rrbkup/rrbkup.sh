#!/bin/bash
DOM=$(date +%d)
FULL=0
DEPLOY="/usr/local/accord/bin/deployfile.sh"
DATABASE="rentroll"
MYSQLOPTS=""
UNAME=$(uname)

if [ ${UNAME} == "Darwin" ]; then
	MYSQLOPTS="--no-defaults"
fi

##############################################################
#   USAGE
##############################################################
usage() {
    cat << ZZEOF
ACCORD RentRoll database backup utility
Usage:   rrbkup.sh [OPTIONS]

OPTIONS:
-h	print this help message

Examples:
Command to backup database
	bash$  ./rrbkup

Command to get help
    bash$  ./rrbkup -h

ZZEOF
	exit 0
}

##############################################################
#   BACKUP - DATA ONLY
##############################################################
bkupData() {
	mysqldump ${MYSQLOPTS} ${DATABASE} > ${DATABASE}db.sql
	gzip ${DATABASE}db.sql
	if [ ${UNAME} != "Darwin" ]; then
		${DEPLOY} ${DATABASE}db.sql.gz ${DATABASE}/db/${DOM}
		rm -f ${DATABASE}db.sql.gz
	fi
}


##############################################################
#   MAIN ROUTINE
##############################################################
while getopts ":hN:" o; do
    case "${o}" in
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

if [ ${FULL} -eq 0 ]; then
	echo "Backing up data on database ${DATABASE}..."
	bkupData
fi

echo
echo "Done."