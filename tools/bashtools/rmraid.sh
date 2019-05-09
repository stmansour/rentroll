#!/bin/bash
#
# USAGE:
# 	rmraid RAID
#
# SYNOPSIS: Completely remove a rental agreement from the Rentroll db
#
# DESCRIPTION:
#	This script generates and executes all the sql steps needed to
#   completely remove a Rental Agreement from the Rentroll database.
#   It removes all references having anything to do with the RA. Only
#   use this if you really know what you're doing.
#
# params::
#   $1  the RAID to remove

doremove() {
#-----------------------------------
# Delete journal entries where:
#   Type = 1 and ID = ASMID
#   Type = 2 and ID = RCPTID
#   Type = 3 and ID = EXPID
#-----------------------------------

A=$(echo "SELECT ASMID from Assessments WHERE RAID=${RAID};" |${MYSQL} --no-defaults rentroll | awk '{if ($1 != "ASMID") print $1;}' | paste -s -d, -)
R=$(echo "SELECT RCPTID from Receipt WHERE RAID=${RAID};"    |${MYSQL} --no-defaults rentroll | awk '{if ($1 != "RCPTID") print $1;}' | paste -s -d, -)
E=$(echo "SELECT EXPID from Expense WHERE RAID=${RAID};"     |${MYSQL} --no-defaults rentroll | awk '{if ($1 != "EXPID") print $1;}' | paste -s -d, -)

delA=""
if [ "x${A}" != "x" ]; then
    delA="DELETE FROM Journal WHERE Type=1 AND ID IN (${A});"
fi

delB=""
if [ "x${R}" != "x" ]; then
    delB="DELETE FROM Journal WHERE Type=2 AND ID IN (${R});"
fi

delC=""
if [ "x${E}" != "x" ]; then
    delC="DELETE FROM Journal WHERE Type=3 AND ID IN ();"
fi

cat >RMRAID << EOF
${delA}
${delB}
${delC}
DELETE FROM LedgerEntry WHERE RAID=${RAID};
DELETE FROM RentalAgreementPayors WHERE RAID=${RAID};
DELETE FROM Assessments WHERE RAID=${RAID};
DELETE FROM ReceiptAllocation WHERE RAID=${RAID};
DELETE FROM RentalAgreement where RAID=${RAID};
EOF

${MYSQL} --no-defaults rentroll < RMRAID
}

RAID=${1}
if [ "x${RAID}" = "x" ]; then
    echo "You must supply the RAID you want to remove."
    exit 0
fi

MYSQL=$(which mysql)
if [ "x${MYSQL}" = "x" ]; then
    echo "Could not find mysql executable!!"
fi

doremove
