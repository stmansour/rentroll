#!/bin/bash

TESTNAME="Business Delete Test"
TESTSUMMARY="Creates many records for a business, then tests DeleteBusinessFromDB to remove them"

source ../share/base.sh
RRCTX="-G ${BUD} -g 12/1/15,1/1/16"

# Create a bunch of content
${CSVLOAD} -b nb.csv >>${LOGFILE} 2>&1
${CSVLOAD} -R rt.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -u custom.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -s specialties.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -D bldg.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -p people.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -r rentable.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -T rat.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -C ra.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -E pets.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -c coa.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -d depository.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -a rp.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -f rprefs.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -n rprtrate.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -t rpsprate.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -l strlists.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -A asmt.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -P pmt.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -e rcpt.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -U assigncustom.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -O nt.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -m depmeth.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -y deposit.csv -S sources.csv ${RRCTX} >>${LOGFILE} 2>&1

# Delete the content with DeleteBusinessFromDB
dorrtest "a00" "-r 22 -b ${BUD}" "DeleteBusiness"

# Validate that all the records were deleted
mysqlverify "a"  "NewBusinesses"	            	"select BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,LastModBy from Business;"
mysqlverify "b"  "StringLists"	            	"select SLID,BID,Name,LastModBy from StringList;"
mysqlverify "c"  "SLString"	            		"select SLSID,SLID,Value,LastModBy from SLString;"
mysqlverify "d"  "RentableTypes"	            	"select RTID,BID,Style,Name,RentCycle,Proration,GSRPC,FLAGS,LastModBy from RentableTypes;"
mysqlverify "e"  "RentableMarketRates"	    	"select * from RentableMarketRate;"
mysqlverify "f"  "Deposit Methods"           	"select * from DepositMethod;"
mysqlverify "g"  "Sources"	            		"select SourceSLSID,BID,Name,Industry from DemandSource;"
mysqlverify "h"  "RentableSpecialtyTypes"    	"select * from RentableSpecialty;"
mysqlverify "i"  "Buildings"	            		"select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
mysqlverify "j"  "Depositories"	            	"select DEPID,BID,Name,AccountNo,LastModBy from Depository;"
mysqlverify "n"  "Transactants"	            	"select TCID,BID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
mysqlverify "o"  "Users"	                    	"select TCID,Points,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyContactEmail,AlternateEmailAddress,EligibleFutureUser,Industry,SourceSLSID from User;"
mysqlverify "p"  "Payors"	                    "select TCID,CreditLimit,TaxpayorID,LastModBy from Payor;"
mysqlverify "q"  "Prospects"	            		"select TCID,CompanyAddress,CompanyCity,CompanyState,CompanyPostalCode,CompanyEmail,CompanyPhone,Occupation,ThirdPartySource,LastModBy from Prospect;"
mysqlverify "na" "Vehicles"	            		"select VID,TCID,VehicleType,VehicleMake,VehicleModel,VehicleColor,VehicleYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DtStart,DtStop,LastModBy from Vehicle;"
mysqlverify "k"  "Rentables"	            		"select RID,BID,RentableName,AssignmentTime,LastModBy from Rentable;"
mysqlverify "l"  "RentableTypeRef"	    		"select RID,BID,RTID,OverrideRentCycle,OverrideProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
mysqlverify "m"  "RentableStatus"	            "select RID,UseStatus,LeaseStatus,DtStart,DtStop,LastModBy from RentableStatus;"
mysqlverify "r"  "RentalAgreementTemplates"   	"select RATID,BID,RATemplateName,LastModBy from RentalAgreementTemplate;"
mysqlverify "s"  "RentalAgreements"	    		"select RAID,RATID,BID,AgreementStart,AgreementStop,Renewal,SpecialProvisions,UnspecifiedAdults,UnspecifiedChildren,LastModBy from RentalAgreement;"
mysqlverify "t"  "Pets"	                    	"select PETID,BID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from Pets;"
mysqlverify "u"  "Notes"	                    	"select NID,PNID,Comment,LastModBy from Notes;"
mysqlverify "v"  "AgreementRentables"	    	"select * from RentalAgreementRentables;"
mysqlverify "w"  "AgreementPayors"	    		"select * from RentalAgreementPayors;"
mysqlverify "x"  "ChartOfAccounts"	    		"select LID,PLID,BID,RAID,GLNumber,FLAGS,Name,AcctType,AllowPost,LastModBy from GLAccount;"
mysqlverify "y"   	"LedgerMarkers"	            	"select LMID,LID,BID,RAID,RID,TCID,Dt,Balance,State,LastModBy from LedgerMarker;"
mysqlverify "z"  "RatePlan"	            		"select RPID,BID,Name,LastModBy from RatePlan;"
mysqlverify "a1" "RatePlanRef"	                "select RPRID,BID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,LastModBy from RatePlanRef;"
mysqlverify "b1" "RatePlanRefRTRate"	    		"select * from RatePlanRefRTRate;"
mysqlverify "c1" "RatePlanRefSPRate"	    		"select * from RatePlanRefSPRate;"
mysqlverify "d1" "Assessments"	            	"select ASMID,BID,RID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
mysqlverify "e1" "PaymentTypes"	            	"select PMTID,BID,Name,Description,LastModBy from PaymentType;"
mysqlverify "f1" "PaymentAllocations"	    	"select * from ReceiptAllocation order by Amount ASC;"
mysqlverify "g1" "Receipts"	            		"select RCPTID,BID,PMTID,DEPID,DID,Dt,Amount,AcctRuleReceive,Comment,LastModBy from Receipt;"
mysqlverify "h1" "CustomAttributes"	    		"select CID,BID,Type,Name,Value,LastModBy from CustomAttr;"
mysqlverify "i1" "CustomAttributesAssignment" 	"select * from CustomAttrRef;"
mysqlverify "j1" "NoteTypes"	            		"select NTID,BID,Name,LastModBy from NoteType;"
mysqlverify "k1" "Deposits"	            		"select DID,BID,Dt,DEPID,Amount,LastModBy from Deposit;"


logcheck
