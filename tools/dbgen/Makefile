TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THIS="dbgen"

dbgen: *.go config.json
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@if [ ! -f "bizerr.csv" ]; then ln -s ../../bizlogic/bizerr.csv; fi
	@rm -f fail
	@echo "*** Completed in ${THIS} ***"

clean:
	rm -f rentroll.log log llog ${THIS} fail c*.json
	@echo "*** CLEAN completed in ${THIS} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test:
	@echo "*** TEST completed in ${THIS} ***"

package:
	@echo "*** PACKAGE completed in ${THIS} ***"

secure:
	@rm -f config.json confdev.json confprod.json
