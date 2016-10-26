#!/bin/bash
DOM=$(date +%d)
FULL=0
DEPLOY="/usr/local/accord/bin/deployfile.sh"
DATABASE="rentroll"
MYSQLOPTS=""
UNAME=$(uname)

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
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
-f  fname
    Make fname the base file name.  The database backup file will
    be named <fname>.sql.gz .  If this option is not specified
    the file will be named:  rentroll<dayOfMonth>.sql.gz

-h	print this help message

Examples:
Command to backup database to the default filename
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
	mysqldump ${MYSQLOPTS} ${DATABASE} > ${BKUPNAME}
	gzip ${BKUPNAME}
}


##############################################################
#   MAIN ROUTINE
##############################################################
while getopts ":f:hN:" o; do
    case "${o}" in
        f)  USERFNAME=${OPTARG}
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

if [ "${USERFNAME}x" == "x" ]; then
    BKUPNAME=${DATABASE}${DOM}.sql
else
    BKUPNAME=${USERFNAME}.sql
fi

echo "Backing up data on from ${DATABASE} to ./bkup/${BKUPNAME}.gz"
bkupData

mkdir -p bkup
mv ${BKUPNAME}.gz bkup/
ls -l bkup/${BKUPNAME}.gz