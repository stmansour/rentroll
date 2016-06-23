package main

import "fmt"

// UILedgerTextReport prints a report of data that will be used to format a ledger UI.
// This routine is primarily for testing
func UILedgerTextReport(ui *RRuiSupport) {
	fmt.Printf("%40s  %10s  %12s\n", "Name", "GLNumber", "Balance")
	for i := 0; i < len(ui.LDG.XL); i++ {
		fmt.Printf("%40s  %10s  %12.2f\n", ui.LDG.XL[i].G.Name, ui.LDG.XL[i].G.GLNumber, ui.LDG.XL[i].LM.Balance)
	}
	s := ""
	for i := 0; i < 66; i++ {
		s += "-"
	}
	fmt.Println(s)
	fmt.Printf("%40s  %10s  %12.2f\n", " ", " ", LMSum(&ui.LDG.XL))
}
