DIRS=setup newbiz crypto workerasm mrr rrr rr1 rr rr_use_cases jm1 gsr notes ccc upd acctbal gap importers bizdelete testdb bizlogic ws websvc1 websvc2 websvc3 payorstmt roller tws tws3 receipts webclient
# DIRS=setup newbiz crypto workerasm mrr rrr rr1 rr rr_use_cases jm1 gsr notes ccc upd acctbal gap importers bizdelete testdb bizlogic ws websvc1 websvc2 websvc3 payorstmt roller tws tws3 receipts
TESTREPORT="testreport.txt"

db:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** MAKE completed in test ***"

clean:
	rm -rf dbreport.txt testreport.txt dbqqqmods.sql cmds x y t refdb checkdb def1.sh def2.sh
	for dir in $(DIRS); do make -C $$dir clean;done
	cd ./testdb;make clean
	cd ./importers/onesite/okc;make clean
	cd ./tws2;make clean
	cd ./webclient;make clean
	@echo "*** CLEAN completed in test ***"

test: validate functest

functest:
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

validate:
	@echo "FUNCTIONAL TEST Report" > ${TESTREPORT}
	@echo "VALIDATING LOCAL DATABASE SCHEMAS" | tee -a ${TESTREPORT}
	@./localcmp.sh

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in test ***"

secure:
	for dir in $(DIRS); do make -C $${dir} secure;done
	@rm -f config.json confdev.json confprod.json
