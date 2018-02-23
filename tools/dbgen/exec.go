package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/kardianos/osext"
)

// InitEmptyDB executes an external command and returns its return code
//
// INPUTS
//  cmdname  - the name of the command to execute
//	args     - a slice of command line args for cmdname
//
// RETURNS
//  exit code from command execution
//-----------------------------------------------------------------------------
func InitEmptyDB() int {
	// rlib.Console("Reading empty.sql\n")
	cmdname := "mysql"
	var args = []string{"--no-defaults", "rentroll"}
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fname := folderPath + "/empty.sql"
	// rlib.Console("empty db read from: %q\n", fname)
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	// rlib.Console("Read %d bytes\n", len(bytes))
	cmd := exec.Command(cmdname, args...)
	cmdIn, _ := cmd.StdinPipe()
	cmdOut, _ := cmd.StdoutPipe()
	cmd.Start()
	cmdIn.Write(bytes)
	cmdIn.Close()
	cmdBytes, _ := ioutil.ReadAll(cmdOut)
	cmd.Wait()
	cmdoutput := string(cmdBytes)
	fmt.Printf("%s", cmdoutput)
	return 0
}
