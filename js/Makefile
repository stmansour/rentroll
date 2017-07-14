TOP=..
THISDIR="js"
BINDIR=${TOP}/tmp/rentroll

js:
	webpack -p && webpack --config ./webpackdev.config.js
	@echo "*** Webpack done ***"

clean:
	@echo "*** Clean js ***"

test:
	@touch fail
	@NODE_ENV=test ./../node_modules/mocha/bin/mocha -b tests
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS ***"
	@rm -f fail

package:
	@touch fail
	mkdir -p ${BINDIR}/js
	cp bundle.js ${BINDIR}
	cp bundle.dev.js ${BINDIR}
	@rm -f fail
	@echo "*** PACKAGE completed in js ***"
