TEMPCSVSTORE = "../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite:
	@echo "*** Completed in test/importers/onesite ***"

clean:
	rm -f [a-z] onesite.log log *.txt ${TEMPCSVSTORE}/*.csv fail
	@echo "*** CLEAN completed in test/importers/onesite ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite ***"
