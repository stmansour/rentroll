#!/bin/bash
pushd ../../db/schema;make newdb;popd
./newbiz -b nb.csv -a asmt.csv -R rt.csv -s specialties.csv -D bldg.csv -r rentable.csv -p people.csv -T rat.csv -C ra.csv >log 2>&1


# $1 = base file name
# $2 = title
# $3 = mysql select statement
dotest () {
cat >xxqq <<EOF
use rentroll;
${3}
EOF
	echo -n $2
	mysql --no-defaults <xxqq >${1}
	if [ ! -f ${1}.gold -o ! -f ${1} ]; then
		echo "Missing file: two files are required for checking this phase: ${1}.gold and ${1}"
		exit 1
	fi
	UDIFFS=$(diff ${1} ${1}.gold | wc -l)
	if [ ${UDIFFS} -eq 0 ]; then
		echo "PASSED"
	else
		echo "FAILED:  differences in ${1} are as follows:"
		diff ${1}.gold ${1}
		exit 1
	fi
}

dotest "x" "PHASE  1: New Businesses...  " "select BID,DES,Name,DefaultOccupancyType,ParkingPermitInUse,LastModBy from business;"
dotest "y" "PHASE  2: Assessment Types...  " "select Name,Description,LastModBy from assessmenttypes;"
dotest "z" "PHASE  3: Rentable Types...  " "select RTID,BID,Style,Name,Frequency,Proration,Report,ManageToBudget,LastModBy from rentabletypes;"
dotest "w" "PHASE  4: Rentable Market Rates...  " "select RTID,MarketRate,DtStop from rentablemarketrate;"
dotest "v" "PHASE  5: Rentable Specialty Types...  " "select * from rentablespecialtytypes;"
dotest "u" "PHASE  6: Buildings...  " "select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from building;"
dotest "t" "PHASE  7: Rentables...  " "select RID,RTID,BID,Name,Assignment,Report,DefaultOccType,OccType,LastModBy from rentable;"
dotest "s" "PHASE  8: Transactants...  " "select TCID,TID,PID,PRSPID,FirstName,MiddleName,LastName,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from transactant;"
dotest "r" "PHASE  9: Tenants...  " "select TID,TCID,Points,CarMake,CarModel,CarColor,CarYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,AccountRep,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,ElibigleForFutureOccupancy,Industry,Source,InvoicingCustomerNumber from tenant;"
dotest "q" "PHASE 10: Payors...  " "select PID,TCID,CreditLimit,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,LastModBy from payor;"
dotest "p" "PHASE 11: Prospects...  " "select PRSPID,TCID,ApplicationFee,LastModBy from prospect;"
dotest "o" "PHASE 12: Rental Agreement Templates...  " "select RATID,ReferenceNumber,RentalAgreementType,LastModBy from rentalagreementtemplate;"
dotest "n" "PHASE 13: Rental Agreements...  " "select RAID,RATID,BID,PrimaryTenant,RentalStart,RentalStop,Renewal,SpecialProvisions,LastModBy from rentalagreement;"
dotest "m" "PHASE 14: Agreement Rentables...  " "select * from agreementrentables;"
dotest "l" "PHASE 15: Agreement Payors...  " "select * from agreementpayors;"


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
	echo "FAILED:  differences are as follows:"
	diff ll.g llog
	exit 1
fi

echo "NEWBIZ TESTS PASSED"
exit 0