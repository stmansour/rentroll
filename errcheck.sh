#!/bin/bash
ERRCOUNT=$(find . -name fail | wc -l)
if (( ERRCOUNT > 0 )); then
	echo "TESTS HAD ERRORS"
	find . -name fail
	exit 2
else
	echo "ALL TESTS PASSED"
fi
exit 0
