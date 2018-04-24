all:
	@echo "*** Completed in svc ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z1-9] qq? fail request w.[a-z] restore.sql ws[0-9]*
	@echo "*** CLEAN completed in svc ***"

test: restore.sql
	@touch fail
	./functest.sh
	@echo "*** TEST completed in svc ***"
	@rm -f fail

restore.sql:
	cd ../testdb;make test

package:
	@echo "*** PACKAGE completed in svc ***"

get:
	@touch fail
	../../tools/bashtools/getdb.sh
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
