TOP=../../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

onesite: *.go
	@touch fail
	if [ ! -f ./config.json ]; then cp ${TOP}/confdev.json ./config.json; fi
	if [ ! -f ./mapper.json ]; then cp ${TOP}/importers/onesite/mapper.json .; fi
	chmod 400 ./mapper.json
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail

clean:
	rm -f onesite config.json mapper.json
	@echo "*** CLEAN completed in rrimporters/onesite ***"

test:
	@echo "*** TEST completed in rrimporters/onesite ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: onesite
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/onesite/
	if [ -f ${TOP}/tmp/rentroll/importers/onesite/mapper.json ]; then rm -f ${TOP}/tmp/rentroll/importers/onesite/mapper.json; fi
	cp ./config.json ${TOP}/tmp/rentroll/importers/onesite/config.json
	cp ./mapper.json ${TOP}/tmp/rentroll/importers/onesite/
	cp ./onesite ${TOP}/tmp/rentroll/importers/onesite/onesiteload
	@echo "*** PACKAGE completed in rrimporters/onesite ***"
	@rm -f fail

secure:
	@rm -f config.json confdev.json confprod.json
