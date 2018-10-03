DIRS = onesite roomkey
THISDIR=rrimporters

rrimporters:
	for dir in $(DIRS); do make -C $$dir; done

relink:
	for dir in $(DIRS); do make -C $$dir relink; done
	@echo "*** Relink completed in ${THISDIR} ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in ${THISDIR} ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in ${THISDIR} ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: ${THISDIR}
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
	@echo "*** SECURE completed in ${THISDIR} ***"
