TOP=../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

genjs: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	if [ ! -f ./conf.json ]; then cp ${TOP}/confdev.json ./conf.json; fi
	go build
	@rm -f fail

clean:
	rm -f genjs conf.json genjs fail
	@echo "*** CLEAN completed in genjs ***"

test:
	@echo "*** TEST completed in genjs ***"

man:
	nroff -man genjs.1
	cp genjs.1 /usr/local/share/man/man1

# package:
# 	./genjs > ${TOP}/js/lists.js
# 	# cp genjs.1 ${TOP}/tmp/rentroll/man/man1
# 	@echo "*** PACKAGE completed in genjs ***"
