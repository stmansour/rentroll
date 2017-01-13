TEMPCSVSTORE = "../../../tmp/rentroll/importers/roomkey/temp_CSVs"

roomkey:
	@echo "*** Completed in test/importers/roomkey ***"

clean:
	rm -f [a-z] rentroll.log log *.txt ${TEMPCSVSTORE}/*.csv
	@echo "*** CLEAN completed in test/importers/roomkey ***"

test:
	./functest.sh
	@echo "*** TEST completed in test/importers/roomkey ***"

package:
	@echo "*** PACKAGE completed in test/importers/roomkey ***"
