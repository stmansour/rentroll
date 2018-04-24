DIRS = test01 test02 test03 test04 test05 test06

onesite:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** Completed in test/rr_use_cases ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in test/rr_use_cases ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in test/rr_use_cases ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGING COMPLETE IN test/rr_use_cases ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
