upd: *.go
	go build
	@echo "*** Completed in upd ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? upd fail
	@echo "*** CLEAN completed in upd ***"

test: upd
	touch fail
	./functest.sh
	@echo "*** TEST completed in upd ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in upd ***"
