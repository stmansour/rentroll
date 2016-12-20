DIRS = onesite

rrimporters:
	for dir in $(DIRS); do make -C $$dir; done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in rrimporters ***"

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in rrimporters ***"

# man:
# 	nroff -man rrloadcsv.1
# 	cp rrloadcsv.1 /usr/local/share/man/man1

package: rrimporters
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in rrimporters ***"
