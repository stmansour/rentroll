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

./newbiz -b nb.csv -a asmttype.csv -R rt.csv -u custom.csv -s specialties.csv -D bldg.csv -r rentable.csv -p people.csv -T rat.csv -C ra.csv -E pets.csv -c coa.csv -A asmt.csv -P pmt.csv -e rcpt.csv -U assigncustom.csv >log 2>&1

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

dotest "x"  "-b nb.csv"           "PHASE  1: New Businesses...  " "select BID,DES,Name,DefaultRentalPeriod,ParkingPermitInUse,LastModBy from business;"
dotest "y"  "-a asmttype.csv"     "PHASE  2: Assessment Types...  " "select Name,Description,LastModBy from assessmenttypes;"
dotest "z"  "-R rt.csv"           "PHASE  3: Rentable Types...  " "select RTID,BID,Style,Name,RentalPeriod,Proration,ManageToBudget,LastModBy from rentabletypes;"
dotest "w"  "-R rt.csv"           "PHASE  4: Rentable Market Rates...  " "select * from rentablemarketrate;"
dotest "v"  "-s specialties.csv"  "PHASE  5: Rentable Specialty Types...  " "select * from rentablespecialtytypes;"
dotest "u"  "-D bldg.csv"         "PHASE  6: Buildings...  " "select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from building;"
dotest "t"  "-r rentable.csv"     "PHASE  7: Rentables...  " "select RID,RTID,BID,Name,AssignmentTime,RentalPeriodDefault,RentalPeriod,LastModBy from rentable;"
dotest "s"  "-p people.csv"       "PHASE  8: Transactants...  " "select TCID,RENTERID,PID,PRSPID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from transactant;"
dotest "r"  "-p people.csv"       "PHASE  9: Renters...  " "select RENTERID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureRenter,Industry,Source from renter;"
dotest "q"  "-p people.csv"       "PHASE 10: Payors...  " "select PID,TCID,CreditLimit,TaxpayorID,AccountRep,LastModBy from payor;"
dotest "p"  "-p people.csv"       "PHASE 11: Prospects...  " "select PRSPID,TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModBy from prospect;"
dotest "o"  "-T rat.csv"          "PHASE 12: Rental Agreement Templates...  " "select RATID,RentalTemplateNumber,RentalAgreementType,LastModBy from rentalagreementtemplate;"
dotest "n"  "-C ra.csv"           "PHASE 13: Rental Agreements...  " "select RAID,RATID,BID,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModBy from rentalagreement;"
dotest "n1" "-E pet.csv"          "PHASE 13s: Pets...  " "select PETID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from agreementpets;"
dotest "m"  "-C ra.csv"           "PHASE 14: Agreement Rentables...  " "select * from agreementrentables;"
dotest "l"  "-C ra.csv"           "PHASE 15: Agreement Payors...  " "select * from agreementpayors;"
dotest "k"  "-c coa.csv"          "PHASE 16a: Chart of Accounts...  " "select LID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,LastModBy from ledger;"
dotest "k1" "-c coa.csv"          "PHASE 16b: Ledger Markers...  " "select LMID,LID,BID,DtStart,DtStop,Balance,State,LastModBy from ledgermarker;"
dotest "j"  "-A asmt.csv"         "PHASE 17: Assessments...  " "select ASMID,BID,RID,ASMTID,RAID,Amount,Start,Stop,RentalPeriod,ProrationMethod,AcctRule,Comment,LastModBy from assessments;"
dotest "i"  "-P pmt.csv"          "PHASE 18: Payment types...  " "select PMTID,BID,Name,Description,LastModBy from paymenttypes;"
dotest "h"  "-e rcpt.csv"         "PHASE 19: Payment allocations...  " "select * from receiptallocation order by Amount ASC;"
dotest "g"  "-e rcpt.csv"         "PHASE 20: Receipts... " "select RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy from receipt;"
dotest "f"  "-u custom.csv"       "PHASE 21: CustomAttributes... " "select CID,Type,Name,Value,LastModBy from customattr;"
dotest "e"  "-U assigncustom.csv" "PHASE 22: CustomAttributes AssignmentTime... " "select * from customattrref;"

echo -n "PHASE x: Log file check...  "
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