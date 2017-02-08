TEMPCSVSTORE = "../../../../tmp/rentroll/importers/roomkey/temp_CSVs"

roomkey:
	@echo "*** Completed in test/importers/roomkey_sample ***"

clean:
	rm -f [a-z] roomkey.log log *.txt ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g fail
	@echo "*** CLEAN completed in test/importers/roomkey_sample ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/roomkey_sample ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/roomkey_sample ***"
