package rlib

// RAStates is an array that maps the Rental Agreement state number to a
// string.
//-----------------------------------------------------------------------------
var RAStates = []string{
	"Application Being Completed",
	"Pending First Approval",
	"Pending Second Approval",
	"Move-In / Execute Modification",
	"Active",
	"Terminated",
	"Notice To Move",
}

// RAActions is an array that maps the Rental Agreement Action number to a
// string.
//-----------------------------------------------------------------------------
var RAActions = []string{
	"Edit Rental Agreement Information",
	"Authorize First Approval",
	"Authorize Second Approval",
	"Complete Move In",
	"Terminate",
	"Received Notice To Move",
}

// GetStatusString returns a string representation of the Rental Agreement
// state.
//-----------------------------------------------------------------------------
func (ra *RentalAgreement) GetStatusString() string {
	return RAStates[int(ra.FLAGS&0xf)]
}
