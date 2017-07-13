package worker

import (
	"os"
	"rentroll/rlib"
	"time"
	"tws"
)

var workers = []struct {
	Name   string
	Worker func(*tws.Item)
}{
	{"CreateAssessmentInstances", CreateAssessmentInstances},
}

// Init registers the TWS functions needed by RentRoll
func Init() {
	funcname := "worker.init"
	for i := 0; i < len(workers); i++ {
		tws.RegisterWorker(workers[i].Name, workers[i].Worker) // first, register our handler
		m, err := tws.FindItem(workers[i].Name)                // next, see if we are already registered
		if err != nil {
			rlib.LogAndPrintError(funcname, err)
			os.Exit(1)
		}
		if len(m) < 1 {
			// No task in the schedule.  Add it
			item := tws.Item{
				Owner:        workers[i].Name,
				OwnerData:    "",
				WorkerName:   workers[i].Name,
				ActivateTime: time.Now(),
			}
			tws.InsertItem(&item)
			rlib.Ulog("Registered Worker: %s\n", workers[i].Name)
		}
	}
}
