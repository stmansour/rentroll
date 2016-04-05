all:
	@echo "+-----------------------------"
	@echo "|  Creating new RentRoll DB"
	@echo "+-----------------------------"
	cd ../../db/schema;make newdb
	@echo "+----------------------------------------"
	@echo "|  Initializing DB with EXAMPLE 1 data"
	@echo "+----------------------------------------"
	mysql --no-defaults < init.sql

clean:
	@echo "*** CLEAN completed in ledger1 ***"

test:
	@echo "*** TEST completed in ledger1 ***"
