rrbkup: *.sh

clean:
	rm -f rrbkup
	@echo "*** Clean completed in rrbkup ***"

install: rrbkup
	@echo "*** Packaging completed in rrbkup ***"

package: rrbkup
	cp rrbkup ../../tmp/rentroll/
	cp *.1 ../../tmp/rentroll/man/man1/
	@echo "*** Packaging completed in rrbkup ***"

packageqa: rrbkup
	cp rrbkup ../../tmp/rentroll/
	cp *.1 ../../tmp/rentrollqa/man/man1/
	@echo "*** Packaging completed in rrbkup ***"

test:
	@echo "*** Testing completed in rrbkup ***"

man:
	nroff -man rrbkup.1
