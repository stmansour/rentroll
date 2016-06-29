DIRS = db rlib rcsv admin test
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
	rm -f rentroll ver.go conf.json rentroll.log

test: package
	for dir in $(DIRS); do make -C $$dir test;done
	go test

man: rentroll.1
	cp rentroll.1 /usr/local/share/man/man1

package: rentroll
	rm -rf tmp
	mkdir -p tmp/rentroll
	mkdir -p tmp/rentroll/man/man1/
	for dir in $(DIRS); do make -C $$dir package;done
	cp rentroll ./tmp/rentroll/
	cp conf.json ./tmp/rentroll/
	cp -r html ./tmp/rentroll/
	@echo "*** PACKAGE COMPLETED ***"

t:
	./mdb
	./rentroll

all: clean rentroll test


try: clean rentroll package
