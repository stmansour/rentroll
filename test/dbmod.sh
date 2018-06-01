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
DBNAME="rentroll"

#=====================================================
#  History of db mods
#=====================================================
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

# # 13 Dec, 2017
# ALTER TABLE CustomAttrRef ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE CustomAttrRef ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentalAgreementRentables ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentalAgreementRentables ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentalAgreementPayors ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentalAgreementPayors ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentableUsers ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentableUsers ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentalAgreementTax ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentalAgreementTax ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE CommissionLedger ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE CommissionLedger ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RatePlanRefRTRate ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RatePlanRefRTRate ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RatePlanRefSPRate ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RatePlanRefSPRate ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RatePlanOD ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RatePlanOD ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE OtherDeliverables ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE OtherDeliverables ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentableMarketRate ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentableMarketRate ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentableTypeTax ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentableTypeTax ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE RentableSpecialty ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE RentableSpecialty ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE AvailabilityTypes ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE AvailabilityTypes ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE BusinessAssessments ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE BusinessAssessments ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE BusinessPaymentTypes ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE BusinessPaymentTypes ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE InvoiceAssessment ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE InvoiceAssessment ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE InvoicePayor ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE InvoicePayor ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE JournalAllocation ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE JournalAllocation ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE JournalAudit DROP COLUMN ModTime;
# ALTER TABLE JournalAudit ADD CreateTS TIMESTAMP DEFAULT CURRENT_TIMESTAMP AFTER UID;
# ALTER TABLE JournalAudit ADD CreateBy BIGINT NOT NULL DEFAULT 0 AFTER CreateTS;
# ALTER TABLE JournalAudit ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE JournalAudit ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE JournalMarkerAudit DROP COLUMN ModTime;
# ALTER TABLE JournalMarkerAudit ADD CreateTS TIMESTAMP DEFAULT CURRENT_TIMESTAMP AFTER UID;
# ALTER TABLE JournalMarkerAudit ADD CreateBy BIGINT NOT NULL DEFAULT 0 AFTER CreateTS;
# ALTER TABLE JournalMarkerAudit ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE JournalMarkerAudit ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE LedgerAudit DROP COLUMN ModTime;
# ALTER TABLE LedgerAudit ADD CreateTS TIMESTAMP DEFAULT CURRENT_TIMESTAMP AFTER UID;
# ALTER TABLE LedgerAudit ADD CreateBy BIGINT NOT NULL DEFAULT 0 AFTER CreateTS;
# ALTER TABLE LedgerAudit ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE LedgerAudit ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;
# ALTER TABLE LedgerMarkerAudit DROP COLUMN ModTime;
# ALTER TABLE LedgerMarkerAudit ADD CreateTS TIMESTAMP DEFAULT CURRENT_TIMESTAMP AFTER UID;
# ALTER TABLE LedgerMarkerAudit ADD CreateBy BIGINT NOT NULL DEFAULT 0 AFTER CreateTS;
# ALTER TABLE LedgerMarkerAudit ADD LastModTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER CreateBy;
# ALTER TABLE LedgerMarkerAudit ADD LastModBy BIGINT NOT NULL DEFAULT 0 AFTER LastModTime;

# # 1 Jan, 2018
# ALTER TABLE rentroll.CustomAttrRef ADD CARID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.RatePlanRefRTRate ADD RPRRTRateID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.RatePlanRefSPRate ADD RPRSPRateID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.RentableSpecialtyRef ADD RSPRefID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.Prospect MODIFY TCID BIGINT NOT NULL;
# ALTER TABLE rentroll.User MODIFY TCID BIGINT NOT NULL;
# ALTER TABLE rentroll.Payor MODIFY TCID BIGINT NOT NULL;
# ALTER TABLE rentroll.InvoiceAssessment ADD InvoiceASMID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;
# ALTER TABLE rentroll.InvoicePayor ADD InvoicePayorID BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY;

# # 15 Feb, 2018
# ALTER TABLE rentroll.Business ADD FLAGS BIGINT NOT NULL DEFAULT 0 AFTER DefaultGSRPC;


# 11-MAR-2018
# CREATE TABLE Task (
#     TID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     TLID BIGINT NOT NULL DEFAULT 0,                             -- the TaskList to which this task belongs
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- Task text
#     Worker VARCHAR(80) NOT NULL DEFAULT '',                     -- Name of the associated work function
#     DtDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',     -- Task Due Date
#     DtPreDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',  -- Pre Completion due date
#     DtDone TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',    -- Task completion Date
#     DtPreDone TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00', -- Task Pre Completion Date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 pre-completion required (if 0 then there is no pre-completion required)
#                                                                 -- 1<<1 PreCompletion done (if 0 it is not yet done)
#                                                                 -- 1<<2 Completion done (if 0 it is not yet done)
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TID)
# );

# CREATE TABLE TaskList (
#     TLID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- TaskList name
#     Cycle BIGINT NOT NULL DEFAULT 0,                            -- recurrence frequency (not editable)
#     DtDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',     -- All tasks in task list are due on this date
#     DtPreDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',  -- All tasks in task list pre-completion date
#     DtDone TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',    -- Task completion Date
#     DtPreDone TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00', -- Task Pre Completion Date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TLID)
# );

# CREATE TABLE TaskListDefinition (
#     TLDID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- TaskList name
#     Cycle BIGINT NOT NULL DEFAULT 0,                            -- recurrence frequency (editable)
#     DtDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',     -- All tasks in task list are due on this date
#     DtPreDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',  -- All tasks in task list pre-completion date
#     DtDone TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',    -- Task completion Date
#     DtPreDone TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00', -- Task Pre Completion Date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TLDID)
# );

# CREATE TABLE TaskDescriptor (
#     TDID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     TLDID BIGINT NOT NULL DEFAULT 0,                            -- the TaskListDefinition to which this taskDescr belongs
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- Task text
#     Worker VARCHAR(80) NOT NULL DEFAULT '',                     -- Name of the associated work function
#     EpochDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',  -- Task Due Date
#     EpochPreDue TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00', -- Pre Completion due date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 pre-completion required (if 0 then there is no pre-completion required)
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TDID)
# );

# # March 12, 2018 -- AWS production mysql server required DATETIME rather than TIMESTAMP for Default val
# ALTER TABLE Task MODIFY DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE Task MODIFY DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE Task MODIFY DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE Task MODIFY DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';

# ALTER TABLE TaskList MODIFY DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskList MODIFY DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskList MODIFY DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskList MODIFY DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';

# ALTER TABLE TaskListDefinition MODIFY DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskListDefinition MODIFY DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskListDefinition MODIFY DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskListDefinition MODIFY DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';

# ALTER TABLE TaskDescriptor MODIFY EpochDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';
# ALTER TABLE TaskDescriptor MODIFY EpochPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00';

# # March 14, 2018
# DROP TABLE IF EXISTS TaskListDefinition;
# CREATE TABLE TaskListDefinition (
#     TLDID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- TaskList name
#     Cycle BIGINT NOT NULL DEFAULT 0,                            -- recurrence frequency (editable)
#     Epoch DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',      -- TaskList start Date
#     EpochDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- TaskList Due Date
#     EpochPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00', -- Pre Completion due date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TLDID)
# );

# March 16, 2018
# DROP TABLE IF EXISTS Task;
# CREATE TABLE Task (
#     TID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     TLID BIGINT NOT NULL DEFAULT 0,                             -- the TaskList to which this task belongs
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- Task text
#     Worker VARCHAR(80) NOT NULL DEFAULT '',                     -- Name of the associated work function
#     DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',      -- Task Due Date
#     DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- Pre Completion due date
#     DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',     -- Task completion Date
#     DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- Task Pre Completion Date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 pre-completion required (if 0 then there is no pre-completion required)
#                                                                 -- 1<<1 PreCompletion done (if 0 it is not yet done)
#                                                                 -- 1<<2 Completion done (if 0 it is not yet done)
#     DoneUID BIGINT NOT NULL DEFAULT 0,                          -- user who marked this task done
#     PreDoneUID BIGINT NOT NULL DEFAULT 0,                       -- user who marked this task predone
#     Comment VARCHAR(2048) NOT NULL DEFAULT '',                  -- any user comments
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TID)
# );

# DROP TABLE IF EXISTS TaskList;
# CREATE TABLE TaskList (
#     TLID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,
#     Name VARCHAR(256) NOT NULL DEFAULT '',                      -- TaskList name
#     Cycle BIGINT NOT NULL DEFAULT 0,                            -- recurrence frequency (not editable)
#     DtDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',      -- All tasks in task list are due on this date
#     DtPreDue DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',   -- All tasks in task list pre-completion date
#     DtDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',     -- Task completion Date
#     DtPreDone DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',  -- Task Pre Completion Date
#     FLAGS BIGINT NOT NULL DEFAULT 0,                            -- 1<<0 - 0 = active, 1 = inactive
#     DoneUID BIGINT NOT NULL DEFAULT 0,                          -- user who marked this task done
#     PreDoneUID BIGINT NOT NULL DEFAULT 0,                       -- user who marked this task predone
#     Comment VARCHAR(2048) NOT NULL DEFAULT '',                  -- any user comments
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY(TLID)
# );

# 23 Mar, 2018
# DROP TABLE IF EXISTS FlowPart;
# CREATE TABLE FlowPart (
#     FlowPartID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,                                                         -- Business id
#     Flow VARCHAR(50) NOT NULL DEFAULT '',                                                  -- for which flow we're storing data ("RA=Rental Agreement Flow")
#     FlowID VARCHAR(50) NOT NULL DEFAULT '',                                                -- unique random flow ID for which we will store relavant json data
#     PartType SMALLINT NOT NULL DEFAULT 0,                                                  -- for which part type ("ASM", "PET", "VEHICLE")
#     Data JSON DEFAULT NULL,                                                                -- JSON Data for each flow type
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was it last updated
#     LastModBy BIGINT NOT NULL DEFAULT 0,                                                   -- who modified it last
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,                                 -- when was it created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                                                    -- who created it
#     PRIMARY KEY(FlowPartID),
#     UNIQUE KEY FlowPartUnique (FlowPartID, BID, FlowID)
# );

# April 4, 2018
# ALTER TABLE TaskListDefinition ADD Comment VARCHAR(2048) NOT NULL DEFAULT '' AFTER FLAGS;
# ALTER TABLE TaskDescriptor ADD Comment VARCHAR(2048) NOT NULL DEFAULT '' AFTER FLAGS;

# 16th March, 2018
# ALTER TABLE Rentable ADD Comment VARCHAR(2048) NOT NULL DEFAULT ''; -- Add Comment textfield to Rentable table

# May 5, 2018
# ALTER TABLE TaskList ADD EmailList VARCHAR(2048) NOT NULL DEFAULT '' AFTER PreDoneUID;

# May 5, 2018
#     Somehow, phonebook schema is getting grafted onto the rentroll database
# DROP TABLE IF EXISTS classes;
# DROP TABLE IF EXISTS companies;
# DROP TABLE IF EXISTS compensation;
# DROP TABLE IF EXISTS counters;
# DROP TABLE IF EXISTS deductionlist;
# DROP TABLE IF EXISTS deductions;
# DROP TABLE IF EXISTS departments;
# DROP TABLE IF EXISTS fieldperms;
# DROP TABLE IF EXISTS jobtitles;
# DROP TABLE IF EXISTS people;
# DROP TABLE IF EXISTS roles;
# DROP TABLE IF EXISTS sessions;

# May 8, 2018
# ALTER TABLE TaskList ADD DtLastNotify DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00' AFTER EmailList;
# ALTER TABLE TaskList ADD DurWait BIGINT NOT NULL DEFAULT 0 AFTER DtLastNotify;
# ALTER TABLE TaskListDefinition ADD EmailList VARCHAR(2048) NOT NULL DEFAULT '' AFTER FLAGS;

# May 9, 2018
# ALTER TABLE TaskList CHANGE DurWait DurWait BIGINT NOT NULL DEFAULT 86400000000000;

# May 11, 2018
# ALTER TABLE TaskList ADD TLDID BIGINT NOT NULL DEFAULT 0 AFTER BID;

# May 14, 2018
# ALTER TABLE TaskList ADD PTLID BIGINT NOT NULL DEFAULT 0 AFTER BID;

# May 16, 2018
# ALTER TABLE RentableTypes ADD ARID BIGINT NOT NULL DEFAULT 0 AFTER FLAGS;

# May 25, 2018
# ALTER TABLE Business ADD ClosePeriodTLID BIGINT NOT NULL DEFAULT 0 AFTER DefaultGSRPC;

# May 25, 2018
# DROP TABLE IF EXISTS FlowPart;
# DROP TABLE IF EXISTS Flow;
# CREATE TABLE Flow (
#     FlowID BIGINT NOT NULL AUTO_INCREMENT,
#     BID BIGINT NOT NULL DEFAULT 0,                                                         -- Business id
#     FlowType VARCHAR(50) NOT NULL DEFAULT '',                                              -- for which flow we're storing data ("RA=Rental Agreement Flow")
#     Data JSON DEFAULT NULL,                                                                -- JSON Data for each flow type
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was it last updated
#     LastModBy BIGINT NOT NULL DEFAULT 0,                                                   -- who modified it last
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,                                 -- when was it created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                                                    -- who created it
#     PRIMARY KEY(FlowID)
# );

# May 28, 2018
# DROP TABLE IF EXISTS FlowPart;

# May 29, 2018
# DROP TABLE IF EXISTS ClosePeriod;
# CREATE TABLE ClosePeriod (
#     CPID BIGINT NOT NULL AUTO_INCREMENT,                        -- Close Period ID
#     BID BIGINT NOT NULL DEFAULT 0,                              -- Business id
#     TLID BIGINT NOT NULL DEFAULT 0,                             -- Task List that was used for close
#     Dt DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',         -- Date/Time of close
#     LastModTime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- when was this record last written
#     LastModBy BIGINT NOT NULL DEFAULT 0,                        -- employee UID (from phonebook) that modified it
#     CreateTS TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,      -- when was this record created
#     CreateBy BIGINT NOT NULL DEFAULT 0,                         -- employee UID (from phonebook) that created this record
#     PRIMARY KEY (CPID)
# );

# Jun 1, 2018
# ALTER TABLE RentalAgreementRentables ADD PRID BIGINT NOT NULL DEFAULT 0 AFTER RID;
# ALTER TABLE RentableTypes DROP COLUMN ManageToBudget;

#=====================================================
#  Put modifications to schema in the lines below
#=====================================================
cat >${MODFILE} <<EOF
EOF

#=====================================================
#  Put dir/sqlfilename in the list below
#=====================================================
declare -a dbs=(
	../tools/dbgen/empty.sql
	acctbal/baltest.sql
	closeperiod/rr.sql
	payorstmt/pstmt.sql
	rfix/rcptfixed.sql
	rfix/receipts.sql
	roller/prodrr.sql
	rr/rr.sql
	tws/rr.sql
	tws2/rrtl.sql
	tws2/moonshine.sql
	tws3/rr.sql
	webclient/accord.sql
	webclient/webclientTest.sql
	websvc1/asmtest.sql
	websvc3/tasks.sql
	workerasm/rex.sql
	workerasm/rr.sql
)

for f in "${dbs[@]}"
do
    if [ -f ${f} ]; then
    	echo "DROP DATABASE IF EXIST ${DBNAME}; create database rentroll"
		echo -n "${f}: loading... "
		${MYSQL} ${DBNAME} < ${f}
		echo -n "updating... "
		${MYSQL} ${DBNAME} < ${MODFILE}
		echo -n "saving... "
		${MYSQLDUMP} ${DBNAME} > ${f}
		echo "done"
    else
		echo "file not found: ${f}"
    fi
done
