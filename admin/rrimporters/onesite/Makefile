TOP=../../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

onesite: *.go
	@touch fail
	if [ ! -f ./conf.json ]; then cp ${TOP}/confdev.json ./conf.json; fi
	if [ ! -f ./mapper.json ]; then cp ${TOP}/importers/onesite/mapper.json .; fi
	chmod 400 ./mapper.json
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail

clean:
	rm -f onesite conf.json mapper.json
	@echo "*** CLEAN completed in rrimporters/onesite ***"

test:
	@echo "*** TEST completed in rrimporters/onesite ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: onesite
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/onesite/
	cp ./conf.json ${TOP}/tmp/rentroll/importers/onesite/conf.json
	cp ./mapper.json ${TOP}/tmp/rentroll/importers/onesite/
	cp ./onesite ${TOP}/tmp/rentroll/importers/onesite/onesiteload
	@echo "*** PACKAGE completed in rrimporters/onesite ***"
	@rm -f fail
