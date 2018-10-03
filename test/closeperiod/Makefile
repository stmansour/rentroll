TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="closeperiod"

closeperiod:
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z0-9] [a-z][A-Za-z0-9][A-Za-z0-9] request serverreply
	@echo "*** CLEAN completed in ${THISDIR} ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
