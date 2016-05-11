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


SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"


##############################################################
#   USAGE
##############################################################
usage() {
    cat << ZZEOF
ACCORD RentRoll new database utility
Usage:   rrnewdb.sh [OPTIONS]

Note: 

OPTIONS:
-h	print this help message

Examples:
Command to destroy current database and create a new blank database
	bash$  ./rrnewdb

Command to get help
    bash$  ./rrnewdb -h

ZZEOF
	exit 0
}

##############################################################
#   CREATE THE NEW DATABASE
##############################################################
newdb() {
	pushd ${DIR}
	mysql ${MYSQLOPTS} <schema.sql
	popd
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

echo "Creating new database ${DATABASE}..."
newdb
echo "Done."