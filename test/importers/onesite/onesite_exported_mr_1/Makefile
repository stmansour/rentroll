TEMPCSVSTORE = "../../../../tmp/rentroll/importers/onesite/temp_CSVs"

onesite_exported_1:
	@echo "*** Completed in test/importers/onesite/onesite_exported_mr_1 ***"

clean:
	rm -f [a-z] onesite.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail coa.csv iso.sql
	@echo "*** CLEAN completed in test/importers/onesite/onesite_exported_mr_1 ***"

test:
	touch fail
	if [ ! -f coa.csv ]; then cp ../../../jm1/coa.csv .;cat coa.csv|sed 's/^REX/ISO/' > x; mv x coa.csv; fi
	./functest.sh
	@echo "*** TEST completed in test/importers/onesite/onesite_exported_mr_1 ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/onesite/onesite_exported_mr_1 ***"

secure:
	@rm -f config.json confdev.json confprod.json
