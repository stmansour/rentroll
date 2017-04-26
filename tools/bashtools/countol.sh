#!/bin/bash
#
# USAGE:
# 	countol [params]
#
# SYNOPSIS: Count Output Lines
#
# DESCRIPTION:
#	Many of the syntax checkers that we run have an exit status of 0 even
#   when the source files they examine contain a number of errors and
#   warnings. This is not ideal for use with make(1) -- it requires that
#   a process it starts must exit with a status other than 0 in order to
#	indicate failure.  In this source tree, we want to fail the build
#	when these checkers detect issues. This routine runs the supplied
#	command and analyzes the output. If the output is empty, then the run
#	is considered successful and the exit status is 0.  If the output
#	is not empty, then it is printed, and the exit status is 2.
# 
# params::
#   $1  a command to execute.

echo "${1}"
${1} > qqtmpqq
RC=$?
x=$(${1} | wc -l)
if [ ${x} -gt 0 ]; then
	cat qqtmpqq
	rm -f qqtmpqq
	exit 2
fi 
if [ ${RC} != 0 ]; then
	echo "Detected exit status = ${RC}"
	exit 2
fi 
rm -f qqtmpqq
exit 0