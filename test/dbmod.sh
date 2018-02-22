#!/bin/bash

#==========================================================================
#  This script performs SQL schema changes on the test databases that are
#  saved as SQL files in the test directory. It loads them, performs the
#  ALTER commands, then saves the sql file.
#
#  If the test file uses its own database saved as a .sql file, make sure
#  it is listed in the dbs array
#==========================================================================

MODFILE="dbqqqmods.sql"
MYSQL="mysql --no-defaults"
MYSQLDUMP="mysqldump --no-defaults"

#=====================================================
#  Put modifications to schema in the lines below
#=====================================================
cat >${MODFILE} <<EOF
# # Sep 25, 2017
# ALTER TABLE RentalAgreement ADD COLUMN FLAGS BIGINT NOT NULL DEFAULT 0 AFTER RightOfFirstRefusal;
# # Sep 26, 2017
# ALTER TABLE AR ADD COLUMN FLAGS BIGINT NOT NULL DEFAULT 0 AFTER DtStop;
# ALTER TABLE AR ADD COLUMN DefaultAmount DECIMAL(19,4) NOT NULL DEFAULT 0.0 AFTER FLAGS;
# # Sep 27, 2017
# ALTER TABLE Receipt ADD COLUMN RAID BIGINT NOT NULL DEFAULT 0 AFTER DID;
# # Oct 9, 2017
# ALTER TABLE Rentable ADD COLUMN MRStatus SMALLINT NOT NULL DEFAULT 0 AFTER AssignmentTime;
# ALTER TABLE Rentable ADD DtMRStart TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP AFTER MRStatus;
# ALTER TABLE RentableStatus CHANGE Status UseStatus SMALLINT NOT NULL DEFAULT 0;
# ALTER TABLE RentableStatus ADD COLUMN LeaseStatus SMALLINT NOT NULL DEFAULT 0 AFTER UseStatus;
# DROP TABLE IF EXISTS SubAR;
# CREATE TABLE SubAR (
#     SARID BIGINT NOT NULL AUTO_INCREMENT,
#     ARID BIGINT NOT NULL DEFAULT 0,                         -- Which ARID
#     SubARID BIGINT NOT NULL DEFAULT 0,                      -- ARID of the sub-account rule
#     BID BIGINT NOT NULL DEFAULT 0,
#     LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                    -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP DEFAULT CURRENT_TIMESTAMP,           -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                     -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(SARID)
# );
# ALTER TABLE Assessments ADD COLUMN AGRCPTID BIGINT NOT NULL DEFAULT 0 AFTER RPASMID;
# 1 Jan, 2018
# ALTER TABLE rentroll.CustomAttrRef ADD CARID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.RatePlanRefRTRate ADD RPRRTRateID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.RatePlanRefSPRate ADD RPRSPRateID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.RentableSpecialtyRef ADD RSPRefID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.Prospect MODIFY TCID BIGINT NOT NULL;
# ALTER TABLE rentroll.User MODIFY TCID BIGINT NOT NULL;
# ALTER TABLE rentroll.Payor MODIFY TCID BIGINT NOT NULL;
# ALTER TABLE rentroll.InvoiceAssessment ADD InvoiceASMID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.InvoicePayor ADD InvoicePayorID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# 15 Feb, 2018
ALTER TABLE rentroll.Business ADD FLAGS BIGINT NOT NULL DEFAULT 0;
EOF

#=====================================================
#  Put dir/sqlfilename in the list below
#=====================================================
declare -a dbs=(
	'acctbal/baltest.sql'
	'websvc1/asmtest.sql'
	'payorstmt/pstmt.sql'
	'rr/rr.sql'
	'webclient/webclientTest.sql'
	'roller/prodrr.sql'
	'workerasm/rr.sql'
)

for f in "${dbs[@]}"
do
	echo -n "${f}: loading... "
	${MYSQL} rentroll < ${f}
	echo -n "updating... "
	${MYSQL} rentroll < ${MODFILE}
	echo -n "saving... "
	${MYSQLDUMP} rentroll > ${f}
	echo "done"
done
