#!/bin/bash

#---------------------------------------------------------------
# TOP is the directory where RentRoll begins. It is used
# in base.sh to set other useful directories such as ${BASHDIR}
#---------------------------------------------------------------
TOP=../..

TESTNAME="Web Services"
TESTSUMMARY="Test Web Services"
RRDATERANGE="-j 2016-07-01 -k 2016-08-01"

CREATENEWDB=0

#---------------------------------------------------------------
#  Use the testdb for these tests...
#---------------------------------------------------------------
echo "Create new database..."
mysql --no-defaults rentroll < wsdb.sql
source ../share/base.sh

echo "STARTING RENTROLL SERVER"
RENTROLLSERVERAUTH="-noauth"
RENTROLLSERVERNOW="-testDtNow 1/1/2015"

startRentRollServer

TFILES="a"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    # get GLAccounts list for the business
    dojsonGET "http://localhost:8270/v1/accountlist/2" "ws1" "WebService--GetAccountsListForBusiness"

    # get parent accounts list for the business
    dojsonGET "http://localhost:8270/v1/parentaccounts/2" "ws2" "WebService--GetParentAccountsListForBusiness"

    # get post accounts list for the business
    dojsonGET "http://localhost:8270/v1/postaccounts/2" "ws3" "WebService--GetPostAccountsListForBusiness"

    # Get Chart of Accounts
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/accounts/2" "request" "ws4"  "WebService--ChartOfAccounts"

    # Get Account details
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22accountForm%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/account/2/97" "request" "ws5" "WebService--ChartOfAccounts-detail"

    # Create new Account
    echo "request%3D%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22%22%2C%22record%22%3A%7B%22LID%22%3A0%2C%22BID%22%3A2%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%22123456789%22%2C%22Name%22%3A%22SmokeTest%20GLAccount%22%2C%22AcctType%22%3A%22%22%2C%22Description%22%3A%22%22%2C%22LastModTime%22%3A%221%2F1%2F1900%22%2C%22LastModBy%22%3A0%2C%22BUD%22%3A%22%22%2C%22PLID%22%3A0%2C%22Status%22%3A0%2C%22Type%22%3A0%2C%22AllowPost%22%3Afalse%2C%22ManageToBudget%22%3Afalse%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/account/2/0" "request" "ws6"  "WebService--CreateGLAccount"

    # Update Account details
    echo "request%3D%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22%22%2C%22record%22%3A%7B%22LID%22%3A108%2C%22BID%22%3A2%2C%22RAID%22%3A0%2C%22TCID%22%3A0%2C%22GLNumber%22%3A%229876543210%22%2C%22Name%22%3A%22SmokeTest%20GLAccount%22%2C%22AcctType%22%3A%22%22%2C%22Description%22%3A%22Update%20this%20Account%20(Smoke%20Test)%22%2C%22LastModTime%22%3A%221%2F1%2F1900%22%2C%22LastModBy%22%3A0%2C%22BUD%22%3A%22%22%2C%22PLID%22%3A0%2C%22Status%22%3A0%2C%22Type%22%3A0%2C%22AllowPost%22%3Atrue%2C%22ManageToBudget%22%3Afalse%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/account/2/108" "request" "ws7"  "WebService--UpdateGLAccount"

    # Delete Account
    echo "request%3D%7B%22cmd%22%3A%22delete%22%2C%22LID%22%3A97%7D" > request
    dojsonPOST "http://localhost:8270/v1/account/2/" "request" "ws8"  "WebService--DeleteGLAccount"

    # Get Transactants
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22sort%22%3A%5B%7B%22field%22%3A%22LastName%22%2C%22direction%22%3A%22asc%22%7D%5D%7D" > request
    dojsonPOST "http://localhost:8270/v1/transactants/1" "request" "ws9"  "WebService--GetTransactants"

    # Get Rentables
    #  Get Rentables for Business 1, between 8/1/2016 and 9/1/2016
    #
    #  Expected Results:
    #  There are 4 rentables: 309 Rexford, 309 1/2 Rexford, 311 Rexford, 311 1/2 Rexford
    #------------------------------------------------------------------------------------
    # echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%228%2F1%2F2016%22%2C%22searchDtStop%22%3A%229%2F1%2F2016%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentables/1" "request" "ws10"  "WebService--GetRentables"

    # Get Receipts
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22searchDtStart%22%3A%222016-08-01%22%2C%22searchDtStop%22%3A%222016-09-01%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/receipts/1" "request" "ws11"  "WebService--GetReceipts"

    # Get Assessments
    echo "request%3d%7b%22cmd%22%3a%22get%22%2c%22selected%22%3a%5b%5d%2c%22limit%22%3a100%2c%22offset%22%3a0%7d" > request
    dojsonPOST "http://localhost:8270/v1/asms/1" "request" "ws12"  "WebService--GetAssessments"

    # Get Assessment 1 from REX
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/asm/REX/1" "request" "ws13"  "WebService--GetAnAssessment"

    # Save the Assessment with an updated comment
    echo "request=%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmInstForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22ASMID%22%3A43%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PASMID%22%3A14%2C%22RID%22%3A1%2C%22Rentable%22%3A%22309%20Rexford%22%2C%22RAID%22%3A1%2C%22Amount%22%3A3750%2C%22Start%22%3A%2212%2F1%2F2016%22%2C%22Stop%22%3A%2212%2F2%2F2016%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22InvoiceNo%22%3A0%2C%22ARID%22%3A0%2C%22Comment%22%3A%22comment%20by%20sman%22%2C%22LastModTime%22%3A%226%2F6%2F2017%22%2C%22LastModBy%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/asm/REX/1" "request" "ws14"  "WebService--SaveAnAssessment"

    # Get Receipt 5 from REX
    echo "%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/receipt/REX/5" "request" "ws15"  "WebService--GetAReceipt"

    # Save the Receipt 5 with an updated comment
    # echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22AcctRule%22%3A%22%22%2C%22Amount%22%3A3550%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Comment%22%3A%22This%20comment%20was%20updated%20by%20a%20web%20service%20test%22%2C%22DocNo%22%3A%221631%22%2C%22Dt%22%3A%221%2F4%2F2016%22%2C%22LastModBy%22%3A0%2C%22LastModTime%22%3A%222%2F23%2F2017%22%2C%22OtherPayorName%22%3A%22%22%2C%22PMTID%22%3A1%2C%22PRCPTID%22%3A0%2C%22RAID%22%3A2%2C%22RCPTID%22%3A5%2C%22recid%22%3A0%7D%7D" > request
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A5%2C%22PRCPTID%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A1%2C%22Payor%22%3A%22Rita+Costea+(TCID%3A+3)%22%2C%22TCID%22%3A3%2C%22Dt%22%3A%221%2F4%2F2016%22%2C%22DocNo%22%3A%221631%22%2C%22Amount%22%3A3550%2C%22ARID%22%3A0%2C%22Comment%22%3A%22update+comment%22%2C%22OtherPayorName%22%3A%22%22%2C%22LastModTime%22%3A%227%2F18%2F2017%22%2C%22LastModBy%22%3A0%2C%22FLAGS%22%3A0%2C%22PmtTypeName%22%3A1%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/receipt/REX/5" "request" "ws16"  "WebService--SaveAReceipt"

    # Create a NEW RECEIPT
    # echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22AcctRule%22%3A%22%22%2C%22Amount%22%3A1590.32%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Comment%22%3A%22This%20is%20a%20NEW%20RECEIPT%20added%20by%20a%20web%20test%22%2C%22DocNo%22%3A%229876%22%2C%22Dt%22%3A%222%2F24%2F2017%22%2C%22LastModBy%22%3A0%2C%22LastModTime%22%3A%222%2F24%2F2017%22%2C%22OtherPayorName%22%3A%22%22%2C%22PMTID%22%3A1%2C%22PRCPTID%22%3A0%2C%22RAID%22%3A2%2C%22RCPTID%22%3A0%2C%22recid%22%3A0%7D%7D" > request
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22Dt%22%3A%227%2F18%2F2017%22%2C%22DocNo%22%3A%222345%22%2C%22Payor%22%3A%22Aaron+Read+(TCID%3A+1)%22%2C%22TCID%22%3A1%2C%22Amount%22%3A3750%2C%22Comment%22%3A%22rent%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/receipt/REX/0" "request" "ws17"  "WebService--InsertAReceipt"

    # This receipt should FAIL -- TCID is 0
    echo "%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22receiptForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22RCPTID%22%3A0%2C%22PRCPTID%22%3A0%2C%22ARID%22%3A3%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22PMTID%22%3A2%2C%22PmtTypeName%22%3A2%2C%22Dt%22%3A%227%2F19%2F2017%22%2C%22DocNo%22%3A%222345%22%2C%22Payor%22%3A%22%22%2C%22TCID%22%3A0%2C%22Amount%22%3A40%2C%22Comment%22%3A%22%22%2C%22OtherPayorName%22%3A%22%22%2C%22FLAGS%22%3A0%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/receipt/REX/0" "request" "ws18"  "WebService--InsertAReceipt-Fai"

    # Create a NEW ASSESSMENT
    echo "request=%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22asmEpochForm%22%2C%22record%22%3A%7B%22ARID%22%3A2%2C%22recid%22%3A0%2C%22RID%22%3A1%2C%22ASMID%22%3A0%2C%22PASMID%22%3A0%2C%22ATypeLID%22%3A0%2C%22InvoiceNo%22%3A0%2C%22RAID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22Start%22%3A%226%2F6%2F2017%22%2C%22Stop%22%3A%226%2F6%2F2017%22%2C%22RentCycle%22%3A6%2C%22ProrationCycle%22%3A4%2C%22TCID%22%3A0%2C%22Amount%22%3A60%2C%22AcctRule%22%3A%22%22%2C%22Comment%22%3A%22%22%2C%22LastModTime%22%3A%226%2F6%2F2017%22%2C%22LastModBy%22%3A0%2C%22Rentable%22%3A%22309%2BRexford%22%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/asm/1/0" "request" "ws19"  "WebService--InsertAnAssessment"

    # Test Transactant Typedown
    dojsonGET "http://localhost:8270/v1/transactantstd/ISO?request=%7B%22search%22%3A%22s%22%2C%22max%22%3A250%7D" "ws20" "WebService--GetTransactantTypeDown"

    # Test Transactant Typedown
    dojsonGET "http://localhost:8270/v1/rentablestd/ISO?request%3D%7B%22search%22%3A%226%22%2C%22max%22%3A250%7D" "ws30" "WebService--GetRentableTypeDown"

    # Search Payment Types
    echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/pmts/1" "request" "ws31"  "WebService--PaymentTypes-SearchAll"

    # Get Specificy PaymentType - FORCE ERROR - no PMTID provided
    echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22paymentTypeGrid%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/pmts/1" "request" "ws32"  "WebService--PaymentTypes-Get-ForceError"

    # Get Specificy PaymentType
    echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22paymentTypeGrid%22%7D" > request
    dojsonPOST "http://localhost:8270/v1/pmts/1/1" "request" "ws33"  "WebService--PaymentTypes-Get"

    # Get Specificy PaymentType
    # echo "request%3D%7B%22cmd%22%3A%22get%22%2C%22recid%22%3A0%2C%22name%22%3A%22paymentTypeGrid%22%7D" > request
    # dojsonPOST "http://localhost:8270/v1/pmt/1/1" "request" "ws34"  "WebService--PaymentTypes"

    # get UI Values...
    doPlainGET "http://localhost:8270/v1/uival/REX/app.AssessmentRules" "ws35" "WebService--GetUIValue-app.AssessmentRules"
    doPlainGET "http://localhost:8270/v1/uival/REX/app.ReceiptRules" "ws36" "WebService--GetUIValue-app.ReceiptRules"

    # rental Agreement typedown...
    dojsonGET "http://localhost:8270/v1/rentalagrtd/CCC?request=%7B%22search%22%3A%22s%22%2C%22max%22%3A250%7D" "ws37" "WebService--GetRentalAgreementTypeDown"

    # get Rentable types list for a business
    dojsonGET "http://localhost:8270/v1/rtlist/2" "ws38" "WebService--GetRentableTypesForBusiness"

    # save rentable
    echo "request=%7B%22cmd%22%3A%22save%22%2C%22recid%22%3A0%2C%22name%22%3A%22rentableForm%22%2C%22record%22%3A%7B%22recid%22%3A0%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A0%2C%22RentableName%22%3A%22REX-Test-1%22%2C%22AssignmentTime%22%3A1%7D%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentable/1/0" "request" "ws39"  "WebService--SaveRentable"

    # save rentable UseStatus with usestatus: 2(Administrative)
    encodeRequest '{"cmd":"save","selected":[],"limit":100,"offset":0,"changes":[{"recid":0,"RSID":0,"BID":1,"BUD":"REX","RID":1129,"UseStatus":2,"DtStart":"1/1/2016","DtStop":"1/1/9999","DtNoticeToVacate":"1/1/1900","DtNoticeToVacateIsSet":false,"CreateBy":0,"LastModBy":0}]}' > request
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1129" "request" "ws40"  "WebService--SaveRentableUseStatus-Rentable(1129)"

    # save rentable UseType with usetype: 101(Administrative)
    encodeRequest '{"cmd":"save","selected":[],"limit":100,"offset":0,"changes":[{"recid":0,"UTID":0,"BID":1,"BUD":"REX","RID":1129,"UseType":101,"DtStart":"1/1/2016","DtStop":"1/1/9999","CreateBy":0,"LastModBy":0}]}' > request
    dojsonPOST "http://localhost:8270/v1/rentableusetype/1/1129" "request" "ws40a"  "WebService--SaveRentableUseType-Rentable(1129)"

    # save rentable Lease status with LeaseStatus: Reserved
    encodeRequest '{"cmd":"save","recid":0,"name":"bookResForm","record":{"BID":1,"BUD":"","RTID":1,"RID":1129,"Nights":1,"DtStart":"1/1/2016","DtStop":"12/31/9999","RLID":0,"FirstName":"Steve","LastName":"Mansour","Email":"sman@stevemansour.com","Phone":"4089219957","Street":"","City":"Milpitas","Country":"United+States","State":"CA","PostalCode":"95035","CCName":"STEVE MANSOUR","CCType":"VISA","CCNumber":"1234111112341111","CCExpMonth":"09","CCExpYear":"2014","Comment":"","RentableName":"Rentable010","Address":"2215+Wellington+Drive","Address2":""}}' > request
    dojsonPOST "http://localhost:8270/v1/reservation/1/1129" "request" "ws40b"  "WebService--SaveRentableLeaseStatus-Rentable(1129)"

    # save rentable type ref
    echo "%7B%22cmd%22%3A%22save%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%2C%22changes%22%3A%5B%7B%22recid%22%3A0%2C%22RTRID%22%3A0%2C%22RTID%22%3A1%2C%22BID%22%3A1%2C%22BUD%22%3A%22REX%22%2C%22RID%22%3A1129%2C%22OverrideRentCycle%22%3A0%2C%22OverrideProrationCycle%22%3A0%2C%22DtStart%22%3A%221%2F1%2F2016%22%2C%22DtStop%22%3A%221%2F1%2F9999%22%2C%22CreateBy%22%3A0%2C%22LastModBy%22%3A0%7D%5D%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1129" "request" "ws41"  "WebService--SaveRentableTypeRef-Rentable(1129)"

    # Get Rentables
    # echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    encodeRequest '{"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"8/1/2016","searchDtStop":"9/1/2016"}' > request
    dojsonPOST "http://localhost:8270/v1/rentables/1" "request" "ws42"  "WebService--GetRentables"

    # Get Rentable Use Status list for Rentable(1129)
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1129" "request" "ws43"  "WebService--GetRentableUseStatus"

    # Get Rentable Type ref list for Rentable(1129)
    echo "request=%7B%22cmd%22%3A%22get%22%2C%22selected%22%3A%5B%5D%2C%22limit%22%3A100%2C%22offset%22%3A0%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1129" "request" "ws44"  "WebService--GetRentableTypeRef"

    # delete rentable use status recently created (1129)
    echo "%7B%22cmd%22%3A%22delete%22%2C%22RSIDList%22%3A%5B1129%5D%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentableusestatus/1/1129" "request" "ws45" "WebService--DeleteRentableUseStatus(1129)"

    # delete rentable type ref recently created (1129)
    echo "%7B%22cmd%22%3A%22delete%22%2C%22RTRIDList%22%3A%5B1129%5D%7D" > request
    dojsonPOST "http://localhost:8270/v1/rentabletyperef/1/1129" "request" "ws46" "WebService--DeleteRentableTypeRef(1129)"
fi

#------------------------------------------------------------------------------
#  TEST b
#
#  Test typedown, transactant and transactantdet
#
#  Scenario:
#  Just make a query and validate that the return names and data are correct.
#
#  Expected Results:
#   1.  typedown is done as a "contains" match. So any substring match
#       on the names of transactants must match
#------------------------------------------------------------------------------
TFILES="b"
STEP=0
if [ "${SINGLETEST}${TFILES}" = "${TFILES}" -o "${SINGLETEST}${TFILES}" = "${TFILES}${TFILES}" ]; then
    #------------------------------------------------------------------
    # Test Transactant Typedown for details.  Search string is "ga".
    # Limit is set to 10
    #------------------------------------------------------------------
    dojsonGET "http://localhost:8270/v1/transactantsdettd/ISO?request=%7B%22search%22%3A%22ga%22%2C%22max%22%3A10%7D" "${TFILES}${STEP}" "WebService--GetTransactantDetailsTypeDown"

fi

stopRentRollServer
echo "RENTROLL SERVER STOPPED"

echo "Restoring test database..."
mysql --no-defaults rentroll < wsdb.sql

logcheck
