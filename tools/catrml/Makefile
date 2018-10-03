TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THIS="catrml"

catrml: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail
	@echo "*** Completed in ${THIS} ***"

clean:
	rm -f rentroll.log log llog ${THIS} fail c*.json catrml
	@echo "*** CLEAN completed in ${THIS} ***"

relink:
	go build
	@echo "*** Relink completed in ${THIS} ***"

test:
	@echo "*** TEST completed in ${THIS} ***"

package:
	@echo "*** PACKAGE completed in ${THIS} ***"

secure:
	@rm -f config.json confdev.json confprod.json
