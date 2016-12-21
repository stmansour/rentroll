onesite: *.go
	if [ ! -f ./conf.json ]; then cp ../../../confdev.json ./conf.json; fi
	if [ ! -f ./mapper.json ]; then cp ../../../importers/onesite/mapper.json .; fi
	chmod 400 ./mapper.json
	go vet
	golint
	go build

clean:
	rm -f onesite conf.json mapper.json
	@echo "*** CLEAN completed in rrimporters/onesite ***"

test:
	@echo "*** TEST completed in rrimporters/onesite ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: onesite
	mkdir -p ../../../tmp/rentroll/importers/onesite/
	cp ./conf.json ../../../tmp/rentroll/importers/onesite/conf.json
	cp ./mapper.json ../../../tmp/rentroll/importers/onesite/
	cp ./onesite ../../../tmp/rentroll/importers/onesite/onesiteload
	@echo "*** PACKAGE completed in rrimporters/onesite ***"
