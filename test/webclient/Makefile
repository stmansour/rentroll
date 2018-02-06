TOP = ./../..
DBGENDIR = ${TOP}/tools/dbgen

.PHONY: test

package:
	@echo "*** PACKAGE completed in test/webclient ***"

test:
	@touch fail
	./functest.sh
	@rm -f fail
	@echo "*** headless UI testing with cypress.io done in test/webclient ***"

jshint:
	@touch fail
	jshint  --config .jshintrc  cypress/integration/roller_spec.js
	@rm -rf fail

clean:
	@rm -f *log *.g [a-z] err.txt fail

all: clean jshint test

link:
	ln -s

