DIRS = onesite

importers:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** Completed in importers ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in importers ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in importers ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGING COMPLETE IN importers ***"
