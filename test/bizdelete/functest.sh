#!/bin/bash

TESTNAME="Business Delete Test"
TESTSUMMARY="Creates many records for a business, then tests DeleteBusinessFromDB to remove them"

source ../share/base.sh
RRCTX="-G ${BUD} -g 12/1/15,1/1/16"

# Create a bunch of content
${CSVLOAD} -b nb.csv  >>${LOGFILE} 2>&1
${CSVLOAD} -f rprefs.csv -n rprtrate.csv -t rpsprate.csv -l strlists.csv -R rt.csv -u custom.csv -d depository.csv -s specialties.csv -D bldg.csv -p people.csv -r rentable.csv -T rat.csv -C ra.csv -E pets.csv -a rp.csv -c coa.csv -A asmt.csv -P pmt.csv -e rcpt.csv -U assigncustom.csv -O nt.csv -m depmeth.csv -y deposit.csv -S sources.csv ${RRCTX} >>${LOGFILE} 2>&1

# Delete the content with DeleteBusinessFromDB
dorrtest "a00" "-r 22 -b ${BUD}" "DeleteBusiness"

# Validate that all the records were deleted
mysqlverify "a"  "${RRCTX}"		"NewBusinesses"	            	"select BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,LastModBy from Business;"
mysqlverify "b"  "${RRCTX}"		"StringLists"	            	"select SLID,BID,Name,LastModBy from StringList;"
mysqlverify "c"  "${RRCTX}"		"SLString"	            		"select SLSID,SLID,Value,LastModBy from SLString;"
mysqlverify "d"  "${RRCTX}"		"RentableTypes"	            	"select RTID,BID,Style,Name,RentCycle,Proration,GSRPC,ManageToBudget,LastModBy from RentableTypes;"
mysqlverify "e"  "${RRCTX}"		"RentableMarketRates"	    	"select * from RentableMarketRate;"
mysqlverify "f"  "${RRCTX}"		"Deposit Methods"           	"select * from DepositMethod;"
mysqlverify "g"  "${RRCTX}"		"Sources"	            		"select SourceSLSID,BID,Name,Industry from DemandSource;"
mysqlverify "h"  "${RRCTX}"		"RentableSpecialtyTypes"    	"select * from RentableSpecialty;"
mysqlverify "i"  "${RRCTX}"		"Buildings"	            		"select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
mysqlverify "j"  "${RRCTX}"		"Depositories"	            	"select DEPID,BID,Name,AccountNo,LastModBy from Depository;"
mysqlverify "n"  "${RRCTX}"		"Transactants"	            	"select TCID,BID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
mysqlverify "o"  "${RRCTX}"		"Users"	                    	"select TCID,Points,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyEmail,AlternateAddress,EligibleFutureUser,Industry,SourceSLSID from User;"
mysqlverify "p"  "${RRCTX}"		"Payors"	                    "select TCID,CreditLimit,TaxpayorID,AccountRep,LastModBy from Payor;"
mysqlverify "q"  "${RRCTX}"		"Prospects"	            		"select TCID,EmployerName,EmployerStreetAddress,EmployerCity,EmployerState,EmployerPostalCode,EmployerEmail,EmployerPhone,Occupation,ApplicationFee,LastModBy from Prospect;"
mysqlverify "na" "${RRCTX}"		"Vehicles"	            		"select VID,TCID,VehicleType,VehicleMake,VehicleModel,VehicleColor,VehicleYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DtStart,DtStop,LastModBy from Vehicle;"
mysqlverify "k"  "${RRCTX}"		"Rentables"	            		"select RID,BID,Name,AssignmentTime,LastModBy from Rentable;"
mysqlverify "l"  "${RRCTX}"		"RentableTypeRef"	    		"select RID,BID,RTID,OverrideRentCycle,OverrideProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
mysqlverify "m"  "${RRCTX}"		"RentableStatus"	            "select RID,Status,DtStart,DtStop,LastModBy from RentableStatus;"
mysqlverify "r"  "${RRCTX}"		"RentalAgreementTemplates"   	"select RATID,BID,RATemplateName,LastModBy from RentalAgreementTemplate;"
mysqlverify "s"  "${RRCTX}"		"RentalAgreements"	    		"select RAID,RATID,BID,AgreementStart,AgreementStop,Renewal,SpecialProvisions,UnspecifiedAdults,UnspecifiedChildren,LastModBy from RentalAgreement;"
mysqlverify "t"  "${RRCTX}"		"Pets"	                    	"select PETID,BID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from RentalAgreementPets;"
mysqlverify "u"  "${RRCTX}"		"Notes"	                    	"select NID,PNID,Comment,LastModBy from Notes;"
mysqlverify "v"  "${RRCTX}"		"AgreementRentables"	    	"select * from RentalAgreementRentables;"
mysqlverify "w"  "${RRCTX}"		"AgreementPayors"	    		"select * from RentalAgreementPayors;"
mysqlverify "x"  "${RRCTX}"		"ChartOfAccounts"	    		"select LID,PLID,BID,RAID,GLNumber,Status,Type,Name,AcctType,RAAssociated,AllowPost,LastModBy from GLAccount;"
mysqlverify "y"  "${RRCTX}"		"LedgerMarkers"	            	"select LMID,LID,BID,Dt,Balance,State,LastModBy from LedgerMarker;"
mysqlverify "z"  "${RRCTX}"		"RatePlan"	            		"select RPID,BID,Name,LastModBy from RatePlan;"
mysqlverify "a1" "${RRCTX}"		"RatePlanRef"	                "select RPRID,BID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,LastModBy from RatePlanRef;"
mysqlverify "b1" "${RRCTX}"		"RatePlanRefRTRate"	    		"select * from RatePlanRefRTRate;"
mysqlverify "c1" "${RRCTX}"		"RatePlanRefSPRate"	    		"select * from RatePlanRefSPRate;"
mysqlverify "d1" "${RRCTX}"		"Assessments"	            	"select ASMID,BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
mysqlverify "e1" "${RRCTX}"		"PaymentTypes"	            	"select PMTID,BID,Name,Description,LastModBy from PaymentTypes;"
mysqlverify "f1" "${RRCTX}"		"PaymentAllocations"	    	"select * from ReceiptAllocation order by Amount ASC;"
mysqlverify "g1" "${RRCTX}"		"Receipts"	            		"select RCPTID,BID,RAID,PMTID,Dt,Amount,AcctRule,Comment,LastModBy from Receipt;"
mysqlverify "h1" "${RRCTX}"		"CustomAttributes"	    		"select CID,BID,Type,Name,Value,LastModBy from CustomAttr;"
mysqlverify "i1" "${RRCTX}"		"CustomAttributesAssignment" 	"select * from CustomAttrRef;"
mysqlverify "j1" "${RRCTX}"		"NoteTypes"	            		"select NTID,BID,Name,LastModBy from NoteType;"
mysqlverify "k1" "${RRCTX}"		"Deposits"	            		"select DID,BID,Dt,DEPID,Amount,LastModBy from Deposit;"


logcheck
