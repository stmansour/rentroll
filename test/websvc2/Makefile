THISDIR="websvc2"


all:
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][0-9] [a-z0-9][0-9][0-9] qq? fail request w.[a-z]
	@echo "*** CLEAN completed in ${THISDIR} ***"

test:
	@touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

get:
	@touch fail
	../../tools/bashtools/getdb.sh
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
