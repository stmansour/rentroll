#!/bin/bash

TESTNAME="Assessment Instance Generator Worker"
TESTSUMMARY="Test the generation of asm instances"

CREATENEWDB=0
RRDATERANGE="-j 2018-03-01 -k 2018-04-01"
BUD="REX"

echo "Create new database..."
mysql --no-defaults rentroll < rex.sql

source ../share/base.sh

WLOG=workasm.log

#------------------------------------------------------------------------------
#  FIX
#  Add assessment instances for May 2018
#
#------------------------------------------------------------------------------
RRDATERANGE="-j 2018-05-01 -k 2018-06-01"
./workerasm -dt "5/1/2018" > ${WLOG}


exit 0
