package ws

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"rentroll/bizlogic"
	"rentroll/rlib"
	"rentroll/worker"
	"sort"
	"strings"

	"github.com/kardianos/osext"
)

// rentroll report row flags
var rrFLAGS = rlib.Str2Int64Map{
	"RentRollMainRow":       rlib.RentRollMainRow,
	"RentRollSubTotalRow":   rlib.RentRollSubTotalRow,
	"RentRollBlankRow":      rlib.RentRollBlankRow,
	"RentRollGrandTotalRow": rlib.RentRollGrandTotalRow,
}

// EpochInstance is a list of strings that describe a recurring entity.
// The entity can be either an "epoch", the definition of the recurring
// series, or an "instance", one of the members of the recurring series.
var EpochInstance = []string{"Epoch", "Instance"}

// USStateAbbr is an array of American state abbreviations for the UI.
var USStateAbbr = []string{"AK", "AL", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}

var yesno = []string{"no", "yes"}

var smapToJS = []struct {
	name   string
	valmap *rlib.Str2Int64Map
}{
	{"assignmentTimeList", &rlib.AssignmentTimeMap},
	{"businesses", &rlib.RRdb.BUDlist},
	{"cycleFreq", &rlib.CycleFreqMap},
}

var idTextMapList = []struct {
	name   string
	valmap *rlib.Str2Int64Map
}{
	{"renewalMap", &rlib.RenewalMap},
	{"companyOrPerson", &rlib.CompanyOrPersonMap},
}

var ssliceToJS = []struct {
	name     string
	valslice *[]string
}{
	{"epochInstance", &EpochInstance},
	{"yesNoList", &yesno},
	{"usStateAbbr", &USStateAbbr},
	{"RSUseStatus", &rlib.RSUseStatus},
	{"RSLeaseStatus", &rlib.RSLeaseStatus},
	{"qbAcctType", &rlib.QBAcctType},
}

// bizMap struct, used in app.BizMap on front-end side
type bizMap struct {
	BID int64
	BUD string
}

// pmtMap struct, used in app.pmtTypes on front-end side
type pmtMap struct {
	PMTID int64
	Name  string
}

// IDTextMap drop list for w2ui {id: ID, text: Text}
type IDTextMap struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

// DepMethMap drop down list for deposit methods
type DepMethMap struct {
	DPMID int64  `json:"id"`
	Text  string `json:"text"`
}

// GetJSDepositMethods builds a datastructure suitable for javascript parsing that
// holds the DepositMethods for all businesses
func GetJSDepositMethods(ctx context.Context) map[string][]DepMethMap {
	var budDepMethods = make(map[string][]DepMethMap)
	for bud, bid := range rlib.RRdb.BUDlist {
		depmethList := []DepMethMap{}
		m, _ := rlib.GetAllDepositMethods(ctx, bid) // get the payment types for this business
		for i := 0; i < len(m); i++ {
			depmethList = append(depmethList, DepMethMap{DPMID: m[i].DPMID, Text: m[i].Method})
		}
		budDepMethods[bud] = depmethList
	}
	return budDepMethods
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
//	@Synopsis Return jsonReponse lists that are used in the UI
//  @Desc Return data can be processed by eval() to create the string lists used in the UI.
//  @Desc LANGUAGE is optional, it defaults to "en-us".  TEMPLATE is also options, and it
//  @Desc defaults to "default"
//	@Input WebGridSearchRequest
//  @Response JSONResponse
// wsdoc }
//-----------------------------------------------------------------------------
func SvcUILists(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcUILists"
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
	rlib.Console("Language: %s\nTemplate: %s\n", language, template)

	// appData will hold the map of data with key as string
	// value would be any type, typically it will be used in front-end side
	// to hold the config values in `app` variable
	appData := make(map[string]interface{})

	// --------------- LIST DOWN BUSINESSESS ----------------------
	// set business map in app
	businessList := []bizMap{}
	for bud, bid := range rlib.RRdb.BUDlist {
		businessList = append(businessList, bizMap{BID: bid, BUD: bud})
	}
	appData["BizMap"] = businessList

	// --------------- LIST DOWN Workers ----------------------
	var wa []IDTextMap
	ss := worker.GetWorkerList()
	for i := 0; i < len(ss); i++ {
		var p = IDTextMap{ID: int64(i), Text: ss[i]}
		wa = append(wa, p)
	}
	appData["workers"] = wa

	// --------------- LIST DOWN Business FLAGS ----------------------
	bizFLAGS := make(map[string]int64)
	for _, bcache := range rlib.RRdb.BizCache {
		bizFLAGS[bcache.BUD] = bcache.FLAGS
	}
	appData["bizFLAGS"] = bizFLAGS

	// --------------- LIST DOWN Assessment FLAGS ----------------------
	appData["asmFLAGS"] = rlib.AsmFLAGS

	// --------------- LIST DOWN Receipt FLAGS ----------------------
	appData["rcptFLAGS"] = rlib.RcptFLAGS

	// --------------- LIST DOWN rentroll report FLAGS --------------
	appData["rrFLAGS"] = rrFLAGS

	// --------------- LIST DOWN account rules FLAGS --------------
	appData["arFLAGS"] = bizlogic.ARFLAGS

	// --------------- LIST DOWN rentable type FLAGS --------------
	appData["rtFLAGS"] = bizlogic.RTFLAGS

	// --------------- LIST DOWN ra flow part types --------------
	appData["raFlowPartTypes"] = rlib.RAFlowPartsMap

	// --------------- MAPPING - smapToJS ----------------------
	for i := 0; i < len(smapToJS); i++ {
		list := []string{}

		// sort out items of valmap, based on int64 value
		sortList := rlib.Int64Range{}
		i64ToSMap := map[int64]string{}
		for s, i64 := range *smapToJS[i].valmap {
			sortList = append(sortList, i64)
			i64ToSMap[i64] = s
		}
		sort.Sort(sortList)

		for _, i64 := range sortList {
			v := i64ToSMap[i64]
			list = append(list, v)
		}
		appData[smapToJS[i].name] = list
	}

	// --------------- LIST DOWN ID TEXT MAPS ----------------------
	for i := 0; i < len(idTextMapList); i++ {
		list := []IDTextMap{}

		// sort out items of valmap, based on int64 value
		sortList := rlib.Int64Range{}
		i64ToSMap := map[int64]string{}
		for txt, id := range *idTextMapList[i].valmap {
			sortList = append(sortList, id)
			i64ToSMap[id] = txt
		}
		sort.Sort(sortList)

		for _, id := range sortList {
			txt := i64ToSMap[id]
			list = append(list, IDTextMap{ID: id, Text: txt})
		}
		appData[idTextMapList[i].name] = list
	}

	// --------------- LIST DOWN PAYMENT TYPES ----------------------
	var pmtTypes = make(map[string][]pmtMap)
	for bud, bid := range rlib.RRdb.BUDlist {
		bizPmtList := []pmtMap{}
		m, err := rlib.GetPaymentTypesByBusiness(r.Context(), bid) // get the payment types for this business
		if err != nil {
			rlib.Console("Error gettint PaymentTypes: %s\n", err.Error())
			SvcErrorReturn(w, err, funcname)
			return
		}
		for pmt, a := range m {
			bizPmtList = append(bizPmtList, pmtMap{PMTID: pmt, Name: a.Name})
		}
		pmtTypes[bud] = bizPmtList
	}
	appData["pmtTypes"] = pmtTypes

	// --------------- LIST DOWN DEPOSIT METHODS ----------------------
	appData["depmeth"] = GetJSDepositMethods(r.Context())

	// --------------- LIST DOWN SLICES ----------------------
	for i := 0; i < len(ssliceToJS); i++ {
		appData[ssliceToJS[i].name] = ssliceToJS[i].valslice
	}

	// --------------- LIST DOWN LANGUAGE/TEMPLATE STRINGS ----------------------
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	fname := folderPath + "/webclient/html/" + language + "/" + template + "/strings.csv"
	fmt.Printf("fname = %s\n", fname)
	_, err = os.Stat(fname)
	if nil != err {
		e := fmt.Errorf("Unknown language / template :   %s / %s", language, template)
		SvcErrorReturn(w, e, funcname)
		return
	}
	m := rlib.LoadCSV(fname)
	for i := 0; i < len(m); i++ {
		appData[m[i][0]] = m[i][1]
	}

	// --------------- LIST DOWN ACCOUNT STUFF ----------------------
	accountStuffData := getAccountThingJSList()
	accountStuff := make(map[string][]IDTextMap)
	for k, v := range accountStuffData {
		list := []IDTextMap{}
		for i, txt := range v {
			list = append(list, IDTextMap{ID: i, Text: txt})
		}
		accountStuff[k] = list
	}
	appData["account_stuff"] = accountStuff

	// --------------- TASK LIST DEFINITIONS ----------------------
	appData["TaskListDefinitions"] = new([]string)

	// b, err := json.Marshal(appData)
	// if err != nil {
	// 	e := fmt.Errorf("Error marshaling json data: %s", err.Error())
	// 	rlib.Ulog("Error in AJAXLIST marshaling: %s\n", e.Error())
	// } else {
	// 	rlib.Console("AJAXLIST sent to server:  %s\n", string(b))
	// }

	// // send down then json stuff
	// if err := json.NewEncoder(w).Encode(appData); err != nil {
	// 	SvcErrorReturn(w, err, funcname)
	// 	return
	// }
	SvcWriteResponse(-1, appData, w)
}
