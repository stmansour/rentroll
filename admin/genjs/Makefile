genjs: *.go
	if [ ! -f ./conf.json ]; then cp ../../confdev.json ./conf.json; fi
	go vet
	golint
	go build

clean:
	rm -f genjs conf.json genjs
	@echo "*** CLEAN completed in genjs ***"

test:
	@echo "*** TEST completed in genjs ***"

man:
	nroff -man genjs.1
	cp genjs.1 /usr/local/share/man/man1

package:
	../../tmp/rentroll/genjs > ../../js/lists.js
	cp genjs ../../tmp/rentroll/genjs
	# cp genjs.1 ../../tmp/rentroll/man/man1
	@echo "*** PACKAGE completed in genjs ***"
