TEMPCSVSTORE = "../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite:
	@echo "*** Completed in test/importers/onesite ***"

clean:
	rm -f [a-z] rentroll.log log *.txt ${TEMPCSVSTORE}/*.csv
	@echo "*** CLEAN completed in test/importers/onesite ***"

test:
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite ***"

package:
	@echo "*** PACKAGE completed in test/importers/onesite ***"
