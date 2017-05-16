TOP=..
THISDIR="bizlogic"
COUNTOL=${TOP}/tools/bashtools/countol.sh

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
	rm -f fail GoAnalyzerError.log

package:
	@echo "*** PACKAGE completed in rlib ***"
