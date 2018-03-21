TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite_exported_1:
	@echo "*** Completed in test/importers/onesite/onesite_exported_1 ***"

clean:
	rm -f [a-z] onesite.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail coa.csv
	@echo "*** CLEAN completed in test/importers/onesite/onesite_exported_1 ***"

test:
	if [ ! -f coa.csv ]; then cp ../../../jm1/coa.csv .;cat coa.csv|sed 's/^REX/ISO/' > x; mv x coa.csv; fi
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/onesite_exported_1 ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/onesite_exported_1 ***"
