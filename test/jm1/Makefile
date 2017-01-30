all:
	@echo "*** Completed in jm1 ***"

clean:
	rm -f rentroll.log log llog ll.g err.txt [a-z] [a-z][a-z0-9] [a-z][a-z1-9][0-9] fail
	@echo "*** CLEAN completed in jm1 ***"

test:
	touch fail
	./functest.sh
	@echo "*** TEST completed in jm1 ***"
	rm -f fail

package:
	touch fail
	cp *.csv ../../tmp/rentroll/example/csv/
	cp *.html ../../tmp/rentroll/example/csv/
	@echo "*** PACKAGE completed in jm1 ***"
	rm -f fail

pubsandboxdb:
	mysqldump --no-defaults rentroll >rentrolldb.sql ; gzip rentrolldb.sql ; /usr/local/accord/bin/deployfile.sh rentrolldb.sql.gz accord/db
