roomkey: *.go
	go vet
	golint
	go build
	go test
	go install

clean:
	go clean
	@echo "*** CLEAN completed in importers/roomkey ***"

test:
	go test
	@echo "*** TEST completed in importers/roomkey ***"

package: roomkey
	@echo "*** PACKAGE completed in importers/roomkey ***"
