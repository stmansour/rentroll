TOP=../../..
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=onesite
CONF=config.json

onesite: *.go
	@touch fail
	if [ ! -f ./${CONF} ]; then cp ${TOP}/confdev.json ./${CONF}; fi
	if [ ! -f ./mapper.json ]; then cp ${TOP}/importers/onesite/mapper.json .; fi
	chmod 400 ./mapper.json
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -f fail
	@echo "*** Relink completed in ${THISDIR} ***"

relink:
	go build
	@echo "*** Relink completed in ${THISDIR} ***"

clean:
	rm -f onesite ${CONF} mapper.json
	@echo "*** CLEAN completed in ${THISDIR} ***"

test:
	@echo "*** TEST completed in ${THISDIR} ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: onesite
	@touch fail
	mkdir -p ${TOP}/tmp/rentroll/importers/${THISDIR}/
	if [ -f ${TOP}/tmp/rentroll/importers/${THISDIR}/mapper.json ]; then rm -f ${TOP}/tmp/rentroll/importers/${THISDIR}/mapper.json; fi
	if [ -f ./${CONF} ]; then cp ./${CONF} ${TOP}/tmp/rentroll/importers/${THISDIR}/${CONF}; fi
	cp ./mapper.json ${TOP}/tmp/rentroll/importers/${THISDIR}/
	cp ./${THISDIR} ${TOP}/tmp/rentroll/importers/${THISDIR}/onesiteload
	@echo "*** PACKAGE completed in ${THISDIR} ***"
	@rm -f fail

secure:
	@rm -f ${CONF} confdev.json confprod.json
