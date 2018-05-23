TEMPCSVSTORE = "../../../../tmp/rentroll/importers/roomkey/temp_CSVs"

roomkey:
	@echo "*** Completed in test/importers/roomkey_exported_guest ***"

clean:
	rm -f [a-z] roomkey.log log ${TEMPCSVSTORE}/*.csv *.g ./gold/*.g err.txt fail
	@echo "*** CLEAN completed in test/importers/roomkey_exported_guest ***"

test:
	if [ ! -f coa.csv ]; then cp ../../../jm1/coa.csv .;cat coa.csv|sed 's/^REX/OKC/' > x; mv x coa.csv; fi
	if [ ! -f ar.csv ]; then cp ../../../jm1/ar.csv .;cat ar.csv|sed 's/^REX/OKC/' > x; mv x ar.csv; fi
	touch fail
	./functest.sh
	@echo "*** TEST completed in test/importers/roomkey_exported_guest ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in test/importers/roomkey_exported_guest ***"

secure:
	@rm -f config.json confdev.json confprod.json
