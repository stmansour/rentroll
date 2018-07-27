#!/bin/bash

TESTNAME="Crypto - Encrypt and Decrypt"
TESTSUMMARY="Test encrypt / decrypt functions"
DBGEN=../../tools/dbgen 

source ../share/base.sh

./flow2ra z

genericlogcheck "z"  ""  "Validations"

logcheck

exit 0
