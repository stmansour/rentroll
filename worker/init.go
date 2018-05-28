package worker

import (
	"os"
	"rentroll/rlib"
	"time"
	"tws"
)

// Worker describes a tws-based worker function
type Worker struct {
	Bot     rlib.BotRegistryEntry // has uid,
	FLAGS   uint64                // 1<<0 = task availability: 0 means it cannot be used as a task, 1 means it can
	Handler func(*tws.Item)       // function that does the work
}

// WorkerRegistry is where workers register themselves to the infrastructure
var WorkerRegistry = map[string]Worker{
	//------------------------------------------------------------------
	// The following workers ARE NOT available to users for tasklists
	//------------------------------------------------------------------
	rlib.BotReg[rlib.WorkerAsmt].Designator:        {rlib.BotReg[rlib.WorkerAsmt], uint64(0), CreateAssessmentInstances},
	rlib.BotReg[rlib.RARBcacheBot].Designator:      {rlib.BotReg[rlib.RARBcacheBot], uint64(0), CleanRARBalanceCache},
	rlib.BotReg[rlib.SecDepCacheBot].Designator:    {rlib.BotReg[rlib.SecDepCacheBot], uint64(0), CleanSecDepBalanceCache},
	rlib.BotReg[rlib.AcctSliceCacheBot].Designator: {rlib.BotReg[rlib.AcctSliceCacheBot], uint64(0), CleanAcctSliceCache},
	rlib.BotReg[rlib.ARSliceCacheBot].Designator:   {rlib.BotReg[rlib.ARSliceCacheBot], uint64(0), CleanARSliceCache},
	rlib.BotReg[rlib.TLReportBot].Designator:       {rlib.BotReg[rlib.TLReportBot], uint64(0), TLChecker},
	rlib.BotReg[rlib.TLInstanceBot].Designator:     {rlib.BotReg[rlib.TLInstanceBot], uint64(0), TLInstanceBot},

	//------------------------------------------------------------------
	// The following workers ARE available to users for tasklists
	//------------------------------------------------------------------
	rlib.BotReg[rlib.TaskManual].Designator: {rlib.BotReg[rlib.TaskManual], uint64(1), ProcessManualTask},
}

// Init registers the TWS functions needed by RentRoll
//-----------------------------------------------------------------------------
func Init() {
	InitCore(WorkerRegistry)
}

// InitCore enables workers defined in the array of workers, w
//
// INPUTS
//  w   - map of workers
//
// RETURNS
//  nothing
//-----------------------------------------------------------------------------
func InitCore(w map[string]Worker) {
	funcname := "worker.InitCore"
	for k, v := range w {
		tws.RegisterWorker(k, v.Handler) // first, register our handler
		m, err := tws.FindItem(k)        // next, see if we are already registered
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			os.Exit(1)
		}
		if len(m) < 1 {
			// No task in the schedule.  Add it
			item := tws.Item{
				Owner:        k,
				OwnerData:    "",
				WorkerName:   k,
				ActivateTime: time.Now(),
			}
			tws.InsertItem(&item)
			rlib.Ulog("Registered Worker: %s\n", v.Bot.Name)
		}
	}
}

// GetWorkerList returns an array of worker designators
//
// INPUTS
//  nothing
//
// RETURNS
//  []string with the register worker designators
//-----------------------------------------------------------------------------
func GetWorkerList() []string {
	var m []string
	for k, v := range WorkerRegistry {
		if v.FLAGS&1 > 0 {
			m = append(m, k)
		}
	}
	return m
}
