package rlib

import "fmt"

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
