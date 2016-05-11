newbiz: *.go
	go vet
	golint
	../../mkver.sh
	go build

clean:
	rm -f newbiz [a-z] llog llog.g xxqq ver.go log
	@echo "*** CLEAN completed in newbiz ***"

test: newbiz
	./functest.sh
	@echo "*** TEST completed in newbiz ***"

package:
	cp newbiz ../../tmp/rentroll/rrloadcvs
	@echo "*** PACKAGE completed in newbiz ***"
