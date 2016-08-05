newbiz: *.go
	go vet
	golint
	../../mkver.sh
	go build

clean:
	rm -f newbiz [a-z] [a-z][1-9] llog llog.g xxqq ver.go log rentroll.log
	@echo "*** CLEAN completed in newbiz ***"

test: newbiz
	cp ../../tmp/rentroll/conf.json .
	./functest.sh
	@echo "*** TEST completed in newbiz ***"

man:
	nroff -man rrloadcsv.1
	cp rrloadcsv.1 /usr/local/share/man/man1

package:
	cp newbiz ../../tmp/rentroll/rrloadcsv
	cp rrloadcsv.1 ../../tmp/rentroll/man/man1
	@echo "*** PACKAGE completed in newbiz ***"
