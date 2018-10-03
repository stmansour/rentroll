package rlib

import (
	"fmt"
	"time"
)

var consoleEnable = true

// EnableConsole causes Console statements to print
func EnableConsole() { consoleEnable = true }

// DisableConsole causes Console statements to print
func DisableConsole() { consoleEnable = false }

// Console is the standard logger
func Console(format string, a ...interface{}) {
	if consoleEnable {
		fmt.Printf(format, a...)
	}
}

// ConsoleDRange formats a date range string.
// format date range
// print date range
// format dates, print dates
func ConsoleDRange(d1, d2 *time.Time) string {
	return fmt.Sprintf("%s, %s", d1.Format(RRDATEFMT3), d2.Format(RRDATEFMT3))
}
