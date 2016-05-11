rrnewdb: *.sh

clean:
	rm -f rrnewdb
	@echo "*** Clean completed in admin/rrnewdb ***"

install: rrnewdb
	@echo "*** Packaging completed in admin/rrnewdb ***"

package: rrnewdb
	cp rrnewdb ../../tmp/rentroll/
	cp *.1 ../../tmp/rentroll/man/man1/
	@echo "*** Packaging completed in admin/rrnewdb ***"

packageqa: rrnewdb
	cp rrnewdb ../../tmp/rentroll/
	cp *.1 ../../tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in admin/rrnewdb ***"

test:
	@echo "*** Testing completed in admin/rrnewdb ***"

man:
	nroff -man rrnewdb.1
