package rcsv

import (
	"fmt"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// CSVColumn defines a column of the CSV file
type CSVColumn struct {
	Name  string
	Index int
}

// ValidateCSVColumns verifies the column titles with the supplied, expected titles.
// Returns:
//   0 = everything is OK
//   1 = at least 1 column is wrong, error message already printed
func ValidateCSVColumns(csvCols []CSVColumn, sa []string, funcname string, lineno int) int {
	required := len(csvCols)
	if len(sa) < required {
		fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
		l := len(sa)
		for i := 0; i < len(csvCols); i++ {
			if i < l {
				s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
				if s != strings.ToLower(csvCols[i].Name) {
					fmt.Printf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
					return 1
				}
			}
		}
		return 1
	}

	if lineno == 1 {
		for i := 0; i < len(csvCols); i++ {
			s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
			if s != strings.ToLower(csvCols[i].Name) {
				fmt.Printf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
				return 1
			}
		}
	}
	return 0
}

// CSV file format:
//  |<------------------------------------------------------------------  TRANSACTANT ----------------------------------------------------------------------------->|  |<-------------------------------------------------------------------------------------------------------------  rlib.User  ------------------------------------------------------------------------------------------------------------------------------------------------------------------------>|<----------------------------------------------------------------------------- rlib.Payor ------------------------------------------------->|
//   0   1          2           3         4            5          6             7               8          9          10       11        12    13     14          15       16      17       18        19        20       21                 22                  23                   24          25           26                    27                       28                         29              30                31                          32        33            34           35         36            37                     38            39             40                  41             42             43          44             45    46                     47                      48        49                  50                51            52       53            54               55
// 	BUD, FirstName, MiddleName, LastName, CompanyName, IsCompany, PrimaryEmail, SecondaryEmail, WorkPhone, CellPhone, Address, Address2, City, State, PostalCode, Country, Points, CarMake, CarModel, CarColor, CarYear, LicensePlateState, LicensePlateNumber, ParkingPermitNumber, AccountRep, DateofBirth, EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyEmail, AlternateAddress, EligibleFutureUser, Industry, SourceSLSID, CreditLimit, TaxpayorID, EmployerName, EmployerStreetAddress, EmployerCity, EmployerState, EmployerPostalCode, EmployerEmail, EmployerPhone, Occupation, ApplicationFee,Notes,DesiredUsageStartDate, RentableTypePreference, Approver, DeclineReasonSLSID, OtherPreferences, FollowUpDate, CSAgent, OutcomeSLSID, FloatingDeposit, RAID
// 	Edna,,Krabappel,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Ned,,Flanders,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Moe,,Szyslak,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Montgomery,,Burns,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Nelson,,Muntz,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Milhouse,,Van Houten,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Clancey,,Wiggum,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Homer,J,Simpson,homerj@springfield.com,,408-654-8732,,744 Evergreen Terrace,,Springfield,MO,64001,USA,5987,,Canyonero,red,,MO,BR549,,,,Marge Simpson,744 Evergreen Terrace,654=183-7946,,,,,,,,,,,,,,,,,"note: Homer is an idiot"

// CreatePeopleFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
// If the return value is not 0, abort the csv load
func CreatePeopleFromCSV(sa []string, lineno int) int {
	funcname := "CreatePeopleFromCSV"

	var (
		err      error
		tr       rlib.Transactant
		t        rlib.User
		p        rlib.Payor
		pr       rlib.Prospect
		x        float64
		userNote string
	)

	const (
		BUD                       = 0
		FirstName                 = iota
		MiddleName                = iota
		LastName                  = iota
		CompanyName               = iota
		IsCompany                 = iota
		PrimaryEmail              = iota
		SecondaryEmail            = iota
		WorkPhone                 = iota
		CellPhone                 = iota
		Address                   = iota
		Address2                  = iota
		City                      = iota
		State                     = iota
		PostalCode                = iota
		Country                   = iota
		Points                    = iota
		CarMake                   = iota
		CarModel                  = iota
		CarColor                  = iota
		CarYear                   = iota
		LicensePlateState         = iota
		LicensePlateNumber        = iota
		ParkingPermitNumber       = iota
		AccountRep                = iota
		DateofBirth               = iota
		EmergencyContactName      = iota
		EmergencyContactAddress   = iota
		EmergencyContactTelephone = iota
		EmergencyEmail            = iota
		AlternateAddress          = iota
		EligibleFutureUser        = iota
		Industry                  = iota
		SourceSLSID               = iota
		CreditLimit               = iota
		TaxpayorID                = iota
		EmployerName              = iota
		EmployerStreetAddress     = iota
		EmployerCity              = iota
		EmployerState             = iota
		EmployerPostalCode        = iota
		EmployerEmail             = iota
		EmployerPhone             = iota
		Occupation                = iota
		ApplicationFee            = iota
		Notes                     = iota
		DesiredUsageStartDate     = iota
		RentableTypePreference    = iota
		Approver                  = iota
		DeclineReasonSLSID        = iota
		OtherPreferences          = iota
		FollowUpDate              = iota
		CSAgent                   = iota
		OutcomeSLSID              = iota
		FloatingDeposit           = iota
		RAID                      = iota
	)

	// csvCols is an array that defines all the columns that should be in this csv file
	var csvCols = []CSVColumn{
		{"BUD", BUD},
		{"FirstName", FirstName},
		{"MiddleName", MiddleName},
		{"LastName", LastName},
		{"CompanyName", CompanyName},
		{"IsCompany", IsCompany},
		{"PrimaryEmail", PrimaryEmail},
		{"SecondaryEmail", SecondaryEmail},
		{"WorkPhone", WorkPhone},
		{"CellPhone", CellPhone},
		{"Address", Address},
		{"Address2", Address2},
		{"City", City},
		{"State", State},
		{"PostalCode", PostalCode},
		{"Country", Country},
		{"Points", Points},
		{"CarMake", CarMake},
		{"CarModel", CarModel},
		{"CarColor", CarColor},
		{"CarYear", CarYear},
		{"LicensePlateState", LicensePlateState},
		{"LicensePlateNumber", LicensePlateNumber},
		{"ParkingPermitNumber", ParkingPermitNumber},
		{"AccountRep", AccountRep},
		{"DateofBirth", DateofBirth},
		{"EmergencyContactName", EmergencyContactName},
		{"EmergencyContactAddress", EmergencyContactAddress},
		{"EmergencyContactTelephone", EmergencyContactTelephone},
		{"EmergencyEmail", EmergencyEmail},
		{"AlternateAddress", AlternateAddress},
		{"EligibleFutureUser", EligibleFutureUser},
		{"Industry", Industry},
		{"SourceSLSID", SourceSLSID},
		{"CreditLimit", CreditLimit},
		{"TaxpayorID", TaxpayorID},
		{"EmployerName", EmployerName},
		{"EmployerStreetAddress", EmployerStreetAddress},
		{"EmployerCity", EmployerCity},
		{"EmployerState", EmployerState},
		{"EmployerPostalCode", EmployerPostalCode},
		{"EmployerEmail", EmployerEmail},
		{"EmployerPhone", EmployerPhone},
		{"Occupation", Occupation},
		{"ApplicationFee", ApplicationFee},
		{"Notes", Notes},
		{"DesiredUsageStartDate", DesiredUsageStartDate},
		{"RentableTypePreference", RentableTypePreference},
		{"Approver", Approver},
		{"DeclineReasonSLSID", DeclineReasonSLSID},
		{"OtherPreferences", OtherPreferences},
		{"FollowUpDate", FollowUpDate},
		{"CSAgent", CSAgent},
		{"OutcomeSLSID", OutcomeSLSID},
		{"FloatingDeposit", FloatingDeposit},
		{"RAID", RAID},
	}

	if ValidateCSVColumns(csvCols, sa, funcname, lineno) > 0 {
		return 1
	}
	if lineno == 1 {
		return 0
	}

	// //------------------------------------
	// // Validate the columns...
	// //------------------------------------
	// required := len(csvCols)
	// if len(sa) < required {
	// 	fmt.Printf("%s: line %d - found %d values, there must be at least %d\n", funcname, lineno, len(sa), required)
	// 	l := len(sa)
	// 	for i := 0; i < len(csvCols); i++ {
	// 		if i < l {
	// 			s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
	// 			if s != strings.ToLower(csvCols[i].Name) {
	// 				fmt.Printf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
	// 				return 1
	// 			}
	// 		}
	// 	}
	// 	return 1
	// }

	// if lineno == 1 {
	// 	for i := 0; i < len(csvCols); i++ {
	// 		s := rlib.Stripchars(strings.ToLower(strings.TrimSpace(sa[i])), " ")
	// 		if s != strings.ToLower(csvCols[i].Name) {
	// 			fmt.Printf("%s: line %d - Error at column heading %d, expected %s, found %s\n", funcname, lineno, i, csvCols[i].Name, sa[i])
	// 			return 1
	// 		}
	// 	}
	// 	return 0
	// }

	dateform := "2006-01-02"
	pr.OtherPreferences = ""

	for i := 0; i < len(sa); i++ {
		s := strings.TrimSpace(sa[i])
		switch i {
		case BUD: // business
			des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD

			//-------------------------------------------------------------------
			// Make sure the rlib.Business is in the database
			//-------------------------------------------------------------------
			if len(des) > 0 { // make sure it's not empty
				b1 := rlib.GetBusinessByDesignation(des) // see if we can find the biz
				if len(b1.Designation) == 0 {
					rlib.Ulog("%s: line %d, Business with designation %s does net exist\n", funcname, lineno, sa[0])
					return 2
				}
				tr.BID = b1.BID
			}
		case FirstName:
			tr.FirstName = s
		case MiddleName:
			tr.MiddleName = s
		case LastName:
			tr.LastName = s
		case CompanyName:
			tr.CompanyName = s
		case IsCompany:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - IsCompany value is invalid: %s\n", funcname, lineno, s)
					return 2
				}
				if i < 0 || i > 1 {
					fmt.Printf("%s: line %d - IsCompany value is invalid: %s\n", funcname, lineno, s)
					return 2
				}
				tr.IsCompany = i
			}
		case PrimaryEmail:
			tr.PrimaryEmail = s
		case SecondaryEmail:
			tr.SecondaryEmail = s
		case WorkPhone:
			tr.WorkPhone = s
		case CellPhone:
			tr.CellPhone = s
		case Address:
			tr.Address = s
		case Address2:
			tr.Address2 = s
		case City:
			tr.City = s
		case State:
			tr.State = s
		case PostalCode:
			tr.PostalCode = s
		case Country:
			tr.Country = s
		case Points:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - Points value is invalid: %s\n", funcname, lineno, s)
					return 2
				}
				t.Points = int64(i)
			}
		case CarMake:
			t.CarMake = s
		case CarModel:
			t.CarModel = s
		case CarColor:
			t.CarColor = s
		case CarYear:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - CarYear value is invalid: %s\n", funcname, lineno, s)
					return 2
				}
				t.CarYear = int64(i)
			}
		case LicensePlateState:
			t.LicensePlateState = s
		case LicensePlateNumber:
			t.LicensePlateNumber = s
		case ParkingPermitNumber:
			t.ParkingPermitNumber = s
		case AccountRep:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					fmt.Printf("%s: line %d - AccountRep value is invalid: %s\n", funcname, lineno, s)
					return 2
				}
				p.AccountRep = int64(i)
			}
		case DateofBirth:
			if len(s) > 0 {
				t.DateofBirth, _ = time.Parse(dateform, s)
			}
		case EmergencyContactName:
			t.EmergencyContactName = s
		case EmergencyContactAddress:
			t.EmergencyContactAddress = s
		case EmergencyContactTelephone:
			t.EmergencyContactTelephone = s
		case EmergencyEmail:
			t.EmergencyEmail = s
		case AlternateAddress:
			t.AlternateAddress = s
		case EligibleFutureUser:
			if len(s) > 0 {
				var err error
				t.EligibleFutureUser, err = rlib.YesNoToInt(s)
				if err != nil {
					fmt.Printf("%s: line %d - %s\n", funcname, lineno, err.Error())
				}
			}
		case Industry:
			t.Industry = s
		case SourceSLSID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid SourceSLSID value: %s\n", funcname, lineno, s)
					return 2
				}
				t.SourceSLSID = y
			}
		case CreditLimit:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid Credit Limit value: %s\n", funcname, lineno, s)
					return 2
				}
				p.CreditLimit = x
			}
		case TaxpayorID:
			p.TaxpayorID = s
		case EmployerName:
			pr.EmployerName = s
		case EmployerStreetAddress:
			pr.EmployerStreetAddress = s
		case EmployerCity:
			pr.EmployerCity = s
		case EmployerState:
			pr.EmployerState = s
		case EmployerPostalCode:
			pr.EmployerPostalCode = s
		case EmployerEmail:
			pr.EmployerEmail = s
		case EmployerPhone:
			pr.EmployerPhone = s
		case Occupation:
			pr.Occupation = s
		case ApplicationFee:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid ApplicationFee value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.ApplicationFee = x
			}
		case Notes:
			if len(s) > 0 {
				userNote = s
			}
		case DesiredUsageStartDate:
			if len(s) > 0 {
				pr.DesiredUsageStartDate, err = rlib.StringToDate(s)
				if err != nil {
					rlib.Ulog("%s: line %d - Invalid DesiredUsageStartDate value: %s\n", funcname, lineno, s)
					return 2
				}
			}
		case RentableTypePreference:
			if len(s) > 0 {
				rt, err := rlib.GetRentableTypeByStyle(s, tr.BID)
				if err != nil || rt.RTID == 0 {
					rlib.Ulog("%s: line %d - Invalid DesiredUsageStartDate value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.RentableTypePreference = rt.RTID
			}
		case Approver: // Approver ID
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid Approver UID value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.Approver = y
			}
		case DeclineReasonSLSID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid DeclineReasonSLSID value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.DeclineReasonSLSID = y
			}
		case OtherPreferences:
			if len(s) > 0 {
				pr.OtherPreferences = s
			}
		case FollowUpDate:
			if len(s) > 0 {
				pr.FollowUpDate, _ = time.Parse(dateform, s)
			}
		case CSAgent:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid CSAgent ID value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.CSAgent = y
			}
		case OutcomeSLSID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid OutcomeSLSID value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.OutcomeSLSID = y
			}

		case FloatingDeposit:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid FloatingDeposit value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.FloatingDeposit = x
			}
		case RAID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					rlib.Ulog("%s: line %d - Invalid RAID value: %s\n", funcname, lineno, s)
					return 2
				}
				pr.RAID = y
			}
		default:
			fmt.Printf("i = %d, unknown field\n", i)
		}
	}
	//-------------------------------------------------------------------
	// Make sure this person doesn't already exist...
	//-------------------------------------------------------------------
	if len(tr.PrimaryEmail) > 0 {
		t1 := rlib.GetTransactantByPhoneOrEmail(tr.PrimaryEmail)
		if t1.TCID > 0 {
			rlib.Ulog("%s: line %d - rlib.Transactant with PrimaryEmail address = %s already exists\n", funcname, lineno, tr.PrimaryEmail)
			return 2
		}
	}
	if len(tr.CellPhone) > 0 {
		t1 := rlib.GetTransactantByPhoneOrEmail(tr.CellPhone)
		if t1.TCID > 0 {
			rlib.Ulog("%s: line %d - rlib.Transactant with CellPhone number = %s already exists\n", funcname, lineno, tr.CellPhone)
			return 2
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
			return 2
		}
		var n rlib.Note
		n.Comment = userNote
		n.NTID = 1 // first comment type
		n.NLID = nl.NLID
		_, err = rlib.InsertNote(&n)
		if err != nil {
			fmt.Printf("%s: line %d - error creating NoteList = %s\n", funcname, lineno, err.Error())
			return 2
		}
		tr.NLID = nl.NLID // start a notelist for this transactant
	}

	//-------------------------------------------------------------------
	// OK, just insert the records and we're done
	//-------------------------------------------------------------------
	tcid, err := rlib.InsertTransactant(&tr)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting Transactant = %v\n", funcname, lineno, err)
		return 3
	}
	tr.TCID = tcid
	t.TCID = tcid
	p.TCID = tcid
	pr.TCID = tcid

	if tcid == 0 {
		fmt.Printf("%s: line %d - after InsertTransactant tcid = %d\n", funcname, lineno, tcid)
		return 4
	}
	_, err = rlib.InsertUser(&t)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.User = %v\n", funcname, lineno, err)
		return 5
	}

	_, err = rlib.InsertPayor(&p)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.Payor = %v\n", funcname, lineno, err)
		return 6
	}

	_, err = rlib.InsertProspect(&pr)
	if nil != err {
		fmt.Printf("%s: line %d - error inserting rlib.Prospect = %v\n", funcname, lineno, err)
		return 7
	}
	return 0
}

// LoadPeopleCSV loads a csv file with rental specialty types and processes each one
func LoadPeopleCSV(fname string) {
	t := rlib.LoadCSV(fname)
	for i := 0; i < len(t); i++ {
		if 0 != CreatePeopleFromCSV(t[i], i+1) {
			return
		}
	}
}
