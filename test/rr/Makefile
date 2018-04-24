TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="rr"

rr: *.go config.json
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	mv rr.sql keepme
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] [a-z][A-Za-z0-9][A-Za-z0-9] \
			qq? ${THISDIR} fail conf*.json request *.sql
	mv keepme rr.sql
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
