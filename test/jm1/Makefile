all:
	@echo "*** Completed in jm1 ***"

clean:
	rm -f rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z0-9] [a-z][a-z1-9][0-9] fail rex.sql
	@echo "*** CLEAN completed in jm1 ***"

test:
	touch fail
	./functest.sh
	mysqldump --no-defaults rentroll >rex.sql
	cp rex.sql ../testdb/
	@rm -f fail
	@echo "*** TEST completed in jm1 ***"

package:
	touch fail
	cp *.csv ../../tmp/rentroll/example/csv/
	cp *.html ../../tmp/rentroll/example/csv/
	@echo "*** PACKAGE completed in jm1 ***"
	@rm -f fail

pubsandboxdb:
	mysqldump --no-defaults rentroll >rentrolldb.sql ; gzip rentrolldb.sql ; /usr/local/accord/bin/deployfile.sh rentrolldb.sql.gz accord/db
