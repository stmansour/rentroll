package worker

import (
	"os"
	"rentroll/rlib"
	"time"
	"tws"
)

// TWSWorker defines a timed worker
type TWSWorker struct {
	Name   string
	Worker func(*tws.Item)
}

var workers = []TWSWorker{
	{"CreateAssessmentInstances", CreateAssessmentInstances},
	{"CleanRARBalanceCache", CleanRARBalanceCache},
	{"CleanSecDepBalanceCache", CleanSecDepBalanceCache},
	{"CleanAcctSliceCache", CleanAcctSliceCache},
	{"CleanARSliceCache", CleanARSliceCache},
}

// Init registers the TWS functions needed by RentRoll
//-----------------------------------------------------------------------------
func Init() {
	InitCore(workers)
}

// InitCore enables workers defined in the array of workers, w
//
// INPUTS
//  w   - slice of TWSWorkers
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func InitCore(w []TWSWorker) {
	funcname := "worker.init"
	for i := 0; i < len(w); i++ {
		tws.RegisterWorker(w[i].Name, w[i].Worker) // first, register our handler
		m, err := tws.FindItem(w[i].Name)          // next, see if we are already registered
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			os.Exit(1)
		}
		if len(m) < 1 {
			// No task in the schedule.  Add it
			item := tws.Item{
				Owner:        w[i].Name,
				OwnerData:    "",
				WorkerName:   w[i].Name,
				ActivateTime: time.Now(),
			}
			tws.InsertItem(&item)
			rlib.Ulog("Registered Worker: %s\n", workers[i].Name)
		}
	}
}
