TOP=../..
COUNTOL=${TOP}/test/share/countol.sh

core: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail

clean:
	go clean
	@rm -f fail
	@echo "*** CLEAN completed in importers/core ***"

test:
	@touch fail
	go test
	@echo "*** TEST completed in importers/core ***"
	@rm -f fail

#man:
#	nroff -man importers/core.1
#	cp importers/core.1 /usr/local/share/man/man1

package: core
	@echo "*** PACKAGE completed in importers/core ***"
