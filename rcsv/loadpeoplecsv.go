package rcsv

import (
	"context"
	"fmt"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"strconv"
	"strings"
	"time"
)

// BUD et all are constants used by multiple programs
// for the column headings in csv files.
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
	ThirdPartySource          = iota
	DateofBirth               = iota
	EmergencyContactName      = iota
	EmergencyContactAddress   = iota
	EmergencyContactTelephone = iota
	EmergencyContactEmail     = iota
	AlternateAddress          = iota
	EligibleFutureUser        = iota
	Industry                  = iota
	SourceSLSID               = iota
	CreditLimit               = iota
	TaxpayorID                = iota
	CompanyAddress            = iota
	CompanyCity               = iota
	CompanyState              = iota
	CompanyPostalCode         = iota
	CompanyEmail              = iota
	CompanyPhone              = iota
	Occupation                = iota
	Notes                     = iota
	DesiredUsageStartDate     = iota
	RentableTypePreference    = iota
	Approver                  = iota
	DeclineReasonSLSID        = iota
	OtherPreferences          = iota
	FollowUpDate              = iota
	CSAgent                   = iota
	OutcomeSLSID              = iota
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
	{"ThirdPartySource", ThirdPartySource},
	{"DateofBirth", DateofBirth},
	{"EmergencyContactName", EmergencyContactName},
	{"EmergencyContactAddress", EmergencyContactAddress},
	{"EmergencyContactTelephone", EmergencyContactTelephone},
	{"EmergencyContactEmail", EmergencyContactEmail},
	{"AlternateAddress", AlternateAddress},
	{"EligibleFutureUser", EligibleFutureUser},
	{"Industry", Industry},
	{"SourceSLSID", SourceSLSID},
	{"CreditLimit", CreditLimit},
	{"TaxpayorID", TaxpayorID},
	{"CompanyAddress", CompanyAddress},
	{"CompanyCity", CompanyCity},
	{"CompanyState", CompanyState},
	{"CompanyPostalCode", CompanyPostalCode},
	{"CompanyEmail", CompanyEmail},
	{"CompanyPhone", CompanyPhone},
	{"Occupation", Occupation},
	{"Notes", Notes},
	{"DesiredUsageStartDate", DesiredUsageStartDate},
	{"RentableTypePreference", RentableTypePreference},
	{"Approver", Approver},
	{"DeclineReasonSLSID", DeclineReasonSLSID},
	{"OtherPreferences", OtherPreferences},
	{"FollowUpDate", FollowUpDate},
	{"CSAgent", CSAgent},
	{"OutcomeSLSID", OutcomeSLSID},
}

func rcsvCopyString(p *string, s string) error {
	*p = s
	return nil
}

// CSV file format:
//  |<------------------------------------------------------------------  TRANSACTANT ----------------------------------------------------------------------------->|  |<-------------------------------------------------------------------------------------------------------------  rlib.User  ------------------------------------------------------------------------------------------------------------------------------------------------------------------------>|<----------------------------------------------------------------------------- rlib.Payor ------------------------------------------------->|
//   0   1          2           3         4            5          6             7               8          9          10       11        12    13     14          15       16      17       18        19        20       21                 22                  23                   24          25           26                    27                       28                         29              30                31                          32        33            34           35         36            37                     38            39             40                  41             42             43          44             45    46                     47                      48        49                  50                51            52       53
// 	BUD, FirstName, MiddleName, LastName, CompanyName, IsCompany, PrimaryEmail, SecondaryEmail, WorkPhone, CellPhone, Address, Address2, City, State, PostalCode, Country, Points, VehicleMake, VehicleModel, VehicleColor, VehicleYear, LicensePlateState, LicensePlateNumber, ParkingPermitNumber, ThirdPartySource, DateofBirth, EmergencyContactName, EmergencyContactAddress, EmergencyContactTelephone, EmergencyContactEmail, AlternateAddress, EligibleFutureUser, Industry, SourceSLSID, CreditLimit, TaxpayorID, CompanyAddress, CompanyCity, CompanyState, CompanyPostalCode, CompanyEmail, CompanyPhone, Occupation, Notes,DesiredUsageStartDate, RentableTypePreference, Approver, DeclineReasonSLSID, OtherPreferences, FollowUpDate, CSAgent, OutcomeSLSID
// 	Edna,,Krabappel,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Ned,,Flanders,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Moe,,Szyslak,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Montgomery,,Burns,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Nelson,,Muntz,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Milhouse,,Van Houten,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Clancey,,Wiggum,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
// 	Homer,J,Simpson,homerj@springfield.com,,408-654-8732,,744 Evergreen Terrace,,Springfield,MO,64001,USA,5987,,Canyonero,red,,MO,BR549,,,,Marge Simpson,744 Evergreen Terrace,654=183-7946,,,,,,,,,,,,,,,,,"note: Homer is an idiot"

// CreatePeopleFromCSV reads a rental specialty type string array and creates a database record for the rental specialty type.
//
// Return Values
// int   -->  0 = everything is fine, process the next line;  1 abort the csv load
// error -->  nil if no problems
func CreatePeopleFromCSV(ctx context.Context, sa []string, lineno int) (int, error) {
	const funcname = "CreatePeopleFromCSV"

	var (
		err      error
		tr       rlib.Transactant
		t        rlib.User
		p        rlib.Payor
		pr       rlib.Prospect
		x        float64
		userNote string
	)

	var rcsvPeopleHandlers = []struct {
		ID      int
		Handler func(*string, string) error
		p       *string
	}{
		{BUD, nil, nil},
		{FirstName, rcsvCopyString, &tr.FirstName},
		{MiddleName, rcsvCopyString, &tr.MiddleName},
		{LastName, rcsvCopyString, &tr.LastName},
		{CompanyName, rcsvCopyString, &tr.CompanyName},
		{IsCompany, nil, nil},
		{PrimaryEmail, rcsvCopyString, &tr.PrimaryEmail},
		{SecondaryEmail, rcsvCopyString, &tr.SecondaryEmail},
		{WorkPhone, rcsvCopyString, &tr.WorkPhone},
		{CellPhone, nil, nil},
		{Address, rcsvCopyString, &tr.Address},
		{Address2, rcsvCopyString, &tr.Address2},
		{City, rcsvCopyString, &tr.City},
		{State, rcsvCopyString, &tr.State},
		{PostalCode, rcsvCopyString, &tr.PostalCode},
		{Country, rcsvCopyString, &tr.Country},
		{Points, nil, nil},
		{ThirdPartySource, nil, nil},
		{DateofBirth, nil, nil},
		{EmergencyContactName, rcsvCopyString, &t.EmergencyContactName},
		{EmergencyContactAddress, rcsvCopyString, &t.EmergencyContactAddress},
		{EmergencyContactTelephone, rcsvCopyString, &t.EmergencyContactTelephone},
		{EmergencyContactEmail, rcsvCopyString, &t.EmergencyContactEmail},
		{AlternateAddress, rcsvCopyString, &t.AlternateAddress},
		{EligibleFutureUser, nil, nil},
		{Industry, rcsvCopyString, &t.Industry},
		{SourceSLSID, nil, nil},
		{CreditLimit, nil, nil},
		{TaxpayorID, rcsvCopyString, &p.TaxpayorID},
		{CompanyAddress, rcsvCopyString, &pr.CompanyAddress},
		{CompanyCity, rcsvCopyString, &pr.CompanyCity},
		{CompanyState, rcsvCopyString, &pr.CompanyState},
		{CompanyPostalCode, rcsvCopyString, &pr.CompanyPostalCode},
		{CompanyEmail, rcsvCopyString, &pr.CompanyEmail},
		{CompanyPhone, rcsvCopyString, &pr.CompanyPhone},
		{Occupation, rcsvCopyString, &pr.Occupation},
		{Notes, nil, nil},
		{DesiredUsageStartDate, nil, nil},
		{RentableTypePreference, nil, nil},
		{Approver, nil, nil},
		{DeclineReasonSLSID, nil, nil},
		{OtherPreferences, nil, nil},
		{FollowUpDate, nil, nil},
		{CSAgent, nil, nil},
		{OutcomeSLSID, nil, nil},
	}

	ignoreDupPhone := false

	y, err := ValidateCSVColumnsErr(csvCols, sa, funcname, lineno)
	if y {
		return 1, err
	}
	if lineno == 1 {
		return 0, nil // we've validated the col headings, all is good, send the next line
	}

	dateform := "2006-01-02"
	pr.OtherPreferences = ""

	for i := 0; i < len(sa); i++ {
		s := strings.TrimSpace(sa[i])
		if i >= len(rcsvPeopleHandlers) {
			rlib.Ulog("Could not find handler for column: %s  col. index = %d\n", sa[i], i)
			continue
		}
		if rcsvPeopleHandlers[i].p != nil {
			rcsvPeopleHandlers[i].Handler(rcsvPeopleHandlers[i].p, s)
			continue
		}
		switch i {
		case BUD: // business
			des := strings.ToLower(strings.TrimSpace(sa[0])) // this should be BUD

			//-------------------------------------------------------------------
			// Make sure the rlib.Business is in the database
			//-------------------------------------------------------------------
			if len(des) > 0 { // make sure it's not empty
				b1, err := rlib.GetBusinessByDesignation(ctx, des) // see if we can find the biz
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d, error while getting business by designation(%s): %s", funcname, lineno, sa[0], err.Error())
				}
				if len(b1.Designation) == 0 {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Business with designation %s does not exist", funcname, lineno, sa[0])
				}
				tr.BID = b1.BID
			}
		case IsCompany:
			if len(s) > 0 {
				ic, err := rlib.YesNoToBool(s)
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - IsCompany value is invalid: %s", funcname, lineno, s)
				}
				tr.IsCompany = ic
			}
		case CellPhone:
			if len(s) > 0 && s[0] == '*' {
				s = s[1:]
				ignoreDupPhone = true
			}
			tr.CellPhone = s
		case Points:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Points value is invalid: %s", funcname, lineno, s)
				}
				t.Points = int64(i)
			}
		case ThirdPartySource:
			if len(s) > 0 {
				i, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - ThirdPartySource value is invalid: %s", funcname, lineno, s)
				}
				p.ThirdPartySource = int64(i)
			}
		case DateofBirth:
			if len(s) > 0 {
				t.DateofBirth, err = rlib.StringToDate(s)
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Bad date of birth: %s, error = %s", funcname, lineno, s, err.Error())
				}
			}
		case EligibleFutureUser:
			if len(s) > 0 {
				var err error
				t.EligibleFutureUser, err = rlib.YesNoToBool(s)
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s", funcname, lineno, err.Error())
				}
			}
		case SourceSLSID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid SourceSLSID value: %s", funcname, lineno, s)
				}
				t.SourceSLSID = y
			}
		case CreditLimit:
			if len(s) > 0 {
				if x, err = strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid Credit Limit value: %s", funcname, lineno, s)
				}
				p.CreditLimit = x
			}
		case Notes:
			if len(s) > 0 {
				userNote = s
			}
		case DesiredUsageStartDate:
			if len(s) > 0 {
				pr.DesiredUsageStartDate, err = rlib.StringToDate(s)
				if err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid DesiredUsageStartDate value: %s", funcname, lineno, s)
				}
			}
		case RentableTypePreference:
			if len(s) > 0 {
				rt, err := rlib.GetRentableTypeByStyle(ctx, s, tr.BID)
				if err != nil || rt.RTID == 0 {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid DesiredUsageStartDate value: %s", funcname, lineno, s)
				}
				pr.RentableTypePreference = rt.RTID
			}
		case Approver: // Approver ID
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid Approver UID value: %s", funcname, lineno, s)
				}
				pr.Approver = y
			}
		case DeclineReasonSLSID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid DeclineReasonSLSID value: %s", funcname, lineno, s)
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
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid CSAgent ID value: %s", funcname, lineno, s)
				}
				pr.CSAgent = y
			}
		case OutcomeSLSID:
			if len(s) > 0 {
				var y int64
				if y, err = strconv.ParseInt(strings.TrimSpace(s), 10, 64); err != nil {
					return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Invalid OutcomeSLSID value: %s", funcname, lineno, s)
				}
				pr.OutcomeSLSID = y
			}
		default:
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Unknown field, column %d", funcname, lineno, i)
		}
	}

	//-------------------------------------------------------------------
	// Make sure BID is not 0
	//-------------------------------------------------------------------
	if tr.BID == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - No Business found for BUD = %s", funcname, lineno, sa[BUD])
	}

	//-------------------------------------------------------------------
	// Make sure this person doesn't already exist...
	//-------------------------------------------------------------------
	if len(tr.PrimaryEmail) > 0 {
		t1, err := rlib.GetTransactantByPhoneOrEmail(ctx, tr.BID, tr.PrimaryEmail)
		if err != nil { // if not "no rows error" then MUST return
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error while verifying Transactant with PrimaryEmail address = %s: %s", funcname, lineno, tr.PrimaryEmail, err.Error())
		}
		if t1.TCID > 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s:: Transactant with PrimaryEmail address = %s ", funcname, lineno, DupTransactant, tr.PrimaryEmail)
		}
	}
	if len(tr.CellPhone) > 0 && !ignoreDupPhone {
		t1, err := rlib.GetTransactantByPhoneOrEmail(ctx, tr.BID, tr.CellPhone)
		if err != nil { // if not "no rows error" then MUST return
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - Error verifying Transactant with CellPhone number = %s: %s", funcname, lineno, tr.CellPhone, err.Error())
		}
		if t1.TCID > 0 {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - %s:: Transactant with CellPhone number = %s already exists", funcname, lineno, DupTransactant, tr.CellPhone)
		}
	}

	//-------------------------------------------------------------------
	// Make sure there's a name... if it's not a Company, then it needs
	// a first & last name.  If it is a company, then it needs a Company
	// name.
	//-------------------------------------------------------------------
	if (!tr.IsCompany) && len(tr.FirstName) == 0 && len(tr.LastName) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - FirstName and LastName are required for a person", funcname, lineno)
	}
	if (tr.IsCompany) && len(tr.CompanyName) == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - CompanyName is required for a company", funcname, lineno)
	}

	//-------------------------------------------------------------------
	// If there's a notelist, create it now...
	//-------------------------------------------------------------------
	if len(userNote) > 0 {
		var nl rlib.NoteList
		nl.BID = tr.BID
		nl.NLID, err = rlib.InsertNoteList(ctx, &nl)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error creating NoteList = %s", funcname, lineno, err.Error())
		}
		var n rlib.Note
		n.Comment = userNote
		n.NTID = 1 // first comment type
		n.NLID = nl.NLID
		n.BID = nl.BID
		_, err = rlib.InsertNote(ctx, &n)
		if err != nil {
			return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error creating NoteList = %s", funcname, lineno, err.Error())
		}
		tr.NLID = nl.NLID // start a notelist for this transactant
	}

	//-------------------------------------------------------------------
	// OK, just insert the records and we're done
	//-------------------------------------------------------------------
	tcid, err := rlib.InsertTransactant(ctx, &tr)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting Transactant = %v", funcname, lineno, err)
	}
	tr.TCID = tcid
	t.TCID = tcid
	t.BID = tr.BID
	p.TCID = tcid
	p.BID = tr.BID
	pr.TCID = tcid
	pr.BID = tr.BID

	if tcid == 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - after InsertTransactant tcid = %d", funcname, lineno, tcid)
	}
	// fmt.Printf("tcid = %d\n", tcid)
	// fmt.Printf("inserting user = %#v\n", t)
	_, err = rlib.InsertUser(ctx, &t)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting rlib.User = %v", funcname, lineno, err)
	}

	_, err = rlib.InsertPayor(ctx, &p)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting rlib.Payor = %v", funcname, lineno, err)
	}

	_, err = rlib.InsertProspect(ctx, &pr)
	if nil != err {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting rlib.Prospect = %v", funcname, lineno, err)
	}
	errlist := bizlogic.FinalizeTransactant(ctx, &tr)
	if len(errlist) > 0 {
		return CsvErrorSensitivity, fmt.Errorf("%s: line %d - error inserting Transactant LedgerMarker = %s", funcname, lineno, errlist[0].Message)
	}
	return 0, nil
}

// LoadPeopleCSV loads a csv file with people information
func LoadPeopleCSV(ctx context.Context, fname string) []error {
	return LoadRentRollCSV(ctx, fname, CreatePeopleFromCSV)
}
