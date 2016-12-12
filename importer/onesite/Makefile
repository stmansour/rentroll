onesite: *.go
	go vet
	golint
	go build
	go install

clean:
	@echo "*** CLEAN completed in importer/onesite ***"

test:
	@echo "*** TEST completed in importer/onesite ***"

#man:
#	nroff -man importer/onesite.1
#	cp importer/onesite.1 /usr/local/share/man/man1

package:
	mkdir -p ../../tmp/rentroll/onesite/
	cp mapper.json ../../tmp/rentroll/onesite/
	@echo "*** PACKAGE completed in importer/onesite ***"
