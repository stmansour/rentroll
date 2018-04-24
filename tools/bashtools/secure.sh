#!/bin/bash

# This script recurses through the current directory and all directories
# beneath it and adds a 'secure' target to every Makefile if one does not
# already exist.  The default actions under the target are listed in the
# here-document section below and can be adjusted as needed.

#############################################################################
# secureIt
#   Description:
#       Pass the first param is the Makefile filename, fully qualified
#   	Add a default secure target to the Makefile
#
#   Params:
#		$1 fully qualified path to the Makefile
#
#	Returns:
#		nothing
#
#############################################################################
function secureIt {
	COUNT=$(egrep "^secure:" ${1} | wc -l)
	if [ ${COUNT} -eq 0 ]; then
		echo "secure:" >> ${1}

		#-------------------------------------------------
		# Handle the case where the Makefile has DIRS...
		#-------------------------------------------------
		COUNT=$(egrep "^DIRS[ \t]*=" ${f} | wc -l)
		if [ ${COUNT} -gt 0 ]; then
			cat >> ${1} << FEOF
	for dir in \$(DIRS); do make -C \$\${dir} secure;done
FEOF
		fi

		#-------------------------------------------------
		# and now the code for all Makefiles...
		#-------------------------------------------------
		cat >> ${1} << EOF
	@rm -f config.json confdev.json confprod.json
EOF
	fi

}

#-------------------------------------------------------

for f in $(find . -name Makefile)
do
	echo "Updating: ${f}"
	secureIt ${f}
done