DIRS = core onesite

importer:
	for dir in $(DIRS); do make -C $$dir; done
	go vet
	golint
	go build
	
clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	rm -f importer

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** ALL TESTS PASSED ***"

package: importer
	for dir in $(DIRS); do make -C $$dir package;done
	cp importer ../tmp/rentroll/onesiteLoad
	@echo "*** PACKAGING COMPLETE IN importer ***"
