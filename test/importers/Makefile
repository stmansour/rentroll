DIRS = onesite roomkey
THISDIR = importers
RLNK =

importers:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** Completed in test/importers ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in test/importers ***"

relink:
	for dir in $(RLNK); do make -C $$dir relink;done
	@echo "*** RELINK completed in ${THISDIR} ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in test/importers ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGING COMPLETE IN test/importers ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
