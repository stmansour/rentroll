TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="tws2"

tws2: *.go config.json bizerr.csv
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f tws2 rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? tws2 fail con*.json bizerr.csv
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

bizerr.csv:
	@ln -s ${TOP}/bizlogic/bizerr.csv

test: tws2 config.json
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail
	
package:
	cp tws2 ${BINDIR}/
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
