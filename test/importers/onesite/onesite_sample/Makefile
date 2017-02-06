TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

sample:
	@echo "*** Completed in test/importers/onesite/sample ***"

clean:
	rm -f [a-z] onesite.log log *.txt ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g fail
	@echo "*** CLEAN completed in test/importers/onesite/sample ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/sample ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/sample ***"
