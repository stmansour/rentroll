TOP=../..
TOOLDIR=${TOP}/tools
BASHDIR=${TOOLDIR}/bashtools
BINDIR=../../tmp/rentroll
BKUP=${BINDIR}/rrbkup

testdb:
	@echo "**** build complete in testdb ***"

test:
	@touch fail
	./functest.sh
	mysqldump --no-defaults rentroll >restore.sql
	mv restore.sql ../ws/
	@rm -f fail

package:
	@echo "*** package completed in testdb ***"


clean:
	rm -f xyz rex.sql onesite.log log llog ll.g err.txt [a-z] [a-z][a-z1-9] [a-z].g gold/*.g fail
	@echo "*** CLEAN completed in testdb ***"

# This target deploys the file used as the database for interactive test nodes.
dbbackup:
	@touch fail
	pushd ${BINDIR};./rrbkup -f rentrolldb;cd bkup;/usr/local/accord/bin/deployfile.sh rentrolldb.sql.gz accord/db;cd ..; rm -rf bkup;popd
	@rm -f fail

get:
	${BASHDIR}/getdb.sh

secure:
	@rm -f config.json confdev.json confprod.json
