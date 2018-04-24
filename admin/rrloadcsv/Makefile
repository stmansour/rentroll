TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh

rrloadcsv: *.go config.json
	touch fail
	${COUNTOL} "go vet"
	${COUNTOL} golint
	go build
	go test
	@rm -f fail

clean:
	rm -f rrloadcsv ver.go fail conf*.json
	echo "*** CLEAN completed in rrloadcsv ***"

config.json:
	/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	cp confdev.json config.json

test:
	echo "*** TEST completed in rrloadcsv ***"

man:
	nroff -man rrloadcsv.1
	cp rrloadcsv.1 /usr/local/share/man/man1

package:
	touch fail
	cp rrloadcsv config.json ${BINDIR}/
	cp rrloadcsv.1 ${BINDIR}/man/man1
	echo "*** PACKAGE completed in rrloadcsv ***"
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
