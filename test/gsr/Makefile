TOP=../..
THISDIR="gsr"
COUNTOL=${TOP}/tools/bashtools/countol.sh

gsr: *.go config.json config.json
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail
	@echo "*** Completed in ${THISDIR} ***"

unalloc:
	./unalloc.sh

more: unalloc
	./gsr -funds more -db

less: unalloc
	./gsr -funds less -db

eq: unalloc
	./gsr -funds eq -db

config.json:
	/usr/local/accord/bin/getfile.sh accord/db/confdev.json ; cp confdev.json config.json

clean:
	rm -f rentroll.log log ll.g llog [a-z] [a-z][a-z] [a-z][0-9] [a-z][0-9][a-z] fail *.json unalloc.sql gsr
	@echo "*** CLEAN completed in ${THISDIR} ***"

test: gsr
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"
