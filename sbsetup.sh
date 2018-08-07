#!/bin/bash
#------------------------------------------------------------------------------
#  sbsetup.sh
#  Set up a sandbox node before rentroll is launched for the first time.
#  The tasks in this file need only be done once, before rentroll is launched.
#
#  Tasks:
#	1. Pull down the development config file for sandboxes
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

#--------------------------------------------------------------
#  For QA, Sandbox, and Production nodes, go through the
#  laundry list of details...
#  1. Set up permissions for the database on QA and Sandbox nodes
#  2. Install a database with some data for testing
#  3. For PDF printing, install wkhtmltopdf
#--------------------------------------------------------------
setupAppNode() {
	#---------------------
	# database
	#---------------------
	RRDB=$(echo "show databases;" | mysql | grep rentroll | wc -l)
	if [ ${RRDB} -gt "0" ]; then
	    rm -rf ${DATABASENAME}db*  >log.out 2>&1
	    ${GETFILE} accord/db/${DATABASENAME}db.sql.gz  >log.out 2>&1
	    gunzip ${DATABASENAME}db.sql  >log.out 2>&1
	    echo "DROP DATABASE IF EXISTS ${DATABASENAME}; CREATE DATABASE ${DATABASENAME}; USE ${DATABASENAME};" > restore.sql
	    echo "source ${DATABASENAME}db.sql" >> restore.sql
	    echo "GRANT ALL PRIVILEGES ON ${DATABASENAME} TO 'ec2-user'@'localhost' WITH GRANT OPTION;" >> restore.sql
	    mysql ${MYSQLOPTS} < restore.sql  >log.out 2>&1
	fi

	#---------------------
	# wkhtmltopdf
	#---------------------
	./pdfinstall.sh  >log.out 2>&1

	#-----------------------------------------------------------------
	#  If no config.json exists, pull the development environment
	#  version and use it.  The Env values mean the following:
	#    0 = development environment
	#    1 = production environment
	#    2 = QA environment
	#-----------------------------------------------------------------
	if [ ! -f ./config.json ]; then
		${GETFILE} accord/db/confdev.json  >log.out 2>&1
		mv confdev.json config.json
	fi
}

if [ ! -f "/usr/local/share/man/man1/rentroll.1" ]; then
	./installman.sh >installman.log 2>&1
fi

setupAppNode
