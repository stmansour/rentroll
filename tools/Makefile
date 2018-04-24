DIRS = dbgen schemacmp bashtools

tools:
	for dir in $(DIRS); do make -C $$dir; done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** ALL TESTS PASSED ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
