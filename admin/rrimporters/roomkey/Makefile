TOP=../../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

roomkey: *.go
	@touch fail
	if [ ! -f ./config.json ]; then cp ${TOP}/confdev.json ./config.json; fi
	if [ ! -f ./mapper.json ]; then cp ${TOP}/importers/roomkey/mapper.json .; fi
	chmod 400 ./mapper.json
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail

clean:
	rm -f roomkey config.json mapper.json fail
	@echo "*** CLEAN completed in rrimporters/roomkey ***"

test:
	@echo "*** TEST completed in rrimporters/roomkey ***"

package: roomkey
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/roomkey/
	cp ./config.json ${TOP}/tmp/rentroll/importers/roomkey/config.json
	rm -f ${TOP}/tmp/rentroll/importers/roomkey/mapper.json
	cp ./mapper.json ${TOP}/tmp/rentroll/importers/roomkey/mapper.json
	cp ./roomkey ${TOP}/tmp/rentroll/importers/roomkey/roomkeyload
	@echo "*** PACKAGE completed in rrimporters/roomkey ***"
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
