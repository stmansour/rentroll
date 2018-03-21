TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

sample:
	@echo "*** Completed in test/importers/onesite/sample ***"

clean:
	rm -f [a-z] onesite.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail coa.csv
	@echo "*** CLEAN completed in test/importers/onesite/sample ***"

test:
	if [ ! -f coa.csv ]; then cp ../../../jm1/coa.csv .; fi
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/sample ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/sample ***"
