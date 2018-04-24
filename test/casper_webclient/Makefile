CAPTURES_DIR = "CAPTURES"
TOP = ./../..
DBGENDIR = ${TOP}/tools/dbgen
# WEBCLIENTTESTDIR = $(shell pwd)

.PHONY: test

captures:
	rm -rf ${CAPTURES_DIR}
	mkdir -p ${CAPTURES_DIR}
	@echo "*** packaging done in test/webclient ***"

package:
	@echo "*** PACKAGE completed in test/webclient ***"

test: captures db4
	@touch fail
	./functest.sh
	@rm -f fail
	@echo "*** headless UI testing with casperjs done in test/webclient ***"

# test-dev:
# 	@echo "running casper with logging in debug mode..."
# 	casperjs test index.js --log-level=debug --verbose

jshint:
	@touch fail
	jshint  --config .jshintrc  *.js
	@rm -rf fail

clean:
	@rm -f *log *sql *.g [a-z] err.txt fail
	@rm -rf ${CAPTURES_DIR}

all: clean jshint test

link:
	ln -s

db4:
	@echo "*** loading database with db4.json ***"
	cd ${DBGENDIR};./dbgen -f db4.json -noauth

secure:
	@rm -f config.json confdev.json confprod.json
