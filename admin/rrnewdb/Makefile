TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

rrnewdb: *.sh

clean:
	rm -f rrnewdb fail
	@echo "*** Clean completed in admin/rrnewdb ***"

install: rrnewdb
	@echo "*** Packaging completed in admin/rrnewdb ***"

package: rrnewdb
	touch fail
	cp rrnewdb ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentroll/man/man1/
	@echo "*** Packaging completed in admin/rrnewdb ***"
	@rm -f fail

packageqa: rrnewdb
	touch fail
	cp rrnewdb ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in admin/rrnewdb ***"
	@rm -f fail

test:
	@echo "*** Testing completed in admin/rrnewdb ***"

man:
	nroff -man rrnewdb.1

secure:
	@rm -f config.json confdev.json confprod.json
