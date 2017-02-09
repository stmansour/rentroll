DIRS = db rlib rrpt rcsv admin importers test
TOP = .
COUNTOL=${TOP}/test/share/countol.sh

# 

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
	test/share/buildcheck.sh BUILD

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
	@test/share/buildcheck.sh TEST
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
	@test/share/buildcheck.sh PACKAGE

publish: package
	cd tmp;tar cvf rentroll.tar rentroll; gzip rentroll.tar
	cd tmp;/usr/local/accord/bin/deployfile.sh rentroll.tar.gz jenkins-snapshot/rentroll/latest

pubimages:
	cd tmp/rentroll;find . -name "*.png" | tar -cf rrimages.tar -T - ;gzip rrimages.tar ;/usr/local/accord/bin/deployfile.sh rrimages.tar.gz jenkins-snapshot/rentroll/latest

pubjs:
	cd tmp/rentroll;tar czvf rrjs.tar.gz ./js;/usr/local/accord/bin/deployfile.sh rrjs.tar.gz jenkins-snapshot/rentroll/latest

pubdb:
	cd ./test/testdb;make dbbackup

pub: pubjs pubimages pubdb


all: clean rentroll test
	@echo
	@echo "GO SOURCE CODE STATISTICS"
	@echo "----------------------------------------"
	@find . -name "*.go" | srcstats
	@echo "----------------------------------------"

try: clean rentroll package
