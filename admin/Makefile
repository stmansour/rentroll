DIRS = rrbkup rrnewdb rrrestore rrloadcsv

admin:
	for dir in $(DIRS); do make -C $$dir; done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** ALL TESTS PASSED ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	cp installman.sh ../tmp/rentroll/
