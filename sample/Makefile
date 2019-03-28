TOP=./..
BINDIR=${TOP}/tmp/rentroll
THISDIR=sample

sample: 
	@touch fail
	# put work steps here
	@rm -f fail
	@echo "*** BUILD done in ${THISDIR} ***"

clean:
	@touch fail
	# put work steps here
	@rm -f fail
	@echo "*** CLEAN done in ${THISDIR} ***"

package:
	@touch fail
	mkdir -p ${BINDIR}/sample/
	cp *.csv ${BINDIR}/sample/
	@rm -f fail
	@echo "*** PACKAGE completed in clean ***"

secure:
	@rm -f config.json confdev.json confprod.json
