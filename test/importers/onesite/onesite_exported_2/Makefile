TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite_exported_2:
	@echo "*** Completed in test/importers/onesite/onesite_exported_2 ***"

clean:
	rm -f [a-z] onesite.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail
	@echo "*** CLEAN completed in test/importers/onesite/onesite_exported_2 ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/onesite_exported_2 ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/onesite_exported_2 ***"
