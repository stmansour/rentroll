package main

// This is a simple program to read two files, a lineNumbers files, and an
// Text file.  The program outputs the Text file`lines but it
// filters out the lines listed in the lineNumbers file.

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// App is the global application structure
var App struct {
	Lines   []int  // line numbers to remove
	FLines  string // file name of line numbers to remove
	FSource string // source text file
}

func readCommandLineArgs() {
	flptr := flag.String("l", "", "file with line numbers to remove from source")
	fsptr := flag.String("f", "", "file to remove lines from")

	flag.Parse()

	App.FLines = *flptr
	App.FSource = *fsptr
	if len(App.FLines) == 0 {
		fmt.Printf("You must supply the -l option, the name of the file containing line numbers to remove\n")
		os.Exit(1)
	}
	if len(App.FSource) == 0 {
		fmt.Printf("You must supply the -f option, the name of the file to filter\n")
		os.Exit(1)
	}
}

func main() {
	readCommandLineArgs()

	//-------------------------------------
	// First, read in the line numbers...
	//-------------------------------------
	fl, err := os.Open(App.FLines)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	flscan := bufio.NewScanner(fl)
	lineno := 0
	for flscan.Scan() {
		lineno++
		s := flscan.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%s: line %d  -  error converting %s to a number: %s\n", App.FLines, lineno, s, err.Error())
			os.Exit(1)
		}
		App.Lines = append(App.Lines, i)
	}

	//----------------------------------------
	// Read the text, filter lines to output
	//----------------------------------------
	f, err := os.Open(App.FSource)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	fsscan := bufio.NewScanner(f)
	lineno = 0 // line number just read
	j := 0     // index of next line number to omit
	maxj := len(App.Lines)
	// rlib.Console("App.Lines has %d elements\n", maxj)
	for fsscan.Scan() {
		lineno++
		s := fsscan.Text()
		prt := true
		if j < maxj {
			// rlib.Console("j = %d, lineno = %d, App.Lines[j] = %d\n", j, lineno, App.Lines[j])
			prt = App.Lines[j] != lineno
			// rlib.Console("prt = %t\n", prt)
			if !prt {
				j++
			}
		}
		if prt {
			fmt.Println(s)
		}
	}
}
