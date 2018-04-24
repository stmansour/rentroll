TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR="worker"

worker: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail
	@echo "*** ${THISDIR}: completed build ***"

test:
	@touch fail
	go test
	@rm -f fail
	@echo "*** ${THISDIR}: completed testing - ALL TESTS PASS ***"

clean:
	rm -f ver.go
	go clean
	@rm -f fail GoAnalyzerError.log
	@echo "*** ${THISDIR}: completed clean ***"

package:
	@cp -r doc ${BINDIR}/
	@echo "*** ${THISDIR}: completed package ***"

secure:
	@rm -f config.json confdev.json confprod.json
