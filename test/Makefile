DIRS = ledger1

db:
	for dir in $(DIRS); do make -C $$dir;done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done
	@echo "*** CLEAN completed in test ***"
	
test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in test ***"