TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite_exported_1:
	@echo "*** Completed in test/importers/onesite/onesite_exported_1 ***"

clean:
	rm -f [a-z] onesite.log log *.txt ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g fail
	@echo "*** CLEAN completed in test/importers/onesite/onesite_exported_1 ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/onesite_exported_1 ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/onesite_exported_1 ***"
