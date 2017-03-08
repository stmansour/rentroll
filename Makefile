DIRS = db rlib rrpt rcsv admin importers tools test
TOP = .
COUNTOL=${TOP}/tools/bashtools/countol.sh

.PHONY:  test

rentroll: ver.go *.go
	@touch fail
	cp confdev.json conf.json
	for dir in $(DIRS); do make -C $$dir;done
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	./mkver.sh
	go build
	@rm -rf fail
	@tools/bashtools/buildcheck.sh BUILD

all: clean rentroll test stats

jshint:
	@touch fail
	@${COUNTOL} "jshint --extract=always html/*.html html/test/*.html js/rutil.js"
	@rm -rf fail

try: clean rentroll package

testdb:
	cd test/svc;mysql --no-defaults < restore.sql

dbschemachange:
	cd test/testdb;make clean test dbbackup;cd ../svc;make get
	@tools/bashtools/buildcheck.sh SCHEMA_UPDATE

rebuild: try testdb

stats:
	@echo "GO SOURCE CODE STATISTICS"
	@echo "----------------------------------------"
	@find . -name "*.go" | srcstats
	@echo "----------------------------------------"

ver.go:
	./mkver.sh

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	go clean
	rm -f rentroll ver.go conf.json rentroll.log *.out restore.sql rrbkup rrnewdb rrrestore example fail GoAnalyzerError.log

test: package
	@rm -f test/*/err.txt
	for dir in $(DIRS); do make -C $$dir test;done
	@tools/bashtools/buildcheck.sh TEST
	@./errcheck.sh

man: rentroll.1
	cp rentroll.1 /usr/local/share/man/man1

instman:
	pushd tmp/rentroll;./installman.sh;popd

package: rentroll
	@touch fail
	rm -rf tmp
	mkdir -p tmp/rentroll
	mkdir -p tmp/rentroll/man/man1/
	mkdir -p tmp/rentroll/example/csv
	cp rentroll.1 tmp/rentroll/man/man1
	for dir in $(DIRS); do make -C $$dir package;done
	cp rentroll ./tmp/rentroll/
	cp conf.json ./tmp/rentroll/
	cp -r html ./tmp/rentroll/
	if [ -e js ]; then cp -r js ./tmp/rentroll/ ; fi
	cp activate.sh update.sh ./tmp/rentroll/
	rm -f ./rrnewdb ./rrbkup ./rrrestore
	ln -s tmp/rentroll/rrnewdb
	ln -s tmp/rentroll/rrbkup
	ln -s tmp/rentroll/rrrestore
	@echo "*** PACKAGE COMPLETED ***"
	@rm -f fail
	@tools/bashtools/buildcheck.sh PACKAGE

publish: package
	cd tmp;tar cvf rentroll.tar rentroll; gzip rentroll.tar
	cd tmp;/usr/local/accord/bin/deployfile.sh rentroll.tar.gz jenkins-snapshot/rentroll/latest

pubimages:
	cd tmp/rentroll;find . -name "*.png" | tar -cf rrimages.tar -T - ;gzip rrimages.tar ;/usr/local/accord/bin/deployfile.sh rrimages.tar.gz jenkins-snapshot/rentroll/latest

pubjs:
	cd tmp/rentroll;tar czvf rrjs.tar.gz ./js;/usr/local/accord/bin/deployfile.sh rrjs.tar.gz jenkins-snapshot/rentroll/latest

pubdb:
	# testing db
	cd ./test/testdb;make dbbackup

pubfa:
	# font awesome
	cd tmp/rentroll;tar czvf fa.tar.gz ./html/fa;/usr/local/accord/bin/deployfile.sh fa.tar.gz jenkins-snapshot/rentroll/latest

# publish all the non-os-dependent files to the repo
pub: pubjs pubimages pubdb pubfa


