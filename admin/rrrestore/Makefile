TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh
rrrestore: *.sh

clean:
	rm -f rrrestore fail
	@echo "*** Clean completed in rrrestore ***"

install: rrrestore
	@echo "*** Packaging completed in rrrestore ***"

package: rrrestore
	touch fail
	cp rrrestore ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentroll/man/man1/
	@echo "*** Packaging completed in rrrestore ***"
	@rm -f fail

packageqa: rrrestore
	touch fail
	cp rrrestore ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in rrrestore ***"
	@rm -f fail

test:
	@echo "*** Testing completed in rrrestore ***"

man:
	nroff -man rrrestore.1

secure:
	@rm -f config.json confdev.json confprod.json
