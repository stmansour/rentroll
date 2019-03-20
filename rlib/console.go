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

// ConsoleDate is a quick way to get a formated date...
func ConsoleDate(d *time.Time) string {
	return d.Format(RRDATEFMT3)
}

// ConsoleJSONDate is a quick way to get a formated date...
func ConsoleJSONDate(x *JSONDateTime) string {
	d := time.Time(*x)
	return d.Format(RRDATEFMT3)
}

// ConJDt is a shortcut name for ConsoleJSONDate
func ConJDt(x *JSONDateTime) string { return ConsoleJSONDate(x) }

// ConDt is a shortcut name for ConsoleDate
func ConDt(d *time.Time) string {
	return ConsoleDate(d)
}

// ConsoleJSONDRange is a wrapper around ConsoleDRange for JSONDates
func ConsoleJSONDRange(d1, d2 *JSONDate) string {
	dt1 := time.Time(*d1)
	dt2 := time.Time(*d2)
	return ConsoleDRange(&dt1, &dt2)
}

// ConJDRange is a shortcut name for ConsoleJSONDRange
func ConJDRange(x, y *JSONDate) string { return ConsoleJSONDRange(x, y) }

// ConsoleJSONDtRange is a wrapper around ConsoleDRange for JSONDates
func ConsoleJSONDtRange(d1, d2 *JSONDateTime) string {
	dt1 := time.Time(*d1)
	dt2 := time.Time(*d2)
	return ConsoleDRange(&dt1, &dt2)
}

// ConJDtRange is a shortcut name for ConsoleJSONDtRange
func ConJDtRange(x, y *JSONDateTime) string { return ConsoleJSONDtRange(x, y) }
