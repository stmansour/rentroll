DIRS = setup newbiz mrr rrr rr1 rr rr_use_cases jm1 gsr notes ccc upd acctbal gap importers bizdelete testdb bizlogic ws websvc1 websvc2 webclient payorstmt
DIRS = setup newbiz mrr rrr rr1 rr rr_use_cases jm1 gsr notes ccc upd acctbal gap importers bizdelete testdb bizlogic ws websvc1 websvc2 payorstmt
db:
	for dir in $(DIRS); do make -C $$dir;done
	@echo "*** MAKE completed in test ***"

clean:
	rm -f testreport.txt dbqqqmods.sql
	for dir in $(DIRS); do make -C $$dir clean;done
	cd ./testdb;make clean
	@echo "*** CLEAN completed in test ***"

test:
	@echo "FUNCTIONAL TEST Report" > testreport.txt
	@echo "Report Start: \c" >> testreport.txt
	@date >> testreport.txt
	@echo >> testreport.txt
	@echo "Result  Path                  Test Name                                 Phases" >> testreport.txt
	@echo "------  --------------------  ----------------------------------------  ------" >> testreport.txt
	for dir in $(DIRS); do make -C $$dir test;done
	@./totals.sh
	@echo >> testreport.txt
	@echo "Report Complete: \c" >> testreport.txt
	@date >> testreport.txt
	@echo "*** TEST completed in test ***"
	@cat testreport.txt

package:
	for dir in $(DIRS); do make -C $$dir package;done
	@echo "*** PACKAGE completed in test ***"
