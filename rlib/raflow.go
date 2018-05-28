package rlib

// RAFlow etc.. all are list of all flows exist in the system
const (
	RAFlow string = "RA"
)

// RAFlowPartType is type of rental agreement flow part
type RAFlowPartType int

// DatesRAFlowPart etc. all are constants for rental agreement flow part
const (
	DatesRAFlowPart RAFlowPartType = 1 + iota // must start from 1
	PeopleRAFlowPart
	PetsRAFlowPart
	VehiclesRAFlowPart
	RentablesRAFlowPart
	ParentChildRAFlowPart
	TieRAFlowPart
)

// IsValid checks the validity of RAFlowPartType raftp
func (raftp RAFlowPartType) IsValid() bool {
	if raftp < DatesRAFlowPart || raftp > RentablesRAFlowPart {
		return false
	}

	return true
}

// String representation of RAFlowPartType
func (raftp RAFlowPartType) String() string {
	names := [...]string{
		"Agreement Dates",
		"People with background info",
		"Pets",
		"Vehicles",
		"Rentables with fees",
		"Parent/Child",
		"Tie",
	}

	// if not valid then return unknown
	if !(raftp.IsValid()) {
		return "Unknown RA FlowPart"
	}

	return names[raftp-1]
}

// RAFlowPartsMap parts of a rental agreement flow
var RAFlowPartsMap = Str2Int64Map{
	"dates":       int64(DatesRAFlowPart),
	"people":      int64(PeopleRAFlowPart),
	"pets":        int64(PetsRAFlowPart),
	"vehicles":    int64(VehiclesRAFlowPart),
	"rentables":   int64(RentablesRAFlowPart),
	"parentchild": int64(ParentChildRAFlowPart),
	"tie":         int64(TieRAFlowPart),
}
