TOP=../..
THISDIR="irex"
COUNTOL=${TOP}/tools/bashtools/countol.sh

irex:
	@touch fail
	@rm -f fail
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f rentroll.log log ll.g llog [a-z] [a-z][a-z] [a-z][0-9] [a-z][0-9][a-z] fail *.json unalloc.sql
	@echo "*** CLEAN completed in ${THISDIR} ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"
