TOP=../../..
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=roomkey

roomkey: *.go
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/roomkey/
	if [ ! -f ./config.json ]; then cp ${TOP}/confdev.json ./config.json; fi
	if [ ! -f ./mapper.json ]; then cp ${TOP}/importers/roomkey/mapper.json .; fi
	chmod 400 ./mapper.json
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail

clean:
	rm -f roomkey config.json mapper.json fail
	@echo "*** CLEAN completed in ${THISDIR}  ***"

relink:
	go build
	@echo "*** Relink completed in ${THISDIR} ***"

test:
	@echo "*** TEST completed in ${THISDIR}  ***"

package: roomkey
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/roomkey/
	if [ -f config.json ]; then cp config.json ${TOP}/tmp/rentroll/importers/roomkey/config.json; fi
	rm -f ${TOP}/tmp/rentroll/importers/roomkey/mapper.json
	cp ./mapper.json ${TOP}/tmp/rentroll/importers/roomkey/mapper.json
	cp ./roomkey ${TOP}/tmp/rentroll/importers/roomkey/roomkeyload
	@echo "*** PACKAGE completed in ${THISDIR}  ***"
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
