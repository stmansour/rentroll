test03:
	@echo "*** Completed in test/rr_use_cases/test03 ***"

clean:
	rm -f [a-z] [a-z][a-z1-9] ws* *.sql log *.log *.g ./gold/*.g err.txt fail request
	@echo "*** CLEAN completed in test/rr_use_cases/test03 ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/rr_use_cases/test03 ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/rr_use_cases/test03 ***"

secure:
	@rm -f config.json confdev.json confprod.json
