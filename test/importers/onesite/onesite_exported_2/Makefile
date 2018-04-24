TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite_exported_2:
	@echo "*** Completed in test/importers/onesite/onesite_exported_2 ***"

clean:
	rm -f [a-z] onesite.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail
	@echo "*** CLEAN completed in test/importers/onesite/onesite_exported_2 ***"

test:
	if [ ! -f coa.csv ]; then cp ../../../jm1/coa.csv .;cat coa.csv|sed 's/^REX/ISO/' > x; mv x coa.csv; fi
	if [ ! -f ar.csv ]; then cp ../../../jm1/ar.csv .;cat ar.csv|sed 's/^REX/ISO/' > x; mv x ar.csv; fi
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/onesite_exported_2 ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/onesite_exported_2 ***"

secure:
	@rm -f config.json confdev.json confprod.json
