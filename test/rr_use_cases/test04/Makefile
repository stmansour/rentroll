test04:
	@echo "*** Completed in test/rr_use_cases/test04 ***"

clean:
	rm -f [a-z] [a-z][a-z1-9] ws* *.sql log *.log *.g ./gold/*.g err.txt fail request conf*.json
	@echo "*** CLEAN completed in test/rr_use_cases/test04 ***"

test: config.json
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/rr_use_cases/test04 ***"
	@rm -f fail

config.json:
	/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	grep -v Timezone confdev.json > config.json

package:
	@echo "*** PACKAGE completed in test/rr_use_cases/test04 ***"

secure:
	@rm -f config.json confdev.json confprod.json
