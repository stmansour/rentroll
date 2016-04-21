newbiz: *.go
	go vet
	golint
	../../mkver.sh
	go build

clean:
	rm -f newbiz [w-z] xxqq ver.go
	@echo "*** CLEAN completed in newbiz ***"

test: newbiz
	./functest.sh
	@echo "*** TEST completed in newbiz ***"
