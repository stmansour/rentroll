onesite: *.go
	go vet
	golint
	go build
	go test
	go install

clean:
	go clean
	@echo "*** CLEAN completed in importers/onesite ***"

test:
	go test
	@echo "*** TEST completed in importers/onesite ***"

#man:
#	nroff -man importers/onesite.1
#	cp importers/onesite.1 /usr/local/share/man/man1

package:
	@echo "*** PACKAGE completed in importers/onesite ***"
