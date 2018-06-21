package core

// RentableTypeCSV is struct that is used
// to parse fields from onesite csv and create
// temporary file to import the data using rcsv
// routine in rentabletype db
type RentableTypeCSV struct {
	BUD            string
	Style          string
	Name           string
	RentCycle      string
	Proration      string
	GSRPC          string
	ManageToBudget string
	MarketRate     string
	DtStart        string
	DtStop         string
}

// PeopleCSV is struct that is used
// to parse fields from onesite csv and create
// temporary file to import the data using rcsv
// routine in people db
type PeopleCSV struct {
	BUD                       string
	FirstName                 string
	MiddleName                string
	LastName                  string
	CompanyName               string
	IsCompany                 string
	PrimaryEmail              string
	SecondaryEmail            string
	WorkPhone                 string
	CellPhone                 string
	Address                   string
	Address2                  string
	City                      string
	State                     string
	PostalCode                string
	Country                   string
	Points                    string
	ThirdPartySource          string
	DateofBirth               string
	EmergencyContactName      string
	EmergencyContactAddress   string
	EmergencyContactTelephone string
	EmergencyContactEmail     string
	AlternateAddress          string
	EligibleFutureUser        string
	Industry                  string
	SourceSLSID               string
	CreditLimit               string
	TaxpayorID                string
	CompanyAddress            string
	CompanyCity               string
	CompanyState              string
	CompanyPostalCode         string
	CompanyEmail              string
	CompanyPhone              string
	Occupation                string
	Notes                     string
	DesiredUsageStartDate     string
	RentableTypePreference    string
	Approver                  string
	DeclineReasonSLSID        string
	OtherPreferences          string
	FollowUpDate              string
	CSAgent                   string
	OutcomeSLSID              string
}

// RentableCSV is struct that is used
// to parse fields from onesite csv and create
// temporary file to import the data using rcsv
// routine in rentable db
type RentableCSV struct {
	BUD             string
	Name            string
	AssignmentTime  string
	RUserSpec       string
	RentableStatus  string
	RentableTypeRef string
}

// RentalAgreementCSV is struct that is used
// to parse fields from onesite csv and create
// temporary file to import the data using rcsv
// routine in rental agreement db
type RentalAgreementCSV struct {
	BUD                 string
	RATemplateName      string
	AgreementStart      string
	AgreementStop       string
	PossessionStart     string
	PossessionStop      string
	RentStart           string
	RentStop            string
	RentCycleEpoch      string
	PayorSpec           string
	UserSpec            string
	UnspecifiedAdults   string
	UnspecifiedChildren string
	Renewal             string
	SpecialProvisions   string
	RentableSpec        string
	Notes               string
}

// CustomAttributeCSV is struct that is used
// to parse fields from onesite csv and create
// temporary file to import the data using rcsv
// routine in custom attribute db
type CustomAttributeCSV struct {
	BUD       string
	Name      string
	ValueType string
	Value     string
	Units     string
}
