TOP=../..
BINDIR=${TOP}/tmp/rentroll
COUNTOL=${TOP}/tools/bashtools/countol.sh
THIS="schemacmp"

schcmp:
	@echo "*** Completed in ${THIS} ***"

clean:
	rm -rf t cmds *1 def*.sh u local remote *.json v report.txt x y
	@echo "*** CLEAN completed in ${THIS} ***"

config.json:
	@/usr/local/accord/bin/getfile.sh accord/db/confprod.json

test:
	@echo "*** TEST completed in ${THIS} ***"

package:
	@echo "*** PACKAGE completed in ${THIS} ***"

secure:
	@rm -f config.json confdev.json confprod.json
