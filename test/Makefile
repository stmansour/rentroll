DIRS = engine newbiz rex rsp gsr gsr1 gsr2 notes rrr rr1

db:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** MAKE completed in test ***"

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in test ***"
	
test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in test ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in test ***"
