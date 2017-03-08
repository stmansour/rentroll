TEMPCSVSTORE = "../../../../tmp/rentroll/importers/roomkey/temp_CSVs"

roomkey:
	@echo "*** Completed in test/importers/roomkey_exported_guest ***"

clean:
	rm -f [a-z] roomkey.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail
	@echo "*** CLEAN completed in test/importers/roomkey_exported_guest ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/roomkey_exported_guest ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/roomkey_exported_guest ***"
