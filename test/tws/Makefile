TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="tws"

tws: *.go config.json
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f tws rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? tws fail conf*.json
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test: tws tws.sql
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

tws.sql:
	cp ${TOP}/../tws/db/schema/tws.sql .
	
package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"
