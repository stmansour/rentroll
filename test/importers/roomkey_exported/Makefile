TEMPCSVSTORE = "../../../tmp/rentroll/importers/roomkey/temp_CSVs"

roomkey:
	@echo "*** Completed in test/importers/roomkey_exported ***"

clean:
	rm -f [a-z] rentroll.log log *.txt ${TEMPCSVSTORE}/*.csv roomkey.log fail
	@echo "*** CLEAN completed in test/importers/roomkey_exported ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/roomkey_exported ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/roomkey_exported ***"
