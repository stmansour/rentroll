TOP=..
COUNTOL=${TOP}/tools/bashtools/countol.sh

rtags: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail

test:
	@touch fail
	go test
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS in rtags ***"
	@rm -f fail

clean:
	go clean

package:
	@echo "*** PACKAGE completed in rtags ***"

secure:
	@rm -f config.json confdev.json confprod.json
