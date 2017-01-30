#!/bin/bash
#
# SYNOPSIS:
# 	buildcheck [params]
# params::
#   $1  a string describing the portion of the build being checked.
#       "BUILD" refers to the compilation phase
#       "TEST" refers to the testing phase
#       "PACKAGE" refers to the packaging phase

F=$(find . -name fail | wc -l)

if (( $F > 0 )); then
	echo "${1} failed.  See the following directories"
	find . -name fail
	exit 2
fi
exit 0