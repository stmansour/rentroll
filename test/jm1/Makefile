all:
	@echo "*** Completed in jm1 ***"

clean:
	rm -f rentroll.log log llog ll.g err.txt [a-z] [a-z][a-z0-9] [a-z][a-z1-9][0-9]
	@echo "*** CLEAN completed in jm1 ***"

test:
	./functest.sh
	@echo "*** TEST completed in jm1 ***"

package:
	cp *.csv ../../tmp/rentroll/example/csv/
	cp *.html ../../tmp/rentroll/example/csv/
	@echo "*** PACKAGE completed in jm1 ***"

pubsandboxdb:
	mysqldump --no-defaults rentroll >rentrolldb.sql ; gzip rentrolldb.sql ; /usr/local/accord/bin/deployfile.sh rentrolldb.sql.gz accord/db
