package rcsv

import (
	"rentroll/rlib"
	"time"
)

/*// CSVBusiness et. al., are indeces of the functions that load a csv file with
// the type of information described in the constant's name.
const (
	CSVAssessments              = 0
	CSVReceipts                 = iota
	CSVBusiness                 = iota
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
	CSVDeposit                  = iota
	CSVNoteTypes                = iota
	CSVInvoices                 = iota
)

// CSVLoader is a struct to define a csv loading function
type CSVLoader struct {
	Name   string
	Index  int // which loader
	Loader func(context.Context, string) []error
}

// CSVLoaders is an array of functions that load CSV files that are indexed
// by the associated Index value
var CSVLoaders = []CSVLoader{
	{Name: "Assessments", Index: CSVAssessments, Loader: LoadAssessmentsCSV},
	{Name: "Receipts", Index: CSVReceipts, Loader: LoadReceiptsCSV},
	// {Name: "Business", Index: CSVBusiness, Loader: LoadBusinessCSV},
	// {Name: "StringTables", Index: CSVStringTables, Loader: LoadStringTablesCSV},
	// {Name: "PaymentTypes", Index: CSVPaymentTypes, Loader: LoadPaymentTypesCSV},
	// {Name: "DepositMethods", Index: CSVDepositMethods, Loader: LoadDepositMethodsCSV},
	// {Name: "Sources", Index: CSVSources, Loader: LoadSourcesCSV},
	// {Name: "RentableTypes", Index: CSVRentableTypes, Loader: LoadRentableTypesCSV},
	// {Name: "CustomAttributes", Index: CSVCustomAttributes, Loader: LoadCustomAttributesCSV},
	// {Name: "Depository", Index: CSVDepository, Loader: LoadDepositoryCSV},
	// {Name: "RentalSpecialties", Index: CSVRentalSpecialties, Loader: LoadRentalSpecialtiesCSV},
	// {Name: "Building", Index: CSVBuilding, Loader: LoadBuildingCSV},
	// {Name: "People", Index: CSVPeople, Loader: LoadPeopleCSV},
	// {Name: "Rentables", Index: CSVRentables, Loader: LoadRentablesCSV},
	// {Name: "RentableSpecialtyRefs", Index: CSVRentableSpecialtyRefs, Loader: LoadRentableSpecialtyRefsCSV},
	// {Name: "RentalAgreementTemplates", Index: CSVRentalAgreementTemplates, Loader: LoadRentalAgreementTemplatesCSV},
	// {Name: "RentalAgreement", Index: CSVRentalAgreement, Loader: LoadRentalAgreementCSV},
	// {Name: "Pets", Index: CSVPets, Loader: LoadPetsCSV},
	// {Name: "ChartOfAccounts", Index: CSVChartOfAccounts, Loader: LoadChartOfAccountsCSV},
	// {Name: "RatePlans", Index: CSVRatePlans, Loader: LoadRatePlansCSV},
	// {Name: "RatePlanRefs", Index: CSVRatePlanRefs, Loader: LoadRatePlanRefsCSV},
	// {Name: "RatePlanRefRTRates", Index: CSVRatePlanRefRTRates, Loader: LoadRatePlanRefRTRatesCSV},
	// {Name: "RatePlanRefSPRates", Index: CSVRatePlanRefSPRates, Loader: LoadRatePlanRefSPRatesCSV},
	// {Name: "Deposit", Index: CSVDeposit, Loader: LoadDepositCSV},
	// {Name: "CustomAttributeRefs", Index: CSVCustomAttributeRefs, Loader: LoadCustomAttributeRefsCSV},
	// {Name: "NoteTypes", Index: CSVNoteTypes, Loader: LoadNoteTypesCSV},
	// {Name: "Invoices", Index: CSVInvoices, Loader: LoadInvoicesCSV},
	// {Name: "Building", Index: CSVBuilding, Loader: LoadBuildingCSV},
}*/

// Rcsv contains the shared data used by the RCS loaders
var Rcsv struct {
	DtStart time.Time
	DtStop  time.Time
	Xbiz    *rlib.XBusiness
}

// InitRCSV initializes the shared data used by they RCS loaders.
func InitRCSV(d1, d2 *time.Time, xbiz *rlib.XBusiness) {
	Rcsv.DtStart = *d1
	Rcsv.DtStop = *d2
	Rcsv.Xbiz = xbiz

	// if dateMode is on then change the stopDate value for search op
	rlib.HandleFrontEndDates(Rcsv.Xbiz.P.BID, &Rcsv.DtStart, &Rcsv.DtStop)
}

/*// DispatchCSV is the generic CSV loader call. It will call a csv loader with the supplied
// file name based on the supplied index.
func DispatchCSV(ctx context.Context, index int, fname string) string {
	for i := 0; i < len(CSVLoaders); i++ {
		if CSVLoaders[i].Index == index {
			m := CSVLoaders[i].Loader(ctx, fname)
			return ErrlistToString(&m)
		}
	}
	return fmt.Sprintf("CSV Loader %d not found", index)
}*/
