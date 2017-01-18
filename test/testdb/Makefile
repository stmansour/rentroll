BINDIR=../../tmp/rentroll
BKUP=${BINDIR}/rrbkup

test:
	./functest.sh

clean:
	rm -f onesite.log log llog ll.g err.txt [a-z] [a-z][a-z1-9] [a-z].g gold/*.g
	@echo "*** CLEAN completed in 1site ***"

# This target deploys the file used as the database for interactive test nodes.
dbbackup:
	pushd ${BINDIR};./rrbkup -f rentrolldb;cd bkup;/usr/local/accord/bin/deployfile.sh rentrolldb.sql.gz accord/db;cd ..; rm -rf bkup;popd

