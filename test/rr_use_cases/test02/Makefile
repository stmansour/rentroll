test02:
	@echo "*** Completed in test/rr_use_cases/test02 ***"

clean:
	rm -f [a-z] [a-z][a-z1-9] ws* *.sql log *.log *.g ./gold/*.g err.txt fail request
	@echo "*** CLEAN completed in test/rr_use_cases/test02 ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/rr_use_cases/test02 ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/rr_use_cases/test02 ***"
