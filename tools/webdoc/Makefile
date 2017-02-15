TOP=../..
COUNTOL=${TOP}/test/share/countol.sh

webdoc: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	@rm -f fail

clean:
	rm -f webdoc ver.go fail
	@echo "*** CLEAN completed in webdoc ***"

test:
	@echo "*** TEST completed in webdoc ***"

man:
	nroff -man webdoc.1
	cp webdoc.1 /usr/local/share/man/man1

package:
	@echo "*** PACKAGE completed in webdoc ***"
