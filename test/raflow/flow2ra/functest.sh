#!/bin/bash

TESTNAME="Flow2RA"
TESTSUMMARY="Test Flow data to permanent tables"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 10/24/2018"
DB2LOADED=0

startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  An existing rental agreement (RAID=1) is being amended. This test
#  verifies that all assessments from RAID=1 to the new RAID (24) are correct.
#  It also verifies that payments are properly filtered. For example,
#  if the original agreement there is a Security Deposit request in
#  September. This security deposit is not in the fees for the amended
#  Rental Agreement, so it should be reversed in the old rental agreement
#
#  Scenario:
#  The flow for RAID 1 is updated to the active state causing an amended
#  Rental Agreement (RAID=24) to be created.
#  RAID  1 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#  RAID 24 - AgreementStart = 8/20/2018,  AgreementStop = 3/1/2020
#            The flow used to create RAID 24 has no links between its fees and
#            the assessments in RAID 1. So, the handling tests how "unlinked"
#            assessments are handled when amending a rental agreement.
#
#  Expected Results:
#   1.  All RAID 1 recurring assessment definitions that overlap the period
#       8/8/2018 - 3/1/2020 must have their stop date set to 8/8/201
#   2.  The RAID 1 rent assessment has already occured, and it has been paid.
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  There is a Security Deposit assessment for RAID 1 due on 9/20 in the
#       old rental agreement. It is not in the fees list for the RefNo, so it
#       should be reversed in RAID 1 and not present in RAID 24
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    RAID1REFNO="T7LYN5K18Z7F756KE64C"
    RAIDAMENDEDID="24"

    # Send the command to change the flow to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAID1REFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "a0"  "WebService--Action-setTo-ACTIVE"

    # Generate an assessment report from Aug 1 to Oct 1. The security deposit
    # assessment for RAID 1 should no longer be present
    docsvtest "a1" "-G ${BUD} -g 8/1/18,10/1/18 -L 11,${BUD}" "Assessments-2018-AUG"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%228%2F31%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "a2"  "PayorStatement--StmtInfo"
fi

#------------------------------------------------------------------------------
#  TEST b
#  This is just like test a except that the $4500 security deposit assessment
#  from the origin RA (RAID 2) is kept.  Since its time frame falls into that
#  of the amended Rental Agreement, it becomes part of that rental agreement.
#
#  Scenario:
#  RAID  2 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#  RAID 25 - AgreementStart = 8/20/2018,  AgreementStop = 3/1/2020
#            Verify that the Security Deposit on 9/20 is linked to the new
#            rental agreement.
#
#  Expected Results:
#   1.  ASMID 402 (which was charged to RAID 1) should be reversed and a new
#       one should be created (ASMID 412) associated with the amended RAID (25)
#       8/8/2018 - 3/1/2020 must have their stop date set to 8/8/201
#   2.  The RAID 1 rent assessment has already occured, and it has been paid.
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  There is a Security Deposit assessment (ASMID=402) due on 9/20 in the
#       old rental agreement. It is not in the fees list for the RefNo, so it
#       should be reversed
#------------------------------------------------------------------------------
TFILES="b"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    RAIDREFNO="NZXY8FS6NHJ34N383950"
    RAIDAMENDEDID="25"

    # Send the command to change the RefNo to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "b0"  "WebService--Action-setTo-ACTIVE"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "b1"  "PayorStatement--StmtInfo"
fi

#------------------------------------------------------------------------------
#  TEST d
#  This tests updating a rental agreement with the term dates the same as
#  the original rental agreement. In this case we want the old rental agreement
#  to start and stop on the same day (so there is a record of this agreement)
#  and the amended agreement to start on the same day and end on the date that
#  the parent agreement eneded. For good measure, I made one change to a pet:
#  the cat named BatMan is renamed to Crappy.
#
#  Scenario:
#  RAID  4 - AgreementStart = 2/13/2018,  AgreementStop = 3/1/2020
#            Change pet BatMan to Crappy
#  RAID 25 - AgreementStart = 8/20/2018,  AgreementStop = 3/1/2020
#            Pet should be named Crappy
#
#  Expected Results:
#   1.  RAID 4 should start and stop on 2/13.
#   2.  All assessments associated with RA
#       Same for the RAID 1 pet rent. The assessments must be reversed and the
#       payments must become available.
#   3.  There is a Security Deposit assessment (ASMID=402) due on 9/20 in the
#       old rental agreement. It is not in the fees list for the RefNo, so it
#       should be reversed
#------------------------------------------------------------------------------
TFILES="a"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    RAIDREFNO="7K9B2FD9293R0RN67PSE"
    RAIDAMENDEDID="25"

    # Send the command to change the RefNo to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "d0"  "WebService--Action-setTo-ACTIVE"

    # Generate a payor statement -- ensure that 2 RAs are there and have correct
    # info.
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
    dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "d1"  "PayorStatement--StmtInfo"
fi

#------------------------------------------------------------------------------
#  TEST c
#  Validate that a new owner of a pet is properly handled in the TBind
#  records.  Also validate that new pets are handled properly for both
#  TBind and newly created Occupant who becomes their contact point.
#  And do the same tests for new owners of a vehicle and a new vehicle
#
#  Scenario:
#  RAID  1 - Add a new user, Sally. Add a new pet, Rocky. Add a new car.
#            Make Sally the contact for both pets and both vehicles going
#            forward.
#
#  Expected Results:
#   1.  Pet 2 is created and TCID 2 is the contact.
#   2.  Vehicle 2 is created and TCID 2 is the contact.
#   3.  TBind record for Pet 1 will be split at 8/23/2018.  TCID 1 was the
#       contact person before the split.  TCID 2 is the contact going forward.
#   4.  TBind record for Vehicle 1 will be split at 8/23/2018. TCID 1 was the
#       contact person before the split.  TCID 2 is the contact going forward.
#------------------------------------------------------------------------------
TFILES="c"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"
    echo "Create new database..."
    mysql --no-defaults rentroll < rrsm1.sql
    DB2LOADED=1

    RAIDREFNO="8VMAH0O53D6R4W25P5V0"

    # Send the command to change the RefNo to Active:
    echo "%7B%22UserRefNo%22%3A%22${RAIDREFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "c0"  "WebService--Action-setTo-ACTIVE"

    # make sure the TBinds are correct
    mysqlverify "c1" "TBind-Pets" "SELECT TBID,BID,SourceElemType,SourceElemID,AssocElemType,AssocElemID,DtStart,DtStop,FLAGS FROM TBind;"

    # make sure the transactants are correct
    mysqlverify "c2" "flow2ra-Transactants" "SELECT TCID,BID,PreferredName,LastName FROM Transactant;"
fi

# import rr.sql again to test update existing RA
stopRentRollServer
echo "RENTROLL SERVER STOPPED"

if [ "${DB2LOADED}" = "1" ]; then
    echo "Create new database..."
    mysql --no-defaults rentroll < rr.sql
fi

RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#----------------------------------------------------------------------------------
# TEST z
# This test is for verifying update RAflow of existing RAFlow
#
# Scenario:
# Edit one existing RA Application from its view mode only.
# Update Pet/Vehicle/People information fields value
# After updating information, move RAApplication to state 'Complete Move-In'
#
# Expected Result:
# Check same RA Application flow's data. It must be match with the updated information
# Outdated RA Application must be terminated
#---------------------------------------------------------------------------------
TFILES="z"
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    echo "Test ${TFILES}"

    RAID3REFNO="A02QXP7Z2D5SC8U74Y79"

# send command to Edit existing Rental Agreement with RAID: 3
#echo "%7B%22cmd%22%3A%22get%22%2C%22FlowType%22%3A%22RA%22%2C%22RAID%22%3A3%2C%22UserRefNo%22%3Anull%2C%22Version%22%3A%22refno%22%7D" > request
#dojsonPOST "http://localhost:8270/v1/flow/1/0" "request" "z0" "Rental Agreement--Edit--RAID:3"

# edit pets information
encodeRequest '{"cmd":"save","FlowType":"RA","FlowID":3,"FlowPartKey":"pets","BID":1,"Data":[{"Fees":[{"recid":1,"TMPASMID":3,"ARID":24,"ASMID":38,"ARName":"Pet Rent","ContractAmount":3.55,"RentCycle":0,"ProrationCycle":0,"Start":"8/21/2018","Stop":"8/21/2018","AtSigningPreTax":0,"SalesTax":0,"TransOccTax":0,"Comment":"prorated for 11 of 31 days","w2ui":{"class":"","style":{}}},{"recid":2,"TMPASMID":4,"ARID":24,"ASMID":38,"ARName":"Pet Rent","ContractAmount":10,"RentCycle":6,"ProrationCycle":4,"Start":"9/1/2018","Stop":"3/1/2020","AtSigningPreTax":0,"SalesTax":0,"TransOccTax":0,"Comment":"","w2ui":{"class":"","style":{}}}],"Name":"Bosky","Type":"Cat","Breed":"Neapolitan Mastiff","Color":"Black","PETID":2,"DtStop":"3/1/2020","Weight":11,"DtStart":"8/21/2018","TMPTCID":1,"TMPPETID":1}]}' > request
dojsonPOST "http://localhost:8270/v1/flow/1/3/" "request" "z0" "Rental Agreement--RAID:3--Update Pet Information"

# edit vehicle information
encodeRequest '{"cmd":"save","FlowType":"RA","FlowID":3,"FlowPartKey":"vehicles","BID":1,"Data":[{"VID":3,"VIN":"2BPDY2OYZM4YMRTC","Fees":[],"DtStop":"3/1/2020","TMPVID":1,"DtStart":"8/22/2018","TMPTCID":1,"VehicleMake":"Suzuki","VehicleType":"Bike","VehicleYear":2008,"VehicleColor":"White","VehicleModel":"12BBT","LicensePlateState":"GJ","LicensePlateNumber":"1T9RW28","ParkingPermitNumber":"7916444"}]}' > request
dojsonPOST "http://localhost:8270/v1/flow/1/3/" "request" "z1" "Rental Agreement--RAID:3--Update Vehicle Information"

# edit people information
encodeRequest '{"cmd":"save","FlowType":"RA","FlowID":3,"FlowPartKey":"people","BID":1,"Data":[{"City":"Denton","TCID":3,"State":"ME","Points":0,"Address":"66789 Shore","Comment":"","Country":"USA","Evicted":false,"TMPTCID":1,"Website":"","Address2":"","Industry":344,"IsRenter":true,"LastName":"Bosamiya","CellPhone":"(314) 860-0587","Convicted":false,"FirstName":"Akshay","IsCompany":false,"WorkPhone":"(607) 954-3966","Bankruptcy":false,"EvictedDes":"","IsOccupant":true,"MiddleName":"","Occupation":"the hygiene service assistant (hygiene service assistant)","PostalCode":"28162","TaxpayorID":"08114320","CompanyCity":"Elizabeth","CompanyName":"Western Digital Inc","CreditLimit":15343,"DateofBirth":"2/15/1957","GrossIncome":22028,"IsGuarantor":false,"SourceSLSID":17,"CompanyEmail":"WElizabeth7089@aol.com","CompanyPhone":"(255) 339-0248","CompanyState":"VA","ConvictedDes":"","PrimaryEmail":"akshay@yopmail.com","PriorAddress":"66787 Hampton, Cambridge, NC 31445","SpecialNeeds":"","BankruptcyDes":"","PreferredName":"Denisha","CompanyAddress":"71590 Fifth","CurrentAddress":"81062 Wood, Santa Rosa, FL 11211","DriversLicense":"D7626933","SecondaryEmail":"akshay@yopmail.com","OtherPreferences":"","ThirdPartySource":"Stacia Robertson","CompanyPostalCode":"33274","PriorLandLordName":"Kali Graves","EligibleFutureUser":false,"CurrentLandLordName":"Eugenia Dunn","EligibleFuturePayor":true,"EmergencyContactName":"Reyna Ramirez","PriorLandLordPhoneNo":"(823) 260-2871","PriorReasonForMoving":117,"AlternateEmailAddress":"14557 Lakeview,Jefferson,GA 24728","EmergencyContactEmail":"RRamirez8989@bdiddy.com","CurrentLandLordPhoneNo":"(510) 858-4871","CurrentReasonForMoving":129,"PriorLengthOfResidency":"8 years 6 months","EmergencyContactAddress":"84390 Rhode Island,Durham,AK 55089","CurrentLengthOfResidency":"2 years 7 months","EmergencyContactTelephone":"(506) 681-2584","recid":1,"w2ui":{"class":"","style":{}},"BID":1,"BUD":"REX","NLID":"","CreateBy":"","CreateTS":"","LastModBy":"","LastModTime":""}]}' > request
dojsonPOST "http://localhost:8270/v1/flow/1/3/" "request" "z2" "Rental Agreement--RAID:3--Update People Information"

    # validate updated flow's data
    encodeRequest '{"cmd":"get","FlowID":3}' > request
    dojsonPOST "http://localhost:8270/v1/validate-raflow/1/3/" "request" "z3" "Rental Agreement--RAID:3--Validate update RAFlow"

    # RAAction: Complete to Move In
    echo "%7B%22UserRefNo%22%3A%22${RAID3REFNO}%22%2C%22RAID%22%3A3%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/raactions/1/3/" "request" "z4" "Rental Agreement--RAID:3--Complete Move In"

    # Get updated flow: It'll have new RAID: 24
    # TODO: z5.gold must require to update after fixing bug in the code Bug: Duplicate entry of the pets/vehicles
    encodeRequest '{"cmd":"get","UserRefNo":null,"RAID":24,"Version":"raid","FlowType":"RA"}' > request
    dojsonPOST "http://localhost:8270/v1/flow/1/0/" "request" "z5" "Rental Agreement--RAID:24--Get updated flow"

    # Check old RAID:3 Rental agreement must be terminated due to update rental agreement
    # TODO: z6.gold must require to update after fixing bug in the code Bug: Duplicate entry of the pets/vehicles
    encodeRequest '{"cmd":"get","UserRefNo":null,"RAID":3,"Version":"raid","FlowType":"RA"}' > request
    dojsonPOST "http://localhost:8270/v1/flow/1/0/" "request" "z6" "Rental Agreement--RAID:3--Terminated"
fi


stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
