engine:
	@echo "*** COMPLETED in test/engine ***"

clean:
	rm -f *.txt [a-z] [a-z][a-z1-9] log *.log fail
	@echo "*** CLEAN COMPLETE in test/engine ***"

test:
	touch fail
	rm -f *.txt
	./func.sh
	@echo "*** TEST COMPLETE - ALL TESTS PASSED in test/engine ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/engine ***"
