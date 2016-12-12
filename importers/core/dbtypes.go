package core

// RentableTypeCSV is struct that is used
// to parse fields from onesite csv and create
// temporary file to import the data using rcsv
// routine in system
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
