DIRS=setup newbiz workerasm mrr rrr rr1 rr rr_use_cases jm1 gsr notes ccc upd acctbal gap importers bizdelete testdb bizlogic ws websvc1 websvc2 websvc3 payorstmt roller receipts webclient
TESTREPORT="testreport.txt"

db:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** MAKE completed in test ***"

clean:
	rm -f testreport.txt dbqqqmods.sql
	for dir in $(DIRS); do make -C $$dir clean;done
	cd ./testdb;make clean
	cd ./importers/onesite/okc;make clean
	@echo "*** CLEAN completed in test ***"

test:
	@echo "FUNCTIONAL TEST Report" > ${TESTREPORT}
	@echo "Report Start: \c" >> ${TESTREPORT}
	@date >> ${TESTREPORT}
	@echo >> ${TESTREPORT}
	@echo "Result  Path                  Test Name                                 Phases" >> ${TESTREPORT}
	@echo "------  --------------------  ----------------------------------------  ------" >> ${TESTREPORT}
	for dir in $(DIRS); do make -C $$dir test;done
	@./totals.sh
	@echo >> ${TESTREPORT}
	@echo "Report Complete: \c" >> ${TESTREPORT}
	@date >> ${TESTREPORT}
	@echo "*** TEST completed in test ***"
	@cat ${TESTREPORT}

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in test ***"
