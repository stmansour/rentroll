package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"rentroll/rlib"
	"rentroll/ws"
	"strings"
	"time"
)

// Playground: https://play.golang.org/p/R0xCuJXkzL

// Creator is an integral part of the factory implementation.
// A creator function returns a new struct of different types
// by returning an interface.
type Creator func() interface{}

// ProtocolJSON describes an individual field in the JSON protocol
// for this web service
type ProtocolJSON struct {
	Field    string
	DataType string
	Optional bool
}

// DirectiveData is a struct of data describing the web service. Its members
// are set as the comments are parsed. The data in this struct is used to
// create an html file describing the web service.
type DirectiveData struct {
	Title       string         // name of web service
	URL         string         // programming url format
	Synopsis    string         // One line explanation
	Method      string         // POST, GET, ...
	Description string         // detailed explanation
	Input       []ProtocolJSON // JSON input data
	Response    []ProtocolJSON // JSON response data
	Filename    string         // the name of the html file describing the web service
	ID          string         // a unique id used in the UI
}

// Directive is a struct describing a particular Cmd within the WS DOC comments
// and the Handler function to process the data following the Cmd
type Directive struct {
	Cmd     string
	Handler func(string, *Directive)
	D       *DirectiveData
}

// IndexData is an array of DirectiveData structs used to generate an index page.
var IndexData struct {
	TOC     []DirectiveData
	Date    string
	Version string
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

// WSTypeFactory is a map for creating new data types used by the
// web services routines based on the supplied name.
var WSTypeFactory = map[string]Creator{
	"ColSort":                    NewColSort,
	"GenSearch":                  NewGenSearch,
	"GetRentableResponse":        NewGetRentableResponse,
	"GLAccount":                  NewGLAccount,
	"PrRentableOther":            NewPrRentableOther,
	"RAPeople":                   NewRAPeople,
	"RAPeopleResponse":           NewRAPeopleResponse,
	"RAPets":                     NewRAPets,
	"RAR":                        NewWSRAR,
	"RentalAgr":                  NewRentalAgr,
	"RentalAgreementPet":         NewRentalAgreementPet,
	"RentalAgrSearchResponse":    NewRentalAgrSearchResponse,
	"SearchGLAccountsResponse":   NewSearchGLAccountsResponse,
	"SearchRentablesResponse":    NewSearchRentablesResponse,
	"SearchTransactantsResponse": NewSearchTransactantsResponse,
	"SvcStatusResponse":          NewSvcStatusResponse,
	"WebRequest":                 NewWebRequest,
	"GetRentalAgreementResponse": NewGetRentalAgreementResponse,
}

// NewGetRentableResponse is a factory for GetRentableResponse structs
func NewGetRentableResponse() interface{} {
	return new(ws.GetRentableResponse)
}

// NewRAPets is a factory for RAPets structs
func NewRAPets() interface{} {
	return new(ws.RAPets)
}

// NewRAPeopleResponse is a factory for RAPeopleResponse structs
func NewRAPeopleResponse() interface{} {
	return new(ws.RAPeopleResponse)
}

// NewSearchTransactantsResponse is a factory for SearchTransactantsResponse structs
func NewSearchTransactantsResponse() interface{} {
	return new(ws.SearchTransactantsResponse)
}

// NewSearchGLAccountsResponse is a factory for SearchGLAccountsResponse structs
func NewSearchGLAccountsResponse() interface{} {
	return new(ws.SearchGLAccountsResponse)
}

// NewSvcStatusResponse is a factory for SvcStatusResponse structs
func NewSvcStatusResponse() interface{} {
	return new(ws.SvcStatusResponse)
}

// NewGetRentalAgreementResponse is a factory for GetRentalAgreementResponse structs
func NewGetRentalAgreementResponse() interface{} {
	return new(ws.GetRentalAgreementResponse)
}

// NewRentalAgrSearchResponse is a factory for RentalAgrSearchResponse structs
func NewRentalAgrSearchResponse() interface{} {
	return new(ws.RentalAgrSearchResponse)
}

// NewColSort is a factory for ColSort structs
func NewColSort() interface{} {
	return new(ws.ColSort)
}

// NewGenSearch is a factory for GenSearch structs
func NewGenSearch() interface{} {
	return new(ws.GenSearch)
}

// NewRentalAgr is a factory for RentalAgr structs
func NewRentalAgr() interface{} {
	return new(ws.RentalAgr)
}

// NewPrRentableOther is a factory for PrRentableOther structs
func NewPrRentableOther() interface{} {
	return new(ws.PrRentableOther)
}

// NewSearchRentablesResponse is a factory for SearchRentablesResponse structs
func NewSearchRentablesResponse() interface{} {
	return new(ws.SearchRentablesResponse)
}

// NewRentalAgreementPet is a factory for RentalAgreementPet structs
func NewRentalAgreementPet() interface{} {
	return new(rlib.RentalAgreementPet)
}

// NewWSPets is a factory for RAPets structs
func NewWSPets() interface{} {
	return new(ws.RAPets)
}

// NewRAPeople is a factory for RAPeople structs
func NewRAPeople() interface{} {
	return new(ws.RAPeople)
}

// NewWebRequest is a factory for WebRequest structs
func NewWebRequest() interface{} {
	return new(ws.WebRequest)
}

// NewWSRAR is a factory for RAR structs
func NewWSRAR() interface{} {
	return new(ws.RAR)
}

// NewGLAccount is a factory for GLAccount structs
func NewGLAccount() interface{} {
	return new(rlib.GLAccount)
}

// AnalyzeType determines:
//		if the field is a slice
//		if the type requires recursion
//		the type name to use for the factory if recursion is needed
// The return values are:
//		IsSlice bool    -- true if the field is a slice
// 		Recurse bool	-- true if recursion is required
//		Tname   string	-- the data type for a call to the factory
func AnalyzeType(t string) (bool, bool, string) {
	Tname := t
	IsSlice := false
	if pos := strings.Index(Tname, "[]"); pos >= 0 {
		Tname = t[2+pos:]
		IsSlice = true
	}
	if i := strings.Index(Tname, "."); i >= 0 {
		Tname = Tname[i+1:]
	}
	// Is Tname in our factory?
	_, Recursion := WSTypeFactory[Tname]
	return IsSlice, Recursion, Tname
}

// ListVars lists the names of the variables within a struct and their types
func ListVars(a interface{}, d *Directive, depth int) []ProtocolJSON {
	var m []ProtocolJSON
	v := reflect.ValueOf(a).Elem()
	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "...."
	}
	for j := 0; j < v.NumField(); j++ {
		var p ProtocolJSON
		f := v.Field(j)
		p.Field = prefix + v.Type().Field(j).Name
		p.DataType = f.Type().String()
		isSlice, recurse, rtype := AnalyzeType(p.DataType)
		sl := ""
		if isSlice {
			sl = "[]"
		}
		p.DataType = sl + rtype
		fmt.Printf("Name = %s, Recurse = %t,  Kind = %s,  type = %s\n", p.Field, recurse, f.Kind().String(), rtype)
		m = append(m, p)
		if recurse {
			fmt.Printf("Recursing into type = %s\n", rtype)
			x := WSTypeFactory[rtype]()
			n := ListVars(x, d, depth+1)
			m = append(m, n...)
			fmt.Printf("Appended %d members to m\n", len(n))
		}
	}
	return m
}

func handleURL(s string, d *Directive) {
	d.D.URL = strings.TrimSpace(s[len(d.Cmd):])
}

func handleTitle(s string, d *Directive) {
	d.D.Title = strings.TrimSpace(s[len(d.Cmd):])
	d.D.ID = rlib.Stripchars(strings.ToLower(d.D.Title), " ")
	d.D.Filename = d.D.ID + ".html"
}

func handleSynopsis(s string, d *Directive) {
	d.D.Synopsis = strings.TrimSpace(s[len(d.Cmd):])
}

func handleDescription(s string, d *Directive) {
	d.D.Description = strings.TrimSpace(s[len(d.Cmd):])
}

func handleMethod(s string, d *Directive) {
	d.D.Method = strings.TrimSpace(s[len(d.Cmd):])
}

func handleInput(s string, d *Directive) {
	s1 := strings.TrimSpace(s[len(d.Cmd):])
	d.D.Input = getStructDef(s1, d)
}

func handleResponse(s string, d *Directive) {
	s1 := strings.TrimSpace(s[len(d.Cmd):])
	d.D.Response = getStructDef(s1, d)
}

func getStructDef(s string, d *Directive) []ProtocolJSON {
	ss := strings.Split(s, " ")
	for i := 0; i < len(ss); i++ {
		t := strings.TrimSpace(ss[i])
		if len(t) == 0 {
			continue
		}
		_, ok := WSTypeFactory[t]
		if ok {
			x := WSTypeFactory[t]()
			return ListVars(x, d, 0)
		}
		if strings.ToLower(t) == "string" {
			var p ProtocolJSON
			p.Field = "data"
			p.DataType = "string"
			var m []ProtocolJSON
			m = append(m, p)
			return m
		}
	}
	return []ProtocolJSON{}
}

func generateHTMLRefPage(d *DirectiveData) error {
	path := "./doc"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|0777)
	}
	f, err := os.Create(path + "/" + d.Filename)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("doc.html").ParseFiles("doc.html")
	if nil != err {
		fmt.Printf("Error loading template: %v\n", err)
	}
	if err = t.Execute(f, d); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}
	return err
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
	var d DirectiveData
	for i := 0; i < len(sa); i++ {
		ss := strings.Split(sa[i], "//")
		if len(ss) < 2 {
			continue
		}
		for j := 0; j < len(Directives); j++ {
			s := strings.TrimSpace(ss[1])
			sl := strings.ToLower(s)
			if strings.Index(sl, Directives[j].Cmd) == 0 {
				Directives[j].D = &d
				Directives[j].Handler(s, &Directives[j])
				break
			}
		}
	}
	if err := generateHTMLRefPage(&d); err != nil {
		fmt.Printf("Error generating reference page: %s\n", err.Error())
	}
	IndexData.TOC = append(IndexData.TOC, d)
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

func generateDocs() error {
	f, err := os.Create("./doc/docs.html")
	if err != nil {
		return err
	}
	defer f.Close()

	IndexData.Date = time.Now().Format("Jan 2, 2006  3:04PM MST")
	IndexData.Version = "1.0"
	t, err := template.New("docs.html").ParseFiles("docs.html")
	if nil != err {
		fmt.Printf("Error loading template: %v\n", err)
	}
	if err = t.Execute(f, &IndexData); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}

	return err
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
	if scanner.Err() != nil {
		fmt.Printf("Error scanning file: %s\n", scanner.Err().Error())
		return err
	}
	err = generateDocs()
	return err
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
