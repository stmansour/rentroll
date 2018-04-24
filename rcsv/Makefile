TOP=..
COUNTOL=${TOP}/tools/bashtools/countol.sh

rrlib: *.go
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
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS ***"
	@rm -f fail

clean:
	go clean
	@rm -f fail

package:
	@echo "*** PACKAGE completed in rcsv ***"

secure:
	@rm -f config.json confdev.json confprod.json
