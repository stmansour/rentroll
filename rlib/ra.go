package rlib

// RAActionApplicationBeingCompleted are all constants used for RA actions
const (
	RAActionApplicationBeingCompleted = 0
	RAActionSetToFirstApproval        = 1
	RAActionSetToSecondApproval       = 2
	RAActionSetToMoveIn               = 3
	RAActionCompleteMoveIn            = 4
	RAActionReceivedNoticeToMove      = 5
	RAActionTerminate                 = 6
	RAActionVoid                      = 7 // special case of Terminate
)

// RAStates is an array that maps the Rental Agreement state number to a
// string.
//-----------------------------------------------------------------------------
var RAStates = []string{
	"Application Being Completed",
	"Pending First Approval",
	"Pending Second Approval",
	"Move-In / Execute Modification",
	"Active",
	"Notice To Move",
	"Terminated",
}

// RAActions is an array that maps the Rental Agreement Action number to a
// string.
//-----------------------------------------------------------------------------
var RAActions = []string{
	"Application Being Completed",
	"Set To First Approval",
	"Set To Second Approval",
	"Set To Move-In",
	"Make Active / Execute",
	"Received Notice-To-Move",
	"Terminate",
}

// GetStatusString returns a string representation of the Rental Agreement
// state.
//-----------------------------------------------------------------------------
func (ra *RentalAgreement) GetStatusString() string {
	return RAStates[int(ra.FLAGS&0xf)]
}
