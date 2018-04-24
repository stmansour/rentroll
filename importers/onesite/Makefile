TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

onesite: *.go config.json
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail

clean:
	go clean
	@rm -f fail conf*.json
	@echo "*** CLEAN completed in importers/onesite ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

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

secure:
	@rm -f config.json confdev.json confprod.json
