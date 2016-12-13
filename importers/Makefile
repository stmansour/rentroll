DIRS = core onesite

importers:
	for dir in $(DIRS); do make -C $$dir; done
	go vet
	golint
	go build
	
clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	rm -f importers

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** ALL TESTS PASSED ***"

package: importers
	for dir in $(DIRS); do make -C $$dir package;done
	cp importers ../tmp/rentroll/onesiteLoad
	@echo "*** PACKAGING COMPLETE IN importers ***"
