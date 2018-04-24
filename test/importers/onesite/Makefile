DIRS = onesite_sample onesite_exported_1 onesite_exported_2 onesite_exported_mr_1

onesite:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** Completed in test/importers/onesite ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in test/importers/onesite ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in test/importers/onesite ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGING COMPLETE IN test/importers/onesite ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
