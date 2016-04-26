newbiz: *.go
	go vet
	golint
	../../mkver.sh
	go build

clean:
	rm -f newbiz [t-z] llog llog.gold xxqq ver.go log
	@echo "*** CLEAN completed in newbiz ***"

test: newbiz
	./functest.sh
	@echo "*** TEST completed in newbiz ***"
