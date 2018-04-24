TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="receipts"

roller:
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -rf rentroll.log log llog err.txt [a-z] [a-z][a-z0-9] [a-z][a-z0-9][a-z0-9] fail conf*.json *.sql request
	@echo "*** CLEAN completed in ${THISDIR} ***"

test: roller
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
