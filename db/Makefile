DIRS = schema

db:
	for dir in $(DIRS); do make -C $$dir;done

clean:
	for dir in $(DIRS); do make -C $$dir clean;done

test:
	for dir in $(DIRS); do make -C $$dir test;done
	@echo "*** TEST completed in db ***"

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in db/schema ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
