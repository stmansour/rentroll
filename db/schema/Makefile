newdb: doit
	@echo "*** completed in db/schema ***"

doit:
	@touch fail
	id
	whoami
	mysql --no-defaults < schema.sql
	@rm -rf fail

clean:
	@echo "*** CLEAN completed in db/schema ***"

test:
	@echo "*** TEST completed in db/schema ***"

package:
	cp schema.sql ../../tmp/rentroll/
	@echo "*** PACKAGE completed in db/schema ***"
