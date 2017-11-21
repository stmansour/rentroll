TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="dbgen"

dbgen: *.go config.json
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	if [ ! -f "bizerr.csv" ]; then ln -s ../../bizlogic/bizerr.csv; fi
	@rm -f fail
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f rentroll.log log llog ${THISDIR} fail
	@echo "*** CLEAN completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test: acctbal
	@echo "*** TEST completed in ${THISDIR} ***"

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"
