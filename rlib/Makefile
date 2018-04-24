TOP=..
COUNTOL=${TOP}/tools/bashtools/countol.sh

rrlib: ver.go *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail

ver.go:
	./mkver.sh rlib

test:
	@touch fail
	go test
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS ***"
	@rm -f fail

clean:
	go clean
	@rm -f fail GoAnalyzerError.log ver.go

package:
	@echo "*** PACKAGE completed in rlib ***"

secure:
	@rm -f config.json confdev.json confprod.json
