#!/bin/bash

TESTNAME="StringLists - Create and Validate"
TESTSUMMARY="Ensures roller messages are created and maintained"


source ../share/base.sh

MYSQL="mysql --no-defaults"
${MYSQL} rentroll < refdb.sql

# First time should create the table
./strlist > z

# Second time pulls from existing table
./strlist > z

genericlogcheck "z"  ""  "Validations"

logcheck

exit 0
