TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=rrbkup

rrbkup: *.sh

clean:
	rm -f ${THISDIR} fail
	@echo "*** Clean completed in ${THISDIR} ***"

install: ${THISDIR}
	@echo "*** Packaging completed in ${THISDIR} ***"

relink:
	@echo "*** Relink completed in ${THISDIR} ***"

package: ${THISDIR}
	@touch fail
	cp ${THISDIR} ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentroll/man/man1/
	@echo "*** Packaging completed in ${THISDIR} ***"
	@rm -f fail

packageqa: ${THISDIR}
	@touch fail
	cp ${THISDIR} ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in ${THISDIR} ***"
	@rm -f fail

test:
	@echo "*** Testing completed in ${THISDIR} ***"

man:
	nroff -man ${THISDIR}.1

secure:
	@rm -f config.json confdev.json confprod.json
