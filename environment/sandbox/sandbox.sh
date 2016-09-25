#!/bin/bash
# This creates a RentRoll sandbox environment
# It consists of the releases for both RentRoll and Phonebook

UPORT=8251
ENV_DESCR="rrsb.json"
SYS_TEST_DIR=$(pwd)

#---------------------------------------------------------------------
#  Find accord bin...
#---------------------------------------------------------------------
if [ -d /usr/local/accord/bin ]; then
	ACCORDBIN=/usr/local/accord/bin
	TOOLS_DIR="/usr/local/accord/testtools"
elif [ -d /c/Accord/bin ]; then
	ACCORDBIN=/c/Accord/bin
	TOOLS_DIR="/c/Accord/testtools"
else
	echo "*** ERROR: Required directory /usr/local/accord/bin or /c/Accord/bin does not exist."
	echo "           Please repair installation and try again."
	exit 2
fi


rm -f qm* *.log *.out
${ACCORDBIN}/uhura -p ${UPORT} -d -k -e ${SYS_TEST_DIR}/${ENV_DESCR} >uhura.out 2>&1 
