package rcsv

// CSVBusiness et. al., are indeces of the functions that load a csv file with
// the type of information described in the constant's name.
const (
	CSVBusiness                 = 0
	CSVChartOfAccounts          = iota
	CSVStringTables             = iota
	CSVPaymentTypes             = iota
	CSVDepositMethods           = iota
	CSVSources                  = iota
	CSVRentableTypes            = iota
	CSVRentalSpecialties        = iota
	CSVBuilding                 = iota
	CSVDepository               = iota
	CSVPeople                   = iota
	CSVRentables                = iota
	CSVRentableSpecialtyRefs    = iota
	CSVRentalAgreementTemplates = iota
	CSVRentalAgreement          = iota
	CSVPets                     = iota
	CSVCustomAttributes         = iota
	CSVCustomAttributeRefs      = iota
	CSVRatePlans                = iota
	CSVRatePlanRefs             = iota
	CSVRatePlanRefRTRates       = iota
	CSVRatePlanRefSPRates       = iota
	CSVAssessments              = iota
	CSVReceipts                 = iota
	CSVDeposit                  = iota
	CSVNoteTypes                = iota
	CSVInvoices                 = iota
)

// CSVLoader is a struct to define a csv loading function
type CSVLoader struct {
	Name   string
	Index  int // which loader
	Loader func(string)
}

// CSVLoaders is an array of functions that load CSV files that are indexed
// by the associated Index value.
var CSVLoaders = []CSVLoader{
	{Name: "Business", Index: CSVBusiness, Loader: LoadBusinessCSV},
	{Name: "StringTables", Index: CSVStringTables, Loader: LoadStringTablesCSV},
	{Name: "PaymentTypes", Index: CSVPaymentTypes, Loader: LoadPaymentTypesCSV},
	{Name: "DepositMethods", Index: CSVDepositMethods, Loader: LoadDepositMethodsCSV},
	{Name: "Sources", Index: CSVSources, Loader: LoadSourcesCSV},
	{Name: "RentableTypes", Index: CSVRentableTypes, Loader: LoadRentableTypesCSV},
	{Name: "CustomAttributes", Index: CSVCustomAttributes, Loader: LoadCustomAttributesCSV},
	{Name: "Depository", Index: CSVDepository, Loader: LoadDepositoryCSV},
	{Name: "RentalSpecialties", Index: CSVRentalSpecialties, Loader: LoadRentalSpecialtiesCSV},
	{Name: "Building", Index: CSVBuilding, Loader: LoadBuildingCSV},
	{Name: "People", Index: CSVPeople, Loader: LoadPeopleCSV},
	{Name: "Rentables", Index: CSVRentables, Loader: LoadRentablesCSV},
	{Name: "RentableSpecialtyRefs", Index: CSVRentableSpecialtyRefs, Loader: LoadRentableSpecialtyRefsCSV},
	{Name: "RentalAgreementTemplates", Index: CSVRentalAgreementTemplates, Loader: LoadRentalAgreementTemplatesCSV},
	{Name: "RentalAgreement", Index: CSVRentalAgreement, Loader: LoadRentalAgreementCSV},
	{Name: "Pets", Index: CSVPets, Loader: LoadPetsCSV},
	{Name: "ChartOfAccounts", Index: CSVChartOfAccounts, Loader: LoadChartOfAccountsCSV},
	{Name: "RatePlans", Index: CSVRatePlans, Loader: LoadRatePlansCSV},
	{Name: "RatePlanRefs", Index: CSVRatePlanRefs, Loader: LoadRatePlanRefsCSV},
	{Name: "RatePlanRefRTRates", Index: CSVRatePlanRefRTRates, Loader: LoadRatePlanRefRTRatesCSV},
	{Name: "RatePlanRefSPRates", Index: CSVRatePlanRefSPRates, Loader: LoadRatePlanRefSPRatesCSV},
	{Name: "Assessments", Index: CSVAssessments, Loader: LoadAssessmentsCSV},
	{Name: "Receipts", Index: CSVReceipts, Loader: LoadReceiptsCSV},
	{Name: "Deposit", Index: CSVDeposit, Loader: LoadDepositCSV},
	{Name: "CustomAttributeRefs", Index: CSVCustomAttributeRefs, Loader: LoadCustomAttributeRefsCSV},
	{Name: "NoteTypes", Index: CSVNoteTypes, Loader: LoadNoteTypesCSV},
	{Name: "Invoices", Index: CSVInvoices, Loader: LoadInvoicesCSV},
}

// LoadCSV is the generic CSV loader call. It will call a csv loader with the supplied
// file name based on the supplied index.
func LoadCSV(index int, fname string) {
	for i := 0; i < len(CSVLoaders); i++ {
		if CSVLoaders[i].Index == index {
			CSVLoaders[i].Loader(fname)
			return
		}
	}
}
