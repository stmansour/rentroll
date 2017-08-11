CAPTURES_DIR = "CAPTURES"

.PHONY: test

captures:
	rm -rf ${CAPTURES_DIR}
	mkdir -p ${CAPTURES_DIR}
	@echo "*** packaging done in test/webclient ***"

package:
	@echo "*** PACKAGE completed in test/webclient ***"

test: captures
	@touch fail
	./functest.sh
	@rm -f fail
	@echo "*** headless UI testing with casperjs done in test/webclient ***"

# test-dev:
# 	@echo "running casper with logging in debug mode..."
# 	casperjs test index.js --log-level=debug --verbose

jshint:
	@touch fail
	jshint --extract=always --show-non-errors .
	@rm -rf fail

clean:
	@rm -f *log *sql *.g [a-z] err.txt fail
	@rm -rf ${CAPTURES_DIR}

all: clean jshint test

link:
	ln -s 
