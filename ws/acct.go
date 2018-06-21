package ws

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"rentroll/bizlogic"
	"rentroll/importers/core"
	"rentroll/rlib"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ListedAccount is struct to list down individual ledger Account record
type ListedAccount struct {
	LID  int64  `json:"id"`   // Ledger account ID
	Name string `json:"text"` // Ledger account name
}

// AccountListResponse is the response to list down all ledger accounts
type AccountListResponse struct {
	Status  string          `json:"status"`
	Total   int64           `json:"total"`
	Records []ListedAccount `json:"records"`
}

// // w2uiChild struct used to build subgrid
// type w2uiChild struct {
// 	Children []GLAccount `json:"children"`
// }

// GLAccount describes the static (or mostly static) attributes of a Ledger
type GLAccount struct {
	Recid       int               `json:"recid"` // this is for the grid widget
	LID         int64             // unique id for this GLAccount
	PLID        int64             // unique id of Parent, 0 if no parent
	BID         int64             // Business unit associated with this GLAccount
	RAID        int64             // associated rental agreement, this field is only used when Type = 1
	TCID        int64             // associated payor, this field is only used when Type = 1
	GLNumber    string            // acct system name
	Name        string            // descriptive name for the GLAccount
	AcctType    string            // QB Acct Type: Income, Expense, Fixed Asset, Bank, Loan, Credit Card, Equity, Accounts Receivable, Other Current Asset, Other Asset, Accounts Payable, Other Current Liability, Cost of Goods Sold, Other Income, Other Expense
	AllowPost   bool              // 0 = no posting, 1 = posting is allowed
	RARequired  int64             // 0 = during rental period, 1 = valid prior or during, 2 = valid during or after, 3 = valid before, during, and after
	Description string            // description for this account
	FLAGS       uint64            // 1<<0 = inactive flag:  0 = active account, 1 = inactive account
	LastModTime rlib.JSONDateTime // auto updated
	LastModBy   int64             // user making the mod
	// W2UIChild      w2uiChild `json:"w2ui"`
}

// SearchGLAccountsResponse is the response data to a request for GLAccounts
type SearchGLAccountsResponse struct {
	Status  string      `json:"status"`
	Total   int64       `json:"total"`
	Records []GLAccount `json:"records"`
}

// AcctDetailsForm is the response data to request for a GLAccount
type AcctDetailsForm struct {
	LID           int64
	PLID          int64
	BID           int64
	BUD           rlib.XJSONBud
	RAID          int64
	TCID          int64
	GLNumber      string
	Name          string
	AcctType      string
	AllowPost     bool
	Description   string
	FLAGS         uint64 // 1<<0 = inactive flag:  0 = active account, 1 = inactive account
	OffsetAccount int    // 0 = not offset-account, 1 = offset account
	LastModTime   rlib.JSONDateTime
	LastModBy     int64
	CreateTS      rlib.JSONDateTime
	CreateBy      int64
}

// AcctSaveForm used save inputs directly
type AcctSaveForm struct {
	LID           int64
	BID           int64
	RAID          int64
	TCID          int64
	GLNumber      string
	Name          string
	AcctType      string
	Description   string
	BUD           rlib.XJSONBud
	PLID          int64
	AllowPost     bool
	FLAGS         uint64 //
	OffsetAccount int    // the UI value for bit 0 of FLAGS
}

// SaveAcctInput is the input data format for a Save command
type SaveAcctInput struct {
	Cmd      string       `json:"cmd"` // get, save, delete
	Recid    int64        `json:"recid"`
	FormName string       `json:"name"`
	Record   AcctSaveForm `json:"record"`
}

// GetAccountResponse is the response to get details of an account for the requested Account LID
type GetAccountResponse struct {
	Status string          `json:"status"`
	Record AcctDetailsForm `json:"record"`
}

// AcctDeleteForm is struct used to delete Account
type AcctDeleteForm struct {
	LID int64
}

// acctStatus map
// Status cannot have an Unknown state. It's either active or inactive. Default state is Active.
var acctStatus = map[int64]string{
	//0: "Unknown",
	0: "Active",
	1: "Inactive",
}

// account type
var acctType = map[int64]string{
	0: "Normal Account",
	// 1: "balance for this particular RentalAgreement",
	// 2: "balance for this payor",
	// 3:  "Reserved",
	// 4:  "Reserved",
	// 5:  "Reserved",
	// 6:  "Reserved",
	// 7:  "Reserved",
	// 8:  "Reserved",
	// 9:  "Reserved",
	10: "Default Cash",
	11: "Default General Receivables",
	12: "Default Gross Scheduled Rent",
	13: "Default Loss To Lease",
	14: "Default Vacancy",
	16: "Default Security Deposit",
	17: "Default Owner Equity",
}

// // account allow posts
// var acctAllowPosts = map[int64]string{
// 	0: "Summary Account only, do not allow posts to this ledger",
// 	1: "Allow posts",
// }

// getAccountThingJSList sending down list related with accounts info
func getAccountThingJSList() map[string]map[int64]string {
	accountStuff := make(map[string]map[int64]string)
	accountStuff["typeList"] = acctType
	accountStuff["statusList"] = acctStatus
	return accountStuff
}

// SvcAccountsList generates a list of all Accounts with respect of business id specified by d.BID
// wsdoc {
//  @Title Get list of accounts
//  @URL /v1/accountlist/:BUI
//  @Method  GET
//  @Synopsis Get account list
//  @Description Get all General Ledger Account's list for the requested business
//  @Input WebGridSearchRequest
//  @Response AccountListResponse
// wsdoc }
func SvcAccountsList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcAccountsList"
	var (
		g AccountListResponse
	)
	fmt.Printf("Entered %s\n", funcname)

	// get rentable types for a business
	m, err := rlib.GetLedgerList(r.Context(), d.BID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	fmt.Printf("rlib.GetLedgerList returned %d records\n", len(g.Records))

	// append records in ascending order
	var glAccountList []ListedAccount
	for _, acct := range m {
		glAccountList = append(glAccountList,
			ListedAccount{LID: acct.LID, Name: fmt.Sprintf("%s (%s)", acct.GLNumber, acct.Name)},
		)
	}

	// sort based on name, needs version 1.8 later of golang
	sort.Slice(glAccountList, func(i, j int) bool { return glAccountList[i].Name < glAccountList[j].Name })

	g.Records = glAccountList
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// SvcParentAccountsList generates a list of all possible Parent Accounts with respect of business id specified by d.BID
// wsdoc {
//  @Title Get list of parent accounts
//  @URL /v1/parentaccounts/:BUI
//  @Method  GET
//  @Synopsis Get parent account list
//  @Description Get all Parent Account's list for the requested business
//  @Input
//  @Response AccountListResponse
// wsdoc }
func SvcParentAccountsList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcParentAccountsList"
	var (
		err error
		g   AccountListResponse
	)
	fmt.Printf("Entered %s\n", funcname)

	// Need to init some internals for Business
	var xbiz rlib.XBusiness
	err = rlib.InitBizInternals(d.BID, &xbiz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get rentable types for a business
	m := bizlogic.PossibleParentAccounts(d.BID)
	fmt.Printf("bizlogic.PossibleParentAccounts returned %d records\n", len(g.Records))

	// append records in ascending order
	var glAccountList []ListedAccount
	for _, acct := range m {
		glAccountList = append(glAccountList,
			ListedAccount{LID: acct.LID, Name: fmt.Sprintf("%s (%s)", acct.GLNumber, acct.Name)},
		)
	}

	// sort based on name, needs version 1.8 later of golang
	sort.Slice(glAccountList, func(i, j int) bool { return glAccountList[i].Name < glAccountList[j].Name })

	g.Records = glAccountList
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// SvcPostAccountsList generates a list of all Accounts
// that are permissible to use in Assessment/Receipt Rules
// wsdoc {
//  @Title Get list of post accounts
//  @URL /v1/postaccounts/:BUI
//  @Method  GET
//  @Synopsis Get post account list
//  @Description Get all Post Account's list for the requested business
//  @Input
//  @Response AccountListResponse
// wsdoc }
func SvcPostAccountsList(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcPostAccountsList"
	var (
		err error
		g   AccountListResponse
	)
	fmt.Printf("Entered %s\n", funcname)

	// Need to init some internals for Business
	var xbiz rlib.XBusiness
	err = rlib.InitBizInternals(d.BID, &xbiz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get rentable types for a business
	m := bizlogic.PossiblePostAccounts(d.BID)
	fmt.Printf("bizlogic.PossiblePostAccounts returned %d records\n", len(g.Records))

	// append records in ascending order
	var glAccountList []ListedAccount
	for _, acct := range m {
		glAccountList = append(glAccountList,
			ListedAccount{LID: acct.LID, Name: fmt.Sprintf("%s (%s)", acct.GLNumber, acct.Name)},
		)
	}

	// sort based on name, needs version 1.8 later of golang
	sort.Slice(glAccountList, func(i, j int) bool { return glAccountList[i].Name < glAccountList[j].Name })

	g.Records = glAccountList
	g.Total = int64(len(g.Records))
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// SvcSearchHandlerGLAccounts generates a report of all GLAccounts for a the business unit
// called out in d.BID
// wsdoc {
//  @Title  Search General Ledger Accounts
//  @URL /v1/accounts/:BUI
//  @Method POST
//  @Description This service returns a list of General Ledger accounts
//  @Input WebGridSearchRequest
//  @Response SearchGLAccountsResponse
// wsdoc }
func SvcSearchHandlerGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcSearchHandlerGLAccounts"
	var (
		err error
		p   rlib.GLAccount
		g   SearchGLAccountsResponse
	)

	fmt.Printf("Entered %s\n", funcname)

	srch := fmt.Sprintf("BID=%d", d.BID) // default WHERE clause
	order := "GLNumber ASC, Name ASC"    // default ORDER
	q, qw := gridBuildQuery("GLAccount", srch, order, d, &p)

	// set g.Total to the total number of rows of this data...
	g.Total, err = GetRowCount("GLAccount", qw)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	fmt.Printf("db query = %s\n", q)

	rows, err := rlib.RRdb.Dbrr.Query(q)
	defer rows.Close()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	count := 0
	for rows.Next() {
		var p GLAccount
		var q rlib.GLAccount
		err = rlib.ReadGLAccounts(rows, &q)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		rlib.MigrateStructVals(&q, &p)
		p.Recid = count

		g.Records = append(g.Records, p)

		count++ // update the count only after adding the record
		if count >= d.wsSearchReq.Limit {
			break // if we've added the max number requested, then exit
		}
	}
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// ============================================== //
// This following routine is arranging accounts in parent-child
// fashion (according to w2ui subgrid style), as of now just
// disable it, later it will need to be fixed
// ============================================== //

// // SvcSearchHandlerGLAccounts generates a report of all GLAccounts for a the business unit
// // called out in d.BID
// // wsdoc {
// //  @Title  Search General Ledger Accounts
// //	@URL /v1/accounts/:BUI
// //  @Method  GET, POST
// //	@Synopsis Return a list of General Ledger Accounts
// //  @Description This service returns a list of General Ledger accounts
// //	@Input WebGridSearchRequest
// //  @Response SearchGLAccountsResponse
// // wsdoc }
// func SvcSearchHandlerGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {

// 	var (
// 		funcname = "SvcSearchHandlerGLAccounts"
// 		p        rlib.GLAccount
// 		err      error
// 		g        SearchGLAccountsResponse
// 	)

// 	fmt.Printf("Entered %s\n", funcname)

// 	srch := fmt.Sprintf("BID=%d", d.BID)                 // default WHERE clause
// 	order := "PLID ASC, LID ASC, GLNumber ASC, Name ASC" // default ORDER
// 	q, qw := gridBuildQuery("GLAccount", srch, order, d, &p)

// 	// set g.Total to the total number of rows of this data...
// 	g.Total, err = GetRowCount("GLAccount", qw)
// 	if err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}
// 	fmt.Printf("db query = %s\n", q)

// 	rows, err := rlib.RRdb.Dbrr.Query(q)
// 	defer rows.Close()
// 	if err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	// this holds LID keys in ascending order
// 	var sortedLIDKeys rlib.Int64Range

// 	// this map holds values LID -> PLID
// 	acctParentMap := make(map[int64]int64)

// 	// account link: LID -> GLAccount
// 	acctMap := make(map[int64]GLAccount)

// 	count := 0
// 	for rows.Next() {
// 		var p GLAccount
// 		var q rlib.GLAccount
// 		rlib.ReadGLAccounts(rows, &q)
// 		rlib.MigrateStructVals(&q, &p)

// 		// map the account with its LID
// 		acctMap[p.LID] = p
// 		// map account's parent account
// 		acctParentMap[p.LID] = p.PLID
// 		// append LID in sorted slice
// 		sortedLIDKeys = append(sortedLIDKeys, p.LID)

// 		count++ // update the count only after adding the record
// 		if count >= d.wsSearchReq.Limit {
// 			break // if we've added the max number requested, then exit
// 		}
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		SvcErrorReturn(w, err, funcname)
// 		return
// 	}

// 	// this holds the list of deleting account from map, after parent-child relation build-up
// 	deleteAcctKeys := []int64{}

// 	// descending order of LID
// 	sort.Sort(sort.Reverse(sortedLIDKeys))

// 	// find child accounts of parent account, fit it in tree
// 	for _, lid := range sortedLIDKeys {

// 		// get parent LID
// 		plid := acctParentMap[lid]

// 		// if this account is at most parent level then keep continue
// 		if plid == 0 {
// 			continue
// 		}

// 		// get parent account
// 		parentAcct, _ := acctMap[plid]

// 		// get account
// 		childAcct := acctMap[lid]

// 		parentAcct.W2UIChild.Children = append(parentAcct.W2UIChild.Children, childAcct)
// 		acctMap[plid] = parentAcct
// 		deleteAcctKeys = append(deleteAcctKeys, lid)
// 	}

// 	// now delete records which has been put as in child of other account
// 	for _, id := range deleteAcctKeys {
// 		delete(acctMap, id)
// 	}

// 	// this holds PLID keys in ascending order
// 	var sortedPLIDKeys rlib.Int64Range

// 	for plid := range acctMap {
// 		sortedPLIDKeys = append(sortedPLIDKeys, plid)
// 	}

// 	// now sort it in ascending order
// 	sort.Sort(sortedPLIDKeys)

// 	// setRecid is internal function to set Recid used in w2ui grid
// 	setRecid := func(acctMap map[int64]GLAccount) {

// 		// recursive routine
// 		// first declare the function signature, so that we can call it recursively
// 		var childAcctRecid func(acct GLAccount, recid int)

// 		childAcctRecid = func(acct GLAccount, recid int) {
// 			if len(acct.W2UIChild.Children) > 0 {
// 				for id, childAcct := range acct.W2UIChild.Children {
// 					recidx := id
// 					// childID would be parentID + incremental id
// 					childID, _ := strconv.Atoi(strconv.Itoa(acct.Recid) + strconv.Itoa(recidx))
// 					childAcct.Recid = childID
// 					acct.W2UIChild.Children[id] = childAcct
// 					childAcctRecid(childAcct, childID)
// 				}
// 				// TODO: what if someone want to see in ascending order
// 			}
// 		}

// 		mostParentCount := 1
// 		for _, plid := range sortedPLIDKeys {
// 			acct := acctMap[plid]
// 			acct.Recid = mostParentCount
// 			acctMap[plid] = acct
// 			childAcctRecid(acct, mostParentCount)
// 			mostParentCount++
// 		}
// 	}

// 	setRecid(acctMap)

// 	// web response
// 	var records []GLAccount
// 	for _, plid := range sortedPLIDKeys {
// 		acct := acctMap[plid]
// 		records = append(records, acct)
// 	}
// 	g.Records = records
// 	g.Total = int64(len(g.Records))

// 	g.Status = "success"
// 	w.Header().Set("Content-Type", "application/json")
// 	SvcWriteResponse(d.BID,&g, w)
// }

// ======================================================= //

// SvcFormHandlerGLAccounts formats a complete data record for a gl account suitable for use with the w2ui Form
// For this call, we expect the URI to contain the BID and the LID as follows:
//           0    1     2   3
// uri 		/v1/account/BUI/LID
// The server command can be:
//      get
//      save
//      delete
//-----------------------------------------------------------------------------------
func SvcFormHandlerGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcFormHandlerGLAccounts"
	var (
		err error
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("Request: %s:  BID = %d,  LID = %d\n", d.wsSearchReq.Cmd, d.BID, d.ID)

	switch d.wsSearchReq.Cmd {
	case "get":
		if d.ID < 0 {
			err = fmt.Errorf("GLAccount ID is required but was not specified")
			SvcErrorReturn(w, err, funcname)
			return
		}
		getGLAccount(w, r, d)
		break
	case "save":
		saveGLAccount(w, r, d)
		break
	case "delete":
		deleteGLAccount(w, r, d)
		break
	default:
		err = fmt.Errorf("Unhandled command: %s", d.wsSearchReq.Cmd)
		SvcErrorReturn(w, err, funcname)
		return
	}
}

// saveGLAccount returns the requested receipt
// wsdoc {
//  @Title Save GLAccount
//  @URL /v1/account/:BUI/:LID
//  @Method POST
//  @Synopsis Saves a GLAccount details
//  @Description This service saves a GLAccount.  If :LID exists, it will
//  @Description be updated with the information supplied. All fields must
//  @Description be supplied. If LID is 0, then a new GLAccount is created.
//  @Input SaveAcctInput
//  @Response SvcStatusResponse
// wsdoc }
func saveGLAccount(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "saveGLAccount"
	var (
		foo SaveAcctInput
		a   rlib.GLAccount
		err error
	)

	rlib.Console("Entered %s\n", funcname)
	rlib.Console("record data = %s\n", d.data)

	// get data
	data := []byte(d.data)

	if err = json.Unmarshal(data, &foo); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//--------------------------------------------------------------------------
	// Make no assumptions about what the caller has done to the FLAGS
	// Only accept those bits that a caller is authorized to edit. The bizlogic
	// check takes care of not overwriting the existing FLAG values if this
	// is an update operation. Vut we're going to limit the FLAGS to only what
	// clients are authorized to change...
	//--------------------------------------------------------------------------
	rlib.Console("foo.Record.OffsetAccount = %d\n", foo.Record.OffsetAccount)

	// migrate foo.Record data to a struct's fields
	rlib.MigrateStructVals(&foo.Record, &a) // the variables that don't need special handling
	fmt.Printf("saveAcct - first migrate: a = %#v\n", a)
	a.FLAGS = 0 // reset anything that the UI sent

	// data validation
	if a.Name == "" {
		err := fmt.Errorf("Provide account name")
		SvcErrorReturn(w, err, funcname)
		return
	}
	if a.GLNumber == "" {
		err := fmt.Errorf("Provide value of GLNumber")
		SvcErrorReturn(w, err, funcname)
		return
	}

	// save or update
	if a.LID == 0 && d.ID == 0 {
		//-------------------------------------------------------------------
		// check that given name is already exists for business, or GLNumber
		// both name and GLNumber should be unique
		// VALIDATION 2
		//-------------------------------------------------------------------
		existQuery := `SELECT LID FROM GLAccount WHERE {{.WhereClause}};`
		qc := rlib.QueryClause{
			"WhereClause": fmt.Sprintf("BID=%d AND (Name=\"%s\" OR GLNumber=\"%s\")",
				d.BID, strings.ToLower(a.Name), strings.ToLower(a.GLNumber)),
		}
		q := rlib.RenderSQLQuery(existQuery, qc)
		fmt.Printf("db query = %s\n", q)
		rows, err := rlib.RRdb.Dbrr.Query(q)
		defer rows.Close()
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		for rows.Next() {
			err := fmt.Errorf("GLAccount is already exists for given name or GLNumber")
			SvcErrorReturn(w, err, funcname)
			return
		}

		//-------------------------------------------------------------------
		// OK, it's a new account.  Do the bizlogic checks and save...
		//-------------------------------------------------------------------
		errlist := bizlogic.SaveGLAccount(r.Context(), &a)
		if len(errlist) > 0 {
			SvcErrListReturn(w, errlist, funcname)
			return
		}

		//-------------------------------------------------------------------
		// Since it is a new Account, we need a LedgerMarker for it...
		//-------------------------------------------------------------------
		var lm = rlib.LedgerMarker{
			BID:   a.BID,
			LID:   a.LID,
			Dt:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			State: rlib.LMINITIAL,
		}
		_, err = rlib.InsertLedgerMarker(r.Context(), &lm)
		if err != nil {
			e := fmt.Errorf("Error saving Account %s LedgerMarker, Error:= %s", a.Name, err.Error())
			SvcErrorReturn(w, e, funcname)
			return
		}
	} else {
		// update existing record
		errlist := bizlogic.SaveGLAccount(r.Context(), &a)
		if len(errlist) > 0 {
			SvcErrListReturn(w, errlist, funcname)
			return
		}
	}

	SvcWriteSuccessResponseWithID(d.BID, w, a.LID)
}

// which fields needs to be fetched for SQL query for receipts grid
var getAcctQuerySelectFields = rlib.SelectQueryFields{
	"GLAccount.LID",
	"GLAccount.PLID",
	"GLAccount.RAID",
	"GLAccount.TCID",
	"GLAccount.GLNumber",
	"GLAccount.Name",
	"GLAccount.AcctType",
	"GLAccount.AllowPost",
	"GLAccount.Description",
	"GLAccount.FLAGS",
	"GLAccount.LastModTime",
	"GLAccount.LastModBy",
	"GLAccount.CreateTS",
	"GLAccount.CreateBy",
}

// getGLAccount returns the requested glaccount
// wsdoc {
//  @Title  Get account details
//  @URL /v1/account/:BUI/:LID
//  @Method POST
//  @Synopsis Get details about an account
//  @Description  Return all fields for account :LID
//  @Input WebGridSearchRequest
//  @Response GetAccountResponse
// wsdoc }
func getGLAccount(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "getGLAccount"
	var (
		g     GetAccountResponse
		err   error
		order = `GLAccount.LID ASC`
		whr   = fmt.Sprintf(`GLAccount.BID=%d AND GLAccount.LID=%d`, d.BID, d.ID)
	)

	fmt.Printf("entered %s\n", funcname)

	glQuery := `
	SELECT
		{{.SelectClause}}
	FROM GLAccount
	WHERE {{.WhereClause}}
	ORDER BY {{.OrderClause}};`

	qc := rlib.QueryClause{
		"SelectClause": strings.Join(getAcctQuerySelectFields, ","),
		"WhereClause":  whr,
		"OrderClause":  order,
	}

	// get formatted query with substitution of select, where, order clause
	q := rlib.RenderSQLQuery(glQuery, qc)
	fmt.Printf("db query = %s\n", q)

	// execute the query
	rows, err := rlib.RRdb.Dbrr.Query(q)
	defer rows.Close()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	for rows.Next() {
		var gg AcctDetailsForm
		gg.BID = d.BID
		gg.BUD = rlib.GetBUDFromBIDList(d.BID)

		err = rows.Scan(&gg.LID, &gg.PLID, &gg.RAID, &gg.TCID, &gg.GLNumber, &gg.Name, &gg.AcctType, &gg.AllowPost, &gg.Description, &gg.FLAGS, &gg.LastModTime, &gg.LastModBy, &gg.CreateTS, &gg.CreateBy)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		gg.OffsetAccount = int(gg.FLAGS & 0x1)

		g.Record = gg
	}

	// error check
	err = rows.Err()
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}

// deleteGLAccount request delete GLAccount from database
// wsdoc {
//  @Title  Delete GLAccount
//  @URL /v1/account/:BUI/:LID
//  @Method  DELETE
//  @Synopsis Delete record for a GL Account
//  @Description  Delete the GL Account for a database and delete its
//  @Description  associated LedgerMarkers.  Use with caution. Only use
//  @Description  this command if you really understand what you're doing.
//  @Input AcctDeleteForm
//  @Response SvcStatusResponse
// wsdoc }
func deleteGLAccount(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "deleteGLAccount"
	var (
		del AcctDeleteForm
	)
	fmt.Printf("Entered %s\n", funcname)
	fmt.Printf("record data = %s\n", d.data)

	if err := json.Unmarshal([]byte(d.data), &del); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	//----------------------------------------
	// First check, account exists or not
	//----------------------------------------
	gl, err := rlib.GetLedger(r.Context(), del.LID)
	if err != nil || gl.LID == 0 {
		// if you want to log error then separate this if clause condition
		err := fmt.Errorf("No such account exists with ID: %d", del.LID)
		SvcErrorReturn(w, err, funcname)
		return
	}

	//----------------------------------------
	// do biz logic checks...
	//----------------------------------------
	l, err := rlib.GetLedger(r.Context(), del.LID)
	if err != nil {
		err := fmt.Errorf("No such account exists with ID: %d", del.LID)
		SvcErrorReturn(w, err, funcname)
		return
	}

	ok, errlist := bizlogic.OKToDelete(r.Context(), &l)
	if !ok {
		SvcErrListReturn(w, errlist, funcname)
		return
	}

	//-----------------------------------------------
	// Passed all the checks... OK to remove it.
	// Remove LedgerMarkers for this LID
	//-----------------------------------------------
	// ODO(Steve): ignore error?
	lm, _ := rlib.GetLatestLedgerMarkerByLID(r.Context(), d.BID, del.LID)
	if lm.State != rlib.LMINITIAL {
		e := fmt.Errorf("This account (LID = %d) cannot be deleted because Ledger Markers exist beyond the origin", del.LID)
		SvcErrorReturn(w, e, funcname)
		return
	}
	if err := rlib.DeleteLedgerMarker(r.Context(), lm.LMID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	if err := rlib.DeleteLedger(r.Context(), del.LID); err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	SvcWriteSuccessResponse(d.BID, w)
}

// SvcExportGLAccounts used to export glaccounts for a business in csv format
func SvcExportGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcExportGLAccounts"
	var (
		err error
		buf = bytes.Buffer{}
		wr  = csv.NewWriter(&buf)
	)
	fmt.Printf("Entered %s", funcname)

	// Need to init some internals for Business
	var xbiz rlib.XBusiness
	err = rlib.InitBizInternals(d.BID, &xbiz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get list of all accounts
	accts, err := rlib.GetLedgerList(r.Context(), d.BID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// write csv file headers
	wr.Write([]string{"BUD", "Name", "GLNumber", "Parent GLNumber", "Account Type",
		"Balance", "Account Status", "Date", "Description"})

	for _, a := range accts {
		bud := rlib.GetBUDFromBIDList(a.BID)
		rec := []string{string(bud), a.Name, a.GLNumber}

		// get parent account GLNumber
		var paGLNumber string
		if a.PLID > 0 {
			pa, err := rlib.GetLedger(r.Context(), a.PLID)
			if err != nil {
				SvcErrorReturn(w, err, funcname)
				return
			}
			paGLNumber = pa.GLNumber
		}
		rec = append(rec, paGLNumber)

		// append account type
		rec = append(rec, a.AcctType)

		// append balance
		now := time.Now()
		bal, err := rlib.GetAccountBalance(r.Context(), d.BID, a.LID, &now)
		if err != nil {
			SvcErrorReturn(w, err, funcname)
			return
		}
		s64Bal := strconv.FormatFloat(bal, 'f', 2, 64)
		rec = append(rec, s64Bal)

		// append Status, CreateDate, Description
		rec = append(rec, acctStatus[int64(a.FLAGS&1)])
		rec = append(rec, a.CreateTS.Format(rlib.RRDATEFMT3))
		rec = append(rec, a.Description)

		// write to buffer
		wr.Write(rec)
	}
	wr.Flush()

	expFileName := fmt.Sprintf("%s_%s.csv", rlib.GetBUDFromBIDList(d.BID), "GLAccounts")
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", expFileName))
	w.Write(buf.Bytes())
}

// ImportGLAccountRow struct used to load data from imported csv file
type ImportGLAccountRow struct {
	BUD            string
	Name           string
	GLNumber       string
	ParentGLNumber string
	AccountType    string
	Balance        float64
	AccountStatus  string
	Date           time.Time
	Description    string
}

// SvcImportGLAccounts used to import glaccounts for a business from csv format
func SvcImportGLAccounts(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcImportGLAccounts"
	var (
		err           error
		inf           multipart.File
		recs          = [][]string{}
		skipRowsCount int
		// it tells which headers holds which column index
		acctCSVIndexMap = map[string]int{
			"bud":            -1,
			"name":           -1,
			"glnumber":       -1,
			"parentglnumber": -1,
			"accounttype":    -1,
			"balance":        -1,
			"accountstatus":  -1,
			"date":           -1,
			"description":    -1,
		}
	)
	fmt.Printf("Entered %s\n", funcname)

	// Need to init some internals for Business
	var xbiz rlib.XBusiness
	err = rlib.InitBizInternals(d.BID, &xbiz)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get BUD from formData value
	var bud = d.MFValues["BUD"][0] //get first value from slice
	bid, ok := rlib.RRdb.BUDlist[bud]
	if !ok {
		err = fmt.Errorf("Supplied Business (%s) not found", bud)
		SvcErrorReturn(w, err, funcname)
		return
	}

	fheaders, ok := d.Files["GLAccountFile"]
	if !ok { // if not found file then just return
		err = fmt.Errorf("file is missing")
		SvcErrorReturn(w, err, funcname)
		return
	}

	fh := fheaders[0]    // get one file
	inf, err = fh.Open() // get File (multipart.File)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}
	// check extension/content-type
	if fh.Header["Content-Type"][0] != "text/csv" {
		err = fmt.Errorf("Provided file is not type of csv")
		SvcErrorReturn(w, err, funcname)
		return
	}

	cr := csv.NewReader(inf) // csv NewReader (since, inf composed io.Reader)
	recs, err = cr.ReadAll()
	if err != nil {
		err = fmt.Errorf("Unable to read the file")
		SvcErrorReturn(w, err, funcname)
		return
	}

	// get headers index
	for ri := 0; ri < len(recs); ri++ {
		for ci := 0; ci < len(recs[ri]); ci++ {
			cell := strings.ToLower(core.SpecialCharsReplacer.Replace(recs[ri][ci]))
			// if header is exist in map then overwrite it position
			if _, ok := acctCSVIndexMap[cell]; ok {
				acctCSVIndexMap[cell] = ci
			}
		}
		// check after row columns parsing that headers are found or not
		headersFound := true
		for _, v := range acctCSVIndexMap {
			if v == -1 {
				headersFound = false
				break
			}
		}

		if headersFound {
			ri++
			skipRowsCount = ri
			break
		} else {
			missingFields := []string{}
			for h, i := range acctCSVIndexMap {
				if i == -1 {
					missingFields = append(missingFields, h)
				}
			}
			err = fmt.Errorf("Required fields(%v) are not present in file", missingFields)
			SvcErrorReturn(w, err, funcname)
			return
		}
	}

	// iterate over csv rows
	for ri := skipRowsCount; ri < len(recs); ri++ {
		// new gl account
		var ngl rlib.GLAccount

		// first check if account is already exists, if yes then continue to next
		// 1. By GLNumber,
		glNo := recs[ri][acctCSVIndexMap["glnumber"]]
		ngl, _ = rlib.GetLedgerByGLNo(r.Context(), bid, glNo)
		if ngl.LID > 0 {
			continue
		}
		// 2. By Name
		name := recs[ri][acctCSVIndexMap["name"]]
		ngl, _ = rlib.GetLedgerByName(r.Context(), bid, name)
		if ngl.LID > 0 {
			continue
		}

		// NOTE: ignore bud, date
		ngl.BID = bid
		ngl.Name = name
		ngl.GLNumber = glNo
		ngl.AcctType = recs[ri][acctCSVIndexMap["accounttype"]]
		ngl.Description = recs[ri][acctCSVIndexMap["description"]]

		// set status
		strStatus := recs[ri][acctCSVIndexMap["accountstatus"]]
		for n, s := range acctStatus {
			if strings.ToLower(s) == strings.ToLower(strStatus) { // if both match as case-insensitive the mark
				ngl.FLAGS &= uint64(n)
			}
		}

		// if ngl.Status == 0 { // set default as active
		// 	ngl.Status = 2 // 0=unknown, 1=inactive, 2=active
		// }

		// set PLID from parent glnumber if available
		pglNo := recs[ri][acctCSVIndexMap["parentglnumber"]]
		if pglNo != "" { // set parent account LID -> PLID
			pgl, _ := rlib.GetLedgerByGLNo(r.Context(), bid, pglNo)
			if pgl.LID > 0 {
				ngl.PLID = pgl.LID // set parent account lid in PLID field
			}
		}

		// now hit this in database to insert new record
		ngl.LID, err = rlib.InsertLedger(r.Context(), &ngl)
		if err != nil {
			// continue to next one, if current one fails
			// process as much as possible, we can mark failure one in future
			continue
		}

		// if succeed then, process for balance
		lm := rlib.LedgerMarker{
			BID:     bid,
			State:   3,
			Dt:      time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			LID:     ngl.LID,
			Balance: float64(0),
		}
		// now if balance provided, then parse it to float64
		balStr := recs[ri][acctCSVIndexMap["balance"]]
		bal, _ := strconv.ParseFloat(balStr, 64)
		lm.Balance = bal                                   // don't worry about balance if can't parsed from string, default will be 0
		_, err = rlib.InsertLedgerMarker(r.Context(), &lm) // insert ledger marker
		if err != nil {
			rlib.Ulog("%s: Error while inserting ledger marker: %s\n", funcname, err.Error())
			// continue to next one, if current one fails
			// process as much as possible, we can mark failure one in future
			continue

		}
	}

	// if all passed then return success response
	SvcWriteSuccessResponse(d.BID, w)
}
