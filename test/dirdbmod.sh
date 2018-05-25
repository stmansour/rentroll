#!/bin/bash

#==========================================================================
#  This is just like dbmod.sh only it's for directory (phonebook) dbs
#
#==========================================================================

MODFILE="dbqqqmods.sql"
MYSQL="mysql --no-defaults"
MYSQLDUMP="mysqldump --no-defaults"
DBNAME="accord"

#=====================================================
#  History of db mods
#=====================================================

# May 26, 2018 - found all the rentroll tables in the accord db!!??
# DROP TABLE IF EXISTS AR;
# DROP TABLE IF EXISTS AssessmentTax;
# DROP TABLE IF EXISTS Assessments;
# DROP TABLE IF EXISTS AvailabilityTypes;
# DROP TABLE IF EXISTS Building;
# DROP TABLE IF EXISTS Business;
# DROP TABLE IF EXISTS BusinessAssessments;
# DROP TABLE IF EXISTS BusinessPaymentTypes;
# DROP TABLE IF EXISTS CommissionLedger;
# DROP TABLE IF EXISTS CustomAttr;
# DROP TABLE IF EXISTS CustomAttrRef;
# DROP TABLE IF EXISTS DemandSource;
# DROP TABLE IF EXISTS Deposit;
# DROP TABLE IF EXISTS DepositMethod;
# DROP TABLE IF EXISTS DepositPart;
# DROP TABLE IF EXISTS Depository;
# DROP TABLE IF EXISTS Expense;
# DROP TABLE IF EXISTS FlowPart;
# DROP TABLE IF EXISTS GLAccount;
# DROP TABLE IF EXISTS Invoice;
# DROP TABLE IF EXISTS InvoiceAssessment;
# DROP TABLE IF EXISTS InvoicePayor;
# DROP TABLE IF EXISTS Journal;
# DROP TABLE IF EXISTS JournalAllocation;
# DROP TABLE IF EXISTS JournalAudit;
# DROP TABLE IF EXISTS JournalMarker;
# DROP TABLE IF EXISTS JournalMarkerAudit;
# DROP TABLE IF EXISTS LeadSource;
# DROP TABLE IF EXISTS LedgerAudit;
# DROP TABLE IF EXISTS LedgerEntry;
# DROP TABLE IF EXISTS LedgerMarker;
# DROP TABLE IF EXISTS LedgerMarkerAudit;
# DROP TABLE IF EXISTS MRHistory;
# DROP TABLE IF EXISTS NoteList;
# DROP TABLE IF EXISTS NoteType;
# DROP TABLE IF EXISTS Notes;
# DROP TABLE IF EXISTS OtherDeliverables;
# DROP TABLE IF EXISTS PaymentType;
# DROP TABLE IF EXISTS Payor;
# DROP TABLE IF EXISTS Prospect;
# DROP TABLE IF EXISTS RatePlan;
# DROP TABLE IF EXISTS RatePlanOD;
# DROP TABLE IF EXISTS RatePlanRef;
# DROP TABLE IF EXISTS RatePlanRefRTRate;
# DROP TABLE IF EXISTS RatePlanRefSPRate;
# DROP TABLE IF EXISTS Receipt;
# DROP TABLE IF EXISTS ReceiptAllocation;
# DROP TABLE IF EXISTS Rentable;
# DROP TABLE IF EXISTS RentableMarketRate;
# DROP TABLE IF EXISTS RentableSpecialty;
# DROP TABLE IF EXISTS RentableSpecialtyRef;
# DROP TABLE IF EXISTS RentableStatus;
# DROP TABLE IF EXISTS RentableTypeRef;
# DROP TABLE IF EXISTS RentableTypeTax;
# DROP TABLE IF EXISTS RentableTypes;
# DROP TABLE IF EXISTS RentableUsers;
# DROP TABLE IF EXISTS RentalAgreement;
# DROP TABLE IF EXISTS RentalAgreementPayors;
# DROP TABLE IF EXISTS RentalAgreementPets;
# DROP TABLE IF EXISTS RentalAgreementRentables;
# DROP TABLE IF EXISTS RentalAgreementTax;
# DROP TABLE IF EXISTS RentalAgreementTemplate;
# DROP TABLE IF EXISTS SLString;
# DROP TABLE IF EXISTS StringList;
# DROP TABLE IF EXISTS SubAR;
# DROP TABLE IF EXISTS TWS;
# DROP TABLE IF EXISTS Task;
# DROP TABLE IF EXISTS TaskDescriptor;
# DROP TABLE IF EXISTS TaskList;
# DROP TABLE IF EXISTS TaskListDefinition;
# DROP TABLE IF EXISTS Tax;
# DROP TABLE IF EXISTS TaxRate;
# DROP TABLE IF EXISTS Transactant;
# DROP TABLE IF EXISTS User;
# DROP TABLE IF EXISTS Vehicle;

#=====================================================
#  Put modifications to schema in the lines below
#=====================================================
cat >${MODFILE} <<EOF

EOF

#=====================================================
#  Put dir/sqlfilename in the list below
#=====================================================
declare -a dbs=(
	setup/accord.sql
)

for f in "${dbs[@]}"
do
    if [ -f ${f} ]; then
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
