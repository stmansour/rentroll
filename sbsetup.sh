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
DBFILENAME="sandbox"
DBUSER="ec2-user"
IAM=$(whoami)
OS=$(uname)
MYSQL="mysql --no-defaults"
SBLOG="sblog.out"
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
	RRDB=$(grep "Env" config.json | awk '{print $2}' | sed 's/,//')
	# RRDB=$(echo "show databases;" | ${MYSQL} | grep rentroll | wc -l)
	if [ ${RRDB} == "0" ]; then
	    rm -rf ${DBFILENAME}db*  >${SBLOG} 2>&1
	    ${GETFILE} accord/db/${DBFILENAME}db.sql.gz  >${SBLOG} 2>&1
	    gunzip ${DBFILENAME}db.sql  >${SBLOG} 2>&1
		echo "DROP USER IF EXISTS 'ec2-user'@'localhost';CREATE USER 'ec2-user'@'localhost';GRANT ALL PRIVILEGES ON *.* TO 'ec2-user'@'localhost' WITH GRANT OPTION;" | mysql > ${SBLOG} 2>&1
	    echo "DROP DATABASE IF EXISTS ${DATABASENAME}; CREATE DATABASE ${DATABASENAME}; USE ${DATABASENAME};" > restore.sql
	    echo "source ${DBFILENAME}db.sql" >> restore.sql
	    ${MYSQL} < restore.sql  >${SBLOG} 2>&1
	fi

	#---------------------
	# wkhtmltopdf
	#---------------------
	./pdfinstall.sh  >${SBLOG} 2>&1

	#-----------------------------------------------------------------
	#  If no config.json exists, pull the development environment
	#  version and use it.  The Env values mean the following:
	#    0 = development environment
	#    1 = production environment
	#    2 = QA environment
	#-----------------------------------------------------------------
	if [ ! -f ./config.json ]; then
		${GETFILE} accord/db/conflocal.json  >${SBLOG} 2>&1
		mv confdev.json config.json
	fi
}

if [ ! -f "/usr/local/share/man/man1/rentroll.1" ]; then
	./installman.sh >installman.log 2>&1
fi

setupAppNode
