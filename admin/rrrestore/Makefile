rrrestore: *.sh

clean:
	rm -f rrrestore
	@echo "*** Clean completed in rrrestore ***"

install: rrrestore
	@echo "*** Packaging completed in rrrestore ***"

package: rrrestore
	cp rrrestore ../../tmp/rentroll/
	cp *.1 ../../tmp/rentroll/man/man1/
	@echo "*** Packaging completed in rrrestore ***"

packageqa: rrrestore
	cp rrrestore ../../tmp/rentroll/
	cp *.1 ../../tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in rrrestore ***"

test:
	@echo "*** Testing completed in rrrestore ***"

man:
	nroff -man rrrestore.1
