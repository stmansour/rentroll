DIRS = invalid_flow proper_sequence_flow improper_sequence_flow raid_version

THISDIR="actions"

build:
	go build
	@echo "*** Completed in ${THISDIR} ***"

clean:
	rm -f rrlog rentroll.log log llog *.g ./gold/*.g err.txt [a-z] [a-z][a-z0-9] a[0-9]* qq? ${THISDIR} fail conf*.json request serverreply svrresponse
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in ${THISDIR} ***"

relink:
	go build
	@echo "*** RELINK completed in ${THISDIR} ***"

test: build
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in ${THISDIR} ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
