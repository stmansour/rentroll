rrlib: *.go
	go vet
	golint
	go build
	go test
	go install

test:
	go test
	@echo "*** TESTS COMPLETE in rrpt -- ALL TESTS PASS ***"

clean:
	go clean

package:
	@echo "*** PACKAGE completed in rrpt ***"
