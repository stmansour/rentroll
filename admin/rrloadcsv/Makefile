TOP=../..
COUNTOL=${TOP}/test/share/countol.sh

rrloadcsv: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	${TOP}/mkver.sh
	go build
	go test
	@rm -f fail

clean:
	rm -f rrloadcsv ver.go fail
	@echo "*** CLEAN completed in rrloadcsv ***"

test:
	@echo "*** TEST completed in rrloadcsv ***"

man:
	nroff -man rrloadcsv.1
	cp rrloadcsv.1 /usr/local/share/man/man1

package:
	touch fail
	cp rrloadcsv ${TOP}/tmp/rentroll/rrloadcsv
	cp rrloadcsv.1 ${TOP}/tmp/rentroll/man/man1
	@echo "*** PACKAGE completed in rrloadcsv ***"
	rm -f fail
