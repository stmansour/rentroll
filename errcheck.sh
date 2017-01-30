#!/bin/bash
ERRORS=$(ls test/*/err.txt 2>/dev/null)
ERRCOUNT=$(ls test/*/err.txt 2>/dev/null | wc -l)
if (( ERRCOUNT > 0 )); then
	echo "TESTS HAD ERRORS"
	echo "${ERRORS}"
	exit 2
else
	echo "ALL TESTS PASSED"
fi
exit 0
