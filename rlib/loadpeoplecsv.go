package rlib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// PeopleSpecialty is the structure for attributes of a rentable specialty

// CSV file format:
//  |<------------------------------------------------------------------  TRANSACTANT ----------------------------------------------------------------------------->|  |<-------------------------------------------------------------------------------------------------------------  Renter  ----------------------------------------------------------------------------------------------------------------------------------------------------------------->|<------------------------------------------------------------------------- Payor ------------------------------------------------------>|  -- prospect --
//   0           1          2          3          4          5             6               7          8          9        10        11    12     13          14       15      16       17        18        19       20                 21                  22                   23          24           25                    26                       27                          28             29                30                          31        32      33                   34           35               36            37            38             39                  40              41          42
// 	FirstName, MiddleName, LastName, CompanyName, IsCompany, PrimaryEmail, SecondaryEmail, WorkPhone, CellPhone, Address, Address2, City, State, PostalCode, Country, Points, CarMake, CarModel, CarColor, CarYear, LicensePlateState, LicensePlateNumber, ParkingPermitNumber, AccountRep, DateofBirth, EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail, AlternateAddress, EligibleFutureRenter, Industry, Source, CreditLimit, EmployerName, EmployerStreetAddress, EmployerCity, EmployerState, EmployerPostalCode, EmployerEmail, EmployerPhone, Occupation, ApplicationFee
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
	var t Renter
	var p Payor
	var pr Prospect
	var x float64
	dateform := "2006-01-02"

	for i := 0; i < len(sa); i++ {
		s := strings.TrimSpace(sa[i])
		// fmt.Printf("%d. sa[%d] = \"%s\"\n", i, i, sa[i])
		switch {
		case i == 0: // transactant FirstName
			tr.FirstName = s
		case i == 1:
			tr.MiddleName = s
		case i == 2:
			tr.LastName = s
		case i == 3:
			tr.CompanyName = s
		case i == 4:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: IsCompany value is invalid: %s\n", s)
					return
				}
				if i < 0 || i > 1 {
					fmt.Printf("CreatePeopleFromCSV: IsCompany value is invalid: %s\n", s)
					return
				}
				tr.IsCompany = i
			}
		case i == 5:
			tr.PrimaryEmail = s
		case i == 6:
			tr.SecondaryEmail = s
		case i == 7:
			tr.WorkPhone = s
		case i == 8:
			tr.CellPhone = s
		case i == 9:
			tr.Address = s
		case i == 10:
			tr.Address2 = s
		case i == 11:
			tr.City = s
		case i == 12:
			tr.State = s
		case i == 13:
			tr.PostalCode = s
		case i == 14:
			tr.Country = s
		case i == 15:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: Points value is invalid: %s\n", s)
					return
				}
				t.Points = int64(i)
			}
		case i == 16:
			t.CarMake = s
		case i == 17:
			t.CarModel = s
		case i == 18:
			t.CarColor = s
		case i == 19:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: CarYear value is invalid: %s\n", s)
					return
				}
				t.CarYear = int64(i)
			}
		case i == 20:
			t.LicensePlateState = s
		case i == 21:
			t.LicensePlateNumber = s
		case i == 22:
			t.ParkingPermitNumber = s
		case i == 23:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: AccountRep value is invalid: %s\n", s)
					return
				}
				p.AccountRep = int64(i)
			}
		case i == 24:
			if len(s) > 0 {
				t.DateofBirth, _ = time.Parse(dateform, s)
			}
		case i == 25:
			t.EmergencyContactName = s
		case i == 26:
			t.EmergencyContactAddress = s
		case i == 27:
			t.EmergencyContactTelephone = s
		case i == 28:
			t.EmergencyEmail = s
		case i == 29:
			t.AlternateAddress = s
		case i == 30:
			if len(s) > 0 {
				var err error
				t.EligibleFutureRenter, err = yesnoToInt(s)
				if err != nil {
					fmt.Printf("CreatePeopleFromCSV: %s\n", err.Error())
				}
			}
		case i == 31:
			t.Industry = s
		case i == 32:
			t.Source = s
		case i == 33:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					Ulog("CreatePeopleFromCSV: Invalid Credit Limit value: %s\n", s)
					return
				}
				p.CreditLimit = x
			}
		case i == 34:
			pr.EmployerName = s
		case i == 35:
			pr.EmployerStreetAddress = s
		case i == 36:
			pr.EmployerCity = s
		case i == 37:
			pr.EmployerState = s
		case i == 38:
			pr.EmployerPostalCode = s
		case i == 39:
			pr.EmployerEmail = s
		case i == 40:
			pr.EmployerPhone = s
		case i == 41:
			pr.Occupation = s
		case i == 42:
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
	// Make sure this person doesn't already exist...
	//-------------------------------------------------------------------
	if len(tr.PrimaryEmail) > 0 {
		t1, err := GetTransactantByPhoneOrEmail(tr.PrimaryEmail)
		if err != nil && !IsSQLNoResultsError(err) {
			Ulog("CreatePeopleFromCSV: error retrieving transactant by email: %v\n", err)
			return
		}
		if t1.TCID > 0 {
			Ulog("CreatePeopleFromCSV: Transactant with PrimaryEmail address = %s already exists\n", tr.PrimaryEmail)
			return
		}
	}
	if len(tr.CellPhone) > 0 {
		t1, err := GetTransactantByPhoneOrEmail(tr.CellPhone)
		if err != nil && !IsSQLNoResultsError(err) {
			Ulog("CreatePeopleFromCSV: error retrieving transactant by phone: %v\n", err)
			return
		}
		if t1.TCID > 0 {
			Ulog("CreatePeopleFromCSV: Transactant with CellPhone number = %s already exists\n", tr.CellPhone)
			return
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

	tid, err := InsertRenter(&t)
	if nil != err {
		fmt.Printf("CreatePeople: error inserting Renter = %v\n", err)
		return
	}
	tr.RENTERID = tid

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
