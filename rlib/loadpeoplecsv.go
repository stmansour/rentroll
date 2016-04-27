package rlib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// PeopleSpecialty is the structure for attributes of a rentable specialty

// CSV file format:
//  |<---------------------------------------------------  TRANSACTANT ------------------------------------------------------------------->|  |<-------------------------------------------------------------------------------------------------------------  Tenant  ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------->|  |<------------------------------------------------------------------------- Payor ---------------------------------------------------->|  -- prospect --
//   0           1          2          3              4              5          6          7        8        9     10      11          12      13       14       15        16       17        18                  19                  20                 21             22           23                    24                      25                           26                27                28                      29       30         31                     32           33              34                  35              36              37             38               39           40           41
// 	FirstName, MiddleName, LastName, PrimaryEmail, SecondaryEmail, WorkPhone, CellPhone, Address, Address2, City, State, PostalCode, Country, Points, CarMake, CarModel, CarColor, CarYear, LicensePlateState, LicensePlateNumber, ParkingPermitNumber, AccountRep, DateofBirth, EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail, AlternateAddress, ElibigleForFutureOccupancy, Industry, Source, InvoicingCustomerNumber, CreditLimit, EmployerName, EmployerStreetAddress, EmployerCity, EmployerState, EmployerPostalCode, EmployerEmail, EmployerPhone, Occupation, ApplicationFee
// 	Edna,,Krabappel,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Ned,,Flanders,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Moe,,Szyslak,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Montgomery,,Burns,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Nelson,,Muntz,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Milhouse,,Van Houten,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Clancey,,Wiggum,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Homer,J,Simpson,homerj@springfield.com,,408-654-8732,,744 Evergreen Terrace,,Springfield,MO,64001,USA,5987,,Canyonero,red,,MO,BR549,,,,Marge Simpson,744 Evergreen Terrace,654=183-7946,,,,,,,,,,,,,,,,

// CreatePeopleFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreatePeopleFromCSV(sa []string) {
	// skip the header line
	if sa[0] == "FirstName" {
		return
	}
	var err error
	var tr Transactant
	var t Tenant
	var p Payor
	var pr Prospect
	var x float64
	dateform := "2006-01-02"

	for i := 0; i < len(sa); i++ {
		s := strings.TrimSpace(sa[i])
		switch {
		case i == 0: // transactant FirstName
			tr.FirstName = s
		case i == 1:
			tr.MiddleName = s
		case i == 2:
			tr.LastName = s
		case i == 3:
			tr.PrimaryEmail = s
		case i == 4:
			tr.SecondaryEmail = s
		case i == 5:
			tr.WorkPhone = s
		case i == 6:
			tr.CellPhone = s
		case i == 7:
			tr.Address = s
		case i == 8:
			tr.Address2 = s
		case i == 9:
			tr.City = s
		case i == 10:
			tr.State = s
		case i == 11:
			tr.PostalCode = s
		case i == 12:
			tr.Country = s
		case i == 13:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: Points value is invalid: %s\n", s)
					return
				}
				t.Points = int64(i)
			}
		case i == 14:
			t.CarMake = s
		case i == 15:
			t.CarModel = s
		case i == 16:
			t.CarColor = s
		case i == 17:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: CarYear value is invalid: %s\n", s)
					return
				}
				t.CarYear = int64(i)
			}
		case i == 18:
			t.LicensePlateState = s
		case i == 19:
			t.LicensePlateNumber = s
		case i == 20:
			t.ParkingPermitNumber = s
		case i == 21:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: AccountRep value is invalid: %s\n", s)
					return
				}
				t.AccountRep = int64(i)
			}
		case i == 22:
			if len(s) > 0 {
				t.DateofBirth, _ = time.Parse(dateform, s)
			}
		case i == 23:
			t.EmergencyContactName = s
		case i == 24:
			t.EmergencyContactAddress = s
		case i == 25:
			t.EmergencyContactTelephone = s
		case i == 26:
			t.EmergencyEmail = s
		case i == 27:
			t.AlternateAddress = s
		case i == 28:
			if len(s) > 0 {
				t.ElibigleForFutureOccupancy = yesnoToInt(s)
			}
		case i == 29:
			t.Industry = s
		case i == 30:
			t.Source = s
		case i == 31:
			t.InvoicingCustomerNumber = s
		case i == 32:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					Ulog("CreatePeopleFromCSV: Invalid Credit Limit value: %s\n", s)
					return
				}
				p.CreditLimit = x
			}
		case i == 33:
			p.EmployerName = s
		case i == 34:
			p.EmployerStreetAddress = s
		case i == 35:
			p.EmployerCity = s
		case i == 36:
			p.EmployerState = s
		case i == 37:
			p.EmployerPostalCode = s
		case i == 38:
			p.EmployerEmail = s
		case i == 39:
			p.EmployerPhone = s
		case i == 40:
			p.Occupation = s
		case i == 41:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					Ulog("CreatePeopleFromCSV: Invalid ApplicationFee value: %s\n", s)
					return
				}
				pr.ApplicationFee = x
			}
		default:
			fmt.Printf("i = %d, unknown field\n", i)
		}
	}
	//-------------------------------------------------------------------
	// OK, just insert the records and we're done
	//-------------------------------------------------------------------
	tcid, err := InsertTransactant(&tr)
	if nil != err {
		fmt.Printf("CreatePeople: error inserting Transactant = %v\n", err)
		return
	}
	tr.TCID = tcid
	t.TCID = tcid
	p.TCID = tcid
	pr.TCID = tcid

	tid, err := InsertTenant(&t)
	if nil != err {
		fmt.Printf("CreatePeople: error inserting Tenant = %v\n", err)
		return
	}
	tr.TID = tid

	pid, err := InsertPayor(&p)
	if nil != err {
		fmt.Printf("CreatePeople: error inserting Payor = %v\n", err)
		return
	}
	tr.PID = pid

	prid, err := InsertProspect(&pr)
	if nil != err {
		fmt.Printf("CreatePeople: error inserting Prospect = %v\n", err)
		return
	}
	tr.PRSPID = prid

	// now that we have all the other ids, update the Transactant record
	UpdateTransactant(&tr)

}

// LoadPeopleCSV loads a csv file with rental specialty types and processes each one
func LoadPeopleCSV(fname string) {
	t := LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreatePeopleFromCSV(t[i])
	}
}
