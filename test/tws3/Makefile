TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="tws3"

tws3: *.go config.json bizerr.csv
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f tws3 rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? tws3 fail con*.json bizerr.csv
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

bizerr.csv:
	@ln -s ${TOP}/bizlogic/bizerr.csv

test: tws3
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail
	
package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
