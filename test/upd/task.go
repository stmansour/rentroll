package main

import (
	"context"
	"fmt"
	"rentroll/rlib"
	"time"
)

// Tasks is used to test the task related
func Tasks(ctx context.Context, biz *rlib.Business) {
	var tldef rlib.TaskListDefinition
	// First... define a task list that has a "pre-due-date" on the 20th
	// and a due date at 5pm on the last day of the month.
	tldef.BID = biz.BID
	tldef.Cycle = rlib.CYCLEMONTHLY
	tldef.Name = "Monthly Close"
	tldef.DtDue = time.Date(2018, time.January, 31, 17, 0, 0, 0, time.UTC)
	tldef.DtPreDue = time.Date(2018, time.January, 20, 17, 0, 0, 0, time.UTC)

	err := rlib.InsertTaskListDefinition(ctx, &tldef)
	if err != nil {
		fmt.Printf("rlib.InsertTaskListDefinition: error = %s\n", err.Error())
		return
	}

	var due = time.Date(2018, time.January, 31, 20, 0, 0, 0, time.UTC) //
	var predue = time.Date(2018, time.January, 20, 20, 0, 0, 0, time.UTC)
	var t = []rlib.TaskDescriptor{
		{Name: "Delinquency Report", Worker: "Manual", EpochDue: due, EpochPreDue: predue},
		{Name: "Walk the Units", Worker: "Manual", EpochDue: due, EpochPreDue: predue},
		{Name: "Generate Offsets", Worker: "OffsetBot", EpochDue: due, EpochPreDue: predue},
	}

	for i := 0; i < len(t); i++ {
		t[i].TLDID = tldef.TLDID
		err := rlib.InsertTaskDescriptor(ctx, &t[i])
		if err != nil {
			fmt.Printf("rlib.InsertTaskDescriptor: error = %s\n", err.Error())
			return
		}
	}

	//----------------------------------------------
	// Now, create an instance of this task list.
	//----------------------------------------------
	err = CreateTaskListInstance(ctx, tldef.TLDID)
	if err != nil {
		fmt.Printf("CreateTaskListInstance:  error = %s\n", err.Error())
	}

}

// CreateTaskListInstance creates a new task list based on the supplied
// definition and Epoch Date.
//
// INPUTS
//  ctx    - context for database transactions
//  TLDID  - Task List Definition ID
//
// RETURNS
//  error  - any error encountered
//
//-----------------------------------------------------------------------------
func CreateTaskListInstance(ctx context.Context, TLDID int64) error {
	tld, err := rlib.GetTaskListDefinition(ctx, TLDID)
	if err != nil {
		return err
	}

	rlib.Console("Found tld.TLDID = %d, name = %s\n", tld.TLDID, tld.Name)

	td, err := rlib.GetTaskListDescriptors(ctx, tld.TLDID)
	if err != nil {
		return err
	}

	return nil
}
