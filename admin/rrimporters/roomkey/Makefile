roomkey: *.go
	if [ ! -f ./conf.json ]; then cp ../../../confdev.json ./conf.json; fi
	if [ ! -f ./mapper.json ]; then cp ../../../importers/roomkey/mapper.json .; fi
	chmod 400 ./mapper.json
	go vet
	golint
	go build

clean:
	rm -f roomkey conf.json mapper.json
	@echo "*** CLEAN completed in rrimporters/roomkey ***"

test:
	@echo "*** TEST completed in rrimporters/roomkey ***"

package: roomkey
	mkdir -p ../../../tmp/rentroll/importers/roomkey/
	cp ./conf.json ../../../tmp/rentroll/importers/roomkey/conf.json
	cp ./mapper.json ../../../tmp/rentroll/importers/roomkey/mapper.json
	cp ./roomkey ../../../tmp/rentroll/importers/roomkey/roomkeyload
	@echo "*** PACKAGE completed in rrimporters/roomkey ***"
