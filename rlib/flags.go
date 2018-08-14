package rlib

// ARApplyFundsToReceiveAccts etc all account rule related flags bit
const (
	ARApplyFundsToReceiveAccts = 0
	ARAutoPopulateToNewRA      = 1
	ARRAIDRequired             = 2
	ARSubARIDsOnly             = 3
	ARIsRentASM                = 4
	ARIsSecDepASM              = 5
	ARIsNonRecurCharge         = 6
	ARPETIDReq                 = 7
	ARVIDReq                   = 8
)

// ARFLAGS account rules FLAGS
var ARFLAGS = Str2Int64Map{
	"ApplyFundsToReceiveAccts": ARApplyFundsToReceiveAccts,
	"AutoPopulateToNewRA":      ARAutoPopulateToNewRA,
	"RAIDRequired":             ARRAIDRequired,
	"SubARIDsOnly":             ARSubARIDsOnly,
	"IsRentASM":                ARIsRentASM,
	"IsSecDepASM":              ARIsSecDepASM,
	"IsNonRecurCharge":         ARIsNonRecurCharge,
	"PETIDReq":                 ARPETIDReq,
	"VIDReq":                   ARVIDReq,
}
