TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=rrloadcsv

rrloadcsv: *.go config.json
	@touch fail
	${COUNTOL} "go vet"
	${COUNTOL} golint
	go build
	go test
	@rm -f fail

clean:
	rm -f ${THISDIR} ver.go fail conf*.json
	echo "*** CLEAN completed in ${THISDIR} ***"

relink:
	go build
	@echo "*** Relink completed in ${THISDIR} ***"

config.json:
	/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	cp confdev.json config.json

test:
	echo "*** TEST completed in ${THISDIR} ***"

man:
	nroff -man ${THISDIR}.1
	cp ${THISDIR}.1 /usr/local/share/man/man1

package:
	@touch fail
	cp ${THISDIR} config.json ${BINDIR}/
	cp ${THISDIR}.1 ${BINDIR}/man/man1
	echo "*** PACKAGE completed in ${THISDIR} ***"
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
