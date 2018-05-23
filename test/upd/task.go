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
	tldef.Cycle = rlib.RECURMONTHLY
	tldef.Name = "Monthly Close"
	tldef.Epoch = time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	tldef.EpochDue = time.Date(2018, time.January, 31, 17, 0, 0, 0, time.UTC)
	tldef.EpochPreDue = time.Date(2018, time.January, 20, 17, 0, 0, 0, time.UTC)

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
	pivot := time.Date(2018, time.February, 3, 12, 32, 13, 0, time.UTC)
	_, err = rlib.CreateTaskListInstance(ctx, tldef.TLDID, 0, &pivot)
	if err != nil {
		fmt.Printf("CreateTaskListInstance:  error = %s\n", err.Error())
	}
}
