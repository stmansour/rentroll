engine:
	@echo "*** COMPLETED in rsp ***"

clean:
	rm -f *.txt [a-z] log rsplog rentroll.log fail
	@echo "*** CLEAN COMPLETE in rsp ***"

test:
	touch fail
	rm -f *.txt
	./func.sh
	@echo "*** TEST COMPLETE - ALL TESTS PASSED in rsp ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in rsp ***"
