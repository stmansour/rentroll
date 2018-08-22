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
startRentRollServer

#------------------------------------------------------------------------------
#  TEST a
#  An existing rental agreement (RAID=1) is being amended. This test
#  verifies that all assessments from RAID=1 to the new RAID (2) are correct.
#  It also verifies that payments are properly filtered. For example,
#  I the original agreement there is a Security Deposit request in
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
RAID1REFNO="NZXY8FS6NHJ34N383950"
RAIDAMENDEDID="25"

# Send the command to change the RefNo to Active:
echo "%7B%22UserRefNo%22%3A%22${RAID1REFNO}%22%2C%22RAID%22%3A1%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/" "request" "b0"  "WebService--Action-setTo-ACTIVE"

# Generate a payor statement -- ensure that 2 RAs are there and have correct
# info.
echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2018%22%2C%22searchDtStop%22%3A%229%2F30%2F2018%22%2C%22Bool1%22%3Afalse%7D" > request
dojsonPOST "http://localhost:8270/v1/payorstmt/1/1" "request" "b1"  "PayorStatement--StmtInfo"


# import rr.sql again to test update existing RA
echo "Create new database..."
mysql --no-defaults rentroll < rr.sql

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
#---------------------------------------------------------------------------------

RAID3REFNO="A02QXP7Z2D5SC8U74Y79"

# send command to Edit existing Rental Agreement with RAID: 3
#echo "%7B%22cmd%22%3A%22get%22%2C%22FlowType%22%3A%22RA%22%2C%22RAID%22%3A3%2C%22UserRefNo%22%3Anull%2C%22Version%22%3A%22refno%22%7D" > request
#dojsonPOST "http://localhost:8270/v1/flow/1/0" "request" "z0" "Rental Agreement--Edit--RAID:3"

# edit pets information
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A3%2C%22FlowPartKey%22%3A%22pets%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22Fees%22%3A%5B%7B%22recid%22%3A1%2C%22TMPASMID%22%3A3%2C%22ARID%22%3A24%2C%22ASMID%22%3A38%2C%22ARName%22%3A%22Pet%20Rent%22%2C%22ContractAmount%22%3A3.55%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22Start%22%3A%228%2F21%2F2018%22%2C%22Stop%22%3A%228%2F21%2F2018%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22prorated%20for%2011%20of%2031%20days%22%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%2C%7B%22recid%22%3A2%2C%22TMPASMID%22%3A4%2C%22ARID%22%3A24%2C%22ASMID%22%3A38%2C%22ARName%22%3A%22Pet%20Rent%22%2C%22ContractAmount%22%3A10%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22Start%22%3A%229%2F1%2F2018%22%2C%22Stop%22%3A%223%2F1%2F2020%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22%22%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%5D%2C%22Name%22%3A%22Bosky%22%2C%22Type%22%3A%22Cat%22%2C%22Breed%22%3A%22Neapolitan%20Mastiff%22%2C%22Color%22%3A%22Black%22%2C%22PETID%22%3A2%2C%22DtStop%22%3A%223%2F1%2F2020%22%2C%22Weight%22%3A11%2C%22DtStart%22%3A%228%2F21%2F2018%22%2C%22TMPTCID%22%3A1%2C%22TMPPETID%22%3A1%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/3/" "request" "z0" "Rental Agreement--RAID:3--Update Pet Information"

# edit vehicle information
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A3%2C%22FlowPartKey%22%3A%22vehicles%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22VID%22%3A3%2C%22VIN%22%3A%222BPDY2OYZM4YMRTC%22%2C%22Fees%22%3A%5B%5D%2C%22DtStop%22%3A%223%2F1%2F2020%22%2C%22TMPVID%22%3A1%2C%22DtStart%22%3A%228%2F22%2F2018%22%2C%22TMPTCID%22%3A1%2C%22VehicleMake%22%3A%22Suzuki%22%2C%22VehicleType%22%3A%22Bike%22%2C%22VehicleYear%22%3A2008%2C%22VehicleColor%22%3A%22White%22%2C%22VehicleModel%22%3A%2212BBT%22%2C%22LicensePlateState%22%3A%22GJ%22%2C%22LicensePlateNumber%22%3A%221T9RW28%22%2C%22ParkingPermitNumber%22%3A%227916444%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/3/" "request" "z1" "Rental Agreement--RAID:3--Update Vehicle Information"

# edit people information
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A3%2C%22FlowPartKey%22%3A%22people%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22City%22%3A%22Denton%22%2C%22TCID%22%3A3%2C%22State%22%3A%22ME%22%2C%22Points%22%3A0%2C%22Address%22%3A%2266789%20Shore%22%2C%22Comment%22%3A%22%22%2C%22Country%22%3A%22USA%22%2C%22Evicted%22%3Afalse%2C%22TMPTCID%22%3A1%2C%22Website%22%3A%22%22%2C%22Address2%22%3A%22%22%2C%22Industry%22%3A344%2C%22IsRenter%22%3Atrue%2C%22LastName%22%3A%22Bosamiya%22%2C%22CellPhone%22%3A%22(314)%20860-0587%22%2C%22Convicted%22%3Afalse%2C%22FirstName%22%3A%22Akshay%22%2C%22IsCompany%22%3Afalse%2C%22WorkPhone%22%3A%22(607)%20954-3966%22%2C%22Bankruptcy%22%3Afalse%2C%22EvictedDes%22%3A%22%22%2C%22IsOccupant%22%3Atrue%2C%22MiddleName%22%3A%22%22%2C%22Occupation%22%3A%22the%20hygiene%20service%20assistant%20(hygiene%20service%20assistant)%22%2C%22PostalCode%22%3A%2228162%22%2C%22TaxpayorID%22%3A%2208114320%22%2C%22CompanyCity%22%3A%22Elizabeth%22%2C%22CompanyName%22%3A%22Western%20Digital%20Inc%22%2C%22CreditLimit%22%3A15343%2C%22DateofBirth%22%3A%222%2F15%2F1957%22%2C%22GrossIncome%22%3A22028%2C%22IsGuarantor%22%3Afalse%2C%22SourceSLSID%22%3A17%2C%22CompanyEmail%22%3A%22WElizabeth7089%40aol.com%22%2C%22CompanyPhone%22%3A%22(255)%20339-0248%22%2C%22CompanyState%22%3A%22VA%22%2C%22ConvictedDes%22%3A%22%22%2C%22PrimaryEmail%22%3A%22akshay%40yopmail.com%22%2C%22PriorAddress%22%3A%2266787%20Hampton%2C%20Cambridge%2C%20NC%2031445%22%2C%22SpecialNeeds%22%3A%22%22%2C%22BankruptcyDes%22%3A%22%22%2C%22PreferredName%22%3A%22Denisha%22%2C%22CompanyAddress%22%3A%2271590%20Fifth%22%2C%22CurrentAddress%22%3A%2281062%20Wood%2C%20Santa%20Rosa%2C%20FL%2011211%22%2C%22DriversLicense%22%3A%22D7626933%22%2C%22SecondaryEmail%22%3A%22akshay%40yopmail.com%22%2C%22OtherPreferences%22%3A%22%22%2C%22ThirdPartySource%22%3A%22Stacia%20Robertson%22%2C%22CompanyPostalCode%22%3A%2233274%22%2C%22PriorLandLordName%22%3A%22Kali%20Graves%22%2C%22EligibleFutureUser%22%3Afalse%2C%22CurrentLandLordName%22%3A%22Eugenia%20Dunn%22%2C%22EligibleFuturePayor%22%3Atrue%2C%22EmergencyContactName%22%3A%22Reyna%20Ramirez%22%2C%22PriorLandLordPhoneNo%22%3A%22(823)%20260-2871%22%2C%22PriorReasonForMoving%22%3A117%2C%22AlternateEmailAddress%22%3A%2214557%20Lakeview%2CJefferson%2CGA%2024728%22%2C%22EmergencyContactEmail%22%3A%22RRamirez8989%40bdiddy.com%22%2C%22CurrentLandLordPhoneNo%22%3A%22(510)%20858-4871%22%2C%22CurrentReasonForMoving%22%3A129%2C%22PriorLengthOfResidency%22%3A%228%20years%206%20months%22%2C%22EmergencyContactAddress%22%3A%2284390%20Rhode%20Island%2CDurham%2CAK%2055089%22%2C%22CurrentLengthOfResidency%22%3A%222%20years%207%20months%22%2C%22EmergencyContactTelephone%22%3A%22(506)%20681-2584%22%2C%22recid%22%3A1%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22NLID%22%3A%22%22%2C%22CreateBy%22%3A%22%22%2C%22CreateTS%22%3A%22%22%2C%22LastModBy%22%3A%22%22%2C%22LastModTime%22%3A%22%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/3/" "request" "z2" "Rental Agreement--RAID:3--Update People Information"

# validate updated flow's data
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A3%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/3/" "request" "z3" "Rental Agreement--RAID:3--Validate update RAFlow"

# RAAction: Complete to Move In
echo "%7B%22UserRefNo%22%3A%22${RAID3REFNO}%22%2C%22RAID%22%3A3%2C%22Version%22%3A%22refno%22%2C%22Action%22%3A4%2C%22Mode%22%3A%22Action%22%7D" > request
dojsonPOST "http://localhost:8270/v1/raactions/1/3/" "request" "z4" "Rental Agreement--RAID:3--Complete Move In"

# Get updated flow
#echo "" > request
#dojsonPOST "" "request" "z6" "Rental Agreement--RAID:3--Get updated flow"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
