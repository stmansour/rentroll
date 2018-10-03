THISDIR = notes

notes: *.go config.json
	touch fail
	go vet
	golint
	go build
	@echo "*** Completed in ${THISDIR} ***"
	@rm -f fail

clean:
	rm -f notes.log log *.g ./gold/*.g err.txt notes [a-z] fail conf*.json
	@echo "*** CLEAN completed in ${THISDIR} ***"

relink:
	go build
	@echo "*** RELINK completed in ${THISDIR} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confdev.json
	@cp confdev.json config.json

test:	notes
	touch fail
	./functest.sh
	@echo "*** TEST completed in ${THISDIR} ***"
	@rm -f fail

package:
	@echo "*** PACKAGE completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
