#!/bin/bash
RRBIN="../../tmp/rentroll"
SCRIPTLOG="f.log"
APP="${RRBIN}/rentroll -A"
MYSQLOPTS=""
UNAME=$(uname)

if [ "${UNAME}" == "Darwin" -o "${IAMJENKINS}" == "jenkins" ]; then
	MYSQLOPTS="--no-defaults"
fi


########################################
# start with a clean database
########################################
${RRBIN}/rrnewdb

./newbiz -b nb.csv -a asmttype.csv -R rt.csv -u custom.csv -s specialties.csv -D bldg.csv -p people.csv -r rentable.csv -T rat.csv -C ra.csv -E pets.csv -c coa.csv -A asmt.csv -P pmt.csv -e rcpt.csv -U assigncustom.csv >log 2>&1

########################################
# dotest()
#	Parameters:
# 		$1 = base file name
#		$2 = app options to reproduce
# 		$3 = title
# 		$4 = mysql select statement
########################################
dotest () {
cat >xxqq <<EOF
use rentroll;
${4}
EOF
	echo -n $3
	mysql --no-defaults <xxqq >${1}
	if [ ! -f ${1}.gold -o ! -f ${1} ]; then
		echo "Missing file: two files are required for checking this phase: ${1}.gold and ${1}"
		exit 1
	fi
	UDIFFS=$(diff ${1} ${1}.gold | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else
		echo "FAILED...   if correct:  mv ${1} ${1}.gold"
		echo "Command to reproduce:  ./newbiz ${2}"
		echo "Differences in ${1} are as follows:"
		diff ${1}.gold ${1}
		exit 1
	fi
}

dotest "x"  "-b nb.csv"           "PHASE  1: New Businesses...  " "select BID,BUD,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModBy from Business;"
dotest "y"  "-a asmttype.csv"     "PHASE  2: Assessment Types...  " "select Name,Description,LastModBy from AssessmentTypes;"
dotest "z"  "-R rt.csv"           "PHASE  3: Rentable Types...  " "select RTID,BID,Style,Name,RentCycle,Proration,ManageToBudget,LastModBy from RentableTypes;"
dotest "w"  "-R rt.csv"           "PHASE  4: Rentable Market Rates...  " "select * from RentableMarketrate;"
dotest "v"  "-s specialties.csv"  "PHASE  5: Rentable Specialty Types...  " "select * from RentableSpecialtyType;"
dotest "u"  "-D bldg.csv"         "PHASE  6: Buildings...  " "select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
dotest "t"  "-r rentable.csv"     "PHASE  7: Rentables...  " "select RID,BID,Name,AssignmentTime,LastModBy from Rentable;"
dotest "t1" "-r rentable.csv"     "PHASE  8: RentableTypeRef...  " "select RID,RTID,RentCycle,ProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
dotest "t2" "-r rentable.csv"     "PHASE  9: RentableStatus...  " "select RID,Status,DtStart,DtStop,LastModBy from RentableStatus;"
dotest "s"  "-p people.csv"       "PHASE 10: Transactants...  " "select TCID,USERID,PID,PRSPID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
dotest "r"  "-p people.csv"       "PHASE 11: Users...  " "select USERID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureUser,Industry,Source from User;"
dotest "q"  "-p people.csv"       "PHASE 12: Payors...  " "select PID,TCID,CreditLimit,TaxpayorID,AccountRep,LastModBy from Payor;"
dotest "p"  "-p people.csv"       "PHASE 13: Prospects...  " "select PRSPID,TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModBy from Prospect;"
dotest "o"  "-T rat.csv"          "PHASE 14: Rental Agreement Templates...  " "select RATID,BID,RentalTemplateNumber,LastModBy from RentalAgreementTemplate;"
dotest "n"  "-C ra.csv"           "PHASE 15: Rental Agreements...  " "select RAID,RATID,BID,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModBy from RentalAgreement;"
dotest "n1" "-E pet.csv"          "PHASE 16: Pets...  " "select PETID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from RentalAgreementPets;"
dotest "m"  "-C ra.csv"           "PHASE 17: Agreement Rentables...  " "select * from RentalAgreementRentables;"
dotest "l"  "-C ra.csv"           "PHASE 18: Agreement Payors...  " "select * from RentalAgreementPayors;"
dotest "k"  "-c coa.csv"          "PHASE 19: Chart of Accounts...  " "select LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModBy from Ledger;"
dotest "k1" "-c coa.csv"          "PHASE 20: Ledger Markers...  " "select LMID,LID,BID,DtStart,DtStop,Balance,State,LastModBy from LedgerMarker;"
dotest "j"  "-A asmt.csv"         "PHASE 21: Assessments...  " "select ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
dotest "i"  "-P pmt.csv"          "PHASE 22: Payment types...  " "select PMTID,BID,Name,Description,LastModBy from PaymentTypes;"
dotest "h"  "-e rcpt.csv"         "PHASE 23: Payment allocations...  " "select * from ReceiptAllocation order by Amount ASC;"
dotest "g"  "-e rcpt.csv"         "PHASE 24: Receipts... " "select RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy from Receipt;"
dotest "f"  "-u custom.csv"       "PHASE 25: CustomAttributes... " "select CID,Type,Name,Value,LastModBy from CustomAttr;"
dotest "e"  "-U assigncustom.csv" "PHASE 26: CustomAttributes AssignmentTime... " "select * from CustomAttrRef;"

echo -n "PHASE FINAL: Log file check...  "
if [ ! -f log.gold -o ! -f log ]; then
	echo "Missing file -- Required files for this check: log.gold and log"
	exit 1
fi
declare -a out_filters=(
	's/(20[1-4][0-9]\/[0-1][0-9]\/[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9] )(.*)/$2/'	
)
cp log.gold ll.g
cp log llog
for f in "${out_filters[@]}"
do
	perl -pe "$f" ll.g > x1; mv x1 ll.g
	perl -pe "$f" llog > y1; mv y1 llog
done
UDIFFS=$(diff llog ll.g | wc -l)
if [ ${UDIFFS} -eq 0 ]; then
	echo "PASSED"
	rm -f ll.g llog
else
	echo "FAILED...   if correct:   mv log log.gold"
	echo "Differences are as follows:"
	diff ll.g llog
	exit 1
fi

echo "NEWBIZ TESTS PASSED"
exit 0