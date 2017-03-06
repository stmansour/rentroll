package ws

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"rentroll/rlib"
	"strings"

	"github.com/kardianos/osext"
)

// EpochInstance is a list of strings that describe a recurring entity.
// The entity can be either an "epoch", the definition of the recurring
// series, or an "instance", one of the members of the recurring series.
var EpochInstance = []string{"Epoch", "Instance"}

// USStateAbbr is an array of American state abbreviations for the UI.
var USStateAbbr = []string{"AK", "AL", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}

var yesno = []string{"no", "yes"}

// String2Int64MapToJSList generates a string of JS code that assigns
// all the map strings in m to an array.  Suitable for a JS eval call.
func String2Int64MapToJSList(name string, m *rlib.Str2Int64Map) string {
	s := name + "=["
	l := len(*m)
	i := 0
	for k := range *m {
		s += "'" + k + "'"
		if i+1 < l {
			s += ","
		}
		i++
	}
	s += "];\n"
	return s
}

// StringListToJSList generates a string of JS code that assigns
// all the map strings in m to an array.  Suitable for a JS eval call.
func StringListToJSList(name string, m *[]string) string {
	s := name + "=["
	l := len(*m)
	i := 0
	for k := 0; k < l; k++ {
		s += "'" + (*m)[k] + "'"
		if i+1 < l {
			s += ","
		}
		i++
	}
	s += "];\n"
	return s
}

var smapToJS = []struct {
	name   string
	valmap *rlib.Str2Int64Map
}{
	{"assignmentTimeList", &rlib.AssignmentTimeMap},
	{"businesses", &rlib.RRdb.BUDlist},
	{"companyOrPerson", &rlib.CompanyOrPersonMap},
	{"renewalMap", &rlib.RenewalMap},
	{"cycleFreq", &rlib.CycleFreqMap},
}

var ssliceToJS = []struct {
	name     string
	valslice *[]string
}{
	{"epochInstance", &EpochInstance},
	{"yesNoList", &yesno},
	{"usStateAbbr", &USStateAbbr},
}

// SvcUILists returns JSON for the Javascript lists needed for the UI.  Typically,
// these lists are put into a map such as rlib.Str2Int64Map or a slice of strings.
// Then the map or slice is entered into either smapToJS or ssliceToJS so that it
// will be automatically sent to the UI.  Within the UI, all the application
// related data is maintained in a single global variable named "app". So, adding
// a entry with name "foo" either of the structs above will result in the JS array
// being named "app.foo"
//
// wsdoc {
//  @Title  Get UI Lists
//	@URL /v1/uilists/:LANGUAGE/:TEMPLATE
//  @Method  GET
//	@Synopsis Return string lists that are used in the UI
//  @Desc Return data can be processed by eval() to create the string lists used in the UI.
//  @Desc LANGUAGE is optional, it defaults to "en-us".  TEMPLATE is also options, and it
//  @Desc defaults to "default"
//	@Input WebGridSearchRequest
//  @Response string
// wsdoc }
func SvcUILists(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcUILists"
	fmt.Printf("Entered %s\n", funcname)
	language := "en-us"   // start with default
	template := "default" // start with default
	s, err := url.QueryUnescape(strings.TrimSpace(r.URL.String()))
	if err != nil {
		e := fmt.Errorf("%s: Error with url.QueryUnescape:  %s", funcname, err.Error())
		rlib.Ulog(e.Error()) // this is not fatal, we'll just use the default values
	} else {
		if strings.HasSuffix(s, "/") { // /v1/uilists/language/template/
			s = s[:len(s)-1] // /v1/uilists/language/template
		}
		ss := strings.Split(s[1:], "/") //   ss = ["v1" "uilists" "language" "template"]
		if len(ss) > 2 {
			language = ss[2]
		}
		if len(ss) > 3 {
			template = ss[3]
		}
	}
	fmt.Printf("Language: %s\nTemplate: %s\n", language, template)

	for i := 0; i < len(smapToJS); i++ {
		io.WriteString(w, String2Int64MapToJSList("app."+smapToJS[i].name, smapToJS[i].valmap))
	}

	for i := 0; i < len(ssliceToJS); i++ {
		io.WriteString(w, StringListToJSList("app."+ssliceToJS[i].name, ssliceToJS[i].valslice))
	}

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		SvcGridErrorReturn(w, err)
		return
	}
	fname := folderPath + "/html/" + language + "/" + template + "/strings.csv"
	fmt.Printf("fname = %s\n", fname)
	_, err = os.Stat(fname)
	if nil != err {
		e := fmt.Errorf("Unknown language / template :   %s / %s", language, template)
		SvcGridErrorReturn(w, e)
		return
	}
	m := rlib.LoadCSV(fname)
	for i := 0; i < len(m); i++ {
		s := fmt.Sprintf("%s.%s='%s';\n", "app", m[i][0], m[i][1])
		io.WriteString(w, s)
	}
}
