TOP=../..
TOOLDIR=${TOP}/tools
BASHDIR=${TOOLDIR}/bashtools
BINDIR=../../tmp/rentroll
BKUP=${BINDIR}/rrbkup

test:
	@touch fail
	./functest.sh
	@rm -f fail

clean:
	rm -f onesite.log log llog ll.g err.txt [a-z] [a-z][a-z1-9] [a-z].g gold/*.g fail
	@echo "*** CLEAN completed in 1site ***"

# This target deploys the file used as the database for interactive test nodes.
dbbackup:
	@touch fail
	pushd ${BINDIR};./rrbkup -f rentrolldb;cd bkup;/usr/local/accord/bin/deployfile.sh rentrolldb.sql.gz accord/db;cd ..; rm -rf bkup;popd
	@rm -f fail

get:
	${BASHDIR}/getdb.sh
