DIRS = db rtags rlib rrpt rcsv worker bizlogic ws admin importers webclient tools test
TOP = .
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=.
DIST=tmp
GITHOOKDIR=.git/hooks

.PHONY:  test

rentroll: *.go config.json
	@find . -name "fail" -exec rm -r "{}" \;
	@touch fail
	for dir in $(DIRS); do make -C $$dir;done
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail
	@tools/bashtools/buildcheck.sh BUILD

all: clean rentroll test stats

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

jshint:
	@touch fail
	@${COUNTOL} "jshint --extract=always ./webclient/html/*.html ./webclient/html/test/*.html ./webclient/js/elems/*.js"
	@rm -rf fail

try: build db4 githook

db4:
	cd tools/dbgen;./dbgen -f db4.json

build: clean rentroll package

testdb:
	cd test/ws;mysql --no-defaults rentroll < restore.sql

receiptsdb:
	mysql --no-defaults rentroll < receipts.sql

rrdb:
	mysql --no-defaults rentroll < rr.sql

dbschemachange:
	cd test/testdb;make clean test dbbackup;cd ../ws;make get
	@tools/bashtools/buildcheck.sh SCHEMA_UPDATE

stats:
	@echo "GO SOURCE CODE STATISTICS"
	@echo "----------------------------------------"
	@find . -name "*.go" | srcstats
	@echo "----------------------------------------"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	go clean
	rm -f rentroll ver.go config.json rentroll.log *.out restore.sql rrbkup rrnewdb rrrestore example fail GoAnalyzerError.log *.json

test: package
	@find . -name "fail" -exec rm -r "{}" \;
	@rm -f test/*/err.txt
	for dir in $(DIRS); do make -C $$dir test;done
	@tools/bashtools/buildcheck.sh TEST
	@./errcheck.sh

man: rentroll.1
	cp rentroll.1 /usr/local/share/man/man1

dev:
	ln -s ./webclient/js
	ln -s ./webclient/html

instman:
	pushd ${DIST}/rentroll;./installman.sh;popd

package: rentroll
	@find . -name "fail" -exec rm -r "{}" \;
	@touch fail
	rm -rf ${DIST}
	mkdir -p ${DIST}/rentroll
	mkdir -p ${DIST}/rentroll/man/man1/
	mkdir -p ${DIST}/rentroll/example/csv
	cp rentroll.1 ${DIST}/rentroll/man/man1
	for dir in $(DIRS); do make -C $$dir package;done
	cp rentroll ./${DIST}/rentroll/
	cp ../gotable/pdfinstall.sh ${DIST}/rentroll/
	cp activate.sh update.sh ./${DIST}/rentroll/
	rm -f ./rrnewdb ./rrbkup ./rrrestore
	ln -s ${DIST}/rentroll/rrnewdb
	ln -s ${DIST}/rentroll/rrbkup
	ln -s ${DIST}/rentroll/rrrestore
	@rm -f fail
	@echo "*** PACKAGE COMPLETED ***"
	@tools/bashtools/buildcheck.sh PACKAGE

publish: package
	cd ${DIST};if [ -f ./rentroll/config.json ]; then mv ./rentroll/config.json .; fi
	cd ${DIST};tar cvf rentroll.tar rentroll; gzip rentroll.tar
	cd ${DIST};/usr/local/accord/bin/deployfile.sh rentroll.tar.gz jenkins-snapshot/rentroll/latest
	cd ${DIST};if [ -f ./config.json ]; then mv ./config.json ./rentroll/config.json; fi

pubimages:
	cd ${DIST}/rentroll;find . -name "*.png" | tar -cf rrimages.tar -T - ;gzip rrimages.tar ;/usr/local/accord/bin/deployfile.sh rrimages.tar.gz jenkins-snapshot/rentroll/latest

pubjs:
	cd ${DIST}/rentroll;mv js/bundle*.js .;tar czvf rrjs.tar.gz ./js;mv bundle*.js js/;/usr/local/accord/bin/deployfile.sh rrjs.tar.gz jenkins-snapshot/rentroll/latest

pubdb:
	# testing db
	cd ./test/testdb;make dbbackup

pubfa:
	# font awesome
	cd ${DIST}/rentroll;tar czvf fa.tar.gz ./webclient/html/fa;/usr/local/accord/bin/deployfile.sh fa.tar.gz jenkins-snapshot/rentroll/latest

# publish all the non-os-dependent files to the repo
pub: pubjs pubimages pubdb pubfa

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
	@if [ -d ${DIST} ]; then find ${DIST}/ -name config.json -exec rm {} \;; fi

githook: addgspattern
	./linkgithook.sh

addgspattern:
	./addgspattern.sh
