upd: *.go config.json
	go build
	@echo "*** Completed in upd ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? upd fail conf*.json
	@echo "*** CLEAN completed in upd ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test: upd
	touch fail
	./functest.sh
	@echo "*** TEST completed in upd ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in upd ***"

secure:
	@rm -f config.json confdev.json confprod.json
