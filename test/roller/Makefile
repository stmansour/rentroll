TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="roller"

roller:
	@echo "*** Completed in ${THISDIR} ***"

clean:
	mv gold savethisgolddir
	rm -rf rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] [a-z][A-Z][A-Z0-9][A-Z0-9] qq? ${THISDIR} fail conf*.json rex.sql request tmpcoa tmpar
	mv savethisgolddir gold
	@echo "*** CLEAN completed in ${THISDIR} ***"

test: roller
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"
