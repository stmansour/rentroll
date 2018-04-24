TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

roomkey: *.go config.json
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
	@echo "*** CLEAN completed in importers/roomkey ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test:
	@touch fail
	go test
	@echo "*** TEST completed in importers/roomkey ***"
	@rm -f fail

package: roomkey
	@echo "*** PACKAGE completed in importers/roomkey ***"

secure:
	@rm -f config.json confdev.json confprod.json
