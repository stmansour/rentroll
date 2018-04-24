all:
	@echo "*** Completed in mrr ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? fail request
	@echo "*** CLEAN completed in mrr ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in mrr ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in mrr ***"

secure:
	@rm -f config.json confdev.json confprod.json
