#!/bin/bash

#############################################################################
# removeSecureTarget
#   Description:
#       Pass this routine a Makefile and it will produce a new
#   	Makefile with the secure target and steps removed.
#
#   Params:
#		$1 fully qualified path to the Makefile
#
#	Returns:
#		nothing
#
#############################################################################
function removeSecureTarget {
	sed '/^secure:/,$d' ${1} > qqxxyyzz123;mv qqxxyyzz123 ${1}
}

# -----

for f in $(find . -name Makefile)
do
	echo "Updating: ${f}"
	removeSecureTarget ${f}
done