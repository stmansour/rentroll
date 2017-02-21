TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

onesite: *.go
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
	@echo "*** CLEAN completed in importers/onesite ***"

test:
	@touch fail
	go test
	@echo "*** TEST completed in importers/onesite ***"
	@rm -f fail

#man:
#	nroff -man importers/onesite.1
#	cp importers/onesite.1 /usr/local/share/man/man1

package: onesite
	@echo "*** PACKAGE completed in importers/onesite ***"
