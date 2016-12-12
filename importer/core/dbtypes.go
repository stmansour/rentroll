package core

// RentableTypeJSON is rentroll type struct
// contains field which maps to onesite field
// defined in mapper.json
type RentableTypeJSON struct {
	RTID           string // unique identifier for this RentableType
	BID            string // the business unit to which this RentableType belongs
	Style          string // a short name
	Name           string // longer name
	RentCycle      string // frequency at which rent accrues, 0 = not set or n/a, 1 = secondly, 2=minutely, 3=hourly, 4=daily, 5=weekly, 6=monthly...
	Proration      string // frequency for prorating rent if the full rentcycle is not used
	GSRPC          string // Time increments in which GSR is calculated to account for rate changes
	ManageToBudget string // 0=no, 1 = yes
	MR             string // array of time sensitive market rates
	CA             string // index by Name of attribute, associated custom attributes
	MRCurrent      string // the current market rate (historical values are in MR)
	LastModTime    string
	LastModBy      string
}
