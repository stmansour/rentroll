onesite: *.go
	go vet
	golint
	# ../../mkver.sh
	go build

clean:
	rm -f onesite
	@echo "*** CLEAN completed in rrimporters/onesite ***"

test:
	@echo "*** TEST completed in rrimporters/onesite ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: onesite
	mkdir -p ../../../tmp/rentroll/importers/onesite/
	# cp ../../../importers/onesite/mapper.json ../../../tmp/rentroll/importers/onesite/
	cp onesite ../../../tmp/rentroll/onesiteload
	# cp ../../../conf.json ../../../tmp/rentroll/conf.json
	@echo "*** PACKAGE completed in rrimporters/onesite ***"
