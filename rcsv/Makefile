rrlib: *.go
	go vet
	golint
	go build
	go test
	go install

test:
	go test
	@echo "*** TESTS COMPLETE -- ALL TESTS PASS ***"

clean:
	go clean

package:
	@echo "*** PACKAGE completed in rlib ***"
