TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=rrwatchdog

rrwatchdog: *.sh

clean:
	rm -f ${THISDIR}
	@echo "*** Clean completed in ${THISDIR} ***"

relink:
	@echo "*** Relink completed in ${THISDIR} ***"

install: ${THISDIR}
	@echo "*** Packaging completed in ${THISDIR} ***"

package: ${THISDIR}
	cp ${THISDIR} ${BINDIR}/
	@echo "*** Packaging completed in ${THISDIR} ***"

test:
	@echo "*** Testing completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
