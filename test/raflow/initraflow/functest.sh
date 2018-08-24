#!/bin/bash

TESTNAME="Init RAFlow"
TESTSUMMARY="Test to init raflow and saving data component wise"
DBGEN=../../../tools/dbgen
CREATENEWDB=0
RRBIN="../../../tmp/rentroll"

echo "Create new database..."
mysql --no-defaults rentroll < ra0.sql

source ../../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
startRentRollServer

#------------------------------------------------------------------------------
#  Following test cases will be performed using only web service calls.
#
#  Test cases will be performed in a manner of sequences.
#  It will add/init data section by section as follows:
#       1. Dates/Agent
#       2. People
#       3. Pets
#       4. Vehicles
#       5. Rentables
#       6. Parent/Child
#       7. Tie
#------------------------------------------------------------------------------

# INIITIATE A NEW RAFLOW WITH SECTIONS BASIC DATA
echo "%7B%22cmd%22%3A%22init%22%2C%22FlowType%22%3A%22RA%22%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/0/" "request" "a0"  "RAFlow--initiate_brand_new_raflow"

# CHANGE ALL DATES(TERM, RENT, POSSESSION) START: 11 JAN, 2018 | STOP: 1 JAN, 2020, CSAGENT TO 1
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22dates%22%2C%22BID%22%3A1%2C%22Data%22%3A%7B%22CSAgent%22%3A1%2C%22RentStop%22%3A%221%2F1%2F2020%22%2C%22RentStart%22%3A%221%2F11%2F2018%22%2C%22AgreementStop%22%3A%221%2F1%2F2020%22%2C%22AgreementStart%22%3A%221%2F11%2F2018%22%2C%22PossessionStop%22%3A%221%2F1%2F2020%22%2C%22PossessionStart%22%3A%221%2F11%2F2018%22%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a1"  "RAFlow--modify_dates_csagent"

# ADD EXISTING PERSON WITH TCID 3 (Elke Sanders) (HAVING ONE PET & ONE VEHICLE) (Renter & Occupant both)
echo "%7B%22cmd%22%3A%22save%22%2C%22TCID%22%3A3%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-person/1/1/" "request" "a2"  "RAFlow--add_person--Elke_Sanders_TCID_3_existing"

# ADD NEW PERSON WITH TCID 0 (John M Doe, john.doe@earth.com) (Occupant only)
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22people%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22City%22%3A%22Akron%22%2C%22TCID%22%3A3%2C%22State%22%3A%22AL%22%2C%22Points%22%3A0%2C%22Address%22%3A%2228658%20Pioneer%22%2C%22Comment%22%3A%22%22%2C%22Country%22%3A%22USA%22%2C%22Evicted%22%3Afalse%2C%22TMPTCID%22%3A1%2C%22Website%22%3A%22%22%2C%22Address2%22%3A%22%22%2C%22Industry%22%3A234%2C%22IsRenter%22%3Atrue%2C%22LastName%22%3A%22Sanders%22%2C%22CellPhone%22%3A%22(841)%20877-5574%22%2C%22Convicted%22%3Afalse%2C%22FirstName%22%3A%22Elke%22%2C%22IsCompany%22%3Afalse%2C%22WorkPhone%22%3A%22(675)%20462-1665%22%2C%22Bankruptcy%22%3Afalse%2C%22EvictedDes%22%3A%22%22%2C%22IsOccupant%22%3Atrue%2C%22MiddleName%22%3A%22Dolly%22%2C%22Occupation%22%3A%22travel%20agency%20clerk%22%2C%22PostalCode%22%3A%2228162%22%2C%22TaxpayorID%22%3A%2201428008%22%2C%22CompanyCity%22%3A%22FortWorth%22%2C%22CompanyName%22%3A%22Group%201%20Automotive%20Inc.%22%2C%22CreditLimit%22%3A22755%2C%22DateofBirth%22%3A%226%2F7%2F1964%22%2C%22GrossIncome%22%3A67090%2C%22IsGuarantor%22%3Afalse%2C%22SourceSLSID%22%3A11%2C%22CompanyEmail%22%3A%22Group1AutomotiveIncFortWorth569%40abiz.com%22%2C%22CompanyPhone%22%3A%22(499)%20551-5800%22%2C%22CompanyState%22%3A%22PA%22%2C%22ConvictedDes%22%3A%22%22%2C%22PrimaryEmail%22%3A%22ElkeSanders551%40aol.com%22%2C%22PriorAddress%22%3A%2231277%20Magnolia%2C%20Shreveport%2C%20ME%2031445%22%2C%22SpecialNeeds%22%3A%22%22%2C%22BankruptcyDes%22%3A%22%22%2C%22PreferredName%22%3A%22Flo%22%2C%22CompanyAddress%22%3A%2234626%20County%20Line%22%2C%22CurrentAddress%22%3A%2235942%20Dogwood%2C%20Montgomery%2C%20IN%2011211%22%2C%22DriversLicense%22%3A%22G8252891%22%2C%22SecondaryEmail%22%3A%22ESanders9207%40aol.com%22%2C%22OtherPreferences%22%3A%22%22%2C%22ThirdPartySource%22%3A%22Elois%20Erickson%22%2C%22CompanyPostalCode%22%3A%2233274%22%2C%22PriorLandLordName%22%3A%22Cheri%20Bentley%22%2C%22EligibleFutureUser%22%3Afalse%2C%22CurrentLandLordName%22%3A%22Mandi%20Jackson%22%2C%22EligibleFuturePayor%22%3Atrue%2C%22EmergencyContactName%22%3A%22Sherron%20Jensen%22%2C%22PriorLandLordPhoneNo%22%3A%22(606)%20217-7001%22%2C%22PriorReasonForMoving%22%3A93%2C%22AlternateEmailAddress%22%3A%2275871%20Navajo%2CLexington%2COH%2024728%22%2C%22EmergencyContactEmail%22%3A%22SJensen3151%40comcast.net%22%2C%22CurrentLandLordPhoneNo%22%3A%22(163)%20990-9639%22%2C%22CurrentReasonForMoving%22%3A94%2C%22PriorLengthOfResidency%22%3A%224%20years%207%20months%22%2C%22EmergencyContactAddress%22%3A%2233391%20Church%2CCanton%2CAR%2055089%22%2C%22CurrentLengthOfResidency%22%3A%223%20years%2010%20months%22%2C%22EmergencyContactTelephone%22%3A%22(853)%20692-7008%22%2C%22recid%22%3A1%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%2C%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22NLID%22%3A0%2C%22TCID%22%3A0%2C%22TMPTCID%22%3A0%2C%22IsRenter%22%3Afalse%2C%22IsOccupant%22%3Atrue%2C%22IsGuarantor%22%3Afalse%2C%22FirstName%22%3A%22John%22%2C%22MiddleName%22%3A%22M%22%2C%22LastName%22%3A%22Doe%22%2C%22PreferredName%22%3A%22%22%2C%22IsCompany%22%3Afalse%2C%22CompanyName%22%3A%22%22%2C%22PrimaryEmail%22%3A%22johndoe%40earth.co%22%2C%22SecondaryEmail%22%3A%22%22%2C%22WorkPhone%22%3A%22%22%2C%22CellPhone%22%3A%22%2B1-999-000-9999%22%2C%22Address%22%3A%223118%20%20Doctors%20Drive%22%2C%22Address2%22%3A%22%22%2C%22City%22%3A%22Los%20Angeles%22%2C%22State%22%3A%22CA%22%2C%22PostalCode%22%3A%2290017%22%2C%22Country%22%3A%22USA%22%2C%22Website%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22Points%22%3A0%2C%22DateofBirth%22%3A%226%2F1%2F1951%22%2C%22EmergencyContactAddress%22%3A%2233391%20Church%2CCanton%2CAR%2055089%22%2C%22EmergencyContactEmail%22%3A%22friend%40earth.world%22%2C%22EmergencyContactName%22%3A%22Earth%20Man%22%2C%22EmergencyContactTelephone%22%3A%22%2B1-099-000-9999%22%2C%22AlternateEmailAddress%22%3A%22%22%2C%22EligibleFutureUser%22%3Atrue%2C%22Industry%22%3A183%2C%22SourceSLSID%22%3A11%2C%22CreditLimit%22%3A100%2C%22TaxpayorID%22%3A%2209090909%22%2C%22GrossIncome%22%3A10012%2C%22DriversLicense%22%3A%22G8XX9EXX%22%2C%22ThirdPartySource%22%3A%22%22%2C%22EligibleFuturePayor%22%3Atrue%2C%22CompanyAddress%22%3A%22%22%2C%22CompanyCity%22%3A%22%22%2C%22CompanyState%22%3A%22%22%2C%22CompanyPostalCode%22%3A%22%22%2C%22CompanyEmail%22%3A%22%22%2C%22CompanyPhone%22%3A%22%22%2C%22Occupation%22%3A%22free%2C%20jobless%2C%20happy%20man%22%2C%22CurrentAddress%22%3A%2235942%20Dogwood%2C%20Montgomery%2C%20IN%2011211%22%2C%22CurrentLandLordName%22%3A%22Steve%20Roger%22%2C%22CurrentLandLordPhoneNo%22%3A%22%2B1-000-911-9000%22%2C%22CurrentLengthOfResidency%22%3A%220%20years%2010%20months%22%2C%22CurrentReasonForMoving%22%3A94%2C%22PriorAddress%22%3A%22%22%2C%22PriorLandLordName%22%3A%22%22%2C%22PriorLandLordPhoneNo%22%3A%22%22%2C%22PriorLengthOfResidency%22%3A%22%22%2C%22PriorReasonForMoving%22%3A0%2C%22Evicted%22%3Afalse%2C%22EvictedDes%22%3A%22%22%2C%22Convicted%22%3Afalse%2C%22ConvictedDes%22%3A%22%22%2C%22Bankruptcy%22%3Afalse%2C%22BankruptcyDes%22%3A%22%22%2C%22OtherPreferences%22%3A%22%22%2C%22SpecialNeeds%22%3A%22%22%2C%22LastModTime%22%3A%222018-08-21T11%3A25%3A03.697Z%22%2C%22LastModBy%22%3A0%2C%22CreateBy%22%3A%22%22%2C%22CreateTS%22%3A%22%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a3"  "RAFlow--add_person--John_Doe_TCID_0_new"

# REQUEST TO CREATE NEW ENTRY FOR A PET (FOR John Doe: TCID=0)
echo "%7B%22cmd%22%3A%22new%22%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-pets/1/1/" "request" "a4" "RAFlow--create_new_pet_entry"

# UPDATE PET INFO AND ASSIGN CONTACT PERSON (John Doe: TCID=0)
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22pets%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22Fees%22%3A%5B%7B%22ARID%22%3A23%2C%22Stop%22%3A%221%2F11%2F2018%22%2C%22ASMID%22%3A0%2C%22Start%22%3A%221%2F11%2F2018%22%2C%22ARName%22%3A%22Pet%20Fee%22%2C%22Comment%22%3A%22%22%2C%22SalesTax%22%3A0%2C%22TMPASMID%22%3A1%2C%22RentCycle%22%3A0%2C%22TransOccTax%22%3A0%2C%22ContractAmount%22%3A50%2C%22ProrationCycle%22%3A0%2C%22AtSigningPreTax%22%3A0%7D%2C%7B%22ARID%22%3A24%2C%22Stop%22%3A%221%2F11%2F2018%22%2C%22ASMID%22%3A0%2C%22Start%22%3A%221%2F11%2F2018%22%2C%22ARName%22%3A%22Pet%20Rent%22%2C%22Comment%22%3A%22prorated%20for%2021%20of%2031%20days%22%2C%22SalesTax%22%3A0%2C%22TMPASMID%22%3A2%2C%22RentCycle%22%3A0%2C%22TransOccTax%22%3A0%2C%22ContractAmount%22%3A6.77%2C%22ProrationCycle%22%3A0%2C%22AtSigningPreTax%22%3A0%7D%2C%7B%22ARID%22%3A24%2C%22Stop%22%3A%221%2F1%2F2020%22%2C%22ASMID%22%3A0%2C%22Start%22%3A%222%2F1%2F2018%22%2C%22ARName%22%3A%22Pet%20Rent%22%2C%22Comment%22%3A%22%22%2C%22SalesTax%22%3A0%2C%22TMPASMID%22%3A3%2C%22RentCycle%22%3A6%2C%22TransOccTax%22%3A0%2C%22ContractAmount%22%3A10%2C%22ProrationCycle%22%3A4%2C%22AtSigningPreTax%22%3A0%7D%5D%2C%22Name%22%3A%22Gizmo%20%22%2C%22Type%22%3A%22cat%22%2C%22Breed%22%3A%22American%20Ringtail%22%2C%22Color%22%3A%22chocolate%20point%22%2C%22PETID%22%3A1%2C%22Weight%22%3A14%2C%22TMPTCID%22%3A1%2C%22TMPPETID%22%3A1%7D%2C%7B%22Fees%22%3A%5B%7B%22recid%22%3A1%2C%22TMPASMID%22%3A5%2C%22ARID%22%3A23%2C%22ASMID%22%3A0%2C%22ARName%22%3A%22Pet%20Fee%22%2C%22ContractAmount%22%3A50%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22Start%22%3A%221%2F11%2F2018%22%2C%22Stop%22%3A%221%2F11%2F2018%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22%22%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%2C%7B%22recid%22%3A2%2C%22TMPASMID%22%3A6%2C%22ARID%22%3A24%2C%22ASMID%22%3A0%2C%22ARName%22%3A%22Pet%20Rent%22%2C%22ContractAmount%22%3A6.77%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22Start%22%3A%221%2F11%2F2018%22%2C%22Stop%22%3A%221%2F11%2F2018%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22prorated%20for%2021%20of%2031%20days%22%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%2C%7B%22recid%22%3A3%2C%22TMPASMID%22%3A7%2C%22ARID%22%3A24%2C%22ASMID%22%3A0%2C%22ARName%22%3A%22Pet%20Rent%22%2C%22ContractAmount%22%3A10%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22Start%22%3A%222%2F1%2F2018%22%2C%22Stop%22%3A%221%2F1%2F2020%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22%22%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%5D%2C%22Name%22%3A%22Micky%22%2C%22Type%22%3A%22dog%22%2C%22Breed%22%3A%22Rough%20Collie%22%2C%22Color%22%3A%22brown%22%2C%22PETID%22%3A0%2C%22Weight%22%3A45%2C%22TMPTCID%22%3A2%2C%22TMPPETID%22%3A2%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a5" "RAFlow--update_new_pet_info--Micky--John_Doe_TCID_0"

# REQUEST TO CREATE NEW ENTRY FOR A VEHICLE (FOR John Doe: TCID=0)
echo "%7B%22cmd%22%3A%22new%22%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-vehicles/1/1/" "request" "a6" "RAFlow--create_new_vehicle_entry"

# UPDATE VEHICLE INFO AND ASSIGN CONTACT PERSON (John Doe: TCID=0)
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22vehicles%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22VID%22%3A3%2C%22VIN%22%3A%22SMT4S7HCVPX8CPL8%22%2C%22Fees%22%3A%5B%7B%22ARID%22%3A39%2C%22Stop%22%3A%221%2F11%2F2018%22%2C%22ASMID%22%3A0%2C%22Start%22%3A%221%2F11%2F2018%22%2C%22ARName%22%3A%22Vehicle%20Registration%20Fee%22%2C%22Comment%22%3A%22%22%2C%22SalesTax%22%3A0%2C%22TMPASMID%22%3A4%2C%22RentCycle%22%3A0%2C%22TransOccTax%22%3A0%2C%22ContractAmount%22%3A10%2C%22ProrationCycle%22%3A0%2C%22AtSigningPreTax%22%3A0%7D%5D%2C%22TMPVID%22%3A1%2C%22TMPTCID%22%3A1%2C%22VehicleMake%22%3A%22Pontiac%22%2C%22VehicleType%22%3A%22car%22%2C%22VehicleYear%22%3A1995%2C%22VehicleColor%22%3A%22Metallic%22%2C%22VehicleModel%22%3A%22Firebird%22%2C%22LicensePlateState%22%3A%22CA%22%2C%22LicensePlateNumber%22%3A%22GW1M627%22%2C%22ParkingPermitNumber%22%3A%223177786%22%7D%2C%7B%22VID%22%3A0%2C%22VIN%22%3A%22%22%2C%22Fees%22%3A%5B%7B%22recid%22%3A1%2C%22TMPASMID%22%3A8%2C%22ARID%22%3A39%2C%22ASMID%22%3A0%2C%22ARName%22%3A%22Vehicle%20Registration%20Fee%22%2C%22ContractAmount%22%3A10%2C%22RentCycle%22%3A0%2C%22ProrationCycle%22%3A0%2C%22Start%22%3A%221%2F11%2F2018%22%2C%22Stop%22%3A%221%2F11%2F2018%22%2C%22AtSigningPreTax%22%3A0%2C%22SalesTax%22%3A0%2C%22TransOccTax%22%3A0%2C%22Comment%22%3A%22%22%2C%22w2ui%22%3A%7B%22class%22%3A%22%22%2C%22style%22%3A%7B%7D%7D%7D%5D%2C%22TMPVID%22%3A2%2C%22TMPTCID%22%3A2%2C%22VehicleMake%22%3A%22Tesla%22%2C%22VehicleType%22%3A%224%20wheeler%22%2C%22VehicleYear%22%3A2018%2C%22VehicleColor%22%3A%22black%22%2C%22VehicleModel%22%3A%22S%22%2C%22LicensePlateState%22%3A%22CA%22%2C%22LicensePlateNumber%22%3A%22E9V9999%22%2C%22ParkingPermitNumber%22%3A%22%22%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a7" "RAFlow--update_new_vehicle_info--Tesla_S_2018--John_Doe_TCID_0"

# ADD ONE RENTABLE (Rentable001:RID=1) FOR (Elke Sanders: TCID=3)
echo "%7B%22cmd%22%3A%22save%22%2C%22RID%22%3A1%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-rentable/1/1/" "request" "a8" "RAFlow--add_rentable--Rentable001_RID_1--Elke_Sanders_TCID_3"

# ADD ONE MORE RENTABLE (Rentable002:RID=2) FOR (John Doe: TCID=0)
echo "%7B%22cmd%22%3A%22save%22%2C%22RID%22%3A2%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-rentable/1/1/" "request" "a9" "RAFlow--add_rentable--Rentable002_RID_2--John_Doe_TCID_0"

# ADD ONE CHILD RENTABLE (CP001:RID=9) FOR (Elke Sanders: TCID=3)
echo "%7B%22cmd%22%3A%22save%22%2C%22RID%22%3A9%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/raflow-rentable/1/1/" "request" "a10" "RAFlow--add_rentable--CP001_RID_9--Elke_Sanders_TCID_3"

# ASSIGN CHILD RENTABLE (CP001:RID=9) TO RENTABLE (Rentable001:RID=1)
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22parentchild%22%2C%22BID%22%3A1%2C%22Data%22%3A%5B%7B%22CRID%22%3A9%2C%22PRID%22%3A1%7D%5D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a11" "RAFlow--assign_child_rentable--CP001_RID_9--to_rentable--Rentable001_RID_1"

# ASSIGN RENTABLE(Rentable002:RID=2) TO (John Doe: TCID=0)
echo "%7B%22cmd%22%3A%22save%22%2C%22FlowType%22%3A%22RA%22%2C%22FlowID%22%3A1%2C%22FlowPartKey%22%3A%22tie%22%2C%22BID%22%3A1%2C%22Data%22%3A%7B%22people%22%3A%5B%7B%22PRID%22%3A1%2C%22TMPTCID%22%3A1%7D%2C%7B%22PRID%22%3A2%2C%22TMPTCID%22%3A2%7D%5D%7D%7D" > request
dojsonPOST "http://localhost:8270/v1/flow/1/1/" "request" "a12" "RAFlow--assign_rentable--Rentable001_RID_1--John_Doe_TCID_0"

# ALL DATA ARE SAVE, VALIDATE THE FLOW (TOTAL:0), WILL SET THE STATE TO PENDING FIRST APPROVAL
echo "%7B%22cmd%22%3A%22get%22%2C%22FlowID%22%3A1%7D" > request
dojsonPOST "http://localhost:8270/v1/validate-raflow/1/1/" "request" "a13" "RAFlow--validate_raflow"

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

logcheck

exit 0
