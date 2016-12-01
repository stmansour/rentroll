rrloadcsv: *.go
	go vet
	golint
	../../mkver.sh
	go build

clean:
	rm -f rrloadcsv ver.go
	@echo "*** CLEAN completed in rrloadcsv ***"

test:
	@echo "*** TEST completed in rrloadcsv ***"

man:
	nroff -man rrloadcsv.1
	cp rrloadcsv.1 /usr/local/share/man/man1

package:
	cp rrloadcsv ../../tmp/rentroll/rrloadcsv
	cp rrloadcsv.1 ../../tmp/rentroll/man/man1
	@echo "*** PACKAGE completed in rrloadcsv ***"
