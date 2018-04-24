DIRS = roomkey_sample_guest roomkey_exported_guest

roomkey:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** Completed in test/importers/roomkey ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in test/importers/roomkey ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in test/importers/roomkey ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGING COMPLETE IN test/importers/roomkey ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
