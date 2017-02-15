package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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

//var WSTypes map[string]reflect.Type // so that we can instance variables of a paraticular top

// ListVars lists the names of the variables within a struct and their types
func ListVars(a interface{}) {
	fmt.Printf("Type Of a = %s\n", reflect.TypeOf(a).Name())
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		fmt.Printf("%d. %s   %s\n", j, v.Type().Field(j).Name, f.Type().Name())
	}
}

// Directive is a struct describing a particular directive within the WS DOC comments
// and the handler function to process the data following the directive
type Directive struct {
	directive string
	handler   func(string, *Directive)
}

// Directives contains the list of Directive struct for fields recognized within the
// WS DOC comments and the handlers for each.
var Directives = []Directive{
	{directive: "@url", handler: handleURL},
	{directive: "@synopsis", handler: handleSynopsis},
	{directive: "@description", handler: handleDescription},
	{directive: "@input", handler: handleInput},
	{directive: "@response", handler: handleResponse},
}

func handleURL(s string, d *Directive) {
	s1 := s[len(d.directive):]
	fmt.Printf("URL: %s\n", s1)
}

func handleSynopsis(s string, d *Directive) {
	s1 := s[len(d.directive):]
	fmt.Printf("Synopsis: %s\n", s1)
}

func handleDescription(s string, d *Directive) {
	s1 := s[len(d.directive):]
	fmt.Printf("Description: %s\n", s1)
}

func handleInput(s string, d *Directive) {
	s1 := s[len(d.directive):]
	fmt.Printf("Input: %s\n", s1)
}
func handleResponse(s string, d *Directive) {
	s1 := s[len(d.directive):]
	fmt.Printf("Response: %s\n", s1)
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
			if strings.Index(sl, Directives[j].directive) == 0 {
				Directives[j].handler(s, &Directives[j])
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
	fmt.Printf("Process %s\n", path)

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
	var a AtestStruct
	flag.Parse()
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}
	if err := filepath.Walk(root, processGoFiles); err != nil {
		fmt.Printf("Error walking file path = %s]n", err)
	}
	fmt.Printf("\n\n")
	ListVars(&a)
}
