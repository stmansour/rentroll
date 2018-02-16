package worker

import "tws"

// WorkerAsmt et al., are UIDs for automated processes. The negative number
// space is reserved for automated processes.
const (
	WorkerAsmt    = -1
	WorkerAsmtDes = "AsmtCreator"
)

// Worker describes a tws-based worker function
type Worker struct {
	Designator string // short name
	Name       string // name of this worker
	UID        int64  // uid should be a negative number, unique among all workers
	FLAGS      uint64 // 1<<0 = task availability: 0 means it cannot be used as a task, 1 means it can
	Handler    func(*tws.Item)
}

// WorkerRegistry is where workers register themselves to the infrastructure
var WorkerRegistry = map[string]Worker{
	WorkerAsmtDes: {WorkerAsmtDes, "Assessment Instance Creator", WorkerAsmt, uint64(0), CreateAssessmentInstances},
}

// Register is used if there are workers created outside this package
//
// INPUT
//  d       designator - a short name
//  n	    human readable identifer
//  u       user id of the automated process used for database updates
//  FLAGS   indicates availability for tasks, and more
//  Handler callable function
//
// RETURNS
//  nothing at this time
//-----------------------------------------------------------------------------
func Register(d, n string, u int64, FLAGS uint64, h func(*tws.Item)) {
	var w = Worker{
		Designator: d,
		Name:       n,
		UID:        u,
		FLAGS:      FLAGS,
		Handler:    h,
	}
	WorkerRegistry[n] = w
}
