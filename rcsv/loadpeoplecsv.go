package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// PeopleSpecialty is the structure for attributes of a rlib.Rentable specialty

// CSV file format:
//  |<------------------------------------------------------------------  TRANSACTANT ----------------------------------------------------------------------------->|  |<-------------------------------------------------------------------------------------------------------------  rlib.User  ----------------------------------------------------------------------------------------------------------------------------------------------------------------->|<------------------------------------------------------------------------- rlib.Payor ------------------------------------------------------------------>|
//   0           1          2          3          4          5             6               7          8          9        10        11    12     13          14       15      16       17        18        19       20                 21                  22                   23          24           25                    26                       27                          28             29                30                  31        32   33           34            35                     36            37             38                  39             40             41          42              43
// 	FirstName, MiddleName, LastName, CompanyName, IsCompany, PrimaryEmail, SecondaryEmail, WorkPhone, CellPhone, Address, Address2, City, State, PostalCode, Country, Points, CarMake, CarModel, CarColor, CarYear, LicensePlateState, LicensePlateNumber, ParkingPermitNumber, AccountRep, DateofBirth, EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail, AlternateAddress, EligibleFutureUser, Industry, SID, CreditLimit, EmployerName, EmployerStreetAddress, EmployerCity, EmployerState, EmployerPostalCode, EmployerEmail, EmployerPhone, Occupation, ApplicationFee, Notes
// 	Edna,,Krabappel,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Ned,,Flanders,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Moe,,Szyslak,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Montgomery,,Burns,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Nelson,,Muntz,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Milhouse,,Van Houten,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Clancey,,Wiggum,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Homer,J,Simpson,homerj@springfield.com,,408-654-8732,,744 Evergreen Terrace,,Springfield,MO,64001,USA,5987,,Canyonero,red,,MO,BR549,,,,Marge Simpson,744 Evergreen Terrace,654=183-7946,,,,,,,,,,,,,,,,,"note: Homer is an idiot"

// CreatePeopleFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
func CreatePeopleFromCSV(sa []string, lineno int) {
	funcname := "CreatePeopleFromCSV"
	// skip the header line
	if sa[0] == "FirstName" {
		return
	}
	// fmt.Printf("line %d, sa = %#v\n", lineno, sa)
	required := 44
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		return
	}

	var (
		err      error
		tr       rlib.Transactant
		t        rlib.User
		p        rlib.Payor
		pr       rlib.Prospect
		x        float64
		userNote string
	)
	dateform := "2006-01-02"

	for i := 0; i < len(sa); i++ {
		s := strings.TrimSpace(sa[i])
		// fmt.Printf("%d. sa[%d] = \"%s\"\n", i, i, sa[i])
		switch i {
		case 0: // rlib.Transactant FirstName
			tr.FirstName = s
		case 1:
			tr.MiddleName = s
		case 2:
			tr.LastName = s
		case 3:
			tr.CompanyName = s
		case 4:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - IsCompany value is invalid: %s\n", funcname, lineno, s)
					return
				}
				if i < 0 || i > 1 {
					fmt.Printf("%s: line %d - IsCompany value is invalid: %s\n", funcname, lineno, s)
					return
				}
				tr.IsCompany = i
			}
		case 5:
			tr.PrimaryEmail = s
		case 6:
			tr.SecondaryEmail = s
		case 7:
			tr.WorkPhone = s
		case 8:
			tr.CellPhone = s
		case 9:
			tr.Address = s
		case 10:
			tr.Address2 = s
		case 11:
			tr.City = s
		case 12:
			tr.State = s
		case 13:
			tr.PostalCode = s
		case 14:
			tr.Country = s
		case 15:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - Points value is invalid: %s\n", funcname, lineno, s)
					return
				}
				t.Points = int64(i)
			}
		case 16:
			t.CarMake = s
		case 17:
			t.CarModel = s
		case 18:
			t.CarColor = s
		case 19:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - CarYear value is invalid: %s\n", funcname, lineno, s)
					return
				}
				t.CarYear = int64(i)
			}
		case 20:
			t.LicensePlateState = s
		case 21:
			t.LicensePlateNumber = s
		case 22:
			t.ParkingPermitNumber = s
		case 23:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - AccountRep value is invalid: %s\n", funcname, lineno, s)
					return
				}
				p.AccountRep = int64(i)
			}
		case 24:
			if len(s) > 0 {
				t.DateofBirth, _ = time.Parse(dateform, s)
			}
		case 25:
			t.EmergencyContactName = s
		case 26:
			t.EmergencyContactAddress = s
		case 27:
			t.EmergencyContactTelephone = s
		case 28:
			t.EmergencyEmail = s
		case 29:
			t.AlternateAddress = s
		case 30:
			if len(s) > 0 {
				var err error
				t.EligibleFutureUser, err = rlib.YesNoToInt(s)
				if err != nil {
					fmt.Printf("%s: line %d - %s\n", funcname, lineno, err.Error())
				}
			}
		case 31:
			t.Industry = s
		case 32:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid ApplicationFee value: %s\n", funcname, lineno, s)
					return
				}
				t.SID = y
			}
		case 33:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid Credit Limit value: %s\n", funcname, lineno, s)
					return
				}
				p.CreditLimit = x
			}
		case 34:
			pr.EmployerName = s
		case 35:
			pr.EmployerStreetAddress = s
		case 36:
			pr.EmployerCity = s
		case 37:
			pr.EmployerState = s
		case 38:
			pr.EmployerPostalCode = s
		case 39:
			pr.EmployerEmail = s
		case 40:
			pr.EmployerPhone = s
		case 41:
			pr.Occupation = s
		case 42:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid ApplicationFee value: %s\n", funcname, lineno, s)
					return
				}
				pr.ApplicationFee = x
			}
		case 43:
			if len(s) > 0 {
				userNote = s
			}
		default:
			fmt.Printf("i = %d, unknown field\n", i)
		}
	}
	//-------------------------------------------------------------------
	// Make sure this person doesn't already exist...
	//-------------------------------------------------------------------
	if len(tr.PrimaryEmail) > 0 {
		t1, err := rlib.GetTransactantByPhoneOrEmail(tr.PrimaryEmail)
		if err != nil && !rlib.IsSQLNoResultsError(err) {
			rlib.Ulog("%s: line %d - error retrieving rlib.Transactant by email: %v\n", funcname, lineno, err)
			return
		}
		if t1.TCID > 0 {
			rlib.Ulog("%s: line %d - rlib.Transactant with PrimaryEmail address = %s already exists\n", funcname, lineno, tr.PrimaryEmail)
			return
		}
	}
	if len(tr.CellPhone) > 0 {
		t1, err := rlib.GetTransactantByPhoneOrEmail(tr.CellPhone)
		if err != nil && !rlib.IsSQLNoResultsError(err) {
			rlib.Ulog("%s: line %d - error retrieving rlib.Transactant by phone: %v\n", funcname, lineno, err)
			return
		}
		if t1.TCID > 0 {
			rlib.Ulog("%s: line %d - rlib.Transactant with CellPhone number = %s already exists\n", funcname, lineno, tr.CellPhone)
			return
		}
	}

	//-------------------------------------------------------------------
	// If there's a notelist, create it now...
	//-------------------------------------------------------------------
	if len(userNote) > 0 {
		var nl rlib.NoteList
		nl.NLID, err = rlib.InsertNoteList(&nl)
		if err != nil {
			fmt.Printf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
			return
		}
		var n rlib.Note
		n.Comment = userNote
		n.NTID = 1 // first comment type
		n.NLID = nl.NLID
		_, err = rlib.InsertNote(&n)
		if err != nil {
			fmt.Printf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
			return
		}
		tr.NLID = nl.NLID // start a notelist for this transactant
	}

	//-------------------------------------------------------------------
	// OK, just insert the records and we're done
	//-------------------------------------------------------------------
	tcid, err := rlib.InsertTransactant(&tr)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.Transactant = %v\n", funcname, lineno, err)
		return
	}
	tr.TCID = tcid
	t.TCID = tcid
	p.TCID = tcid
	pr.TCID = tcid

	tid, err := rlib.InsertUser(&t)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.User = %v\n", funcname, lineno, err)
		return
	}
	tr.USERID = tid

	pid, err := rlib.InsertPayor(&p)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.Payor = %v\n", funcname, lineno, err)
		return
	}
	tr.PID = pid

	prid, err := rlib.InsertProspect(&pr)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.Prospect = %v\n", funcname, lineno, err)
		return
	}
	tr.PRSPID = prid

	// now that we have all the other ids, update the rlib.Transactant record
	rlib.UpdateTransactant(&tr)

}

// LoadPeopleCSV loads a csv file with rental specialty types and processes each one
func LoadPeopleCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		CreatePeopleFromCSV(t[i], i+1)
	}
}
