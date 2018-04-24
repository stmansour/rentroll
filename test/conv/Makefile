TOP = ../..
COUNTOL=${TOP}/tools/bashtools/countol.sh

all: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -rf fail

clean:
	rm -f conv

secure:
	@rm -f config.json confdev.json confprod.json
