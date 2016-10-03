DIRS = db rlib rcsv rrpt admin test
.PHONY:  test

rentroll: *.go ver.go
	cp confdev.json conf.json
	for dir in $(DIRS); do make -C $$dir;done
	go vet
	golint
	./mkver.sh
	go build

ver.go:
	./mkver.sh

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	go clean
	rm -f rentroll ver.go conf.json rentroll.log *.out

test: package
	rm -f test/*/err.txt
	for dir in $(DIRS); do make -C $$dir test;done
	@./errcheck.sh

man: rentroll.1
	cp rentroll.1 /usr/local/share/man/man1

instman:
	pushd tmp/rentroll;./installman.sh;popd

package: rentroll
	rm -rf tmp
	mkdir -p tmp/rentroll
	mkdir -p tmp/rentroll/man/man1/
	cp rentroll.1 tmp/rentroll/man/man1
	for dir in $(DIRS); do make -C $$dir package;done
	cp rentroll ./tmp/rentroll/
	cp conf.json ./tmp/rentroll/
	cp -r html ./tmp/rentroll/
	cp activate.sh ./tmp/rentroll/
	@echo "*** PACKAGE COMPLETED ***"

publish: package
	cd tmp;tar cvf rentroll.tar rentroll; gzip rentroll.tar
	cd tmp;/usr/local/accord/bin/deployfile.sh rentroll.tar.gz jenkins-snapshot/rentroll/latest

all: clean rentroll test


try: clean rentroll package
