TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh
rrbkup: *.sh

clean:
	rm -f rrbkup fail
	@echo "*** Clean completed in rrbkup ***"

install: rrbkup
	@echo "*** Packaging completed in rrbkup ***"

package: rrbkup
	@touch fail
	cp rrbkup ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentroll/man/man1/
	@echo "*** Packaging completed in rrbkup ***"
	@rm -f fail

packageqa: rrbkup
	@touch fail
	cp rrbkup ${TOP}/tmp/rentroll/
	cp *.1 ${TOP}/tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in rrbkup ***"
	@rm -f fail

test:
	@echo "*** Testing completed in rrbkup ***"

man:
	nroff -man rrbkup.1

secure:
	@rm -f config.json confdev.json confprod.json
