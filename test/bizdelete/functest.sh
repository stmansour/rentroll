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
mysqlverify "a"  "${RRCTX}"		"NewBusinesses"	            	"select BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,LastModBy from Business;"
mysqlverify "b"  "${RRCTX}"		"StringLists"	            	"select SLID,BID,Name,LastModBy from StringList;"
mysqlverify "c"  "${RRCTX}"		"SLString"	            		"select SLSID,SLID,Value,LastModBy from SLString;"
mysqlverify "d"  "${RRCTX}"		"RentableTypes"	            	"select RTID,BID,Style,Name,RentCycle,Proration,GSRPC,FLAGS,LastModBy from RentableTypes;"
mysqlverify "e"  "${RRCTX}"		"RentableMarketRates"	    	"select * from RentableMarketRate;"
mysqlverify "f"  "${RRCTX}"		"Deposit Methods"           	"select * from DepositMethod;"
mysqlverify "g"  "${RRCTX}"		"Sources"	            		"select SourceSLSID,BID,Name,Industry from DemandSource;"
mysqlverify "h"  "${RRCTX}"		"RentableSpecialtyTypes"    	"select * from RentableSpecialty;"
mysqlverify "i"  "${RRCTX}"		"Buildings"	            		"select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
mysqlverify "j"  "${RRCTX}"		"Depositories"	            	"select DEPID,BID,Name,AccountNo,LastModBy from Depository;"
mysqlverify "n"  "${RRCTX}"		"Transactants"	            	"select TCID,BID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
mysqlverify "o"  "${RRCTX}"		"Users"	                    	"select TCID,Points,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyContactEmail,AlternateAddress,EligibleFutureUser,Industry,SourceSLSID from User;"
mysqlverify "p"  "${RRCTX}"		"Payors"	                    "select TCID,CreditLimit,TaxpayorID,ThirdPartySource,LastModBy from Payor;"
mysqlverify "q"  "${RRCTX}"		"Prospects"	            		"select TCID,CompanyAddress,CompanyCity,CompanyState,CompanyPostalCode,CompanyEmail,CompanyPhone,Occupation,LastModBy from Prospect;"
mysqlverify "na" "${RRCTX}"		"Vehicles"	            		"select VID,TCID,VehicleType,VehicleMake,VehicleModel,VehicleColor,VehicleYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DtStart,DtStop,LastModBy from Vehicle;"
mysqlverify "k"  "${RRCTX}"		"Rentables"	            		"select RID,BID,RentableName,AssignmentTime,LastModBy from Rentable;"
mysqlverify "l"  "${RRCTX}"		"RentableTypeRef"	    		"select RID,BID,RTID,OverrideRentCycle,OverrideProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
mysqlverify "m"  "${RRCTX}"		"RentableStatus"	            "select RID,UseStatus,LeaseStatus,DtStart,DtStop,LastModBy from RentableStatus;"
mysqlverify "r"  "${RRCTX}"		"RentalAgreementTemplates"   	"select RATID,BID,RATemplateName,LastModBy from RentalAgreementTemplate;"
mysqlverify "s"  "${RRCTX}"		"RentalAgreements"	    		"select RAID,RATID,BID,AgreementStart,AgreementStop,Renewal,SpecialProvisions,UnspecifiedAdults,UnspecifiedChildren,LastModBy from RentalAgreement;"
mysqlverify "t"  "${RRCTX}"		"Pets"	                    	"select PETID,BID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from RentalAgreementPets;"
mysqlverify "u"  "${RRCTX}"		"Notes"	                    	"select NID,PNID,Comment,LastModBy from Notes;"
mysqlverify "v"  "${RRCTX}"		"AgreementRentables"	    	"select * from RentalAgreementRentables;"
mysqlverify "w"  "${RRCTX}"		"AgreementPayors"	    		"select * from RentalAgreementPayors;"
mysqlverify "x"  "${RRCTX}"		"ChartOfAccounts"	    		"select LID,PLID,BID,RAID,GLNumber,Status,Name,AcctType,AllowPost,LastModBy from GLAccount;"
mysqlverify "y"  "${RRCTX}"    	"LedgerMarkers"	            	"select LMID,LID,BID,RAID,RID,TCID,Dt,Balance,State,LastModBy from LedgerMarker;"
mysqlverify "z"  "${RRCTX}"		"RatePlan"	            		"select RPID,BID,Name,LastModBy from RatePlan;"
mysqlverify "a1" "${RRCTX}"		"RatePlanRef"	                "select RPRID,BID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,LastModBy from RatePlanRef;"
mysqlverify "b1" "${RRCTX}"		"RatePlanRefRTRate"	    		"select * from RatePlanRefRTRate;"
mysqlverify "c1" "${RRCTX}"		"RatePlanRefSPRate"	    		"select * from RatePlanRefSPRate;"
mysqlverify "d1" "${RRCTX}"		"Assessments"	            	"select ASMID,BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
mysqlverify "e1" "${RRCTX}"		"PaymentTypes"	            	"select PMTID,BID,Name,Description,LastModBy from PaymentType;"
mysqlverify "f1" "${RRCTX}"		"PaymentAllocations"	    	"select * from ReceiptAllocation order by Amount ASC;"
mysqlverify "g1" "${RRCTX}"		"Receipts"	            		"select RCPTID,BID,PMTID,DEPID,DID,Dt,Amount,AcctRuleReceive,Comment,LastModBy from Receipt;"
mysqlverify "h1" "${RRCTX}"		"CustomAttributes"	    		"select CID,BID,Type,Name,Value,LastModBy from CustomAttr;"
mysqlverify "i1" "${RRCTX}"		"CustomAttributesAssignment" 	"select * from CustomAttrRef;"
mysqlverify "j1" "${RRCTX}"		"NoteTypes"	            		"select NTID,BID,Name,LastModBy from NoteType;"
mysqlverify "k1" "${RRCTX}"		"Deposits"	            		"select DID,BID,Dt,DEPID,Amount,LastModBy from Deposit;"


logcheck
