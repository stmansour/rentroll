TOP = ./../..
ACCORDDB=${TOP}/test/setup/accord.sql
THISDIR="webclient"

.PHONY: test

webclient: config.json accord.sql
	@echo "*** Completed in ${THISDIR} ***"

test: webclient
	@touch fail
	./functest.sh
	@rm -f fail
	@echo "*** headless UI testing with cypress.io done in test/webclient ***"

jshint:
	@touch fail
	jshint  --config .jshintrc  cypress/integration/roller_spec.js
	@rm -rf fail

clean:
	@rm -f *log *.g [a-z] err.txt fail ./accord.sql ./conf*.json

all: clean jshint test

link:
	ln -s

config.json:
	@/usr/local/accord/bin/getfile.sh accord/webclient/confdev.json
	@cp confdev.json config.json

accord.sql:
	@cp ${ACCORDDB} .

package:
	@echo "*** PACKAGE completed in test/webclient ***"
