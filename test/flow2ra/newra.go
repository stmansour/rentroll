package main

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"rentroll/ws"
)

var nra = string(`{
    "dates": {
        "AgreementStart": "7/25/2018",
        "AgreementStop": "7/25/2019",
        "BID": 1,
        "CSAgent": 72,
        "PossessionStart": "7/25/2018",
        "PossessionStop": "7/25/2019",
        "RentStart": "7/25/2018",
        "RentStop": "7/25/2019"
    },
    "meta": {
        "Approver1": 72,
        "Approver1Name": "Yolanda Hernandez",
        "Approver2": 296,
        "Approver2Name": "Fritch David",
        "DecisionDate1": "2018-07-22 18:46:00 UTC",
        "DecisionDate2": "2018-07-22 23:12:00 UTC",
        "DeclineReason1": 0,
        "DeclineReason2": 0,
        "DocumentDate": "2018-07-25 17:00:00 UTC",
        "HavePets": true,
        "HaveVehicles": true,
        "LastTMPASMID": 9,
        "LastTMPPETID": 2,
        "LastTMPTCID": 1,
        "LastTMPVID": 1,
        "LeaseTerminationReason": 0,
        "NoticeToMoveDate": "1900-01-01 00:00:00 UTC",
        "NoticeToMoveName": "",
        "NoticeToMoveReported": "1900-01-01 00:00:00 UTC",
        "NoticeToMoveUID": 0,
        "RAFLAGS": 52,
        "RAID": 0,
        "TerminationDate": "1900-01-01 00:00:00 UTC",
        "TerminatorName": "",
        "TerminatorUID": 0
    },
    "parentchild": [],
    "people": [
        {
            "Address": "",
            "Address2": "",
            "AlternateEmailAddress": "",
            "BID": 1,
            "Bankruptcy": false,
            "BankruptcyDes": "",
            "CellPhone": "111-867-5309",
            "City": "",
            "Comment": "",
            "CompanyAddress": "123 Elm Street",
            "CompanyCity": "Anytown",
            "CompanyEmail": "zeke@gunsrus.com",
            "CompanyName": "",
            "CompanyPhone": "987-654-3210",
            "CompanyPostalCode": "54321",
            "CompanyState": "IL",
            "Convicted": false,
            "ConvictedDes": "",
            "Country": "",
            "CreditLimit": 150,
            "CurrentAddress": "",
            "CurrentLandLordName": "",
            "CurrentLandLordPhoneNo": "",
            "CurrentLengthOfResidency": "",
            "CurrentReasonForMoving": 0,
            "DateofBirth": "3/4/1956",
            "DriversLicense": "IL123456789",
            "EligibleFuturePayor": true,
            "EligibleFutureUser": true,
            "EmergencyContactAddress": "543 Oak St",
            "EmergencyContactEmail": "sallybb@hicks.com",
            "EmergencyContactName": "Sally Bob Bueford",
            "EmergencyContactTelephone": "643-432-3421",
            "Evicted": false,
            "EvictedDes": "",
            "FirstName": "William",
            "GrossIncome": 28500,
            "Industry": 367,
            "IsCompany": false,
            "IsGuarantor": false,
            "IsOccupant": true,
            "IsRenter": true,
            "LastName": "Thorton",
            "MiddleName": "Robert",
            "Occupation": "",
            "OtherPreferences": "",
            "Points": 0,
            "PostalCode": "",
            "PreferredName": "Billybob",
            "PrimaryEmail": "pyro@fireguy.com",
            "PriorAddress": "",
            "PriorLandLordName": "",
            "PriorLandLordPhoneNo": "",
            "PriorLengthOfResidency": "",
            "PriorReasonForMoving": 0,
            "SecondaryEmail": "",
            "SourceSLSID": 28,
            "SpecialNeeds": "I need lots of stuff",
            "State": "",
            "TCID": 0,
            "TMPTCID": 1,
            "TaxpayorID": "123456789",
            "ThirdPartySource": 0,
            "Website": "",
            "WorkPhone": "123-456-7890"
        }
    ],
    "pets": [
        {
            "BID": 0,
            "Breed": "dog",
            "Color": "brown",
            "DtStart": "7/25/2018",
            "DtStop": "7/25/2019",
            "Fees": [
                {
                    "ARID": 23,
                    "ARName": "Pet Fee",
                    "ASMID": 0,
                    "AtSigningPreTax": 0,
                    "Comment": "",
                    "ContractAmount": 50,
                    "RentCycle": 0,
                    "SalesTax": 0,
                    "Start": "7/25/2018",
                    "Stop": "7/25/2018",
                    "TMPASMID": 1,
                    "TransOccTax": 0
                },
                {
                    "ARID": 24,
                    "ARName": "Pet Rent",
                    "ASMID": 0,
                    "AtSigningPreTax": 0,
                    "Comment": "prorated for 7 of 31 days",
                    "ContractAmount": 2.26,
                    "RentCycle": 6,
                    "SalesTax": 0,
                    "Start": "7/25/2018",
                    "Stop": "7/25/2018",
                    "TMPASMID": 2,
                    "TransOccTax": 0
                },
                {
                    "ARID": 24,
                    "ARName": "Pet Rent",
                    "ASMID": 0,
                    "AtSigningPreTax": 0,
                    "Comment": "",
                    "ContractAmount": 10,
                    "RentCycle": 6,
                    "SalesTax": 0,
                    "Start": "8/1/2018",
                    "Stop": "7/25/2019",
                    "TMPASMID": 3,
                    "TransOccTax": 0
                }
            ],
            "Name": "Beauregard",
            "PETID": 0,
            "TMPPETID": 1,
            "TMPTCID": 1,
            "Type": "bloodhound",
            "Weight": 90
        }
    ],
    "rentables": [
        {
            "AtSigningPreTax": 0,
            "BID": 1,
            "Fees": [
                {
                    "ARID": 16,
                    "ARName": "Gas Base Fee",
                    "ASMID": 0,
                    "AtSigningPreTax": 0,
                    "Comment": "",
                    "ContractAmount": 50,
                    "RentCycle": 6,
                    "SalesTax": 0,
                    "Start": "7/25/2018",
                    "Stop": "7/25/2019",
                    "TMPASMID": 9,
                    "TransOccTax": 0
                },
                {
                    "ARID": 43,
                    "ARName": "Rent ST003",
                    "ASMID": 0,
                    "AtSigningPreTax": 0,
                    "Comment": "",
                    "ContractAmount": 2500,
                    "RentCycle": 6,
                    "SalesTax": 0,
                    "Start": "7/25/2018",
                    "Stop": "7/25/2019",
                    "TMPASMID": 8,
                    "TransOccTax": 0
                }
            ],
            "RID": 5,
            "RTFLAGS": 4,
            "RTID": 1,
            "RentCycle": 6,
            "RentableName": "Rentable001",
            "SalesTax": 0,
            "TransOccTax": 0
        }
    ],
    "tie": {
        "people": [
            {
                "BID": 1,
                "PRID": 5,
                "TMPTCID": 1
            }
        ]
    },
    "vehicles": [
        {
            "BID": 0,
            "DtStart": "7/25/2018",
            "DtStop": "7/25/2019",
            "Fees": [
                {
                    "ARID": 39,
                    "ARName": "Vehicle Registration Fee",
                    "ASMID": 0,
                    "AtSigningPreTax": 0,
                    "Comment": "",
                    "ContractAmount": 10,
                    "RentCycle": 0,
                    "SalesTax": 0,
                    "Start": "7/25/2018",
                    "Stop": "7/25/2018",
                    "TMPASMID": 7,
                    "TransOccTax": 0
                }
            ],
            "LicensePlateNumber": "BR 549",
            "LicensePlateState": "LA",
            "ParkingPermitNumber": "",
            "TMPTCID": 1,
            "TMPVID": 1,
            "VID": 0,
            "VIN": "29385723987235253",
            "VehicleColor": "Silver",
            "VehicleMake": "Chevrolet",
            "VehicleModel": "Silverado",
            "VehicleType": "Pickup Truck",
            "VehicleYear": 2017
        }
    ]
}`)

// DoNewRA creates a new ORIGIN rental agreement from the JSON data in nra.
//
// INPUTS
//     ctx  - database context for transactions
//     s    - a session
//
// RETURNS
//     nraid - new RAID
//     err   - any errors encountered
//--------------------------------------------------------------------------
func DoNewRA(ctx context.Context, s *rlib.Session) {
	var bid = int64(1)
	rlib.Console("New Flow\n")
	b := []byte(nra)
	a := rlib.Flow{
		BID:       bid,
		FlowID:    0, // it's new flowID,
		UserRefNo: rlib.GenerateUserRefNo(),
		FlowType:  rlib.RAFlow,
		Data:      b,
		CreateBy:  s.UID,
		LastModBy: s.UID,
	}

	// insert new flow
	flowID, err := rlib.InsertFlow(ctx, &a)
	if err != nil {
		fmt.Printf("Error while inserting Flow: %s\n", err.Error())
		return
	}

	tx, tctx, err := rlib.NewTransactionWithContext(ctx)
	if err != nil {
		fmt.Printf("Could not create transaction context: %s\n", err.Error())
		return
	}
	nraid, err := ws.Flow2RA(tctx, flowID)
	if err != nil {
		tx.Rollback()
		rlib.Console("Flow2RA error\n")
		fmt.Printf("Could not write Flow back to db: %s\n", err.Error())
		return
	}
	if err = tx.Commit(); err != nil {
		fmt.Printf("Error committing transaction: %s\n", err.Error())
		return
	}
	rlib.Console("Successfully created new Rental Agreement, RAID = %d\n", nraid)
	rlib.Console("Removing flow: %d\n", flowID)
	if err = rlib.DeleteFlow(ctx, flowID); err != nil {
		fmt.Printf("Error deleting flow: %s\n", err.Error())
		return
	}
	rlib.Console("Completed without errors\n")

}
