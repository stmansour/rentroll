#!/bin/bash
#------------------------------------------------------------------------------
#  prodsetup.sh
#  Set up a production node before rentroll is launched for the first time.
#  The tasks in this file need only be done once, before rentroll is launched.
#
#  Tasks:
#	1. Pull down the production config
#
#  Notes:
#   1. Env values:
#      0 = development
#      1 = production
#      2 = QA
#------------------------------------------------------------------------------

HOST=localhost
PROGNAME="rentroll"
PORT=8270
WATCHDOGOPTS=""
GETFILE="/usr/local/accord/bin/getfile.sh"
RENTROLLHOME="/home/ec2-user/apps/${PROGNAME}"
DATABASENAME="${PROGNAME}"
DBUSER="ec2-user"
IAM=$(whoami)
OS=$(uname)

makeProdNode() {
	${GETFILE} accord/db/confprod.json ; mv confprod.json config.json  >log.out 2>&1
    if [ -f config.json ]; then
        chmod 600 config.json
    fi
}

makeProdNode
./pdfinstall.sh  >log.out 2>&1
if [ ! -f "/usr/local/share/man/man1/rentroll.1" ]; then
	./installman.sh >installman.log 2>&1
fi
