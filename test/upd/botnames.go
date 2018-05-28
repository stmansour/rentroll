package main

import (
	"fmt"
	"rentroll/rlib"
	"sort"
)

// CheckBotNames prints out the entries in rlib.BotRange
// in a repeatable,predictable order.  If we just use the
// range function and list the contents it will be in different
// orders all the time and the test will fail.
//-----------------------------------------------------------------------------
func CheckBotNames() {
	var idx []rlib.BotRegistryEntry
	for _, v := range rlib.BotReg {
		idx = append(idx, v)
	}
	sort.Slice(idx[:], func(i, j int) bool {
		return idx[i].BotID < idx[j].BotID
	})

	for i := 0; i < len(idx); i++ {
		fmt.Printf("%3d. %3d, %-20s %s\n", i, idx[i].BotID, idx[i].Designator, idx[i].Name)
	}
}
