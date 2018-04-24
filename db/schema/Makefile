TOPDIR="../.."

newdb: tws.sql doit
	@echo "*** completed in db/schema ***"

doit:
	@touch fail
	id
	whoami
	cat schema.sql tws.sql > db.sql
	mysql --no-defaults < db.sql
	@rm -rf fail

tws.sql:
	cp ${TOPDIR}/../tws/db/schema/tws.sql .

clean:
	rm -f tws.sql db.sql fail
	@echo "*** CLEAN completed in db/schema ***"

test:
	@echo "*** TEST completed in db/schema ***"

package:
	cat schema.sql tws.sql > db.sql
	cp db.sql ../../tmp/rentroll/schema.sql
	@echo "*** PACKAGE completed in db/schema ***"

secure:
	@rm -f config.json confdev.json confprod.json
