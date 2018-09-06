DIRS = dbgen schemacmp bashtools
RLNK = dbgen
THISDIR=tools

tools:
	for dir in $(DIRS); do make -C $$dir; done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** clean completed in ${THISDIR} ***"

relink:
	for dir in ${RLNK}; do make -C $$dir relink;done
	@echo "*** Relink completed in ${THISDIR} ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TESTS completed in ${THISDIR} ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** package completed in ${THISDIR} ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
