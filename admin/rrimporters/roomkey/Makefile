TOP=../../..
COUNTOL=${TOP}/test/share/countol.sh

roomkey: *.go
	@touch fail
	if [ ! -f ./conf.json ]; then cp ${TOP}/confdev.json ./conf.json; fi
	if [ ! -f ./mapper.json ]; then cp ${TOP}/importers/roomkey/mapper.json .; fi
	chmod 400 ./mapper.json
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail

clean:
	rm -f roomkey conf.json mapper.json
	@echo "*** CLEAN completed in rrimporters/roomkey ***"

test:
	@echo "*** TEST completed in rrimporters/roomkey ***"

package: roomkey
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/roomkey/
	cp ./conf.json ${TOP}/tmp/rentroll/importers/roomkey/conf.json
	rm -f ${TOP}/tmp/rentroll/importers/roomkey/mapper.json
	cp ./mapper.json ${TOP}/tmp/rentroll/importers/roomkey/mapper.json
	cp ./roomkey ${TOP}/tmp/rentroll/importers/roomkey/roomkeyload
	@echo "*** PACKAGE completed in rrimporters/roomkey ***"
	@rm -f fail
