DIRS = db rlib test
.PHONY:  test

rentroll: *.go mkver.sh
	for dir in $(DIRS); do make -C $$dir;done
	go vet
	golint
	./mkver.sh
	go build

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	rm -f rentroll ver.go

test:
	for dir in $(DIRS); do make -C $$dir test;done
	go test
	@echo "*** TESTING COMPLETE -- ALL TESTS PASSED ***"

man: rentroll.1
	cp rentroll.1 /usr/local/share/man/man1
