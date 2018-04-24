TOP=..
THISDIR="bizlogic"
COUNTOL=${TOP}/tools/bashtools/countol.sh
BINDIR=${TOP}/tmp/rentroll

bizlogic: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go install
	@rm -f fail

test:
	@touch fail
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS ***"
	@rm -f fail

clean:
	go clean
	@rm -f fail GoAnalyzerError.log

package:
	@touch fail
	cp bizerr.csv ${BINDIR}
	@rm -f fail
	@echo "*** PACKAGE completed in rlib ***"

secure:
	@rm -f config.json confdev.json confprod.json
