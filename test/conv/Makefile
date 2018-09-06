TOP = ../..
COUNTOL=${TOP}/tools/bashtools/countol.sh
THISDIR=conv

all: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -rf fail

clean:
	rm -f conv

relink:
	go build
	@echo "*** RELINK completed in ${THISDIR} ***"

secure:
	@rm -f config.json confdev.json confprod.json
