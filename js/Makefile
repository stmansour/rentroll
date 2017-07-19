TOP=..
THISDIR="js"
BINDIR=${TOP}/tmp/rentroll

js:
	@touch fail
	webpack -p && webpack --config ./webpackdev.config.js
	@rm -f fail
	@echo "*** Webpack done ***"

clean:
	@rm -f bundle*
	@echo "*** Clean js ***"

test:
	@touch fail
	@cd ..;/usr/local/accord/bin/getfile.sh jenkins-snapshot/rentroll/latest/rrjs.tar.gz; tar xzvf rrjs.tar.gz;cd ..; rm -f rrjs.tar.gz
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
