TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="workerasm"

workerasm: *.go config.json
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f workerasm rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z][0-9] [a-z][a-z][0-9][0-9] qq? workerasm fail conf*.json *.log request
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test: workerasm
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail


package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
