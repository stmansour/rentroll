rentroll:
	cd lib; make
	go vet
	golint
	./mkver.sh
	go compile


man: rentroll.1
	cp rentroll.1 /usr/local/share/man/man1
