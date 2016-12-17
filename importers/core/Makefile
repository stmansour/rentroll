core: *.go
	go vet
	golint
	go build
	go test
	go install

clean:
	go clean
	@echo "*** CLEAN completed in importers/core ***"

test:
	go test
	@echo "*** TEST completed in importers/core ***"

#man:
#	nroff -man importers/core.1
#	cp importers/core.1 /usr/local/share/man/man1

package:
	@echo "*** PACKAGE completed in importers/core ***"
