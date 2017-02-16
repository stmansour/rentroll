package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"rentroll/ws"
	"strings"
)

var wsDocFactory = map[string]interface{}{
//"foo": ws.RentalAgr{},
}

// AtestStruct is a testing type
type AtestStruct struct {
	Foo string
	Bar int
	X   float64
	I   int64
}

// Creator is an integral part of the factory implementation.
// A creator function returns a new struct of different types
// by returning an interface.
type Creator func() interface{}

// WSTypeFactory is a map for creating new data types used by the
// web services routines based on the supplied name.
var WSTypeFactory = map[string]Creator{
	"RentalAgr":  NewRentalAgr,
	"WebRequest": NewWebRequest,
	"WSRAR":      NewWSRAR,
}

// NewRentalAgr is a factory for RentalAgr structs
func NewRentalAgr() interface{} {
	return new(ws.RentalAgr)
}

// NewWebRequest is a factory for WebRequest structs
func NewWebRequest() interface{} {
	return new(ws.WebRequest)
}

// NewWSRAR is a factory for WSRAR structs
func NewWSRAR() interface{} {
	return new(ws.WSRAR)
}

// ListVars lists the names of the variables within a struct and their types
func ListVars(a interface{}) string {
	s := fmt.Sprintf("{\n")
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		s += fmt.Sprintf("    %q: %s", v.Type().Field(j).Name, f.Type().Name())
		if j+1 < v.NumField() {
			s += fmt.Sprintf(",")
		}
		s += fmt.Sprintf("\n")
	}
	s += fmt.Sprintf("}\n")
	return s
}

// Directive is a struct describing a particular Cmd within the WS DOC comments
// and the Handler function to process the data following the Cmd
type Directive struct {
	Cmd     string
	Handler func(string, *Directive)
}

// Directives contains the list of Directive struct for fields recognized within the
// WS DOC comments and the Handlers for each.
var Directives = []Directive{
	{Cmd: "@title", Handler: handleTitle},
	{Cmd: "@url", Handler: handleURL},
	{Cmd: "@synopsis", Handler: handleSynopsis},
	{Cmd: "@method", Handler: handleMethod},
	{Cmd: "@description", Handler: handleDescription},
	{Cmd: "@input", Handler: handleInput},
	{Cmd: "@response", Handler: handleResponse},
}

func handleURL(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("URL:         %s\n", strings.TrimSpace(s1))
}

func handleTitle(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("Title:       %s\n", strings.TrimSpace(s1))
}

func handleSynopsis(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("Synopsis:    %s\n", strings.TrimSpace(s1))
}

func handleDescription(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("Description: %s\n", strings.TrimSpace(s1))
}

func handleMethod(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("Method:      %s\n", strings.TrimSpace(s1))
}

func handleInput(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("Input:\n%s\n", getStructDef(s1))
}

func handleResponse(s string, d *Directive) {
	s1 := s[len(d.Cmd):]
	fmt.Printf("Response:\n%s\n", getStructDef(s1))
}

func getStructDef(s1 string) string {
	ss1 := strings.Split(s1, " ")
	for i := 0; i < len(ss1); i++ {
		t := strings.TrimSpace(ss1[i])
		if len(t) == 0 {
			continue
		}
		_, ok := WSTypeFactory[t]
		if !ok {
			fmt.Printf("Could not find type %s in factory\n", t)
			return ""
		}
		x := WSTypeFactory[t]()
		return ListVars(x)
	}
	return ""
}

// processWebDocLines builds the documentation for a single web service call. The content
// defining the document is contained in the supplied array of strings.  In particular,
// the definitions it looks for are:
//		@URL		 - the format of the url
//		@Synopsis	 - 1 sentence description
//		@Description - detailed explanation of the web service
//		@Input		 - format and content of data required as input
//		@Response	 - format and content of data returned -- if it is a Go data type it
//						will be expanded
func processWebDocLines(sa []string) {
	if len(sa) == 0 {
		return
	}
	fmt.Printf("\n\nnew web service\n")
	for i := 0; i < len(sa); i++ {
		ss := strings.Split(sa[i], "//")
		if len(ss) < 2 {
			continue
		}
		for j := 0; j < len(Directives); j++ {
			s := strings.TrimSpace(ss[1])
			sl := strings.ToLower(s)
			if strings.Index(sl, Directives[j].Cmd) == 0 {
				Directives[j].Handler(s, &Directives[j])
				break
			}
		}
	}
}

// isCommentContaining first checks to see if the supplied line is a go comment
// (that is, if it contains "//").  If so, it looks for the supplied target
// string to be the first thing is sees after the comment chars (ignoring whitespace).
// If target is found it returns true, otherwise it returns false.
func isCommentContaining(s, target string) bool {
	ss := strings.Split(s, "//")
	if len(ss) < 2 {
		return false // it's not a comment
	}
	return strings.Index(strings.TrimSpace(ss[1]), target) == 0
}

// processGoFiles searches for go files, exclude go unit test files
// It then opens the file and scans for comment lines containing markers
// for Web Services Docs.  The markers surrounding these lines are:
//  	wsdoc {
//      wsdoc }
// All lines between these two markers are sent for further processing.
func processGoFiles(path string, f os.FileInfo, err error) error {
	if f.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(f.Name(), "_test.go") {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		s := scanner.Text()
		if !found { // search for start of web service doc
			found = isCommentContaining(s, "wsdoc {") // look for start of ws docs
			continue
		}
		if isCommentContaining(s, "wsdoc }") { // look for end of ws docs
			processWebDocLines(lines) // transform into document
			found = false
			lines = []string{}
			continue
		}
		lines = append(lines, s) // save all lines between start and end of ws docs
	}

	return scanner.Err()
}

func main() {
	root := "."
	flag.Parse()
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}
	if err := filepath.Walk(root, processGoFiles); err != nil {
		fmt.Printf("Error walking file path = %s]n", err)
	}
	fmt.Printf("\n")
}
