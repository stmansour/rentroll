all:
	@echo "*** Completed in setup ***"

clean:
	@echo "*** CLEAN completed in setup ***"

test:
	./dosetup.sh
	@echo "*** TEST completed in setup ***"

package:
	@echo "*** PACKAGE completed in setup ***"

newdb:
	mysqldump --no-defaults accord > accord.sql

secure:
	@rm -f config.json confdev.json confprod.json
