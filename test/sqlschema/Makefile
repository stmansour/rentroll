TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="sqlschema"

sqlschema:
	@echo "*** Completed in ${THISDIR} ***"

clean:
	@rm -f fixrcpts.sql fixrr.sql dbqqqmods.sql fixrcpts.sql
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	# ensure that we use UTC as the timezone
	@grep -v Timezone confdev.json >config.json

test:
	touch fail
	#./fixprod.sh
	#./fixrcpt.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
