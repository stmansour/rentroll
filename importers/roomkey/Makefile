TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

roomkey: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail

clean:
	go clean
	rm -f fail
	@echo "*** CLEAN completed in importers/roomkey ***"

test:
	@touch fail
	go test
	@echo "*** TEST completed in importers/roomkey ***"
	@rm -f fail

package: roomkey
	@echo "*** PACKAGE completed in importers/roomkey ***"
