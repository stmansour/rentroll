TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
rrwatchdog: *.sh

clean:
	rm -f rrwatchdog
	@echo "*** Clean completed in rrwatchdog ***"

install: rrwatchdog
	@echo "*** Packaging completed in rrwatchdog ***"

package: rrwatchdog
	cp rrwatchdog ${BINDIR}/
	@echo "*** Packaging completed in rrwatchdog ***"

test:
	@echo "*** Testing completed in rrwatchdog ***"

secure:
	@rm -f config.json confdev.json confprod.json
