TOP = ../..
COUNTOL=${TOP}/test/share/countol.sh

all: *.go
	@touch fail
	@${COUNTOL} "go vet"
	@${COUNTOL} golint
	go build
	@rm -rf fail

clean:
	rm -f conv
