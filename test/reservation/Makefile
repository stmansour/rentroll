TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="reservation"

reservation:
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f db1.json rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z0-9] [a-z][a-z0-9][a-z0-9] qq? ${THISDIR} fail conf*.json request serverreply svrresponse xxqq xxxe.sql
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
