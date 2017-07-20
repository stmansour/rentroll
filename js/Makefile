TOP=..
THISDIR="js"
BINDIR=${TOP}/tmp/rentroll

js:
	@touch fail
	grunt
	@rm -f fail
	@echo "*** Webpack done ***"

clean:
	@touch fail
	grunt clean
	@rm -f fail
	@echo "*** Clean js ***"

test:
	@touch fail
	mkdir -p ./../temp
	grunt qunit
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS ***"
	@rm -f fail

package:
	@touch fail
	mkdir -p ${BINDIR}/js
	cp bundle.js ${BINDIR}/js/
	cp bundle.min.js ${BINDIR}/js/
	@rm -f fail
	@echo "*** PACKAGE completed in js ***"
