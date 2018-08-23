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
mysqlverify  "${RRCTX}"  "NewBusinesses"	           "select BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,LastModBy from Business;"
mysqlverify  "${RRCTX}"  "StringLists"	               "select SLID,BID,Name,LastModBy from StringList;"
mysqlverify  "${RRCTX}"  "SLString"	            	   "select SLSID,SLID,Value,LastModBy from SLString;"
mysqlverify  "${RRCTX}"  "RentableTypes"	           "select RTID,BID,Style,Name,RentCycle,Proration,GSRPC,FLAGS,LastModBy from RentableTypes;"
mysqlverify  "${RRCTX}"  "RentableMarketRates"	       "select * from RentableMarketRate;"
mysqlverify  "${RRCTX}"  "Deposit Methods"             "select * from DepositMethod;"
mysqlverify  "${RRCTX}"  "Sources"	            	   "select SourceSLSID,BID,Name,Industry from DemandSource;"
mysqlverify  "${RRCTX}"  "RentableSpecialtyTypes"      "select * from RentableSpecialty;"
mysqlverify  "${RRCTX}"  "Buildings"	               "select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
mysqlverify  "${RRCTX}"  "Depositories"	               "select DEPID,BID,Name,AccountNo,LastModBy from Depository;"
mysqlverify  "${RRCTX}"  "Transactants"	               "select TCID,BID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
mysqlverify  "${RRCTX}"  "Users"	                   "select TCID,Points,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyContactEmail,AlternateEmailAddress,EligibleFutureUser,Industry,SourceSLSID from User;"
mysqlverify  "${RRCTX}"  "Payors"	                   "select TCID,CreditLimit,TaxpayorID,LastModBy from Payor;"
mysqlverify  "${RRCTX}"  "Prospects"	               "select TCID,CompanyAddress,CompanyCity,CompanyState,CompanyPostalCode,CompanyEmail,CompanyPhone,Occupation,ThirdPartySource,LastModBy from Prospect;"
mysqlverify  "${RRCTX}"  "Vehicles"	            	   "select VID,TCID,VehicleType,VehicleMake,VehicleModel,VehicleColor,VehicleYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DtStart,DtStop,LastModBy from Vehicle;"
mysqlverify  "${RRCTX}"  "Rentables"	               "select RID,BID,RentableName,AssignmentTime,LastModBy from Rentable;"
mysqlverify  "${RRCTX}"  "RentableTypeRef"	    	   "select RID,BID,RTID,OverrideRentCycle,OverrideProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
mysqlverify  "${RRCTX}"  "RentableStatus"	           "select RID,UseStatus,LeaseStatus,DtStart,DtStop,LastModBy from RentableStatus;"
mysqlverify  "${RRCTX}"  "RentalAgreementTemplates"    "select RATID,BID,RATemplateName,LastModBy from RentalAgreementTemplate;"
mysqlverify  "${RRCTX}"  "RentalAgreements"	    	   "select RAID,RATID,BID,AgreementStart,AgreementStop,Renewal,SpecialProvisions,UnspecifiedAdults,UnspecifiedChildren,LastModBy from RentalAgreement;"
mysqlverify  "${RRCTX}"  "Pets"	                       "select PETID,BID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from Pets;"
mysqlverify  "${RRCTX}"  "Notes"	                   "select NID,PNID,Comment,LastModBy from Notes;"
mysqlverify  "${RRCTX}"  "AgreementRentables"	       "select * from RentalAgreementRentables;"
mysqlverify  "${RRCTX}"  "AgreementPayors"	    	   "select * from RentalAgreementPayors;"
mysqlverify  "${RRCTX}"  "ChartOfAccounts"	    	   "select LID,PLID,BID,RAID,GLNumber,FLAGS,Name,AcctType,AllowPost,LastModBy from GLAccount;"
mysqlverify  "${RRCTX}"  "LedgerMarkers"	           "select LMID,LID,BID,RAID,RID,TCID,Dt,Balance,State,LastModBy from LedgerMarker;"
mysqlverify  "${RRCTX}"  "RatePlan"	            	   "select RPID,BID,Name,LastModBy from RatePlan;"
mysqlverify  "${RRCTX}"  "RatePlanRef"	               "select RPRID,BID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,LastModBy from RatePlanRef;"
mysqlverify  "${RRCTX}"  "RatePlanRefRTRate"	 	   "select * from RatePlanRefRTRate;"
mysqlverify  "${RRCTX}"  "RatePlanRefSPRate"	  	   "select * from RatePlanRefSPRate;"
mysqlverify  "${RRCTX}"  "Assessments"	               "select ASMID,BID,RID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
mysqlverify  "${RRCTX}"  "PaymentTypes"	               "select PMTID,BID,Name,Description,LastModBy from PaymentType;"
mysqlverify  "${RRCTX}"  "PaymentAllocations"	       "select * from ReceiptAllocation order by Amount ASC;"
mysqlverify  "${RRCTX}"  "Receipts"	            	   "select RCPTID,BID,PMTID,DEPID,DID,Dt,Amount,AcctRuleReceive,Comment,LastModBy from Receipt;"
mysqlverify  "${RRCTX}"  "CustomAttributes"	    	   "select CID,BID,Type,Name,Value,LastModBy from CustomAttr;"
mysqlverify  "${RRCTX}"  "CustomAttributesAssignment"  "select * from CustomAttrRef;"
mysqlverify  "${RRCTX}"  "NoteTypes"	               "select NTID,BID,Name,LastModBy from NoteType;"
mysqlverify  "${RRCTX}"  "Deposits"	            	   "select DID,BID,Dt,DEPID,Amount,LastModBy from Deposit;"


logcheck
