#!/bin/bash

TESTNAME="Crypto - Encrypt and Decrypt"
TESTSUMMARY="Test encrypt / decrypt functions"


source ../share/base.sh

./crypto -noauth > z

genericlogcheck "z"  ""  "Validations"

logcheck

exit 0
