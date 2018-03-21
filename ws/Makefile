TOP=..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
WEBDOCDIR=${TOP}/../webdoc
WEBDOC=${WEBDOCDIR}/webdoc
WEBCSS=${WEBDOCDIR}/doc.css
THISDIR="ws"

ws: *.go webdoc
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	go test
	go install
	@rm -f fail
	@echo "*** ${THISDIR}: completed build ***"

webdoc:
	rm -rf doc
	${WEBDOC} 
	@cp ${WEBDOCDIR}/doc.css doc/
	@cp ${WEBDOCDIR}/tester.html doc/
	rm -rf ${BINDIR}/doc
	cp -r doc ${BINDIR}/

test:
	@touch fail
	go test
	@rm -f fail
	@echo "*** ${THISDIR}: completed testing - ALL TESTS PASS ***"

clean:
	rm -f ver.go buildno
	go clean
	@rm -f fail GoAnalyzerError.log
	@echo "*** ${THISDIR}: completed clean ***"

package:
	@cp -r doc ${BINDIR}/
	@echo "*** ${THISDIR}: completed package ***"
