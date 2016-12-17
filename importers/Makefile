DIRS = core onesite

importers:
	for dir in $(DIRS); do make -C $$dir; done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	rm -f importers

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** ALL TESTS PASSED in importers ***"

package: importers
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGING COMPLETE in importers ***"
