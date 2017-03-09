notes: *.go
	touch fail
	cp ../../confdev.json conf.json
	go vet
	golint
	go build
	@echo "*** Completed in Notes ***"
	rm -f fail

clean:
	rm -f notes.log log *.g ./gold/*.g err.txt notes [a-z] fail
	@echo "*** CLEAN completed in Notes ***"

test:	notes
	touch fail
	./functest.sh
	@echo "*** TEST completed in Notes ***"
	rm -f fail

package:
	@echo "*** PACKAGE completed in Notes ***"
