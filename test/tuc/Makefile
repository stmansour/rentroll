all:
	@echo "create new db"
	cd ../../db/schema;make newdb
	@echo "initialize with tucasa info"
	mysql --no-defaults < init.sql

clean:
	@echo "*** CLEAN completed in tuc ***"

test:
	@echo "*** TEST completed in tuc ***"
