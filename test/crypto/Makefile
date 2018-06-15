TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="crypto"

crypto: *.go config.json
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? ${THISDIR} fail conf*.json rex.sql request
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	# ensure that we use UTC as the timezone
	@grep -v Timezone confdev.json >config.json

test: crypto
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
