onesite: *.go
	go vet
	golint
	go build

clean:
	@echo "*** CLEAN completed in importers/onesite ***"

test:
	@echo "*** TEST completed in importers/onesite ***"

#man:
#	nroff -man importers/onesite.1
#	cp importers/onesite.1 /usr/local/share/man/man1

package:
	mkdir -p ../../tmp/rentroll/importers/onesite/
	cp mapper.json ../../tmp/rentroll/importers/onesite/
	@echo "*** PACKAGE completed in importers/onesite ***"
