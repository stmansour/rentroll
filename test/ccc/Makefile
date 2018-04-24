all:
	@echo "*** Completed in ccc ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][1-9] fail
	@echo "*** CLEAN completed in ccc ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in ccc ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ccc ***"

secure:
	@rm -f config.json confdev.json confprod.json
