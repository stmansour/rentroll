package worker

import (
	"os"
	"rentroll/rlib"
	"time"
	"tws"
)

// WorkerAsmt et al., are UIDs and Designators for automated processes. The negative number
// space is reserved for automated processes.
const (
	WorkerAsmt           = -1
	WorkerAsmtDes        = "CreateAssessmentInstances"
	TaskManual           = -2
	WorkerTskMan         = "ManualTaskBot"
	RARBcacheBot         = -3
	RARBcacheBotDes      = "RARBcacheBot"
	SecDepCacheBot       = -4
	SecDepCacheBotDes    = "SecDepCacheBot"
	AcctSliceCacheBot    = -5
	AcctSliceCacheBotDes = "AcctSliceCacheBot"
	ARSliceCacheBot      = -6
	ARSliceCacheBotDes   = "ARSliceCacheBot"
)

// Worker describes a tws-based worker function
type Worker struct {
	Designator string          // short name
	Name       string          // name of this worker
	UID        int64           // uid should be a negative number, unique among all workers
	FLAGS      uint64          // 1<<0 = task availability: 0 means it cannot be used as a task, 1 means it can
	Handler    func(*tws.Item) // function that does the work
}

// WorkerRegistry is where workers register themselves to the infrastructure
var WorkerRegistry = map[string]Worker{
	WorkerAsmtDes:        {WorkerAsmtDes, "Assessment Instance Bot", WorkerAsmt, uint64(0), CreateAssessmentInstances},
	WorkerTskMan:         {WorkerTskMan, "Manual Task Bot", TaskManual, uint64(1), ProcessManualTask},
	RARBcacheBotDes:      {RARBcacheBotDes, "Clean RARBalance Cache", RARBcacheBot, uint64(0), CleanRARBalanceCache},
	SecDepCacheBotDes:    {SecDepCacheBotDes, "Clean SecDepBalance Cache", SecDepCacheBot, uint64(0), CleanSecDepBalanceCache},
	AcctSliceCacheBotDes: {AcctSliceCacheBotDes, "Clean AcctSlice Cache", AcctSliceCacheBot, uint64(0), CleanAcctSliceCache},
	ARSliceCacheBotDes:   {ARSliceCacheBotDes, "Clean ARSlice Cache", ARSliceCacheBot, uint64(0), CleanARSliceCache},
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
			rlib.Ulog("Registered Worker: %s\n", v.Name)
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
