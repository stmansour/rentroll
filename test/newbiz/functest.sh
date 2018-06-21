#!/bin/bash

TESTNAME="CSV Loader Test"
TESTSUMMARY="Load all csv files through loader and validate the database after loading"

source ../share/base.sh
RRCTX="-G ${BUD} -g 12/1/15,1/1/16"

echo "CSVLOAD = ${CSVLOAD}"

# Create a bunch of content
${CSVLOAD} -b nb.csv >>${LOGFILE} 2>&1
${CSVLOAD} -c coa.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -R rt.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -u custom.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -d depository.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -s specialties.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -D bldg.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -p people.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -r rentable.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -T rat.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -C ra.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -E pets.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -V vehicle.csv  ${RRCTX} >>${LOGFILE} 2>&1
${CSVLOAD} -ar ar.csv  ${RRCTX} >>${LOGFILE} 2>&1
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


mysqlverify "a"  "-b nb.csv"           		"NewBusinesses"	            	"select BID,BUD,Name,DefaultRentCycle,DefaultProrationCycle,DefaultGSRPC,LastModBy from Business;"
mysqlverify "b"  "-l strlists.csv"     		"StringLists"	            	"select SLID,BID,Name,LastModBy from StringList;"
mysqlverify "c"  " "	               		"SLString"	            		"select SLSID,SLID,Value,LastModBy from SLString;"
mysqlverify "d"  "-R rt.csv"           		"RentableTypes"	            	"select RTID,BID,Style,Name,RentCycle,Proration,GSRPC,FLAGS,LastModBy from RentableTypes;"
mysqlverify "e"  " "                   		"RentableMarketRates"	    	"select * from RentableMarketRate;"
mysqlverify "f"  "-m depmeth.csv"      		"Deposit Methods"           	"select * from DepositMethod;"
mysqlverify "g"  "-S sources.csv"      		"Sources"	            		"select SourceSLSID,BID,Name,Industry from DemandSource;"
mysqlverify "h"  "-s specialties.csv"  		"RentableSpecialtyTypes"    	"select * from RentableSpecialty;"
mysqlverify "i"  "-D bldg.csv"         		"Buildings"	            		"select BLDGID,BID,Address,Address2,City,State,PostalCode,Country,LastModBy from Building;"
mysqlverify "j"  "-d depository.csv"   		"Depositories"	            	"select DEPID,BID,LID,Name,AccountNo,LastModBy from Depository;"
mysqlverify "n"  "-p people.csv"       		"Transactants"	            	"select TCID,BID,FirstName,MiddleName,LastName,CompanyName,IsCompany,PrimaryEmail,SecondaryEmail,WorkPhone,CellPhone,Address,Address2,City,State,PostalCode,Country,LastModBy from Transactant;"
mysqlverify "o"  ""			                "Users"	                        "select TCID,Points,DateofBirth,EmergencyContactName,EmergencyContactAddress,EmergencyContactTelephone,EmergencyContactEmail,AlternateAddress,EligibleFutureUser,Industry,SourceSLSID from User;"
mysqlverify "p"  ""			                "Payors"	                	"select TCID,CreditLimit,TaxpayorID,ThirdPartySource,LastModBy from Payor;"
mysqlverify "q"  ""			                "Prospects"	            		"select TCID,CompanyAddress,CompanyCity,CompanyState,CompanyPostalCode,CompanyEmail,CompanyPhone,Occupation,LastModBy from Prospect;"
mysqlverify "na"  "-V vehicle.csv"       	"Vehicles"	            		"select VID,TCID,VehicleType,VehicleMake,VehicleModel,VehicleColor,VehicleYear,LicensePlateState,LicensePlateNumber,ParkingPermitNumber,DtStart,DtStop,LastModBy from Vehicle;"
mysqlverify "k"  "-r rentable.csv"     		"Rentables"	            		"select RID,BID,RentableName,AssignmentTime,LastModBy from Rentable;"
mysqlverify "l"  " "                   		"RentableTypeRef"	    		"select RID,BID,RTID,OverrideRentCycle,OverrideProrationCycle,DtStart,DtStop,LastModBy from RentableTypeRef;"
mysqlverify "m"  " "                   		"RentableStatus"	        	"select RID,UseStatus,LeaseStatus,DtStart,DtStop,LastModBy from RentableStatus;"
mysqlverify "r"  "-T rat.csv"          		"RentalAgreementTemplates"   	"select RATID,BID,RATemplateName,LastModBy from RentalAgreementTemplate;"
mysqlverify "s"  "-C ra.csv"           		"RentalAgreements"	    		"select RAID,RATID,BID,AgreementStart,AgreementStop,Renewal,SpecialProvisions,UnspecifiedAdults,UnspecifiedChildren,LastModBy from RentalAgreement;"
mysqlverify "t"  "-E pets.csv"         		"Pets"	                    	"select PETID,BID,RAID,Type,Breed,Color,Weight,Name,DtStart,DtStop,LastModBy from RentalAgreementPets;"
mysqlverify "u"  ""           		   	    "Notes"	                    	"select NID,PNID,Comment,LastModBy from Notes;"
mysqlverify "v"  " "                   		"AgreementRentables"	    	"select * from RentalAgreementRentables;"
mysqlverify "w"  " "                   		"AgreementPayors"	    		"select * from RentalAgreementPayors;"
mysqlverify "x"  "-c coa.csv"          		"ChartOfAccounts"	    		"select LID,PLID,BID,RAID,GLNumber,Name,AcctType,AllowPost,LastModBy from GLAccount;"
mysqlverify "xa"  "-ar ar.csv"          	"AccountRules"	    			"select ARID,BID,Name,ARType,DebitLID,CreditLID,Description,LastModBy from AR;"
mysqlverify "y"  " "                   		"LedgerMarkers"	            	"select LMID,LID,BID,RAID,RID,TCID,Dt,Balance,State,LastModBy from LedgerMarker;"
mysqlverify "z"  "-a rp.csv"           		"RatePlan"	            		"select RPID,BID,Name,LastModBy from RatePlan;"
mysqlverify "a1" "-f rprefs.csv"       		"RatePlanRef"	                "select RPRID,BID,RPID,DtStart,DtStop,FeeAppliesAge,MaxNoFeeUsers,AdditionalUserFee,PromoCode,CancellationFee,FLAGS,LastModBy from RatePlanRef;"
mysqlverify "b1" "-n rprtrate.csv"     		"RatePlanRefRTRate"	    		"select * from RatePlanRefRTRate;"
mysqlverify "c1" "-t rpsprate.csv"     		"RatePlanRefSPRate"	    		"select * from RatePlanRefSPRate;"
mysqlverify "d1" "-A asmt.csv ${RRCTX}"     "Assessments"	            	"select ASMID,BID,RID,ATypeLID,RAID,Amount,Start,Stop,RentCycle,ProrationCycle,AcctRule,Comment,LastModBy from Assessments;"
mysqlverify "e1" "-P pmt.csv"          		"PaymentTypes"	            	"select PMTID,BID,Name,Description,LastModBy from PaymentType;"
mysqlverify "f1" "-e rcpt.csv ${RRCTX}"     "ReceiptAllocations"	    	"select RCPTID,BID,RAID,Dt,Amount,ASMID,AcctRule from ReceiptAllocation order by Amount ASC;"
mysqlverify "g1" " "                   		"Receipts"	            		"select RCPTID,BID,TCID,PMTID,DEPID,DID,Dt,Amount,AcctRuleApply,Comment,LastModBy from Receipt;"
mysqlverify "h1" "-u custom.csv"       		"CustomAttributes"	    		"select CID,BID,Type,Name,Value,LastModBy from CustomAttr;"
mysqlverify "i1" "-U assigncustom.csv" 		"CustomAttributesAssignment" 	"select * from CustomAttrRef;"
mysqlverify "j1" "-O nt.csv"           		"NoteTypes"	            		"select NTID,BID,Name,LastModBy from NoteType;"
mysqlverify "k1" "-y deposit.csv ${RRCTX}"  "Deposits"		            	"select DID,BID,Dt,DEPID,Amount,LastModBy from Deposit;"


logcheck
